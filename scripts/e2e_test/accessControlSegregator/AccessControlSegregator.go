// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package AccessControlSegregator

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

// AccessControlSegregatorABI is the input ABI used to generate the binding from.
const AccessControlSegregatorABI = "[{\"inputs\":[{\"internalType\":\"bytes4[]\",\"name\":\"functions\",\"type\":\"bytes4[]\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"GRANT_ACCESS_SIG\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"name\":\"functionAccess\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantAccess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasAccess\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// AccessControlSegregatorBin is the compiled bytecode used for deploying new contracts.
var AccessControlSegregatorBin = "0x60806040523480156200001157600080fd5b5060405162000bdf38038062000bdf8339818101604052810190620000379190620004c3565b81518151146200007e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200007590620005a9565b60405180910390fd5b6200009763a973ec9360e01b336200010f60201b60201c565b60005b81518110156200010657620000f0838281518110620000be57620000bd620005cb565b5b6020026020010151838381518110620000dc57620000db620005cb565b5b60200260200101516200010f60201b60201c565b8080620000fd9062000633565b9150506200009a565b50505062000681565b80600080847bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b6000604051905090565b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6200020682620001bb565b810181811067ffffffffffffffff82111715620002285762000227620001cc565b5b80604052505050565b60006200023d620001a2565b90506200024b8282620001fb565b919050565b600067ffffffffffffffff8211156200026e576200026d620001cc565b5b602082029050602081019050919050565b600080fd5b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b620002bb8162000284565b8114620002c757600080fd5b50565b600081519050620002db81620002b0565b92915050565b6000620002f8620002f28462000250565b62000231565b905080838252602082019050602084028301858111156200031e576200031d6200027f565b5b835b818110156200034b5780620003368882620002ca565b84526020840193505060208101905062000320565b5050509392505050565b600082601f8301126200036d576200036c620001b6565b5b81516200037f848260208601620002e1565b91505092915050565b600067ffffffffffffffff821115620003a657620003a5620001cc565b5b602082029050602081019050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620003e482620003b7565b9050919050565b620003f681620003d7565b81146200040257600080fd5b50565b6000815190506200041681620003eb565b92915050565b6000620004336200042d8462000388565b62000231565b905080838252602082019050602084028301858111156200045957620004586200027f565b5b835b8181101562000486578062000471888262000405565b8452602084019350506020810190506200045b565b5050509392505050565b600082601f830112620004a857620004a7620001b6565b5b8151620004ba8482602086016200041c565b91505092915050565b60008060408385031215620004dd57620004dc620001ac565b5b600083015167ffffffffffffffff811115620004fe57620004fd620001b1565b5b6200050c8582860162000355565b925050602083015167ffffffffffffffff81111562000530576200052f620001b1565b5b6200053e8582860162000490565b9150509250929050565b600082825260208201905092915050565b7f6172726179206c656e6774682073686f756c6420626520657175616c00000000600082015250565b600062000591601c8362000548565b91506200059e8262000559565b602082019050919050565b60006020820190508181036000830152620005c48162000582565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000819050919050565b6000620006408262000629565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415620006765762000675620005fa565b5b600182019050919050565b61054e80620006916000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806371c9521c146100515780639bec8dd314610081578063a973ec931461009f578063f2bff788146100bb575b600080fd5b61006b6004803603810190610066919061037e565b6100eb565b60405161007891906103d9565b60405180910390f35b610089610194565b6040516100969190610403565b60405180910390f35b6100b960048036038101906100b4919061037e565b61019f565b005b6100d560048036038101906100d0919061041e565b6101fd565b6040516100e2919061045a565b60405180910390f35b60008173ffffffffffffffffffffffffffffffffffffffff16600080857bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614905092915050565b63a973ec9360e01b81565b6101b063a973ec9360e01b336100eb565b6101ef576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101e6906104f8565b60405180910390fd5b6101f98282610230565b5050565b60006020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b80600080847bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b600080fd5b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6102fd816102c8565b811461030857600080fd5b50565b60008135905061031a816102f4565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061034b82610320565b9050919050565b61035b81610340565b811461036657600080fd5b50565b60008135905061037881610352565b92915050565b60008060408385031215610395576103946102c3565b5b60006103a38582860161030b565b92505060206103b485828601610369565b9150509250929050565b60008115159050919050565b6103d3816103be565b82525050565b60006020820190506103ee60008301846103ca565b92915050565b6103fd816102c8565b82525050565b600060208201905061041860008301846103f4565b92915050565b600060208284031215610434576104336102c3565b5b60006104428482850161030b565b91505092915050565b61045481610340565b82525050565b600060208201905061046f600083018461044b565b92915050565b600082825260208201905092915050565b7f73656e64657220646f65736e27742068617665206772616e742061636365737360008201527f2072696768747300000000000000000000000000000000000000000000000000602082015250565b60006104e2602783610475565b91506104ed82610486565b604082019050919050565b60006020820190508181036000830152610511816104d5565b905091905056fea2646970667358221220bab3228cb515e058b1ec2b6db860641f8d9a3f423ec9b7befeea304cd146ad1f64736f6c634300080b0033"

// DeployAccessControlSegregator deploys a new Ethereum contract, binding an instance of AccessControlSegregator to it.
func DeployAccessControlSegregator(auth *bind.TransactOpts, backend bind.ContractBackend, functions [][4]byte, accounts []common.Address) (common.Address, *types.Transaction, *AccessControlSegregator, error) {
	parsed, err := abi.JSON(strings.NewReader(AccessControlSegregatorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AccessControlSegregatorBin), backend, functions, accounts)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AccessControlSegregator{AccessControlSegregatorCaller: AccessControlSegregatorCaller{contract: contract}, AccessControlSegregatorTransactor: AccessControlSegregatorTransactor{contract: contract}, AccessControlSegregatorFilterer: AccessControlSegregatorFilterer{contract: contract}}, nil
}

// AccessControlSegregator is an auto generated Go binding around an Ethereum contract.
type AccessControlSegregator struct {
	AccessControlSegregatorCaller     // Read-only binding to the contract
	AccessControlSegregatorTransactor // Write-only binding to the contract
	AccessControlSegregatorFilterer   // Log filterer for contract events
}

// AccessControlSegregatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccessControlSegregatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlSegregatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccessControlSegregatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlSegregatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccessControlSegregatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlSegregatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccessControlSegregatorSession struct {
	Contract     *AccessControlSegregator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AccessControlSegregatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccessControlSegregatorCallerSession struct {
	Contract *AccessControlSegregatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// AccessControlSegregatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccessControlSegregatorTransactorSession struct {
	Contract     *AccessControlSegregatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// AccessControlSegregatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccessControlSegregatorRaw struct {
	Contract *AccessControlSegregator // Generic contract binding to access the raw methods on
}

// AccessControlSegregatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccessControlSegregatorCallerRaw struct {
	Contract *AccessControlSegregatorCaller // Generic read-only contract binding to access the raw methods on
}

// AccessControlSegregatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccessControlSegregatorTransactorRaw struct {
	Contract *AccessControlSegregatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccessControlSegregator creates a new instance of AccessControlSegregator, bound to a specific deployed contract.
func NewAccessControlSegregator(address common.Address, backend bind.ContractBackend) (*AccessControlSegregator, error) {
	contract, err := bindAccessControlSegregator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccessControlSegregator{AccessControlSegregatorCaller: AccessControlSegregatorCaller{contract: contract}, AccessControlSegregatorTransactor: AccessControlSegregatorTransactor{contract: contract}, AccessControlSegregatorFilterer: AccessControlSegregatorFilterer{contract: contract}}, nil
}

// NewAccessControlSegregatorCaller creates a new read-only instance of AccessControlSegregator, bound to a specific deployed contract.
func NewAccessControlSegregatorCaller(address common.Address, caller bind.ContractCaller) (*AccessControlSegregatorCaller, error) {
	contract, err := bindAccessControlSegregator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlSegregatorCaller{contract: contract}, nil
}

// NewAccessControlSegregatorTransactor creates a new write-only instance of AccessControlSegregator, bound to a specific deployed contract.
func NewAccessControlSegregatorTransactor(address common.Address, transactor bind.ContractTransactor) (*AccessControlSegregatorTransactor, error) {
	contract, err := bindAccessControlSegregator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlSegregatorTransactor{contract: contract}, nil
}

// NewAccessControlSegregatorFilterer creates a new log filterer instance of AccessControlSegregator, bound to a specific deployed contract.
func NewAccessControlSegregatorFilterer(address common.Address, filterer bind.ContractFilterer) (*AccessControlSegregatorFilterer, error) {
	contract, err := bindAccessControlSegregator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccessControlSegregatorFilterer{contract: contract}, nil
}

// bindAccessControlSegregator binds a generic wrapper to an already deployed contract.
func bindAccessControlSegregator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccessControlSegregatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControlSegregator *AccessControlSegregatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControlSegregator.Contract.AccessControlSegregatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControlSegregator *AccessControlSegregatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControlSegregator.Contract.AccessControlSegregatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControlSegregator *AccessControlSegregatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControlSegregator.Contract.AccessControlSegregatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControlSegregator *AccessControlSegregatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControlSegregator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControlSegregator *AccessControlSegregatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControlSegregator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControlSegregator *AccessControlSegregatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControlSegregator.Contract.contract.Transact(opts, method, params...)
}

// GRANTACCESSSIG is a free data retrieval call binding the contract method 0x9bec8dd3.
//
// Solidity: function GRANT_ACCESS_SIG() view returns(bytes4)
func (_AccessControlSegregator *AccessControlSegregatorCaller) GRANTACCESSSIG(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _AccessControlSegregator.contract.Call(opts, &out, "GRANT_ACCESS_SIG")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// GRANTACCESSSIG is a free data retrieval call binding the contract method 0x9bec8dd3.
//
// Solidity: function GRANT_ACCESS_SIG() view returns(bytes4)
func (_AccessControlSegregator *AccessControlSegregatorSession) GRANTACCESSSIG() ([4]byte, error) {
	return _AccessControlSegregator.Contract.GRANTACCESSSIG(&_AccessControlSegregator.CallOpts)
}

// GRANTACCESSSIG is a free data retrieval call binding the contract method 0x9bec8dd3.
//
// Solidity: function GRANT_ACCESS_SIG() view returns(bytes4)
func (_AccessControlSegregator *AccessControlSegregatorCallerSession) GRANTACCESSSIG() ([4]byte, error) {
	return _AccessControlSegregator.Contract.GRANTACCESSSIG(&_AccessControlSegregator.CallOpts)
}

// FunctionAccess is a free data retrieval call binding the contract method 0xf2bff788.
//
// Solidity: function functionAccess(bytes4 ) view returns(address)
func (_AccessControlSegregator *AccessControlSegregatorCaller) FunctionAccess(opts *bind.CallOpts, arg0 [4]byte) (common.Address, error) {
	var out []interface{}
	err := _AccessControlSegregator.contract.Call(opts, &out, "functionAccess", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FunctionAccess is a free data retrieval call binding the contract method 0xf2bff788.
//
// Solidity: function functionAccess(bytes4 ) view returns(address)
func (_AccessControlSegregator *AccessControlSegregatorSession) FunctionAccess(arg0 [4]byte) (common.Address, error) {
	return _AccessControlSegregator.Contract.FunctionAccess(&_AccessControlSegregator.CallOpts, arg0)
}

// FunctionAccess is a free data retrieval call binding the contract method 0xf2bff788.
//
// Solidity: function functionAccess(bytes4 ) view returns(address)
func (_AccessControlSegregator *AccessControlSegregatorCallerSession) FunctionAccess(arg0 [4]byte) (common.Address, error) {
	return _AccessControlSegregator.Contract.FunctionAccess(&_AccessControlSegregator.CallOpts, arg0)
}

// HasAccess is a free data retrieval call binding the contract method 0x71c9521c.
//
// Solidity: function hasAccess(bytes4 sig, address account) view returns(bool)
func (_AccessControlSegregator *AccessControlSegregatorCaller) HasAccess(opts *bind.CallOpts, sig [4]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AccessControlSegregator.contract.Call(opts, &out, "hasAccess", sig, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasAccess is a free data retrieval call binding the contract method 0x71c9521c.
//
// Solidity: function hasAccess(bytes4 sig, address account) view returns(bool)
func (_AccessControlSegregator *AccessControlSegregatorSession) HasAccess(sig [4]byte, account common.Address) (bool, error) {
	return _AccessControlSegregator.Contract.HasAccess(&_AccessControlSegregator.CallOpts, sig, account)
}

// HasAccess is a free data retrieval call binding the contract method 0x71c9521c.
//
// Solidity: function hasAccess(bytes4 sig, address account) view returns(bool)
func (_AccessControlSegregator *AccessControlSegregatorCallerSession) HasAccess(sig [4]byte, account common.Address) (bool, error) {
	return _AccessControlSegregator.Contract.HasAccess(&_AccessControlSegregator.CallOpts, sig, account)
}

// GrantAccess is a paid mutator transaction binding the contract method 0xa973ec93.
//
// Solidity: function grantAccess(bytes4 sig, address account) returns()
func (_AccessControlSegregator *AccessControlSegregatorTransactor) GrantAccess(opts *bind.TransactOpts, sig [4]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlSegregator.contract.Transact(opts, "grantAccess", sig, account)
}

// GrantAccess is a paid mutator transaction binding the contract method 0xa973ec93.
//
// Solidity: function grantAccess(bytes4 sig, address account) returns()
func (_AccessControlSegregator *AccessControlSegregatorSession) GrantAccess(sig [4]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlSegregator.Contract.GrantAccess(&_AccessControlSegregator.TransactOpts, sig, account)
}

// GrantAccess is a paid mutator transaction binding the contract method 0xa973ec93.
//
// Solidity: function grantAccess(bytes4 sig, address account) returns()
func (_AccessControlSegregator *AccessControlSegregatorTransactorSession) GrantAccess(sig [4]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlSegregator.Contract.GrantAccess(&_AccessControlSegregator.TransactOpts, sig, account)
}
