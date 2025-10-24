// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKDevice */


/* debug [class_header]: Header for HKDevice */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKDevice */
// An interface definition for the [HKDevice] class.
type IHKDevice interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HKDevice */
	// properties:
	FirmwareVersion() objc.IObject /* cross-framework: NSString */
	HardwareVersion() objc.IObject /* cross-framework: NSString */
	LocalIdentifier() objc.IObject /* cross-framework: NSString */
	Manufacturer() objc.IObject /* cross-framework: NSString */
	Model() objc.IObject /* cross-framework: NSString */
	Name() objc.IObject /* cross-framework: NSString */
	SoftwareVersion() objc.IObject /* cross-framework: NSString */
	UDIDeviceIdentifier() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKDevice */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKDevice */
// Alloc allocates a new instance without initialization.
func (hc _HKDeviceClass) Alloc() HKDevice {
	rv := objc.Send[HKDevice](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKDevice */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKDevice */

// Initializes a new device object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKDevice/init(name:manufacturer:model:hardwareVersion:firmwareVersion:softwareVersion:localIdentifier:udiDeviceIdentifier:)
func NewHKDeviceWithNameManufacturerModelHardwareVersionFirmwareVersionSoftwareVersionLocalIdentifierUDIDeviceIdentifier(name objc.IObject /* cross-framework: NSString */, manufacturer objc.IObject /* cross-framework: NSString */, model objc.IObject /* cross-framework: NSString */, hardwareVersion objc.IObject /* cross-framework: NSString */, firmwareVersion objc.IObject /* cross-framework: NSString */, softwareVersion objc.IObject /* cross-framework: NSString */, localIdentifier objc.IObject /* cross-framework: NSString */, UDIDeviceIdentifier objc.IObject /* cross-framework: NSString */) HKDevice {
	instance := getHKDeviceClass().Alloc()
	rv := objc.Send[HKDevice](instance.ID, objc.Sel("initWithName:manufacturer:model:hardwareVersion:firmwareVersion:softwareVersion:localIdentifier:UDIDeviceIdentifier:"), name, manufacturer, model, hardwareVersion, firmwareVersion, softwareVersion, localIdentifier, UDIDeviceIdentifier)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewHKDeviceWithNameManufacturerModelHardwareVersionFirmwareVersionSoftwareVersionLocalIdentifierUDIDeviceIdentifier */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKDevice */

// returns a device object that represents the current device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKDevice/local()
func (hc _HKDeviceClass) LocalDevice() HKDevice {
	rv := objc.Send[HKDevice](objc.ID(hc.class), objc.Sel("localDevice"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LocalDevice) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKDevice */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKDevice */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKDevice */

// An arbitrary string representing the current version of the firmware running on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKDevice/firmwareVersion
func (h_ HKDevice) FirmwareVersion() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("firmwareVersion"))
	return rv
}/* debug [instance_properties/getter]: firmwareVersion */


// An arbitrary string representing the hardware version of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKDevice/hardwareVersion
func (h_ HKDevice) HardwareVersion() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("hardwareVersion"))
	return rv
}/* debug [instance_properties/getter]: hardwareVersion */


// An identifier that uniquely identifies the device object on the hardware running this code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKDevice/localIdentifier
func (h_ HKDevice) LocalIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("localIdentifier"))
	return rv
}/* debug [instance_properties/getter]: localIdentifier */


// A string representing the device’s manufacturer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKDevice/manufacturer
func (h_ HKDevice) Manufacturer() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("manufacturer"))
	return rv
}/* debug [instance_properties/getter]: manufacturer */


// A string representing the device’s model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKDevice/model
func (h_ HKDevice) Model() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("model"))
	return rv
}/* debug [instance_properties/getter]: model */


// The user-facing name for the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKDevice/name
func (h_ HKDevice) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// An arbitrary string representing the version of the software running on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKDevice/softwareVersion
func (h_ HKDevice) SoftwareVersion() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("softwareVersion"))
	return rv
}/* debug [instance_properties/getter]: softwareVersion */


// The device identifier portion of the US Food and Drug Administration’s Unique Device Identifier (UDI).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKDevice/udiDeviceIdentifier
func (h_ HKDevice) UDIDeviceIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("UDIDeviceIdentifier"))
	return rv
}/* debug [instance_properties/getter]: UDIDeviceIdentifier */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKDevice */


