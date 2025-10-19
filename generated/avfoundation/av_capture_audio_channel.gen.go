// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVCaptureAudioChannel] class.
var aVCaptureAudioChannelClass = _AVCaptureAudioChannelClass{objc.GetClass("AVCaptureAudioChannel")}

type _AVCaptureAudioChannelClass struct {
	class objc.Class
}

// An object that monitors average and peak power levels for an audio channel in a capture connection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioChannel

type AVCaptureAudioChannel struct {
	objectivec.Object
}

// AVCaptureAudioChannelFrom constructs a [AVCaptureAudioChannel] from an unsafe.Pointer.
//
// An object that monitors average and peak power levels for an audio channel in a capture connection.
func AVCaptureAudioChannelFrom(ptr unsafe.Pointer) AVCaptureAudioChannel {
	return AVCaptureAudioChannel{objectivec.Object{objc.ID(ptr)}}
}



