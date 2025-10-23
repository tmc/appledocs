// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	LocalIdentifier() string /* primitive/slice/pointer. */
	FirmwareVersion() string /* primitive/slice/pointer. */
	SetFirmwareVersion(value string /* primitive/slice/pointer. */)
	HardwareVersion() string /* primitive/slice/pointer. */
	SetHardwareVersion(value string /* primitive/slice/pointer. */)
	Manufacturer() string /* primitive/slice/pointer. */
	SetManufacturer(value string /* primitive/slice/pointer. */)
	Model() string /* primitive/slice/pointer. */
	SetModel(value string /* primitive/slice/pointer. */)
	Name() string /* primitive/slice/pointer. */
	SetName(value string /* primitive/slice/pointer. */)
	SoftwareVersion() string /* primitive/slice/pointer. */
	SetSoftwareVersion(value string /* primitive/slice/pointer. */)
	UdiDeviceIdentifier() string /* primitive/slice/pointer. */
	SetUdiDeviceIdentifier(value string /* primitive/slice/pointer. */)
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
func (h_ HKDevice) LocalIdentifier() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](h_.ID, objc.Sel("localIdentifier"))
	return rv
}


// An arbitrary string representing the current version of the firmware running on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdevice/firmwareversion
func (h_ HKDevice) FirmwareVersion() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](h_.ID, objc.Sel("firmwareVersion"))
	return rv
}


// An arbitrary string representing the current version of the firmware running on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdevice/firmwareversion
func (h_ HKDevice) SetFirmwareVersion(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setFirmwareVersion:"), objc.String(value))
}


// An arbitrary string representing the hardware version of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdevice/hardwareversion
func (h_ HKDevice) HardwareVersion() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](h_.ID, objc.Sel("hardwareVersion"))
	return rv
}


// An arbitrary string representing the hardware version of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdevice/hardwareversion
func (h_ HKDevice) SetHardwareVersion(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setHardwareVersion:"), objc.String(value))
}


// A string representing the device’s manufacturer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdevice/manufacturer
func (h_ HKDevice) Manufacturer() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](h_.ID, objc.Sel("manufacturer"))
	return rv
}


// A string representing the device’s manufacturer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdevice/manufacturer
func (h_ HKDevice) SetManufacturer(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setManufacturer:"), objc.String(value))
}


// A string representing the device’s model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdevice/model
func (h_ HKDevice) Model() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](h_.ID, objc.Sel("model"))
	return rv
}


// A string representing the device’s model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdevice/model
func (h_ HKDevice) SetModel(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setModel:"), objc.String(value))
}


// The user-facing name for the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdevice/name
func (h_ HKDevice) Name() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](h_.ID, objc.Sel("name"))
	return rv
}


// The user-facing name for the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdevice/name
func (h_ HKDevice) SetName(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setName:"), objc.String(value))
}


// An arbitrary string representing the version of the software running on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdevice/softwareversion
func (h_ HKDevice) SoftwareVersion() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](h_.ID, objc.Sel("softwareVersion"))
	return rv
}


// An arbitrary string representing the version of the software running on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdevice/softwareversion
func (h_ HKDevice) SetSoftwareVersion(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setSoftwareVersion:"), objc.String(value))
}


// The device identifier portion of the US Food and Drug Administration’s Unique Device Identifier (UDI).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdevice/udideviceidentifier
func (h_ HKDevice) UdiDeviceIdentifier() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](h_.ID, objc.Sel("udiDeviceIdentifier"))
	return rv
}


// The device identifier portion of the US Food and Drug Administration’s Unique Device Identifier (UDI).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdevice/udideviceidentifier
func (h_ HKDevice) SetUdiDeviceIdentifier(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setUdiDeviceIdentifier:"), objc.String(value))
}



