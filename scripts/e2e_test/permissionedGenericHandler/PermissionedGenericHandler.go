// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PermissionedGenericHandler

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PermissionedGenericHandlerABI is the input ABI used to generate the binding from.
const PermissionedGenericHandlerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"FailedHandlerExecution\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_tokenContractAddressToTokenProperties\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes4\",\"name\":\"depositFunctionSignature\",\"type\":\"bytes4\"},{\"internalType\":\"uint16\",\"name\":\"depositFunctionDepositorOffset\",\"type\":\"uint16\"},{\"internalType\":\"bytes4\",\"name\":\"executeFunctionSignature\",\"type\":\"bytes4\"},{\"internalType\":\"bool\",\"name\":\"isWhitelisted\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"args\",\"type\":\"bytes\"}],\"name\":\"setResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PermissionedGenericHandlerBin is the compiled bytecode used for deploying new contracts.
var PermissionedGenericHandlerBin = "0x60a06040523480156200001157600080fd5b50604051620015ac380380620015ac8339818101604052810190620000379190620000dc565b8073ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1681525050506200010e565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620000a48262000077565b9050919050565b620000b68162000097565b8114620000c257600080fd5b50565b600081519050620000d681620000ab565b92915050565b600060208284031215620000f557620000f462000072565b5b60006200010584828501620000c5565b91505092915050565b60805161147b6200013160003960008181610153015261090b015261147b6000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c8063318c136e14610067578063ac607c2114610085578063b07e54bb146100b9578063c54c2a11146100e9578063e248cff214610119578063fa8675b014610135575b600080fd5b61006f610151565b60405161007c9190610be9565b60405180910390f35b61009f600480360381019061009a9190610c3a565b610175565b6040516100b0959493929190610cf3565b60405180910390f35b6100d360048036038101906100ce9190610dd7565b6101e0565b6040516100e09190610ee4565b60405180910390f35b61010360048036038101906100fe9190610f06565b610598565b6040516101109190610be9565b60405180910390f35b610133600480360381019061012e9190610f33565b6105cb565b005b61014f600480360381019061014a9190610dd7565b61087f565b005b7f000000000000000000000000000000000000000000000000000000000000000081565b60006020528060005260406000206000915090508060000154908060010160009054906101000a900460e01b908060010160049054906101000a900461ffff16908060010160069054906101000a900460e01b9080600101600a9054906101000a900460ff16905085565b60606101ea610909565b6000606084848101906101fd9190610fc9565b915084846020908460206102119190611025565b9261021e93929190611085565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050905060006001600089815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160049054906101000a900461ffff16905060008161ffff16111561037b576000816020850101519050606081901c73ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff1614610379576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103709061111d565b60405180910390fd5b505b6000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101600a9054906101000a900460ff16610409576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610400906111af565b60405180910390fd5b60008060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460e01b9050600060e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161461058a57600081856040516020016104ba92919061122c565b60405160208183030381529060405290506000808573ffffffffffffffffffffffffffffffffffffffff16836040516104f39190611254565b6000604051808303816000865af19150503d8060008114610530576040519150601f19603f3d011682016040523d82523d6000602084013e610535565b606091505b50915091508161057a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610571906112b7565b60405180910390fd5b8098505050505050505050610590565b50505050505b949350505050565b60016020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6105d3610909565b6000606083838101906105e69190610fc9565b915083836020908460206105fa9190611025565b9261060793929190611085565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050905060006001600087815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506000808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101600a9054906101000a900460ff16610712576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610709906111af565b60405180910390fd5b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160069054906101000a900460e01b9050600060e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161461087657600081846040516020016107c392919061122c565b604051602081830303815290604052905060008373ffffffffffffffffffffffffffffffffffffffff16826040516107fb9190611254565b6000604051808303816000865af19150503d8060008114610838576040519150601f19603f3d011682016040523d82523d6000602084013e61083d565b606091505b5050905080610873577f3f51c54fb334903765b4de0daee60de6eff2a87aef23d5a1f311f4ea2e1b7e9660405160405180910390a15b50505b50505050505050565b610887610909565b6000828260009060049261089d93929190611085565b906108a891906112ef565b9050600083836004906006926108c093929190611085565b906108cb919061137a565b60f01c905060008484600690600a926108e693929190611085565b906108f191906112ef565b90506109008787858585610999565b50505050505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610997576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161098e90611425565b60405180910390fd5b565b836001600087815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550846000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000181905550826000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160006101000a81548163ffffffff021916908360e01c0217905550816000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160046101000a81548161ffff021916908361ffff160217905550806000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160066101000a81548163ffffffff021916908360e01c021790555060016000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101600a6101000a81548160ff0219169083151502179055505050505050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610bd382610ba8565b9050919050565b610be381610bc8565b82525050565b6000602082019050610bfe6000830184610bda565b92915050565b600080fd5b600080fd5b610c1781610bc8565b8114610c2257600080fd5b50565b600081359050610c3481610c0e565b92915050565b600060208284031215610c5057610c4f610c04565b5b6000610c5e84828501610c25565b91505092915050565b6000819050919050565b610c7a81610c67565b82525050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b610cb581610c80565b82525050565b600061ffff82169050919050565b610cd281610cbb565b82525050565b60008115159050919050565b610ced81610cd8565b82525050565b600060a082019050610d086000830188610c71565b610d156020830187610cac565b610d226040830186610cc9565b610d2f6060830185610cac565b610d3c6080830184610ce4565b9695505050505050565b610d4f81610c67565b8114610d5a57600080fd5b50565b600081359050610d6c81610d46565b92915050565b600080fd5b600080fd5b600080fd5b60008083601f840112610d9757610d96610d72565b5b8235905067ffffffffffffffff811115610db457610db3610d77565b5b602083019150836001820283011115610dd057610dcf610d7c565b5b9250929050565b60008060008060608587031215610df157610df0610c04565b5b6000610dff87828801610d5d565b9450506020610e1087828801610c25565b935050604085013567ffffffffffffffff811115610e3157610e30610c09565b5b610e3d87828801610d81565b925092505092959194509250565b600081519050919050565b600082825260208201905092915050565b60005b83811015610e85578082015181840152602081019050610e6a565b83811115610e94576000848401525b50505050565b6000601f19601f8301169050919050565b6000610eb682610e4b565b610ec08185610e56565b9350610ed0818560208601610e67565b610ed981610e9a565b840191505092915050565b60006020820190508181036000830152610efe8184610eab565b905092915050565b600060208284031215610f1c57610f1b610c04565b5b6000610f2a84828501610d5d565b91505092915050565b600080600060408486031215610f4c57610f4b610c04565b5b6000610f5a86828701610d5d565b935050602084013567ffffffffffffffff811115610f7b57610f7a610c09565b5b610f8786828701610d81565b92509250509250925092565b6000819050919050565b610fa681610f93565b8114610fb157600080fd5b50565b600081359050610fc381610f9d565b92915050565b600060208284031215610fdf57610fde610c04565b5b6000610fed84828501610fb4565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061103082610f93565b915061103b83610f93565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156110705761106f610ff6565b5b828201905092915050565b600080fd5b600080fd5b600080858511156110995761109861107b565b5b838611156110aa576110a9611080565b5b6001850283019150848603905094509492505050565b600082825260208201905092915050565b7f696e636f7272656374206465706f7369746f7220696e20746865206461746100600082015250565b6000611107601f836110c0565b9150611112826110d1565b602082019050919050565b60006020820190508181036000830152611136816110fa565b9050919050565b7f70726f766964656420636f6e747261637441646472657373206973206e6f742060008201527f77686974656c6973746564000000000000000000000000000000000000000000602082015250565b6000611199602b836110c0565b91506111a48261113d565b604082019050919050565b600060208201905081810360008301526111c88161118c565b9050919050565b6000819050919050565b6111ea6111e582610c80565b6111cf565b82525050565b600081905092915050565b600061120682610e4b565b61121081856111f0565b9350611220818560208601610e67565b80840191505092915050565b600061123882856111d9565b60048201915061124882846111fb565b91508190509392505050565b600061126082846111fb565b915081905092915050565b7f63616c6c20746f20636f6e747261637441646472657373206661696c65640000600082015250565b60006112a1601e836110c0565b91506112ac8261126b565b602082019050919050565b600060208201905081810360008301526112d081611294565b9050919050565b600082905092915050565b600082821b905092915050565b60006112fb83836112d7565b826113068135610c80565b92506004821015611346576113417fffffffff00000000000000000000000000000000000000000000000000000000836004036008026112e2565b831692505b505092915050565b60007fffff00000000000000000000000000000000000000000000000000000000000082169050919050565b600061138683836112d7565b82611391813561134e565b925060028210156113d1576113cc7fffff000000000000000000000000000000000000000000000000000000000000836002036008026112e2565b831692505b505092915050565b7f73656e646572206d7573742062652062726964676520636f6e74726163740000600082015250565b600061140f601e836110c0565b915061141a826113d9565b602082019050919050565b6000602082019050818103600083015261143e81611402565b905091905056fea2646970667358221220b301bdd22fb7201ec622223d0d6a6c4496e59a91e270846d02c15ea3c54ea0c564736f6c634300080b0033"

// DeployPermissionedGenericHandler deploys a new Ethereum contract, binding an instance of PermissionedGenericHandler to it.
func DeployPermissionedGenericHandler(auth *bind.TransactOpts, backend bind.ContractBackend, bridgeAddress common.Address) (common.Address, *types.Transaction, *PermissionedGenericHandler, error) {
	parsed, err := abi.JSON(strings.NewReader(PermissionedGenericHandlerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PermissionedGenericHandlerBin), backend, bridgeAddress)
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

// TokenContractAddressToTokenProperties is a free data retrieval call binding the contract method 0xac607c21.
//
// Solidity: function _tokenContractAddressToTokenProperties(address ) view returns(bytes32 resourceID, bytes4 depositFunctionSignature, uint16 depositFunctionDepositorOffset, bytes4 executeFunctionSignature, bool isWhitelisted)
func (_PermissionedGenericHandler *PermissionedGenericHandlerCaller) TokenContractAddressToTokenProperties(opts *bind.CallOpts, arg0 common.Address) (struct {
	ResourceID                     [32]byte
	DepositFunctionSignature       [4]byte
	DepositFunctionDepositorOffset uint16
	ExecuteFunctionSignature       [4]byte
	IsWhitelisted                  bool
}, error) {
	var out []interface{}
	err := _PermissionedGenericHandler.contract.Call(opts, &out, "_tokenContractAddressToTokenProperties", arg0)

	outstruct := new(struct {
		ResourceID                     [32]byte
		DepositFunctionSignature       [4]byte
		DepositFunctionDepositorOffset uint16
		ExecuteFunctionSignature       [4]byte
		IsWhitelisted                  bool
	})

	outstruct.ResourceID = out[0].([32]byte)
	outstruct.DepositFunctionSignature = out[1].([4]byte)
	outstruct.DepositFunctionDepositorOffset = out[2].(uint16)
	outstruct.ExecuteFunctionSignature = out[3].([4]byte)
	outstruct.IsWhitelisted = out[4].(bool)

	return *outstruct, err

}

// TokenContractAddressToTokenProperties is a free data retrieval call binding the contract method 0xac607c21.
//
// Solidity: function _tokenContractAddressToTokenProperties(address ) view returns(bytes32 resourceID, bytes4 depositFunctionSignature, uint16 depositFunctionDepositorOffset, bytes4 executeFunctionSignature, bool isWhitelisted)
func (_PermissionedGenericHandler *PermissionedGenericHandlerSession) TokenContractAddressToTokenProperties(arg0 common.Address) (struct {
	ResourceID                     [32]byte
	DepositFunctionSignature       [4]byte
	DepositFunctionDepositorOffset uint16
	ExecuteFunctionSignature       [4]byte
	IsWhitelisted                  bool
}, error) {
	return _PermissionedGenericHandler.Contract.TokenContractAddressToTokenProperties(&_PermissionedGenericHandler.CallOpts, arg0)
}

// TokenContractAddressToTokenProperties is a free data retrieval call binding the contract method 0xac607c21.
//
// Solidity: function _tokenContractAddressToTokenProperties(address ) view returns(bytes32 resourceID, bytes4 depositFunctionSignature, uint16 depositFunctionDepositorOffset, bytes4 executeFunctionSignature, bool isWhitelisted)
func (_PermissionedGenericHandler *PermissionedGenericHandlerCallerSession) TokenContractAddressToTokenProperties(arg0 common.Address) (struct {
	ResourceID                     [32]byte
	DepositFunctionSignature       [4]byte
	DepositFunctionDepositorOffset uint16
	ExecuteFunctionSignature       [4]byte
	IsWhitelisted                  bool
}, error) {
	return _PermissionedGenericHandler.Contract.TokenContractAddressToTokenProperties(&_PermissionedGenericHandler.CallOpts, arg0)
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
