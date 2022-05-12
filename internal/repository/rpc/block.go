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
	"fmt"
	"math/big"
	"techpay-api-graphql/internal/types"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

// BlockTypeLatest represents the latest available block in blockchain.
const (
	BlockTypeLatest   = "latest"
	BlockTypeEarliest = "earliest"
)

// MustBlockHeight returns the current block height
// of the blockchain. It returns nil if the block height can not be pulled.
func (tpc *TpcBridge) MustBlockHeight() *big.Int {
	var val hexutil.Big
	if err := tpc.rpc.Call(&val, "tpc_blockNumber"); err != nil {
		tpc.log.Errorf("failed block height check; %s", err.Error())
		return nil
	}
	return val.ToInt()
}

// BlockHeight returns the current block height of the Photon blockchain.
func (tpc *TpcBridge) BlockHeight() (*hexutil.Big, error) {
	// keep track of the operation
	tpc.log.Debugf("checking current block height")

	// call for data
	var height hexutil.Big
	err := tpc.rpc.Call(&height, "tpc_blockNumber")
	if err != nil {
		tpc.log.Error("block height could not be obtained")
		return nil, err
	}

	// inform and return
	tpc.log.Debugf("current block height is %s", height.String())
	return &height, nil
}

// Block returns information about a blockchain block by encoded hex number, or by a type tag.
// For tag based loading use predefined BlockType contacts.
func (tpc *TpcBridge) Block(numTag *string) (*types.Block, error) {
	// keep track of the operation
	tpc.log.Debugf("loading details of block num/tag %s", *numTag)

	// call for data
	var block types.Block
	err := tpc.rpc.Call(&block, "tpc_getBlockByNumber", numTag, false)
	if err != nil {
		tpc.log.Error("block could not be extracted")
		return nil, err
	}

	// detect block not found situation; block number is zero and the hash is also zero
	if uint64(block.Number) == 0 && block.Hash.Big().Cmp(big.NewInt(0)) == 0 {
		tpc.log.Debugf("block [%s] not found", *numTag)
		return nil, fmt.Errorf("block not found")
	}

	// keep track of the operation
	tpc.log.Debugf("block #%d found at mark %s",
		uint64(block.Number), time.Unix(int64(block.TimeStamp), 0).String())
	return &block, nil
}

// BlockByHash returns information about a blockchain block by hash.
func (tpc *TpcBridge) BlockByHash(hash *string) (*types.Block, error) {
	// keep track of the operation
	tpc.log.Debugf("loading details of block %s", *hash)

	// call for data
	var block types.Block
	err := tpc.rpc.Call(&block, "tpc_getBlockByHash", hash, false)
	if err != nil {
		tpc.log.Error("block could not be extracted")
		return nil, err
	}

	// detect block not found situation
	if uint64(block.Number) == 0 {
		tpc.log.Debugf("block [%s] not found", *hash)
		return nil, fmt.Errorf("block not found")
	}

	// inform and return
	tpc.log.Debugf("block #%d found at mark %s by hash %s",
		uint64(block.Number), time.Unix(int64(block.TimeStamp), 0).String(), *hash)
	return &block, nil
}
