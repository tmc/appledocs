// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVCaptureInput] class.
var (
	aVCaptureInputClass     _AVCaptureInputClass
	aVCaptureInputClassOnce sync.Once
)

func getAVCaptureInputClass() _AVCaptureInputClass {
	aVCaptureInputClassOnce.Do(func() {
		aVCaptureInputClass = _AVCaptureInputClass{objc.GetClass("AVCaptureInput")}
	})
	return aVCaptureInputClass
}

type _AVCaptureInputClass struct {
	class objc.Class
}

// An interface definition for the [AVCaptureInput] class.
type IAVCaptureInput interface {
	objectivec.IObject
}

// An abstract superclass for objects that provide input data to a capture session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureInput
type AVCaptureInput struct {
	objectivec.Object
}

// AVCaptureInputFrom constructs a [AVCaptureInput] from an unsafe.Pointer.
//
// An abstract superclass for objects that provide input data to a capture session.
func AVCaptureInputFrom(ptr unsafe.Pointer) AVCaptureInput {
	return AVCaptureInput{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AVCaptureInputClass) Alloc() AVCaptureInput {
	rv := objc.Send[AVCaptureInput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AVCaptureInputClass) New() AVCaptureInput {
	rv := objc.Send[AVCaptureInput](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVCaptureInput) Init() AVCaptureInput {
	rv := objc.Send[AVCaptureInput](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVCaptureInput) Autorelease() AVCaptureInput {
	rv := objc.Send[AVCaptureInput](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureInput creates a new AVCaptureInput instance.
func NewAVCaptureInput() AVCaptureInput {
	return getAVCaptureInputClass().New()
}




