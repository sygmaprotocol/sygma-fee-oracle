// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package FeeHandlerRouter

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

// FeeHandlerRouterMetaData contains all meta data concerning the FeeHandlerRouter contract.
var FeeHandlerRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"FeeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"fromDomainID\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"destinationDomainID\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"FeeCollected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeeDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_domainResourceIDToFeeHandlerAddress\",\"outputs\":[{\"internalType\":\"contractIFeeHandler\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"destinationDomainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"contractIFeeHandler\",\"name\":\"handlerAddress\",\"type\":\"address\"}],\"name\":\"adminSetResourceHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"fromDomainID\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"destinationDomainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"depositData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"name\":\"calculateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"fromDomainID\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"destinationDomainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"depositData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"name\":\"collectFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getRoleMemberIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b5060405162001a3538038062001a35833981810160405281019062000037919062000287565b8073ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1681525050620000806000801b336200008760201b60201c565b50620002b9565b6200009982826200009d60201b60201c565b5050565b620000cb816000808581526020019081526020016000206000016200014060201b620008ae1790919060201c565b156200013c57620000e16200017860201b60201c565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b600062000170836000018373ffffffffffffffffffffffffffffffffffffffff1660001b6200018060201b60201c565b905092915050565b600033905090565b6000620001948383620001fa60201b60201c565b620001ef578260000182908060018154018082558091505060019003906000526020600020016000909190919091505582600001805490508360010160008481526020019081526020016000208190555060019050620001f4565b600090505b92915050565b600080836001016000848152602001908152602001600020541415905092915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006200024f8262000222565b9050919050565b620002618162000242565b81146200026d57600080fd5b50565b600081519050620002818162000256565b92915050565b600060208284031215620002a0576200029f6200021d565b5b6000620002b08482850162000270565b91505092915050565b608051611759620002dc6000396000818161058a015261093501526117596000f3fe6080604052600436106100dd5760003560e01c80634e0df3f61161007f578063a217fddf11610059578063a217fddf146102d5578063ca15c87314610300578063d547741f1461033d578063ef4f081f14610366576100dd565b80634e0df3f61461021e5780639010d07c1461025b57806391d1485414610298576100dd565b80632f2ff15d116100bb5780632f2ff15d14610164578063318c136e1461018d57806336568abe146101b85780633d94ebc6146101e1576100dd565b80631af5fe38146100e2578063248a9ca31461010b5780632530706514610148575b600080fd5b3480156100ee57600080fd5b5061010960048036038101906101049190610e4c565b6103a4565b005b34801561011757600080fd5b50610132600480360381019061012d9190610e9f565b61041a565b60405161013f9190610edb565b60405180910390f35b610162600480360381019061015d9190610f87565b610439565b005b34801561017057600080fd5b5061018b60048036038101906101869190611056565b610515565b005b34801561019957600080fd5b506101a2610588565b6040516101af91906110a5565b60405180910390f35b3480156101c457600080fd5b506101df60048036038101906101da9190611056565b6105ac565b005b3480156101ed57600080fd5b50610208600480360381019061020391906110c0565b61062f565b604051610215919061115f565b60405180910390f35b34801561022a57600080fd5b5061024560048036038101906102409190611056565b610671565b6040516102529190611193565b60405180910390f35b34801561026757600080fd5b50610282600480360381019061027d91906111da565b6106c1565b60405161028f91906110a5565b60405180910390f35b3480156102a457600080fd5b506102bf60048036038101906102ba9190611056565b6106f2565b6040516102cc9190611235565b60405180910390f35b3480156102e157600080fd5b506102ea610723565b6040516102f79190610edb565b60405180910390f35b34801561030c57600080fd5b5061032760048036038101906103229190610e9f565b61072a565b6040516103349190611193565b60405180910390f35b34801561034957600080fd5b50610364600480360381019061035f9190611056565b610750565b005b34801561037257600080fd5b5061038d60048036038101906103889190610f87565b6107c3565b60405161039b929190611250565b60405180910390f35b6103ac6108de565b80600160008560ff1660ff168152602001908152602001600020600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050565b6000806000838152602001908152602001600020600201549050919050565b610441610933565b6000600160008860ff1660ff168152602001908152602001600020600087815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff166325307065348b8b8b8b8b8b8b8b6040518a63ffffffff1660e01b81526004016104d89897969594939291906112e6565b6000604051808303818588803b1580156104f157600080fd5b505af1158015610505573d6000803e3d6000fd5b5050505050505050505050505050565b61053b600080848152602001908152602001600020600201546105366109c3565b6106f2565b61057a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610571906113dc565b60405180910390fd5b61058482826109cb565b5050565b7f000000000000000000000000000000000000000000000000000000000000000081565b6105b46109c3565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610621576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106189061146e565b60405180910390fd5b61062b8282610a5e565b5050565b60016020528160005260406000206020528060005260406000206000915091509054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600080600084815260200190815260200160002060000160000160010160008373ffffffffffffffffffffffffffffffffffffffff1660001b815260200190815260200160002054905092915050565b60006106ea82600080868152602001908152602001600020600001610af190919063ffffffff16565b905092915050565b600061071b82600080868152602001908152602001600020600001610b0b90919063ffffffff16565b905092915050565b6000801b81565b6000610749600080848152602001908152602001600020600001610b3b565b9050919050565b610776600080848152602001908152602001600020600201546107716109c3565b6106f2565b6107b5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107ac90611500565b60405180910390fd5b6107bf8282610a5e565b5050565b6000806000600160008a60ff1660ff168152602001908152602001600020600089815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff1663ef4f081f8c8c8c8c8c8c8c8c6040518963ffffffff1660e01b815260040161085c9897969594939291906112e6565b6040805180830381865afa158015610878573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061089c919061154a565b92509250509850989650505050505050565b60006108d6836000018373ffffffffffffffffffffffffffffffffffffffff1660001b610b50565b905092915050565b6108f26000801b6108ed6109c3565b6106f2565b610931576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610928906115d6565b60405180910390fd5b565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146109c1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109b890611642565b60405180910390fd5b565b600033905090565b6109f2816000808581526020019081526020016000206000016108ae90919063ffffffff16565b15610a5a576109ff6109c3565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b610a8581600080858152602001908152602001600020600001610bc090919063ffffffff16565b15610aed57610a926109c3565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45b5050565b6000610b008360000183610bf0565b60001c905092915050565b6000610b33836000018373ffffffffffffffffffffffffffffffffffffffff1660001b610c1b565b905092915050565b6000610b4982600001610c3e565b9050919050565b6000610b5c8383610c1b565b610bb5578260000182908060018154018082558091505060019003906000526020600020016000909190919091505582600001805490508360010160008481526020019081526020016000208190555060019050610bba565b600090505b92915050565b6000610be8836000018373ffffffffffffffffffffffffffffffffffffffff1660001b610c4f565b905092915050565b6000826000018281548110610c0857610c07611662565b5b9060005260206000200154905092915050565b600080836001016000848152602001908152602001600020541415905092915050565b600081600001805490509050919050565b60008083600101600084815260200190815260200160002054905060008114610d57576000600182610c8191906116c0565b9050600060018660000180549050610c9991906116c0565b9050818114610d08576000866000018281548110610cba57610cb9611662565b5b9060005260206000200154905080876000018481548110610cde57610cdd611662565b5b90600052602060002001819055508387600101600083815260200190815260200160002081905550505b85600001805480610d1c57610d1b6116f4565b5b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610d5d565b60009150505b92915050565b600080fd5b600080fd5b600060ff82169050919050565b610d8381610d6d565b8114610d8e57600080fd5b50565b600081359050610da081610d7a565b92915050565b6000819050919050565b610db981610da6565b8114610dc457600080fd5b50565b600081359050610dd681610db0565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610e0782610ddc565b9050919050565b6000610e1982610dfc565b9050919050565b610e2981610e0e565b8114610e3457600080fd5b50565b600081359050610e4681610e20565b92915050565b600080600060608486031215610e6557610e64610d63565b5b6000610e7386828701610d91565b9350506020610e8486828701610dc7565b9250506040610e9586828701610e37565b9150509250925092565b600060208284031215610eb557610eb4610d63565b5b6000610ec384828501610dc7565b91505092915050565b610ed581610da6565b82525050565b6000602082019050610ef06000830184610ecc565b92915050565b610eff81610dfc565b8114610f0a57600080fd5b50565b600081359050610f1c81610ef6565b92915050565b600080fd5b600080fd5b600080fd5b60008083601f840112610f4757610f46610f22565b5b8235905067ffffffffffffffff811115610f6457610f63610f27565b5b602083019150836001820283011115610f8057610f7f610f2c565b5b9250929050565b60008060008060008060008060c0898b031215610fa757610fa6610d63565b5b6000610fb58b828c01610f0d565b9850506020610fc68b828c01610d91565b9750506040610fd78b828c01610d91565b9650506060610fe88b828c01610dc7565b955050608089013567ffffffffffffffff81111561100957611008610d68565b5b6110158b828c01610f31565b945094505060a089013567ffffffffffffffff81111561103857611037610d68565b5b6110448b828c01610f31565b92509250509295985092959890939650565b6000806040838503121561106d5761106c610d63565b5b600061107b85828601610dc7565b925050602061108c85828601610f0d565b9150509250929050565b61109f81610dfc565b82525050565b60006020820190506110ba6000830184611096565b92915050565b600080604083850312156110d7576110d6610d63565b5b60006110e585828601610d91565b92505060206110f685828601610dc7565b9150509250929050565b6000819050919050565b600061112561112061111b84610ddc565b611100565b610ddc565b9050919050565b60006111378261110a565b9050919050565b60006111498261112c565b9050919050565b6111598161113e565b82525050565b60006020820190506111746000830184611150565b92915050565b6000819050919050565b61118d8161117a565b82525050565b60006020820190506111a86000830184611184565b92915050565b6111b78161117a565b81146111c257600080fd5b50565b6000813590506111d4816111ae565b92915050565b600080604083850312156111f1576111f0610d63565b5b60006111ff85828601610dc7565b9250506020611210858286016111c5565b9150509250929050565b60008115159050919050565b61122f8161121a565b82525050565b600060208201905061124a6000830184611226565b92915050565b60006040820190506112656000830185611184565b6112726020830184611096565b9392505050565b61128281610d6d565b82525050565b600082825260208201905092915050565b82818337600083830152505050565b6000601f19601f8301169050919050565b60006112c58385611288565b93506112d2838584611299565b6112db836112a8565b840190509392505050565b600060c0820190506112fb600083018b611096565b611308602083018a611279565b6113156040830189611279565b6113226060830188610ecc565b81810360808301526113358186886112b9565b905081810360a083015261134a8184866112b9565b90509998505050505050505050565b600082825260208201905092915050565b7f416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e60008201527f2061646d696e20746f206772616e740000000000000000000000000000000000602082015250565b60006113c6602f83611359565b91506113d18261136a565b604082019050919050565b600060208201905081810360008301526113f5816113b9565b9050919050565b7f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560008201527f20726f6c657320666f722073656c660000000000000000000000000000000000602082015250565b6000611458602f83611359565b9150611463826113fc565b604082019050919050565b600060208201905081810360008301526114878161144b565b9050919050565b7f416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e60008201527f2061646d696e20746f207265766f6b6500000000000000000000000000000000602082015250565b60006114ea603083611359565b91506114f58261148e565b604082019050919050565b60006020820190508181036000830152611519816114dd565b9050919050565b60008151905061152f816111ae565b92915050565b60008151905061154481610ef6565b92915050565b6000806040838503121561156157611560610d63565b5b600061156f85828601611520565b925050602061158085828601611535565b9150509250929050565b7f73656e64657220646f65736e277420686176652061646d696e20726f6c650000600082015250565b60006115c0601e83611359565b91506115cb8261158a565b602082019050919050565b600060208201905081810360008301526115ef816115b3565b9050919050565b7f73656e646572206d7573742062652062726964676520636f6e74726163740000600082015250565b600061162c601e83611359565b9150611637826115f6565b602082019050919050565b6000602082019050818103600083015261165b8161161f565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006116cb8261117a565b91506116d68361117a565b9250828210156116e9576116e8611691565b5b828203905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea26469706673582212204f6aeb36bf9f51c53e48bc9a7c45d47ef77ce6a8ed26314abcbecf73fb7b89e364736f6c634300080b0033",
}

// FeeHandlerRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use FeeHandlerRouterMetaData.ABI instead.
var FeeHandlerRouterABI = FeeHandlerRouterMetaData.ABI

// FeeHandlerRouterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FeeHandlerRouterMetaData.Bin instead.
var FeeHandlerRouterBin = FeeHandlerRouterMetaData.Bin

// DeployFeeHandlerRouter deploys a new Ethereum contract, binding an instance of FeeHandlerRouter to it.
func DeployFeeHandlerRouter(auth *bind.TransactOpts, backend bind.ContractBackend, bridgeAddress common.Address) (common.Address, *types.Transaction, *FeeHandlerRouter, error) {
	parsed, err := FeeHandlerRouterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FeeHandlerRouterBin), backend, bridgeAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FeeHandlerRouter{FeeHandlerRouterCaller: FeeHandlerRouterCaller{contract: contract}, FeeHandlerRouterTransactor: FeeHandlerRouterTransactor{contract: contract}, FeeHandlerRouterFilterer: FeeHandlerRouterFilterer{contract: contract}}, nil
}

// FeeHandlerRouter is an auto generated Go binding around an Ethereum contract.
type FeeHandlerRouter struct {
	FeeHandlerRouterCaller     // Read-only binding to the contract
	FeeHandlerRouterTransactor // Write-only binding to the contract
	FeeHandlerRouterFilterer   // Log filterer for contract events
}

// FeeHandlerRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type FeeHandlerRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeHandlerRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FeeHandlerRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeHandlerRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FeeHandlerRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeHandlerRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FeeHandlerRouterSession struct {
	Contract     *FeeHandlerRouter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FeeHandlerRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FeeHandlerRouterCallerSession struct {
	Contract *FeeHandlerRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// FeeHandlerRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FeeHandlerRouterTransactorSession struct {
	Contract     *FeeHandlerRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// FeeHandlerRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type FeeHandlerRouterRaw struct {
	Contract *FeeHandlerRouter // Generic contract binding to access the raw methods on
}

// FeeHandlerRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FeeHandlerRouterCallerRaw struct {
	Contract *FeeHandlerRouterCaller // Generic read-only contract binding to access the raw methods on
}

// FeeHandlerRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FeeHandlerRouterTransactorRaw struct {
	Contract *FeeHandlerRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFeeHandlerRouter creates a new instance of FeeHandlerRouter, bound to a specific deployed contract.
func NewFeeHandlerRouter(address common.Address, backend bind.ContractBackend) (*FeeHandlerRouter, error) {
	contract, err := bindFeeHandlerRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeeHandlerRouter{FeeHandlerRouterCaller: FeeHandlerRouterCaller{contract: contract}, FeeHandlerRouterTransactor: FeeHandlerRouterTransactor{contract: contract}, FeeHandlerRouterFilterer: FeeHandlerRouterFilterer{contract: contract}}, nil
}

// NewFeeHandlerRouterCaller creates a new read-only instance of FeeHandlerRouter, bound to a specific deployed contract.
func NewFeeHandlerRouterCaller(address common.Address, caller bind.ContractCaller) (*FeeHandlerRouterCaller, error) {
	contract, err := bindFeeHandlerRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeeHandlerRouterCaller{contract: contract}, nil
}

// NewFeeHandlerRouterTransactor creates a new write-only instance of FeeHandlerRouter, bound to a specific deployed contract.
func NewFeeHandlerRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*FeeHandlerRouterTransactor, error) {
	contract, err := bindFeeHandlerRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeeHandlerRouterTransactor{contract: contract}, nil
}

// NewFeeHandlerRouterFilterer creates a new log filterer instance of FeeHandlerRouter, bound to a specific deployed contract.
func NewFeeHandlerRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*FeeHandlerRouterFilterer, error) {
	contract, err := bindFeeHandlerRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeeHandlerRouterFilterer{contract: contract}, nil
}

// bindFeeHandlerRouter binds a generic wrapper to an already deployed contract.
func bindFeeHandlerRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FeeHandlerRouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeHandlerRouter *FeeHandlerRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeHandlerRouter.Contract.FeeHandlerRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeHandlerRouter *FeeHandlerRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeHandlerRouter.Contract.FeeHandlerRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeHandlerRouter *FeeHandlerRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeHandlerRouter.Contract.FeeHandlerRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeHandlerRouter *FeeHandlerRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeHandlerRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeHandlerRouter *FeeHandlerRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeHandlerRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeHandlerRouter *FeeHandlerRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeHandlerRouter.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_FeeHandlerRouter *FeeHandlerRouterCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FeeHandlerRouter.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_FeeHandlerRouter *FeeHandlerRouterSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _FeeHandlerRouter.Contract.DEFAULTADMINROLE(&_FeeHandlerRouter.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_FeeHandlerRouter *FeeHandlerRouterCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _FeeHandlerRouter.Contract.DEFAULTADMINROLE(&_FeeHandlerRouter.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_FeeHandlerRouter *FeeHandlerRouterCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeHandlerRouter.contract.Call(opts, &out, "_bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_FeeHandlerRouter *FeeHandlerRouterSession) BridgeAddress() (common.Address, error) {
	return _FeeHandlerRouter.Contract.BridgeAddress(&_FeeHandlerRouter.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_FeeHandlerRouter *FeeHandlerRouterCallerSession) BridgeAddress() (common.Address, error) {
	return _FeeHandlerRouter.Contract.BridgeAddress(&_FeeHandlerRouter.CallOpts)
}

// DomainResourceIDToFeeHandlerAddress is a free data retrieval call binding the contract method 0x3d94ebc6.
//
// Solidity: function _domainResourceIDToFeeHandlerAddress(uint8 , bytes32 ) view returns(address)
func (_FeeHandlerRouter *FeeHandlerRouterCaller) DomainResourceIDToFeeHandlerAddress(opts *bind.CallOpts, arg0 uint8, arg1 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _FeeHandlerRouter.contract.Call(opts, &out, "_domainResourceIDToFeeHandlerAddress", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DomainResourceIDToFeeHandlerAddress is a free data retrieval call binding the contract method 0x3d94ebc6.
//
// Solidity: function _domainResourceIDToFeeHandlerAddress(uint8 , bytes32 ) view returns(address)
func (_FeeHandlerRouter *FeeHandlerRouterSession) DomainResourceIDToFeeHandlerAddress(arg0 uint8, arg1 [32]byte) (common.Address, error) {
	return _FeeHandlerRouter.Contract.DomainResourceIDToFeeHandlerAddress(&_FeeHandlerRouter.CallOpts, arg0, arg1)
}

// DomainResourceIDToFeeHandlerAddress is a free data retrieval call binding the contract method 0x3d94ebc6.
//
// Solidity: function _domainResourceIDToFeeHandlerAddress(uint8 , bytes32 ) view returns(address)
func (_FeeHandlerRouter *FeeHandlerRouterCallerSession) DomainResourceIDToFeeHandlerAddress(arg0 uint8, arg1 [32]byte) (common.Address, error) {
	return _FeeHandlerRouter.Contract.DomainResourceIDToFeeHandlerAddress(&_FeeHandlerRouter.CallOpts, arg0, arg1)
}

// CalculateFee is a free data retrieval call binding the contract method 0xef4f081f.
//
// Solidity: function calculateFee(address sender, uint8 fromDomainID, uint8 destinationDomainID, bytes32 resourceID, bytes depositData, bytes feeData) view returns(uint256 fee, address tokenAddress)
func (_FeeHandlerRouter *FeeHandlerRouterCaller) CalculateFee(opts *bind.CallOpts, sender common.Address, fromDomainID uint8, destinationDomainID uint8, resourceID [32]byte, depositData []byte, feeData []byte) (struct {
	Fee          *big.Int
	TokenAddress common.Address
}, error) {
	var out []interface{}
	err := _FeeHandlerRouter.contract.Call(opts, &out, "calculateFee", sender, fromDomainID, destinationDomainID, resourceID, depositData, feeData)

	outstruct := new(struct {
		Fee          *big.Int
		TokenAddress common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fee = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TokenAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// CalculateFee is a free data retrieval call binding the contract method 0xef4f081f.
//
// Solidity: function calculateFee(address sender, uint8 fromDomainID, uint8 destinationDomainID, bytes32 resourceID, bytes depositData, bytes feeData) view returns(uint256 fee, address tokenAddress)
func (_FeeHandlerRouter *FeeHandlerRouterSession) CalculateFee(sender common.Address, fromDomainID uint8, destinationDomainID uint8, resourceID [32]byte, depositData []byte, feeData []byte) (struct {
	Fee          *big.Int
	TokenAddress common.Address
}, error) {
	return _FeeHandlerRouter.Contract.CalculateFee(&_FeeHandlerRouter.CallOpts, sender, fromDomainID, destinationDomainID, resourceID, depositData, feeData)
}

// CalculateFee is a free data retrieval call binding the contract method 0xef4f081f.
//
// Solidity: function calculateFee(address sender, uint8 fromDomainID, uint8 destinationDomainID, bytes32 resourceID, bytes depositData, bytes feeData) view returns(uint256 fee, address tokenAddress)
func (_FeeHandlerRouter *FeeHandlerRouterCallerSession) CalculateFee(sender common.Address, fromDomainID uint8, destinationDomainID uint8, resourceID [32]byte, depositData []byte, feeData []byte) (struct {
	Fee          *big.Int
	TokenAddress common.Address
}, error) {
	return _FeeHandlerRouter.Contract.CalculateFee(&_FeeHandlerRouter.CallOpts, sender, fromDomainID, destinationDomainID, resourceID, depositData, feeData)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_FeeHandlerRouter *FeeHandlerRouterCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _FeeHandlerRouter.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_FeeHandlerRouter *FeeHandlerRouterSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _FeeHandlerRouter.Contract.GetRoleAdmin(&_FeeHandlerRouter.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_FeeHandlerRouter *FeeHandlerRouterCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _FeeHandlerRouter.Contract.GetRoleAdmin(&_FeeHandlerRouter.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_FeeHandlerRouter *FeeHandlerRouterCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _FeeHandlerRouter.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_FeeHandlerRouter *FeeHandlerRouterSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _FeeHandlerRouter.Contract.GetRoleMember(&_FeeHandlerRouter.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_FeeHandlerRouter *FeeHandlerRouterCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _FeeHandlerRouter.Contract.GetRoleMember(&_FeeHandlerRouter.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_FeeHandlerRouter *FeeHandlerRouterCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _FeeHandlerRouter.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_FeeHandlerRouter *FeeHandlerRouterSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _FeeHandlerRouter.Contract.GetRoleMemberCount(&_FeeHandlerRouter.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_FeeHandlerRouter *FeeHandlerRouterCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _FeeHandlerRouter.Contract.GetRoleMemberCount(&_FeeHandlerRouter.CallOpts, role)
}

// GetRoleMemberIndex is a free data retrieval call binding the contract method 0x4e0df3f6.
//
// Solidity: function getRoleMemberIndex(bytes32 role, address account) view returns(uint256)
func (_FeeHandlerRouter *FeeHandlerRouterCaller) GetRoleMemberIndex(opts *bind.CallOpts, role [32]byte, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FeeHandlerRouter.contract.Call(opts, &out, "getRoleMemberIndex", role, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberIndex is a free data retrieval call binding the contract method 0x4e0df3f6.
//
// Solidity: function getRoleMemberIndex(bytes32 role, address account) view returns(uint256)
func (_FeeHandlerRouter *FeeHandlerRouterSession) GetRoleMemberIndex(role [32]byte, account common.Address) (*big.Int, error) {
	return _FeeHandlerRouter.Contract.GetRoleMemberIndex(&_FeeHandlerRouter.CallOpts, role, account)
}

// GetRoleMemberIndex is a free data retrieval call binding the contract method 0x4e0df3f6.
//
// Solidity: function getRoleMemberIndex(bytes32 role, address account) view returns(uint256)
func (_FeeHandlerRouter *FeeHandlerRouterCallerSession) GetRoleMemberIndex(role [32]byte, account common.Address) (*big.Int, error) {
	return _FeeHandlerRouter.Contract.GetRoleMemberIndex(&_FeeHandlerRouter.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_FeeHandlerRouter *FeeHandlerRouterCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _FeeHandlerRouter.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_FeeHandlerRouter *FeeHandlerRouterSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _FeeHandlerRouter.Contract.HasRole(&_FeeHandlerRouter.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_FeeHandlerRouter *FeeHandlerRouterCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _FeeHandlerRouter.Contract.HasRole(&_FeeHandlerRouter.CallOpts, role, account)
}

// AdminSetResourceHandler is a paid mutator transaction binding the contract method 0x1af5fe38.
//
// Solidity: function adminSetResourceHandler(uint8 destinationDomainID, bytes32 resourceID, address handlerAddress) returns()
func (_FeeHandlerRouter *FeeHandlerRouterTransactor) AdminSetResourceHandler(opts *bind.TransactOpts, destinationDomainID uint8, resourceID [32]byte, handlerAddress common.Address) (*types.Transaction, error) {
	return _FeeHandlerRouter.contract.Transact(opts, "adminSetResourceHandler", destinationDomainID, resourceID, handlerAddress)
}

// AdminSetResourceHandler is a paid mutator transaction binding the contract method 0x1af5fe38.
//
// Solidity: function adminSetResourceHandler(uint8 destinationDomainID, bytes32 resourceID, address handlerAddress) returns()
func (_FeeHandlerRouter *FeeHandlerRouterSession) AdminSetResourceHandler(destinationDomainID uint8, resourceID [32]byte, handlerAddress common.Address) (*types.Transaction, error) {
	return _FeeHandlerRouter.Contract.AdminSetResourceHandler(&_FeeHandlerRouter.TransactOpts, destinationDomainID, resourceID, handlerAddress)
}

// AdminSetResourceHandler is a paid mutator transaction binding the contract method 0x1af5fe38.
//
// Solidity: function adminSetResourceHandler(uint8 destinationDomainID, bytes32 resourceID, address handlerAddress) returns()
func (_FeeHandlerRouter *FeeHandlerRouterTransactorSession) AdminSetResourceHandler(destinationDomainID uint8, resourceID [32]byte, handlerAddress common.Address) (*types.Transaction, error) {
	return _FeeHandlerRouter.Contract.AdminSetResourceHandler(&_FeeHandlerRouter.TransactOpts, destinationDomainID, resourceID, handlerAddress)
}

// CollectFee is a paid mutator transaction binding the contract method 0x25307065.
//
// Solidity: function collectFee(address sender, uint8 fromDomainID, uint8 destinationDomainID, bytes32 resourceID, bytes depositData, bytes feeData) payable returns()
func (_FeeHandlerRouter *FeeHandlerRouterTransactor) CollectFee(opts *bind.TransactOpts, sender common.Address, fromDomainID uint8, destinationDomainID uint8, resourceID [32]byte, depositData []byte, feeData []byte) (*types.Transaction, error) {
	return _FeeHandlerRouter.contract.Transact(opts, "collectFee", sender, fromDomainID, destinationDomainID, resourceID, depositData, feeData)
}

// CollectFee is a paid mutator transaction binding the contract method 0x25307065.
//
// Solidity: function collectFee(address sender, uint8 fromDomainID, uint8 destinationDomainID, bytes32 resourceID, bytes depositData, bytes feeData) payable returns()
func (_FeeHandlerRouter *FeeHandlerRouterSession) CollectFee(sender common.Address, fromDomainID uint8, destinationDomainID uint8, resourceID [32]byte, depositData []byte, feeData []byte) (*types.Transaction, error) {
	return _FeeHandlerRouter.Contract.CollectFee(&_FeeHandlerRouter.TransactOpts, sender, fromDomainID, destinationDomainID, resourceID, depositData, feeData)
}

// CollectFee is a paid mutator transaction binding the contract method 0x25307065.
//
// Solidity: function collectFee(address sender, uint8 fromDomainID, uint8 destinationDomainID, bytes32 resourceID, bytes depositData, bytes feeData) payable returns()
func (_FeeHandlerRouter *FeeHandlerRouterTransactorSession) CollectFee(sender common.Address, fromDomainID uint8, destinationDomainID uint8, resourceID [32]byte, depositData []byte, feeData []byte) (*types.Transaction, error) {
	return _FeeHandlerRouter.Contract.CollectFee(&_FeeHandlerRouter.TransactOpts, sender, fromDomainID, destinationDomainID, resourceID, depositData, feeData)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_FeeHandlerRouter *FeeHandlerRouterTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FeeHandlerRouter.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_FeeHandlerRouter *FeeHandlerRouterSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FeeHandlerRouter.Contract.GrantRole(&_FeeHandlerRouter.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_FeeHandlerRouter *FeeHandlerRouterTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FeeHandlerRouter.Contract.GrantRole(&_FeeHandlerRouter.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_FeeHandlerRouter *FeeHandlerRouterTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FeeHandlerRouter.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_FeeHandlerRouter *FeeHandlerRouterSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FeeHandlerRouter.Contract.RenounceRole(&_FeeHandlerRouter.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_FeeHandlerRouter *FeeHandlerRouterTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FeeHandlerRouter.Contract.RenounceRole(&_FeeHandlerRouter.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_FeeHandlerRouter *FeeHandlerRouterTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FeeHandlerRouter.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_FeeHandlerRouter *FeeHandlerRouterSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FeeHandlerRouter.Contract.RevokeRole(&_FeeHandlerRouter.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_FeeHandlerRouter *FeeHandlerRouterTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FeeHandlerRouter.Contract.RevokeRole(&_FeeHandlerRouter.TransactOpts, role, account)
}

// FeeHandlerRouterFeeChangedIterator is returned from FilterFeeChanged and is used to iterate over the raw logs and unpacked data for FeeChanged events raised by the FeeHandlerRouter contract.
type FeeHandlerRouterFeeChangedIterator struct {
	Event *FeeHandlerRouterFeeChanged // Event containing the contract specifics and raw log

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
func (it *FeeHandlerRouterFeeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeHandlerRouterFeeChanged)
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
		it.Event = new(FeeHandlerRouterFeeChanged)
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
func (it *FeeHandlerRouterFeeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeHandlerRouterFeeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeHandlerRouterFeeChanged represents a FeeChanged event raised by the FeeHandlerRouter contract.
type FeeHandlerRouterFeeChanged struct {
	NewFee *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeeChanged is a free log retrieval operation binding the contract event 0x6bbc57480a46553fa4d156ce702beef5f3ad66303b0ed1a5d4cb44966c6584c3.
//
// Solidity: event FeeChanged(uint256 newFee)
func (_FeeHandlerRouter *FeeHandlerRouterFilterer) FilterFeeChanged(opts *bind.FilterOpts) (*FeeHandlerRouterFeeChangedIterator, error) {

	logs, sub, err := _FeeHandlerRouter.contract.FilterLogs(opts, "FeeChanged")
	if err != nil {
		return nil, err
	}
	return &FeeHandlerRouterFeeChangedIterator{contract: _FeeHandlerRouter.contract, event: "FeeChanged", logs: logs, sub: sub}, nil
}

// WatchFeeChanged is a free log subscription operation binding the contract event 0x6bbc57480a46553fa4d156ce702beef5f3ad66303b0ed1a5d4cb44966c6584c3.
//
// Solidity: event FeeChanged(uint256 newFee)
func (_FeeHandlerRouter *FeeHandlerRouterFilterer) WatchFeeChanged(opts *bind.WatchOpts, sink chan<- *FeeHandlerRouterFeeChanged) (event.Subscription, error) {

	logs, sub, err := _FeeHandlerRouter.contract.WatchLogs(opts, "FeeChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeHandlerRouterFeeChanged)
				if err := _FeeHandlerRouter.contract.UnpackLog(event, "FeeChanged", log); err != nil {
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

// ParseFeeChanged is a log parse operation binding the contract event 0x6bbc57480a46553fa4d156ce702beef5f3ad66303b0ed1a5d4cb44966c6584c3.
//
// Solidity: event FeeChanged(uint256 newFee)
func (_FeeHandlerRouter *FeeHandlerRouterFilterer) ParseFeeChanged(log types.Log) (*FeeHandlerRouterFeeChanged, error) {
	event := new(FeeHandlerRouterFeeChanged)
	if err := _FeeHandlerRouter.contract.UnpackLog(event, "FeeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeHandlerRouterFeeCollectedIterator is returned from FilterFeeCollected and is used to iterate over the raw logs and unpacked data for FeeCollected events raised by the FeeHandlerRouter contract.
type FeeHandlerRouterFeeCollectedIterator struct {
	Event *FeeHandlerRouterFeeCollected // Event containing the contract specifics and raw log

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
func (it *FeeHandlerRouterFeeCollectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeHandlerRouterFeeCollected)
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
		it.Event = new(FeeHandlerRouterFeeCollected)
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
func (it *FeeHandlerRouterFeeCollectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeHandlerRouterFeeCollectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeHandlerRouterFeeCollected represents a FeeCollected event raised by the FeeHandlerRouter contract.
type FeeHandlerRouterFeeCollected struct {
	Sender              common.Address
	FromDomainID        uint8
	DestinationDomainID uint8
	ResourceID          [32]byte
	Fee                 *big.Int
	TokenAddress        common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterFeeCollected is a free log retrieval operation binding the contract event 0xbd231b7fa4103e15e7a238c72f07e8aff310701af121895aa6c793b80245e433.
//
// Solidity: event FeeCollected(address sender, uint8 fromDomainID, uint8 destinationDomainID, bytes32 resourceID, uint256 fee, address tokenAddress)
func (_FeeHandlerRouter *FeeHandlerRouterFilterer) FilterFeeCollected(opts *bind.FilterOpts) (*FeeHandlerRouterFeeCollectedIterator, error) {

	logs, sub, err := _FeeHandlerRouter.contract.FilterLogs(opts, "FeeCollected")
	if err != nil {
		return nil, err
	}
	return &FeeHandlerRouterFeeCollectedIterator{contract: _FeeHandlerRouter.contract, event: "FeeCollected", logs: logs, sub: sub}, nil
}

// WatchFeeCollected is a free log subscription operation binding the contract event 0xbd231b7fa4103e15e7a238c72f07e8aff310701af121895aa6c793b80245e433.
//
// Solidity: event FeeCollected(address sender, uint8 fromDomainID, uint8 destinationDomainID, bytes32 resourceID, uint256 fee, address tokenAddress)
func (_FeeHandlerRouter *FeeHandlerRouterFilterer) WatchFeeCollected(opts *bind.WatchOpts, sink chan<- *FeeHandlerRouterFeeCollected) (event.Subscription, error) {

	logs, sub, err := _FeeHandlerRouter.contract.WatchLogs(opts, "FeeCollected")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeHandlerRouterFeeCollected)
				if err := _FeeHandlerRouter.contract.UnpackLog(event, "FeeCollected", log); err != nil {
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

// ParseFeeCollected is a log parse operation binding the contract event 0xbd231b7fa4103e15e7a238c72f07e8aff310701af121895aa6c793b80245e433.
//
// Solidity: event FeeCollected(address sender, uint8 fromDomainID, uint8 destinationDomainID, bytes32 resourceID, uint256 fee, address tokenAddress)
func (_FeeHandlerRouter *FeeHandlerRouterFilterer) ParseFeeCollected(log types.Log) (*FeeHandlerRouterFeeCollected, error) {
	event := new(FeeHandlerRouterFeeCollected)
	if err := _FeeHandlerRouter.contract.UnpackLog(event, "FeeCollected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeHandlerRouterFeeDistributedIterator is returned from FilterFeeDistributed and is used to iterate over the raw logs and unpacked data for FeeDistributed events raised by the FeeHandlerRouter contract.
type FeeHandlerRouterFeeDistributedIterator struct {
	Event *FeeHandlerRouterFeeDistributed // Event containing the contract specifics and raw log

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
func (it *FeeHandlerRouterFeeDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeHandlerRouterFeeDistributed)
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
		it.Event = new(FeeHandlerRouterFeeDistributed)
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
func (it *FeeHandlerRouterFeeDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeHandlerRouterFeeDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeHandlerRouterFeeDistributed represents a FeeDistributed event raised by the FeeHandlerRouter contract.
type FeeHandlerRouterFeeDistributed struct {
	TokenAddress common.Address
	Recipient    common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFeeDistributed is a free log retrieval operation binding the contract event 0xaaa40a232aaf133fdd28f3485f6fdd163514cfadbffa981f3610f42398efe34b.
//
// Solidity: event FeeDistributed(address tokenAddress, address recipient, uint256 amount)
func (_FeeHandlerRouter *FeeHandlerRouterFilterer) FilterFeeDistributed(opts *bind.FilterOpts) (*FeeHandlerRouterFeeDistributedIterator, error) {

	logs, sub, err := _FeeHandlerRouter.contract.FilterLogs(opts, "FeeDistributed")
	if err != nil {
		return nil, err
	}
	return &FeeHandlerRouterFeeDistributedIterator{contract: _FeeHandlerRouter.contract, event: "FeeDistributed", logs: logs, sub: sub}, nil
}

// WatchFeeDistributed is a free log subscription operation binding the contract event 0xaaa40a232aaf133fdd28f3485f6fdd163514cfadbffa981f3610f42398efe34b.
//
// Solidity: event FeeDistributed(address tokenAddress, address recipient, uint256 amount)
func (_FeeHandlerRouter *FeeHandlerRouterFilterer) WatchFeeDistributed(opts *bind.WatchOpts, sink chan<- *FeeHandlerRouterFeeDistributed) (event.Subscription, error) {

	logs, sub, err := _FeeHandlerRouter.contract.WatchLogs(opts, "FeeDistributed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeHandlerRouterFeeDistributed)
				if err := _FeeHandlerRouter.contract.UnpackLog(event, "FeeDistributed", log); err != nil {
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

// ParseFeeDistributed is a log parse operation binding the contract event 0xaaa40a232aaf133fdd28f3485f6fdd163514cfadbffa981f3610f42398efe34b.
//
// Solidity: event FeeDistributed(address tokenAddress, address recipient, uint256 amount)
func (_FeeHandlerRouter *FeeHandlerRouterFilterer) ParseFeeDistributed(log types.Log) (*FeeHandlerRouterFeeDistributed, error) {
	event := new(FeeHandlerRouterFeeDistributed)
	if err := _FeeHandlerRouter.contract.UnpackLog(event, "FeeDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeHandlerRouterRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the FeeHandlerRouter contract.
type FeeHandlerRouterRoleGrantedIterator struct {
	Event *FeeHandlerRouterRoleGranted // Event containing the contract specifics and raw log

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
func (it *FeeHandlerRouterRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeHandlerRouterRoleGranted)
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
		it.Event = new(FeeHandlerRouterRoleGranted)
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
func (it *FeeHandlerRouterRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeHandlerRouterRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeHandlerRouterRoleGranted represents a RoleGranted event raised by the FeeHandlerRouter contract.
type FeeHandlerRouterRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_FeeHandlerRouter *FeeHandlerRouterFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*FeeHandlerRouterRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _FeeHandlerRouter.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &FeeHandlerRouterRoleGrantedIterator{contract: _FeeHandlerRouter.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_FeeHandlerRouter *FeeHandlerRouterFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *FeeHandlerRouterRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _FeeHandlerRouter.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeHandlerRouterRoleGranted)
				if err := _FeeHandlerRouter.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_FeeHandlerRouter *FeeHandlerRouterFilterer) ParseRoleGranted(log types.Log) (*FeeHandlerRouterRoleGranted, error) {
	event := new(FeeHandlerRouterRoleGranted)
	if err := _FeeHandlerRouter.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeHandlerRouterRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the FeeHandlerRouter contract.
type FeeHandlerRouterRoleRevokedIterator struct {
	Event *FeeHandlerRouterRoleRevoked // Event containing the contract specifics and raw log

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
func (it *FeeHandlerRouterRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeHandlerRouterRoleRevoked)
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
		it.Event = new(FeeHandlerRouterRoleRevoked)
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
func (it *FeeHandlerRouterRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeHandlerRouterRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeHandlerRouterRoleRevoked represents a RoleRevoked event raised by the FeeHandlerRouter contract.
type FeeHandlerRouterRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_FeeHandlerRouter *FeeHandlerRouterFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*FeeHandlerRouterRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _FeeHandlerRouter.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &FeeHandlerRouterRoleRevokedIterator{contract: _FeeHandlerRouter.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_FeeHandlerRouter *FeeHandlerRouterFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *FeeHandlerRouterRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _FeeHandlerRouter.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeHandlerRouterRoleRevoked)
				if err := _FeeHandlerRouter.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_FeeHandlerRouter *FeeHandlerRouterFilterer) ParseRoleRevoked(log types.Log) (*FeeHandlerRouterRoleRevoked, error) {
	event := new(FeeHandlerRouterRoleRevoked)
	if err := _FeeHandlerRouter.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
