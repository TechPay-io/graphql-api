/*
Package rpc implements bridge to Sirius full node API interface.

We recommend using local IPC for fast and the most efficient inter-process communication between the API server
and an Photon/Sirius node. Any remote RPC connection will work, but the performance may be significantly degraded
by extra networking overhead of remote RPC calls.

You should also consider security implications of opening Sirius RPC interface for remote access.
If you considering it as your deployment strategy, you should establish encrypted channel between the API server
and Sirius RPC interface with connection limited to specified endpoints.

We strongly discourage opening Sirius RPC interface for unrestricted Internet access.
*/
package rpc

import (
	"context"
	"strings"
	"sync"
	"techpay-api-graphql/internal/config"
	"techpay-api-graphql/internal/logger"
	"techpay-api-graphql/internal/repository/rpc/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	etc "github.com/ethereum/go-ethereum/core/types"
	eth "github.com/ethereum/go-ethereum/ethclient"
	tpc "github.com/ethereum/go-ethereum/rpc"
	"golang.org/x/sync/singleflight"
)

// rpcHeadProxyChannelCapacity represents the capacity of the new received blocks proxy channel.
const rpcHeadProxyChannelCapacity = 10000

// TpcBridge represents Sirius RPC abstraction layer.
type TpcBridge struct {
	rpc *tpc.Client
	eth *eth.Client
	log logger.Logger
	cg  *singleflight.Group

	// fMintCfg represents the configuration of the fMint protocol
	sigConfig     *config.ServerSignature
	sfcConfig     *config.Staking
	uniswapConfig *config.DeFiUniswap

	// extended minter config
	fMintCfg fMintConfig
	fLendCfg fLendConfig

	// common contracts
	sfcAbi      *abi.ABI
	sfcContract *contracts.SfcContract

	// received blocks proxy
	wg       *sync.WaitGroup
	sigClose chan bool
	headers  chan *etc.Header
}

// New creates new Sirius RPC connection bridge.
func New(cfg *config.Config, log logger.Logger) (*TpcBridge, error) {
	cli, con, err := connect(cfg, log)
	if err != nil {
		log.Criticalf("can not open connection; %s", err.Error())
		return nil, err
	}

	// build the bridge structure using the con we have
	br := &TpcBridge{
		rpc: cli,
		eth: con,
		log: log,
		cg:  new(singleflight.Group),

		// special configuration options below this line
		sigConfig:     &cfg.MySignature,
		sfcConfig:     &cfg.Staking,
		uniswapConfig: &cfg.DeFi.Uniswap,
		fMintCfg: fMintConfig{
			addressProvider: cfg.DeFi.FMint.AddressProvider,
		},
		fLendCfg: fLendConfig{lendigPoolAddress: cfg.DeFi.FLend.LendingPool},

		// configure block observation loop
		wg:       new(sync.WaitGroup),
		sigClose: make(chan bool, 1),
		headers:  make(chan *etc.Header, rpcHeadProxyChannelCapacity),
	}

	// inform about the local address of the API node
	log.Noticef("using signature address %s", br.sigConfig.Address.String())

	// add the bridge ref to the fMintCfg and return the instance
	br.fMintCfg.bridge = br
	br.run()
	return br, nil
}

// connect opens connections we need to communicate with the blockchain node.
func connect(cfg *config.Config, log logger.Logger) (*tpc.Client, *eth.Client, error) {
	// log what we do
	log.Debugf("connecting blockchain node at %s", cfg.Sirius.Url)

	// try to establish a connection
	client, err := tpc.Dial(cfg.Sirius.Url)
	if err != nil {
		log.Critical(err)
		return nil, nil, err
	}

	// try to establish a for smart contract interaction
	con, err := eth.Dial(cfg.Sirius.Url)
	if err != nil {
		log.Critical(err)
		return nil, nil, err
	}

	// log
	log.Notice("node connection open")
	return client, con, nil
}

// run starts the bridge threads required to collect blockchain data.
func (tpc *TpcBridge) run() {
	tpc.wg.Add(1)
	go tpc.observeBlocks()
}

// terminate kills the bridge threads to end the bridge gracefully.
func (tpc *TpcBridge) terminate() {
	tpc.sigClose <- true
	tpc.wg.Wait()
	tpc.log.Noticef("rpc threads terminated")
}

// Close will finish all pending operations and terminate the Sirius RPC connection
func (tpc *TpcBridge) Close() {
	// terminate threads before we close connections
	tpc.terminate()

	// do we have a connection?
	if tpc.rpc != nil {
		tpc.rpc.Close()
		tpc.eth.Close()
		tpc.log.Info("blockchain connections are closed")
	}
}

// Connection returns open Photon/Sirius connection.
func (tpc *TpcBridge) Connection() *tpc.Client {
	return tpc.rpc
}

// DefaultCallOpts creates a default record for call options.
func (tpc *TpcBridge) DefaultCallOpts() *bind.CallOpts {
	// get the default call opts only once if called in parallel
	co, _, _ := tpc.cg.Do("default-call-opts", func() (interface{}, error) {
		return &bind.CallOpts{
			Pending:     false,
			From:        tpc.sigConfig.Address,
			BlockNumber: nil,
			Context:     context.Background(),
		}, nil
	})
	return co.(*bind.CallOpts)
}

// SfcContract returns instance of SFC contract for interaction.
func (tpc *TpcBridge) SfcContract() *contracts.SfcContract {
	// lazy create SFC contract instance
	if nil == tpc.sfcContract {
		// instantiate the contract and display its name
		var err error
		tpc.sfcContract, err = contracts.NewSfcContract(tpc.sfcConfig.SFCContract, tpc.eth)
		if err != nil {
			tpc.log.Criticalf("failed to instantiate SFC contract; %s", err.Error())
			panic(err)
		}
	}
	return tpc.sfcContract
}

// SfcAbi returns a parse ABI of the AFC contract.
func (tpc *TpcBridge) SfcAbi() *abi.ABI {
	if nil == tpc.sfcAbi {
		ab, err := abi.JSON(strings.NewReader(contracts.SfcContractABI))
		if err != nil {
			tpc.log.Criticalf("failed to parse SFC contract ABI; %s", err.Error())
			panic(err)
		}
		tpc.sfcAbi = &ab
	}
	return tpc.sfcAbi
}

// ObservedBlockProxy provides a channel fed with new headers observed
// by the connected blockchain node.
func (tpc *TpcBridge) ObservedBlockProxy() chan *etc.Header {
	return tpc.headers
}
