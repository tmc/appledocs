// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTL4ComputePipelineDescriptor] class.
var (
	MTL4ComputePipelineDescriptorClass     _MTL4ComputePipelineDescriptorClass
	MTL4ComputePipelineDescriptorClassOnce sync.Once
)

func getMTL4ComputePipelineDescriptorClass() _MTL4ComputePipelineDescriptorClass {
	MTL4ComputePipelineDescriptorClassOnce.Do(func() {
		MTL4ComputePipelineDescriptorClass = _MTL4ComputePipelineDescriptorClass{objc.GetClass("MTL4ComputePipelineDescriptor")}
	})
	return MTL4ComputePipelineDescriptorClass
}

type _MTL4ComputePipelineDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [MTL4ComputePipelineDescriptor] class.
type IMTL4ComputePipelineDescriptor interface {
	IMTL4PipelineDescriptor
}

// Describes a compute pipeline state.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4ComputePipelineDescriptor
type MTL4ComputePipelineDescriptor struct {
	MTL4PipelineDescriptor
}

// MTL4ComputePipelineDescriptorFrom constructs a [MTL4ComputePipelineDescriptor] from an unsafe.Pointer.
//
// Describes a compute pipeline state.
func MTL4ComputePipelineDescriptorFrom(ptr unsafe.Pointer) MTL4ComputePipelineDescriptor {
	return MTL4ComputePipelineDescriptor{
		MTL4PipelineDescriptor: MTL4PipelineDescriptorFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTL4ComputePipelineDescriptorClass) Alloc() MTL4ComputePipelineDescriptor {
	rv := objc.Send[MTL4ComputePipelineDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTL4ComputePipelineDescriptorClass) New() MTL4ComputePipelineDescriptor {
	rv := objc.Send[MTL4ComputePipelineDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4ComputePipelineDescriptor) Init() MTL4ComputePipelineDescriptor {
	rv := objc.Send[MTL4ComputePipelineDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4ComputePipelineDescriptor) Autorelease() MTL4ComputePipelineDescriptor {
	rv := objc.Send[MTL4ComputePipelineDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4ComputePipelineDescriptor creates a new MTL4ComputePipelineDescriptor instance.
func NewMTL4ComputePipelineDescriptor() MTL4ComputePipelineDescriptor {
	return getMTL4ComputePipelineDescriptorClass().New()
}


// A boolean value indicating whether the compute pipeline supports linking binary functions.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4computepipelinedescriptor/supportbinarylinking
func (m_ MTL4ComputePipelineDescriptor) SupportBinaryLinking() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("supportBinaryLinking"))
	return rv
}


// SetSupportBinaryLinking sets the value of the supportBinaryLinking property.
// A boolean value indicating whether the compute pipeline supports linking binary functions.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4computepipelinedescriptor/supportbinarylinking
func (m_ MTL4ComputePipelineDescriptor) SetSupportBinaryLinking(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSupportBinaryLinking:"), value)
}

// A descriptor representing the compute pipeline’s function.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4computepipelinedescriptor/computefunctiondescriptor
func (m_ MTL4ComputePipelineDescriptor) ComputeFunctionDescriptor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("computeFunctionDescriptor"))
	return rv
}


// SetComputeFunctionDescriptor sets the value of the computeFunctionDescriptor property.
// A descriptor representing the compute pipeline’s function.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4computepipelinedescriptor/computefunctiondescriptor
func (m_ MTL4ComputePipelineDescriptor) SetComputeFunctionDescriptor(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setComputeFunctionDescriptor:"), value)
}

// An object that contains information about functions to link to the compute pipeline.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4computepipelinedescriptor/staticlinkingdescriptor
func (m_ MTL4ComputePipelineDescriptor) StaticLinkingDescriptor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("staticLinkingDescriptor"))
	return rv
}


// SetStaticLinkingDescriptor sets the value of the staticLinkingDescriptor property.
// An object that contains information about functions to link to the compute pipeline.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4computepipelinedescriptor/staticlinkingdescriptor
func (m_ MTL4ComputePipelineDescriptor) SetStaticLinkingDescriptor(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStaticLinkingDescriptor:"), value)
}

// A value indicating whether the pipeline supports Metal indirect command buffers.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4computepipelinedescriptor/supportindirectcommandbuffers
func (m_ MTL4ComputePipelineDescriptor) SupportIndirectCommandBuffers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("supportIndirectCommandBuffers"))
	return rv
}


// SetSupportIndirectCommandBuffers sets the value of the supportIndirectCommandBuffers property.
// A value indicating whether the pipeline supports Metal indirect command buffers.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4computepipelinedescriptor/supportindirectcommandbuffers
func (m_ MTL4ComputePipelineDescriptor) SetSupportIndirectCommandBuffers(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSupportIndirectCommandBuffers:"), value)
}

// The required number of threads per threadgroup for compute dispatches.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4computepipelinedescriptor/requiredthreadsperthreadgroup
func (m_ MTL4ComputePipelineDescriptor) RequiredThreadsPerThreadgroup() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("requiredThreadsPerThreadgroup"))
	return rv
}


// SetRequiredThreadsPerThreadgroup sets the value of the requiredThreadsPerThreadgroup property.
// The required number of threads per threadgroup for compute dispatches.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4computepipelinedescriptor/requiredthreadsperthreadgroup
func (m_ MTL4ComputePipelineDescriptor) SetRequiredThreadsPerThreadgroup(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRequiredThreadsPerThreadgroup:"), value)
}

// A boolean value indicating whether each dimension of the threadgroup size is a multiple of its
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4computepipelinedescriptor/threadgroupsizeismultipleofthreadexecutionwidth
func (m_ MTL4ComputePipelineDescriptor) ThreadGroupSizeIsMultipleOfThreadExecutionWidth() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("threadGroupSizeIsMultipleOfThreadExecutionWidth"))
	return rv
}


// SetThreadGroupSizeIsMultipleOfThreadExecutionWidth sets the value of the threadGroupSizeIsMultipleOfThreadExecutionWidth property.
// A boolean value indicating whether each dimension of the threadgroup size is a multiple of its

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4computepipelinedescriptor/threadgroupsizeismultipleofthreadexecutionwidth
func (m_ MTL4ComputePipelineDescriptor) SetThreadGroupSizeIsMultipleOfThreadExecutionWidth(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setThreadGroupSizeIsMultipleOfThreadExecutionWidth:"), value)
}

// The maximum total number of threads that Metal can execute in a single threadgroup for the compute function.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4ComputePipelineDescriptor/maxTotalThreadsPerThreadgroup
func (m_ MTL4ComputePipelineDescriptor) MaxTotalThreadsPerThreadgroup() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("maxTotalThreadsPerThreadgroup"))
	return rv
}


// SetMaxTotalThreadsPerThreadgroup sets the value of the maxTotalThreadsPerThreadgroup property.
// The maximum total number of threads that Metal can execute in a single threadgroup for the compute function.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4ComputePipelineDescriptor/maxTotalThreadsPerThreadgroup
func (m_ MTL4ComputePipelineDescriptor) SetMaxTotalThreadsPerThreadgroup(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxTotalThreadsPerThreadgroup:"), value)
}



