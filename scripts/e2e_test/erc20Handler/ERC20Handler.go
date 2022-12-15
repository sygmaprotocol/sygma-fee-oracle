// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ERC20Handler

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

// ERC20HandlerMetaData contains all meta data concerning the ERC20Handler contract.
var ERC20HandlerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_burnList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToTokenContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_tokenContractAddressToResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setBurnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"args\",\"type\":\"bytes\"}],\"name\":\"setResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b50604051620018cb380380620018cb8339818101604052810190620000379190620000de565b808073ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1681525050505062000110565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620000a68262000079565b9050919050565b620000b88162000099565b8114620000c457600080fd5b50565b600081519050620000d881620000ad565b92915050565b600060208284031215620000f757620000f662000074565b5b60006200010784828501620000c7565b91505092915050565b60805161179862000133600039600081816102a8015261066601526117986000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c80637f79bea8116100665780637f79bea814610159578063b07e54bb14610189578063c8ba6c87146101b9578063e248cff2146101e9578063fa8675b0146102055761009e565b806307b7ed99146100a35780630968f264146100bf5780630a6d55d8146100db578063318c136e1461010b5780636a70d08114610129575b600080fd5b6100bd60048036038101906100b89190610cc8565b610221565b005b6100d960048036038101906100d49190610e3b565b610235565b005b6100f560048036038101906100f09190610eba565b610273565b6040516101029190610ef6565b60405180910390f35b6101136102a6565b6040516101209190610ef6565b60405180910390f35b610143600480360381019061013e9190610cc8565b6102ca565b6040516101509190610f2c565b60405180910390f35b610173600480360381019061016e9190610cc8565b6102ea565b6040516101809190610f2c565b60405180910390f35b6101a3600480360381019061019e9190610fa7565b61030a565b6040516101b091906110a3565b60405180910390f35b6101d360048036038101906101ce9190610cc8565b610463565b6040516101e091906110d4565b60405180910390f35b61020360048036038101906101fe91906110ef565b61047b565b005b61021f600480360381019061021a9190610fa7565b61064c565b005b610229610664565b610232816106f4565b50565b61023d610664565b60008060008380602001905181019061025691906111c3565b80935081945082955050505061026d8383836107db565b50505050565b60006020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b7f000000000000000000000000000000000000000000000000000000000000000081565b60036020528060005260406000206000915054906101000a900460ff1681565b60026020528060005260406000206000915054906101000a900460ff1681565b6060610314610664565b60008383810190610325919061122b565b9050600080600088815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600260008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166103ea576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103e1906112db565b60405180910390fd5b600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161561044c576104478187846107f1565b610459565b61045881873085610869565b5b5050949350505050565b60016020528060005260406000206000915090505481565b610483610664565b6000806060848481019061049791906112fb565b809350819450505084846040908460406104b1919061136a565b926104be939291906113ca565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050905060008060008089815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060208301519150600260008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166105ce576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105c5906112db565b60405180910390fd5b600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16156106335761062e818360601c87610881565b610642565b610641818360601c876107db565b5b5050505050505050565b610654610664565b61065e84846108f9565b50505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106f2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e990611451565b60405180910390fd5b565b600260008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610780576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610777906114e3565b60405180910390fd5b6001600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b60008390506107eb8184846109ea565b50505050565b60008390508073ffffffffffffffffffffffffffffffffffffffff166379cc679084846040518363ffffffff1660e01b8152600401610831929190611512565b600060405180830381600087803b15801561084b57600080fd5b505af115801561085f573d6000803e3d6000fd5b5050505050505050565b600084905061087a81858585610a70565b5050505050565b60008390508073ffffffffffffffffffffffffffffffffffffffff166340c10f1984846040518363ffffffff1660e01b81526004016108c1929190611512565b600060405180830381600087803b1580156108db57600080fd5b505af11580156108ef573d6000803e3d6000fd5b5050505050505050565b8060008084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b610a6b8363a9059cbb60e01b8484604051602401610a09929190611512565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610af9565b505050565b610af3846323b872dd60e01b858585604051602401610a919392919061153b565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610af9565b50505050565b6000823b905060008111610b42576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b39906115be565b60405180910390fd5b6000808473ffffffffffffffffffffffffffffffffffffffff1684604051610b6a919061161a565b6000604051808303816000865af19150503d8060008114610ba7576040519150601f19603f3d011682016040523d82523d6000602084013e610bac565b606091505b509150915081610bf1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610be89061167d565b60405180910390fd5b600081511115610c4f5780806020019051810190610c0f91906116c9565b610c4e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c4590611742565b60405180910390fd5b5b5050505050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610c9582610c6a565b9050919050565b610ca581610c8a565b8114610cb057600080fd5b50565b600081359050610cc281610c9c565b92915050565b600060208284031215610cde57610cdd610c60565b5b6000610cec84828501610cb3565b91505092915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610d4882610cff565b810181811067ffffffffffffffff82111715610d6757610d66610d10565b5b80604052505050565b6000610d7a610c56565b9050610d868282610d3f565b919050565b600067ffffffffffffffff821115610da657610da5610d10565b5b610daf82610cff565b9050602081019050919050565b82818337600083830152505050565b6000610dde610dd984610d8b565b610d70565b905082815260208101848484011115610dfa57610df9610cfa565b5b610e05848285610dbc565b509392505050565b600082601f830112610e2257610e21610cf5565b5b8135610e32848260208601610dcb565b91505092915050565b600060208284031215610e5157610e50610c60565b5b600082013567ffffffffffffffff811115610e6f57610e6e610c65565b5b610e7b84828501610e0d565b91505092915050565b6000819050919050565b610e9781610e84565b8114610ea257600080fd5b50565b600081359050610eb481610e8e565b92915050565b600060208284031215610ed057610ecf610c60565b5b6000610ede84828501610ea5565b91505092915050565b610ef081610c8a565b82525050565b6000602082019050610f0b6000830184610ee7565b92915050565b60008115159050919050565b610f2681610f11565b82525050565b6000602082019050610f416000830184610f1d565b92915050565b600080fd5b600080fd5b60008083601f840112610f6757610f66610cf5565b5b8235905067ffffffffffffffff811115610f8457610f83610f47565b5b602083019150836001820283011115610fa057610f9f610f4c565b5b9250929050565b60008060008060608587031215610fc157610fc0610c60565b5b6000610fcf87828801610ea5565b9450506020610fe087828801610cb3565b935050604085013567ffffffffffffffff81111561100157611000610c65565b5b61100d87828801610f51565b925092505092959194509250565b600081519050919050565b600082825260208201905092915050565b60005b8381101561105557808201518184015260208101905061103a565b83811115611064576000848401525b50505050565b60006110758261101b565b61107f8185611026565b935061108f818560208601611037565b61109881610cff565b840191505092915050565b600060208201905081810360008301526110bd818461106a565b905092915050565b6110ce81610e84565b82525050565b60006020820190506110e960008301846110c5565b92915050565b60008060006040848603121561110857611107610c60565b5b600061111686828701610ea5565b935050602084013567ffffffffffffffff81111561113757611136610c65565b5b61114386828701610f51565b92509250509250925092565b600061115a82610c6a565b9050919050565b61116a8161114f565b811461117557600080fd5b50565b60008151905061118781611161565b92915050565b6000819050919050565b6111a08161118d565b81146111ab57600080fd5b50565b6000815190506111bd81611197565b92915050565b6000806000606084860312156111dc576111db610c60565b5b60006111ea86828701611178565b93505060206111fb86828701611178565b925050604061120c868287016111ae565b9150509250925092565b60008135905061122581611197565b92915050565b60006020828403121561124157611240610c60565b5b600061124f84828501611216565b91505092915050565b600082825260208201905092915050565b7f70726f766964656420746f6b656e41646472657373206973206e6f742077686960008201527f74656c6973746564000000000000000000000000000000000000000000000000602082015250565b60006112c5602883611258565b91506112d082611269565b604082019050919050565b600060208201905081810360008301526112f4816112b8565b9050919050565b6000806040838503121561131257611311610c60565b5b600061132085828601611216565b925050602061133185828601611216565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006113758261118d565b91506113808361118d565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156113b5576113b461133b565b5b828201905092915050565b600080fd5b600080fd5b600080858511156113de576113dd6113c0565b5b838611156113ef576113ee6113c5565b5b6001850283019150848603905094509492505050565b7f73656e646572206d7573742062652062726964676520636f6e74726163740000600082015250565b600061143b601e83611258565b915061144682611405565b602082019050919050565b6000602082019050818103600083015261146a8161142e565b9050919050565b7f70726f766964656420636f6e7472616374206973206e6f742077686974656c6960008201527f7374656400000000000000000000000000000000000000000000000000000000602082015250565b60006114cd602483611258565b91506114d882611471565b604082019050919050565b600060208201905081810360008301526114fc816114c0565b9050919050565b61150c8161118d565b82525050565b60006040820190506115276000830185610ee7565b6115346020830184611503565b9392505050565b60006060820190506115506000830186610ee7565b61155d6020830185610ee7565b61156a6040830184611503565b949350505050565b7f45524332303a206e6f74206120636f6e74726163740000000000000000000000600082015250565b60006115a8601583611258565b91506115b382611572565b602082019050919050565b600060208201905081810360008301526115d78161159b565b9050919050565b600081905092915050565b60006115f48261101b565b6115fe81856115de565b935061160e818560208601611037565b80840191505092915050565b600061162682846115e9565b915081905092915050565b7f45524332303a2063616c6c206661696c65640000000000000000000000000000600082015250565b6000611667601283611258565b915061167282611631565b602082019050919050565b600060208201905081810360008301526116968161165a565b9050919050565b6116a681610f11565b81146116b157600080fd5b50565b6000815190506116c38161169d565b92915050565b6000602082840312156116df576116de610c60565b5b60006116ed848285016116b4565b91505092915050565b7f45524332303a206f7065726174696f6e20646964206e6f742073756363656564600082015250565b600061172c602083611258565b9150611737826116f6565b602082019050919050565b6000602082019050818103600083015261175b8161171f565b905091905056fea2646970667358221220c2a50cacc0fdf31a7eebc37abb47bd85c208f4a45ac46c14473f6e32f7c5ed1364736f6c634300080b0033",
}

// ERC20HandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20HandlerMetaData.ABI instead.
var ERC20HandlerABI = ERC20HandlerMetaData.ABI

// ERC20HandlerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC20HandlerMetaData.Bin instead.
var ERC20HandlerBin = ERC20HandlerMetaData.Bin

// DeployERC20Handler deploys a new Ethereum contract, binding an instance of ERC20Handler to it.
func DeployERC20Handler(auth *bind.TransactOpts, backend bind.ContractBackend, bridgeAddress common.Address) (common.Address, *types.Transaction, *ERC20Handler, error) {
	parsed, err := ERC20HandlerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20HandlerBin), backend, bridgeAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Handler{ERC20HandlerCaller: ERC20HandlerCaller{contract: contract}, ERC20HandlerTransactor: ERC20HandlerTransactor{contract: contract}, ERC20HandlerFilterer: ERC20HandlerFilterer{contract: contract}}, nil
}

// ERC20Handler is an auto generated Go binding around an Ethereum contract.
type ERC20Handler struct {
	ERC20HandlerCaller     // Read-only binding to the contract
	ERC20HandlerTransactor // Write-only binding to the contract
	ERC20HandlerFilterer   // Log filterer for contract events
}

// ERC20HandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20HandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20HandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20HandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20HandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20HandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20HandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20HandlerSession struct {
	Contract     *ERC20Handler     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20HandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20HandlerCallerSession struct {
	Contract *ERC20HandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ERC20HandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20HandlerTransactorSession struct {
	Contract     *ERC20HandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ERC20HandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20HandlerRaw struct {
	Contract *ERC20Handler // Generic contract binding to access the raw methods on
}

// ERC20HandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20HandlerCallerRaw struct {
	Contract *ERC20HandlerCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20HandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20HandlerTransactorRaw struct {
	Contract *ERC20HandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Handler creates a new instance of ERC20Handler, bound to a specific deployed contract.
func NewERC20Handler(address common.Address, backend bind.ContractBackend) (*ERC20Handler, error) {
	contract, err := bindERC20Handler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Handler{ERC20HandlerCaller: ERC20HandlerCaller{contract: contract}, ERC20HandlerTransactor: ERC20HandlerTransactor{contract: contract}, ERC20HandlerFilterer: ERC20HandlerFilterer{contract: contract}}, nil
}

// NewERC20HandlerCaller creates a new read-only instance of ERC20Handler, bound to a specific deployed contract.
func NewERC20HandlerCaller(address common.Address, caller bind.ContractCaller) (*ERC20HandlerCaller, error) {
	contract, err := bindERC20Handler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20HandlerCaller{contract: contract}, nil
}

// NewERC20HandlerTransactor creates a new write-only instance of ERC20Handler, bound to a specific deployed contract.
func NewERC20HandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20HandlerTransactor, error) {
	contract, err := bindERC20Handler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20HandlerTransactor{contract: contract}, nil
}

// NewERC20HandlerFilterer creates a new log filterer instance of ERC20Handler, bound to a specific deployed contract.
func NewERC20HandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20HandlerFilterer, error) {
	contract, err := bindERC20Handler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20HandlerFilterer{contract: contract}, nil
}

// bindERC20Handler binds a generic wrapper to an already deployed contract.
func bindERC20Handler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20HandlerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Handler *ERC20HandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Handler.Contract.ERC20HandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Handler *ERC20HandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Handler.Contract.ERC20HandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Handler *ERC20HandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Handler.Contract.ERC20HandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Handler *ERC20HandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Handler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Handler *ERC20HandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Handler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Handler *ERC20HandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Handler.Contract.contract.Transact(opts, method, params...)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_ERC20Handler *ERC20HandlerCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Handler.contract.Call(opts, &out, "_bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_ERC20Handler *ERC20HandlerSession) BridgeAddress() (common.Address, error) {
	return _ERC20Handler.Contract.BridgeAddress(&_ERC20Handler.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_ERC20Handler *ERC20HandlerCallerSession) BridgeAddress() (common.Address, error) {
	return _ERC20Handler.Contract.BridgeAddress(&_ERC20Handler.CallOpts)
}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_ERC20Handler *ERC20HandlerCaller) BurnList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ERC20Handler.contract.Call(opts, &out, "_burnList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_ERC20Handler *ERC20HandlerSession) BurnList(arg0 common.Address) (bool, error) {
	return _ERC20Handler.Contract.BurnList(&_ERC20Handler.CallOpts, arg0)
}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_ERC20Handler *ERC20HandlerCallerSession) BurnList(arg0 common.Address) (bool, error) {
	return _ERC20Handler.Contract.BurnList(&_ERC20Handler.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_ERC20Handler *ERC20HandlerCaller) ContractWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ERC20Handler.contract.Call(opts, &out, "_contractWhitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_ERC20Handler *ERC20HandlerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _ERC20Handler.Contract.ContractWhitelist(&_ERC20Handler.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_ERC20Handler *ERC20HandlerCallerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _ERC20Handler.Contract.ContractWhitelist(&_ERC20Handler.CallOpts, arg0)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_ERC20Handler *ERC20HandlerCaller) ResourceIDToTokenContractAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ERC20Handler.contract.Call(opts, &out, "_resourceIDToTokenContractAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_ERC20Handler *ERC20HandlerSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _ERC20Handler.Contract.ResourceIDToTokenContractAddress(&_ERC20Handler.CallOpts, arg0)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_ERC20Handler *ERC20HandlerCallerSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _ERC20Handler.Contract.ResourceIDToTokenContractAddress(&_ERC20Handler.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_ERC20Handler *ERC20HandlerCaller) TokenContractAddressToResourceID(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _ERC20Handler.contract.Call(opts, &out, "_tokenContractAddressToResourceID", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_ERC20Handler *ERC20HandlerSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _ERC20Handler.Contract.TokenContractAddressToResourceID(&_ERC20Handler.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_ERC20Handler *ERC20HandlerCallerSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _ERC20Handler.Contract.TokenContractAddressToResourceID(&_ERC20Handler.CallOpts, arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0xb07e54bb.
//
// Solidity: function deposit(bytes32 resourceID, address depositor, bytes data) returns(bytes)
func (_ERC20Handler *ERC20HandlerTransactor) Deposit(opts *bind.TransactOpts, resourceID [32]byte, depositor common.Address, data []byte) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "deposit", resourceID, depositor, data)
}

// Deposit is a paid mutator transaction binding the contract method 0xb07e54bb.
//
// Solidity: function deposit(bytes32 resourceID, address depositor, bytes data) returns(bytes)
func (_ERC20Handler *ERC20HandlerSession) Deposit(resourceID [32]byte, depositor common.Address, data []byte) (*types.Transaction, error) {
	return _ERC20Handler.Contract.Deposit(&_ERC20Handler.TransactOpts, resourceID, depositor, data)
}

// Deposit is a paid mutator transaction binding the contract method 0xb07e54bb.
//
// Solidity: function deposit(bytes32 resourceID, address depositor, bytes data) returns(bytes)
func (_ERC20Handler *ERC20HandlerTransactorSession) Deposit(resourceID [32]byte, depositor common.Address, data []byte) (*types.Transaction, error) {
	return _ERC20Handler.Contract.Deposit(&_ERC20Handler.TransactOpts, resourceID, depositor, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_ERC20Handler *ERC20HandlerTransactor) ExecuteProposal(opts *bind.TransactOpts, resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "executeProposal", resourceID, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_ERC20Handler *ERC20HandlerSession) ExecuteProposal(resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _ERC20Handler.Contract.ExecuteProposal(&_ERC20Handler.TransactOpts, resourceID, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) ExecuteProposal(resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _ERC20Handler.Contract.ExecuteProposal(&_ERC20Handler.TransactOpts, resourceID, data)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_ERC20Handler *ERC20HandlerTransactor) SetBurnable(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "setBurnable", contractAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_ERC20Handler *ERC20HandlerSession) SetBurnable(contractAddress common.Address) (*types.Transaction, error) {
	return _ERC20Handler.Contract.SetBurnable(&_ERC20Handler.TransactOpts, contractAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) SetBurnable(contractAddress common.Address) (*types.Transaction, error) {
	return _ERC20Handler.Contract.SetBurnable(&_ERC20Handler.TransactOpts, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xfa8675b0.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress, bytes args) returns()
func (_ERC20Handler *ERC20HandlerTransactor) SetResource(opts *bind.TransactOpts, resourceID [32]byte, contractAddress common.Address, args []byte) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "setResource", resourceID, contractAddress, args)
}

// SetResource is a paid mutator transaction binding the contract method 0xfa8675b0.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress, bytes args) returns()
func (_ERC20Handler *ERC20HandlerSession) SetResource(resourceID [32]byte, contractAddress common.Address, args []byte) (*types.Transaction, error) {
	return _ERC20Handler.Contract.SetResource(&_ERC20Handler.TransactOpts, resourceID, contractAddress, args)
}

// SetResource is a paid mutator transaction binding the contract method 0xfa8675b0.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress, bytes args) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) SetResource(resourceID [32]byte, contractAddress common.Address, args []byte) (*types.Transaction, error) {
	return _ERC20Handler.Contract.SetResource(&_ERC20Handler.TransactOpts, resourceID, contractAddress, args)
}

// Withdraw is a paid mutator transaction binding the contract method 0x0968f264.
//
// Solidity: function withdraw(bytes data) returns()
func (_ERC20Handler *ERC20HandlerTransactor) Withdraw(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "withdraw", data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x0968f264.
//
// Solidity: function withdraw(bytes data) returns()
func (_ERC20Handler *ERC20HandlerSession) Withdraw(data []byte) (*types.Transaction, error) {
	return _ERC20Handler.Contract.Withdraw(&_ERC20Handler.TransactOpts, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x0968f264.
//
// Solidity: function withdraw(bytes data) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) Withdraw(data []byte) (*types.Transaction, error) {
	return _ERC20Handler.Contract.Withdraw(&_ERC20Handler.TransactOpts, data)
}
