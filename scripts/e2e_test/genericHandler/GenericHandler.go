// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package GenericHandler

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

// GenericHandlerABI is the input ABI used to generate the binding from.
const GenericHandlerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"FailedHandlerExecution\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractAddressToDepositFunctionDepositerOffset\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractAddressToDepositFunctionSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractAddressToExecuteFunctionSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractAddressToResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"depositFunctionSig\",\"type\":\"bytes4\"},{\"internalType\":\"uint256\",\"name\":\"depositFunctionDepositerOffset\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"executeFunctionSig\",\"type\":\"bytes4\"}],\"name\":\"setResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GenericHandlerBin is the compiled bytecode used for deploying new contracts.
var GenericHandlerBin = "0x60a06040523480156200001157600080fd5b50604051620015cc380380620015cc8339818101604052810190620000379190620000dc565b8073ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1681525050506200010e565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620000a48262000077565b9050919050565b620000b68162000097565b8114620000c257600080fd5b50565b600081519050620000d681620000ab565b92915050565b600060208284031215620000f557620000f462000072565b5b60006200010584828501620000c5565b91505092915050565b60805161149b620001316000396000818161024b015261099c015261149b6000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c8063c54c2a1111610066578063c54c2a1114610181578063cb624463146101b1578063de319d99146101e1578063e248cff2146101fd578063ec97d3b4146102195761009e565b8063318c136e146100a35780637f79bea8146100c1578063a5c3a985146100f1578063aa50800b14610121578063b07e54bb14610151575b600080fd5b6100ab610249565b6040516100b89190610c59565b60405180910390f35b6100db60048036038101906100d69190610caa565b61026d565b6040516100e89190610cf2565b60405180910390f35b61010b60048036038101906101069190610caa565b61028d565b6040516101189190610d48565b60405180910390f35b61013b60048036038101906101369190610caa565b6102ad565b6040516101489190610d7c565b60405180910390f35b61016b60048036038101906101669190610e32565b6102c5565b6040516101789190610f3f565b60405180910390f35b61019b60048036038101906101969190610f61565b610664565b6040516101a89190610c59565b60405180910390f35b6101cb60048036038101906101c69190610caa565b610697565b6040516101d89190610d48565b60405180910390f35b6101fb60048036038101906101f69190610fe6565b6106b7565b005b61021760048036038101906102129190611061565b6106d3565b005b610233600480360381019061022e9190610caa565b610982565b60405161024091906110d0565b60405180910390f35b7f000000000000000000000000000000000000000000000000000000000000000081565b60056020528060005260406000206000915054906101000a900460ff1681565b60046020528060005260406000206000915054906101000a900460e01b81565b60036020528060005260406000206000915090505481565b60606102cf61099a565b6000606084848101906102e291906110eb565b915084846020908460206102f69190611147565b92610303939291906111a7565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509050600080600089815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506000600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050600081111561044b576000816020850101519050606081901c73ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff1614610449576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104409061123f565b60405180910390fd5b505b600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166104d7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104ce906112d1565b60405180910390fd5b6000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460e01b9050600060e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614610656576000818560405160200161058692919061134e565b60405160208183030381529060405290506000808573ffffffffffffffffffffffffffffffffffffffff16836040516105bf9190611376565b6000604051808303816000865af19150503d80600081146105fc576040519150601f19603f3d011682016040523d82523d6000602084013e610601565b606091505b509150915081610646576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161063d906113d9565b60405180910390fd5b809850505050505050505061065c565b50505050505b949350505050565b60006020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60026020528060005260406000206000915054906101000a900460e01b81565b6106bf61099a565b6106cc8585858585610a2a565b5050505050565b6106db61099a565b6000606083838101906106ee91906110eb565b915083836020908460206107029190611147565b9261070f939291906111a7565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509050600080600087815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600560008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610817576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161080e906112d1565b60405180910390fd5b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460e01b9050600060e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161461097957600081846040516020016108c692919061134e565b604051602081830303815290604052905060008373ffffffffffffffffffffffffffffffffffffffff16826040516108fe9190611376565b6000604051808303816000865af19150503d806000811461093b576040519150601f19603f3d011682016040523d82523d6000602084013e610940565b606091505b5050905080610976577f3f51c54fb334903765b4de0daee60de6eff2a87aef23d5a1f311f4ea2e1b7e9660405160405180910390a15b50505b50505050505050565b60016020528060005260406000206000915090505481565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a28576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a1f90611445565b60405180910390fd5b565b8360008087815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555084600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555082600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908360e01c021790555081600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555080600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908360e01c02179055506001600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050505050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610c4382610c18565b9050919050565b610c5381610c38565b82525050565b6000602082019050610c6e6000830184610c4a565b92915050565b600080fd5b600080fd5b610c8781610c38565b8114610c9257600080fd5b50565b600081359050610ca481610c7e565b92915050565b600060208284031215610cc057610cbf610c74565b5b6000610cce84828501610c95565b91505092915050565b60008115159050919050565b610cec81610cd7565b82525050565b6000602082019050610d076000830184610ce3565b92915050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b610d4281610d0d565b82525050565b6000602082019050610d5d6000830184610d39565b92915050565b6000819050919050565b610d7681610d63565b82525050565b6000602082019050610d916000830184610d6d565b92915050565b6000819050919050565b610daa81610d97565b8114610db557600080fd5b50565b600081359050610dc781610da1565b92915050565b600080fd5b600080fd5b600080fd5b60008083601f840112610df257610df1610dcd565b5b8235905067ffffffffffffffff811115610e0f57610e0e610dd2565b5b602083019150836001820283011115610e2b57610e2a610dd7565b5b9250929050565b60008060008060608587031215610e4c57610e4b610c74565b5b6000610e5a87828801610db8565b9450506020610e6b87828801610c95565b935050604085013567ffffffffffffffff811115610e8c57610e8b610c79565b5b610e9887828801610ddc565b925092505092959194509250565b600081519050919050565b600082825260208201905092915050565b60005b83811015610ee0578082015181840152602081019050610ec5565b83811115610eef576000848401525b50505050565b6000601f19601f8301169050919050565b6000610f1182610ea6565b610f1b8185610eb1565b9350610f2b818560208601610ec2565b610f3481610ef5565b840191505092915050565b60006020820190508181036000830152610f598184610f06565b905092915050565b600060208284031215610f7757610f76610c74565b5b6000610f8584828501610db8565b91505092915050565b610f9781610d0d565b8114610fa257600080fd5b50565b600081359050610fb481610f8e565b92915050565b610fc381610d63565b8114610fce57600080fd5b50565b600081359050610fe081610fba565b92915050565b600080600080600060a0868803121561100257611001610c74565b5b600061101088828901610db8565b955050602061102188828901610c95565b945050604061103288828901610fa5565b935050606061104388828901610fd1565b925050608061105488828901610fa5565b9150509295509295909350565b60008060006040848603121561107a57611079610c74565b5b600061108886828701610db8565b935050602084013567ffffffffffffffff8111156110a9576110a8610c79565b5b6110b586828701610ddc565b92509250509250925092565b6110ca81610d97565b82525050565b60006020820190506110e560008301846110c1565b92915050565b60006020828403121561110157611100610c74565b5b600061110f84828501610fd1565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061115282610d63565b915061115d83610d63565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561119257611191611118565b5b828201905092915050565b600080fd5b600080fd5b600080858511156111bb576111ba61119d565b5b838611156111cc576111cb6111a2565b5b6001850283019150848603905094509492505050565b600082825260208201905092915050565b7f696e636f7272656374206465706f736974657220696e20746865206461746100600082015250565b6000611229601f836111e2565b9150611234826111f3565b602082019050919050565b600060208201905081810360008301526112588161121c565b9050919050565b7f70726f766964656420636f6e747261637441646472657373206973206e6f742060008201527f77686974656c6973746564000000000000000000000000000000000000000000602082015250565b60006112bb602b836111e2565b91506112c68261125f565b604082019050919050565b600060208201905081810360008301526112ea816112ae565b9050919050565b6000819050919050565b61130c61130782610d0d565b6112f1565b82525050565b600081905092915050565b600061132882610ea6565b6113328185611312565b9350611342818560208601610ec2565b80840191505092915050565b600061135a82856112fb565b60048201915061136a828461131d565b91508190509392505050565b6000611382828461131d565b915081905092915050565b7f63616c6c20746f20636f6e747261637441646472657373206661696c65640000600082015250565b60006113c3601e836111e2565b91506113ce8261138d565b602082019050919050565b600060208201905081810360008301526113f2816113b6565b9050919050565b7f73656e646572206d7573742062652062726964676520636f6e74726163740000600082015250565b600061142f601e836111e2565b915061143a826113f9565b602082019050919050565b6000602082019050818103600083015261145e81611422565b905091905056fea26469706673582212209d23cf10e74bed9280aed08c0119c1d6652dde02ddeb15b1ecd23bc765f5d92364736f6c634300080b0033"

// DeployGenericHandler deploys a new Ethereum contract, binding an instance of GenericHandler to it.
func DeployGenericHandler(auth *bind.TransactOpts, backend bind.ContractBackend, bridgeAddress common.Address) (common.Address, *types.Transaction, *GenericHandler, error) {
	parsed, err := abi.JSON(strings.NewReader(GenericHandlerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GenericHandlerBin), backend, bridgeAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GenericHandler{GenericHandlerCaller: GenericHandlerCaller{contract: contract}, GenericHandlerTransactor: GenericHandlerTransactor{contract: contract}, GenericHandlerFilterer: GenericHandlerFilterer{contract: contract}}, nil
}

// GenericHandler is an auto generated Go binding around an Ethereum contract.
type GenericHandler struct {
	GenericHandlerCaller     // Read-only binding to the contract
	GenericHandlerTransactor // Write-only binding to the contract
	GenericHandlerFilterer   // Log filterer for contract events
}

// GenericHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type GenericHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenericHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GenericHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenericHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GenericHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenericHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GenericHandlerSession struct {
	Contract     *GenericHandler   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GenericHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GenericHandlerCallerSession struct {
	Contract *GenericHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// GenericHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GenericHandlerTransactorSession struct {
	Contract     *GenericHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// GenericHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type GenericHandlerRaw struct {
	Contract *GenericHandler // Generic contract binding to access the raw methods on
}

// GenericHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GenericHandlerCallerRaw struct {
	Contract *GenericHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// GenericHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GenericHandlerTransactorRaw struct {
	Contract *GenericHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGenericHandler creates a new instance of GenericHandler, bound to a specific deployed contract.
func NewGenericHandler(address common.Address, backend bind.ContractBackend) (*GenericHandler, error) {
	contract, err := bindGenericHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GenericHandler{GenericHandlerCaller: GenericHandlerCaller{contract: contract}, GenericHandlerTransactor: GenericHandlerTransactor{contract: contract}, GenericHandlerFilterer: GenericHandlerFilterer{contract: contract}}, nil
}

// NewGenericHandlerCaller creates a new read-only instance of GenericHandler, bound to a specific deployed contract.
func NewGenericHandlerCaller(address common.Address, caller bind.ContractCaller) (*GenericHandlerCaller, error) {
	contract, err := bindGenericHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GenericHandlerCaller{contract: contract}, nil
}

// NewGenericHandlerTransactor creates a new write-only instance of GenericHandler, bound to a specific deployed contract.
func NewGenericHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*GenericHandlerTransactor, error) {
	contract, err := bindGenericHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GenericHandlerTransactor{contract: contract}, nil
}

// NewGenericHandlerFilterer creates a new log filterer instance of GenericHandler, bound to a specific deployed contract.
func NewGenericHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*GenericHandlerFilterer, error) {
	contract, err := bindGenericHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GenericHandlerFilterer{contract: contract}, nil
}

// bindGenericHandler binds a generic wrapper to an already deployed contract.
func bindGenericHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GenericHandlerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GenericHandler *GenericHandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GenericHandler.Contract.GenericHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GenericHandler *GenericHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GenericHandler.Contract.GenericHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GenericHandler *GenericHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GenericHandler.Contract.GenericHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GenericHandler *GenericHandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GenericHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GenericHandler *GenericHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GenericHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GenericHandler *GenericHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GenericHandler.Contract.contract.Transact(opts, method, params...)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_GenericHandler *GenericHandlerCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GenericHandler.contract.Call(opts, &out, "_bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_GenericHandler *GenericHandlerSession) BridgeAddress() (common.Address, error) {
	return _GenericHandler.Contract.BridgeAddress(&_GenericHandler.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_GenericHandler *GenericHandlerCallerSession) BridgeAddress() (common.Address, error) {
	return _GenericHandler.Contract.BridgeAddress(&_GenericHandler.CallOpts)
}

// ContractAddressToDepositFunctionDepositerOffset is a free data retrieval call binding the contract method 0xaa50800b.
//
// Solidity: function _contractAddressToDepositFunctionDepositerOffset(address ) view returns(uint256)
func (_GenericHandler *GenericHandlerCaller) ContractAddressToDepositFunctionDepositerOffset(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GenericHandler.contract.Call(opts, &out, "_contractAddressToDepositFunctionDepositerOffset", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ContractAddressToDepositFunctionDepositerOffset is a free data retrieval call binding the contract method 0xaa50800b.
//
// Solidity: function _contractAddressToDepositFunctionDepositerOffset(address ) view returns(uint256)
func (_GenericHandler *GenericHandlerSession) ContractAddressToDepositFunctionDepositerOffset(arg0 common.Address) (*big.Int, error) {
	return _GenericHandler.Contract.ContractAddressToDepositFunctionDepositerOffset(&_GenericHandler.CallOpts, arg0)
}

// ContractAddressToDepositFunctionDepositerOffset is a free data retrieval call binding the contract method 0xaa50800b.
//
// Solidity: function _contractAddressToDepositFunctionDepositerOffset(address ) view returns(uint256)
func (_GenericHandler *GenericHandlerCallerSession) ContractAddressToDepositFunctionDepositerOffset(arg0 common.Address) (*big.Int, error) {
	return _GenericHandler.Contract.ContractAddressToDepositFunctionDepositerOffset(&_GenericHandler.CallOpts, arg0)
}

// ContractAddressToDepositFunctionSignature is a free data retrieval call binding the contract method 0xcb624463.
//
// Solidity: function _contractAddressToDepositFunctionSignature(address ) view returns(bytes4)
func (_GenericHandler *GenericHandlerCaller) ContractAddressToDepositFunctionSignature(opts *bind.CallOpts, arg0 common.Address) ([4]byte, error) {
	var out []interface{}
	err := _GenericHandler.contract.Call(opts, &out, "_contractAddressToDepositFunctionSignature", arg0)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// ContractAddressToDepositFunctionSignature is a free data retrieval call binding the contract method 0xcb624463.
//
// Solidity: function _contractAddressToDepositFunctionSignature(address ) view returns(bytes4)
func (_GenericHandler *GenericHandlerSession) ContractAddressToDepositFunctionSignature(arg0 common.Address) ([4]byte, error) {
	return _GenericHandler.Contract.ContractAddressToDepositFunctionSignature(&_GenericHandler.CallOpts, arg0)
}

// ContractAddressToDepositFunctionSignature is a free data retrieval call binding the contract method 0xcb624463.
//
// Solidity: function _contractAddressToDepositFunctionSignature(address ) view returns(bytes4)
func (_GenericHandler *GenericHandlerCallerSession) ContractAddressToDepositFunctionSignature(arg0 common.Address) ([4]byte, error) {
	return _GenericHandler.Contract.ContractAddressToDepositFunctionSignature(&_GenericHandler.CallOpts, arg0)
}

// ContractAddressToExecuteFunctionSignature is a free data retrieval call binding the contract method 0xa5c3a985.
//
// Solidity: function _contractAddressToExecuteFunctionSignature(address ) view returns(bytes4)
func (_GenericHandler *GenericHandlerCaller) ContractAddressToExecuteFunctionSignature(opts *bind.CallOpts, arg0 common.Address) ([4]byte, error) {
	var out []interface{}
	err := _GenericHandler.contract.Call(opts, &out, "_contractAddressToExecuteFunctionSignature", arg0)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// ContractAddressToExecuteFunctionSignature is a free data retrieval call binding the contract method 0xa5c3a985.
//
// Solidity: function _contractAddressToExecuteFunctionSignature(address ) view returns(bytes4)
func (_GenericHandler *GenericHandlerSession) ContractAddressToExecuteFunctionSignature(arg0 common.Address) ([4]byte, error) {
	return _GenericHandler.Contract.ContractAddressToExecuteFunctionSignature(&_GenericHandler.CallOpts, arg0)
}

// ContractAddressToExecuteFunctionSignature is a free data retrieval call binding the contract method 0xa5c3a985.
//
// Solidity: function _contractAddressToExecuteFunctionSignature(address ) view returns(bytes4)
func (_GenericHandler *GenericHandlerCallerSession) ContractAddressToExecuteFunctionSignature(arg0 common.Address) ([4]byte, error) {
	return _GenericHandler.Contract.ContractAddressToExecuteFunctionSignature(&_GenericHandler.CallOpts, arg0)
}

// ContractAddressToResourceID is a free data retrieval call binding the contract method 0xec97d3b4.
//
// Solidity: function _contractAddressToResourceID(address ) view returns(bytes32)
func (_GenericHandler *GenericHandlerCaller) ContractAddressToResourceID(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _GenericHandler.contract.Call(opts, &out, "_contractAddressToResourceID", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ContractAddressToResourceID is a free data retrieval call binding the contract method 0xec97d3b4.
//
// Solidity: function _contractAddressToResourceID(address ) view returns(bytes32)
func (_GenericHandler *GenericHandlerSession) ContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _GenericHandler.Contract.ContractAddressToResourceID(&_GenericHandler.CallOpts, arg0)
}

// ContractAddressToResourceID is a free data retrieval call binding the contract method 0xec97d3b4.
//
// Solidity: function _contractAddressToResourceID(address ) view returns(bytes32)
func (_GenericHandler *GenericHandlerCallerSession) ContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _GenericHandler.Contract.ContractAddressToResourceID(&_GenericHandler.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_GenericHandler *GenericHandlerCaller) ContractWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _GenericHandler.contract.Call(opts, &out, "_contractWhitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_GenericHandler *GenericHandlerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _GenericHandler.Contract.ContractWhitelist(&_GenericHandler.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_GenericHandler *GenericHandlerCallerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _GenericHandler.Contract.ContractWhitelist(&_GenericHandler.CallOpts, arg0)
}

// ResourceIDToContractAddress is a free data retrieval call binding the contract method 0xc54c2a11.
//
// Solidity: function _resourceIDToContractAddress(bytes32 ) view returns(address)
func (_GenericHandler *GenericHandlerCaller) ResourceIDToContractAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _GenericHandler.contract.Call(opts, &out, "_resourceIDToContractAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResourceIDToContractAddress is a free data retrieval call binding the contract method 0xc54c2a11.
//
// Solidity: function _resourceIDToContractAddress(bytes32 ) view returns(address)
func (_GenericHandler *GenericHandlerSession) ResourceIDToContractAddress(arg0 [32]byte) (common.Address, error) {
	return _GenericHandler.Contract.ResourceIDToContractAddress(&_GenericHandler.CallOpts, arg0)
}

// ResourceIDToContractAddress is a free data retrieval call binding the contract method 0xc54c2a11.
//
// Solidity: function _resourceIDToContractAddress(bytes32 ) view returns(address)
func (_GenericHandler *GenericHandlerCallerSession) ResourceIDToContractAddress(arg0 [32]byte) (common.Address, error) {
	return _GenericHandler.Contract.ResourceIDToContractAddress(&_GenericHandler.CallOpts, arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0xb07e54bb.
//
// Solidity: function deposit(bytes32 resourceID, address depositer, bytes data) returns(bytes)
func (_GenericHandler *GenericHandlerTransactor) Deposit(opts *bind.TransactOpts, resourceID [32]byte, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _GenericHandler.contract.Transact(opts, "deposit", resourceID, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0xb07e54bb.
//
// Solidity: function deposit(bytes32 resourceID, address depositer, bytes data) returns(bytes)
func (_GenericHandler *GenericHandlerSession) Deposit(resourceID [32]byte, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _GenericHandler.Contract.Deposit(&_GenericHandler.TransactOpts, resourceID, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0xb07e54bb.
//
// Solidity: function deposit(bytes32 resourceID, address depositer, bytes data) returns(bytes)
func (_GenericHandler *GenericHandlerTransactorSession) Deposit(resourceID [32]byte, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _GenericHandler.Contract.Deposit(&_GenericHandler.TransactOpts, resourceID, depositer, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_GenericHandler *GenericHandlerTransactor) ExecuteProposal(opts *bind.TransactOpts, resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _GenericHandler.contract.Transact(opts, "executeProposal", resourceID, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_GenericHandler *GenericHandlerSession) ExecuteProposal(resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _GenericHandler.Contract.ExecuteProposal(&_GenericHandler.TransactOpts, resourceID, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_GenericHandler *GenericHandlerTransactorSession) ExecuteProposal(resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _GenericHandler.Contract.ExecuteProposal(&_GenericHandler.TransactOpts, resourceID, data)
}

// SetResource is a paid mutator transaction binding the contract method 0xde319d99.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress, bytes4 depositFunctionSig, uint256 depositFunctionDepositerOffset, bytes4 executeFunctionSig) returns()
func (_GenericHandler *GenericHandlerTransactor) SetResource(opts *bind.TransactOpts, resourceID [32]byte, contractAddress common.Address, depositFunctionSig [4]byte, depositFunctionDepositerOffset *big.Int, executeFunctionSig [4]byte) (*types.Transaction, error) {
	return _GenericHandler.contract.Transact(opts, "setResource", resourceID, contractAddress, depositFunctionSig, depositFunctionDepositerOffset, executeFunctionSig)
}

// SetResource is a paid mutator transaction binding the contract method 0xde319d99.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress, bytes4 depositFunctionSig, uint256 depositFunctionDepositerOffset, bytes4 executeFunctionSig) returns()
func (_GenericHandler *GenericHandlerSession) SetResource(resourceID [32]byte, contractAddress common.Address, depositFunctionSig [4]byte, depositFunctionDepositerOffset *big.Int, executeFunctionSig [4]byte) (*types.Transaction, error) {
	return _GenericHandler.Contract.SetResource(&_GenericHandler.TransactOpts, resourceID, contractAddress, depositFunctionSig, depositFunctionDepositerOffset, executeFunctionSig)
}

// SetResource is a paid mutator transaction binding the contract method 0xde319d99.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress, bytes4 depositFunctionSig, uint256 depositFunctionDepositerOffset, bytes4 executeFunctionSig) returns()
func (_GenericHandler *GenericHandlerTransactorSession) SetResource(resourceID [32]byte, contractAddress common.Address, depositFunctionSig [4]byte, depositFunctionDepositerOffset *big.Int, executeFunctionSig [4]byte) (*types.Transaction, error) {
	return _GenericHandler.Contract.SetResource(&_GenericHandler.TransactOpts, resourceID, contractAddress, depositFunctionSig, depositFunctionDepositerOffset, executeFunctionSig)
}

// GenericHandlerFailedHandlerExecutionIterator is returned from FilterFailedHandlerExecution and is used to iterate over the raw logs and unpacked data for FailedHandlerExecution events raised by the GenericHandler contract.
type GenericHandlerFailedHandlerExecutionIterator struct {
	Event *GenericHandlerFailedHandlerExecution // Event containing the contract specifics and raw log

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
func (it *GenericHandlerFailedHandlerExecutionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GenericHandlerFailedHandlerExecution)
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
		it.Event = new(GenericHandlerFailedHandlerExecution)
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
func (it *GenericHandlerFailedHandlerExecutionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GenericHandlerFailedHandlerExecutionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GenericHandlerFailedHandlerExecution represents a FailedHandlerExecution event raised by the GenericHandler contract.
type GenericHandlerFailedHandlerExecution struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterFailedHandlerExecution is a free log retrieval operation binding the contract event 0x3f51c54fb334903765b4de0daee60de6eff2a87aef23d5a1f311f4ea2e1b7e96.
//
// Solidity: event FailedHandlerExecution()
func (_GenericHandler *GenericHandlerFilterer) FilterFailedHandlerExecution(opts *bind.FilterOpts) (*GenericHandlerFailedHandlerExecutionIterator, error) {

	logs, sub, err := _GenericHandler.contract.FilterLogs(opts, "FailedHandlerExecution")
	if err != nil {
		return nil, err
	}
	return &GenericHandlerFailedHandlerExecutionIterator{contract: _GenericHandler.contract, event: "FailedHandlerExecution", logs: logs, sub: sub}, nil
}

// WatchFailedHandlerExecution is a free log subscription operation binding the contract event 0x3f51c54fb334903765b4de0daee60de6eff2a87aef23d5a1f311f4ea2e1b7e96.
//
// Solidity: event FailedHandlerExecution()
func (_GenericHandler *GenericHandlerFilterer) WatchFailedHandlerExecution(opts *bind.WatchOpts, sink chan<- *GenericHandlerFailedHandlerExecution) (event.Subscription, error) {

	logs, sub, err := _GenericHandler.contract.WatchLogs(opts, "FailedHandlerExecution")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GenericHandlerFailedHandlerExecution)
				if err := _GenericHandler.contract.UnpackLog(event, "FailedHandlerExecution", log); err != nil {
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
func (_GenericHandler *GenericHandlerFilterer) ParseFailedHandlerExecution(log types.Log) (*GenericHandlerFailedHandlerExecution, error) {
	event := new(GenericHandlerFailedHandlerExecution)
	if err := _GenericHandler.contract.UnpackLog(event, "FailedHandlerExecution", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
