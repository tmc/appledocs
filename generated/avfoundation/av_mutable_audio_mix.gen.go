// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVMutableAudioMix */


/* debug [class_header]: Header for AVMutableAudioMix */
// The class instance for the [MutableAudioMix] class.
var (
	MutableAudioMixClass     _MutableAudioMixClass
	MutableAudioMixClassOnce sync.Once
)

func getMutableAudioMixClass() _MutableAudioMixClass {
	MutableAudioMixClassOnce.Do(func() {
		MutableAudioMixClass = _MutableAudioMixClass{objc.GetClass("AVMutableAudioMix")}
	})
	return MutableAudioMixClass
}

type _MutableAudioMixClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MutableAudioMix */
// An interface definition for the [MutableAudioMix] class.
type IMutableAudioMix interface {
	IAudioMix
	
/* debug [class_interface_properties]: Properties for MutableAudioMix */
	// properties:
	InputParameters() []AudioMixInputParameters
	SetInputParameters(value []AudioMixInputParameters)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MutableAudioMix */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MutableAudioMix */
// Alloc allocates a new instance without initialization.
func (mc _MutableAudioMixClass) Alloc() MutableAudioMix {
	rv := objc.Send[MutableAudioMix](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MutableAudioMixClass) New() MutableAudioMix {
	rv := objc.Send[MutableAudioMix](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MutableAudioMix) Init() MutableAudioMix {
	rv := objc.Send[MutableAudioMix](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MutableAudioMix) Autorelease() MutableAudioMix {
	rv := objc.Send[MutableAudioMix](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMutableAudioMix creates a new MutableAudioMix instance.
func NewMutableAudioMix() MutableAudioMix {
	return getMutableAudioMixClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MutableAudioMix */
// An object that manages the input parameters for mixing audio tracks.


// An object that manages the input parameters for mixing audio tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableAudioMix
type MutableAudioMix struct {
	AudioMix
}

// MutableAudioMixFrom constructs a [MutableAudioMix] from an unsafe.Pointer.
//
// An object that manages the input parameters for mixing audio tracks.
func MutableAudioMixFrom(ptr unsafe.Pointer) MutableAudioMix {
	return MutableAudioMix{
		AudioMix: AudioMixFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MutableAudioMix *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MutableAudioMix */

// Returns a new mutable audio mix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableAudioMix/audioMix
func (mc _MutableAudioMixClass) AudioMix() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("audioMix"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AudioMix) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MutableAudioMix */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MutableAudioMix */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MutableAudioMix */

// An array of input parameters for the mix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableAudioMix/inputParameters
func (m_ MutableAudioMix) InputParameters() []AudioMixInputParameters {
	rv := objc.Send[[]AudioMixInputParameters](m_.ID, objc.Sel("inputParameters"))
	return rv
}/* debug [instance_properties/getter]: inputParameters */


// An array of input parameters for the mix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableAudioMix/inputParameters
func (m_ MutableAudioMix) SetInputParameters(value []AudioMixInputParameters) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setInputParameters:"), nsArray)
}/* debug [instance_properties/setter]: inputParameters */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMutableAudioMix */



