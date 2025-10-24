// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVAudioUnitReverb */


/* debug [class_header]: Header for AVAudioUnitReverb */
// The class instance for the [AudioUnitReverb] class.
var (
	AudioUnitReverbClass     _AudioUnitReverbClass
	AudioUnitReverbClassOnce sync.Once
)

func getAudioUnitReverbClass() _AudioUnitReverbClass {
	AudioUnitReverbClassOnce.Do(func() {
		AudioUnitReverbClass = _AudioUnitReverbClass{objc.GetClass("AVAudioUnitReverb")}
	})
	return AudioUnitReverbClass
}

type _AudioUnitReverbClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AudioUnitReverb */
// An interface definition for the [AudioUnitReverb] class.
type IAudioUnitReverb interface {
	IAudioUnitEffect
	
/* debug [class_interface_properties]: Properties for AudioUnitReverb */
	// properties:
	WetDryMix() float32
	SetWetDryMix(value float32)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AudioUnitReverb */
	// methods:
	LoadFactoryPreset(preset AudioUnitReverbPreset)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AudioUnitReverb */
// Alloc allocates a new instance without initialization.
func (ac _AudioUnitReverbClass) Alloc() AudioUnitReverb {
	rv := objc.Send[AudioUnitReverb](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioUnitReverbClass) New() AudioUnitReverb {
	rv := objc.Send[AudioUnitReverb](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioUnitReverb) Init() AudioUnitReverb {
	rv := objc.Send[AudioUnitReverb](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioUnitReverb) Autorelease() AudioUnitReverb {
	rv := objc.Send[AudioUnitReverb](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioUnitReverb creates a new AudioUnitReverb instance.
func NewAudioUnitReverb() AudioUnitReverb {
	return getAudioUnitReverbClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AudioUnitReverb */
// An object that implements a reverb effect.
//
// A reverb simulates the acoustic characteristics of a particular environment. Use the different presets to simulate a particular space and blend it in with the original signal using the property.


// An object that implements a reverb effect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitReverb
type AudioUnitReverb struct {
	AudioUnitEffect
}

// AudioUnitReverbFrom constructs a [AudioUnitReverb] from an unsafe.Pointer.
//
// An object that implements a reverb effect.
func AudioUnitReverbFrom(ptr unsafe.Pointer) AudioUnitReverb {
	return AudioUnitReverb{
		AudioUnitEffect: AudioUnitEffectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AudioUnitReverb *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AudioUnitReverb */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AudioUnitReverb */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AudioUnitReverb */

// Configures the audio unit as a reverb preset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitReverb/loadFactoryPreset(_:)
func (a_ AudioUnitReverb) LoadFactoryPreset(preset AudioUnitReverbPreset) {
	objc.Send[objc.ID](a_.ID, objc.Sel("loadFactoryPreset:"), preset)
}/* debug [instance_methods/method]: LoadFactoryPreset */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AudioUnitReverb */

// The blend of the wet and dry signals.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitReverb/wetDryMix
func (a_ AudioUnitReverb) WetDryMix() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("wetDryMix"))
	return rv
}/* debug [instance_properties/getter]: wetDryMix */


// The blend of the wet and dry signals.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitReverb/wetDryMix
func (a_ AudioUnitReverb) SetWetDryMix(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setWetDryMix:"), value)
}/* debug [instance_properties/setter]: wetDryMix */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAudioUnitReverb */



