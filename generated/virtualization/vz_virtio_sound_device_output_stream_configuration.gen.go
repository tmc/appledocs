// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [VZVirtioSoundDeviceOutputStreamConfiguration] class.
var (
	VZVirtioSoundDeviceOutputStreamConfigurationClass     _VZVirtioSoundDeviceOutputStreamConfigurationClass
	VZVirtioSoundDeviceOutputStreamConfigurationClassOnce sync.Once
)

func getVZVirtioSoundDeviceOutputStreamConfigurationClass() _VZVirtioSoundDeviceOutputStreamConfigurationClass {
	VZVirtioSoundDeviceOutputStreamConfigurationClassOnce.Do(func() {
		VZVirtioSoundDeviceOutputStreamConfigurationClass = _VZVirtioSoundDeviceOutputStreamConfigurationClass{objc.GetClass("VZVirtioSoundDeviceOutputStreamConfiguration")}
	})
	return VZVirtioSoundDeviceOutputStreamConfigurationClass
}

type _VZVirtioSoundDeviceOutputStreamConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [VZVirtioSoundDeviceOutputStreamConfiguration] class.
type IVZVirtioSoundDeviceOutputStreamConfiguration interface {
	IVZVirtioSoundDeviceStreamConfiguration
}

// An object that defines a Virtio sound device output stream configuration.
//
// A PCM stream of output audio data, such as to a speaker.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioSoundDeviceOutputStreamConfiguration
type VZVirtioSoundDeviceOutputStreamConfiguration struct {
	VZVirtioSoundDeviceStreamConfiguration
}

// VZVirtioSoundDeviceOutputStreamConfigurationFrom constructs a [VZVirtioSoundDeviceOutputStreamConfiguration] from an unsafe.Pointer.
//
// An object that defines a Virtio sound device output stream configuration.
func VZVirtioSoundDeviceOutputStreamConfigurationFrom(ptr unsafe.Pointer) VZVirtioSoundDeviceOutputStreamConfiguration {
	return VZVirtioSoundDeviceOutputStreamConfiguration{
		VZVirtioSoundDeviceStreamConfiguration: VZVirtioSoundDeviceStreamConfigurationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (vc _VZVirtioSoundDeviceOutputStreamConfigurationClass) Alloc() VZVirtioSoundDeviceOutputStreamConfiguration {
	rv := objc.Send[VZVirtioSoundDeviceOutputStreamConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZVirtioSoundDeviceOutputStreamConfigurationClass) New() VZVirtioSoundDeviceOutputStreamConfiguration {
	rv := objc.Send[VZVirtioSoundDeviceOutputStreamConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZVirtioSoundDeviceOutputStreamConfiguration) Init() VZVirtioSoundDeviceOutputStreamConfiguration {
	rv := objc.Send[VZVirtioSoundDeviceOutputStreamConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZVirtioSoundDeviceOutputStreamConfiguration) Autorelease() VZVirtioSoundDeviceOutputStreamConfiguration {
	rv := objc.Send[VZVirtioSoundDeviceOutputStreamConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioSoundDeviceOutputStreamConfiguration creates a new VZVirtioSoundDeviceOutputStreamConfiguration instance.
func NewVZVirtioSoundDeviceOutputStreamConfiguration() VZVirtioSoundDeviceOutputStreamConfiguration {
	return getVZVirtioSoundDeviceOutputStreamConfigurationClass().New()
}



// An audio stream sink that defines how the host handles audio data produced by the guest.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioSoundDeviceOutputStreamConfiguration/sink
func (v_ VZVirtioSoundDeviceOutputStreamConfiguration) Sink() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("sink"))
	return rv
}


// SetSink sets the value of the sink property.
// An audio stream sink that defines how the host handles audio data produced by the guest.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioSoundDeviceOutputStreamConfiguration/sink
func (v_ VZVirtioSoundDeviceOutputStreamConfiguration) SetSink(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setSink:"), value)
}

