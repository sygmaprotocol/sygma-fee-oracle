// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ERC20Handler

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

// ERC20HandlerABI is the input ABI used to generate the binding from.
const ERC20HandlerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToTokenContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_tokenContractAddressToTokenProperties\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isWhitelisted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isBurnable\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"externalDecimals\",\"type\":\"uint8\"}],\"internalType\":\"structERCHandlerHelpers.Decimals\",\"name\":\"decimals\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setBurnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"args\",\"type\":\"bytes\"}],\"name\":\"setResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20HandlerBin is the compiled bytecode used for deploying new contracts.
var ERC20HandlerBin = "0x60a06040523480156200001157600080fd5b50604051620020a1380380620020a18339818101604052810190620000379190620000de565b808073ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1681525050505062000110565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620000a68262000079565b9050919050565b620000b88162000099565b8114620000c457600080fd5b50565b600081519050620000d881620000ad565b92915050565b600060208284031215620000f757620000f662000074565b5b60006200010784828501620000c7565b91505092915050565b608051611f6e620001336000396000818161026301526106ca0152611f6e6000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c8063318c136e11610066578063318c136e1461011e578063ac607c211461013c578063b07e54bb1461016f578063e248cff21461019f578063fa8675b0146101bb57610093565b806307b7ed99146100985780630968f264146100b45780630a6d55d8146100d057806330f08abd14610100575b600080fd5b6100b260048036038101906100ad91906110c0565b6101d7565b005b6100ce60048036038101906100c99190611233565b6101eb565b005b6100ea60048036038101906100e591906112b2565b610229565b6040516100f791906112ee565b60405180910390f35b61010861025c565b6040516101159190611325565b60405180910390f35b610126610261565b60405161013391906112ee565b60405180910390f35b610156600480360381019061015191906110c0565b610285565b60405161016694939291906113b7565b60405180910390f35b6101896004803603810190610184919061145c565b610311565b6040516101969190611558565b60405180910390f35b6101b960048036038101906101b4919061157a565b61049b565b005b6101d560048036038101906101d0919061145c565b610684565b005b6101df6106c8565b6101e881610758565b50565b6101f36106c8565b60008060008380602001905181019061020c919061164e565b809350819450829550505050610223838383610844565b50505050565b60006020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b601281565b7f000000000000000000000000000000000000000000000000000000000000000081565b60016020528060005260406000206000915090508060000154908060010160009054906101000a900460ff16908060010160019054906101000a900460ff1690806002016040518060400160405290816000820160009054906101000a900460ff161515151581526020016000820160019054906101000a900460ff1660ff1660ff1681525050905084565b606061031b6106c8565b6000838381019061032c91906116b6565b9050600080600088815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff166103f4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103eb90611766565b60405180910390fd5b600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160019054906101000a900460ff16156104595761045481878461085a565b610466565b610465818730856108d2565b5b61047081836108ea565b60405160200161048091906117c2565b60405160208183030381529060405292505050949350505050565b6104a36106c8565b600080606084848101906104b791906117d9565b809350819450505084846040908460406104d19190611848565b926104de939291906118a8565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050905060008060008089815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060208301519150600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff166105f1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105e890611766565b60405180910390fd5b600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160019054906101000a900460ff16156106625761065d818360601c6106588489610a30565b610b40565b61067a565b610679818360601c6106748489610a30565b610844565b5b5050505050505050565b61068c6106c8565b6106968484610bb8565b60008282905011156106c25760008282906106b19190611927565b60f81c90506106c08482610cae565b505b50505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610756576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161074d906119d2565b60405180910390fd5b565b600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff166107e7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107de90611a64565b60405180910390fd5b60018060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160016101000a81548160ff02191690831515021790555050565b6000839050610854818484610de2565b50505050565b60008390508073ffffffffffffffffffffffffffffffffffffffff166379cc679084846040518363ffffffff1660e01b815260040161089a929190611a93565b600060405180830381600087803b1580156108b457600080fd5b505af11580156108c8573d6000803e3d6000fd5b5050505050505050565b60008490506108e381858585610e68565b5050505050565b60606000600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206002016040518060400160405290816000820160009054906101000a900460ff161515151581526020016000820160019054906101000a900460ff1660ff1660ff1681525050905060008160000151610999576040518060200160405280600081525092505050610a2a565b601260ff16826020015160ff16106109da57601282602001516109bc9190611abc565b600a6109c89190611c23565b846109d39190611c9d565b9050610a05565b816020015160126109eb9190611abc565b600a6109f79190611c23565b84610a029190611cce565b90505b80604051602001610a169190611d49565b604051602081830303815290604052925050505b92915050565b600080600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206002016040518060400160405290816000820160009054906101000a900460ff161515151581526020016000820160019054906101000a900460ff1660ff1660ff168152505090508060000151610acc5782915050610b3a565b601260ff16816020015160ff1610610b0e5760128160200151610aef9190611abc565b600a610afb9190611c23565b83610b069190611cce565b915050610b3a565b80602001516012610b1f9190611abc565b600a610b2b9190611c23565b83610b369190611c9d565b9150505b92915050565b60008390508073ffffffffffffffffffffffffffffffffffffffff166340c10f1984846040518363ffffffff1660e01b8152600401610b80929190611a93565b600060405180830381600087803b158015610b9a57600080fd5b505af1158015610bae573d6000803e3d6000fd5b5050505050505050565b8060008084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000018190555060018060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160006101000a81548160ff0219169083151502179055505050565b600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff16610d3d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d3490611a64565b60405180910390fd5b60405180604001604052806001151581526020018260ff16815250600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a81548160ff021916908360ff1602179055509050505050565b610e638363a9059cbb60e01b8484604051602401610e01929190611a93565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610ef1565b505050565b610eeb846323b872dd60e01b858585604051602401610e8993929190611d64565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610ef1565b50505050565b6000823b905060008111610f3a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f3190611de7565b60405180910390fd5b6000808473ffffffffffffffffffffffffffffffffffffffff1684604051610f6291906117c2565b6000604051808303816000865af19150503d8060008114610f9f576040519150601f19603f3d011682016040523d82523d6000602084013e610fa4565b606091505b509150915081610fe9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fe090611e53565b60405180910390fd5b60008151111561104757808060200190518101906110079190611e9f565b611046576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161103d90611f18565b60405180910390fd5b5b5050505050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061108d82611062565b9050919050565b61109d81611082565b81146110a857600080fd5b50565b6000813590506110ba81611094565b92915050565b6000602082840312156110d6576110d5611058565b5b60006110e4848285016110ab565b91505092915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b611140826110f7565b810181811067ffffffffffffffff8211171561115f5761115e611108565b5b80604052505050565b600061117261104e565b905061117e8282611137565b919050565b600067ffffffffffffffff82111561119e5761119d611108565b5b6111a7826110f7565b9050602081019050919050565b82818337600083830152505050565b60006111d66111d184611183565b611168565b9050828152602081018484840111156111f2576111f16110f2565b5b6111fd8482856111b4565b509392505050565b600082601f83011261121a576112196110ed565b5b813561122a8482602086016111c3565b91505092915050565b60006020828403121561124957611248611058565b5b600082013567ffffffffffffffff8111156112675761126661105d565b5b61127384828501611205565b91505092915050565b6000819050919050565b61128f8161127c565b811461129a57600080fd5b50565b6000813590506112ac81611286565b92915050565b6000602082840312156112c8576112c7611058565b5b60006112d68482850161129d565b91505092915050565b6112e881611082565b82525050565b600060208201905061130360008301846112df565b92915050565b600060ff82169050919050565b61131f81611309565b82525050565b600060208201905061133a6000830184611316565b92915050565b6113498161127c565b82525050565b60008115159050919050565b6113648161134f565b82525050565b6113738161134f565b82525050565b61138281611309565b82525050565b60408201600082015161139e600085018261136a565b5060208201516113b16020850182611379565b50505050565b600060a0820190506113cc6000830187611340565b6113d9602083018661135b565b6113e6604083018561135b565b6113f36060830184611388565b95945050505050565b600080fd5b600080fd5b60008083601f84011261141c5761141b6110ed565b5b8235905067ffffffffffffffff811115611439576114386113fc565b5b60208301915083600182028301111561145557611454611401565b5b9250929050565b6000806000806060858703121561147657611475611058565b5b60006114848782880161129d565b9450506020611495878288016110ab565b935050604085013567ffffffffffffffff8111156114b6576114b561105d565b5b6114c287828801611406565b925092505092959194509250565b600081519050919050565b600082825260208201905092915050565b60005b8381101561150a5780820151818401526020810190506114ef565b83811115611519576000848401525b50505050565b600061152a826114d0565b61153481856114db565b93506115448185602086016114ec565b61154d816110f7565b840191505092915050565b60006020820190508181036000830152611572818461151f565b905092915050565b60008060006040848603121561159357611592611058565b5b60006115a18682870161129d565b935050602084013567ffffffffffffffff8111156115c2576115c161105d565b5b6115ce86828701611406565b92509250509250925092565b60006115e582611062565b9050919050565b6115f5816115da565b811461160057600080fd5b50565b600081519050611612816115ec565b92915050565b6000819050919050565b61162b81611618565b811461163657600080fd5b50565b60008151905061164881611622565b92915050565b60008060006060848603121561166757611666611058565b5b600061167586828701611603565b935050602061168686828701611603565b925050604061169786828701611639565b9150509250925092565b6000813590506116b081611622565b92915050565b6000602082840312156116cc576116cb611058565b5b60006116da848285016116a1565b91505092915050565b600082825260208201905092915050565b7f70726f766964656420746f6b656e41646472657373206973206e6f742077686960008201527f74656c6973746564000000000000000000000000000000000000000000000000602082015250565b60006117506028836116e3565b915061175b826116f4565b604082019050919050565b6000602082019050818103600083015261177f81611743565b9050919050565b600081905092915050565b600061179c826114d0565b6117a68185611786565b93506117b68185602086016114ec565b80840191505092915050565b60006117ce8284611791565b915081905092915050565b600080604083850312156117f0576117ef611058565b5b60006117fe858286016116a1565b925050602061180f858286016116a1565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061185382611618565b915061185e83611618565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561189357611892611819565b5b828201905092915050565b600080fd5b600080fd5b600080858511156118bc576118bb61189e565b5b838611156118cd576118cc6118a3565b5b6001850283019150848603905094509492505050565b600082905092915050565b60007fff0000000000000000000000000000000000000000000000000000000000000082169050919050565b600082821b905092915050565b600061193383836118e3565b8261193e81356118ee565b9250600182101561197e576119797fff000000000000000000000000000000000000000000000000000000000000008360010360080261191a565b831692505b505092915050565b7f73656e646572206d7573742062652062726964676520636f6e74726163740000600082015250565b60006119bc601e836116e3565b91506119c782611986565b602082019050919050565b600060208201905081810360008301526119eb816119af565b9050919050565b7f70726f766964656420636f6e7472616374206973206e6f742077686974656c6960008201527f7374656400000000000000000000000000000000000000000000000000000000602082015250565b6000611a4e6024836116e3565b9150611a59826119f2565b604082019050919050565b60006020820190508181036000830152611a7d81611a41565b9050919050565b611a8d81611618565b82525050565b6000604082019050611aa860008301856112df565b611ab56020830184611a84565b9392505050565b6000611ac782611309565b9150611ad283611309565b925082821015611ae557611ae4611819565b5b828203905092915050565b60008160011c9050919050565b6000808291508390505b6001851115611b4757808604811115611b2357611b22611819565b5b6001851615611b325780820291505b8081029050611b4085611af0565b9450611b07565b94509492505050565b600082611b605760019050611c1c565b81611b6e5760009050611c1c565b8160018114611b845760028114611b8e57611bbd565b6001915050611c1c565b60ff841115611ba057611b9f611819565b5b8360020a915084821115611bb757611bb6611819565b5b50611c1c565b5060208310610133831016604e8410600b8410161715611bf25782820a905083811115611bed57611bec611819565b5b611c1c565b611bff8484846001611afd565b92509050818404811115611c1657611c15611819565b5b81810290505b9392505050565b6000611c2e82611618565b9150611c3983611309565b9250611c667fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484611b50565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000611ca882611618565b9150611cb383611618565b925082611cc357611cc2611c6e565b5b828204905092915050565b6000611cd982611618565b9150611ce483611618565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611d1d57611d1c611819565b5b828202905092915050565b6000819050919050565b611d43611d3e82611618565b611d28565b82525050565b6000611d558284611d32565b60208201915081905092915050565b6000606082019050611d7960008301866112df565b611d8660208301856112df565b611d936040830184611a84565b949350505050565b7f45524332303a206e6f74206120636f6e74726163740000000000000000000000600082015250565b6000611dd16015836116e3565b9150611ddc82611d9b565b602082019050919050565b60006020820190508181036000830152611e0081611dc4565b9050919050565b7f45524332303a2063616c6c206661696c65640000000000000000000000000000600082015250565b6000611e3d6012836116e3565b9150611e4882611e07565b602082019050919050565b60006020820190508181036000830152611e6c81611e30565b9050919050565b611e7c8161134f565b8114611e8757600080fd5b50565b600081519050611e9981611e73565b92915050565b600060208284031215611eb557611eb4611058565b5b6000611ec384828501611e8a565b91505092915050565b7f45524332303a206f7065726174696f6e20646964206e6f742073756363656564600082015250565b6000611f026020836116e3565b9150611f0d82611ecc565b602082019050919050565b60006020820190508181036000830152611f3181611ef5565b905091905056fea264697066735822122073d45daeafaae9e9ee8756b183b6fab6a6ec9a693c0de337b3de53abd15c16bd64736f6c634300080b0033"

// DeployERC20Handler deploys a new Ethereum contract, binding an instance of ERC20Handler to it.
func DeployERC20Handler(auth *bind.TransactOpts, backend bind.ContractBackend, bridgeAddress common.Address) (common.Address, *types.Transaction, *ERC20Handler, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20HandlerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20HandlerBin), backend, bridgeAddress)
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

// TokenContractAddressToTokenProperties is a free data retrieval call binding the contract method 0xac607c21.
//
// Solidity: function _tokenContractAddressToTokenProperties(address ) view returns(bytes32 resourceID, bool isWhitelisted, bool isBurnable, (bool,uint8) decimals)
func (_ERC20Handler *ERC20HandlerCaller) TokenContractAddressToTokenProperties(opts *bind.CallOpts, arg0 common.Address) (struct {
	ResourceID    [32]byte
	IsWhitelisted bool
	IsBurnable    bool
	Decimals      ERCHandlerHelpersDecimals
}, error) {
	var out []interface{}
	err := _ERC20Handler.contract.Call(opts, &out, "_tokenContractAddressToTokenProperties", arg0)

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
func (_ERC20Handler *ERC20HandlerSession) TokenContractAddressToTokenProperties(arg0 common.Address) (struct {
	ResourceID    [32]byte
	IsWhitelisted bool
	IsBurnable    bool
	Decimals      ERCHandlerHelpersDecimals
}, error) {
	return _ERC20Handler.Contract.TokenContractAddressToTokenProperties(&_ERC20Handler.CallOpts, arg0)
}

// TokenContractAddressToTokenProperties is a free data retrieval call binding the contract method 0xac607c21.
//
// Solidity: function _tokenContractAddressToTokenProperties(address ) view returns(bytes32 resourceID, bool isWhitelisted, bool isBurnable, (bool,uint8) decimals)
func (_ERC20Handler *ERC20HandlerCallerSession) TokenContractAddressToTokenProperties(arg0 common.Address) (struct {
	ResourceID    [32]byte
	IsWhitelisted bool
	IsBurnable    bool
	Decimals      ERCHandlerHelpersDecimals
}, error) {
	return _ERC20Handler.Contract.TokenContractAddressToTokenProperties(&_ERC20Handler.CallOpts, arg0)
}

// DefaultDecimals is a free data retrieval call binding the contract method 0x30f08abd.
//
// Solidity: function defaultDecimals() view returns(uint8)
func (_ERC20Handler *ERC20HandlerCaller) DefaultDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20Handler.contract.Call(opts, &out, "defaultDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DefaultDecimals is a free data retrieval call binding the contract method 0x30f08abd.
//
// Solidity: function defaultDecimals() view returns(uint8)
func (_ERC20Handler *ERC20HandlerSession) DefaultDecimals() (uint8, error) {
	return _ERC20Handler.Contract.DefaultDecimals(&_ERC20Handler.CallOpts)
}

// DefaultDecimals is a free data retrieval call binding the contract method 0x30f08abd.
//
// Solidity: function defaultDecimals() view returns(uint8)
func (_ERC20Handler *ERC20HandlerCallerSession) DefaultDecimals() (uint8, error) {
	return _ERC20Handler.Contract.DefaultDecimals(&_ERC20Handler.CallOpts)
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
