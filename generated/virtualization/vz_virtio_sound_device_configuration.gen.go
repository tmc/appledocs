// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [VZVirtioSoundDeviceConfiguration] class.
var (
	VZVirtioSoundDeviceConfigurationClass     _VZVirtioSoundDeviceConfigurationClass
	VZVirtioSoundDeviceConfigurationClassOnce sync.Once
)

func getVZVirtioSoundDeviceConfigurationClass() _VZVirtioSoundDeviceConfigurationClass {
	VZVirtioSoundDeviceConfigurationClassOnce.Do(func() {
		VZVirtioSoundDeviceConfigurationClass = _VZVirtioSoundDeviceConfigurationClass{objc.GetClass("VZVirtioSoundDeviceConfiguration")}
	})
	return VZVirtioSoundDeviceConfigurationClass
}

type _VZVirtioSoundDeviceConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [VZVirtioSoundDeviceConfiguration] class.
type IVZVirtioSoundDeviceConfiguration interface {
	IVZAudioDeviceConfiguration
	// properties:
	Streams() IVZVirtioSoundDeviceStreamConfiguration
	SetStreams(value IVZVirtioSoundDeviceStreamConfiguration)
	AudioDevices() IVZAudioDeviceConfiguration
	SetAudioDevices(value IVZAudioDeviceConfiguration)
	// methods:
}

// An object that defines a Virtio sound device configuration.
//
// Use a object to configure an audio device for your VM. After creating this object, assign appropriate values to the array property which defines the behaviors of the underlying audio streams for this audio device. After creating and configuring a object, assign it to the property of your VM’s configuration.


// An object that defines a Virtio sound device configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioSoundDeviceConfiguration
type VZVirtioSoundDeviceConfiguration struct {
	VZAudioDeviceConfiguration
}

// VZVirtioSoundDeviceConfigurationFrom constructs a [VZVirtioSoundDeviceConfiguration] from an unsafe.Pointer.
//
// An object that defines a Virtio sound device configuration.
func VZVirtioSoundDeviceConfigurationFrom(ptr unsafe.Pointer) VZVirtioSoundDeviceConfiguration {
	return VZVirtioSoundDeviceConfiguration{
		VZAudioDeviceConfiguration: VZAudioDeviceConfigurationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (vc _VZVirtioSoundDeviceConfigurationClass) Alloc() VZVirtioSoundDeviceConfiguration {
	rv := objc.Send[VZVirtioSoundDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZVirtioSoundDeviceConfigurationClass) New() VZVirtioSoundDeviceConfiguration {
	rv := objc.Send[VZVirtioSoundDeviceConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZVirtioSoundDeviceConfiguration) Init() VZVirtioSoundDeviceConfiguration {
	rv := objc.Send[VZVirtioSoundDeviceConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZVirtioSoundDeviceConfiguration) Autorelease() VZVirtioSoundDeviceConfiguration {
	rv := objc.Send[VZVirtioSoundDeviceConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioSoundDeviceConfiguration creates a new VZVirtioSoundDeviceConfiguration instance.
func NewVZVirtioSoundDeviceConfiguration() VZVirtioSoundDeviceConfiguration {
	return getVZVirtioSoundDeviceConfigurationClass().New()
}



// List of audio streams exposed by this device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtiosounddeviceconfiguration/streams
func (v_ VZVirtioSoundDeviceConfiguration) Streams() IVZVirtioSoundDeviceStreamConfiguration {
	rv := objc.Send[VZVirtioSoundDeviceStreamConfiguration](v_.ID, objc.Sel("streams"))
	return rv
}


// List of audio streams exposed by this device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtiosounddeviceconfiguration/streams
func (v_ VZVirtioSoundDeviceConfiguration) SetStreams(value IVZVirtioSoundDeviceStreamConfiguration) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setStreams:"), value)
}


// The list of audio devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/audiodevices
func (v_ VZVirtioSoundDeviceConfiguration) AudioDevices() IVZAudioDeviceConfiguration {
	rv := objc.Send[VZAudioDeviceConfiguration](v_.ID, objc.Sel("audioDevices"))
	return rv
}


// The list of audio devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/audiodevices
func (v_ VZVirtioSoundDeviceConfiguration) SetAudioDevices(value IVZAudioDeviceConfiguration) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAudioDevices:"), value)
}



