// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CaptureScreenInput] class.
var (
	CaptureScreenInputClass     _CaptureScreenInputClass
	CaptureScreenInputClassOnce sync.Once
)

func getCaptureScreenInputClass() _CaptureScreenInputClass {
	CaptureScreenInputClassOnce.Do(func() {
		CaptureScreenInputClass = _CaptureScreenInputClass{objc.GetClass("AVCaptureScreenInput")}
	})
	return CaptureScreenInputClass
}

type _CaptureScreenInputClass struct {
	class objc.Class
}

// An interface definition for the [CaptureScreenInput] class.
type ICaptureScreenInput interface {
	ICaptureInput
}

// A capture input for recording from a screen in macOS.
//
// This class is a concrete capture input subclass that provides an interface to capture media from a screen or a portion of a screen. Use instances of this class as input sources for objects that provide media data from one of the screens connected to the system, represented by .
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureScreenInput
type CaptureScreenInput struct {
	CaptureInput
}

// CaptureScreenInputFrom constructs a [CaptureScreenInput] from an unsafe.Pointer.
//
// A capture input for recording from a screen in macOS.
func CaptureScreenInputFrom(ptr unsafe.Pointer) CaptureScreenInput {
	return CaptureScreenInput{
		CaptureInput: CaptureInputFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CaptureScreenInputClass) Alloc() CaptureScreenInput {
	rv := objc.Send[CaptureScreenInput](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CaptureScreenInputClass) New() CaptureScreenInput {
	rv := objc.Send[CaptureScreenInput](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureScreenInput) Init() CaptureScreenInput {
	rv := objc.Send[CaptureScreenInput](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureScreenInput) Autorelease() CaptureScreenInput {
	rv := objc.Send[CaptureScreenInput](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureScreenInput creates a new CaptureScreenInput instance.
func NewCaptureScreenInput() CaptureScreenInput {
	return getCaptureScreenInputClass().New()
}




