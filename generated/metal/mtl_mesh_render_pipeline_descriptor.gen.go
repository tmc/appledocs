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
	AlphaToCoverageEnabled() bool
	SetAlphaToCoverageEnabled(value bool)
	AlphaToOneEnabled() bool
	SetAlphaToOneEnabled(value bool)
	RasterizationEnabled() bool
	SetRasterizationEnabled(value bool)
	Label() foundation.foundation.INSString
	SetLabel(value foundation.foundation.INSString)
	MaxTotalThreadgroupsPerMeshGrid() uint
	SetMaxTotalThreadgroupsPerMeshGrid(value uint)
	MaxTotalThreadsPerMeshThreadgroup() uint
	SetMaxTotalThreadsPerMeshThreadgroup(value uint)
	MaxTotalThreadsPerObjectThreadgroup() uint
	SetMaxTotalThreadsPerObjectThreadgroup(value uint)
	MaxVertexAmplificationCount() uint
	SetMaxVertexAmplificationCount(value uint)
	MeshBuffers() IMTLPipelineBufferDescriptorArray
	MeshFunction() unsafe.Pointer
	SetMeshFunction(value unsafe.Pointer)
	MeshLinkedFunctions() IMTLLinkedFunctions
	SetMeshLinkedFunctions(value IMTLLinkedFunctions)
	MeshThreadgroupSizeIsMultipleOfThreadExecutionWidth() bool
	SetMeshThreadgroupSizeIsMultipleOfThreadExecutionWidth(value bool)
	ObjectBuffers() IMTLPipelineBufferDescriptorArray
	ObjectFunction() unsafe.Pointer
	SetObjectFunction(value unsafe.Pointer)
	ObjectLinkedFunctions() IMTLLinkedFunctions
	SetObjectLinkedFunctions(value IMTLLinkedFunctions)
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
	ShaderValidation() ShaderValidation
	SetShaderValidation(value ShaderValidation)
	StencilAttachmentPixelFormat() PixelFormat
	SetStencilAttachmentPixelFormat(value PixelFormat)
	SupportIndirectCommandBuffers() bool
	SetSupportIndirectCommandBuffers(value bool)
	IsAlphaToCoverageEnabled() bool
	SetIsAlphaToCoverageEnabled(value bool)
	IsAlphaToOneEnabled() bool
	SetIsAlphaToOneEnabled(value bool)
	IsRasterizationEnabled() bool
	SetIsRasterizationEnabled(value bool)


	

	// methods:
	Reset()


}





// Alloc allocates a new instance without initialization.
func (mc _MeshRenderPipelineDescriptorClass) Alloc() MeshRenderPipelineDescriptor {
	rv := objc.Send[MeshRenderPipelineDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// An object that configures new render pipeline state objects for mesh shading.


// An object that configures new render pipeline state objects for mesh shading.
//
// [Full Topic]
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




















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/reset()
func (m_ MeshRenderPipelineDescriptor) Reset() {
	objc.Send[objc.ID](m_.ID, objc.Sel("reset"))
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/binaryArchives
func (m_ MeshRenderPipelineDescriptor) BinaryArchives() []objc.ID {
	rv := objc.Send[[]objc.ID](m_.ID, objc.Sel("binaryArchives"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/binaryArchives
func (m_ MeshRenderPipelineDescriptor) SetBinaryArchives(value []objc.ID) {
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


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/colorAttachments
func (m_ MeshRenderPipelineDescriptor) ColorAttachments() IMTLRenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Send[RenderPipelineColorAttachmentDescriptorArray](m_.ID, objc.Sel("colorAttachments"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/depthAttachmentPixelFormat
func (m_ MeshRenderPipelineDescriptor) DepthAttachmentPixelFormat() PixelFormat {
	rv := objc.Send[PixelFormat](m_.ID, objc.Sel("depthAttachmentPixelFormat"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/depthAttachmentPixelFormat
func (m_ MeshRenderPipelineDescriptor) SetDepthAttachmentPixelFormat(value PixelFormat) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDepthAttachmentPixelFormat:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/fragmentBuffers
func (m_ MeshRenderPipelineDescriptor) FragmentBuffers() IMTLPipelineBufferDescriptorArray {
	rv := objc.Send[PipelineBufferDescriptorArray](m_.ID, objc.Sel("fragmentBuffers"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/fragmentFunction
func (m_ MeshRenderPipelineDescriptor) FragmentFunction() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("fragmentFunction"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/fragmentFunction
func (m_ MeshRenderPipelineDescriptor) SetFragmentFunction(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFragmentFunction:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/fragmentLinkedFunctions
func (m_ MeshRenderPipelineDescriptor) FragmentLinkedFunctions() IMTLLinkedFunctions {
	rv := objc.Send[LinkedFunctions](m_.ID, objc.Sel("fragmentLinkedFunctions"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/fragmentLinkedFunctions
func (m_ MeshRenderPipelineDescriptor) SetFragmentLinkedFunctions(value IMTLLinkedFunctions) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFragmentLinkedFunctions:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/isAlphaToCoverageEnabled
func (m_ MeshRenderPipelineDescriptor) AlphaToCoverageEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("alphaToCoverageEnabled"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/isAlphaToCoverageEnabled
func (m_ MeshRenderPipelineDescriptor) SetAlphaToCoverageEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlphaToCoverageEnabled:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/isAlphaToOneEnabled
func (m_ MeshRenderPipelineDescriptor) AlphaToOneEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("alphaToOneEnabled"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/isAlphaToOneEnabled
func (m_ MeshRenderPipelineDescriptor) SetAlphaToOneEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlphaToOneEnabled:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/isRasterizationEnabled
func (m_ MeshRenderPipelineDescriptor) RasterizationEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("rasterizationEnabled"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/isRasterizationEnabled
func (m_ MeshRenderPipelineDescriptor) SetRasterizationEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRasterizationEnabled:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/label
func (m_ MeshRenderPipelineDescriptor) Label() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("label"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/label
func (m_ MeshRenderPipelineDescriptor) SetLabel(value foundation.foundation.INSString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/maxTotalThreadgroupsPerMeshGrid
func (m_ MeshRenderPipelineDescriptor) MaxTotalThreadgroupsPerMeshGrid() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("maxTotalThreadgroupsPerMeshGrid"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/maxTotalThreadgroupsPerMeshGrid
func (m_ MeshRenderPipelineDescriptor) SetMaxTotalThreadgroupsPerMeshGrid(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxTotalThreadgroupsPerMeshGrid:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/maxTotalThreadsPerMeshThreadgroup
func (m_ MeshRenderPipelineDescriptor) MaxTotalThreadsPerMeshThreadgroup() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("maxTotalThreadsPerMeshThreadgroup"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/maxTotalThreadsPerMeshThreadgroup
func (m_ MeshRenderPipelineDescriptor) SetMaxTotalThreadsPerMeshThreadgroup(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxTotalThreadsPerMeshThreadgroup:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/maxTotalThreadsPerObjectThreadgroup
func (m_ MeshRenderPipelineDescriptor) MaxTotalThreadsPerObjectThreadgroup() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("maxTotalThreadsPerObjectThreadgroup"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/maxTotalThreadsPerObjectThreadgroup
func (m_ MeshRenderPipelineDescriptor) SetMaxTotalThreadsPerObjectThreadgroup(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxTotalThreadsPerObjectThreadgroup:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/maxVertexAmplificationCount
func (m_ MeshRenderPipelineDescriptor) MaxVertexAmplificationCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("maxVertexAmplificationCount"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/maxVertexAmplificationCount
func (m_ MeshRenderPipelineDescriptor) SetMaxVertexAmplificationCount(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxVertexAmplificationCount:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/meshBuffers
func (m_ MeshRenderPipelineDescriptor) MeshBuffers() IMTLPipelineBufferDescriptorArray {
	rv := objc.Send[PipelineBufferDescriptorArray](m_.ID, objc.Sel("meshBuffers"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/meshFunction
func (m_ MeshRenderPipelineDescriptor) MeshFunction() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("meshFunction"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/meshFunction
func (m_ MeshRenderPipelineDescriptor) SetMeshFunction(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMeshFunction:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/meshLinkedFunctions
func (m_ MeshRenderPipelineDescriptor) MeshLinkedFunctions() IMTLLinkedFunctions {
	rv := objc.Send[LinkedFunctions](m_.ID, objc.Sel("meshLinkedFunctions"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/meshLinkedFunctions
func (m_ MeshRenderPipelineDescriptor) SetMeshLinkedFunctions(value IMTLLinkedFunctions) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMeshLinkedFunctions:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/meshThreadgroupSizeIsMultipleOfThreadExecutionWidth
func (m_ MeshRenderPipelineDescriptor) MeshThreadgroupSizeIsMultipleOfThreadExecutionWidth() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("meshThreadgroupSizeIsMultipleOfThreadExecutionWidth"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/meshThreadgroupSizeIsMultipleOfThreadExecutionWidth
func (m_ MeshRenderPipelineDescriptor) SetMeshThreadgroupSizeIsMultipleOfThreadExecutionWidth(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMeshThreadgroupSizeIsMultipleOfThreadExecutionWidth:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/objectBuffers
func (m_ MeshRenderPipelineDescriptor) ObjectBuffers() IMTLPipelineBufferDescriptorArray {
	rv := objc.Send[PipelineBufferDescriptorArray](m_.ID, objc.Sel("objectBuffers"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/objectFunction
func (m_ MeshRenderPipelineDescriptor) ObjectFunction() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("objectFunction"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/objectFunction
func (m_ MeshRenderPipelineDescriptor) SetObjectFunction(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setObjectFunction:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/objectLinkedFunctions
func (m_ MeshRenderPipelineDescriptor) ObjectLinkedFunctions() IMTLLinkedFunctions {
	rv := objc.Send[LinkedFunctions](m_.ID, objc.Sel("objectLinkedFunctions"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/objectLinkedFunctions
func (m_ MeshRenderPipelineDescriptor) SetObjectLinkedFunctions(value IMTLLinkedFunctions) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setObjectLinkedFunctions:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/objectThreadgroupSizeIsMultipleOfThreadExecutionWidth
func (m_ MeshRenderPipelineDescriptor) ObjectThreadgroupSizeIsMultipleOfThreadExecutionWidth() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("objectThreadgroupSizeIsMultipleOfThreadExecutionWidth"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/objectThreadgroupSizeIsMultipleOfThreadExecutionWidth
func (m_ MeshRenderPipelineDescriptor) SetObjectThreadgroupSizeIsMultipleOfThreadExecutionWidth(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setObjectThreadgroupSizeIsMultipleOfThreadExecutionWidth:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/payloadMemoryLength
func (m_ MeshRenderPipelineDescriptor) PayloadMemoryLength() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("payloadMemoryLength"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/payloadMemoryLength
func (m_ MeshRenderPipelineDescriptor) SetPayloadMemoryLength(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPayloadMemoryLength:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/rasterSampleCount
func (m_ MeshRenderPipelineDescriptor) RasterSampleCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("rasterSampleCount"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/rasterSampleCount
func (m_ MeshRenderPipelineDescriptor) SetRasterSampleCount(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRasterSampleCount:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/requiredThreadsPerMeshThreadgroup
func (m_ MeshRenderPipelineDescriptor) RequiredThreadsPerMeshThreadgroup() MTLSize {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("requiredThreadsPerMeshThreadgroup"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/requiredThreadsPerMeshThreadgroup
func (m_ MeshRenderPipelineDescriptor) SetRequiredThreadsPerMeshThreadgroup(value MTLSize) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRequiredThreadsPerMeshThreadgroup:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/requiredThreadsPerObjectThreadgroup
func (m_ MeshRenderPipelineDescriptor) RequiredThreadsPerObjectThreadgroup() MTLSize {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("requiredThreadsPerObjectThreadgroup"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/requiredThreadsPerObjectThreadgroup
func (m_ MeshRenderPipelineDescriptor) SetRequiredThreadsPerObjectThreadgroup(value MTLSize) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRequiredThreadsPerObjectThreadgroup:"), value)
}


// A value that enables or disables shader validation for the pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/shaderValidation
func (m_ MeshRenderPipelineDescriptor) ShaderValidation() ShaderValidation {
	rv := objc.Send[ShaderValidation](m_.ID, objc.Sel("shaderValidation"))
	return rv
}


// A value that enables or disables shader validation for the pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/shaderValidation
func (m_ MeshRenderPipelineDescriptor) SetShaderValidation(value ShaderValidation) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShaderValidation:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/stencilAttachmentPixelFormat
func (m_ MeshRenderPipelineDescriptor) StencilAttachmentPixelFormat() PixelFormat {
	rv := objc.Send[PixelFormat](m_.ID, objc.Sel("stencilAttachmentPixelFormat"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/stencilAttachmentPixelFormat
func (m_ MeshRenderPipelineDescriptor) SetStencilAttachmentPixelFormat(value PixelFormat) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStencilAttachmentPixelFormat:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/supportIndirectCommandBuffers
func (m_ MeshRenderPipelineDescriptor) SupportIndirectCommandBuffers() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("supportIndirectCommandBuffers"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/supportIndirectCommandBuffers
func (m_ MeshRenderPipelineDescriptor) SetSupportIndirectCommandBuffers(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSupportIndirectCommandBuffers:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlmeshrenderpipelinedescriptor/isalphatocoverageenabled
func (m_ MeshRenderPipelineDescriptor) IsAlphaToCoverageEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isAlphaToCoverageEnabled"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlmeshrenderpipelinedescriptor/isalphatocoverageenabled
func (m_ MeshRenderPipelineDescriptor) SetIsAlphaToCoverageEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsAlphaToCoverageEnabled:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlmeshrenderpipelinedescriptor/isalphatooneenabled
func (m_ MeshRenderPipelineDescriptor) IsAlphaToOneEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isAlphaToOneEnabled"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlmeshrenderpipelinedescriptor/isalphatooneenabled
func (m_ MeshRenderPipelineDescriptor) SetIsAlphaToOneEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsAlphaToOneEnabled:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlmeshrenderpipelinedescriptor/israsterizationenabled
func (m_ MeshRenderPipelineDescriptor) IsRasterizationEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isRasterizationEnabled"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlmeshrenderpipelinedescriptor/israsterizationenabled
func (m_ MeshRenderPipelineDescriptor) SetIsRasterizationEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsRasterizationEnabled:"), value)
}








