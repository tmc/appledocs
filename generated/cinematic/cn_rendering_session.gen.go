// Code generated from Apple documentation for Cinematic. DO NOT EDIT.

package cinematic

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CNRenderingSession */


/* debug [class_header]: Header for CNRenderingSession */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNRenderingSession */
// An interface definition for the [CNRenderingSession] class.
type ICNRenderingSession interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CNRenderingSession */
	// properties:
	CommandQueue() unsafe.Pointer
	PreferredTransform() corefoundation.CGAffineTransform
	Quality() CNRenderingQuality
	SessionAttributes() ICNRenderingSessionAttributes
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNRenderingSession */
	// methods:
	EncodeRenderToCommandBufferFrameAttributesSourceImageSourceDisparityDestinationImage(commandBuffer unsafe.Pointer, frameAttributes ICNRenderingSessionFrameAttributes, sourceImage PixelBufferRef /* not a class type */, sourceDisparity PixelBufferRef /* not a class type */, destinationImage PixelBufferRef /* not a class type */) bool
	EncodeRenderToCommandBufferFrameAttributesSourceImageSourceDisparityDestinationLumaDestinationChroma(commandBuffer unsafe.Pointer, frameAttributes ICNRenderingSessionFrameAttributes, sourceImage PixelBufferRef /* not a class type */, sourceDisparity PixelBufferRef /* not a class type */, destinationLuma unsafe.Pointer, destinationChroma unsafe.Pointer) bool
	EncodeRenderToCommandBufferFrameAttributesSourceImageSourceDisparityDestinationRGBA(commandBuffer unsafe.Pointer, frameAttributes ICNRenderingSessionFrameAttributes, sourceImage PixelBufferRef /* not a class type */, sourceDisparity PixelBufferRef /* not a class type */, destinationRGBA unsafe.Pointer) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNRenderingSession */
// Alloc allocates a new instance without initialization.
func (cc _CNRenderingSessionClass) Alloc() CNRenderingSession {
	rv := objc.Send[CNRenderingSession](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNRenderingSession */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNRenderingSession */

// Intializes an object for a rendering session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNRenderingSession-8yghc/initWithCommandQueue:sessionAttributes:preferredTransform:quality:
func NewCNRenderingSessionWithCommandQueueSessionAttributesPreferredTransformQuality(commandQueue unsafe.Pointer, sessionAttributes ICNRenderingSessionAttributes, preferredTransform corefoundation.CGAffineTransform, quality CNRenderingQuality) CNRenderingSession {
	instance := getCNRenderingSessionClass().Alloc()
	rv := objc.Send[CNRenderingSession](instance.ID, objc.Sel("initWithCommandQueue:sessionAttributes:preferredTransform:quality:"), commandQueue, sessionAttributes, preferredTransform, quality)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNRenderingSessionWithCommandQueueSessionAttributesPreferredTransformQuality */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNRenderingSession */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNRenderingSession */

// A static number representing the video compositor’s required pixel buffer attributes context dictionary when implementing video compositing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNRenderingSession-8yghc/destinationPixelFormatTypes
func (cc _CNRenderingSessionClass) DestinationPixelFormatTypes() []foundation.Number {
	rv := objc.Send[[]foundation.Number](objc.ID(cc.class), objc.Sel("destinationPixelFormatTypes"))
	return rv
}/* debug [class_properties_class/property]: destinationPixelFormatTypes */

// The pixel format types supported for the output destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNRenderingSession-8yghc/sourcePixelFormatTypes
func (cc _CNRenderingSessionClass) SourcePixelFormatTypes() []foundation.Number {
	rv := objc.Send[[]foundation.Number](objc.ID(cc.class), objc.Sel("sourcePixelFormatTypes"))
	return rv
}/* debug [class_properties_class/property]: sourcePixelFormatTypes */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNRenderingSession */

// Encodes a command to render a shallow depth of field (SDoF) image to a pixel buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNRenderingSession-8yghc/encodeRenderToCommandBuffer:frameAttributes:sourceImage:sourceDisparity:destinationImage:
func (c_ CNRenderingSession) EncodeRenderToCommandBufferFrameAttributesSourceImageSourceDisparityDestinationImage(commandBuffer unsafe.Pointer, frameAttributes ICNRenderingSessionFrameAttributes, sourceImage PixelBufferRef /* not a class type */, sourceDisparity PixelBufferRef /* not a class type */, destinationImage PixelBufferRef /* not a class type */) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("encodeRenderToCommandBuffer:frameAttributes:sourceImage:sourceDisparity:destinationImage:"), commandBuffer, frameAttributes, sourceImage, sourceDisparity, destinationImage)
	return rv
}/* debug [instance_methods/method]: EncodeRenderToCommandBufferFrameAttributesSourceImageSourceDisparityDestinationImage */


// Encodes a command to render a shallow depth of field (SDoF) image to two metal textures as luma and chroma.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNRenderingSession-8yghc/encodeRenderToCommandBuffer:frameAttributes:sourceImage:sourceDisparity:destinationLuma:destinationChroma:
func (c_ CNRenderingSession) EncodeRenderToCommandBufferFrameAttributesSourceImageSourceDisparityDestinationLumaDestinationChroma(commandBuffer unsafe.Pointer, frameAttributes ICNRenderingSessionFrameAttributes, sourceImage PixelBufferRef /* not a class type */, sourceDisparity PixelBufferRef /* not a class type */, destinationLuma unsafe.Pointer, destinationChroma unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("encodeRenderToCommandBuffer:frameAttributes:sourceImage:sourceDisparity:destinationLuma:destinationChroma:"), commandBuffer, frameAttributes, sourceImage, sourceDisparity, destinationLuma, destinationChroma)
	return rv
}/* debug [instance_methods/method]: EncodeRenderToCommandBufferFrameAttributesSourceImageSourceDisparityDestinationLumaDestinationChroma */


// Encodes a command to render a shallow depth of field (SDoF) image to a metal texture as RGBA.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNRenderingSession-8yghc/encodeRenderToCommandBuffer:frameAttributes:sourceImage:sourceDisparity:destinationRGBA:
func (c_ CNRenderingSession) EncodeRenderToCommandBufferFrameAttributesSourceImageSourceDisparityDestinationRGBA(commandBuffer unsafe.Pointer, frameAttributes ICNRenderingSessionFrameAttributes, sourceImage PixelBufferRef /* not a class type */, sourceDisparity PixelBufferRef /* not a class type */, destinationRGBA unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("encodeRenderToCommandBuffer:frameAttributes:sourceImage:sourceDisparity:destinationRGBA:"), commandBuffer, frameAttributes, sourceImage, sourceDisparity, destinationRGBA)
	return rv
}/* debug [instance_methods/method]: EncodeRenderToCommandBufferFrameAttributesSourceImageSourceDisparityDestinationRGBA */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNRenderingSession */

// The command queue of a Metal device that creates the command buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNRenderingSession-8yghc/commandQueue
func (c_ CNRenderingSession) CommandQueue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("commandQueue"))
	return rv
}/* debug [instance_properties/getter]: commandQueue */


// A static number representing the video compositor’s required pixel buffer attributes context dictionary when implementing video compositing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNRenderingSession-8yghc/destinationPixelFormatTypes
func (c_ CNRenderingSession) DestinationPixelFormatTypes() []foundation.Number {
	rv := objc.Send[[]foundation.Number](c_.ID, objc.Sel("destinationPixelFormatTypes"))
	return rv
}/* debug [instance_properties/getter]: destinationPixelFormatTypes */


// The preferred transform of the rendered image for display purposes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNRenderingSession-8yghc/preferredTransform
func (c_ CNRenderingSession) PreferredTransform() corefoundation.CGAffineTransform {
	rv := objc.Send[corefoundation.CGAffineTransform](c_.ID, objc.Sel("preferredTransform"))
	return rv
}/* debug [instance_properties/getter]: preferredTransform */


// The quality of rendering desired for a session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNRenderingSession-8yghc/quality
func (c_ CNRenderingSession) Quality() CNRenderingQuality {
	rv := objc.Send[CNRenderingQuality](c_.ID, objc.Sel("quality"))
	return rv
}/* debug [instance_properties/getter]: quality */


// Rendering session attributes for a Cinematic asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNRenderingSession-8yghc/sessionAttributes
func (c_ CNRenderingSession) SessionAttributes() ICNRenderingSessionAttributes {
	rv := objc.Send[CNRenderingSessionAttributes](c_.ID, objc.Sel("sessionAttributes"))
	return rv
}/* debug [instance_properties/getter]: sessionAttributes */


// The pixel format types supported for the output destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNRenderingSession-8yghc/sourcePixelFormatTypes
func (c_ CNRenderingSession) SourcePixelFormatTypes() []foundation.Number {
	rv := objc.Send[[]foundation.Number](c_.ID, objc.Sel("sourcePixelFormatTypes"))
	return rv
}/* debug [instance_properties/getter]: sourcePixelFormatTypes */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNRenderingSession */


