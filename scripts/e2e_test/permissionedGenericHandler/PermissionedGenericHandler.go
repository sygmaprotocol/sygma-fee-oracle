// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PermissionedGenericHandler

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

// PermissionedGenericHandlerMetaData contains all meta data concerning the PermissionedGenericHandler contract.
var PermissionedGenericHandlerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"FailedHandlerExecution\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractAddressToDepositFunctionDepositorOffset\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractAddressToDepositFunctionSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractAddressToExecuteFunctionSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractAddressToResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"args\",\"type\":\"bytes\"}],\"name\":\"setResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b5060405162001669380380620016698339818101604052810190620000379190620000dc565b8073ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1681525050506200010e565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620000a48262000077565b9050919050565b620000b68162000097565b8114620000c257600080fd5b50565b600081519050620000d681620000ab565b92915050565b600060208284031215620000f557620000f462000072565b5b60006200010584828501620000c5565b91505092915050565b608051611538620001316000396000818161024b0152610a0a01526115386000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c8063c54c2a1111610066578063c54c2a1114610181578063cb624463146101b1578063e248cff2146101e1578063ec97d3b4146101fd578063fa8675b01461022d5761009e565b8063318c136e146100a35780637f79bea8146100c15780638843a70a146100f1578063a5c3a98514610121578063b07e54bb14610151575b600080fd5b6100ab610249565b6040516100b89190610cc7565b60405180910390f35b6100db60048036038101906100d69190610d18565b61026d565b6040516100e89190610d60565b60405180910390f35b61010b60048036038101906101069190610d18565b61028d565b6040516101189190610d94565b60405180910390f35b61013b60048036038101906101369190610d18565b6102a5565b6040516101489190610dea565b60405180910390f35b61016b60048036038101906101669190610ea0565b6102c5565b6040516101789190610fad565b60405180910390f35b61019b60048036038101906101969190610fcf565b610664565b6040516101a89190610cc7565b60405180910390f35b6101cb60048036038101906101c69190610d18565b610697565b6040516101d89190610dea565b60405180910390f35b6101fb60048036038101906101f69190610ffc565b6106b7565b005b61021760048036038101906102129190610d18565b610966565b604051610224919061106b565b60405180910390f35b61024760048036038101906102429190610ea0565b61097e565b005b7f000000000000000000000000000000000000000000000000000000000000000081565b60056020528060005260406000206000915054906101000a900460ff1681565b60036020528060005260406000206000915090505481565b60046020528060005260406000206000915054906101000a900460e01b81565b60606102cf610a08565b6000606084848101906102e291906110b2565b915084846020908460206102f6919061110e565b926103039392919061116e565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509050600080600089815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506000600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050600081111561044b576000816020850101519050606081901c73ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff1614610449576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161044090611206565b60405180910390fd5b505b600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166104d7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104ce90611298565b60405180910390fd5b6000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460e01b9050600060e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916146106565760008185604051602001610586929190611315565b60405160208183030381529060405290506000808573ffffffffffffffffffffffffffffffffffffffff16836040516105bf919061133d565b6000604051808303816000865af19150503d80600081146105fc576040519150601f19603f3d011682016040523d82523d6000602084013e610601565b606091505b509150915081610646576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161063d906113a0565b60405180910390fd5b809850505050505050505061065c565b50505050505b949350505050565b60006020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60026020528060005260406000206000915054906101000a900460e01b81565b6106bf610a08565b6000606083838101906106d291906110b2565b915083836020908460206106e6919061110e565b926106f39392919061116e565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509050600080600087815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600560008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166107fb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107f290611298565b60405180910390fd5b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460e01b9050600060e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161461095d57600081846040516020016108aa929190611315565b604051602081830303815290604052905060008373ffffffffffffffffffffffffffffffffffffffff16826040516108e2919061133d565b6000604051808303816000865af19150503d806000811461091f576040519150601f19603f3d011682016040523d82523d6000602084013e610924565b606091505b505090508061095a577f3f51c54fb334903765b4de0daee60de6eff2a87aef23d5a1f311f4ea2e1b7e9660405160405180910390a15b50505b50505050505050565b60016020528060005260406000206000915090505481565b610986610a08565b6000828260009060049261099c9392919061116e565b906109a791906113d8565b9050600083836004906024926109bf9392919061116e565b906109ca9190611437565b60001c9050600084846024906028926109e59392919061116e565b906109f091906113d8565b90506109ff8787858585610a98565b50505050505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a96576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a8d906114e2565b60405180910390fd5b565b8360008087815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555084600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555082600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908360e01c021790555081600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555080600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908360e01c02179055506001600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050505050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610cb182610c86565b9050919050565b610cc181610ca6565b82525050565b6000602082019050610cdc6000830184610cb8565b92915050565b600080fd5b600080fd5b610cf581610ca6565b8114610d0057600080fd5b50565b600081359050610d1281610cec565b92915050565b600060208284031215610d2e57610d2d610ce2565b5b6000610d3c84828501610d03565b91505092915050565b60008115159050919050565b610d5a81610d45565b82525050565b6000602082019050610d756000830184610d51565b92915050565b6000819050919050565b610d8e81610d7b565b82525050565b6000602082019050610da96000830184610d85565b92915050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b610de481610daf565b82525050565b6000602082019050610dff6000830184610ddb565b92915050565b6000819050919050565b610e1881610e05565b8114610e2357600080fd5b50565b600081359050610e3581610e0f565b92915050565b600080fd5b600080fd5b600080fd5b60008083601f840112610e6057610e5f610e3b565b5b8235905067ffffffffffffffff811115610e7d57610e7c610e40565b5b602083019150836001820283011115610e9957610e98610e45565b5b9250929050565b60008060008060608587031215610eba57610eb9610ce2565b5b6000610ec887828801610e26565b9450506020610ed987828801610d03565b935050604085013567ffffffffffffffff811115610efa57610ef9610ce7565b5b610f0687828801610e4a565b925092505092959194509250565b600081519050919050565b600082825260208201905092915050565b60005b83811015610f4e578082015181840152602081019050610f33565b83811115610f5d576000848401525b50505050565b6000601f19601f8301169050919050565b6000610f7f82610f14565b610f898185610f1f565b9350610f99818560208601610f30565b610fa281610f63565b840191505092915050565b60006020820190508181036000830152610fc78184610f74565b905092915050565b600060208284031215610fe557610fe4610ce2565b5b6000610ff384828501610e26565b91505092915050565b60008060006040848603121561101557611014610ce2565b5b600061102386828701610e26565b935050602084013567ffffffffffffffff81111561104457611043610ce7565b5b61105086828701610e4a565b92509250509250925092565b61106581610e05565b82525050565b6000602082019050611080600083018461105c565b92915050565b61108f81610d7b565b811461109a57600080fd5b50565b6000813590506110ac81611086565b92915050565b6000602082840312156110c8576110c7610ce2565b5b60006110d68482850161109d565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061111982610d7b565b915061112483610d7b565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611159576111586110df565b5b828201905092915050565b600080fd5b600080fd5b6000808585111561118257611181611164565b5b8386111561119357611192611169565b5b6001850283019150848603905094509492505050565b600082825260208201905092915050565b7f696e636f7272656374206465706f7369746f7220696e20746865206461746100600082015250565b60006111f0601f836111a9565b91506111fb826111ba565b602082019050919050565b6000602082019050818103600083015261121f816111e3565b9050919050565b7f70726f766964656420636f6e747261637441646472657373206973206e6f742060008201527f77686974656c6973746564000000000000000000000000000000000000000000602082015250565b6000611282602b836111a9565b915061128d82611226565b604082019050919050565b600060208201905081810360008301526112b181611275565b9050919050565b6000819050919050565b6112d36112ce82610daf565b6112b8565b82525050565b600081905092915050565b60006112ef82610f14565b6112f981856112d9565b9350611309818560208601610f30565b80840191505092915050565b600061132182856112c2565b60048201915061133182846112e4565b91508190509392505050565b600061134982846112e4565b915081905092915050565b7f63616c6c20746f20636f6e747261637441646472657373206661696c65640000600082015250565b600061138a601e836111a9565b915061139582611354565b602082019050919050565b600060208201905081810360008301526113b98161137d565b9050919050565b600082905092915050565b600082821b905092915050565b60006113e483836113c0565b826113ef8135610daf565b9250600482101561142f5761142a7fffffffff00000000000000000000000000000000000000000000000000000000836004036008026113cb565b831692505b505092915050565b600061144383836113c0565b8261144e8135610e05565b9250602082101561148e576114897fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff836020036008026113cb565b831692505b505092915050565b7f73656e646572206d7573742062652062726964676520636f6e74726163740000600082015250565b60006114cc601e836111a9565b91506114d782611496565b602082019050919050565b600060208201905081810360008301526114fb816114bf565b905091905056fea26469706673582212206c83827d9106b12a5707141314556c0228427414e4cb8173b7fe7d3e3ba6d85664736f6c634300080b0033",
}

// PermissionedGenericHandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use PermissionedGenericHandlerMetaData.ABI instead.
var PermissionedGenericHandlerABI = PermissionedGenericHandlerMetaData.ABI

// PermissionedGenericHandlerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PermissionedGenericHandlerMetaData.Bin instead.
var PermissionedGenericHandlerBin = PermissionedGenericHandlerMetaData.Bin

// DeployPermissionedGenericHandler deploys a new Ethereum contract, binding an instance of PermissionedGenericHandler to it.
func DeployPermissionedGenericHandler(auth *bind.TransactOpts, backend bind.ContractBackend, bridgeAddress common.Address) (common.Address, *types.Transaction, *PermissionedGenericHandler, error) {
	parsed, err := PermissionedGenericHandlerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PermissionedGenericHandlerBin), backend, bridgeAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PermissionedGenericHandler{PermissionedGenericHandlerCaller: PermissionedGenericHandlerCaller{contract: contract}, PermissionedGenericHandlerTransactor: PermissionedGenericHandlerTransactor{contract: contract}, PermissionedGenericHandlerFilterer: PermissionedGenericHandlerFilterer{contract: contract}}, nil
}

// PermissionedGenericHandler is an auto generated Go binding around an Ethereum contract.
type PermissionedGenericHandler struct {
	PermissionedGenericHandlerCaller     // Read-only binding to the contract
	PermissionedGenericHandlerTransactor // Write-only binding to the contract
	PermissionedGenericHandlerFilterer   // Log filterer for contract events
}

// PermissionedGenericHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type PermissionedGenericHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissionedGenericHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PermissionedGenericHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissionedGenericHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PermissionedGenericHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissionedGenericHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PermissionedGenericHandlerSession struct {
	Contract     *PermissionedGenericHandler // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// PermissionedGenericHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PermissionedGenericHandlerCallerSession struct {
	Contract *PermissionedGenericHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// PermissionedGenericHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PermissionedGenericHandlerTransactorSession struct {
	Contract     *PermissionedGenericHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// PermissionedGenericHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type PermissionedGenericHandlerRaw struct {
	Contract *PermissionedGenericHandler // Generic contract binding to access the raw methods on
}

// PermissionedGenericHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PermissionedGenericHandlerCallerRaw struct {
	Contract *PermissionedGenericHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// PermissionedGenericHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PermissionedGenericHandlerTransactorRaw struct {
	Contract *PermissionedGenericHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPermissionedGenericHandler creates a new instance of PermissionedGenericHandler, bound to a specific deployed contract.
func NewPermissionedGenericHandler(address common.Address, backend bind.ContractBackend) (*PermissionedGenericHandler, error) {
	contract, err := bindPermissionedGenericHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PermissionedGenericHandler{PermissionedGenericHandlerCaller: PermissionedGenericHandlerCaller{contract: contract}, PermissionedGenericHandlerTransactor: PermissionedGenericHandlerTransactor{contract: contract}, PermissionedGenericHandlerFilterer: PermissionedGenericHandlerFilterer{contract: contract}}, nil
}

// NewPermissionedGenericHandlerCaller creates a new read-only instance of PermissionedGenericHandler, bound to a specific deployed contract.
func NewPermissionedGenericHandlerCaller(address common.Address, caller bind.ContractCaller) (*PermissionedGenericHandlerCaller, error) {
	contract, err := bindPermissionedGenericHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PermissionedGenericHandlerCaller{contract: contract}, nil
}

// NewPermissionedGenericHandlerTransactor creates a new write-only instance of PermissionedGenericHandler, bound to a specific deployed contract.
func NewPermissionedGenericHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*PermissionedGenericHandlerTransactor, error) {
	contract, err := bindPermissionedGenericHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PermissionedGenericHandlerTransactor{contract: contract}, nil
}

// NewPermissionedGenericHandlerFilterer creates a new log filterer instance of PermissionedGenericHandler, bound to a specific deployed contract.
func NewPermissionedGenericHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*PermissionedGenericHandlerFilterer, error) {
	contract, err := bindPermissionedGenericHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PermissionedGenericHandlerFilterer{contract: contract}, nil
}

// bindPermissionedGenericHandler binds a generic wrapper to an already deployed contract.
func bindPermissionedGenericHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PermissionedGenericHandlerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PermissionedGenericHandler *PermissionedGenericHandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PermissionedGenericHandler.Contract.PermissionedGenericHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PermissionedGenericHandler *PermissionedGenericHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermissionedGenericHandler.Contract.PermissionedGenericHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PermissionedGenericHandler *PermissionedGenericHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PermissionedGenericHandler.Contract.PermissionedGenericHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PermissionedGenericHandler *PermissionedGenericHandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PermissionedGenericHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PermissionedGenericHandler *PermissionedGenericHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermissionedGenericHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PermissionedGenericHandler *PermissionedGenericHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PermissionedGenericHandler.Contract.contract.Transact(opts, method, params...)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_PermissionedGenericHandler *PermissionedGenericHandlerCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PermissionedGenericHandler.contract.Call(opts, &out, "_bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_PermissionedGenericHandler *PermissionedGenericHandlerSession) BridgeAddress() (common.Address, error) {
	return _PermissionedGenericHandler.Contract.BridgeAddress(&_PermissionedGenericHandler.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_PermissionedGenericHandler *PermissionedGenericHandlerCallerSession) BridgeAddress() (common.Address, error) {
	return _PermissionedGenericHandler.Contract.BridgeAddress(&_PermissionedGenericHandler.CallOpts)
}

// ContractAddressToDepositFunctionDepositorOffset is a free data retrieval call binding the contract method 0x8843a70a.
//
// Solidity: function _contractAddressToDepositFunctionDepositorOffset(address ) view returns(uint256)
func (_PermissionedGenericHandler *PermissionedGenericHandlerCaller) ContractAddressToDepositFunctionDepositorOffset(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PermissionedGenericHandler.contract.Call(opts, &out, "_contractAddressToDepositFunctionDepositorOffset", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ContractAddressToDepositFunctionDepositorOffset is a free data retrieval call binding the contract method 0x8843a70a.
//
// Solidity: function _contractAddressToDepositFunctionDepositorOffset(address ) view returns(uint256)
func (_PermissionedGenericHandler *PermissionedGenericHandlerSession) ContractAddressToDepositFunctionDepositorOffset(arg0 common.Address) (*big.Int, error) {
	return _PermissionedGenericHandler.Contract.ContractAddressToDepositFunctionDepositorOffset(&_PermissionedGenericHandler.CallOpts, arg0)
}

// ContractAddressToDepositFunctionDepositorOffset is a free data retrieval call binding the contract method 0x8843a70a.
//
// Solidity: function _contractAddressToDepositFunctionDepositorOffset(address ) view returns(uint256)
func (_PermissionedGenericHandler *PermissionedGenericHandlerCallerSession) ContractAddressToDepositFunctionDepositorOffset(arg0 common.Address) (*big.Int, error) {
	return _PermissionedGenericHandler.Contract.ContractAddressToDepositFunctionDepositorOffset(&_PermissionedGenericHandler.CallOpts, arg0)
}

// ContractAddressToDepositFunctionSignature is a free data retrieval call binding the contract method 0xcb624463.
//
// Solidity: function _contractAddressToDepositFunctionSignature(address ) view returns(bytes4)
func (_PermissionedGenericHandler *PermissionedGenericHandlerCaller) ContractAddressToDepositFunctionSignature(opts *bind.CallOpts, arg0 common.Address) ([4]byte, error) {
	var out []interface{}
	err := _PermissionedGenericHandler.contract.Call(opts, &out, "_contractAddressToDepositFunctionSignature", arg0)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// ContractAddressToDepositFunctionSignature is a free data retrieval call binding the contract method 0xcb624463.
//
// Solidity: function _contractAddressToDepositFunctionSignature(address ) view returns(bytes4)
func (_PermissionedGenericHandler *PermissionedGenericHandlerSession) ContractAddressToDepositFunctionSignature(arg0 common.Address) ([4]byte, error) {
	return _PermissionedGenericHandler.Contract.ContractAddressToDepositFunctionSignature(&_PermissionedGenericHandler.CallOpts, arg0)
}

// ContractAddressToDepositFunctionSignature is a free data retrieval call binding the contract method 0xcb624463.
//
// Solidity: function _contractAddressToDepositFunctionSignature(address ) view returns(bytes4)
func (_PermissionedGenericHandler *PermissionedGenericHandlerCallerSession) ContractAddressToDepositFunctionSignature(arg0 common.Address) ([4]byte, error) {
	return _PermissionedGenericHandler.Contract.ContractAddressToDepositFunctionSignature(&_PermissionedGenericHandler.CallOpts, arg0)
}

// ContractAddressToExecuteFunctionSignature is a free data retrieval call binding the contract method 0xa5c3a985.
//
// Solidity: function _contractAddressToExecuteFunctionSignature(address ) view returns(bytes4)
func (_PermissionedGenericHandler *PermissionedGenericHandlerCaller) ContractAddressToExecuteFunctionSignature(opts *bind.CallOpts, arg0 common.Address) ([4]byte, error) {
	var out []interface{}
	err := _PermissionedGenericHandler.contract.Call(opts, &out, "_contractAddressToExecuteFunctionSignature", arg0)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// ContractAddressToExecuteFunctionSignature is a free data retrieval call binding the contract method 0xa5c3a985.
//
// Solidity: function _contractAddressToExecuteFunctionSignature(address ) view returns(bytes4)
func (_PermissionedGenericHandler *PermissionedGenericHandlerSession) ContractAddressToExecuteFunctionSignature(arg0 common.Address) ([4]byte, error) {
	return _PermissionedGenericHandler.Contract.ContractAddressToExecuteFunctionSignature(&_PermissionedGenericHandler.CallOpts, arg0)
}

// ContractAddressToExecuteFunctionSignature is a free data retrieval call binding the contract method 0xa5c3a985.
//
// Solidity: function _contractAddressToExecuteFunctionSignature(address ) view returns(bytes4)
func (_PermissionedGenericHandler *PermissionedGenericHandlerCallerSession) ContractAddressToExecuteFunctionSignature(arg0 common.Address) ([4]byte, error) {
	return _PermissionedGenericHandler.Contract.ContractAddressToExecuteFunctionSignature(&_PermissionedGenericHandler.CallOpts, arg0)
}

// ContractAddressToResourceID is a free data retrieval call binding the contract method 0xec97d3b4.
//
// Solidity: function _contractAddressToResourceID(address ) view returns(bytes32)
func (_PermissionedGenericHandler *PermissionedGenericHandlerCaller) ContractAddressToResourceID(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _PermissionedGenericHandler.contract.Call(opts, &out, "_contractAddressToResourceID", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ContractAddressToResourceID is a free data retrieval call binding the contract method 0xec97d3b4.
//
// Solidity: function _contractAddressToResourceID(address ) view returns(bytes32)
func (_PermissionedGenericHandler *PermissionedGenericHandlerSession) ContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _PermissionedGenericHandler.Contract.ContractAddressToResourceID(&_PermissionedGenericHandler.CallOpts, arg0)
}

// ContractAddressToResourceID is a free data retrieval call binding the contract method 0xec97d3b4.
//
// Solidity: function _contractAddressToResourceID(address ) view returns(bytes32)
func (_PermissionedGenericHandler *PermissionedGenericHandlerCallerSession) ContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _PermissionedGenericHandler.Contract.ContractAddressToResourceID(&_PermissionedGenericHandler.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_PermissionedGenericHandler *PermissionedGenericHandlerCaller) ContractWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _PermissionedGenericHandler.contract.Call(opts, &out, "_contractWhitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_PermissionedGenericHandler *PermissionedGenericHandlerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _PermissionedGenericHandler.Contract.ContractWhitelist(&_PermissionedGenericHandler.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_PermissionedGenericHandler *PermissionedGenericHandlerCallerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _PermissionedGenericHandler.Contract.ContractWhitelist(&_PermissionedGenericHandler.CallOpts, arg0)
}

// ResourceIDToContractAddress is a free data retrieval call binding the contract method 0xc54c2a11.
//
// Solidity: function _resourceIDToContractAddress(bytes32 ) view returns(address)
func (_PermissionedGenericHandler *PermissionedGenericHandlerCaller) ResourceIDToContractAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _PermissionedGenericHandler.contract.Call(opts, &out, "_resourceIDToContractAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResourceIDToContractAddress is a free data retrieval call binding the contract method 0xc54c2a11.
//
// Solidity: function _resourceIDToContractAddress(bytes32 ) view returns(address)
func (_PermissionedGenericHandler *PermissionedGenericHandlerSession) ResourceIDToContractAddress(arg0 [32]byte) (common.Address, error) {
	return _PermissionedGenericHandler.Contract.ResourceIDToContractAddress(&_PermissionedGenericHandler.CallOpts, arg0)
}

// ResourceIDToContractAddress is a free data retrieval call binding the contract method 0xc54c2a11.
//
// Solidity: function _resourceIDToContractAddress(bytes32 ) view returns(address)
func (_PermissionedGenericHandler *PermissionedGenericHandlerCallerSession) ResourceIDToContractAddress(arg0 [32]byte) (common.Address, error) {
	return _PermissionedGenericHandler.Contract.ResourceIDToContractAddress(&_PermissionedGenericHandler.CallOpts, arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0xb07e54bb.
//
// Solidity: function deposit(bytes32 resourceID, address depositor, bytes data) returns(bytes)
func (_PermissionedGenericHandler *PermissionedGenericHandlerTransactor) Deposit(opts *bind.TransactOpts, resourceID [32]byte, depositor common.Address, data []byte) (*types.Transaction, error) {
	return _PermissionedGenericHandler.contract.Transact(opts, "deposit", resourceID, depositor, data)
}

// Deposit is a paid mutator transaction binding the contract method 0xb07e54bb.
//
// Solidity: function deposit(bytes32 resourceID, address depositor, bytes data) returns(bytes)
func (_PermissionedGenericHandler *PermissionedGenericHandlerSession) Deposit(resourceID [32]byte, depositor common.Address, data []byte) (*types.Transaction, error) {
	return _PermissionedGenericHandler.Contract.Deposit(&_PermissionedGenericHandler.TransactOpts, resourceID, depositor, data)
}

// Deposit is a paid mutator transaction binding the contract method 0xb07e54bb.
//
// Solidity: function deposit(bytes32 resourceID, address depositor, bytes data) returns(bytes)
func (_PermissionedGenericHandler *PermissionedGenericHandlerTransactorSession) Deposit(resourceID [32]byte, depositor common.Address, data []byte) (*types.Transaction, error) {
	return _PermissionedGenericHandler.Contract.Deposit(&_PermissionedGenericHandler.TransactOpts, resourceID, depositor, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_PermissionedGenericHandler *PermissionedGenericHandlerTransactor) ExecuteProposal(opts *bind.TransactOpts, resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _PermissionedGenericHandler.contract.Transact(opts, "executeProposal", resourceID, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_PermissionedGenericHandler *PermissionedGenericHandlerSession) ExecuteProposal(resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _PermissionedGenericHandler.Contract.ExecuteProposal(&_PermissionedGenericHandler.TransactOpts, resourceID, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_PermissionedGenericHandler *PermissionedGenericHandlerTransactorSession) ExecuteProposal(resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _PermissionedGenericHandler.Contract.ExecuteProposal(&_PermissionedGenericHandler.TransactOpts, resourceID, data)
}

// SetResource is a paid mutator transaction binding the contract method 0xfa8675b0.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress, bytes args) returns()
func (_PermissionedGenericHandler *PermissionedGenericHandlerTransactor) SetResource(opts *bind.TransactOpts, resourceID [32]byte, contractAddress common.Address, args []byte) (*types.Transaction, error) {
	return _PermissionedGenericHandler.contract.Transact(opts, "setResource", resourceID, contractAddress, args)
}

// SetResource is a paid mutator transaction binding the contract method 0xfa8675b0.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress, bytes args) returns()
func (_PermissionedGenericHandler *PermissionedGenericHandlerSession) SetResource(resourceID [32]byte, contractAddress common.Address, args []byte) (*types.Transaction, error) {
	return _PermissionedGenericHandler.Contract.SetResource(&_PermissionedGenericHandler.TransactOpts, resourceID, contractAddress, args)
}

// SetResource is a paid mutator transaction binding the contract method 0xfa8675b0.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress, bytes args) returns()
func (_PermissionedGenericHandler *PermissionedGenericHandlerTransactorSession) SetResource(resourceID [32]byte, contractAddress common.Address, args []byte) (*types.Transaction, error) {
	return _PermissionedGenericHandler.Contract.SetResource(&_PermissionedGenericHandler.TransactOpts, resourceID, contractAddress, args)
}

// PermissionedGenericHandlerFailedHandlerExecutionIterator is returned from FilterFailedHandlerExecution and is used to iterate over the raw logs and unpacked data for FailedHandlerExecution events raised by the PermissionedGenericHandler contract.
type PermissionedGenericHandlerFailedHandlerExecutionIterator struct {
	Event *PermissionedGenericHandlerFailedHandlerExecution // Event containing the contract specifics and raw log

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
func (it *PermissionedGenericHandlerFailedHandlerExecutionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionedGenericHandlerFailedHandlerExecution)
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
		it.Event = new(PermissionedGenericHandlerFailedHandlerExecution)
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
func (it *PermissionedGenericHandlerFailedHandlerExecutionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionedGenericHandlerFailedHandlerExecutionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionedGenericHandlerFailedHandlerExecution represents a FailedHandlerExecution event raised by the PermissionedGenericHandler contract.
type PermissionedGenericHandlerFailedHandlerExecution struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterFailedHandlerExecution is a free log retrieval operation binding the contract event 0x3f51c54fb334903765b4de0daee60de6eff2a87aef23d5a1f311f4ea2e1b7e96.
//
// Solidity: event FailedHandlerExecution()
func (_PermissionedGenericHandler *PermissionedGenericHandlerFilterer) FilterFailedHandlerExecution(opts *bind.FilterOpts) (*PermissionedGenericHandlerFailedHandlerExecutionIterator, error) {

	logs, sub, err := _PermissionedGenericHandler.contract.FilterLogs(opts, "FailedHandlerExecution")
	if err != nil {
		return nil, err
	}
	return &PermissionedGenericHandlerFailedHandlerExecutionIterator{contract: _PermissionedGenericHandler.contract, event: "FailedHandlerExecution", logs: logs, sub: sub}, nil
}

// WatchFailedHandlerExecution is a free log subscription operation binding the contract event 0x3f51c54fb334903765b4de0daee60de6eff2a87aef23d5a1f311f4ea2e1b7e96.
//
// Solidity: event FailedHandlerExecution()
func (_PermissionedGenericHandler *PermissionedGenericHandlerFilterer) WatchFailedHandlerExecution(opts *bind.WatchOpts, sink chan<- *PermissionedGenericHandlerFailedHandlerExecution) (event.Subscription, error) {

	logs, sub, err := _PermissionedGenericHandler.contract.WatchLogs(opts, "FailedHandlerExecution")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionedGenericHandlerFailedHandlerExecution)
				if err := _PermissionedGenericHandler.contract.UnpackLog(event, "FailedHandlerExecution", log); err != nil {
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

// ParseFailedHandlerExecution is a log parse operation binding the contract event 0x3f51c54fb334903765b4de0daee60de6eff2a87aef23d5a1f311f4ea2e1b7e96.
//
// Solidity: event FailedHandlerExecution()
func (_PermissionedGenericHandler *PermissionedGenericHandlerFilterer) ParseFailedHandlerExecution(log types.Log) (*PermissionedGenericHandlerFailedHandlerExecution, error) {
	event := new(PermissionedGenericHandlerFailedHandlerExecution)
	if err := _PermissionedGenericHandler.contract.UnpackLog(event, "FailedHandlerExecution", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
