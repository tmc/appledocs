// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [MTL4RenderPipelineColorAttachmentDescriptor] class.
type IMTL4RenderPipelineColorAttachmentDescriptor interface {
	objectivec.IObject
	

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


	

	// methods:
	Reset()


}





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







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor
type MTL4RenderPipelineColorAttachmentDescriptor struct {
	objectivec.Object
}

// MTL4RenderPipelineColorAttachmentDescriptorFrom constructs a [MTL4RenderPipelineColorAttachmentDescriptor] from an unsafe.Pointer.
func MTL4RenderPipelineColorAttachmentDescriptorFrom(ptr unsafe.Pointer) MTL4RenderPipelineColorAttachmentDescriptor {
	return MTL4RenderPipelineColorAttachmentDescriptor{objectivec.Object{objc.ID(ptr)}}
}




















// Resets this descriptor to its default state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/reset()
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) Reset() {
	objc.Send[objc.ID](m_.ID, objc.Sel("reset"))
}







// Configures the alpha blending operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/alphaBlendOperation
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) AlphaBlendOperation() BlendOperation {
	rv := objc.Send[BlendOperation](m_.ID, objc.Sel("alphaBlendOperation"))
	return rv
}


// Configures the alpha blending operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/alphaBlendOperation
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) SetAlphaBlendOperation(value BlendOperation) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlphaBlendOperation:"), value)
}


// Configure the blend state for color attachments the pipeline state uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/blendingState
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) BlendingState() MTL4BlendState {
	rv := objc.Send[MTL4BlendState](m_.ID, objc.Sel("blendingState"))
	return rv
}


// Configure the blend state for color attachments the pipeline state uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/blendingState
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) SetBlendingState(value MTL4BlendState) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBlendingState:"), value)
}


// Configures the destination-alpha blend factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/destinationAlphaBlendFactor
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) DestinationAlphaBlendFactor() BlendFactor {
	rv := objc.Send[BlendFactor](m_.ID, objc.Sel("destinationAlphaBlendFactor"))
	return rv
}


// Configures the destination-alpha blend factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/destinationAlphaBlendFactor
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) SetDestinationAlphaBlendFactor(value BlendFactor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDestinationAlphaBlendFactor:"), value)
}


// Configures the destination RGB blend factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/destinationRGBBlendFactor
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) DestinationRGBBlendFactor() BlendFactor {
	rv := objc.Send[BlendFactor](m_.ID, objc.Sel("destinationRGBBlendFactor"))
	return rv
}


// Configures the destination RGB blend factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/destinationRGBBlendFactor
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) SetDestinationRGBBlendFactor(value BlendFactor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDestinationRGBBlendFactor:"), value)
}


// Configures the pixel format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/pixelFormat
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) PixelFormat() PixelFormat {
	rv := objc.Send[PixelFormat](m_.ID, objc.Sel("pixelFormat"))
	return rv
}


// Configures the pixel format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/pixelFormat
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) SetPixelFormat(value PixelFormat) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPixelFormat:"), value)
}


// Configures the RGB blend operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/rgbBlendOperation
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) RgbBlendOperation() BlendOperation {
	rv := objc.Send[BlendOperation](m_.ID, objc.Sel("rgbBlendOperation"))
	return rv
}


// Configures the RGB blend operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/rgbBlendOperation
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) SetRgbBlendOperation(value BlendOperation) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRgbBlendOperation:"), value)
}


// Configures the source-alpha blend factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/sourceAlphaBlendFactor
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) SourceAlphaBlendFactor() BlendFactor {
	rv := objc.Send[BlendFactor](m_.ID, objc.Sel("sourceAlphaBlendFactor"))
	return rv
}


// Configures the source-alpha blend factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/sourceAlphaBlendFactor
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) SetSourceAlphaBlendFactor(value BlendFactor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceAlphaBlendFactor:"), value)
}


// Configures the source RGB blend factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/sourceRGBBlendFactor
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) SourceRGBBlendFactor() BlendFactor {
	rv := objc.Send[BlendFactor](m_.ID, objc.Sel("sourceRGBBlendFactor"))
	return rv
}


// Configures the source RGB blend factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/sourceRGBBlendFactor
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) SetSourceRGBBlendFactor(value BlendFactor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceRGBBlendFactor:"), value)
}


// Configures the color write mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/writeMask
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) WriteMask() ColorWriteMask {
	rv := objc.Send[ColorWriteMask](m_.ID, objc.Sel("writeMask"))
	return rv
}


// Configures the color write mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/writeMask
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) SetWriteMask(value ColorWriteMask) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWriteMask:"), value)
}








