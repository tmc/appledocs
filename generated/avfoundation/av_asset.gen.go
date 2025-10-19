// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVAsset] class.
var (
	aVAssetClass     _AVAssetClass
	aVAssetClassOnce sync.Once
)

func getAVAssetClass() _AVAssetClass {
	aVAssetClassOnce.Do(func() {
		aVAssetClass = _AVAssetClass{objc.GetClass("AVAsset")}
	})
	return aVAssetClass
}

type _AVAssetClass struct {
	class objc.Class
}

// An interface definition for the [AVAsset] class.
type IAVAsset interface {
	objectivec.IObject
	LoadTracksWithMediaCharacteristicCompletionHandler(mediaCharacteristic unsafe.Pointer, completionHandler unsafe.Pointer)
}

// An object that models timed audiovisual media.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset
type AVAsset struct {
	objectivec.Object
}

// AVAssetFrom constructs a [AVAsset] from an unsafe.Pointer.
//
// An object that models timed audiovisual media.
func AVAssetFrom(ptr unsafe.Pointer) AVAsset {
	return AVAsset{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AVAssetClass) Alloc() AVAsset {
	rv := objc.Send[AVAsset](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AVAssetClass) New() AVAsset {
	rv := objc.Send[AVAsset](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVAsset) Init() AVAsset {
	rv := objc.Send[AVAsset](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVAsset) Autorelease() AVAsset {
	rv := objc.Send[AVAsset](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAsset creates a new AVAsset instance.
func NewAVAsset() AVAsset {
	return getAVAssetClass().New()
}


// Loads tracks that contain media of a specified characteristic.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/loadTracks(withMediaCharacteristic:completionHandler:)
func (a_ AVAsset) LoadTracksWithMediaCharacteristicCompletionHandler(mediaCharacteristic unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("loadTracksWithMediaCharacteristic:completionHandler:"), mediaCharacteristic, completionHandler)
}


