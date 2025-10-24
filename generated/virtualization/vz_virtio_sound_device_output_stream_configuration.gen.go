// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZVirtioSoundDeviceOutputStreamConfiguration */


/* debug [class_header]: Header for VZVirtioSoundDeviceOutputStreamConfiguration */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZVirtioSoundDeviceOutputStreamConfiguration */
// An interface definition for the [VZVirtioSoundDeviceOutputStreamConfiguration] class.
type IVZVirtioSoundDeviceOutputStreamConfiguration interface {
	IVZVirtioSoundDeviceStreamConfiguration
	
/* debug [class_interface_properties]: Properties for VZVirtioSoundDeviceOutputStreamConfiguration */
	// properties:
	Sink() IVZAudioOutputStreamSink
	SetSink(value IVZAudioOutputStreamSink)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZVirtioSoundDeviceOutputStreamConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZVirtioSoundDeviceOutputStreamConfiguration */
// Alloc allocates a new instance without initialization.
func (vc _VZVirtioSoundDeviceOutputStreamConfigurationClass) Alloc() VZVirtioSoundDeviceOutputStreamConfiguration {
	rv := objc.Send[VZVirtioSoundDeviceOutputStreamConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZVirtioSoundDeviceOutputStreamConfiguration */
// An object that defines a Virtio sound device output stream configuration.
//
// A PCM stream of output audio data, such as to a speaker.


// An object that defines a Virtio sound device output stream configuration.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZVirtioSoundDeviceOutputStreamConfiguration */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZVirtioSoundDeviceOutputStreamConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZVirtioSoundDeviceOutputStreamConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZVirtioSoundDeviceOutputStreamConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZVirtioSoundDeviceOutputStreamConfiguration */

// An audio stream sink that defines how the host handles audio data produced by the guest.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioSoundDeviceOutputStreamConfiguration/sink
func (v_ VZVirtioSoundDeviceOutputStreamConfiguration) Sink() IVZAudioOutputStreamSink {
	rv := objc.Send[VZAudioOutputStreamSink](v_.ID, objc.Sel("sink"))
	return rv
}/* debug [instance_properties/getter]: sink */


// An audio stream sink that defines how the host handles audio data produced by the guest.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioSoundDeviceOutputStreamConfiguration/sink
func (v_ VZVirtioSoundDeviceOutputStreamConfiguration) SetSink(value IVZAudioOutputStreamSink) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setSink:"), value)
}/* debug [instance_properties/setter]: sink */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZVirtioSoundDeviceOutputStreamConfiguration */


