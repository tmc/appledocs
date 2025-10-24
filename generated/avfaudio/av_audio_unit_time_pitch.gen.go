// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVAudioUnitTimePitch */


/* debug [class_header]: Header for AVAudioUnitTimePitch */
// The class instance for the [AudioUnitTimePitch] class.
var (
	AudioUnitTimePitchClass     _AudioUnitTimePitchClass
	AudioUnitTimePitchClassOnce sync.Once
)

func getAudioUnitTimePitchClass() _AudioUnitTimePitchClass {
	AudioUnitTimePitchClassOnce.Do(func() {
		AudioUnitTimePitchClass = _AudioUnitTimePitchClass{objc.GetClass("AVAudioUnitTimePitch")}
	})
	return AudioUnitTimePitchClass
}

type _AudioUnitTimePitchClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AudioUnitTimePitch */
// An interface definition for the [AudioUnitTimePitch] class.
type IAudioUnitTimePitch interface {
	IAudioUnitTimeEffect
	
/* debug [class_interface_properties]: Properties for AudioUnitTimePitch */
	// properties:
	Overlap() float32
	SetOverlap(value float32)
	Pitch() float32
	SetPitch(value float32)
	Rate() float32
	SetRate(value float32)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AudioUnitTimePitch */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AudioUnitTimePitch */
// Alloc allocates a new instance without initialization.
func (ac _AudioUnitTimePitchClass) Alloc() AudioUnitTimePitch {
	rv := objc.Send[AudioUnitTimePitch](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioUnitTimePitchClass) New() AudioUnitTimePitch {
	rv := objc.Send[AudioUnitTimePitch](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioUnitTimePitch) Init() AudioUnitTimePitch {
	rv := objc.Send[AudioUnitTimePitch](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioUnitTimePitch) Autorelease() AudioUnitTimePitch {
	rv := objc.Send[AudioUnitTimePitch](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioUnitTimePitch creates a new AudioUnitTimePitch instance.
func NewAudioUnitTimePitch() AudioUnitTimePitch {
	return getAudioUnitTimePitchClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AudioUnitTimePitch */
// An object that provides a good-quality playback rate and pitch shifting independently of each other.


// An object that provides a good-quality playback rate and pitch shifting independently of each other.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitTimePitch
type AudioUnitTimePitch struct {
	AudioUnitTimeEffect
}

// AudioUnitTimePitchFrom constructs a [AudioUnitTimePitch] from an unsafe.Pointer.
//
// An object that provides a good-quality playback rate and pitch shifting independently of each other.
func AudioUnitTimePitchFrom(ptr unsafe.Pointer) AudioUnitTimePitch {
	return AudioUnitTimePitch{
		AudioUnitTimeEffect: AudioUnitTimeEffectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AudioUnitTimePitch *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AudioUnitTimePitch */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AudioUnitTimePitch */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AudioUnitTimePitch */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AudioUnitTimePitch */

// The amount of overlap between segments of the input audio signal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitTimePitch/overlap
func (a_ AudioUnitTimePitch) Overlap() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("overlap"))
	return rv
}/* debug [instance_properties/getter]: overlap */


// The amount of overlap between segments of the input audio signal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitTimePitch/overlap
func (a_ AudioUnitTimePitch) SetOverlap(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOverlap:"), value)
}/* debug [instance_properties/setter]: overlap */


// The amount to use to pitch shift the input signal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitTimePitch/pitch
func (a_ AudioUnitTimePitch) Pitch() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("pitch"))
	return rv
}/* debug [instance_properties/getter]: pitch */


// The amount to use to pitch shift the input signal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitTimePitch/pitch
func (a_ AudioUnitTimePitch) SetPitch(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPitch:"), value)
}/* debug [instance_properties/setter]: pitch */


// The playback rate of the input signal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitTimePitch/rate
func (a_ AudioUnitTimePitch) Rate() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("rate"))
	return rv
}/* debug [instance_properties/getter]: rate */


// The playback rate of the input signal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitTimePitch/rate
func (a_ AudioUnitTimePitch) SetRate(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRate:"), value)
}/* debug [instance_properties/setter]: rate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAudioUnitTimePitch */



