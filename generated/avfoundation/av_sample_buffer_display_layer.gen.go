// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AVSampleBufferDisplayLayer] class.
var aVSampleBufferDisplayLayerClass = _AVSampleBufferDisplayLayerClass{objc.GetClass("AVSampleBufferDisplayLayer")}

type _AVSampleBufferDisplayLayerClass struct {
	class objc.Class
}

// An object that displays compressed or uncompressed video frames. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer

type AVSampleBufferDisplayLayer struct {
	Layer
}

// AVSampleBufferDisplayLayerFrom constructs a [AVSampleBufferDisplayLayer] from an unsafe.Pointer.
//
// An object that displays compressed or uncompressed video frames.
func AVSampleBufferDisplayLayerFrom(ptr unsafe.Pointer) AVSampleBufferDisplayLayer {
	return AVSampleBufferDisplayLayer{
		Layer: LayerFrom(ptr),
	}
}



