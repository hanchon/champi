// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package storecore

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

// StorecoreMetaData contains all meta data concerning the Storecore contract.
var StorecoreMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"ResourceId\",\"name\":\"tableId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"keyTuple\",\"type\":\"bytes32[]\"}],\"name\":\"Store_DeleteRecord\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"ResourceId\",\"name\":\"tableId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"keyTuple\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"staticData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"PackedCounter\",\"name\":\"encodedLengths\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"dynamicData\",\"type\":\"bytes\"}],\"name\":\"Store_SetRecord\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"ResourceId\",\"name\":\"tableId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"keyTuple\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"start\",\"type\":\"uint48\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"deleteCount\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"PackedCounter\",\"name\":\"encodedLengths\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Store_SpliceDynamicData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"ResourceId\",\"name\":\"tableId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"keyTuple\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"start\",\"type\":\"uint48\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Store_SpliceStaticData\",\"type\":\"event\"}]",
}

// StorecoreABI is the input ABI used to generate the binding from.
// Deprecated: Use StorecoreMetaData.ABI instead.
var StorecoreABI = StorecoreMetaData.ABI

// Storecore is an auto generated Go binding around an Ethereum contract.
type Storecore struct {
	StorecoreCaller     // Read-only binding to the contract
	StorecoreTransactor // Write-only binding to the contract
	StorecoreFilterer   // Log filterer for contract events
}

// StorecoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorecoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorecoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorecoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorecoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorecoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorecoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorecoreSession struct {
	Contract     *Storecore        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StorecoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorecoreCallerSession struct {
	Contract *StorecoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// StorecoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorecoreTransactorSession struct {
	Contract     *StorecoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// StorecoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorecoreRaw struct {
	Contract *Storecore // Generic contract binding to access the raw methods on
}

// StorecoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorecoreCallerRaw struct {
	Contract *StorecoreCaller // Generic read-only contract binding to access the raw methods on
}

// StorecoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorecoreTransactorRaw struct {
	Contract *StorecoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorecore creates a new instance of Storecore, bound to a specific deployed contract.
func NewStorecore(address common.Address, backend bind.ContractBackend) (*Storecore, error) {
	contract, err := bindStorecore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Storecore{StorecoreCaller: StorecoreCaller{contract: contract}, StorecoreTransactor: StorecoreTransactor{contract: contract}, StorecoreFilterer: StorecoreFilterer{contract: contract}}, nil
}

// NewStorecoreCaller creates a new read-only instance of Storecore, bound to a specific deployed contract.
func NewStorecoreCaller(address common.Address, caller bind.ContractCaller) (*StorecoreCaller, error) {
	contract, err := bindStorecore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorecoreCaller{contract: contract}, nil
}

// NewStorecoreTransactor creates a new write-only instance of Storecore, bound to a specific deployed contract.
func NewStorecoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StorecoreTransactor, error) {
	contract, err := bindStorecore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorecoreTransactor{contract: contract}, nil
}

// NewStorecoreFilterer creates a new log filterer instance of Storecore, bound to a specific deployed contract.
func NewStorecoreFilterer(address common.Address, filterer bind.ContractFilterer) (*StorecoreFilterer, error) {
	contract, err := bindStorecore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorecoreFilterer{contract: contract}, nil
}

// bindStorecore binds a generic wrapper to an already deployed contract.
func bindStorecore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StorecoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storecore *StorecoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storecore.Contract.StorecoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storecore *StorecoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storecore.Contract.StorecoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storecore *StorecoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storecore.Contract.StorecoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storecore *StorecoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storecore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storecore *StorecoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storecore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storecore *StorecoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storecore.Contract.contract.Transact(opts, method, params...)
}

// StorecoreStoreDeleteRecordIterator is returned from FilterStoreDeleteRecord and is used to iterate over the raw logs and unpacked data for StoreDeleteRecord events raised by the Storecore contract.
type StorecoreStoreDeleteRecordIterator struct {
	Event *StorecoreStoreDeleteRecord // Event containing the contract specifics and raw log

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
func (it *StorecoreStoreDeleteRecordIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorecoreStoreDeleteRecord)
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
		it.Event = new(StorecoreStoreDeleteRecord)
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
func (it *StorecoreStoreDeleteRecordIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorecoreStoreDeleteRecordIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorecoreStoreDeleteRecord represents a StoreDeleteRecord event raised by the Storecore contract.
type StorecoreStoreDeleteRecord struct {
	TableId  [32]byte
	KeyTuple [][32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStoreDeleteRecord is a free log retrieval operation binding the contract event 0x0e1f72f429eb97e64878619984a91e687ae91610348b9ff4216782cc96e49d07.
//
// Solidity: event Store_DeleteRecord(bytes32 indexed tableId, bytes32[] keyTuple)
func (_Storecore *StorecoreFilterer) FilterStoreDeleteRecord(opts *bind.FilterOpts, tableId [][32]byte) (*StorecoreStoreDeleteRecordIterator, error) {

	var tableIdRule []interface{}
	for _, tableIdItem := range tableId {
		tableIdRule = append(tableIdRule, tableIdItem)
	}

	logs, sub, err := _Storecore.contract.FilterLogs(opts, "Store_DeleteRecord", tableIdRule)
	if err != nil {
		return nil, err
	}
	return &StorecoreStoreDeleteRecordIterator{contract: _Storecore.contract, event: "Store_DeleteRecord", logs: logs, sub: sub}, nil
}

// WatchStoreDeleteRecord is a free log subscription operation binding the contract event 0x0e1f72f429eb97e64878619984a91e687ae91610348b9ff4216782cc96e49d07.
//
// Solidity: event Store_DeleteRecord(bytes32 indexed tableId, bytes32[] keyTuple)
func (_Storecore *StorecoreFilterer) WatchStoreDeleteRecord(opts *bind.WatchOpts, sink chan<- *StorecoreStoreDeleteRecord, tableId [][32]byte) (event.Subscription, error) {

	var tableIdRule []interface{}
	for _, tableIdItem := range tableId {
		tableIdRule = append(tableIdRule, tableIdItem)
	}

	logs, sub, err := _Storecore.contract.WatchLogs(opts, "Store_DeleteRecord", tableIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorecoreStoreDeleteRecord)
				if err := _Storecore.contract.UnpackLog(event, "Store_DeleteRecord", log); err != nil {
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

// ParseStoreDeleteRecord is a log parse operation binding the contract event 0x0e1f72f429eb97e64878619984a91e687ae91610348b9ff4216782cc96e49d07.
//
// Solidity: event Store_DeleteRecord(bytes32 indexed tableId, bytes32[] keyTuple)
func (_Storecore *StorecoreFilterer) ParseStoreDeleteRecord(log types.Log) (*StorecoreStoreDeleteRecord, error) {
	event := new(StorecoreStoreDeleteRecord)
	if err := _Storecore.contract.UnpackLog(event, "Store_DeleteRecord", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorecoreStoreSetRecordIterator is returned from FilterStoreSetRecord and is used to iterate over the raw logs and unpacked data for StoreSetRecord events raised by the Storecore contract.
type StorecoreStoreSetRecordIterator struct {
	Event *StorecoreStoreSetRecord // Event containing the contract specifics and raw log

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
func (it *StorecoreStoreSetRecordIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorecoreStoreSetRecord)
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
		it.Event = new(StorecoreStoreSetRecord)
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
func (it *StorecoreStoreSetRecordIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorecoreStoreSetRecordIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorecoreStoreSetRecord represents a StoreSetRecord event raised by the Storecore contract.
type StorecoreStoreSetRecord struct {
	TableId        [32]byte
	KeyTuple       [][32]byte
	StaticData     []byte
	EncodedLengths [32]byte
	DynamicData    []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterStoreSetRecord is a free log retrieval operation binding the contract event 0x8dbb3a9672eebfd3773e72dd9c102393436816d832c7ba9e1e1ac8fcadcac7a9.
//
// Solidity: event Store_SetRecord(bytes32 indexed tableId, bytes32[] keyTuple, bytes staticData, bytes32 encodedLengths, bytes dynamicData)
func (_Storecore *StorecoreFilterer) FilterStoreSetRecord(opts *bind.FilterOpts, tableId [][32]byte) (*StorecoreStoreSetRecordIterator, error) {

	var tableIdRule []interface{}
	for _, tableIdItem := range tableId {
		tableIdRule = append(tableIdRule, tableIdItem)
	}

	logs, sub, err := _Storecore.contract.FilterLogs(opts, "Store_SetRecord", tableIdRule)
	if err != nil {
		return nil, err
	}
	return &StorecoreStoreSetRecordIterator{contract: _Storecore.contract, event: "Store_SetRecord", logs: logs, sub: sub}, nil
}

// WatchStoreSetRecord is a free log subscription operation binding the contract event 0x8dbb3a9672eebfd3773e72dd9c102393436816d832c7ba9e1e1ac8fcadcac7a9.
//
// Solidity: event Store_SetRecord(bytes32 indexed tableId, bytes32[] keyTuple, bytes staticData, bytes32 encodedLengths, bytes dynamicData)
func (_Storecore *StorecoreFilterer) WatchStoreSetRecord(opts *bind.WatchOpts, sink chan<- *StorecoreStoreSetRecord, tableId [][32]byte) (event.Subscription, error) {

	var tableIdRule []interface{}
	for _, tableIdItem := range tableId {
		tableIdRule = append(tableIdRule, tableIdItem)
	}

	logs, sub, err := _Storecore.contract.WatchLogs(opts, "Store_SetRecord", tableIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorecoreStoreSetRecord)
				if err := _Storecore.contract.UnpackLog(event, "Store_SetRecord", log); err != nil {
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

// ParseStoreSetRecord is a log parse operation binding the contract event 0x8dbb3a9672eebfd3773e72dd9c102393436816d832c7ba9e1e1ac8fcadcac7a9.
//
// Solidity: event Store_SetRecord(bytes32 indexed tableId, bytes32[] keyTuple, bytes staticData, bytes32 encodedLengths, bytes dynamicData)
func (_Storecore *StorecoreFilterer) ParseStoreSetRecord(log types.Log) (*StorecoreStoreSetRecord, error) {
	event := new(StorecoreStoreSetRecord)
	if err := _Storecore.contract.UnpackLog(event, "Store_SetRecord", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorecoreStoreSpliceDynamicDataIterator is returned from FilterStoreSpliceDynamicData and is used to iterate over the raw logs and unpacked data for StoreSpliceDynamicData events raised by the Storecore contract.
type StorecoreStoreSpliceDynamicDataIterator struct {
	Event *StorecoreStoreSpliceDynamicData // Event containing the contract specifics and raw log

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
func (it *StorecoreStoreSpliceDynamicDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorecoreStoreSpliceDynamicData)
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
		it.Event = new(StorecoreStoreSpliceDynamicData)
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
func (it *StorecoreStoreSpliceDynamicDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorecoreStoreSpliceDynamicDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorecoreStoreSpliceDynamicData represents a StoreSpliceDynamicData event raised by the Storecore contract.
type StorecoreStoreSpliceDynamicData struct {
	TableId        [32]byte
	KeyTuple       [][32]byte
	Start          *big.Int
	DeleteCount    *big.Int
	EncodedLengths [32]byte
	Data           []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterStoreSpliceDynamicData is a free log retrieval operation binding the contract event 0xaa63765a776145e5e6492f471ae097dfed11cd57a61bc2679dd43180422385b4.
//
// Solidity: event Store_SpliceDynamicData(bytes32 indexed tableId, bytes32[] keyTuple, uint48 start, uint40 deleteCount, bytes32 encodedLengths, bytes data)
func (_Storecore *StorecoreFilterer) FilterStoreSpliceDynamicData(opts *bind.FilterOpts, tableId [][32]byte) (*StorecoreStoreSpliceDynamicDataIterator, error) {

	var tableIdRule []interface{}
	for _, tableIdItem := range tableId {
		tableIdRule = append(tableIdRule, tableIdItem)
	}

	logs, sub, err := _Storecore.contract.FilterLogs(opts, "Store_SpliceDynamicData", tableIdRule)
	if err != nil {
		return nil, err
	}
	return &StorecoreStoreSpliceDynamicDataIterator{contract: _Storecore.contract, event: "Store_SpliceDynamicData", logs: logs, sub: sub}, nil
}

// WatchStoreSpliceDynamicData is a free log subscription operation binding the contract event 0xaa63765a776145e5e6492f471ae097dfed11cd57a61bc2679dd43180422385b4.
//
// Solidity: event Store_SpliceDynamicData(bytes32 indexed tableId, bytes32[] keyTuple, uint48 start, uint40 deleteCount, bytes32 encodedLengths, bytes data)
func (_Storecore *StorecoreFilterer) WatchStoreSpliceDynamicData(opts *bind.WatchOpts, sink chan<- *StorecoreStoreSpliceDynamicData, tableId [][32]byte) (event.Subscription, error) {

	var tableIdRule []interface{}
	for _, tableIdItem := range tableId {
		tableIdRule = append(tableIdRule, tableIdItem)
	}

	logs, sub, err := _Storecore.contract.WatchLogs(opts, "Store_SpliceDynamicData", tableIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorecoreStoreSpliceDynamicData)
				if err := _Storecore.contract.UnpackLog(event, "Store_SpliceDynamicData", log); err != nil {
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

// ParseStoreSpliceDynamicData is a log parse operation binding the contract event 0xaa63765a776145e5e6492f471ae097dfed11cd57a61bc2679dd43180422385b4.
//
// Solidity: event Store_SpliceDynamicData(bytes32 indexed tableId, bytes32[] keyTuple, uint48 start, uint40 deleteCount, bytes32 encodedLengths, bytes data)
func (_Storecore *StorecoreFilterer) ParseStoreSpliceDynamicData(log types.Log) (*StorecoreStoreSpliceDynamicData, error) {
	event := new(StorecoreStoreSpliceDynamicData)
	if err := _Storecore.contract.UnpackLog(event, "Store_SpliceDynamicData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorecoreStoreSpliceStaticDataIterator is returned from FilterStoreSpliceStaticData and is used to iterate over the raw logs and unpacked data for StoreSpliceStaticData events raised by the Storecore contract.
type StorecoreStoreSpliceStaticDataIterator struct {
	Event *StorecoreStoreSpliceStaticData // Event containing the contract specifics and raw log

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
func (it *StorecoreStoreSpliceStaticDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorecoreStoreSpliceStaticData)
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
		it.Event = new(StorecoreStoreSpliceStaticData)
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
func (it *StorecoreStoreSpliceStaticDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorecoreStoreSpliceStaticDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorecoreStoreSpliceStaticData represents a StoreSpliceStaticData event raised by the Storecore contract.
type StorecoreStoreSpliceStaticData struct {
	TableId  [32]byte
	KeyTuple [][32]byte
	Start    *big.Int
	Data     []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStoreSpliceStaticData is a free log retrieval operation binding the contract event 0x8c0b5119d4cec7b284c6b1b39252a03d1e2f2d7451a5895562524c113bb952be.
//
// Solidity: event Store_SpliceStaticData(bytes32 indexed tableId, bytes32[] keyTuple, uint48 start, bytes data)
func (_Storecore *StorecoreFilterer) FilterStoreSpliceStaticData(opts *bind.FilterOpts, tableId [][32]byte) (*StorecoreStoreSpliceStaticDataIterator, error) {

	var tableIdRule []interface{}
	for _, tableIdItem := range tableId {
		tableIdRule = append(tableIdRule, tableIdItem)
	}

	logs, sub, err := _Storecore.contract.FilterLogs(opts, "Store_SpliceStaticData", tableIdRule)
	if err != nil {
		return nil, err
	}
	return &StorecoreStoreSpliceStaticDataIterator{contract: _Storecore.contract, event: "Store_SpliceStaticData", logs: logs, sub: sub}, nil
}

// WatchStoreSpliceStaticData is a free log subscription operation binding the contract event 0x8c0b5119d4cec7b284c6b1b39252a03d1e2f2d7451a5895562524c113bb952be.
//
// Solidity: event Store_SpliceStaticData(bytes32 indexed tableId, bytes32[] keyTuple, uint48 start, bytes data)
func (_Storecore *StorecoreFilterer) WatchStoreSpliceStaticData(opts *bind.WatchOpts, sink chan<- *StorecoreStoreSpliceStaticData, tableId [][32]byte) (event.Subscription, error) {

	var tableIdRule []interface{}
	for _, tableIdItem := range tableId {
		tableIdRule = append(tableIdRule, tableIdItem)
	}

	logs, sub, err := _Storecore.contract.WatchLogs(opts, "Store_SpliceStaticData", tableIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorecoreStoreSpliceStaticData)
				if err := _Storecore.contract.UnpackLog(event, "Store_SpliceStaticData", log); err != nil {
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

// ParseStoreSpliceStaticData is a log parse operation binding the contract event 0x8c0b5119d4cec7b284c6b1b39252a03d1e2f2d7451a5895562524c113bb952be.
//
// Solidity: event Store_SpliceStaticData(bytes32 indexed tableId, bytes32[] keyTuple, uint48 start, bytes data)
func (_Storecore *StorecoreFilterer) ParseStoreSpliceStaticData(log types.Log) (*StorecoreStoreSpliceStaticData, error) {
	event := new(StorecoreStoreSpliceStaticData)
	if err := _Storecore.contract.UnpackLog(event, "Store_SpliceStaticData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
