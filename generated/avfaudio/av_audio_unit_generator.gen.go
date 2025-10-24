// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/audiotoolbox"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAudioUnitGenerator */


/* debug [class_header]: Header for AVAudioUnitGenerator */
// The class instance for the [AudioUnitGenerator] class.
var (
	AudioUnitGeneratorClass     _AudioUnitGeneratorClass
	AudioUnitGeneratorClassOnce sync.Once
)

func getAudioUnitGeneratorClass() _AudioUnitGeneratorClass {
	AudioUnitGeneratorClassOnce.Do(func() {
		AudioUnitGeneratorClass = _AudioUnitGeneratorClass{objc.GetClass("AVAudioUnitGenerator")}
	})
	return AudioUnitGeneratorClass
}

type _AudioUnitGeneratorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AudioUnitGenerator */
// An interface definition for the [AudioUnitGenerator] class.
type IAudioUnitGenerator interface {
	IAudioUnit
	
/* debug [class_interface_properties]: Properties for AudioUnitGenerator */
	// properties:
	Bypass() bool
	SetBypass(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AudioUnitGenerator */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AudioUnitGenerator */
// Alloc allocates a new instance without initialization.
func (ac _AudioUnitGeneratorClass) Alloc() AudioUnitGenerator {
	rv := objc.Send[AudioUnitGenerator](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioUnitGeneratorClass) New() AudioUnitGenerator {
	rv := objc.Send[AudioUnitGenerator](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioUnitGenerator) Init() AudioUnitGenerator {
	rv := objc.Send[AudioUnitGenerator](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioUnitGenerator) Autorelease() AudioUnitGenerator {
	rv := objc.Send[AudioUnitGenerator](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioUnitGenerator creates a new AudioUnitGenerator instance.
func NewAudioUnitGenerator() AudioUnitGenerator {
	return getAudioUnitGeneratorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AudioUnitGenerator */
// An object that generates audio output.
//
// A generator represents an of type or . A generator has no audio input, but produces audio output. An example is a tone generator.


// An object that generates audio output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitGenerator
type AudioUnitGenerator struct {
	AudioUnit
}

// AudioUnitGeneratorFrom constructs a [AudioUnitGenerator] from an unsafe.Pointer.
//
// An object that generates audio output.
func AudioUnitGeneratorFrom(ptr unsafe.Pointer) AudioUnitGenerator {
	return AudioUnitGenerator{
		AudioUnit: AudioUnitFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AudioUnitGenerator */

// Creates a generator audio unit with the specified description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitGenerator/init(audioComponentDescription:)
func NewAudioUnitGeneratorWithAudioComponentDescription(audioComponentDescription audiotoolbox.AudioComponentDescription) AudioUnitGenerator {
	instance := getAudioUnitGeneratorClass().Alloc()
	rv := objc.Send[AudioUnitGenerator](instance.ID, objc.Sel("initWithAudioComponentDescription:"), audioComponentDescription)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAudioUnitGeneratorWithAudioComponentDescription */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AudioUnitGenerator */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AudioUnitGenerator */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AudioUnitGenerator */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AudioUnitGenerator */

// The bypass state of the audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitGenerator/bypass
func (a_ AudioUnitGenerator) Bypass() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("bypass"))
	return rv
}/* debug [instance_properties/getter]: bypass */


// The bypass state of the audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitGenerator/bypass
func (a_ AudioUnitGenerator) SetBypass(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setBypass:"), value)
}/* debug [instance_properties/setter]: bypass */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAudioUnitGenerator */


