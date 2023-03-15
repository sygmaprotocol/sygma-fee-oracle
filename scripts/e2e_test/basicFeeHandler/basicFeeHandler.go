// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package basicFeeHandler

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

// BasicFeeHandlerABI is the input ABI used to generate the binding from.
const BasicFeeHandlerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeHandlerRouterAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"FeeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"fromDomainID\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"destinationDomainID\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"FeeCollected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeeDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_feeHandlerRouterAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"fromDomainID\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"destinationDomainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"depositData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"name\":\"calculateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"changeFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"fromDomainID\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"destinationDomainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"depositData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"name\":\"collectFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getRoleMemberIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"renounceAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable[]\",\"name\":\"addrs\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"transferFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BasicFeeHandlerBin is the compiled bytecode used for deploying new contracts.
var BasicFeeHandlerBin = "0x60c06040523480156200001157600080fd5b506040516200203f3803806200203f8339818101604052810190620000379190620002bc565b8173ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250508073ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff1681525050620000b46000801b33620000bc60201b60201c565b505062000303565b620000ce8282620000d260201b60201c565b5050565b62000100816000808581526020019081526020016000206000016200017560201b62000b281790919060201c565b15620001715762000116620001ad60201b60201c565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b6000620001a5836000018373ffffffffffffffffffffffffffffffffffffffff1660001b620001b560201b60201c565b905092915050565b600033905090565b6000620001c983836200022f60201b60201c565b6200022457826000018290806001815401808255809150506001900390600052602060002001600090919091909150558260000180549050836001016000848152602001908152602001600020819055506001905062000229565b600090505b92915050565b600080836001016000848152602001908152602001600020541415905092915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620002848262000257565b9050919050565b620002968162000277565b8114620002a257600080fd5b50565b600081519050620002b6816200028b565b92915050565b60008060408385031215620002d657620002d562000252565b5b6000620002e685828601620002a5565b9250506020620002f985828601620002a5565b9150509250929050565b60805160a051611d0862000337600039600081816107c00152610baf01526000818161055e0152610b5a0152611d086000f3fe6080604052600436106100fe5760003560e01c8063745e6b6111610095578063a217fddf11610064578063a217fddf14610336578063c5b37c2214610361578063ca15c8731461038c578063d547741f146103c9578063ef4f081f146103f2576100fe565b8063745e6b61146102685780639010d07c14610293578063904e9b2b146102d057806391d14854146102f9576100fe565b806336568abe116100d157806336568abe146101b05780634e0df3f6146101d95780635e1fab0f146102165780636a1db1bf1461023f576100fe565b8063248a9ca31461010357806325307065146101405780632f2ff15d1461015c578063318c136e14610185575b600080fd5b34801561010f57600080fd5b5061012a6004803603810190610125919061101e565b610430565b604051610137919061105a565b60405180910390f35b61015a60048036038101906101559190611171565b61044f565b005b34801561016857600080fd5b50610183600480360381019061017e9190611240565b6104e9565b005b34801561019157600080fd5b5061019a61055c565b6040516101a7919061128f565b60405180910390f35b3480156101bc57600080fd5b506101d760048036038101906101d29190611240565b610580565b005b3480156101e557600080fd5b5061020060048036038101906101fb9190611240565b610603565b60405161020d91906112c3565b60405180910390f35b34801561022257600080fd5b5061023d600480360381019061023891906112de565b610653565b005b34801561024b57600080fd5b5061026660048036038101906102619190611337565b6106ec565b005b34801561027457600080fd5b5061027d6107be565b60405161028a919061128f565b60405180910390f35b34801561029f57600080fd5b506102ba60048036038101906102b59190611364565b6107e2565b6040516102c7919061128f565b60405180910390f35b3480156102dc57600080fd5b506102f760048036038101906102f29190611450565b610813565b005b34801561030557600080fd5b50610320600480360381019061031b9190611240565b610a38565b60405161032d91906114ec565b60405180910390f35b34801561034257600080fd5b5061034b610a69565b604051610358919061105a565b60405180910390f35b34801561036d57600080fd5b50610376610a70565b60405161038391906112c3565b60405180910390f35b34801561039857600080fd5b506103b360048036038101906103ae919061101e565b610a76565b6040516103c091906112c3565b60405180910390f35b3480156103d557600080fd5b506103f060048036038101906103eb9190611240565b610a9c565b005b3480156103fe57600080fd5b5061041960048036038101906104149190611171565b610b0f565b604051610427929190611507565b60405180910390f35b6000806000838152602001908152602001600020600201549050919050565b610457610b58565b600154341461049b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104929061158d565b60405180910390fd5b7fbd231b7fa4103e15e7a238c72f07e8aff310701af121895aa6c793b80245e4338888888860015460006040516104d7969594939291906115bc565b60405180910390a15050505050505050565b61050f6000808481526020019081526020016000206002015461050a610c3e565b610a38565b61054e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105459061168f565b60405180910390fd5b6105588282610c46565b5050565b7f000000000000000000000000000000000000000000000000000000000000000081565b610588610c3e565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146105f5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105ec90611721565b60405180910390fd5b6105ff8282610cd9565b5050565b600080600084815260200190815260200160002060000160000160010160008373ffffffffffffffffffffffffffffffffffffffff1660001b815260200190815260200160002054905092915050565b600061065d610c3e565b90508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156106ce576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106c59061178d565b60405180910390fd5b6106db6000801b836104e9565b6106e86000801b82610580565b5050565b6106f96000801b33610a38565b610738576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161072f906117f9565b60405180910390fd5b80600154141561077d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161077490611865565b60405180910390fd5b806001819055507f6bbc57480a46553fa4d156ce702beef5f3ad66303b0ed1a5d4cb44966c6584c3816040516107b391906112c3565b60405180910390a150565b7f000000000000000000000000000000000000000000000000000000000000000081565b600061080b82600080868152602001908152602001600020600001610d6c90919063ffffffff16565b905092915050565b6108206000801b33610a38565b61085f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610856906117f9565b60405180910390fd5b8181905084849050146108a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161089e906118d1565b60405180910390fd5b60005b84849050811015610a315760008585838181106108ca576108c96118f1565b5b90506020020160208101906108df919061195e565b73ffffffffffffffffffffffffffffffffffffffff16848484818110610908576109076118f1565b5b9050602002013560405161091b906119bc565b60006040518083038185875af1925050503d8060008114610958576040519150601f19603f3d011682016040523d82523d6000602084013e61095d565b606091505b50509050806109a1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161099890611a1d565b60405180910390fd5b7faaa40a232aaf133fdd28f3485f6fdd163514cfadbffa981f3610f42398efe34b60008787858181106109d7576109d66118f1565b5b90506020020160208101906109ec919061195e565b8686868181106109ff576109fe6118f1565b5b90506020020135604051610a1593929190611a9c565b60405180910390a1508080610a2990611b02565b9150506108aa565b5050505050565b6000610a6182600080868152602001908152602001600020600001610d8690919063ffffffff16565b905092915050565b6000801b81565b60015481565b6000610a95600080848152602001908152602001600020600001610db6565b9050919050565b610ac260008084815260200190815260200160002060020154610abd610c3e565b610a38565b610b01576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610af890611bbd565b60405180910390fd5b610b0b8282610cd9565b5050565b6000806001546000915091509850989650505050505050565b6000610b50836000018373ffffffffffffffffffffffffffffffffffffffff1660001b610dcb565b905092915050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161480610bfd57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b610c3c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c3390611c4f565b60405180910390fd5b565b600033905090565b610c6d81600080858152602001908152602001600020600001610b2890919063ffffffff16565b15610cd557610c7a610c3e565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b610d0081600080858152602001908152602001600020600001610e3b90919063ffffffff16565b15610d6857610d0d610c3e565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45b5050565b6000610d7b8360000183610e6b565b60001c905092915050565b6000610dae836000018373ffffffffffffffffffffffffffffffffffffffff1660001b610e96565b905092915050565b6000610dc482600001610eb9565b9050919050565b6000610dd78383610e96565b610e30578260000182908060018154018082558091505060019003906000526020600020016000909190919091505582600001805490508360010160008481526020019081526020016000208190555060019050610e35565b600090505b92915050565b6000610e63836000018373ffffffffffffffffffffffffffffffffffffffff1660001b610eca565b905092915050565b6000826000018281548110610e8357610e826118f1565b5b9060005260206000200154905092915050565b600080836001016000848152602001908152602001600020541415905092915050565b600081600001805490509050919050565b60008083600101600084815260200190815260200160002054905060008114610fd2576000600182610efc9190611c6f565b9050600060018660000180549050610f149190611c6f565b9050818114610f83576000866000018281548110610f3557610f346118f1565b5b9060005260206000200154905080876000018481548110610f5957610f586118f1565b5b90600052602060002001819055508387600101600083815260200190815260200160002081905550505b85600001805480610f9757610f96611ca3565b5b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610fd8565b60009150505b92915050565b600080fd5b600080fd5b6000819050919050565b610ffb81610fe8565b811461100657600080fd5b50565b60008135905061101881610ff2565b92915050565b60006020828403121561103457611033610fde565b5b600061104284828501611009565b91505092915050565b61105481610fe8565b82525050565b600060208201905061106f600083018461104b565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006110a082611075565b9050919050565b6110b081611095565b81146110bb57600080fd5b50565b6000813590506110cd816110a7565b92915050565b600060ff82169050919050565b6110e9816110d3565b81146110f457600080fd5b50565b600081359050611106816110e0565b92915050565b600080fd5b600080fd5b600080fd5b60008083601f8401126111315761113061110c565b5b8235905067ffffffffffffffff81111561114e5761114d611111565b5b60208301915083600182028301111561116a57611169611116565b5b9250929050565b60008060008060008060008060c0898b03121561119157611190610fde565b5b600061119f8b828c016110be565b98505060206111b08b828c016110f7565b97505060406111c18b828c016110f7565b96505060606111d28b828c01611009565b955050608089013567ffffffffffffffff8111156111f3576111f2610fe3565b5b6111ff8b828c0161111b565b945094505060a089013567ffffffffffffffff81111561122257611221610fe3565b5b61122e8b828c0161111b565b92509250509295985092959890939650565b6000806040838503121561125757611256610fde565b5b600061126585828601611009565b9250506020611276858286016110be565b9150509250929050565b61128981611095565b82525050565b60006020820190506112a46000830184611280565b92915050565b6000819050919050565b6112bd816112aa565b82525050565b60006020820190506112d860008301846112b4565b92915050565b6000602082840312156112f4576112f3610fde565b5b6000611302848285016110be565b91505092915050565b611314816112aa565b811461131f57600080fd5b50565b6000813590506113318161130b565b92915050565b60006020828403121561134d5761134c610fde565b5b600061135b84828501611322565b91505092915050565b6000806040838503121561137b5761137a610fde565b5b600061138985828601611009565b925050602061139a85828601611322565b9150509250929050565b60008083601f8401126113ba576113b961110c565b5b8235905067ffffffffffffffff8111156113d7576113d6611111565b5b6020830191508360208202830111156113f3576113f2611116565b5b9250929050565b60008083601f8401126114105761140f61110c565b5b8235905067ffffffffffffffff81111561142d5761142c611111565b5b60208301915083602082028301111561144957611448611116565b5b9250929050565b6000806000806040858703121561146a57611469610fde565b5b600085013567ffffffffffffffff81111561148857611487610fe3565b5b611494878288016113a4565b9450945050602085013567ffffffffffffffff8111156114b7576114b6610fe3565b5b6114c3878288016113fa565b925092505092959194509250565b60008115159050919050565b6114e6816114d1565b82525050565b600060208201905061150160008301846114dd565b92915050565b600060408201905061151c60008301856112b4565b6115296020830184611280565b9392505050565b600082825260208201905092915050565b7f496e636f72726563742066656520737570706c69656400000000000000000000600082015250565b6000611577601683611530565b915061158282611541565b602082019050919050565b600060208201905081810360008301526115a68161156a565b9050919050565b6115b6816110d3565b82525050565b600060c0820190506115d16000830189611280565b6115de60208301886115ad565b6115eb60408301876115ad565b6115f8606083018661104b565b61160560808301856112b4565b61161260a0830184611280565b979650505050505050565b7f416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e60008201527f2061646d696e20746f206772616e740000000000000000000000000000000000602082015250565b6000611679602f83611530565b91506116848261161d565b604082019050919050565b600060208201905081810360008301526116a88161166c565b9050919050565b7f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560008201527f20726f6c657320666f722073656c660000000000000000000000000000000000602082015250565b600061170b602f83611530565b9150611716826116af565b604082019050919050565b6000602082019050818103600083015261173a816116fe565b9050919050565b7f43616e6e6f742072656e6f756e6365206f6e6573656c66000000000000000000600082015250565b6000611777601783611530565b915061178282611741565b602082019050919050565b600060208201905081810360008301526117a68161176a565b9050919050565b7f73656e64657220646f65736e277420686176652061646d696e20726f6c650000600082015250565b60006117e3601e83611530565b91506117ee826117ad565b602082019050919050565b60006020820190508181036000830152611812816117d6565b9050919050565b7f43757272656e742066656520697320657175616c20746f206e65772066656500600082015250565b600061184f601f83611530565b915061185a82611819565b602082019050919050565b6000602082019050818103600083015261187e81611842565b9050919050565b7f61646472735b5d2c20616d6f756e74735b5d3a2064696666206c656e67746800600082015250565b60006118bb601f83611530565b91506118c682611885565b602082019050919050565b600060208201905081810360008301526118ea816118ae565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600061192b82611075565b9050919050565b61193b81611920565b811461194657600080fd5b50565b60008135905061195881611932565b92915050565b60006020828403121561197457611973610fde565b5b600061198284828501611949565b91505092915050565b600081905092915050565b50565b60006119a660008361198b565b91506119b182611996565b600082019050919050565b60006119c782611999565b9150819050919050565b7f466565206574686572207472616e73666572206661696c656400000000000000600082015250565b6000611a07601983611530565b9150611a12826119d1565b602082019050919050565b60006020820190508181036000830152611a36816119fa565b9050919050565b6000819050919050565b6000611a62611a5d611a5884611075565b611a3d565b611075565b9050919050565b6000611a7482611a47565b9050919050565b6000611a8682611a69565b9050919050565b611a9681611a7b565b82525050565b6000606082019050611ab16000830186611280565b611abe6020830185611a8d565b611acb60408301846112b4565b949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611b0d826112aa565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611b4057611b3f611ad3565b5b600182019050919050565b7f416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e60008201527f2061646d696e20746f207265766f6b6500000000000000000000000000000000602082015250565b6000611ba7603083611530565b9150611bb282611b4b565b604082019050919050565b60006020820190508181036000830152611bd681611b9a565b9050919050565b7f73656e646572206d75737420626520627269646765206f722066656520726f7560008201527f74657220636f6e74726163740000000000000000000000000000000000000000602082015250565b6000611c39602c83611530565b9150611c4482611bdd565b604082019050919050565b60006020820190508181036000830152611c6881611c2c565b9050919050565b6000611c7a826112aa565b9150611c85836112aa565b925082821015611c9857611c97611ad3565b5b828203905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea2646970667358221220df12b34d4aaa792caad3d5838f24f5d030918bc88101414baae2311897d7c12f64736f6c634300080b0033"

// DeployBasicFeeHandler deploys a new Ethereum contract, binding an instance of BasicFeeHandler to it.
func DeployBasicFeeHandler(auth *bind.TransactOpts, backend bind.ContractBackend, bridgeAddress common.Address, feeHandlerRouterAddress common.Address) (common.Address, *types.Transaction, *BasicFeeHandler, error) {
	parsed, err := abi.JSON(strings.NewReader(BasicFeeHandlerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BasicFeeHandlerBin), backend, bridgeAddress, feeHandlerRouterAddress)
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

// FeeHandlerRouterAddress is a free data retrieval call binding the contract method 0x745e6b61.
//
// Solidity: function _feeHandlerRouterAddress() view returns(address)
func (_BasicFeeHandler *BasicFeeHandlerCaller) FeeHandlerRouterAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasicFeeHandler.contract.Call(opts, &out, "_feeHandlerRouterAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeHandlerRouterAddress is a free data retrieval call binding the contract method 0x745e6b61.
//
// Solidity: function _feeHandlerRouterAddress() view returns(address)
func (_BasicFeeHandler *BasicFeeHandlerSession) FeeHandlerRouterAddress() (common.Address, error) {
	return _BasicFeeHandler.Contract.FeeHandlerRouterAddress(&_BasicFeeHandler.CallOpts)
}

// FeeHandlerRouterAddress is a free data retrieval call binding the contract method 0x745e6b61.
//
// Solidity: function _feeHandlerRouterAddress() view returns(address)
func (_BasicFeeHandler *BasicFeeHandlerCallerSession) FeeHandlerRouterAddress() (common.Address, error) {
	return _BasicFeeHandler.Contract.FeeHandlerRouterAddress(&_BasicFeeHandler.CallOpts)
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

// RenounceAdmin is a paid mutator transaction binding the contract method 0x5e1fab0f.
//
// Solidity: function renounceAdmin(address newAdmin) returns()
func (_BasicFeeHandler *BasicFeeHandlerTransactor) RenounceAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _BasicFeeHandler.contract.Transact(opts, "renounceAdmin", newAdmin)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x5e1fab0f.
//
// Solidity: function renounceAdmin(address newAdmin) returns()
func (_BasicFeeHandler *BasicFeeHandlerSession) RenounceAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _BasicFeeHandler.Contract.RenounceAdmin(&_BasicFeeHandler.TransactOpts, newAdmin)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x5e1fab0f.
//
// Solidity: function renounceAdmin(address newAdmin) returns()
func (_BasicFeeHandler *BasicFeeHandlerTransactorSession) RenounceAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _BasicFeeHandler.Contract.RenounceAdmin(&_BasicFeeHandler.TransactOpts, newAdmin)
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
