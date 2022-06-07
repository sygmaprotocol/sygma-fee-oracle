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
const GenericHandlerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractAddressToDepositFunctionDepositerOffset\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractAddressToDepositFunctionSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractAddressToExecuteFunctionSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractAddressToResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"depositFunctionSig\",\"type\":\"bytes4\"},{\"internalType\":\"uint256\",\"name\":\"depositFunctionDepositerOffset\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"executeFunctionSig\",\"type\":\"bytes4\"}],\"name\":\"setResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GenericHandlerBin is the compiled bytecode used for deploying new contracts.
var GenericHandlerBin = "0x60a06040523480156200001157600080fd5b506040516200166c3803806200166c8339818101604052810190620000379190620000dc565b8073ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1681525050506200010e565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620000a48262000077565b9050919050565b620000b68162000097565b8114620000c257600080fd5b50565b600081519050620000d681620000ab565b92915050565b600060208284031215620000f557620000f462000072565b5b60006200010584828501620000c5565b91505092915050565b60805161153b620001316000396000818161024b01526109aa015261153b6000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c8063c54c2a1111610066578063c54c2a1114610181578063cb624463146101b1578063de319d99146101e1578063e248cff2146101fd578063ec97d3b4146102195761009e565b8063318c136e146100a35780637f79bea8146100c1578063a5c3a985146100f1578063aa50800b14610121578063b07e54bb14610151575b600080fd5b6100ab610249565b6040516100b89190610c67565b60405180910390f35b6100db60048036038101906100d69190610cb8565b61026d565b6040516100e89190610d00565b60405180910390f35b61010b60048036038101906101069190610cb8565b61028d565b6040516101189190610d56565b60405180910390f35b61013b60048036038101906101369190610cb8565b6102ad565b6040516101489190610d8a565b60405180910390f35b61016b60048036038101906101669190610e40565b6102c5565b6040516101789190610f4d565b60405180910390f35b61019b60048036038101906101969190610f6f565b610664565b6040516101a89190610c67565b60405180910390f35b6101cb60048036038101906101c69190610cb8565b610697565b6040516101d89190610d56565b60405180910390f35b6101fb60048036038101906101f69190610ff4565b6106b7565b005b6102176004803603810190610212919061106f565b6106d3565b005b610233600480360381019061022e9190610cb8565b610990565b60405161024091906110de565b60405180910390f35b7f000000000000000000000000000000000000000000000000000000000000000081565b60056020528060005260406000206000915054906101000a900460ff1681565b60046020528060005260406000206000915054906101000a900460e01b81565b60036020528060005260406000206000915090505481565b60606102cf6109a8565b6000606084848101906102e291906110f9565b915084846020908460206102f69190611155565b92610303939291906111b5565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509050600080600089815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506000600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050600081111561044b576000816020850101519050606081901c73ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff1614610449576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104409061124d565b60405180910390fd5b505b600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166104d7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104ce906112df565b60405180910390fd5b6000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460e01b9050600060e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614610656576000818560405160200161058692919061135c565b60405160208183030381529060405290506000808573ffffffffffffffffffffffffffffffffffffffff16836040516105bf9190611384565b6000604051808303816000865af19150503d80600081146105fc576040519150601f19603f3d011682016040523d82523d6000602084013e610601565b606091505b509150915081610646576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161063d906113e7565b60405180910390fd5b809850505050505050505061065c565b50505050505b949350505050565b60006020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60026020528060005260406000206000915054906101000a900460e01b81565b6106bf6109a8565b6106cc8585858585610a38565b5050505050565b6106db6109a8565b6000606083838101906106ee91906110f9565b915083836020908460206107029190611155565b9261070f939291906111b5565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509050600080600087815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600560008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610817576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161080e906112df565b60405180910390fd5b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460e01b9050600060e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161461098757600081846040516020016108c692919061135c565b604051602081830303815290604052905060008373ffffffffffffffffffffffffffffffffffffffff16826040516108fe9190611384565b6000604051808303816000865af19150503d806000811461093b576040519150601f19603f3d011682016040523d82523d6000602084013e610940565b606091505b5050905080610984576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097b90611479565b60405180910390fd5b50505b50505050505050565b60016020528060005260406000206000915090505481565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a36576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a2d906114e5565b60405180910390fd5b565b8360008087815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555084600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555082600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908360e01c021790555081600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555080600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908360e01c02179055506001600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050505050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610c5182610c26565b9050919050565b610c6181610c46565b82525050565b6000602082019050610c7c6000830184610c58565b92915050565b600080fd5b600080fd5b610c9581610c46565b8114610ca057600080fd5b50565b600081359050610cb281610c8c565b92915050565b600060208284031215610cce57610ccd610c82565b5b6000610cdc84828501610ca3565b91505092915050565b60008115159050919050565b610cfa81610ce5565b82525050565b6000602082019050610d156000830184610cf1565b92915050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b610d5081610d1b565b82525050565b6000602082019050610d6b6000830184610d47565b92915050565b6000819050919050565b610d8481610d71565b82525050565b6000602082019050610d9f6000830184610d7b565b92915050565b6000819050919050565b610db881610da5565b8114610dc357600080fd5b50565b600081359050610dd581610daf565b92915050565b600080fd5b600080fd5b600080fd5b60008083601f840112610e0057610dff610ddb565b5b8235905067ffffffffffffffff811115610e1d57610e1c610de0565b5b602083019150836001820283011115610e3957610e38610de5565b5b9250929050565b60008060008060608587031215610e5a57610e59610c82565b5b6000610e6887828801610dc6565b9450506020610e7987828801610ca3565b935050604085013567ffffffffffffffff811115610e9a57610e99610c87565b5b610ea687828801610dea565b925092505092959194509250565b600081519050919050565b600082825260208201905092915050565b60005b83811015610eee578082015181840152602081019050610ed3565b83811115610efd576000848401525b50505050565b6000601f19601f8301169050919050565b6000610f1f82610eb4565b610f298185610ebf565b9350610f39818560208601610ed0565b610f4281610f03565b840191505092915050565b60006020820190508181036000830152610f678184610f14565b905092915050565b600060208284031215610f8557610f84610c82565b5b6000610f9384828501610dc6565b91505092915050565b610fa581610d1b565b8114610fb057600080fd5b50565b600081359050610fc281610f9c565b92915050565b610fd181610d71565b8114610fdc57600080fd5b50565b600081359050610fee81610fc8565b92915050565b600080600080600060a086880312156110105761100f610c82565b5b600061101e88828901610dc6565b955050602061102f88828901610ca3565b945050604061104088828901610fb3565b935050606061105188828901610fdf565b925050608061106288828901610fb3565b9150509295509295909350565b60008060006040848603121561108857611087610c82565b5b600061109686828701610dc6565b935050602084013567ffffffffffffffff8111156110b7576110b6610c87565b5b6110c386828701610dea565b92509250509250925092565b6110d881610da5565b82525050565b60006020820190506110f360008301846110cf565b92915050565b60006020828403121561110f5761110e610c82565b5b600061111d84828501610fdf565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061116082610d71565b915061116b83610d71565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156111a05761119f611126565b5b828201905092915050565b600080fd5b600080fd5b600080858511156111c9576111c86111ab565b5b838611156111da576111d96111b0565b5b6001850283019150848603905094509492505050565b600082825260208201905092915050565b7f696e636f7272656374206465706f736974657220696e20746865206461746100600082015250565b6000611237601f836111f0565b915061124282611201565b602082019050919050565b600060208201905081810360008301526112668161122a565b9050919050565b7f70726f766964656420636f6e747261637441646472657373206973206e6f742060008201527f77686974656c6973746564000000000000000000000000000000000000000000602082015250565b60006112c9602b836111f0565b91506112d48261126d565b604082019050919050565b600060208201905081810360008301526112f8816112bc565b9050919050565b6000819050919050565b61131a61131582610d1b565b6112ff565b82525050565b600081905092915050565b600061133682610eb4565b6113408185611320565b9350611350818560208601610ed0565b80840191505092915050565b60006113688285611309565b600482019150611378828461132b565b91508190509392505050565b6000611390828461132b565b915081905092915050565b7f63616c6c20746f20636f6e747261637441646472657373206661696c65640000600082015250565b60006113d1601e836111f0565b91506113dc8261139b565b602082019050919050565b60006020820190508181036000830152611400816113c4565b9050919050565b7f64656c656761746563616c6c20746f20636f6e7472616374416464726573732060008201527f6661696c65640000000000000000000000000000000000000000000000000000602082015250565b60006114636026836111f0565b915061146e82611407565b604082019050919050565b6000602082019050818103600083015261149281611456565b9050919050565b7f73656e646572206d7573742062652062726964676520636f6e74726163740000600082015250565b60006114cf601e836111f0565b91506114da82611499565b602082019050919050565b600060208201905081810360008301526114fe816114c2565b905091905056fea26469706673582212203b185c419cdcb74b5464333fdbe25ae4ddde21c598c761bf2ba54e36b601422f64736f6c634300080b0033"

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
