// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVCaptureOutput] class.
var aVCaptureOutputClass = _AVCaptureOutputClass{objc.GetClass("AVCaptureOutput")}

type _AVCaptureOutputClass struct {
	class objc.Class
}

// An interface definition for the [AVCaptureOutput] class.
type IAVCaptureOutput interface {
	objectivec.IObject
}

// An abstract superclass for objects that provide media output destinations for a capture session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureOutput

type AVCaptureOutput struct {
	objectivec.Object
}

// AVCaptureOutputFrom constructs a [AVCaptureOutput] from an unsafe.Pointer.
//
// An abstract superclass for objects that provide media output destinations for a capture session.
func AVCaptureOutputFrom(ptr unsafe.Pointer) AVCaptureOutput {
	return AVCaptureOutput{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (ac _AVCaptureOutputClass) Alloc() AVCaptureOutput {
	rv := objc.Send[AVCaptureOutput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (ac _AVCaptureOutputClass) New() AVCaptureOutput {
	rv := objc.Send[AVCaptureOutput](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVCaptureOutput) Init() AVCaptureOutput {
	rv := objc.Send[AVCaptureOutput](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVCaptureOutput) Autorelease() AVCaptureOutput {
	rv := objc.Send[AVCaptureOutput](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureOutput creates a new AVCaptureOutput instance.
func NewAVCaptureOutput() AVCaptureOutput {
	return aVCaptureOutputClass.New()
}




