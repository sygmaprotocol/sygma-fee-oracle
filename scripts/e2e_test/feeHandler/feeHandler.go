// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package feeHandler

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

// FeeHandlerMetaData contains all meta data concerning the FeeHandler contract.
var FeeHandlerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeHandlerRouterAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"fromDomainID\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"destinationDomainID\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"FeeCollected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeeDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_feeHandlerRouterAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_feePercent\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_gasUsed\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_oracleAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"fromDomainID\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"destinationDomainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"depositData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"name\":\"calculateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"fromDomainID\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"destinationDomainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"depositData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"feeData\",\"type\":\"bytes\"}],\"name\":\"collectFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"renounceAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracleAddress\",\"type\":\"address\"}],\"name\":\"setFeeOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"gasUsed\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"feePercent\",\"type\":\"uint16\"}],\"name\":\"setFeeProperties\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"transferFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "60c06040523480156200001157600080fd5b5060405162001f2738038062001f2783398101604081905262000034916200012d565b6001600160a01b03808316608052811660a05281816200005660003362000060565b5050505062000165565b6200006c828262000070565b5050565b6000828152602081815260408083206001600160a01b038516845290915290205460ff166200006c576000828152602081815260408083206001600160a01b03851684529091529020805460ff19166001179055620000cc3390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b80516001600160a01b03811681146200012857600080fd5b919050565b600080604083850312156200014157600080fd5b6200014c8362000110565b91506200015c6020840162000110565b90509250929050565b60805160a051611d80620001a7600039600081816102a90152610a6c0152600081816101c80152818161078f01528181610a3a0152610d1b0152611d806000f3fe6080604052600436106101095760003560e01c8063745e6b6111610095578063bff4275511610064578063bff4275514610320578063c297983f14610340578063d547741f14610360578063ef4f081f14610380578063fc818cfb146103bd57600080fd5b8063745e6b611461029757806391d14854146102cb578063a217fddf146102eb578063a8a989621461030057600080fd5b8063318c136e116100dc578063318c136e146101b657806336568abe146102025780635e1fab0f1461022257806369222948146102425780636fb7cb571461027757600080fd5b806301ffc9a71461010e578063248a9ca31461014357806325307065146101815780632f2ff15d14610196575b600080fd5b34801561011a57600080fd5b5061012e610129366004611777565b6103f6565b60405190151581526020015b60405180910390f35b34801561014f57600080fd5b5061017361015e3660046117a1565b60009081526020819052604090206001015490565b60405190815260200161013a565b61019461018f366004611820565b61042d565b005b3480156101a257600080fd5b506101946101b13660046118d1565b610575565b3480156101c257600080fd5b506101ea7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b03909116815260200161013a565b34801561020e57600080fd5b5061019461021d3660046118d1565b6105a0565b34801561022e57600080fd5b5061019461023d366004611901565b61061e565b34801561024e57600080fd5b5060015461026490600160c01b900461ffff1681565b60405161ffff909116815260200161013a565b34801561028357600080fd5b506001546101ea906001600160a01b031681565b3480156102a357600080fd5b506101ea7f000000000000000000000000000000000000000000000000000000000000000081565b3480156102d757600080fd5b5061012e6102e63660046118d1565b61068e565b3480156102f757600080fd5b50610173600081565b34801561030c57600080fd5b5061019461031b366004611901565b6106b7565b34801561032c57600080fd5b5061019461033b366004611963565b610700565b34801561034c57600080fd5b5061019461035b3660046119dd565b61097b565b34801561036c57600080fd5b5061019461037b3660046118d1565b6109e5565b34801561038c57600080fd5b506103a061039b366004611820565b610a0b565b604080519283526001600160a01b0390911660208301520161013a565b3480156103c957600080fd5b506001546103e190600160a01b900463ffffffff1681565b60405163ffffffff909116815260200161013a565b60006001600160e01b03198216637965db0b60e01b148061042757506301ffc9a760e01b6001600160e01b03198316145b92915050565b610435610a2f565b6000806104488a8a8a8a8a8a8a8a610af1565b90925090506001600160a01b0381166104ad578134146104a85760405162461bcd60e51b8152602060048201526016602482015275125b98dbdc9c9958dd08199959481cdd5c1c1b1a595960521b60448201526064015b60405180910390fd5b610507565b34156104fb5760405162461bcd60e51b815260206004820152601a60248201527f636f6c6c6563744665653a206d73672e76616c756520213d2030000000000000604482015260640161049f565b610507818b3085610e99565b604080516001600160a01b038c8116825260ff8c811660208401528b1682840152606082018a905260808201859052831660a082015290517fbd231b7fa4103e15e7a238c72f07e8aff310701af121895aa6c793b80245e4339181900360c00190a150505050505050505050565b6000828152602081905260409020600101546105918133610ead565b61059b8383610f11565b505050565b6001600160a01b03811633146106105760405162461bcd60e51b815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201526e103937b632b9903337b91039b2b63360891b606482015260840161049f565b61061a8282610f95565b5050565b336001600160a01b0382168114156106785760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f742072656e6f756e6365206f6e6573656c66000000000000000000604482015260640161049f565b610683600083610575565b61061a6000826105a0565b6000918252602082815260408084206001600160a01b0393909316845291905290205460ff1690565b6106c260003361068e565b6106de5760405162461bcd60e51b815260040161049f90611a1b565b600180546001600160a01b0319166001600160a01b0392909216919091179055565b61070b60003361068e565b6107275760405162461bcd60e51b815260040161049f90611a1b565b8281146107765760405162461bcd60e51b815260206004820152601f60248201527f61646472735b5d2c20616d6f756e74735b5d3a2064696666206c656e67746800604482015260640161049f565b6040516384db809f60e01b8152600481018690526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906384db809f90602401602060405180830381865afa1580156107de573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108029190611a52565b60405163014daabb60e31b8152600481018890529091506000906001600160a01b03831690630a6d55d890602401602060405180830381865afa15801561084d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108719190611a52565b905060005b85811015610971576108c78288888481811061089457610894611a6f565b90506020020160208101906108a99190611901565b8787858181106108bb576108bb611a6f565b90506020020135610ffa565b7faaa40a232aaf133fdd28f3485f6fdd163514cfadbffa981f3610f42398efe34b828888848181106108fb576108fb611a6f565b90506020020160208101906109109190611901565b87878581811061092257610922611a6f565b90506020020135604051610957939291906001600160a01b039384168152919092166020820152604081019190915260600190565b60405180910390a18061096981611a9b565b915050610876565b5050505050505050565b61098660003361068e565b6109a25760405162461bcd60e51b815260040161049f90611a1b565b6001805461ffff909216600160c01b0261ffff60c01b1963ffffffff909416600160a01b029390931665ffffffffffff60a01b1990921691909117919091179055565b600082815260208190526040902060010154610a018133610ead565b61059b8383610f95565b600080610a1e8a8a8a8a8a8a8a8a610af1565b915091509850989650505050505050565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480610a8e5750336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016145b610aef5760405162461bcd60e51b815260206004820152602c60248201527f73656e646572206d75737420626520627269646765206f722066656520726f7560448201526b1d195c8818dbdb9d1c9858dd60a21b606482015260840161049f565b565b6000806101618314610b455760405162461bcd60e51b815260206004820152601860248201527f496e636f72726563742066656544617461206c656e6774680000000000000000604482015260640161049f565b610b6960405180606001604052806060815260200160608152602001600081525090565b6000610b79610100828789611ab6565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505090835250610bc06101416101008789611ab6565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920182905250602080880195909552865180519195610c13955091810182019350019050611b06565b90508060600151421115610c605760405162461bcd60e51b81526020600482015260146024820152734f62736f6c657465206f7261636c65206461746160601b604482015260640161049f565b8b60ff16816080015160ff16148015610c8257508a60ff168160a0015160ff16145b8015610c915750898160c00151145b610cdd5760405162461bcd60e51b815260206004820152601860248201527f496e636f7272656374206465706f73697420706172616d730000000000000000604482015260640161049f565b8251805160209182012090840151600154610d029183916001600160a01b031661100c565b6040516384db809f60e01b8152600481018c90526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906384db809f90602401602060405180830381865afa158015610d6a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d8e9190611a52565b60405163014daabb60e31b8152600481018e90529091506001600160a01b03821690630a6d55d890602401602060405180830381865afa158015610dd6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dfa9190611a52565b60208401516001546040860151929850670de0b6b3a764000092610e2b91600160a01b900463ffffffff1690611bb2565b610e359190611bb2565b610e3f9190611bd1565b93506000610e4f8b8d018d6117a1565b60015490915061271090610e6e90600160c01b900461ffff1683611bb2565b610e789190611bd1565b975084881015610e86578497505b5050505050509850989650505050505050565b83610ea68185858561106f565b5050505050565b610eb7828261068e565b61061a57610ecf816001600160a01b031660146110da565b610eda8360206110da565b604051602001610eeb929190611c1f565b60408051601f198184030181529082905262461bcd60e51b825261049f91600401611c94565b610f1b828261068e565b61061a576000828152602081815260408083206001600160a01b03851684529091529020805460ff19166001179055610f513390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b610f9f828261068e565b1561061a576000828152602081815260408083206001600160a01b0385168085529252808320805460ff1916905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b8261100681848461127d565b50505050565b600061101884846112ad565b9050816001600160a01b0316816001600160a01b0316146110065760405162461bcd60e51b8152602060048201526011602482015270496e76616c6964207369676e617475726560781b604482015260640161049f565b6040516001600160a01b03808516602483015283166044820152606481018290526110069085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b0319909316929092179091526112d1565b606060006110e9836002611bb2565b6110f4906002611cc7565b67ffffffffffffffff81111561110c5761110c611ae0565b6040519080825280601f01601f191660200182016040528015611136576020820181803683370190505b509050600360fc1b8160008151811061115157611151611a6f565b60200101906001600160f81b031916908160001a905350600f60fb1b8160018151811061118057611180611a6f565b60200101906001600160f81b031916908160001a90535060006111a4846002611bb2565b6111af906001611cc7565b90505b6001811115611227576f181899199a1a9b1b9c1cb0b131b232b360811b85600f16601081106111e3576111e3611a6f565b1a60f81b8282815181106111f9576111f9611a6f565b60200101906001600160f81b031916908160001a90535060049490941c9361122081611cdf565b90506111b2565b5083156112765760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e74604482015260640161049f565b9392505050565b6040516001600160a01b03831660248201526044810182905261059b90849063a9059cbb60e01b906064016110a3565b60008060006112bc8585611423565b915091506112c981611493565b509392505050565b813b806113185760405162461bcd60e51b8152602060048201526015602482015274115490cc8c0e881b9bdd08184818dbdb9d1c9858dd605a1b604482015260640161049f565b600080846001600160a01b0316846040516113339190611cf6565b6000604051808303816000865af19150503d8060008114611370576040519150601f19603f3d011682016040523d82523d6000602084013e611375565b606091505b5091509150816113bc5760405162461bcd60e51b8152602060048201526012602482015271115490cc8c0e8818d85b1b0819985a5b195960721b604482015260640161049f565b805115610ea657808060200190518101906113d79190611d12565b610ea65760405162461bcd60e51b815260206004820181905260248201527f45524332303a206f7065726174696f6e20646964206e6f742073756363656564604482015260640161049f565b60008082516041141561145a5760208301516040840151606085015160001a61144e87828585611651565b9450945050505061148c565b825160401415611484576020830151604084015161147986838361173e565b93509350505061148c565b506000905060025b9250929050565b60008160048111156114a7576114a7611d34565b14156114b05750565b60018160048111156114c4576114c4611d34565b14156115125760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015260640161049f565b600281600481111561152657611526611d34565b14156115745760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015260640161049f565b600381600481111561158857611588611d34565b14156115e15760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b606482015260840161049f565b60048160048111156115f5576115f5611d34565b141561164e5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b606482015260840161049f565b50565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156116885750600090506003611735565b8460ff16601b141580156116a057508460ff16601c14155b156116b15750600090506004611735565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015611705573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b03811661172e57600060019250925050611735565b9150600090505b94509492505050565b6000806001600160ff1b0383168161175b60ff86901c601b611cc7565b905061176987828885611651565b935093505050935093915050565b60006020828403121561178957600080fd5b81356001600160e01b03198116811461127657600080fd5b6000602082840312156117b357600080fd5b5035919050565b6001600160a01b038116811461164e57600080fd5b60ff8116811461164e57600080fd5b60008083601f8401126117f057600080fd5b50813567ffffffffffffffff81111561180857600080fd5b60208301915083602082850101111561148c57600080fd5b60008060008060008060008060c0898b03121561183c57600080fd5b8835611847816117ba565b97506020890135611857816117cf565b96506040890135611867816117cf565b955060608901359450608089013567ffffffffffffffff8082111561188b57600080fd5b6118978c838d016117de565b909650945060a08b01359150808211156118b057600080fd5b506118bd8b828c016117de565b999c989b5096995094979396929594505050565b600080604083850312156118e457600080fd5b8235915060208301356118f6816117ba565b809150509250929050565b60006020828403121561191357600080fd5b8135611276816117ba565b60008083601f84011261193057600080fd5b50813567ffffffffffffffff81111561194857600080fd5b6020830191508360208260051b850101111561148c57600080fd5b60008060008060006060868803121561197b57600080fd5b85359450602086013567ffffffffffffffff8082111561199a57600080fd5b6119a689838a0161191e565b909650945060408801359150808211156119bf57600080fd5b506119cc8882890161191e565b969995985093965092949392505050565b600080604083850312156119f057600080fd5b823563ffffffff81168114611a0457600080fd5b9150602083013561ffff811681146118f657600080fd5b6020808252601e908201527f73656e64657220646f65736e277420686176652061646d696e20726f6c650000604082015260600190565b600060208284031215611a6457600080fd5b8151611276816117ba565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000600019821415611aaf57611aaf611a85565b5060010190565b60008085851115611ac657600080fd5b83861115611ad357600080fd5b5050820193919092039150565b634e487b7160e01b600052604160045260246000fd5b8051611b01816117cf565b919050565b6000610100808385031215611b1a57600080fd5b6040519081019067ffffffffffffffff82118183101715611b4b57634e487b7160e01b600052604160045260246000fd5b8160405283518152602084015160208201526040840151604082015260608401516060820152611b7d60808501611af6565b6080820152611b8e60a08501611af6565b60a082015260c084015160c082015260e084015160e0820152809250505092915050565b6000816000190483118215151615611bcc57611bcc611a85565b500290565b600082611bee57634e487b7160e01b600052601260045260246000fd5b500490565b60005b83811015611c0e578181015183820152602001611bf6565b838111156110065750506000910152565b7f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000815260008351611c57816017850160208801611bf3565b7001034b99036b4b9b9b4b733903937b6329607d1b6017918401918201528351611c88816028840160208801611bf3565b01602801949350505050565b6020815260008251806020840152611cb3816040850160208701611bf3565b601f01601f19169190910160400192915050565b60008219821115611cda57611cda611a85565b500190565b600081611cee57611cee611a85565b506000190190565b60008251611d08818460208701611bf3565b9190910192915050565b600060208284031215611d2457600080fd5b8151801515811461127657600080fd5b634e487b7160e01b600052602160045260246000fdfea26469706673582212204e86d92d57bf08ba21eb6ab0a75e4ca7b38bb09b447e19e5d3e535c235af336f64736f6c634300080b0033",
}

// FeeHandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use FeeHandlerMetaData.ABI instead.
var FeeHandlerABI = FeeHandlerMetaData.ABI

// FeeHandlerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FeeHandlerMetaData.Bin instead.
var FeeHandlerBin = FeeHandlerMetaData.Bin

// DeployFeeHandler deploys a new Ethereum contract, binding an instance of FeeHandler to it.
func DeployFeeHandler(auth *bind.TransactOpts, backend bind.ContractBackend, bridgeAddress common.Address, feeHandlerRouterAddress common.Address) (common.Address, *types.Transaction, *FeeHandler, error) {
	parsed, err := FeeHandlerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FeeHandlerBin), backend, bridgeAddress, feeHandlerRouterAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FeeHandler{FeeHandlerCaller: FeeHandlerCaller{contract: contract}, FeeHandlerTransactor: FeeHandlerTransactor{contract: contract}, FeeHandlerFilterer: FeeHandlerFilterer{contract: contract}}, nil
}

// FeeHandler is an auto generated Go binding around an Ethereum contract.
type FeeHandler struct {
	FeeHandlerCaller     // Read-only binding to the contract
	FeeHandlerTransactor // Write-only binding to the contract
	FeeHandlerFilterer   // Log filterer for contract events
}

// FeeHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type FeeHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FeeHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FeeHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FeeHandlerSession struct {
	Contract     *FeeHandler       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FeeHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FeeHandlerCallerSession struct {
	Contract *FeeHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// FeeHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FeeHandlerTransactorSession struct {
	Contract     *FeeHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// FeeHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type FeeHandlerRaw struct {
	Contract *FeeHandler // Generic contract binding to access the raw methods on
}

// FeeHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FeeHandlerCallerRaw struct {
	Contract *FeeHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// FeeHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FeeHandlerTransactorRaw struct {
	Contract *FeeHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFeeHandler creates a new instance of FeeHandler, bound to a specific deployed contract.
func NewFeeHandler(address common.Address, backend bind.ContractBackend) (*FeeHandler, error) {
	contract, err := bindFeeHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeeHandler{FeeHandlerCaller: FeeHandlerCaller{contract: contract}, FeeHandlerTransactor: FeeHandlerTransactor{contract: contract}, FeeHandlerFilterer: FeeHandlerFilterer{contract: contract}}, nil
}

// NewFeeHandlerCaller creates a new read-only instance of FeeHandler, bound to a specific deployed contract.
func NewFeeHandlerCaller(address common.Address, caller bind.ContractCaller) (*FeeHandlerCaller, error) {
	contract, err := bindFeeHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeeHandlerCaller{contract: contract}, nil
}

// NewFeeHandlerTransactor creates a new write-only instance of FeeHandler, bound to a specific deployed contract.
func NewFeeHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*FeeHandlerTransactor, error) {
	contract, err := bindFeeHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeeHandlerTransactor{contract: contract}, nil
}

// NewFeeHandlerFilterer creates a new log filterer instance of FeeHandler, bound to a specific deployed contract.
func NewFeeHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*FeeHandlerFilterer, error) {
	contract, err := bindFeeHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeeHandlerFilterer{contract: contract}, nil
}

// bindFeeHandler binds a generic wrapper to an already deployed contract.
func bindFeeHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FeeHandlerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeHandler *FeeHandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeHandler.Contract.FeeHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeHandler *FeeHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeHandler.Contract.FeeHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeHandler *FeeHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeHandler.Contract.FeeHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeHandler *FeeHandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeHandler *FeeHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeHandler *FeeHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeHandler.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_FeeHandler *FeeHandlerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FeeHandler.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_FeeHandler *FeeHandlerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _FeeHandler.Contract.DEFAULTADMINROLE(&_FeeHandler.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_FeeHandler *FeeHandlerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _FeeHandler.Contract.DEFAULTADMINROLE(&_FeeHandler.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_FeeHandler *FeeHandlerCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeHandler.contract.Call(opts, &out, "_bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_FeeHandler *FeeHandlerSession) BridgeAddress() (common.Address, error) {
	return _FeeHandler.Contract.BridgeAddress(&_FeeHandler.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_FeeHandler *FeeHandlerCallerSession) BridgeAddress() (common.Address, error) {
	return _FeeHandler.Contract.BridgeAddress(&_FeeHandler.CallOpts)
}

// FeeHandlerRouterAddress is a free data retrieval call binding the contract method 0x745e6b61.
//
// Solidity: function _feeHandlerRouterAddress() view returns(address)
func (_FeeHandler *FeeHandlerCaller) FeeHandlerRouterAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeHandler.contract.Call(opts, &out, "_feeHandlerRouterAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeHandlerRouterAddress is a free data retrieval call binding the contract method 0x745e6b61.
//
// Solidity: function _feeHandlerRouterAddress() view returns(address)
func (_FeeHandler *FeeHandlerSession) FeeHandlerRouterAddress() (common.Address, error) {
	return _FeeHandler.Contract.FeeHandlerRouterAddress(&_FeeHandler.CallOpts)
}

// FeeHandlerRouterAddress is a free data retrieval call binding the contract method 0x745e6b61.
//
// Solidity: function _feeHandlerRouterAddress() view returns(address)
func (_FeeHandler *FeeHandlerCallerSession) FeeHandlerRouterAddress() (common.Address, error) {
	return _FeeHandler.Contract.FeeHandlerRouterAddress(&_FeeHandler.CallOpts)
}

// FeePercent is a free data retrieval call binding the contract method 0x69222948.
//
// Solidity: function _feePercent() view returns(uint16)
func (_FeeHandler *FeeHandlerCaller) FeePercent(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _FeeHandler.contract.Call(opts, &out, "_feePercent")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// FeePercent is a free data retrieval call binding the contract method 0x69222948.
//
// Solidity: function _feePercent() view returns(uint16)
func (_FeeHandler *FeeHandlerSession) FeePercent() (uint16, error) {
	return _FeeHandler.Contract.FeePercent(&_FeeHandler.CallOpts)
}

// FeePercent is a free data retrieval call binding the contract method 0x69222948.
//
// Solidity: function _feePercent() view returns(uint16)
func (_FeeHandler *FeeHandlerCallerSession) FeePercent() (uint16, error) {
	return _FeeHandler.Contract.FeePercent(&_FeeHandler.CallOpts)
}

// GasUsed is a free data retrieval call binding the contract method 0xfc818cfb.
//
// Solidity: function _gasUsed() view returns(uint32)
func (_FeeHandler *FeeHandlerCaller) GasUsed(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _FeeHandler.contract.Call(opts, &out, "_gasUsed")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GasUsed is a free data retrieval call binding the contract method 0xfc818cfb.
//
// Solidity: function _gasUsed() view returns(uint32)
func (_FeeHandler *FeeHandlerSession) GasUsed() (uint32, error) {
	return _FeeHandler.Contract.GasUsed(&_FeeHandler.CallOpts)
}

// GasUsed is a free data retrieval call binding the contract method 0xfc818cfb.
//
// Solidity: function _gasUsed() view returns(uint32)
func (_FeeHandler *FeeHandlerCallerSession) GasUsed() (uint32, error) {
	return _FeeHandler.Contract.GasUsed(&_FeeHandler.CallOpts)
}

// OracleAddress is a free data retrieval call binding the contract method 0x6fb7cb57.
//
// Solidity: function _oracleAddress() view returns(address)
func (_FeeHandler *FeeHandlerCaller) OracleAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeHandler.contract.Call(opts, &out, "_oracleAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OracleAddress is a free data retrieval call binding the contract method 0x6fb7cb57.
//
// Solidity: function _oracleAddress() view returns(address)
func (_FeeHandler *FeeHandlerSession) OracleAddress() (common.Address, error) {
	return _FeeHandler.Contract.OracleAddress(&_FeeHandler.CallOpts)
}

// OracleAddress is a free data retrieval call binding the contract method 0x6fb7cb57.
//
// Solidity: function _oracleAddress() view returns(address)
func (_FeeHandler *FeeHandlerCallerSession) OracleAddress() (common.Address, error) {
	return _FeeHandler.Contract.OracleAddress(&_FeeHandler.CallOpts)
}

// CalculateFee is a free data retrieval call binding the contract method 0xef4f081f.
//
// Solidity: function calculateFee(address sender, uint8 fromDomainID, uint8 destinationDomainID, bytes32 resourceID, bytes depositData, bytes feeData) view returns(uint256 fee, address tokenAddress)
func (_FeeHandler *FeeHandlerCaller) CalculateFee(opts *bind.CallOpts, sender common.Address, fromDomainID uint8, destinationDomainID uint8, resourceID [32]byte, depositData []byte, feeData []byte) (struct {
	Fee          *big.Int
	TokenAddress common.Address
}, error) {
	var out []interface{}
	err := _FeeHandler.contract.Call(opts, &out, "calculateFee", sender, fromDomainID, destinationDomainID, resourceID, depositData, feeData)

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
func (_FeeHandler *FeeHandlerSession) CalculateFee(sender common.Address, fromDomainID uint8, destinationDomainID uint8, resourceID [32]byte, depositData []byte, feeData []byte) (struct {
	Fee          *big.Int
	TokenAddress common.Address
}, error) {
	return _FeeHandler.Contract.CalculateFee(&_FeeHandler.CallOpts, sender, fromDomainID, destinationDomainID, resourceID, depositData, feeData)
}

// CalculateFee is a free data retrieval call binding the contract method 0xef4f081f.
//
// Solidity: function calculateFee(address sender, uint8 fromDomainID, uint8 destinationDomainID, bytes32 resourceID, bytes depositData, bytes feeData) view returns(uint256 fee, address tokenAddress)
func (_FeeHandler *FeeHandlerCallerSession) CalculateFee(sender common.Address, fromDomainID uint8, destinationDomainID uint8, resourceID [32]byte, depositData []byte, feeData []byte) (struct {
	Fee          *big.Int
	TokenAddress common.Address
}, error) {
	return _FeeHandler.Contract.CalculateFee(&_FeeHandler.CallOpts, sender, fromDomainID, destinationDomainID, resourceID, depositData, feeData)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_FeeHandler *FeeHandlerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _FeeHandler.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_FeeHandler *FeeHandlerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _FeeHandler.Contract.GetRoleAdmin(&_FeeHandler.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_FeeHandler *FeeHandlerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _FeeHandler.Contract.GetRoleAdmin(&_FeeHandler.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_FeeHandler *FeeHandlerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _FeeHandler.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_FeeHandler *FeeHandlerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _FeeHandler.Contract.HasRole(&_FeeHandler.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_FeeHandler *FeeHandlerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _FeeHandler.Contract.HasRole(&_FeeHandler.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FeeHandler *FeeHandlerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _FeeHandler.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FeeHandler *FeeHandlerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FeeHandler.Contract.SupportsInterface(&_FeeHandler.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FeeHandler *FeeHandlerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FeeHandler.Contract.SupportsInterface(&_FeeHandler.CallOpts, interfaceId)
}

// CollectFee is a paid mutator transaction binding the contract method 0x25307065.
//
// Solidity: function collectFee(address sender, uint8 fromDomainID, uint8 destinationDomainID, bytes32 resourceID, bytes depositData, bytes feeData) payable returns()
func (_FeeHandler *FeeHandlerTransactor) CollectFee(opts *bind.TransactOpts, sender common.Address, fromDomainID uint8, destinationDomainID uint8, resourceID [32]byte, depositData []byte, feeData []byte) (*types.Transaction, error) {
	return _FeeHandler.contract.Transact(opts, "collectFee", sender, fromDomainID, destinationDomainID, resourceID, depositData, feeData)
}

// CollectFee is a paid mutator transaction binding the contract method 0x25307065.
//
// Solidity: function collectFee(address sender, uint8 fromDomainID, uint8 destinationDomainID, bytes32 resourceID, bytes depositData, bytes feeData) payable returns()
func (_FeeHandler *FeeHandlerSession) CollectFee(sender common.Address, fromDomainID uint8, destinationDomainID uint8, resourceID [32]byte, depositData []byte, feeData []byte) (*types.Transaction, error) {
	return _FeeHandler.Contract.CollectFee(&_FeeHandler.TransactOpts, sender, fromDomainID, destinationDomainID, resourceID, depositData, feeData)
}

// CollectFee is a paid mutator transaction binding the contract method 0x25307065.
//
// Solidity: function collectFee(address sender, uint8 fromDomainID, uint8 destinationDomainID, bytes32 resourceID, bytes depositData, bytes feeData) payable returns()
func (_FeeHandler *FeeHandlerTransactorSession) CollectFee(sender common.Address, fromDomainID uint8, destinationDomainID uint8, resourceID [32]byte, depositData []byte, feeData []byte) (*types.Transaction, error) {
	return _FeeHandler.Contract.CollectFee(&_FeeHandler.TransactOpts, sender, fromDomainID, destinationDomainID, resourceID, depositData, feeData)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_FeeHandler *FeeHandlerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FeeHandler.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_FeeHandler *FeeHandlerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FeeHandler.Contract.GrantRole(&_FeeHandler.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_FeeHandler *FeeHandlerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FeeHandler.Contract.GrantRole(&_FeeHandler.TransactOpts, role, account)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x5e1fab0f.
//
// Solidity: function renounceAdmin(address newAdmin) returns()
func (_FeeHandler *FeeHandlerTransactor) RenounceAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _FeeHandler.contract.Transact(opts, "renounceAdmin", newAdmin)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x5e1fab0f.
//
// Solidity: function renounceAdmin(address newAdmin) returns()
func (_FeeHandler *FeeHandlerSession) RenounceAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _FeeHandler.Contract.RenounceAdmin(&_FeeHandler.TransactOpts, newAdmin)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x5e1fab0f.
//
// Solidity: function renounceAdmin(address newAdmin) returns()
func (_FeeHandler *FeeHandlerTransactorSession) RenounceAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _FeeHandler.Contract.RenounceAdmin(&_FeeHandler.TransactOpts, newAdmin)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_FeeHandler *FeeHandlerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FeeHandler.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_FeeHandler *FeeHandlerSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FeeHandler.Contract.RenounceRole(&_FeeHandler.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_FeeHandler *FeeHandlerTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FeeHandler.Contract.RenounceRole(&_FeeHandler.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_FeeHandler *FeeHandlerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FeeHandler.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_FeeHandler *FeeHandlerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FeeHandler.Contract.RevokeRole(&_FeeHandler.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_FeeHandler *FeeHandlerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FeeHandler.Contract.RevokeRole(&_FeeHandler.TransactOpts, role, account)
}

// SetFeeOracle is a paid mutator transaction binding the contract method 0xa8a98962.
//
// Solidity: function setFeeOracle(address oracleAddress) returns()
func (_FeeHandler *FeeHandlerTransactor) SetFeeOracle(opts *bind.TransactOpts, oracleAddress common.Address) (*types.Transaction, error) {
	return _FeeHandler.contract.Transact(opts, "setFeeOracle", oracleAddress)
}

// SetFeeOracle is a paid mutator transaction binding the contract method 0xa8a98962.
//
// Solidity: function setFeeOracle(address oracleAddress) returns()
func (_FeeHandler *FeeHandlerSession) SetFeeOracle(oracleAddress common.Address) (*types.Transaction, error) {
	return _FeeHandler.Contract.SetFeeOracle(&_FeeHandler.TransactOpts, oracleAddress)
}

// SetFeeOracle is a paid mutator transaction binding the contract method 0xa8a98962.
//
// Solidity: function setFeeOracle(address oracleAddress) returns()
func (_FeeHandler *FeeHandlerTransactorSession) SetFeeOracle(oracleAddress common.Address) (*types.Transaction, error) {
	return _FeeHandler.Contract.SetFeeOracle(&_FeeHandler.TransactOpts, oracleAddress)
}

// SetFeeProperties is a paid mutator transaction binding the contract method 0xc297983f.
//
// Solidity: function setFeeProperties(uint32 gasUsed, uint16 feePercent) returns()
func (_FeeHandler *FeeHandlerTransactor) SetFeeProperties(opts *bind.TransactOpts, gasUsed uint32, feePercent uint16) (*types.Transaction, error) {
	return _FeeHandler.contract.Transact(opts, "setFeeProperties", gasUsed, feePercent)
}

// SetFeeProperties is a paid mutator transaction binding the contract method 0xc297983f.
//
// Solidity: function setFeeProperties(uint32 gasUsed, uint16 feePercent) returns()
func (_FeeHandler *FeeHandlerSession) SetFeeProperties(gasUsed uint32, feePercent uint16) (*types.Transaction, error) {
	return _FeeHandler.Contract.SetFeeProperties(&_FeeHandler.TransactOpts, gasUsed, feePercent)
}

// SetFeeProperties is a paid mutator transaction binding the contract method 0xc297983f.
//
// Solidity: function setFeeProperties(uint32 gasUsed, uint16 feePercent) returns()
func (_FeeHandler *FeeHandlerTransactorSession) SetFeeProperties(gasUsed uint32, feePercent uint16) (*types.Transaction, error) {
	return _FeeHandler.Contract.SetFeeProperties(&_FeeHandler.TransactOpts, gasUsed, feePercent)
}

// TransferFee is a paid mutator transaction binding the contract method 0xbff42755.
//
// Solidity: function transferFee(bytes32 resourceID, address[] addrs, uint256[] amounts) returns()
func (_FeeHandler *FeeHandlerTransactor) TransferFee(opts *bind.TransactOpts, resourceID [32]byte, addrs []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _FeeHandler.contract.Transact(opts, "transferFee", resourceID, addrs, amounts)
}

// TransferFee is a paid mutator transaction binding the contract method 0xbff42755.
//
// Solidity: function transferFee(bytes32 resourceID, address[] addrs, uint256[] amounts) returns()
func (_FeeHandler *FeeHandlerSession) TransferFee(resourceID [32]byte, addrs []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _FeeHandler.Contract.TransferFee(&_FeeHandler.TransactOpts, resourceID, addrs, amounts)
}

// TransferFee is a paid mutator transaction binding the contract method 0xbff42755.
//
// Solidity: function transferFee(bytes32 resourceID, address[] addrs, uint256[] amounts) returns()
func (_FeeHandler *FeeHandlerTransactorSession) TransferFee(resourceID [32]byte, addrs []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _FeeHandler.Contract.TransferFee(&_FeeHandler.TransactOpts, resourceID, addrs, amounts)
}

// FeeHandlerFeeCollectedIterator is returned from FilterFeeCollected and is used to iterate over the raw logs and unpacked data for FeeCollected events raised by the FeeHandler contract.
type FeeHandlerFeeCollectedIterator struct {
	Event *FeeHandlerFeeCollected // Event containing the contract specifics and raw log

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
func (it *FeeHandlerFeeCollectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeHandlerFeeCollected)
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
		it.Event = new(FeeHandlerFeeCollected)
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
func (it *FeeHandlerFeeCollectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeHandlerFeeCollectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeHandlerFeeCollected represents a FeeCollected event raised by the FeeHandler contract.
type FeeHandlerFeeCollected struct {
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
func (_FeeHandler *FeeHandlerFilterer) FilterFeeCollected(opts *bind.FilterOpts) (*FeeHandlerFeeCollectedIterator, error) {

	logs, sub, err := _FeeHandler.contract.FilterLogs(opts, "FeeCollected")
	if err != nil {
		return nil, err
	}
	return &FeeHandlerFeeCollectedIterator{contract: _FeeHandler.contract, event: "FeeCollected", logs: logs, sub: sub}, nil
}

// WatchFeeCollected is a free log subscription operation binding the contract event 0xbd231b7fa4103e15e7a238c72f07e8aff310701af121895aa6c793b80245e433.
//
// Solidity: event FeeCollected(address sender, uint8 fromDomainID, uint8 destinationDomainID, bytes32 resourceID, uint256 fee, address tokenAddress)
func (_FeeHandler *FeeHandlerFilterer) WatchFeeCollected(opts *bind.WatchOpts, sink chan<- *FeeHandlerFeeCollected) (event.Subscription, error) {

	logs, sub, err := _FeeHandler.contract.WatchLogs(opts, "FeeCollected")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeHandlerFeeCollected)
				if err := _FeeHandler.contract.UnpackLog(event, "FeeCollected", log); err != nil {
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
func (_FeeHandler *FeeHandlerFilterer) ParseFeeCollected(log types.Log) (*FeeHandlerFeeCollected, error) {
	event := new(FeeHandlerFeeCollected)
	if err := _FeeHandler.contract.UnpackLog(event, "FeeCollected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeHandlerFeeDistributedIterator is returned from FilterFeeDistributed and is used to iterate over the raw logs and unpacked data for FeeDistributed events raised by the FeeHandler contract.
type FeeHandlerFeeDistributedIterator struct {
	Event *FeeHandlerFeeDistributed // Event containing the contract specifics and raw log

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
func (it *FeeHandlerFeeDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeHandlerFeeDistributed)
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
		it.Event = new(FeeHandlerFeeDistributed)
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
func (it *FeeHandlerFeeDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeHandlerFeeDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeHandlerFeeDistributed represents a FeeDistributed event raised by the FeeHandler contract.
type FeeHandlerFeeDistributed struct {
	TokenAddress common.Address
	Recipient    common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFeeDistributed is a free log retrieval operation binding the contract event 0xaaa40a232aaf133fdd28f3485f6fdd163514cfadbffa981f3610f42398efe34b.
//
// Solidity: event FeeDistributed(address tokenAddress, address recipient, uint256 amount)
func (_FeeHandler *FeeHandlerFilterer) FilterFeeDistributed(opts *bind.FilterOpts) (*FeeHandlerFeeDistributedIterator, error) {

	logs, sub, err := _FeeHandler.contract.FilterLogs(opts, "FeeDistributed")
	if err != nil {
		return nil, err
	}
	return &FeeHandlerFeeDistributedIterator{contract: _FeeHandler.contract, event: "FeeDistributed", logs: logs, sub: sub}, nil
}

// WatchFeeDistributed is a free log subscription operation binding the contract event 0xaaa40a232aaf133fdd28f3485f6fdd163514cfadbffa981f3610f42398efe34b.
//
// Solidity: event FeeDistributed(address tokenAddress, address recipient, uint256 amount)
func (_FeeHandler *FeeHandlerFilterer) WatchFeeDistributed(opts *bind.WatchOpts, sink chan<- *FeeHandlerFeeDistributed) (event.Subscription, error) {

	logs, sub, err := _FeeHandler.contract.WatchLogs(opts, "FeeDistributed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeHandlerFeeDistributed)
				if err := _FeeHandler.contract.UnpackLog(event, "FeeDistributed", log); err != nil {
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
func (_FeeHandler *FeeHandlerFilterer) ParseFeeDistributed(log types.Log) (*FeeHandlerFeeDistributed, error) {
	event := new(FeeHandlerFeeDistributed)
	if err := _FeeHandler.contract.UnpackLog(event, "FeeDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeHandlerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the FeeHandler contract.
type FeeHandlerRoleAdminChangedIterator struct {
	Event *FeeHandlerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *FeeHandlerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeHandlerRoleAdminChanged)
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
		it.Event = new(FeeHandlerRoleAdminChanged)
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
func (it *FeeHandlerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeHandlerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeHandlerRoleAdminChanged represents a RoleAdminChanged event raised by the FeeHandler contract.
type FeeHandlerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_FeeHandler *FeeHandlerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*FeeHandlerRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _FeeHandler.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &FeeHandlerRoleAdminChangedIterator{contract: _FeeHandler.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_FeeHandler *FeeHandlerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *FeeHandlerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _FeeHandler.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeHandlerRoleAdminChanged)
				if err := _FeeHandler.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_FeeHandler *FeeHandlerFilterer) ParseRoleAdminChanged(log types.Log) (*FeeHandlerRoleAdminChanged, error) {
	event := new(FeeHandlerRoleAdminChanged)
	if err := _FeeHandler.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeHandlerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the FeeHandler contract.
type FeeHandlerRoleGrantedIterator struct {
	Event *FeeHandlerRoleGranted // Event containing the contract specifics and raw log

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
func (it *FeeHandlerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeHandlerRoleGranted)
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
		it.Event = new(FeeHandlerRoleGranted)
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
func (it *FeeHandlerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeHandlerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeHandlerRoleGranted represents a RoleGranted event raised by the FeeHandler contract.
type FeeHandlerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_FeeHandler *FeeHandlerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*FeeHandlerRoleGrantedIterator, error) {

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

	logs, sub, err := _FeeHandler.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &FeeHandlerRoleGrantedIterator{contract: _FeeHandler.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_FeeHandler *FeeHandlerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *FeeHandlerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _FeeHandler.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeHandlerRoleGranted)
				if err := _FeeHandler.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_FeeHandler *FeeHandlerFilterer) ParseRoleGranted(log types.Log) (*FeeHandlerRoleGranted, error) {
	event := new(FeeHandlerRoleGranted)
	if err := _FeeHandler.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeHandlerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the FeeHandler contract.
type FeeHandlerRoleRevokedIterator struct {
	Event *FeeHandlerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *FeeHandlerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeHandlerRoleRevoked)
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
		it.Event = new(FeeHandlerRoleRevoked)
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
func (it *FeeHandlerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeHandlerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeHandlerRoleRevoked represents a RoleRevoked event raised by the FeeHandler contract.
type FeeHandlerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_FeeHandler *FeeHandlerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*FeeHandlerRoleRevokedIterator, error) {

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

	logs, sub, err := _FeeHandler.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &FeeHandlerRoleRevokedIterator{contract: _FeeHandler.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_FeeHandler *FeeHandlerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *FeeHandlerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _FeeHandler.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeHandlerRoleRevoked)
				if err := _FeeHandler.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_FeeHandler *FeeHandlerFilterer) ParseRoleRevoked(log types.Log) (*FeeHandlerRoleRevoked, error) {
	event := new(FeeHandlerRoleRevoked)
	if err := _FeeHandler.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
