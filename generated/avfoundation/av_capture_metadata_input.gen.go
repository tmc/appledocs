// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CaptureMetadataInput] class.
var (
	CaptureMetadataInputClass     _CaptureMetadataInputClass
	CaptureMetadataInputClassOnce sync.Once
)

func getCaptureMetadataInputClass() _CaptureMetadataInputClass {
	CaptureMetadataInputClassOnce.Do(func() {
		CaptureMetadataInputClass = _CaptureMetadataInputClass{objc.GetClass("AVCaptureMetadataInput")}
	})
	return CaptureMetadataInputClass
}

type _CaptureMetadataInputClass struct {
	class objc.Class
}

// An interface definition for the [CaptureMetadataInput] class.
type ICaptureMetadataInput interface {
	ICaptureInput
}

// A capture input for providing timed metadata to a capture session.
//
// This class provides input to an . An instance of can present one and only one connected to an . Provide metadata through the input port by conforming to a and supplying objects in an .


// A capture input for providing timed metadata to a capture session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMetadataInput
type CaptureMetadataInput struct {
	CaptureInput
}

// CaptureMetadataInputFrom constructs a [CaptureMetadataInput] from an unsafe.Pointer.
//
// A capture input for providing timed metadata to a capture session.
func CaptureMetadataInputFrom(ptr unsafe.Pointer) CaptureMetadataInput {
	return CaptureMetadataInput{
		CaptureInput: CaptureInputFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CaptureMetadataInputClass) Alloc() CaptureMetadataInput {
	rv := objc.Send[CaptureMetadataInput](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CaptureMetadataInputClass) New() CaptureMetadataInput {
	rv := objc.Send[CaptureMetadataInput](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureMetadataInput) Init() CaptureMetadataInput {
	rv := objc.Send[CaptureMetadataInput](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureMetadataInput) Autorelease() CaptureMetadataInput {
	rv := objc.Send[CaptureMetadataInput](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureMetadataInput creates a new CaptureMetadataInput instance.
func NewCaptureMetadataInput() CaptureMetadataInput {
	return getCaptureMetadataInputClass().New()
}




