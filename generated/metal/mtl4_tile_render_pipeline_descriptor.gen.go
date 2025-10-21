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
	Reset()
}

// Groups together properties you use to create a tile render pipeline state object.
//
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

// Alloc allocates a new instance without initialization.
func (mc _MTL4TileRenderPipelineDescriptorClass) Alloc() MTL4TileRenderPipelineDescriptor {
	rv := objc.Send[MTL4TileRenderPipelineDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Resets the descriptor to the default state.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4TileRenderPipelineDescriptor/reset()
func (m_ MTL4TileRenderPipelineDescriptor) Reset() {
	objc.Send[objc.ID](m_.ID, objc.Sel("reset"))
}

// Access an array of descriptors that configure the properties of each color attachment in the tile render pipeline.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4TileRenderPipelineDescriptor/colorAttachments
func (m_ MTL4TileRenderPipelineDescriptor) ColorAttachments() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("colorAttachments"))
	return rv
}

// Sets the maximum number of threads that the GPU can execute simultaneously within a single threadgroup in the tile render pipeline.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4TileRenderPipelineDescriptor/maxTotalThreadsPerThreadgroup
func (m_ MTL4TileRenderPipelineDescriptor) MaxTotalThreadsPerThreadgroup() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("maxTotalThreadsPerThreadgroup"))
	return rv
}


// SetMaxTotalThreadsPerThreadgroup sets the value of the maxTotalThreadsPerThreadgroup property.
// Sets the maximum number of threads that the GPU can execute simultaneously within a single threadgroup in the tile render pipeline.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4TileRenderPipelineDescriptor/maxTotalThreadsPerThreadgroup
func (m_ MTL4TileRenderPipelineDescriptor) SetMaxTotalThreadsPerThreadgroup(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxTotalThreadsPerThreadgroup:"), value)
}
// Configures the number of samples per pixel used for multisampling.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4TileRenderPipelineDescriptor/rasterSampleCount
func (m_ MTL4TileRenderPipelineDescriptor) RasterSampleCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("rasterSampleCount"))
	return rv
}


// SetRasterSampleCount sets the value of the rasterSampleCount property.
// Configures the number of samples per pixel used for multisampling.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4TileRenderPipelineDescriptor/rasterSampleCount
func (m_ MTL4TileRenderPipelineDescriptor) SetRasterSampleCount(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRasterSampleCount:"), value)
}
// Sets the required number of threads per threadgroup for tile dispatches.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4TileRenderPipelineDescriptor/requiredThreadsPerThreadgroup
func (m_ MTL4TileRenderPipelineDescriptor) RequiredThreadsPerThreadgroup() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("requiredThreadsPerThreadgroup"))
	return rv
}


// SetRequiredThreadsPerThreadgroup sets the value of the requiredThreadsPerThreadgroup property.
// Sets the required number of threads per threadgroup for tile dispatches.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4TileRenderPipelineDescriptor/requiredThreadsPerThreadgroup
func (m_ MTL4TileRenderPipelineDescriptor) SetRequiredThreadsPerThreadgroup(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRequiredThreadsPerThreadgroup:"), value)
}
// Configures an object that contains information about functions to link to the tile render pipeline when Metal builds it.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4TileRenderPipelineDescriptor/staticLinkingDescriptor
func (m_ MTL4TileRenderPipelineDescriptor) StaticLinkingDescriptor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("staticLinkingDescriptor"))
	return rv
}


// SetStaticLinkingDescriptor sets the value of the staticLinkingDescriptor property.
// Configures an object that contains information about functions to link to the tile render pipeline when Metal builds it.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4TileRenderPipelineDescriptor/staticLinkingDescriptor
func (m_ MTL4TileRenderPipelineDescriptor) SetStaticLinkingDescriptor(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStaticLinkingDescriptor:"), value)
}
// Indicates whether the pipeline supports linking binary functions.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4TileRenderPipelineDescriptor/supportBinaryLinking
func (m_ MTL4TileRenderPipelineDescriptor) SupportBinaryLinking() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("supportBinaryLinking"))
	return rv
}


// SetSupportBinaryLinking sets the value of the supportBinaryLinking property.
// Indicates whether the pipeline supports linking binary functions.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4TileRenderPipelineDescriptor/supportBinaryLinking
func (m_ MTL4TileRenderPipelineDescriptor) SetSupportBinaryLinking(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSupportBinaryLinking:"), value)
}
// Indicating whether the size of the threadgroup matches the size of a tile in the render pipeline.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4TileRenderPipelineDescriptor/threadgroupSizeMatchesTileSize
func (m_ MTL4TileRenderPipelineDescriptor) ThreadgroupSizeMatchesTileSize() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("threadgroupSizeMatchesTileSize"))
	return rv
}


// SetThreadgroupSizeMatchesTileSize sets the value of the threadgroupSizeMatchesTileSize property.
// Indicating whether the size of the threadgroup matches the size of a tile in the render pipeline.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4TileRenderPipelineDescriptor/threadgroupSizeMatchesTileSize
func (m_ MTL4TileRenderPipelineDescriptor) SetThreadgroupSizeMatchesTileSize(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setThreadgroupSizeMatchesTileSize:"), value)
}
// Configures the tile function that the render pipeline executes for each tile in the tile shader stage.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4TileRenderPipelineDescriptor/tileFunctionDescriptor
func (m_ MTL4TileRenderPipelineDescriptor) TileFunctionDescriptor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("tileFunctionDescriptor"))
	return rv
}


// SetTileFunctionDescriptor sets the value of the tileFunctionDescriptor property.
// Configures the tile function that the render pipeline executes for each tile in the tile shader stage.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4TileRenderPipelineDescriptor/tileFunctionDescriptor
func (m_ MTL4TileRenderPipelineDescriptor) SetTileFunctionDescriptor(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTileFunctionDescriptor:"), value)
}


