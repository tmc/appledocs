// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/audiotoolbox"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAudioUnitTimeEffect */


/* debug [class_header]: Header for AVAudioUnitTimeEffect */
// The class instance for the [AudioUnitTimeEffect] class.
var (
	AudioUnitTimeEffectClass     _AudioUnitTimeEffectClass
	AudioUnitTimeEffectClassOnce sync.Once
)

func getAudioUnitTimeEffectClass() _AudioUnitTimeEffectClass {
	AudioUnitTimeEffectClassOnce.Do(func() {
		AudioUnitTimeEffectClass = _AudioUnitTimeEffectClass{objc.GetClass("AVAudioUnitTimeEffect")}
	})
	return AudioUnitTimeEffectClass
}

type _AudioUnitTimeEffectClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AudioUnitTimeEffect */
// An interface definition for the [AudioUnitTimeEffect] class.
type IAudioUnitTimeEffect interface {
	IAudioUnit
	
/* debug [class_interface_properties]: Properties for AudioUnitTimeEffect */
	// properties:
	Bypass() bool
	SetBypass(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AudioUnitTimeEffect */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AudioUnitTimeEffect */
// Alloc allocates a new instance without initialization.
func (ac _AudioUnitTimeEffectClass) Alloc() AudioUnitTimeEffect {
	rv := objc.Send[AudioUnitTimeEffect](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioUnitTimeEffectClass) New() AudioUnitTimeEffect {
	rv := objc.Send[AudioUnitTimeEffect](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioUnitTimeEffect) Init() AudioUnitTimeEffect {
	rv := objc.Send[AudioUnitTimeEffect](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioUnitTimeEffect) Autorelease() AudioUnitTimeEffect {
	rv := objc.Send[AudioUnitTimeEffect](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioUnitTimeEffect creates a new AudioUnitTimeEffect instance.
func NewAudioUnitTimeEffect() AudioUnitTimeEffect {
	return getAudioUnitTimeEffectClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AudioUnitTimeEffect */
// An object that processes audio in nonreal time.
//
// A time effect audio unit represents an with a type ( . These effects don’t process audio in real time. The class is an example of a time effect unit.


// An object that processes audio in nonreal time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitTimeEffect
type AudioUnitTimeEffect struct {
	AudioUnit
}

// AudioUnitTimeEffectFrom constructs a [AudioUnitTimeEffect] from an unsafe.Pointer.
//
// An object that processes audio in nonreal time.
func AudioUnitTimeEffectFrom(ptr unsafe.Pointer) AudioUnitTimeEffect {
	return AudioUnitTimeEffect{
		AudioUnit: AudioUnitFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AudioUnitTimeEffect */

// Creates a time effect audio unit with the specified description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitTimeEffect/init(audioComponentDescription:)
func NewAudioUnitTimeEffectWithAudioComponentDescription(audioComponentDescription audiotoolbox.AudioComponentDescription) AudioUnitTimeEffect {
	instance := getAudioUnitTimeEffectClass().Alloc()
	rv := objc.Send[AudioUnitTimeEffect](instance.ID, objc.Sel("initWithAudioComponentDescription:"), audioComponentDescription)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAudioUnitTimeEffectWithAudioComponentDescription */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AudioUnitTimeEffect */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AudioUnitTimeEffect */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AudioUnitTimeEffect */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AudioUnitTimeEffect */

// The bypass state of the audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitTimeEffect/bypass
func (a_ AudioUnitTimeEffect) Bypass() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("bypass"))
	return rv
}/* debug [instance_properties/getter]: bypass */


// The bypass state of the audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitTimeEffect/bypass
func (a_ AudioUnitTimeEffect) SetBypass(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setBypass:"), value)
}/* debug [instance_properties/setter]: bypass */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAudioUnitTimeEffect */


