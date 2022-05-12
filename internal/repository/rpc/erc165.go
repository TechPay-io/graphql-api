/*
Package rpc implements bridge to Sirius full node API interface.

We recommend using local IPC for fast and the most efficient inter-process communication between the API server
and an Photon/Sirius node. Any remote RPC connection will work, but the performance may be significantly degraded
by extra networking overhead of remote RPC calls.

You should also consider security implications of opening Sirius RPC interface for a remote access.
If you considering it as your deployment strategy, you should establish encrypted channel between the API server
and Sirius RPC interface with connection limited to specified endpoints.

We strongly discourage opening Sirius RPC interface for unrestricted Internet access.
*/
package rpc

import (
	"techpay-api-graphql/internal/repository/rpc/contracts"

	"github.com/ethereum/go-ethereum/common"
)

//go:generate tools/abigen.sh --abi ./contracts/abi/erc165.abi --pkg contracts --type ERC165 --out ./contracts/erc165.go

func (tpc *TpcBridge) Erc165SupportsInterface(address *common.Address, interfaceID [4]byte) (bool, error) {
	// connect the contract
	contract, err := contracts.NewERC165(*address, tpc.eth)
	if err != nil {
		tpc.log.Errorf("can not contact ERC165 contract; %s", err.Error())
		return false, err
	}

	supports, err := contract.SupportsInterface(nil, interfaceID)
	if err != nil {
		tpc.log.Noticef("interface support by ERC165 for contract %s cannot be detected; %s", address.String(), err.Error())
		return false, err
	}

	return supports, nil
}
