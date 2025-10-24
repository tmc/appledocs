// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VZVirtioSoundDeviceConfiguration */

/* debug [class_header]: Header for VZVirtioSoundDeviceConfiguration */
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

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for VZVirtioSoundDeviceConfiguration */
// An interface definition for the [VZVirtioSoundDeviceConfiguration] class.
type IVZVirtioSoundDeviceConfiguration interface {
	IVZAudioDeviceConfiguration

	/* debug [class_interface_properties]: Properties for VZVirtioSoundDeviceConfiguration */
	// properties:
	Streams() []VZVirtioSoundDeviceStreamConfiguration
	SetStreams(value []VZVirtioSoundDeviceStreamConfiguration)
	AudioDevices() IVZAudioDeviceConfiguration
	SetAudioDevices(value IVZAudioDeviceConfiguration)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for VZVirtioSoundDeviceConfiguration */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for VZVirtioSoundDeviceConfiguration */
// Alloc allocates a new instance without initialization.
func (vc _VZVirtioSoundDeviceConfigurationClass) Alloc() VZVirtioSoundDeviceConfiguration {
	rv := objc.Send[VZVirtioSoundDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for VZVirtioSoundDeviceConfiguration */
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

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for VZVirtioSoundDeviceConfiguration */
/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for VZVirtioSoundDeviceConfiguration */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for VZVirtioSoundDeviceConfiguration */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for VZVirtioSoundDeviceConfiguration */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for VZVirtioSoundDeviceConfiguration */

// List of audio streams exposed by this device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioSoundDeviceConfiguration/streams
func (v_ VZVirtioSoundDeviceConfiguration) Streams() []VZVirtioSoundDeviceStreamConfiguration {
	rv := objc.Send[[]VZVirtioSoundDeviceStreamConfiguration](v_.ID, objc.Sel("streams"))
	return rv
} /* debug [instance_properties/getter]: streams */

// List of audio streams exposed by this device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioSoundDeviceConfiguration/streams
func (v_ VZVirtioSoundDeviceConfiguration) SetStreams(value []VZVirtioSoundDeviceStreamConfiguration) {
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
} /* debug [instance_properties/setter]: streams */

// The list of audio devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/audiodevices
func (v_ VZVirtioSoundDeviceConfiguration) AudioDevices() IVZAudioDeviceConfiguration {
	rv := objc.Send[VZAudioDeviceConfiguration](v_.ID, objc.Sel("audioDevices"))
	return rv
} /* debug [instance_properties/getter]: audioDevices */

// The list of audio devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/audiodevices
func (v_ VZVirtioSoundDeviceConfiguration) SetAudioDevices(value IVZAudioDeviceConfiguration) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAudioDevices:"), value)
} /* debug [instance_properties/setter]: audioDevices */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class VZVirtioSoundDeviceConfiguration */
