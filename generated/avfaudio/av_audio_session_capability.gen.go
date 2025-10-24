// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAudioSessionCapability */


/* debug [class_header]: Header for AVAudioSessionCapability */
// The class instance for the [AudioSessionCapability] class.
var (
	AudioSessionCapabilityClass     _AudioSessionCapabilityClass
	AudioSessionCapabilityClassOnce sync.Once
)

func getAudioSessionCapabilityClass() _AudioSessionCapabilityClass {
	AudioSessionCapabilityClassOnce.Do(func() {
		AudioSessionCapabilityClass = _AudioSessionCapabilityClass{objc.GetClass("AVAudioSessionCapability")}
	})
	return AudioSessionCapabilityClass
}

type _AudioSessionCapabilityClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AudioSessionCapability */
// An interface definition for the [AudioSessionCapability] class.
type IAudioSessionCapability interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AudioSessionCapability */
	// properties:
	Enabled() bool
	Supported() bool
	IsEnabled() bool
	SetIsEnabled(value bool)
	IsSupported() bool
	SetIsSupported(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AudioSessionCapability */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AudioSessionCapability */
// Alloc allocates a new instance without initialization.
func (ac _AudioSessionCapabilityClass) Alloc() AudioSessionCapability {
	rv := objc.Send[AudioSessionCapability](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioSessionCapabilityClass) New() AudioSessionCapability {
	rv := objc.Send[AudioSessionCapability](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioSessionCapability) Init() AudioSessionCapability {
	rv := objc.Send[AudioSessionCapability](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioSessionCapability) Autorelease() AudioSessionCapability {
	rv := objc.Send[AudioSessionCapability](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioSessionCapability creates a new AudioSessionCapability instance.
func NewAudioSessionCapability() AudioSessionCapability {
	return getAudioSessionCapabilityClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AudioSessionCapability */
// Describes whether a specific capability is supported and if that capability is currently enabled


// Describes whether a specific capability is supported and if that capability is currently enabled
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionCapability
type AudioSessionCapability struct {
	objectivec.Object
}

// AudioSessionCapabilityFrom constructs a [AudioSessionCapability] from an unsafe.Pointer.
//
// Describes whether a specific capability is supported and if that capability is currently enabled
func AudioSessionCapabilityFrom(ptr unsafe.Pointer) AudioSessionCapability {
	return AudioSessionCapability{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AudioSessionCapability *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AudioSessionCapability */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AudioSessionCapability */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AudioSessionCapability */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AudioSessionCapability */

// A Boolean value that indicates whether the capability is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionCapability/isEnabled
func (a_ AudioSessionCapability) Enabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("enabled"))
	return rv
}/* debug [instance_properties/getter]: enabled */


// A Boolean value that indicates whether the capability is supported.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionCapability/isSupported
func (a_ AudioSessionCapability) Supported() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("supported"))
	return rv
}/* debug [instance_properties/getter]: supported */


// A Boolean value that indicates whether the capability is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessioncapability/isenabled
func (a_ AudioSessionCapability) IsEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isEnabled"))
	return rv
}/* debug [instance_properties/getter]: isEnabled */


// A Boolean value that indicates whether the capability is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessioncapability/isenabled
func (a_ AudioSessionCapability) SetIsEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsEnabled:"), value)
}/* debug [instance_properties/setter]: isEnabled */


// A Boolean value that indicates whether the capability is supported.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessioncapability/issupported
func (a_ AudioSessionCapability) IsSupported() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isSupported"))
	return rv
}/* debug [instance_properties/getter]: isSupported */


// A Boolean value that indicates whether the capability is supported.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessioncapability/issupported
func (a_ AudioSessionCapability) SetIsSupported(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsSupported:"), value)
}/* debug [instance_properties/setter]: isSupported */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAudioSessionCapability */



