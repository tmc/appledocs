// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [VZStorageDevice] class.
var (
	VZStorageDeviceClass     _VZStorageDeviceClass
	VZStorageDeviceClassOnce sync.Once
)

func getVZStorageDeviceClass() _VZStorageDeviceClass {
	VZStorageDeviceClassOnce.Do(func() {
		VZStorageDeviceClass = _VZStorageDeviceClass{objc.GetClass("VZStorageDevice")}
	})
	return VZStorageDeviceClass
}

type _VZStorageDeviceClass struct {
	class objc.Class
}

// An interface definition for the [VZStorageDevice] class.
type IVZStorageDevice interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A class that represents a storage device in a VM.
//
// Don’t create a directly. Use one of its subclasses, such as , instead.


// A class that represents a storage device in a VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZStorageDevice
type VZStorageDevice struct {
	objectivec.Object
}

// VZStorageDeviceFrom constructs a [VZStorageDevice] from an unsafe.Pointer.
//
// A class that represents a storage device in a VM.
func VZStorageDeviceFrom(ptr unsafe.Pointer) VZStorageDevice {
	return VZStorageDevice{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VZStorageDeviceClass) Alloc() VZStorageDevice {
	rv := objc.Send[VZStorageDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZStorageDeviceClass) New() VZStorageDevice {
	rv := objc.Send[VZStorageDevice](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZStorageDevice) Init() VZStorageDevice {
	rv := objc.Send[VZStorageDevice](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZStorageDevice) Autorelease() VZStorageDevice {
	rv := objc.Send[VZStorageDevice](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZStorageDevice creates a new VZStorageDevice instance.
func NewVZStorageDevice() VZStorageDevice {
	return getVZStorageDeviceClass().New()
}




