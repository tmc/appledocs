// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AVPlayerItemRenderedLegibleOutput] class.
var (
	aVPlayerItemRenderedLegibleOutputClass     _AVPlayerItemRenderedLegibleOutputClass
	aVPlayerItemRenderedLegibleOutputClassOnce sync.Once
)

func getAVPlayerItemRenderedLegibleOutputClass() _AVPlayerItemRenderedLegibleOutputClass {
	aVPlayerItemRenderedLegibleOutputClassOnce.Do(func() {
		aVPlayerItemRenderedLegibleOutputClass = _AVPlayerItemRenderedLegibleOutputClass{objc.GetClass("AVPlayerItemRenderedLegibleOutput")}
	})
	return aVPlayerItemRenderedLegibleOutputClass
}

type _AVPlayerItemRenderedLegibleOutputClass struct {
	class objc.Class
}

// An interface definition for the [AVPlayerItemRenderedLegibleOutput] class.
type IAVPlayerItemRenderedLegibleOutput interface {
	IAVPlayerItemOutput
}

// A player item output that vends media with a legible characteristic as rendered pixel buffers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemRenderedLegibleOutput
type AVPlayerItemRenderedLegibleOutput struct {
	AVPlayerItemOutput
}

// AVPlayerItemRenderedLegibleOutputFrom constructs a [AVPlayerItemRenderedLegibleOutput] from an unsafe.Pointer.
//
// A player item output that vends media with a legible characteristic as rendered pixel buffers.
func AVPlayerItemRenderedLegibleOutputFrom(ptr unsafe.Pointer) AVPlayerItemRenderedLegibleOutput {
	return AVPlayerItemRenderedLegibleOutput{
		AVPlayerItemOutput: AVPlayerItemOutputFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AVPlayerItemRenderedLegibleOutputClass) Alloc() AVPlayerItemRenderedLegibleOutput {
	rv := objc.Send[AVPlayerItemRenderedLegibleOutput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AVPlayerItemRenderedLegibleOutputClass) New() AVPlayerItemRenderedLegibleOutput {
	rv := objc.Send[AVPlayerItemRenderedLegibleOutput](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVPlayerItemRenderedLegibleOutput) Init() AVPlayerItemRenderedLegibleOutput {
	rv := objc.Send[AVPlayerItemRenderedLegibleOutput](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVPlayerItemRenderedLegibleOutput) Autorelease() AVPlayerItemRenderedLegibleOutput {
	rv := objc.Send[AVPlayerItemRenderedLegibleOutput](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlayerItemRenderedLegibleOutput creates a new AVPlayerItemRenderedLegibleOutput instance.
func NewAVPlayerItemRenderedLegibleOutput() AVPlayerItemRenderedLegibleOutput {
	return getAVPlayerItemRenderedLegibleOutputClass().New()
}




