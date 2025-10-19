// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AVCaptureVideoDataOutput] class.
var (
	aVCaptureVideoDataOutputClass     _AVCaptureVideoDataOutputClass
	aVCaptureVideoDataOutputClassOnce sync.Once
)

func getAVCaptureVideoDataOutputClass() _AVCaptureVideoDataOutputClass {
	aVCaptureVideoDataOutputClassOnce.Do(func() {
		aVCaptureVideoDataOutputClass = _AVCaptureVideoDataOutputClass{objc.GetClass("AVCaptureVideoDataOutput")}
	})
	return aVCaptureVideoDataOutputClass
}

type _AVCaptureVideoDataOutputClass struct {
	class objc.Class
}

// An interface definition for the [AVCaptureVideoDataOutput] class.
type IAVCaptureVideoDataOutput interface {
	IAVCaptureOutput
	SetSampleBufferDelegateQueue(sampleBufferDelegate unsafe.Pointer, sampleBufferCallbackQueue unsafe.Pointer)
}

// A capture output that records video and provides access to video frames for processing.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput
type AVCaptureVideoDataOutput struct {
	AVCaptureOutput
}

// AVCaptureVideoDataOutputFrom constructs a [AVCaptureVideoDataOutput] from an unsafe.Pointer.
//
// A capture output that records video and provides access to video frames for processing.
func AVCaptureVideoDataOutputFrom(ptr unsafe.Pointer) AVCaptureVideoDataOutput {
	return AVCaptureVideoDataOutput{
		AVCaptureOutput: AVCaptureOutputFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AVCaptureVideoDataOutputClass) Alloc() AVCaptureVideoDataOutput {
	rv := objc.Send[AVCaptureVideoDataOutput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AVCaptureVideoDataOutputClass) New() AVCaptureVideoDataOutput {
	rv := objc.Send[AVCaptureVideoDataOutput](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVCaptureVideoDataOutput) Init() AVCaptureVideoDataOutput {
	rv := objc.Send[AVCaptureVideoDataOutput](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVCaptureVideoDataOutput) Autorelease() AVCaptureVideoDataOutput {
	rv := objc.Send[AVCaptureVideoDataOutput](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureVideoDataOutput creates a new AVCaptureVideoDataOutput instance.
func NewAVCaptureVideoDataOutput() AVCaptureVideoDataOutput {
	return getAVCaptureVideoDataOutputClass().New()
}


// Sets the sample buffer delegate and the queue for invoking callbacks.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/setSampleBufferDelegate(_:queue:)
func (a_ AVCaptureVideoDataOutput) SetSampleBufferDelegateQueue(sampleBufferDelegate unsafe.Pointer, sampleBufferCallbackQueue unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSampleBufferDelegate:queue:"), sampleBufferDelegate, sampleBufferCallbackQueue)
}


