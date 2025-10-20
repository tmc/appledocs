// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [VZVirtioConsolePortConfiguration] class.
var (
	VZVirtioConsolePortConfigurationClass     _VZVirtioConsolePortConfigurationClass
	VZVirtioConsolePortConfigurationClassOnce sync.Once
)

func getVZVirtioConsolePortConfigurationClass() _VZVirtioConsolePortConfigurationClass {
	VZVirtioConsolePortConfigurationClassOnce.Do(func() {
		VZVirtioConsolePortConfigurationClass = _VZVirtioConsolePortConfigurationClass{objc.GetClass("VZVirtioConsolePortConfiguration")}
	})
	return VZVirtioConsolePortConfigurationClass
}

type _VZVirtioConsolePortConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [VZVirtioConsolePortConfiguration] class.
type IVZVirtioConsolePortConfiguration interface {
	IVZConsolePortConfiguration
}

// A class that represents the configuration options you can set on a Virtio console port.
//
// A console port is a two-way communication channel between a host and a VM console port. A Virtio device can have one or more attached console devices. Optionally, you can set a name for a console port and also configure a console port that the guest can use as the system console.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePortConfiguration
type VZVirtioConsolePortConfiguration struct {
	VZConsolePortConfiguration
}

// VZVirtioConsolePortConfigurationFrom constructs a [VZVirtioConsolePortConfiguration] from an unsafe.Pointer.
//
// A class that represents the configuration options you can set on a Virtio console port.
func VZVirtioConsolePortConfigurationFrom(ptr unsafe.Pointer) VZVirtioConsolePortConfiguration {
	return VZVirtioConsolePortConfiguration{
		VZConsolePortConfiguration: VZConsolePortConfigurationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (vc _VZVirtioConsolePortConfigurationClass) Alloc() VZVirtioConsolePortConfiguration {
	rv := objc.Send[VZVirtioConsolePortConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZVirtioConsolePortConfigurationClass) New() VZVirtioConsolePortConfiguration {
	rv := objc.Send[VZVirtioConsolePortConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZVirtioConsolePortConfiguration) Init() VZVirtioConsolePortConfiguration {
	rv := objc.Send[VZVirtioConsolePortConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZVirtioConsolePortConfiguration) Autorelease() VZVirtioConsolePortConfiguration {
	rv := objc.Send[VZVirtioConsolePortConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioConsolePortConfiguration creates a new VZVirtioConsolePortConfiguration instance.
func NewVZVirtioConsolePortConfiguration() VZVirtioConsolePortConfiguration {
	return getVZVirtioConsolePortConfigurationClass().New()
}



// A Boolean value that indicates whether this port is a console.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePortConfiguration/isConsole
func (v_ VZVirtioConsolePortConfiguration) IsConsole() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isConsole"))
	return rv
}


// SetIsConsole sets the value of the isConsole property.
// A Boolean value that indicates whether this port is a console.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePortConfiguration/isConsole
func (v_ VZVirtioConsolePortConfiguration) SetIsConsole(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setIsConsole:"), value)
}
// The name of the port.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePortConfiguration/name
func (v_ VZVirtioConsolePortConfiguration) Name() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
// The name of the port.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePortConfiguration/name
func (v_ VZVirtioConsolePortConfiguration) SetName(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setName:"), value)
}

