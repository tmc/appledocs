// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MTL4RenderPipelineDescriptor */


/* debug [class_header]: Header for MTL4RenderPipelineDescriptor */
// The class instance for the [MTL4RenderPipelineDescriptor] class.
var (
	MTL4RenderPipelineDescriptorClass     _MTL4RenderPipelineDescriptorClass
	MTL4RenderPipelineDescriptorClassOnce sync.Once
)

func getMTL4RenderPipelineDescriptorClass() _MTL4RenderPipelineDescriptorClass {
	MTL4RenderPipelineDescriptorClassOnce.Do(func() {
		MTL4RenderPipelineDescriptorClass = _MTL4RenderPipelineDescriptorClass{objc.GetClass("MTL4RenderPipelineDescriptor")}
	})
	return MTL4RenderPipelineDescriptorClass
}

type _MTL4RenderPipelineDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTL4RenderPipelineDescriptor */
// An interface definition for the [MTL4RenderPipelineDescriptor] class.
type IMTL4RenderPipelineDescriptor interface {
	IMTL4PipelineDescriptor
	
/* debug [class_interface_properties]: Properties for MTL4RenderPipelineDescriptor */
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
	InputPrimitiveTopology() PrimitiveTopologyClass
	SetInputPrimitiveTopology(value PrimitiveTopologyClass)
	RasterizationEnabled() bool
	SetRasterizationEnabled(value bool)
	MaxVertexAmplificationCount() uint
	SetMaxVertexAmplificationCount(value uint)
	RasterSampleCount() uint
	SetRasterSampleCount(value uint)
	SupportFragmentBinaryLinking() bool
	SetSupportFragmentBinaryLinking(value bool)
	SupportIndirectCommandBuffers() MTL4IndirectCommandBufferSupportState
	SetSupportIndirectCommandBuffers(value MTL4IndirectCommandBufferSupportState)
	SupportVertexBinaryLinking() bool
	SetSupportVertexBinaryLinking(value bool)
	VertexDescriptor() IMTLVertexDescriptor
	SetVertexDescriptor(value IMTLVertexDescriptor)
	VertexFunctionDescriptor() IMTL4FunctionDescriptor
	SetVertexFunctionDescriptor(value IMTL4FunctionDescriptor)
	VertexStaticLinkingDescriptor() IMTL4StaticLinkingDescriptor
	SetVertexStaticLinkingDescriptor(value IMTL4StaticLinkingDescriptor)
	IsRasterizationEnabled() bool
	SetIsRasterizationEnabled(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTL4RenderPipelineDescriptor */
	// methods:
	Reset()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTL4RenderPipelineDescriptor */
// Alloc allocates a new instance without initialization.
func (mc _MTL4RenderPipelineDescriptorClass) Alloc() MTL4RenderPipelineDescriptor {
	rv := objc.Send[MTL4RenderPipelineDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTL4RenderPipelineDescriptorClass) New() MTL4RenderPipelineDescriptor {
	rv := objc.Send[MTL4RenderPipelineDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4RenderPipelineDescriptor) Init() MTL4RenderPipelineDescriptor {
	rv := objc.Send[MTL4RenderPipelineDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4RenderPipelineDescriptor) Autorelease() MTL4RenderPipelineDescriptor {
	rv := objc.Send[MTL4RenderPipelineDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4RenderPipelineDescriptor creates a new MTL4RenderPipelineDescriptor instance.
func NewMTL4RenderPipelineDescriptor() MTL4RenderPipelineDescriptor {
	return getMTL4RenderPipelineDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTL4RenderPipelineDescriptor */
// Groups together properties to create a render pipeline state object.
//
// Compared to , this interface doesn’t offer a mechanism to hint to Metal mutability of vertex and fragment buffers. Additionally, using this descriptor, you don’t specify binary archives.


// Groups together properties to create a render pipeline state object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor
type MTL4RenderPipelineDescriptor struct {
	MTL4PipelineDescriptor
}

// MTL4RenderPipelineDescriptorFrom constructs a [MTL4RenderPipelineDescriptor] from an unsafe.Pointer.
//
// Groups together properties to create a render pipeline state object.
func MTL4RenderPipelineDescriptorFrom(ptr unsafe.Pointer) MTL4RenderPipelineDescriptor {
	return MTL4RenderPipelineDescriptor{
		MTL4PipelineDescriptor: MTL4PipelineDescriptorFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTL4RenderPipelineDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTL4RenderPipelineDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTL4RenderPipelineDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTL4RenderPipelineDescriptor */

// Resets this descriptor to its default state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/reset()
func (m_ MTL4RenderPipelineDescriptor) Reset() {
	objc.Send[objc.ID](m_.ID, objc.Sel("reset"))
}/* debug [instance_methods/method]: Reset */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTL4RenderPipelineDescriptor */

// Indicates whether to read and use the alpha channel fragment output of color attachments to compute a sample coverage mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/alphaToCoverageState
func (m_ MTL4RenderPipelineDescriptor) AlphaToCoverageState() MTL4AlphaToCoverageState {
	rv := objc.Send[MTL4AlphaToCoverageState](m_.ID, objc.Sel("alphaToCoverageState"))
	return rv
}/* debug [instance_properties/getter]: alphaToCoverageState */


// Indicates whether to read and use the alpha channel fragment output of color attachments to compute a sample coverage mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/alphaToCoverageState
func (m_ MTL4RenderPipelineDescriptor) SetAlphaToCoverageState(value MTL4AlphaToCoverageState) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlphaToCoverageState:"), value)
}/* debug [instance_properties/setter]: alphaToCoverageState */


// Indicates whether the pipeline forces alpha channel values of color attachments to the largest representable value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/alphaToOneState
func (m_ MTL4RenderPipelineDescriptor) AlphaToOneState() MTL4AlphaToOneState {
	rv := objc.Send[MTL4AlphaToOneState](m_.ID, objc.Sel("alphaToOneState"))
	return rv
}/* debug [instance_properties/getter]: alphaToOneState */


// Indicates whether the pipeline forces alpha channel values of color attachments to the largest representable value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/alphaToOneState
func (m_ MTL4RenderPipelineDescriptor) SetAlphaToOneState(value MTL4AlphaToOneState) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlphaToOneState:"), value)
}/* debug [instance_properties/setter]: alphaToOneState */


// Configures a logical-to-physical rendering remap state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/colorAttachmentMappingState
func (m_ MTL4RenderPipelineDescriptor) ColorAttachmentMappingState() MTL4LogicalToPhysicalColorAttachmentMappingState {
	rv := objc.Send[MTL4LogicalToPhysicalColorAttachmentMappingState](m_.ID, objc.Sel("colorAttachmentMappingState"))
	return rv
}/* debug [instance_properties/getter]: colorAttachmentMappingState */


// Configures a logical-to-physical rendering remap state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/colorAttachmentMappingState
func (m_ MTL4RenderPipelineDescriptor) SetColorAttachmentMappingState(value MTL4LogicalToPhysicalColorAttachmentMappingState) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setColorAttachmentMappingState:"), value)
}/* debug [instance_properties/setter]: colorAttachmentMappingState */


// Accesses an array containing descriptions of the color attachments this pipeline writes to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/colorAttachments
func (m_ MTL4RenderPipelineDescriptor) ColorAttachments() IMTL4RenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Send[MTL4RenderPipelineColorAttachmentDescriptorArray](m_.ID, objc.Sel("colorAttachments"))
	return rv
}/* debug [instance_properties/getter]: colorAttachments */


// Assigns the shader function that this pipeline executes for each fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/fragmentFunctionDescriptor
func (m_ MTL4RenderPipelineDescriptor) FragmentFunctionDescriptor() IMTL4FunctionDescriptor {
	rv := objc.Send[MTL4FunctionDescriptor](m_.ID, objc.Sel("fragmentFunctionDescriptor"))
	return rv
}/* debug [instance_properties/getter]: fragmentFunctionDescriptor */


// Assigns the shader function that this pipeline executes for each fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/fragmentFunctionDescriptor
func (m_ MTL4RenderPipelineDescriptor) SetFragmentFunctionDescriptor(value IMTL4FunctionDescriptor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFragmentFunctionDescriptor:"), value)
}/* debug [instance_properties/setter]: fragmentFunctionDescriptor */


// Provides static linking information for the fragment stage of the render pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/fragmentStaticLinkingDescriptor
func (m_ MTL4RenderPipelineDescriptor) FragmentStaticLinkingDescriptor() IMTL4StaticLinkingDescriptor {
	rv := objc.Send[MTL4StaticLinkingDescriptor](m_.ID, objc.Sel("fragmentStaticLinkingDescriptor"))
	return rv
}/* debug [instance_properties/getter]: fragmentStaticLinkingDescriptor */


// Provides static linking information for the fragment stage of the render pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/fragmentStaticLinkingDescriptor
func (m_ MTL4RenderPipelineDescriptor) SetFragmentStaticLinkingDescriptor(value IMTL4StaticLinkingDescriptor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFragmentStaticLinkingDescriptor:"), value)
}/* debug [instance_properties/setter]: fragmentStaticLinkingDescriptor */


// Assigns type of primitive topology this pipeline renders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/inputPrimitiveTopology
func (m_ MTL4RenderPipelineDescriptor) InputPrimitiveTopology() PrimitiveTopologyClass {
	rv := objc.Send[PrimitiveTopologyClass](m_.ID, objc.Sel("inputPrimitiveTopology"))
	return rv
}/* debug [instance_properties/getter]: inputPrimitiveTopology */


// Assigns type of primitive topology this pipeline renders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/inputPrimitiveTopology
func (m_ MTL4RenderPipelineDescriptor) SetInputPrimitiveTopology(value PrimitiveTopologyClass) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInputPrimitiveTopology:"), value)
}/* debug [instance_properties/setter]: inputPrimitiveTopology */


// Determines whether the pipeline rasterizes primitives.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/isRasterizationEnabled
func (m_ MTL4RenderPipelineDescriptor) RasterizationEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("rasterizationEnabled"))
	return rv
}/* debug [instance_properties/getter]: rasterizationEnabled */


// Determines whether the pipeline rasterizes primitives.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/isRasterizationEnabled
func (m_ MTL4RenderPipelineDescriptor) SetRasterizationEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRasterizationEnabled:"), value)
}/* debug [instance_properties/setter]: rasterizationEnabled */


// Determines the maximum value that can you can pass as the pipeline’s amplification count.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/maxVertexAmplificationCount
func (m_ MTL4RenderPipelineDescriptor) MaxVertexAmplificationCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("maxVertexAmplificationCount"))
	return rv
}/* debug [instance_properties/getter]: maxVertexAmplificationCount */


// Determines the maximum value that can you can pass as the pipeline’s amplification count.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/maxVertexAmplificationCount
func (m_ MTL4RenderPipelineDescriptor) SetMaxVertexAmplificationCount(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxVertexAmplificationCount:"), value)
}/* debug [instance_properties/setter]: maxVertexAmplificationCount */


// Controls the number of samples this pipeline applies for each fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/rasterSampleCount
func (m_ MTL4RenderPipelineDescriptor) RasterSampleCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("rasterSampleCount"))
	return rv
}/* debug [instance_properties/getter]: rasterSampleCount */


// Controls the number of samples this pipeline applies for each fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/rasterSampleCount
func (m_ MTL4RenderPipelineDescriptor) SetRasterSampleCount(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRasterSampleCount:"), value)
}/* debug [instance_properties/setter]: rasterSampleCount */


// Indicates whether you can use the pipeline to create new pipelines by adding binary functions to the fragment shader function’s callable functions list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/supportFragmentBinaryLinking
func (m_ MTL4RenderPipelineDescriptor) SupportFragmentBinaryLinking() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("supportFragmentBinaryLinking"))
	return rv
}/* debug [instance_properties/getter]: supportFragmentBinaryLinking */


// Indicates whether you can use the pipeline to create new pipelines by adding binary functions to the fragment shader function’s callable functions list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/supportFragmentBinaryLinking
func (m_ MTL4RenderPipelineDescriptor) SetSupportFragmentBinaryLinking(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSupportFragmentBinaryLinking:"), value)
}/* debug [instance_properties/setter]: supportFragmentBinaryLinking */


// Indicates whether the pipeline supports indirect command buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/supportIndirectCommandBuffers
func (m_ MTL4RenderPipelineDescriptor) SupportIndirectCommandBuffers() MTL4IndirectCommandBufferSupportState {
	rv := objc.Send[MTL4IndirectCommandBufferSupportState](m_.ID, objc.Sel("supportIndirectCommandBuffers"))
	return rv
}/* debug [instance_properties/getter]: supportIndirectCommandBuffers */


// Indicates whether the pipeline supports indirect command buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/supportIndirectCommandBuffers
func (m_ MTL4RenderPipelineDescriptor) SetSupportIndirectCommandBuffers(value MTL4IndirectCommandBufferSupportState) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSupportIndirectCommandBuffers:"), value)
}/* debug [instance_properties/setter]: supportIndirectCommandBuffers */


// Indicates whether you can use the render pipeline to create new pipelines by adding binary functions to the vertex shader function’s callable functions list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/supportVertexBinaryLinking
func (m_ MTL4RenderPipelineDescriptor) SupportVertexBinaryLinking() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("supportVertexBinaryLinking"))
	return rv
}/* debug [instance_properties/getter]: supportVertexBinaryLinking */


// Indicates whether you can use the render pipeline to create new pipelines by adding binary functions to the vertex shader function’s callable functions list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/supportVertexBinaryLinking
func (m_ MTL4RenderPipelineDescriptor) SetSupportVertexBinaryLinking(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSupportVertexBinaryLinking:"), value)
}/* debug [instance_properties/setter]: supportVertexBinaryLinking */


// Configures an optional vertex descriptor for the vertex input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/vertexDescriptor
func (m_ MTL4RenderPipelineDescriptor) VertexDescriptor() IMTLVertexDescriptor {
	rv := objc.Send[VertexDescriptor](m_.ID, objc.Sel("vertexDescriptor"))
	return rv
}/* debug [instance_properties/getter]: vertexDescriptor */


// Configures an optional vertex descriptor for the vertex input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/vertexDescriptor
func (m_ MTL4RenderPipelineDescriptor) SetVertexDescriptor(value IMTLVertexDescriptor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVertexDescriptor:"), value)
}/* debug [instance_properties/setter]: vertexDescriptor */


// Assigns the shader function that this pipeline executes for each vertex.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/vertexFunctionDescriptor
func (m_ MTL4RenderPipelineDescriptor) VertexFunctionDescriptor() IMTL4FunctionDescriptor {
	rv := objc.Send[MTL4FunctionDescriptor](m_.ID, objc.Sel("vertexFunctionDescriptor"))
	return rv
}/* debug [instance_properties/getter]: vertexFunctionDescriptor */


// Assigns the shader function that this pipeline executes for each vertex.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/vertexFunctionDescriptor
func (m_ MTL4RenderPipelineDescriptor) SetVertexFunctionDescriptor(value IMTL4FunctionDescriptor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVertexFunctionDescriptor:"), value)
}/* debug [instance_properties/setter]: vertexFunctionDescriptor */


// Provides static linking information for the vertex stage of the render pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/vertexStaticLinkingDescriptor
func (m_ MTL4RenderPipelineDescriptor) VertexStaticLinkingDescriptor() IMTL4StaticLinkingDescriptor {
	rv := objc.Send[MTL4StaticLinkingDescriptor](m_.ID, objc.Sel("vertexStaticLinkingDescriptor"))
	return rv
}/* debug [instance_properties/getter]: vertexStaticLinkingDescriptor */


// Provides static linking information for the vertex stage of the render pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/vertexStaticLinkingDescriptor
func (m_ MTL4RenderPipelineDescriptor) SetVertexStaticLinkingDescriptor(value IMTL4StaticLinkingDescriptor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVertexStaticLinkingDescriptor:"), value)
}/* debug [instance_properties/setter]: vertexStaticLinkingDescriptor */


// Determines whether the pipeline rasterizes primitives.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4renderpipelinedescriptor/israsterizationenabled
func (m_ MTL4RenderPipelineDescriptor) IsRasterizationEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isRasterizationEnabled"))
	return rv
}/* debug [instance_properties/getter]: isRasterizationEnabled */


// Determines whether the pipeline rasterizes primitives.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4renderpipelinedescriptor/israsterizationenabled
func (m_ MTL4RenderPipelineDescriptor) SetIsRasterizationEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsRasterizationEnabled:"), value)
}/* debug [instance_properties/setter]: isRasterizationEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTL4RenderPipelineDescriptor */



