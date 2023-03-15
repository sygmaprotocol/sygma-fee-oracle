// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package TestContracts

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

// TestContractsABI is the input ABI used to generate the binding from.
const TestContractsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"asset\",\"type\":\"bytes32\"}],\"name\":\"AssetStored\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_assetsStored\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"asset\",\"type\":\"bytes32\"}],\"name\":\"store\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"asset\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"depositorCheck\",\"type\":\"address\"}],\"name\":\"storeWithDepositor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TestContractsBin is the compiled bytecode used for deploying new contracts.
var TestContractsBin = "0x608060405234801561001057600080fd5b50610522806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063654cf88c1461004657806396add60014610062578063ea287d1514610092575b600080fd5b610060600480360381019061005b91906102ef565b6100ae565b005b61007c600480360381019061007791906102ef565b610169565b6040516100899190610337565b60405180910390f35b6100ac60048036038101906100a791906103b0565b610189565b005b60008082815260200190815260200160002060009054906101000a900460ff161561010e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161010590610460565b60405180910390fd5b600160008083815260200190815260200160002060006101000a81548160ff021916908315150217905550807f08ae553713effae7116be03743b167b8b803449ee8fb912c2ec43dc2c824f53560405160405180910390a250565b60006020528060005260406000206000915054906101000a900460ff1681565b60008083815260200190815260200160002060009054906101000a900460ff16156101e9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101e090610460565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614610257576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161024e906104cc565b60405180910390fd5b600160008084815260200190815260200160002060006101000a81548160ff021916908315150217905550817f08ae553713effae7116be03743b167b8b803449ee8fb912c2ec43dc2c824f53560405160405180910390a2505050565b600080fd5b6000819050919050565b6102cc816102b9565b81146102d757600080fd5b50565b6000813590506102e9816102c3565b92915050565b600060208284031215610305576103046102b4565b5b6000610313848285016102da565b91505092915050565b60008115159050919050565b6103318161031c565b82525050565b600060208201905061034c6000830184610328565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061037d82610352565b9050919050565b61038d81610372565b811461039857600080fd5b50565b6000813590506103aa81610384565b92915050565b6000806000606084860312156103c9576103c86102b4565b5b60006103d78682870161039b565b93505060206103e8868287016102da565b92505060406103f98682870161039b565b9150509250925092565b600082825260208201905092915050565b7f617373657420697320616c72656164792073746f726564000000000000000000600082015250565b600061044a601783610403565b915061045582610414565b602082019050919050565b600060208201905081810360008301526104798161043d565b9050919050565b7f696e76616c6964206465706f7369746f72206164647265737300000000000000600082015250565b60006104b6601983610403565b91506104c182610480565b602082019050919050565b600060208201905081810360008301526104e5816104a9565b905091905056fea26469706673582212207d589e0d582ae248bf7ffac7dbc42a838d5b705703a50b567ae9b5ed0b346e9d64736f6c634300080b0033"

// DeployTestContracts deploys a new Ethereum contract, binding an instance of TestContracts to it.
func DeployTestContracts(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestContracts, error) {
	parsed, err := abi.JSON(strings.NewReader(TestContractsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TestContractsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestContracts{TestContractsCaller: TestContractsCaller{contract: contract}, TestContractsTransactor: TestContractsTransactor{contract: contract}, TestContractsFilterer: TestContractsFilterer{contract: contract}}, nil
}

// TestContracts is an auto generated Go binding around an Ethereum contract.
type TestContracts struct {
	TestContractsCaller     // Read-only binding to the contract
	TestContractsTransactor // Write-only binding to the contract
	TestContractsFilterer   // Log filterer for contract events
}

// TestContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestContractsSession struct {
	Contract     *TestContracts    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestContractsCallerSession struct {
	Contract *TestContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TestContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestContractsTransactorSession struct {
	Contract     *TestContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TestContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestContractsRaw struct {
	Contract *TestContracts // Generic contract binding to access the raw methods on
}

// TestContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestContractsCallerRaw struct {
	Contract *TestContractsCaller // Generic read-only contract binding to access the raw methods on
}

// TestContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestContractsTransactorRaw struct {
	Contract *TestContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestContracts creates a new instance of TestContracts, bound to a specific deployed contract.
func NewTestContracts(address common.Address, backend bind.ContractBackend) (*TestContracts, error) {
	contract, err := bindTestContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestContracts{TestContractsCaller: TestContractsCaller{contract: contract}, TestContractsTransactor: TestContractsTransactor{contract: contract}, TestContractsFilterer: TestContractsFilterer{contract: contract}}, nil
}

// NewTestContractsCaller creates a new read-only instance of TestContracts, bound to a specific deployed contract.
func NewTestContractsCaller(address common.Address, caller bind.ContractCaller) (*TestContractsCaller, error) {
	contract, err := bindTestContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestContractsCaller{contract: contract}, nil
}

// NewTestContractsTransactor creates a new write-only instance of TestContracts, bound to a specific deployed contract.
func NewTestContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*TestContractsTransactor, error) {
	contract, err := bindTestContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestContractsTransactor{contract: contract}, nil
}

// NewTestContractsFilterer creates a new log filterer instance of TestContracts, bound to a specific deployed contract.
func NewTestContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*TestContractsFilterer, error) {
	contract, err := bindTestContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestContractsFilterer{contract: contract}, nil
}

// bindTestContracts binds a generic wrapper to an already deployed contract.
func bindTestContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestContractsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestContracts *TestContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestContracts.Contract.TestContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestContracts *TestContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestContracts.Contract.TestContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestContracts *TestContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestContracts.Contract.TestContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestContracts *TestContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestContracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestContracts *TestContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestContracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestContracts *TestContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestContracts.Contract.contract.Transact(opts, method, params...)
}

// AssetsStored is a free data retrieval call binding the contract method 0x96add600.
//
// Solidity: function _assetsStored(bytes32 ) view returns(bool)
func (_TestContracts *TestContractsCaller) AssetsStored(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _TestContracts.contract.Call(opts, &out, "_assetsStored", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AssetsStored is a free data retrieval call binding the contract method 0x96add600.
//
// Solidity: function _assetsStored(bytes32 ) view returns(bool)
func (_TestContracts *TestContractsSession) AssetsStored(arg0 [32]byte) (bool, error) {
	return _TestContracts.Contract.AssetsStored(&_TestContracts.CallOpts, arg0)
}

// AssetsStored is a free data retrieval call binding the contract method 0x96add600.
//
// Solidity: function _assetsStored(bytes32 ) view returns(bool)
func (_TestContracts *TestContractsCallerSession) AssetsStored(arg0 [32]byte) (bool, error) {
	return _TestContracts.Contract.AssetsStored(&_TestContracts.CallOpts, arg0)
}

// Store is a paid mutator transaction binding the contract method 0x654cf88c.
//
// Solidity: function store(bytes32 asset) returns()
func (_TestContracts *TestContractsTransactor) Store(opts *bind.TransactOpts, asset [32]byte) (*types.Transaction, error) {
	return _TestContracts.contract.Transact(opts, "store", asset)
}

// Store is a paid mutator transaction binding the contract method 0x654cf88c.
//
// Solidity: function store(bytes32 asset) returns()
func (_TestContracts *TestContractsSession) Store(asset [32]byte) (*types.Transaction, error) {
	return _TestContracts.Contract.Store(&_TestContracts.TransactOpts, asset)
}

// Store is a paid mutator transaction binding the contract method 0x654cf88c.
//
// Solidity: function store(bytes32 asset) returns()
func (_TestContracts *TestContractsTransactorSession) Store(asset [32]byte) (*types.Transaction, error) {
	return _TestContracts.Contract.Store(&_TestContracts.TransactOpts, asset)
}

// StoreWithDepositor is a paid mutator transaction binding the contract method 0xea287d15.
//
// Solidity: function storeWithDepositor(address depositor, bytes32 asset, address depositorCheck) returns()
func (_TestContracts *TestContractsTransactor) StoreWithDepositor(opts *bind.TransactOpts, depositor common.Address, asset [32]byte, depositorCheck common.Address) (*types.Transaction, error) {
	return _TestContracts.contract.Transact(opts, "storeWithDepositor", depositor, asset, depositorCheck)
}

// StoreWithDepositor is a paid mutator transaction binding the contract method 0xea287d15.
//
// Solidity: function storeWithDepositor(address depositor, bytes32 asset, address depositorCheck) returns()
func (_TestContracts *TestContractsSession) StoreWithDepositor(depositor common.Address, asset [32]byte, depositorCheck common.Address) (*types.Transaction, error) {
	return _TestContracts.Contract.StoreWithDepositor(&_TestContracts.TransactOpts, depositor, asset, depositorCheck)
}

// StoreWithDepositor is a paid mutator transaction binding the contract method 0xea287d15.
//
// Solidity: function storeWithDepositor(address depositor, bytes32 asset, address depositorCheck) returns()
func (_TestContracts *TestContractsTransactorSession) StoreWithDepositor(depositor common.Address, asset [32]byte, depositorCheck common.Address) (*types.Transaction, error) {
	return _TestContracts.Contract.StoreWithDepositor(&_TestContracts.TransactOpts, depositor, asset, depositorCheck)
}

// TestContractsAssetStoredIterator is returned from FilterAssetStored and is used to iterate over the raw logs and unpacked data for AssetStored events raised by the TestContracts contract.
type TestContractsAssetStoredIterator struct {
	Event *TestContractsAssetStored // Event containing the contract specifics and raw log

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
func (it *TestContractsAssetStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestContractsAssetStored)
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
		it.Event = new(TestContractsAssetStored)
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
func (it *TestContractsAssetStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestContractsAssetStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestContractsAssetStored represents a AssetStored event raised by the TestContracts contract.
type TestContractsAssetStored struct {
	Asset [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAssetStored is a free log retrieval operation binding the contract event 0x08ae553713effae7116be03743b167b8b803449ee8fb912c2ec43dc2c824f535.
//
// Solidity: event AssetStored(bytes32 indexed asset)
func (_TestContracts *TestContractsFilterer) FilterAssetStored(opts *bind.FilterOpts, asset [][32]byte) (*TestContractsAssetStoredIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _TestContracts.contract.FilterLogs(opts, "AssetStored", assetRule)
	if err != nil {
		return nil, err
	}
	return &TestContractsAssetStoredIterator{contract: _TestContracts.contract, event: "AssetStored", logs: logs, sub: sub}, nil
}

// WatchAssetStored is a free log subscription operation binding the contract event 0x08ae553713effae7116be03743b167b8b803449ee8fb912c2ec43dc2c824f535.
//
// Solidity: event AssetStored(bytes32 indexed asset)
func (_TestContracts *TestContractsFilterer) WatchAssetStored(opts *bind.WatchOpts, sink chan<- *TestContractsAssetStored, asset [][32]byte) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _TestContracts.contract.WatchLogs(opts, "AssetStored", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestContractsAssetStored)
				if err := _TestContracts.contract.UnpackLog(event, "AssetStored", log); err != nil {
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

// ParseAssetStored is a log parse operation binding the contract event 0x08ae553713effae7116be03743b167b8b803449ee8fb912c2ec43dc2c824f535.
//
// Solidity: event AssetStored(bytes32 indexed asset)
func (_TestContracts *TestContractsFilterer) ParseAssetStored(log types.Log) (*TestContractsAssetStored, error) {
	event := new(TestContractsAssetStored)
	if err := _TestContracts.contract.UnpackLog(event, "AssetStored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
