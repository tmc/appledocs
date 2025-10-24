// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZVirtioEntropyDeviceConfiguration */


/* debug [class_header]: Header for VZVirtioEntropyDeviceConfiguration */
// The class instance for the [VZVirtioEntropyDeviceConfiguration] class.
var (
	VZVirtioEntropyDeviceConfigurationClass     _VZVirtioEntropyDeviceConfigurationClass
	VZVirtioEntropyDeviceConfigurationClassOnce sync.Once
)

func getVZVirtioEntropyDeviceConfigurationClass() _VZVirtioEntropyDeviceConfigurationClass {
	VZVirtioEntropyDeviceConfigurationClassOnce.Do(func() {
		VZVirtioEntropyDeviceConfigurationClass = _VZVirtioEntropyDeviceConfigurationClass{objc.GetClass("VZVirtioEntropyDeviceConfiguration")}
	})
	return VZVirtioEntropyDeviceConfigurationClass
}

type _VZVirtioEntropyDeviceConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZVirtioEntropyDeviceConfiguration */
// An interface definition for the [VZVirtioEntropyDeviceConfiguration] class.
type IVZVirtioEntropyDeviceConfiguration interface {
	IVZEntropyDeviceConfiguration
	
/* debug [class_interface_properties]: Properties for VZVirtioEntropyDeviceConfiguration */
	// properties:
	EntropyDevices() IVZEntropyDeviceConfiguration
	SetEntropyDevices(value IVZEntropyDeviceConfiguration)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZVirtioEntropyDeviceConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZVirtioEntropyDeviceConfiguration */
// Alloc allocates a new instance without initialization.
func (vc _VZVirtioEntropyDeviceConfigurationClass) Alloc() VZVirtioEntropyDeviceConfiguration {
	rv := objc.Send[VZVirtioEntropyDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZVirtioEntropyDeviceConfigurationClass) New() VZVirtioEntropyDeviceConfiguration {
	rv := objc.Send[VZVirtioEntropyDeviceConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZVirtioEntropyDeviceConfiguration) Init() VZVirtioEntropyDeviceConfiguration {
	rv := objc.Send[VZVirtioEntropyDeviceConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZVirtioEntropyDeviceConfiguration) Autorelease() VZVirtioEntropyDeviceConfiguration {
	rv := objc.Send[VZVirtioEntropyDeviceConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioEntropyDeviceConfiguration creates a new VZVirtioEntropyDeviceConfiguration instance.
func NewVZVirtioEntropyDeviceConfiguration() VZVirtioEntropyDeviceConfiguration {
	return getVZVirtioEntropyDeviceConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZVirtioEntropyDeviceConfiguration */
// A source of entropy for the guest’s random number generator.
//
// Use a object to expose a source of entropy for the guest operating system’s random-number generator. When you create this object and add it to your virtual machine’s configuration, the virtual machine configures a Virtio-compliant entropy device. The guest operating system uses this device as a seed to generate random numbers. Create a object and add it to the property of your virtual machine’s configuration.


// A source of entropy for the guest’s random number generator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioEntropyDeviceConfiguration
type VZVirtioEntropyDeviceConfiguration struct {
	VZEntropyDeviceConfiguration
}

// VZVirtioEntropyDeviceConfigurationFrom constructs a [VZVirtioEntropyDeviceConfiguration] from an unsafe.Pointer.
//
// A source of entropy for the guest’s random number generator.
func VZVirtioEntropyDeviceConfigurationFrom(ptr unsafe.Pointer) VZVirtioEntropyDeviceConfiguration {
	return VZVirtioEntropyDeviceConfiguration{
		VZEntropyDeviceConfiguration: VZEntropyDeviceConfigurationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZVirtioEntropyDeviceConfiguration */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZVirtioEntropyDeviceConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZVirtioEntropyDeviceConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZVirtioEntropyDeviceConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZVirtioEntropyDeviceConfiguration */

// The array of randomization devices that you expose to the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/entropydevices
func (v_ VZVirtioEntropyDeviceConfiguration) EntropyDevices() IVZEntropyDeviceConfiguration {
	rv := objc.Send[VZEntropyDeviceConfiguration](v_.ID, objc.Sel("entropyDevices"))
	return rv
}/* debug [instance_properties/getter]: entropyDevices */


// The array of randomization devices that you expose to the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/entropydevices
func (v_ VZVirtioEntropyDeviceConfiguration) SetEntropyDevices(value IVZEntropyDeviceConfiguration) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setEntropyDevices:"), value)
}/* debug [instance_properties/setter]: entropyDevices */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZVirtioEntropyDeviceConfiguration */


