// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVAsset] class.
var aVAssetClass = _AVAssetClass{objc.GetClass("AVAsset")}

type _AVAssetClass struct {
	class objc.Class
}

// An object that models timed audiovisual media. [Full Topic]
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

// Loads tracks that contain media of a specified characteristic. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsset/loadTracks(withMediaCharacteristic:completionHandler:)
func (a_ AVAsset) LoadTracksWithMediaCharacteristicCompletionHandler(mediaCharacteristic unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("loadTracksWithMediaCharacteristic:completionHandler:"), mediaCharacteristic, completionHandler)
}


