// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [VZConsoleDevice] class.
var (
	VZConsoleDeviceClass     _VZConsoleDeviceClass
	VZConsoleDeviceClassOnce sync.Once
)

func getVZConsoleDeviceClass() _VZConsoleDeviceClass {
	VZConsoleDeviceClassOnce.Do(func() {
		VZConsoleDeviceClass = _VZConsoleDeviceClass{objc.GetClass("VZConsoleDevice")}
	})
	return VZConsoleDeviceClass
}

type _VZConsoleDeviceClass struct {
	class objc.Class
}

// An interface definition for the [VZConsoleDevice] class.
type IVZConsoleDevice interface {
	objectivec.IObject
}

// A class that represents a console device in a VM.
//
// Don’t instantiate a directly: You first configure console devices on the through a subclass of . After you create from the configuration, the console devices are available through the property. The actual type of corresponds to the type that the configuration uses. For example, a is a device of type .
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZConsoleDevice
type VZConsoleDevice struct {
	objectivec.Object
}

// VZConsoleDeviceFrom constructs a [VZConsoleDevice] from an unsafe.Pointer.
//
// A class that represents a console device in a VM.
func VZConsoleDeviceFrom(ptr unsafe.Pointer) VZConsoleDevice {
	return VZConsoleDevice{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VZConsoleDeviceClass) Alloc() VZConsoleDevice {
	rv := objc.Send[VZConsoleDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZConsoleDeviceClass) New() VZConsoleDevice {
	rv := objc.Send[VZConsoleDevice](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZConsoleDevice) Init() VZConsoleDevice {
	rv := objc.Send[VZConsoleDevice](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZConsoleDevice) Autorelease() VZConsoleDevice {
	rv := objc.Send[VZConsoleDevice](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZConsoleDevice creates a new VZConsoleDevice instance.
func NewVZConsoleDevice() VZConsoleDevice {
	return getVZConsoleDeviceClass().New()
}




