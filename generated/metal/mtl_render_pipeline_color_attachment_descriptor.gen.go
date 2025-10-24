// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLRenderPipelineColorAttachmentDescriptor */


/* debug [class_header]: Header for MTLRenderPipelineColorAttachmentDescriptor */
// The class instance for the [RenderPipelineColorAttachmentDescriptor] class.
var (
	RenderPipelineColorAttachmentDescriptorClass     _RenderPipelineColorAttachmentDescriptorClass
	RenderPipelineColorAttachmentDescriptorClassOnce sync.Once
)

func getRenderPipelineColorAttachmentDescriptorClass() _RenderPipelineColorAttachmentDescriptorClass {
	RenderPipelineColorAttachmentDescriptorClassOnce.Do(func() {
		RenderPipelineColorAttachmentDescriptorClass = _RenderPipelineColorAttachmentDescriptorClass{objc.GetClass("MTLRenderPipelineColorAttachmentDescriptor")}
	})
	return RenderPipelineColorAttachmentDescriptorClass
}

type _RenderPipelineColorAttachmentDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RenderPipelineColorAttachmentDescriptor */
// An interface definition for the [RenderPipelineColorAttachmentDescriptor] class.
type IRenderPipelineColorAttachmentDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for RenderPipelineColorAttachmentDescriptor */
	// properties:
	AlphaBlendOperation() BlendOperation
	SetAlphaBlendOperation(value BlendOperation)
	DestinationAlphaBlendFactor() BlendFactor
	SetDestinationAlphaBlendFactor(value BlendFactor)
	DestinationRGBBlendFactor() BlendFactor
	SetDestinationRGBBlendFactor(value BlendFactor)
	BlendingEnabled() bool
	SetBlendingEnabled(value bool)
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
	IsBlendingEnabled() bool
	SetIsBlendingEnabled(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RenderPipelineColorAttachmentDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RenderPipelineColorAttachmentDescriptor */
// Alloc allocates a new instance without initialization.
func (rc _RenderPipelineColorAttachmentDescriptorClass) Alloc() RenderPipelineColorAttachmentDescriptor {
	rv := objc.Send[RenderPipelineColorAttachmentDescriptor](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RenderPipelineColorAttachmentDescriptorClass) New() RenderPipelineColorAttachmentDescriptor {
	rv := objc.Send[RenderPipelineColorAttachmentDescriptor](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RenderPipelineColorAttachmentDescriptor) Init() RenderPipelineColorAttachmentDescriptor {
	rv := objc.Send[RenderPipelineColorAttachmentDescriptor](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RenderPipelineColorAttachmentDescriptor) Autorelease() RenderPipelineColorAttachmentDescriptor {
	rv := objc.Send[RenderPipelineColorAttachmentDescriptor](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRenderPipelineColorAttachmentDescriptor creates a new RenderPipelineColorAttachmentDescriptor instance.
func NewRenderPipelineColorAttachmentDescriptor() RenderPipelineColorAttachmentDescriptor {
	return getRenderPipelineColorAttachmentDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RenderPipelineColorAttachmentDescriptor */
// A color render target that specifies the color configuration and color operations for a render pipeline.
//
// An instance defines the configuration of a color attachment associated with a rendering pipeline. The property must be specified for the rendering pipeline state at the color attachment. Blend operations determine how a source fragment is combined with a destination value in a color attachment to determine the pixel value to be written. The following properties define whether and how blending is performed: The property enables blending. The default value is . The property identifies which color channels are blended. The default value is , which allows all color channels to be blended. The and properties assign the blend operations for RGB and alpha pixel data. The default value for both properties is . The , , , and properties assign the source and destination blend factors. The default value for and is . The default value for and is .


// A color render target that specifies the color configuration and color operations for a render pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor
type RenderPipelineColorAttachmentDescriptor struct {
	objectivec.Object
}

// RenderPipelineColorAttachmentDescriptorFrom constructs a [RenderPipelineColorAttachmentDescriptor] from an unsafe.Pointer.
//
// A color render target that specifies the color configuration and color operations for a render pipeline.
func RenderPipelineColorAttachmentDescriptorFrom(ptr unsafe.Pointer) RenderPipelineColorAttachmentDescriptor {
	return RenderPipelineColorAttachmentDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RenderPipelineColorAttachmentDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RenderPipelineColorAttachmentDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RenderPipelineColorAttachmentDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RenderPipelineColorAttachmentDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RenderPipelineColorAttachmentDescriptor */

// The blend operation assigned for the alpha data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor/alphaBlendOperation
func (r_ RenderPipelineColorAttachmentDescriptor) AlphaBlendOperation() BlendOperation {
	rv := objc.Send[BlendOperation](r_.ID, objc.Sel("alphaBlendOperation"))
	return rv
}/* debug [instance_properties/getter]: alphaBlendOperation */


// The blend operation assigned for the alpha data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor/alphaBlendOperation
func (r_ RenderPipelineColorAttachmentDescriptor) SetAlphaBlendOperation(value BlendOperation) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setAlphaBlendOperation:"), value)
}/* debug [instance_properties/setter]: alphaBlendOperation */


// The destination blend factor (DBF) used by the alpha blend operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor/destinationAlphaBlendFactor
func (r_ RenderPipelineColorAttachmentDescriptor) DestinationAlphaBlendFactor() BlendFactor {
	rv := objc.Send[BlendFactor](r_.ID, objc.Sel("destinationAlphaBlendFactor"))
	return rv
}/* debug [instance_properties/getter]: destinationAlphaBlendFactor */


// The destination blend factor (DBF) used by the alpha blend operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor/destinationAlphaBlendFactor
func (r_ RenderPipelineColorAttachmentDescriptor) SetDestinationAlphaBlendFactor(value BlendFactor) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setDestinationAlphaBlendFactor:"), value)
}/* debug [instance_properties/setter]: destinationAlphaBlendFactor */


// The destination blend factor (DBF) used by the RGB blend operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor/destinationRGBBlendFactor
func (r_ RenderPipelineColorAttachmentDescriptor) DestinationRGBBlendFactor() BlendFactor {
	rv := objc.Send[BlendFactor](r_.ID, objc.Sel("destinationRGBBlendFactor"))
	return rv
}/* debug [instance_properties/getter]: destinationRGBBlendFactor */


// The destination blend factor (DBF) used by the RGB blend operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor/destinationRGBBlendFactor
func (r_ RenderPipelineColorAttachmentDescriptor) SetDestinationRGBBlendFactor(value BlendFactor) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setDestinationRGBBlendFactor:"), value)
}/* debug [instance_properties/setter]: destinationRGBBlendFactor */


// A Boolean value that determines whether blending is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor/isBlendingEnabled
func (r_ RenderPipelineColorAttachmentDescriptor) BlendingEnabled() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("blendingEnabled"))
	return rv
}/* debug [instance_properties/getter]: blendingEnabled */


// A Boolean value that determines whether blending is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor/isBlendingEnabled
func (r_ RenderPipelineColorAttachmentDescriptor) SetBlendingEnabled(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setBlendingEnabled:"), value)
}/* debug [instance_properties/setter]: blendingEnabled */


// The pixel format of the color attachment’s texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor/pixelFormat
func (r_ RenderPipelineColorAttachmentDescriptor) PixelFormat() PixelFormat {
	rv := objc.Send[PixelFormat](r_.ID, objc.Sel("pixelFormat"))
	return rv
}/* debug [instance_properties/getter]: pixelFormat */


// The pixel format of the color attachment’s texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor/pixelFormat
func (r_ RenderPipelineColorAttachmentDescriptor) SetPixelFormat(value PixelFormat) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setPixelFormat:"), value)
}/* debug [instance_properties/setter]: pixelFormat */


// The blend operation assigned for the RGB data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor/rgbBlendOperation
func (r_ RenderPipelineColorAttachmentDescriptor) RgbBlendOperation() BlendOperation {
	rv := objc.Send[BlendOperation](r_.ID, objc.Sel("rgbBlendOperation"))
	return rv
}/* debug [instance_properties/getter]: rgbBlendOperation */


// The blend operation assigned for the RGB data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor/rgbBlendOperation
func (r_ RenderPipelineColorAttachmentDescriptor) SetRgbBlendOperation(value BlendOperation) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRgbBlendOperation:"), value)
}/* debug [instance_properties/setter]: rgbBlendOperation */


// The source blend factor (SBF) used by the alpha blend operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor/sourceAlphaBlendFactor
func (r_ RenderPipelineColorAttachmentDescriptor) SourceAlphaBlendFactor() BlendFactor {
	rv := objc.Send[BlendFactor](r_.ID, objc.Sel("sourceAlphaBlendFactor"))
	return rv
}/* debug [instance_properties/getter]: sourceAlphaBlendFactor */


// The source blend factor (SBF) used by the alpha blend operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor/sourceAlphaBlendFactor
func (r_ RenderPipelineColorAttachmentDescriptor) SetSourceAlphaBlendFactor(value BlendFactor) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setSourceAlphaBlendFactor:"), value)
}/* debug [instance_properties/setter]: sourceAlphaBlendFactor */


// The source blend factor (SBF) used by the RGB blend operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor/sourceRGBBlendFactor
func (r_ RenderPipelineColorAttachmentDescriptor) SourceRGBBlendFactor() BlendFactor {
	rv := objc.Send[BlendFactor](r_.ID, objc.Sel("sourceRGBBlendFactor"))
	return rv
}/* debug [instance_properties/getter]: sourceRGBBlendFactor */


// The source blend factor (SBF) used by the RGB blend operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor/sourceRGBBlendFactor
func (r_ RenderPipelineColorAttachmentDescriptor) SetSourceRGBBlendFactor(value BlendFactor) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setSourceRGBBlendFactor:"), value)
}/* debug [instance_properties/setter]: sourceRGBBlendFactor */


// A bitmask that restricts which color channels are written into the texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor/writeMask
func (r_ RenderPipelineColorAttachmentDescriptor) WriteMask() ColorWriteMask {
	rv := objc.Send[ColorWriteMask](r_.ID, objc.Sel("writeMask"))
	return rv
}/* debug [instance_properties/getter]: writeMask */


// A bitmask that restricts which color channels are written into the texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor/writeMask
func (r_ RenderPipelineColorAttachmentDescriptor) SetWriteMask(value ColorWriteMask) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setWriteMask:"), value)
}/* debug [instance_properties/setter]: writeMask */


// A Boolean value that determines whether blending is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinecolorattachmentdescriptor/isblendingenabled
func (r_ RenderPipelineColorAttachmentDescriptor) IsBlendingEnabled() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isBlendingEnabled"))
	return rv
}/* debug [instance_properties/getter]: isBlendingEnabled */


// A Boolean value that determines whether blending is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinecolorattachmentdescriptor/isblendingenabled
func (r_ RenderPipelineColorAttachmentDescriptor) SetIsBlendingEnabled(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsBlendingEnabled:"), value)
}/* debug [instance_properties/setter]: isBlendingEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLRenderPipelineColorAttachmentDescriptor */



