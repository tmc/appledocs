// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVPlayerItemOutput] class.
var (
	aVPlayerItemOutputClass     _AVPlayerItemOutputClass
	aVPlayerItemOutputClassOnce sync.Once
)

func getAVPlayerItemOutputClass() _AVPlayerItemOutputClass {
	aVPlayerItemOutputClassOnce.Do(func() {
		aVPlayerItemOutputClass = _AVPlayerItemOutputClass{objc.GetClass("AVPlayerItemOutput")}
	})
	return aVPlayerItemOutputClass
}

type _AVPlayerItemOutputClass struct {
	class objc.Class
}

// An interface definition for the [AVPlayerItemOutput] class.
type IAVPlayerItemOutput interface {
	objectivec.IObject
}

// A parent class referenced by other AVFoundation classes. [Full Topic]
type AVPlayerItemOutput struct {
	objectivec.Object
}

// AVPlayerItemOutputFrom constructs a [AVPlayerItemOutput] from an unsafe.Pointer.
//
// A parent class referenced by other AVFoundation classes.
func AVPlayerItemOutputFrom(ptr unsafe.Pointer) AVPlayerItemOutput {
	return AVPlayerItemOutput{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AVPlayerItemOutputClass) Alloc() AVPlayerItemOutput {
	rv := objc.Send[AVPlayerItemOutput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AVPlayerItemOutputClass) New() AVPlayerItemOutput {
	rv := objc.Send[AVPlayerItemOutput](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVPlayerItemOutput) Init() AVPlayerItemOutput {
	rv := objc.Send[AVPlayerItemOutput](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVPlayerItemOutput) Autorelease() AVPlayerItemOutput {
	rv := objc.Send[AVPlayerItemOutput](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlayerItemOutput creates a new AVPlayerItemOutput instance.
func NewAVPlayerItemOutput() AVPlayerItemOutput {
	return getAVPlayerItemOutputClass().New()
}




