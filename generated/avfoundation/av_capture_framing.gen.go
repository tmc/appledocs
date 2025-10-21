// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CaptureFraming] class.
var (
	CaptureFramingClass     _CaptureFramingClass
	CaptureFramingClassOnce sync.Once
)

func getCaptureFramingClass() _CaptureFramingClass {
	CaptureFramingClassOnce.Do(func() {
		CaptureFramingClass = _CaptureFramingClass{objc.GetClass("AVCaptureFraming")}
	})
	return CaptureFramingClass
}

type _CaptureFramingClass struct {
	class objc.Class
}

// An interface definition for the [CaptureFraming] class.
type ICaptureFraming interface {
	objectivec.IObject
}

// A framing, consisting of an aspect ratio and a zoom factor.
//
// An provides framing recommendations using this object.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureFraming
type CaptureFraming struct {
	objectivec.Object
}

// CaptureFramingFrom constructs a [CaptureFraming] from an unsafe.Pointer.
//
// A framing, consisting of an aspect ratio and a zoom factor.
func CaptureFramingFrom(ptr unsafe.Pointer) CaptureFraming {
	return CaptureFraming{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CaptureFramingClass) Alloc() CaptureFraming {
	rv := objc.Send[CaptureFraming](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CaptureFramingClass) New() CaptureFraming {
	rv := objc.Send[CaptureFraming](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureFraming) Init() CaptureFraming {
	rv := objc.Send[CaptureFraming](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureFraming) Autorelease() CaptureFraming {
	rv := objc.Send[CaptureFraming](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureFraming creates a new CaptureFraming instance.
func NewCaptureFraming() CaptureFraming {
	return getCaptureFramingClass().New()
}


// A zoom factor.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureFraming/zoomFactor
func (c_ CaptureFraming) ZoomFactor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("zoomFactor"))
	return rv
}



