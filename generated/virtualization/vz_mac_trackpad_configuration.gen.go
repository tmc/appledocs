// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [VZMacTrackpadConfiguration] class.
type IVZMacTrackpadConfiguration interface {
	IVZPointingDeviceConfiguration
}

// The class that represents the configuration for a Mac trackpad.
//
// The uses this device to send pointer events and multi-touch trackpad gestures to the virtual machine. In macOS 13 and later, guests use the multi-touch trackpad device, while earlier versions of macOS uses the USB pointing device.
//
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

// Alloc allocates a new instance without initialization.
func (vc _VZMacTrackpadConfigurationClass) Alloc() VZMacTrackpadConfiguration {
	rv := objc.Send[VZMacTrackpadConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The list of pointing devices.
//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/pointingdevices
func (v_ VZMacTrackpadConfiguration) PointingDevices() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("pointingDevices"))
	return rv
}


// SetPointingDevices sets the value of the pointingDevices property.
// The list of pointing devices.

//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/pointingdevices
func (v_ VZMacTrackpadConfiguration) SetPointingDevices(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setPointingDevices:"), value)
}


