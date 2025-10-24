// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/audiotoolbox"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAudioUnitEffect */


/* debug [class_header]: Header for AVAudioUnitEffect */
// The class instance for the [AudioUnitEffect] class.
var (
	AudioUnitEffectClass     _AudioUnitEffectClass
	AudioUnitEffectClassOnce sync.Once
)

func getAudioUnitEffectClass() _AudioUnitEffectClass {
	AudioUnitEffectClassOnce.Do(func() {
		AudioUnitEffectClass = _AudioUnitEffectClass{objc.GetClass("AVAudioUnitEffect")}
	})
	return AudioUnitEffectClass
}

type _AudioUnitEffectClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AudioUnitEffect */
// An interface definition for the [AudioUnitEffect] class.
type IAudioUnitEffect interface {
	IAudioUnit
	
/* debug [class_interface_properties]: Properties for AudioUnitEffect */
	// properties:
	Bypass() bool
	SetBypass(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AudioUnitEffect */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AudioUnitEffect */
// Alloc allocates a new instance without initialization.
func (ac _AudioUnitEffectClass) Alloc() AudioUnitEffect {
	rv := objc.Send[AudioUnitEffect](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioUnitEffectClass) New() AudioUnitEffect {
	rv := objc.Send[AudioUnitEffect](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioUnitEffect) Init() AudioUnitEffect {
	rv := objc.Send[AudioUnitEffect](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioUnitEffect) Autorelease() AudioUnitEffect {
	rv := objc.Send[AudioUnitEffect](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioUnitEffect creates a new AudioUnitEffect instance.
func NewAudioUnitEffect() AudioUnitEffect {
	return getAudioUnitEffectClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AudioUnitEffect */
// An object that processes audio in real time.
//
// This processing uses of type effect, music effect, panner, remote effect, or remote music effect. These effects run in real time and process some number of audio input samples to produce several audio output samples. A delay unit is an example of an effect unit.


// An object that processes audio in real time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEffect
type AudioUnitEffect struct {
	AudioUnit
}

// AudioUnitEffectFrom constructs a [AudioUnitEffect] from an unsafe.Pointer.
//
// An object that processes audio in real time.
func AudioUnitEffectFrom(ptr unsafe.Pointer) AudioUnitEffect {
	return AudioUnitEffect{
		AudioUnit: AudioUnitFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AudioUnitEffect */

// Creates an audio unit effect object with the specified description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEffect/init(audioComponentDescription:)
func NewAudioUnitEffectWithAudioComponentDescription(audioComponentDescription audiotoolbox.AudioComponentDescription) AudioUnitEffect {
	instance := getAudioUnitEffectClass().Alloc()
	rv := objc.Send[AudioUnitEffect](instance.ID, objc.Sel("initWithAudioComponentDescription:"), audioComponentDescription)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAudioUnitEffectWithAudioComponentDescription */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AudioUnitEffect */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AudioUnitEffect */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AudioUnitEffect */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AudioUnitEffect */

// The bypass state of the audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEffect/bypass
func (a_ AudioUnitEffect) Bypass() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("bypass"))
	return rv
}/* debug [instance_properties/getter]: bypass */


// The bypass state of the audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEffect/bypass
func (a_ AudioUnitEffect) SetBypass(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setBypass:"), value)
}/* debug [instance_properties/setter]: bypass */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAudioUnitEffect */


