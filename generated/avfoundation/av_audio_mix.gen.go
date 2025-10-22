// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [AudioMix] class.
type IAudioMix interface {
	objectivec.IObject
	InputParameters() []AudioMixInputParameters
}

// An object that manages the input parameters for mixing audio tracks.
//
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

// Alloc allocates a new instance without initialization.
func (ac _AudioMixClass) Alloc() AudioMix {
	rv := objc.Send[AudioMix](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// An array of input parameters for the mix.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAudioMix/inputParameters
func (a_ AudioMix) InputParameters() []AudioMixInputParameters {
	rv := objc.Send[[]AudioMixInputParameters](a_.ID, objc.Sel("inputParameters"))
	return rv
}



