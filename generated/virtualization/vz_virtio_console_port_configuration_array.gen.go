// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [VZVirtioConsolePortConfigurationArray] class.
var (
	VZVirtioConsolePortConfigurationArrayClass     _VZVirtioConsolePortConfigurationArrayClass
	VZVirtioConsolePortConfigurationArrayClassOnce sync.Once
)

func getVZVirtioConsolePortConfigurationArrayClass() _VZVirtioConsolePortConfigurationArrayClass {
	VZVirtioConsolePortConfigurationArrayClassOnce.Do(func() {
		VZVirtioConsolePortConfigurationArrayClass = _VZVirtioConsolePortConfigurationArrayClass{objc.GetClass("VZVirtioConsolePortConfigurationArray")}
	})
	return VZVirtioConsolePortConfigurationArrayClass
}

type _VZVirtioConsolePortConfigurationArrayClass struct {
	class objc.Class
}

// An interface definition for the [VZVirtioConsolePortConfigurationArray] class.
type IVZVirtioConsolePortConfigurationArray interface {
	objectivec.IObject
	SetObjectAtIndexedSubscript(configuration IVZVirtioConsolePortConfiguration, portIndex uint)
	ObjectAtIndexedSubscript(portIndex uint) VZVirtioConsolePortConfiguration
	MaximumPortCount() uint32
	SetMaximumPortCount(value Iuint32)
}

// A class that represents a collection of Virtio console port configurations.
//
// This array stores a collection of port configurations for a . The index in the array corresponds to the port index that the VM uses. You can set a value, but the value must be larger than the highest indexed port. If there’s no value set, the framework uses the value the highest indexed port.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePortConfigurationArray
type VZVirtioConsolePortConfigurationArray struct {
	objectivec.Object
}

// VZVirtioConsolePortConfigurationArrayFrom constructs a [VZVirtioConsolePortConfigurationArray] from an unsafe.Pointer.
//
// A class that represents a collection of Virtio console port configurations.
func VZVirtioConsolePortConfigurationArrayFrom(ptr unsafe.Pointer) VZVirtioConsolePortConfigurationArray {
	return VZVirtioConsolePortConfigurationArray{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VZVirtioConsolePortConfigurationArrayClass) Alloc() VZVirtioConsolePortConfigurationArray {
	rv := objc.Send[VZVirtioConsolePortConfigurationArray](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZVirtioConsolePortConfigurationArrayClass) New() VZVirtioConsolePortConfigurationArray {
	rv := objc.Send[VZVirtioConsolePortConfigurationArray](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZVirtioConsolePortConfigurationArray) Init() VZVirtioConsolePortConfigurationArray {
	rv := objc.Send[VZVirtioConsolePortConfigurationArray](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZVirtioConsolePortConfigurationArray) Autorelease() VZVirtioConsolePortConfigurationArray {
	rv := objc.Send[VZVirtioConsolePortConfigurationArray](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioConsolePortConfigurationArray creates a new VZVirtioConsolePortConfigurationArray instance.
func NewVZVirtioConsolePortConfigurationArray() VZVirtioConsolePortConfigurationArray {
	return getVZVirtioConsolePortConfigurationArrayClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePortConfigurationArray/setObject:atIndexedSubscript:
func (v_ VZVirtioConsolePortConfigurationArray) SetObjectAtIndexedSubscript(configuration IVZVirtioConsolePortConfiguration, portIndex uint) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setObject:atIndexedSubscript:"), configuration, portIndex)
}

// Returns the Virtio console port configuration as the specified index.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePortConfigurationArray/subscript(_:)
func (v_ VZVirtioConsolePortConfigurationArray) ObjectAtIndexedSubscript(portIndex uint) VZVirtioConsolePortConfiguration {
	rv := objc.Send[VZVirtioConsolePortConfiguration](v_.ID, objc.Sel("objectAtIndexedSubscript:"), portIndex)
	return rv
}

// An unsigned integer that represents the maximum number of ports allocated by this device.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePortConfigurationArray/maximumPortCount
func (v_ VZVirtioConsolePortConfigurationArray) MaximumPortCount() uint32 {
	rv := objc.Send[uint32](v_.ID, objc.Sel("maximumPortCount"))
	return rv
}


// SetMaximumPortCount sets the value of the maximumPortCount property.
// An unsigned integer that represents the maximum number of ports allocated by this device.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePortConfigurationArray/maximumPortCount
func (v_ VZVirtioConsolePortConfigurationArray) SetMaximumPortCount(value Iuint32) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setMaximumPortCount:"), value)
}



