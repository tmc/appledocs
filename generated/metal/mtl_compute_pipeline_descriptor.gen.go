// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ComputePipelineDescriptor] class.
var (
	ComputePipelineDescriptorClass     _ComputePipelineDescriptorClass
	ComputePipelineDescriptorClassOnce sync.Once
)

func getComputePipelineDescriptorClass() _ComputePipelineDescriptorClass {
	ComputePipelineDescriptorClassOnce.Do(func() {
		ComputePipelineDescriptorClass = _ComputePipelineDescriptorClass{objc.GetClass("MTLComputePipelineDescriptor")}
	})
	return ComputePipelineDescriptorClass
}

type _ComputePipelineDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [ComputePipelineDescriptor] class.
type IComputePipelineDescriptor interface {
	objectivec.IObject
}

// An instance describing the desired GPU state for a kernel call in a compute pass.
//
// A pipeline descriptor provides information necessary for creating an instance.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor
type ComputePipelineDescriptor struct {
	objectivec.Object
}

// ComputePipelineDescriptorFrom constructs a [ComputePipelineDescriptor] from an unsafe.Pointer.
//
// An instance describing the desired GPU state for a kernel call in a compute pass.
func ComputePipelineDescriptorFrom(ptr unsafe.Pointer) ComputePipelineDescriptor {
	return ComputePipelineDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _ComputePipelineDescriptorClass) Alloc() ComputePipelineDescriptor {
	rv := objc.Send[ComputePipelineDescriptor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ComputePipelineDescriptorClass) New() ComputePipelineDescriptor {
	rv := objc.Send[ComputePipelineDescriptor](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ComputePipelineDescriptor) Init() ComputePipelineDescriptor {
	rv := objc.Send[ComputePipelineDescriptor](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ComputePipelineDescriptor) Autorelease() ComputePipelineDescriptor {
	rv := objc.Send[ComputePipelineDescriptor](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewComputePipelineDescriptor creates a new ComputePipelineDescriptor instance.
func NewComputePipelineDescriptor() ComputePipelineDescriptor {
	return getComputePipelineDescriptorClass().New()
}


// A string that identifies the instance.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/label
func (c_ ComputePipelineDescriptor) Label() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("label"))
	return rv
}


// SetLabel sets the value of the label property.
// A string that identifies the instance.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/label
func (c_ ComputePipelineDescriptor) SetLabel(value appkit.string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLabel:"), value)
}

// The binary archives that contain any precompiled shader functions to link.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/binaryarchives
func (c_ ComputePipelineDescriptor) BinaryArchives() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("binaryArchives"))
	return rv
}


// SetBinaryArchives sets the value of the binaryArchives property.
// The binary archives that contain any precompiled shader functions to link.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/binaryarchives
func (c_ ComputePipelineDescriptor) SetBinaryArchives(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBinaryArchives:"), value)
}

// The buffer mutability options to apply to the next kernel call.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/buffers
func (c_ ComputePipelineDescriptor) Buffers() MTLPipelineBufferDescriptorArray {
	rv := objc.Send[MTLPipelineBufferDescriptorArray](c_.ID, objc.Sel("buffers"))
	return rv
}


// SetBuffers sets the value of the buffers property.
// The buffer mutability options to apply to the next kernel call.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/buffers
func (c_ ComputePipelineDescriptor) SetBuffers(value IMTLPipelineBufferDescriptorArray) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBuffers:"), value)
}

// The compute kernel the pipeline calls.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/computefunction
func (c_ ComputePipelineDescriptor) ComputeFunction() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("computeFunction"))
	return rv
}


// SetComputeFunction sets the value of the computeFunction property.
// The compute kernel the pipeline calls.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/computefunction
func (c_ ComputePipelineDescriptor) SetComputeFunction(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setComputeFunction:"), value)
}

// The dynamic libraries that contain precompiled shader functions you want to link.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/insertlibraries
func (c_ ComputePipelineDescriptor) InsertLibraries() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("insertLibraries"))
	return rv
}


// SetInsertLibraries sets the value of the insertLibraries property.
// The dynamic libraries that contain precompiled shader functions you want to link.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/insertlibraries
func (c_ ComputePipelineDescriptor) SetInsertLibraries(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInsertLibraries:"), value)
}

// The functions with available function pointers for the next kernel call.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/linkedfunctions
func (c_ ComputePipelineDescriptor) LinkedFunctions() MTLLinkedFunctions {
	rv := objc.Send[MTLLinkedFunctions](c_.ID, objc.Sel("linkedFunctions"))
	return rv
}


// SetLinkedFunctions sets the value of the linkedFunctions property.
// The functions with available function pointers for the next kernel call.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/linkedfunctions
func (c_ ComputePipelineDescriptor) SetLinkedFunctions(value IMTLLinkedFunctions) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLinkedFunctions:"), value)
}

// The maximum recursive call depth for dynamic library, visible, and intersection functions.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/maxcallstackdepth
func (c_ ComputePipelineDescriptor) MaxCallStackDepth() int {
	rv := objc.Send[int](c_.ID, objc.Sel("maxCallStackDepth"))
	return rv
}


// SetMaxCallStackDepth sets the value of the maxCallStackDepth property.
// The maximum recursive call depth for dynamic library, visible, and intersection functions.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/maxcallstackdepth
func (c_ ComputePipelineDescriptor) SetMaxCallStackDepth(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxCallStackDepth:"), value)
}

// The maximum number of threads in a threadgroup that you can dispatch to the compute function.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/maxtotalthreadsperthreadgroup
func (c_ ComputePipelineDescriptor) MaxTotalThreadsPerThreadgroup() int {
	rv := objc.Send[int](c_.ID, objc.Sel("maxTotalThreadsPerThreadgroup"))
	return rv
}


// SetMaxTotalThreadsPerThreadgroup sets the value of the maxTotalThreadsPerThreadgroup property.
// The maximum number of threads in a threadgroup that you can dispatch to the compute function.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/maxtotalthreadsperthreadgroup
func (c_ ComputePipelineDescriptor) SetMaxTotalThreadsPerThreadgroup(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxTotalThreadsPerThreadgroup:"), value)
}

// The dynamic libraries that contain precompiled shader functions you want to link.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/preloadedlibraries
func (c_ ComputePipelineDescriptor) PreloadedLibraries() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("preloadedLibraries"))
	return rv
}


// SetPreloadedLibraries sets the value of the preloadedLibraries property.
// The dynamic libraries that contain precompiled shader functions you want to link.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/preloadedlibraries
func (c_ ComputePipelineDescriptor) SetPreloadedLibraries(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreloadedLibraries:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/requiredthreadsperthreadgroup
func (c_ ComputePipelineDescriptor) RequiredThreadsPerThreadgroup() coregraphics.Size {
	rv := objc.Send[coregraphics.Size](c_.ID, objc.Sel("requiredThreadsPerThreadgroup"))
	return rv
}


// SetRequiredThreadsPerThreadgroup sets the value of the requiredThreadsPerThreadgroup property.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/requiredthreadsperthreadgroup
func (c_ ComputePipelineDescriptor) SetRequiredThreadsPerThreadgroup(value coregraphics.ISize) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRequiredThreadsPerThreadgroup:"), value)
}

// A value that enables or disables shader validation for the pipeline.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/shadervalidation
func (c_ ComputePipelineDescriptor) ShaderValidation() ShaderValidation {
	rv := objc.Send[ShaderValidation](c_.ID, objc.Sel("shaderValidation"))
	return rv
}


// SetShaderValidation sets the value of the shaderValidation property.
// A value that enables or disables shader validation for the pipeline.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/shadervalidation
func (c_ ComputePipelineDescriptor) SetShaderValidation(value IShaderValidation) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShaderValidation:"), value)
}

// The organization of input and output data for the next kernel call.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/stageinputdescriptor
func (c_ ComputePipelineDescriptor) StageInputDescriptor() MTLStageInputOutputDescriptor {
	rv := objc.Send[MTLStageInputOutputDescriptor](c_.ID, objc.Sel("stageInputDescriptor"))
	return rv
}


// SetStageInputDescriptor sets the value of the stageInputDescriptor property.
// The organization of input and output data for the next kernel call.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/stageinputdescriptor
func (c_ ComputePipelineDescriptor) SetStageInputDescriptor(value IMTLStageInputOutputDescriptor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStageInputDescriptor:"), value)
}

// A Boolean value that indicates whether you can use the pipeline to create new pipelines by adding binary functions to its callable functions list.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/supportaddingbinaryfunctions
func (c_ ComputePipelineDescriptor) SupportAddingBinaryFunctions() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("supportAddingBinaryFunctions"))
	return rv
}


// SetSupportAddingBinaryFunctions sets the value of the supportAddingBinaryFunctions property.
// A Boolean value that indicates whether you can use the pipeline to create new pipelines by adding binary functions to its callable functions list.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/supportaddingbinaryfunctions
func (c_ ComputePipelineDescriptor) SetSupportAddingBinaryFunctions(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupportAddingBinaryFunctions:"), value)
}

// A Boolean value that indicates whether you can encode commands that reference the pipeline state object into an indirect command buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/supportindirectcommandbuffers
func (c_ ComputePipelineDescriptor) SupportIndirectCommandBuffers() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("supportIndirectCommandBuffers"))
	return rv
}


// SetSupportIndirectCommandBuffers sets the value of the supportIndirectCommandBuffers property.
// A Boolean value that indicates whether you can encode commands that reference the pipeline state object into an indirect command buffer.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/supportindirectcommandbuffers
func (c_ ComputePipelineDescriptor) SetSupportIndirectCommandBuffers(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupportIndirectCommandBuffers:"), value)
}

// A Boolean value that indicates whether the threadgroup size is always a multiple of the thread execution width.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/threadgroupsizeismultipleofthreadexecutionwidth
func (c_ ComputePipelineDescriptor) ThreadGroupSizeIsMultipleOfThreadExecutionWidth() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("threadGroupSizeIsMultipleOfThreadExecutionWidth"))
	return rv
}


// SetThreadGroupSizeIsMultipleOfThreadExecutionWidth sets the value of the threadGroupSizeIsMultipleOfThreadExecutionWidth property.
// A Boolean value that indicates whether the threadgroup size is always a multiple of the thread execution width.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/threadgroupsizeismultipleofthreadexecutionwidth
func (c_ ComputePipelineDescriptor) SetThreadGroupSizeIsMultipleOfThreadExecutionWidth(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setThreadGroupSizeIsMultipleOfThreadExecutionWidth:"), value)
}



