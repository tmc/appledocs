// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CaptureOutput] class.
var (
	CaptureOutputClass     _CaptureOutputClass
	CaptureOutputClassOnce sync.Once
)

func getCaptureOutputClass() _CaptureOutputClass {
	CaptureOutputClassOnce.Do(func() {
		CaptureOutputClass = _CaptureOutputClass{objc.GetClass("AVCaptureOutput")}
	})
	return CaptureOutputClass
}

type _CaptureOutputClass struct {
	class objc.Class
}

// An interface definition for the [CaptureOutput] class.
type ICaptureOutput interface {
	objectivec.IObject
	Connections() AVCaptureConnection
	SetConnections(value IAVCaptureConnection)
	IsDeferredStartEnabled() bool
	SetIsDeferredStartEnabled(value bool)
	IsDeferredStartSupported() bool
	SetIsDeferredStartSupported(value bool)
}

// An abstract superclass for objects that provide media output destinations for a capture session.
//
// This class provides an abstract interface to connect capture output destinations, such as files and streams, to a capture session. A capture output can have multiple connections, one for each stream of media that it receives from a capture input. A capture output doesn’t have any connections when you create it. When you add it to a capture session, the session automatically forms connections between compatible inputs and outputs.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureOutput
type CaptureOutput struct {
	objectivec.Object
}

// CaptureOutputFrom constructs a [CaptureOutput] from an unsafe.Pointer.
//
// An abstract superclass for objects that provide media output destinations for a capture session.
func CaptureOutputFrom(ptr unsafe.Pointer) CaptureOutput {
	return CaptureOutput{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CaptureOutputClass) Alloc() CaptureOutput {
	rv := objc.Send[CaptureOutput](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CaptureOutputClass) New() CaptureOutput {
	rv := objc.Send[CaptureOutput](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureOutput) Init() CaptureOutput {
	rv := objc.Send[CaptureOutput](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureOutput) Autorelease() CaptureOutput {
	rv := objc.Send[CaptureOutput](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureOutput creates a new CaptureOutput instance.
func NewCaptureOutput() CaptureOutput {
	return getCaptureOutputClass().New()
}


// The capture output object’s connections.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureoutput/connections
func (c_ CaptureOutput) Connections() AVCaptureConnection {
	rv := objc.Send[AVCaptureConnection](c_.ID, objc.Sel("connections"))
	return rv
}


// SetConnections sets the value of the connections property.
// The capture output object’s connections.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureoutput/connections
func (c_ CaptureOutput) SetConnections(value IAVCaptureConnection) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConnections:"), value)
}

// A Boolean value that indicates whether to defer starting this capture output.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureoutput/isdeferredstartenabled
func (c_ CaptureOutput) IsDeferredStartEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isDeferredStartEnabled"))
	return rv
}


// SetIsDeferredStartEnabled sets the value of the isDeferredStartEnabled property.
// A Boolean value that indicates whether to defer starting this capture output.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureoutput/isdeferredstartenabled
func (c_ CaptureOutput) SetIsDeferredStartEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsDeferredStartEnabled:"), value)
}

// A
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureoutput/isdeferredstartsupported
func (c_ CaptureOutput) IsDeferredStartSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isDeferredStartSupported"))
	return rv
}


// SetIsDeferredStartSupported sets the value of the isDeferredStartSupported property.
// A

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureoutput/isdeferredstartsupported
func (c_ CaptureOutput) SetIsDeferredStartSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsDeferredStartSupported:"), value)
}



