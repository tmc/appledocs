// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [MTL4MeshRenderPipelineDescriptor] class.
var (
	MTL4MeshRenderPipelineDescriptorClass     _MTL4MeshRenderPipelineDescriptorClass
	MTL4MeshRenderPipelineDescriptorClassOnce sync.Once
)

func getMTL4MeshRenderPipelineDescriptorClass() _MTL4MeshRenderPipelineDescriptorClass {
	MTL4MeshRenderPipelineDescriptorClassOnce.Do(func() {
		MTL4MeshRenderPipelineDescriptorClass = _MTL4MeshRenderPipelineDescriptorClass{objc.GetClass("MTL4MeshRenderPipelineDescriptor")}
	})
	return MTL4MeshRenderPipelineDescriptorClass
}

type _MTL4MeshRenderPipelineDescriptorClass struct {
	class objc.Class
}





// An interface definition for the [MTL4MeshRenderPipelineDescriptor] class.
type IMTL4MeshRenderPipelineDescriptor interface {
	IMTL4PipelineDescriptor
	

	// properties:
	AlphaToCoverageState() MTL4AlphaToCoverageState
	SetAlphaToCoverageState(value MTL4AlphaToCoverageState)
	AlphaToOneState() MTL4AlphaToOneState
	SetAlphaToOneState(value MTL4AlphaToOneState)
	ColorAttachmentMappingState() MTL4LogicalToPhysicalColorAttachmentMappingState
	SetColorAttachmentMappingState(value MTL4LogicalToPhysicalColorAttachmentMappingState)
	ColorAttachments() IMTL4RenderPipelineColorAttachmentDescriptorArray
	FragmentFunctionDescriptor() IMTL4FunctionDescriptor
	SetFragmentFunctionDescriptor(value IMTL4FunctionDescriptor)
	FragmentStaticLinkingDescriptor() IMTL4StaticLinkingDescriptor
	SetFragmentStaticLinkingDescriptor(value IMTL4StaticLinkingDescriptor)
	RasterizationEnabled() bool
	SetRasterizationEnabled(value bool)
	MaxTotalThreadgroupsPerMeshGrid() uint
	SetMaxTotalThreadgroupsPerMeshGrid(value uint)
	MaxTotalThreadsPerMeshThreadgroup() uint
	SetMaxTotalThreadsPerMeshThreadgroup(value uint)
	MaxTotalThreadsPerObjectThreadgroup() uint
	SetMaxTotalThreadsPerObjectThreadgroup(value uint)
	MaxVertexAmplificationCount() uint
	SetMaxVertexAmplificationCount(value uint)
	MeshFunctionDescriptor() IMTL4FunctionDescriptor
	SetMeshFunctionDescriptor(value IMTL4FunctionDescriptor)
	MeshStaticLinkingDescriptor() IMTL4StaticLinkingDescriptor
	SetMeshStaticLinkingDescriptor(value IMTL4StaticLinkingDescriptor)
	MeshThreadgroupSizeIsMultipleOfThreadExecutionWidth() bool
	SetMeshThreadgroupSizeIsMultipleOfThreadExecutionWidth(value bool)
	ObjectFunctionDescriptor() IMTL4FunctionDescriptor
	SetObjectFunctionDescriptor(value IMTL4FunctionDescriptor)
	ObjectStaticLinkingDescriptor() IMTL4StaticLinkingDescriptor
	SetObjectStaticLinkingDescriptor(value IMTL4StaticLinkingDescriptor)
	ObjectThreadgroupSizeIsMultipleOfThreadExecutionWidth() bool
	SetObjectThreadgroupSizeIsMultipleOfThreadExecutionWidth(value bool)
	PayloadMemoryLength() uint
	SetPayloadMemoryLength(value uint)
	RasterSampleCount() uint
	SetRasterSampleCount(value uint)
	RequiredThreadsPerMeshThreadgroup() MTLSize
	SetRequiredThreadsPerMeshThreadgroup(value MTLSize)
	RequiredThreadsPerObjectThreadgroup() MTLSize
	SetRequiredThreadsPerObjectThreadgroup(value MTLSize)
	SupportFragmentBinaryLinking() bool
	SetSupportFragmentBinaryLinking(value bool)
	SupportIndirectCommandBuffers() MTL4IndirectCommandBufferSupportState
	SetSupportIndirectCommandBuffers(value MTL4IndirectCommandBufferSupportState)
	SupportMeshBinaryLinking() bool
	SetSupportMeshBinaryLinking(value bool)
	SupportObjectBinaryLinking() bool
	SetSupportObjectBinaryLinking(value bool)
	IsRasterizationEnabled() bool
	SetIsRasterizationEnabled(value bool)


	

	// methods:
	Reset()


}





// Alloc allocates a new instance without initialization.
func (mc _MTL4MeshRenderPipelineDescriptorClass) Alloc() MTL4MeshRenderPipelineDescriptor {
	rv := objc.Send[MTL4MeshRenderPipelineDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTL4MeshRenderPipelineDescriptorClass) New() MTL4MeshRenderPipelineDescriptor {
	rv := objc.Send[MTL4MeshRenderPipelineDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4MeshRenderPipelineDescriptor) Init() MTL4MeshRenderPipelineDescriptor {
	rv := objc.Send[MTL4MeshRenderPipelineDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4MeshRenderPipelineDescriptor) Autorelease() MTL4MeshRenderPipelineDescriptor {
	rv := objc.Send[MTL4MeshRenderPipelineDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4MeshRenderPipelineDescriptor creates a new MTL4MeshRenderPipelineDescriptor instance.
func NewMTL4MeshRenderPipelineDescriptor() MTL4MeshRenderPipelineDescriptor {
	return getMTL4MeshRenderPipelineDescriptorClass().New()
}





// Groups together properties you use to create a mesh render pipeline state object.
//
// Compared to , this interface doesn’t offer a mechanism to hint to Metal mutability of object, mesh, or fragment buffers. Additionally, when you use this descriptor, you don’t specify binary archives.


// Groups together properties you use to create a mesh render pipeline state object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor
type MTL4MeshRenderPipelineDescriptor struct {
	MTL4PipelineDescriptor
}

// MTL4MeshRenderPipelineDescriptorFrom constructs a [MTL4MeshRenderPipelineDescriptor] from an unsafe.Pointer.
//
// Groups together properties you use to create a mesh render pipeline state object.
func MTL4MeshRenderPipelineDescriptorFrom(ptr unsafe.Pointer) MTL4MeshRenderPipelineDescriptor {
	return MTL4MeshRenderPipelineDescriptor{
		MTL4PipelineDescriptor: MTL4PipelineDescriptorFrom(ptr),
	}
}




















// Resets this descriptor to its default state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/reset()
func (m_ MTL4MeshRenderPipelineDescriptor) Reset() {
	objc.Send[objc.ID](m_.ID, objc.Sel("reset"))
}







// Indicates whether to read and use the alpha channel fragment output of color attachments to compute a sample coverage mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/alphaToCoverageState
func (m_ MTL4MeshRenderPipelineDescriptor) AlphaToCoverageState() MTL4AlphaToCoverageState {
	rv := objc.Send[MTL4AlphaToCoverageState](m_.ID, objc.Sel("alphaToCoverageState"))
	return rv
}


// Indicates whether to read and use the alpha channel fragment output of color attachments to compute a sample coverage mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/alphaToCoverageState
func (m_ MTL4MeshRenderPipelineDescriptor) SetAlphaToCoverageState(value MTL4AlphaToCoverageState) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlphaToCoverageState:"), value)
}


// Indicates whether the pipeline forces alpha channel values of color attachments to the largest representable value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/alphaToOneState
func (m_ MTL4MeshRenderPipelineDescriptor) AlphaToOneState() MTL4AlphaToOneState {
	rv := objc.Send[MTL4AlphaToOneState](m_.ID, objc.Sel("alphaToOneState"))
	return rv
}


// Indicates whether the pipeline forces alpha channel values of color attachments to the largest representable value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/alphaToOneState
func (m_ MTL4MeshRenderPipelineDescriptor) SetAlphaToOneState(value MTL4AlphaToOneState) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlphaToOneState:"), value)
}


// Sets the logical-to-physical rendering remap state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/colorAttachmentMappingState
func (m_ MTL4MeshRenderPipelineDescriptor) ColorAttachmentMappingState() MTL4LogicalToPhysicalColorAttachmentMappingState {
	rv := objc.Send[MTL4LogicalToPhysicalColorAttachmentMappingState](m_.ID, objc.Sel("colorAttachmentMappingState"))
	return rv
}


// Sets the logical-to-physical rendering remap state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/colorAttachmentMappingState
func (m_ MTL4MeshRenderPipelineDescriptor) SetColorAttachmentMappingState(value MTL4LogicalToPhysicalColorAttachmentMappingState) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setColorAttachmentMappingState:"), value)
}


// Accesses an array containing descriptions of the color attachments this pipeline writes to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/colorAttachments
func (m_ MTL4MeshRenderPipelineDescriptor) ColorAttachments() IMTL4RenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Send[MTL4RenderPipelineColorAttachmentDescriptorArray](m_.ID, objc.Sel("colorAttachments"))
	return rv
}


// Assigns a function descriptor representing the function this pipeline executes for each fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/fragmentFunctionDescriptor
func (m_ MTL4MeshRenderPipelineDescriptor) FragmentFunctionDescriptor() IMTL4FunctionDescriptor {
	rv := objc.Send[MTL4FunctionDescriptor](m_.ID, objc.Sel("fragmentFunctionDescriptor"))
	return rv
}


// Assigns a function descriptor representing the function this pipeline executes for each fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/fragmentFunctionDescriptor
func (m_ MTL4MeshRenderPipelineDescriptor) SetFragmentFunctionDescriptor(value IMTL4FunctionDescriptor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFragmentFunctionDescriptor:"), value)
}


// Provides static linking information for the fragment stage of the render pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/fragmentStaticLinkingDescriptor
func (m_ MTL4MeshRenderPipelineDescriptor) FragmentStaticLinkingDescriptor() IMTL4StaticLinkingDescriptor {
	rv := objc.Send[MTL4StaticLinkingDescriptor](m_.ID, objc.Sel("fragmentStaticLinkingDescriptor"))
	return rv
}


// Provides static linking information for the fragment stage of the render pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/fragmentStaticLinkingDescriptor
func (m_ MTL4MeshRenderPipelineDescriptor) SetFragmentStaticLinkingDescriptor(value IMTL4StaticLinkingDescriptor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFragmentStaticLinkingDescriptor:"), value)
}


// Determines whether the pipeline rasterizes primitives.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/isRasterizationEnabled
func (m_ MTL4MeshRenderPipelineDescriptor) RasterizationEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("rasterizationEnabled"))
	return rv
}


// Determines whether the pipeline rasterizes primitives.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/isRasterizationEnabled
func (m_ MTL4MeshRenderPipelineDescriptor) SetRasterizationEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRasterizationEnabled:"), value)
}


// Controls the largest number of threads the pipeline state can execute when the object stage of a mesh render pipeline you create from this descriptor dispatches its mesh stage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/maxTotalThreadgroupsPerMeshGrid
func (m_ MTL4MeshRenderPipelineDescriptor) MaxTotalThreadgroupsPerMeshGrid() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("maxTotalThreadgroupsPerMeshGrid"))
	return rv
}


// Controls the largest number of threads the pipeline state can execute when the object stage of a mesh render pipeline you create from this descriptor dispatches its mesh stage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/maxTotalThreadgroupsPerMeshGrid
func (m_ MTL4MeshRenderPipelineDescriptor) SetMaxTotalThreadgroupsPerMeshGrid(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxTotalThreadgroupsPerMeshGrid:"), value)
}


// Controls the largest number of threads the pipeline state can execute in a single mesh shader threadgroup dispatch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/maxTotalThreadsPerMeshThreadgroup
func (m_ MTL4MeshRenderPipelineDescriptor) MaxTotalThreadsPerMeshThreadgroup() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("maxTotalThreadsPerMeshThreadgroup"))
	return rv
}


// Controls the largest number of threads the pipeline state can execute in a single mesh shader threadgroup dispatch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/maxTotalThreadsPerMeshThreadgroup
func (m_ MTL4MeshRenderPipelineDescriptor) SetMaxTotalThreadsPerMeshThreadgroup(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxTotalThreadsPerMeshThreadgroup:"), value)
}


// Controls the largest number of threads the pipeline state can execute in a single object shader threadgroup dispatch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/maxTotalThreadsPerObjectThreadgroup
func (m_ MTL4MeshRenderPipelineDescriptor) MaxTotalThreadsPerObjectThreadgroup() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("maxTotalThreadsPerObjectThreadgroup"))
	return rv
}


// Controls the largest number of threads the pipeline state can execute in a single object shader threadgroup dispatch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/maxTotalThreadsPerObjectThreadgroup
func (m_ MTL4MeshRenderPipelineDescriptor) SetMaxTotalThreadsPerObjectThreadgroup(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxTotalThreadsPerObjectThreadgroup:"), value)
}


// Determines the maximum value that can you can pass as the pipeline’s amplification count.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/maxVertexAmplificationCount
func (m_ MTL4MeshRenderPipelineDescriptor) MaxVertexAmplificationCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("maxVertexAmplificationCount"))
	return rv
}


// Determines the maximum value that can you can pass as the pipeline’s amplification count.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/maxVertexAmplificationCount
func (m_ MTL4MeshRenderPipelineDescriptor) SetMaxVertexAmplificationCount(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxVertexAmplificationCount:"), value)
}


// Assigns a function descriptor representing the function this pipeline executes for each primitive in the mesh shader stage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/meshFunctionDescriptor
func (m_ MTL4MeshRenderPipelineDescriptor) MeshFunctionDescriptor() IMTL4FunctionDescriptor {
	rv := objc.Send[MTL4FunctionDescriptor](m_.ID, objc.Sel("meshFunctionDescriptor"))
	return rv
}


// Assigns a function descriptor representing the function this pipeline executes for each primitive in the mesh shader stage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/meshFunctionDescriptor
func (m_ MTL4MeshRenderPipelineDescriptor) SetMeshFunctionDescriptor(value IMTL4FunctionDescriptor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMeshFunctionDescriptor:"), value)
}


// Provides static linking information for the mesh stage of the render pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/meshStaticLinkingDescriptor
func (m_ MTL4MeshRenderPipelineDescriptor) MeshStaticLinkingDescriptor() IMTL4StaticLinkingDescriptor {
	rv := objc.Send[MTL4StaticLinkingDescriptor](m_.ID, objc.Sel("meshStaticLinkingDescriptor"))
	return rv
}


// Provides static linking information for the mesh stage of the render pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/meshStaticLinkingDescriptor
func (m_ MTL4MeshRenderPipelineDescriptor) SetMeshStaticLinkingDescriptor(value IMTL4StaticLinkingDescriptor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMeshStaticLinkingDescriptor:"), value)
}


// Provides a guarantee to Metal regarding the number of threadgroup threads for the mesh stage of a pipeline you create from this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/meshThreadgroupSizeIsMultipleOfThreadExecutionWidth
func (m_ MTL4MeshRenderPipelineDescriptor) MeshThreadgroupSizeIsMultipleOfThreadExecutionWidth() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("meshThreadgroupSizeIsMultipleOfThreadExecutionWidth"))
	return rv
}


// Provides a guarantee to Metal regarding the number of threadgroup threads for the mesh stage of a pipeline you create from this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/meshThreadgroupSizeIsMultipleOfThreadExecutionWidth
func (m_ MTL4MeshRenderPipelineDescriptor) SetMeshThreadgroupSizeIsMultipleOfThreadExecutionWidth(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMeshThreadgroupSizeIsMultipleOfThreadExecutionWidth:"), value)
}


// Assigns a function descriptor representing the function this pipeline executes for each in the object shader stage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/objectFunctionDescriptor
func (m_ MTL4MeshRenderPipelineDescriptor) ObjectFunctionDescriptor() IMTL4FunctionDescriptor {
	rv := objc.Send[MTL4FunctionDescriptor](m_.ID, objc.Sel("objectFunctionDescriptor"))
	return rv
}


// Assigns a function descriptor representing the function this pipeline executes for each in the object shader stage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/objectFunctionDescriptor
func (m_ MTL4MeshRenderPipelineDescriptor) SetObjectFunctionDescriptor(value IMTL4FunctionDescriptor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setObjectFunctionDescriptor:"), value)
}


// Provides static linking information for the object stage of the render pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/objectStaticLinkingDescriptor
func (m_ MTL4MeshRenderPipelineDescriptor) ObjectStaticLinkingDescriptor() IMTL4StaticLinkingDescriptor {
	rv := objc.Send[MTL4StaticLinkingDescriptor](m_.ID, objc.Sel("objectStaticLinkingDescriptor"))
	return rv
}


// Provides static linking information for the object stage of the render pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/objectStaticLinkingDescriptor
func (m_ MTL4MeshRenderPipelineDescriptor) SetObjectStaticLinkingDescriptor(value IMTL4StaticLinkingDescriptor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setObjectStaticLinkingDescriptor:"), value)
}


// Provides a guarantee to Metal regarding the number of threadgroup threads for the object stage of a pipeline you create from this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/objectThreadgroupSizeIsMultipleOfThreadExecutionWidth
func (m_ MTL4MeshRenderPipelineDescriptor) ObjectThreadgroupSizeIsMultipleOfThreadExecutionWidth() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("objectThreadgroupSizeIsMultipleOfThreadExecutionWidth"))
	return rv
}


// Provides a guarantee to Metal regarding the number of threadgroup threads for the object stage of a pipeline you create from this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/objectThreadgroupSizeIsMultipleOfThreadExecutionWidth
func (m_ MTL4MeshRenderPipelineDescriptor) SetObjectThreadgroupSizeIsMultipleOfThreadExecutionWidth(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setObjectThreadgroupSizeIsMultipleOfThreadExecutionWidth:"), value)
}


// Reserves storage for the object-to-mesh stage payload.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/payloadMemoryLength
func (m_ MTL4MeshRenderPipelineDescriptor) PayloadMemoryLength() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("payloadMemoryLength"))
	return rv
}


// Reserves storage for the object-to-mesh stage payload.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/payloadMemoryLength
func (m_ MTL4MeshRenderPipelineDescriptor) SetPayloadMemoryLength(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPayloadMemoryLength:"), value)
}


// Sets number of samples this pipeline applies for each fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/rasterSampleCount
func (m_ MTL4MeshRenderPipelineDescriptor) RasterSampleCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("rasterSampleCount"))
	return rv
}


// Sets number of samples this pipeline applies for each fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/rasterSampleCount
func (m_ MTL4MeshRenderPipelineDescriptor) SetRasterSampleCount(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRasterSampleCount:"), value)
}


// Controls the required number of mesh threads-per-threadgroup when drawing with a mesh shader pipeline you create from this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/requiredThreadsPerMeshThreadgroup
func (m_ MTL4MeshRenderPipelineDescriptor) RequiredThreadsPerMeshThreadgroup() MTLSize {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("requiredThreadsPerMeshThreadgroup"))
	return rv
}


// Controls the required number of mesh threads-per-threadgroup when drawing with a mesh shader pipeline you create from this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/requiredThreadsPerMeshThreadgroup
func (m_ MTL4MeshRenderPipelineDescriptor) SetRequiredThreadsPerMeshThreadgroup(value MTLSize) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRequiredThreadsPerMeshThreadgroup:"), value)
}


// Controls the required number of object threads-per-threadgroup when drawing with a mesh shader pipeline you create from this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/requiredThreadsPerObjectThreadgroup
func (m_ MTL4MeshRenderPipelineDescriptor) RequiredThreadsPerObjectThreadgroup() MTLSize {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("requiredThreadsPerObjectThreadgroup"))
	return rv
}


// Controls the required number of object threads-per-threadgroup when drawing with a mesh shader pipeline you create from this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/requiredThreadsPerObjectThreadgroup
func (m_ MTL4MeshRenderPipelineDescriptor) SetRequiredThreadsPerObjectThreadgroup(value MTLSize) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRequiredThreadsPerObjectThreadgroup:"), value)
}


// Indicates whether you can use the render pipeline to create new pipelines by adding binary functions to the fragment shader function’s callable functions list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/supportFragmentBinaryLinking
func (m_ MTL4MeshRenderPipelineDescriptor) SupportFragmentBinaryLinking() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("supportFragmentBinaryLinking"))
	return rv
}


// Indicates whether you can use the render pipeline to create new pipelines by adding binary functions to the fragment shader function’s callable functions list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/supportFragmentBinaryLinking
func (m_ MTL4MeshRenderPipelineDescriptor) SetSupportFragmentBinaryLinking(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSupportFragmentBinaryLinking:"), value)
}


// Indicates whether the pipeline supports indirect command buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/supportIndirectCommandBuffers
func (m_ MTL4MeshRenderPipelineDescriptor) SupportIndirectCommandBuffers() MTL4IndirectCommandBufferSupportState {
	rv := objc.Send[MTL4IndirectCommandBufferSupportState](m_.ID, objc.Sel("supportIndirectCommandBuffers"))
	return rv
}


// Indicates whether the pipeline supports indirect command buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/supportIndirectCommandBuffers
func (m_ MTL4MeshRenderPipelineDescriptor) SetSupportIndirectCommandBuffers(value MTL4IndirectCommandBufferSupportState) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSupportIndirectCommandBuffers:"), value)
}


// Indicates whether you can use the render pipeline to create new pipelines by adding binary functions to the mesh shader function’s callable functions list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/supportMeshBinaryLinking
func (m_ MTL4MeshRenderPipelineDescriptor) SupportMeshBinaryLinking() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("supportMeshBinaryLinking"))
	return rv
}


// Indicates whether you can use the render pipeline to create new pipelines by adding binary functions to the mesh shader function’s callable functions list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/supportMeshBinaryLinking
func (m_ MTL4MeshRenderPipelineDescriptor) SetSupportMeshBinaryLinking(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSupportMeshBinaryLinking:"), value)
}


// Indicates whether you can use the render pipeline to create new pipelines by adding binary functions to the object shader function’s callable functions list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/supportObjectBinaryLinking
func (m_ MTL4MeshRenderPipelineDescriptor) SupportObjectBinaryLinking() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("supportObjectBinaryLinking"))
	return rv
}


// Indicates whether you can use the render pipeline to create new pipelines by adding binary functions to the object shader function’s callable functions list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/supportObjectBinaryLinking
func (m_ MTL4MeshRenderPipelineDescriptor) SetSupportObjectBinaryLinking(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSupportObjectBinaryLinking:"), value)
}


// Determines whether the pipeline rasterizes primitives.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4meshrenderpipelinedescriptor/israsterizationenabled
func (m_ MTL4MeshRenderPipelineDescriptor) IsRasterizationEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isRasterizationEnabled"))
	return rv
}


// Determines whether the pipeline rasterizes primitives.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4meshrenderpipelinedescriptor/israsterizationenabled
func (m_ MTL4MeshRenderPipelineDescriptor) SetIsRasterizationEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsRasterizationEnabled:"), value)
}








