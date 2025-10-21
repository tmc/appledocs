// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [VZDirectorySharingDevice] class.
var (
	VZDirectorySharingDeviceClass     _VZDirectorySharingDeviceClass
	VZDirectorySharingDeviceClassOnce sync.Once
)

func getVZDirectorySharingDeviceClass() _VZDirectorySharingDeviceClass {
	VZDirectorySharingDeviceClassOnce.Do(func() {
		VZDirectorySharingDeviceClass = _VZDirectorySharingDeviceClass{objc.GetClass("VZDirectorySharingDevice")}
	})
	return VZDirectorySharingDeviceClass
}

type _VZDirectorySharingDeviceClass struct {
	class objc.Class
}

// An interface definition for the [VZDirectorySharingDevice] class.
type IVZDirectorySharingDevice interface {
	objectivec.IObject
}

// The base class that represents a directory sharing device in a VM.
//
// Don’t instantiate directly; configure a directory sharing device first by using through a subclass of . When you create a from the configuration, the directory sharing devices are available through the property. The real type of corresponds to the type used by the configuration. For example, a leads to a device of type .
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZDirectorySharingDevice
type VZDirectorySharingDevice struct {
	objectivec.Object
}

// VZDirectorySharingDeviceFrom constructs a [VZDirectorySharingDevice] from an unsafe.Pointer.
//
// The base class that represents a directory sharing device in a VM.
func VZDirectorySharingDeviceFrom(ptr unsafe.Pointer) VZDirectorySharingDevice {
	return VZDirectorySharingDevice{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VZDirectorySharingDeviceClass) Alloc() VZDirectorySharingDevice {
	rv := objc.Send[VZDirectorySharingDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZDirectorySharingDeviceClass) New() VZDirectorySharingDevice {
	rv := objc.Send[VZDirectorySharingDevice](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZDirectorySharingDevice) Init() VZDirectorySharingDevice {
	rv := objc.Send[VZDirectorySharingDevice](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZDirectorySharingDevice) Autorelease() VZDirectorySharingDevice {
	rv := objc.Send[VZDirectorySharingDevice](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZDirectorySharingDevice creates a new VZDirectorySharingDevice instance.
func NewVZDirectorySharingDevice() VZDirectorySharingDevice {
	return getVZDirectorySharingDeviceClass().New()
}




