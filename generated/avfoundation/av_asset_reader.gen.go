// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVAssetReader] class.
var aVAssetReaderClass = _AVAssetReaderClass{objc.GetClass("AVAssetReader")}

type _AVAssetReaderClass struct {
	class objc.Class
}

// An object that reads media data from an asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReader

type AVAssetReader struct {
	objectivec.Object
}

// AVAssetReaderFrom constructs a [AVAssetReader] from an unsafe.Pointer.
//
// An object that reads media data from an asset.
func AVAssetReaderFrom(ptr unsafe.Pointer) AVAssetReader {
	return AVAssetReader{objectivec.Object{objc.ID(ptr)}}
}



