// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AVAssetReaderTrackOutput] class.
var aVAssetReaderTrackOutputClass = _AVAssetReaderTrackOutputClass{objc.GetClass("AVAssetReaderTrackOutput")}

type _AVAssetReaderTrackOutputClass struct {
	class objc.Class
}

// An object that reads media data from a single track of an asset. [Full Topic]
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



