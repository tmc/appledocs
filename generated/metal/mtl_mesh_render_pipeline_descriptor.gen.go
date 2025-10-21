// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MeshRenderPipelineDescriptor] class.
var (
	MeshRenderPipelineDescriptorClass     _MeshRenderPipelineDescriptorClass
	MeshRenderPipelineDescriptorClassOnce sync.Once
)

func getMeshRenderPipelineDescriptorClass() _MeshRenderPipelineDescriptorClass {
	MeshRenderPipelineDescriptorClassOnce.Do(func() {
		MeshRenderPipelineDescriptorClass = _MeshRenderPipelineDescriptorClass{objc.GetClass("MTLMeshRenderPipelineDescriptor")}
	})
	return MeshRenderPipelineDescriptorClass
}

type _MeshRenderPipelineDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [MeshRenderPipelineDescriptor] class.
type IMeshRenderPipelineDescriptor interface {
	objectivec.IObject
	Reset()
}

// An object that configures new render pipeline state objects for mesh shading.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor
type MeshRenderPipelineDescriptor struct {
	objectivec.Object
}

// MeshRenderPipelineDescriptorFrom constructs a [MeshRenderPipelineDescriptor] from an unsafe.Pointer.
//
// An object that configures new render pipeline state objects for mesh shading.
func MeshRenderPipelineDescriptorFrom(ptr unsafe.Pointer) MeshRenderPipelineDescriptor {
	return MeshRenderPipelineDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MeshRenderPipelineDescriptorClass) Alloc() MeshRenderPipelineDescriptor {
	rv := objc.Send[MeshRenderPipelineDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MeshRenderPipelineDescriptorClass) New() MeshRenderPipelineDescriptor {
	rv := objc.Send[MeshRenderPipelineDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MeshRenderPipelineDescriptor) Init() MeshRenderPipelineDescriptor {
	rv := objc.Send[MeshRenderPipelineDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MeshRenderPipelineDescriptor) Autorelease() MeshRenderPipelineDescriptor {
	rv := objc.Send[MeshRenderPipelineDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMeshRenderPipelineDescriptor creates a new MeshRenderPipelineDescriptor instance.
func NewMeshRenderPipelineDescriptor() MeshRenderPipelineDescriptor {
	return getMeshRenderPipelineDescriptorClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/reset()
func (m_ MeshRenderPipelineDescriptor) Reset() {
	objc.Send[objc.ID](m_.ID, objc.Sel("reset"))
}

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/binaryArchives
func (m_ MeshRenderPipelineDescriptor) BinaryArchives() []objc.ID {
	rv := objc.Send[[]objc.ID](m_.ID, objc.Sel("binaryArchives"))
	return rv
}


// SetBinaryArchives sets the value of the binaryArchives property.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/binaryArchives
func (m_ MeshRenderPipelineDescriptor) SetBinaryArchives(value []objc.ID) {
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
	objc.Send[objc.ID](m_.ID, objc.Sel("setBinaryArchives:"), nsArray)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/colorAttachments
func (m_ MeshRenderPipelineDescriptor) ColorAttachments() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("colorAttachments"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/depthAttachmentPixelFormat
func (m_ MeshRenderPipelineDescriptor) DepthAttachmentPixelFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("depthAttachmentPixelFormat"))
	return rv
}


// SetDepthAttachmentPixelFormat sets the value of the depthAttachmentPixelFormat property.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/depthAttachmentPixelFormat
func (m_ MeshRenderPipelineDescriptor) SetDepthAttachmentPixelFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDepthAttachmentPixelFormat:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/fragmentBuffers
func (m_ MeshRenderPipelineDescriptor) FragmentBuffers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("fragmentBuffers"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/fragmentFunction
func (m_ MeshRenderPipelineDescriptor) FragmentFunction() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("fragmentFunction"))
	return rv
}


// SetFragmentFunction sets the value of the fragmentFunction property.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/fragmentFunction
func (m_ MeshRenderPipelineDescriptor) SetFragmentFunction(value objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFragmentFunction:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/fragmentLinkedFunctions
func (m_ MeshRenderPipelineDescriptor) FragmentLinkedFunctions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("fragmentLinkedFunctions"))
	return rv
}


// SetFragmentLinkedFunctions sets the value of the fragmentLinkedFunctions property.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/fragmentLinkedFunctions
func (m_ MeshRenderPipelineDescriptor) SetFragmentLinkedFunctions(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFragmentLinkedFunctions:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/isAlphaToCoverageEnabled
func (m_ MeshRenderPipelineDescriptor) AlphaToCoverageEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("alphaToCoverageEnabled"))
	return rv
}


// SetAlphaToCoverageEnabled sets the value of the alphaToCoverageEnabled property.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/isAlphaToCoverageEnabled
func (m_ MeshRenderPipelineDescriptor) SetAlphaToCoverageEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlphaToCoverageEnabled:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/isAlphaToOneEnabled
func (m_ MeshRenderPipelineDescriptor) AlphaToOneEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("alphaToOneEnabled"))
	return rv
}


// SetAlphaToOneEnabled sets the value of the alphaToOneEnabled property.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/isAlphaToOneEnabled
func (m_ MeshRenderPipelineDescriptor) SetAlphaToOneEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlphaToOneEnabled:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/isRasterizationEnabled
func (m_ MeshRenderPipelineDescriptor) RasterizationEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("rasterizationEnabled"))
	return rv
}


// SetRasterizationEnabled sets the value of the rasterizationEnabled property.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/isRasterizationEnabled
func (m_ MeshRenderPipelineDescriptor) SetRasterizationEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRasterizationEnabled:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/label
func (m_ MeshRenderPipelineDescriptor) Label() string {
	rv := objc.Send[string](m_.ID, objc.Sel("label"))
	return rv
}


// SetLabel sets the value of the label property.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/label
func (m_ MeshRenderPipelineDescriptor) SetLabel(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/maxTotalThreadgroupsPerMeshGrid
func (m_ MeshRenderPipelineDescriptor) MaxTotalThreadgroupsPerMeshGrid() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("maxTotalThreadgroupsPerMeshGrid"))
	return rv
}


// SetMaxTotalThreadgroupsPerMeshGrid sets the value of the maxTotalThreadgroupsPerMeshGrid property.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/maxTotalThreadgroupsPerMeshGrid
func (m_ MeshRenderPipelineDescriptor) SetMaxTotalThreadgroupsPerMeshGrid(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxTotalThreadgroupsPerMeshGrid:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/maxTotalThreadsPerMeshThreadgroup
func (m_ MeshRenderPipelineDescriptor) MaxTotalThreadsPerMeshThreadgroup() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("maxTotalThreadsPerMeshThreadgroup"))
	return rv
}


// SetMaxTotalThreadsPerMeshThreadgroup sets the value of the maxTotalThreadsPerMeshThreadgroup property.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/maxTotalThreadsPerMeshThreadgroup
func (m_ MeshRenderPipelineDescriptor) SetMaxTotalThreadsPerMeshThreadgroup(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxTotalThreadsPerMeshThreadgroup:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/maxTotalThreadsPerObjectThreadgroup
func (m_ MeshRenderPipelineDescriptor) MaxTotalThreadsPerObjectThreadgroup() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("maxTotalThreadsPerObjectThreadgroup"))
	return rv
}


// SetMaxTotalThreadsPerObjectThreadgroup sets the value of the maxTotalThreadsPerObjectThreadgroup property.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/maxTotalThreadsPerObjectThreadgroup
func (m_ MeshRenderPipelineDescriptor) SetMaxTotalThreadsPerObjectThreadgroup(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxTotalThreadsPerObjectThreadgroup:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/maxVertexAmplificationCount
func (m_ MeshRenderPipelineDescriptor) MaxVertexAmplificationCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("maxVertexAmplificationCount"))
	return rv
}


// SetMaxVertexAmplificationCount sets the value of the maxVertexAmplificationCount property.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/maxVertexAmplificationCount
func (m_ MeshRenderPipelineDescriptor) SetMaxVertexAmplificationCount(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxVertexAmplificationCount:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/meshBuffers
func (m_ MeshRenderPipelineDescriptor) MeshBuffers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("meshBuffers"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/meshFunction
func (m_ MeshRenderPipelineDescriptor) MeshFunction() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("meshFunction"))
	return rv
}


// SetMeshFunction sets the value of the meshFunction property.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/meshFunction
func (m_ MeshRenderPipelineDescriptor) SetMeshFunction(value objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMeshFunction:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/meshLinkedFunctions
func (m_ MeshRenderPipelineDescriptor) MeshLinkedFunctions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("meshLinkedFunctions"))
	return rv
}


// SetMeshLinkedFunctions sets the value of the meshLinkedFunctions property.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/meshLinkedFunctions
func (m_ MeshRenderPipelineDescriptor) SetMeshLinkedFunctions(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMeshLinkedFunctions:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/meshThreadgroupSizeIsMultipleOfThreadExecutionWidth
func (m_ MeshRenderPipelineDescriptor) MeshThreadgroupSizeIsMultipleOfThreadExecutionWidth() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("meshThreadgroupSizeIsMultipleOfThreadExecutionWidth"))
	return rv
}


// SetMeshThreadgroupSizeIsMultipleOfThreadExecutionWidth sets the value of the meshThreadgroupSizeIsMultipleOfThreadExecutionWidth property.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/meshThreadgroupSizeIsMultipleOfThreadExecutionWidth
func (m_ MeshRenderPipelineDescriptor) SetMeshThreadgroupSizeIsMultipleOfThreadExecutionWidth(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMeshThreadgroupSizeIsMultipleOfThreadExecutionWidth:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/objectBuffers
func (m_ MeshRenderPipelineDescriptor) ObjectBuffers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("objectBuffers"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/objectFunction
func (m_ MeshRenderPipelineDescriptor) ObjectFunction() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("objectFunction"))
	return rv
}


// SetObjectFunction sets the value of the objectFunction property.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/objectFunction
func (m_ MeshRenderPipelineDescriptor) SetObjectFunction(value objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setObjectFunction:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/objectLinkedFunctions
func (m_ MeshRenderPipelineDescriptor) ObjectLinkedFunctions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("objectLinkedFunctions"))
	return rv
}


// SetObjectLinkedFunctions sets the value of the objectLinkedFunctions property.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/objectLinkedFunctions
func (m_ MeshRenderPipelineDescriptor) SetObjectLinkedFunctions(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setObjectLinkedFunctions:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/objectThreadgroupSizeIsMultipleOfThreadExecutionWidth
func (m_ MeshRenderPipelineDescriptor) ObjectThreadgroupSizeIsMultipleOfThreadExecutionWidth() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("objectThreadgroupSizeIsMultipleOfThreadExecutionWidth"))
	return rv
}


// SetObjectThreadgroupSizeIsMultipleOfThreadExecutionWidth sets the value of the objectThreadgroupSizeIsMultipleOfThreadExecutionWidth property.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/objectThreadgroupSizeIsMultipleOfThreadExecutionWidth
func (m_ MeshRenderPipelineDescriptor) SetObjectThreadgroupSizeIsMultipleOfThreadExecutionWidth(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setObjectThreadgroupSizeIsMultipleOfThreadExecutionWidth:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/payloadMemoryLength
func (m_ MeshRenderPipelineDescriptor) PayloadMemoryLength() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("payloadMemoryLength"))
	return rv
}


// SetPayloadMemoryLength sets the value of the payloadMemoryLength property.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/payloadMemoryLength
func (m_ MeshRenderPipelineDescriptor) SetPayloadMemoryLength(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPayloadMemoryLength:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/rasterSampleCount
func (m_ MeshRenderPipelineDescriptor) RasterSampleCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("rasterSampleCount"))
	return rv
}


// SetRasterSampleCount sets the value of the rasterSampleCount property.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/rasterSampleCount
func (m_ MeshRenderPipelineDescriptor) SetRasterSampleCount(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRasterSampleCount:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/requiredThreadsPerMeshThreadgroup
func (m_ MeshRenderPipelineDescriptor) RequiredThreadsPerMeshThreadgroup() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("requiredThreadsPerMeshThreadgroup"))
	return rv
}


// SetRequiredThreadsPerMeshThreadgroup sets the value of the requiredThreadsPerMeshThreadgroup property.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/requiredThreadsPerMeshThreadgroup
func (m_ MeshRenderPipelineDescriptor) SetRequiredThreadsPerMeshThreadgroup(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRequiredThreadsPerMeshThreadgroup:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/requiredThreadsPerObjectThreadgroup
func (m_ MeshRenderPipelineDescriptor) RequiredThreadsPerObjectThreadgroup() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("requiredThreadsPerObjectThreadgroup"))
	return rv
}


// SetRequiredThreadsPerObjectThreadgroup sets the value of the requiredThreadsPerObjectThreadgroup property.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/requiredThreadsPerObjectThreadgroup
func (m_ MeshRenderPipelineDescriptor) SetRequiredThreadsPerObjectThreadgroup(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRequiredThreadsPerObjectThreadgroup:"), value)
}

// A value that enables or disables shader validation for the pipeline.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/shaderValidation
func (m_ MeshRenderPipelineDescriptor) ShaderValidation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("shaderValidation"))
	return rv
}


// SetShaderValidation sets the value of the shaderValidation property.
// A value that enables or disables shader validation for the pipeline.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/shaderValidation
func (m_ MeshRenderPipelineDescriptor) SetShaderValidation(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShaderValidation:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/stencilAttachmentPixelFormat
func (m_ MeshRenderPipelineDescriptor) StencilAttachmentPixelFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("stencilAttachmentPixelFormat"))
	return rv
}


// SetStencilAttachmentPixelFormat sets the value of the stencilAttachmentPixelFormat property.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/stencilAttachmentPixelFormat
func (m_ MeshRenderPipelineDescriptor) SetStencilAttachmentPixelFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStencilAttachmentPixelFormat:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/supportIndirectCommandBuffers
func (m_ MeshRenderPipelineDescriptor) SupportIndirectCommandBuffers() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("supportIndirectCommandBuffers"))
	return rv
}


// SetSupportIndirectCommandBuffers sets the value of the supportIndirectCommandBuffers property.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/supportIndirectCommandBuffers
func (m_ MeshRenderPipelineDescriptor) SetSupportIndirectCommandBuffers(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSupportIndirectCommandBuffers:"), value)
}



