// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TileRenderPipelineDescriptor] class.
var (
	TileRenderPipelineDescriptorClass     _TileRenderPipelineDescriptorClass
	TileRenderPipelineDescriptorClassOnce sync.Once
)

func getTileRenderPipelineDescriptorClass() _TileRenderPipelineDescriptorClass {
	TileRenderPipelineDescriptorClassOnce.Do(func() {
		TileRenderPipelineDescriptorClass = _TileRenderPipelineDescriptorClass{objc.GetClass("MTLTileRenderPipelineDescriptor")}
	})
	return TileRenderPipelineDescriptorClass
}

type _TileRenderPipelineDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [TileRenderPipelineDescriptor] class.
type ITileRenderPipelineDescriptor interface {
	objectivec.IObject
	Reset()
}

// An object that configures new render pipeline state objects for tile shading.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor
type TileRenderPipelineDescriptor struct {
	objectivec.Object
}

// TileRenderPipelineDescriptorFrom constructs a [TileRenderPipelineDescriptor] from an unsafe.Pointer.
//
// An object that configures new render pipeline state objects for tile shading.
func TileRenderPipelineDescriptorFrom(ptr unsafe.Pointer) TileRenderPipelineDescriptor {
	return TileRenderPipelineDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TileRenderPipelineDescriptorClass) Alloc() TileRenderPipelineDescriptor {
	rv := objc.Send[TileRenderPipelineDescriptor](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TileRenderPipelineDescriptorClass) New() TileRenderPipelineDescriptor {
	rv := objc.Send[TileRenderPipelineDescriptor](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TileRenderPipelineDescriptor) Init() TileRenderPipelineDescriptor {
	rv := objc.Send[TileRenderPipelineDescriptor](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TileRenderPipelineDescriptor) Autorelease() TileRenderPipelineDescriptor {
	rv := objc.Send[TileRenderPipelineDescriptor](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTileRenderPipelineDescriptor creates a new TileRenderPipelineDescriptor instance.
func NewTileRenderPipelineDescriptor() TileRenderPipelineDescriptor {
	return getTileRenderPipelineDescriptorClass().New()
}


// Specifies the default rendering pipeline state values for the descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/reset()
func (t_ TileRenderPipelineDescriptor) Reset() {
	objc.Send[objc.ID](t_.ID, objc.Sel("reset"))
}

// An array of binary archives to search for precompiled versions of the shader.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/binaryArchives
func (t_ TileRenderPipelineDescriptor) BinaryArchives() []objc.ID {
	rv := objc.Send[[]objc.ID](t_.ID, objc.Sel("binaryArchives"))
	return rv
}


// SetBinaryArchives sets the value of the binaryArchives property.
// An array of binary archives to search for precompiled versions of the shader.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/binaryArchives
func (t_ TileRenderPipelineDescriptor) SetBinaryArchives(value []objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBinaryArchives:"), value)
}
// An array of attachments that store color data.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/colorAttachments
func (t_ TileRenderPipelineDescriptor) ColorAttachments() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("colorAttachments"))
	return rv
}

// A string that identifies the tile pipeline descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/label
func (t_ TileRenderPipelineDescriptor) Label() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("label"))
	return rv
}


// SetLabel sets the value of the label property.
// A string that identifies the tile pipeline descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/label
func (t_ TileRenderPipelineDescriptor) SetLabel(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLabel:"), value)
}
// Functions that you can specify as function arguments for the tile shader when encoding commands that use the pipeline.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/linkedFunctions
func (t_ TileRenderPipelineDescriptor) LinkedFunctions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("linkedFunctions"))
	return rv
}


// SetLinkedFunctions sets the value of the linkedFunctions property.
// Functions that you can specify as function arguments for the tile shader when encoding commands that use the pipeline.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/linkedFunctions
func (t_ TileRenderPipelineDescriptor) SetLinkedFunctions(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLinkedFunctions:"), value)
}
// The maximum function call depth from the top-most shader function.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/maxCallStackDepth
func (t_ TileRenderPipelineDescriptor) MaxCallStackDepth() uint {
	rv := objc.Send[uint](t_.ID, objc.Sel("maxCallStackDepth"))
	return rv
}


// SetMaxCallStackDepth sets the value of the maxCallStackDepth property.
// The maximum function call depth from the top-most shader function.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/maxCallStackDepth
func (t_ TileRenderPipelineDescriptor) SetMaxCallStackDepth(value uint) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMaxCallStackDepth:"), value)
}
// The maximum number of threads in a threadgroup when dispatching a command using the pipeline.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/maxTotalThreadsPerThreadgroup
func (t_ TileRenderPipelineDescriptor) MaxTotalThreadsPerThreadgroup() uint {
	rv := objc.Send[uint](t_.ID, objc.Sel("maxTotalThreadsPerThreadgroup"))
	return rv
}


// SetMaxTotalThreadsPerThreadgroup sets the value of the maxTotalThreadsPerThreadgroup property.
// The maximum number of threads in a threadgroup when dispatching a command using the pipeline.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/maxTotalThreadsPerThreadgroup
func (t_ TileRenderPipelineDescriptor) SetMaxTotalThreadsPerThreadgroup(value uint) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMaxTotalThreadsPerThreadgroup:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/preloadedLibraries
func (t_ TileRenderPipelineDescriptor) PreloadedLibraries() []objc.ID {
	rv := objc.Send[[]objc.ID](t_.ID, objc.Sel("preloadedLibraries"))
	return rv
}


// SetPreloadedLibraries sets the value of the preloadedLibraries property.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/preloadedLibraries
func (t_ TileRenderPipelineDescriptor) SetPreloadedLibraries(value []objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPreloadedLibraries:"), value)
}
// The number of samples in each fragment.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/rasterSampleCount
func (t_ TileRenderPipelineDescriptor) RasterSampleCount() uint {
	rv := objc.Send[uint](t_.ID, objc.Sel("rasterSampleCount"))
	return rv
}


// SetRasterSampleCount sets the value of the rasterSampleCount property.
// The number of samples in each fragment.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/rasterSampleCount
func (t_ TileRenderPipelineDescriptor) SetRasterSampleCount(value uint) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRasterSampleCount:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/requiredThreadsPerThreadgroup
func (t_ TileRenderPipelineDescriptor) RequiredThreadsPerThreadgroup() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("requiredThreadsPerThreadgroup"))
	return rv
}


// SetRequiredThreadsPerThreadgroup sets the value of the requiredThreadsPerThreadgroup property.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/requiredThreadsPerThreadgroup
func (t_ TileRenderPipelineDescriptor) SetRequiredThreadsPerThreadgroup(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRequiredThreadsPerThreadgroup:"), value)
}
// A value that enables or disables shader validation for the pipeline.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/shaderValidation
func (t_ TileRenderPipelineDescriptor) ShaderValidation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("shaderValidation"))
	return rv
}


// SetShaderValidation sets the value of the shaderValidation property.
// A value that enables or disables shader validation for the pipeline.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/shaderValidation
func (t_ TileRenderPipelineDescriptor) SetShaderValidation(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setShaderValidation:"), value)
}
// A Boolean value that indicates whether you can use the pipeline to create new pipelines by adding binary functions to its callable functions list.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/supportAddingBinaryFunctions
func (t_ TileRenderPipelineDescriptor) SupportAddingBinaryFunctions() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("supportAddingBinaryFunctions"))
	return rv
}


// SetSupportAddingBinaryFunctions sets the value of the supportAddingBinaryFunctions property.
// A Boolean value that indicates whether you can use the pipeline to create new pipelines by adding binary functions to its callable functions list.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/supportAddingBinaryFunctions
func (t_ TileRenderPipelineDescriptor) SetSupportAddingBinaryFunctions(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSupportAddingBinaryFunctions:"), value)
}
// A Boolean value that indicates whether all threadgroups for this pipeline completely cover tiles.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/threadgroupSizeMatchesTileSize
func (t_ TileRenderPipelineDescriptor) ThreadgroupSizeMatchesTileSize() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("threadgroupSizeMatchesTileSize"))
	return rv
}


// SetThreadgroupSizeMatchesTileSize sets the value of the threadgroupSizeMatchesTileSize property.
// A Boolean value that indicates whether all threadgroups for this pipeline completely cover tiles.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/threadgroupSizeMatchesTileSize
func (t_ TileRenderPipelineDescriptor) SetThreadgroupSizeMatchesTileSize(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setThreadgroupSizeMatchesTileSize:"), value)
}
// An array that contains the buffer mutability options for a render pipeline’s tile function.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/tileBuffers
func (t_ TileRenderPipelineDescriptor) TileBuffers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("tileBuffers"))
	return rv
}

// The compute kernel or fragment function the pipeline calls.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/tileFunction
func (t_ TileRenderPipelineDescriptor) TileFunction() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("tileFunction"))
	return rv
}


// SetTileFunction sets the value of the tileFunction property.
// The compute kernel or fragment function the pipeline calls.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/tileFunction
func (t_ TileRenderPipelineDescriptor) SetTileFunction(value objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTileFunction:"), value)
}


