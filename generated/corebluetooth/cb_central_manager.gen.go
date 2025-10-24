// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CBCentralManager] class.
var (
	CBCentralManagerClass     _CBCentralManagerClass
	CBCentralManagerClassOnce sync.Once
)

func getCBCentralManagerClass() _CBCentralManagerClass {
	CBCentralManagerClassOnce.Do(func() {
		CBCentralManagerClass = _CBCentralManagerClass{objc.GetClass("CBCentralManager")}
	})
	return CBCentralManagerClass
}

type _CBCentralManagerClass struct {
	class objc.Class
}

// An interface definition for the [CBCentralManager] class.
type ICBCentralManager interface {
	ICBManager
	// properties:
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	IsScanning() bool
	// methods:
	CancelPeripheralConnection(peripheral ICBPeripheral)
	ConnectPeripheralOptions(peripheral ICBPeripheral, options foundation.IDictionary)
	RetrieveConnectedPeripheralsWithServices(serviceUUIDs []ICBUUID) []ICBPeripheral
	RetrievePeripheralsWithIdentifiers(identifiers []objc.IObject /* cross-framework: UUID */) []ICBPeripheral
	ScanForPeripheralsWithServicesOptions(serviceUUIDs []ICBUUID, options foundation.IDictionary)
	StopScan()
}

// An object that scans for, discovers, connects to, and manages peripherals.
//
// objects manage discovered or connected remote peripheral devices (represented by objects), including scanning for, discovering, and connecting to advertising peripherals. Before calling the methods, set the state of the central manager object to powered on, as indicated by the constant. This state indicates that the central device (your iPhone or iPad, for instance) supports Bluetooth low energy and that Bluetooth is on and available for use.


// An object that scans for, discovers, connects to, and manages peripherals.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManager
type CBCentralManager struct {
	CBManager
}

// CBCentralManagerFrom constructs a [CBCentralManager] from an unsafe.Pointer.
//
// An object that scans for, discovers, connects to, and manages peripherals.
func CBCentralManagerFrom(ptr unsafe.Pointer) CBCentralManager {
	return CBCentralManager{
		CBManager: CBManagerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CBCentralManagerClass) Alloc() CBCentralManager {
	rv := objc.Send[CBCentralManager](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CBCentralManagerClass) New() CBCentralManager {
	rv := objc.Send[CBCentralManager](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CBCentralManager) Init() CBCentralManager {
	rv := objc.Send[CBCentralManager](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CBCentralManager) Autorelease() CBCentralManager {
	rv := objc.Send[CBCentralManager](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCBCentralManager creates a new CBCentralManager instance.
func NewCBCentralManager() CBCentralManager {
	return getCBCentralManagerClass().New()
}



// Initializes the central manager with a specified delegate and dispatch queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManager/init(delegate:queue:)
func NewCBCentralManagerWithDelegateQueue(delegate objectivec.IObject, queue unsafe.Pointer) CBCentralManager {
	instance := getCBCentralManagerClass().Alloc()
	rv := objc.Send[CBCentralManager](instance.ID, objc.Sel("initWithDelegate:queue:"), delegate, queue)
	rv.Autorelease()
	return rv
}


// Initializes the central manager with specified delegate, dispatch queue, and initialization options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManager/init(delegate:queue:options:)
func NewCBCentralManagerWithDelegateQueueOptions(delegate objectivec.IObject, queue unsafe.Pointer, options foundation.IDictionary) CBCentralManager {
	instance := getCBCentralManagerClass().Alloc()
	rv := objc.Send[CBCentralManager](instance.ID, objc.Sel("initWithDelegate:queue:options:"), delegate, queue, options)
	rv.Autorelease()
	return rv
}



// Returns a Boolean that indicates whether the device supports a specific set of features.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManager/supports(_:)
func (cc _CBCentralManagerClass) SupportsFeatures(features CBCentralManagerFeature) bool {
	rv := objc.Send[bool](objc.ID(cc.class), objc.Sel("supportsFeatures:"), features)
	return rv
}


// Cancels an active or pending local connection to a peripheral.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManager/cancelPeripheralConnection(_:)
func (c_ CBCentralManager) CancelPeripheralConnection(peripheral ICBPeripheral) {
	objc.Send[objc.ID](c_.ID, objc.Sel("cancelPeripheralConnection:"), peripheral)
}


// Establishes a local connection to a peripheral.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManager/connect(_:options:)
func (c_ CBCentralManager) ConnectPeripheralOptions(peripheral ICBPeripheral, options foundation.IDictionary) {
	objc.Send[objc.ID](c_.ID, objc.Sel("connectPeripheral:options:"), peripheral, options)
}


// Returns a list of the peripherals connected to the system whose services match a given set of criteria.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManager/retrieveConnectedPeripherals(withServices:)
func (c_ CBCentralManager) RetrieveConnectedPeripheralsWithServices(serviceUUIDs []ICBUUID) []ICBPeripheral {
	rv := objc.Send[[]CBPeripheral](c_.ID, objc.Sel("retrieveConnectedPeripheralsWithServices:"), serviceUUIDs)
	return rv
}


// Returns a list of known peripherals by their identifiers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManager/retrievePeripherals(withIdentifiers:)
func (c_ CBCentralManager) RetrievePeripheralsWithIdentifiers(identifiers []objc.IObject /* cross-framework: UUID */) []ICBPeripheral {
	rv := objc.Send[[]CBPeripheral](c_.ID, objc.Sel("retrievePeripheralsWithIdentifiers:"), identifiers)
	return rv
}


// Scans for peripherals that are advertising services.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManager/scanForPeripherals(withServices:options:)
func (c_ CBCentralManager) ScanForPeripheralsWithServicesOptions(serviceUUIDs []ICBUUID, options foundation.IDictionary) {
	objc.Send[objc.ID](c_.ID, objc.Sel("scanForPeripheralsWithServices:options:"), serviceUUIDs, options)
}


// Asks the central manager to stop scanning for peripherals.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManager/stopScan()
func (c_ CBCentralManager) StopScan() {
	objc.Send[objc.ID](c_.ID, objc.Sel("stopScan"))
}


// The delegate object that you want to receive central manager events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManager/delegate
func (c_ CBCentralManager) Delegate() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("delegate"))
	return rv
}


// The delegate object that you want to receive central manager events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManager/delegate
func (c_ CBCentralManager) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegate:"), value)
}


// A Boolean value that indicates whether the central is currently scanning.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManager/isScanning
func (c_ CBCentralManager) IsScanning() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isScanning"))
	return rv
}


