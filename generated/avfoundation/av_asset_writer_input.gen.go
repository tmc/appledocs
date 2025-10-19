// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVAssetWriterInput] class.
var aVAssetWriterInputClass = _AVAssetWriterInputClass{objc.GetClass("AVAssetWriterInput")}

type _AVAssetWriterInputClass struct {
	class objc.Class
}

// An object that appends media samples to a track in an asset writer’s output file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput

type AVAssetWriterInput struct {
	objectivec.Object
}

// AVAssetWriterInputFrom constructs a [AVAssetWriterInput] from an unsafe.Pointer.
//
// An object that appends media samples to a track in an asset writer’s output file.
func AVAssetWriterInputFrom(ptr unsafe.Pointer) AVAssetWriterInput {
	return AVAssetWriterInput{objectivec.Object{objc.ID(ptr)}}
}



