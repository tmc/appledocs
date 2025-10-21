// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [VZVirtioConsolePortArray] class.
var (
	VZVirtioConsolePortArrayClass     _VZVirtioConsolePortArrayClass
	VZVirtioConsolePortArrayClassOnce sync.Once
)

func getVZVirtioConsolePortArrayClass() _VZVirtioConsolePortArrayClass {
	VZVirtioConsolePortArrayClassOnce.Do(func() {
		VZVirtioConsolePortArrayClass = _VZVirtioConsolePortArrayClass{objc.GetClass("VZVirtioConsolePortArray")}
	})
	return VZVirtioConsolePortArrayClass
}

type _VZVirtioConsolePortArrayClass struct {
	class objc.Class
}

// An interface definition for the [VZVirtioConsolePortArray] class.
type IVZVirtioConsolePortArray interface {
	objectivec.IObject
	ObjectAtIndexedSubscript(portIndex uint) unsafe.Pointer
}

// A class that represents a collection of Virtio console ports.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePortArray
type VZVirtioConsolePortArray struct {
	objectivec.Object
}

// VZVirtioConsolePortArrayFrom constructs a [VZVirtioConsolePortArray] from an unsafe.Pointer.
//
// A class that represents a collection of Virtio console ports.
func VZVirtioConsolePortArrayFrom(ptr unsafe.Pointer) VZVirtioConsolePortArray {
	return VZVirtioConsolePortArray{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VZVirtioConsolePortArrayClass) Alloc() VZVirtioConsolePortArray {
	rv := objc.Send[VZVirtioConsolePortArray](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZVirtioConsolePortArrayClass) New() VZVirtioConsolePortArray {
	rv := objc.Send[VZVirtioConsolePortArray](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZVirtioConsolePortArray) Init() VZVirtioConsolePortArray {
	rv := objc.Send[VZVirtioConsolePortArray](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZVirtioConsolePortArray) Autorelease() VZVirtioConsolePortArray {
	rv := objc.Send[VZVirtioConsolePortArray](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioConsolePortArray creates a new VZVirtioConsolePortArray instance.
func NewVZVirtioConsolePortArray() VZVirtioConsolePortArray {
	return getVZVirtioConsolePortArrayClass().New()
}


// Returns the Virtio console port at the specified index.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePortArray/subscript(_:)
func (v_ VZVirtioConsolePortArray) ObjectAtIndexedSubscript(portIndex uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("objectAtIndexedSubscript:"), portIndex)
	return rv
}

// An unsigned integer that represents the maximum number of ports allocated by this device.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePortArray/maximumPortCount
func (v_ VZVirtioConsolePortArray) MaximumPortCount() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("maximumPortCount"))
	return rv
}



