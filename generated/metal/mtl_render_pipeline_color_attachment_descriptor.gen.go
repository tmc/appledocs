// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [RenderPipelineColorAttachmentDescriptor] class.
type IRenderPipelineColorAttachmentDescriptor interface {
	objectivec.IObject
}

// A color render target that specifies the color configuration and color operations for a render pipeline.
//
// An instance defines the configuration of a color attachment associated with a rendering pipeline. The property must be specified for the rendering pipeline state at the color attachment. Blend operations determine how a source fragment is combined with a destination value in a color attachment to determine the pixel value to be written. The following properties define whether and how blending is performed: The property enables blending. The default value is . The property identifies which color channels are blended. The default value is , which allows all color channels to be blended. The and properties assign the blend operations for RGB and alpha pixel data. The default value for both properties is . The , , , and properties assign the source and destination blend factors. The default value for and is . The default value for and is .
//
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

// Alloc allocates a new instance without initialization.
func (rc _RenderPipelineColorAttachmentDescriptorClass) Alloc() RenderPipelineColorAttachmentDescriptor {
	rv := objc.Send[RenderPipelineColorAttachmentDescriptor](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The blend operation assigned for the alpha data.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor/alphaBlendOperation
func (r_ RenderPipelineColorAttachmentDescriptor) AlphaBlendOperation() BlendOperation {
	rv := objc.Send[BlendOperation](r_.ID, objc.Sel("alphaBlendOperation"))
	return rv
}


// SetAlphaBlendOperation sets the value of the alphaBlendOperation property.
// The blend operation assigned for the alpha data.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor/alphaBlendOperation
func (r_ RenderPipelineColorAttachmentDescriptor) SetAlphaBlendOperation(value IBlendOperation) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setAlphaBlendOperation:"), value)
}

// The destination blend factor (DBF) used by the alpha blend operation.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor/destinationAlphaBlendFactor
func (r_ RenderPipelineColorAttachmentDescriptor) DestinationAlphaBlendFactor() BlendFactor {
	rv := objc.Send[BlendFactor](r_.ID, objc.Sel("destinationAlphaBlendFactor"))
	return rv
}


// SetDestinationAlphaBlendFactor sets the value of the destinationAlphaBlendFactor property.
// The destination blend factor (DBF) used by the alpha blend operation.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor/destinationAlphaBlendFactor
func (r_ RenderPipelineColorAttachmentDescriptor) SetDestinationAlphaBlendFactor(value IBlendFactor) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setDestinationAlphaBlendFactor:"), value)
}

// The destination blend factor (DBF) used by the RGB blend operation.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor/destinationRGBBlendFactor
func (r_ RenderPipelineColorAttachmentDescriptor) DestinationRGBBlendFactor() BlendFactor {
	rv := objc.Send[BlendFactor](r_.ID, objc.Sel("destinationRGBBlendFactor"))
	return rv
}


// SetDestinationRGBBlendFactor sets the value of the destinationRGBBlendFactor property.
// The destination blend factor (DBF) used by the RGB blend operation.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor/destinationRGBBlendFactor
func (r_ RenderPipelineColorAttachmentDescriptor) SetDestinationRGBBlendFactor(value IBlendFactor) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setDestinationRGBBlendFactor:"), value)
}

// A Boolean value that determines whether blending is enabled.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor/isBlendingEnabled
func (r_ RenderPipelineColorAttachmentDescriptor) BlendingEnabled() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("blendingEnabled"))
	return rv
}


// SetBlendingEnabled sets the value of the blendingEnabled property.
// A Boolean value that determines whether blending is enabled.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor/isBlendingEnabled
func (r_ RenderPipelineColorAttachmentDescriptor) SetBlendingEnabled(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setBlendingEnabled:"), value)
}

// The pixel format of the color attachment’s texture.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor/pixelFormat
func (r_ RenderPipelineColorAttachmentDescriptor) PixelFormat() PixelFormat {
	rv := objc.Send[PixelFormat](r_.ID, objc.Sel("pixelFormat"))
	return rv
}


// SetPixelFormat sets the value of the pixelFormat property.
// The pixel format of the color attachment’s texture.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor/pixelFormat
func (r_ RenderPipelineColorAttachmentDescriptor) SetPixelFormat(value PixelFormat) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setPixelFormat:"), value)
}

// The blend operation assigned for the RGB data.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor/rgbBlendOperation
func (r_ RenderPipelineColorAttachmentDescriptor) RgbBlendOperation() BlendOperation {
	rv := objc.Send[BlendOperation](r_.ID, objc.Sel("rgbBlendOperation"))
	return rv
}


// SetRgbBlendOperation sets the value of the rgbBlendOperation property.
// The blend operation assigned for the RGB data.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor/rgbBlendOperation
func (r_ RenderPipelineColorAttachmentDescriptor) SetRgbBlendOperation(value IBlendOperation) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRgbBlendOperation:"), value)
}

// The source blend factor (SBF) used by the alpha blend operation.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor/sourceAlphaBlendFactor
func (r_ RenderPipelineColorAttachmentDescriptor) SourceAlphaBlendFactor() BlendFactor {
	rv := objc.Send[BlendFactor](r_.ID, objc.Sel("sourceAlphaBlendFactor"))
	return rv
}


// SetSourceAlphaBlendFactor sets the value of the sourceAlphaBlendFactor property.
// The source blend factor (SBF) used by the alpha blend operation.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor/sourceAlphaBlendFactor
func (r_ RenderPipelineColorAttachmentDescriptor) SetSourceAlphaBlendFactor(value IBlendFactor) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setSourceAlphaBlendFactor:"), value)
}

// The source blend factor (SBF) used by the RGB blend operation.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor/sourceRGBBlendFactor
func (r_ RenderPipelineColorAttachmentDescriptor) SourceRGBBlendFactor() BlendFactor {
	rv := objc.Send[BlendFactor](r_.ID, objc.Sel("sourceRGBBlendFactor"))
	return rv
}


// SetSourceRGBBlendFactor sets the value of the sourceRGBBlendFactor property.
// The source blend factor (SBF) used by the RGB blend operation.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor/sourceRGBBlendFactor
func (r_ RenderPipelineColorAttachmentDescriptor) SetSourceRGBBlendFactor(value IBlendFactor) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setSourceRGBBlendFactor:"), value)
}

// A bitmask that restricts which color channels are written into the texture.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor/writeMask
func (r_ RenderPipelineColorAttachmentDescriptor) WriteMask() ColorWriteMask {
	rv := objc.Send[ColorWriteMask](r_.ID, objc.Sel("writeMask"))
	return rv
}


// SetWriteMask sets the value of the writeMask property.
// A bitmask that restricts which color channels are written into the texture.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor/writeMask
func (r_ RenderPipelineColorAttachmentDescriptor) SetWriteMask(value ColorWriteMask) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setWriteMask:"), value)
}

// A Boolean value that determines whether blending is enabled.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinecolorattachmentdescriptor/isblendingenabled
func (r_ RenderPipelineColorAttachmentDescriptor) IsBlendingEnabled() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isBlendingEnabled"))
	return rv
}


// SetIsBlendingEnabled sets the value of the isBlendingEnabled property.
// A Boolean value that determines whether blending is enabled.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinecolorattachmentdescriptor/isblendingenabled
func (r_ RenderPipelineColorAttachmentDescriptor) SetIsBlendingEnabled(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsBlendingEnabled:"), value)
}



