// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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
	_ = abi.ConvertType
)

// MockEndPointMetaData contains all meta data concerning the MockEndPoint contract.
var MockEndPointMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_protocolId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_consensusTargetRate\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"}],\"name\":\"addAllowedProtocol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"protocolId\",\"type\":\"bytes32\"}],\"name\":\"allowedProtocolInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isCreated\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"consensusTargetRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"proocolId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedTransmitters\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"protocolId\",\"type\":\"bytes32\"}],\"name\":\"numberOfAllowedTransmitters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600f57600080fd5b506106438061001f6000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806343304c0f146100515780634eeb72de1461008257806355b5190b1461009e578063e20aeb2d146100ce575b600080fd5b61006b600480360381019061006691906102d7565b6100fe565b604051610079929190610338565b60405180910390f35b61009c600480360381019061009791906103f2565b61012f565b005b6100b860048036038101906100b391906104c4565b610250565b6040516100c59190610504565b60405180910390f35b6100e860048036038101906100e391906102d7565b61027f565b6040516100f5919061051f565b60405180910390f35b60026020528060005260406000206000915090508060000160009054906101000a900460ff16908060010154905082565b60016002600086815260200190815260200160002060000160006101000a81548160ff02191690831515021790555082600260008681526020019081526020016000206001018190555060005b8282905081101561024957600180600087815260200190815260200160002060008585858181106101b0576101af61053a565b5b90506020020160208101906101c59190610569565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506000808681526020019081526020016000206000815480929190610239906105c5565b919050555080600101905061017c565b5050505050565b60016020528160005260406000206020528060005260406000206000915091509054906101000a900460ff1681565b60006020528060005260406000206000915090505481565b600080fd5b600080fd5b6000819050919050565b6102b4816102a1565b81146102bf57600080fd5b50565b6000813590506102d1816102ab565b92915050565b6000602082840312156102ed576102ec610297565b5b60006102fb848285016102c2565b91505092915050565b60008115159050919050565b61031981610304565b82525050565b6000819050919050565b6103328161031f565b82525050565b600060408201905061034d6000830185610310565b61035a6020830184610329565b9392505050565b61036a8161031f565b811461037557600080fd5b50565b60008135905061038781610361565b92915050565b600080fd5b600080fd5b600080fd5b60008083601f8401126103b2576103b161038d565b5b8235905067ffffffffffffffff8111156103cf576103ce610392565b5b6020830191508360208202830111156103eb576103ea610397565b5b9250929050565b6000806000806060858703121561040c5761040b610297565b5b600061041a878288016102c2565b945050602061042b87828801610378565b935050604085013567ffffffffffffffff81111561044c5761044b61029c565b5b6104588782880161039c565b925092505092959194509250565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061049182610466565b9050919050565b6104a181610486565b81146104ac57600080fd5b50565b6000813590506104be81610498565b92915050565b600080604083850312156104db576104da610297565b5b60006104e9858286016102c2565b92505060206104fa858286016104af565b9150509250929050565b60006020820190506105196000830184610310565b92915050565b60006020820190506105346000830184610329565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60006020828403121561057f5761057e610297565b5b600061058d848285016104af565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006105d08261031f565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361060257610601610596565b5b60018201905091905056fea26469706673582212200c0d79c2472a9bb2c43c36b90674bcccdf7bdfe3f968b251e1d5738fe08b28ad64736f6c63430008190033",
}

// MockEndPointABI is the input ABI used to generate the binding from.
// Deprecated: Use MockEndPointMetaData.ABI instead.
var MockEndPointABI = MockEndPointMetaData.ABI

// MockEndPointBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MockEndPointMetaData.Bin instead.
var MockEndPointBin = MockEndPointMetaData.Bin

// DeployMockEndPoint deploys a new Ethereum contract, binding an instance of MockEndPoint to it.
func DeployMockEndPoint(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MockEndPoint, error) {
	parsed, err := MockEndPointMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockEndPointBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockEndPoint{MockEndPointCaller: MockEndPointCaller{contract: contract}, MockEndPointTransactor: MockEndPointTransactor{contract: contract}, MockEndPointFilterer: MockEndPointFilterer{contract: contract}}, nil
}

// MockEndPoint is an auto generated Go binding around an Ethereum contract.
type MockEndPoint struct {
	MockEndPointCaller     // Read-only binding to the contract
	MockEndPointTransactor // Write-only binding to the contract
	MockEndPointFilterer   // Log filterer for contract events
}

// MockEndPointCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockEndPointCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockEndPointTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockEndPointTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockEndPointFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockEndPointFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockEndPointSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockEndPointSession struct {
	Contract     *MockEndPoint     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MockEndPointCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockEndPointCallerSession struct {
	Contract *MockEndPointCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// MockEndPointTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockEndPointTransactorSession struct {
	Contract     *MockEndPointTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MockEndPointRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockEndPointRaw struct {
	Contract *MockEndPoint // Generic contract binding to access the raw methods on
}

// MockEndPointCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockEndPointCallerRaw struct {
	Contract *MockEndPointCaller // Generic read-only contract binding to access the raw methods on
}

// MockEndPointTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockEndPointTransactorRaw struct {
	Contract *MockEndPointTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockEndPoint creates a new instance of MockEndPoint, bound to a specific deployed contract.
func NewMockEndPoint(address common.Address, backend bind.ContractBackend) (*MockEndPoint, error) {
	contract, err := bindMockEndPoint(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockEndPoint{MockEndPointCaller: MockEndPointCaller{contract: contract}, MockEndPointTransactor: MockEndPointTransactor{contract: contract}, MockEndPointFilterer: MockEndPointFilterer{contract: contract}}, nil
}

// NewMockEndPointCaller creates a new read-only instance of MockEndPoint, bound to a specific deployed contract.
func NewMockEndPointCaller(address common.Address, caller bind.ContractCaller) (*MockEndPointCaller, error) {
	contract, err := bindMockEndPoint(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockEndPointCaller{contract: contract}, nil
}

// NewMockEndPointTransactor creates a new write-only instance of MockEndPoint, bound to a specific deployed contract.
func NewMockEndPointTransactor(address common.Address, transactor bind.ContractTransactor) (*MockEndPointTransactor, error) {
	contract, err := bindMockEndPoint(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockEndPointTransactor{contract: contract}, nil
}

// NewMockEndPointFilterer creates a new log filterer instance of MockEndPoint, bound to a specific deployed contract.
func NewMockEndPointFilterer(address common.Address, filterer bind.ContractFilterer) (*MockEndPointFilterer, error) {
	contract, err := bindMockEndPoint(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockEndPointFilterer{contract: contract}, nil
}

// bindMockEndPoint binds a generic wrapper to an already deployed contract.
func bindMockEndPoint(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockEndPointMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockEndPoint *MockEndPointRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockEndPoint.Contract.MockEndPointCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockEndPoint *MockEndPointRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockEndPoint.Contract.MockEndPointTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockEndPoint *MockEndPointRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockEndPoint.Contract.MockEndPointTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockEndPoint *MockEndPointCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockEndPoint.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockEndPoint *MockEndPointTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockEndPoint.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockEndPoint *MockEndPointTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockEndPoint.Contract.contract.Transact(opts, method, params...)
}

// AllowedProtocolInfo is a free data retrieval call binding the contract method 0x43304c0f.
//
// Solidity: function allowedProtocolInfo(bytes32 protocolId) view returns(bool isCreated, uint256 consensusTargetRate)
func (_MockEndPoint *MockEndPointCaller) AllowedProtocolInfo(opts *bind.CallOpts, protocolId [32]byte) (struct {
	IsCreated           bool
	ConsensusTargetRate *big.Int
}, error) {
	var out []interface{}
	err := _MockEndPoint.contract.Call(opts, &out, "allowedProtocolInfo", protocolId)

	outstruct := new(struct {
		IsCreated           bool
		ConsensusTargetRate *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsCreated = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ConsensusTargetRate = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// AllowedProtocolInfo is a free data retrieval call binding the contract method 0x43304c0f.
//
// Solidity: function allowedProtocolInfo(bytes32 protocolId) view returns(bool isCreated, uint256 consensusTargetRate)
func (_MockEndPoint *MockEndPointSession) AllowedProtocolInfo(protocolId [32]byte) (struct {
	IsCreated           bool
	ConsensusTargetRate *big.Int
}, error) {
	return _MockEndPoint.Contract.AllowedProtocolInfo(&_MockEndPoint.CallOpts, protocolId)
}

// AllowedProtocolInfo is a free data retrieval call binding the contract method 0x43304c0f.
//
// Solidity: function allowedProtocolInfo(bytes32 protocolId) view returns(bool isCreated, uint256 consensusTargetRate)
func (_MockEndPoint *MockEndPointCallerSession) AllowedProtocolInfo(protocolId [32]byte) (struct {
	IsCreated           bool
	ConsensusTargetRate *big.Int
}, error) {
	return _MockEndPoint.Contract.AllowedProtocolInfo(&_MockEndPoint.CallOpts, protocolId)
}

// AllowedTransmitters is a free data retrieval call binding the contract method 0x55b5190b.
//
// Solidity: function allowedTransmitters(bytes32 proocolId, address ) view returns(bool)
func (_MockEndPoint *MockEndPointCaller) AllowedTransmitters(opts *bind.CallOpts, proocolId [32]byte, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _MockEndPoint.contract.Call(opts, &out, "allowedTransmitters", proocolId, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedTransmitters is a free data retrieval call binding the contract method 0x55b5190b.
//
// Solidity: function allowedTransmitters(bytes32 proocolId, address ) view returns(bool)
func (_MockEndPoint *MockEndPointSession) AllowedTransmitters(proocolId [32]byte, arg1 common.Address) (bool, error) {
	return _MockEndPoint.Contract.AllowedTransmitters(&_MockEndPoint.CallOpts, proocolId, arg1)
}

// AllowedTransmitters is a free data retrieval call binding the contract method 0x55b5190b.
//
// Solidity: function allowedTransmitters(bytes32 proocolId, address ) view returns(bool)
func (_MockEndPoint *MockEndPointCallerSession) AllowedTransmitters(proocolId [32]byte, arg1 common.Address) (bool, error) {
	return _MockEndPoint.Contract.AllowedTransmitters(&_MockEndPoint.CallOpts, proocolId, arg1)
}

// NumberOfAllowedTransmitters is a free data retrieval call binding the contract method 0xe20aeb2d.
//
// Solidity: function numberOfAllowedTransmitters(bytes32 protocolId) view returns(uint256)
func (_MockEndPoint *MockEndPointCaller) NumberOfAllowedTransmitters(opts *bind.CallOpts, protocolId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _MockEndPoint.contract.Call(opts, &out, "numberOfAllowedTransmitters", protocolId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberOfAllowedTransmitters is a free data retrieval call binding the contract method 0xe20aeb2d.
//
// Solidity: function numberOfAllowedTransmitters(bytes32 protocolId) view returns(uint256)
func (_MockEndPoint *MockEndPointSession) NumberOfAllowedTransmitters(protocolId [32]byte) (*big.Int, error) {
	return _MockEndPoint.Contract.NumberOfAllowedTransmitters(&_MockEndPoint.CallOpts, protocolId)
}

// NumberOfAllowedTransmitters is a free data retrieval call binding the contract method 0xe20aeb2d.
//
// Solidity: function numberOfAllowedTransmitters(bytes32 protocolId) view returns(uint256)
func (_MockEndPoint *MockEndPointCallerSession) NumberOfAllowedTransmitters(protocolId [32]byte) (*big.Int, error) {
	return _MockEndPoint.Contract.NumberOfAllowedTransmitters(&_MockEndPoint.CallOpts, protocolId)
}

// AddAllowedProtocol is a paid mutator transaction binding the contract method 0x4eeb72de.
//
// Solidity: function addAllowedProtocol(bytes32 _protocolId, uint256 _consensusTargetRate, address[] _transmitters) returns()
func (_MockEndPoint *MockEndPointTransactor) AddAllowedProtocol(opts *bind.TransactOpts, _protocolId [32]byte, _consensusTargetRate *big.Int, _transmitters []common.Address) (*types.Transaction, error) {
	return _MockEndPoint.contract.Transact(opts, "addAllowedProtocol", _protocolId, _consensusTargetRate, _transmitters)
}

// AddAllowedProtocol is a paid mutator transaction binding the contract method 0x4eeb72de.
//
// Solidity: function addAllowedProtocol(bytes32 _protocolId, uint256 _consensusTargetRate, address[] _transmitters) returns()
func (_MockEndPoint *MockEndPointSession) AddAllowedProtocol(_protocolId [32]byte, _consensusTargetRate *big.Int, _transmitters []common.Address) (*types.Transaction, error) {
	return _MockEndPoint.Contract.AddAllowedProtocol(&_MockEndPoint.TransactOpts, _protocolId, _consensusTargetRate, _transmitters)
}

// AddAllowedProtocol is a paid mutator transaction binding the contract method 0x4eeb72de.
//
// Solidity: function addAllowedProtocol(bytes32 _protocolId, uint256 _consensusTargetRate, address[] _transmitters) returns()
func (_MockEndPoint *MockEndPointTransactorSession) AddAllowedProtocol(_protocolId [32]byte, _consensusTargetRate *big.Int, _transmitters []common.Address) (*types.Transaction, error) {
	return _MockEndPoint.Contract.AddAllowedProtocol(&_MockEndPoint.TransactOpts, _protocolId, _consensusTargetRate, _transmitters)
}
