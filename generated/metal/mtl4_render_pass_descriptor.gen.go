// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTL4RenderPassDescriptor] class.
var (
	MTL4RenderPassDescriptorClass     _MTL4RenderPassDescriptorClass
	MTL4RenderPassDescriptorClassOnce sync.Once
)

func getMTL4RenderPassDescriptorClass() _MTL4RenderPassDescriptorClass {
	MTL4RenderPassDescriptorClassOnce.Do(func() {
		MTL4RenderPassDescriptorClass = _MTL4RenderPassDescriptorClass{objc.GetClass("MTL4RenderPassDescriptor")}
	})
	return MTL4RenderPassDescriptorClass
}

type _MTL4RenderPassDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [MTL4RenderPassDescriptor] class.
type IMTL4RenderPassDescriptor interface {
	objectivec.IObject
	GetSamplePositionsCount(positions unsafe.Pointer, count uint) uint
	SetSamplePositionsCount(positions unsafe.Pointer, count uint)
}

// Describes a render pass.
//
// You use render pass descriptors to create instances of and encode draw commands into instances of . To create render command encoders, you typically call . The variant of this method allows you to specify additional options to encode a render pass in parallel from multiple CPU cores by creating and render passes.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor
type MTL4RenderPassDescriptor struct {
	objectivec.Object
}

// MTL4RenderPassDescriptorFrom constructs a [MTL4RenderPassDescriptor] from an unsafe.Pointer.
//
// Describes a render pass.
func MTL4RenderPassDescriptorFrom(ptr unsafe.Pointer) MTL4RenderPassDescriptor {
	return MTL4RenderPassDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTL4RenderPassDescriptorClass) Alloc() MTL4RenderPassDescriptor {
	rv := objc.Send[MTL4RenderPassDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTL4RenderPassDescriptorClass) New() MTL4RenderPassDescriptor {
	rv := objc.Send[MTL4RenderPassDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4RenderPassDescriptor) Init() MTL4RenderPassDescriptor {
	rv := objc.Send[MTL4RenderPassDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4RenderPassDescriptor) Autorelease() MTL4RenderPassDescriptor {
	rv := objc.Send[MTL4RenderPassDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4RenderPassDescriptor creates a new MTL4RenderPassDescriptor instance.
func NewMTL4RenderPassDescriptor() MTL4RenderPassDescriptor {
	return getMTL4RenderPassDescriptorClass().New()
}


// Retrieves the previously-configured custom sample positions.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/getSamplePositions:count:
func (m_ MTL4RenderPassDescriptor) GetSamplePositionsCount(positions unsafe.Pointer, count uint) uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("getSamplePositions:count:"), positions, count)
	return rv
}

// Configures the custom sample positions to use in MSAA rendering.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/setSamplePositions:count:
func (m_ MTL4RenderPassDescriptor) SetSamplePositionsCount(positions unsafe.Pointer, count uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSamplePositions:count:"), positions, count)
}

// Sets the default raster sample count for the render pass when it references no attachments.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/defaultRasterSampleCount
func (m_ MTL4RenderPassDescriptor) DefaultRasterSampleCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("defaultRasterSampleCount"))
	return rv
}


// SetDefaultRasterSampleCount sets the value of the defaultRasterSampleCount property.
// Sets the default raster sample count for the render pass when it references no attachments.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/defaultRasterSampleCount
func (m_ MTL4RenderPassDescriptor) SetDefaultRasterSampleCount(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDefaultRasterSampleCount:"), value)
}

// Accesses state information for a render attachment that stores depth data.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/depthAttachment
func (m_ MTL4RenderPassDescriptor) DepthAttachment() MTLRenderPassDepthAttachmentDescriptor {
	rv := objc.Send[MTLRenderPassDepthAttachmentDescriptor](m_.ID, objc.Sel("depthAttachment"))
	return rv
}


// SetDepthAttachment sets the value of the depthAttachment property.
// Accesses state information for a render attachment that stores depth data.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/depthAttachment
func (m_ MTL4RenderPassDescriptor) SetDepthAttachment(value IMTLRenderPassDepthAttachmentDescriptor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDepthAttachment:"), value)
}

// Assigns the per-sample size, in bytes, of the largest explicit imageblock layout in the render pass.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/imageblockSampleLength
func (m_ MTL4RenderPassDescriptor) ImageblockSampleLength() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("imageblockSampleLength"))
	return rv
}


// SetImageblockSampleLength sets the value of the imageblockSampleLength property.
// Assigns the per-sample size, in bytes, of the largest explicit imageblock layout in the render pass.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/imageblockSampleLength
func (m_ MTL4RenderPassDescriptor) SetImageblockSampleLength(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setImageblockSampleLength:"), value)
}

// Assigns an optional variable rasterization rate map that Metal uses in the render pass.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/rasterizationRateMap
func (m_ MTL4RenderPassDescriptor) RasterizationRateMap() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("rasterizationRateMap"))
	return rv
}


// SetRasterizationRateMap sets the value of the rasterizationRateMap property.
// Assigns an optional variable rasterization rate map that Metal uses in the render pass.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/rasterizationRateMap
func (m_ MTL4RenderPassDescriptor) SetRasterizationRateMap(value objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRasterizationRateMap:"), value)
}

// Assigns the number of layers that all attachments this descriptor references have.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/renderTargetArrayLength
func (m_ MTL4RenderPassDescriptor) RenderTargetArrayLength() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("renderTargetArrayLength"))
	return rv
}


// SetRenderTargetArrayLength sets the value of the renderTargetArrayLength property.
// Assigns the number of layers that all attachments this descriptor references have.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/renderTargetArrayLength
func (m_ MTL4RenderPassDescriptor) SetRenderTargetArrayLength(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRenderTargetArrayLength:"), value)
}

// Sets the height, in pixels, to which Metal constrains the render target.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/renderTargetHeight
func (m_ MTL4RenderPassDescriptor) RenderTargetHeight() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("renderTargetHeight"))
	return rv
}


// SetRenderTargetHeight sets the value of the renderTargetHeight property.
// Sets the height, in pixels, to which Metal constrains the render target.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/renderTargetHeight
func (m_ MTL4RenderPassDescriptor) SetRenderTargetHeight(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRenderTargetHeight:"), value)
}

// Sets the width, in pixels, to which Metal constrains the render target.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/renderTargetWidth
func (m_ MTL4RenderPassDescriptor) RenderTargetWidth() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("renderTargetWidth"))
	return rv
}


// SetRenderTargetWidth sets the value of the renderTargetWidth property.
// Sets the width, in pixels, to which Metal constrains the render target.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/renderTargetWidth
func (m_ MTL4RenderPassDescriptor) SetRenderTargetWidth(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRenderTargetWidth:"), value)
}

// Accesses state information for a render attachment that stores stencil data.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/stencilAttachment
func (m_ MTL4RenderPassDescriptor) StencilAttachment() MTLRenderPassStencilAttachmentDescriptor {
	rv := objc.Send[MTLRenderPassStencilAttachmentDescriptor](m_.ID, objc.Sel("stencilAttachment"))
	return rv
}


// SetStencilAttachment sets the value of the stencilAttachment property.
// Accesses state information for a render attachment that stores stencil data.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/stencilAttachment
func (m_ MTL4RenderPassDescriptor) SetStencilAttachment(value IMTLRenderPassStencilAttachmentDescriptor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStencilAttachment:"), value)
}

// Controls if the render pass supports color attachment mapping.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/supportColorAttachmentMapping
func (m_ MTL4RenderPassDescriptor) SupportColorAttachmentMapping() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("supportColorAttachmentMapping"))
	return rv
}


// SetSupportColorAttachmentMapping sets the value of the supportColorAttachmentMapping property.
// Controls if the render pass supports color attachment mapping.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/supportColorAttachmentMapping
func (m_ MTL4RenderPassDescriptor) SetSupportColorAttachmentMapping(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSupportColorAttachmentMapping:"), value)
}

// Assigns the per-tile size, in bytes, of the persistent threadgroup memory allocation of this render pass.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/threadgroupMemoryLength
func (m_ MTL4RenderPassDescriptor) ThreadgroupMemoryLength() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("threadgroupMemoryLength"))
	return rv
}


// SetThreadgroupMemoryLength sets the value of the threadgroupMemoryLength property.
// Assigns the per-tile size, in bytes, of the persistent threadgroup memory allocation of this render pass.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/threadgroupMemoryLength
func (m_ MTL4RenderPassDescriptor) SetThreadgroupMemoryLength(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setThreadgroupMemoryLength:"), value)
}

// The height of the tiles, in pixels, a render pass you create with this descriptor applies to its attachments.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/tileHeight
func (m_ MTL4RenderPassDescriptor) TileHeight() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("tileHeight"))
	return rv
}


// SetTileHeight sets the value of the tileHeight property.
// The height of the tiles, in pixels, a render pass you create with this descriptor applies to its attachments.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/tileHeight
func (m_ MTL4RenderPassDescriptor) SetTileHeight(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTileHeight:"), value)
}

// The width of the tiles, in pixels, a render pass you create with this descriptor applies to its attachments.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/tileWidth
func (m_ MTL4RenderPassDescriptor) TileWidth() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("tileWidth"))
	return rv
}


// SetTileWidth sets the value of the tileWidth property.
// The width of the tiles, in pixels, a render pass you create with this descriptor applies to its attachments.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/tileWidth
func (m_ MTL4RenderPassDescriptor) SetTileWidth(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTileWidth:"), value)
}

// Configures a buffer into which Metal writes counts of fragments (pixels) passing the depth and stencil tests.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/visibilityResultBuffer
func (m_ MTL4RenderPassDescriptor) VisibilityResultBuffer() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("visibilityResultBuffer"))
	return rv
}


// SetVisibilityResultBuffer sets the value of the visibilityResultBuffer property.
// Configures a buffer into which Metal writes counts of fragments (pixels) passing the depth and stencil tests.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/visibilityResultBuffer
func (m_ MTL4RenderPassDescriptor) SetVisibilityResultBuffer(value objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVisibilityResultBuffer:"), value)
}

// Determines if Metal accumulates visibility results between render encoders or resets them.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/visibilityResultType
func (m_ MTL4RenderPassDescriptor) VisibilityResultType() VisibilityResultType {
	rv := objc.Send[VisibilityResultType](m_.ID, objc.Sel("visibilityResultType"))
	return rv
}


// SetVisibilityResultType sets the value of the visibilityResultType property.
// Determines if Metal accumulates visibility results between render encoders or resets them.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/visibilityResultType
func (m_ MTL4RenderPassDescriptor) SetVisibilityResultType(value VisibilityResultType) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVisibilityResultType:"), value)
}

// Accesses the array of state information for render attachments that store color data.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4renderpassdescriptor/colorattachments
func (m_ MTL4RenderPassDescriptor) ColorAttachments() MTLRenderPassColorAttachmentDescriptorArray {
	rv := objc.Send[MTLRenderPassColorAttachmentDescriptorArray](m_.ID, objc.Sel("colorAttachments"))
	return rv
}


// SetColorAttachments sets the value of the colorAttachments property.
// Accesses the array of state information for render attachments that store color data.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4renderpassdescriptor/colorattachments
func (m_ MTL4RenderPassDescriptor) SetColorAttachments(value IMTLRenderPassColorAttachmentDescriptorArray) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setColorAttachments:"), value)
}

// Configures the custom sample positions to use in MSAA rendering.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4renderpassdescriptor/samplepositions
func (m_ MTL4RenderPassDescriptor) SamplePositions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("samplePositions"))
	return rv
}


// SetSamplePositions sets the value of the samplePositions property.
// Configures the custom sample positions to use in MSAA rendering.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4renderpassdescriptor/samplepositions
func (m_ MTL4RenderPassDescriptor) SetSamplePositions(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSamplePositions:"), value)
}



