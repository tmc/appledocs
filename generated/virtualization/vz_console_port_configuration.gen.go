// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [VZConsolePortConfiguration] class.
var (
	VZConsolePortConfigurationClass     _VZConsolePortConfigurationClass
	VZConsolePortConfigurationClassOnce sync.Once
)

func getVZConsolePortConfigurationClass() _VZConsolePortConfigurationClass {
	VZConsolePortConfigurationClassOnce.Do(func() {
		VZConsolePortConfigurationClass = _VZConsolePortConfigurationClass{objc.GetClass("VZConsolePortConfiguration")}
	})
	return VZConsolePortConfigurationClass
}

type _VZConsolePortConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [VZConsolePortConfiguration] class.
type IVZConsolePortConfiguration interface {
	objectivec.IObject
}

// The base class for a console port configuration.
//
// Don’t instantiate directly, instead use one of its subclasses like .
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZConsolePortConfiguration
type VZConsolePortConfiguration struct {
	objectivec.Object
}

// VZConsolePortConfigurationFrom constructs a [VZConsolePortConfiguration] from an unsafe.Pointer.
//
// The base class for a console port configuration.
func VZConsolePortConfigurationFrom(ptr unsafe.Pointer) VZConsolePortConfiguration {
	return VZConsolePortConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VZConsolePortConfigurationClass) Alloc() VZConsolePortConfiguration {
	rv := objc.Send[VZConsolePortConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZConsolePortConfigurationClass) New() VZConsolePortConfiguration {
	rv := objc.Send[VZConsolePortConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZConsolePortConfiguration) Init() VZConsolePortConfiguration {
	rv := objc.Send[VZConsolePortConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZConsolePortConfiguration) Autorelease() VZConsolePortConfiguration {
	rv := objc.Send[VZConsolePortConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZConsolePortConfiguration creates a new VZConsolePortConfiguration instance.
func NewVZConsolePortConfiguration() VZConsolePortConfiguration {
	return getVZConsolePortConfigurationClass().New()
}


// The serial port attachment.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZConsolePortConfiguration/attachment
func (v_ VZConsolePortConfiguration) Attachment() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("attachment"))
	return rv
}


// SetAttachment sets the value of the attachment property.
// The serial port attachment.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZConsolePortConfiguration/attachment
func (v_ VZConsolePortConfiguration) SetAttachment(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAttachment:"), value)
}



