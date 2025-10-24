// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAudioMix */


/* debug [class_header]: Header for AVAudioMix */
// The class instance for the [AudioMix] class.
var (
	AudioMixClass     _AudioMixClass
	AudioMixClassOnce sync.Once
)

func getAudioMixClass() _AudioMixClass {
	AudioMixClassOnce.Do(func() {
		AudioMixClass = _AudioMixClass{objc.GetClass("AVAudioMix")}
	})
	return AudioMixClass
}

type _AudioMixClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AudioMix */
// An interface definition for the [AudioMix] class.
type IAudioMix interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AudioMix */
	// properties:
	InputParameters() []AudioMixInputParameters
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AudioMix */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AudioMix */
// Alloc allocates a new instance without initialization.
func (ac _AudioMixClass) Alloc() AudioMix {
	rv := objc.Send[AudioMix](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioMixClass) New() AudioMix {
	rv := objc.Send[AudioMix](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioMix) Init() AudioMix {
	rv := objc.Send[AudioMix](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioMix) Autorelease() AudioMix {
	rv := objc.Send[AudioMix](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioMix creates a new AudioMix instance.
func NewAudioMix() AudioMix {
	return getAudioMixClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AudioMix */
// An object that manages the input parameters for mixing audio tracks.


// An object that manages the input parameters for mixing audio tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAudioMix
type AudioMix struct {
	objectivec.Object
}

// AudioMixFrom constructs a [AudioMix] from an unsafe.Pointer.
//
// An object that manages the input parameters for mixing audio tracks.
func AudioMixFrom(ptr unsafe.Pointer) AudioMix {
	return AudioMix{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AudioMix *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AudioMix */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AudioMix */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AudioMix */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AudioMix */

// An array of input parameters for the mix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAudioMix/inputParameters
func (a_ AudioMix) InputParameters() []AudioMixInputParameters {
	rv := objc.Send[[]AudioMixInputParameters](a_.ID, objc.Sel("inputParameters"))
	return rv
}/* debug [instance_properties/getter]: inputParameters */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAudioMix */



