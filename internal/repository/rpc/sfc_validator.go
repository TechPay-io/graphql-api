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
	"fmt"
	"math/big"
	"techpay-api-graphql/internal/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// ValidatorDowntime pulls information about validator downtime from the RPC interface.
func (tpc *TpcBridge) ValidatorDowntime(valID *hexutil.Big) (uint64, uint64, error) {
	// use rather the public API, it should be faster since it does not involve contract call
	var dt struct {
		Blocks hexutil.Uint64 `json:"offlineBlocks"`
		Time   hexutil.Uint64 `json:"offlineTime"`
	}
	if err := tpc.rpc.Call(&dt, "abft_getDowntime", valID); err != nil {
		tpc.log.Errorf("failed to get downtime of validator #%d; %s", valID.ToInt().Uint64(), err.Error())
		return 0, 0, err
	}
	// get the values
	return uint64(dt.Time), uint64(dt.Blocks), nil
}

// ValidatorEpochUptime pulls information about validator uptime on the given epoch.
func (tpc *TpcBridge) ValidatorEpochUptime(valID *hexutil.Big) (uint64, error) {
	// use rather the public API, it should be faster since it does not involve contract call
	var ut hexutil.Uint64
	if err := tpc.rpc.Call(&ut, "abft_getEpochUptime", valID); err != nil {
		tpc.log.Errorf("failed to get epoch uptime of validator #%d; %s", valID.ToInt().Uint64(), err.Error())
		return 0, err
	}
	// get the values
	return uint64(ut), nil
}

// LastValidatorId returns the last staker id in Photon blockchain.
func (tpc *TpcBridge) LastValidatorId() (uint64, error) {
	// get the value from the contract
	sl, err := tpc.SfcContract().LastValidatorID(nil)
	if err != nil {
		tpc.log.Errorf("failed to get the last staker ID: %s", err.Error())
		return 0, err
	}
	return sl.Uint64(), nil
}

// ValidatorsCount returns the number of validators in Photon blockchain.
func (tpc *TpcBridge) ValidatorsCount() (uint64, error) {
	// get the value from the contract
	epoch, err := tpc.SfcContract().CurrentEpoch(tpc.DefaultCallOpts())
	if err != nil {
		tpc.log.Errorf("failed to get the current sealed epoch: %s", err.Error())
		return 0, err
	}

	// get the value from the contract
	val, err := tpc.SfcContract().GetEpochValidatorIDs(nil, epoch)
	if err != nil {
		tpc.log.Errorf("failed to get the list of validators; %s", err.Error())
		return 0, err
	}

	// get the value
	return uint64(len(val)), nil
}

// Validator extract a staker information by numeric id.
func (tpc *TpcBridge) Validator(valID *big.Int) (*types.Validator, error) {
	// no validator id?
	if valID == nil {
		return nil, fmt.Errorf("validator ID not provided")
	}

	// keep track of the operation
	tpc.log.Debugf("loading validator #%d", valID.Uint64())
	return tpc.validatorById(valID)
}

// validatorById loads details of a validator with the specified ID.
func (tpc *TpcBridge) validatorById(valID *big.Int) (*types.Validator, error) {
	// call for data
	val, err := tpc.SfcContract().GetValidator(nil, valID)
	if err != nil {
		tpc.log.Criticalf("failed to load validator #%d from SFC; %s", valID.Uint64(), err.Error())
		return nil, err
	}

	// any creation record?
	if 0 == val.CreatedTime.Uint64() {
		tpc.log.Errorf("validator #%d has zero created time, assuming empty record", valID.Uint64())
		return nil, fmt.Errorf("validator #%d not found", valID.Uint64())
	}

	// any deactivation epoch?
	var deaEpoch hexutil.Uint64
	if nil != val.DeactivatedEpoch {
		deaEpoch = hexutil.Uint64(val.DeactivatedEpoch.Uint64())
	}

	// any deactivation time?
	var deaTime hexutil.Uint64
	if nil != val.DeactivatedTime {
		deaTime = hexutil.Uint64(val.DeactivatedTime.Uint64())
	}

	// keep track of the operation
	tpc.log.Debugf("validator #%d is %s", valID.Uint64(), val.Auth.String())
	return &types.Validator{
		Id:               (hexutil.Big)(*valID),
		StakerAddress:    val.Auth,
		TotalStake:       (*hexutil.Big)(val.ReceivedStake),
		Status:           hexutil.Uint64(val.Status.Uint64()),
		CreatedEpoch:     hexutil.Uint64(val.CreatedEpoch.Uint64()),
		CreatedTime:      hexutil.Uint64(val.CreatedTime.Uint64()),
		DeactivatedEpoch: deaEpoch,
		DeactivatedTime:  deaTime,
	}, nil
}

// ValidatorAddress extract a staker address for the given staker ID.
func (tpc *TpcBridge) ValidatorAddress(valID *big.Int) (*common.Address, error) {
	// do we have an address call?
	val, err := tpc.SfcContract().GetValidator(nil, valID)
	if err != nil {
		tpc.log.Error("validator information could not be extracted")
		return nil, err
	}

	// any creation record?
	if 0 == val.CreatedTime.Uint64() {
		tpc.log.Errorf("validator #%d has zero created time, assuming empty record", valID.Uint64())
		return nil, fmt.Errorf("validator #%d not found", valID.Uint64())
	}

	tpc.log.Debugf("validator #%d is %s", valID.Uint64(), val.Auth.String())
	return &val.Auth, nil
}

// IsValidator returns if the given address is an SFC validator.
func (tpc *TpcBridge) IsValidator(addr *common.Address) (bool, error) {
	// keep track of the operation
	tpc.log.Debugf("verifying validator address %s", addr.String())

	// try to get the id
	id, err := tpc.SfcContract().GetValidatorID(nil, *addr)
	if err != nil {
		tpc.log.Criticalf("can not check validator at %s; %s", addr.String(), err.Error())
		return false, err
	}

	return 0 < id.Uint64(), nil
}

// ValidatorByAddress extracts a validator information by address.
func (tpc *TpcBridge) ValidatorByAddress(addr *common.Address) (*types.Validator, error) {
	// no validator id?
	if addr == nil {
		return nil, fmt.Errorf("validator address not provided")
	}

	// keep track of the operation
	tpc.log.Debugf("loading validator with address %s", addr.String())

	// try to get the staker id
	id, err := tpc.SfcContract().GetValidatorID(tpc.DefaultCallOpts(), *addr)
	if err != nil {
		tpc.log.Criticalf("can not check validator at %s; %s", addr.String(), err.Error())
		return nil, err
	}

	// do we have the ID?
	if 0 == id.Uint64() {
		tpc.log.Debugf("validator not found for address %s", addr.String())
		return nil, nil
	}
	return tpc.validatorById(id)
}
