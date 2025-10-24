// Code generated from Apple documentation for Cinematic. DO NOT EDIT.

package cinematic

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CNRenderingSession] class.
var (
	CNRenderingSessionClass     _CNRenderingSessionClass
	CNRenderingSessionClassOnce sync.Once
)

func getCNRenderingSessionClass() _CNRenderingSessionClass {
	CNRenderingSessionClassOnce.Do(func() {
		CNRenderingSessionClass = _CNRenderingSessionClass{objc.GetClass("CNRenderingSession")}
	})
	return CNRenderingSessionClass
}

type _CNRenderingSessionClass struct {
	class objc.Class
}

// An interface definition for the [CNRenderingSession] class.
type ICNRenderingSession interface {
	objectivec.IObject
	// properties:
	// methods:
	EncodeRenderToCommandBufferFrameAttributesSourceImageSourceDisparityDestinationLumaDestinationChroma(commandBuffer objectivec.IObject, frameAttributes ICNRenderingSessionFrameAttributes, sourceImage PixelBufferRef /* not a class type */, sourceDisparity PixelBufferRef /* not a class type */, destinationLuma objectivec.IObject, destinationChroma objectivec.IObject) bool
}

// An object representing the context in which rendering occurs.


// An object representing the context in which rendering occurs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNRenderingSession-8yghc
type CNRenderingSession struct {
	objectivec.Object
}

// CNRenderingSessionFrom constructs a [CNRenderingSession] from an unsafe.Pointer.
//
// An object representing the context in which rendering occurs.
func CNRenderingSessionFrom(ptr unsafe.Pointer) CNRenderingSession {
	return CNRenderingSession{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CNRenderingSessionClass) Alloc() CNRenderingSession {
	rv := objc.Send[CNRenderingSession](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNRenderingSessionClass) New() CNRenderingSession {
	rv := objc.Send[CNRenderingSession](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNRenderingSession) Init() CNRenderingSession {
	rv := objc.Send[CNRenderingSession](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNRenderingSession) Autorelease() CNRenderingSession {
	rv := objc.Send[CNRenderingSession](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNRenderingSession creates a new CNRenderingSession instance.
func NewCNRenderingSession() CNRenderingSession {
	return getCNRenderingSessionClass().New()
}



// The pixel format types supported for the output destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNRenderingSession-8yghc/sourcePixelFormatTypes
func (cc _CNRenderingSessionClass) SourcePixelFormatTypes() []objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[[]foundation.Number](objc.ID(cc.class), objc.Sel("sourcePixelFormatTypes"))
	return rv
}

// Encodes a command to render a shallow depth of field (SDoF) image to two metal textures as luma and chroma.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNRenderingSession-8yghc/encodeRenderToCommandBuffer:frameAttributes:sourceImage:sourceDisparity:destinationLuma:destinationChroma:
func (c_ CNRenderingSession) EncodeRenderToCommandBufferFrameAttributesSourceImageSourceDisparityDestinationLumaDestinationChroma(commandBuffer objectivec.IObject, frameAttributes ICNRenderingSessionFrameAttributes, sourceImage PixelBufferRef /* not a class type */, sourceDisparity PixelBufferRef /* not a class type */, destinationLuma objectivec.IObject, destinationChroma objectivec.IObject) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("encodeRenderToCommandBuffer:frameAttributes:sourceImage:sourceDisparity:destinationLuma:destinationChroma:"), commandBuffer, frameAttributes, sourceImage, sourceDisparity, destinationLuma, destinationChroma)
	return rv
}


// The pixel format types supported for the output destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNRenderingSession-8yghc/sourcePixelFormatTypes
func (c_ CNRenderingSession) SourcePixelFormatTypes() []objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[[]foundation.Number](c_.ID, objc.Sel("sourcePixelFormatTypes"))
	return rv
}



