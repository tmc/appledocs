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

// An interface definition for the [AVCaptureMovieFileOutput] class.
type IAVCaptureMovieFileOutput interface {
	IAVCaptureFileOutput
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
// Alloc allocates a new instance without initialization.
func (ac _AVCaptureMovieFileOutputClass) Alloc() AVCaptureMovieFileOutput {
	rv := objc.Send[AVCaptureMovieFileOutput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (ac _AVCaptureMovieFileOutputClass) New() AVCaptureMovieFileOutput {
	rv := objc.Send[AVCaptureMovieFileOutput](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVCaptureMovieFileOutput) Init() AVCaptureMovieFileOutput {
	rv := objc.Send[AVCaptureMovieFileOutput](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVCaptureMovieFileOutput) Autorelease() AVCaptureMovieFileOutput {
	rv := objc.Send[AVCaptureMovieFileOutput](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureMovieFileOutput creates a new AVCaptureMovieFileOutput instance.
func NewAVCaptureMovieFileOutput() AVCaptureMovieFileOutput {
	return aVCaptureMovieFileOutputClass.New()
}




