// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AVCaptureDeviceInput] class.
var (
	aVCaptureDeviceInputClass     _AVCaptureDeviceInputClass
	aVCaptureDeviceInputClassOnce sync.Once
)

func getAVCaptureDeviceInputClass() _AVCaptureDeviceInputClass {
	aVCaptureDeviceInputClassOnce.Do(func() {
		aVCaptureDeviceInputClass = _AVCaptureDeviceInputClass{objc.GetClass("AVCaptureDeviceInput")}
	})
	return aVCaptureDeviceInputClass
}

type _AVCaptureDeviceInputClass struct {
	class objc.Class
}

// An interface definition for the [AVCaptureDeviceInput] class.
type IAVCaptureDeviceInput interface {
	IAVCaptureInput
}

// An object that provides media input from a capture device to a capture session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput
type AVCaptureDeviceInput struct {
	AVCaptureInput
}

// AVCaptureDeviceInputFrom constructs a [AVCaptureDeviceInput] from an unsafe.Pointer.
//
// An object that provides media input from a capture device to a capture session.
func AVCaptureDeviceInputFrom(ptr unsafe.Pointer) AVCaptureDeviceInput {
	return AVCaptureDeviceInput{
		AVCaptureInput: AVCaptureInputFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AVCaptureDeviceInputClass) Alloc() AVCaptureDeviceInput {
	rv := objc.Send[AVCaptureDeviceInput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AVCaptureDeviceInputClass) New() AVCaptureDeviceInput {
	rv := objc.Send[AVCaptureDeviceInput](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVCaptureDeviceInput) Init() AVCaptureDeviceInput {
	rv := objc.Send[AVCaptureDeviceInput](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVCaptureDeviceInput) Autorelease() AVCaptureDeviceInput {
	rv := objc.Send[AVCaptureDeviceInput](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureDeviceInput creates a new AVCaptureDeviceInput instance.
func NewAVCaptureDeviceInput() AVCaptureDeviceInput {
	return getAVCaptureDeviceInputClass().New()
}




