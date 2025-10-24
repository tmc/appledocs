// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZUSBControllerConfiguration */

/* debug [class_header]: Header for VZUSBControllerConfiguration */
// The class instance for the [VZUSBControllerConfiguration] class.
var (
	VZUSBControllerConfigurationClass     _VZUSBControllerConfigurationClass
	VZUSBControllerConfigurationClassOnce sync.Once
)

func getVZUSBControllerConfigurationClass() _VZUSBControllerConfigurationClass {
	VZUSBControllerConfigurationClassOnce.Do(func() {
		VZUSBControllerConfigurationClass = _VZUSBControllerConfigurationClass{objc.GetClass("VZUSBControllerConfiguration")}
	})
	return VZUSBControllerConfigurationClass
}

type _VZUSBControllerConfigurationClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for VZUSBControllerConfiguration */
// An interface definition for the [VZUSBControllerConfiguration] class.
type IVZUSBControllerConfiguration interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for VZUSBControllerConfiguration */
	// properties:
	UsbDevices() []objc.ID
	SetUsbDevices(value []objc.ID)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for VZUSBControllerConfiguration */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for VZUSBControllerConfiguration */
// Alloc allocates a new instance without initialization.
func (vc _VZUSBControllerConfigurationClass) Alloc() VZUSBControllerConfiguration {
	rv := objc.Send[VZUSBControllerConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZUSBControllerConfigurationClass) New() VZUSBControllerConfiguration {
	rv := objc.Send[VZUSBControllerConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZUSBControllerConfiguration) Init() VZUSBControllerConfiguration {
	rv := objc.Send[VZUSBControllerConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZUSBControllerConfiguration) Autorelease() VZUSBControllerConfiguration {
	rv := objc.Send[VZUSBControllerConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZUSBControllerConfiguration creates a new VZUSBControllerConfiguration instance.
func NewVZUSBControllerConfiguration() VZUSBControllerConfiguration {
	return getVZUSBControllerConfigurationClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for VZUSBControllerConfiguration */
// The base class for a USB controller configuration.
//
// Don’t create objects directly. Use one of its subclasses, such as , instead.

// The base class for a USB controller configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZUSBControllerConfiguration
type VZUSBControllerConfiguration struct {
	objectivec.Object
}

// VZUSBControllerConfigurationFrom constructs a [VZUSBControllerConfiguration] from an unsafe.Pointer.
//
// The base class for a USB controller configuration.
func VZUSBControllerConfigurationFrom(ptr unsafe.Pointer) VZUSBControllerConfiguration {
	return VZUSBControllerConfiguration{objectivec.Object{objc.ID(ptr)}}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for VZUSBControllerConfiguration */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for VZUSBControllerConfiguration */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for VZUSBControllerConfiguration */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for VZUSBControllerConfiguration */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for VZUSBControllerConfiguration */

// The list of USB devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZUSBControllerConfiguration/usbDevices
func (v_ VZUSBControllerConfiguration) UsbDevices() []objc.ID {
	rv := objc.Send[[]objc.ID](v_.ID, objc.Sel("usbDevices"))
	return rv
} /* debug [instance_properties/getter]: usbDevices */

// The list of USB devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZUSBControllerConfiguration/usbDevices
func (v_ VZUSBControllerConfiguration) SetUsbDevices(value []objc.ID) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](v_.ID, objc.Sel("setUsbDevices:"), nsArray)
} /* debug [instance_properties/setter]: usbDevices */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class VZUSBControllerConfiguration */
