// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CBPeripheral] class.
var (
	CBPeripheralClass     _CBPeripheralClass
	CBPeripheralClassOnce sync.Once
)

func getCBPeripheralClass() _CBPeripheralClass {
	CBPeripheralClassOnce.Do(func() {
		CBPeripheralClass = _CBPeripheralClass{objc.GetClass("CBPeripheral")}
	})
	return CBPeripheralClass
}

type _CBPeripheralClass struct {
	class objc.Class
}

// An interface definition for the [CBPeripheral] class.
type ICBPeripheral interface {
	ICBPeer
	// properties:
	CanSendWriteWithoutResponse() bool
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	Name() objc.IObject /* cross-framework: NSString */
	RSSI() objc.IObject /* cross-framework: NSNumber */
	Services() []ICBService
	State() CBPeripheralState
	// methods:
	DiscoverCharacteristicsForService(characteristicUUIDs []ICBUUID, service ICBService)
	DiscoverDescriptorsForCharacteristic(characteristic ICBCharacteristic)
	DiscoverIncludedServicesForService(includedServiceUUIDs []ICBUUID, service ICBService)
	DiscoverServices(serviceUUIDs []ICBUUID)
	MaximumWriteValueLengthForType(type_ CBCharacteristicWriteType) uint
	OpenL2CAPChannel(PSM CBL2CAPPSM /* typedef */)
	ReadRSSI()
	ReadValueForCharacteristic(characteristic ICBCharacteristic)
	ReadValueForDescriptor(descriptor ICBDescriptor)
	SetNotifyValueForCharacteristic(enabled bool, characteristic ICBCharacteristic)
	WriteValueForDescriptor(data objc.IObject /* cross-framework: NSData */, descriptor ICBDescriptor)
	WriteValueForCharacteristicType(data objc.IObject /* cross-framework: NSData */, characteristic ICBCharacteristic, type_ CBCharacteristicWriteType)
}

// A remote peripheral device.
//
// The class represents remote peripheral devices that your app discovers with a central manager (an instance of ). Peripherals use universally unique identifiers (UUIDs), represented by objects, to identify themselves. Peripherals may contain one or more services or provide useful information about their connected signal strength. You use this class to discover, explore, and interact with the services available on a remote peripheral that supports Bluetooth low energy. A service encapsulates the way part of the device behaves. For example, one service of a heart rate monitor may be to expose heart rate data from a sensor. Services themselves contain of characteristics or included services (references to other services). Characteristics provide further details about a peripheral’s service. For example, the heart rate service may contain multiple characteristics. One characteristic could describe the intended body location of the device’s heart rate sensor, and another characteristic could transmit the heart rate measurement data. Finally, characteristics contain any number of descriptors that provide more information about the characteristic’s value, such as a human-readable description and a way to format the value.


// A remote peripheral device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheral
type CBPeripheral struct {
	CBPeer
}

// CBPeripheralFrom constructs a [CBPeripheral] from an unsafe.Pointer.
//
// A remote peripheral device.
func CBPeripheralFrom(ptr unsafe.Pointer) CBPeripheral {
	return CBPeripheral{
		CBPeer: CBPeerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CBPeripheralClass) Alloc() CBPeripheral {
	rv := objc.Send[CBPeripheral](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CBPeripheralClass) New() CBPeripheral {
	rv := objc.Send[CBPeripheral](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CBPeripheral) Init() CBPeripheral {
	rv := objc.Send[CBPeripheral](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CBPeripheral) Autorelease() CBPeripheral {
	rv := objc.Send[CBPeripheral](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCBPeripheral creates a new CBPeripheral instance.
func NewCBPeripheral() CBPeripheral {
	return getCBPeripheralClass().New()
}



// Discovers the specified characteristics of a service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheral/discoverCharacteristics(_:for:)
func (c_ CBPeripheral) DiscoverCharacteristicsForService(characteristicUUIDs []ICBUUID, service ICBService) {
	objc.Send[objc.ID](c_.ID, objc.Sel("discoverCharacteristics:forService:"), characteristicUUIDs, service)
}


// Discovers the descriptors of a characteristic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheral/discoverDescriptors(for:)
func (c_ CBPeripheral) DiscoverDescriptorsForCharacteristic(characteristic ICBCharacteristic) {
	objc.Send[objc.ID](c_.ID, objc.Sel("discoverDescriptorsForCharacteristic:"), characteristic)
}


// Discovers the specified included services of a previously-discovered service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheral/discoverIncludedServices(_:for:)
func (c_ CBPeripheral) DiscoverIncludedServicesForService(includedServiceUUIDs []ICBUUID, service ICBService) {
	objc.Send[objc.ID](c_.ID, objc.Sel("discoverIncludedServices:forService:"), includedServiceUUIDs, service)
}


// Discovers the specified services of the peripheral.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheral/discoverServices(_:)
func (c_ CBPeripheral) DiscoverServices(serviceUUIDs []ICBUUID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("discoverServices:"), serviceUUIDs)
}


// The maximum amount of data, in bytes, you can send to a characteristic in a single write type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheral/maximumWriteValueLength(for:)
func (c_ CBPeripheral) MaximumWriteValueLengthForType(type_ CBCharacteristicWriteType) uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("maximumWriteValueLengthForType:"), type_)
	return rv
}


// Attempts to open an L2CAP channel to the peripheral using the supplied Protocol/Service Multiplexer (PSM).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheral/openL2CAPChannel(_:)
func (c_ CBPeripheral) OpenL2CAPChannel(PSM CBL2CAPPSM /* typedef */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("openL2CAPChannel:"), PSM)
}


// Retrieves the current RSSI value for the peripheral while connected to the central manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheral/readRSSI()
func (c_ CBPeripheral) ReadRSSI() {
	objc.Send[objc.ID](c_.ID, objc.Sel("readRSSI"))
}


// Retrieves the value of a specified characteristic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheral/readValue(for:)-6u2kr
func (c_ CBPeripheral) ReadValueForCharacteristic(characteristic ICBCharacteristic) {
	objc.Send[objc.ID](c_.ID, objc.Sel("readValueForCharacteristic:"), characteristic)
}


// Retrieves the value of a specified characteristic descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheral/readValue(for:)-91hhp
func (c_ CBPeripheral) ReadValueForDescriptor(descriptor ICBDescriptor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("readValueForDescriptor:"), descriptor)
}


// Sets notifications or indications for the value of a specified characteristic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheral/setNotifyValue(_:for:)
func (c_ CBPeripheral) SetNotifyValueForCharacteristic(enabled bool, characteristic ICBCharacteristic) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNotifyValue:forCharacteristic:"), enabled, characteristic)
}


// Writes the value of a characteristic descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheral/writeValue(_:for:)
func (c_ CBPeripheral) WriteValueForDescriptor(data objc.IObject /* cross-framework: NSData */, descriptor ICBDescriptor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("writeValue:forDescriptor:"), data, descriptor)
}


// Writes the value of a characteristic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheral/writeValue(_:for:type:)
func (c_ CBPeripheral) WriteValueForCharacteristicType(data objc.IObject /* cross-framework: NSData */, characteristic ICBCharacteristic, type_ CBCharacteristicWriteType) {
	objc.Send[objc.ID](c_.ID, objc.Sel("writeValue:forCharacteristic:type:"), data, characteristic, type_)
}


// A Boolean value that indicates whether the remote device can send a write without a response.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheral/canSendWriteWithoutResponse
func (c_ CBPeripheral) CanSendWriteWithoutResponse() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("canSendWriteWithoutResponse"))
	return rv
}


// The delegate object specified to receive peripheral events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheral/delegate
func (c_ CBPeripheral) Delegate() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("delegate"))
	return rv
}


// The delegate object specified to receive peripheral events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheral/delegate
func (c_ CBPeripheral) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegate:"), value)
}


// The name of the peripheral.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheral/name
func (c_ CBPeripheral) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("name"))
	return rv
}


// The Received Signal Strength Indicator (RSSI), in decibels, of the peripheral.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheral/rssi
func (c_ CBPeripheral) RSSI() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("RSSI"))
	return rv
}


// A list of a peripheral’s discovered services.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheral/services
func (c_ CBPeripheral) Services() []ICBService {
	rv := objc.Send[[]CBService](c_.ID, objc.Sel("services"))
	return rv
}


// The connection state of the peripheral.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheral/state
func (c_ CBPeripheral) State() CBPeripheralState {
	rv := objc.Send[CBPeripheralState](c_.ID, objc.Sel("state"))
	return rv
}


