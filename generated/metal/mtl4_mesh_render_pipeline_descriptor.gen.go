// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MTL4MeshRenderPipelineDescriptor */


/* debug [class_header]: Header for MTL4MeshRenderPipelineDescriptor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTL4MeshRenderPipelineDescriptor */
// An interface definition for the [MTL4MeshRenderPipelineDescriptor] class.
type IMTL4MeshRenderPipelineDescriptor interface {
	IMTL4PipelineDescriptor
	
/* debug [class_interface_properties]: Properties for MTL4MeshRenderPipelineDescriptor */
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
	RequiredThreadsPerMeshThreadgroup() objc.IObject /* cross-framework: MTLSize */
	SetRequiredThreadsPerMeshThreadgroup(value objc.IObject /* cross-framework: MTLSize */)
	RequiredThreadsPerObjectThreadgroup() objc.IObject /* cross-framework: MTLSize */
	SetRequiredThreadsPerObjectThreadgroup(value objc.IObject /* cross-framework: MTLSize */)
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTL4MeshRenderPipelineDescriptor */
	// methods:
	Reset()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTL4MeshRenderPipelineDescriptor */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTL4MeshRenderPipelineDescriptor */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTL4MeshRenderPipelineDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTL4MeshRenderPipelineDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTL4MeshRenderPipelineDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTL4MeshRenderPipelineDescriptor */

// Resets this descriptor to its default state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/reset()
func (m_ MTL4MeshRenderPipelineDescriptor) Reset() {
	objc.Send[objc.ID](m_.ID, objc.Sel("reset"))
}/* debug [instance_methods/method]: Reset */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTL4MeshRenderPipelineDescriptor */

// Indicates whether to read and use the alpha channel fragment output of color attachments to compute a sample coverage mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/alphaToCoverageState
func (m_ MTL4MeshRenderPipelineDescriptor) AlphaToCoverageState() MTL4AlphaToCoverageState {
	rv := objc.Send[MTL4AlphaToCoverageState](m_.ID, objc.Sel("alphaToCoverageState"))
	return rv
}/* debug [instance_properties/getter]: alphaToCoverageState */


// Indicates whether to read and use the alpha channel fragment output of color attachments to compute a sample coverage mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/alphaToCoverageState
func (m_ MTL4MeshRenderPipelineDescriptor) SetAlphaToCoverageState(value MTL4AlphaToCoverageState) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlphaToCoverageState:"), value)
}/* debug [instance_properties/setter]: alphaToCoverageState */


// Indicates whether the pipeline forces alpha channel values of color attachments to the largest representable value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/alphaToOneState
func (m_ MTL4MeshRenderPipelineDescriptor) AlphaToOneState() MTL4AlphaToOneState {
	rv := objc.Send[MTL4AlphaToOneState](m_.ID, objc.Sel("alphaToOneState"))
	return rv
}/* debug [instance_properties/getter]: alphaToOneState */


// Indicates whether the pipeline forces alpha channel values of color attachments to the largest representable value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/alphaToOneState
func (m_ MTL4MeshRenderPipelineDescriptor) SetAlphaToOneState(value MTL4AlphaToOneState) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlphaToOneState:"), value)
}/* debug [instance_properties/setter]: alphaToOneState */


// Sets the logical-to-physical rendering remap state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/colorAttachmentMappingState
func (m_ MTL4MeshRenderPipelineDescriptor) ColorAttachmentMappingState() MTL4LogicalToPhysicalColorAttachmentMappingState {
	rv := objc.Send[MTL4LogicalToPhysicalColorAttachmentMappingState](m_.ID, objc.Sel("colorAttachmentMappingState"))
	return rv
}/* debug [instance_properties/getter]: colorAttachmentMappingState */


// Sets the logical-to-physical rendering remap state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/colorAttachmentMappingState
func (m_ MTL4MeshRenderPipelineDescriptor) SetColorAttachmentMappingState(value MTL4LogicalToPhysicalColorAttachmentMappingState) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setColorAttachmentMappingState:"), value)
}/* debug [instance_properties/setter]: colorAttachmentMappingState */


// Accesses an array containing descriptions of the color attachments this pipeline writes to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/colorAttachments
func (m_ MTL4MeshRenderPipelineDescriptor) ColorAttachments() IMTL4RenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Send[MTL4RenderPipelineColorAttachmentDescriptorArray](m_.ID, objc.Sel("colorAttachments"))
	return rv
}/* debug [instance_properties/getter]: colorAttachments */


// Assigns a function descriptor representing the function this pipeline executes for each fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/fragmentFunctionDescriptor
func (m_ MTL4MeshRenderPipelineDescriptor) FragmentFunctionDescriptor() IMTL4FunctionDescriptor {
	rv := objc.Send[MTL4FunctionDescriptor](m_.ID, objc.Sel("fragmentFunctionDescriptor"))
	return rv
}/* debug [instance_properties/getter]: fragmentFunctionDescriptor */


// Assigns a function descriptor representing the function this pipeline executes for each fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/fragmentFunctionDescriptor
func (m_ MTL4MeshRenderPipelineDescriptor) SetFragmentFunctionDescriptor(value IMTL4FunctionDescriptor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFragmentFunctionDescriptor:"), value)
}/* debug [instance_properties/setter]: fragmentFunctionDescriptor */


// Provides static linking information for the fragment stage of the render pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/fragmentStaticLinkingDescriptor
func (m_ MTL4MeshRenderPipelineDescriptor) FragmentStaticLinkingDescriptor() IMTL4StaticLinkingDescriptor {
	rv := objc.Send[MTL4StaticLinkingDescriptor](m_.ID, objc.Sel("fragmentStaticLinkingDescriptor"))
	return rv
}/* debug [instance_properties/getter]: fragmentStaticLinkingDescriptor */


// Provides static linking information for the fragment stage of the render pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/fragmentStaticLinkingDescriptor
func (m_ MTL4MeshRenderPipelineDescriptor) SetFragmentStaticLinkingDescriptor(value IMTL4StaticLinkingDescriptor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFragmentStaticLinkingDescriptor:"), value)
}/* debug [instance_properties/setter]: fragmentStaticLinkingDescriptor */


// Determines whether the pipeline rasterizes primitives.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/isRasterizationEnabled
func (m_ MTL4MeshRenderPipelineDescriptor) RasterizationEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("rasterizationEnabled"))
	return rv
}/* debug [instance_properties/getter]: rasterizationEnabled */


// Determines whether the pipeline rasterizes primitives.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/isRasterizationEnabled
func (m_ MTL4MeshRenderPipelineDescriptor) SetRasterizationEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRasterizationEnabled:"), value)
}/* debug [instance_properties/setter]: rasterizationEnabled */


// Controls the largest number of threads the pipeline state can execute when the object stage of a mesh render pipeline you create from this descriptor dispatches its mesh stage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/maxTotalThreadgroupsPerMeshGrid
func (m_ MTL4MeshRenderPipelineDescriptor) MaxTotalThreadgroupsPerMeshGrid() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("maxTotalThreadgroupsPerMeshGrid"))
	return rv
}/* debug [instance_properties/getter]: maxTotalThreadgroupsPerMeshGrid */


// Controls the largest number of threads the pipeline state can execute when the object stage of a mesh render pipeline you create from this descriptor dispatches its mesh stage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/maxTotalThreadgroupsPerMeshGrid
func (m_ MTL4MeshRenderPipelineDescriptor) SetMaxTotalThreadgroupsPerMeshGrid(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxTotalThreadgroupsPerMeshGrid:"), value)
}/* debug [instance_properties/setter]: maxTotalThreadgroupsPerMeshGrid */


// Controls the largest number of threads the pipeline state can execute in a single mesh shader threadgroup dispatch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/maxTotalThreadsPerMeshThreadgroup
func (m_ MTL4MeshRenderPipelineDescriptor) MaxTotalThreadsPerMeshThreadgroup() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("maxTotalThreadsPerMeshThreadgroup"))
	return rv
}/* debug [instance_properties/getter]: maxTotalThreadsPerMeshThreadgroup */


// Controls the largest number of threads the pipeline state can execute in a single mesh shader threadgroup dispatch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/maxTotalThreadsPerMeshThreadgroup
func (m_ MTL4MeshRenderPipelineDescriptor) SetMaxTotalThreadsPerMeshThreadgroup(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxTotalThreadsPerMeshThreadgroup:"), value)
}/* debug [instance_properties/setter]: maxTotalThreadsPerMeshThreadgroup */


// Controls the largest number of threads the pipeline state can execute in a single object shader threadgroup dispatch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/maxTotalThreadsPerObjectThreadgroup
func (m_ MTL4MeshRenderPipelineDescriptor) MaxTotalThreadsPerObjectThreadgroup() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("maxTotalThreadsPerObjectThreadgroup"))
	return rv
}/* debug [instance_properties/getter]: maxTotalThreadsPerObjectThreadgroup */


// Controls the largest number of threads the pipeline state can execute in a single object shader threadgroup dispatch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/maxTotalThreadsPerObjectThreadgroup
func (m_ MTL4MeshRenderPipelineDescriptor) SetMaxTotalThreadsPerObjectThreadgroup(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxTotalThreadsPerObjectThreadgroup:"), value)
}/* debug [instance_properties/setter]: maxTotalThreadsPerObjectThreadgroup */


// Determines the maximum value that can you can pass as the pipeline’s amplification count.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/maxVertexAmplificationCount
func (m_ MTL4MeshRenderPipelineDescriptor) MaxVertexAmplificationCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("maxVertexAmplificationCount"))
	return rv
}/* debug [instance_properties/getter]: maxVertexAmplificationCount */


// Determines the maximum value that can you can pass as the pipeline’s amplification count.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/maxVertexAmplificationCount
func (m_ MTL4MeshRenderPipelineDescriptor) SetMaxVertexAmplificationCount(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxVertexAmplificationCount:"), value)
}/* debug [instance_properties/setter]: maxVertexAmplificationCount */


// Assigns a function descriptor representing the function this pipeline executes for each primitive in the mesh shader stage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/meshFunctionDescriptor
func (m_ MTL4MeshRenderPipelineDescriptor) MeshFunctionDescriptor() IMTL4FunctionDescriptor {
	rv := objc.Send[MTL4FunctionDescriptor](m_.ID, objc.Sel("meshFunctionDescriptor"))
	return rv
}/* debug [instance_properties/getter]: meshFunctionDescriptor */


// Assigns a function descriptor representing the function this pipeline executes for each primitive in the mesh shader stage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/meshFunctionDescriptor
func (m_ MTL4MeshRenderPipelineDescriptor) SetMeshFunctionDescriptor(value IMTL4FunctionDescriptor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMeshFunctionDescriptor:"), value)
}/* debug [instance_properties/setter]: meshFunctionDescriptor */


// Provides static linking information for the mesh stage of the render pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/meshStaticLinkingDescriptor
func (m_ MTL4MeshRenderPipelineDescriptor) MeshStaticLinkingDescriptor() IMTL4StaticLinkingDescriptor {
	rv := objc.Send[MTL4StaticLinkingDescriptor](m_.ID, objc.Sel("meshStaticLinkingDescriptor"))
	return rv
}/* debug [instance_properties/getter]: meshStaticLinkingDescriptor */


// Provides static linking information for the mesh stage of the render pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/meshStaticLinkingDescriptor
func (m_ MTL4MeshRenderPipelineDescriptor) SetMeshStaticLinkingDescriptor(value IMTL4StaticLinkingDescriptor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMeshStaticLinkingDescriptor:"), value)
}/* debug [instance_properties/setter]: meshStaticLinkingDescriptor */


// Provides a guarantee to Metal regarding the number of threadgroup threads for the mesh stage of a pipeline you create from this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/meshThreadgroupSizeIsMultipleOfThreadExecutionWidth
func (m_ MTL4MeshRenderPipelineDescriptor) MeshThreadgroupSizeIsMultipleOfThreadExecutionWidth() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("meshThreadgroupSizeIsMultipleOfThreadExecutionWidth"))
	return rv
}/* debug [instance_properties/getter]: meshThreadgroupSizeIsMultipleOfThreadExecutionWidth */


// Provides a guarantee to Metal regarding the number of threadgroup threads for the mesh stage of a pipeline you create from this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/meshThreadgroupSizeIsMultipleOfThreadExecutionWidth
func (m_ MTL4MeshRenderPipelineDescriptor) SetMeshThreadgroupSizeIsMultipleOfThreadExecutionWidth(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMeshThreadgroupSizeIsMultipleOfThreadExecutionWidth:"), value)
}/* debug [instance_properties/setter]: meshThreadgroupSizeIsMultipleOfThreadExecutionWidth */


// Assigns a function descriptor representing the function this pipeline executes for each in the object shader stage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/objectFunctionDescriptor
func (m_ MTL4MeshRenderPipelineDescriptor) ObjectFunctionDescriptor() IMTL4FunctionDescriptor {
	rv := objc.Send[MTL4FunctionDescriptor](m_.ID, objc.Sel("objectFunctionDescriptor"))
	return rv
}/* debug [instance_properties/getter]: objectFunctionDescriptor */


// Assigns a function descriptor representing the function this pipeline executes for each in the object shader stage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/objectFunctionDescriptor
func (m_ MTL4MeshRenderPipelineDescriptor) SetObjectFunctionDescriptor(value IMTL4FunctionDescriptor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setObjectFunctionDescriptor:"), value)
}/* debug [instance_properties/setter]: objectFunctionDescriptor */


// Provides static linking information for the object stage of the render pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/objectStaticLinkingDescriptor
func (m_ MTL4MeshRenderPipelineDescriptor) ObjectStaticLinkingDescriptor() IMTL4StaticLinkingDescriptor {
	rv := objc.Send[MTL4StaticLinkingDescriptor](m_.ID, objc.Sel("objectStaticLinkingDescriptor"))
	return rv
}/* debug [instance_properties/getter]: objectStaticLinkingDescriptor */


// Provides static linking information for the object stage of the render pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/objectStaticLinkingDescriptor
func (m_ MTL4MeshRenderPipelineDescriptor) SetObjectStaticLinkingDescriptor(value IMTL4StaticLinkingDescriptor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setObjectStaticLinkingDescriptor:"), value)
}/* debug [instance_properties/setter]: objectStaticLinkingDescriptor */


// Provides a guarantee to Metal regarding the number of threadgroup threads for the object stage of a pipeline you create from this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/objectThreadgroupSizeIsMultipleOfThreadExecutionWidth
func (m_ MTL4MeshRenderPipelineDescriptor) ObjectThreadgroupSizeIsMultipleOfThreadExecutionWidth() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("objectThreadgroupSizeIsMultipleOfThreadExecutionWidth"))
	return rv
}/* debug [instance_properties/getter]: objectThreadgroupSizeIsMultipleOfThreadExecutionWidth */


// Provides a guarantee to Metal regarding the number of threadgroup threads for the object stage of a pipeline you create from this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/objectThreadgroupSizeIsMultipleOfThreadExecutionWidth
func (m_ MTL4MeshRenderPipelineDescriptor) SetObjectThreadgroupSizeIsMultipleOfThreadExecutionWidth(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setObjectThreadgroupSizeIsMultipleOfThreadExecutionWidth:"), value)
}/* debug [instance_properties/setter]: objectThreadgroupSizeIsMultipleOfThreadExecutionWidth */


// Reserves storage for the object-to-mesh stage payload.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/payloadMemoryLength
func (m_ MTL4MeshRenderPipelineDescriptor) PayloadMemoryLength() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("payloadMemoryLength"))
	return rv
}/* debug [instance_properties/getter]: payloadMemoryLength */


// Reserves storage for the object-to-mesh stage payload.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/payloadMemoryLength
func (m_ MTL4MeshRenderPipelineDescriptor) SetPayloadMemoryLength(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPayloadMemoryLength:"), value)
}/* debug [instance_properties/setter]: payloadMemoryLength */


// Sets number of samples this pipeline applies for each fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/rasterSampleCount
func (m_ MTL4MeshRenderPipelineDescriptor) RasterSampleCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("rasterSampleCount"))
	return rv
}/* debug [instance_properties/getter]: rasterSampleCount */


// Sets number of samples this pipeline applies for each fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/rasterSampleCount
func (m_ MTL4MeshRenderPipelineDescriptor) SetRasterSampleCount(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRasterSampleCount:"), value)
}/* debug [instance_properties/setter]: rasterSampleCount */


// Controls the required number of mesh threads-per-threadgroup when drawing with a mesh shader pipeline you create from this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/requiredThreadsPerMeshThreadgroup
func (m_ MTL4MeshRenderPipelineDescriptor) RequiredThreadsPerMeshThreadgroup() objc.IObject /* cross-framework: MTLSize */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("requiredThreadsPerMeshThreadgroup"))
	return rv
}/* debug [instance_properties/getter]: requiredThreadsPerMeshThreadgroup */


// Controls the required number of mesh threads-per-threadgroup when drawing with a mesh shader pipeline you create from this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/requiredThreadsPerMeshThreadgroup
func (m_ MTL4MeshRenderPipelineDescriptor) SetRequiredThreadsPerMeshThreadgroup(value objc.IObject /* cross-framework: MTLSize */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRequiredThreadsPerMeshThreadgroup:"), value)
}/* debug [instance_properties/setter]: requiredThreadsPerMeshThreadgroup */


// Controls the required number of object threads-per-threadgroup when drawing with a mesh shader pipeline you create from this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/requiredThreadsPerObjectThreadgroup
func (m_ MTL4MeshRenderPipelineDescriptor) RequiredThreadsPerObjectThreadgroup() objc.IObject /* cross-framework: MTLSize */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("requiredThreadsPerObjectThreadgroup"))
	return rv
}/* debug [instance_properties/getter]: requiredThreadsPerObjectThreadgroup */


// Controls the required number of object threads-per-threadgroup when drawing with a mesh shader pipeline you create from this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/requiredThreadsPerObjectThreadgroup
func (m_ MTL4MeshRenderPipelineDescriptor) SetRequiredThreadsPerObjectThreadgroup(value objc.IObject /* cross-framework: MTLSize */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRequiredThreadsPerObjectThreadgroup:"), value)
}/* debug [instance_properties/setter]: requiredThreadsPerObjectThreadgroup */


// Indicates whether you can use the render pipeline to create new pipelines by adding binary functions to the fragment shader function’s callable functions list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/supportFragmentBinaryLinking
func (m_ MTL4MeshRenderPipelineDescriptor) SupportFragmentBinaryLinking() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("supportFragmentBinaryLinking"))
	return rv
}/* debug [instance_properties/getter]: supportFragmentBinaryLinking */


// Indicates whether you can use the render pipeline to create new pipelines by adding binary functions to the fragment shader function’s callable functions list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/supportFragmentBinaryLinking
func (m_ MTL4MeshRenderPipelineDescriptor) SetSupportFragmentBinaryLinking(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSupportFragmentBinaryLinking:"), value)
}/* debug [instance_properties/setter]: supportFragmentBinaryLinking */


// Indicates whether the pipeline supports indirect command buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/supportIndirectCommandBuffers
func (m_ MTL4MeshRenderPipelineDescriptor) SupportIndirectCommandBuffers() MTL4IndirectCommandBufferSupportState {
	rv := objc.Send[MTL4IndirectCommandBufferSupportState](m_.ID, objc.Sel("supportIndirectCommandBuffers"))
	return rv
}/* debug [instance_properties/getter]: supportIndirectCommandBuffers */


// Indicates whether the pipeline supports indirect command buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/supportIndirectCommandBuffers
func (m_ MTL4MeshRenderPipelineDescriptor) SetSupportIndirectCommandBuffers(value MTL4IndirectCommandBufferSupportState) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSupportIndirectCommandBuffers:"), value)
}/* debug [instance_properties/setter]: supportIndirectCommandBuffers */


// Indicates whether you can use the render pipeline to create new pipelines by adding binary functions to the mesh shader function’s callable functions list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/supportMeshBinaryLinking
func (m_ MTL4MeshRenderPipelineDescriptor) SupportMeshBinaryLinking() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("supportMeshBinaryLinking"))
	return rv
}/* debug [instance_properties/getter]: supportMeshBinaryLinking */


// Indicates whether you can use the render pipeline to create new pipelines by adding binary functions to the mesh shader function’s callable functions list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/supportMeshBinaryLinking
func (m_ MTL4MeshRenderPipelineDescriptor) SetSupportMeshBinaryLinking(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSupportMeshBinaryLinking:"), value)
}/* debug [instance_properties/setter]: supportMeshBinaryLinking */


// Indicates whether you can use the render pipeline to create new pipelines by adding binary functions to the object shader function’s callable functions list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/supportObjectBinaryLinking
func (m_ MTL4MeshRenderPipelineDescriptor) SupportObjectBinaryLinking() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("supportObjectBinaryLinking"))
	return rv
}/* debug [instance_properties/getter]: supportObjectBinaryLinking */


// Indicates whether you can use the render pipeline to create new pipelines by adding binary functions to the object shader function’s callable functions list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/supportObjectBinaryLinking
func (m_ MTL4MeshRenderPipelineDescriptor) SetSupportObjectBinaryLinking(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSupportObjectBinaryLinking:"), value)
}/* debug [instance_properties/setter]: supportObjectBinaryLinking */


// Determines whether the pipeline rasterizes primitives.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4meshrenderpipelinedescriptor/israsterizationenabled
func (m_ MTL4MeshRenderPipelineDescriptor) IsRasterizationEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isRasterizationEnabled"))
	return rv
}/* debug [instance_properties/getter]: isRasterizationEnabled */


// Determines whether the pipeline rasterizes primitives.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4meshrenderpipelinedescriptor/israsterizationenabled
func (m_ MTL4MeshRenderPipelineDescriptor) SetIsRasterizationEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsRasterizationEnabled:"), value)
}/* debug [instance_properties/setter]: isRasterizationEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTL4MeshRenderPipelineDescriptor */



