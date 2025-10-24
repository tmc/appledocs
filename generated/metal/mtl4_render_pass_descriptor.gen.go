// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTL4RenderPassDescriptor */


/* debug [class_header]: Header for MTL4RenderPassDescriptor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTL4RenderPassDescriptor */
// An interface definition for the [MTL4RenderPassDescriptor] class.
type IMTL4RenderPassDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTL4RenderPassDescriptor */
	// properties:
	ColorAttachments() IMTLRenderPassColorAttachmentDescriptorArray
	DefaultRasterSampleCount() uint
	SetDefaultRasterSampleCount(value uint)
	DepthAttachment() IMTLRenderPassDepthAttachmentDescriptor
	SetDepthAttachment(value IMTLRenderPassDepthAttachmentDescriptor)
	ImageblockSampleLength() uint
	SetImageblockSampleLength(value uint)
	RasterizationRateMap() unsafe.Pointer
	SetRasterizationRateMap(value unsafe.Pointer)
	RenderTargetArrayLength() uint
	SetRenderTargetArrayLength(value uint)
	RenderTargetHeight() uint
	SetRenderTargetHeight(value uint)
	RenderTargetWidth() uint
	SetRenderTargetWidth(value uint)
	StencilAttachment() IMTLRenderPassStencilAttachmentDescriptor
	SetStencilAttachment(value IMTLRenderPassStencilAttachmentDescriptor)
	SupportColorAttachmentMapping() bool
	SetSupportColorAttachmentMapping(value bool)
	ThreadgroupMemoryLength() uint
	SetThreadgroupMemoryLength(value uint)
	TileHeight() uint
	SetTileHeight(value uint)
	TileWidth() uint
	SetTileWidth(value uint)
	VisibilityResultBuffer() unsafe.Pointer
	SetVisibilityResultBuffer(value unsafe.Pointer)
	VisibilityResultType() VisibilityResultType
	SetVisibilityResultType(value VisibilityResultType)
	SamplePositions() objc.IObject /* cross-framework: MTLSamplePosition */
	SetSamplePositions(value objc.IObject /* cross-framework: MTLSamplePosition */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTL4RenderPassDescriptor */
	// methods:
	GetSamplePositionsCount(positions objc.IObject /* cross-framework: MTLSamplePosition */, count uint) uint
	SetSamplePositionsCount(positions objc.IObject /* cross-framework: MTLSamplePosition */, count uint)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTL4RenderPassDescriptor */
// Alloc allocates a new instance without initialization.
func (mc _MTL4RenderPassDescriptorClass) Alloc() MTL4RenderPassDescriptor {
	rv := objc.Send[MTL4RenderPassDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTL4RenderPassDescriptor */
// Describes a render pass.
//
// You use render pass descriptors to create instances of and encode draw commands into instances of . To create render command encoders, you typically call . The variant of this method allows you to specify additional options to encode a render pass in parallel from multiple CPU cores by creating and render passes.


// Describes a render pass.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTL4RenderPassDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTL4RenderPassDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTL4RenderPassDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTL4RenderPassDescriptor */

// Retrieves the previously-configured custom sample positions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/getSamplePositions:count:
func (m_ MTL4RenderPassDescriptor) GetSamplePositionsCount(positions objc.IObject /* cross-framework: MTLSamplePosition */, count uint) uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("getSamplePositions:count:"), positions, count)
	return rv
}/* debug [instance_methods/method]: GetSamplePositionsCount */


// Configures the custom sample positions to use in MSAA rendering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/setSamplePositions:count:
func (m_ MTL4RenderPassDescriptor) SetSamplePositionsCount(positions objc.IObject /* cross-framework: MTLSamplePosition */, count uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSamplePositions:count:"), positions, count)
}/* debug [instance_methods/method]: SetSamplePositionsCount */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTL4RenderPassDescriptor */

// Accesses the array of state information for render attachments that store color data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/colorAttachments
func (m_ MTL4RenderPassDescriptor) ColorAttachments() IMTLRenderPassColorAttachmentDescriptorArray {
	rv := objc.Send[RenderPassColorAttachmentDescriptorArray](m_.ID, objc.Sel("colorAttachments"))
	return rv
}/* debug [instance_properties/getter]: colorAttachments */


// Sets the default raster sample count for the render pass when it references no attachments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/defaultRasterSampleCount
func (m_ MTL4RenderPassDescriptor) DefaultRasterSampleCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("defaultRasterSampleCount"))
	return rv
}/* debug [instance_properties/getter]: defaultRasterSampleCount */


// Sets the default raster sample count for the render pass when it references no attachments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/defaultRasterSampleCount
func (m_ MTL4RenderPassDescriptor) SetDefaultRasterSampleCount(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDefaultRasterSampleCount:"), value)
}/* debug [instance_properties/setter]: defaultRasterSampleCount */


// Accesses state information for a render attachment that stores depth data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/depthAttachment
func (m_ MTL4RenderPassDescriptor) DepthAttachment() IMTLRenderPassDepthAttachmentDescriptor {
	rv := objc.Send[RenderPassDepthAttachmentDescriptor](m_.ID, objc.Sel("depthAttachment"))
	return rv
}/* debug [instance_properties/getter]: depthAttachment */


// Accesses state information for a render attachment that stores depth data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/depthAttachment
func (m_ MTL4RenderPassDescriptor) SetDepthAttachment(value IMTLRenderPassDepthAttachmentDescriptor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDepthAttachment:"), value)
}/* debug [instance_properties/setter]: depthAttachment */


// Assigns the per-sample size, in bytes, of the largest explicit imageblock layout in the render pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/imageblockSampleLength
func (m_ MTL4RenderPassDescriptor) ImageblockSampleLength() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("imageblockSampleLength"))
	return rv
}/* debug [instance_properties/getter]: imageblockSampleLength */


// Assigns the per-sample size, in bytes, of the largest explicit imageblock layout in the render pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/imageblockSampleLength
func (m_ MTL4RenderPassDescriptor) SetImageblockSampleLength(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setImageblockSampleLength:"), value)
}/* debug [instance_properties/setter]: imageblockSampleLength */


// Assigns an optional variable rasterization rate map that Metal uses in the render pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/rasterizationRateMap
func (m_ MTL4RenderPassDescriptor) RasterizationRateMap() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("rasterizationRateMap"))
	return rv
}/* debug [instance_properties/getter]: rasterizationRateMap */


// Assigns an optional variable rasterization rate map that Metal uses in the render pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/rasterizationRateMap
func (m_ MTL4RenderPassDescriptor) SetRasterizationRateMap(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRasterizationRateMap:"), value)
}/* debug [instance_properties/setter]: rasterizationRateMap */


// Assigns the number of layers that all attachments this descriptor references have.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/renderTargetArrayLength
func (m_ MTL4RenderPassDescriptor) RenderTargetArrayLength() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("renderTargetArrayLength"))
	return rv
}/* debug [instance_properties/getter]: renderTargetArrayLength */


// Assigns the number of layers that all attachments this descriptor references have.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/renderTargetArrayLength
func (m_ MTL4RenderPassDescriptor) SetRenderTargetArrayLength(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRenderTargetArrayLength:"), value)
}/* debug [instance_properties/setter]: renderTargetArrayLength */


// Sets the height, in pixels, to which Metal constrains the render target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/renderTargetHeight
func (m_ MTL4RenderPassDescriptor) RenderTargetHeight() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("renderTargetHeight"))
	return rv
}/* debug [instance_properties/getter]: renderTargetHeight */


// Sets the height, in pixels, to which Metal constrains the render target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/renderTargetHeight
func (m_ MTL4RenderPassDescriptor) SetRenderTargetHeight(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRenderTargetHeight:"), value)
}/* debug [instance_properties/setter]: renderTargetHeight */


// Sets the width, in pixels, to which Metal constrains the render target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/renderTargetWidth
func (m_ MTL4RenderPassDescriptor) RenderTargetWidth() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("renderTargetWidth"))
	return rv
}/* debug [instance_properties/getter]: renderTargetWidth */


// Sets the width, in pixels, to which Metal constrains the render target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/renderTargetWidth
func (m_ MTL4RenderPassDescriptor) SetRenderTargetWidth(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRenderTargetWidth:"), value)
}/* debug [instance_properties/setter]: renderTargetWidth */


// Accesses state information for a render attachment that stores stencil data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/stencilAttachment
func (m_ MTL4RenderPassDescriptor) StencilAttachment() IMTLRenderPassStencilAttachmentDescriptor {
	rv := objc.Send[RenderPassStencilAttachmentDescriptor](m_.ID, objc.Sel("stencilAttachment"))
	return rv
}/* debug [instance_properties/getter]: stencilAttachment */


// Accesses state information for a render attachment that stores stencil data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/stencilAttachment
func (m_ MTL4RenderPassDescriptor) SetStencilAttachment(value IMTLRenderPassStencilAttachmentDescriptor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStencilAttachment:"), value)
}/* debug [instance_properties/setter]: stencilAttachment */


// Controls if the render pass supports color attachment mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/supportColorAttachmentMapping
func (m_ MTL4RenderPassDescriptor) SupportColorAttachmentMapping() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("supportColorAttachmentMapping"))
	return rv
}/* debug [instance_properties/getter]: supportColorAttachmentMapping */


// Controls if the render pass supports color attachment mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/supportColorAttachmentMapping
func (m_ MTL4RenderPassDescriptor) SetSupportColorAttachmentMapping(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSupportColorAttachmentMapping:"), value)
}/* debug [instance_properties/setter]: supportColorAttachmentMapping */


// Assigns the per-tile size, in bytes, of the persistent threadgroup memory allocation of this render pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/threadgroupMemoryLength
func (m_ MTL4RenderPassDescriptor) ThreadgroupMemoryLength() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("threadgroupMemoryLength"))
	return rv
}/* debug [instance_properties/getter]: threadgroupMemoryLength */


// Assigns the per-tile size, in bytes, of the persistent threadgroup memory allocation of this render pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/threadgroupMemoryLength
func (m_ MTL4RenderPassDescriptor) SetThreadgroupMemoryLength(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setThreadgroupMemoryLength:"), value)
}/* debug [instance_properties/setter]: threadgroupMemoryLength */


// The height of the tiles, in pixels, a render pass you create with this descriptor applies to its attachments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/tileHeight
func (m_ MTL4RenderPassDescriptor) TileHeight() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("tileHeight"))
	return rv
}/* debug [instance_properties/getter]: tileHeight */


// The height of the tiles, in pixels, a render pass you create with this descriptor applies to its attachments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/tileHeight
func (m_ MTL4RenderPassDescriptor) SetTileHeight(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTileHeight:"), value)
}/* debug [instance_properties/setter]: tileHeight */


// The width of the tiles, in pixels, a render pass you create with this descriptor applies to its attachments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/tileWidth
func (m_ MTL4RenderPassDescriptor) TileWidth() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("tileWidth"))
	return rv
}/* debug [instance_properties/getter]: tileWidth */


// The width of the tiles, in pixels, a render pass you create with this descriptor applies to its attachments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/tileWidth
func (m_ MTL4RenderPassDescriptor) SetTileWidth(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTileWidth:"), value)
}/* debug [instance_properties/setter]: tileWidth */


// Configures a buffer into which Metal writes counts of fragments (pixels) passing the depth and stencil tests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/visibilityResultBuffer
func (m_ MTL4RenderPassDescriptor) VisibilityResultBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("visibilityResultBuffer"))
	return rv
}/* debug [instance_properties/getter]: visibilityResultBuffer */


// Configures a buffer into which Metal writes counts of fragments (pixels) passing the depth and stencil tests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/visibilityResultBuffer
func (m_ MTL4RenderPassDescriptor) SetVisibilityResultBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVisibilityResultBuffer:"), value)
}/* debug [instance_properties/setter]: visibilityResultBuffer */


// Determines if Metal accumulates visibility results between render encoders or resets them.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/visibilityResultType
func (m_ MTL4RenderPassDescriptor) VisibilityResultType() VisibilityResultType {
	rv := objc.Send[VisibilityResultType](m_.ID, objc.Sel("visibilityResultType"))
	return rv
}/* debug [instance_properties/getter]: visibilityResultType */


// Determines if Metal accumulates visibility results between render encoders or resets them.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPassDescriptor/visibilityResultType
func (m_ MTL4RenderPassDescriptor) SetVisibilityResultType(value VisibilityResultType) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVisibilityResultType:"), value)
}/* debug [instance_properties/setter]: visibilityResultType */


// Configures the custom sample positions to use in MSAA rendering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4renderpassdescriptor/samplepositions
func (m_ MTL4RenderPassDescriptor) SamplePositions() objc.IObject /* cross-framework: MTLSamplePosition */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("samplePositions"))
	return rv
}/* debug [instance_properties/getter]: samplePositions */


// Configures the custom sample positions to use in MSAA rendering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4renderpassdescriptor/samplepositions
func (m_ MTL4RenderPassDescriptor) SetSamplePositions(value objc.IObject /* cross-framework: MTLSamplePosition */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSamplePositions:"), value)
}/* debug [instance_properties/setter]: samplePositions */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTL4RenderPassDescriptor */



