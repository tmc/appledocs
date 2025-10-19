// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVSampleBufferVideoRenderer] class.
var aVSampleBufferVideoRendererClass = _AVSampleBufferVideoRendererClass{objc.GetClass("AVSampleBufferVideoRenderer")}

type _AVSampleBufferVideoRendererClass struct {
	class objc.Class
}

// An object that enqueues video sample buffers for rendering. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferVideoRenderer

type AVSampleBufferVideoRenderer struct {
	objectivec.Object
}

// AVSampleBufferVideoRendererFrom constructs a [AVSampleBufferVideoRenderer] from an unsafe.Pointer.
//
// An object that enqueues video sample buffers for rendering.
func AVSampleBufferVideoRendererFrom(ptr unsafe.Pointer) AVSampleBufferVideoRenderer {
	return AVSampleBufferVideoRenderer{objectivec.Object{objc.ID(ptr)}}
}



