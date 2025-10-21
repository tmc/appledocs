// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [RenderPipelineDescriptor] class.
type IRenderPipelineDescriptor interface {
	objectivec.IObject
	Reset()
}

// An argument of options you pass to a GPU device to get a render pipeline state.
//
// An instance configures the state of the pipeline to use during a rendering pass, including rasterization (such as multisampling), visibility, blending, tessellation, and graphics function state. Use standard allocation and initialization techniques to create an object. Then configure and use the descriptor to create an object. To specify the vertex or fragment function in the rendering pipeline descriptor, set the or property, respectively, to the desired object. The system ignores the tessellation stage properties if you don’t set the property to a post-tessellation vertex function. A vertex function is a post-tessellation vertex function if the attribute precedes the function’s signature in your Metal Shading Language source. See the “Post-Tessellation Vertex Functions” section of for more information. Setting the property to disables the rasterization of pixels into the color attachment. This action is typically for outputting vertex function data into a buffer object, or for depth-only rendering. If the vertex shader has an argument with per-vertex input attributes, set the property to an object that describes the organization of that vertex data.
//
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

// Alloc allocates a new instance without initialization.
func (rc _RenderPipelineDescriptorClass) Alloc() RenderPipelineDescriptor {
	rv := objc.Send[RenderPipelineDescriptor](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Specifies the default rendering pipeline state values for the descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/reset()
func (r_ RenderPipelineDescriptor) Reset() {
	objc.Send[objc.ID](r_.ID, objc.Sel("reset"))
}

// An array of binary archives to search for precompiled versions of the shader.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/binaryArchives
func (r_ RenderPipelineDescriptor) BinaryArchives() []objc.ID {
	rv := objc.Send[[]objc.ID](r_.ID, objc.Sel("binaryArchives"))
	return rv
}


// SetBinaryArchives sets the value of the binaryArchives property.
// An array of binary archives to search for precompiled versions of the shader.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/binaryArchives
func (r_ RenderPipelineDescriptor) SetBinaryArchives(value []objc.ID) {
	// Convert Go slice to NSArray
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
}
// An array of attachments that store color data.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/colorAttachments
func (r_ RenderPipelineDescriptor) ColorAttachments() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("colorAttachments"))
	return rv
}

// The pixel format of the attachment that stores depth data.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/depthAttachmentPixelFormat
func (r_ RenderPipelineDescriptor) DepthAttachmentPixelFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("depthAttachmentPixelFormat"))
	return rv
}


// SetDepthAttachmentPixelFormat sets the value of the depthAttachmentPixelFormat property.
// The pixel format of the attachment that stores depth data.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/depthAttachmentPixelFormat
func (r_ RenderPipelineDescriptor) SetDepthAttachmentPixelFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setDepthAttachmentPixelFormat:"), value)
}
// An array that contains the buffer mutability options for a render pipeline’s fragment function.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/fragmentBuffers
func (r_ RenderPipelineDescriptor) FragmentBuffers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("fragmentBuffers"))
	return rv
}

// The fragment function the pipeline calls to process fragments.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/fragmentFunction
func (r_ RenderPipelineDescriptor) FragmentFunction() objc.ID {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("fragmentFunction"))
	return rv
}


// SetFragmentFunction sets the value of the fragmentFunction property.
// The fragment function the pipeline calls to process fragments.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/fragmentFunction
func (r_ RenderPipelineDescriptor) SetFragmentFunction(value objc.ID) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setFragmentFunction:"), value)
}
// Functions that you can specify as function arguments for the fragment shader when encoding commands that use the pipeline.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/fragmentLinkedFunctions
func (r_ RenderPipelineDescriptor) FragmentLinkedFunctions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("fragmentLinkedFunctions"))
	return rv
}


// SetFragmentLinkedFunctions sets the value of the fragmentLinkedFunctions property.
// Functions that you can specify as function arguments for the fragment shader when encoding commands that use the pipeline.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/fragmentLinkedFunctions
func (r_ RenderPipelineDescriptor) SetFragmentLinkedFunctions(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setFragmentLinkedFunctions:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/fragmentPreloadedLibraries
func (r_ RenderPipelineDescriptor) FragmentPreloadedLibraries() []objc.ID {
	rv := objc.Send[[]objc.ID](r_.ID, objc.Sel("fragmentPreloadedLibraries"))
	return rv
}


// SetFragmentPreloadedLibraries sets the value of the fragmentPreloadedLibraries property.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/fragmentPreloadedLibraries
func (r_ RenderPipelineDescriptor) SetFragmentPreloadedLibraries(value []objc.ID) {
	// Convert Go slice to NSArray
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
}
// The type of primitive topology the pipeline renders.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/inputPrimitiveTopology
func (r_ RenderPipelineDescriptor) InputPrimitiveTopology() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("inputPrimitiveTopology"))
	return rv
}


// SetInputPrimitiveTopology sets the value of the inputPrimitiveTopology property.
// The type of primitive topology the pipeline renders.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/inputPrimitiveTopology
func (r_ RenderPipelineDescriptor) SetInputPrimitiveTopology(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setInputPrimitiveTopology:"), value)
}
// A Boolean value that indicates whether to read and use the alpha channel fragment output for color attachments to compute a sample coverage mask.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/isAlphaToCoverageEnabled
func (r_ RenderPipelineDescriptor) AlphaToCoverageEnabled() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("alphaToCoverageEnabled"))
	return rv
}


// SetAlphaToCoverageEnabled sets the value of the alphaToCoverageEnabled property.
// A Boolean value that indicates whether to read and use the alpha channel fragment output for color attachments to compute a sample coverage mask.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/isAlphaToCoverageEnabled
func (r_ RenderPipelineDescriptor) SetAlphaToCoverageEnabled(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setAlphaToCoverageEnabled:"), value)
}
// A Boolean value that indicates whether to force alpha channel values for color attachments to the largest representable value.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/isAlphaToOneEnabled
func (r_ RenderPipelineDescriptor) AlphaToOneEnabled() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("alphaToOneEnabled"))
	return rv
}


// SetAlphaToOneEnabled sets the value of the alphaToOneEnabled property.
// A Boolean value that indicates whether to force alpha channel values for color attachments to the largest representable value.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/isAlphaToOneEnabled
func (r_ RenderPipelineDescriptor) SetAlphaToOneEnabled(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setAlphaToOneEnabled:"), value)
}
// A Boolean value that determines whether the pipeline rasterizes primitives.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/isRasterizationEnabled
func (r_ RenderPipelineDescriptor) RasterizationEnabled() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("rasterizationEnabled"))
	return rv
}


// SetRasterizationEnabled sets the value of the rasterizationEnabled property.
// A Boolean value that determines whether the pipeline rasterizes primitives.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/isRasterizationEnabled
func (r_ RenderPipelineDescriptor) SetRasterizationEnabled(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRasterizationEnabled:"), value)
}
// A Boolean value that determines whether the pipeline scales the tessellation factor.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/isTessellationFactorScaleEnabled
func (r_ RenderPipelineDescriptor) TessellationFactorScaleEnabled() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("tessellationFactorScaleEnabled"))
	return rv
}


// SetTessellationFactorScaleEnabled sets the value of the tessellationFactorScaleEnabled property.
// A Boolean value that determines whether the pipeline scales the tessellation factor.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/isTessellationFactorScaleEnabled
func (r_ RenderPipelineDescriptor) SetTessellationFactorScaleEnabled(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setTessellationFactorScaleEnabled:"), value)
}
// A string that identifies the render pipeline descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/label
func (r_ RenderPipelineDescriptor) Label() string {
	rv := objc.Send[string](r_.ID, objc.Sel("label"))
	return rv
}


// SetLabel sets the value of the label property.
// A string that identifies the render pipeline descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/label
func (r_ RenderPipelineDescriptor) SetLabel(value string) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setLabel:"), objc.String(value))
}
// The maximum function call depth from the top-most fragment shader function.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/maxFragmentCallStackDepth
func (r_ RenderPipelineDescriptor) MaxFragmentCallStackDepth() uint {
	rv := objc.Send[uint](r_.ID, objc.Sel("maxFragmentCallStackDepth"))
	return rv
}


// SetMaxFragmentCallStackDepth sets the value of the maxFragmentCallStackDepth property.
// The maximum function call depth from the top-most fragment shader function.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/maxFragmentCallStackDepth
func (r_ RenderPipelineDescriptor) SetMaxFragmentCallStackDepth(value uint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setMaxFragmentCallStackDepth:"), value)
}
// The maximum tessellation factor that the tessellator uses when tessellating patches.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/maxTessellationFactor
func (r_ RenderPipelineDescriptor) MaxTessellationFactor() uint {
	rv := objc.Send[uint](r_.ID, objc.Sel("maxTessellationFactor"))
	return rv
}


// SetMaxTessellationFactor sets the value of the maxTessellationFactor property.
// The maximum tessellation factor that the tessellator uses when tessellating patches.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/maxTessellationFactor
func (r_ RenderPipelineDescriptor) SetMaxTessellationFactor(value uint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setMaxTessellationFactor:"), value)
}
// The maximum vertex amplification count you can set when encoding render commands.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/maxVertexAmplificationCount
func (r_ RenderPipelineDescriptor) MaxVertexAmplificationCount() uint {
	rv := objc.Send[uint](r_.ID, objc.Sel("maxVertexAmplificationCount"))
	return rv
}


// SetMaxVertexAmplificationCount sets the value of the maxVertexAmplificationCount property.
// The maximum vertex amplification count you can set when encoding render commands.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/maxVertexAmplificationCount
func (r_ RenderPipelineDescriptor) SetMaxVertexAmplificationCount(value uint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setMaxVertexAmplificationCount:"), value)
}
// The maximum function call depth from the top-most vertex shader function.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/maxVertexCallStackDepth
func (r_ RenderPipelineDescriptor) MaxVertexCallStackDepth() uint {
	rv := objc.Send[uint](r_.ID, objc.Sel("maxVertexCallStackDepth"))
	return rv
}


// SetMaxVertexCallStackDepth sets the value of the maxVertexCallStackDepth property.
// The maximum function call depth from the top-most vertex shader function.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/maxVertexCallStackDepth
func (r_ RenderPipelineDescriptor) SetMaxVertexCallStackDepth(value uint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setMaxVertexCallStackDepth:"), value)
}
// The number of samples the pipeline applies for each fragment.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/rasterSampleCount
func (r_ RenderPipelineDescriptor) RasterSampleCount() uint {
	rv := objc.Send[uint](r_.ID, objc.Sel("rasterSampleCount"))
	return rv
}


// SetRasterSampleCount sets the value of the rasterSampleCount property.
// The number of samples the pipeline applies for each fragment.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/rasterSampleCount
func (r_ RenderPipelineDescriptor) SetRasterSampleCount(value uint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRasterSampleCount:"), value)
}
// The number of samples the pipeline applies for each fragment.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/sampleCount
func (r_ RenderPipelineDescriptor) SampleCount() uint {
	rv := objc.Send[uint](r_.ID, objc.Sel("sampleCount"))
	return rv
}


// SetSampleCount sets the value of the sampleCount property.
// The number of samples the pipeline applies for each fragment.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/sampleCount
func (r_ RenderPipelineDescriptor) SetSampleCount(value uint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setSampleCount:"), value)
}
// A value that enables or disables shader validation for the pipeline.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/shaderValidation
func (r_ RenderPipelineDescriptor) ShaderValidation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("shaderValidation"))
	return rv
}


// SetShaderValidation sets the value of the shaderValidation property.
// A value that enables or disables shader validation for the pipeline.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/shaderValidation
func (r_ RenderPipelineDescriptor) SetShaderValidation(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setShaderValidation:"), value)
}
// The pixel format of the attachment that stores stencil data.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/stencilAttachmentPixelFormat
func (r_ RenderPipelineDescriptor) StencilAttachmentPixelFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("stencilAttachmentPixelFormat"))
	return rv
}


// SetStencilAttachmentPixelFormat sets the value of the stencilAttachmentPixelFormat property.
// The pixel format of the attachment that stores stencil data.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/stencilAttachmentPixelFormat
func (r_ RenderPipelineDescriptor) SetStencilAttachmentPixelFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setStencilAttachmentPixelFormat:"), value)
}
// A Boolean value that indicates whether you can use the pipeline to create new pipelines by adding binary functions to the fragment shader’s callable functions list.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/supportAddingFragmentBinaryFunctions
func (r_ RenderPipelineDescriptor) SupportAddingFragmentBinaryFunctions() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("supportAddingFragmentBinaryFunctions"))
	return rv
}


// SetSupportAddingFragmentBinaryFunctions sets the value of the supportAddingFragmentBinaryFunctions property.
// A Boolean value that indicates whether you can use the pipeline to create new pipelines by adding binary functions to the fragment shader’s callable functions list.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/supportAddingFragmentBinaryFunctions
func (r_ RenderPipelineDescriptor) SetSupportAddingFragmentBinaryFunctions(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setSupportAddingFragmentBinaryFunctions:"), value)
}
// A Boolean value that indicates whether you can use the pipeline to create new pipelines by adding binary functions to the vertex shader’s callable functions list.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/supportAddingVertexBinaryFunctions
func (r_ RenderPipelineDescriptor) SupportAddingVertexBinaryFunctions() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("supportAddingVertexBinaryFunctions"))
	return rv
}


// SetSupportAddingVertexBinaryFunctions sets the value of the supportAddingVertexBinaryFunctions property.
// A Boolean value that indicates whether you can use the pipeline to create new pipelines by adding binary functions to the vertex shader’s callable functions list.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/supportAddingVertexBinaryFunctions
func (r_ RenderPipelineDescriptor) SetSupportAddingVertexBinaryFunctions(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setSupportAddingVertexBinaryFunctions:"), value)
}
// A Boolean value that determines whether you can encode commands into an indirect command buffer using the render pipeline.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/supportIndirectCommandBuffers
func (r_ RenderPipelineDescriptor) SupportIndirectCommandBuffers() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("supportIndirectCommandBuffers"))
	return rv
}


// SetSupportIndirectCommandBuffers sets the value of the supportIndirectCommandBuffers property.
// A Boolean value that determines whether you can encode commands into an indirect command buffer using the render pipeline.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/supportIndirectCommandBuffers
func (r_ RenderPipelineDescriptor) SetSupportIndirectCommandBuffers(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setSupportIndirectCommandBuffers:"), value)
}
// The size of the control point indices in a control point index buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/tessellationControlPointIndexType
func (r_ RenderPipelineDescriptor) TessellationControlPointIndexType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("tessellationControlPointIndexType"))
	return rv
}


// SetTessellationControlPointIndexType sets the value of the tessellationControlPointIndexType property.
// The size of the control point indices in a control point index buffer.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/tessellationControlPointIndexType
func (r_ RenderPipelineDescriptor) SetTessellationControlPointIndexType(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setTessellationControlPointIndexType:"), value)
}
// The format of the tessellation factors in the tessellation factor buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/tessellationFactorFormat
func (r_ RenderPipelineDescriptor) TessellationFactorFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("tessellationFactorFormat"))
	return rv
}


// SetTessellationFactorFormat sets the value of the tessellationFactorFormat property.
// The format of the tessellation factors in the tessellation factor buffer.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/tessellationFactorFormat
func (r_ RenderPipelineDescriptor) SetTessellationFactorFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setTessellationFactorFormat:"), value)
}
// The step function for determining the tessellation factors for a patch from the tessellation factor buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/tessellationFactorStepFunction
func (r_ RenderPipelineDescriptor) TessellationFactorStepFunction() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("tessellationFactorStepFunction"))
	return rv
}


// SetTessellationFactorStepFunction sets the value of the tessellationFactorStepFunction property.
// The step function for determining the tessellation factors for a patch from the tessellation factor buffer.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/tessellationFactorStepFunction
func (r_ RenderPipelineDescriptor) SetTessellationFactorStepFunction(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setTessellationFactorStepFunction:"), value)
}
// The winding order of triangles from the tessellator.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/tessellationOutputWindingOrder
func (r_ RenderPipelineDescriptor) TessellationOutputWindingOrder() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("tessellationOutputWindingOrder"))
	return rv
}


// SetTessellationOutputWindingOrder sets the value of the tessellationOutputWindingOrder property.
// The winding order of triangles from the tessellator.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/tessellationOutputWindingOrder
func (r_ RenderPipelineDescriptor) SetTessellationOutputWindingOrder(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setTessellationOutputWindingOrder:"), value)
}
// The partitioning mode that the tessellator uses to derive the number and spacing of segments for subdividing a corresponding edge.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/tessellationPartitionMode
func (r_ RenderPipelineDescriptor) TessellationPartitionMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("tessellationPartitionMode"))
	return rv
}


// SetTessellationPartitionMode sets the value of the tessellationPartitionMode property.
// The partitioning mode that the tessellator uses to derive the number and spacing of segments for subdividing a corresponding edge.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/tessellationPartitionMode
func (r_ RenderPipelineDescriptor) SetTessellationPartitionMode(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setTessellationPartitionMode:"), value)
}
// An array that contains the buffer mutability options for a render pipeline’s vertex function.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/vertexBuffers
func (r_ RenderPipelineDescriptor) VertexBuffers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("vertexBuffers"))
	return rv
}

// The organization of vertex data in an attribute’s argument table.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/vertexDescriptor
func (r_ RenderPipelineDescriptor) VertexDescriptor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("vertexDescriptor"))
	return rv
}


// SetVertexDescriptor sets the value of the vertexDescriptor property.
// The organization of vertex data in an attribute’s argument table.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/vertexDescriptor
func (r_ RenderPipelineDescriptor) SetVertexDescriptor(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setVertexDescriptor:"), value)
}
// Functions that you can specify as function arguments for the vertex shader when encoding commands that use the pipeline.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/vertexLinkedFunctions
func (r_ RenderPipelineDescriptor) VertexLinkedFunctions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("vertexLinkedFunctions"))
	return rv
}


// SetVertexLinkedFunctions sets the value of the vertexLinkedFunctions property.
// Functions that you can specify as function arguments for the vertex shader when encoding commands that use the pipeline.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/vertexLinkedFunctions
func (r_ RenderPipelineDescriptor) SetVertexLinkedFunctions(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setVertexLinkedFunctions:"), value)
}


