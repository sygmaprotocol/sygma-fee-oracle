// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ERC721Handler

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

// ERCHandlerHelpersDecimals is an auto generated low-level Go binding around an user-defined struct.
type ERCHandlerHelpersDecimals struct {
	IsSet            bool
	ExternalDecimals uint8
}

// ERC721HandlerABI is the input ABI used to generate the binding from.
const ERC721HandlerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToTokenContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_tokenContractAddressToTokenProperties\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isWhitelisted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isBurnable\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"externalDecimals\",\"type\":\"uint8\"}],\"internalType\":\"structERCHandlerHelpers.Decimals\",\"name\":\"decimals\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"metaData\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setBurnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"args\",\"type\":\"bytes\"}],\"name\":\"setResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC721HandlerBin is the compiled bytecode used for deploying new contracts.
var ERC721HandlerBin = "0x60a06040523480156200001157600080fd5b5060405162001cc938038062001cc98339818101604052810190620000379190620000de565b808073ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1681525050505062000110565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620000a68262000079565b9050919050565b620000b88162000099565b8114620000c457600080fd5b50565b600081519050620000d881620000ad565b92915050565b600060208284031215620000f757620000f662000074565b5b60006200010784828501620000c7565b91505092915050565b608051611b96620001336000396000818161026401526107c80152611b966000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c8063318c136e11610066578063318c136e1461011e578063ac607c211461013c578063b07e54bb1461016f578063e248cff21461019f578063fa8675b0146101bb57610093565b806307b7ed99146100985780630968f264146100b45780630a6d55d8146100d057806330f08abd14610100575b600080fd5b6100b260048036038101906100ad9190610f15565b6101d7565b005b6100ce60048036038101906100c99190611088565b6101eb565b005b6100ea60048036038101906100e59190611107565b61022a565b6040516100f79190611143565b60405180910390f35b61010861025d565b604051610115919061117a565b60405180910390f35b610126610262565b6040516101339190611143565b60405180910390f35b61015660048036038101906101519190610f15565b610286565b604051610166949392919061120c565b60405180910390f35b610189600480360381019061018491906112b1565b610312565b60405161019691906113ad565b60405180910390f35b6101b960048036038101906101b491906113cf565b61052e565b005b6101d560048036038101906101d091906112b1565b6107ae565b005b6101df6107c6565b6101e881610856565b50565b6101f36107c6565b60008060008380602001905181019061020c91906114a3565b80935081945082955050505061022483308484610942565b50505050565b60006020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b601281565b7f000000000000000000000000000000000000000000000000000000000000000081565b60016020528060005260406000206000915090508060000154908060010160009054906101000a900460ff16908060010160019054906101000a900460ff1690806002016040518060400160405290816000820160009054906101000a900460ff161515151581526020016000820160019054906101000a900460ff1660ff1660ff1681525050905084565b606061031c6107c6565b6000838381019061032d919061150b565b9050600080600088815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff166103f5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103ec906115bb565b60405180910390fd5b610425635b5e139f60e01b8273ffffffffffffffffffffffffffffffffffffffff166109bd90919063ffffffff16565b156104b25760008190508073ffffffffffffffffffffffffffffffffffffffff1663c87b56dd846040518263ffffffff1660e01b815260040161046891906115ea565b600060405180830381865afa158015610485573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906104ae91906116a6565b9350505b600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160019054906101000a900460ff1615610517576105128187846109e2565b610524565b61052381873085610b3f565b5b5050949350505050565b6105366107c6565b60008060606000806060878781019061054f91906116ef565b8096508197505050846040610564919061175e565b925087876040908592610579939291906117be565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505093508787849080926105d0939291906117be565b8101906105dd919061150b565b915087876020856105ee919061175e565b90846020876105fd919061175e565b610607919061175e565b92610614939291906117be565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050905060006020850151905060008060008c815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff16610728576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161071f906115bb565b60405180910390fd5b600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160019054906101000a900460ff16156107915761078c818360601c8a86610bba565b6107a1565b6107a081308460601c8b610942565b5b5050505050505050505050565b6107b66107c6565b6107c08484610c35565b50505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610854576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161084b90611845565b60405180910390fd5b565b600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff166108e5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108dc906118d7565b60405180910390fd5b60018060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160016101000a81548160ff02191690831515021790555050565b60008490508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8585856040518463ffffffff1660e01b8152600401610984939291906118f7565b600060405180830381600087803b15801561099e57600080fd5b505af11580156109b2573d6000803e3d6000fd5b505050505050505050565b60006109c883610d2b565b80156109da57506109d98383610d78565b5b905092915050565b60008390508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16636352211e846040518263ffffffff1660e01b8152600401610a3791906115ea565b602060405180830381865afa158015610a54573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a789190611943565b73ffffffffffffffffffffffffffffffffffffffff1614610ace576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ac5906119bc565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166342966c68836040518263ffffffff1660e01b8152600401610b0791906115ea565b600060405180830381600087803b158015610b2157600080fd5b505af1158015610b35573d6000803e3d6000fd5b5050505050505050565b60008490508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8585856040518463ffffffff1660e01b8152600401610b81939291906118f7565b600060405180830381600087803b158015610b9b57600080fd5b505af1158015610baf573d6000803e3d6000fd5b505050505050505050565b60008490508073ffffffffffffffffffffffffffffffffffffffff1663d3fc98648585856040518463ffffffff1660e01b8152600401610bfc93929190611a20565b600060405180830381600087803b158015610c1657600080fd5b505af1158015610c2a573d6000803e3d6000fd5b505050505050505050565b8060008084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000018190555060018060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160006101000a81548160ff0219169083151502179055505050565b6000610d57827f01ffc9a700000000000000000000000000000000000000000000000000000000610d78565b8015610d715750610d6f8263ffffffff60e01b610d78565b155b9050919050565b6000806301ffc9a760e01b83604051602401610d949190611a99565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505090506000808573ffffffffffffffffffffffffffffffffffffffff1661753084604051610e1e9190611af0565b6000604051808303818686fa925050503d8060008114610e5a576040519150601f19603f3d011682016040523d82523d6000602084013e610e5f565b606091505b5091509150602081511015610e7a5760009350505050610e9d565b818015610e97575080806020019051810190610e969190611b33565b5b93505050505b92915050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610ee282610eb7565b9050919050565b610ef281610ed7565b8114610efd57600080fd5b50565b600081359050610f0f81610ee9565b92915050565b600060208284031215610f2b57610f2a610ead565b5b6000610f3984828501610f00565b91505092915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610f9582610f4c565b810181811067ffffffffffffffff82111715610fb457610fb3610f5d565b5b80604052505050565b6000610fc7610ea3565b9050610fd38282610f8c565b919050565b600067ffffffffffffffff821115610ff357610ff2610f5d565b5b610ffc82610f4c565b9050602081019050919050565b82818337600083830152505050565b600061102b61102684610fd8565b610fbd565b90508281526020810184848401111561104757611046610f47565b5b611052848285611009565b509392505050565b600082601f83011261106f5761106e610f42565b5b813561107f848260208601611018565b91505092915050565b60006020828403121561109e5761109d610ead565b5b600082013567ffffffffffffffff8111156110bc576110bb610eb2565b5b6110c88482850161105a565b91505092915050565b6000819050919050565b6110e4816110d1565b81146110ef57600080fd5b50565b600081359050611101816110db565b92915050565b60006020828403121561111d5761111c610ead565b5b600061112b848285016110f2565b91505092915050565b61113d81610ed7565b82525050565b60006020820190506111586000830184611134565b92915050565b600060ff82169050919050565b6111748161115e565b82525050565b600060208201905061118f600083018461116b565b92915050565b61119e816110d1565b82525050565b60008115159050919050565b6111b9816111a4565b82525050565b6111c8816111a4565b82525050565b6111d78161115e565b82525050565b6040820160008201516111f360008501826111bf565b50602082015161120660208501826111ce565b50505050565b600060a0820190506112216000830187611195565b61122e60208301866111b0565b61123b60408301856111b0565b61124860608301846111dd565b95945050505050565b600080fd5b600080fd5b60008083601f84011261127157611270610f42565b5b8235905067ffffffffffffffff81111561128e5761128d611251565b5b6020830191508360018202830111156112aa576112a9611256565b5b9250929050565b600080600080606085870312156112cb576112ca610ead565b5b60006112d9878288016110f2565b94505060206112ea87828801610f00565b935050604085013567ffffffffffffffff81111561130b5761130a610eb2565b5b6113178782880161125b565b925092505092959194509250565b600081519050919050565b600082825260208201905092915050565b60005b8381101561135f578082015181840152602081019050611344565b8381111561136e576000848401525b50505050565b600061137f82611325565b6113898185611330565b9350611399818560208601611341565b6113a281610f4c565b840191505092915050565b600060208201905081810360008301526113c78184611374565b905092915050565b6000806000604084860312156113e8576113e7610ead565b5b60006113f6868287016110f2565b935050602084013567ffffffffffffffff81111561141757611416610eb2565b5b6114238682870161125b565b92509250509250925092565b600061143a82610eb7565b9050919050565b61144a8161142f565b811461145557600080fd5b50565b60008151905061146781611441565b92915050565b6000819050919050565b6114808161146d565b811461148b57600080fd5b50565b60008151905061149d81611477565b92915050565b6000806000606084860312156114bc576114bb610ead565b5b60006114ca86828701611458565b93505060206114db86828701611458565b92505060406114ec8682870161148e565b9150509250925092565b60008135905061150581611477565b92915050565b60006020828403121561152157611520610ead565b5b600061152f848285016114f6565b91505092915050565b600082825260208201905092915050565b7f70726f766964656420746f6b656e41646472657373206973206e6f742077686960008201527f74656c6973746564000000000000000000000000000000000000000000000000602082015250565b60006115a5602883611538565b91506115b082611549565b604082019050919050565b600060208201905081810360008301526115d481611598565b9050919050565b6115e48161146d565b82525050565b60006020820190506115ff60008301846115db565b92915050565b600067ffffffffffffffff8211156116205761161f610f5d565b5b61162982610f4c565b9050602081019050919050565b600061164961164484611605565b610fbd565b90508281526020810184848401111561166557611664610f47565b5b611670848285611341565b509392505050565b600082601f83011261168d5761168c610f42565b5b815161169d848260208601611636565b91505092915050565b6000602082840312156116bc576116bb610ead565b5b600082015167ffffffffffffffff8111156116da576116d9610eb2565b5b6116e684828501611678565b91505092915050565b6000806040838503121561170657611705610ead565b5b6000611714858286016114f6565b9250506020611725858286016114f6565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006117698261146d565b91506117748361146d565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156117a9576117a861172f565b5b828201905092915050565b600080fd5b600080fd5b600080858511156117d2576117d16117b4565b5b838611156117e3576117e26117b9565b5b6001850283019150848603905094509492505050565b7f73656e646572206d7573742062652062726964676520636f6e74726163740000600082015250565b600061182f601e83611538565b915061183a826117f9565b602082019050919050565b6000602082019050818103600083015261185e81611822565b9050919050565b7f70726f766964656420636f6e7472616374206973206e6f742077686974656c6960008201527f7374656400000000000000000000000000000000000000000000000000000000602082015250565b60006118c1602483611538565b91506118cc82611865565b604082019050919050565b600060208201905081810360008301526118f0816118b4565b9050919050565b600060608201905061190c6000830186611134565b6119196020830185611134565b61192660408301846115db565b949350505050565b60008151905061193d81610ee9565b92915050565b60006020828403121561195957611958610ead565b5b60006119678482850161192e565b91505092915050565b7f4275726e206e6f742066726f6d206f776e657200000000000000000000000000600082015250565b60006119a6601383611538565b91506119b182611970565b602082019050919050565b600060208201905081810360008301526119d581611999565b9050919050565b600081519050919050565b60006119f2826119dc565b6119fc8185611538565b9350611a0c818560208601611341565b611a1581610f4c565b840191505092915050565b6000606082019050611a356000830186611134565b611a4260208301856115db565b8181036040830152611a5481846119e7565b9050949350505050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b611a9381611a5e565b82525050565b6000602082019050611aae6000830184611a8a565b92915050565b600081905092915050565b6000611aca82611325565b611ad48185611ab4565b9350611ae4818560208601611341565b80840191505092915050565b6000611afc8284611abf565b915081905092915050565b611b10816111a4565b8114611b1b57600080fd5b50565b600081519050611b2d81611b07565b92915050565b600060208284031215611b4957611b48610ead565b5b6000611b5784828501611b1e565b9150509291505056fea26469706673582212202d6448c150b443cfba381bf4c0a15f1105817a414658a2cd534514eb9ebe739564736f6c634300080b0033"

// DeployERC721Handler deploys a new Ethereum contract, binding an instance of ERC721Handler to it.
func DeployERC721Handler(auth *bind.TransactOpts, backend bind.ContractBackend, bridgeAddress common.Address) (common.Address, *types.Transaction, *ERC721Handler, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721HandlerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC721HandlerBin), backend, bridgeAddress)
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

// TokenContractAddressToTokenProperties is a free data retrieval call binding the contract method 0xac607c21.
//
// Solidity: function _tokenContractAddressToTokenProperties(address ) view returns(bytes32 resourceID, bool isWhitelisted, bool isBurnable, (bool,uint8) decimals)
func (_ERC721Handler *ERC721HandlerCaller) TokenContractAddressToTokenProperties(opts *bind.CallOpts, arg0 common.Address) (struct {
	ResourceID    [32]byte
	IsWhitelisted bool
	IsBurnable    bool
	Decimals      ERCHandlerHelpersDecimals
}, error) {
	var out []interface{}
	err := _ERC721Handler.contract.Call(opts, &out, "_tokenContractAddressToTokenProperties", arg0)

	outstruct := new(struct {
		ResourceID    [32]byte
		IsWhitelisted bool
		IsBurnable    bool
		Decimals      ERCHandlerHelpersDecimals
	})

	outstruct.ResourceID = out[0].([32]byte)
	outstruct.IsWhitelisted = out[1].(bool)
	outstruct.IsBurnable = out[2].(bool)
	outstruct.Decimals = out[3].(ERCHandlerHelpersDecimals)

	return *outstruct, err

}

// TokenContractAddressToTokenProperties is a free data retrieval call binding the contract method 0xac607c21.
//
// Solidity: function _tokenContractAddressToTokenProperties(address ) view returns(bytes32 resourceID, bool isWhitelisted, bool isBurnable, (bool,uint8) decimals)
func (_ERC721Handler *ERC721HandlerSession) TokenContractAddressToTokenProperties(arg0 common.Address) (struct {
	ResourceID    [32]byte
	IsWhitelisted bool
	IsBurnable    bool
	Decimals      ERCHandlerHelpersDecimals
}, error) {
	return _ERC721Handler.Contract.TokenContractAddressToTokenProperties(&_ERC721Handler.CallOpts, arg0)
}

// TokenContractAddressToTokenProperties is a free data retrieval call binding the contract method 0xac607c21.
//
// Solidity: function _tokenContractAddressToTokenProperties(address ) view returns(bytes32 resourceID, bool isWhitelisted, bool isBurnable, (bool,uint8) decimals)
func (_ERC721Handler *ERC721HandlerCallerSession) TokenContractAddressToTokenProperties(arg0 common.Address) (struct {
	ResourceID    [32]byte
	IsWhitelisted bool
	IsBurnable    bool
	Decimals      ERCHandlerHelpersDecimals
}, error) {
	return _ERC721Handler.Contract.TokenContractAddressToTokenProperties(&_ERC721Handler.CallOpts, arg0)
}

// DefaultDecimals is a free data retrieval call binding the contract method 0x30f08abd.
//
// Solidity: function defaultDecimals() view returns(uint8)
func (_ERC721Handler *ERC721HandlerCaller) DefaultDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC721Handler.contract.Call(opts, &out, "defaultDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DefaultDecimals is a free data retrieval call binding the contract method 0x30f08abd.
//
// Solidity: function defaultDecimals() view returns(uint8)
func (_ERC721Handler *ERC721HandlerSession) DefaultDecimals() (uint8, error) {
	return _ERC721Handler.Contract.DefaultDecimals(&_ERC721Handler.CallOpts)
}

// DefaultDecimals is a free data retrieval call binding the contract method 0x30f08abd.
//
// Solidity: function defaultDecimals() view returns(uint8)
func (_ERC721Handler *ERC721HandlerCallerSession) DefaultDecimals() (uint8, error) {
	return _ERC721Handler.Contract.DefaultDecimals(&_ERC721Handler.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xb07e54bb.
//
// Solidity: function deposit(bytes32 resourceID, address depositor, bytes data) returns(bytes metaData)
func (_ERC721Handler *ERC721HandlerTransactor) Deposit(opts *bind.TransactOpts, resourceID [32]byte, depositor common.Address, data []byte) (*types.Transaction, error) {
	return _ERC721Handler.contract.Transact(opts, "deposit", resourceID, depositor, data)
}

// Deposit is a paid mutator transaction binding the contract method 0xb07e54bb.
//
// Solidity: function deposit(bytes32 resourceID, address depositor, bytes data) returns(bytes metaData)
func (_ERC721Handler *ERC721HandlerSession) Deposit(resourceID [32]byte, depositor common.Address, data []byte) (*types.Transaction, error) {
	return _ERC721Handler.Contract.Deposit(&_ERC721Handler.TransactOpts, resourceID, depositor, data)
}

// Deposit is a paid mutator transaction binding the contract method 0xb07e54bb.
//
// Solidity: function deposit(bytes32 resourceID, address depositor, bytes data) returns(bytes metaData)
func (_ERC721Handler *ERC721HandlerTransactorSession) Deposit(resourceID [32]byte, depositor common.Address, data []byte) (*types.Transaction, error) {
	return _ERC721Handler.Contract.Deposit(&_ERC721Handler.TransactOpts, resourceID, depositor, data)
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

// SetResource is a paid mutator transaction binding the contract method 0xfa8675b0.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress, bytes args) returns()
func (_ERC721Handler *ERC721HandlerTransactor) SetResource(opts *bind.TransactOpts, resourceID [32]byte, contractAddress common.Address, args []byte) (*types.Transaction, error) {
	return _ERC721Handler.contract.Transact(opts, "setResource", resourceID, contractAddress, args)
}

// SetResource is a paid mutator transaction binding the contract method 0xfa8675b0.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress, bytes args) returns()
func (_ERC721Handler *ERC721HandlerSession) SetResource(resourceID [32]byte, contractAddress common.Address, args []byte) (*types.Transaction, error) {
	return _ERC721Handler.Contract.SetResource(&_ERC721Handler.TransactOpts, resourceID, contractAddress, args)
}

// SetResource is a paid mutator transaction binding the contract method 0xfa8675b0.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress, bytes args) returns()
func (_ERC721Handler *ERC721HandlerTransactorSession) SetResource(resourceID [32]byte, contractAddress common.Address, args []byte) (*types.Transaction, error) {
	return _ERC721Handler.Contract.SetResource(&_ERC721Handler.TransactOpts, resourceID, contractAddress, args)
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
