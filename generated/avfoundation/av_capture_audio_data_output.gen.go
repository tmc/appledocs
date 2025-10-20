// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CaptureAudioDataOutput] class.
var (
	CaptureAudioDataOutputClass     _CaptureAudioDataOutputClass
	CaptureAudioDataOutputClassOnce sync.Once
)

func getCaptureAudioDataOutputClass() _CaptureAudioDataOutputClass {
	CaptureAudioDataOutputClassOnce.Do(func() {
		CaptureAudioDataOutputClass = _CaptureAudioDataOutputClass{objc.GetClass("AVCaptureAudioDataOutput")}
	})
	return CaptureAudioDataOutputClass
}

type _CaptureAudioDataOutputClass struct {
	class objc.Class
}

// An interface definition for the [CaptureAudioDataOutput] class.
type ICaptureAudioDataOutput interface {
	ICaptureOutput
}

// A capture output that records audio and provides access to audio sample buffers as they are recorded.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioDataOutput
type CaptureAudioDataOutput struct {
	CaptureOutput
}

// CaptureAudioDataOutputFrom constructs a [CaptureAudioDataOutput] from an unsafe.Pointer.
//
// A capture output that records audio and provides access to audio sample buffers as they are recorded.
func CaptureAudioDataOutputFrom(ptr unsafe.Pointer) CaptureAudioDataOutput {
	return CaptureAudioDataOutput{
		CaptureOutput: CaptureOutputFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CaptureAudioDataOutputClass) Alloc() CaptureAudioDataOutput {
	rv := objc.Send[CaptureAudioDataOutput](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CaptureAudioDataOutputClass) New() CaptureAudioDataOutput {
	rv := objc.Send[CaptureAudioDataOutput](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureAudioDataOutput) Init() CaptureAudioDataOutput {
	rv := objc.Send[CaptureAudioDataOutput](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureAudioDataOutput) Autorelease() CaptureAudioDataOutput {
	rv := objc.Send[CaptureAudioDataOutput](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureAudioDataOutput creates a new CaptureAudioDataOutput instance.
func NewCaptureAudioDataOutput() CaptureAudioDataOutput {
	return getCaptureAudioDataOutputClass().New()
}




