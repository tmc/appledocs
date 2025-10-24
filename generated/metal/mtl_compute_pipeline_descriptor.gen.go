// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLComputePipelineDescriptor */


/* debug [class_header]: Header for MTLComputePipelineDescriptor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ComputePipelineDescriptor */
// An interface definition for the [ComputePipelineDescriptor] class.
type IComputePipelineDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ComputePipelineDescriptor */
	// properties:
	BinaryArchives() []objc.ID
	SetBinaryArchives(value []objc.ID)
	Buffers() IMTLPipelineBufferDescriptorArray
	ComputeFunction() unsafe.Pointer
	SetComputeFunction(value unsafe.Pointer)
	InsertLibraries() []objc.ID
	SetInsertLibraries(value []objc.ID)
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	LinkedFunctions() IMTLLinkedFunctions
	SetLinkedFunctions(value IMTLLinkedFunctions)
	MaxCallStackDepth() uint
	SetMaxCallStackDepth(value uint)
	MaxTotalThreadsPerThreadgroup() uint
	SetMaxTotalThreadsPerThreadgroup(value uint)
	PreloadedLibraries() []objc.ID
	SetPreloadedLibraries(value []objc.ID)
	RequiredThreadsPerThreadgroup() objc.IObject /* cross-framework: MTLSize */
	SetRequiredThreadsPerThreadgroup(value objc.IObject /* cross-framework: MTLSize */)
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ComputePipelineDescriptor */
	// methods:
	Reset()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ComputePipelineDescriptor */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ComputePipelineDescriptor */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ComputePipelineDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ComputePipelineDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ComputePipelineDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ComputePipelineDescriptor */

// Resets all compute pipeline descriptor properties to their default values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/reset()
func (c_ ComputePipelineDescriptor) Reset() {
	objc.Send[objc.ID](c_.ID, objc.Sel("reset"))
}/* debug [instance_methods/method]: Reset */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ComputePipelineDescriptor */

// The binary archives that contain any precompiled shader functions to link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/binaryArchives
func (c_ ComputePipelineDescriptor) BinaryArchives() []objc.ID {
	rv := objc.Send[[]objc.ID](c_.ID, objc.Sel("binaryArchives"))
	return rv
}/* debug [instance_properties/getter]: binaryArchives */


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
}/* debug [instance_properties/setter]: binaryArchives */


// The buffer mutability options to apply to the next kernel call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/buffers
func (c_ ComputePipelineDescriptor) Buffers() IMTLPipelineBufferDescriptorArray {
	rv := objc.Send[PipelineBufferDescriptorArray](c_.ID, objc.Sel("buffers"))
	return rv
}/* debug [instance_properties/getter]: buffers */


// The compute kernel the pipeline calls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/computeFunction
func (c_ ComputePipelineDescriptor) ComputeFunction() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("computeFunction"))
	return rv
}/* debug [instance_properties/getter]: computeFunction */


// The compute kernel the pipeline calls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/computeFunction
func (c_ ComputePipelineDescriptor) SetComputeFunction(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setComputeFunction:"), value)
}/* debug [instance_properties/setter]: computeFunction */


// The dynamic libraries that contain precompiled shader functions you want to link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/insertLibraries
func (c_ ComputePipelineDescriptor) InsertLibraries() []objc.ID {
	rv := objc.Send[[]objc.ID](c_.ID, objc.Sel("insertLibraries"))
	return rv
}/* debug [instance_properties/getter]: insertLibraries */


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
}/* debug [instance_properties/setter]: insertLibraries */


// A string that identifies the instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/label
func (c_ ComputePipelineDescriptor) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// A string that identifies the instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/label
func (c_ ComputePipelineDescriptor) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLabel:"), value)
}/* debug [instance_properties/setter]: label */


// The functions with available function pointers for the next kernel call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/linkedFunctions
func (c_ ComputePipelineDescriptor) LinkedFunctions() IMTLLinkedFunctions {
	rv := objc.Send[LinkedFunctions](c_.ID, objc.Sel("linkedFunctions"))
	return rv
}/* debug [instance_properties/getter]: linkedFunctions */


// The functions with available function pointers for the next kernel call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/linkedFunctions
func (c_ ComputePipelineDescriptor) SetLinkedFunctions(value IMTLLinkedFunctions) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLinkedFunctions:"), value)
}/* debug [instance_properties/setter]: linkedFunctions */


// The maximum recursive call depth for dynamic library, visible, and intersection functions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/maxCallStackDepth
func (c_ ComputePipelineDescriptor) MaxCallStackDepth() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("maxCallStackDepth"))
	return rv
}/* debug [instance_properties/getter]: maxCallStackDepth */


// The maximum recursive call depth for dynamic library, visible, and intersection functions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/maxCallStackDepth
func (c_ ComputePipelineDescriptor) SetMaxCallStackDepth(value uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxCallStackDepth:"), value)
}/* debug [instance_properties/setter]: maxCallStackDepth */


// The maximum number of threads in a threadgroup that you can dispatch to the compute function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/maxTotalThreadsPerThreadgroup
func (c_ ComputePipelineDescriptor) MaxTotalThreadsPerThreadgroup() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("maxTotalThreadsPerThreadgroup"))
	return rv
}/* debug [instance_properties/getter]: maxTotalThreadsPerThreadgroup */


// The maximum number of threads in a threadgroup that you can dispatch to the compute function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/maxTotalThreadsPerThreadgroup
func (c_ ComputePipelineDescriptor) SetMaxTotalThreadsPerThreadgroup(value uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxTotalThreadsPerThreadgroup:"), value)
}/* debug [instance_properties/setter]: maxTotalThreadsPerThreadgroup */


// The dynamic libraries that contain precompiled shader functions you want to link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/preloadedLibraries
func (c_ ComputePipelineDescriptor) PreloadedLibraries() []objc.ID {
	rv := objc.Send[[]objc.ID](c_.ID, objc.Sel("preloadedLibraries"))
	return rv
}/* debug [instance_properties/getter]: preloadedLibraries */


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
}/* debug [instance_properties/setter]: preloadedLibraries */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/requiredThreadsPerThreadgroup
func (c_ ComputePipelineDescriptor) RequiredThreadsPerThreadgroup() objc.IObject /* cross-framework: MTLSize */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("requiredThreadsPerThreadgroup"))
	return rv
}/* debug [instance_properties/getter]: requiredThreadsPerThreadgroup */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/requiredThreadsPerThreadgroup
func (c_ ComputePipelineDescriptor) SetRequiredThreadsPerThreadgroup(value objc.IObject /* cross-framework: MTLSize */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRequiredThreadsPerThreadgroup:"), value)
}/* debug [instance_properties/setter]: requiredThreadsPerThreadgroup */


// A value that enables or disables shader validation for the pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/shaderValidation
func (c_ ComputePipelineDescriptor) ShaderValidation() ShaderValidation {
	rv := objc.Send[ShaderValidation](c_.ID, objc.Sel("shaderValidation"))
	return rv
}/* debug [instance_properties/getter]: shaderValidation */


// A value that enables or disables shader validation for the pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/shaderValidation
func (c_ ComputePipelineDescriptor) SetShaderValidation(value ShaderValidation) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShaderValidation:"), value)
}/* debug [instance_properties/setter]: shaderValidation */


// The organization of input and output data for the next kernel call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/stageInputDescriptor
func (c_ ComputePipelineDescriptor) StageInputDescriptor() IMTLStageInputOutputDescriptor {
	rv := objc.Send[StageInputOutputDescriptor](c_.ID, objc.Sel("stageInputDescriptor"))
	return rv
}/* debug [instance_properties/getter]: stageInputDescriptor */


// The organization of input and output data for the next kernel call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/stageInputDescriptor
func (c_ ComputePipelineDescriptor) SetStageInputDescriptor(value IMTLStageInputOutputDescriptor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStageInputDescriptor:"), value)
}/* debug [instance_properties/setter]: stageInputDescriptor */


// A Boolean value that indicates whether you can use the pipeline to create new pipelines by adding binary functions to its callable functions list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/supportAddingBinaryFunctions
func (c_ ComputePipelineDescriptor) SupportAddingBinaryFunctions() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("supportAddingBinaryFunctions"))
	return rv
}/* debug [instance_properties/getter]: supportAddingBinaryFunctions */


// A Boolean value that indicates whether you can use the pipeline to create new pipelines by adding binary functions to its callable functions list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/supportAddingBinaryFunctions
func (c_ ComputePipelineDescriptor) SetSupportAddingBinaryFunctions(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupportAddingBinaryFunctions:"), value)
}/* debug [instance_properties/setter]: supportAddingBinaryFunctions */


// A Boolean value that indicates whether you can encode commands that reference the pipeline state object into an indirect command buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/supportIndirectCommandBuffers
func (c_ ComputePipelineDescriptor) SupportIndirectCommandBuffers() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("supportIndirectCommandBuffers"))
	return rv
}/* debug [instance_properties/getter]: supportIndirectCommandBuffers */


// A Boolean value that indicates whether you can encode commands that reference the pipeline state object into an indirect command buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/supportIndirectCommandBuffers
func (c_ ComputePipelineDescriptor) SetSupportIndirectCommandBuffers(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupportIndirectCommandBuffers:"), value)
}/* debug [instance_properties/setter]: supportIndirectCommandBuffers */


// A Boolean value that indicates whether the threadgroup size is always a multiple of the thread execution width.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/threadGroupSizeIsMultipleOfThreadExecutionWidth
func (c_ ComputePipelineDescriptor) ThreadGroupSizeIsMultipleOfThreadExecutionWidth() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("threadGroupSizeIsMultipleOfThreadExecutionWidth"))
	return rv
}/* debug [instance_properties/getter]: threadGroupSizeIsMultipleOfThreadExecutionWidth */


// A Boolean value that indicates whether the threadgroup size is always a multiple of the thread execution width.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineDescriptor/threadGroupSizeIsMultipleOfThreadExecutionWidth
func (c_ ComputePipelineDescriptor) SetThreadGroupSizeIsMultipleOfThreadExecutionWidth(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setThreadGroupSizeIsMultipleOfThreadExecutionWidth:"), value)
}/* debug [instance_properties/setter]: threadGroupSizeIsMultipleOfThreadExecutionWidth */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLComputePipelineDescriptor */



