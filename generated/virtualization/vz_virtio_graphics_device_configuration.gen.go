// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [VZVirtioGraphicsDeviceConfiguration] class.
var (
	VZVirtioGraphicsDeviceConfigurationClass     _VZVirtioGraphicsDeviceConfigurationClass
	VZVirtioGraphicsDeviceConfigurationClassOnce sync.Once
)

func getVZVirtioGraphicsDeviceConfigurationClass() _VZVirtioGraphicsDeviceConfigurationClass {
	VZVirtioGraphicsDeviceConfigurationClassOnce.Do(func() {
		VZVirtioGraphicsDeviceConfigurationClass = _VZVirtioGraphicsDeviceConfigurationClass{objc.GetClass("VZVirtioGraphicsDeviceConfiguration")}
	})
	return VZVirtioGraphicsDeviceConfigurationClass
}

type _VZVirtioGraphicsDeviceConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [VZVirtioGraphicsDeviceConfiguration] class.
type IVZVirtioGraphicsDeviceConfiguration interface {
	objectivec.IObject
}

// Configuration that represents the configuration of a Virtio graphics device for a Linux VM.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioGraphicsDeviceConfiguration
type VZVirtioGraphicsDeviceConfiguration struct {
	objectivec.Object
}

// VZVirtioGraphicsDeviceConfigurationFrom constructs a [VZVirtioGraphicsDeviceConfiguration] from an unsafe.Pointer.
//
// Configuration that represents the configuration of a Virtio graphics device for a Linux VM.
func VZVirtioGraphicsDeviceConfigurationFrom(ptr unsafe.Pointer) VZVirtioGraphicsDeviceConfiguration {
	return VZVirtioGraphicsDeviceConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VZVirtioGraphicsDeviceConfigurationClass) Alloc() VZVirtioGraphicsDeviceConfiguration {
	rv := objc.Send[VZVirtioGraphicsDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZVirtioGraphicsDeviceConfigurationClass) New() VZVirtioGraphicsDeviceConfiguration {
	rv := objc.Send[VZVirtioGraphicsDeviceConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZVirtioGraphicsDeviceConfiguration) Init() VZVirtioGraphicsDeviceConfiguration {
	rv := objc.Send[VZVirtioGraphicsDeviceConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZVirtioGraphicsDeviceConfiguration) Autorelease() VZVirtioGraphicsDeviceConfiguration {
	rv := objc.Send[VZVirtioGraphicsDeviceConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioGraphicsDeviceConfiguration creates a new VZVirtioGraphicsDeviceConfiguration instance.
func NewVZVirtioGraphicsDeviceConfiguration() VZVirtioGraphicsDeviceConfiguration {
	return getVZVirtioGraphicsDeviceConfigurationClass().New()
}

// The array of graphics scanout configurations.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioGraphicsDeviceConfiguration/scanouts
func (v_ VZVirtioGraphicsDeviceConfiguration) Scanouts() []VZVirtioGraphicsScanoutConfiguration {
	rv := objc.Send[[]VZVirtioGraphicsScanoutConfiguration](v_.ID, objc.Sel("scanouts"))
	return rv
}

// SetScanouts sets the value of the scanouts property.
// The array of graphics scanout configurations.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioGraphicsDeviceConfiguration/scanouts
func (v_ VZVirtioGraphicsDeviceConfiguration) SetScanouts(value []VZVirtioGraphicsScanoutConfiguration) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setScanouts:"), value)
}



