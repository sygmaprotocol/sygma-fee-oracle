// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ERC721Handler

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

// ERC721HandlerMetaData contains all meta data concerning the ERC721Handler contract.
var ERC721HandlerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_burnList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToTokenContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_tokenContractAddressToResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"metaData\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setBurnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b5060405162001ad838038062001ad88339818101604052810190620000379190620000de565b808073ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1681525050505062000110565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620000a68262000079565b9050919050565b620000b88162000099565b8114620000c457600080fd5b50565b600081519050620000d881620000ad565b92915050565b600060208284031215620000f757620000f662000074565b5b60006200010784828501620000c7565b91505092915050565b6080516119a562000133600039600081816102a901526107ca01526119a56000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c80637f79bea8116100665780637f79bea814610159578063b07e54bb14610189578063b8fa3736146101b9578063c8ba6c87146101d5578063e248cff2146102055761009e565b806307b7ed99146100a35780630968f264146100bf5780630a6d55d8146100db578063318c136e1461010b5780636a70d08114610129575b600080fd5b6100bd60048036038101906100b89190610e25565b610221565b005b6100d960048036038101906100d49190610f98565b610235565b005b6100f560048036038101906100f09190611017565b610274565b6040516101029190611053565b60405180910390f35b6101136102a7565b6040516101209190611053565b60405180910390f35b610143600480360381019061013e9190610e25565b6102cb565b6040516101509190611089565b60405180910390f35b610173600480360381019061016e9190610e25565b6102eb565b6040516101809190611089565b60405180910390f35b6101a3600480360381019061019e9190611104565b61030b565b6040516101b09190611200565b60405180910390f35b6101d360048036038101906101ce9190611222565b610520565b005b6101ef60048036038101906101ea9190610e25565b610536565b6040516101fc9190611271565b60405180910390f35b61021f600480360381019061021a919061128c565b61054e565b005b6102296107c8565b61023281610858565b50565b61023d6107c8565b6000806000838060200190518101906102569190611360565b80935081945082955050505061026e8330848461093f565b50505050565b60006020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b7f000000000000000000000000000000000000000000000000000000000000000081565b60036020528060005260406000206000915054906101000a900460ff1681565b60026020528060005260406000206000915054906101000a900460ff1681565b60606103156107c8565b6000838381019061032691906113c8565b9050600080600088815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600260008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166103eb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103e290611478565b60405180910390fd5b61041b635b5e139f60e01b8273ffffffffffffffffffffffffffffffffffffffff166109ba90919063ffffffff16565b156104a85760008190508073ffffffffffffffffffffffffffffffffffffffff1663c87b56dd846040518263ffffffff1660e01b815260040161045e91906114a7565b600060405180830381865afa15801561047b573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906104a49190611563565b9350505b600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16156105095761050481836109df565b610516565b61051581873085610a54565b5b5050949350505050565b6105286107c8565b6105328282610acf565b5050565b60016020528060005260406000206000915090505481565b6105566107c8565b60008060606000806060878781019061056f91906115ac565b8096508197505050846040610584919061161b565b9250878760409085926105999392919061167b565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505093508787849080926105f09392919061167b565b8101906105fd91906113c8565b9150878760208561060e919061161b565b908460208761061d919061161b565b610627919061161b565b926106349392919061167b565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050905060006020850151905060008060008c815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600260008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610745576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161073c90611478565b60405180910390fd5b600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16156107ab576107a6818360601c8a86610bc0565b6107bb565b6107ba81308460601c8b61093f565b5b5050505050505050505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610856576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161084d90611702565b60405180910390fd5b565b600260008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166108e4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108db90611794565b60405180910390fd5b6001600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b60008490508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8585856040518463ffffffff1660e01b8152600401610981939291906117b4565b600060405180830381600087803b15801561099b57600080fd5b505af11580156109af573d6000803e3d6000fd5b505050505050505050565b60006109c583610c3b565b80156109d757506109d68383610c88565b5b905092915050565b60008290508073ffffffffffffffffffffffffffffffffffffffff166342966c68836040518263ffffffff1660e01b8152600401610a1d91906114a7565b600060405180830381600087803b158015610a3757600080fd5b505af1158015610a4b573d6000803e3d6000fd5b50505050505050565b60008490508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8585856040518463ffffffff1660e01b8152600401610a96939291906117b4565b600060405180830381600087803b158015610ab057600080fd5b505af1158015610ac4573d6000803e3d6000fd5b505050505050505050565b8060008084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b60008490508073ffffffffffffffffffffffffffffffffffffffff1663d3fc98648585856040518463ffffffff1660e01b8152600401610c029392919061182f565b600060405180830381600087803b158015610c1c57600080fd5b505af1158015610c30573d6000803e3d6000fd5b505050505050505050565b6000610c67827f01ffc9a700000000000000000000000000000000000000000000000000000000610c88565b8015610c815750610c7f8263ffffffff60e01b610c88565b155b9050919050565b6000806301ffc9a760e01b83604051602401610ca491906118a8565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505090506000808573ffffffffffffffffffffffffffffffffffffffff1661753084604051610d2e91906118ff565b6000604051808303818686fa925050503d8060008114610d6a576040519150601f19603f3d011682016040523d82523d6000602084013e610d6f565b606091505b5091509150602081511015610d8a5760009350505050610dad565b818015610da7575080806020019051810190610da69190611942565b5b93505050505b92915050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610df282610dc7565b9050919050565b610e0281610de7565b8114610e0d57600080fd5b50565b600081359050610e1f81610df9565b92915050565b600060208284031215610e3b57610e3a610dbd565b5b6000610e4984828501610e10565b91505092915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610ea582610e5c565b810181811067ffffffffffffffff82111715610ec457610ec3610e6d565b5b80604052505050565b6000610ed7610db3565b9050610ee38282610e9c565b919050565b600067ffffffffffffffff821115610f0357610f02610e6d565b5b610f0c82610e5c565b9050602081019050919050565b82818337600083830152505050565b6000610f3b610f3684610ee8565b610ecd565b905082815260208101848484011115610f5757610f56610e57565b5b610f62848285610f19565b509392505050565b600082601f830112610f7f57610f7e610e52565b5b8135610f8f848260208601610f28565b91505092915050565b600060208284031215610fae57610fad610dbd565b5b600082013567ffffffffffffffff811115610fcc57610fcb610dc2565b5b610fd884828501610f6a565b91505092915050565b6000819050919050565b610ff481610fe1565b8114610fff57600080fd5b50565b60008135905061101181610feb565b92915050565b60006020828403121561102d5761102c610dbd565b5b600061103b84828501611002565b91505092915050565b61104d81610de7565b82525050565b60006020820190506110686000830184611044565b92915050565b60008115159050919050565b6110838161106e565b82525050565b600060208201905061109e600083018461107a565b92915050565b600080fd5b600080fd5b60008083601f8401126110c4576110c3610e52565b5b8235905067ffffffffffffffff8111156110e1576110e06110a4565b5b6020830191508360018202830111156110fd576110fc6110a9565b5b9250929050565b6000806000806060858703121561111e5761111d610dbd565b5b600061112c87828801611002565b945050602061113d87828801610e10565b935050604085013567ffffffffffffffff81111561115e5761115d610dc2565b5b61116a878288016110ae565b925092505092959194509250565b600081519050919050565b600082825260208201905092915050565b60005b838110156111b2578082015181840152602081019050611197565b838111156111c1576000848401525b50505050565b60006111d282611178565b6111dc8185611183565b93506111ec818560208601611194565b6111f581610e5c565b840191505092915050565b6000602082019050818103600083015261121a81846111c7565b905092915050565b6000806040838503121561123957611238610dbd565b5b600061124785828601611002565b925050602061125885828601610e10565b9150509250929050565b61126b81610fe1565b82525050565b60006020820190506112866000830184611262565b92915050565b6000806000604084860312156112a5576112a4610dbd565b5b60006112b386828701611002565b935050602084013567ffffffffffffffff8111156112d4576112d3610dc2565b5b6112e0868287016110ae565b92509250509250925092565b60006112f782610dc7565b9050919050565b611307816112ec565b811461131257600080fd5b50565b600081519050611324816112fe565b92915050565b6000819050919050565b61133d8161132a565b811461134857600080fd5b50565b60008151905061135a81611334565b92915050565b60008060006060848603121561137957611378610dbd565b5b600061138786828701611315565b935050602061139886828701611315565b92505060406113a98682870161134b565b9150509250925092565b6000813590506113c281611334565b92915050565b6000602082840312156113de576113dd610dbd565b5b60006113ec848285016113b3565b91505092915050565b600082825260208201905092915050565b7f70726f766964656420746f6b656e41646472657373206973206e6f742077686960008201527f74656c6973746564000000000000000000000000000000000000000000000000602082015250565b60006114626028836113f5565b915061146d82611406565b604082019050919050565b6000602082019050818103600083015261149181611455565b9050919050565b6114a18161132a565b82525050565b60006020820190506114bc6000830184611498565b92915050565b600067ffffffffffffffff8211156114dd576114dc610e6d565b5b6114e682610e5c565b9050602081019050919050565b6000611506611501846114c2565b610ecd565b90508281526020810184848401111561152257611521610e57565b5b61152d848285611194565b509392505050565b600082601f83011261154a57611549610e52565b5b815161155a8482602086016114f3565b91505092915050565b60006020828403121561157957611578610dbd565b5b600082015167ffffffffffffffff81111561159757611596610dc2565b5b6115a384828501611535565b91505092915050565b600080604083850312156115c3576115c2610dbd565b5b60006115d1858286016113b3565b92505060206115e2858286016113b3565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006116268261132a565b91506116318361132a565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611666576116656115ec565b5b828201905092915050565b600080fd5b600080fd5b6000808585111561168f5761168e611671565b5b838611156116a05761169f611676565b5b6001850283019150848603905094509492505050565b7f73656e646572206d7573742062652062726964676520636f6e74726163740000600082015250565b60006116ec601e836113f5565b91506116f7826116b6565b602082019050919050565b6000602082019050818103600083015261171b816116df565b9050919050565b7f70726f766964656420636f6e7472616374206973206e6f742077686974656c6960008201527f7374656400000000000000000000000000000000000000000000000000000000602082015250565b600061177e6024836113f5565b915061178982611722565b604082019050919050565b600060208201905081810360008301526117ad81611771565b9050919050565b60006060820190506117c96000830186611044565b6117d66020830185611044565b6117e36040830184611498565b949350505050565b600081519050919050565b6000611801826117eb565b61180b81856113f5565b935061181b818560208601611194565b61182481610e5c565b840191505092915050565b60006060820190506118446000830186611044565b6118516020830185611498565b818103604083015261186381846117f6565b9050949350505050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6118a28161186d565b82525050565b60006020820190506118bd6000830184611899565b92915050565b600081905092915050565b60006118d982611178565b6118e381856118c3565b93506118f3818560208601611194565b80840191505092915050565b600061190b82846118ce565b915081905092915050565b61191f8161106e565b811461192a57600080fd5b50565b60008151905061193c81611916565b92915050565b60006020828403121561195857611957610dbd565b5b60006119668482850161192d565b9150509291505056fea26469706673582212208a470c41a24213cefaf601d9f0f7fb4a9782ed76c941b9d22ada0ce2d6d0ef8664736f6c634300080b0033",
}

// ERC721HandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC721HandlerMetaData.ABI instead.
var ERC721HandlerABI = ERC721HandlerMetaData.ABI

// ERC721HandlerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC721HandlerMetaData.Bin instead.
var ERC721HandlerBin = ERC721HandlerMetaData.Bin

// DeployERC721Handler deploys a new Ethereum contract, binding an instance of ERC721Handler to it.
func DeployERC721Handler(auth *bind.TransactOpts, backend bind.ContractBackend, bridgeAddress common.Address) (common.Address, *types.Transaction, *ERC721Handler, error) {
	parsed, err := ERC721HandlerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC721HandlerBin), backend, bridgeAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC721Handler{ERC721HandlerCaller: ERC721HandlerCaller{contract: contract}, ERC721HandlerTransactor: ERC721HandlerTransactor{contract: contract}, ERC721HandlerFilterer: ERC721HandlerFilterer{contract: contract}}, nil
}

// ERC721Handler is an auto generated Go binding around an Ethereum contract.
type ERC721Handler struct {
	ERC721HandlerCaller     // Read-only binding to the contract
	ERC721HandlerTransactor // Write-only binding to the contract
	ERC721HandlerFilterer   // Log filterer for contract events
}

// ERC721HandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721HandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721HandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721HandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721HandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721HandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721HandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721HandlerSession struct {
	Contract     *ERC721Handler    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721HandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721HandlerCallerSession struct {
	Contract *ERC721HandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ERC721HandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721HandlerTransactorSession struct {
	Contract     *ERC721HandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ERC721HandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721HandlerRaw struct {
	Contract *ERC721Handler // Generic contract binding to access the raw methods on
}

// ERC721HandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721HandlerCallerRaw struct {
	Contract *ERC721HandlerCaller // Generic read-only contract binding to access the raw methods on
}

// ERC721HandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721HandlerTransactorRaw struct {
	Contract *ERC721HandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721Handler creates a new instance of ERC721Handler, bound to a specific deployed contract.
func NewERC721Handler(address common.Address, backend bind.ContractBackend) (*ERC721Handler, error) {
	contract, err := bindERC721Handler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721Handler{ERC721HandlerCaller: ERC721HandlerCaller{contract: contract}, ERC721HandlerTransactor: ERC721HandlerTransactor{contract: contract}, ERC721HandlerFilterer: ERC721HandlerFilterer{contract: contract}}, nil
}

// NewERC721HandlerCaller creates a new read-only instance of ERC721Handler, bound to a specific deployed contract.
func NewERC721HandlerCaller(address common.Address, caller bind.ContractCaller) (*ERC721HandlerCaller, error) {
	contract, err := bindERC721Handler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721HandlerCaller{contract: contract}, nil
}

// NewERC721HandlerTransactor creates a new write-only instance of ERC721Handler, bound to a specific deployed contract.
func NewERC721HandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721HandlerTransactor, error) {
	contract, err := bindERC721Handler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721HandlerTransactor{contract: contract}, nil
}

// NewERC721HandlerFilterer creates a new log filterer instance of ERC721Handler, bound to a specific deployed contract.
func NewERC721HandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721HandlerFilterer, error) {
	contract, err := bindERC721Handler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721HandlerFilterer{contract: contract}, nil
}

// bindERC721Handler binds a generic wrapper to an already deployed contract.
func bindERC721Handler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721HandlerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Handler *ERC721HandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721Handler.Contract.ERC721HandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Handler *ERC721HandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Handler.Contract.ERC721HandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Handler *ERC721HandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Handler.Contract.ERC721HandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Handler *ERC721HandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721Handler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Handler *ERC721HandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Handler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Handler *ERC721HandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Handler.Contract.contract.Transact(opts, method, params...)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_ERC721Handler *ERC721HandlerCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC721Handler.contract.Call(opts, &out, "_bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_ERC721Handler *ERC721HandlerSession) BridgeAddress() (common.Address, error) {
	return _ERC721Handler.Contract.BridgeAddress(&_ERC721Handler.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_ERC721Handler *ERC721HandlerCallerSession) BridgeAddress() (common.Address, error) {
	return _ERC721Handler.Contract.BridgeAddress(&_ERC721Handler.CallOpts)
}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_ERC721Handler *ERC721HandlerCaller) BurnList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ERC721Handler.contract.Call(opts, &out, "_burnList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_ERC721Handler *ERC721HandlerSession) BurnList(arg0 common.Address) (bool, error) {
	return _ERC721Handler.Contract.BurnList(&_ERC721Handler.CallOpts, arg0)
}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_ERC721Handler *ERC721HandlerCallerSession) BurnList(arg0 common.Address) (bool, error) {
	return _ERC721Handler.Contract.BurnList(&_ERC721Handler.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_ERC721Handler *ERC721HandlerCaller) ContractWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ERC721Handler.contract.Call(opts, &out, "_contractWhitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_ERC721Handler *ERC721HandlerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _ERC721Handler.Contract.ContractWhitelist(&_ERC721Handler.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_ERC721Handler *ERC721HandlerCallerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _ERC721Handler.Contract.ContractWhitelist(&_ERC721Handler.CallOpts, arg0)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_ERC721Handler *ERC721HandlerCaller) ResourceIDToTokenContractAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ERC721Handler.contract.Call(opts, &out, "_resourceIDToTokenContractAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_ERC721Handler *ERC721HandlerSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _ERC721Handler.Contract.ResourceIDToTokenContractAddress(&_ERC721Handler.CallOpts, arg0)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_ERC721Handler *ERC721HandlerCallerSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _ERC721Handler.Contract.ResourceIDToTokenContractAddress(&_ERC721Handler.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_ERC721Handler *ERC721HandlerCaller) TokenContractAddressToResourceID(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _ERC721Handler.contract.Call(opts, &out, "_tokenContractAddressToResourceID", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_ERC721Handler *ERC721HandlerSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _ERC721Handler.Contract.TokenContractAddressToResourceID(&_ERC721Handler.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_ERC721Handler *ERC721HandlerCallerSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _ERC721Handler.Contract.TokenContractAddressToResourceID(&_ERC721Handler.CallOpts, arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0xb07e54bb.
//
// Solidity: function deposit(bytes32 resourceID, address depositer, bytes data) returns(bytes metaData)
func (_ERC721Handler *ERC721HandlerTransactor) Deposit(opts *bind.TransactOpts, resourceID [32]byte, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _ERC721Handler.contract.Transact(opts, "deposit", resourceID, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0xb07e54bb.
//
// Solidity: function deposit(bytes32 resourceID, address depositer, bytes data) returns(bytes metaData)
func (_ERC721Handler *ERC721HandlerSession) Deposit(resourceID [32]byte, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _ERC721Handler.Contract.Deposit(&_ERC721Handler.TransactOpts, resourceID, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0xb07e54bb.
//
// Solidity: function deposit(bytes32 resourceID, address depositer, bytes data) returns(bytes metaData)
func (_ERC721Handler *ERC721HandlerTransactorSession) Deposit(resourceID [32]byte, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _ERC721Handler.Contract.Deposit(&_ERC721Handler.TransactOpts, resourceID, depositer, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_ERC721Handler *ERC721HandlerTransactor) ExecuteProposal(opts *bind.TransactOpts, resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _ERC721Handler.contract.Transact(opts, "executeProposal", resourceID, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_ERC721Handler *ERC721HandlerSession) ExecuteProposal(resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _ERC721Handler.Contract.ExecuteProposal(&_ERC721Handler.TransactOpts, resourceID, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_ERC721Handler *ERC721HandlerTransactorSession) ExecuteProposal(resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _ERC721Handler.Contract.ExecuteProposal(&_ERC721Handler.TransactOpts, resourceID, data)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_ERC721Handler *ERC721HandlerTransactor) SetBurnable(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC721Handler.contract.Transact(opts, "setBurnable", contractAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_ERC721Handler *ERC721HandlerSession) SetBurnable(contractAddress common.Address) (*types.Transaction, error) {
	return _ERC721Handler.Contract.SetBurnable(&_ERC721Handler.TransactOpts, contractAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_ERC721Handler *ERC721HandlerTransactorSession) SetBurnable(contractAddress common.Address) (*types.Transaction, error) {
	return _ERC721Handler.Contract.SetBurnable(&_ERC721Handler.TransactOpts, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_ERC721Handler *ERC721HandlerTransactor) SetResource(opts *bind.TransactOpts, resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC721Handler.contract.Transact(opts, "setResource", resourceID, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_ERC721Handler *ERC721HandlerSession) SetResource(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC721Handler.Contract.SetResource(&_ERC721Handler.TransactOpts, resourceID, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_ERC721Handler *ERC721HandlerTransactorSession) SetResource(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC721Handler.Contract.SetResource(&_ERC721Handler.TransactOpts, resourceID, contractAddress)
}

// Withdraw is a paid mutator transaction binding the contract method 0x0968f264.
//
// Solidity: function withdraw(bytes data) returns()
func (_ERC721Handler *ERC721HandlerTransactor) Withdraw(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _ERC721Handler.contract.Transact(opts, "withdraw", data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x0968f264.
//
// Solidity: function withdraw(bytes data) returns()
func (_ERC721Handler *ERC721HandlerSession) Withdraw(data []byte) (*types.Transaction, error) {
	return _ERC721Handler.Contract.Withdraw(&_ERC721Handler.TransactOpts, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x0968f264.
//
// Solidity: function withdraw(bytes data) returns()
func (_ERC721Handler *ERC721HandlerTransactorSession) Withdraw(data []byte) (*types.Transaction, error) {
	return _ERC721Handler.Contract.Withdraw(&_ERC721Handler.TransactOpts, data)
}
