// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [CBPeripheralManager] class.
type ICBPeripheralManager interface {
	ICBManager
	AddService(service ICBMutableService)
	PublishL2CAPChannelWithEncryption(encryptionRequired bool)
	RemoveService(service ICBMutableService)
	RemoveAllServices()
	RespondToRequestWithResult(request ICBATTRequest, result ICBATTError)
	SetDesiredConnectionLatencyForCentral(latency ICBPeripheralManagerConnectionLatency, central ICBCentral)
	StartAdvertising(advertisementData unsafe.Pointer)
	StopAdvertising()
	UnpublishL2CAPChannel(PSM ICBL2CAPPSM)
	UpdateValueForCharacteristicOnSubscribedCentrals(value foundation.IData, characteristic ICBMutableCharacteristic, centrals []CBCentral) bool
}

// An object that manages and advertises peripheral services exposed by this app.
//
// Core Bluetooth uses objects to manage published services within the local peripheral’s Generic Attribute Profile (GATT) database and to advertise these services to central devices (represented by objects). While a service is in the database, any connected central can see and connect to it. That said, if your app hasn’t specified the background mode, the contents of its services become disabled when it’s in the background or in a suspended state. In this scenario, any remote central trying to access the service’s characteristic value or characteristic descriptors receives an error. Before you call methods, the peripheral manager object must be in the powered-on state, as indicated by the . This state indicates that the device (your iPhone or iPad, for instance) supports Bluetooth low energy and that its Bluetooth is on and available for use. In watchOS, tvOS, and visionOS, you can’t advertise services using a object because support for doing so is unavailable.
//
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

// Alloc allocates a new instance without initialization.
func (cc _CBPeripheralManagerClass) Alloc() CBPeripheralManager {
	rv := objc.Send[CBPeripheralManager](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Initializes the peripheral manager with a specified delegate and dispatch queue.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/init(delegate:queue:)
func NewCBPeripheralManagerWithDelegateQueue(delegate objectivec.IObject, queue unsafe.Pointer) CBPeripheralManager {
	instance := getCBPeripheralManagerClass().Alloc()
	rv := objc.Send[CBPeripheralManager](instance.ID, objc.Sel("initWithDelegate:queue:"), delegate, queue)
	rv.Autorelease()
	return rv
}



// Initializes the peripheral manager with a specified delegate, dispatch queue, and initialization options.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/init(delegate:queue:options:)
func NewCBPeripheralManagerWithDelegateQueueOptions(delegate objectivec.IObject, queue unsafe.Pointer, options unsafe.Pointer) CBPeripheralManager {
	instance := getCBPeripheralManagerClass().Alloc()
	rv := objc.Send[CBPeripheralManager](instance.ID, objc.Sel("initWithDelegate:queue:options:"), delegate, queue, options)
	rv.Autorelease()
	return rv
}


// Returns the app’s authorization status for sharing data while in the background.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/authorizationStatus()
func (cc _CBPeripheralManagerClass) AuthorizationStatus() CBPeripheralManagerAuthorizationStatus {
	rv := objc.Send[CBPeripheralManagerAuthorizationStatus](objc.ID(cc.class), objc.Sel("authorizationStatus"))
	return rv
}

// Publishes a service and any of its associated characteristics and characteristic descriptors to the local GATT database.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/add(_:)
func (c_ CBPeripheralManager) AddService(service ICBMutableService) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addService:"), service)
}

// Creates a listener for incoming L2CAP channel connections.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/publishL2CAPChannel(withEncryption:)
func (c_ CBPeripheralManager) PublishL2CAPChannelWithEncryption(encryptionRequired bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("publishL2CAPChannelWithEncryption:"), encryptionRequired)
}

// Removes a specified published service from the local GATT database.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/remove(_:)
func (c_ CBPeripheralManager) RemoveService(service ICBMutableService) {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeService:"), service)
}

// Removes all published services from the local GATT database.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/removeAllServices()
func (c_ CBPeripheralManager) RemoveAllServices() {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeAllServices"))
}

// Responds to a read or write request from a connected central.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/respond(to:withResult:)
func (c_ CBPeripheralManager) RespondToRequestWithResult(request ICBATTRequest, result ICBATTError) {
	objc.Send[objc.ID](c_.ID, objc.Sel("respondToRequest:withResult:"), request, result)
}

// Sets the desired connection latency for an existing connection to a central device.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/setDesiredConnectionLatency(_:for:)
func (c_ CBPeripheralManager) SetDesiredConnectionLatencyForCentral(latency ICBPeripheralManagerConnectionLatency, central ICBCentral) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDesiredConnectionLatency:forCentral:"), latency, central)
}

// Advertises peripheral manager data.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/startAdvertising(_:)
func (c_ CBPeripheralManager) StartAdvertising(advertisementData unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("startAdvertising:"), advertisementData)
}

// Stops advertising peripheral manager data.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/stopAdvertising()
func (c_ CBPeripheralManager) StopAdvertising() {
	objc.Send[objc.ID](c_.ID, objc.Sel("stopAdvertising"))
}

// Removes a published service from the local system.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/unpublishL2CAPChannel(_:)
func (c_ CBPeripheralManager) UnpublishL2CAPChannel(PSM ICBL2CAPPSM) {
	objc.Send[objc.ID](c_.ID, objc.Sel("unpublishL2CAPChannel:"), PSM)
}

// Send an updated characteristic value to one or more subscribed centrals, using a notification or indication.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/updateValue(_:for:onSubscribedCentrals:)
func (c_ CBPeripheralManager) UpdateValueForCharacteristicOnSubscribedCentrals(value foundation.IData, characteristic ICBMutableCharacteristic, centrals []CBCentral) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("updateValue:forCharacteristic:onSubscribedCentrals:"), value, characteristic, centrals)
	return rv
}

// The delegate object specified to receive peripheral events.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/delegate
func (c_ CBPeripheralManager) Delegate() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The delegate object specified to receive peripheral events.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/delegate
func (c_ CBPeripheralManager) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegate:"), value)
}

// A Boolean value that indicates whether the peripheral is advertising data.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManager/isAdvertising
func (c_ CBPeripheralManager) IsAdvertising() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isAdvertising"))
	return rv
}


