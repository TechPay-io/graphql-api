// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ErcWrappedTpcMetaData contains all meta data concerning the ErcWrappedTpc contract.
var ErcWrappedTpcMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_INVALID_ZERO_VALUE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_NO_ERROR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPauser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renouncePauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ErcWrappedTpcABI is the input ABI used to generate the binding from.
// Deprecated: Use ErcWrappedTpcMetaData.ABI instead.
var ErcWrappedTpcABI = ErcWrappedTpcMetaData.ABI

// ErcWrappedTpc is an auto generated Go binding around an Ethereum contract.
type ErcWrappedTpc struct {
	ErcWrappedTpcCaller     // Read-only binding to the contract
	ErcWrappedTpcTransactor // Write-only binding to the contract
	ErcWrappedTpcFilterer   // Log filterer for contract events
}

// ErcWrappedTpcCaller is an auto generated read-only Go binding around an Ethereum contract.
type ErcWrappedTpcCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ErcWrappedTpcTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ErcWrappedTpcTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ErcWrappedTpcFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ErcWrappedTpcFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ErcWrappedTpcSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ErcWrappedTpcSession struct {
	Contract     *ErcWrappedTpc    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ErcWrappedTpcCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ErcWrappedTpcCallerSession struct {
	Contract *ErcWrappedTpcCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ErcWrappedTpcTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ErcWrappedTpcTransactorSession struct {
	Contract     *ErcWrappedTpcTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ErcWrappedTpcRaw is an auto generated low-level Go binding around an Ethereum contract.
type ErcWrappedTpcRaw struct {
	Contract *ErcWrappedTpc // Generic contract binding to access the raw methods on
}

// ErcWrappedTpcCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ErcWrappedTpcCallerRaw struct {
	Contract *ErcWrappedTpcCaller // Generic read-only contract binding to access the raw methods on
}

// ErcWrappedTpcTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ErcWrappedTpcTransactorRaw struct {
	Contract *ErcWrappedTpcTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErcWrappedTpc creates a new instance of ErcWrappedTpc, bound to a specific deployed contract.
func NewErcWrappedTpc(address common.Address, backend bind.ContractBackend) (*ErcWrappedTpc, error) {
	contract, err := bindErcWrappedTpc(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ErcWrappedTpc{ErcWrappedTpcCaller: ErcWrappedTpcCaller{contract: contract}, ErcWrappedTpcTransactor: ErcWrappedTpcTransactor{contract: contract}, ErcWrappedTpcFilterer: ErcWrappedTpcFilterer{contract: contract}}, nil
}

// NewErcWrappedTpcCaller creates a new read-only instance of ErcWrappedTpc, bound to a specific deployed contract.
func NewErcWrappedTpcCaller(address common.Address, caller bind.ContractCaller) (*ErcWrappedTpcCaller, error) {
	contract, err := bindErcWrappedTpc(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ErcWrappedTpcCaller{contract: contract}, nil
}

// NewErcWrappedTpcTransactor creates a new write-only instance of ErcWrappedTpc, bound to a specific deployed contract.
func NewErcWrappedTpcTransactor(address common.Address, transactor bind.ContractTransactor) (*ErcWrappedTpcTransactor, error) {
	contract, err := bindErcWrappedTpc(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ErcWrappedTpcTransactor{contract: contract}, nil
}

// NewErcWrappedTpcFilterer creates a new log filterer instance of ErcWrappedTpc, bound to a specific deployed contract.
func NewErcWrappedTpcFilterer(address common.Address, filterer bind.ContractFilterer) (*ErcWrappedTpcFilterer, error) {
	contract, err := bindErcWrappedTpc(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ErcWrappedTpcFilterer{contract: contract}, nil
}

// bindErcWrappedTpc binds a generic wrapper to an already deployed contract.
func bindErcWrappedTpc(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ErcWrappedTpcABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ErcWrappedTpc *ErcWrappedTpcRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ErcWrappedTpc.Contract.ErcWrappedTpcCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ErcWrappedTpc *ErcWrappedTpcRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ErcWrappedTpc.Contract.ErcWrappedTpcTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ErcWrappedTpc *ErcWrappedTpcRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ErcWrappedTpc.Contract.ErcWrappedTpcTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ErcWrappedTpc *ErcWrappedTpcCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ErcWrappedTpc.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ErcWrappedTpc *ErcWrappedTpcTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ErcWrappedTpc.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ErcWrappedTpc *ErcWrappedTpcTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ErcWrappedTpc.Contract.contract.Transact(opts, method, params...)
}

// ERRINVALIDZEROVALUE is a free data retrieval call binding the contract method 0x6d7497b3.
//
// Solidity: function ERR_INVALID_ZERO_VALUE() view returns(uint256)
func (_ErcWrappedTpc *ErcWrappedTpcCaller) ERRINVALIDZEROVALUE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ErcWrappedTpc.contract.Call(opts, &out, "ERR_INVALID_ZERO_VALUE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ERRINVALIDZEROVALUE is a free data retrieval call binding the contract method 0x6d7497b3.
//
// Solidity: function ERR_INVALID_ZERO_VALUE() view returns(uint256)
func (_ErcWrappedTpc *ErcWrappedTpcSession) ERRINVALIDZEROVALUE() (*big.Int, error) {
	return _ErcWrappedTpc.Contract.ERRINVALIDZEROVALUE(&_ErcWrappedTpc.CallOpts)
}

// ERRINVALIDZEROVALUE is a free data retrieval call binding the contract method 0x6d7497b3.
//
// Solidity: function ERR_INVALID_ZERO_VALUE() view returns(uint256)
func (_ErcWrappedTpc *ErcWrappedTpcCallerSession) ERRINVALIDZEROVALUE() (*big.Int, error) {
	return _ErcWrappedTpc.Contract.ERRINVALIDZEROVALUE(&_ErcWrappedTpc.CallOpts)
}

// ERRNOERROR is a free data retrieval call binding the contract method 0x35052d6e.
//
// Solidity: function ERR_NO_ERROR() view returns(uint256)
func (_ErcWrappedTpc *ErcWrappedTpcCaller) ERRNOERROR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ErcWrappedTpc.contract.Call(opts, &out, "ERR_NO_ERROR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ERRNOERROR is a free data retrieval call binding the contract method 0x35052d6e.
//
// Solidity: function ERR_NO_ERROR() view returns(uint256)
func (_ErcWrappedTpc *ErcWrappedTpcSession) ERRNOERROR() (*big.Int, error) {
	return _ErcWrappedTpc.Contract.ERRNOERROR(&_ErcWrappedTpc.CallOpts)
}

// ERRNOERROR is a free data retrieval call binding the contract method 0x35052d6e.
//
// Solidity: function ERR_NO_ERROR() view returns(uint256)
func (_ErcWrappedTpc *ErcWrappedTpcCallerSession) ERRNOERROR() (*big.Int, error) {
	return _ErcWrappedTpc.Contract.ERRNOERROR(&_ErcWrappedTpc.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ErcWrappedTpc *ErcWrappedTpcCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ErcWrappedTpc.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ErcWrappedTpc *ErcWrappedTpcSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ErcWrappedTpc.Contract.Allowance(&_ErcWrappedTpc.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ErcWrappedTpc *ErcWrappedTpcCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ErcWrappedTpc.Contract.Allowance(&_ErcWrappedTpc.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ErcWrappedTpc *ErcWrappedTpcCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ErcWrappedTpc.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ErcWrappedTpc *ErcWrappedTpcSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ErcWrappedTpc.Contract.BalanceOf(&_ErcWrappedTpc.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ErcWrappedTpc *ErcWrappedTpcCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ErcWrappedTpc.Contract.BalanceOf(&_ErcWrappedTpc.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ErcWrappedTpc *ErcWrappedTpcCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ErcWrappedTpc.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ErcWrappedTpc *ErcWrappedTpcSession) Decimals() (uint8, error) {
	return _ErcWrappedTpc.Contract.Decimals(&_ErcWrappedTpc.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ErcWrappedTpc *ErcWrappedTpcCallerSession) Decimals() (uint8, error) {
	return _ErcWrappedTpc.Contract.Decimals(&_ErcWrappedTpc.CallOpts)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_ErcWrappedTpc *ErcWrappedTpcCaller) IsPauser(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _ErcWrappedTpc.contract.Call(opts, &out, "isPauser", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_ErcWrappedTpc *ErcWrappedTpcSession) IsPauser(account common.Address) (bool, error) {
	return _ErcWrappedTpc.Contract.IsPauser(&_ErcWrappedTpc.CallOpts, account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_ErcWrappedTpc *ErcWrappedTpcCallerSession) IsPauser(account common.Address) (bool, error) {
	return _ErcWrappedTpc.Contract.IsPauser(&_ErcWrappedTpc.CallOpts, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ErcWrappedTpc *ErcWrappedTpcCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ErcWrappedTpc.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ErcWrappedTpc *ErcWrappedTpcSession) Name() (string, error) {
	return _ErcWrappedTpc.Contract.Name(&_ErcWrappedTpc.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ErcWrappedTpc *ErcWrappedTpcCallerSession) Name() (string, error) {
	return _ErcWrappedTpc.Contract.Name(&_ErcWrappedTpc.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ErcWrappedTpc *ErcWrappedTpcCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ErcWrappedTpc.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ErcWrappedTpc *ErcWrappedTpcSession) Paused() (bool, error) {
	return _ErcWrappedTpc.Contract.Paused(&_ErcWrappedTpc.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ErcWrappedTpc *ErcWrappedTpcCallerSession) Paused() (bool, error) {
	return _ErcWrappedTpc.Contract.Paused(&_ErcWrappedTpc.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ErcWrappedTpc *ErcWrappedTpcCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ErcWrappedTpc.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ErcWrappedTpc *ErcWrappedTpcSession) Symbol() (string, error) {
	return _ErcWrappedTpc.Contract.Symbol(&_ErcWrappedTpc.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ErcWrappedTpc *ErcWrappedTpcCallerSession) Symbol() (string, error) {
	return _ErcWrappedTpc.Contract.Symbol(&_ErcWrappedTpc.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ErcWrappedTpc *ErcWrappedTpcCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ErcWrappedTpc.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ErcWrappedTpc *ErcWrappedTpcSession) TotalSupply() (*big.Int, error) {
	return _ErcWrappedTpc.Contract.TotalSupply(&_ErcWrappedTpc.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ErcWrappedTpc *ErcWrappedTpcCallerSession) TotalSupply() (*big.Int, error) {
	return _ErcWrappedTpc.Contract.TotalSupply(&_ErcWrappedTpc.CallOpts)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_ErcWrappedTpc *ErcWrappedTpcTransactor) AddPauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ErcWrappedTpc.contract.Transact(opts, "addPauser", account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_ErcWrappedTpc *ErcWrappedTpcSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _ErcWrappedTpc.Contract.AddPauser(&_ErcWrappedTpc.TransactOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_ErcWrappedTpc *ErcWrappedTpcTransactorSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _ErcWrappedTpc.Contract.AddPauser(&_ErcWrappedTpc.TransactOpts, account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ErcWrappedTpc *ErcWrappedTpcTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ErcWrappedTpc.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ErcWrappedTpc *ErcWrappedTpcSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ErcWrappedTpc.Contract.Approve(&_ErcWrappedTpc.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ErcWrappedTpc *ErcWrappedTpcTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ErcWrappedTpc.Contract.Approve(&_ErcWrappedTpc.TransactOpts, spender, value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ErcWrappedTpc *ErcWrappedTpcTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ErcWrappedTpc.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ErcWrappedTpc *ErcWrappedTpcSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ErcWrappedTpc.Contract.DecreaseAllowance(&_ErcWrappedTpc.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ErcWrappedTpc *ErcWrappedTpcTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ErcWrappedTpc.Contract.DecreaseAllowance(&_ErcWrappedTpc.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns(uint256)
func (_ErcWrappedTpc *ErcWrappedTpcTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ErcWrappedTpc.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns(uint256)
func (_ErcWrappedTpc *ErcWrappedTpcSession) Deposit() (*types.Transaction, error) {
	return _ErcWrappedTpc.Contract.Deposit(&_ErcWrappedTpc.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns(uint256)
func (_ErcWrappedTpc *ErcWrappedTpcTransactorSession) Deposit() (*types.Transaction, error) {
	return _ErcWrappedTpc.Contract.Deposit(&_ErcWrappedTpc.TransactOpts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ErcWrappedTpc *ErcWrappedTpcTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ErcWrappedTpc.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ErcWrappedTpc *ErcWrappedTpcSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ErcWrappedTpc.Contract.IncreaseAllowance(&_ErcWrappedTpc.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ErcWrappedTpc *ErcWrappedTpcTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ErcWrappedTpc.Contract.IncreaseAllowance(&_ErcWrappedTpc.TransactOpts, spender, addedValue)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ErcWrappedTpc *ErcWrappedTpcTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ErcWrappedTpc.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ErcWrappedTpc *ErcWrappedTpcSession) Pause() (*types.Transaction, error) {
	return _ErcWrappedTpc.Contract.Pause(&_ErcWrappedTpc.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ErcWrappedTpc *ErcWrappedTpcTransactorSession) Pause() (*types.Transaction, error) {
	return _ErcWrappedTpc.Contract.Pause(&_ErcWrappedTpc.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_ErcWrappedTpc *ErcWrappedTpcTransactor) RenouncePauser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ErcWrappedTpc.contract.Transact(opts, "renouncePauser")
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_ErcWrappedTpc *ErcWrappedTpcSession) RenouncePauser() (*types.Transaction, error) {
	return _ErcWrappedTpc.Contract.RenouncePauser(&_ErcWrappedTpc.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_ErcWrappedTpc *ErcWrappedTpcTransactorSession) RenouncePauser() (*types.Transaction, error) {
	return _ErcWrappedTpc.Contract.RenouncePauser(&_ErcWrappedTpc.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ErcWrappedTpc *ErcWrappedTpcTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ErcWrappedTpc.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ErcWrappedTpc *ErcWrappedTpcSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ErcWrappedTpc.Contract.Transfer(&_ErcWrappedTpc.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ErcWrappedTpc *ErcWrappedTpcTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ErcWrappedTpc.Contract.Transfer(&_ErcWrappedTpc.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ErcWrappedTpc *ErcWrappedTpcTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ErcWrappedTpc.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ErcWrappedTpc *ErcWrappedTpcSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ErcWrappedTpc.Contract.TransferFrom(&_ErcWrappedTpc.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ErcWrappedTpc *ErcWrappedTpcTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ErcWrappedTpc.Contract.TransferFrom(&_ErcWrappedTpc.TransactOpts, from, to, value)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ErcWrappedTpc *ErcWrappedTpcTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ErcWrappedTpc.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ErcWrappedTpc *ErcWrappedTpcSession) Unpause() (*types.Transaction, error) {
	return _ErcWrappedTpc.Contract.Unpause(&_ErcWrappedTpc.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ErcWrappedTpc *ErcWrappedTpcTransactorSession) Unpause() (*types.Transaction, error) {
	return _ErcWrappedTpc.Contract.Unpause(&_ErcWrappedTpc.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns(uint256)
func (_ErcWrappedTpc *ErcWrappedTpcTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ErcWrappedTpc.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns(uint256)
func (_ErcWrappedTpc *ErcWrappedTpcSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _ErcWrappedTpc.Contract.Withdraw(&_ErcWrappedTpc.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns(uint256)
func (_ErcWrappedTpc *ErcWrappedTpcTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _ErcWrappedTpc.Contract.Withdraw(&_ErcWrappedTpc.TransactOpts, amount)
}

// ErcWrappedTpcApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ErcWrappedTpc contract.
type ErcWrappedTpcApprovalIterator struct {
	Event *ErcWrappedTpcApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ErcWrappedTpcApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ErcWrappedTpcApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ErcWrappedTpcApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ErcWrappedTpcApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ErcWrappedTpcApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ErcWrappedTpcApproval represents a Approval event raised by the ErcWrappedTpc contract.
type ErcWrappedTpcApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ErcWrappedTpc *ErcWrappedTpcFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ErcWrappedTpcApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ErcWrappedTpc.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ErcWrappedTpcApprovalIterator{contract: _ErcWrappedTpc.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ErcWrappedTpc *ErcWrappedTpcFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ErcWrappedTpcApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ErcWrappedTpc.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ErcWrappedTpcApproval)
				if err := _ErcWrappedTpc.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ErcWrappedTpc *ErcWrappedTpcFilterer) ParseApproval(log types.Log) (*ErcWrappedTpcApproval, error) {
	event := new(ErcWrappedTpcApproval)
	if err := _ErcWrappedTpc.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ErcWrappedTpcPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ErcWrappedTpc contract.
type ErcWrappedTpcPausedIterator struct {
	Event *ErcWrappedTpcPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ErcWrappedTpcPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ErcWrappedTpcPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ErcWrappedTpcPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ErcWrappedTpcPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ErcWrappedTpcPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ErcWrappedTpcPaused represents a Paused event raised by the ErcWrappedTpc contract.
type ErcWrappedTpcPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ErcWrappedTpc *ErcWrappedTpcFilterer) FilterPaused(opts *bind.FilterOpts) (*ErcWrappedTpcPausedIterator, error) {

	logs, sub, err := _ErcWrappedTpc.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ErcWrappedTpcPausedIterator{contract: _ErcWrappedTpc.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ErcWrappedTpc *ErcWrappedTpcFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ErcWrappedTpcPaused) (event.Subscription, error) {

	logs, sub, err := _ErcWrappedTpc.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ErcWrappedTpcPaused)
				if err := _ErcWrappedTpc.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ErcWrappedTpc *ErcWrappedTpcFilterer) ParsePaused(log types.Log) (*ErcWrappedTpcPaused, error) {
	event := new(ErcWrappedTpcPaused)
	if err := _ErcWrappedTpc.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ErcWrappedTpcPauserAddedIterator is returned from FilterPauserAdded and is used to iterate over the raw logs and unpacked data for PauserAdded events raised by the ErcWrappedTpc contract.
type ErcWrappedTpcPauserAddedIterator struct {
	Event *ErcWrappedTpcPauserAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ErcWrappedTpcPauserAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ErcWrappedTpcPauserAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ErcWrappedTpcPauserAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ErcWrappedTpcPauserAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ErcWrappedTpcPauserAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ErcWrappedTpcPauserAdded represents a PauserAdded event raised by the ErcWrappedTpc contract.
type ErcWrappedTpcPauserAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserAdded is a free log retrieval operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address indexed account)
func (_ErcWrappedTpc *ErcWrappedTpcFilterer) FilterPauserAdded(opts *bind.FilterOpts, account []common.Address) (*ErcWrappedTpcPauserAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ErcWrappedTpc.contract.FilterLogs(opts, "PauserAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &ErcWrappedTpcPauserAddedIterator{contract: _ErcWrappedTpc.contract, event: "PauserAdded", logs: logs, sub: sub}, nil
}

// WatchPauserAdded is a free log subscription operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address indexed account)
func (_ErcWrappedTpc *ErcWrappedTpcFilterer) WatchPauserAdded(opts *bind.WatchOpts, sink chan<- *ErcWrappedTpcPauserAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ErcWrappedTpc.contract.WatchLogs(opts, "PauserAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ErcWrappedTpcPauserAdded)
				if err := _ErcWrappedTpc.contract.UnpackLog(event, "PauserAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePauserAdded is a log parse operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address indexed account)
func (_ErcWrappedTpc *ErcWrappedTpcFilterer) ParsePauserAdded(log types.Log) (*ErcWrappedTpcPauserAdded, error) {
	event := new(ErcWrappedTpcPauserAdded)
	if err := _ErcWrappedTpc.contract.UnpackLog(event, "PauserAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ErcWrappedTpcPauserRemovedIterator is returned from FilterPauserRemoved and is used to iterate over the raw logs and unpacked data for PauserRemoved events raised by the ErcWrappedTpc contract.
type ErcWrappedTpcPauserRemovedIterator struct {
	Event *ErcWrappedTpcPauserRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ErcWrappedTpcPauserRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ErcWrappedTpcPauserRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ErcWrappedTpcPauserRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ErcWrappedTpcPauserRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ErcWrappedTpcPauserRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ErcWrappedTpcPauserRemoved represents a PauserRemoved event raised by the ErcWrappedTpc contract.
type ErcWrappedTpcPauserRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserRemoved is a free log retrieval operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address indexed account)
func (_ErcWrappedTpc *ErcWrappedTpcFilterer) FilterPauserRemoved(opts *bind.FilterOpts, account []common.Address) (*ErcWrappedTpcPauserRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ErcWrappedTpc.contract.FilterLogs(opts, "PauserRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &ErcWrappedTpcPauserRemovedIterator{contract: _ErcWrappedTpc.contract, event: "PauserRemoved", logs: logs, sub: sub}, nil
}

// WatchPauserRemoved is a free log subscription operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address indexed account)
func (_ErcWrappedTpc *ErcWrappedTpcFilterer) WatchPauserRemoved(opts *bind.WatchOpts, sink chan<- *ErcWrappedTpcPauserRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ErcWrappedTpc.contract.WatchLogs(opts, "PauserRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ErcWrappedTpcPauserRemoved)
				if err := _ErcWrappedTpc.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePauserRemoved is a log parse operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address indexed account)
func (_ErcWrappedTpc *ErcWrappedTpcFilterer) ParsePauserRemoved(log types.Log) (*ErcWrappedTpcPauserRemoved, error) {
	event := new(ErcWrappedTpcPauserRemoved)
	if err := _ErcWrappedTpc.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ErcWrappedTpcTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ErcWrappedTpc contract.
type ErcWrappedTpcTransferIterator struct {
	Event *ErcWrappedTpcTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ErcWrappedTpcTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ErcWrappedTpcTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ErcWrappedTpcTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ErcWrappedTpcTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ErcWrappedTpcTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ErcWrappedTpcTransfer represents a Transfer event raised by the ErcWrappedTpc contract.
type ErcWrappedTpcTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ErcWrappedTpc *ErcWrappedTpcFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ErcWrappedTpcTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ErcWrappedTpc.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ErcWrappedTpcTransferIterator{contract: _ErcWrappedTpc.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ErcWrappedTpc *ErcWrappedTpcFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ErcWrappedTpcTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ErcWrappedTpc.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ErcWrappedTpcTransfer)
				if err := _ErcWrappedTpc.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ErcWrappedTpc *ErcWrappedTpcFilterer) ParseTransfer(log types.Log) (*ErcWrappedTpcTransfer, error) {
	event := new(ErcWrappedTpcTransfer)
	if err := _ErcWrappedTpc.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ErcWrappedTpcUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ErcWrappedTpc contract.
type ErcWrappedTpcUnpausedIterator struct {
	Event *ErcWrappedTpcUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ErcWrappedTpcUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ErcWrappedTpcUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ErcWrappedTpcUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ErcWrappedTpcUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ErcWrappedTpcUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ErcWrappedTpcUnpaused represents a Unpaused event raised by the ErcWrappedTpc contract.
type ErcWrappedTpcUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ErcWrappedTpc *ErcWrappedTpcFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ErcWrappedTpcUnpausedIterator, error) {

	logs, sub, err := _ErcWrappedTpc.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ErcWrappedTpcUnpausedIterator{contract: _ErcWrappedTpc.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ErcWrappedTpc *ErcWrappedTpcFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ErcWrappedTpcUnpaused) (event.Subscription, error) {

	logs, sub, err := _ErcWrappedTpc.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ErcWrappedTpcUnpaused)
				if err := _ErcWrappedTpc.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ErcWrappedTpc *ErcWrappedTpcFilterer) ParseUnpaused(log types.Log) (*ErcWrappedTpcUnpaused, error) {
	event := new(ErcWrappedTpcUnpaused)
	if err := _ErcWrappedTpc.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
