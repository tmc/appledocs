// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AVAssetReaderTrackOutput] class.
var (
	aVAssetReaderTrackOutputClass     _AVAssetReaderTrackOutputClass
	aVAssetReaderTrackOutputClassOnce sync.Once
)

func getAVAssetReaderTrackOutputClass() _AVAssetReaderTrackOutputClass {
	aVAssetReaderTrackOutputClassOnce.Do(func() {
		aVAssetReaderTrackOutputClass = _AVAssetReaderTrackOutputClass{objc.GetClass("AVAssetReaderTrackOutput")}
	})
	return aVAssetReaderTrackOutputClass
}

type _AVAssetReaderTrackOutputClass struct {
	class objc.Class
}

// An interface definition for the [AVAssetReaderTrackOutput] class.
type IAVAssetReaderTrackOutput interface {
	IAVAssetReaderOutput
}

// An object that reads media data from a single track of an asset.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderTrackOutput
type AVAssetReaderTrackOutput struct {
	AVAssetReaderOutput
}

// AVAssetReaderTrackOutputFrom constructs a [AVAssetReaderTrackOutput] from an unsafe.Pointer.
//
// An object that reads media data from a single track of an asset.
func AVAssetReaderTrackOutputFrom(ptr unsafe.Pointer) AVAssetReaderTrackOutput {
	return AVAssetReaderTrackOutput{
		AVAssetReaderOutput: AVAssetReaderOutputFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AVAssetReaderTrackOutputClass) Alloc() AVAssetReaderTrackOutput {
	rv := objc.Send[AVAssetReaderTrackOutput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AVAssetReaderTrackOutputClass) New() AVAssetReaderTrackOutput {
	rv := objc.Send[AVAssetReaderTrackOutput](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVAssetReaderTrackOutput) Init() AVAssetReaderTrackOutput {
	rv := objc.Send[AVAssetReaderTrackOutput](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVAssetReaderTrackOutput) Autorelease() AVAssetReaderTrackOutput {
	rv := objc.Send[AVAssetReaderTrackOutput](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetReaderTrackOutput creates a new AVAssetReaderTrackOutput instance.
func NewAVAssetReaderTrackOutput() AVAssetReaderTrackOutput {
	return getAVAssetReaderTrackOutputClass().New()
}




