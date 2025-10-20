// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CaptureDepthDataOutput] class.
var (
	CaptureDepthDataOutputClass     _CaptureDepthDataOutputClass
	CaptureDepthDataOutputClassOnce sync.Once
)

func getCaptureDepthDataOutputClass() _CaptureDepthDataOutputClass {
	CaptureDepthDataOutputClassOnce.Do(func() {
		CaptureDepthDataOutputClass = _CaptureDepthDataOutputClass{objc.GetClass("AVCaptureDepthDataOutput")}
	})
	return CaptureDepthDataOutputClass
}

type _CaptureDepthDataOutputClass struct {
	class objc.Class
}

// An interface definition for the [CaptureDepthDataOutput] class.
type ICaptureDepthDataOutput interface {
	ICaptureOutput
}

// A capture output that records scene depth information on compatible camera devices.
//
// This output type captures objects containing per-pixel depth or disparity information, following a streaming delivery model similar to that used by . Alternatively, you can capture depth data alongside photos using (see the property). This object always provides depth data in the format expressed by the source object’s property. If you wish to receive depth data in another format, choose a new value for that property from those listed in the array of the device’s object.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDepthDataOutput
type CaptureDepthDataOutput struct {
	CaptureOutput
}

// CaptureDepthDataOutputFrom constructs a [CaptureDepthDataOutput] from an unsafe.Pointer.
//
// A capture output that records scene depth information on compatible camera devices.
func CaptureDepthDataOutputFrom(ptr unsafe.Pointer) CaptureDepthDataOutput {
	return CaptureDepthDataOutput{
		CaptureOutput: CaptureOutputFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CaptureDepthDataOutputClass) Alloc() CaptureDepthDataOutput {
	rv := objc.Send[CaptureDepthDataOutput](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CaptureDepthDataOutputClass) New() CaptureDepthDataOutput {
	rv := objc.Send[CaptureDepthDataOutput](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureDepthDataOutput) Init() CaptureDepthDataOutput {
	rv := objc.Send[CaptureDepthDataOutput](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureDepthDataOutput) Autorelease() CaptureDepthDataOutput {
	rv := objc.Send[CaptureDepthDataOutput](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureDepthDataOutput creates a new CaptureDepthDataOutput instance.
func NewCaptureDepthDataOutput() CaptureDepthDataOutput {
	return getCaptureDepthDataOutputClass().New()
}




