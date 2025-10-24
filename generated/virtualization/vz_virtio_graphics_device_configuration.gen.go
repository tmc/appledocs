// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	IVZGraphicsDeviceConfiguration
	// properties:
	Scanouts() []IVZVirtioGraphicsScanoutConfiguration
	SetScanouts(value []IVZVirtioGraphicsScanoutConfiguration)
	// methods:
}

// Configuration that represents the configuration of a Virtio graphics device for a Linux VM.


// Configuration that represents the configuration of a Virtio graphics device for a Linux VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioGraphicsDeviceConfiguration
type VZVirtioGraphicsDeviceConfiguration struct {
	VZGraphicsDeviceConfiguration
}

// VZVirtioGraphicsDeviceConfigurationFrom constructs a [VZVirtioGraphicsDeviceConfiguration] from an unsafe.Pointer.
//
// Configuration that represents the configuration of a Virtio graphics device for a Linux VM.
func VZVirtioGraphicsDeviceConfigurationFrom(ptr unsafe.Pointer) VZVirtioGraphicsDeviceConfiguration {
	return VZVirtioGraphicsDeviceConfiguration{
		VZGraphicsDeviceConfiguration: VZGraphicsDeviceConfigurationFrom(ptr),
	}
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




// The array of output devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioGraphicsDeviceConfiguration/scanouts
func (v_ VZVirtioGraphicsDeviceConfiguration) Scanouts() []IVZVirtioGraphicsScanoutConfiguration {
	rv := objc.Send[[]VZVirtioGraphicsScanoutConfiguration](v_.ID, objc.Sel("scanouts"))
	return rv
}


// The array of output devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioGraphicsDeviceConfiguration/scanouts
func (v_ VZVirtioGraphicsDeviceConfiguration) SetScanouts(value []IVZVirtioGraphicsScanoutConfiguration) {
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
	objc.Send[objc.ID](v_.ID, objc.Sel("setScanouts:"), nsArray)
}


