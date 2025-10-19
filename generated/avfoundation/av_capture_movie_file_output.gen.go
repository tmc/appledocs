// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AVCaptureMovieFileOutput] class.
var aVCaptureMovieFileOutputClass = _AVCaptureMovieFileOutputClass{objc.GetClass("AVCaptureMovieFileOutput")}

type _AVCaptureMovieFileOutputClass struct {
	class objc.Class
}

// A capture output that records video and audio to a QuickTime movie file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMovieFileOutput

type AVCaptureMovieFileOutput struct {
	AVCaptureFileOutput
}

// AVCaptureMovieFileOutputFrom constructs a [AVCaptureMovieFileOutput] from an unsafe.Pointer.
//
// A capture output that records video and audio to a QuickTime movie file.
func AVCaptureMovieFileOutputFrom(ptr unsafe.Pointer) AVCaptureMovieFileOutput {
	return AVCaptureMovieFileOutput{
		AVCaptureFileOutput: AVCaptureFileOutputFrom(ptr),
	}
}



