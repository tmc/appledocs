// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTL4RenderPipelineColorAttachmentDescriptor */


/* debug [class_header]: Header for MTL4RenderPipelineColorAttachmentDescriptor */
// The class instance for the [MTL4RenderPipelineColorAttachmentDescriptor] class.
var (
	MTL4RenderPipelineColorAttachmentDescriptorClass     _MTL4RenderPipelineColorAttachmentDescriptorClass
	MTL4RenderPipelineColorAttachmentDescriptorClassOnce sync.Once
)

func getMTL4RenderPipelineColorAttachmentDescriptorClass() _MTL4RenderPipelineColorAttachmentDescriptorClass {
	MTL4RenderPipelineColorAttachmentDescriptorClassOnce.Do(func() {
		MTL4RenderPipelineColorAttachmentDescriptorClass = _MTL4RenderPipelineColorAttachmentDescriptorClass{objc.GetClass("MTL4RenderPipelineColorAttachmentDescriptor")}
	})
	return MTL4RenderPipelineColorAttachmentDescriptorClass
}

type _MTL4RenderPipelineColorAttachmentDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTL4RenderPipelineColorAttachmentDescriptor */
// An interface definition for the [MTL4RenderPipelineColorAttachmentDescriptor] class.
type IMTL4RenderPipelineColorAttachmentDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTL4RenderPipelineColorAttachmentDescriptor */
	// properties:
	AlphaBlendOperation() BlendOperation
	SetAlphaBlendOperation(value BlendOperation)
	BlendingState() MTL4BlendState
	SetBlendingState(value MTL4BlendState)
	DestinationAlphaBlendFactor() BlendFactor
	SetDestinationAlphaBlendFactor(value BlendFactor)
	DestinationRGBBlendFactor() BlendFactor
	SetDestinationRGBBlendFactor(value BlendFactor)
	PixelFormat() PixelFormat
	SetPixelFormat(value PixelFormat)
	RgbBlendOperation() BlendOperation
	SetRgbBlendOperation(value BlendOperation)
	SourceAlphaBlendFactor() BlendFactor
	SetSourceAlphaBlendFactor(value BlendFactor)
	SourceRGBBlendFactor() BlendFactor
	SetSourceRGBBlendFactor(value BlendFactor)
	WriteMask() ColorWriteMask
	SetWriteMask(value ColorWriteMask)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTL4RenderPipelineColorAttachmentDescriptor */
	// methods:
	Reset()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTL4RenderPipelineColorAttachmentDescriptor */
// Alloc allocates a new instance without initialization.
func (mc _MTL4RenderPipelineColorAttachmentDescriptorClass) Alloc() MTL4RenderPipelineColorAttachmentDescriptor {
	rv := objc.Send[MTL4RenderPipelineColorAttachmentDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTL4RenderPipelineColorAttachmentDescriptorClass) New() MTL4RenderPipelineColorAttachmentDescriptor {
	rv := objc.Send[MTL4RenderPipelineColorAttachmentDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) Init() MTL4RenderPipelineColorAttachmentDescriptor {
	rv := objc.Send[MTL4RenderPipelineColorAttachmentDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) Autorelease() MTL4RenderPipelineColorAttachmentDescriptor {
	rv := objc.Send[MTL4RenderPipelineColorAttachmentDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4RenderPipelineColorAttachmentDescriptor creates a new MTL4RenderPipelineColorAttachmentDescriptor instance.
func NewMTL4RenderPipelineColorAttachmentDescriptor() MTL4RenderPipelineColorAttachmentDescriptor {
	return getMTL4RenderPipelineColorAttachmentDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTL4RenderPipelineColorAttachmentDescriptor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor
type MTL4RenderPipelineColorAttachmentDescriptor struct {
	objectivec.Object
}

// MTL4RenderPipelineColorAttachmentDescriptorFrom constructs a [MTL4RenderPipelineColorAttachmentDescriptor] from an unsafe.Pointer.
func MTL4RenderPipelineColorAttachmentDescriptorFrom(ptr unsafe.Pointer) MTL4RenderPipelineColorAttachmentDescriptor {
	return MTL4RenderPipelineColorAttachmentDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTL4RenderPipelineColorAttachmentDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTL4RenderPipelineColorAttachmentDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTL4RenderPipelineColorAttachmentDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTL4RenderPipelineColorAttachmentDescriptor */

// Resets this descriptor to its default state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/reset()
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) Reset() {
	objc.Send[objc.ID](m_.ID, objc.Sel("reset"))
}/* debug [instance_methods/method]: Reset */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTL4RenderPipelineColorAttachmentDescriptor */

// Configures the alpha blending operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/alphaBlendOperation
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) AlphaBlendOperation() BlendOperation {
	rv := objc.Send[BlendOperation](m_.ID, objc.Sel("alphaBlendOperation"))
	return rv
}/* debug [instance_properties/getter]: alphaBlendOperation */


// Configures the alpha blending operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/alphaBlendOperation
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) SetAlphaBlendOperation(value BlendOperation) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlphaBlendOperation:"), value)
}/* debug [instance_properties/setter]: alphaBlendOperation */


// Configure the blend state for color attachments the pipeline state uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/blendingState
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) BlendingState() MTL4BlendState {
	rv := objc.Send[MTL4BlendState](m_.ID, objc.Sel("blendingState"))
	return rv
}/* debug [instance_properties/getter]: blendingState */


// Configure the blend state for color attachments the pipeline state uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/blendingState
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) SetBlendingState(value MTL4BlendState) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBlendingState:"), value)
}/* debug [instance_properties/setter]: blendingState */


// Configures the destination-alpha blend factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/destinationAlphaBlendFactor
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) DestinationAlphaBlendFactor() BlendFactor {
	rv := objc.Send[BlendFactor](m_.ID, objc.Sel("destinationAlphaBlendFactor"))
	return rv
}/* debug [instance_properties/getter]: destinationAlphaBlendFactor */


// Configures the destination-alpha blend factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/destinationAlphaBlendFactor
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) SetDestinationAlphaBlendFactor(value BlendFactor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDestinationAlphaBlendFactor:"), value)
}/* debug [instance_properties/setter]: destinationAlphaBlendFactor */


// Configures the destination RGB blend factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/destinationRGBBlendFactor
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) DestinationRGBBlendFactor() BlendFactor {
	rv := objc.Send[BlendFactor](m_.ID, objc.Sel("destinationRGBBlendFactor"))
	return rv
}/* debug [instance_properties/getter]: destinationRGBBlendFactor */


// Configures the destination RGB blend factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/destinationRGBBlendFactor
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) SetDestinationRGBBlendFactor(value BlendFactor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDestinationRGBBlendFactor:"), value)
}/* debug [instance_properties/setter]: destinationRGBBlendFactor */


// Configures the pixel format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/pixelFormat
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) PixelFormat() PixelFormat {
	rv := objc.Send[PixelFormat](m_.ID, objc.Sel("pixelFormat"))
	return rv
}/* debug [instance_properties/getter]: pixelFormat */


// Configures the pixel format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/pixelFormat
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) SetPixelFormat(value PixelFormat) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPixelFormat:"), value)
}/* debug [instance_properties/setter]: pixelFormat */


// Configures the RGB blend operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/rgbBlendOperation
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) RgbBlendOperation() BlendOperation {
	rv := objc.Send[BlendOperation](m_.ID, objc.Sel("rgbBlendOperation"))
	return rv
}/* debug [instance_properties/getter]: rgbBlendOperation */


// Configures the RGB blend operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/rgbBlendOperation
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) SetRgbBlendOperation(value BlendOperation) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRgbBlendOperation:"), value)
}/* debug [instance_properties/setter]: rgbBlendOperation */


// Configures the source-alpha blend factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/sourceAlphaBlendFactor
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) SourceAlphaBlendFactor() BlendFactor {
	rv := objc.Send[BlendFactor](m_.ID, objc.Sel("sourceAlphaBlendFactor"))
	return rv
}/* debug [instance_properties/getter]: sourceAlphaBlendFactor */


// Configures the source-alpha blend factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/sourceAlphaBlendFactor
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) SetSourceAlphaBlendFactor(value BlendFactor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceAlphaBlendFactor:"), value)
}/* debug [instance_properties/setter]: sourceAlphaBlendFactor */


// Configures the source RGB blend factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/sourceRGBBlendFactor
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) SourceRGBBlendFactor() BlendFactor {
	rv := objc.Send[BlendFactor](m_.ID, objc.Sel("sourceRGBBlendFactor"))
	return rv
}/* debug [instance_properties/getter]: sourceRGBBlendFactor */


// Configures the source RGB blend factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/sourceRGBBlendFactor
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) SetSourceRGBBlendFactor(value BlendFactor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceRGBBlendFactor:"), value)
}/* debug [instance_properties/setter]: sourceRGBBlendFactor */


// Configures the color write mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/writeMask
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) WriteMask() ColorWriteMask {
	rv := objc.Send[ColorWriteMask](m_.ID, objc.Sel("writeMask"))
	return rv
}/* debug [instance_properties/getter]: writeMask */


// Configures the color write mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/writeMask
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) SetWriteMask(value ColorWriteMask) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWriteMask:"), value)
}/* debug [instance_properties/setter]: writeMask */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTL4RenderPipelineColorAttachmentDescriptor */



