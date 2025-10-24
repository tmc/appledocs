// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MTL4ComputePipelineDescriptor */


/* debug [class_header]: Header for MTL4ComputePipelineDescriptor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTL4ComputePipelineDescriptor */
// An interface definition for the [MTL4ComputePipelineDescriptor] class.
type IMTL4ComputePipelineDescriptor interface {
	IMTL4PipelineDescriptor
	
/* debug [class_interface_properties]: Properties for MTL4ComputePipelineDescriptor */
	// properties:
	ComputeFunctionDescriptor() IMTL4FunctionDescriptor
	SetComputeFunctionDescriptor(value IMTL4FunctionDescriptor)
	MaxTotalThreadsPerThreadgroup() uint
	SetMaxTotalThreadsPerThreadgroup(value uint)
	RequiredThreadsPerThreadgroup() objc.IObject /* cross-framework: MTLSize */
	SetRequiredThreadsPerThreadgroup(value objc.IObject /* cross-framework: MTLSize */)
	StaticLinkingDescriptor() IMTL4StaticLinkingDescriptor
	SetStaticLinkingDescriptor(value IMTL4StaticLinkingDescriptor)
	SupportBinaryLinking() bool
	SetSupportBinaryLinking(value bool)
	SupportIndirectCommandBuffers() MTL4IndirectCommandBufferSupportState
	SetSupportIndirectCommandBuffers(value MTL4IndirectCommandBufferSupportState)
	ThreadGroupSizeIsMultipleOfThreadExecutionWidth() bool
	SetThreadGroupSizeIsMultipleOfThreadExecutionWidth(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTL4ComputePipelineDescriptor */
	// methods:
	Reset()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTL4ComputePipelineDescriptor */
// Alloc allocates a new instance without initialization.
func (mc _MTL4ComputePipelineDescriptorClass) Alloc() MTL4ComputePipelineDescriptor {
	rv := objc.Send[MTL4ComputePipelineDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTL4ComputePipelineDescriptor */
// Describes a compute pipeline state.


// Describes a compute pipeline state.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTL4ComputePipelineDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTL4ComputePipelineDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTL4ComputePipelineDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTL4ComputePipelineDescriptor */

// Resets the descriptor to its default values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4ComputePipelineDescriptor/reset()
func (m_ MTL4ComputePipelineDescriptor) Reset() {
	objc.Send[objc.ID](m_.ID, objc.Sel("reset"))
}/* debug [instance_methods/method]: Reset */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTL4ComputePipelineDescriptor */

// A descriptor representing the compute pipeline’s function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4ComputePipelineDescriptor/computeFunctionDescriptor
func (m_ MTL4ComputePipelineDescriptor) ComputeFunctionDescriptor() IMTL4FunctionDescriptor {
	rv := objc.Send[MTL4FunctionDescriptor](m_.ID, objc.Sel("computeFunctionDescriptor"))
	return rv
}/* debug [instance_properties/getter]: computeFunctionDescriptor */


// A descriptor representing the compute pipeline’s function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4ComputePipelineDescriptor/computeFunctionDescriptor
func (m_ MTL4ComputePipelineDescriptor) SetComputeFunctionDescriptor(value IMTL4FunctionDescriptor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setComputeFunctionDescriptor:"), value)
}/* debug [instance_properties/setter]: computeFunctionDescriptor */


// The maximum total number of threads that Metal can execute in a single threadgroup for the compute function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4ComputePipelineDescriptor/maxTotalThreadsPerThreadgroup
func (m_ MTL4ComputePipelineDescriptor) MaxTotalThreadsPerThreadgroup() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("maxTotalThreadsPerThreadgroup"))
	return rv
}/* debug [instance_properties/getter]: maxTotalThreadsPerThreadgroup */


// The maximum total number of threads that Metal can execute in a single threadgroup for the compute function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4ComputePipelineDescriptor/maxTotalThreadsPerThreadgroup
func (m_ MTL4ComputePipelineDescriptor) SetMaxTotalThreadsPerThreadgroup(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxTotalThreadsPerThreadgroup:"), value)
}/* debug [instance_properties/setter]: maxTotalThreadsPerThreadgroup */


// The required number of threads per threadgroup for compute dispatches.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4ComputePipelineDescriptor/requiredThreadsPerThreadgroup
func (m_ MTL4ComputePipelineDescriptor) RequiredThreadsPerThreadgroup() objc.IObject /* cross-framework: MTLSize */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("requiredThreadsPerThreadgroup"))
	return rv
}/* debug [instance_properties/getter]: requiredThreadsPerThreadgroup */


// The required number of threads per threadgroup for compute dispatches.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4ComputePipelineDescriptor/requiredThreadsPerThreadgroup
func (m_ MTL4ComputePipelineDescriptor) SetRequiredThreadsPerThreadgroup(value objc.IObject /* cross-framework: MTLSize */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRequiredThreadsPerThreadgroup:"), value)
}/* debug [instance_properties/setter]: requiredThreadsPerThreadgroup */


// An object that contains information about functions to link to the compute pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4ComputePipelineDescriptor/staticLinkingDescriptor
func (m_ MTL4ComputePipelineDescriptor) StaticLinkingDescriptor() IMTL4StaticLinkingDescriptor {
	rv := objc.Send[MTL4StaticLinkingDescriptor](m_.ID, objc.Sel("staticLinkingDescriptor"))
	return rv
}/* debug [instance_properties/getter]: staticLinkingDescriptor */


// An object that contains information about functions to link to the compute pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4ComputePipelineDescriptor/staticLinkingDescriptor
func (m_ MTL4ComputePipelineDescriptor) SetStaticLinkingDescriptor(value IMTL4StaticLinkingDescriptor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStaticLinkingDescriptor:"), value)
}/* debug [instance_properties/setter]: staticLinkingDescriptor */


// A boolean value indicating whether the compute pipeline supports linking binary functions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4ComputePipelineDescriptor/supportBinaryLinking
func (m_ MTL4ComputePipelineDescriptor) SupportBinaryLinking() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("supportBinaryLinking"))
	return rv
}/* debug [instance_properties/getter]: supportBinaryLinking */


// A boolean value indicating whether the compute pipeline supports linking binary functions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4ComputePipelineDescriptor/supportBinaryLinking
func (m_ MTL4ComputePipelineDescriptor) SetSupportBinaryLinking(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSupportBinaryLinking:"), value)
}/* debug [instance_properties/setter]: supportBinaryLinking */


// A value indicating whether the pipeline supports Metal indirect command buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4ComputePipelineDescriptor/supportIndirectCommandBuffers
func (m_ MTL4ComputePipelineDescriptor) SupportIndirectCommandBuffers() MTL4IndirectCommandBufferSupportState {
	rv := objc.Send[MTL4IndirectCommandBufferSupportState](m_.ID, objc.Sel("supportIndirectCommandBuffers"))
	return rv
}/* debug [instance_properties/getter]: supportIndirectCommandBuffers */


// A value indicating whether the pipeline supports Metal indirect command buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4ComputePipelineDescriptor/supportIndirectCommandBuffers
func (m_ MTL4ComputePipelineDescriptor) SetSupportIndirectCommandBuffers(value MTL4IndirectCommandBufferSupportState) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSupportIndirectCommandBuffers:"), value)
}/* debug [instance_properties/setter]: supportIndirectCommandBuffers */


// A boolean value indicating whether each dimension of the threadgroup size is a multiple of its corresponding thread execution width.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4ComputePipelineDescriptor/threadGroupSizeIsMultipleOfThreadExecutionWidth
func (m_ MTL4ComputePipelineDescriptor) ThreadGroupSizeIsMultipleOfThreadExecutionWidth() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("threadGroupSizeIsMultipleOfThreadExecutionWidth"))
	return rv
}/* debug [instance_properties/getter]: threadGroupSizeIsMultipleOfThreadExecutionWidth */


// A boolean value indicating whether each dimension of the threadgroup size is a multiple of its corresponding thread execution width.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4ComputePipelineDescriptor/threadGroupSizeIsMultipleOfThreadExecutionWidth
func (m_ MTL4ComputePipelineDescriptor) SetThreadGroupSizeIsMultipleOfThreadExecutionWidth(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setThreadGroupSizeIsMultipleOfThreadExecutionWidth:"), value)
}/* debug [instance_properties/setter]: threadGroupSizeIsMultipleOfThreadExecutionWidth */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTL4ComputePipelineDescriptor */



