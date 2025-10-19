// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AVURLAsset] class.
var (
	aVURLAssetClass     _AVURLAssetClass
	aVURLAssetClassOnce sync.Once
)

func getAVURLAssetClass() _AVURLAssetClass {
	aVURLAssetClassOnce.Do(func() {
		aVURLAssetClass = _AVURLAssetClass{objc.GetClass("AVURLAsset")}
	})
	return aVURLAssetClass
}

type _AVURLAssetClass struct {
	class objc.Class
}

// An interface definition for the [AVURLAsset] class.
type IAVURLAsset interface {
	IAVAsset
}

// An asset that represents media at a local or remote URL.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVURLAsset
type AVURLAsset struct {
	AVAsset
}

// AVURLAssetFrom constructs a [AVURLAsset] from an unsafe.Pointer.
//
// An asset that represents media at a local or remote URL.
func AVURLAssetFrom(ptr unsafe.Pointer) AVURLAsset {
	return AVURLAsset{
		AVAsset: AVAssetFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AVURLAssetClass) Alloc() AVURLAsset {
	rv := objc.Send[AVURLAsset](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AVURLAssetClass) New() AVURLAsset {
	rv := objc.Send[AVURLAsset](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVURLAsset) Init() AVURLAsset {
	rv := objc.Send[AVURLAsset](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVURLAsset) Autorelease() AVURLAsset {
	rv := objc.Send[AVURLAsset](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVURLAsset creates a new AVURLAsset instance.
func NewAVURLAsset() AVURLAsset {
	return getAVURLAssetClass().New()
}




