// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [MTL4TileRenderPipelineDescriptor] class.
var (
	MTL4TileRenderPipelineDescriptorClass     _MTL4TileRenderPipelineDescriptorClass
	MTL4TileRenderPipelineDescriptorClassOnce sync.Once
)

func getMTL4TileRenderPipelineDescriptorClass() _MTL4TileRenderPipelineDescriptorClass {
	MTL4TileRenderPipelineDescriptorClassOnce.Do(func() {
		MTL4TileRenderPipelineDescriptorClass = _MTL4TileRenderPipelineDescriptorClass{objc.GetClass("MTL4TileRenderPipelineDescriptor")}
	})
	return MTL4TileRenderPipelineDescriptorClass
}

type _MTL4TileRenderPipelineDescriptorClass struct {
	class objc.Class
}





// An interface definition for the [MTL4TileRenderPipelineDescriptor] class.
type IMTL4TileRenderPipelineDescriptor interface {
	IMTL4PipelineDescriptor
	

	// properties:
	ColorAttachments() IMTLTileRenderPipelineColorAttachmentDescriptorArray
	MaxTotalThreadsPerThreadgroup() uint
	SetMaxTotalThreadsPerThreadgroup(value uint)
	RasterSampleCount() uint
	SetRasterSampleCount(value uint)
	RequiredThreadsPerThreadgroup() MTLSize
	SetRequiredThreadsPerThreadgroup(value MTLSize)
	StaticLinkingDescriptor() IMTL4StaticLinkingDescriptor
	SetStaticLinkingDescriptor(value IMTL4StaticLinkingDescriptor)
	SupportBinaryLinking() bool
	SetSupportBinaryLinking(value bool)
	ThreadgroupSizeMatchesTileSize() bool
	SetThreadgroupSizeMatchesTileSize(value bool)
	TileFunctionDescriptor() IMTL4FunctionDescriptor
	SetTileFunctionDescriptor(value IMTL4FunctionDescriptor)


	

	// methods:
	Reset()


}





// Alloc allocates a new instance without initialization.
func (mc _MTL4TileRenderPipelineDescriptorClass) Alloc() MTL4TileRenderPipelineDescriptor {
	rv := objc.Send[MTL4TileRenderPipelineDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTL4TileRenderPipelineDescriptorClass) New() MTL4TileRenderPipelineDescriptor {
	rv := objc.Send[MTL4TileRenderPipelineDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4TileRenderPipelineDescriptor) Init() MTL4TileRenderPipelineDescriptor {
	rv := objc.Send[MTL4TileRenderPipelineDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4TileRenderPipelineDescriptor) Autorelease() MTL4TileRenderPipelineDescriptor {
	rv := objc.Send[MTL4TileRenderPipelineDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4TileRenderPipelineDescriptor creates a new MTL4TileRenderPipelineDescriptor instance.
func NewMTL4TileRenderPipelineDescriptor() MTL4TileRenderPipelineDescriptor {
	return getMTL4TileRenderPipelineDescriptorClass().New()
}





// Groups together properties you use to create a tile render pipeline state object.


// Groups together properties you use to create a tile render pipeline state object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4TileRenderPipelineDescriptor
type MTL4TileRenderPipelineDescriptor struct {
	MTL4PipelineDescriptor
}

// MTL4TileRenderPipelineDescriptorFrom constructs a [MTL4TileRenderPipelineDescriptor] from an unsafe.Pointer.
//
// Groups together properties you use to create a tile render pipeline state object.
func MTL4TileRenderPipelineDescriptorFrom(ptr unsafe.Pointer) MTL4TileRenderPipelineDescriptor {
	return MTL4TileRenderPipelineDescriptor{
		MTL4PipelineDescriptor: MTL4PipelineDescriptorFrom(ptr),
	}
}




















// Resets the descriptor to the default state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4TileRenderPipelineDescriptor/reset()
func (m_ MTL4TileRenderPipelineDescriptor) Reset() {
	objc.Send[objc.ID](m_.ID, objc.Sel("reset"))
}







// Access an array of descriptors that configure the properties of each color attachment in the tile render pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4TileRenderPipelineDescriptor/colorAttachments
func (m_ MTL4TileRenderPipelineDescriptor) ColorAttachments() IMTLTileRenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Send[TileRenderPipelineColorAttachmentDescriptorArray](m_.ID, objc.Sel("colorAttachments"))
	return rv
}


// Sets the maximum number of threads that the GPU can execute simultaneously within a single threadgroup in the tile render pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4TileRenderPipelineDescriptor/maxTotalThreadsPerThreadgroup
func (m_ MTL4TileRenderPipelineDescriptor) MaxTotalThreadsPerThreadgroup() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("maxTotalThreadsPerThreadgroup"))
	return rv
}


// Sets the maximum number of threads that the GPU can execute simultaneously within a single threadgroup in the tile render pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4TileRenderPipelineDescriptor/maxTotalThreadsPerThreadgroup
func (m_ MTL4TileRenderPipelineDescriptor) SetMaxTotalThreadsPerThreadgroup(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxTotalThreadsPerThreadgroup:"), value)
}


// Configures the number of samples per pixel used for multisampling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4TileRenderPipelineDescriptor/rasterSampleCount
func (m_ MTL4TileRenderPipelineDescriptor) RasterSampleCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("rasterSampleCount"))
	return rv
}


// Configures the number of samples per pixel used for multisampling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4TileRenderPipelineDescriptor/rasterSampleCount
func (m_ MTL4TileRenderPipelineDescriptor) SetRasterSampleCount(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRasterSampleCount:"), value)
}


// Sets the required number of threads per threadgroup for tile dispatches.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4TileRenderPipelineDescriptor/requiredThreadsPerThreadgroup
func (m_ MTL4TileRenderPipelineDescriptor) RequiredThreadsPerThreadgroup() MTLSize {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("requiredThreadsPerThreadgroup"))
	return rv
}


// Sets the required number of threads per threadgroup for tile dispatches.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4TileRenderPipelineDescriptor/requiredThreadsPerThreadgroup
func (m_ MTL4TileRenderPipelineDescriptor) SetRequiredThreadsPerThreadgroup(value MTLSize) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRequiredThreadsPerThreadgroup:"), value)
}


// Configures an object that contains information about functions to link to the tile render pipeline when Metal builds it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4TileRenderPipelineDescriptor/staticLinkingDescriptor
func (m_ MTL4TileRenderPipelineDescriptor) StaticLinkingDescriptor() IMTL4StaticLinkingDescriptor {
	rv := objc.Send[MTL4StaticLinkingDescriptor](m_.ID, objc.Sel("staticLinkingDescriptor"))
	return rv
}


// Configures an object that contains information about functions to link to the tile render pipeline when Metal builds it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4TileRenderPipelineDescriptor/staticLinkingDescriptor
func (m_ MTL4TileRenderPipelineDescriptor) SetStaticLinkingDescriptor(value IMTL4StaticLinkingDescriptor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStaticLinkingDescriptor:"), value)
}


// Indicates whether the pipeline supports linking binary functions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4TileRenderPipelineDescriptor/supportBinaryLinking
func (m_ MTL4TileRenderPipelineDescriptor) SupportBinaryLinking() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("supportBinaryLinking"))
	return rv
}


// Indicates whether the pipeline supports linking binary functions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4TileRenderPipelineDescriptor/supportBinaryLinking
func (m_ MTL4TileRenderPipelineDescriptor) SetSupportBinaryLinking(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSupportBinaryLinking:"), value)
}


// Indicating whether the size of the threadgroup matches the size of a tile in the render pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4TileRenderPipelineDescriptor/threadgroupSizeMatchesTileSize
func (m_ MTL4TileRenderPipelineDescriptor) ThreadgroupSizeMatchesTileSize() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("threadgroupSizeMatchesTileSize"))
	return rv
}


// Indicating whether the size of the threadgroup matches the size of a tile in the render pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4TileRenderPipelineDescriptor/threadgroupSizeMatchesTileSize
func (m_ MTL4TileRenderPipelineDescriptor) SetThreadgroupSizeMatchesTileSize(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setThreadgroupSizeMatchesTileSize:"), value)
}


// Configures the tile function that the render pipeline executes for each tile in the tile shader stage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4TileRenderPipelineDescriptor/tileFunctionDescriptor
func (m_ MTL4TileRenderPipelineDescriptor) TileFunctionDescriptor() IMTL4FunctionDescriptor {
	rv := objc.Send[MTL4FunctionDescriptor](m_.ID, objc.Sel("tileFunctionDescriptor"))
	return rv
}


// Configures the tile function that the render pipeline executes for each tile in the tile shader stage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4TileRenderPipelineDescriptor/tileFunctionDescriptor
func (m_ MTL4TileRenderPipelineDescriptor) SetTileFunctionDescriptor(value IMTL4FunctionDescriptor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTileFunctionDescriptor:"), value)
}








