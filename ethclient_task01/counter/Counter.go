// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package counter

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
	_ = abi.ConvertType
)

// CounterMetaData contains all meta data concerning the Counter contract.
var CounterMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"operation\",\"type\":\"string\"}],\"name\":\"CountChangeEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"decrement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506104c88061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c80632baeceb71461004e578063a87d942c14610058578063d09de08a14610076578063d826f88f14610080575b5f5ffd5b61005661008a565b005b610060610131565b60405161006d919061020c565b60405180910390f35b61007e610139565b005b61008861019d565b005b5f5f54116100cd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100c49061027f565b60405180910390fd5b60015f546100db91906102ca565b5f819055503373ffffffffffffffffffffffffffffffffffffffff167fb8a4f7b9920a91e295f46ecd734a2ec096aae65e512624541b4296e821423c925f546040516101279190610347565b60405180910390a2565b5f5f54905090565b60015f546101479190610373565b5f819055503373ffffffffffffffffffffffffffffffffffffffff167fb8a4f7b9920a91e295f46ecd734a2ec096aae65e512624541b4296e821423c925f5460405161019391906103f0565b60405180910390a2565b5f5f819055503373ffffffffffffffffffffffffffffffffffffffff167fb8a4f7b9920a91e295f46ecd734a2ec096aae65e512624541b4296e821423c925f546040516101ea9190610466565b60405180910390a2565b5f819050919050565b610206816101f4565b82525050565b5f60208201905061021f5f8301846101fd565b92915050565b5f82825260208201905092915050565b7f636f756e74206d757374206265203e20300000000000000000000000000000005f82015250565b5f610269601183610225565b915061027482610235565b602082019050919050565b5f6020820190508181035f8301526102968161025d565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6102d4826101f4565b91506102df836101f4565b92508282039050818111156102f7576102f661029d565b5b92915050565b7f64656372656d656e7400000000000000000000000000000000000000000000005f82015250565b5f610331600983610225565b915061033c826102fd565b602082019050919050565b5f60408201905061035a5f8301846101fd565b818103602083015261036b81610325565b905092915050565b5f61037d826101f4565b9150610388836101f4565b92508282019050808211156103a05761039f61029d565b5b92915050565b7f696e6372656d656e7400000000000000000000000000000000000000000000005f82015250565b5f6103da600983610225565b91506103e5826103a6565b602082019050919050565b5f6040820190506104035f8301846101fd565b8181036020830152610414816103ce565b905092915050565b7f72657365740000000000000000000000000000000000000000000000000000005f82015250565b5f610450600583610225565b915061045b8261041c565b602082019050919050565b5f6040820190506104795f8301846101fd565b818103602083015261048a81610444565b90509291505056fea264697066735822122090b06b0198e29cca20d939e9258208a9d1167261b47af2ef9712155989d4e90364736f6c634300081e0033",
}

// CounterABI is the input ABI used to generate the binding from.
// Deprecated: Use CounterMetaData.ABI instead.
var CounterABI = CounterMetaData.ABI

// CounterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CounterMetaData.Bin instead.
var CounterBin = CounterMetaData.Bin

// DeployCounter deploys a new Ethereum contract, binding an instance of Counter to it.
func DeployCounter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Counter, error) {
	parsed, err := CounterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CounterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Counter{CounterCaller: CounterCaller{contract: contract}, CounterTransactor: CounterTransactor{contract: contract}, CounterFilterer: CounterFilterer{contract: contract}}, nil
}

// Counter is an auto generated Go binding around an Ethereum contract.
type Counter struct {
	CounterCaller     // Read-only binding to the contract
	CounterTransactor // Write-only binding to the contract
	CounterFilterer   // Log filterer for contract events
}

// CounterCaller is an auto generated read-only Go binding around an Ethereum contract.
type CounterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CounterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CounterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CounterSession struct {
	Contract     *Counter          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CounterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CounterCallerSession struct {
	Contract *CounterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// CounterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CounterTransactorSession struct {
	Contract     *CounterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CounterRaw is an auto generated low-level Go binding around an Ethereum contract.
type CounterRaw struct {
	Contract *Counter // Generic contract binding to access the raw methods on
}

// CounterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CounterCallerRaw struct {
	Contract *CounterCaller // Generic read-only contract binding to access the raw methods on
}

// CounterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CounterTransactorRaw struct {
	Contract *CounterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCounter creates a new instance of Counter, bound to a specific deployed contract.
func NewCounter(address common.Address, backend bind.ContractBackend) (*Counter, error) {
	contract, err := bindCounter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Counter{CounterCaller: CounterCaller{contract: contract}, CounterTransactor: CounterTransactor{contract: contract}, CounterFilterer: CounterFilterer{contract: contract}}, nil
}

// NewCounterCaller creates a new read-only instance of Counter, bound to a specific deployed contract.
func NewCounterCaller(address common.Address, caller bind.ContractCaller) (*CounterCaller, error) {
	contract, err := bindCounter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CounterCaller{contract: contract}, nil
}

// NewCounterTransactor creates a new write-only instance of Counter, bound to a specific deployed contract.
func NewCounterTransactor(address common.Address, transactor bind.ContractTransactor) (*CounterTransactor, error) {
	contract, err := bindCounter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CounterTransactor{contract: contract}, nil
}

// NewCounterFilterer creates a new log filterer instance of Counter, bound to a specific deployed contract.
func NewCounterFilterer(address common.Address, filterer bind.ContractFilterer) (*CounterFilterer, error) {
	contract, err := bindCounter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CounterFilterer{contract: contract}, nil
}

// bindCounter binds a generic wrapper to an already deployed contract.
func bindCounter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CounterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Counter *CounterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Counter.Contract.CounterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Counter *CounterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.Contract.CounterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Counter *CounterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Counter.Contract.CounterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Counter *CounterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Counter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Counter *CounterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Counter *CounterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Counter.Contract.contract.Transact(opts, method, params...)
}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_Counter *CounterCaller) GetCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Counter.contract.Call(opts, &out, "getCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_Counter *CounterSession) GetCount() (*big.Int, error) {
	return _Counter.Contract.GetCount(&_Counter.CallOpts)
}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_Counter *CounterCallerSession) GetCount() (*big.Int, error) {
	return _Counter.Contract.GetCount(&_Counter.CallOpts)
}

// Decrement is a paid mutator transaction binding the contract method 0x2baeceb7.
//
// Solidity: function decrement() returns()
func (_Counter *CounterTransactor) Decrement(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "decrement")
}

// Decrement is a paid mutator transaction binding the contract method 0x2baeceb7.
//
// Solidity: function decrement() returns()
func (_Counter *CounterSession) Decrement() (*types.Transaction, error) {
	return _Counter.Contract.Decrement(&_Counter.TransactOpts)
}

// Decrement is a paid mutator transaction binding the contract method 0x2baeceb7.
//
// Solidity: function decrement() returns()
func (_Counter *CounterTransactorSession) Decrement() (*types.Transaction, error) {
	return _Counter.Contract.Decrement(&_Counter.TransactOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_Counter *CounterTransactor) Increment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "increment")
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_Counter *CounterSession) Increment() (*types.Transaction, error) {
	return _Counter.Contract.Increment(&_Counter.TransactOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_Counter *CounterTransactorSession) Increment() (*types.Transaction, error) {
	return _Counter.Contract.Increment(&_Counter.TransactOpts)
}

// Reset is a paid mutator transaction binding the contract method 0xd826f88f.
//
// Solidity: function reset() returns()
func (_Counter *CounterTransactor) Reset(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "reset")
}

// Reset is a paid mutator transaction binding the contract method 0xd826f88f.
//
// Solidity: function reset() returns()
func (_Counter *CounterSession) Reset() (*types.Transaction, error) {
	return _Counter.Contract.Reset(&_Counter.TransactOpts)
}

// Reset is a paid mutator transaction binding the contract method 0xd826f88f.
//
// Solidity: function reset() returns()
func (_Counter *CounterTransactorSession) Reset() (*types.Transaction, error) {
	return _Counter.Contract.Reset(&_Counter.TransactOpts)
}

// CounterCountChangeEventIterator is returned from FilterCountChangeEvent and is used to iterate over the raw logs and unpacked data for CountChangeEvent events raised by the Counter contract.
type CounterCountChangeEventIterator struct {
	Event *CounterCountChangeEvent // Event containing the contract specifics and raw log

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
func (it *CounterCountChangeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CounterCountChangeEvent)
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
		it.Event = new(CounterCountChangeEvent)
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
func (it *CounterCountChangeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CounterCountChangeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CounterCountChangeEvent represents a CountChangeEvent event raised by the Counter contract.
type CounterCountChangeEvent struct {
	Operator  common.Address
	NewCount  *big.Int
	Operation string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCountChangeEvent is a free log retrieval operation binding the contract event 0xb8a4f7b9920a91e295f46ecd734a2ec096aae65e512624541b4296e821423c92.
//
// Solidity: event CountChangeEvent(address indexed operator, uint256 newCount, string operation)
func (_Counter *CounterFilterer) FilterCountChangeEvent(opts *bind.FilterOpts, operator []common.Address) (*CounterCountChangeEventIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Counter.contract.FilterLogs(opts, "CountChangeEvent", operatorRule)
	if err != nil {
		return nil, err
	}
	return &CounterCountChangeEventIterator{contract: _Counter.contract, event: "CountChangeEvent", logs: logs, sub: sub}, nil
}

// WatchCountChangeEvent is a free log subscription operation binding the contract event 0xb8a4f7b9920a91e295f46ecd734a2ec096aae65e512624541b4296e821423c92.
//
// Solidity: event CountChangeEvent(address indexed operator, uint256 newCount, string operation)
func (_Counter *CounterFilterer) WatchCountChangeEvent(opts *bind.WatchOpts, sink chan<- *CounterCountChangeEvent, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Counter.contract.WatchLogs(opts, "CountChangeEvent", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CounterCountChangeEvent)
				if err := _Counter.contract.UnpackLog(event, "CountChangeEvent", log); err != nil {
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

// ParseCountChangeEvent is a log parse operation binding the contract event 0xb8a4f7b9920a91e295f46ecd734a2ec096aae65e512624541b4296e821423c92.
//
// Solidity: event CountChangeEvent(address indexed operator, uint256 newCount, string operation)
func (_Counter *CounterFilterer) ParseCountChangeEvent(log types.Log) (*CounterCountChangeEvent, error) {
	event := new(CounterCountChangeEvent)
	if err := _Counter.contract.UnpackLog(event, "CountChangeEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
