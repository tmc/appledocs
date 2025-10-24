// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVAudioUnitDistortion */


/* debug [class_header]: Header for AVAudioUnitDistortion */
// The class instance for the [AudioUnitDistortion] class.
var (
	AudioUnitDistortionClass     _AudioUnitDistortionClass
	AudioUnitDistortionClassOnce sync.Once
)

func getAudioUnitDistortionClass() _AudioUnitDistortionClass {
	AudioUnitDistortionClassOnce.Do(func() {
		AudioUnitDistortionClass = _AudioUnitDistortionClass{objc.GetClass("AVAudioUnitDistortion")}
	})
	return AudioUnitDistortionClass
}

type _AudioUnitDistortionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AudioUnitDistortion */
// An interface definition for the [AudioUnitDistortion] class.
type IAudioUnitDistortion interface {
	IAudioUnitEffect
	
/* debug [class_interface_properties]: Properties for AudioUnitDistortion */
	// properties:
	PreGain() float32
	SetPreGain(value float32)
	WetDryMix() float32
	SetWetDryMix(value float32)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AudioUnitDistortion */
	// methods:
	LoadFactoryPreset(preset AudioUnitDistortionPreset)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AudioUnitDistortion */
// Alloc allocates a new instance without initialization.
func (ac _AudioUnitDistortionClass) Alloc() AudioUnitDistortion {
	rv := objc.Send[AudioUnitDistortion](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioUnitDistortionClass) New() AudioUnitDistortion {
	rv := objc.Send[AudioUnitDistortion](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioUnitDistortion) Init() AudioUnitDistortion {
	rv := objc.Send[AudioUnitDistortion](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioUnitDistortion) Autorelease() AudioUnitDistortion {
	rv := objc.Send[AudioUnitDistortion](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioUnitDistortion creates a new AudioUnitDistortion instance.
func NewAudioUnitDistortion() AudioUnitDistortion {
	return getAudioUnitDistortionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AudioUnitDistortion */
// An object that implements a multistage distortion effect.


// An object that implements a multistage distortion effect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDistortion
type AudioUnitDistortion struct {
	AudioUnitEffect
}

// AudioUnitDistortionFrom constructs a [AudioUnitDistortion] from an unsafe.Pointer.
//
// An object that implements a multistage distortion effect.
func AudioUnitDistortionFrom(ptr unsafe.Pointer) AudioUnitDistortion {
	return AudioUnitDistortion{
		AudioUnitEffect: AudioUnitEffectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AudioUnitDistortion *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AudioUnitDistortion */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AudioUnitDistortion */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AudioUnitDistortion */

// Configures the audio distortion unit by loading a distortion preset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDistortion/loadFactoryPreset(_:)
func (a_ AudioUnitDistortion) LoadFactoryPreset(preset AudioUnitDistortionPreset) {
	objc.Send[objc.ID](a_.ID, objc.Sel("loadFactoryPreset:"), preset)
}/* debug [instance_methods/method]: LoadFactoryPreset */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AudioUnitDistortion */

// The gain that the audio unit applies to the signal before distortion, in decibels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDistortion/preGain
func (a_ AudioUnitDistortion) PreGain() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("preGain"))
	return rv
}/* debug [instance_properties/getter]: preGain */


// The gain that the audio unit applies to the signal before distortion, in decibels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDistortion/preGain
func (a_ AudioUnitDistortion) SetPreGain(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPreGain:"), value)
}/* debug [instance_properties/setter]: preGain */


// The blend of the distorted and dry signals.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDistortion/wetDryMix
func (a_ AudioUnitDistortion) WetDryMix() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("wetDryMix"))
	return rv
}/* debug [instance_properties/getter]: wetDryMix */


// The blend of the distorted and dry signals.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDistortion/wetDryMix
func (a_ AudioUnitDistortion) SetWetDryMix(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setWetDryMix:"), value)
}/* debug [instance_properties/setter]: wetDryMix */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAudioUnitDistortion */



