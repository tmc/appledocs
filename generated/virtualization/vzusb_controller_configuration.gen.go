// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [VZUSBControllerConfiguration] class.
type IVZUSBControllerConfiguration interface {
	objectivec.IObject
	UsbDevices() []objc.ID
	SetUsbDevices(value []objc.ID)
}

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

// Alloc allocates a new instance without initialization.
func (vc _VZUSBControllerConfigurationClass) Alloc() VZUSBControllerConfiguration {
	rv := objc.Send[VZUSBControllerConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The list of USB devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZUSBControllerConfiguration/usbDevices
func (v_ VZUSBControllerConfiguration) UsbDevices() []objc.ID {
	rv := objc.Send[[]objc.ID](v_.ID, objc.Sel("usbDevices"))
	return rv
}


// The list of USB devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZUSBControllerConfiguration/usbDevices
func (v_ VZUSBControllerConfiguration) SetUsbDevices(value []objc.ID) {
	// Convert Go slice to NSArray
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
}



