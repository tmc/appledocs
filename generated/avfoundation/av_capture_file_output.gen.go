// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CaptureFileOutput] class.
var (
	CaptureFileOutputClass     _CaptureFileOutputClass
	CaptureFileOutputClassOnce sync.Once
)

func getCaptureFileOutputClass() _CaptureFileOutputClass {
	CaptureFileOutputClassOnce.Do(func() {
		CaptureFileOutputClass = _CaptureFileOutputClass{objc.GetClass("AVCaptureFileOutput")}
	})
	return CaptureFileOutputClass
}

type _CaptureFileOutputClass struct {
	class objc.Class
}

// An interface definition for the [CaptureFileOutput] class.
type ICaptureFileOutput interface {
	ICaptureOutput
	StartRecordingToOutputFileURLRecordingDelegate(outputFileURL unsafe.Pointer, delegate objc.ID)
}

// The abstract superclass for capture outputs that can record captured data to a file.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutput
type CaptureFileOutput struct {
	CaptureOutput
}

// CaptureFileOutputFrom constructs a [CaptureFileOutput] from an unsafe.Pointer.
//
// The abstract superclass for capture outputs that can record captured data to a file.
func CaptureFileOutputFrom(ptr unsafe.Pointer) CaptureFileOutput {
	return CaptureFileOutput{
		CaptureOutput: CaptureOutputFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CaptureFileOutputClass) Alloc() CaptureFileOutput {
	rv := objc.Send[CaptureFileOutput](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CaptureFileOutputClass) New() CaptureFileOutput {
	rv := objc.Send[CaptureFileOutput](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureFileOutput) Init() CaptureFileOutput {
	rv := objc.Send[CaptureFileOutput](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureFileOutput) Autorelease() CaptureFileOutput {
	rv := objc.Send[CaptureFileOutput](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureFileOutput creates a new CaptureFileOutput instance.
func NewCaptureFileOutput() CaptureFileOutput {
	return getCaptureFileOutputClass().New()
}


// Starts recording media to the specified output URL.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutput/startRecording(to:recordingDelegate:)
func (c_ CaptureFileOutput) StartRecordingToOutputFileURLRecordingDelegate(outputFileURL unsafe.Pointer, delegate objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("startRecordingToOutputFileURL:recordingDelegate:"), outputFileURL, delegate)
}



