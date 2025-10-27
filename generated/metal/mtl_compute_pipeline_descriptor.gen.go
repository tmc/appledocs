// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	

	// properties:
	BinaryArchives() []objc.ID
	SetBinaryArchives(value []objc.ID)
	Buffers() IMTLPipelineBufferDescriptorArray
	ComputeFunction() unsafe.Pointer
	SetComputeFunction(value unsafe.Pointer)
	InsertLibraries() []objc.ID
	SetInsertLibraries(value []objc.ID)
	Label() foundation.foundation.INSString
	SetLabel(value foundation.foundation.INSString)
	LinkedFunctions() IMTLLinkedFunctions
	SetLinkedFunctions(value IMTLLinkedFunctions)
	MaxCallStackDepth() uint
	SetMaxCallStackDepth(value uint)
	MaxTotalThreadsPerThreadgroup() uint
	SetMaxTotalThreadsPerThreadgroup(value uint)
	PreloadedLibraries() []objc.ID
	SetPreloadedLibraries(value []objc.ID)
	RequiredThreadsPerThreadgroup() MTLSize
	SetRequiredThreadsPerThreadgroup(value MTLSize)
	ShaderValidation() ShaderValidation
	SetShaderValidation(value ShaderValidation)
	StageInputDescriptor() IMTLStageInputOutputDescriptor
	SetStageInputDescriptor(value IMTLStageInputOutputDescriptor)
	SupportAddingBinaryFunctions() bool
	SetSupportAddingBinaryFunctions(value bool)
	SupportIndirectCommandBuffers() bool
	SetSupportIndirectCommandBuffers(value bool)
	ThreadGroupSizeIsMultipleOfThreadExecutionWidth() bool
	SetThreadGroupSizeIsMultipleOfThreadExecutionWidth(value bool)


	

	// methods:
	Reset()


}





// Alloc allocates a new instance without initialization.
func (cc _ComputePipelineDescriptorClass) Alloc() ComputePipelineDescriptor {
	rv := objc.Send[ComputePipelineDescriptor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// An instance describing the desired GPU state for a kernel call in a compute pass.
//
// A pipeline descriptor provides information necessary for creating an instance.


// An instance describing the desired GPU state for a kernel call in a compute pass.
//
// [Full Topic]
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




















// Resets all compute pipeline descriptor properties to their default values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/reset()
func (c_ ComputePipelineDescriptor) Reset() {
	objc.Send[objc.ID](c_.ID, objc.Sel("reset"))
}







// The binary archives that contain any precompiled shader functions to link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/binaryArchives
func (c_ ComputePipelineDescriptor) BinaryArchives() []objc.ID {
	rv := objc.Send[[]objc.ID](c_.ID, objc.Sel("binaryArchives"))
	return rv
}


// The binary archives that contain any precompiled shader functions to link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/binaryArchives
func (c_ ComputePipelineDescriptor) SetBinaryArchives(value []objc.ID) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setBinaryArchives:"), nsArray)
}


// The buffer mutability options to apply to the next kernel call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/buffers
func (c_ ComputePipelineDescriptor) Buffers() IMTLPipelineBufferDescriptorArray {
	rv := objc.Send[PipelineBufferDescriptorArray](c_.ID, objc.Sel("buffers"))
	return rv
}


// The compute kernel the pipeline calls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/computeFunction
func (c_ ComputePipelineDescriptor) ComputeFunction() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("computeFunction"))
	return rv
}


// The compute kernel the pipeline calls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/computeFunction
func (c_ ComputePipelineDescriptor) SetComputeFunction(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setComputeFunction:"), value)
}


// The dynamic libraries that contain precompiled shader functions you want to link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/insertLibraries
func (c_ ComputePipelineDescriptor) InsertLibraries() []objc.ID {
	rv := objc.Send[[]objc.ID](c_.ID, objc.Sel("insertLibraries"))
	return rv
}


// The dynamic libraries that contain precompiled shader functions you want to link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/insertLibraries
func (c_ ComputePipelineDescriptor) SetInsertLibraries(value []objc.ID) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setInsertLibraries:"), nsArray)
}


// A string that identifies the instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/label
func (c_ ComputePipelineDescriptor) Label() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("label"))
	return rv
}


// A string that identifies the instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/label
func (c_ ComputePipelineDescriptor) SetLabel(value foundation.foundation.INSString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLabel:"), value)
}


// The functions with available function pointers for the next kernel call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/linkedFunctions
func (c_ ComputePipelineDescriptor) LinkedFunctions() IMTLLinkedFunctions {
	rv := objc.Send[LinkedFunctions](c_.ID, objc.Sel("linkedFunctions"))
	return rv
}


// The functions with available function pointers for the next kernel call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/linkedFunctions
func (c_ ComputePipelineDescriptor) SetLinkedFunctions(value IMTLLinkedFunctions) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLinkedFunctions:"), value)
}


// The maximum recursive call depth for dynamic library, visible, and intersection functions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/maxCallStackDepth
func (c_ ComputePipelineDescriptor) MaxCallStackDepth() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("maxCallStackDepth"))
	return rv
}


// The maximum recursive call depth for dynamic library, visible, and intersection functions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/maxCallStackDepth
func (c_ ComputePipelineDescriptor) SetMaxCallStackDepth(value uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxCallStackDepth:"), value)
}


// The maximum number of threads in a threadgroup that you can dispatch to the compute function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/maxTotalThreadsPerThreadgroup
func (c_ ComputePipelineDescriptor) MaxTotalThreadsPerThreadgroup() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("maxTotalThreadsPerThreadgroup"))
	return rv
}


// The maximum number of threads in a threadgroup that you can dispatch to the compute function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/maxTotalThreadsPerThreadgroup
func (c_ ComputePipelineDescriptor) SetMaxTotalThreadsPerThreadgroup(value uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxTotalThreadsPerThreadgroup:"), value)
}


// The dynamic libraries that contain precompiled shader functions you want to link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/preloadedLibraries
func (c_ ComputePipelineDescriptor) PreloadedLibraries() []objc.ID {
	rv := objc.Send[[]objc.ID](c_.ID, objc.Sel("preloadedLibraries"))
	return rv
}


// The dynamic libraries that contain precompiled shader functions you want to link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/preloadedLibraries
func (c_ ComputePipelineDescriptor) SetPreloadedLibraries(value []objc.ID) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreloadedLibraries:"), nsArray)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/requiredThreadsPerThreadgroup
func (c_ ComputePipelineDescriptor) RequiredThreadsPerThreadgroup() MTLSize {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("requiredThreadsPerThreadgroup"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/requiredThreadsPerThreadgroup
func (c_ ComputePipelineDescriptor) SetRequiredThreadsPerThreadgroup(value MTLSize) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRequiredThreadsPerThreadgroup:"), value)
}


// A value that enables or disables shader validation for the pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/shaderValidation
func (c_ ComputePipelineDescriptor) ShaderValidation() ShaderValidation {
	rv := objc.Send[ShaderValidation](c_.ID, objc.Sel("shaderValidation"))
	return rv
}


// A value that enables or disables shader validation for the pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/shaderValidation
func (c_ ComputePipelineDescriptor) SetShaderValidation(value ShaderValidation) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShaderValidation:"), value)
}


// The organization of input and output data for the next kernel call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/stageInputDescriptor
func (c_ ComputePipelineDescriptor) StageInputDescriptor() IMTLStageInputOutputDescriptor {
	rv := objc.Send[StageInputOutputDescriptor](c_.ID, objc.Sel("stageInputDescriptor"))
	return rv
}


// The organization of input and output data for the next kernel call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/stageInputDescriptor
func (c_ ComputePipelineDescriptor) SetStageInputDescriptor(value IMTLStageInputOutputDescriptor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStageInputDescriptor:"), value)
}


// A Boolean value that indicates whether you can use the pipeline to create new pipelines by adding binary functions to its callable functions list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/supportAddingBinaryFunctions
func (c_ ComputePipelineDescriptor) SupportAddingBinaryFunctions() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("supportAddingBinaryFunctions"))
	return rv
}


// A Boolean value that indicates whether you can use the pipeline to create new pipelines by adding binary functions to its callable functions list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/supportAddingBinaryFunctions
func (c_ ComputePipelineDescriptor) SetSupportAddingBinaryFunctions(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupportAddingBinaryFunctions:"), value)
}


// A Boolean value that indicates whether you can encode commands that reference the pipeline state object into an indirect command buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/supportIndirectCommandBuffers
func (c_ ComputePipelineDescriptor) SupportIndirectCommandBuffers() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("supportIndirectCommandBuffers"))
	return rv
}


// A Boolean value that indicates whether you can encode commands that reference the pipeline state object into an indirect command buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/supportIndirectCommandBuffers
func (c_ ComputePipelineDescriptor) SetSupportIndirectCommandBuffers(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupportIndirectCommandBuffers:"), value)
}


// A Boolean value that indicates whether the threadgroup size is always a multiple of the thread execution width.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/threadGroupSizeIsMultipleOfThreadExecutionWidth
func (c_ ComputePipelineDescriptor) ThreadGroupSizeIsMultipleOfThreadExecutionWidth() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("threadGroupSizeIsMultipleOfThreadExecutionWidth"))
	return rv
}


// A Boolean value that indicates whether the threadgroup size is always a multiple of the thread execution width.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/threadGroupSizeIsMultipleOfThreadExecutionWidth
func (c_ ComputePipelineDescriptor) SetThreadGroupSizeIsMultipleOfThreadExecutionWidth(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setThreadGroupSizeIsMultipleOfThreadExecutionWidth:"), value)
}








