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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_burnList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToTokenContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_tokenContractAddressToResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"metaData\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setBurnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"args\",\"type\":\"bytes\"}],\"name\":\"setResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b5060405162001c3138038062001c318339818101604052810190620000379190620000de565b808073ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1681525050505062000110565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620000a68262000079565b9050919050565b620000b88162000099565b8114620000c457600080fd5b50565b600081519050620000d881620000ad565b92915050565b600060208284031215620000f757620000f662000074565b5b60006200010784828501620000c7565b91505092915050565b608051611afe62000133600039600081816102a901526107cd0152611afe6000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c80637f79bea8116100665780637f79bea814610159578063b07e54bb14610189578063c8ba6c87146101b9578063e248cff2146101e9578063fa8675b0146102055761009e565b806307b7ed99146100a35780630968f264146100bf5780630a6d55d8146100db578063318c136e1461010b5780636a70d08114610129575b600080fd5b6100bd60048036038101906100b89190610f10565b610221565b005b6100d960048036038101906100d49190611083565b610235565b005b6100f560048036038101906100f09190611102565b610274565b604051610102919061113e565b60405180910390f35b6101136102a7565b604051610120919061113e565b60405180910390f35b610143600480360381019061013e9190610f10565b6102cb565b6040516101509190611174565b60405180910390f35b610173600480360381019061016e9190610f10565b6102eb565b6040516101809190611174565b60405180910390f35b6101a3600480360381019061019e91906111ef565b61030b565b6040516101b091906112eb565b60405180910390f35b6101d360048036038101906101ce9190610f10565b610521565b6040516101e0919061131c565b60405180910390f35b61020360048036038101906101fe9190611337565b610539565b005b61021f600480360381019061021a91906111ef565b6107b3565b005b6102296107cb565b6102328161085b565b50565b61023d6107cb565b600080600083806020019051810190610256919061140b565b80935081945082955050505061026e83308484610942565b50505050565b60006020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b7f000000000000000000000000000000000000000000000000000000000000000081565b60036020528060005260406000206000915054906101000a900460ff1681565b60026020528060005260406000206000915054906101000a900460ff1681565b60606103156107cb565b600083838101906103269190611473565b9050600080600088815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600260008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166103eb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103e290611523565b60405180910390fd5b61041b635b5e139f60e01b8273ffffffffffffffffffffffffffffffffffffffff166109bd90919063ffffffff16565b156104a85760008190508073ffffffffffffffffffffffffffffffffffffffff1663c87b56dd846040518263ffffffff1660e01b815260040161045e9190611552565b600060405180830381865afa15801561047b573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906104a4919061160e565b9350505b600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161561050a576105058187846109e2565b610517565b61051681873085610b3f565b5b5050949350505050565b60016020528060005260406000206000915090505481565b6105416107cb565b60008060606000806060878781019061055a9190611657565b809650819750505084604061056f91906116c6565b92508787604090859261058493929190611726565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505093508787849080926105db93929190611726565b8101906105e89190611473565b915087876020856105f991906116c6565b908460208761060891906116c6565b61061291906116c6565b9261061f93929190611726565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050905060006020850151905060008060008c815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600260008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610730576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161072790611523565b60405180910390fd5b600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161561079657610791818360601c8a86610bba565b6107a6565b6107a581308460601c8b610942565b5b5050505050505050505050565b6107bb6107cb565b6107c58484610c35565b50505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610859576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610850906117ad565b60405180910390fd5b565b600260008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166108e7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108de9061183f565b60405180910390fd5b6001600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b60008490508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8585856040518463ffffffff1660e01b81526004016109849392919061185f565b600060405180830381600087803b15801561099e57600080fd5b505af11580156109b2573d6000803e3d6000fd5b505050505050505050565b60006109c883610d26565b80156109da57506109d98383610d73565b5b905092915050565b60008390508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16636352211e846040518263ffffffff1660e01b8152600401610a379190611552565b602060405180830381865afa158015610a54573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a7891906118ab565b73ffffffffffffffffffffffffffffffffffffffff1614610ace576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ac590611924565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166342966c68836040518263ffffffff1660e01b8152600401610b079190611552565b600060405180830381600087803b158015610b2157600080fd5b505af1158015610b35573d6000803e3d6000fd5b5050505050505050565b60008490508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8585856040518463ffffffff1660e01b8152600401610b819392919061185f565b600060405180830381600087803b158015610b9b57600080fd5b505af1158015610baf573d6000803e3d6000fd5b505050505050505050565b60008490508073ffffffffffffffffffffffffffffffffffffffff1663d3fc98648585856040518463ffffffff1660e01b8152600401610bfc93929190611988565b600060405180830381600087803b158015610c1657600080fd5b505af1158015610c2a573d6000803e3d6000fd5b505050505050505050565b8060008084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b6000610d52827f01ffc9a700000000000000000000000000000000000000000000000000000000610d73565b8015610d6c5750610d6a8263ffffffff60e01b610d73565b155b9050919050565b6000806301ffc9a760e01b83604051602401610d8f9190611a01565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505090506000808573ffffffffffffffffffffffffffffffffffffffff1661753084604051610e199190611a58565b6000604051808303818686fa925050503d8060008114610e55576040519150601f19603f3d011682016040523d82523d6000602084013e610e5a565b606091505b5091509150602081511015610e755760009350505050610e98565b818015610e92575080806020019051810190610e919190611a9b565b5b93505050505b92915050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610edd82610eb2565b9050919050565b610eed81610ed2565b8114610ef857600080fd5b50565b600081359050610f0a81610ee4565b92915050565b600060208284031215610f2657610f25610ea8565b5b6000610f3484828501610efb565b91505092915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610f9082610f47565b810181811067ffffffffffffffff82111715610faf57610fae610f58565b5b80604052505050565b6000610fc2610e9e565b9050610fce8282610f87565b919050565b600067ffffffffffffffff821115610fee57610fed610f58565b5b610ff782610f47565b9050602081019050919050565b82818337600083830152505050565b600061102661102184610fd3565b610fb8565b90508281526020810184848401111561104257611041610f42565b5b61104d848285611004565b509392505050565b600082601f83011261106a57611069610f3d565b5b813561107a848260208601611013565b91505092915050565b60006020828403121561109957611098610ea8565b5b600082013567ffffffffffffffff8111156110b7576110b6610ead565b5b6110c384828501611055565b91505092915050565b6000819050919050565b6110df816110cc565b81146110ea57600080fd5b50565b6000813590506110fc816110d6565b92915050565b60006020828403121561111857611117610ea8565b5b6000611126848285016110ed565b91505092915050565b61113881610ed2565b82525050565b6000602082019050611153600083018461112f565b92915050565b60008115159050919050565b61116e81611159565b82525050565b60006020820190506111896000830184611165565b92915050565b600080fd5b600080fd5b60008083601f8401126111af576111ae610f3d565b5b8235905067ffffffffffffffff8111156111cc576111cb61118f565b5b6020830191508360018202830111156111e8576111e7611194565b5b9250929050565b6000806000806060858703121561120957611208610ea8565b5b6000611217878288016110ed565b945050602061122887828801610efb565b935050604085013567ffffffffffffffff81111561124957611248610ead565b5b61125587828801611199565b925092505092959194509250565b600081519050919050565b600082825260208201905092915050565b60005b8381101561129d578082015181840152602081019050611282565b838111156112ac576000848401525b50505050565b60006112bd82611263565b6112c7818561126e565b93506112d781856020860161127f565b6112e081610f47565b840191505092915050565b6000602082019050818103600083015261130581846112b2565b905092915050565b611316816110cc565b82525050565b6000602082019050611331600083018461130d565b92915050565b6000806000604084860312156113505761134f610ea8565b5b600061135e868287016110ed565b935050602084013567ffffffffffffffff81111561137f5761137e610ead565b5b61138b86828701611199565b92509250509250925092565b60006113a282610eb2565b9050919050565b6113b281611397565b81146113bd57600080fd5b50565b6000815190506113cf816113a9565b92915050565b6000819050919050565b6113e8816113d5565b81146113f357600080fd5b50565b600081519050611405816113df565b92915050565b60008060006060848603121561142457611423610ea8565b5b6000611432868287016113c0565b9350506020611443868287016113c0565b9250506040611454868287016113f6565b9150509250925092565b60008135905061146d816113df565b92915050565b60006020828403121561148957611488610ea8565b5b60006114978482850161145e565b91505092915050565b600082825260208201905092915050565b7f70726f766964656420746f6b656e41646472657373206973206e6f742077686960008201527f74656c6973746564000000000000000000000000000000000000000000000000602082015250565b600061150d6028836114a0565b9150611518826114b1565b604082019050919050565b6000602082019050818103600083015261153c81611500565b9050919050565b61154c816113d5565b82525050565b60006020820190506115676000830184611543565b92915050565b600067ffffffffffffffff82111561158857611587610f58565b5b61159182610f47565b9050602081019050919050565b60006115b16115ac8461156d565b610fb8565b9050828152602081018484840111156115cd576115cc610f42565b5b6115d884828561127f565b509392505050565b600082601f8301126115f5576115f4610f3d565b5b815161160584826020860161159e565b91505092915050565b60006020828403121561162457611623610ea8565b5b600082015167ffffffffffffffff81111561164257611641610ead565b5b61164e848285016115e0565b91505092915050565b6000806040838503121561166e5761166d610ea8565b5b600061167c8582860161145e565b925050602061168d8582860161145e565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006116d1826113d5565b91506116dc836113d5565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561171157611710611697565b5b828201905092915050565b600080fd5b600080fd5b6000808585111561173a5761173961171c565b5b8386111561174b5761174a611721565b5b6001850283019150848603905094509492505050565b7f73656e646572206d7573742062652062726964676520636f6e74726163740000600082015250565b6000611797601e836114a0565b91506117a282611761565b602082019050919050565b600060208201905081810360008301526117c68161178a565b9050919050565b7f70726f766964656420636f6e7472616374206973206e6f742077686974656c6960008201527f7374656400000000000000000000000000000000000000000000000000000000602082015250565b60006118296024836114a0565b9150611834826117cd565b604082019050919050565b600060208201905081810360008301526118588161181c565b9050919050565b6000606082019050611874600083018661112f565b611881602083018561112f565b61188e6040830184611543565b949350505050565b6000815190506118a581610ee4565b92915050565b6000602082840312156118c1576118c0610ea8565b5b60006118cf84828501611896565b91505092915050565b7f4275726e206e6f742066726f6d206f776e657200000000000000000000000000600082015250565b600061190e6013836114a0565b9150611919826118d8565b602082019050919050565b6000602082019050818103600083015261193d81611901565b9050919050565b600081519050919050565b600061195a82611944565b61196481856114a0565b935061197481856020860161127f565b61197d81610f47565b840191505092915050565b600060608201905061199d600083018661112f565b6119aa6020830185611543565b81810360408301526119bc818461194f565b9050949350505050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6119fb816119c6565b82525050565b6000602082019050611a1660008301846119f2565b92915050565b600081905092915050565b6000611a3282611263565b611a3c8185611a1c565b9350611a4c81856020860161127f565b80840191505092915050565b6000611a648284611a27565b915081905092915050565b611a7881611159565b8114611a8357600080fd5b50565b600081519050611a9581611a6f565b92915050565b600060208284031215611ab157611ab0610ea8565b5b6000611abf84828501611a86565b9150509291505056fea2646970667358221220289acf2b07a124596ffac62c6305b8a518d7a8cf42513ee1b92bd1dbe73d018364736f6c634300080b0033",
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
