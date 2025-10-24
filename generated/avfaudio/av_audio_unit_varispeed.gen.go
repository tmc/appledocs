// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVAudioUnitVarispeed */


/* debug [class_header]: Header for AVAudioUnitVarispeed */
// The class instance for the [AudioUnitVarispeed] class.
var (
	AudioUnitVarispeedClass     _AudioUnitVarispeedClass
	AudioUnitVarispeedClassOnce sync.Once
)

func getAudioUnitVarispeedClass() _AudioUnitVarispeedClass {
	AudioUnitVarispeedClassOnce.Do(func() {
		AudioUnitVarispeedClass = _AudioUnitVarispeedClass{objc.GetClass("AVAudioUnitVarispeed")}
	})
	return AudioUnitVarispeedClass
}

type _AudioUnitVarispeedClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AudioUnitVarispeed */
// An interface definition for the [AudioUnitVarispeed] class.
type IAudioUnitVarispeed interface {
	IAudioUnitTimeEffect
	
/* debug [class_interface_properties]: Properties for AudioUnitVarispeed */
	// properties:
	Rate() float32
	SetRate(value float32)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AudioUnitVarispeed */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AudioUnitVarispeed */
// Alloc allocates a new instance without initialization.
func (ac _AudioUnitVarispeedClass) Alloc() AudioUnitVarispeed {
	rv := objc.Send[AudioUnitVarispeed](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioUnitVarispeedClass) New() AudioUnitVarispeed {
	rv := objc.Send[AudioUnitVarispeed](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioUnitVarispeed) Init() AudioUnitVarispeed {
	rv := objc.Send[AudioUnitVarispeed](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioUnitVarispeed) Autorelease() AudioUnitVarispeed {
	rv := objc.Send[AudioUnitVarispeed](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioUnitVarispeed creates a new AudioUnitVarispeed instance.
func NewAudioUnitVarispeed() AudioUnitVarispeed {
	return getAudioUnitVarispeedClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AudioUnitVarispeed */
// An object that allows control of the playback rate.


// An object that allows control of the playback rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitVarispeed
type AudioUnitVarispeed struct {
	AudioUnitTimeEffect
}

// AudioUnitVarispeedFrom constructs a [AudioUnitVarispeed] from an unsafe.Pointer.
//
// An object that allows control of the playback rate.
func AudioUnitVarispeedFrom(ptr unsafe.Pointer) AudioUnitVarispeed {
	return AudioUnitVarispeed{
		AudioUnitTimeEffect: AudioUnitTimeEffectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AudioUnitVarispeed *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AudioUnitVarispeed */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AudioUnitVarispeed */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AudioUnitVarispeed */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AudioUnitVarispeed */

// The audio playback rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitVarispeed/rate
func (a_ AudioUnitVarispeed) Rate() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("rate"))
	return rv
}/* debug [instance_properties/getter]: rate */


// The audio playback rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitVarispeed/rate
func (a_ AudioUnitVarispeed) SetRate(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRate:"), value)
}/* debug [instance_properties/setter]: rate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAudioUnitVarispeed */



