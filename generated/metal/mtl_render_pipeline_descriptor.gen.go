// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLRenderPipelineDescriptor */


/* debug [class_header]: Header for MTLRenderPipelineDescriptor */
// The class instance for the [RenderPipelineDescriptor] class.
var (
	RenderPipelineDescriptorClass     _RenderPipelineDescriptorClass
	RenderPipelineDescriptorClassOnce sync.Once
)

func getRenderPipelineDescriptorClass() _RenderPipelineDescriptorClass {
	RenderPipelineDescriptorClassOnce.Do(func() {
		RenderPipelineDescriptorClass = _RenderPipelineDescriptorClass{objc.GetClass("MTLRenderPipelineDescriptor")}
	})
	return RenderPipelineDescriptorClass
}

type _RenderPipelineDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RenderPipelineDescriptor */
// An interface definition for the [RenderPipelineDescriptor] class.
type IRenderPipelineDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for RenderPipelineDescriptor */
	// properties:
	BinaryArchives() []objc.ID
	SetBinaryArchives(value []objc.ID)
	ColorAttachments() IMTLRenderPipelineColorAttachmentDescriptorArray
	DepthAttachmentPixelFormat() PixelFormat
	SetDepthAttachmentPixelFormat(value PixelFormat)
	FragmentBuffers() IMTLPipelineBufferDescriptorArray
	FragmentFunction() unsafe.Pointer
	SetFragmentFunction(value unsafe.Pointer)
	FragmentLinkedFunctions() IMTLLinkedFunctions
	SetFragmentLinkedFunctions(value IMTLLinkedFunctions)
	FragmentPreloadedLibraries() []objc.ID
	SetFragmentPreloadedLibraries(value []objc.ID)
	InputPrimitiveTopology() PrimitiveTopologyClass
	SetInputPrimitiveTopology(value PrimitiveTopologyClass)
	AlphaToCoverageEnabled() bool
	SetAlphaToCoverageEnabled(value bool)
	AlphaToOneEnabled() bool
	SetAlphaToOneEnabled(value bool)
	RasterizationEnabled() bool
	SetRasterizationEnabled(value bool)
	TessellationFactorScaleEnabled() bool
	SetTessellationFactorScaleEnabled(value bool)
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	MaxFragmentCallStackDepth() uint
	SetMaxFragmentCallStackDepth(value uint)
	MaxTessellationFactor() uint
	SetMaxTessellationFactor(value uint)
	MaxVertexAmplificationCount() uint
	SetMaxVertexAmplificationCount(value uint)
	MaxVertexCallStackDepth() uint
	SetMaxVertexCallStackDepth(value uint)
	RasterSampleCount() uint
	SetRasterSampleCount(value uint)
	SampleCount() uint
	SetSampleCount(value uint)
	ShaderValidation() ShaderValidation
	SetShaderValidation(value ShaderValidation)
	StencilAttachmentPixelFormat() PixelFormat
	SetStencilAttachmentPixelFormat(value PixelFormat)
	SupportAddingFragmentBinaryFunctions() bool
	SetSupportAddingFragmentBinaryFunctions(value bool)
	SupportAddingVertexBinaryFunctions() bool
	SetSupportAddingVertexBinaryFunctions(value bool)
	SupportIndirectCommandBuffers() bool
	SetSupportIndirectCommandBuffers(value bool)
	TessellationControlPointIndexType() TessellationControlPointIndexType
	SetTessellationControlPointIndexType(value TessellationControlPointIndexType)
	TessellationFactorFormat() TessellationFactorFormat
	SetTessellationFactorFormat(value TessellationFactorFormat)
	TessellationFactorStepFunction() TessellationFactorStepFunction
	SetTessellationFactorStepFunction(value TessellationFactorStepFunction)
	TessellationOutputWindingOrder() Winding
	SetTessellationOutputWindingOrder(value Winding)
	TessellationPartitionMode() TessellationPartitionMode
	SetTessellationPartitionMode(value TessellationPartitionMode)
	VertexBuffers() IMTLPipelineBufferDescriptorArray
	VertexDescriptor() IMTLVertexDescriptor
	SetVertexDescriptor(value IMTLVertexDescriptor)
	VertexFunction() unsafe.Pointer
	SetVertexFunction(value unsafe.Pointer)
	VertexLinkedFunctions() IMTLLinkedFunctions
	SetVertexLinkedFunctions(value IMTLLinkedFunctions)
	VertexPreloadedLibraries() []objc.ID
	SetVertexPreloadedLibraries(value []objc.ID)
	IsAlphaToCoverageEnabled() bool
	SetIsAlphaToCoverageEnabled(value bool)
	IsAlphaToOneEnabled() bool
	SetIsAlphaToOneEnabled(value bool)
	IsRasterizationEnabled() bool
	SetIsRasterizationEnabled(value bool)
	IsTessellationFactorScaleEnabled() bool
	SetIsTessellationFactorScaleEnabled(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RenderPipelineDescriptor */
	// methods:
	Reset()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RenderPipelineDescriptor */
// Alloc allocates a new instance without initialization.
func (rc _RenderPipelineDescriptorClass) Alloc() RenderPipelineDescriptor {
	rv := objc.Send[RenderPipelineDescriptor](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RenderPipelineDescriptorClass) New() RenderPipelineDescriptor {
	rv := objc.Send[RenderPipelineDescriptor](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RenderPipelineDescriptor) Init() RenderPipelineDescriptor {
	rv := objc.Send[RenderPipelineDescriptor](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RenderPipelineDescriptor) Autorelease() RenderPipelineDescriptor {
	rv := objc.Send[RenderPipelineDescriptor](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRenderPipelineDescriptor creates a new RenderPipelineDescriptor instance.
func NewRenderPipelineDescriptor() RenderPipelineDescriptor {
	return getRenderPipelineDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RenderPipelineDescriptor */
// An argument of options you pass to a GPU device to get a render pipeline state.
//
// An instance configures the state of the pipeline to use during a rendering pass, including rasterization (such as multisampling), visibility, blending, tessellation, and graphics function state. Use standard allocation and initialization techniques to create an object. Then configure and use the descriptor to create an object. To specify the vertex or fragment function in the rendering pipeline descriptor, set the or property, respectively, to the desired object. The system ignores the tessellation stage properties if you don’t set the property to a post-tessellation vertex function. A vertex function is a post-tessellation vertex function if the attribute precedes the function’s signature in your Metal Shading Language source. See the “Post-Tessellation Vertex Functions” section of for more information. Setting the property to disables the rasterization of pixels into the color attachment. This action is typically for outputting vertex function data into a buffer object, or for depth-only rendering. If the vertex shader has an argument with per-vertex input attributes, set the property to an object that describes the organization of that vertex data.


// An argument of options you pass to a GPU device to get a render pipeline state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor
type RenderPipelineDescriptor struct {
	objectivec.Object
}

// RenderPipelineDescriptorFrom constructs a [RenderPipelineDescriptor] from an unsafe.Pointer.
//
// An argument of options you pass to a GPU device to get a render pipeline state.
func RenderPipelineDescriptorFrom(ptr unsafe.Pointer) RenderPipelineDescriptor {
	return RenderPipelineDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RenderPipelineDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RenderPipelineDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RenderPipelineDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RenderPipelineDescriptor */

// Specifies the default rendering pipeline state values for the descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/reset()
func (r_ RenderPipelineDescriptor) Reset() {
	objc.Send[objc.ID](r_.ID, objc.Sel("reset"))
}/* debug [instance_methods/method]: Reset */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RenderPipelineDescriptor */

// An array of binary archives to search for precompiled versions of the shader.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/binaryArchives
func (r_ RenderPipelineDescriptor) BinaryArchives() []objc.ID {
	rv := objc.Send[[]objc.ID](r_.ID, objc.Sel("binaryArchives"))
	return rv
}/* debug [instance_properties/getter]: binaryArchives */


// An array of binary archives to search for precompiled versions of the shader.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/binaryArchives
func (r_ RenderPipelineDescriptor) SetBinaryArchives(value []objc.ID) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](r_.ID, objc.Sel("setBinaryArchives:"), nsArray)
}/* debug [instance_properties/setter]: binaryArchives */


// An array of attachments that store color data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/colorAttachments
func (r_ RenderPipelineDescriptor) ColorAttachments() IMTLRenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Send[RenderPipelineColorAttachmentDescriptorArray](r_.ID, objc.Sel("colorAttachments"))
	return rv
}/* debug [instance_properties/getter]: colorAttachments */


// The pixel format of the attachment that stores depth data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/depthAttachmentPixelFormat
func (r_ RenderPipelineDescriptor) DepthAttachmentPixelFormat() PixelFormat {
	rv := objc.Send[PixelFormat](r_.ID, objc.Sel("depthAttachmentPixelFormat"))
	return rv
}/* debug [instance_properties/getter]: depthAttachmentPixelFormat */


// The pixel format of the attachment that stores depth data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/depthAttachmentPixelFormat
func (r_ RenderPipelineDescriptor) SetDepthAttachmentPixelFormat(value PixelFormat) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setDepthAttachmentPixelFormat:"), value)
}/* debug [instance_properties/setter]: depthAttachmentPixelFormat */


// An array that contains the buffer mutability options for a render pipeline’s fragment function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/fragmentBuffers
func (r_ RenderPipelineDescriptor) FragmentBuffers() IMTLPipelineBufferDescriptorArray {
	rv := objc.Send[PipelineBufferDescriptorArray](r_.ID, objc.Sel("fragmentBuffers"))
	return rv
}/* debug [instance_properties/getter]: fragmentBuffers */


// The fragment function the pipeline calls to process fragments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/fragmentFunction
func (r_ RenderPipelineDescriptor) FragmentFunction() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("fragmentFunction"))
	return rv
}/* debug [instance_properties/getter]: fragmentFunction */


// The fragment function the pipeline calls to process fragments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/fragmentFunction
func (r_ RenderPipelineDescriptor) SetFragmentFunction(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setFragmentFunction:"), value)
}/* debug [instance_properties/setter]: fragmentFunction */


// Functions that you can specify as function arguments for the fragment shader when encoding commands that use the pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/fragmentLinkedFunctions
func (r_ RenderPipelineDescriptor) FragmentLinkedFunctions() IMTLLinkedFunctions {
	rv := objc.Send[LinkedFunctions](r_.ID, objc.Sel("fragmentLinkedFunctions"))
	return rv
}/* debug [instance_properties/getter]: fragmentLinkedFunctions */


// Functions that you can specify as function arguments for the fragment shader when encoding commands that use the pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/fragmentLinkedFunctions
func (r_ RenderPipelineDescriptor) SetFragmentLinkedFunctions(value IMTLLinkedFunctions) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setFragmentLinkedFunctions:"), value)
}/* debug [instance_properties/setter]: fragmentLinkedFunctions */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/fragmentPreloadedLibraries
func (r_ RenderPipelineDescriptor) FragmentPreloadedLibraries() []objc.ID {
	rv := objc.Send[[]objc.ID](r_.ID, objc.Sel("fragmentPreloadedLibraries"))
	return rv
}/* debug [instance_properties/getter]: fragmentPreloadedLibraries */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/fragmentPreloadedLibraries
func (r_ RenderPipelineDescriptor) SetFragmentPreloadedLibraries(value []objc.ID) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](r_.ID, objc.Sel("setFragmentPreloadedLibraries:"), nsArray)
}/* debug [instance_properties/setter]: fragmentPreloadedLibraries */


// The type of primitive topology the pipeline renders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/inputPrimitiveTopology
func (r_ RenderPipelineDescriptor) InputPrimitiveTopology() PrimitiveTopologyClass {
	rv := objc.Send[PrimitiveTopologyClass](r_.ID, objc.Sel("inputPrimitiveTopology"))
	return rv
}/* debug [instance_properties/getter]: inputPrimitiveTopology */


// The type of primitive topology the pipeline renders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/inputPrimitiveTopology
func (r_ RenderPipelineDescriptor) SetInputPrimitiveTopology(value PrimitiveTopologyClass) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setInputPrimitiveTopology:"), value)
}/* debug [instance_properties/setter]: inputPrimitiveTopology */


// A Boolean value that indicates whether to read and use the alpha channel fragment output for color attachments to compute a sample coverage mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/isAlphaToCoverageEnabled
func (r_ RenderPipelineDescriptor) AlphaToCoverageEnabled() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("alphaToCoverageEnabled"))
	return rv
}/* debug [instance_properties/getter]: alphaToCoverageEnabled */


// A Boolean value that indicates whether to read and use the alpha channel fragment output for color attachments to compute a sample coverage mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/isAlphaToCoverageEnabled
func (r_ RenderPipelineDescriptor) SetAlphaToCoverageEnabled(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setAlphaToCoverageEnabled:"), value)
}/* debug [instance_properties/setter]: alphaToCoverageEnabled */


// A Boolean value that indicates whether to force alpha channel values for color attachments to the largest representable value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/isAlphaToOneEnabled
func (r_ RenderPipelineDescriptor) AlphaToOneEnabled() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("alphaToOneEnabled"))
	return rv
}/* debug [instance_properties/getter]: alphaToOneEnabled */


// A Boolean value that indicates whether to force alpha channel values for color attachments to the largest representable value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/isAlphaToOneEnabled
func (r_ RenderPipelineDescriptor) SetAlphaToOneEnabled(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setAlphaToOneEnabled:"), value)
}/* debug [instance_properties/setter]: alphaToOneEnabled */


// A Boolean value that determines whether the pipeline rasterizes primitives.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/isRasterizationEnabled
func (r_ RenderPipelineDescriptor) RasterizationEnabled() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("rasterizationEnabled"))
	return rv
}/* debug [instance_properties/getter]: rasterizationEnabled */


// A Boolean value that determines whether the pipeline rasterizes primitives.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/isRasterizationEnabled
func (r_ RenderPipelineDescriptor) SetRasterizationEnabled(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRasterizationEnabled:"), value)
}/* debug [instance_properties/setter]: rasterizationEnabled */


// A Boolean value that determines whether the pipeline scales the tessellation factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/isTessellationFactorScaleEnabled
func (r_ RenderPipelineDescriptor) TessellationFactorScaleEnabled() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("tessellationFactorScaleEnabled"))
	return rv
}/* debug [instance_properties/getter]: tessellationFactorScaleEnabled */


// A Boolean value that determines whether the pipeline scales the tessellation factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/isTessellationFactorScaleEnabled
func (r_ RenderPipelineDescriptor) SetTessellationFactorScaleEnabled(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setTessellationFactorScaleEnabled:"), value)
}/* debug [instance_properties/setter]: tessellationFactorScaleEnabled */


// A string that identifies the render pipeline descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/label
func (r_ RenderPipelineDescriptor) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](r_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// A string that identifies the render pipeline descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/label
func (r_ RenderPipelineDescriptor) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setLabel:"), value)
}/* debug [instance_properties/setter]: label */


// The maximum function call depth from the top-most fragment shader function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/maxFragmentCallStackDepth
func (r_ RenderPipelineDescriptor) MaxFragmentCallStackDepth() uint {
	rv := objc.Send[uint](r_.ID, objc.Sel("maxFragmentCallStackDepth"))
	return rv
}/* debug [instance_properties/getter]: maxFragmentCallStackDepth */


// The maximum function call depth from the top-most fragment shader function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/maxFragmentCallStackDepth
func (r_ RenderPipelineDescriptor) SetMaxFragmentCallStackDepth(value uint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setMaxFragmentCallStackDepth:"), value)
}/* debug [instance_properties/setter]: maxFragmentCallStackDepth */


// The maximum tessellation factor that the tessellator uses when tessellating patches.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/maxTessellationFactor
func (r_ RenderPipelineDescriptor) MaxTessellationFactor() uint {
	rv := objc.Send[uint](r_.ID, objc.Sel("maxTessellationFactor"))
	return rv
}/* debug [instance_properties/getter]: maxTessellationFactor */


// The maximum tessellation factor that the tessellator uses when tessellating patches.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/maxTessellationFactor
func (r_ RenderPipelineDescriptor) SetMaxTessellationFactor(value uint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setMaxTessellationFactor:"), value)
}/* debug [instance_properties/setter]: maxTessellationFactor */


// The maximum vertex amplification count you can set when encoding render commands.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/maxVertexAmplificationCount
func (r_ RenderPipelineDescriptor) MaxVertexAmplificationCount() uint {
	rv := objc.Send[uint](r_.ID, objc.Sel("maxVertexAmplificationCount"))
	return rv
}/* debug [instance_properties/getter]: maxVertexAmplificationCount */


// The maximum vertex amplification count you can set when encoding render commands.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/maxVertexAmplificationCount
func (r_ RenderPipelineDescriptor) SetMaxVertexAmplificationCount(value uint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setMaxVertexAmplificationCount:"), value)
}/* debug [instance_properties/setter]: maxVertexAmplificationCount */


// The maximum function call depth from the top-most vertex shader function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/maxVertexCallStackDepth
func (r_ RenderPipelineDescriptor) MaxVertexCallStackDepth() uint {
	rv := objc.Send[uint](r_.ID, objc.Sel("maxVertexCallStackDepth"))
	return rv
}/* debug [instance_properties/getter]: maxVertexCallStackDepth */


// The maximum function call depth from the top-most vertex shader function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/maxVertexCallStackDepth
func (r_ RenderPipelineDescriptor) SetMaxVertexCallStackDepth(value uint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setMaxVertexCallStackDepth:"), value)
}/* debug [instance_properties/setter]: maxVertexCallStackDepth */


// The number of samples the pipeline applies for each fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/rasterSampleCount
func (r_ RenderPipelineDescriptor) RasterSampleCount() uint {
	rv := objc.Send[uint](r_.ID, objc.Sel("rasterSampleCount"))
	return rv
}/* debug [instance_properties/getter]: rasterSampleCount */


// The number of samples the pipeline applies for each fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/rasterSampleCount
func (r_ RenderPipelineDescriptor) SetRasterSampleCount(value uint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRasterSampleCount:"), value)
}/* debug [instance_properties/setter]: rasterSampleCount */


// The number of samples the pipeline applies for each fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/sampleCount
func (r_ RenderPipelineDescriptor) SampleCount() uint {
	rv := objc.Send[uint](r_.ID, objc.Sel("sampleCount"))
	return rv
}/* debug [instance_properties/getter]: sampleCount */


// The number of samples the pipeline applies for each fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/sampleCount
func (r_ RenderPipelineDescriptor) SetSampleCount(value uint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setSampleCount:"), value)
}/* debug [instance_properties/setter]: sampleCount */


// A value that enables or disables shader validation for the pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/shaderValidation
func (r_ RenderPipelineDescriptor) ShaderValidation() ShaderValidation {
	rv := objc.Send[ShaderValidation](r_.ID, objc.Sel("shaderValidation"))
	return rv
}/* debug [instance_properties/getter]: shaderValidation */


// A value that enables or disables shader validation for the pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/shaderValidation
func (r_ RenderPipelineDescriptor) SetShaderValidation(value ShaderValidation) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setShaderValidation:"), value)
}/* debug [instance_properties/setter]: shaderValidation */


// The pixel format of the attachment that stores stencil data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/stencilAttachmentPixelFormat
func (r_ RenderPipelineDescriptor) StencilAttachmentPixelFormat() PixelFormat {
	rv := objc.Send[PixelFormat](r_.ID, objc.Sel("stencilAttachmentPixelFormat"))
	return rv
}/* debug [instance_properties/getter]: stencilAttachmentPixelFormat */


// The pixel format of the attachment that stores stencil data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/stencilAttachmentPixelFormat
func (r_ RenderPipelineDescriptor) SetStencilAttachmentPixelFormat(value PixelFormat) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setStencilAttachmentPixelFormat:"), value)
}/* debug [instance_properties/setter]: stencilAttachmentPixelFormat */


// A Boolean value that indicates whether you can use the pipeline to create new pipelines by adding binary functions to the fragment shader’s callable functions list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/supportAddingFragmentBinaryFunctions
func (r_ RenderPipelineDescriptor) SupportAddingFragmentBinaryFunctions() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("supportAddingFragmentBinaryFunctions"))
	return rv
}/* debug [instance_properties/getter]: supportAddingFragmentBinaryFunctions */


// A Boolean value that indicates whether you can use the pipeline to create new pipelines by adding binary functions to the fragment shader’s callable functions list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/supportAddingFragmentBinaryFunctions
func (r_ RenderPipelineDescriptor) SetSupportAddingFragmentBinaryFunctions(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setSupportAddingFragmentBinaryFunctions:"), value)
}/* debug [instance_properties/setter]: supportAddingFragmentBinaryFunctions */


// A Boolean value that indicates whether you can use the pipeline to create new pipelines by adding binary functions to the vertex shader’s callable functions list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/supportAddingVertexBinaryFunctions
func (r_ RenderPipelineDescriptor) SupportAddingVertexBinaryFunctions() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("supportAddingVertexBinaryFunctions"))
	return rv
}/* debug [instance_properties/getter]: supportAddingVertexBinaryFunctions */


// A Boolean value that indicates whether you can use the pipeline to create new pipelines by adding binary functions to the vertex shader’s callable functions list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/supportAddingVertexBinaryFunctions
func (r_ RenderPipelineDescriptor) SetSupportAddingVertexBinaryFunctions(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setSupportAddingVertexBinaryFunctions:"), value)
}/* debug [instance_properties/setter]: supportAddingVertexBinaryFunctions */


// A Boolean value that determines whether you can encode commands into an indirect command buffer using the render pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/supportIndirectCommandBuffers
func (r_ RenderPipelineDescriptor) SupportIndirectCommandBuffers() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("supportIndirectCommandBuffers"))
	return rv
}/* debug [instance_properties/getter]: supportIndirectCommandBuffers */


// A Boolean value that determines whether you can encode commands into an indirect command buffer using the render pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/supportIndirectCommandBuffers
func (r_ RenderPipelineDescriptor) SetSupportIndirectCommandBuffers(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setSupportIndirectCommandBuffers:"), value)
}/* debug [instance_properties/setter]: supportIndirectCommandBuffers */


// The size of the control point indices in a control point index buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/tessellationControlPointIndexType
func (r_ RenderPipelineDescriptor) TessellationControlPointIndexType() TessellationControlPointIndexType {
	rv := objc.Send[TessellationControlPointIndexType](r_.ID, objc.Sel("tessellationControlPointIndexType"))
	return rv
}/* debug [instance_properties/getter]: tessellationControlPointIndexType */


// The size of the control point indices in a control point index buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/tessellationControlPointIndexType
func (r_ RenderPipelineDescriptor) SetTessellationControlPointIndexType(value TessellationControlPointIndexType) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setTessellationControlPointIndexType:"), value)
}/* debug [instance_properties/setter]: tessellationControlPointIndexType */


// The format of the tessellation factors in the tessellation factor buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/tessellationFactorFormat
func (r_ RenderPipelineDescriptor) TessellationFactorFormat() TessellationFactorFormat {
	rv := objc.Send[TessellationFactorFormat](r_.ID, objc.Sel("tessellationFactorFormat"))
	return rv
}/* debug [instance_properties/getter]: tessellationFactorFormat */


// The format of the tessellation factors in the tessellation factor buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/tessellationFactorFormat
func (r_ RenderPipelineDescriptor) SetTessellationFactorFormat(value TessellationFactorFormat) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setTessellationFactorFormat:"), value)
}/* debug [instance_properties/setter]: tessellationFactorFormat */


// The step function for determining the tessellation factors for a patch from the tessellation factor buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/tessellationFactorStepFunction
func (r_ RenderPipelineDescriptor) TessellationFactorStepFunction() TessellationFactorStepFunction {
	rv := objc.Send[TessellationFactorStepFunction](r_.ID, objc.Sel("tessellationFactorStepFunction"))
	return rv
}/* debug [instance_properties/getter]: tessellationFactorStepFunction */


// The step function for determining the tessellation factors for a patch from the tessellation factor buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/tessellationFactorStepFunction
func (r_ RenderPipelineDescriptor) SetTessellationFactorStepFunction(value TessellationFactorStepFunction) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setTessellationFactorStepFunction:"), value)
}/* debug [instance_properties/setter]: tessellationFactorStepFunction */


// The winding order of triangles from the tessellator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/tessellationOutputWindingOrder
func (r_ RenderPipelineDescriptor) TessellationOutputWindingOrder() Winding {
	rv := objc.Send[Winding](r_.ID, objc.Sel("tessellationOutputWindingOrder"))
	return rv
}/* debug [instance_properties/getter]: tessellationOutputWindingOrder */


// The winding order of triangles from the tessellator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/tessellationOutputWindingOrder
func (r_ RenderPipelineDescriptor) SetTessellationOutputWindingOrder(value Winding) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setTessellationOutputWindingOrder:"), value)
}/* debug [instance_properties/setter]: tessellationOutputWindingOrder */


// The partitioning mode that the tessellator uses to derive the number and spacing of segments for subdividing a corresponding edge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/tessellationPartitionMode
func (r_ RenderPipelineDescriptor) TessellationPartitionMode() TessellationPartitionMode {
	rv := objc.Send[TessellationPartitionMode](r_.ID, objc.Sel("tessellationPartitionMode"))
	return rv
}/* debug [instance_properties/getter]: tessellationPartitionMode */


// The partitioning mode that the tessellator uses to derive the number and spacing of segments for subdividing a corresponding edge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/tessellationPartitionMode
func (r_ RenderPipelineDescriptor) SetTessellationPartitionMode(value TessellationPartitionMode) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setTessellationPartitionMode:"), value)
}/* debug [instance_properties/setter]: tessellationPartitionMode */


// An array that contains the buffer mutability options for a render pipeline’s vertex function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/vertexBuffers
func (r_ RenderPipelineDescriptor) VertexBuffers() IMTLPipelineBufferDescriptorArray {
	rv := objc.Send[PipelineBufferDescriptorArray](r_.ID, objc.Sel("vertexBuffers"))
	return rv
}/* debug [instance_properties/getter]: vertexBuffers */


// The organization of vertex data in an attribute’s argument table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/vertexDescriptor
func (r_ RenderPipelineDescriptor) VertexDescriptor() IMTLVertexDescriptor {
	rv := objc.Send[VertexDescriptor](r_.ID, objc.Sel("vertexDescriptor"))
	return rv
}/* debug [instance_properties/getter]: vertexDescriptor */


// The organization of vertex data in an attribute’s argument table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/vertexDescriptor
func (r_ RenderPipelineDescriptor) SetVertexDescriptor(value IMTLVertexDescriptor) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setVertexDescriptor:"), value)
}/* debug [instance_properties/setter]: vertexDescriptor */


// The vertex function the pipeline calls to process vertices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/vertexFunction
func (r_ RenderPipelineDescriptor) VertexFunction() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("vertexFunction"))
	return rv
}/* debug [instance_properties/getter]: vertexFunction */


// The vertex function the pipeline calls to process vertices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/vertexFunction
func (r_ RenderPipelineDescriptor) SetVertexFunction(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setVertexFunction:"), value)
}/* debug [instance_properties/setter]: vertexFunction */


// Functions that you can specify as function arguments for the vertex shader when encoding commands that use the pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/vertexLinkedFunctions
func (r_ RenderPipelineDescriptor) VertexLinkedFunctions() IMTLLinkedFunctions {
	rv := objc.Send[LinkedFunctions](r_.ID, objc.Sel("vertexLinkedFunctions"))
	return rv
}/* debug [instance_properties/getter]: vertexLinkedFunctions */


// Functions that you can specify as function arguments for the vertex shader when encoding commands that use the pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/vertexLinkedFunctions
func (r_ RenderPipelineDescriptor) SetVertexLinkedFunctions(value IMTLLinkedFunctions) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setVertexLinkedFunctions:"), value)
}/* debug [instance_properties/setter]: vertexLinkedFunctions */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/vertexPreloadedLibraries
func (r_ RenderPipelineDescriptor) VertexPreloadedLibraries() []objc.ID {
	rv := objc.Send[[]objc.ID](r_.ID, objc.Sel("vertexPreloadedLibraries"))
	return rv
}/* debug [instance_properties/getter]: vertexPreloadedLibraries */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/vertexPreloadedLibraries
func (r_ RenderPipelineDescriptor) SetVertexPreloadedLibraries(value []objc.ID) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](r_.ID, objc.Sel("setVertexPreloadedLibraries:"), nsArray)
}/* debug [instance_properties/setter]: vertexPreloadedLibraries */


// A Boolean value that indicates whether to read and use the alpha channel fragment output for color attachments to compute a sample coverage mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/isalphatocoverageenabled
func (r_ RenderPipelineDescriptor) IsAlphaToCoverageEnabled() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isAlphaToCoverageEnabled"))
	return rv
}/* debug [instance_properties/getter]: isAlphaToCoverageEnabled */


// A Boolean value that indicates whether to read and use the alpha channel fragment output for color attachments to compute a sample coverage mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/isalphatocoverageenabled
func (r_ RenderPipelineDescriptor) SetIsAlphaToCoverageEnabled(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsAlphaToCoverageEnabled:"), value)
}/* debug [instance_properties/setter]: isAlphaToCoverageEnabled */


// A Boolean value that indicates whether to force alpha channel values for color attachments to the largest representable value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/isalphatooneenabled
func (r_ RenderPipelineDescriptor) IsAlphaToOneEnabled() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isAlphaToOneEnabled"))
	return rv
}/* debug [instance_properties/getter]: isAlphaToOneEnabled */


// A Boolean value that indicates whether to force alpha channel values for color attachments to the largest representable value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/isalphatooneenabled
func (r_ RenderPipelineDescriptor) SetIsAlphaToOneEnabled(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsAlphaToOneEnabled:"), value)
}/* debug [instance_properties/setter]: isAlphaToOneEnabled */


// A Boolean value that determines whether the pipeline rasterizes primitives.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/israsterizationenabled
func (r_ RenderPipelineDescriptor) IsRasterizationEnabled() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isRasterizationEnabled"))
	return rv
}/* debug [instance_properties/getter]: isRasterizationEnabled */


// A Boolean value that determines whether the pipeline rasterizes primitives.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/israsterizationenabled
func (r_ RenderPipelineDescriptor) SetIsRasterizationEnabled(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsRasterizationEnabled:"), value)
}/* debug [instance_properties/setter]: isRasterizationEnabled */


// A Boolean value that determines whether the pipeline scales the tessellation factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/istessellationfactorscaleenabled
func (r_ RenderPipelineDescriptor) IsTessellationFactorScaleEnabled() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isTessellationFactorScaleEnabled"))
	return rv
}/* debug [instance_properties/getter]: isTessellationFactorScaleEnabled */


// A Boolean value that determines whether the pipeline scales the tessellation factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderpipelinedescriptor/istessellationfactorscaleenabled
func (r_ RenderPipelineDescriptor) SetIsTessellationFactorScaleEnabled(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsTessellationFactorScaleEnabled:"), value)
}/* debug [instance_properties/setter]: isTessellationFactorScaleEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLRenderPipelineDescriptor */



