// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAudioUnitEQ */


/* debug [class_header]: Header for AVAudioUnitEQ */
// The class instance for the [AudioUnitEQ] class.
var (
	AudioUnitEQClass     _AudioUnitEQClass
	AudioUnitEQClassOnce sync.Once
)

func getAudioUnitEQClass() _AudioUnitEQClass {
	AudioUnitEQClassOnce.Do(func() {
		AudioUnitEQClass = _AudioUnitEQClass{objc.GetClass("AVAudioUnitEQ")}
	})
	return AudioUnitEQClass
}

type _AudioUnitEQClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AudioUnitEQ */
// An interface definition for the [AudioUnitEQ] class.
type IAudioUnitEQ interface {
	IAudioUnitEffect
	
/* debug [class_interface_properties]: Properties for AudioUnitEQ */
	// properties:
	Bands() []AudioUnitEQFilterParameters
	GlobalGain() float32
	SetGlobalGain(value float32)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AudioUnitEQ */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AudioUnitEQ */
// Alloc allocates a new instance without initialization.
func (ac _AudioUnitEQClass) Alloc() AudioUnitEQ {
	rv := objc.Send[AudioUnitEQ](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioUnitEQClass) New() AudioUnitEQ {
	rv := objc.Send[AudioUnitEQ](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioUnitEQ) Init() AudioUnitEQ {
	rv := objc.Send[AudioUnitEQ](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioUnitEQ) Autorelease() AudioUnitEQ {
	rv := objc.Send[AudioUnitEQ](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioUnitEQ creates a new AudioUnitEQ instance.
func NewAudioUnitEQ() AudioUnitEQ {
	return getAudioUnitEQClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AudioUnitEQ */
// An object that implements a multiband equalizer.
//
// The class encapsulates the filter parameters that the property array returns.


// An object that implements a multiband equalizer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEQ
type AudioUnitEQ struct {
	AudioUnitEffect
}

// AudioUnitEQFrom constructs a [AudioUnitEQ] from an unsafe.Pointer.
//
// An object that implements a multiband equalizer.
func AudioUnitEQFrom(ptr unsafe.Pointer) AudioUnitEQ {
	return AudioUnitEQ{
		AudioUnitEffect: AudioUnitEffectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AudioUnitEQ */

// Creates an audio unit equalizer object with the specified number of bands.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEQ/init(numberOfBands:)
func NewAudioUnitEQWithNumberOfBands(numberOfBands uint) AudioUnitEQ {
	instance := getAudioUnitEQClass().Alloc()
	rv := objc.Send[AudioUnitEQ](instance.ID, objc.Sel("initWithNumberOfBands:"), numberOfBands)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAudioUnitEQWithNumberOfBands */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AudioUnitEQ */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AudioUnitEQ */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AudioUnitEQ */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AudioUnitEQ */

// An array of equalizer filter parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEQ/bands
func (a_ AudioUnitEQ) Bands() []AudioUnitEQFilterParameters {
	rv := objc.Send[[]AudioUnitEQFilterParameters](a_.ID, objc.Sel("bands"))
	return rv
}/* debug [instance_properties/getter]: bands */


// The overall gain adjustment that the audio unit applies to the signal, in decibels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEQ/globalGain
func (a_ AudioUnitEQ) GlobalGain() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("globalGain"))
	return rv
}/* debug [instance_properties/getter]: globalGain */


// The overall gain adjustment that the audio unit applies to the signal, in decibels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEQ/globalGain
func (a_ AudioUnitEQ) SetGlobalGain(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setGlobalGain:"), value)
}/* debug [instance_properties/setter]: globalGain */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAudioUnitEQ */


