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
	Reset()
}

// Groups together properties you use to create a mesh render pipeline state object.
//
// Compared to , this interface doesn’t offer a mechanism to hint to Metal mutability of object, mesh, or fragment buffers. Additionally, when you use this descriptor, you don’t specify binary archives.
//
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

// Alloc allocates a new instance without initialization.
func (mc _MTL4MeshRenderPipelineDescriptorClass) Alloc() MTL4MeshRenderPipelineDescriptor {
	rv := objc.Send[MTL4MeshRenderPipelineDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Resets this descriptor to its default state.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/reset()
func (m_ MTL4MeshRenderPipelineDescriptor) Reset() {
	objc.Send[objc.ID](m_.ID, objc.Sel("reset"))
}

// Indicates whether to read and use the alpha channel fragment output of color attachments to compute a sample coverage mask.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/alphaToCoverageState
func (m_ MTL4MeshRenderPipelineDescriptor) AlphaToCoverageState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("alphaToCoverageState"))
	return rv
}


// SetAlphaToCoverageState sets the value of the alphaToCoverageState property.
// Indicates whether to read and use the alpha channel fragment output of color attachments to compute a sample coverage mask.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/alphaToCoverageState
func (m_ MTL4MeshRenderPipelineDescriptor) SetAlphaToCoverageState(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlphaToCoverageState:"), value)
}

// Indicates whether the pipeline forces alpha channel values of color attachments to the largest representable value.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/alphaToOneState
func (m_ MTL4MeshRenderPipelineDescriptor) AlphaToOneState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("alphaToOneState"))
	return rv
}


// SetAlphaToOneState sets the value of the alphaToOneState property.
// Indicates whether the pipeline forces alpha channel values of color attachments to the largest representable value.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/alphaToOneState
func (m_ MTL4MeshRenderPipelineDescriptor) SetAlphaToOneState(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlphaToOneState:"), value)
}

// Sets the logical-to-physical rendering remap state.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/colorAttachmentMappingState
func (m_ MTL4MeshRenderPipelineDescriptor) ColorAttachmentMappingState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("colorAttachmentMappingState"))
	return rv
}


// SetColorAttachmentMappingState sets the value of the colorAttachmentMappingState property.
// Sets the logical-to-physical rendering remap state.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/colorAttachmentMappingState
func (m_ MTL4MeshRenderPipelineDescriptor) SetColorAttachmentMappingState(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setColorAttachmentMappingState:"), value)
}

// Accesses an array containing descriptions of the color attachments this pipeline writes to.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/colorAttachments
func (m_ MTL4MeshRenderPipelineDescriptor) ColorAttachments() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("colorAttachments"))
	return rv
}

// Assigns a function descriptor representing the function this pipeline executes for each fragment.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/fragmentFunctionDescriptor
func (m_ MTL4MeshRenderPipelineDescriptor) FragmentFunctionDescriptor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("fragmentFunctionDescriptor"))
	return rv
}


// SetFragmentFunctionDescriptor sets the value of the fragmentFunctionDescriptor property.
// Assigns a function descriptor representing the function this pipeline executes for each fragment.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/fragmentFunctionDescriptor
func (m_ MTL4MeshRenderPipelineDescriptor) SetFragmentFunctionDescriptor(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFragmentFunctionDescriptor:"), value)
}

// Provides static linking information for the fragment stage of the render pipeline.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/fragmentStaticLinkingDescriptor
func (m_ MTL4MeshRenderPipelineDescriptor) FragmentStaticLinkingDescriptor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("fragmentStaticLinkingDescriptor"))
	return rv
}


// SetFragmentStaticLinkingDescriptor sets the value of the fragmentStaticLinkingDescriptor property.
// Provides static linking information for the fragment stage of the render pipeline.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/fragmentStaticLinkingDescriptor
func (m_ MTL4MeshRenderPipelineDescriptor) SetFragmentStaticLinkingDescriptor(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFragmentStaticLinkingDescriptor:"), value)
}

// Determines whether the pipeline rasterizes primitives.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/isRasterizationEnabled
func (m_ MTL4MeshRenderPipelineDescriptor) RasterizationEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("rasterizationEnabled"))
	return rv
}


// SetRasterizationEnabled sets the value of the rasterizationEnabled property.
// Determines whether the pipeline rasterizes primitives.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/isRasterizationEnabled
func (m_ MTL4MeshRenderPipelineDescriptor) SetRasterizationEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRasterizationEnabled:"), value)
}

// Controls the largest number of threads the pipeline state can execute when the object stage of a mesh render pipeline you create from this descriptor dispatches its mesh stage.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/maxTotalThreadgroupsPerMeshGrid
func (m_ MTL4MeshRenderPipelineDescriptor) MaxTotalThreadgroupsPerMeshGrid() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("maxTotalThreadgroupsPerMeshGrid"))
	return rv
}


// SetMaxTotalThreadgroupsPerMeshGrid sets the value of the maxTotalThreadgroupsPerMeshGrid property.
// Controls the largest number of threads the pipeline state can execute when the object stage of a mesh render pipeline you create from this descriptor dispatches its mesh stage.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/maxTotalThreadgroupsPerMeshGrid
func (m_ MTL4MeshRenderPipelineDescriptor) SetMaxTotalThreadgroupsPerMeshGrid(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxTotalThreadgroupsPerMeshGrid:"), value)
}

// Controls the largest number of threads the pipeline state can execute in a single mesh shader threadgroup dispatch.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/maxTotalThreadsPerMeshThreadgroup
func (m_ MTL4MeshRenderPipelineDescriptor) MaxTotalThreadsPerMeshThreadgroup() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("maxTotalThreadsPerMeshThreadgroup"))
	return rv
}


// SetMaxTotalThreadsPerMeshThreadgroup sets the value of the maxTotalThreadsPerMeshThreadgroup property.
// Controls the largest number of threads the pipeline state can execute in a single mesh shader threadgroup dispatch.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/maxTotalThreadsPerMeshThreadgroup
func (m_ MTL4MeshRenderPipelineDescriptor) SetMaxTotalThreadsPerMeshThreadgroup(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxTotalThreadsPerMeshThreadgroup:"), value)
}

// Controls the largest number of threads the pipeline state can execute in a single object shader threadgroup dispatch.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/maxTotalThreadsPerObjectThreadgroup
func (m_ MTL4MeshRenderPipelineDescriptor) MaxTotalThreadsPerObjectThreadgroup() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("maxTotalThreadsPerObjectThreadgroup"))
	return rv
}


// SetMaxTotalThreadsPerObjectThreadgroup sets the value of the maxTotalThreadsPerObjectThreadgroup property.
// Controls the largest number of threads the pipeline state can execute in a single object shader threadgroup dispatch.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/maxTotalThreadsPerObjectThreadgroup
func (m_ MTL4MeshRenderPipelineDescriptor) SetMaxTotalThreadsPerObjectThreadgroup(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxTotalThreadsPerObjectThreadgroup:"), value)
}

// Determines the maximum value that can you can pass as the pipeline’s amplification count.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/maxVertexAmplificationCount
func (m_ MTL4MeshRenderPipelineDescriptor) MaxVertexAmplificationCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("maxVertexAmplificationCount"))
	return rv
}


// SetMaxVertexAmplificationCount sets the value of the maxVertexAmplificationCount property.
// Determines the maximum value that can you can pass as the pipeline’s amplification count.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/maxVertexAmplificationCount
func (m_ MTL4MeshRenderPipelineDescriptor) SetMaxVertexAmplificationCount(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxVertexAmplificationCount:"), value)
}

// Assigns a function descriptor representing the function this pipeline executes for each primitive in the mesh shader stage.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/meshFunctionDescriptor
func (m_ MTL4MeshRenderPipelineDescriptor) MeshFunctionDescriptor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("meshFunctionDescriptor"))
	return rv
}


// SetMeshFunctionDescriptor sets the value of the meshFunctionDescriptor property.
// Assigns a function descriptor representing the function this pipeline executes for each primitive in the mesh shader stage.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/meshFunctionDescriptor
func (m_ MTL4MeshRenderPipelineDescriptor) SetMeshFunctionDescriptor(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMeshFunctionDescriptor:"), value)
}

// Provides static linking information for the mesh stage of the render pipeline.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/meshStaticLinkingDescriptor
func (m_ MTL4MeshRenderPipelineDescriptor) MeshStaticLinkingDescriptor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("meshStaticLinkingDescriptor"))
	return rv
}


// SetMeshStaticLinkingDescriptor sets the value of the meshStaticLinkingDescriptor property.
// Provides static linking information for the mesh stage of the render pipeline.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/meshStaticLinkingDescriptor
func (m_ MTL4MeshRenderPipelineDescriptor) SetMeshStaticLinkingDescriptor(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMeshStaticLinkingDescriptor:"), value)
}

// Provides a guarantee to Metal regarding the number of threadgroup threads for the mesh stage of a pipeline you create from this descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/meshThreadgroupSizeIsMultipleOfThreadExecutionWidth
func (m_ MTL4MeshRenderPipelineDescriptor) MeshThreadgroupSizeIsMultipleOfThreadExecutionWidth() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("meshThreadgroupSizeIsMultipleOfThreadExecutionWidth"))
	return rv
}


// SetMeshThreadgroupSizeIsMultipleOfThreadExecutionWidth sets the value of the meshThreadgroupSizeIsMultipleOfThreadExecutionWidth property.
// Provides a guarantee to Metal regarding the number of threadgroup threads for the mesh stage of a pipeline you create from this descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/meshThreadgroupSizeIsMultipleOfThreadExecutionWidth
func (m_ MTL4MeshRenderPipelineDescriptor) SetMeshThreadgroupSizeIsMultipleOfThreadExecutionWidth(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMeshThreadgroupSizeIsMultipleOfThreadExecutionWidth:"), value)
}

// Assigns a function descriptor representing the function this pipeline executes for each in the object shader stage.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/objectFunctionDescriptor
func (m_ MTL4MeshRenderPipelineDescriptor) ObjectFunctionDescriptor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("objectFunctionDescriptor"))
	return rv
}


// SetObjectFunctionDescriptor sets the value of the objectFunctionDescriptor property.
// Assigns a function descriptor representing the function this pipeline executes for each in the object shader stage.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/objectFunctionDescriptor
func (m_ MTL4MeshRenderPipelineDescriptor) SetObjectFunctionDescriptor(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setObjectFunctionDescriptor:"), value)
}

// Provides static linking information for the object stage of the render pipeline.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/objectStaticLinkingDescriptor
func (m_ MTL4MeshRenderPipelineDescriptor) ObjectStaticLinkingDescriptor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("objectStaticLinkingDescriptor"))
	return rv
}


// SetObjectStaticLinkingDescriptor sets the value of the objectStaticLinkingDescriptor property.
// Provides static linking information for the object stage of the render pipeline.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/objectStaticLinkingDescriptor
func (m_ MTL4MeshRenderPipelineDescriptor) SetObjectStaticLinkingDescriptor(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setObjectStaticLinkingDescriptor:"), value)
}

// Provides a guarantee to Metal regarding the number of threadgroup threads for the object stage of a pipeline you create from this descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/objectThreadgroupSizeIsMultipleOfThreadExecutionWidth
func (m_ MTL4MeshRenderPipelineDescriptor) ObjectThreadgroupSizeIsMultipleOfThreadExecutionWidth() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("objectThreadgroupSizeIsMultipleOfThreadExecutionWidth"))
	return rv
}


// SetObjectThreadgroupSizeIsMultipleOfThreadExecutionWidth sets the value of the objectThreadgroupSizeIsMultipleOfThreadExecutionWidth property.
// Provides a guarantee to Metal regarding the number of threadgroup threads for the object stage of a pipeline you create from this descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/objectThreadgroupSizeIsMultipleOfThreadExecutionWidth
func (m_ MTL4MeshRenderPipelineDescriptor) SetObjectThreadgroupSizeIsMultipleOfThreadExecutionWidth(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setObjectThreadgroupSizeIsMultipleOfThreadExecutionWidth:"), value)
}

// Reserves storage for the object-to-mesh stage payload.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/payloadMemoryLength
func (m_ MTL4MeshRenderPipelineDescriptor) PayloadMemoryLength() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("payloadMemoryLength"))
	return rv
}


// SetPayloadMemoryLength sets the value of the payloadMemoryLength property.
// Reserves storage for the object-to-mesh stage payload.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/payloadMemoryLength
func (m_ MTL4MeshRenderPipelineDescriptor) SetPayloadMemoryLength(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPayloadMemoryLength:"), value)
}

// Sets number of samples this pipeline applies for each fragment.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/rasterSampleCount
func (m_ MTL4MeshRenderPipelineDescriptor) RasterSampleCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("rasterSampleCount"))
	return rv
}


// SetRasterSampleCount sets the value of the rasterSampleCount property.
// Sets number of samples this pipeline applies for each fragment.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/rasterSampleCount
func (m_ MTL4MeshRenderPipelineDescriptor) SetRasterSampleCount(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRasterSampleCount:"), value)
}

// Controls the required number of mesh threads-per-threadgroup when drawing with a mesh shader pipeline you create from this descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/requiredThreadsPerMeshThreadgroup
func (m_ MTL4MeshRenderPipelineDescriptor) RequiredThreadsPerMeshThreadgroup() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("requiredThreadsPerMeshThreadgroup"))
	return rv
}


// SetRequiredThreadsPerMeshThreadgroup sets the value of the requiredThreadsPerMeshThreadgroup property.
// Controls the required number of mesh threads-per-threadgroup when drawing with a mesh shader pipeline you create from this descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/requiredThreadsPerMeshThreadgroup
func (m_ MTL4MeshRenderPipelineDescriptor) SetRequiredThreadsPerMeshThreadgroup(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRequiredThreadsPerMeshThreadgroup:"), value)
}

// Controls the required number of object threads-per-threadgroup when drawing with a mesh shader pipeline you create from this descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/requiredThreadsPerObjectThreadgroup
func (m_ MTL4MeshRenderPipelineDescriptor) RequiredThreadsPerObjectThreadgroup() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("requiredThreadsPerObjectThreadgroup"))
	return rv
}


// SetRequiredThreadsPerObjectThreadgroup sets the value of the requiredThreadsPerObjectThreadgroup property.
// Controls the required number of object threads-per-threadgroup when drawing with a mesh shader pipeline you create from this descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/requiredThreadsPerObjectThreadgroup
func (m_ MTL4MeshRenderPipelineDescriptor) SetRequiredThreadsPerObjectThreadgroup(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRequiredThreadsPerObjectThreadgroup:"), value)
}

// Indicates whether you can use the render pipeline to create new pipelines by adding binary functions to the fragment shader function’s callable functions list.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/supportFragmentBinaryLinking
func (m_ MTL4MeshRenderPipelineDescriptor) SupportFragmentBinaryLinking() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("supportFragmentBinaryLinking"))
	return rv
}


// SetSupportFragmentBinaryLinking sets the value of the supportFragmentBinaryLinking property.
// Indicates whether you can use the render pipeline to create new pipelines by adding binary functions to the fragment shader function’s callable functions list.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/supportFragmentBinaryLinking
func (m_ MTL4MeshRenderPipelineDescriptor) SetSupportFragmentBinaryLinking(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSupportFragmentBinaryLinking:"), value)
}

// Indicates whether the pipeline supports indirect command buffers.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/supportIndirectCommandBuffers
func (m_ MTL4MeshRenderPipelineDescriptor) SupportIndirectCommandBuffers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("supportIndirectCommandBuffers"))
	return rv
}


// SetSupportIndirectCommandBuffers sets the value of the supportIndirectCommandBuffers property.
// Indicates whether the pipeline supports indirect command buffers.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/supportIndirectCommandBuffers
func (m_ MTL4MeshRenderPipelineDescriptor) SetSupportIndirectCommandBuffers(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSupportIndirectCommandBuffers:"), value)
}

// Indicates whether you can use the render pipeline to create new pipelines by adding binary functions to the mesh shader function’s callable functions list.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/supportMeshBinaryLinking
func (m_ MTL4MeshRenderPipelineDescriptor) SupportMeshBinaryLinking() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("supportMeshBinaryLinking"))
	return rv
}


// SetSupportMeshBinaryLinking sets the value of the supportMeshBinaryLinking property.
// Indicates whether you can use the render pipeline to create new pipelines by adding binary functions to the mesh shader function’s callable functions list.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/supportMeshBinaryLinking
func (m_ MTL4MeshRenderPipelineDescriptor) SetSupportMeshBinaryLinking(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSupportMeshBinaryLinking:"), value)
}

// Indicates whether you can use the render pipeline to create new pipelines by adding binary functions to the object shader function’s callable functions list.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/supportObjectBinaryLinking
func (m_ MTL4MeshRenderPipelineDescriptor) SupportObjectBinaryLinking() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("supportObjectBinaryLinking"))
	return rv
}


// SetSupportObjectBinaryLinking sets the value of the supportObjectBinaryLinking property.
// Indicates whether you can use the render pipeline to create new pipelines by adding binary functions to the object shader function’s callable functions list.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/supportObjectBinaryLinking
func (m_ MTL4MeshRenderPipelineDescriptor) SetSupportObjectBinaryLinking(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSupportObjectBinaryLinking:"), value)
}



