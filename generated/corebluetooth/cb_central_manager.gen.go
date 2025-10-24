// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CBCentralManager */


/* debug [class_header]: Header for CBCentralManager */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CBCentralManager */
// An interface definition for the [CBCentralManager] class.
type ICBCentralManager interface {
	ICBManager
	
/* debug [class_interface_properties]: Properties for CBCentralManager */
	// properties:
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	IsScanning() bool
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CBCentralManager */
	// methods:
	CancelPeripheralConnection(peripheral ICBPeripheral)
	ConnectPeripheralOptions(peripheral ICBPeripheral, options foundation.IDictionary)
	RetrieveConnectedPeripheralsWithServices(serviceUUIDs []CBUUID) []CBPeripheral
	RetrievePeripheralsWithIdentifiers(identifiers []foundation.UUID) []CBPeripheral
	ScanForPeripheralsWithServicesOptions(serviceUUIDs []CBUUID, options foundation.IDictionary)
	StopScan()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CBCentralManager */
// Alloc allocates a new instance without initialization.
func (cc _CBCentralManagerClass) Alloc() CBCentralManager {
	rv := objc.Send[CBCentralManager](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CBCentralManager */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CBCentralManager */

// Initializes the central manager with a specified delegate and dispatch queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManager/init(delegate:queue:)
func NewCBCentralManagerWithDelegateQueue(delegate unsafe.Pointer, queue objectivec.IObject) CBCentralManager {
	instance := getCBCentralManagerClass().Alloc()
	rv := objc.Send[CBCentralManager](instance.ID, objc.Sel("initWithDelegate:queue:"), delegate, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCBCentralManagerWithDelegateQueue */


// Initializes the central manager with specified delegate, dispatch queue, and initialization options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManager/init(delegate:queue:options:)
func NewCBCentralManagerWithDelegateQueueOptions(delegate unsafe.Pointer, queue objectivec.IObject, options foundation.IDictionary) CBCentralManager {
	instance := getCBCentralManagerClass().Alloc()
	rv := objc.Send[CBCentralManager](instance.ID, objc.Sel("initWithDelegate:queue:options:"), delegate, queue, options)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCBCentralManagerWithDelegateQueueOptions */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CBCentralManager */

// Returns a Boolean that indicates whether the device supports a specific set of features.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManager/supports(_:)
func (cc _CBCentralManagerClass) SupportsFeatures(features CBCentralManagerFeature) bool {
	rv := objc.Send[bool](objc.ID(cc.class), objc.Sel("supportsFeatures:"), features)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SupportsFeatures) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CBCentralManager */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CBCentralManager */

// Cancels an active or pending local connection to a peripheral.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManager/cancelPeripheralConnection(_:)
func (c_ CBCentralManager) CancelPeripheralConnection(peripheral ICBPeripheral) {
	objc.Send[objc.ID](c_.ID, objc.Sel("cancelPeripheralConnection:"), peripheral)
}/* debug [instance_methods/method]: CancelPeripheralConnection */


// Establishes a local connection to a peripheral.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManager/connect(_:options:)
func (c_ CBCentralManager) ConnectPeripheralOptions(peripheral ICBPeripheral, options foundation.IDictionary) {
	objc.Send[objc.ID](c_.ID, objc.Sel("connectPeripheral:options:"), peripheral, options)
}/* debug [instance_methods/method]: ConnectPeripheralOptions */


// Returns a list of the peripherals connected to the system whose services match a given set of criteria.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManager/retrieveConnectedPeripherals(withServices:)
func (c_ CBCentralManager) RetrieveConnectedPeripheralsWithServices(serviceUUIDs []CBUUID) []CBPeripheral {
	rv := objc.Send[[]CBPeripheral](c_.ID, objc.Sel("retrieveConnectedPeripheralsWithServices:"), serviceUUIDs)
	return rv
}/* debug [instance_methods/method]: RetrieveConnectedPeripheralsWithServices */


// Returns a list of known peripherals by their identifiers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManager/retrievePeripherals(withIdentifiers:)
func (c_ CBCentralManager) RetrievePeripheralsWithIdentifiers(identifiers []foundation.UUID) []CBPeripheral {
	rv := objc.Send[[]CBPeripheral](c_.ID, objc.Sel("retrievePeripheralsWithIdentifiers:"), identifiers)
	return rv
}/* debug [instance_methods/method]: RetrievePeripheralsWithIdentifiers */


// Scans for peripherals that are advertising services.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManager/scanForPeripherals(withServices:options:)
func (c_ CBCentralManager) ScanForPeripheralsWithServicesOptions(serviceUUIDs []CBUUID, options foundation.IDictionary) {
	objc.Send[objc.ID](c_.ID, objc.Sel("scanForPeripheralsWithServices:options:"), serviceUUIDs, options)
}/* debug [instance_methods/method]: ScanForPeripheralsWithServicesOptions */


// Asks the central manager to stop scanning for peripherals.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManager/stopScan()
func (c_ CBCentralManager) StopScan() {
	objc.Send[objc.ID](c_.ID, objc.Sel("stopScan"))
}/* debug [instance_methods/method]: StopScan */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CBCentralManager */

// The delegate object that you want to receive central manager events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManager/delegate
func (c_ CBCentralManager) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The delegate object that you want to receive central manager events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManager/delegate
func (c_ CBCentralManager) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// A Boolean value that indicates whether the central is currently scanning.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManager/isScanning
func (c_ CBCentralManager) IsScanning() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isScanning"))
	return rv
}/* debug [instance_properties/getter]: isScanning */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CBCentralManager */


