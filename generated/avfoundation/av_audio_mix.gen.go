// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVAudioMix] class.
var (
	aVAudioMixClass     _AVAudioMixClass
	aVAudioMixClassOnce sync.Once
)

func getAVAudioMixClass() _AVAudioMixClass {
	aVAudioMixClassOnce.Do(func() {
		aVAudioMixClass = _AVAudioMixClass{objc.GetClass("AVAudioMix")}
	})
	return aVAudioMixClass
}

type _AVAudioMixClass struct {
	class objc.Class
}

// An interface definition for the [AVAudioMix] class.
type IAVAudioMix interface {
	objectivec.IObject
}

// An object that manages the input parameters for mixing audio tracks.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAudioMix
type AVAudioMix struct {
	objectivec.Object
}

// AVAudioMixFrom constructs a [AVAudioMix] from an unsafe.Pointer.
//
// An object that manages the input parameters for mixing audio tracks.
func AVAudioMixFrom(ptr unsafe.Pointer) AVAudioMix {
	return AVAudioMix{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AVAudioMixClass) Alloc() AVAudioMix {
	rv := objc.Send[AVAudioMix](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AVAudioMixClass) New() AVAudioMix {
	rv := objc.Send[AVAudioMix](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVAudioMix) Init() AVAudioMix {
	rv := objc.Send[AVAudioMix](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVAudioMix) Autorelease() AVAudioMix {
	rv := objc.Send[AVAudioMix](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioMix creates a new AVAudioMix instance.
func NewAVAudioMix() AVAudioMix {
	return getAVAudioMixClass().New()
}




