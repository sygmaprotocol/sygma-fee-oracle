// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package basicFeeHandler

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

// BasicFeeHandlerMetaData contains all meta data concerning the BasicFeeHandler contract.
var BasicFeeHandlerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"FeeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"fromDomainID\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"destinationDomainID\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"FeeCollected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeeDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"fromDomainID\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"destinationDomainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"depositData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"name\":\"calculateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"changeFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"fromDomainID\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"destinationDomainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"depositData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"name\":\"collectFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getRoleMemberIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable[]\",\"name\":\"addrs\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"transferFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b5060405162001da838038062001da8833981810160405281019062000037919062000287565b8073ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1681525050620000806000801b336200008760201b60201c565b50620002b9565b6200009982826200009d60201b60201c565b5050565b620000cb816000808581526020019081526020016000206000016200014060201b62000a011790919060201c565b156200013c57620000e16200017860201b60201c565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b600062000170836000018373ffffffffffffffffffffffffffffffffffffffff1660001b6200018060201b60201c565b905092915050565b600033905090565b6000620001948383620001fa60201b60201c565b620001ef578260000182908060018154018082558091505060019003906000526020600020016000909190919091505582600001805490508360010160008481526020019081526020016000208190555060019050620001f4565b600090505b92915050565b600080836001016000848152602001908152602001600020541415905092915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006200024f8262000222565b9050919050565b620002618162000242565b81146200026d57600080fd5b50565b600081519050620002818162000256565b92915050565b600060208284031215620002a0576200029f6200021d565b5b6000620002b08482850162000270565b91505092915050565b608051611acc620002dc600039600081816104f40152610a330152611acc6000f3fe6080604052600436106100e85760003560e01c80639010d07c1161008a578063c5b37c2211610059578063c5b37c22146102f7578063ca15c87314610322578063d547741f1461035f578063ef4f081f14610388576100e8565b80639010d07c14610229578063904e9b2b1461026657806391d148541461028f578063a217fddf146102cc576100e8565b8063318c136e116100c6578063318c136e1461016f57806336568abe1461019a5780634e0df3f6146101c35780636a1db1bf14610200576100e8565b8063248a9ca3146100ed578063253070651461012a5780632f2ff15d14610146575b600080fd5b3480156100f957600080fd5b50610114600480360381019061010f9190610ea1565b6103c6565b6040516101219190610edd565b60405180910390f35b610144600480360381019061013f9190610ff4565b6103e5565b005b34801561015257600080fd5b5061016d600480360381019061016891906110c3565b61047f565b005b34801561017b57600080fd5b506101846104f2565b6040516101919190611112565b60405180910390f35b3480156101a657600080fd5b506101c160048036038101906101bc91906110c3565b610516565b005b3480156101cf57600080fd5b506101ea60048036038101906101e591906110c3565b610599565b6040516101f79190611146565b60405180910390f35b34801561020c57600080fd5b506102276004803603810190610222919061118d565b6105e9565b005b34801561023557600080fd5b50610250600480360381019061024b91906111ba565b6106bb565b60405161025d9190611112565b60405180910390f35b34801561027257600080fd5b5061028d600480360381019061028891906112a6565b6106ec565b005b34801561029b57600080fd5b506102b660048036038101906102b191906110c3565b610911565b6040516102c39190611342565b60405180910390f35b3480156102d857600080fd5b506102e1610942565b6040516102ee9190610edd565b60405180910390f35b34801561030357600080fd5b5061030c610949565b6040516103199190611146565b60405180910390f35b34801561032e57600080fd5b5061034960048036038101906103449190610ea1565b61094f565b6040516103569190611146565b60405180910390f35b34801561036b57600080fd5b50610386600480360381019061038191906110c3565b610975565b005b34801561039457600080fd5b506103af60048036038101906103aa9190610ff4565b6109e8565b6040516103bd92919061135d565b60405180910390f35b6000806000838152602001908152602001600020600201549050919050565b6103ed610a31565b6001543414610431576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610428906113e3565b60405180910390fd5b7fbd231b7fa4103e15e7a238c72f07e8aff310701af121895aa6c793b80245e43388888888600154600060405161046d96959493929190611412565b60405180910390a15050505050505050565b6104a5600080848152602001908152602001600020600201546104a0610ac1565b610911565b6104e4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104db906114e5565b60405180910390fd5b6104ee8282610ac9565b5050565b7f000000000000000000000000000000000000000000000000000000000000000081565b61051e610ac1565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461058b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161058290611577565b60405180910390fd5b6105958282610b5c565b5050565b600080600084815260200190815260200160002060000160000160010160008373ffffffffffffffffffffffffffffffffffffffff1660001b815260200190815260200160002054905092915050565b6105f66000801b33610911565b610635576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161062c906115e3565b60405180910390fd5b80600154141561067a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106719061164f565b60405180910390fd5b806001819055507f6bbc57480a46553fa4d156ce702beef5f3ad66303b0ed1a5d4cb44966c6584c3816040516106b09190611146565b60405180910390a150565b60006106e482600080868152602001908152602001600020600001610bef90919063ffffffff16565b905092915050565b6106f96000801b33610911565b610738576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161072f906115e3565b60405180910390fd5b818190508484905014610780576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610777906116bb565b60405180910390fd5b60005b8484905081101561090a5760008585838181106107a3576107a26116db565b5b90506020020160208101906107b89190611748565b73ffffffffffffffffffffffffffffffffffffffff168484848181106107e1576107e06116db565b5b905060200201356040516107f4906117a6565b60006040518083038185875af1925050503d8060008114610831576040519150601f19603f3d011682016040523d82523d6000602084013e610836565b606091505b505090508061087a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161087190611807565b60405180910390fd5b7faaa40a232aaf133fdd28f3485f6fdd163514cfadbffa981f3610f42398efe34b60008787858181106108b0576108af6116db565b5b90506020020160208101906108c59190611748565b8686868181106108d8576108d76116db565b5b905060200201356040516108ee93929190611886565b60405180910390a1508080610902906118ec565b915050610783565b5050505050565b600061093a82600080868152602001908152602001600020600001610c0990919063ffffffff16565b905092915050565b6000801b81565b60015481565b600061096e600080848152602001908152602001600020600001610c39565b9050919050565b61099b60008084815260200190815260200160002060020154610996610ac1565b610911565b6109da576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109d1906119a7565b60405180910390fd5b6109e48282610b5c565b5050565b6000806001546000915091509850989650505050505050565b6000610a29836000018373ffffffffffffffffffffffffffffffffffffffff1660001b610c4e565b905092915050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610abf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ab690611a13565b60405180910390fd5b565b600033905090565b610af081600080858152602001908152602001600020600001610a0190919063ffffffff16565b15610b5857610afd610ac1565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b610b8381600080858152602001908152602001600020600001610cbe90919063ffffffff16565b15610beb57610b90610ac1565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45b5050565b6000610bfe8360000183610cee565b60001c905092915050565b6000610c31836000018373ffffffffffffffffffffffffffffffffffffffff1660001b610d19565b905092915050565b6000610c4782600001610d3c565b9050919050565b6000610c5a8383610d19565b610cb3578260000182908060018154018082558091505060019003906000526020600020016000909190919091505582600001805490508360010160008481526020019081526020016000208190555060019050610cb8565b600090505b92915050565b6000610ce6836000018373ffffffffffffffffffffffffffffffffffffffff1660001b610d4d565b905092915050565b6000826000018281548110610d0657610d056116db565b5b9060005260206000200154905092915050565b600080836001016000848152602001908152602001600020541415905092915050565b600081600001805490509050919050565b60008083600101600084815260200190815260200160002054905060008114610e55576000600182610d7f9190611a33565b9050600060018660000180549050610d979190611a33565b9050818114610e06576000866000018281548110610db857610db76116db565b5b9060005260206000200154905080876000018481548110610ddc57610ddb6116db565b5b90600052602060002001819055508387600101600083815260200190815260200160002081905550505b85600001805480610e1a57610e19611a67565b5b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610e5b565b60009150505b92915050565b600080fd5b600080fd5b6000819050919050565b610e7e81610e6b565b8114610e8957600080fd5b50565b600081359050610e9b81610e75565b92915050565b600060208284031215610eb757610eb6610e61565b5b6000610ec584828501610e8c565b91505092915050565b610ed781610e6b565b82525050565b6000602082019050610ef26000830184610ece565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610f2382610ef8565b9050919050565b610f3381610f18565b8114610f3e57600080fd5b50565b600081359050610f5081610f2a565b92915050565b600060ff82169050919050565b610f6c81610f56565b8114610f7757600080fd5b50565b600081359050610f8981610f63565b92915050565b600080fd5b600080fd5b600080fd5b60008083601f840112610fb457610fb3610f8f565b5b8235905067ffffffffffffffff811115610fd157610fd0610f94565b5b602083019150836001820283011115610fed57610fec610f99565b5b9250929050565b60008060008060008060008060c0898b03121561101457611013610e61565b5b60006110228b828c01610f41565b98505060206110338b828c01610f7a565b97505060406110448b828c01610f7a565b96505060606110558b828c01610e8c565b955050608089013567ffffffffffffffff81111561107657611075610e66565b5b6110828b828c01610f9e565b945094505060a089013567ffffffffffffffff8111156110a5576110a4610e66565b5b6110b18b828c01610f9e565b92509250509295985092959890939650565b600080604083850312156110da576110d9610e61565b5b60006110e885828601610e8c565b92505060206110f985828601610f41565b9150509250929050565b61110c81610f18565b82525050565b60006020820190506111276000830184611103565b92915050565b6000819050919050565b6111408161112d565b82525050565b600060208201905061115b6000830184611137565b92915050565b61116a8161112d565b811461117557600080fd5b50565b60008135905061118781611161565b92915050565b6000602082840312156111a3576111a2610e61565b5b60006111b184828501611178565b91505092915050565b600080604083850312156111d1576111d0610e61565b5b60006111df85828601610e8c565b92505060206111f085828601611178565b9150509250929050565b60008083601f8401126112105761120f610f8f565b5b8235905067ffffffffffffffff81111561122d5761122c610f94565b5b60208301915083602082028301111561124957611248610f99565b5b9250929050565b60008083601f84011261126657611265610f8f565b5b8235905067ffffffffffffffff81111561128357611282610f94565b5b60208301915083602082028301111561129f5761129e610f99565b5b9250929050565b600080600080604085870312156112c0576112bf610e61565b5b600085013567ffffffffffffffff8111156112de576112dd610e66565b5b6112ea878288016111fa565b9450945050602085013567ffffffffffffffff81111561130d5761130c610e66565b5b61131987828801611250565b925092505092959194509250565b60008115159050919050565b61133c81611327565b82525050565b60006020820190506113576000830184611333565b92915050565b60006040820190506113726000830185611137565b61137f6020830184611103565b9392505050565b600082825260208201905092915050565b7f496e636f72726563742066656520737570706c69656400000000000000000000600082015250565b60006113cd601683611386565b91506113d882611397565b602082019050919050565b600060208201905081810360008301526113fc816113c0565b9050919050565b61140c81610f56565b82525050565b600060c0820190506114276000830189611103565b6114346020830188611403565b6114416040830187611403565b61144e6060830186610ece565b61145b6080830185611137565b61146860a0830184611103565b979650505050505050565b7f416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e60008201527f2061646d696e20746f206772616e740000000000000000000000000000000000602082015250565b60006114cf602f83611386565b91506114da82611473565b604082019050919050565b600060208201905081810360008301526114fe816114c2565b9050919050565b7f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560008201527f20726f6c657320666f722073656c660000000000000000000000000000000000602082015250565b6000611561602f83611386565b915061156c82611505565b604082019050919050565b6000602082019050818103600083015261159081611554565b9050919050565b7f73656e64657220646f65736e277420686176652061646d696e20726f6c650000600082015250565b60006115cd601e83611386565b91506115d882611597565b602082019050919050565b600060208201905081810360008301526115fc816115c0565b9050919050565b7f43757272656e742066656520697320657175616c20746f206e65772066656500600082015250565b6000611639601f83611386565b915061164482611603565b602082019050919050565b600060208201905081810360008301526116688161162c565b9050919050565b7f61646472735b5d2c20616d6f756e74735b5d3a2064696666206c656e67746800600082015250565b60006116a5601f83611386565b91506116b08261166f565b602082019050919050565b600060208201905081810360008301526116d481611698565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600061171582610ef8565b9050919050565b6117258161170a565b811461173057600080fd5b50565b6000813590506117428161171c565b92915050565b60006020828403121561175e5761175d610e61565b5b600061176c84828501611733565b91505092915050565b600081905092915050565b50565b6000611790600083611775565b915061179b82611780565b600082019050919050565b60006117b182611783565b9150819050919050565b7f466565206574686572207472616e73666572206661696c656400000000000000600082015250565b60006117f1601983611386565b91506117fc826117bb565b602082019050919050565b60006020820190508181036000830152611820816117e4565b9050919050565b6000819050919050565b600061184c61184761184284610ef8565b611827565b610ef8565b9050919050565b600061185e82611831565b9050919050565b600061187082611853565b9050919050565b61188081611865565b82525050565b600060608201905061189b6000830186611103565b6118a86020830185611877565b6118b56040830184611137565b949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006118f78261112d565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82141561192a576119296118bd565b5b600182019050919050565b7f416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e60008201527f2061646d696e20746f207265766f6b6500000000000000000000000000000000602082015250565b6000611991603083611386565b915061199c82611935565b604082019050919050565b600060208201905081810360008301526119c081611984565b9050919050565b7f73656e646572206d7573742062652062726964676520636f6e74726163740000600082015250565b60006119fd601e83611386565b9150611a08826119c7565b602082019050919050565b60006020820190508181036000830152611a2c816119f0565b9050919050565b6000611a3e8261112d565b9150611a498361112d565b925082821015611a5c57611a5b6118bd565b5b828203905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea2646970667358221220b23268692ca3f76bf385585c475760e821763b4673b29c5d3b9e240dc7073be464736f6c634300080b0033",
}

// BasicFeeHandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use BasicFeeHandlerMetaData.ABI instead.
var BasicFeeHandlerABI = BasicFeeHandlerMetaData.ABI

// BasicFeeHandlerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BasicFeeHandlerMetaData.Bin instead.
var BasicFeeHandlerBin = BasicFeeHandlerMetaData.Bin

// DeployBasicFeeHandler deploys a new Ethereum contract, binding an instance of BasicFeeHandler to it.
func DeployBasicFeeHandler(auth *bind.TransactOpts, backend bind.ContractBackend, bridgeAddress common.Address) (common.Address, *types.Transaction, *BasicFeeHandler, error) {
	parsed, err := BasicFeeHandlerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BasicFeeHandlerBin), backend, bridgeAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BasicFeeHandler{BasicFeeHandlerCaller: BasicFeeHandlerCaller{contract: contract}, BasicFeeHandlerTransactor: BasicFeeHandlerTransactor{contract: contract}, BasicFeeHandlerFilterer: BasicFeeHandlerFilterer{contract: contract}}, nil
}

// BasicFeeHandler is an auto generated Go binding around an Ethereum contract.
type BasicFeeHandler struct {
	BasicFeeHandlerCaller     // Read-only binding to the contract
	BasicFeeHandlerTransactor // Write-only binding to the contract
	BasicFeeHandlerFilterer   // Log filterer for contract events
}

// BasicFeeHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasicFeeHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicFeeHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasicFeeHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicFeeHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasicFeeHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicFeeHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasicFeeHandlerSession struct {
	Contract     *BasicFeeHandler  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasicFeeHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasicFeeHandlerCallerSession struct {
	Contract *BasicFeeHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// BasicFeeHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasicFeeHandlerTransactorSession struct {
	Contract     *BasicFeeHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// BasicFeeHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasicFeeHandlerRaw struct {
	Contract *BasicFeeHandler // Generic contract binding to access the raw methods on
}

// BasicFeeHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasicFeeHandlerCallerRaw struct {
	Contract *BasicFeeHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// BasicFeeHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasicFeeHandlerTransactorRaw struct {
	Contract *BasicFeeHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasicFeeHandler creates a new instance of BasicFeeHandler, bound to a specific deployed contract.
func NewBasicFeeHandler(address common.Address, backend bind.ContractBackend) (*BasicFeeHandler, error) {
	contract, err := bindBasicFeeHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasicFeeHandler{BasicFeeHandlerCaller: BasicFeeHandlerCaller{contract: contract}, BasicFeeHandlerTransactor: BasicFeeHandlerTransactor{contract: contract}, BasicFeeHandlerFilterer: BasicFeeHandlerFilterer{contract: contract}}, nil
}

// NewBasicFeeHandlerCaller creates a new read-only instance of BasicFeeHandler, bound to a specific deployed contract.
func NewBasicFeeHandlerCaller(address common.Address, caller bind.ContractCaller) (*BasicFeeHandlerCaller, error) {
	contract, err := bindBasicFeeHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasicFeeHandlerCaller{contract: contract}, nil
}

// NewBasicFeeHandlerTransactor creates a new write-only instance of BasicFeeHandler, bound to a specific deployed contract.
func NewBasicFeeHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*BasicFeeHandlerTransactor, error) {
	contract, err := bindBasicFeeHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasicFeeHandlerTransactor{contract: contract}, nil
}

// NewBasicFeeHandlerFilterer creates a new log filterer instance of BasicFeeHandler, bound to a specific deployed contract.
func NewBasicFeeHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*BasicFeeHandlerFilterer, error) {
	contract, err := bindBasicFeeHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasicFeeHandlerFilterer{contract: contract}, nil
}

// bindBasicFeeHandler binds a generic wrapper to an already deployed contract.
func bindBasicFeeHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BasicFeeHandlerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicFeeHandler *BasicFeeHandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasicFeeHandler.Contract.BasicFeeHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicFeeHandler *BasicFeeHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicFeeHandler.Contract.BasicFeeHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicFeeHandler *BasicFeeHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicFeeHandler.Contract.BasicFeeHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicFeeHandler *BasicFeeHandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasicFeeHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicFeeHandler *BasicFeeHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicFeeHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicFeeHandler *BasicFeeHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicFeeHandler.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BasicFeeHandler *BasicFeeHandlerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BasicFeeHandler.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BasicFeeHandler *BasicFeeHandlerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BasicFeeHandler.Contract.DEFAULTADMINROLE(&_BasicFeeHandler.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BasicFeeHandler *BasicFeeHandlerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BasicFeeHandler.Contract.DEFAULTADMINROLE(&_BasicFeeHandler.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_BasicFeeHandler *BasicFeeHandlerCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasicFeeHandler.contract.Call(opts, &out, "_bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_BasicFeeHandler *BasicFeeHandlerSession) BridgeAddress() (common.Address, error) {
	return _BasicFeeHandler.Contract.BridgeAddress(&_BasicFeeHandler.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_BasicFeeHandler *BasicFeeHandlerCallerSession) BridgeAddress() (common.Address, error) {
	return _BasicFeeHandler.Contract.BridgeAddress(&_BasicFeeHandler.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xc5b37c22.
//
// Solidity: function _fee() view returns(uint256)
func (_BasicFeeHandler *BasicFeeHandlerCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BasicFeeHandler.contract.Call(opts, &out, "_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xc5b37c22.
//
// Solidity: function _fee() view returns(uint256)
func (_BasicFeeHandler *BasicFeeHandlerSession) Fee() (*big.Int, error) {
	return _BasicFeeHandler.Contract.Fee(&_BasicFeeHandler.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xc5b37c22.
//
// Solidity: function _fee() view returns(uint256)
func (_BasicFeeHandler *BasicFeeHandlerCallerSession) Fee() (*big.Int, error) {
	return _BasicFeeHandler.Contract.Fee(&_BasicFeeHandler.CallOpts)
}

// CalculateFee is a free data retrieval call binding the contract method 0xef4f081f.
//
// Solidity: function calculateFee(address sender, uint8 fromDomainID, uint8 destinationDomainID, bytes32 resourceID, bytes depositData, bytes feeData) view returns(uint256, address)
func (_BasicFeeHandler *BasicFeeHandlerCaller) CalculateFee(opts *bind.CallOpts, sender common.Address, fromDomainID uint8, destinationDomainID uint8, resourceID [32]byte, depositData []byte, feeData []byte) (*big.Int, common.Address, error) {
	var out []interface{}
	err := _BasicFeeHandler.contract.Call(opts, &out, "calculateFee", sender, fromDomainID, destinationDomainID, resourceID, depositData, feeData)

	if err != nil {
		return *new(*big.Int), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// CalculateFee is a free data retrieval call binding the contract method 0xef4f081f.
//
// Solidity: function calculateFee(address sender, uint8 fromDomainID, uint8 destinationDomainID, bytes32 resourceID, bytes depositData, bytes feeData) view returns(uint256, address)
func (_BasicFeeHandler *BasicFeeHandlerSession) CalculateFee(sender common.Address, fromDomainID uint8, destinationDomainID uint8, resourceID [32]byte, depositData []byte, feeData []byte) (*big.Int, common.Address, error) {
	return _BasicFeeHandler.Contract.CalculateFee(&_BasicFeeHandler.CallOpts, sender, fromDomainID, destinationDomainID, resourceID, depositData, feeData)
}

// CalculateFee is a free data retrieval call binding the contract method 0xef4f081f.
//
// Solidity: function calculateFee(address sender, uint8 fromDomainID, uint8 destinationDomainID, bytes32 resourceID, bytes depositData, bytes feeData) view returns(uint256, address)
func (_BasicFeeHandler *BasicFeeHandlerCallerSession) CalculateFee(sender common.Address, fromDomainID uint8, destinationDomainID uint8, resourceID [32]byte, depositData []byte, feeData []byte) (*big.Int, common.Address, error) {
	return _BasicFeeHandler.Contract.CalculateFee(&_BasicFeeHandler.CallOpts, sender, fromDomainID, destinationDomainID, resourceID, depositData, feeData)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BasicFeeHandler *BasicFeeHandlerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _BasicFeeHandler.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BasicFeeHandler *BasicFeeHandlerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BasicFeeHandler.Contract.GetRoleAdmin(&_BasicFeeHandler.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BasicFeeHandler *BasicFeeHandlerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BasicFeeHandler.Contract.GetRoleAdmin(&_BasicFeeHandler.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_BasicFeeHandler *BasicFeeHandlerCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BasicFeeHandler.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_BasicFeeHandler *BasicFeeHandlerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _BasicFeeHandler.Contract.GetRoleMember(&_BasicFeeHandler.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_BasicFeeHandler *BasicFeeHandlerCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _BasicFeeHandler.Contract.GetRoleMember(&_BasicFeeHandler.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_BasicFeeHandler *BasicFeeHandlerCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _BasicFeeHandler.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_BasicFeeHandler *BasicFeeHandlerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _BasicFeeHandler.Contract.GetRoleMemberCount(&_BasicFeeHandler.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_BasicFeeHandler *BasicFeeHandlerCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _BasicFeeHandler.Contract.GetRoleMemberCount(&_BasicFeeHandler.CallOpts, role)
}

// GetRoleMemberIndex is a free data retrieval call binding the contract method 0x4e0df3f6.
//
// Solidity: function getRoleMemberIndex(bytes32 role, address account) view returns(uint256)
func (_BasicFeeHandler *BasicFeeHandlerCaller) GetRoleMemberIndex(opts *bind.CallOpts, role [32]byte, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BasicFeeHandler.contract.Call(opts, &out, "getRoleMemberIndex", role, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberIndex is a free data retrieval call binding the contract method 0x4e0df3f6.
//
// Solidity: function getRoleMemberIndex(bytes32 role, address account) view returns(uint256)
func (_BasicFeeHandler *BasicFeeHandlerSession) GetRoleMemberIndex(role [32]byte, account common.Address) (*big.Int, error) {
	return _BasicFeeHandler.Contract.GetRoleMemberIndex(&_BasicFeeHandler.CallOpts, role, account)
}

// GetRoleMemberIndex is a free data retrieval call binding the contract method 0x4e0df3f6.
//
// Solidity: function getRoleMemberIndex(bytes32 role, address account) view returns(uint256)
func (_BasicFeeHandler *BasicFeeHandlerCallerSession) GetRoleMemberIndex(role [32]byte, account common.Address) (*big.Int, error) {
	return _BasicFeeHandler.Contract.GetRoleMemberIndex(&_BasicFeeHandler.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BasicFeeHandler *BasicFeeHandlerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _BasicFeeHandler.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BasicFeeHandler *BasicFeeHandlerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BasicFeeHandler.Contract.HasRole(&_BasicFeeHandler.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BasicFeeHandler *BasicFeeHandlerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BasicFeeHandler.Contract.HasRole(&_BasicFeeHandler.CallOpts, role, account)
}

// ChangeFee is a paid mutator transaction binding the contract method 0x6a1db1bf.
//
// Solidity: function changeFee(uint256 newFee) returns()
func (_BasicFeeHandler *BasicFeeHandlerTransactor) ChangeFee(opts *bind.TransactOpts, newFee *big.Int) (*types.Transaction, error) {
	return _BasicFeeHandler.contract.Transact(opts, "changeFee", newFee)
}

// ChangeFee is a paid mutator transaction binding the contract method 0x6a1db1bf.
//
// Solidity: function changeFee(uint256 newFee) returns()
func (_BasicFeeHandler *BasicFeeHandlerSession) ChangeFee(newFee *big.Int) (*types.Transaction, error) {
	return _BasicFeeHandler.Contract.ChangeFee(&_BasicFeeHandler.TransactOpts, newFee)
}

// ChangeFee is a paid mutator transaction binding the contract method 0x6a1db1bf.
//
// Solidity: function changeFee(uint256 newFee) returns()
func (_BasicFeeHandler *BasicFeeHandlerTransactorSession) ChangeFee(newFee *big.Int) (*types.Transaction, error) {
	return _BasicFeeHandler.Contract.ChangeFee(&_BasicFeeHandler.TransactOpts, newFee)
}

// CollectFee is a paid mutator transaction binding the contract method 0x25307065.
//
// Solidity: function collectFee(address sender, uint8 fromDomainID, uint8 destinationDomainID, bytes32 resourceID, bytes depositData, bytes feeData) payable returns()
func (_BasicFeeHandler *BasicFeeHandlerTransactor) CollectFee(opts *bind.TransactOpts, sender common.Address, fromDomainID uint8, destinationDomainID uint8, resourceID [32]byte, depositData []byte, feeData []byte) (*types.Transaction, error) {
	return _BasicFeeHandler.contract.Transact(opts, "collectFee", sender, fromDomainID, destinationDomainID, resourceID, depositData, feeData)
}

// CollectFee is a paid mutator transaction binding the contract method 0x25307065.
//
// Solidity: function collectFee(address sender, uint8 fromDomainID, uint8 destinationDomainID, bytes32 resourceID, bytes depositData, bytes feeData) payable returns()
func (_BasicFeeHandler *BasicFeeHandlerSession) CollectFee(sender common.Address, fromDomainID uint8, destinationDomainID uint8, resourceID [32]byte, depositData []byte, feeData []byte) (*types.Transaction, error) {
	return _BasicFeeHandler.Contract.CollectFee(&_BasicFeeHandler.TransactOpts, sender, fromDomainID, destinationDomainID, resourceID, depositData, feeData)
}

// CollectFee is a paid mutator transaction binding the contract method 0x25307065.
//
// Solidity: function collectFee(address sender, uint8 fromDomainID, uint8 destinationDomainID, bytes32 resourceID, bytes depositData, bytes feeData) payable returns()
func (_BasicFeeHandler *BasicFeeHandlerTransactorSession) CollectFee(sender common.Address, fromDomainID uint8, destinationDomainID uint8, resourceID [32]byte, depositData []byte, feeData []byte) (*types.Transaction, error) {
	return _BasicFeeHandler.Contract.CollectFee(&_BasicFeeHandler.TransactOpts, sender, fromDomainID, destinationDomainID, resourceID, depositData, feeData)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BasicFeeHandler *BasicFeeHandlerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BasicFeeHandler.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BasicFeeHandler *BasicFeeHandlerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BasicFeeHandler.Contract.GrantRole(&_BasicFeeHandler.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BasicFeeHandler *BasicFeeHandlerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BasicFeeHandler.Contract.GrantRole(&_BasicFeeHandler.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_BasicFeeHandler *BasicFeeHandlerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BasicFeeHandler.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_BasicFeeHandler *BasicFeeHandlerSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BasicFeeHandler.Contract.RenounceRole(&_BasicFeeHandler.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_BasicFeeHandler *BasicFeeHandlerTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BasicFeeHandler.Contract.RenounceRole(&_BasicFeeHandler.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BasicFeeHandler *BasicFeeHandlerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BasicFeeHandler.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BasicFeeHandler *BasicFeeHandlerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BasicFeeHandler.Contract.RevokeRole(&_BasicFeeHandler.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BasicFeeHandler *BasicFeeHandlerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BasicFeeHandler.Contract.RevokeRole(&_BasicFeeHandler.TransactOpts, role, account)
}

// TransferFee is a paid mutator transaction binding the contract method 0x904e9b2b.
//
// Solidity: function transferFee(address[] addrs, uint256[] amounts) returns()
func (_BasicFeeHandler *BasicFeeHandlerTransactor) TransferFee(opts *bind.TransactOpts, addrs []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _BasicFeeHandler.contract.Transact(opts, "transferFee", addrs, amounts)
}

// TransferFee is a paid mutator transaction binding the contract method 0x904e9b2b.
//
// Solidity: function transferFee(address[] addrs, uint256[] amounts) returns()
func (_BasicFeeHandler *BasicFeeHandlerSession) TransferFee(addrs []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _BasicFeeHandler.Contract.TransferFee(&_BasicFeeHandler.TransactOpts, addrs, amounts)
}

// TransferFee is a paid mutator transaction binding the contract method 0x904e9b2b.
//
// Solidity: function transferFee(address[] addrs, uint256[] amounts) returns()
func (_BasicFeeHandler *BasicFeeHandlerTransactorSession) TransferFee(addrs []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _BasicFeeHandler.Contract.TransferFee(&_BasicFeeHandler.TransactOpts, addrs, amounts)
}

// BasicFeeHandlerFeeChangedIterator is returned from FilterFeeChanged and is used to iterate over the raw logs and unpacked data for FeeChanged events raised by the BasicFeeHandler contract.
type BasicFeeHandlerFeeChangedIterator struct {
	Event *BasicFeeHandlerFeeChanged // Event containing the contract specifics and raw log

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
func (it *BasicFeeHandlerFeeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasicFeeHandlerFeeChanged)
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
		it.Event = new(BasicFeeHandlerFeeChanged)
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
func (it *BasicFeeHandlerFeeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasicFeeHandlerFeeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasicFeeHandlerFeeChanged represents a FeeChanged event raised by the BasicFeeHandler contract.
type BasicFeeHandlerFeeChanged struct {
	NewFee *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeeChanged is a free log retrieval operation binding the contract event 0x6bbc57480a46553fa4d156ce702beef5f3ad66303b0ed1a5d4cb44966c6584c3.
//
// Solidity: event FeeChanged(uint256 newFee)
func (_BasicFeeHandler *BasicFeeHandlerFilterer) FilterFeeChanged(opts *bind.FilterOpts) (*BasicFeeHandlerFeeChangedIterator, error) {

	logs, sub, err := _BasicFeeHandler.contract.FilterLogs(opts, "FeeChanged")
	if err != nil {
		return nil, err
	}
	return &BasicFeeHandlerFeeChangedIterator{contract: _BasicFeeHandler.contract, event: "FeeChanged", logs: logs, sub: sub}, nil
}

// WatchFeeChanged is a free log subscription operation binding the contract event 0x6bbc57480a46553fa4d156ce702beef5f3ad66303b0ed1a5d4cb44966c6584c3.
//
// Solidity: event FeeChanged(uint256 newFee)
func (_BasicFeeHandler *BasicFeeHandlerFilterer) WatchFeeChanged(opts *bind.WatchOpts, sink chan<- *BasicFeeHandlerFeeChanged) (event.Subscription, error) {

	logs, sub, err := _BasicFeeHandler.contract.WatchLogs(opts, "FeeChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasicFeeHandlerFeeChanged)
				if err := _BasicFeeHandler.contract.UnpackLog(event, "FeeChanged", log); err != nil {
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
func (_BasicFeeHandler *BasicFeeHandlerFilterer) ParseFeeChanged(log types.Log) (*BasicFeeHandlerFeeChanged, error) {
	event := new(BasicFeeHandlerFeeChanged)
	if err := _BasicFeeHandler.contract.UnpackLog(event, "FeeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasicFeeHandlerFeeCollectedIterator is returned from FilterFeeCollected and is used to iterate over the raw logs and unpacked data for FeeCollected events raised by the BasicFeeHandler contract.
type BasicFeeHandlerFeeCollectedIterator struct {
	Event *BasicFeeHandlerFeeCollected // Event containing the contract specifics and raw log

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
func (it *BasicFeeHandlerFeeCollectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasicFeeHandlerFeeCollected)
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
		it.Event = new(BasicFeeHandlerFeeCollected)
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
func (it *BasicFeeHandlerFeeCollectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasicFeeHandlerFeeCollectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasicFeeHandlerFeeCollected represents a FeeCollected event raised by the BasicFeeHandler contract.
type BasicFeeHandlerFeeCollected struct {
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
func (_BasicFeeHandler *BasicFeeHandlerFilterer) FilterFeeCollected(opts *bind.FilterOpts) (*BasicFeeHandlerFeeCollectedIterator, error) {

	logs, sub, err := _BasicFeeHandler.contract.FilterLogs(opts, "FeeCollected")
	if err != nil {
		return nil, err
	}
	return &BasicFeeHandlerFeeCollectedIterator{contract: _BasicFeeHandler.contract, event: "FeeCollected", logs: logs, sub: sub}, nil
}

// WatchFeeCollected is a free log subscription operation binding the contract event 0xbd231b7fa4103e15e7a238c72f07e8aff310701af121895aa6c793b80245e433.
//
// Solidity: event FeeCollected(address sender, uint8 fromDomainID, uint8 destinationDomainID, bytes32 resourceID, uint256 fee, address tokenAddress)
func (_BasicFeeHandler *BasicFeeHandlerFilterer) WatchFeeCollected(opts *bind.WatchOpts, sink chan<- *BasicFeeHandlerFeeCollected) (event.Subscription, error) {

	logs, sub, err := _BasicFeeHandler.contract.WatchLogs(opts, "FeeCollected")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasicFeeHandlerFeeCollected)
				if err := _BasicFeeHandler.contract.UnpackLog(event, "FeeCollected", log); err != nil {
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
func (_BasicFeeHandler *BasicFeeHandlerFilterer) ParseFeeCollected(log types.Log) (*BasicFeeHandlerFeeCollected, error) {
	event := new(BasicFeeHandlerFeeCollected)
	if err := _BasicFeeHandler.contract.UnpackLog(event, "FeeCollected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasicFeeHandlerFeeDistributedIterator is returned from FilterFeeDistributed and is used to iterate over the raw logs and unpacked data for FeeDistributed events raised by the BasicFeeHandler contract.
type BasicFeeHandlerFeeDistributedIterator struct {
	Event *BasicFeeHandlerFeeDistributed // Event containing the contract specifics and raw log

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
func (it *BasicFeeHandlerFeeDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasicFeeHandlerFeeDistributed)
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
		it.Event = new(BasicFeeHandlerFeeDistributed)
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
func (it *BasicFeeHandlerFeeDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasicFeeHandlerFeeDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasicFeeHandlerFeeDistributed represents a FeeDistributed event raised by the BasicFeeHandler contract.
type BasicFeeHandlerFeeDistributed struct {
	TokenAddress common.Address
	Recipient    common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFeeDistributed is a free log retrieval operation binding the contract event 0xaaa40a232aaf133fdd28f3485f6fdd163514cfadbffa981f3610f42398efe34b.
//
// Solidity: event FeeDistributed(address tokenAddress, address recipient, uint256 amount)
func (_BasicFeeHandler *BasicFeeHandlerFilterer) FilterFeeDistributed(opts *bind.FilterOpts) (*BasicFeeHandlerFeeDistributedIterator, error) {

	logs, sub, err := _BasicFeeHandler.contract.FilterLogs(opts, "FeeDistributed")
	if err != nil {
		return nil, err
	}
	return &BasicFeeHandlerFeeDistributedIterator{contract: _BasicFeeHandler.contract, event: "FeeDistributed", logs: logs, sub: sub}, nil
}

// WatchFeeDistributed is a free log subscription operation binding the contract event 0xaaa40a232aaf133fdd28f3485f6fdd163514cfadbffa981f3610f42398efe34b.
//
// Solidity: event FeeDistributed(address tokenAddress, address recipient, uint256 amount)
func (_BasicFeeHandler *BasicFeeHandlerFilterer) WatchFeeDistributed(opts *bind.WatchOpts, sink chan<- *BasicFeeHandlerFeeDistributed) (event.Subscription, error) {

	logs, sub, err := _BasicFeeHandler.contract.WatchLogs(opts, "FeeDistributed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasicFeeHandlerFeeDistributed)
				if err := _BasicFeeHandler.contract.UnpackLog(event, "FeeDistributed", log); err != nil {
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
func (_BasicFeeHandler *BasicFeeHandlerFilterer) ParseFeeDistributed(log types.Log) (*BasicFeeHandlerFeeDistributed, error) {
	event := new(BasicFeeHandlerFeeDistributed)
	if err := _BasicFeeHandler.contract.UnpackLog(event, "FeeDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasicFeeHandlerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the BasicFeeHandler contract.
type BasicFeeHandlerRoleGrantedIterator struct {
	Event *BasicFeeHandlerRoleGranted // Event containing the contract specifics and raw log

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
func (it *BasicFeeHandlerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasicFeeHandlerRoleGranted)
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
		it.Event = new(BasicFeeHandlerRoleGranted)
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
func (it *BasicFeeHandlerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasicFeeHandlerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasicFeeHandlerRoleGranted represents a RoleGranted event raised by the BasicFeeHandler contract.
type BasicFeeHandlerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BasicFeeHandler *BasicFeeHandlerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BasicFeeHandlerRoleGrantedIterator, error) {

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

	logs, sub, err := _BasicFeeHandler.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BasicFeeHandlerRoleGrantedIterator{contract: _BasicFeeHandler.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BasicFeeHandler *BasicFeeHandlerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BasicFeeHandlerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BasicFeeHandler.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasicFeeHandlerRoleGranted)
				if err := _BasicFeeHandler.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_BasicFeeHandler *BasicFeeHandlerFilterer) ParseRoleGranted(log types.Log) (*BasicFeeHandlerRoleGranted, error) {
	event := new(BasicFeeHandlerRoleGranted)
	if err := _BasicFeeHandler.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasicFeeHandlerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the BasicFeeHandler contract.
type BasicFeeHandlerRoleRevokedIterator struct {
	Event *BasicFeeHandlerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *BasicFeeHandlerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasicFeeHandlerRoleRevoked)
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
		it.Event = new(BasicFeeHandlerRoleRevoked)
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
func (it *BasicFeeHandlerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasicFeeHandlerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasicFeeHandlerRoleRevoked represents a RoleRevoked event raised by the BasicFeeHandler contract.
type BasicFeeHandlerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BasicFeeHandler *BasicFeeHandlerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BasicFeeHandlerRoleRevokedIterator, error) {

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

	logs, sub, err := _BasicFeeHandler.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BasicFeeHandlerRoleRevokedIterator{contract: _BasicFeeHandler.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BasicFeeHandler *BasicFeeHandlerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BasicFeeHandlerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BasicFeeHandler.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasicFeeHandlerRoleRevoked)
				if err := _BasicFeeHandler.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_BasicFeeHandler *BasicFeeHandlerFilterer) ParseRoleRevoked(log types.Log) (*BasicFeeHandlerRoleRevoked, error) {
	event := new(BasicFeeHandlerRoleRevoked)
	if err := _BasicFeeHandler.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
