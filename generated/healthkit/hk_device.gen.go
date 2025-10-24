// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HKDevice] class.
var (
	HKDeviceClass     _HKDeviceClass
	HKDeviceClassOnce sync.Once
)

func getHKDeviceClass() _HKDeviceClass {
	HKDeviceClassOnce.Do(func() {
		HKDeviceClass = _HKDeviceClass{objc.GetClass("HKDevice")}
	})
	return HKDeviceClass
}

type _HKDeviceClass struct {
	class objc.Class
}

// An interface definition for the [HKDevice] class.
type IHKDevice interface {
	objectivec.IObject
	// properties:
	LocalIdentifier() objc.IObject /* cross-framework: NSString */
	FirmwareVersion() objc.IObject /* cross-framework: NSString */
	SetFirmwareVersion(value objc.IObject /* cross-framework: NSString */)
	HardwareVersion() objc.IObject /* cross-framework: NSString */
	SetHardwareVersion(value objc.IObject /* cross-framework: NSString */)
	Manufacturer() objc.IObject /* cross-framework: NSString */
	SetManufacturer(value objc.IObject /* cross-framework: NSString */)
	Model() objc.IObject /* cross-framework: NSString */
	SetModel(value objc.IObject /* cross-framework: NSString */)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	SoftwareVersion() objc.IObject /* cross-framework: NSString */
	SetSoftwareVersion(value objc.IObject /* cross-framework: NSString */)
	UdiDeviceIdentifier() objc.IObject /* cross-framework: NSString */
	SetUdiDeviceIdentifier(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

// A device that generates data for HealthKit.
//
// Devices include Apple Watch, iPhone, and any other health or fitness peripherals that produce the sample data stored in HealthKit. Device objects are immutable: You set the device’s properties when you create the object, and they cannot change.


// A device that generates data for HealthKit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKDevice
type HKDevice struct {
	objectivec.Object
}

// HKDeviceFrom constructs a [HKDevice] from an unsafe.Pointer.
//
// A device that generates data for HealthKit.
func HKDeviceFrom(ptr unsafe.Pointer) HKDevice {
	return HKDevice{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HKDeviceClass) Alloc() HKDevice {
	rv := objc.Send[HKDevice](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKDeviceClass) New() HKDevice {
	rv := objc.Send[HKDevice](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKDevice) Init() HKDevice {
	rv := objc.Send[HKDevice](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKDevice) Autorelease() HKDevice {
	rv := objc.Send[HKDevice](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKDevice creates a new HKDevice instance.
func NewHKDevice() HKDevice {
	return getHKDeviceClass().New()
}



// An identifier that uniquely identifies the device object on the hardware running this code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKDevice/localIdentifier
func (h_ HKDevice) LocalIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("localIdentifier"))
	return rv
}


// An arbitrary string representing the current version of the firmware running on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdevice/firmwareversion
func (h_ HKDevice) FirmwareVersion() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("firmwareVersion"))
	return rv
}


// An arbitrary string representing the current version of the firmware running on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdevice/firmwareversion
func (h_ HKDevice) SetFirmwareVersion(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setFirmwareVersion:"), value)
}


// An arbitrary string representing the hardware version of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdevice/hardwareversion
func (h_ HKDevice) HardwareVersion() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("hardwareVersion"))
	return rv
}


// An arbitrary string representing the hardware version of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdevice/hardwareversion
func (h_ HKDevice) SetHardwareVersion(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setHardwareVersion:"), value)
}


// A string representing the device’s manufacturer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdevice/manufacturer
func (h_ HKDevice) Manufacturer() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("manufacturer"))
	return rv
}


// A string representing the device’s manufacturer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdevice/manufacturer
func (h_ HKDevice) SetManufacturer(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setManufacturer:"), value)
}


// A string representing the device’s model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdevice/model
func (h_ HKDevice) Model() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("model"))
	return rv
}


// A string representing the device’s model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdevice/model
func (h_ HKDevice) SetModel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setModel:"), value)
}


// The user-facing name for the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdevice/name
func (h_ HKDevice) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("name"))
	return rv
}


// The user-facing name for the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdevice/name
func (h_ HKDevice) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setName:"), value)
}


// An arbitrary string representing the version of the software running on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdevice/softwareversion
func (h_ HKDevice) SoftwareVersion() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("softwareVersion"))
	return rv
}


// An arbitrary string representing the version of the software running on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdevice/softwareversion
func (h_ HKDevice) SetSoftwareVersion(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setSoftwareVersion:"), value)
}


// The device identifier portion of the US Food and Drug Administration’s Unique Device Identifier (UDI).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdevice/udideviceidentifier
func (h_ HKDevice) UdiDeviceIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("udiDeviceIdentifier"))
	return rv
}


// The device identifier portion of the US Food and Drug Administration’s Unique Device Identifier (UDI).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdevice/udideviceidentifier
func (h_ HKDevice) SetUdiDeviceIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setUdiDeviceIdentifier:"), value)
}



