// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [VZAudioDeviceConfiguration] class.
var (
	VZAudioDeviceConfigurationClass     _VZAudioDeviceConfigurationClass
	VZAudioDeviceConfigurationClassOnce sync.Once
)

func getVZAudioDeviceConfigurationClass() _VZAudioDeviceConfigurationClass {
	VZAudioDeviceConfigurationClassOnce.Do(func() {
		VZAudioDeviceConfigurationClass = _VZAudioDeviceConfigurationClass{objc.GetClass("VZAudioDeviceConfiguration")}
	})
	return VZAudioDeviceConfigurationClass
}

type _VZAudioDeviceConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [VZAudioDeviceConfiguration] class.
type IVZAudioDeviceConfiguration interface {
	objectivec.IObject
}

// The base class for an audio device configuration.
//
// Don’t instantiate this abstract class directly. Instead, instantiate one of its subclasses such as .
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZAudioDeviceConfiguration
type VZAudioDeviceConfiguration struct {
	objectivec.Object
}

// VZAudioDeviceConfigurationFrom constructs a [VZAudioDeviceConfiguration] from an unsafe.Pointer.
//
// The base class for an audio device configuration.
func VZAudioDeviceConfigurationFrom(ptr unsafe.Pointer) VZAudioDeviceConfiguration {
	return VZAudioDeviceConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VZAudioDeviceConfigurationClass) Alloc() VZAudioDeviceConfiguration {
	rv := objc.Send[VZAudioDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZAudioDeviceConfigurationClass) New() VZAudioDeviceConfiguration {
	rv := objc.Send[VZAudioDeviceConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZAudioDeviceConfiguration) Init() VZAudioDeviceConfiguration {
	rv := objc.Send[VZAudioDeviceConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZAudioDeviceConfiguration) Autorelease() VZAudioDeviceConfiguration {
	rv := objc.Send[VZAudioDeviceConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZAudioDeviceConfiguration creates a new VZAudioDeviceConfiguration instance.
func NewVZAudioDeviceConfiguration() VZAudioDeviceConfiguration {
	return getVZAudioDeviceConfigurationClass().New()
}




