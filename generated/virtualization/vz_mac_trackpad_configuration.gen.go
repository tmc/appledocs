// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZMacTrackpadConfiguration */


/* debug [class_header]: Header for VZMacTrackpadConfiguration */
// The class instance for the [VZMacTrackpadConfiguration] class.
var (
	VZMacTrackpadConfigurationClass     _VZMacTrackpadConfigurationClass
	VZMacTrackpadConfigurationClassOnce sync.Once
)

func getVZMacTrackpadConfigurationClass() _VZMacTrackpadConfigurationClass {
	VZMacTrackpadConfigurationClassOnce.Do(func() {
		VZMacTrackpadConfigurationClass = _VZMacTrackpadConfigurationClass{objc.GetClass("VZMacTrackpadConfiguration")}
	})
	return VZMacTrackpadConfigurationClass
}

type _VZMacTrackpadConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZMacTrackpadConfiguration */
// An interface definition for the [VZMacTrackpadConfiguration] class.
type IVZMacTrackpadConfiguration interface {
	IVZPointingDeviceConfiguration
	
/* debug [class_interface_properties]: Properties for VZMacTrackpadConfiguration */
	// properties:
	PointingDevices() IVZPointingDeviceConfiguration
	SetPointingDevices(value IVZPointingDeviceConfiguration)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZMacTrackpadConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZMacTrackpadConfiguration */
// Alloc allocates a new instance without initialization.
func (vc _VZMacTrackpadConfigurationClass) Alloc() VZMacTrackpadConfiguration {
	rv := objc.Send[VZMacTrackpadConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZMacTrackpadConfigurationClass) New() VZMacTrackpadConfiguration {
	rv := objc.Send[VZMacTrackpadConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZMacTrackpadConfiguration) Init() VZMacTrackpadConfiguration {
	rv := objc.Send[VZMacTrackpadConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZMacTrackpadConfiguration) Autorelease() VZMacTrackpadConfiguration {
	rv := objc.Send[VZMacTrackpadConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacTrackpadConfiguration creates a new VZMacTrackpadConfiguration instance.
func NewVZMacTrackpadConfiguration() VZMacTrackpadConfiguration {
	return getVZMacTrackpadConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZMacTrackpadConfiguration */
// The class that represents the configuration for a Mac trackpad.
//
// The uses this device to send pointer events and multi-touch trackpad gestures to the virtual machine. In macOS 13 and later, guests use the multi-touch trackpad device, while earlier versions of macOS uses the USB pointing device.


// The class that represents the configuration for a Mac trackpad.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacTrackpadConfiguration
type VZMacTrackpadConfiguration struct {
	VZPointingDeviceConfiguration
}

// VZMacTrackpadConfigurationFrom constructs a [VZMacTrackpadConfiguration] from an unsafe.Pointer.
//
// The class that represents the configuration for a Mac trackpad.
func VZMacTrackpadConfigurationFrom(ptr unsafe.Pointer) VZMacTrackpadConfiguration {
	return VZMacTrackpadConfiguration{
		VZPointingDeviceConfiguration: VZPointingDeviceConfigurationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZMacTrackpadConfiguration */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZMacTrackpadConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZMacTrackpadConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZMacTrackpadConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZMacTrackpadConfiguration */

// The list of pointing devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/pointingdevices
func (v_ VZMacTrackpadConfiguration) PointingDevices() IVZPointingDeviceConfiguration {
	rv := objc.Send[VZPointingDeviceConfiguration](v_.ID, objc.Sel("pointingDevices"))
	return rv
}/* debug [instance_properties/getter]: pointingDevices */


// The list of pointing devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/pointingdevices
func (v_ VZMacTrackpadConfiguration) SetPointingDevices(value IVZPointingDeviceConfiguration) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setPointingDevices:"), value)
}/* debug [instance_properties/setter]: pointingDevices */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZMacTrackpadConfiguration */


