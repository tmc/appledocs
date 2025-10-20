// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CaptureMetadataOutput] class.
var (
	CaptureMetadataOutputClass     _CaptureMetadataOutputClass
	CaptureMetadataOutputClassOnce sync.Once
)

func getCaptureMetadataOutputClass() _CaptureMetadataOutputClass {
	CaptureMetadataOutputClassOnce.Do(func() {
		CaptureMetadataOutputClass = _CaptureMetadataOutputClass{objc.GetClass("AVCaptureMetadataOutput")}
	})
	return CaptureMetadataOutputClass
}

type _CaptureMetadataOutputClass struct {
	class objc.Class
}

// An interface definition for the [CaptureMetadataOutput] class.
type ICaptureMetadataOutput interface {
	ICaptureOutput
}

// A capture output for processing timed metadata produced by a capture session.
//
// An object intercepts metadata objects emitted by its associated capture connection and forwards them to a delegate object for processing. You can use instances of this class to process specific types of metadata included with the input data. You use this class the way you do other output objects, typically by adding it as an output to an object.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMetadataOutput
type CaptureMetadataOutput struct {
	CaptureOutput
}

// CaptureMetadataOutputFrom constructs a [CaptureMetadataOutput] from an unsafe.Pointer.
//
// A capture output for processing timed metadata produced by a capture session.
func CaptureMetadataOutputFrom(ptr unsafe.Pointer) CaptureMetadataOutput {
	return CaptureMetadataOutput{
		CaptureOutput: CaptureOutputFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CaptureMetadataOutputClass) Alloc() CaptureMetadataOutput {
	rv := objc.Send[CaptureMetadataOutput](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CaptureMetadataOutputClass) New() CaptureMetadataOutput {
	rv := objc.Send[CaptureMetadataOutput](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureMetadataOutput) Init() CaptureMetadataOutput {
	rv := objc.Send[CaptureMetadataOutput](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureMetadataOutput) Autorelease() CaptureMetadataOutput {
	rv := objc.Send[CaptureMetadataOutput](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureMetadataOutput creates a new CaptureMetadataOutput instance.
func NewCaptureMetadataOutput() CaptureMetadataOutput {
	return getCaptureMetadataOutputClass().New()
}




