// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hpblock

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

// HpblockMetaData contains all meta data concerning the Hpblock contract.
var HpblockMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"nodeToLockAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"nodeAddr\",\"type\":\"address\"}],\"name\":\"stake\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"deleteAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAdmins\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"lockBal\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"nodeAddr\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"nodeAddr\",\"type\":\"address\"}],\"name\":\"fetchLockInfoByNodeAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNodeContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"lockTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"adminMap\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setNodeContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ReceivedHpb\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"nodeAddr\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"lockNum\",\"type\":\"uint256\"}],\"name\":\"LockBal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"nodeAddr\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"lockNum\",\"type\":\"uint256\"}],\"name\":\"withdrawBal\",\"type\":\"event\"}]",
}

// HpblockABI is the input ABI used to generate the binding from.
// Deprecated: Use HpblockMetaData.ABI instead.
var HpblockABI = HpblockMetaData.ABI

// Hpblock is an auto generated Go binding around an Ethereum contract.
type Hpblock struct {
	HpblockCaller     // Read-only binding to the contract
	HpblockTransactor // Write-only binding to the contract
	HpblockFilterer   // Log filterer for contract events
}

// HpblockCaller is an auto generated read-only Go binding around an Ethereum contract.
type HpblockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HpblockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HpblockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HpblockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HpblockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HpblockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HpblockSession struct {
	Contract     *Hpblock          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HpblockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HpblockCallerSession struct {
	Contract *HpblockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// HpblockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HpblockTransactorSession struct {
	Contract     *HpblockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// HpblockRaw is an auto generated low-level Go binding around an Ethereum contract.
type HpblockRaw struct {
	Contract *Hpblock // Generic contract binding to access the raw methods on
}

// HpblockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HpblockCallerRaw struct {
	Contract *HpblockCaller // Generic read-only contract binding to access the raw methods on
}

// HpblockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HpblockTransactorRaw struct {
	Contract *HpblockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHpblock creates a new instance of Hpblock, bound to a specific deployed contract.
func NewHpblock(address common.Address, backend bind.ContractBackend) (*Hpblock, error) {
	contract, err := bindHpblock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Hpblock{HpblockCaller: HpblockCaller{contract: contract}, HpblockTransactor: HpblockTransactor{contract: contract}, HpblockFilterer: HpblockFilterer{contract: contract}}, nil
}

// NewHpblockCaller creates a new read-only instance of Hpblock, bound to a specific deployed contract.
func NewHpblockCaller(address common.Address, caller bind.ContractCaller) (*HpblockCaller, error) {
	contract, err := bindHpblock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HpblockCaller{contract: contract}, nil
}

// NewHpblockTransactor creates a new write-only instance of Hpblock, bound to a specific deployed contract.
func NewHpblockTransactor(address common.Address, transactor bind.ContractTransactor) (*HpblockTransactor, error) {
	contract, err := bindHpblock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HpblockTransactor{contract: contract}, nil
}

// NewHpblockFilterer creates a new log filterer instance of Hpblock, bound to a specific deployed contract.
func NewHpblockFilterer(address common.Address, filterer bind.ContractFilterer) (*HpblockFilterer, error) {
	contract, err := bindHpblock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HpblockFilterer{contract: contract}, nil
}

// bindHpblock binds a generic wrapper to an already deployed contract.
func bindHpblock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HpblockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hpblock *HpblockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hpblock.Contract.HpblockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hpblock *HpblockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hpblock.Contract.HpblockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hpblock *HpblockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hpblock.Contract.HpblockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hpblock *HpblockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hpblock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hpblock *HpblockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hpblock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hpblock *HpblockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hpblock.Contract.contract.Transact(opts, method, params...)
}

// AdminMap is a free data retrieval call binding the contract method 0xdbbc830b.
//
// Solidity: function adminMap(address ) view returns(address)
func (_Hpblock *HpblockCaller) AdminMap(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Hpblock.contract.Call(opts, &out, "adminMap", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AdminMap is a free data retrieval call binding the contract method 0xdbbc830b.
//
// Solidity: function adminMap(address ) view returns(address)
func (_Hpblock *HpblockSession) AdminMap(arg0 common.Address) (common.Address, error) {
	return _Hpblock.Contract.AdminMap(&_Hpblock.CallOpts, arg0)
}

// AdminMap is a free data retrieval call binding the contract method 0xdbbc830b.
//
// Solidity: function adminMap(address ) view returns(address)
func (_Hpblock *HpblockCallerSession) AdminMap(arg0 common.Address) (common.Address, error) {
	return _Hpblock.Contract.AdminMap(&_Hpblock.CallOpts, arg0)
}

// FetchLockInfoByNodeAddr is a free data retrieval call binding the contract method 0x695c6269.
//
// Solidity: function fetchLockInfoByNodeAddr(address nodeAddr) view returns(address, uint256)
func (_Hpblock *HpblockCaller) FetchLockInfoByNodeAddr(opts *bind.CallOpts, nodeAddr common.Address) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _Hpblock.contract.Call(opts, &out, "fetchLockInfoByNodeAddr", nodeAddr)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// FetchLockInfoByNodeAddr is a free data retrieval call binding the contract method 0x695c6269.
//
// Solidity: function fetchLockInfoByNodeAddr(address nodeAddr) view returns(address, uint256)
func (_Hpblock *HpblockSession) FetchLockInfoByNodeAddr(nodeAddr common.Address) (common.Address, *big.Int, error) {
	return _Hpblock.Contract.FetchLockInfoByNodeAddr(&_Hpblock.CallOpts, nodeAddr)
}

// FetchLockInfoByNodeAddr is a free data retrieval call binding the contract method 0x695c6269.
//
// Solidity: function fetchLockInfoByNodeAddr(address nodeAddr) view returns(address, uint256)
func (_Hpblock *HpblockCallerSession) FetchLockInfoByNodeAddr(nodeAddr common.Address) (common.Address, *big.Int, error) {
	return _Hpblock.Contract.FetchLockInfoByNodeAddr(&_Hpblock.CallOpts, nodeAddr)
}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[])
func (_Hpblock *HpblockCaller) GetAdmins(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Hpblock.contract.Call(opts, &out, "getAdmins")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[])
func (_Hpblock *HpblockSession) GetAdmins() ([]common.Address, error) {
	return _Hpblock.Contract.GetAdmins(&_Hpblock.CallOpts)
}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[])
func (_Hpblock *HpblockCallerSession) GetAdmins() ([]common.Address, error) {
	return _Hpblock.Contract.GetAdmins(&_Hpblock.CallOpts)
}

// GetNodeContract is a free data retrieval call binding the contract method 0x9fd2e0be.
//
// Solidity: function getNodeContract() view returns(address)
func (_Hpblock *HpblockCaller) GetNodeContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hpblock.contract.Call(opts, &out, "getNodeContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNodeContract is a free data retrieval call binding the contract method 0x9fd2e0be.
//
// Solidity: function getNodeContract() view returns(address)
func (_Hpblock *HpblockSession) GetNodeContract() (common.Address, error) {
	return _Hpblock.Contract.GetNodeContract(&_Hpblock.CallOpts)
}

// GetNodeContract is a free data retrieval call binding the contract method 0x9fd2e0be.
//
// Solidity: function getNodeContract() view returns(address)
func (_Hpblock *HpblockCallerSession) GetNodeContract() (common.Address, error) {
	return _Hpblock.Contract.GetNodeContract(&_Hpblock.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Hpblock *HpblockCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hpblock.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Hpblock *HpblockSession) GetOwner() (common.Address, error) {
	return _Hpblock.Contract.GetOwner(&_Hpblock.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Hpblock *HpblockCallerSession) GetOwner() (common.Address, error) {
	return _Hpblock.Contract.GetOwner(&_Hpblock.CallOpts)
}

// LockBal is a free data retrieval call binding the contract method 0x31dddf75.
//
// Solidity: function lockBal(address ) view returns(uint256)
func (_Hpblock *HpblockCaller) LockBal(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Hpblock.contract.Call(opts, &out, "lockBal", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockBal is a free data retrieval call binding the contract method 0x31dddf75.
//
// Solidity: function lockBal(address ) view returns(uint256)
func (_Hpblock *HpblockSession) LockBal(arg0 common.Address) (*big.Int, error) {
	return _Hpblock.Contract.LockBal(&_Hpblock.CallOpts, arg0)
}

// LockBal is a free data retrieval call binding the contract method 0x31dddf75.
//
// Solidity: function lockBal(address ) view returns(uint256)
func (_Hpblock *HpblockCallerSession) LockBal(arg0 common.Address) (*big.Int, error) {
	return _Hpblock.Contract.LockBal(&_Hpblock.CallOpts, arg0)
}

// LockTime is a free data retrieval call binding the contract method 0xa4beda63.
//
// Solidity: function lockTime(address ) view returns(uint256)
func (_Hpblock *HpblockCaller) LockTime(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Hpblock.contract.Call(opts, &out, "lockTime", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockTime is a free data retrieval call binding the contract method 0xa4beda63.
//
// Solidity: function lockTime(address ) view returns(uint256)
func (_Hpblock *HpblockSession) LockTime(arg0 common.Address) (*big.Int, error) {
	return _Hpblock.Contract.LockTime(&_Hpblock.CallOpts, arg0)
}

// LockTime is a free data retrieval call binding the contract method 0xa4beda63.
//
// Solidity: function lockTime(address ) view returns(uint256)
func (_Hpblock *HpblockCallerSession) LockTime(arg0 common.Address) (*big.Int, error) {
	return _Hpblock.Contract.LockTime(&_Hpblock.CallOpts, arg0)
}

// NodeToLockAddr is a free data retrieval call binding the contract method 0x022277b1.
//
// Solidity: function nodeToLockAddr(address ) view returns(address)
func (_Hpblock *HpblockCaller) NodeToLockAddr(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Hpblock.contract.Call(opts, &out, "nodeToLockAddr", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NodeToLockAddr is a free data retrieval call binding the contract method 0x022277b1.
//
// Solidity: function nodeToLockAddr(address ) view returns(address)
func (_Hpblock *HpblockSession) NodeToLockAddr(arg0 common.Address) (common.Address, error) {
	return _Hpblock.Contract.NodeToLockAddr(&_Hpblock.CallOpts, arg0)
}

// NodeToLockAddr is a free data retrieval call binding the contract method 0x022277b1.
//
// Solidity: function nodeToLockAddr(address ) view returns(address)
func (_Hpblock *HpblockCallerSession) NodeToLockAddr(arg0 common.Address) (common.Address, error) {
	return _Hpblock.Contract.NodeToLockAddr(&_Hpblock.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Hpblock *HpblockCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hpblock.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Hpblock *HpblockSession) Owner() (common.Address, error) {
	return _Hpblock.Contract.Owner(&_Hpblock.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Hpblock *HpblockCallerSession) Owner() (common.Address, error) {
	return _Hpblock.Contract.Owner(&_Hpblock.CallOpts)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address addr) returns(bool)
func (_Hpblock *HpblockTransactor) AddAdmin(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Hpblock.contract.Transact(opts, "addAdmin", addr)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address addr) returns(bool)
func (_Hpblock *HpblockSession) AddAdmin(addr common.Address) (*types.Transaction, error) {
	return _Hpblock.Contract.AddAdmin(&_Hpblock.TransactOpts, addr)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address addr) returns(bool)
func (_Hpblock *HpblockTransactorSession) AddAdmin(addr common.Address) (*types.Transaction, error) {
	return _Hpblock.Contract.AddAdmin(&_Hpblock.TransactOpts, addr)
}

// DeleteAdmin is a paid mutator transaction binding the contract method 0x27e1f7df.
//
// Solidity: function deleteAdmin(address addr) returns(bool)
func (_Hpblock *HpblockTransactor) DeleteAdmin(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Hpblock.contract.Transact(opts, "deleteAdmin", addr)
}

// DeleteAdmin is a paid mutator transaction binding the contract method 0x27e1f7df.
//
// Solidity: function deleteAdmin(address addr) returns(bool)
func (_Hpblock *HpblockSession) DeleteAdmin(addr common.Address) (*types.Transaction, error) {
	return _Hpblock.Contract.DeleteAdmin(&_Hpblock.TransactOpts, addr)
}

// DeleteAdmin is a paid mutator transaction binding the contract method 0x27e1f7df.
//
// Solidity: function deleteAdmin(address addr) returns(bool)
func (_Hpblock *HpblockTransactorSession) DeleteAdmin(addr common.Address) (*types.Transaction, error) {
	return _Hpblock.Contract.DeleteAdmin(&_Hpblock.TransactOpts, addr)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() payable returns(bool)
func (_Hpblock *HpblockTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hpblock.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() payable returns(bool)
func (_Hpblock *HpblockSession) Kill() (*types.Transaction, error) {
	return _Hpblock.Contract.Kill(&_Hpblock.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() payable returns(bool)
func (_Hpblock *HpblockTransactorSession) Kill() (*types.Transaction, error) {
	return _Hpblock.Contract.Kill(&_Hpblock.TransactOpts)
}

// SetNodeContract is a paid mutator transaction binding the contract method 0xf50fd926.
//
// Solidity: function setNodeContract(address addr) returns()
func (_Hpblock *HpblockTransactor) SetNodeContract(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Hpblock.contract.Transact(opts, "setNodeContract", addr)
}

// SetNodeContract is a paid mutator transaction binding the contract method 0xf50fd926.
//
// Solidity: function setNodeContract(address addr) returns()
func (_Hpblock *HpblockSession) SetNodeContract(addr common.Address) (*types.Transaction, error) {
	return _Hpblock.Contract.SetNodeContract(&_Hpblock.TransactOpts, addr)
}

// SetNodeContract is a paid mutator transaction binding the contract method 0xf50fd926.
//
// Solidity: function setNodeContract(address addr) returns()
func (_Hpblock *HpblockTransactorSession) SetNodeContract(addr common.Address) (*types.Transaction, error) {
	return _Hpblock.Contract.SetNodeContract(&_Hpblock.TransactOpts, addr)
}

// Stake is a paid mutator transaction binding the contract method 0x26476204.
//
// Solidity: function stake(address nodeAddr) payable returns(bool)
func (_Hpblock *HpblockTransactor) Stake(opts *bind.TransactOpts, nodeAddr common.Address) (*types.Transaction, error) {
	return _Hpblock.contract.Transact(opts, "stake", nodeAddr)
}

// Stake is a paid mutator transaction binding the contract method 0x26476204.
//
// Solidity: function stake(address nodeAddr) payable returns(bool)
func (_Hpblock *HpblockSession) Stake(nodeAddr common.Address) (*types.Transaction, error) {
	return _Hpblock.Contract.Stake(&_Hpblock.TransactOpts, nodeAddr)
}

// Stake is a paid mutator transaction binding the contract method 0x26476204.
//
// Solidity: function stake(address nodeAddr) payable returns(bool)
func (_Hpblock *HpblockTransactorSession) Stake(nodeAddr common.Address) (*types.Transaction, error) {
	return _Hpblock.Contract.Stake(&_Hpblock.TransactOpts, nodeAddr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns(bool)
func (_Hpblock *HpblockTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Hpblock.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns(bool)
func (_Hpblock *HpblockSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Hpblock.Contract.TransferOwnership(&_Hpblock.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns(bool)
func (_Hpblock *HpblockTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Hpblock.Contract.TransferOwnership(&_Hpblock.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address nodeAddr) payable returns(bool)
func (_Hpblock *HpblockTransactor) Withdraw(opts *bind.TransactOpts, nodeAddr common.Address) (*types.Transaction, error) {
	return _Hpblock.contract.Transact(opts, "withdraw", nodeAddr)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address nodeAddr) payable returns(bool)
func (_Hpblock *HpblockSession) Withdraw(nodeAddr common.Address) (*types.Transaction, error) {
	return _Hpblock.Contract.Withdraw(&_Hpblock.TransactOpts, nodeAddr)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address nodeAddr) payable returns(bool)
func (_Hpblock *HpblockTransactorSession) Withdraw(nodeAddr common.Address) (*types.Transaction, error) {
	return _Hpblock.Contract.Withdraw(&_Hpblock.TransactOpts, nodeAddr)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Hpblock *HpblockTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Hpblock.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Hpblock *HpblockSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Hpblock.Contract.Fallback(&_Hpblock.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Hpblock *HpblockTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Hpblock.Contract.Fallback(&_Hpblock.TransactOpts, calldata)
}

// HpblockLockBalIterator is returned from FilterLockBal and is used to iterate over the raw logs and unpacked data for LockBal events raised by the Hpblock contract.
type HpblockLockBalIterator struct {
	Event *HpblockLockBal // Event containing the contract specifics and raw log

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
func (it *HpblockLockBalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HpblockLockBal)
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
		it.Event = new(HpblockLockBal)
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
func (it *HpblockLockBalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HpblockLockBalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HpblockLockBal represents a LockBal event raised by the Hpblock contract.
type HpblockLockBal struct {
	NodeAddr common.Address
	Addr     common.Address
	LockNum  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLockBal is a free log retrieval operation binding the contract event 0xb87d4f14e7ba38c29817fe73e71e3e3a24821eb336c4ea274b891072dbacdf34.
//
// Solidity: event LockBal(address indexed nodeAddr, address indexed addr, uint256 indexed lockNum)
func (_Hpblock *HpblockFilterer) FilterLockBal(opts *bind.FilterOpts, nodeAddr []common.Address, addr []common.Address, lockNum []*big.Int) (*HpblockLockBalIterator, error) {

	var nodeAddrRule []interface{}
	for _, nodeAddrItem := range nodeAddr {
		nodeAddrRule = append(nodeAddrRule, nodeAddrItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var lockNumRule []interface{}
	for _, lockNumItem := range lockNum {
		lockNumRule = append(lockNumRule, lockNumItem)
	}

	logs, sub, err := _Hpblock.contract.FilterLogs(opts, "LockBal", nodeAddrRule, addrRule, lockNumRule)
	if err != nil {
		return nil, err
	}
	return &HpblockLockBalIterator{contract: _Hpblock.contract, event: "LockBal", logs: logs, sub: sub}, nil
}

// WatchLockBal is a free log subscription operation binding the contract event 0xb87d4f14e7ba38c29817fe73e71e3e3a24821eb336c4ea274b891072dbacdf34.
//
// Solidity: event LockBal(address indexed nodeAddr, address indexed addr, uint256 indexed lockNum)
func (_Hpblock *HpblockFilterer) WatchLockBal(opts *bind.WatchOpts, sink chan<- *HpblockLockBal, nodeAddr []common.Address, addr []common.Address, lockNum []*big.Int) (event.Subscription, error) {

	var nodeAddrRule []interface{}
	for _, nodeAddrItem := range nodeAddr {
		nodeAddrRule = append(nodeAddrRule, nodeAddrItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var lockNumRule []interface{}
	for _, lockNumItem := range lockNum {
		lockNumRule = append(lockNumRule, lockNumItem)
	}

	logs, sub, err := _Hpblock.contract.WatchLogs(opts, "LockBal", nodeAddrRule, addrRule, lockNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HpblockLockBal)
				if err := _Hpblock.contract.UnpackLog(event, "LockBal", log); err != nil {
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

// ParseLockBal is a log parse operation binding the contract event 0xb87d4f14e7ba38c29817fe73e71e3e3a24821eb336c4ea274b891072dbacdf34.
//
// Solidity: event LockBal(address indexed nodeAddr, address indexed addr, uint256 indexed lockNum)
func (_Hpblock *HpblockFilterer) ParseLockBal(log types.Log) (*HpblockLockBal, error) {
	event := new(HpblockLockBal)
	if err := _Hpblock.contract.UnpackLog(event, "LockBal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HpblockReceivedHpbIterator is returned from FilterReceivedHpb and is used to iterate over the raw logs and unpacked data for ReceivedHpb events raised by the Hpblock contract.
type HpblockReceivedHpbIterator struct {
	Event *HpblockReceivedHpb // Event containing the contract specifics and raw log

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
func (it *HpblockReceivedHpbIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HpblockReceivedHpb)
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
		it.Event = new(HpblockReceivedHpb)
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
func (it *HpblockReceivedHpbIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HpblockReceivedHpbIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HpblockReceivedHpb represents a ReceivedHpb event raised by the Hpblock contract.
type HpblockReceivedHpb struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedHpb is a free log retrieval operation binding the contract event 0x7129701436f0cdc265d1e2cda298e8a1ccd6ed5fce7f69343e16530b07a2e06e.
//
// Solidity: event ReceivedHpb(address indexed sender, uint256 amount)
func (_Hpblock *HpblockFilterer) FilterReceivedHpb(opts *bind.FilterOpts, sender []common.Address) (*HpblockReceivedHpbIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Hpblock.contract.FilterLogs(opts, "ReceivedHpb", senderRule)
	if err != nil {
		return nil, err
	}
	return &HpblockReceivedHpbIterator{contract: _Hpblock.contract, event: "ReceivedHpb", logs: logs, sub: sub}, nil
}

// WatchReceivedHpb is a free log subscription operation binding the contract event 0x7129701436f0cdc265d1e2cda298e8a1ccd6ed5fce7f69343e16530b07a2e06e.
//
// Solidity: event ReceivedHpb(address indexed sender, uint256 amount)
func (_Hpblock *HpblockFilterer) WatchReceivedHpb(opts *bind.WatchOpts, sink chan<- *HpblockReceivedHpb, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Hpblock.contract.WatchLogs(opts, "ReceivedHpb", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HpblockReceivedHpb)
				if err := _Hpblock.contract.UnpackLog(event, "ReceivedHpb", log); err != nil {
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

// ParseReceivedHpb is a log parse operation binding the contract event 0x7129701436f0cdc265d1e2cda298e8a1ccd6ed5fce7f69343e16530b07a2e06e.
//
// Solidity: event ReceivedHpb(address indexed sender, uint256 amount)
func (_Hpblock *HpblockFilterer) ParseReceivedHpb(log types.Log) (*HpblockReceivedHpb, error) {
	event := new(HpblockReceivedHpb)
	if err := _Hpblock.contract.UnpackLog(event, "ReceivedHpb", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HpblockWithdrawBalIterator is returned from FilterWithdrawBal and is used to iterate over the raw logs and unpacked data for WithdrawBal events raised by the Hpblock contract.
type HpblockWithdrawBalIterator struct {
	Event *HpblockWithdrawBal // Event containing the contract specifics and raw log

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
func (it *HpblockWithdrawBalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HpblockWithdrawBal)
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
		it.Event = new(HpblockWithdrawBal)
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
func (it *HpblockWithdrawBalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HpblockWithdrawBalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HpblockWithdrawBal represents a WithdrawBal event raised by the Hpblock contract.
type HpblockWithdrawBal struct {
	NodeAddr common.Address
	Addr     common.Address
	LockNum  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawBal is a free log retrieval operation binding the contract event 0x8f2663a76951a6b5fefd39a998d0ff0d6fc9abc3ea3b14431557da631aea1df4.
//
// Solidity: event withdrawBal(address indexed nodeAddr, address indexed addr, uint256 indexed lockNum)
func (_Hpblock *HpblockFilterer) FilterWithdrawBal(opts *bind.FilterOpts, nodeAddr []common.Address, addr []common.Address, lockNum []*big.Int) (*HpblockWithdrawBalIterator, error) {

	var nodeAddrRule []interface{}
	for _, nodeAddrItem := range nodeAddr {
		nodeAddrRule = append(nodeAddrRule, nodeAddrItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var lockNumRule []interface{}
	for _, lockNumItem := range lockNum {
		lockNumRule = append(lockNumRule, lockNumItem)
	}

	logs, sub, err := _Hpblock.contract.FilterLogs(opts, "withdrawBal", nodeAddrRule, addrRule, lockNumRule)
	if err != nil {
		return nil, err
	}
	return &HpblockWithdrawBalIterator{contract: _Hpblock.contract, event: "withdrawBal", logs: logs, sub: sub}, nil
}

// WatchWithdrawBal is a free log subscription operation binding the contract event 0x8f2663a76951a6b5fefd39a998d0ff0d6fc9abc3ea3b14431557da631aea1df4.
//
// Solidity: event withdrawBal(address indexed nodeAddr, address indexed addr, uint256 indexed lockNum)
func (_Hpblock *HpblockFilterer) WatchWithdrawBal(opts *bind.WatchOpts, sink chan<- *HpblockWithdrawBal, nodeAddr []common.Address, addr []common.Address, lockNum []*big.Int) (event.Subscription, error) {

	var nodeAddrRule []interface{}
	for _, nodeAddrItem := range nodeAddr {
		nodeAddrRule = append(nodeAddrRule, nodeAddrItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var lockNumRule []interface{}
	for _, lockNumItem := range lockNum {
		lockNumRule = append(lockNumRule, lockNumItem)
	}

	logs, sub, err := _Hpblock.contract.WatchLogs(opts, "withdrawBal", nodeAddrRule, addrRule, lockNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HpblockWithdrawBal)
				if err := _Hpblock.contract.UnpackLog(event, "withdrawBal", log); err != nil {
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

// ParseWithdrawBal is a log parse operation binding the contract event 0x8f2663a76951a6b5fefd39a998d0ff0d6fc9abc3ea3b14431557da631aea1df4.
//
// Solidity: event withdrawBal(address indexed nodeAddr, address indexed addr, uint256 indexed lockNum)
func (_Hpblock *HpblockFilterer) ParseWithdrawBal(log types.Log) (*HpblockWithdrawBal, error) {
	event := new(HpblockWithdrawBal)
	if err := _Hpblock.contract.UnpackLog(event, "withdrawBal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
