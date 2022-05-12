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

//go:generate tools/abigen.sh --abi ./contracts/abi/sfc-1.1.abi --pkg contracts --type SfcV1Contract --out ./contracts/sfc-v1.go
//go:generate tools/abigen.sh --abi ./contracts/abi/sfc-2.0.4-rc.2.abi --pkg contracts --type SfcV2Contract --out ./contracts/sfc-v2.go
//go:generate tools/abigen.sh --abi ./contracts/abi/sfc-3.0-rc.1.abi --pkg contracts --type SfcContract --out ./contracts/sfc-v3.go
//go:generate tools/abigen.sh --abi ./contracts/abi/sfc-tokenizer.abi --pkg contracts --type SfcTokenizer --out ./contracts/sfc_tokenizer.go

import (
	"math/big"
	"techpay-api-graphql/internal/types"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

// sfcFirstLockEpoch represents the first epoch with stake locking available.
const sfcFirstLockEpoch uint64 = 1600

// SfcVersion returns current version of the SFC contract as a single number.
func (tpc *TpcBridge) SfcVersion() (hexutil.Uint64, error) {
	// get the version information from the contract
	var ver [3]byte
	var err error
	ver, err = tpc.SfcContract().Version(nil)
	if err != nil {
		tpc.log.Criticalf("failed to get the SFC version; %s", err.Error())
		return 0, err
	}
	return hexutil.Uint64((uint64(ver[0]) << 16) | (uint64(ver[1]) << 8) | uint64(ver[2])), nil
}

// CurrentEpoch extract the current epoch id from SFC smart contract.
func (tpc *TpcBridge) CurrentEpoch() (hexutil.Uint64, error) {
	// get the value from the contract
	epoch, err := tpc.SfcContract().CurrentEpoch(tpc.DefaultCallOpts())
	if err != nil {
		tpc.log.Errorf("failed to get the current sealed epoch: %s", err.Error())
		return 0, err
	}
	return hexutil.Uint64(epoch.Uint64()), nil
}

// CurrentSealedEpoch extract the current sealed epoch id from SFC smart contract.
func (tpc *TpcBridge) CurrentSealedEpoch() (hexutil.Uint64, error) {
	// get the value from the contract
	epoch, err := tpc.SfcContract().CurrentSealedEpoch(tpc.DefaultCallOpts())
	if err != nil {
		tpc.log.Errorf("failed to get the current sealed epoch: %s", err.Error())
		return 0, err
	}
	return hexutil.Uint64(epoch.Uint64()), nil
}

// Epoch extract information about an epoch from SFC smart contract.
func (tpc *TpcBridge) Epoch(id hexutil.Uint64) (*types.Epoch, error) {
	// extract epoch snapshot
	epo, err := tpc.SfcContract().GetEpochSnapshot(nil, big.NewInt(int64(id)))
	if err != nil {
		tpc.log.Errorf("failed to extract epoch information: %s", err.Error())
		return nil, err
	}

	return &types.Epoch{
		Id:                    id,
		EndTime:               hexutil.Uint64(epo.EndTime.Uint64()),
		EpochFee:              (hexutil.Big)(*epo.EpochFee),
		TotalBaseRewardWeight: (hexutil.Big)(*epo.TotalBaseRewardWeight),
		TotalTxRewardWeight:   (hexutil.Big)(*epo.TotalTxRewardWeight),
		BaseRewardPerSecond:   (hexutil.Big)(*epo.BaseRewardPerSecond),
		StakeTotalAmount:      (hexutil.Big)(*epo.TotalStake),
		TotalSupply:           (hexutil.Big)(*epo.TotalSupply),
	}, nil
}

// RewardsAllowed returns if the rewards can be manipulated with.
func (tpc *TpcBridge) RewardsAllowed() (bool, error) {
	tpc.log.Debug("rewards lock always open")
	return true, nil
}

// LockingAllowed indicates if the stake locking has been enabled in SFC.
func (tpc *TpcBridge) LockingAllowed() (bool, error) {
	// get the current sealed epoch value from the contract
	epoch, err := tpc.SfcContract().CurrentSealedEpoch(nil)
	if err != nil {
		tpc.log.Errorf("failed to get the current sealed epoch: %s", err.Error())
		return false, err
	}

	return epoch.Uint64() >= sfcFirstLockEpoch, nil
}

// TotalStaked returns the total amount of staked tokens.
func (tpc *TpcBridge) TotalStaked() (*big.Int, error) {
	return tpc.SfcContract().TotalStake(tpc.DefaultCallOpts())
}

// SfcMinValidatorStake extracts a value of minimal validator self stake.
func (tpc *TpcBridge) SfcMinValidatorStake() (*big.Int, error) {
	return tpc.SfcContract().MinSelfStake(tpc.DefaultCallOpts())
}

// SfcMaxDelegatedRatio extracts a ratio between self delegation and received stake.
func (tpc *TpcBridge) SfcMaxDelegatedRatio() (*big.Int, error) {
	return tpc.SfcContract().MaxDelegatedRatio(tpc.DefaultCallOpts())
}

// SfcMinLockupDuration extracts a minimal lockup duration.
func (tpc *TpcBridge) SfcMinLockupDuration() (*big.Int, error) {
	return tpc.SfcContract().MinLockupDuration(tpc.DefaultCallOpts())
}

// SfcMaxLockupDuration extracts a maximal lockup duration.
func (tpc *TpcBridge) SfcMaxLockupDuration() (*big.Int, error) {
	return tpc.SfcContract().MaxLockupDuration(tpc.DefaultCallOpts())
}

// SfcWithdrawalPeriodEpochs extracts a minimal number of epochs between un-delegate and withdraw.
func (tpc *TpcBridge) SfcWithdrawalPeriodEpochs() (*big.Int, error) {
	return tpc.SfcContract().WithdrawalPeriodEpochs(tpc.DefaultCallOpts())
}

// SfcWithdrawalPeriodTime extracts a minimal number of seconds between un-delegate and withdraw.
func (tpc *TpcBridge) SfcWithdrawalPeriodTime() (*big.Int, error) {
	return tpc.SfcContract().WithdrawalPeriodTime(tpc.DefaultCallOpts())
}
