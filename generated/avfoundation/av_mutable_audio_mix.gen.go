// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [MutableAudioMix] class.
type IMutableAudioMix interface {
	IAudioMix
	

	// properties:
	InputParameters() []AudioMixInputParameters
	SetInputParameters(value []AudioMixInputParameters)


	

	// methods:


}





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










// Returns a new mutable audio mix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableAudioMix/audioMix
func (mc _MutableAudioMixClass) AudioMix() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("audioMix"))
	return rv
}

















// An array of input parameters for the mix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableAudioMix/inputParameters
func (m_ MutableAudioMix) InputParameters() []AudioMixInputParameters {
	rv := objc.Send[[]AudioMixInputParameters](m_.ID, objc.Sel("inputParameters"))
	return rv
}


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
}








