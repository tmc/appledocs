// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CBPeripheralManager */


/* debug [class_header]: Header for CBPeripheralManager */
// The class instance for the [CBPeripheralManager] class.
var (
	CBPeripheralManagerClass     _CBPeripheralManagerClass
	CBPeripheralManagerClassOnce sync.Once
)

func getCBPeripheralManagerClass() _CBPeripheralManagerClass {
	CBPeripheralManagerClassOnce.Do(func() {
		CBPeripheralManagerClass = _CBPeripheralManagerClass{objc.GetClass("CBPeripheralManager")}
	})
	return CBPeripheralManagerClass
}

type _CBPeripheralManagerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CBPeripheralManager */
// An interface definition for the [CBPeripheralManager] class.
type ICBPeripheralManager interface {
	ICBManager
	
/* debug [class_interface_properties]: Properties for CBPeripheralManager */
	// properties:
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	IsAdvertising() bool
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CBPeripheralManager */
	// methods:
	AddService(service ICBMutableService)
	PublishL2CAPChannelWithEncryption(encryptionRequired bool)
	RemoveService(service ICBMutableService)
	RemoveAllServices()
	RespondToRequestWithResult(request ICBATTRequest, result CBATTError)
	SetDesiredConnectionLatencyForCentral(latency CBPeripheralManagerConnectionLatency, central ICBCentral)
	StartAdvertising(advertisementData foundation.IDictionary)
	StopAdvertising()
	UnpublishL2CAPChannel(PSM CBL2CAPPSM /* typedef */)
	UpdateValueForCharacteristicOnSubscribedCentrals(value objc.IObject /* cross-framework: NSData */, characteristic ICBMutableCharacteristic, centrals []CBCentral) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CBPeripheralManager */
// Alloc allocates a new instance without initialization.
func (cc _CBPeripheralManagerClass) Alloc() CBPeripheralManager {
	rv := objc.Send[CBPeripheralManager](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CBPeripheralManagerClass) New() CBPeripheralManager {
	rv := objc.Send[CBPeripheralManager](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CBPeripheralManager) Init() CBPeripheralManager {
	rv := objc.Send[CBPeripheralManager](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CBPeripheralManager) Autorelease() CBPeripheralManager {
	rv := objc.Send[CBPeripheralManager](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCBPeripheralManager creates a new CBPeripheralManager instance.
func NewCBPeripheralManager() CBPeripheralManager {
	return getCBPeripheralManagerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CBPeripheralManager */
// An object that manages and advertises peripheral services exposed by this app.
//
// Core Bluetooth uses objects to manage published services within the local peripheral’s Generic Attribute Profile (GATT) database and to advertise these services to central devices (represented by objects). While a service is in the database, any connected central can see and connect to it. That said, if your app hasn’t specified the background mode, the contents of its services become disabled when it’s in the background or in a suspended state. In this scenario, any remote central trying to access the service’s characteristic value or characteristic descriptors receives an error. Before you call methods, the peripheral manager object must be in the powered-on state, as indicated by the . This state indicates that the device (your iPhone or iPad, for instance) supports Bluetooth low energy and that its Bluetooth is on and available for use. In watchOS, tvOS, and visionOS, you can’t advertise services using a object because support for doing so is unavailable.


// An object that manages and advertises peripheral services exposed by this app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager
type CBPeripheralManager struct {
	CBManager
}

// CBPeripheralManagerFrom constructs a [CBPeripheralManager] from an unsafe.Pointer.
//
// An object that manages and advertises peripheral services exposed by this app.
func CBPeripheralManagerFrom(ptr unsafe.Pointer) CBPeripheralManager {
	return CBPeripheralManager{
		CBManager: CBManagerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CBPeripheralManager */

// Initializes the peripheral manager with a specified delegate and dispatch queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/init(delegate:queue:)
func NewCBPeripheralManagerWithDelegateQueue(delegate unsafe.Pointer, queue objectivec.IObject) CBPeripheralManager {
	instance := getCBPeripheralManagerClass().Alloc()
	rv := objc.Send[CBPeripheralManager](instance.ID, objc.Sel("initWithDelegate:queue:"), delegate, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCBPeripheralManagerWithDelegateQueue */


// Initializes the peripheral manager with a specified delegate, dispatch queue, and initialization options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/init(delegate:queue:options:)
func NewCBPeripheralManagerWithDelegateQueueOptions(delegate unsafe.Pointer, queue objectivec.IObject, options foundation.IDictionary) CBPeripheralManager {
	instance := getCBPeripheralManagerClass().Alloc()
	rv := objc.Send[CBPeripheralManager](instance.ID, objc.Sel("initWithDelegate:queue:options:"), delegate, queue, options)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCBPeripheralManagerWithDelegateQueueOptions */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CBPeripheralManager */

// Returns the app’s authorization status for sharing data while in the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/authorizationStatus()
func (cc _CBPeripheralManagerClass) AuthorizationStatus() CBPeripheralManagerAuthorizationStatus {
	rv := objc.Send[CBPeripheralManagerAuthorizationStatus](objc.ID(cc.class), objc.Sel("authorizationStatus"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AuthorizationStatus) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CBPeripheralManager */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CBPeripheralManager */

// Publishes a service and any of its associated characteristics and characteristic descriptors to the local GATT database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/add(_:)
func (c_ CBPeripheralManager) AddService(service ICBMutableService) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addService:"), service)
}/* debug [instance_methods/method]: AddService */


// Creates a listener for incoming L2CAP channel connections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/publishL2CAPChannel(withEncryption:)
func (c_ CBPeripheralManager) PublishL2CAPChannelWithEncryption(encryptionRequired bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("publishL2CAPChannelWithEncryption:"), encryptionRequired)
}/* debug [instance_methods/method]: PublishL2CAPChannelWithEncryption */


// Removes a specified published service from the local GATT database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/remove(_:)
func (c_ CBPeripheralManager) RemoveService(service ICBMutableService) {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeService:"), service)
}/* debug [instance_methods/method]: RemoveService */


// Removes all published services from the local GATT database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/removeAllServices()
func (c_ CBPeripheralManager) RemoveAllServices() {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeAllServices"))
}/* debug [instance_methods/method]: RemoveAllServices */


// Responds to a read or write request from a connected central.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/respond(to:withResult:)
func (c_ CBPeripheralManager) RespondToRequestWithResult(request ICBATTRequest, result CBATTError) {
	objc.Send[objc.ID](c_.ID, objc.Sel("respondToRequest:withResult:"), request, result)
}/* debug [instance_methods/method]: RespondToRequestWithResult */


// Sets the desired connection latency for an existing connection to a central device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/setDesiredConnectionLatency(_:for:)
func (c_ CBPeripheralManager) SetDesiredConnectionLatencyForCentral(latency CBPeripheralManagerConnectionLatency, central ICBCentral) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDesiredConnectionLatency:forCentral:"), latency, central)
}/* debug [instance_methods/method]: SetDesiredConnectionLatencyForCentral */


// Advertises peripheral manager data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/startAdvertising(_:)
func (c_ CBPeripheralManager) StartAdvertising(advertisementData foundation.IDictionary) {
	objc.Send[objc.ID](c_.ID, objc.Sel("startAdvertising:"), advertisementData)
}/* debug [instance_methods/method]: StartAdvertising */


// Stops advertising peripheral manager data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/stopAdvertising()
func (c_ CBPeripheralManager) StopAdvertising() {
	objc.Send[objc.ID](c_.ID, objc.Sel("stopAdvertising"))
}/* debug [instance_methods/method]: StopAdvertising */


// Removes a published service from the local system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/unpublishL2CAPChannel(_:)
func (c_ CBPeripheralManager) UnpublishL2CAPChannel(PSM CBL2CAPPSM /* typedef */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("unpublishL2CAPChannel:"), PSM)
}/* debug [instance_methods/method]: UnpublishL2CAPChannel */


// Send an updated characteristic value to one or more subscribed centrals, using a notification or indication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/updateValue(_:for:onSubscribedCentrals:)
func (c_ CBPeripheralManager) UpdateValueForCharacteristicOnSubscribedCentrals(value objc.IObject /* cross-framework: NSData */, characteristic ICBMutableCharacteristic, centrals []CBCentral) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("updateValue:forCharacteristic:onSubscribedCentrals:"), value, characteristic, centrals)
	return rv
}/* debug [instance_methods/method]: UpdateValueForCharacteristicOnSubscribedCentrals */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CBPeripheralManager */

// The delegate object specified to receive peripheral events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/delegate
func (c_ CBPeripheralManager) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The delegate object specified to receive peripheral events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/delegate
func (c_ CBPeripheralManager) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// A Boolean value that indicates whether the peripheral is advertising data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/isAdvertising
func (c_ CBPeripheralManager) IsAdvertising() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isAdvertising"))
	return rv
}/* debug [instance_properties/getter]: isAdvertising */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CBPeripheralManager */


