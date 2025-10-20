// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [VZVirtioSoundDeviceInputStreamConfiguration] class.
var (
	VZVirtioSoundDeviceInputStreamConfigurationClass     _VZVirtioSoundDeviceInputStreamConfigurationClass
	VZVirtioSoundDeviceInputStreamConfigurationClassOnce sync.Once
)

func getVZVirtioSoundDeviceInputStreamConfigurationClass() _VZVirtioSoundDeviceInputStreamConfigurationClass {
	VZVirtioSoundDeviceInputStreamConfigurationClassOnce.Do(func() {
		VZVirtioSoundDeviceInputStreamConfigurationClass = _VZVirtioSoundDeviceInputStreamConfigurationClass{objc.GetClass("VZVirtioSoundDeviceInputStreamConfiguration")}
	})
	return VZVirtioSoundDeviceInputStreamConfigurationClass
}

type _VZVirtioSoundDeviceInputStreamConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [VZVirtioSoundDeviceInputStreamConfiguration] class.
type IVZVirtioSoundDeviceInputStreamConfiguration interface {
	IVZVirtioSoundDeviceStreamConfiguration
}

// A PCM stream of input audio data, such as from a microphone.
//
// This device represents a PCM stream of audio data. Don’t instantiate directly. Instead, use one of its subclasses such as or .
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioSoundDeviceInputStreamConfiguration
type VZVirtioSoundDeviceInputStreamConfiguration struct {
	VZVirtioSoundDeviceStreamConfiguration
}

// VZVirtioSoundDeviceInputStreamConfigurationFrom constructs a [VZVirtioSoundDeviceInputStreamConfiguration] from an unsafe.Pointer.
//
// A PCM stream of input audio data, such as from a microphone.
func VZVirtioSoundDeviceInputStreamConfigurationFrom(ptr unsafe.Pointer) VZVirtioSoundDeviceInputStreamConfiguration {
	return VZVirtioSoundDeviceInputStreamConfiguration{
		VZVirtioSoundDeviceStreamConfiguration: VZVirtioSoundDeviceStreamConfigurationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (vc _VZVirtioSoundDeviceInputStreamConfigurationClass) Alloc() VZVirtioSoundDeviceInputStreamConfiguration {
	rv := objc.Send[VZVirtioSoundDeviceInputStreamConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZVirtioSoundDeviceInputStreamConfigurationClass) New() VZVirtioSoundDeviceInputStreamConfiguration {
	rv := objc.Send[VZVirtioSoundDeviceInputStreamConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZVirtioSoundDeviceInputStreamConfiguration) Init() VZVirtioSoundDeviceInputStreamConfiguration {
	rv := objc.Send[VZVirtioSoundDeviceInputStreamConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZVirtioSoundDeviceInputStreamConfiguration) Autorelease() VZVirtioSoundDeviceInputStreamConfiguration {
	rv := objc.Send[VZVirtioSoundDeviceInputStreamConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioSoundDeviceInputStreamConfiguration creates a new VZVirtioSoundDeviceInputStreamConfiguration instance.
func NewVZVirtioSoundDeviceInputStreamConfiguration() VZVirtioSoundDeviceInputStreamConfiguration {
	return getVZVirtioSoundDeviceInputStreamConfigurationClass().New()
}




