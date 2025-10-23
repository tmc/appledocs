// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
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
	// properties:
	Connections() []CaptureConnection /* primitive/slice/pointer */
	DeferredStartEnabled() bool /* primitive/slice/pointer */
	SetDeferredStartEnabled(value bool /* primitive/slice/pointer */)
	DeferredStartSupported() bool /* primitive/slice/pointer */
	IsDeferredStartEnabled() bool /* primitive/slice/pointer */
	SetIsDeferredStartEnabled(value bool /* primitive/slice/pointer */)
	IsDeferredStartSupported() bool /* primitive/slice/pointer */
	SetIsDeferredStartSupported(value bool /* primitive/slice/pointer */)
	// methods:
	ConnectionWithMediaType(mediaType AVMediaType /* foo */) IAVCaptureConnection
	MetadataOutputRectOfInterestForRect(rectInOutputCoordinates coregraphics.CGRect) coregraphics.CGRect
	RectForMetadataOutputRectOfInterest(rectInMetadataOutputCoordinates coregraphics.CGRect) coregraphics.CGRect
	TransformedMetadataObjectForMetadataObjectConnection(metadataObject AVMetadataObject /* foo */, connection IAVCaptureConnection) AVMetadataObject /* foo */
}

// An abstract superclass for objects that provide media output destinations for a capture session.
//
// This class provides an abstract interface to connect capture output destinations, such as files and streams, to a capture session. A capture output can have multiple connections, one for each stream of media that it receives from a capture input. A capture output doesn’t have any connections when you create it. When you add it to a capture session, the session automatically forms connections between compatible inputs and outputs.


// An abstract superclass for objects that provide media output destinations for a capture session.
//
// [Full Topic]
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



// Returns the first connection with an input port of a specified media type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureOutput/connection(with:)
func (c_ CaptureOutput) ConnectionWithMediaType(mediaType AVMediaType /* foo */) IAVCaptureConnection {
	rv := objc.Send[CaptureConnection](c_.ID, objc.Sel("connectionWithMediaType:"), mediaType)
	return rv
}


// Converts a rectangle in the capture output object’s coordinate system to one in the coordinate system used for metadata outputs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureOutput/metadataOutputRectConverted(fromOutputRect:)
func (c_ CaptureOutput) MetadataOutputRectOfInterestForRect(rectInOutputCoordinates coregraphics.CGRect) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](c_.ID, objc.Sel("metadataOutputRectOfInterestForRect:"), rectInOutputCoordinates)
	return rv
}


// Converts a rectangle in the coordinate system used for metadata outputs to one in the capture output object’s coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureOutput/outputRectConverted(fromMetadataOutputRect:)
func (c_ CaptureOutput) RectForMetadataOutputRectOfInterest(rectInMetadataOutputCoordinates coregraphics.CGRect) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](c_.ID, objc.Sel("rectForMetadataOutputRectOfInterest:"), rectInMetadataOutputCoordinates)
	return rv
}


// Converts a metadata object’s visual properties to layer coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureOutput/transformedMetadataObject(for:connection:)
func (c_ CaptureOutput) TransformedMetadataObjectForMetadataObjectConnection(metadataObject AVMetadataObject /* foo */, connection IAVCaptureConnection) AVMetadataObject /* foo */ {
	rv := objc.Send[MetadataObject](c_.ID, objc.Sel("transformedMetadataObjectForMetadataObject:connection:"), metadataObject, connection)
	return rv
}


// The capture output object’s connections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureOutput/connections
func (c_ CaptureOutput) Connections() []CaptureConnection /* primitive/slice/pointer */ {
	rv := objc.Send[[]CaptureConnection](c_.ID, objc.Sel("connections"))
	return rv
}


// A Boolean value that indicates whether to defer starting this capture output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureOutput/isDeferredStartEnabled
func (c_ CaptureOutput) DeferredStartEnabled() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("deferredStartEnabled"))
	return rv
}


// A Boolean value that indicates whether to defer starting this capture output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureOutput/isDeferredStartEnabled
func (c_ CaptureOutput) SetDeferredStartEnabled(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDeferredStartEnabled:"), value)
}


// A value that indicates whether the output supports deferred start.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureOutput/isDeferredStartSupported
func (c_ CaptureOutput) DeferredStartSupported() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("deferredStartSupported"))
	return rv
}


// A Boolean value that indicates whether to defer starting this capture output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureoutput/isdeferredstartenabled
func (c_ CaptureOutput) IsDeferredStartEnabled() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isDeferredStartEnabled"))
	return rv
}


// A Boolean value that indicates whether to defer starting this capture output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureoutput/isdeferredstartenabled
func (c_ CaptureOutput) SetIsDeferredStartEnabled(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsDeferredStartEnabled:"), value)
}


// A
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureoutput/isdeferredstartsupported
func (c_ CaptureOutput) IsDeferredStartSupported() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isDeferredStartSupported"))
	return rv
}


// A
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureoutput/isdeferredstartsupported
func (c_ CaptureOutput) SetIsDeferredStartSupported(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsDeferredStartSupported:"), value)
}



