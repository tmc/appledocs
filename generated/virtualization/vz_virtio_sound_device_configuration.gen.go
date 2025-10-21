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
}

// An object that defines a Virtio sound device configuration.
//
// Use a object to configure an audio device for your VM. After creating this object, assign appropriate values to the array property which defines the behaviors of the underlying audio streams for this audio device. After creating and configuring a object, assign it to the property of your VM’s configuration.
//
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



// The list of audio devices.
//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/audiodevices
func (v_ VZVirtioSoundDeviceConfiguration) AudioDevices() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("audioDevices"))
	return rv
}


// SetAudioDevices sets the value of the audioDevices property.
// The list of audio devices.

//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/audiodevices
func (v_ VZVirtioSoundDeviceConfiguration) SetAudioDevices(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAudioDevices:"), value)
}

// List of audio streams exposed by this device.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioSoundDeviceConfiguration/streams
func (v_ VZVirtioSoundDeviceConfiguration) Streams() []VZVirtioSoundDeviceStreamConfiguration {
	rv := objc.Send[[]VZVirtioSoundDeviceStreamConfiguration](v_.ID, objc.Sel("streams"))
	return rv
}


// SetStreams sets the value of the streams property.
// List of audio streams exposed by this device.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioSoundDeviceConfiguration/streams
func (v_ VZVirtioSoundDeviceConfiguration) SetStreams(value []VZVirtioSoundDeviceStreamConfiguration) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](v_.ID, objc.Sel("setStreams:"), nsArray)
}


