// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVAssetResourceLoader] class.
var aVAssetResourceLoaderClass = _AVAssetResourceLoaderClass{objc.GetClass("AVAssetResourceLoader")}

type _AVAssetResourceLoaderClass struct {
	class objc.Class
}

// An object that mediates resource requests from a URL asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoader

type AVAssetResourceLoader struct {
	objectivec.Object
}

// AVAssetResourceLoaderFrom constructs a [AVAssetResourceLoader] from an unsafe.Pointer.
//
// An object that mediates resource requests from a URL asset.
func AVAssetResourceLoaderFrom(ptr unsafe.Pointer) AVAssetResourceLoader {
	return AVAssetResourceLoader{objectivec.Object{objc.ID(ptr)}}
}



