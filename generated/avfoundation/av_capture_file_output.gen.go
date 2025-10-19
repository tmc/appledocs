// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AVCaptureFileOutput] class.
var aVCaptureFileOutputClass = _AVCaptureFileOutputClass{objc.GetClass("AVCaptureFileOutput")}

type _AVCaptureFileOutputClass struct {
	class objc.Class
}

// An interface definition for the [AVCaptureFileOutput] class.
type IAVCaptureFileOutput interface {
	IAVCaptureOutput
}

// The abstract superclass for capture outputs that can record captured data to a file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutput

type AVCaptureFileOutput struct {
	AVCaptureOutput
}

// AVCaptureFileOutputFrom constructs a [AVCaptureFileOutput] from an unsafe.Pointer.
//
// The abstract superclass for capture outputs that can record captured data to a file.
func AVCaptureFileOutputFrom(ptr unsafe.Pointer) AVCaptureFileOutput {
	return AVCaptureFileOutput{
		AVCaptureOutput: AVCaptureOutputFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (ac _AVCaptureFileOutputClass) Alloc() AVCaptureFileOutput {
	rv := objc.Send[AVCaptureFileOutput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (ac _AVCaptureFileOutputClass) New() AVCaptureFileOutput {
	rv := objc.Send[AVCaptureFileOutput](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVCaptureFileOutput) Init() AVCaptureFileOutput {
	rv := objc.Send[AVCaptureFileOutput](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVCaptureFileOutput) Autorelease() AVCaptureFileOutput {
	rv := objc.Send[AVCaptureFileOutput](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureFileOutput creates a new AVCaptureFileOutput instance.
func NewAVCaptureFileOutput() AVCaptureFileOutput {
	return aVCaptureFileOutputClass.New()
}




