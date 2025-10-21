// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
	Reset()
}

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor
type MTL4RenderPipelineColorAttachmentDescriptor struct {
	objectivec.Object
}

// MTL4RenderPipelineColorAttachmentDescriptorFrom constructs a [MTL4RenderPipelineColorAttachmentDescriptor] from an unsafe.Pointer.
func MTL4RenderPipelineColorAttachmentDescriptorFrom(ptr unsafe.Pointer) MTL4RenderPipelineColorAttachmentDescriptor {
	return MTL4RenderPipelineColorAttachmentDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTL4RenderPipelineColorAttachmentDescriptorClass) Alloc() MTL4RenderPipelineColorAttachmentDescriptor {
	rv := objc.Send[MTL4RenderPipelineColorAttachmentDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Resets this descriptor to its default state.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/reset()
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) Reset() {
	objc.Send[objc.ID](m_.ID, objc.Sel("reset"))
}

// Configures the alpha blending operation.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/alphaBlendOperation
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) AlphaBlendOperation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("alphaBlendOperation"))
	return rv
}


// SetAlphaBlendOperation sets the value of the alphaBlendOperation property.
// Configures the alpha blending operation.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/alphaBlendOperation
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) SetAlphaBlendOperation(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlphaBlendOperation:"), value)
}
// Configure the blend state for color attachments the pipeline state uses.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/blendingState
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) BlendingState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("blendingState"))
	return rv
}


// SetBlendingState sets the value of the blendingState property.
// Configure the blend state for color attachments the pipeline state uses.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/blendingState
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) SetBlendingState(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBlendingState:"), value)
}
// Configures the destination-alpha blend factor.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/destinationAlphaBlendFactor
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) DestinationAlphaBlendFactor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("destinationAlphaBlendFactor"))
	return rv
}


// SetDestinationAlphaBlendFactor sets the value of the destinationAlphaBlendFactor property.
// Configures the destination-alpha blend factor.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/destinationAlphaBlendFactor
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) SetDestinationAlphaBlendFactor(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDestinationAlphaBlendFactor:"), value)
}
// Configures the destination RGB blend factor.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/destinationRGBBlendFactor
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) DestinationRGBBlendFactor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("destinationRGBBlendFactor"))
	return rv
}


// SetDestinationRGBBlendFactor sets the value of the destinationRGBBlendFactor property.
// Configures the destination RGB blend factor.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/destinationRGBBlendFactor
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) SetDestinationRGBBlendFactor(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDestinationRGBBlendFactor:"), value)
}
// Configures the pixel format.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/pixelFormat
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) PixelFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("pixelFormat"))
	return rv
}


// SetPixelFormat sets the value of the pixelFormat property.
// Configures the pixel format.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/pixelFormat
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) SetPixelFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPixelFormat:"), value)
}
// Configures the RGB blend operation.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/rgbBlendOperation
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) RgbBlendOperation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("rgbBlendOperation"))
	return rv
}


// SetRgbBlendOperation sets the value of the rgbBlendOperation property.
// Configures the RGB blend operation.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/rgbBlendOperation
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) SetRgbBlendOperation(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRgbBlendOperation:"), value)
}
// Configures the source-alpha blend factor.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/sourceAlphaBlendFactor
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) SourceAlphaBlendFactor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("sourceAlphaBlendFactor"))
	return rv
}


// SetSourceAlphaBlendFactor sets the value of the sourceAlphaBlendFactor property.
// Configures the source-alpha blend factor.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/sourceAlphaBlendFactor
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) SetSourceAlphaBlendFactor(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceAlphaBlendFactor:"), value)
}
// Configures the source RGB blend factor.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/sourceRGBBlendFactor
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) SourceRGBBlendFactor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("sourceRGBBlendFactor"))
	return rv
}


// SetSourceRGBBlendFactor sets the value of the sourceRGBBlendFactor property.
// Configures the source RGB blend factor.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/sourceRGBBlendFactor
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) SetSourceRGBBlendFactor(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceRGBBlendFactor:"), value)
}
// Configures the color write mask.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/writeMask
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) WriteMask() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("writeMask"))
	return rv
}


// SetWriteMask sets the value of the writeMask property.
// Configures the color write mask.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineColorAttachmentDescriptor/writeMask
func (m_ MTL4RenderPipelineColorAttachmentDescriptor) SetWriteMask(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWriteMask:"), value)
}


