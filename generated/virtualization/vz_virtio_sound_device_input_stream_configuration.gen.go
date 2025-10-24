// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VZVirtioSoundDeviceInputStreamConfiguration */

/* debug [class_header]: Header for VZVirtioSoundDeviceInputStreamConfiguration */
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

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for VZVirtioSoundDeviceInputStreamConfiguration */
// An interface definition for the [VZVirtioSoundDeviceInputStreamConfiguration] class.
type IVZVirtioSoundDeviceInputStreamConfiguration interface {
	IVZVirtioSoundDeviceStreamConfiguration

	/* debug [class_interface_properties]: Properties for VZVirtioSoundDeviceInputStreamConfiguration */
	// properties:
	Source() IVZAudioInputStreamSource
	SetSource(value IVZAudioInputStreamSource)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for VZVirtioSoundDeviceInputStreamConfiguration */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for VZVirtioSoundDeviceInputStreamConfiguration */
// Alloc allocates a new instance without initialization.
func (vc _VZVirtioSoundDeviceInputStreamConfigurationClass) Alloc() VZVirtioSoundDeviceInputStreamConfiguration {
	rv := objc.Send[VZVirtioSoundDeviceInputStreamConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for VZVirtioSoundDeviceInputStreamConfiguration */
// A PCM stream of input audio data, such as from a microphone.
//
// This device represents a PCM stream of audio data. Don’t instantiate directly. Instead, use one of its subclasses such as or .

// A PCM stream of input audio data, such as from a microphone.
//
// [Full Topic]
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

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for VZVirtioSoundDeviceInputStreamConfiguration */
/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for VZVirtioSoundDeviceInputStreamConfiguration */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for VZVirtioSoundDeviceInputStreamConfiguration */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for VZVirtioSoundDeviceInputStreamConfiguration */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for VZVirtioSoundDeviceInputStreamConfiguration */

// An audio stream source that defines how the host supplies audio data for the guest.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioSoundDeviceInputStreamConfiguration/source
func (v_ VZVirtioSoundDeviceInputStreamConfiguration) Source() IVZAudioInputStreamSource {
	rv := objc.Send[VZAudioInputStreamSource](v_.ID, objc.Sel("source"))
	return rv
} /* debug [instance_properties/getter]: source */

// An audio stream source that defines how the host supplies audio data for the guest.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioSoundDeviceInputStreamConfiguration/source
func (v_ VZVirtioSoundDeviceInputStreamConfiguration) SetSource(value IVZAudioInputStreamSource) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setSource:"), value)
} /* debug [instance_properties/setter]: source */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class VZVirtioSoundDeviceInputStreamConfiguration */
