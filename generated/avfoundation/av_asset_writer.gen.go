// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVAssetWriter] class.
var aVAssetWriterClass = _AVAssetWriterClass{objc.GetClass("AVAssetWriter")}

type _AVAssetWriterClass struct {
	class objc.Class
}

// An object that writes media data to a container file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter

type AVAssetWriter struct {
	objectivec.Object
}

// AVAssetWriterFrom constructs a [AVAssetWriter] from an unsafe.Pointer.
//
// An object that writes media data to a container file.
func AVAssetWriterFrom(ptr unsafe.Pointer) AVAssetWriter {
	return AVAssetWriter{objectivec.Object{objc.ID(ptr)}}
}



