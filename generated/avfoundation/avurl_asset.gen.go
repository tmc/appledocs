// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AVURLAsset] class.
var aVURLAssetClass = _AVURLAssetClass{objc.GetClass("AVURLAsset")}

type _AVURLAssetClass struct {
	class objc.Class
}

// An asset that represents media at a local or remote URL. [Full Topic]
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



