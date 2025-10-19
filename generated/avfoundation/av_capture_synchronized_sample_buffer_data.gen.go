// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AVCaptureSynchronizedSampleBufferData] class.
var aVCaptureSynchronizedSampleBufferDataClass = _AVCaptureSynchronizedSampleBufferDataClass{objc.GetClass("AVCaptureSynchronizedSampleBufferData")}

type _AVCaptureSynchronizedSampleBufferDataClass struct {
	class objc.Class
}

// A container for video or audio samples collected using synchronized capture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSynchronizedSampleBufferData

type AVCaptureSynchronizedSampleBufferData struct {
	AVCaptureSynchronizedData
}

// AVCaptureSynchronizedSampleBufferDataFrom constructs a [AVCaptureSynchronizedSampleBufferData] from an unsafe.Pointer.
//
// A container for video or audio samples collected using synchronized capture.
func AVCaptureSynchronizedSampleBufferDataFrom(ptr unsafe.Pointer) AVCaptureSynchronizedSampleBufferData {
	return AVCaptureSynchronizedSampleBufferData{
		AVCaptureSynchronizedData: AVCaptureSynchronizedDataFrom(ptr),
	}
}



