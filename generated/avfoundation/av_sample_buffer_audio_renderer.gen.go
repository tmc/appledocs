// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVSampleBufferAudioRenderer] class.
var aVSampleBufferAudioRendererClass = _AVSampleBufferAudioRendererClass{objc.GetClass("AVSampleBufferAudioRenderer")}

type _AVSampleBufferAudioRendererClass struct {
	class objc.Class
}

// An object used to decompress audio and play compressed or uncompressed audio. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAudioRenderer

type AVSampleBufferAudioRenderer struct {
	objectivec.Object
}

// AVSampleBufferAudioRendererFrom constructs a [AVSampleBufferAudioRenderer] from an unsafe.Pointer.
//
// An object used to decompress audio and play compressed or uncompressed audio.
func AVSampleBufferAudioRendererFrom(ptr unsafe.Pointer) AVSampleBufferAudioRenderer {
	return AVSampleBufferAudioRenderer{objectivec.Object{objc.ID(ptr)}}
}



