// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVAudioUnitDelay */


/* debug [class_header]: Header for AVAudioUnitDelay */
// The class instance for the [AudioUnitDelay] class.
var (
	AudioUnitDelayClass     _AudioUnitDelayClass
	AudioUnitDelayClassOnce sync.Once
)

func getAudioUnitDelayClass() _AudioUnitDelayClass {
	AudioUnitDelayClassOnce.Do(func() {
		AudioUnitDelayClass = _AudioUnitDelayClass{objc.GetClass("AVAudioUnitDelay")}
	})
	return AudioUnitDelayClass
}

type _AudioUnitDelayClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AudioUnitDelay */
// An interface definition for the [AudioUnitDelay] class.
type IAudioUnitDelay interface {
	IAudioUnitEffect
	
/* debug [class_interface_properties]: Properties for AudioUnitDelay */
	// properties:
	DelayTime() float64
	SetDelayTime(value float64)
	Feedback() float32
	SetFeedback(value float32)
	LowPassCutoff() float32
	SetLowPassCutoff(value float32)
	WetDryMix() float32
	SetWetDryMix(value float32)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AudioUnitDelay */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AudioUnitDelay */
// Alloc allocates a new instance without initialization.
func (ac _AudioUnitDelayClass) Alloc() AudioUnitDelay {
	rv := objc.Send[AudioUnitDelay](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioUnitDelayClass) New() AudioUnitDelay {
	rv := objc.Send[AudioUnitDelay](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioUnitDelay) Init() AudioUnitDelay {
	rv := objc.Send[AudioUnitDelay](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioUnitDelay) Autorelease() AudioUnitDelay {
	rv := objc.Send[AudioUnitDelay](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioUnitDelay creates a new AudioUnitDelay instance.
func NewAudioUnitDelay() AudioUnitDelay {
	return getAudioUnitDelayClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AudioUnitDelay */
// An object that implements a delay effect.
//
// A delay unit delays the input signal by the specified time interval and then blends it with the input signal. You can also control the amount of high-frequency roll-off to simulate the effect of a tape delay.


// An object that implements a delay effect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDelay
type AudioUnitDelay struct {
	AudioUnitEffect
}

// AudioUnitDelayFrom constructs a [AudioUnitDelay] from an unsafe.Pointer.
//
// An object that implements a delay effect.
func AudioUnitDelayFrom(ptr unsafe.Pointer) AudioUnitDelay {
	return AudioUnitDelay{
		AudioUnitEffect: AudioUnitEffectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AudioUnitDelay *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AudioUnitDelay */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AudioUnitDelay */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AudioUnitDelay */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AudioUnitDelay */

// The time for the input signal to reach the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDelay/delayTime
func (a_ AudioUnitDelay) DelayTime() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("delayTime"))
	return rv
}/* debug [instance_properties/getter]: delayTime */


// The time for the input signal to reach the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDelay/delayTime
func (a_ AudioUnitDelay) SetDelayTime(value float64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDelayTime:"), value)
}/* debug [instance_properties/setter]: delayTime */


// The amount of the output signal that feeds back into the delay line.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDelay/feedback
func (a_ AudioUnitDelay) Feedback() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("feedback"))
	return rv
}/* debug [instance_properties/getter]: feedback */


// The amount of the output signal that feeds back into the delay line.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDelay/feedback
func (a_ AudioUnitDelay) SetFeedback(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFeedback:"), value)
}/* debug [instance_properties/setter]: feedback */


// The cutoff frequency above which high frequency content rolls off, in hertz.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDelay/lowPassCutoff
func (a_ AudioUnitDelay) LowPassCutoff() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("lowPassCutoff"))
	return rv
}/* debug [instance_properties/getter]: lowPassCutoff */


// The cutoff frequency above which high frequency content rolls off, in hertz.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDelay/lowPassCutoff
func (a_ AudioUnitDelay) SetLowPassCutoff(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLowPassCutoff:"), value)
}/* debug [instance_properties/setter]: lowPassCutoff */


// The blend of the wet and dry signals.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDelay/wetDryMix
func (a_ AudioUnitDelay) WetDryMix() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("wetDryMix"))
	return rv
}/* debug [instance_properties/getter]: wetDryMix */


// The blend of the wet and dry signals.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDelay/wetDryMix
func (a_ AudioUnitDelay) SetWetDryMix(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setWetDryMix:"), value)
}/* debug [instance_properties/setter]: wetDryMix */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAudioUnitDelay */



