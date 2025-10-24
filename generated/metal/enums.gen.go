// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

/* debug [enums.gen.go]: Generating 125 enums for Metal */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum MTL4AlphaToCoverageState (2 cases) */
// MTL4AlphaToCoverageState - Enumeration for controlling alpha-to-coverage state of a pipeline state object.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AlphaToCoverageState
type MTL4AlphaToCoverageState uint

const (
	// MTL4AlphaToCoverageStateDisabled - Disables alpha-to-coverage.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AlphaToCoverageState/disabled
	MTL4AlphaToCoverageStateDisabled MTL4AlphaToCoverageState = 0
	// MTL4AlphaToCoverageStateEnabled - Enables alpha-to-coverage.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AlphaToCoverageState/enabled
	MTL4AlphaToCoverageStateEnabled MTL4AlphaToCoverageState = 0
)

/* debug [enums.gen.go]: Processing enum MTL4AlphaToOneState (2 cases) */
// MTL4AlphaToOneState - Enumeration for controlling alpha-to-one state of a pipeline state object.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AlphaToOneState
type MTL4AlphaToOneState uint

const (
	// MTL4AlphaToOneStateDisabled - Disables alpha-to-one.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AlphaToOneState/disabled
	MTL4AlphaToOneStateDisabled MTL4AlphaToOneState = 0
	// MTL4AlphaToOneStateEnabled - Enables alpha-to-one.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AlphaToOneState/enabled
	MTL4AlphaToOneStateEnabled MTL4AlphaToOneState = 0
)

/* debug [enums.gen.go]: Processing enum MTL4BinaryFunctionOptions (2 cases) */
// MTL4BinaryFunctionOptions - Options for configuring the creation of binary functions.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4BinaryFunctionOptions
type MTL4BinaryFunctionOptions uint

const (
	// MTL4BinaryFunctionOptionNone - Represents the default value: no options.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4BinaryFunctionOptions/MTL4BinaryFunctionOptionNone
	MTL4BinaryFunctionOptionNone MTL4BinaryFunctionOptions = 0
	// MTL4BinaryFunctionOptionPipelineIndependent - Compiles the function to have its function handles return a constant MTLResourceID across   all pipeline states. The function needs to be linked to the pipeline that will use this function.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4BinaryFunctionOptions/pipelineIndependent
	MTL4BinaryFunctionOptionPipelineIndependent MTL4BinaryFunctionOptions = 0
)

/* debug [enums.gen.go]: Processing enum MTL4BlendState (3 cases) */
// MTL4BlendState - Enumeration for controlling the blend state of a pipeline state object.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4BlendState
type MTL4BlendState uint

const (
	// MTL4BlendStateDisabled - Disables blending.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4BlendState/disabled
	MTL4BlendStateDisabled MTL4BlendState = 0
	// MTL4BlendStateEnabled - Enables blending.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4BlendState/enabled
	MTL4BlendStateEnabled MTL4BlendState = 0
	// MTL4BlendStateUnspecialized - Defers determining the blending stage.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4BlendState/unspecialized
	MTL4BlendStateUnspecialized MTL4BlendState = 0
)

/* debug [enums.gen.go]: Processing enum MTL4CommandQueueError (7 cases) */
// MTL4CommandQueueError - Enumeration of kinds of errors that committing an array of command buffers instances can produce.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CommandQueueError-swift.struct/Code
type MTL4CommandQueueError uint

const (
	// MTL4CommandQueueErrorAccessRevoked - Indicates that the system revokes GPU access because it’s responsible for too many timeouts or hangs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CommandQueueError-swift.struct/Code/accessRevoked
	MTL4CommandQueueErrorAccessRevoked MTL4CommandQueueError = 0
	// MTL4CommandQueueErrorDeviceRemoved - Indicates the physical removal of the GPU before the command buffer completed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CommandQueueError-swift.struct/Code/deviceRemoved
	MTL4CommandQueueErrorDeviceRemoved MTL4CommandQueueError = 0
	// MTL4CommandQueueErrorInternal - Indicates an internal problem in the Metal framework.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CommandQueueError-swift.struct/Code/internal
	MTL4CommandQueueErrorInternal MTL4CommandQueueError = 0
	// MTL4CommandQueueErrorNone - Indicates the absence of any problems.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CommandQueueError-swift.struct/Code/none
	MTL4CommandQueueErrorNone MTL4CommandQueueError = 0
	// MTL4CommandQueueErrorNotPermitted - Indicates a process doesn’t have access to a GPU device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CommandQueueError-swift.struct/Code/notPermitted
	MTL4CommandQueueErrorNotPermitted MTL4CommandQueueError = 0
	// MTL4CommandQueueErrorOutOfMemory - Indicates the GPU doesn’t have sufficient memory to execute a command buffer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CommandQueueError-swift.struct/Code/outOfMemory
	MTL4CommandQueueErrorOutOfMemory MTL4CommandQueueError = 0
	// MTL4CommandQueueErrorTimeout - Indicates the workload takes longer to execute than the system allows.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CommandQueueError-swift.struct/Code/timeout
	MTL4CommandQueueErrorTimeout MTL4CommandQueueError = 0
)

/* debug [enums.gen.go]: Processing enum MTL4CompilerTaskStatus (4 cases) */
// MTL4CompilerTaskStatus - Represents the status of a compiler task.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CompilerTaskStatus
type MTL4CompilerTaskStatus uint

const (
	// MTL4CompilerTaskStatusCompiling - The compiler task is currently compiling.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CompilerTaskStatus/compiling
	MTL4CompilerTaskStatusCompiling MTL4CompilerTaskStatus = 0
	// MTL4CompilerTaskStatusFinished - The compiler task is finished.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CompilerTaskStatus/finished
	MTL4CompilerTaskStatusFinished MTL4CompilerTaskStatus = 0
	// MTL4CompilerTaskStatusNone - No status.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CompilerTaskStatus/none
	MTL4CompilerTaskStatusNone MTL4CompilerTaskStatus = 0
	// MTL4CompilerTaskStatusScheduled - The compiler task is currently scheduled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CompilerTaskStatus/scheduled
	MTL4CompilerTaskStatusScheduled MTL4CompilerTaskStatus = 0
)

/* debug [enums.gen.go]: Processing enum MTL4CounterHeapType (2 cases) */
// MTL4CounterHeapType - Defines the type of a 
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CounterHeapType
type MTL4CounterHeapType uint

const (
	// MTL4CounterHeapTypeInvalid - Specifies that   entries contain invalid data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CounterHeapType/invalid
	MTL4CounterHeapTypeInvalid MTL4CounterHeapType = 0
	// MTL4CounterHeapTypeTimestamp - Specifies that   entries contain GPU timestamp data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CounterHeapType/timestamp
	MTL4CounterHeapTypeTimestamp MTL4CounterHeapType = 0
)

/* debug [enums.gen.go]: Processing enum MTL4IndirectCommandBufferSupportState (2 cases) */
// MTL4IndirectCommandBufferSupportState - Enumeration for controlling support for 
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectCommandBufferSupportState
type MTL4IndirectCommandBufferSupportState uint

const (
	// MTL4IndirectCommandBufferSupportStateDisabled - Disables support for indirect command buffers.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectCommandBufferSupportState/disabled
	MTL4IndirectCommandBufferSupportStateDisabled MTL4IndirectCommandBufferSupportState = 0
	// MTL4IndirectCommandBufferSupportStateEnabled - Enables support for indirect command buffers.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4IndirectCommandBufferSupportState/enabled
	MTL4IndirectCommandBufferSupportStateEnabled MTL4IndirectCommandBufferSupportState = 0
)

/* debug [enums.gen.go]: Processing enum MTL4LogicalToPhysicalColorAttachmentMappingState (2 cases) */
// MTL4LogicalToPhysicalColorAttachmentMappingState - Enumerates possible behaviors of how a pipeline maps its logical outputs to its color attachments.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4LogicalToPhysicalColorAttachmentMappingState
type MTL4LogicalToPhysicalColorAttachmentMappingState uint

const (
	// MTL4LogicalToPhysicalColorAttachmentMappingStateIdentity - Treats the logical color attachment descriptor array for render and tile render pipelines to match the physical one.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4LogicalToPhysicalColorAttachmentMappingState/identity
	MTL4LogicalToPhysicalColorAttachmentMappingStateIdentity MTL4LogicalToPhysicalColorAttachmentMappingState = 0
	// MTL4LogicalToPhysicalColorAttachmentMappingStateInherited - Deduces the color attachment mapping by inheriting it from the color attachment map of the current encoder.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4LogicalToPhysicalColorAttachmentMappingState/inherited
	MTL4LogicalToPhysicalColorAttachmentMappingStateInherited MTL4LogicalToPhysicalColorAttachmentMappingState = 0
)

/* debug [enums.gen.go]: Processing enum MTL4PipelineDataSetSerializerConfiguration (2 cases) */
// MTL4PipelineDataSetSerializerConfiguration - Configuration options for pipeline dataset serializer objects.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PipelineDataSetSerializerConfiguration
type MTL4PipelineDataSetSerializerConfiguration uint

const (
	// MTL4PipelineDataSetSerializerConfigurationCaptureBinaries - Enables serializing pipeline binary functions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PipelineDataSetSerializerConfiguration/captureBinaries
	MTL4PipelineDataSetSerializerConfigurationCaptureBinaries MTL4PipelineDataSetSerializerConfiguration = 0
	// MTL4PipelineDataSetSerializerConfigurationCaptureDescriptors - Enables serializing pipeline scripts.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PipelineDataSetSerializerConfiguration/captureDescriptors
	MTL4PipelineDataSetSerializerConfigurationCaptureDescriptors MTL4PipelineDataSetSerializerConfiguration = 0
)

/* debug [enums.gen.go]: Processing enum MTL4RenderEncoderOptions (3 cases) */
// MTL4RenderEncoderOptions - Custom render pass options you specify at encoder creation time.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderEncoderOptions
type MTL4RenderEncoderOptions uint

const (
	// MTL4RenderEncoderOptionNone - Declares that this render pass doesn’t suspend nor resume.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderEncoderOptions/MTL4RenderEncoderOptionNone
	MTL4RenderEncoderOptionNone MTL4RenderEncoderOptions = 0
	// MTL4RenderEncoderOptionResuming - Configures the render pass to as  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderEncoderOptions/resuming
	MTL4RenderEncoderOptionResuming MTL4RenderEncoderOptions = 0
	// MTL4RenderEncoderOptionSuspending - Configures the render pass as  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderEncoderOptions/suspending
	MTL4RenderEncoderOptionSuspending MTL4RenderEncoderOptions = 0
)

/* debug [enums.gen.go]: Processing enum MTL4ShaderReflection (3 cases) */
// MTL4ShaderReflection - Option mask for requesting reflection information at pipeline build time.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4ShaderReflection
type MTL4ShaderReflection uint

const (
	// MTL4ShaderReflectionBindingInfo - Requests reflection information for bindings.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4ShaderReflection/bindingInfo
	MTL4ShaderReflectionBindingInfo MTL4ShaderReflection = 0
	// MTL4ShaderReflectionBufferTypeInfo - Requests reflection information for buffer types.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4ShaderReflection/bufferTypeInfo
	MTL4ShaderReflectionBufferTypeInfo MTL4ShaderReflection = 0
	// MTL4ShaderReflectionNone - Requests no information.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4ShaderReflection/MTL4ShaderReflectionNone
	MTL4ShaderReflectionNone MTL4ShaderReflection = 0
)

/* debug [enums.gen.go]: Processing enum MTL4TimestampGranularity (2 cases) */
// MTL4TimestampGranularity - Provides a hint to the system about the desired accuracy when writing GPU counter timestamps.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4TimestampGranularity
type MTL4TimestampGranularity uint

const (
	// MTL4TimestampGranularityPrecise - A timestamp as precise as possible.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4TimestampGranularity/precise
	MTL4TimestampGranularityPrecise MTL4TimestampGranularity = 0
	// MTL4TimestampGranularityRelaxed - A minimally-invasive timestamp which may be less precise.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4TimestampGranularity/relaxed
	MTL4TimestampGranularityRelaxed MTL4TimestampGranularity = 0
)

/* debug [enums.gen.go]: Processing enum MTL4VisibilityOptions (3 cases) */
// MTL4VisibilityOptions - Memory consistency options for synchronization commands.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4VisibilityOptions
type MTL4VisibilityOptions uint

const (
	// MTL4VisibilityOptionDevice - Flushes caches to the GPU (device) memory coherence point.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4VisibilityOptions/device
	MTL4VisibilityOptionDevice MTL4VisibilityOptions = 0
	// MTL4VisibilityOptionNone - Don’t flush caches. When you use this option on a barrier, it turns it into an execution barrier.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4VisibilityOptions/MTL4VisibilityOptionNone
	MTL4VisibilityOptionNone MTL4VisibilityOptions = 0
	// MTL4VisibilityOptionResourceAlias - Flushes caches to ensure that aliased virtual addresses are memory consistent.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4VisibilityOptions/resourceAlias
	MTL4VisibilityOptionResourceAlias MTL4VisibilityOptions = 0
)

/* debug [enums.gen.go]: Processing enum MTLAccelerationStructureInstanceDescriptorType (5 cases) */
// MTLAccelerationStructureInstanceDescriptorType - Options for specifying different kinds of instance types.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureInstanceDescriptorType
type MTLAccelerationStructureInstanceDescriptorType uint

const (
	// MTLAccelerationStructureInstanceDescriptorTypeDefault - An option specifying that the instance uses the default characteristics.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureInstanceDescriptorType/default
	MTLAccelerationStructureInstanceDescriptorTypeDefault MTLAccelerationStructureInstanceDescriptorType = 0
	// MTLAccelerationStructureInstanceDescriptorTypeIndirect - An option that enables using an instance descriptor memory layout that the GPU can populate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureInstanceDescriptorType/indirect
	MTLAccelerationStructureInstanceDescriptorTypeIndirect MTLAccelerationStructureInstanceDescriptorType = 0
	// MTLAccelerationStructureInstanceDescriptorTypeIndirectMotion - An option specifying that the instance contains motion data, and enables using an instance descriptor memory layout that the GPU can populate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureInstanceDescriptorType/indirectMotion
	MTLAccelerationStructureInstanceDescriptorTypeIndirectMotion MTLAccelerationStructureInstanceDescriptorType = 0
	// MTLAccelerationStructureInstanceDescriptorTypeMotion - An option specifying that the instance contains motion data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureInstanceDescriptorType/motion
	MTLAccelerationStructureInstanceDescriptorTypeMotion MTLAccelerationStructureInstanceDescriptorType = 0
	// MTLAccelerationStructureInstanceDescriptorTypeUserID - An option specifying that the instance contains a user identifier.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureInstanceDescriptorType/userID
	MTLAccelerationStructureInstanceDescriptorTypeUserID MTLAccelerationStructureInstanceDescriptorType = 0
)

/* debug [enums.gen.go]: Processing enum MTLAccelerationStructureInstanceOptions (5 cases) */
// MTLAccelerationStructureInstanceOptions - Options for adjusting the behavior of an instanced acceleration structure.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureInstanceOptions
type MTLAccelerationStructureInstanceOptions uint

const (
	// MTLAccelerationStructureInstanceOptionDisableTriangleCulling - An option that turns off culling for this instance if ray intersector has culling enabled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureInstanceOptions/disableTriangleCulling
	MTLAccelerationStructureInstanceOptionDisableTriangleCulling MTLAccelerationStructureInstanceOptions = 0
	// MTLAccelerationStructureInstanceOptionNone - Specifies the default behavior for resulting acceleration structure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureInstanceOptions/MTLAccelerationStructureInstanceOptionNone
	MTLAccelerationStructureInstanceOptionNone MTLAccelerationStructureInstanceOptions = 0
	// MTLAccelerationStructureInstanceOptionNonOpaque - Specifies that intersectors should treat the instance as non-opaque.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureInstanceOptions/nonOpaque
	MTLAccelerationStructureInstanceOptionNonOpaque MTLAccelerationStructureInstanceOptions = 0
	// MTLAccelerationStructureInstanceOptionOpaque - Specifies that intersectors should treat the instance as opaque.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureInstanceOptions/opaque
	MTLAccelerationStructureInstanceOptionOpaque MTLAccelerationStructureInstanceOptions = 0
	// MTLAccelerationStructureInstanceOptionTriangleFrontFacingWindingCounterClockwise - Specifies that the instance specifies front facing triangles in counter-clockwise order.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureInstanceOptions/triangleFrontFacingWindingCounterClockwise
	MTLAccelerationStructureInstanceOptionTriangleFrontFacingWindingCounterClockwise MTLAccelerationStructureInstanceOptions = 0
)

/* debug [enums.gen.go]: Processing enum MTLAccelerationStructureRefitOptions (2 cases) */
// MTLAccelerationStructureRefitOptions enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureRefitOptions
type MTLAccelerationStructureRefitOptions uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureRefitOptions/perPrimitiveData
	MTLAccelerationStructureRefitOptionPerPrimitiveData MTLAccelerationStructureRefitOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureRefitOptions/vertexData
	MTLAccelerationStructureRefitOptionVertexData MTLAccelerationStructureRefitOptions = 0
)

/* debug [enums.gen.go]: Processing enum MTLAccelerationStructureUsage (6 cases) */
// MTLAccelerationStructureUsage - Options that affect how Metal builds an acceleration structure and the behavior
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureUsage
type MTLAccelerationStructureUsage uint

const (
	// MTLAccelerationStructureUsageExtendedLimits - An option that increases an acceleration structure’s storage capacity.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureUsage/extendedLimits
	MTLAccelerationStructureUsageExtendedLimits MTLAccelerationStructureUsage = 0
	// MTLAccelerationStructureUsageMinimizeMemory - An option that instructs Metal to prioritize building an acceleration structure that needs less memory.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureUsage/minimizeMemory
	MTLAccelerationStructureUsageMinimizeMemory MTLAccelerationStructureUsage = 0
	// MTLAccelerationStructureUsageNone - A sentinel option the represents an empty set of options,   which is the default behavior for building new acceleration structures.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureUsage/MTLAccelerationStructureUsageNone
	MTLAccelerationStructureUsageNone MTLAccelerationStructureUsage = 0
	// MTLAccelerationStructureUsagePreferFastBuild - An option that instructs Metal to build an acceleration structure quickly.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureUsage/preferFastBuild
	MTLAccelerationStructureUsagePreferFastBuild MTLAccelerationStructureUsage = 0
	// MTLAccelerationStructureUsagePreferFastIntersection - An option that instructs Metal to prioritize building an acceleration structure with better intersection performance.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureUsage/preferFastIntersection
	MTLAccelerationStructureUsagePreferFastIntersection MTLAccelerationStructureUsage = 0
	// MTLAccelerationStructureUsageRefit - An option that lets you update an acceleration structure after creating it.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureUsage/refit
	MTLAccelerationStructureUsageRefit MTLAccelerationStructureUsage = 0
)

/* debug [enums.gen.go]: Processing enum MTLArgumentBuffersTier (2 cases) */
// MTLArgumentBuffersTier - The values that determine the limits and capabilities of argument buffers.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgumentBuffersTier
type MTLArgumentBuffersTier uint

const (
	// MTLArgumentBuffersTier1 - Support for tier 1 argument buffers.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgumentBuffersTier/tier1
	MTLArgumentBuffersTier1 MTLArgumentBuffersTier = 0
	// MTLArgumentBuffersTier2 - Support for tier 2 argument buffers.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgumentBuffersTier/tier2
	MTLArgumentBuffersTier2 MTLArgumentBuffersTier = 0
)

/* debug [enums.gen.go]: Processing enum MTLArgumentType (10 cases) */
// MTLArgumentType - The resource type for an argument of a function.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgumentType
type MTLArgumentType uint

const (
	// MTLArgumentTypeBuffer - The argument is a buffer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgumentType/buffer
	MTLArgumentTypeBuffer MTLArgumentType = 0
	// MTLArgumentTypeImageblock - The argument is an imageblock.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgumentType/imageblock
	MTLArgumentTypeImageblock MTLArgumentType = 0
	// MTLArgumentTypeImageblockData - The argument is imageblock data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgumentType/imageblockData
	MTLArgumentTypeImageblockData MTLArgumentType = 0
	// MTLArgumentTypeInstanceAccelerationStructure - The argument is a top-level ray tracing acceleration structure for a set of instances.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgumentType/instanceAccelerationStructure
	MTLArgumentTypeInstanceAccelerationStructure MTLArgumentType = 0
	// MTLArgumentTypeIntersectionFunctionTable - The argument is an intersection function table.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgumentType/intersectionFunctionTable
	MTLArgumentTypeIntersectionFunctionTable MTLArgumentType = 0
	// MTLArgumentTypePrimitiveAccelerationStructure - The argument is a bottom-level ray tracing acceleraton structure for a set of primitives.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgumentType/primitiveAccelerationStructure
	MTLArgumentTypePrimitiveAccelerationStructure MTLArgumentType = 0
	// MTLArgumentTypeSampler - The argument is a texture sampler.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgumentType/sampler
	MTLArgumentTypeSampler MTLArgumentType = 0
	// MTLArgumentTypeTexture - The argument is a texture.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgumentType/texture
	MTLArgumentTypeTexture MTLArgumentType = 0
	// MTLArgumentTypeThreadgroupMemory - The argument is a pointer to threadgroup memory.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgumentType/threadgroupMemory
	MTLArgumentTypeThreadgroupMemory MTLArgumentType = 0
	// MTLArgumentTypeVisibleFunctionTable - The argument is a visible function table.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgumentType/visibleFunctionTable
	MTLArgumentTypeVisibleFunctionTable MTLArgumentType = 0
)

/* debug [enums.gen.go]: Processing enum MTLAttributeFormat (54 cases) */
// MTLAttributeFormat - Values indicating the organization and format of data for function attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat
type MTLAttributeFormat uint

const (
	// MTLAttributeFormatChar - One signed 8-bit two’s complement value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/char
	MTLAttributeFormatChar MTLAttributeFormat = 0
	// MTLAttributeFormatChar2 - Two signed 8-bit two’s complement values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/char2
	MTLAttributeFormatChar2 MTLAttributeFormat = 0
	// MTLAttributeFormatChar2Normalized - Two signed normalized 8-bit two’s complement values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/char2Normalized
	MTLAttributeFormatChar2Normalized MTLAttributeFormat = 0
	// MTLAttributeFormatChar3 - Three signed 8-bit two’s complement values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/char3
	MTLAttributeFormatChar3 MTLAttributeFormat = 0
	// MTLAttributeFormatChar3Normalized - Three signed normalized 8-bit two’s complement values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/char3Normalized
	MTLAttributeFormatChar3Normalized MTLAttributeFormat = 0
	// MTLAttributeFormatChar4 - Four signed 8-bit two’s complement values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/char4
	MTLAttributeFormatChar4 MTLAttributeFormat = 0
	// MTLAttributeFormatChar4Normalized - Four signed normalized 8-bit two’s complement values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/char4Normalized
	MTLAttributeFormatChar4Normalized MTLAttributeFormat = 0
	// MTLAttributeFormatCharNormalized - One signed normalized 8-bit two’s complement value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/charNormalized
	MTLAttributeFormatCharNormalized MTLAttributeFormat = 0
	// MTLAttributeFormatFloat - One single-precision floating-point value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/float
	MTLAttributeFormatFloat MTLAttributeFormat = 0
	// MTLAttributeFormatFloat2 - Two single-precision floating-point values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/float2
	MTLAttributeFormatFloat2 MTLAttributeFormat = 0
	// MTLAttributeFormatFloat3 - Three single-precision floating-point values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/float3
	MTLAttributeFormatFloat3 MTLAttributeFormat = 0
	// MTLAttributeFormatFloat4 - Four single-precision floating-point values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/float4
	MTLAttributeFormatFloat4 MTLAttributeFormat = 0
	// MTLAttributeFormatFloatRG11B10 - One packed 32-bit value representing pixel data containing 11-bit float red and green channels, and a 10-bit float blue channel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/floatRG11B10
	MTLAttributeFormatFloatRG11B10 MTLAttributeFormat = 0
	// MTLAttributeFormatFloatRGB9E5 - One packed 32-bit value representing pixel data containing 9-bit float red, green, and blue channels, and a 5-bit float shared exponent channel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/floatRGB9E5
	MTLAttributeFormatFloatRGB9E5 MTLAttributeFormat = 0
	// MTLAttributeFormatHalf - One half-precision floating-point value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/half
	MTLAttributeFormatHalf MTLAttributeFormat = 0
	// MTLAttributeFormatHalf2 - Two half-precision floating-point values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/half2
	MTLAttributeFormatHalf2 MTLAttributeFormat = 0
	// MTLAttributeFormatHalf3 - Three half-precision floating-point values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/half3
	MTLAttributeFormatHalf3 MTLAttributeFormat = 0
	// MTLAttributeFormatHalf4 - Four half-precision floating-point values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/half4
	MTLAttributeFormatHalf4 MTLAttributeFormat = 0
	// MTLAttributeFormatInt - One signed 32-bit two’s complement value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/int
	MTLAttributeFormatInt MTLAttributeFormat = 0
	// MTLAttributeFormatInt1010102Normalized - One packed 32-bit value with four normalized signed two’s complement integer values, arranged as 10 bits, 10 bits, 10 bits, and 2 bits.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/int1010102Normalized
	MTLAttributeFormatInt1010102Normalized MTLAttributeFormat = 0
	// MTLAttributeFormatInt2 - Two signed 32-bit two’s complement values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/int2
	MTLAttributeFormatInt2 MTLAttributeFormat = 0
	// MTLAttributeFormatInt3 - Three signed 32-bit two’s complement values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/int3
	MTLAttributeFormatInt3 MTLAttributeFormat = 0
	// MTLAttributeFormatInt4 - Four signed 32-bit two’s complement values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/int4
	MTLAttributeFormatInt4 MTLAttributeFormat = 0
	// MTLAttributeFormatInvalid - An invalid format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/invalid
	MTLAttributeFormatInvalid MTLAttributeFormat = 0
	// MTLAttributeFormatShort - One signed 16-bit two’s complement value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/short
	MTLAttributeFormatShort MTLAttributeFormat = 0
	// MTLAttributeFormatShort2 - Two signed 16-bit two’s complement values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/short2
	MTLAttributeFormatShort2 MTLAttributeFormat = 0
	// MTLAttributeFormatShort2Normalized - Two signed normalized 16-bit two’s complement values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/short2Normalized
	MTLAttributeFormatShort2Normalized MTLAttributeFormat = 0
	// MTLAttributeFormatShort3 - Three signed 16-bit two’s complement values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/short3
	MTLAttributeFormatShort3 MTLAttributeFormat = 0
	// MTLAttributeFormatShort3Normalized - Three signed normalized 16-bit two’s complement values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/short3Normalized
	MTLAttributeFormatShort3Normalized MTLAttributeFormat = 0
	// MTLAttributeFormatShort4 - Four signed 16-bit two’s complement values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/short4
	MTLAttributeFormatShort4 MTLAttributeFormat = 0
	// MTLAttributeFormatShort4Normalized - Four signed normalized 16-bit two’s complement values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/short4Normalized
	MTLAttributeFormatShort4Normalized MTLAttributeFormat = 0
	// MTLAttributeFormatShortNormalized - One signed normalized 16-bit two’s complement value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/shortNormalized
	MTLAttributeFormatShortNormalized MTLAttributeFormat = 0
	// MTLAttributeFormatUChar - One unsigned 8-bit value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/uchar
	MTLAttributeFormatUChar MTLAttributeFormat = 0
	// MTLAttributeFormatUChar2 - Two unsigned 8-bit values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/uchar2
	MTLAttributeFormatUChar2 MTLAttributeFormat = 0
	// MTLAttributeFormatUChar2Normalized - Two unsigned normalized 8-bit values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/uchar2Normalized
	MTLAttributeFormatUChar2Normalized MTLAttributeFormat = 0
	// MTLAttributeFormatUChar3 - Three unsigned 8-bit values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/uchar3
	MTLAttributeFormatUChar3 MTLAttributeFormat = 0
	// MTLAttributeFormatUChar3Normalized - Three unsigned normalized 8-bit values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/uchar3Normalized
	MTLAttributeFormatUChar3Normalized MTLAttributeFormat = 0
	// MTLAttributeFormatUChar4 - Four unsigned 8-bit values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/uchar4
	MTLAttributeFormatUChar4 MTLAttributeFormat = 0
	// MTLAttributeFormatUChar4Normalized - Four unsigned normalized 8-bit values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/uchar4Normalized
	MTLAttributeFormatUChar4Normalized MTLAttributeFormat = 0
	// MTLAttributeFormatUChar4Normalized_BGRA - Four unsigned normalized 8-bit values, arranged as blue, green, red, and alpha components.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/uchar4Normalized_bgra
	MTLAttributeFormatUChar4Normalized_BGRA MTLAttributeFormat = 0
	// MTLAttributeFormatUCharNormalized - One unsigned normalized 8-bit value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/ucharNormalized
	MTLAttributeFormatUCharNormalized MTLAttributeFormat = 0
	// MTLAttributeFormatUInt - One unsigned 32-bit value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/uint
	MTLAttributeFormatUInt MTLAttributeFormat = 0
	// MTLAttributeFormatUInt1010102Normalized - One packed 32-bit value with four normalized unsigned integer values, arranged as 10 bits, 10 bits, 10 bits, and 2 bits.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/uint1010102Normalized
	MTLAttributeFormatUInt1010102Normalized MTLAttributeFormat = 0
	// MTLAttributeFormatUInt2 - Two unsigned 32-bit values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/uint2
	MTLAttributeFormatUInt2 MTLAttributeFormat = 0
	// MTLAttributeFormatUInt3 - Three unsigned 32-bit values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/uint3
	MTLAttributeFormatUInt3 MTLAttributeFormat = 0
	// MTLAttributeFormatUInt4 - Four unsigned 32-bit values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/uint4
	MTLAttributeFormatUInt4 MTLAttributeFormat = 0
	// MTLAttributeFormatUShort - One unsigned 16-bit value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/ushort
	MTLAttributeFormatUShort MTLAttributeFormat = 0
	// MTLAttributeFormatUShort2 - Two unsigned 16-bit values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/ushort2
	MTLAttributeFormatUShort2 MTLAttributeFormat = 0
	// MTLAttributeFormatUShort2Normalized - Two unsigned normalized 16-bit values
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/ushort2Normalized
	MTLAttributeFormatUShort2Normalized MTLAttributeFormat = 0
	// MTLAttributeFormatUShort3 - Three unsigned 16-bit values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/ushort3
	MTLAttributeFormatUShort3 MTLAttributeFormat = 0
	// MTLAttributeFormatUShort3Normalized - Three unsigned normalized 16-bit values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/ushort3Normalized
	MTLAttributeFormatUShort3Normalized MTLAttributeFormat = 0
	// MTLAttributeFormatUShort4 - Four unsigned 16-bit values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/ushort4
	MTLAttributeFormatUShort4 MTLAttributeFormat = 0
	// MTLAttributeFormatUShort4Normalized - Four unsigned normalized 16-bit values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/ushort4Normalized
	MTLAttributeFormatUShort4Normalized MTLAttributeFormat = 0
	// MTLAttributeFormatUShortNormalized - One unsigned normalized 16-bit value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeFormat/ushortNormalized
	MTLAttributeFormatUShortNormalized MTLAttributeFormat = 0
)

/* debug [enums.gen.go]: Processing enum MTLBarrierScope (3 cases) */
// MTLBarrierScope - Describes the types of resources that a barrier operates on.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBarrierScope
type MTLBarrierScope uint

const (
	// MTLBarrierScopeBuffers - The barrier affects any buffer objects.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBarrierScope/buffers
	MTLBarrierScopeBuffers MTLBarrierScope = 0
	// MTLBarrierScopeRenderTargets - The barrier affects any render targets.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBarrierScope/renderTargets
	MTLBarrierScopeRenderTargets MTLBarrierScope = 0
	// MTLBarrierScopeTextures - The barrier affects textures.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBarrierScope/textures
	MTLBarrierScopeTextures MTLBarrierScope = 0
)

/* debug [enums.gen.go]: Processing enum MTLBinaryArchiveError (5 cases) */
// MTLBinaryArchiveError - Error codes when creating binary archives of compiled shader code.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBinaryArchiveError-swift.struct/Code
type MTLBinaryArchiveError uint

const (
	// MTLBinaryArchiveErrorCompilationFailure - An error code that indicates the archive’s inability to compile its contents, typically when serializing it to a URL.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBinaryArchiveError-swift.struct/Code/compilationFailure
	MTLBinaryArchiveErrorCompilationFailure MTLBinaryArchiveError = 0
	// MTLBinaryArchiveErrorInternalError - An error code that indicates the Metal framework has an internal problem.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBinaryArchiveError-swift.struct/Code/internalError
	MTLBinaryArchiveErrorInternalError MTLBinaryArchiveError = 0
	// MTLBinaryArchiveErrorInvalidFile - An error code that indicates an app is using an invalid reference to an archive file, typically related to a URL.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBinaryArchiveError-swift.struct/Code/invalidFile
	MTLBinaryArchiveErrorInvalidFile MTLBinaryArchiveError = 0
	// MTLBinaryArchiveErrorNone - An error code that represents the absence of any problems.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBinaryArchiveError-swift.struct/Code/none
	MTLBinaryArchiveErrorNone MTLBinaryArchiveError = 0
	// MTLBinaryArchiveErrorUnexpectedElement - An error code that indicates a problem with a configuration, typically in a descriptor or an archive’s inability to add linked functions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBinaryArchiveError-swift.struct/Code/unexpectedElement
	MTLBinaryArchiveErrorUnexpectedElement MTLBinaryArchiveError = 0
)

/* debug [enums.gen.go]: Processing enum MTLBindingAccess (6 cases) */
// MTLBindingAccess enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBindingAccess
type MTLBindingAccess uint

const (
	// MTLArgumentAccessReadOnly - The function can only read its argument data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBindingAccess/MTLArgumentAccessReadOnly
	MTLArgumentAccessReadOnly MTLBindingAccess = 0
	// MTLArgumentAccessReadWrite - The function can either read or write its argument data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBindingAccess/MTLArgumentAccessReadWrite
	MTLArgumentAccessReadWrite MTLBindingAccess = 0
	// MTLArgumentAccessWriteOnly - The function can only write its argument data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBindingAccess/MTLArgumentAccessWriteOnly
	MTLArgumentAccessWriteOnly MTLBindingAccess = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBindingAccess/readOnly
	MTLBindingAccessReadOnly MTLBindingAccess = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBindingAccess/readWrite
	MTLBindingAccessReadWrite MTLBindingAccess = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBindingAccess/writeOnly
	MTLBindingAccessWriteOnly MTLBindingAccess = 0
)

/* debug [enums.gen.go]: Processing enum MTLBindingType (12 cases) */
// MTLBindingType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBindingType
type MTLBindingType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBindingType/buffer
	MTLBindingTypeBuffer MTLBindingType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBindingType/imageblock
	MTLBindingTypeImageblock MTLBindingType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBindingType/imageblockData
	MTLBindingTypeImageblockData MTLBindingType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBindingType/instanceAccelerationStructure
	MTLBindingTypeInstanceAccelerationStructure MTLBindingType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBindingType/intersectionFunctionTable
	MTLBindingTypeIntersectionFunctionTable MTLBindingType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBindingType/objectPayload
	MTLBindingTypeObjectPayload MTLBindingType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBindingType/primitiveAccelerationStructure
	MTLBindingTypePrimitiveAccelerationStructure MTLBindingType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBindingType/sampler
	MTLBindingTypeSampler MTLBindingType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBindingType/tensor
	MTLBindingTypeTensor MTLBindingType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBindingType/texture
	MTLBindingTypeTexture MTLBindingType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBindingType/threadgroupMemory
	MTLBindingTypeThreadgroupMemory MTLBindingType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBindingType/visibleFunctionTable
	MTLBindingTypeVisibleFunctionTable MTLBindingType = 0
)

/* debug [enums.gen.go]: Processing enum MTLBlendFactor (20 cases) */
// MTLBlendFactor - The source and destination blend factors are often needed to complete specification of a blend operation. In most cases, the blend factor for both RGB values (
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendFactor
type MTLBlendFactor uint

const (
	// MTLBlendFactorBlendAlpha - Blend factor of alpha value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendFactor/blendAlpha
	MTLBlendFactorBlendAlpha MTLBlendFactor = 0
	// MTLBlendFactorBlendColor - Blend factor of RGB values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendFactor/blendColor
	MTLBlendFactorBlendColor MTLBlendFactor = 0
	// MTLBlendFactorDestinationAlpha - Blend factor of destination alpha.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendFactor/destinationAlpha
	MTLBlendFactorDestinationAlpha MTLBlendFactor = 0
	// MTLBlendFactorDestinationColor - Blend factor of destination values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendFactor/destinationColor
	MTLBlendFactorDestinationColor MTLBlendFactor = 0
	// MTLBlendFactorOne - Blend factor of one.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendFactor/one
	MTLBlendFactorOne MTLBlendFactor = 0
	// MTLBlendFactorOneMinusBlendAlpha - Blend factor of one minus alpha value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendFactor/oneMinusBlendAlpha
	MTLBlendFactorOneMinusBlendAlpha MTLBlendFactor = 0
	// MTLBlendFactorOneMinusBlendColor - Blend factor of one minus RGB values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendFactor/oneMinusBlendColor
	MTLBlendFactorOneMinusBlendColor MTLBlendFactor = 0
	// MTLBlendFactorOneMinusDestinationAlpha - Blend factor of one minus destination alpha.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendFactor/oneMinusDestinationAlpha
	MTLBlendFactorOneMinusDestinationAlpha MTLBlendFactor = 0
	// MTLBlendFactorOneMinusDestinationColor - Blend factor of one minus destination values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendFactor/oneMinusDestinationColor
	MTLBlendFactorOneMinusDestinationColor MTLBlendFactor = 0
	// MTLBlendFactorOneMinusSource1Alpha - Blend factor of one minus source alpha. This option supports dual-source blending and reads from the second color output of the fragment function.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendFactor/oneMinusSource1Alpha
	MTLBlendFactorOneMinusSource1Alpha MTLBlendFactor = 0
	// MTLBlendFactorOneMinusSource1Color - Blend factor of one minus source values. This option supports dual-source blending and reads from the second color output of the fragment function.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendFactor/oneMinusSource1Color
	MTLBlendFactorOneMinusSource1Color MTLBlendFactor = 0
	// MTLBlendFactorOneMinusSourceAlpha - Blend factor of one minus source alpha.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendFactor/oneMinusSourceAlpha
	MTLBlendFactorOneMinusSourceAlpha MTLBlendFactor = 0
	// MTLBlendFactorOneMinusSourceColor - Blend factor of one minus source values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendFactor/oneMinusSourceColor
	MTLBlendFactorOneMinusSourceColor MTLBlendFactor = 0
	// MTLBlendFactorSource1Alpha - Blend factor of source alpha. This option supports dual-source blending and reads from the second color output of the fragment function.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendFactor/source1Alpha
	MTLBlendFactorSource1Alpha MTLBlendFactor = 0
	// MTLBlendFactorSource1Color - Blend factor of source values. This option supports dual-source blending and reads from the second color output of the fragment function.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendFactor/source1Color
	MTLBlendFactorSource1Color MTLBlendFactor = 0
	// MTLBlendFactorSourceAlpha - Blend factor of source alpha.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendFactor/sourceAlpha
	MTLBlendFactorSourceAlpha MTLBlendFactor = 0
	// MTLBlendFactorSourceAlphaSaturated - Blend factor of the minimum of either source alpha or one minus destination alpha.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendFactor/sourceAlphaSaturated
	MTLBlendFactorSourceAlphaSaturated MTLBlendFactor = 0
	// MTLBlendFactorSourceColor - Blend factor of source values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendFactor/sourceColor
	MTLBlendFactorSourceColor MTLBlendFactor = 0
	// MTLBlendFactorUnspecialized - Defers assigning the blend factor.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendFactor/unspecialized
	MTLBlendFactorUnspecialized MTLBlendFactor = 0
	// MTLBlendFactorZero - Blend factor of zero.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendFactor/zero
	MTLBlendFactorZero MTLBlendFactor = 0
)

/* debug [enums.gen.go]: Processing enum MTLBlendOperation (6 cases) */
// MTLBlendOperation - For every pixel, 
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendOperation
type MTLBlendOperation uint

const (
	// MTLBlendOperationAdd - Add portions of both source and destination pixel values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendOperation/add
	MTLBlendOperationAdd MTLBlendOperation = 0
	// MTLBlendOperationMax - Maximum of the source and destination pixel values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendOperation/max
	MTLBlendOperationMax MTLBlendOperation = 0
	// MTLBlendOperationMin - Minimum of the source and destination pixel values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendOperation/min
	MTLBlendOperationMin MTLBlendOperation = 0
	// MTLBlendOperationReverseSubtract - Subtract a portion of the source values from a portion of the destination pixel values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendOperation/reverseSubtract
	MTLBlendOperationReverseSubtract MTLBlendOperation = 0
	// MTLBlendOperationSubtract - Subtract a portion of the destination pixel values from a portion of the source.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendOperation/subtract
	MTLBlendOperationSubtract MTLBlendOperation = 0
	// MTLBlendOperationUnspecialized - Defers assigning the blend operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendOperation/unspecialized
	MTLBlendOperationUnspecialized MTLBlendOperation = 0
)

/* debug [enums.gen.go]: Processing enum MTLBlitOption (4 cases) */
// MTLBlitOption - The options that enable behavior for some blit operations.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlitOption
type MTLBlitOption uint

const (
	// MTLBlitOptionDepthFromDepthStencil - A blit option that copies the depth portion of a combined depth and stencil texture to or from a buffer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlitOption/depthFromDepthStencil
	MTLBlitOptionDepthFromDepthStencil MTLBlitOption = 0
	// MTLBlitOptionNone - A blit option that clears other blit options, which removes any optional behavior for a blit operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlitOption/MTLBlitOptionNone
	MTLBlitOptionNone MTLBlitOption = 0
	// MTLBlitOptionRowLinearPVRTC - A blit option that copies PVRTC data between a texture and a buffer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlitOption/rowLinearPVRTC
	MTLBlitOptionRowLinearPVRTC MTLBlitOption = 0
	// MTLBlitOptionStencilFromDepthStencil - A blit option that copies the stencil portion of a combined depth and stencil texture to or from a buffer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlitOption/stencilFromDepthStencil
	MTLBlitOptionStencilFromDepthStencil MTLBlitOption = 0
)

/* debug [enums.gen.go]: Processing enum MTLBufferSparseTier (2 cases) */
// MTLBufferSparseTier - Enumerates the different support levels for sparse buffers.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBufferSparseTier
type MTLBufferSparseTier uint

const (
	// MTLBufferSparseTier1 - Indicates support for sparse buffers tier 1.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBufferSparseTier/tier1
	MTLBufferSparseTier1 MTLBufferSparseTier = 0
	// MTLBufferSparseTierNone - Indicates that the buffer is not sparse.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBufferSparseTier/tierNone
	MTLBufferSparseTierNone MTLBufferSparseTier = 0
)

/* debug [enums.gen.go]: Processing enum MTLCaptureDestination (2 cases) */
// MTLCaptureDestination - The kinds of destinations for captured command data.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCaptureDestination
type MTLCaptureDestination uint

const (
	// MTLCaptureDestinationDeveloperTools - An option specifying that data should be captured to Xcode and that execution should stop in Xcode after the data is captured.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCaptureDestination/developerTools
	MTLCaptureDestinationDeveloperTools MTLCaptureDestination = 0
	// MTLCaptureDestinationGPUTraceDocument - An option specifying that the captured command data should be saved to a GPU trace document.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCaptureDestination/gpuTraceDocument
	MTLCaptureDestinationGPUTraceDocument MTLCaptureDestination = 0
)

/* debug [enums.gen.go]: Processing enum MTLCaptureError (3 cases) */
// MTLCaptureError - Errors returned by capture sessions.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCaptureError
type MTLCaptureError uint

const (
	// MTLCaptureErrorAlreadyCapturing - A capture session is already in progress.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCaptureError/alreadyCapturing
	MTLCaptureErrorAlreadyCapturing MTLCaptureError = 0
	// MTLCaptureErrorInvalidDescriptor - The descriptor contained invalid parameters.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCaptureError/invalidDescriptor
	MTLCaptureErrorInvalidDescriptor MTLCaptureError = 0
	// MTLCaptureErrorNotSupported - The requested capture options are not available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCaptureError/notSupported
	MTLCaptureErrorNotSupported MTLCaptureError = 0
)

/* debug [enums.gen.go]: Processing enum MTLColorWriteMask (7 cases) */
// MTLColorWriteMask - Values used to specify a mask to permit or restrict writing to color channels of a color value.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLColorWriteMask
type MTLColorWriteMask uint

const (
	// MTLColorWriteMaskAll - All color channels are enabled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLColorWriteMask/all
	MTLColorWriteMaskAll MTLColorWriteMask = 0
	// MTLColorWriteMaskAlpha - The alpha color channel is enabled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLColorWriteMask/alpha
	MTLColorWriteMaskAlpha MTLColorWriteMask = 0
	// MTLColorWriteMaskBlue - The blue color channel is enabled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLColorWriteMask/blue
	MTLColorWriteMaskBlue MTLColorWriteMask = 0
	// MTLColorWriteMaskGreen - The green color channel is enabled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLColorWriteMask/green
	MTLColorWriteMaskGreen MTLColorWriteMask = 0
	// MTLColorWriteMaskNone - All color channels are disabled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLColorWriteMask/MTLColorWriteMaskNone
	MTLColorWriteMaskNone MTLColorWriteMask = 0
	// MTLColorWriteMaskRed - The red color channel is enabled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLColorWriteMask/red
	MTLColorWriteMaskRed MTLColorWriteMask = 0
	// MTLColorWriteMaskUnspecialized - Defers assigning the color write mask.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLColorWriteMask/unspecialized
	MTLColorWriteMaskUnspecialized MTLColorWriteMask = 0
)

/* debug [enums.gen.go]: Processing enum MTLCommandBufferError (12 cases) */
// MTLCommandBufferError - Error codes that indicate why a GPU is unable to finish running a command buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandBufferError-swift.struct/Code
type MTLCommandBufferError uint

const (
	// MTLCommandBufferErrorAccessRevoked - An error code that indicates the system has revoked the Metal device’s access because it’s responsible for too many timeouts or hangs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandBufferError-c.enum/MTLCommandBufferErrorAccessRevoked
	MTLCommandBufferErrorAccessRevoked MTLCommandBufferError = 0
	// MTLCommandBufferErrorBlacklisted - A former error code that indicates the system has revoked the Metal device’s access because it’s responsible for too many timeouts or hangs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandBufferError-swift.struct/Code/blacklisted
	MTLCommandBufferErrorBlacklisted MTLCommandBufferError = 0
	// MTLCommandBufferErrorDeviceRemoved - An error code that indicates a person physically removed the GPU device before the command buffer finished running.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandBufferError-swift.struct/Code/deviceRemoved
	MTLCommandBufferErrorDeviceRemoved MTLCommandBufferError = 0
	// MTLCommandBufferErrorInternal - An error code that indicates the Metal framework has an internal problem.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandBufferError-swift.struct/Code/internal
	MTLCommandBufferErrorInternal MTLCommandBufferError = 0
	// MTLCommandBufferErrorInvalidResource - An error code that indicates the command buffer has an invalid reference to resource.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandBufferError-swift.struct/Code/invalidResource
	MTLCommandBufferErrorInvalidResource MTLCommandBufferError = 0
	// MTLCommandBufferErrorMemoryless - An error code that indicates the GPU ran out of one or more of its internal resources that support memoryless render pass attachments.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandBufferError-swift.struct/Code/memoryless
	MTLCommandBufferErrorMemoryless MTLCommandBufferError = 0
	// MTLCommandBufferErrorNone - An error code that represents the absence of any problems.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandBufferError-swift.struct/Code/none
	MTLCommandBufferErrorNone MTLCommandBufferError = 0
	// MTLCommandBufferErrorNotPermitted - An error code that indicates a process doesn’t have access to a GPU device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandBufferError-swift.struct/Code/notPermitted
	MTLCommandBufferErrorNotPermitted MTLCommandBufferError = 0
	// MTLCommandBufferErrorOutOfMemory - An error code that indicates the GPU device doesn’t have sufficient memory to execute a command buffer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandBufferError-swift.struct/Code/outOfMemory
	MTLCommandBufferErrorOutOfMemory MTLCommandBufferError = 0
	// MTLCommandBufferErrorPageFault - An error code that indicates the command buffer generated a page fault the GPU can’t service.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandBufferError-swift.struct/Code/pageFault
	MTLCommandBufferErrorPageFault MTLCommandBufferError = 0
	// MTLCommandBufferErrorStackOverflow - An error code that indicates the GPU terminated the command buffer because a kernel function of tile shader used too many stack frames.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandBufferError-swift.struct/Code/stackOverflow
	MTLCommandBufferErrorStackOverflow MTLCommandBufferError = 0
	// MTLCommandBufferErrorTimeout - An error code that indicates the system interrupted and terminated the command buffer because it took more time to execute than the system allows.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandBufferError-swift.struct/Code/timeout
	MTLCommandBufferErrorTimeout MTLCommandBufferError = 0
)

/* debug [enums.gen.go]: Processing enum MTLCommandBufferErrorOption (2 cases) */
// MTLCommandBufferErrorOption - Options for reporting errors from a command buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandBufferErrorOption
type MTLCommandBufferErrorOption uint

const (
	// MTLCommandBufferErrorOptionEncoderExecutionStatus - An option that instructs a command buffer to save additional details about a GPU runtime error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandBufferErrorOption/encoderExecutionStatus
	MTLCommandBufferErrorOptionEncoderExecutionStatus MTLCommandBufferErrorOption = 0
	// MTLCommandBufferErrorOptionNone - An option that clears a command buffer’s error options.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandBufferErrorOption/MTLCommandBufferErrorOptionNone
	MTLCommandBufferErrorOptionNone MTLCommandBufferErrorOption = 0
)

/* debug [enums.gen.go]: Processing enum MTLCommandBufferStatus (6 cases) */
// MTLCommandBufferStatus - The discrete states for a command buffer that represent its life cycle stages.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandBufferStatus
type MTLCommandBufferStatus uint

const (
	// MTLCommandBufferStatusCommitted - A command buffer’s third state, which indicates the command queue is preparing to schedule the command buffer by resolving its dependencies.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandBufferStatus/committed
	MTLCommandBufferStatusCommitted MTLCommandBufferStatus = 0
	// MTLCommandBufferStatusCompleted - A command buffer’s successful, final state, which indicates the GPU finished running the command buffer’s commands without any problems.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandBufferStatus/completed
	MTLCommandBufferStatusCompleted MTLCommandBufferStatus = 0
	// MTLCommandBufferStatusEnqueued - A command buffer’s second state, which indicates its command queue is reserving a place for it.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandBufferStatus/enqueued
	MTLCommandBufferStatusEnqueued MTLCommandBufferStatus = 0
	// MTLCommandBufferStatusError - A command buffer’s unsuccessful, final state, which indicates the GPU stopped running the buffer’s commands because of a runtime issue.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandBufferStatus/error
	MTLCommandBufferStatusError MTLCommandBufferStatus = 0
	// MTLCommandBufferStatusNotEnqueued - A command buffer’s initial state, which indicates its command queue isn’t reserving a place for it.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandBufferStatus/notEnqueued
	MTLCommandBufferStatusNotEnqueued MTLCommandBufferStatus = 0
	// MTLCommandBufferStatusScheduled - A command buffer’s fourth state, which indicates the command buffer has its resources ready and is waiting for the GPU to run its commands.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandBufferStatus/scheduled
	MTLCommandBufferStatusScheduled MTLCommandBufferStatus = 0
)

/* debug [enums.gen.go]: Processing enum MTLCommandEncoderErrorState (5 cases) */
// MTLCommandEncoderErrorState - Possible error conditions for the command encoder’s commands.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandEncoderErrorState
type MTLCommandEncoderErrorState uint

const (
	// MTLCommandEncoderErrorStateAffected - An error state that indicates the GPU failed to fully execute the commands because of an error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandEncoderErrorState/affected
	MTLCommandEncoderErrorStateAffected MTLCommandEncoderErrorState = 0
	// MTLCommandEncoderErrorStateCompleted - A state that indicates the GPU successfully executed the commands without any errors.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandEncoderErrorState/completed
	MTLCommandEncoderErrorStateCompleted MTLCommandEncoderErrorState = 0
	// MTLCommandEncoderErrorStateFaulted - An error state that indicates the commands in the command buffer are the cause of an error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandEncoderErrorState/faulted
	MTLCommandEncoderErrorStateFaulted MTLCommandEncoderErrorState = 0
	// MTLCommandEncoderErrorStatePending - An error state that indicates the GPU didn’t execute the commands.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandEncoderErrorState/pending
	MTLCommandEncoderErrorStatePending MTLCommandEncoderErrorState = 0
	// MTLCommandEncoderErrorStateUnknown - An error state that indicates the command buffer doesn’t know the state of its commands on the GPU.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandEncoderErrorState/unknown
	MTLCommandEncoderErrorStateUnknown MTLCommandEncoderErrorState = 0
)

/* debug [enums.gen.go]: Processing enum MTLCompareFunction (8 cases) */
// MTLCompareFunction - Options used to specify how a sample compare operation should be performed on a depth texture.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompareFunction
type MTLCompareFunction uint

const (
	// MTLCompareFunctionAlways - A new value always passes the comparison test.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompareFunction/always
	MTLCompareFunctionAlways MTLCompareFunction = 0
	// MTLCompareFunctionEqual - A new value passes the comparison test if it is equal to the existing value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompareFunction/equal
	MTLCompareFunctionEqual MTLCompareFunction = 0
	// MTLCompareFunctionGreater - A new value passes the comparison test if it is greater than the existing value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompareFunction/greater
	MTLCompareFunctionGreater MTLCompareFunction = 0
	// MTLCompareFunctionGreaterEqual - A new value passes the comparison test if it is greater than or equal to the existing value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompareFunction/greaterEqual
	MTLCompareFunctionGreaterEqual MTLCompareFunction = 0
	// MTLCompareFunctionLess - A new value passes the comparison test if it is less than the existing value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompareFunction/less
	MTLCompareFunctionLess MTLCompareFunction = 0
	// MTLCompareFunctionLessEqual - A new value passes the comparison test if it is less than or equal to the existing value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompareFunction/lessEqual
	MTLCompareFunctionLessEqual MTLCompareFunction = 0
	// MTLCompareFunctionNever - A new value never passes the comparison test.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompareFunction/never
	MTLCompareFunctionNever MTLCompareFunction = 0
	// MTLCompareFunctionNotEqual - A new value passes the comparison test if it is not equal to the existing value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompareFunction/notEqual
	MTLCompareFunctionNotEqual MTLCompareFunction = 0
)

/* debug [enums.gen.go]: Processing enum MTLCompileSymbolVisibility (2 cases) */
// MTLCompileSymbolVisibility enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileSymbolVisibility
type MTLCompileSymbolVisibility uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileSymbolVisibility/default
	MTLCompileSymbolVisibilityDefault MTLCompileSymbolVisibility = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCompileSymbolVisibility/hidden
	MTLCompileSymbolVisibilityHidden MTLCompileSymbolVisibility = 0
)

/* debug [enums.gen.go]: Processing enum MTLCounterSampleBufferError (3 cases) */
// MTLCounterSampleBufferError - The underlying error code type that indicates why a GPU driver can’t create a counter sample buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCounterSampleBufferError-swift.struct/Code
type MTLCounterSampleBufferError uint

const (
	// MTLCounterSampleBufferErrorInternal - An error code that indicates the Metal framework has an internal problem.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCounterSampleBufferError-swift.struct/Code/internal
	MTLCounterSampleBufferErrorInternal MTLCounterSampleBufferError = 0
	// MTLCounterSampleBufferErrorInvalid - An error code that indicates when a counter-sample buffer descriptor has at   least one invalid property.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCounterSampleBufferError-swift.struct/Code/invalid
	MTLCounterSampleBufferErrorInvalid MTLCounterSampleBufferError = 0
	// MTLCounterSampleBufferErrorOutOfMemory - An error code that indicates the GPU device doesn’t have sufficient memory to create a counter sample buffer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCounterSampleBufferError-swift.struct/Code/outOfMemory
	MTLCounterSampleBufferErrorOutOfMemory MTLCounterSampleBufferError = 0
)

/* debug [enums.gen.go]: Processing enum MTLCounterSamplingPoint (5 cases) */
// MTLCounterSamplingPoint - Options for different times when you can sample GPU counters.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCounterSamplingPoint
type MTLCounterSamplingPoint uint

const (
	// MTLCounterSamplingPointAtBlitBoundary - Counter sampling is allowed between blit commands in a blit pass.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCounterSamplingPoint/atBlitBoundary
	MTLCounterSamplingPointAtBlitBoundary MTLCounterSamplingPoint = 0
	// MTLCounterSamplingPointAtDispatchBoundary - Counter sampling is allowed between kernel dispatches in a compute pass.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCounterSamplingPoint/atDispatchBoundary
	MTLCounterSamplingPointAtDispatchBoundary MTLCounterSamplingPoint = 0
	// MTLCounterSamplingPointAtDrawBoundary - Counter sampling is allowed between draw commands in a render pass.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCounterSamplingPoint/atDrawBoundary
	MTLCounterSamplingPointAtDrawBoundary MTLCounterSamplingPoint = 0
	// MTLCounterSamplingPointAtStageBoundary - Counter sampling is allowed at the start and end of a render pass’s vertex and fragment stages, and at the start and end of compute and blit passes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCounterSamplingPoint/atStageBoundary
	MTLCounterSamplingPointAtStageBoundary MTLCounterSamplingPoint = 0
	// MTLCounterSamplingPointAtTileDispatchBoundary - Counter sampling is allowed between tile dispatches in a render pass.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCounterSamplingPoint/atTileDispatchBoundary
	MTLCounterSamplingPointAtTileDispatchBoundary MTLCounterSamplingPoint = 0
)

/* debug [enums.gen.go]: Processing enum MTLCPUCacheMode (2 cases) */
// MTLCPUCacheMode - Options for the CPU cache mode that define the CPU mapping of the resource.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCPUCacheMode
type MTLCPUCacheMode uint

const (
	// MTLCPUCacheModeDefaultCache - The default CPU cache mode for the resource, which guarantees that read and write operations are executed in the expected order.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCPUCacheMode/defaultCache
	MTLCPUCacheModeDefaultCache MTLCPUCacheMode = 0
	// MTLCPUCacheModeWriteCombined - A write-combined CPU cache mode that is optimized for resources that the CPU writes into, but never reads.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCPUCacheMode/writeCombined
	MTLCPUCacheModeWriteCombined MTLCPUCacheMode = 0
)

/* debug [enums.gen.go]: Processing enum MTLCullMode (3 cases) */
// MTLCullMode - The mode that determines whether to perform culling and which type of primitive to cull.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCullMode
type MTLCullMode uint

const (
	// MTLCullModeBack - Culls back-facing primitives.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCullMode/back
	MTLCullModeBack MTLCullMode = 0
	// MTLCullModeFront - Culls front-facing primitives.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCullMode/front
	MTLCullModeFront MTLCullMode = 0
	// MTLCullModeNone - Does not cull any primitives.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCullMode/none
	MTLCullModeNone MTLCullMode = 0
)

/* debug [enums.gen.go]: Processing enum MTLCurveBasis (4 cases) */
// MTLCurveBasis enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCurveBasis
type MTLCurveBasis uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCurveBasis/bezier
	MTLCurveBasisBezier MTLCurveBasis = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCurveBasis/bSpline
	MTLCurveBasisBSpline MTLCurveBasis = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCurveBasis/catmullRom
	MTLCurveBasisCatmullRom MTLCurveBasis = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCurveBasis/linear
	MTLCurveBasisLinear MTLCurveBasis = 0
)

/* debug [enums.gen.go]: Processing enum MTLCurveEndCaps (3 cases) */
// MTLCurveEndCaps enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCurveEndCaps
type MTLCurveEndCaps uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCurveEndCaps/disk
	MTLCurveEndCapsDisk MTLCurveEndCaps = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCurveEndCaps/none
	MTLCurveEndCapsNone MTLCurveEndCaps = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCurveEndCaps/sphere
	MTLCurveEndCapsSphere MTLCurveEndCaps = 0
)

/* debug [enums.gen.go]: Processing enum MTLCurveType (2 cases) */
// MTLCurveType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCurveType
type MTLCurveType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCurveType/flat
	MTLCurveTypeFlat MTLCurveType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCurveType/round
	MTLCurveTypeRound MTLCurveType = 0
)

/* debug [enums.gen.go]: Processing enum MTLDataType (97 cases) */
// MTLDataType - The types of GPU functions, including shaders and compute kernels.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType
type MTLDataType uint

const (
	// MTLDataTypeArray - An array instance.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/array
	MTLDataTypeArray MTLDataType = 0
	// MTLDataTypeBFloat - A 16-bit, brain floating-point value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/bfloat
	MTLDataTypeBFloat MTLDataType = 0
	// MTLDataTypeBFloat2 - A two-component vector with 16-bit, brain floating-point values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/bfloat2
	MTLDataTypeBFloat2 MTLDataType = 0
	// MTLDataTypeBFloat3 - A three-component vector with 16-bit, brain floating-point values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/bfloat3
	MTLDataTypeBFloat3 MTLDataType = 0
	// MTLDataTypeBFloat4 - A four-component vector with 16-bit, brain floating-point values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/bfloat4
	MTLDataTypeBFloat4 MTLDataType = 0
	// MTLDataTypeBool - A Boolean value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/bool
	MTLDataTypeBool MTLDataType = 0
	// MTLDataTypeBool2 - A two-component Boolean vector.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/bool2
	MTLDataTypeBool2 MTLDataType = 0
	// MTLDataTypeBool3 - A three-component Boolean vector.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/bool3
	MTLDataTypeBool3 MTLDataType = 0
	// MTLDataTypeBool4 - A four-component Boolean vector.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/bool4
	MTLDataTypeBool4 MTLDataType = 0
	// MTLDataTypeChar - An 8-bit, signed integer value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/char
	MTLDataTypeChar MTLDataType = 0
	// MTLDataTypeChar2 - A two-component vector with 8-bit, signed integer values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/char2
	MTLDataTypeChar2 MTLDataType = 0
	// MTLDataTypeChar3 - A three-component vector with 8-bit, signed integer values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/char3
	MTLDataTypeChar3 MTLDataType = 0
	// MTLDataTypeChar4 - A four-component vector with 8-bit, signed integer values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/char4
	MTLDataTypeChar4 MTLDataType = 0
	// MTLDataTypeComputePipeline - A Metal compute pipeline instance.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/computePipeline
	MTLDataTypeComputePipeline MTLDataType = 0
	// MTLDataTypeDepthStencilState - Represents a data type corresponding to a depth-stencil state object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/depthStencilState
	MTLDataTypeDepthStencilState MTLDataType = 0
	// MTLDataTypeFloat - A 32-bit floating-point value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/float
	MTLDataTypeFloat MTLDataType = 0
	// MTLDataTypeFloat2 - A two-component vector with 32-bit floating-point values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/float2
	MTLDataTypeFloat2 MTLDataType = 0
	// MTLDataTypeFloat2x2 - A 2x2 component matrix with 32-bit floating-point values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/float2x2
	MTLDataTypeFloat2x2 MTLDataType = 0
	// MTLDataTypeFloat2x3 - A 2x3 component matrix with 32-bit floating-point values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/float2x3
	MTLDataTypeFloat2x3 MTLDataType = 0
	// MTLDataTypeFloat2x4 - A 2x4 component matrix with 32-bit floating-point values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/float2x4
	MTLDataTypeFloat2x4 MTLDataType = 0
	// MTLDataTypeFloat3 - A three-component vector with 32-bit floating-point values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/float3
	MTLDataTypeFloat3 MTLDataType = 0
	// MTLDataTypeFloat3x2 - A 3x2 component matrix with 32-bit floating-point values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/float3x2
	MTLDataTypeFloat3x2 MTLDataType = 0
	// MTLDataTypeFloat3x3 - A 3x3 component matrix with 32-bit floating-point values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/float3x3
	MTLDataTypeFloat3x3 MTLDataType = 0
	// MTLDataTypeFloat3x4 - A 3x4 component matrix with 32-bit floating-point values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/float3x4
	MTLDataTypeFloat3x4 MTLDataType = 0
	// MTLDataTypeFloat4 - A four-component vector with 32-bit floating-point values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/float4
	MTLDataTypeFloat4 MTLDataType = 0
	// MTLDataTypeFloat4x2 - A 4x2 component matrix with 32-bit floating-point values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/float4x2
	MTLDataTypeFloat4x2 MTLDataType = 0
	// MTLDataTypeFloat4x3 - A 4x3 component matrix with 32-bit floating-point values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/float4x3
	MTLDataTypeFloat4x3 MTLDataType = 0
	// MTLDataTypeFloat4x4 - A 4x4 component matrix with 32-bit floating-point values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/float4x4
	MTLDataTypeFloat4x4 MTLDataType = 0
	// MTLDataTypeHalf - A 16-bit floating-point value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/half
	MTLDataTypeHalf MTLDataType = 0
	// MTLDataTypeHalf2 - A two-component vector with 16-bit floating-point values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/half2
	MTLDataTypeHalf2 MTLDataType = 0
	// MTLDataTypeHalf2x2 - A 2x2 component matrix with 16-bit floating-point values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/half2x2
	MTLDataTypeHalf2x2 MTLDataType = 0
	// MTLDataTypeHalf2x3 - A 2x3 component matrix with 16-bit floating-point values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/half2x3
	MTLDataTypeHalf2x3 MTLDataType = 0
	// MTLDataTypeHalf2x4 - A 2x4 component matrix with 16-bit floating-point values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/half2x4
	MTLDataTypeHalf2x4 MTLDataType = 0
	// MTLDataTypeHalf3 - A three-component vector with 16-bit floating-point values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/half3
	MTLDataTypeHalf3 MTLDataType = 0
	// MTLDataTypeHalf3x2 - A 3x2 component matrix with 16-bit floating-point values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/half3x2
	MTLDataTypeHalf3x2 MTLDataType = 0
	// MTLDataTypeHalf3x3 - A 3x3 component matrix with 16-bit floating-point values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/half3x3
	MTLDataTypeHalf3x3 MTLDataType = 0
	// MTLDataTypeHalf3x4 - A 3x4 component matrix with 16-bit floating-point values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/half3x4
	MTLDataTypeHalf3x4 MTLDataType = 0
	// MTLDataTypeHalf4 - A four-component vector with 16-bit floating-point values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/half4
	MTLDataTypeHalf4 MTLDataType = 0
	// MTLDataTypeHalf4x2 - A 4x2 component matrix with 16-bit floating-point values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/half4x2
	MTLDataTypeHalf4x2 MTLDataType = 0
	// MTLDataTypeHalf4x3 - A 4x3 component matrix with 16-bit floating-point values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/half4x3
	MTLDataTypeHalf4x3 MTLDataType = 0
	// MTLDataTypeHalf4x4 - A 4x4 component matrix with 16-bit floating-point values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/half4x4
	MTLDataTypeHalf4x4 MTLDataType = 0
	// MTLDataTypeIndirectCommandBuffer - An indirect command buffer resource instance.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/indirectCommandBuffer
	MTLDataTypeIndirectCommandBuffer MTLDataType = 0
	// MTLDataTypeInstanceAccelerationStructure - A high-level, ray-tracing acceleration structure for a set of low-level primitive instances.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/instanceAccelerationStructure
	MTLDataTypeInstanceAccelerationStructure MTLDataType = 0
	// MTLDataTypeInt - A 32-bit, signed integer value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/int
	MTLDataTypeInt MTLDataType = 0
	// MTLDataTypeInt2 - A two-component vector with 32-bit, signed integer values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/int2
	MTLDataTypeInt2 MTLDataType = 0
	// MTLDataTypeInt3 - A three-component vector with 32-bit, signed integer values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/int3
	MTLDataTypeInt3 MTLDataType = 0
	// MTLDataTypeInt4 - A four-component vector with 32-bit, signed integer values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/int4
	MTLDataTypeInt4 MTLDataType = 0
	// MTLDataTypeIntersectionFunctionTable - A table of intersection functions that a render or compute pipeline can call.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/intersectionFunctionTable
	MTLDataTypeIntersectionFunctionTable MTLDataType = 0
	// MTLDataTypeLong - A 64-bit, signed integer value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/long
	MTLDataTypeLong MTLDataType = 0
	// MTLDataTypeLong2 - A two-component vector with 64-bit, signed integer values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/long2
	MTLDataTypeLong2 MTLDataType = 0
	// MTLDataTypeLong3 - A three-component vector with 64-bit, signed integer values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/long3
	MTLDataTypeLong3 MTLDataType = 0
	// MTLDataTypeLong4 - A four-component vector with 64-bit, signed integer values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/long4
	MTLDataTypeLong4 MTLDataType = 0
	// MTLDataTypeNone - A placeholder that represents a GPU function parameter that doesn’t have a valid data type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/none
	MTLDataTypeNone MTLDataType = 0
	// MTLDataTypePointer - A pointer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/pointer
	MTLDataTypePointer MTLDataType = 0
	// MTLDataTypePrimitiveAccelerationStructure - A low-level ray-tracing acceleration structure for a set of primitives.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/primitiveAccelerationStructure
	MTLDataTypePrimitiveAccelerationStructure MTLDataType = 0
	// MTLDataTypeR16Snorm - An ordinary pixel with one component that’s a 16-bit, normalized, signed integer value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/r16Snorm
	MTLDataTypeR16Snorm MTLDataType = 0
	// MTLDataTypeR16Unorm - An ordinary pixel with one component that’s a 16-bit, normalized, unsigned integer value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/r16Unorm
	MTLDataTypeR16Unorm MTLDataType = 0
	// MTLDataTypeR8Snorm - An ordinary pixel with one component that’s an 8-bit, normalized, signed integer value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/r8Snorm
	MTLDataTypeR8Snorm MTLDataType = 0
	// MTLDataTypeR8Unorm - An ordinary pixel with one component that’s an 8-bit, normalized, unsigned integer value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/r8Unorm
	MTLDataTypeR8Unorm MTLDataType = 0
	// MTLDataTypeRenderPipeline - A Metal render pipeline instance.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/renderPipeline
	MTLDataTypeRenderPipeline MTLDataType = 0
	// MTLDataTypeRG11B10Float - A packed 32-bit format with three floating-point color components, two of which are 11-bit values, and one is a 10-bit value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/rg11b10Float
	MTLDataTypeRG11B10Float MTLDataType = 0
	// MTLDataTypeRG16Snorm - An ordinary pixel with two components, each of which is a 16-bit, normalized, signed integer value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/rg16Snorm
	MTLDataTypeRG16Snorm MTLDataType = 0
	// MTLDataTypeRG16Unorm - An ordinary pixel with two components, each of which is a 16-bit, normalized, unsigned integer value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/rg16Unorm
	MTLDataTypeRG16Unorm MTLDataType = 0
	// MTLDataTypeRG8Snorm - An ordinary pixel with two components, each of which is an 8-bit, normalized, signed integer value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/rg8Snorm
	MTLDataTypeRG8Snorm MTLDataType = 0
	// MTLDataTypeRG8Unorm - An ordinary pixel with two components, each of which is an 8-bit, normalized, unsigned integer value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/rg8Unorm
	MTLDataTypeRG8Unorm MTLDataType = 0
	// MTLDataTypeRGB10A2Unorm - A packed 32-bit format with three color components, each of which is a 10-bit, normalized, unsigned integer value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/rgb10a2Unorm
	MTLDataTypeRGB10A2Unorm MTLDataType = 0
	// MTLDataTypeRGB9E5Float - A packed 32-bit format with three color components, each of which is a 9-bit floating-point value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/rgb9e5Float
	MTLDataTypeRGB9E5Float MTLDataType = 0
	// MTLDataTypeRGBA16Snorm - An ordinary pixel with four components, each of which is a 16-bit, normalized, signed integer value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/rgba16Snorm
	MTLDataTypeRGBA16Snorm MTLDataType = 0
	// MTLDataTypeRGBA16Unorm - An ordinary pixel with four components, each of which is a 16-bit, normalized, unsigned integer value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/rgba16Unorm
	MTLDataTypeRGBA16Unorm MTLDataType = 0
	// MTLDataTypeRGBA8Snorm - An ordinary pixel with four components, each of which is an 8-bit, normalized, signed integer value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/rgba8Snorm
	MTLDataTypeRGBA8Snorm MTLDataType = 0
	// MTLDataTypeRGBA8Unorm - An ordinary pixel with four components, each of which is an 8-bit, normalized, unsigned integer value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/rgba8Unorm
	MTLDataTypeRGBA8Unorm MTLDataType = 0
	// MTLDataTypeRGBA8Unorm_sRGB - An ordinary pixel with four components, each of which is an 8-bit, normalized, unsigned integer value in the sRGB color space.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/rgba8Unorm_srgb
	MTLDataTypeRGBA8Unorm_sRGB MTLDataType = 0
	// MTLDataTypeSampler - A Metal texture sampler instance.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/sampler
	MTLDataTypeSampler MTLDataType = 0
	// MTLDataTypeShort - A 16-bit, signed integer value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/short
	MTLDataTypeShort MTLDataType = 0
	// MTLDataTypeShort2 - A two-component vector with 16-bit, signed integer values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/short2
	MTLDataTypeShort2 MTLDataType = 0
	// MTLDataTypeShort3 - A three-component vector with 16-bit, signed integer values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/short3
	MTLDataTypeShort3 MTLDataType = 0
	// MTLDataTypeShort4 - A four-component vector with 16-bit, signed integer values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/short4
	MTLDataTypeShort4 MTLDataType = 0
	// MTLDataTypeStruct - A structure instance.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/struct
	MTLDataTypeStruct MTLDataType = 0
	// MTLDataTypeTensor - Represents a data type corresponding to a machine learning tensor.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/tensor
	MTLDataTypeTensor MTLDataType = 0
	// MTLDataTypeTexture - A Metal texture resource instance.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/texture
	MTLDataTypeTexture MTLDataType = 0
	// MTLDataTypeUChar - An 8-bit, unsigned integer value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/uchar
	MTLDataTypeUChar MTLDataType = 0
	// MTLDataTypeUChar2 - A two-component vector with 8-bit, unsigned integer values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/uchar2
	MTLDataTypeUChar2 MTLDataType = 0
	// MTLDataTypeUChar3 - A three-component vector with 8-bit, unsigned integer values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/uchar3
	MTLDataTypeUChar3 MTLDataType = 0
	// MTLDataTypeUChar4 - A four-component vector with 8-bit, unsigned integer values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/uchar4
	MTLDataTypeUChar4 MTLDataType = 0
	// MTLDataTypeUInt - A 32-bit, unsigned integer value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/uint
	MTLDataTypeUInt MTLDataType = 0
	// MTLDataTypeUInt2 - A two-component vector with 32-bit, unsigned integer values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/uint2
	MTLDataTypeUInt2 MTLDataType = 0
	// MTLDataTypeUInt3 - A three-component vector with 32-bit, unsigned integer values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/uint3
	MTLDataTypeUInt3 MTLDataType = 0
	// MTLDataTypeUInt4 - A four-component vector with 32-bit, unsigned integer values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/uint4
	MTLDataTypeUInt4 MTLDataType = 0
	// MTLDataTypeULong - A 64-bit, unsigned integer value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/ulong
	MTLDataTypeULong MTLDataType = 0
	// MTLDataTypeULong2 - A two-component vector with 64-bit, unsigned integer values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/ulong2
	MTLDataTypeULong2 MTLDataType = 0
	// MTLDataTypeULong3 - A three-component vector with 64-bit, unsigned integer values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/ulong3
	MTLDataTypeULong3 MTLDataType = 0
	// MTLDataTypeULong4 - A four-component vector with 64-bit, unsigned integer values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/ulong4
	MTLDataTypeULong4 MTLDataType = 0
	// MTLDataTypeUShort - A 16-bit, unsigned integer value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/ushort
	MTLDataTypeUShort MTLDataType = 0
	// MTLDataTypeUShort2 - A two-component vector with 16-bit, unsigned integer values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/ushort2
	MTLDataTypeUShort2 MTLDataType = 0
	// MTLDataTypeUShort3 - A three-component vector with 16-bit, unsigned integer values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/ushort3
	MTLDataTypeUShort3 MTLDataType = 0
	// MTLDataTypeUShort4 - A four-component vector with 16-bit, unsigned integer values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/ushort4
	MTLDataTypeUShort4 MTLDataType = 0
	// MTLDataTypeVisibleFunctionTable - A table of visible functions that a render or compute pipeline can call.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDataType/visibleFunctionTable
	MTLDataTypeVisibleFunctionTable MTLDataType = 0
)

/* debug [enums.gen.go]: Processing enum MTLDepthClipMode (2 cases) */
// MTLDepthClipMode - The mode that determines how to deal with fragments outside of the near or far planes.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDepthClipMode
type MTLDepthClipMode uint

const (
	// MTLDepthClipModeClamp - Clamp fragments outside the near or far planes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDepthClipMode/clamp
	MTLDepthClipModeClamp MTLDepthClipMode = 0
	// MTLDepthClipModeClip - Clip fragments outside the near or far planes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDepthClipMode/clip
	MTLDepthClipModeClip MTLDepthClipMode = 0
)

/* debug [enums.gen.go]: Processing enum MTLDeviceLocation (4 cases) */
// MTLDeviceLocation - Indicates the location of the GPU relative to the system it’s connect to.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDeviceLocation
type MTLDeviceLocation uint

const (
	// MTLDeviceLocationBuiltIn - A location that indicates the GPU is permanently connected to the system internally.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDeviceLocation/builtIn
	MTLDeviceLocationBuiltIn MTLDeviceLocation = 0
	// MTLDeviceLocationExternal - A GPU location that indicates a person connected the GPU to the system with an external interface, such as Thunderbolt.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDeviceLocation/external
	MTLDeviceLocationExternal MTLDeviceLocation = 0
	// MTLDeviceLocationSlot - A GPU location that indicates a person connected the GPU to a system’s internal slot.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDeviceLocation/slot
	MTLDeviceLocationSlot MTLDeviceLocation = 0
	// MTLDeviceLocationUnspecified - A value that indicates the system can’t determine how the GPU connects to it.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDeviceLocation/unspecified
	MTLDeviceLocationUnspecified MTLDeviceLocation = 0
)

/* debug [enums.gen.go]: Processing enum MTLDispatchType (2 cases) */
// MTLDispatchType - The type of dispatch method to use when calling encoded functions.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDispatchType
type MTLDispatchType uint

const (
	// MTLDispatchTypeConcurrent - Sets a command encoder to dispatch encoded commands concurrently during your pass.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDispatchType/concurrent
	MTLDispatchTypeConcurrent MTLDispatchType = 0
	// MTLDispatchTypeSerial - Sets a command encoder to dispatch encoded commands serially during your pass.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDispatchType/serial
	MTLDispatchTypeSerial MTLDispatchType = 0
)

/* debug [enums.gen.go]: Processing enum MTLDynamicLibraryError (6 cases) */
// MTLDynamicLibraryError - Error codes that Metal can generate when creating dynamic libraries.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDynamicLibraryError-swift.struct/Code
type MTLDynamicLibraryError uint

const (
	// MTLDynamicLibraryErrorCompilationFailure - An error code that indicates Metal couldn’t compile a dynamic library.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDynamicLibraryError-swift.struct/Code/compilationFailure
	MTLDynamicLibraryErrorCompilationFailure MTLDynamicLibraryError = 0
	// MTLDynamicLibraryErrorDependencyLoadFailure - An error code that indicates a dynamic library couldn’t link to other dynamic libraries.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDynamicLibraryError-swift.struct/Code/dependencyLoadFailure
	MTLDynamicLibraryErrorDependencyLoadFailure MTLDynamicLibraryError = 0
	// MTLDynamicLibraryErrorInvalidFile - An error code that indicates an app is using an invalid reference to a library file, typically related to a URL.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDynamicLibraryError-swift.struct/Code/invalidFile
	MTLDynamicLibraryErrorInvalidFile MTLDynamicLibraryError = 0
	// MTLDynamicLibraryErrorNone - An error code that represents the absence of any problems.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDynamicLibraryError-swift.struct/Code/none
	MTLDynamicLibraryErrorNone MTLDynamicLibraryError = 0
	// MTLDynamicLibraryErrorUnresolvedInstallName - An error code that indicates Metal couldn’t resolve the installation name for a new dynamic library.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDynamicLibraryError-swift.struct/Code/unresolvedInstallName
	MTLDynamicLibraryErrorUnresolvedInstallName MTLDynamicLibraryError = 0
	// MTLDynamicLibraryErrorUnsupported - An error code that indicates the GPU device doesn’t support dynamic libraries.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDynamicLibraryError-swift.struct/Code/unsupported
	MTLDynamicLibraryErrorUnsupported MTLDynamicLibraryError = 0
)

/* debug [enums.gen.go]: Processing enum MTLFeatureSet (33 cases) */
// MTLFeatureSet - The device feature sets that define specific platform, hardware, and software configurations.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet
type MTLFeatureSet uint

const (
	// MTLFeatureSet_iOS_GPUFamily1_v1 - The GPU family 1, version 1 feature set for iOS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/iOS_GPUFamily1_v1
	MTLFeatureSet_iOS_GPUFamily1_v1 MTLFeatureSet = 0
	// MTLFeatureSet_iOS_GPUFamily1_v2 - The GPU family 1, version 2 feature set for iOS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/iOS_GPUFamily1_v2
	MTLFeatureSet_iOS_GPUFamily1_v2 MTLFeatureSet = 0
	// MTLFeatureSet_iOS_GPUFamily1_v3 - The GPU family 1, version 3 feature set for iOS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/iOS_GPUFamily1_v3
	MTLFeatureSet_iOS_GPUFamily1_v3 MTLFeatureSet = 0
	// MTLFeatureSet_iOS_GPUFamily1_v4 - The GPU family 1, version 4 feature set for iOS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/iOS_GPUFamily1_v4
	MTLFeatureSet_iOS_GPUFamily1_v4 MTLFeatureSet = 0
	// MTLFeatureSet_iOS_GPUFamily1_v5 - The GPU family 1, version 5 feature set for iOS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/iOS_GPUFamily1_v5
	MTLFeatureSet_iOS_GPUFamily1_v5 MTLFeatureSet = 0
	// MTLFeatureSet_iOS_GPUFamily2_v1 - The GPU family 2, version 1 feature set for iOS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/iOS_GPUFamily2_v1
	MTLFeatureSet_iOS_GPUFamily2_v1 MTLFeatureSet = 0
	// MTLFeatureSet_iOS_GPUFamily2_v2 - The GPU family 2, version 2 feature set for iOS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/iOS_GPUFamily2_v2
	MTLFeatureSet_iOS_GPUFamily2_v2 MTLFeatureSet = 0
	// MTLFeatureSet_iOS_GPUFamily2_v3 - The GPU family 2, version 3 feature set for iOS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/iOS_GPUFamily2_v3
	MTLFeatureSet_iOS_GPUFamily2_v3 MTLFeatureSet = 0
	// MTLFeatureSet_iOS_GPUFamily2_v4 - The GPU family 2, version 4 feature set for iOS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/iOS_GPUFamily2_v4
	MTLFeatureSet_iOS_GPUFamily2_v4 MTLFeatureSet = 0
	// MTLFeatureSet_iOS_GPUFamily2_v5 - The GPU family 2, version 5 feature set for iOS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/iOS_GPUFamily2_v5
	MTLFeatureSet_iOS_GPUFamily2_v5 MTLFeatureSet = 0
	// MTLFeatureSet_iOS_GPUFamily3_v1 - The GPU family 3, version 1 feature set for iOS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/iOS_GPUFamily3_v1
	MTLFeatureSet_iOS_GPUFamily3_v1 MTLFeatureSet = 0
	// MTLFeatureSet_iOS_GPUFamily3_v2 - The GPU family 3, version 2 feature set for iOS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/iOS_GPUFamily3_v2
	MTLFeatureSet_iOS_GPUFamily3_v2 MTLFeatureSet = 0
	// MTLFeatureSet_iOS_GPUFamily3_v3 - The GPU family 3, version 3 feature set for iOS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/iOS_GPUFamily3_v3
	MTLFeatureSet_iOS_GPUFamily3_v3 MTLFeatureSet = 0
	// MTLFeatureSet_iOS_GPUFamily3_v4 - The GPU family 3, version 4 feature set for iOS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/iOS_GPUFamily3_v4
	MTLFeatureSet_iOS_GPUFamily3_v4 MTLFeatureSet = 0
	// MTLFeatureSet_iOS_GPUFamily4_v1 - The GPU family 4, version 1 feature set for iOS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/iOS_GPUFamily4_v1
	MTLFeatureSet_iOS_GPUFamily4_v1 MTLFeatureSet = 0
	// MTLFeatureSet_iOS_GPUFamily4_v2 - The GPU family 4, version 2 feature set for iOS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/iOS_GPUFamily4_v2
	MTLFeatureSet_iOS_GPUFamily4_v2 MTLFeatureSet = 0
	// MTLFeatureSet_iOS_GPUFamily5_v1 - The GPU family 5, version 1 feature set for iOS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/iOS_GPUFamily5_v1
	MTLFeatureSet_iOS_GPUFamily5_v1 MTLFeatureSet = 0
	// MTLFeatureSet_macOS_GPUFamily1_v1 - The GPU family 1, version 1 feature set for macOS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/macOS_GPUFamily1_v1
	MTLFeatureSet_macOS_GPUFamily1_v1 MTLFeatureSet = 0
	// MTLFeatureSet_macOS_GPUFamily1_v2 - The GPU family 1, version 2 feature set for macOS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/macOS_GPUFamily1_v2
	MTLFeatureSet_macOS_GPUFamily1_v2 MTLFeatureSet = 0
	// MTLFeatureSet_macOS_GPUFamily1_v3 - The GPU family 1, version 3 feature set for macOS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/macOS_GPUFamily1_v3
	MTLFeatureSet_macOS_GPUFamily1_v3 MTLFeatureSet = 0
	// MTLFeatureSet_macOS_GPUFamily1_v4 - The GPU family 1, version 4 feature set for macOS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/macOS_GPUFamily1_v4
	MTLFeatureSet_macOS_GPUFamily1_v4 MTLFeatureSet = 0
	// MTLFeatureSet_macOS_GPUFamily2_v1 - The GPU family 2, version 1 feature set for macOS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/macOS_GPUFamily2_v1
	MTLFeatureSet_macOS_GPUFamily2_v1 MTLFeatureSet = 0
	// MTLFeatureSet_macOS_ReadWriteTextureTier2 - The read-write texture, tier 2 feature set for macOS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/macOS_ReadWriteTextureTier2
	MTLFeatureSet_macOS_ReadWriteTextureTier2 MTLFeatureSet = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/osx_GPUFamily1_v1
	MTLFeatureSet_OSX_GPUFamily1_v1 MTLFeatureSet = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/osx_GPUFamily1_v2
	MTLFeatureSet_OSX_GPUFamily1_v2 MTLFeatureSet = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/osx_ReadWriteTextureTier2
	MTLFeatureSet_OSX_ReadWriteTextureTier2 MTLFeatureSet = 0
	// MTLFeatureSet_tvOS_GPUFamily1_v1 - The GPU family 1, version 1 feature set for tvOS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/tvOS_GPUFamily1_v1-swift.enum.case
	MTLFeatureSet_tvOS_GPUFamily1_v1 MTLFeatureSet = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/tvos_GPUFamily1_v1-swift.type.property
	MTLFeatureSet_TVOS_GPUFamily1_v1 MTLFeatureSet = 0
	// MTLFeatureSet_tvOS_GPUFamily1_v2 - The GPU family 1, version 2 feature set for tvOS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/tvOS_GPUFamily1_v2
	MTLFeatureSet_tvOS_GPUFamily1_v2 MTLFeatureSet = 0
	// MTLFeatureSet_tvOS_GPUFamily1_v3 - The GPU family 1, version 3 feature set for tvOS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/tvOS_GPUFamily1_v3
	MTLFeatureSet_tvOS_GPUFamily1_v3 MTLFeatureSet = 0
	// MTLFeatureSet_tvOS_GPUFamily1_v4 - The GPU family 1, version 4 feature set for tvOS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/tvOS_GPUFamily1_v4
	MTLFeatureSet_tvOS_GPUFamily1_v4 MTLFeatureSet = 0
	// MTLFeatureSet_tvOS_GPUFamily2_v1 - The GPU family 2, version 1 feature set for tvOS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/tvOS_GPUFamily2_v1
	MTLFeatureSet_tvOS_GPUFamily2_v1 MTLFeatureSet = 0
	// MTLFeatureSet_tvOS_GPUFamily2_v2 - The GPU family 2, version 2 feature set for tvOS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/tvOS_GPUFamily2_v2
	MTLFeatureSet_tvOS_GPUFamily2_v2 MTLFeatureSet = 0
)

/* debug [enums.gen.go]: Processing enum MTLFunctionLogType (1 cases) */
// MTLFunctionLogType - Options for different kinds of function logs.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionLogType
type MTLFunctionLogType uint

const (
	// MTLFunctionLogTypeValidation - A message related to usage validation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionLogType/validation
	MTLFunctionLogTypeValidation MTLFunctionLogType = 0
)

/* debug [enums.gen.go]: Processing enum MTLFunctionOptions (6 cases) */
// MTLFunctionOptions - Options that define how Metal creates the function object.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionOptions
type MTLFunctionOptions uint

const (
	// MTLFunctionOptionCompileToBinary - An option that tells Metal to compile the function to a binary format for dynamic linking.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionOptions/compileToBinary
	MTLFunctionOptionCompileToBinary MTLFunctionOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionOptions/failOnBinaryArchiveMiss
	MTLFunctionOptionFailOnBinaryArchiveMiss MTLFunctionOptions = 0
	// MTLFunctionOptionNone - An option that specifies that Metal should use its default behavior when creating the function object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionOptions/MTLFunctionOptionNone
	MTLFunctionOptionNone MTLFunctionOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionOptions/pipelineIndependent
	MTLFunctionOptionPipelineIndependent MTLFunctionOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionOptions/storeFunctionInMetalPipelinesScript
	MTLFunctionOptionStoreFunctionInMetalPipelinesScript MTLFunctionOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionOptions/storeFunctionInMetalScript
	MTLFunctionOptionStoreFunctionInMetalScript MTLFunctionOptions = 0
)

/* debug [enums.gen.go]: Processing enum MTLFunctionType (7 cases) */
// MTLFunctionType - The type of a top-level Metal Shading Language (MSL) function.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionType
type MTLFunctionType uint

const (
	// MTLFunctionTypeFragment - A fragment function you can use in a render pipeline state object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionType/fragment
	MTLFunctionTypeFragment MTLFunctionType = 0
	// MTLFunctionTypeIntersection - A function you can use in an intersection function table.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionType/intersection
	MTLFunctionTypeIntersection MTLFunctionType = 0
	// MTLFunctionTypeKernel - A kernel you can use in a compute pipeline state object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionType/kernel
	MTLFunctionTypeKernel MTLFunctionType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionType/mesh
	MTLFunctionTypeMesh MTLFunctionType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionType/object
	MTLFunctionTypeObject MTLFunctionType = 0
	// MTLFunctionTypeVertex - A vertex function you can use in a render pipeline state object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionType/vertex
	MTLFunctionTypeVertex MTLFunctionType = 0
	// MTLFunctionTypeVisible - A function you can use in a visible function table.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionType/visible
	MTLFunctionTypeVisible MTLFunctionType = 0
)

/* debug [enums.gen.go]: Processing enum MTLGPUFamily (19 cases) */
// MTLGPUFamily - Represents the functionality for families of GPUs.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLGPUFamily
type MTLGPUFamily uint

const (
	// MTLGPUFamilyApple1 - Represents the Apple family 1 GPU features that correspond to the Apple A7 GPUs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLGPUFamily/apple1
	MTLGPUFamilyApple1 MTLGPUFamily = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLGPUFamily/apple10
	MTLGPUFamilyApple10 MTLGPUFamily = 0
	// MTLGPUFamilyApple2 - Represents the Apple family 2 GPU features that correspond to the Apple A8 GPUs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLGPUFamily/apple2
	MTLGPUFamilyApple2 MTLGPUFamily = 0
	// MTLGPUFamilyApple3 - Represents the Apple family 3 GPU features that correspond to the Apple A9 and A10 GPUs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLGPUFamily/apple3
	MTLGPUFamilyApple3 MTLGPUFamily = 0
	// MTLGPUFamilyApple4 - Represents the Apple family 4 GPU features that correspond to the Apple A11 GPUs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLGPUFamily/apple4
	MTLGPUFamilyApple4 MTLGPUFamily = 0
	// MTLGPUFamilyApple5 - Represents the Apple family 5 GPU features that correspond to the Apple A12 GPUs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLGPUFamily/apple5
	MTLGPUFamilyApple5 MTLGPUFamily = 0
	// MTLGPUFamilyApple6 - Represents the Apple family 6 GPU features that correspond to the Apple A13 GPUs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLGPUFamily/apple6
	MTLGPUFamilyApple6 MTLGPUFamily = 0
	// MTLGPUFamilyApple7 - Represents the Apple family 7 GPU features that correspond to the Apple A14 and M1 GPUs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLGPUFamily/apple7
	MTLGPUFamilyApple7 MTLGPUFamily = 0
	// MTLGPUFamilyApple8 - Represents the Apple family 8 GPU features that correspond to the Apple A15, A16, and M2 GPUs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLGPUFamily/apple8
	MTLGPUFamilyApple8 MTLGPUFamily = 0
	// MTLGPUFamilyApple9 - Represents the Apple family 9 GPU features that correspond to the Apple A17, M3, and M4 GPUs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLGPUFamily/apple9
	MTLGPUFamilyApple9 MTLGPUFamily = 0
	// MTLGPUFamilyCommon1 - Represents the Common family 1 GPU features.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLGPUFamily/common1
	MTLGPUFamilyCommon1 MTLGPUFamily = 0
	// MTLGPUFamilyCommon2 - Represents the Common family 2 GPU features.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLGPUFamily/common2
	MTLGPUFamilyCommon2 MTLGPUFamily = 0
	// MTLGPUFamilyCommon3 - Represents the Common family 3 GPU features.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLGPUFamily/common3
	MTLGPUFamilyCommon3 MTLGPUFamily = 0
	// MTLGPUFamilyMac1 - Represents the Mac family 1 GPU features.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLGPUFamily/mac1
	MTLGPUFamilyMac1 MTLGPUFamily = 0
	// MTLGPUFamilyMac2 - Represents the Mac family 2 GPU features.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLGPUFamily/mac2
	MTLGPUFamilyMac2 MTLGPUFamily = 0
	// MTLGPUFamilyMacCatalyst1 - Represents a family 1 Mac GPU when running an app you built with Mac Catalyst.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLGPUFamily/macCatalyst1
	MTLGPUFamilyMacCatalyst1 MTLGPUFamily = 0
	// MTLGPUFamilyMacCatalyst2 - Represents a family 2 Mac GPU when running an app you built with Mac Catalyst.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLGPUFamily/macCatalyst2
	MTLGPUFamilyMacCatalyst2 MTLGPUFamily = 0
	// MTLGPUFamilyMetal3 - Represents the Metal 3 features.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLGPUFamily/metal3
	MTLGPUFamilyMetal3 MTLGPUFamily = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLGPUFamily/metal4
	MTLGPUFamilyMetal4 MTLGPUFamily = 0
)

/* debug [enums.gen.go]: Processing enum MTLHazardTrackingMode (3 cases) */
// MTLHazardTrackingMode - The options you use to specify the hazard tracking mode.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLHazardTrackingMode
type MTLHazardTrackingMode uint

const (
	// MTLHazardTrackingModeDefault - An option specifying that the default tracking mode should be used.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLHazardTrackingMode/default
	MTLHazardTrackingModeDefault MTLHazardTrackingMode = 0
	// MTLHazardTrackingModeTracked - An option specifying that Metal prevents hazards when modifying this object’s contents.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLHazardTrackingMode/tracked
	MTLHazardTrackingModeTracked MTLHazardTrackingMode = 0
	// MTLHazardTrackingModeUntracked - An option specifying that the app must prevent hazards when modifying this object’s contents.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLHazardTrackingMode/untracked
	MTLHazardTrackingModeUntracked MTLHazardTrackingMode = 0
)

/* debug [enums.gen.go]: Processing enum MTLHeapType (3 cases) */
// MTLHeapType - The options you use to choose the heap type.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLHeapType
type MTLHeapType uint

const (
	// MTLHeapTypeAutomatic - A heap that automatically places new resource allocations.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLHeapType/automatic
	MTLHeapTypeAutomatic MTLHeapType = 0
	// MTLHeapTypePlacement - The app controls placement of resources on the heap.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLHeapType/placement
	MTLHeapTypePlacement MTLHeapType = 0
	// MTLHeapTypeSparse - The heap contains sparse texture tiles.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLHeapType/sparse
	MTLHeapTypeSparse MTLHeapType = 0
)

/* debug [enums.gen.go]: Processing enum MTLIndexType (2 cases) */
// MTLIndexType - The index type for an index buffer that references vertices of geometric primitives.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndexType
type MTLIndexType uint

const (
	// MTLIndexTypeUInt16 - A 16-bit unsigned integer used as a primitive index.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndexType/uint16
	MTLIndexTypeUInt16 MTLIndexType = 0
	// MTLIndexTypeUInt32 - A 32-bit unsigned integer used as a primitive index.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndexType/uint32
	MTLIndexTypeUInt32 MTLIndexType = 0
)

/* debug [enums.gen.go]: Processing enum MTLIndirectCommandType (8 cases) */
// MTLIndirectCommandType - The types of commands that you can encode into the indirect command buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandType
type MTLIndirectCommandType uint

const (
	// MTLIndirectCommandTypeConcurrentDispatch - A compute command using a grid aligned to threadgroup boundaries.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandType/concurrentDispatch
	MTLIndirectCommandTypeConcurrentDispatch MTLIndirectCommandType = 0
	// MTLIndirectCommandTypeConcurrentDispatchThreads - A compute command using an arbitrarily sized grid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandType/concurrentDispatchThreads
	MTLIndirectCommandTypeConcurrentDispatchThreads MTLIndirectCommandType = 0
	// MTLIndirectCommandTypeDraw - A draw call command.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandType/draw
	MTLIndirectCommandTypeDraw MTLIndirectCommandType = 0
	// MTLIndirectCommandTypeDrawIndexed - An indexed draw call command.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandType/drawIndexed
	MTLIndirectCommandTypeDrawIndexed MTLIndirectCommandType = 0
	// MTLIndirectCommandTypeDrawIndexedPatches - An indexed draw call command for tessellated patches.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandType/drawIndexedPatches
	MTLIndirectCommandTypeDrawIndexedPatches MTLIndirectCommandType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandType/drawMeshThreadgroups
	MTLIndirectCommandTypeDrawMeshThreadgroups MTLIndirectCommandType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandType/drawMeshThreads
	MTLIndirectCommandTypeDrawMeshThreads MTLIndirectCommandType = 0
	// MTLIndirectCommandTypeDrawPatches - A draw call command for tessellated patches.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandType/drawPatches
	MTLIndirectCommandTypeDrawPatches MTLIndirectCommandType = 0
)

/* debug [enums.gen.go]: Processing enum MTLIntersectionFunctionSignature (11 cases) */
// MTLIntersectionFunctionSignature - Constants for specifying different types of custom intersection functions.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionSignature
type MTLIntersectionFunctionSignature uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionSignature/curveData
	MTLIntersectionFunctionSignatureCurveData MTLIntersectionFunctionSignature = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionSignature/extendedLimits
	MTLIntersectionFunctionSignatureExtendedLimits MTLIntersectionFunctionSignature = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionSignature/instanceMotion
	MTLIntersectionFunctionSignatureInstanceMotion MTLIntersectionFunctionSignature = 0
	// MTLIntersectionFunctionSignatureInstancing - A flag indicating that function signature uses instancing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionSignature/instancing
	MTLIntersectionFunctionSignatureInstancing MTLIntersectionFunctionSignature = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionSignature/intersectionFunctionBuffer
	MTLIntersectionFunctionSignatureIntersectionFunctionBuffer MTLIntersectionFunctionSignature = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionSignature/maxLevels
	MTLIntersectionFunctionSignatureMaxLevels MTLIntersectionFunctionSignature = 0
	// MTLIntersectionFunctionSignatureNone - A constant indicating that the function uses the default signature.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionSignature/MTLIntersectionFunctionSignatureNone
	MTLIntersectionFunctionSignatureNone MTLIntersectionFunctionSignature = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionSignature/primitiveMotion
	MTLIntersectionFunctionSignaturePrimitiveMotion MTLIntersectionFunctionSignature = 0
	// MTLIntersectionFunctionSignatureTriangleData - A flag indicating that function signature uses triangle data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionSignature/triangleData
	MTLIntersectionFunctionSignatureTriangleData MTLIntersectionFunctionSignature = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionSignature/userData
	MTLIntersectionFunctionSignatureUserData MTLIntersectionFunctionSignature = 0
	// MTLIntersectionFunctionSignatureWorldSpaceData - A flag indicating that function signature uses world space data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionSignature/worldSpaceData
	MTLIntersectionFunctionSignatureWorldSpaceData MTLIntersectionFunctionSignature = 0
)

/* debug [enums.gen.go]: Processing enum MTLIOCommandQueueType (2 cases) */
// MTLIOCommandQueueType - Designates the queue type for a new input/output command queue.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOCommandQueueType
type MTLIOCommandQueueType uint

const (
	// MTLIOCommandQueueTypeConcurrent - Sets a new input/output command queue’s type to a queue that runs commands concurrently.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOCommandQueueType/concurrent
	MTLIOCommandQueueTypeConcurrent MTLIOCommandQueueType = 0
	// MTLIOCommandQueueTypeSerial - Sets a new input/output command queue’s type to a queue that runs commands serially.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOCommandQueueType/serial
	MTLIOCommandQueueTypeSerial MTLIOCommandQueueType = 0
)

/* debug [enums.gen.go]: Processing enum MTLIOCompressionMethod (5 cases) */
// MTLIOCompressionMethod - The compression codecs that Metal supports for input/output handles.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOCompressionMethod
type MTLIOCompressionMethod uint

const (
	// MTLIOCompressionMethodLZ4 - Indicates that a file uses the LZ4 compression algorithm codec.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOCompressionMethod/lz4
	MTLIOCompressionMethodLZ4 MTLIOCompressionMethod = 0
	// MTLIOCompressionMethodLZBitmap - Indicates that a file uses the LZBitmap compression algorithm codec.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOCompressionMethod/lzBitmap
	MTLIOCompressionMethodLZBitmap MTLIOCompressionMethod = 0
	// MTLIOCompressionMethodLZFSE - Indicates that a file uses the LZFSE compression algorithm codec.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOCompressionMethod/lzfse
	MTLIOCompressionMethodLZFSE MTLIOCompressionMethod = 0
	// MTLIOCompressionMethodLZMA - Indicates that a file uses the LZMA compression algorithm codec.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOCompressionMethod/lzma
	MTLIOCompressionMethodLZMA MTLIOCompressionMethod = 0
	// MTLIOCompressionMethodZlib - Indicates that a file uses the zlib compression algorithm codec.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOCompressionMethod/zlib
	MTLIOCompressionMethodZlib MTLIOCompressionMethod = 0
)

/* debug [enums.gen.go]: Processing enum MTLIOCompressionStatus (2 cases) */
// MTLIOCompressionStatus - Represents the final state of a compression context.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOCompressionStatus
type MTLIOCompressionStatus uint

const (
	// MTLIOCompressionStatusComplete - Indicates the compression API successfully flushed and destroyed a compression context.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOCompressionStatus/complete
	MTLIOCompressionStatusComplete MTLIOCompressionStatus = 0
	// MTLIOCompressionStatusError - Indicates the compression API had an error while flushing and destroying a compression context.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOCompressionStatus/error
	MTLIOCompressionStatusError MTLIOCompressionStatus = 0
)

/* debug [enums.gen.go]: Processing enum MTLIOError (2 cases) */
// MTLIOError - The error codes for creating an input/output file handle.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOError-swift.struct/Code
type MTLIOError uint

const (
	// MTLIOErrorInternal - An error code that represents a problem internal to the Metal framework.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOError-swift.struct/Code/internal
	MTLIOErrorInternal MTLIOError = 0
	// MTLIOErrorURLInvalid - An error code that represents a problem with a file URL.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOError-swift.struct/Code/urlInvalid
	MTLIOErrorURLInvalid MTLIOError = 0
)

/* debug [enums.gen.go]: Processing enum MTLIOPriority (3 cases) */
// MTLIOPriority - Designates the priority for a new input/output command queue.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOPriority
type MTLIOPriority uint

const (
	// MTLIOPriorityHigh - Sets a new input/output command queue’s priority to a high priority.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOPriority/high
	MTLIOPriorityHigh MTLIOPriority = 0
	// MTLIOPriorityLow - Designates the low priority for a new input/output command queue.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOPriority/low
	MTLIOPriorityLow MTLIOPriority = 0
	// MTLIOPriorityNormal - Designates the normal priority for a new input/output command queue.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOPriority/normal
	MTLIOPriorityNormal MTLIOPriority = 0
)

/* debug [enums.gen.go]: Processing enum MTLIOStatus (4 cases) */
// MTLIOStatus - Represents the state of an input/output command buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOStatus
type MTLIOStatus uint

const (
	// MTLIOStatusCancelled - Indicates the GPU has successfully abandoned the input/output command buffer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOStatus/cancelled
	MTLIOStatusCancelled MTLIOStatus = 0
	// MTLIOStatusComplete - Indicates the GPU has successfully finished executing the input/output command buffer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOStatus/complete
	MTLIOStatusComplete MTLIOStatus = 0
	// MTLIOStatusError - Indicates the GPU experienced a problem with the input/output command buffer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOStatus/error
	MTLIOStatusError MTLIOStatus = 0
	// MTLIOStatusPending - Indicates the GPU hasn’t finished executing the input/output command buffer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOStatus/pending
	MTLIOStatusPending MTLIOStatus = 0
)

/* debug [enums.gen.go]: Processing enum MTLLanguageVersion (12 cases) */
// MTLLanguageVersion - Metal shading language versions.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLanguageVersion
type MTLLanguageVersion uint

const (
	// MTLLanguageVersion1_0 - Version 1.0 of the Metal shading language.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLanguageVersion/version1_0
	MTLLanguageVersion1_0 MTLLanguageVersion = 0
	// MTLLanguageVersion1_1 - Version 1.1 of the Metal shading language.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLanguageVersion/version1_1
	MTLLanguageVersion1_1 MTLLanguageVersion = 0
	// MTLLanguageVersion1_2 - Version 1.2 of the Metal shading language.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLanguageVersion/version1_2
	MTLLanguageVersion1_2 MTLLanguageVersion = 0
	// MTLLanguageVersion2_0 - Version 2.0 of the Metal shading language.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLanguageVersion/version2_0
	MTLLanguageVersion2_0 MTLLanguageVersion = 0
	// MTLLanguageVersion2_1 - Version 2.1 of the Metal shading language.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLanguageVersion/version2_1
	MTLLanguageVersion2_1 MTLLanguageVersion = 0
	// MTLLanguageVersion2_2 - Version 2.2 of the Metal shading language.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLanguageVersion/version2_2
	MTLLanguageVersion2_2 MTLLanguageVersion = 0
	// MTLLanguageVersion2_3 - Version 2.3 of the Metal shading language.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLanguageVersion/version2_3
	MTLLanguageVersion2_3 MTLLanguageVersion = 0
	// MTLLanguageVersion2_4 - Version 2.4 of the Metal shading language.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLanguageVersion/version2_4
	MTLLanguageVersion2_4 MTLLanguageVersion = 0
	// MTLLanguageVersion3_0 - Version 3.0 of the Metal shading language.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLanguageVersion/version3_0
	MTLLanguageVersion3_0 MTLLanguageVersion = 0
	// MTLLanguageVersion3_1 - Version 3.1 of the Metal shading language.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLanguageVersion/version3_1
	MTLLanguageVersion3_1 MTLLanguageVersion = 0
	// MTLLanguageVersion3_2 - Version 3.2 of the Metal shading language.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLanguageVersion/version3_2
	MTLLanguageVersion3_2 MTLLanguageVersion = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLanguageVersion/version4_0
	MTLLanguageVersion4_0 MTLLanguageVersion = 0
)

/* debug [enums.gen.go]: Processing enum MTLLibraryError (6 cases) */
// MTLLibraryError - Error codes for Metal library errors.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLibraryError-swift.struct/Code
type MTLLibraryError uint

const (
	// MTLLibraryErrorCompileFailure - The library or function failed to compile.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLibraryError-swift.struct/Code/compileFailure
	MTLLibraryErrorCompileFailure MTLLibraryError = 0
	// MTLLibraryErrorCompileWarning - The library or function compiled successfully but generated warnings.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLibraryError-swift.struct/Code/compileWarning
	MTLLibraryErrorCompileWarning MTLLibraryError = 0
	// MTLLibraryErrorFileNotFound - Metal couldn’t find the Metal source file.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLibraryError-swift.struct/Code/fileNotFound
	MTLLibraryErrorFileNotFound MTLLibraryError = 0
	// MTLLibraryErrorFunctionNotFound - Metal couldn’t find the specified Metal function.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLibraryError-swift.struct/Code/functionNotFound
	MTLLibraryErrorFunctionNotFound MTLLibraryError = 0
	// MTLLibraryErrorInternal - The action caused an internal error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLibraryError-swift.struct/Code/internal
	MTLLibraryErrorInternal MTLLibraryError = 0
	// MTLLibraryErrorUnsupported - Metal couldn’t support the requested action.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLibraryError-swift.struct/Code/unsupported
	MTLLibraryErrorUnsupported MTLLibraryError = 0
)

/* debug [enums.gen.go]: Processing enum MTLLibraryOptimizationLevel (2 cases) */
// MTLLibraryOptimizationLevel - The optimization options for the Metal compiler.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLibraryOptimizationLevel
type MTLLibraryOptimizationLevel uint

const (
	// MTLLibraryOptimizationLevelDefault - An optimization option for the Metal compiler that prioritizes runtime performance.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLibraryOptimizationLevel/default
	MTLLibraryOptimizationLevelDefault MTLLibraryOptimizationLevel = 0
	// MTLLibraryOptimizationLevelSize - An optimization option for the Metal compiler that prioritizes minimizing the size of its output binaries, which may also reduce compile time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLibraryOptimizationLevel/size
	MTLLibraryOptimizationLevelSize MTLLibraryOptimizationLevel = 0
)

/* debug [enums.gen.go]: Processing enum MTLLibraryType (2 cases) */
// MTLLibraryType - A set of options for Metal library types.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLibraryType
type MTLLibraryType uint

const (
	// MTLLibraryTypeDynamic - A library that you can dynamically link to from other libraries.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLibraryType/dynamic
	MTLLibraryTypeDynamic MTLLibraryType = 0
	// MTLLibraryTypeExecutable - A library that can create pipeline state objects.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLibraryType/executable
	MTLLibraryTypeExecutable MTLLibraryType = 0
)

/* debug [enums.gen.go]: Processing enum MTLLoadAction (3 cases) */
// MTLLoadAction - Types of actions performed for an attachment at the start of a rendering pass.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLoadAction
type MTLLoadAction uint

const (
	// MTLLoadActionClear - The GPU writes a value to every pixel in the attachment at the start of the render pass.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLoadAction/clear
	MTLLoadActionClear MTLLoadAction = 0
	// MTLLoadActionDontCare - The GPU has permission to discard the existing contents of the attachment at the start of the render pass, replacing them with arbitrary data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLoadAction/dontCare
	MTLLoadActionDontCare MTLLoadAction = 0
	// MTLLoadActionLoad - The GPU preserves the existing contents of the attachment at the start of the render pass.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLoadAction/load
	MTLLoadActionLoad MTLLoadAction = 0
)

/* debug [enums.gen.go]: Processing enum MTLLogLevel (6 cases) */
// MTLLogLevel - The supported log levels for shader logging.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLogLevel
type MTLLogLevel uint

const (
	// MTLLogLevelDebug - The log level that captures diagnostic information.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLogLevel/debug
	MTLLogLevelDebug MTLLogLevel = 0
	// MTLLogLevelError - The log level that captures error information.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLogLevel/error
	MTLLogLevelError MTLLogLevel = 0
	// MTLLogLevelFault - The log level that captures fault information.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLogLevel/fault
	MTLLogLevelFault MTLLogLevel = 0
	// MTLLogLevelInfo - The log level that captures additional information.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLogLevel/info
	MTLLogLevelInfo MTLLogLevel = 0
	// MTLLogLevelNotice - The log level that captures notifications.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLogLevel/notice
	MTLLogLevelNotice MTLLogLevel = 0
	// MTLLogLevelUndefined - The log level when the log level hasn’t been configured.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLogLevel/undefined
	MTLLogLevelUndefined MTLLogLevel = 0
)

/* debug [enums.gen.go]: Processing enum MTLLogStateError (2 cases) */
// MTLLogStateError enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLogStateError
type MTLLogStateError uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLogStateError/invalid
	MTLLogStateErrorInvalid MTLLogStateError = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLogStateError/invalidSize
	MTLLogStateErrorInvalidSize MTLLogStateError = 0
)

/* debug [enums.gen.go]: Processing enum MTLMathFloatingPointFunctions (2 cases) */
// MTLMathFloatingPointFunctions - Indicates which FP32 math functions Metal uses.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMathFloatingPointFunctions
type MTLMathFloatingPointFunctions uint

const (
	// MTLMathFloatingPointFunctionsFast - An indication that Metal uses the fast version of the 32b floating-point math functions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMathFloatingPointFunctions/fast
	MTLMathFloatingPointFunctionsFast MTLMathFloatingPointFunctions = 0
	// MTLMathFloatingPointFunctionsPrecise - An indication that Metal uses the precise version of the 32b floating-point math functions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMathFloatingPointFunctions/precise
	MTLMathFloatingPointFunctionsPrecise MTLMathFloatingPointFunctions = 0
)

/* debug [enums.gen.go]: Processing enum MTLMathMode (3 cases) */
// MTLMathMode - An indication of whether the compiler can perform optimizations for floating-point arithmetic that may violate the IEEE 754 standard.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMathMode
type MTLMathMode uint

const (
	// MTLMathModeFast - An indicator of the mode the compiler uses to make aggressive, potentially lossy assumptions about floating-point math.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMathMode/fast
	MTLMathModeFast MTLMathMode = 0
	// MTLMathModeRelaxed - An indicator of the mode the compiler uses to make aggressive, potentially lossy assumptions about floating-point math, while honoring Inf/NaN.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMathMode/relaxed
	MTLMathModeRelaxed MTLMathMode = 0
	// MTLMathModeSafe - An indicator of the mode the compiler uses to disable unsafe floating-point optimizations by preventing the compiler from making any transformations that could affect the results.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMathMode/safe
	MTLMathModeSafe MTLMathMode = 0
)

/* debug [enums.gen.go]: Processing enum MTLMatrixLayout (2 cases) */
// MTLMatrixLayout enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMatrixLayout
type MTLMatrixLayout uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMatrixLayout/columnMajor
	MTLMatrixLayoutColumnMajor MTLMatrixLayout = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMatrixLayout/rowMajor
	MTLMatrixLayoutRowMajor MTLMatrixLayout = 0
)

/* debug [enums.gen.go]: Processing enum MTLMotionBorderMode (2 cases) */
// MTLMotionBorderMode - Options for specifying how the acceleration structure handles timestamps that are outside the specified range.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMotionBorderMode
type MTLMotionBorderMode uint

const (
	// MTLMotionBorderModeClamp - A mode that specifies treating times outside the specified endpoint as if they were at the endpoint.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMotionBorderMode/clamp
	MTLMotionBorderModeClamp MTLMotionBorderMode = 0
	// MTLMotionBorderModeVanish - A mode that specifies that times outside the specified endpoint need to prevent any ray-intersections with the primitive.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMotionBorderMode/vanish
	MTLMotionBorderModeVanish MTLMotionBorderMode = 0
)

/* debug [enums.gen.go]: Processing enum MTLMultisampleDepthResolveFilter (3 cases) */
// MTLMultisampleDepthResolveFilter - Filtering options for controlling an MSAA depth resolve operation.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMultisampleDepthResolveFilter
type MTLMultisampleDepthResolveFilter uint

const (
	// MTLMultisampleDepthResolveFilterMax - The GPU compares all depth samples in the pixel and selects the sample with the largest value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMultisampleDepthResolveFilter/max
	MTLMultisampleDepthResolveFilterMax MTLMultisampleDepthResolveFilter = 0
	// MTLMultisampleDepthResolveFilterMin - The GPU compares all depth samples in the pixel and selects the sample with the smallest value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMultisampleDepthResolveFilter/min
	MTLMultisampleDepthResolveFilterMin MTLMultisampleDepthResolveFilter = 0
	// MTLMultisampleDepthResolveFilterSample0 - No filter is applied.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMultisampleDepthResolveFilter/sample0
	MTLMultisampleDepthResolveFilterSample0 MTLMultisampleDepthResolveFilter = 0
)

/* debug [enums.gen.go]: Processing enum MTLMultisampleStencilResolveFilter (2 cases) */
// MTLMultisampleStencilResolveFilter - Constants used to control the multisample stencil resolve operation.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMultisampleStencilResolveFilter
type MTLMultisampleStencilResolveFilter uint

const (
	// MTLMultisampleStencilResolveFilterDepthResolvedSample - Chooses the stencil sample corresponding to the depth sample selected by the depth resolve filter.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMultisampleStencilResolveFilter/depthResolvedSample
	MTLMultisampleStencilResolveFilterDepthResolvedSample MTLMultisampleStencilResolveFilter = 0
	// MTLMultisampleStencilResolveFilterSample0 - Chooses the first stencil sample in the pixel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMultisampleStencilResolveFilter/sample0
	MTLMultisampleStencilResolveFilterSample0 MTLMultisampleStencilResolveFilter = 0
)

/* debug [enums.gen.go]: Processing enum MTLMutability (3 cases) */
// MTLMutability - The options that determine the mutability of a buffer’s contents.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMutability
type MTLMutability uint

const (
	// MTLMutabilityDefault - The default behavior, based on the buffer’s type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMutability/default
	MTLMutabilityDefault MTLMutability = 0
	// MTLMutabilityImmutable - An option that states that you can’t modify the buffer’s contents.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMutability/immutable
	MTLMutabilityImmutable MTLMutability = 0
	// MTLMutabilityMutable - An option that states that you can modify the buffer’s contents.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMutability/mutable
	MTLMutabilityMutable MTLMutability = 0
)

/* debug [enums.gen.go]: Processing enum MTLPatchType (3 cases) */
// MTLPatchType - Types of tessellation patches that can be inputs of a post-tessellation vertex function.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPatchType
type MTLPatchType uint

const (
	// MTLPatchTypeNone - An option that indicates that this isn’t a post-tessellation vertex function.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPatchType/none
	MTLPatchTypeNone MTLPatchType = 0
	// MTLPatchTypeQuad - A quad patch.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPatchType/quad
	MTLPatchTypeQuad MTLPatchType = 0
	// MTLPatchTypeTriangle - A triangle patch.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPatchType/triangle
	MTLPatchTypeTriangle MTLPatchType = 0
)

/* debug [enums.gen.go]: Processing enum MTLPipelineOption (5 cases) */
// MTLPipelineOption - Options that determine how Metal prepares the pipeline.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPipelineOption
type MTLPipelineOption uint

const (
	// MTLPipelineOptionArgumentInfo - An option instance that provides argument information for textures and threadgroup memory.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPipelineOption/argumentInfo
	MTLPipelineOptionArgumentInfo MTLPipelineOption = 0
	// MTLPipelineOptionBindingInfo - An option that provides binding information for pipeline state resources.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPipelineOption/bindingInfo
	MTLPipelineOptionBindingInfo MTLPipelineOption = 0
	// MTLPipelineOptionBufferTypeInfo - An option instance that provides detailed buffer type information for buffer arguments.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPipelineOption/bufferTypeInfo
	MTLPipelineOptionBufferTypeInfo MTLPipelineOption = 0
	// MTLPipelineOptionFailOnBinaryArchiveMiss - An option that specifies that Metal only creates the pipeline state object if the compiled shader is present inside a linked binary archive.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPipelineOption/failOnBinaryArchiveMiss
	MTLPipelineOptionFailOnBinaryArchiveMiss MTLPipelineOption = 0
	// MTLPipelineOptionNone - Don’t provide any reflection information.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPipelineOption/MTLPipelineOptionNone
	MTLPipelineOptionNone MTLPipelineOption = 0
)

/* debug [enums.gen.go]: Processing enum MTLPixelFormat (140 cases) */
// MTLPixelFormat - The data formats that describe the organization and characteristics of individual pixels in a texture.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat
type MTLPixelFormat uint

const (
	// MTLPixelFormatA1BGR5Unorm - Packed 16-bit format with normalized unsigned integer color components: 5 bits each for BGR and 1 for alpha, packed into 16 bits.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/a1bgr5Unorm
	MTLPixelFormatA1BGR5Unorm MTLPixelFormat = 0
	// MTLPixelFormatA8Unorm - Ordinary format with one 8-bit normalized unsigned integer component.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/a8Unorm
	MTLPixelFormatA8Unorm MTLPixelFormat = 0
	// MTLPixelFormatABGR4Unorm - Packed 16-bit format with normalized unsigned integer color components: 4 bits each for ABGR, packed into 16 bits.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/abgr4Unorm
	MTLPixelFormatABGR4Unorm MTLPixelFormat = 0
	// MTLPixelFormatASTC_10x10_HDR - ASTC-compressed format with high-dynamic range content, a block width of 10, and a block height of 10.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/astc_10x10_hdr
	MTLPixelFormatASTC_10x10_HDR MTLPixelFormat = 0
	// MTLPixelFormatASTC_10x10_LDR - ASTC-compressed format with low-dynamic-range content, a block width of 10, and a block height of 10.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/astc_10x10_ldr
	MTLPixelFormatASTC_10x10_LDR MTLPixelFormat = 0
	// MTLPixelFormatASTC_10x10_sRGB - ASTC-compressed format with low-dynamic-range content, conversion between sRGB and linear space, a block width of 10, and a block height of 10.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/astc_10x10_srgb
	MTLPixelFormatASTC_10x10_sRGB MTLPixelFormat = 0
	// MTLPixelFormatASTC_10x5_HDR - ASTC-compressed format with high-dynamic range content, a block width of 10, and a block height of 5.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/astc_10x5_hdr
	MTLPixelFormatASTC_10x5_HDR MTLPixelFormat = 0
	// MTLPixelFormatASTC_10x5_LDR - ASTC-compressed format with low-dynamic-range content, a block width of 10, and a block height of 5.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/astc_10x5_ldr
	MTLPixelFormatASTC_10x5_LDR MTLPixelFormat = 0
	// MTLPixelFormatASTC_10x5_sRGB - ASTC-compressed format with low-dynamic-range content, conversion between sRGB and linear space, a block width of 10, and a block height of 5.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/astc_10x5_srgb
	MTLPixelFormatASTC_10x5_sRGB MTLPixelFormat = 0
	// MTLPixelFormatASTC_10x6_HDR - ASTC-compressed format with high-dynamic range content, a block width of 10, and a block height of 6.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/astc_10x6_hdr
	MTLPixelFormatASTC_10x6_HDR MTLPixelFormat = 0
	// MTLPixelFormatASTC_10x6_LDR - ASTC-compressed format with low-dynamic-range content, a block width of 10, and a block height of 6.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/astc_10x6_ldr
	MTLPixelFormatASTC_10x6_LDR MTLPixelFormat = 0
	// MTLPixelFormatASTC_10x6_sRGB - ASTC-compressed format with low-dynamic-range content, conversion between sRGB and linear space, a block width of 10, and a block height of 6.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/astc_10x6_srgb
	MTLPixelFormatASTC_10x6_sRGB MTLPixelFormat = 0
	// MTLPixelFormatASTC_10x8_HDR - ASTC-compressed format with high-dynamic range content, a block width of 10, and a block height of 8.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/astc_10x8_hdr
	MTLPixelFormatASTC_10x8_HDR MTLPixelFormat = 0
	// MTLPixelFormatASTC_10x8_LDR - ASTC-compressed format with low-dynamic-range content, a block width of 10, and a block height of 8.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/astc_10x8_ldr
	MTLPixelFormatASTC_10x8_LDR MTLPixelFormat = 0
	// MTLPixelFormatASTC_10x8_sRGB - ASTC-compressed format with low-dynamic-range content, conversion between sRGB and linear space, a block width of 10, and a block height of 8.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/astc_10x8_srgb
	MTLPixelFormatASTC_10x8_sRGB MTLPixelFormat = 0
	// MTLPixelFormatASTC_12x10_HDR - ASTC-compressed format with high-dynamic range content, a block width of 12, and a block height of 10.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/astc_12x10_hdr
	MTLPixelFormatASTC_12x10_HDR MTLPixelFormat = 0
	// MTLPixelFormatASTC_12x10_LDR - ASTC-compressed format with low-dynamic-range content, a block width of 12, and a block height of 10.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/astc_12x10_ldr
	MTLPixelFormatASTC_12x10_LDR MTLPixelFormat = 0
	// MTLPixelFormatASTC_12x10_sRGB - ASTC-compressed format with low-dynamic-range content, conversion between sRGB and linear space, a block width of 12, and a block height of 10.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/astc_12x10_srgb
	MTLPixelFormatASTC_12x10_sRGB MTLPixelFormat = 0
	// MTLPixelFormatASTC_12x12_HDR - ASTC-compressed format with high-dynamic range content, a block width of 12, and a block height of 12.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/astc_12x12_hdr
	MTLPixelFormatASTC_12x12_HDR MTLPixelFormat = 0
	// MTLPixelFormatASTC_12x12_LDR - ASTC-compressed format with low-dynamic-range content, a block width of 12, and a block height of 12.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/astc_12x12_ldr
	MTLPixelFormatASTC_12x12_LDR MTLPixelFormat = 0
	// MTLPixelFormatASTC_12x12_sRGB - ASTC-compressed format with low-dynamic-range content, conversion between sRGB and linear space, a block width of 12, and a block height of 12.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/astc_12x12_srgb
	MTLPixelFormatASTC_12x12_sRGB MTLPixelFormat = 0
	// MTLPixelFormatASTC_4x4_HDR - ASTC-compressed format with high-dynamic-range content, a block width of 4, and a block height of 4.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/astc_4x4_hdr
	MTLPixelFormatASTC_4x4_HDR MTLPixelFormat = 0
	// MTLPixelFormatASTC_4x4_LDR - ASTC-compressed format with low-dynamic-range content, a block width of 4, and a block height of 4.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/astc_4x4_ldr
	MTLPixelFormatASTC_4x4_LDR MTLPixelFormat = 0
	// MTLPixelFormatASTC_4x4_sRGB - ASTC-compressed format with low-dynamic-range content, conversion between sRGB and linear space, a block width of 4, and a block height of 4.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/astc_4x4_srgb
	MTLPixelFormatASTC_4x4_sRGB MTLPixelFormat = 0
	// MTLPixelFormatASTC_5x4_HDR - ASTC-compressed format with high-dynamic range content, a block width of 5, and a block height of 4.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/astc_5x4_hdr
	MTLPixelFormatASTC_5x4_HDR MTLPixelFormat = 0
	// MTLPixelFormatASTC_5x4_LDR - ASTC-compressed format with low-dynamic-range content, a block width of 5, and a block height of 4.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/astc_5x4_ldr
	MTLPixelFormatASTC_5x4_LDR MTLPixelFormat = 0
	// MTLPixelFormatASTC_5x4_sRGB - ASTC-compressed format with low-dynamic-range content, conversion between sRGB and linear space, a block width of 5, and a block height of 4.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/astc_5x4_srgb
	MTLPixelFormatASTC_5x4_sRGB MTLPixelFormat = 0
	// MTLPixelFormatASTC_5x5_HDR - ASTC-compressed format with high-dynamic range content, a block width of 5, and a block height of 5.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/astc_5x5_hdr
	MTLPixelFormatASTC_5x5_HDR MTLPixelFormat = 0
	// MTLPixelFormatASTC_5x5_LDR - ASTC-compressed format with low-dynamic-range content, a block width of 5, and a block height of 5.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/astc_5x5_ldr
	MTLPixelFormatASTC_5x5_LDR MTLPixelFormat = 0
	// MTLPixelFormatASTC_5x5_sRGB - ASTC-compressed format with low-dynamic-range content, conversion between sRGB and linear space, a block width of 5, and a block height of 5.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/astc_5x5_srgb
	MTLPixelFormatASTC_5x5_sRGB MTLPixelFormat = 0
	// MTLPixelFormatASTC_6x5_HDR - ASTC-compressed format with high-dynamic range content, a block width of 6, and a block height of 5.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/astc_6x5_hdr
	MTLPixelFormatASTC_6x5_HDR MTLPixelFormat = 0
	// MTLPixelFormatASTC_6x5_LDR - ASTC-compressed format with low-dynamic-range content, a block width of 6, and a block height of 5.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/astc_6x5_ldr
	MTLPixelFormatASTC_6x5_LDR MTLPixelFormat = 0
	// MTLPixelFormatASTC_6x5_sRGB - ASTC-compressed format with low-dynamic-range content, conversion between sRGB and linear space, a block width of 6, and a block height of 5.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/astc_6x5_srgb
	MTLPixelFormatASTC_6x5_sRGB MTLPixelFormat = 0
	// MTLPixelFormatASTC_6x6_HDR - ASTC-compressed format with high-dynamic range content, a block width of 6, and a block height of 6.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/astc_6x6_hdr
	MTLPixelFormatASTC_6x6_HDR MTLPixelFormat = 0
	// MTLPixelFormatASTC_6x6_LDR - ASTC-compressed format with low-dynamic-range content, a block width of 6, and a block height of 6.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/astc_6x6_ldr
	MTLPixelFormatASTC_6x6_LDR MTLPixelFormat = 0
	// MTLPixelFormatASTC_6x6_sRGB - ASTC-compressed format with low-dynamic-range content, conversion between sRGB and linear space, a block width of 6, and a block height of 6.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/astc_6x6_srgb
	MTLPixelFormatASTC_6x6_sRGB MTLPixelFormat = 0
	// MTLPixelFormatASTC_8x5_HDR - ASTC-compressed format with high-dynamic range content, a block width of 8, and a block height of 5.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/astc_8x5_hdr
	MTLPixelFormatASTC_8x5_HDR MTLPixelFormat = 0
	// MTLPixelFormatASTC_8x5_LDR - ASTC-compressed format with low-dynamic-range content, a block width of 8, and a block height of 5.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/astc_8x5_ldr
	MTLPixelFormatASTC_8x5_LDR MTLPixelFormat = 0
	// MTLPixelFormatASTC_8x5_sRGB - ASTC-compressed format with low-dynamic-range content, conversion between sRGB and linear space, a block width of 8, and a block height of 5.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/astc_8x5_srgb
	MTLPixelFormatASTC_8x5_sRGB MTLPixelFormat = 0
	// MTLPixelFormatASTC_8x6_HDR - ASTC-compressed format with high-dynamic range content, a block width of 8, and a block height of 6.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/astc_8x6_hdr
	MTLPixelFormatASTC_8x6_HDR MTLPixelFormat = 0
	// MTLPixelFormatASTC_8x6_LDR - ASTC-compressed format with low-dynamic-range content, a block width of 8, and a block height of 6.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/astc_8x6_ldr
	MTLPixelFormatASTC_8x6_LDR MTLPixelFormat = 0
	// MTLPixelFormatASTC_8x6_sRGB - ASTC-compressed format with low-dynamic-range content, conversion between sRGB and linear space, a block width of 8, and a block height of 6.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/astc_8x6_srgb
	MTLPixelFormatASTC_8x6_sRGB MTLPixelFormat = 0
	// MTLPixelFormatASTC_8x8_HDR - ASTC-compressed format with high-dynamic range content, a block width of 8, and a block height of 8.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/astc_8x8_hdr
	MTLPixelFormatASTC_8x8_HDR MTLPixelFormat = 0
	// MTLPixelFormatASTC_8x8_LDR - ASTC-compressed format with low-dynamic-range content, a block width of 8, and a block height of 8.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/astc_8x8_ldr
	MTLPixelFormatASTC_8x8_LDR MTLPixelFormat = 0
	// MTLPixelFormatASTC_8x8_sRGB - ASTC-compressed format with low-dynamic-range content, conversion between sRGB and linear space, a block width of 8, and a block height of 8.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/astc_8x8_srgb
	MTLPixelFormatASTC_8x8_sRGB MTLPixelFormat = 0
	// MTLPixelFormatB5G6R5Unorm - Packed 16-bit format with normalized unsigned integer color components: 5 bits for blue, 6 bits for green, 5 bits for red, packed into 16 bits.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/b5g6r5Unorm
	MTLPixelFormatB5G6R5Unorm MTLPixelFormat = 0
	// MTLPixelFormatBC1_RGBA - Compressed format with two 16-bit color components and one 32-bit descriptor component.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/bc1_rgba
	MTLPixelFormatBC1_RGBA MTLPixelFormat = 0
	// MTLPixelFormatBC1_RGBA_sRGB - Compressed format with two 16-bit color components and one 32-bit descriptor component, with conversion between sRGB and linear space.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/bc1_rgba_srgb
	MTLPixelFormatBC1_RGBA_sRGB MTLPixelFormat = 0
	// MTLPixelFormatBC2_RGBA - Compressed format with two 64-bit chunks. The first chunk contains two 8-bit alpha components and one 48-bit descriptor component. The second chunk contains two 16-bit color components and one 32-bit descriptor component.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/bc2_rgba
	MTLPixelFormatBC2_RGBA MTLPixelFormat = 0
	// MTLPixelFormatBC2_RGBA_sRGB - Compressed format with two 64-bit chunks, with conversion between sRGB and linear space. The first chunk contains two 8-bit alpha components and one 48-bit descriptor component. The second chunk contains two 16-bit color components and one 32-bit descriptor component.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/bc2_rgba_srgb
	MTLPixelFormatBC2_RGBA_sRGB MTLPixelFormat = 0
	// MTLPixelFormatBC3_RGBA - Compressed format with two 64-bit chunks. The first chunk contains two 8-bit alpha components and one 48-bit descriptor component. The second chunk contains two 16-bit color components and one 32-bit descriptor component.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/bc3_rgba
	MTLPixelFormatBC3_RGBA MTLPixelFormat = 0
	// MTLPixelFormatBC3_RGBA_sRGB - Compressed format with two 64-bit chunks, with conversion between sRGB and linear space. The first chunk contains two 8-bit alpha components and one 48-bit descriptor component. The second chunk contains two 16-bit color components and one 32-bit descriptor component.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/bc3_rgba_srgb
	MTLPixelFormatBC3_RGBA_sRGB MTLPixelFormat = 0
	// MTLPixelFormatBC4_RSnorm - Compressed format with one normalized signed integer component.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/bc4_rSnorm
	MTLPixelFormatBC4_RSnorm MTLPixelFormat = 0
	// MTLPixelFormatBC4_RUnorm - Compressed format with one normalized unsigned integer component.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/bc4_rUnorm
	MTLPixelFormatBC4_RUnorm MTLPixelFormat = 0
	// MTLPixelFormatBC5_RGSnorm - Compressed format with two normalized signed integer components.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/bc5_rgSnorm
	MTLPixelFormatBC5_RGSnorm MTLPixelFormat = 0
	// MTLPixelFormatBC5_RGUnorm - Compressed format with two normalized unsigned integer components.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/bc5_rgUnorm
	MTLPixelFormatBC5_RGUnorm MTLPixelFormat = 0
	// MTLPixelFormatBC6H_RGBFloat - Compressed format with four floating-point components.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/bc6H_rgbFloat
	MTLPixelFormatBC6H_RGBFloat MTLPixelFormat = 0
	// MTLPixelFormatBC6H_RGBUfloat - Compressed format with four unsigned floating-point components.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/bc6H_rgbuFloat
	MTLPixelFormatBC6H_RGBUfloat MTLPixelFormat = 0
	// MTLPixelFormatBC7_RGBAUnorm - Compressed format with four normalized unsigned integer components.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/bc7_rgbaUnorm
	MTLPixelFormatBC7_RGBAUnorm MTLPixelFormat = 0
	// MTLPixelFormatBC7_RGBAUnorm_sRGB - Compressed format with four normalized unsigned integer components, with conversion between sRGB and linear space.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/bc7_rgbaUnorm_srgb
	MTLPixelFormatBC7_RGBAUnorm_sRGB MTLPixelFormat = 0
	// MTLPixelFormatBGR10_XR - A 32-bit extended-range pixel format with three fixed-point components of 10-bit blue, 10-bit green, and 10-bit red.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/bgr10_xr
	MTLPixelFormatBGR10_XR MTLPixelFormat = 0
	// MTLPixelFormatBGR10_XR_sRGB - A 32-bit extended-range pixel format with sRGB conversion and three fixed-point components of 10-bit blue, 10-bit green, and 10-bit red.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/bgr10_xr_srgb
	MTLPixelFormatBGR10_XR_sRGB MTLPixelFormat = 0
	// MTLPixelFormatBGR10A2Unorm - A 32-bit packed pixel format with four normalized unsigned integer components: 10-bit blue, 10-bit green, 10-bit red, and 2-bit alpha.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/bgr10a2Unorm
	MTLPixelFormatBGR10A2Unorm MTLPixelFormat = 0
	// MTLPixelFormatBGR5A1Unorm - Packed 16-bit format with normalized unsigned integer color components: 5 bits each for BGR and 1 for alpha, packed into 16 bits.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/bgr5A1Unorm
	MTLPixelFormatBGR5A1Unorm MTLPixelFormat = 0
	// MTLPixelFormatBGRA10_XR - A 64-bit extended-range pixel format with four fixed-point components of 10-bit blue, 10-bit green, 10-bit red, and 10-bit alpha.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/bgra10_xr
	MTLPixelFormatBGRA10_XR MTLPixelFormat = 0
	// MTLPixelFormatBGRA10_XR_sRGB - A 64-bit extended-range pixel format with sRGB conversion and four fixed-point components of 10-bit blue, 10-bit green, 10-bit red, and 10-bit alpha.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/bgra10_xr_srgb
	MTLPixelFormatBGRA10_XR_sRGB MTLPixelFormat = 0
	// MTLPixelFormatBGRA8Unorm - Ordinary format with four 8-bit normalized unsigned integer components in BGRA order.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/bgra8Unorm
	MTLPixelFormatBGRA8Unorm MTLPixelFormat = 0
	// MTLPixelFormatBGRA8Unorm_sRGB - Ordinary format with four 8-bit normalized unsigned integer components in BGRA order with conversion between sRGB and linear space.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/bgra8Unorm_srgb
	MTLPixelFormatBGRA8Unorm_sRGB MTLPixelFormat = 0
	// MTLPixelFormatBGRG422 - A pixel format where the red and green components are subsampled horizontally.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/bgrg422
	MTLPixelFormatBGRG422 MTLPixelFormat = 0
	// MTLPixelFormatDepth16Unorm - A pixel format for a depth-render target that has a 16-bit normalized, unsigned-integer component.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/depth16Unorm
	MTLPixelFormatDepth16Unorm MTLPixelFormat = 0
	// MTLPixelFormatDepth24Unorm_Stencil8 - A 32-bit combined depth and stencil pixel format with a 24-bit normalized unsigned integer for depth and an 8-bit unsigned integer for stencil.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/depth24Unorm_stencil8
	MTLPixelFormatDepth24Unorm_Stencil8 MTLPixelFormat = 0
	// MTLPixelFormatDepth32Float - A pixel format with one 32-bit floating-point component, used for a depth render target.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/depth32Float
	MTLPixelFormatDepth32Float MTLPixelFormat = 0
	// MTLPixelFormatDepth32Float_Stencil8 - A 40-bit combined depth and stencil pixel format with a 32-bit floating-point value for depth and an 8-bit unsigned integer for stencil.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/depth32Float_stencil8
	MTLPixelFormatDepth32Float_Stencil8 MTLPixelFormat = 0
	// MTLPixelFormatEAC_R11Snorm - Compressed format using EAC compression with one normalized signed integer component.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/eac_r11Snorm
	MTLPixelFormatEAC_R11Snorm MTLPixelFormat = 0
	// MTLPixelFormatEAC_R11Unorm - Compressed format using EAC compression with one normalized unsigned integer component.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/eac_r11Unorm
	MTLPixelFormatEAC_R11Unorm MTLPixelFormat = 0
	// MTLPixelFormatEAC_RG11Snorm - Compressed format using EAC compression with two normalized signed integer components.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/eac_rg11Snorm
	MTLPixelFormatEAC_RG11Snorm MTLPixelFormat = 0
	// MTLPixelFormatEAC_RG11Unorm - Compressed format using EAC compression with two normalized unsigned integer components.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/eac_rg11Unorm
	MTLPixelFormatEAC_RG11Unorm MTLPixelFormat = 0
	// MTLPixelFormatEAC_RGBA8 - Compressed format using EAC compression with four 8-bit components.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/eac_rgba8
	MTLPixelFormatEAC_RGBA8 MTLPixelFormat = 0
	// MTLPixelFormatEAC_RGBA8_sRGB - Compressed format using EAC compression with four 8-bit components with conversion between sRGB and linear space.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/eac_rgba8_srgb
	MTLPixelFormatEAC_RGBA8_sRGB MTLPixelFormat = 0
	// MTLPixelFormatETC2_RGB8 - Compressed format using ETC2 compression with three 8-bit components.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/etc2_rgb8
	MTLPixelFormatETC2_RGB8 MTLPixelFormat = 0
	// MTLPixelFormatETC2_RGB8_sRGB - Compressed format using ETC2 compression with three 8-bit components with conversion between sRGB and linear space.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/etc2_rgb8_srgb
	MTLPixelFormatETC2_RGB8_sRGB MTLPixelFormat = 0
	// MTLPixelFormatETC2_RGB8A1 - Compressed format using ETC2 compression with four 8-bit components.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/etc2_rgb8a1
	MTLPixelFormatETC2_RGB8A1 MTLPixelFormat = 0
	// MTLPixelFormatETC2_RGB8A1_sRGB - Compressed format using ETC2 compression with four 8-bit components with conversion between sRGB and linear space.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/etc2_rgb8a1_srgb
	MTLPixelFormatETC2_RGB8A1_sRGB MTLPixelFormat = 0
	// MTLPixelFormatGBGR422 - A pixel format where the red and green components are subsampled horizontally.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/gbgr422
	MTLPixelFormatGBGR422 MTLPixelFormat = 0
	// MTLPixelFormatInvalid - The default value of the pixel format for the  . You cannot create a texture with this value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/invalid
	MTLPixelFormatInvalid MTLPixelFormat = 0
	// MTLPixelFormatPVRTC_RGB_2BPP - Compressed format using PVRTC compression and 2bpp for RGB components.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/pvrtc_rgb_2bpp
	MTLPixelFormatPVRTC_RGB_2BPP MTLPixelFormat = 0
	// MTLPixelFormatPVRTC_RGB_2BPP_sRGB - Compressed format using PVRTC compression and 2bpp for RGB components with conversion between sRGB and linear space.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/pvrtc_rgb_2bpp_srgb
	MTLPixelFormatPVRTC_RGB_2BPP_sRGB MTLPixelFormat = 0
	// MTLPixelFormatPVRTC_RGB_4BPP - Compressed format using PVRTC compression and 4bpp for RGB components.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/pvrtc_rgb_4bpp
	MTLPixelFormatPVRTC_RGB_4BPP MTLPixelFormat = 0
	// MTLPixelFormatPVRTC_RGB_4BPP_sRGB - Compressed format using PVRTC compression and 4bpp for RGB components with conversion between sRGB and linear space.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/pvrtc_rgb_4bpp_srgb
	MTLPixelFormatPVRTC_RGB_4BPP_sRGB MTLPixelFormat = 0
	// MTLPixelFormatPVRTC_RGBA_2BPP - Compressed format using PVRTC compression and 2bpp for RGBA components.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/pvrtc_rgba_2bpp
	MTLPixelFormatPVRTC_RGBA_2BPP MTLPixelFormat = 0
	// MTLPixelFormatPVRTC_RGBA_2BPP_sRGB - Compressed format using PVRTC compression and 2bpp for RGBA components with conversion between sRGB and linear space.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/pvrtc_rgba_2bpp_srgb
	MTLPixelFormatPVRTC_RGBA_2BPP_sRGB MTLPixelFormat = 0
	// MTLPixelFormatPVRTC_RGBA_4BPP - Compressed format using PVRTC compression and 4bpp for RGBA components.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/pvrtc_rgba_4bpp
	MTLPixelFormatPVRTC_RGBA_4BPP MTLPixelFormat = 0
	// MTLPixelFormatPVRTC_RGBA_4BPP_sRGB - Compressed format using PVRTC compression and 4bpp for RGBA components with conversion between sRGB and linear space.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/pvrtc_rgba_4bpp_srgb
	MTLPixelFormatPVRTC_RGBA_4BPP_sRGB MTLPixelFormat = 0
	// MTLPixelFormatR16Float - Ordinary format with one 16-bit floating-point component.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/r16Float
	MTLPixelFormatR16Float MTLPixelFormat = 0
	// MTLPixelFormatR16Sint - Ordinary format with one 16-bit signed integer component.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/r16Sint
	MTLPixelFormatR16Sint MTLPixelFormat = 0
	// MTLPixelFormatR16Snorm - Ordinary format with one 16-bit normalized signed integer component.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/r16Snorm
	MTLPixelFormatR16Snorm MTLPixelFormat = 0
	// MTLPixelFormatR16Uint - Ordinary format with one 16-bit unsigned integer component.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/r16Uint
	MTLPixelFormatR16Uint MTLPixelFormat = 0
	// MTLPixelFormatR16Unorm - Ordinary format with one 16-bit normalized unsigned integer component.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/r16Unorm
	MTLPixelFormatR16Unorm MTLPixelFormat = 0
	// MTLPixelFormatR32Float - Ordinary format with one 32-bit floating-point component.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/r32Float
	MTLPixelFormatR32Float MTLPixelFormat = 0
	// MTLPixelFormatR32Sint - Ordinary format with one 32-bit signed integer component.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/r32Sint
	MTLPixelFormatR32Sint MTLPixelFormat = 0
	// MTLPixelFormatR32Uint - Ordinary format with one 32-bit unsigned integer component.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/r32Uint
	MTLPixelFormatR32Uint MTLPixelFormat = 0
	// MTLPixelFormatR8Sint - Ordinary format with one 8-bit signed integer component.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/r8Sint
	MTLPixelFormatR8Sint MTLPixelFormat = 0
	// MTLPixelFormatR8Snorm - Ordinary format with one 8-bit normalized signed integer component.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/r8Snorm
	MTLPixelFormatR8Snorm MTLPixelFormat = 0
	// MTLPixelFormatR8Uint - Ordinary format with one 8-bit unsigned integer component.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/r8Uint
	MTLPixelFormatR8Uint MTLPixelFormat = 0
	// MTLPixelFormatR8Unorm - Ordinary format with one 8-bit normalized unsigned integer component.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/r8Unorm
	MTLPixelFormatR8Unorm MTLPixelFormat = 0
	// MTLPixelFormatR8Unorm_sRGB - Ordinary format with one 8-bit normalized unsigned integer component with conversion between sRGB and linear space.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/r8Unorm_srgb
	MTLPixelFormatR8Unorm_sRGB MTLPixelFormat = 0
	// MTLPixelFormatRG11B10Float - 32-bit format with floating-point color components, 11 bits each for red and green and 10 bits for blue.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/rg11b10Float
	MTLPixelFormatRG11B10Float MTLPixelFormat = 0
	// MTLPixelFormatRG16Float - Ordinary format with two 16-bit floating-point components.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/rg16Float
	MTLPixelFormatRG16Float MTLPixelFormat = 0
	// MTLPixelFormatRG16Sint - Ordinary format with two 16-bit signed integer components.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/rg16Sint
	MTLPixelFormatRG16Sint MTLPixelFormat = 0
	// MTLPixelFormatRG16Snorm - Ordinary format with two 16-bit normalized signed integer components.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/rg16Snorm
	MTLPixelFormatRG16Snorm MTLPixelFormat = 0
	// MTLPixelFormatRG16Uint - Ordinary format with two 16-bit unsigned integer components.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/rg16Uint
	MTLPixelFormatRG16Uint MTLPixelFormat = 0
	// MTLPixelFormatRG16Unorm - Ordinary format with two 16-bit normalized unsigned integer components.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/rg16Unorm
	MTLPixelFormatRG16Unorm MTLPixelFormat = 0
	// MTLPixelFormatRG32Float - Ordinary format with two 32-bit floating-point components.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/rg32Float
	MTLPixelFormatRG32Float MTLPixelFormat = 0
	// MTLPixelFormatRG32Sint - Ordinary format with two 32-bit signed integer components.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/rg32Sint
	MTLPixelFormatRG32Sint MTLPixelFormat = 0
	// MTLPixelFormatRG32Uint - Ordinary format with two 32-bit unsigned integer components.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/rg32Uint
	MTLPixelFormatRG32Uint MTLPixelFormat = 0
	// MTLPixelFormatRG8Sint - Ordinary format with two 8-bit signed integer components.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/rg8Sint
	MTLPixelFormatRG8Sint MTLPixelFormat = 0
	// MTLPixelFormatRG8Snorm - Ordinary format with two 8-bit normalized signed integer components.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/rg8Snorm
	MTLPixelFormatRG8Snorm MTLPixelFormat = 0
	// MTLPixelFormatRG8Uint - Ordinary format with two 8-bit unsigned integer components.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/rg8Uint
	MTLPixelFormatRG8Uint MTLPixelFormat = 0
	// MTLPixelFormatRG8Unorm - Ordinary format with two 8-bit normalized unsigned integer components.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/rg8Unorm
	MTLPixelFormatRG8Unorm MTLPixelFormat = 0
	// MTLPixelFormatRG8Unorm_sRGB - Ordinary format with two 8-bit normalized unsigned integer components with conversion between sRGB and linear space.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/rg8Unorm_srgb
	MTLPixelFormatRG8Unorm_sRGB MTLPixelFormat = 0
	// MTLPixelFormatRGB10A2Uint - A 32-bit packed pixel format with four unsigned integer components: 10-bit red, 10-bit green, 10-bit blue, and 2-bit alpha.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/rgb10a2Uint
	MTLPixelFormatRGB10A2Uint MTLPixelFormat = 0
	// MTLPixelFormatRGB10A2Unorm - A 32-bit packed pixel format with four normalized unsigned integer components: 10-bit red, 10-bit green, 10-bit blue, and 2-bit alpha.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/rgb10a2Unorm
	MTLPixelFormatRGB10A2Unorm MTLPixelFormat = 0
	// MTLPixelFormatRGB9E5Float - Packed 32-bit format with floating-point color components: 9 bits each for RGB and 5 bits for an exponent shared by RGB, packed into 32 bits.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/rgb9e5Float
	MTLPixelFormatRGB9E5Float MTLPixelFormat = 0
	// MTLPixelFormatRGBA16Float - Ordinary format with four 16-bit floating-point components in RGBA order.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/rgba16Float
	MTLPixelFormatRGBA16Float MTLPixelFormat = 0
	// MTLPixelFormatRGBA16Sint - Ordinary format with four 16-bit signed integer components in RGBA order.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/rgba16Sint
	MTLPixelFormatRGBA16Sint MTLPixelFormat = 0
	// MTLPixelFormatRGBA16Snorm - Ordinary format with four 16-bit normalized signed integer components in RGBA order.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/rgba16Snorm
	MTLPixelFormatRGBA16Snorm MTLPixelFormat = 0
	// MTLPixelFormatRGBA16Uint - Ordinary format with four 16-bit unsigned integer components in RGBA order.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/rgba16Uint
	MTLPixelFormatRGBA16Uint MTLPixelFormat = 0
	// MTLPixelFormatRGBA16Unorm - Ordinary format with four 16-bit normalized unsigned integer components in RGBA order.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/rgba16Unorm
	MTLPixelFormatRGBA16Unorm MTLPixelFormat = 0
	// MTLPixelFormatRGBA32Float - Ordinary format with four 32-bit floating-point components in RGBA order.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/rgba32Float
	MTLPixelFormatRGBA32Float MTLPixelFormat = 0
	// MTLPixelFormatRGBA32Sint - Ordinary format with four 32-bit signed integer components in RGBA order.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/rgba32Sint
	MTLPixelFormatRGBA32Sint MTLPixelFormat = 0
	// MTLPixelFormatRGBA32Uint - Ordinary format with four 32-bit unsigned integer components in RGBA order.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/rgba32Uint
	MTLPixelFormatRGBA32Uint MTLPixelFormat = 0
	// MTLPixelFormatRGBA8Sint - Ordinary format with four 8-bit signed integer components in RGBA order.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/rgba8Sint
	MTLPixelFormatRGBA8Sint MTLPixelFormat = 0
	// MTLPixelFormatRGBA8Snorm - Ordinary format with four 8-bit normalized signed integer components in RGBA order.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/rgba8Snorm
	MTLPixelFormatRGBA8Snorm MTLPixelFormat = 0
	// MTLPixelFormatRGBA8Uint - Ordinary format with four 8-bit unsigned integer components in RGBA order.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/rgba8Uint
	MTLPixelFormatRGBA8Uint MTLPixelFormat = 0
	// MTLPixelFormatRGBA8Unorm - Ordinary format with four 8-bit normalized unsigned integer components in RGBA order.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/rgba8Unorm
	MTLPixelFormatRGBA8Unorm MTLPixelFormat = 0
	// MTLPixelFormatRGBA8Unorm_sRGB - Ordinary format with four 8-bit normalized unsigned integer components in RGBA order with conversion between sRGB and linear space.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/rgba8Unorm_srgb
	MTLPixelFormatRGBA8Unorm_sRGB MTLPixelFormat = 0
	// MTLPixelFormatStencil8 - A pixel format with an 8-bit unsigned integer component, used for a stencil render target.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/stencil8
	MTLPixelFormatStencil8 MTLPixelFormat = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/unspecialized
	MTLPixelFormatUnspecialized MTLPixelFormat = 0
	// MTLPixelFormatX24_Stencil8 - A stencil pixel format used to read the stencil value from a texture with a combined 24-bit depth and 8-bit stencil value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/x24_stencil8
	MTLPixelFormatX24_Stencil8 MTLPixelFormat = 0
	// MTLPixelFormatX32_Stencil8 - A stencil pixel format used to read the stencil value from a texture with a combined 32-bit depth and 8-bit stencil value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/x32_stencil8
	MTLPixelFormatX32_Stencil8 MTLPixelFormat = 0
)

/* debug [enums.gen.go]: Processing enum MTLPrimitiveTopologyClass (4 cases) */
// MTLPrimitiveTopologyClass - The primitive topologies available for rendering.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPrimitiveTopologyClass
type MTLPrimitiveTopologyClass uint

const (
	// MTLPrimitiveTopologyClassLine - A line primitive.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPrimitiveTopologyClass/line
	MTLPrimitiveTopologyClassLine MTLPrimitiveTopologyClass = 0
	// MTLPrimitiveTopologyClassPoint - A point primitive.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPrimitiveTopologyClass/point
	MTLPrimitiveTopologyClassPoint MTLPrimitiveTopologyClass = 0
	// MTLPrimitiveTopologyClassTriangle - A triangle primitive.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPrimitiveTopologyClass/triangle
	MTLPrimitiveTopologyClassTriangle MTLPrimitiveTopologyClass = 0
	// MTLPrimitiveTopologyClassUnspecified - An unspecified primitive.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPrimitiveTopologyClass/unspecified
	MTLPrimitiveTopologyClassUnspecified MTLPrimitiveTopologyClass = 0
)

/* debug [enums.gen.go]: Processing enum MTLPrimitiveType (5 cases) */
// MTLPrimitiveType - The geometric primitive type for drawing commands.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPrimitiveType
type MTLPrimitiveType uint

const (
	// MTLPrimitiveTypeLine - Rasterize a line between each separate pair of vertices, resulting in a series of unconnected lines. If there are an odd number of vertices, the last vertex is ignored.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPrimitiveType/line
	MTLPrimitiveTypeLine MTLPrimitiveType = 0
	// MTLPrimitiveTypeLineStrip - Rasterize a line between each pair of adjacent vertices, resulting in a series of connected lines (also called a polyline).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPrimitiveType/lineStrip
	MTLPrimitiveTypeLineStrip MTLPrimitiveType = 0
	// MTLPrimitiveTypePoint - Rasterize a point at each vertex. The vertex shader must provide  , or the point size is undefined.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPrimitiveType/point
	MTLPrimitiveTypePoint MTLPrimitiveType = 0
	// MTLPrimitiveTypeTriangle - For every separate set of three vertices, rasterize a triangle. If the number of vertices is not a multiple of three, either one or two vertices is ignored.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPrimitiveType/triangle
	MTLPrimitiveTypeTriangle MTLPrimitiveType = 0
	// MTLPrimitiveTypeTriangleStrip - For every three adjacent vertices, rasterize a triangle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPrimitiveType/triangleStrip
	MTLPrimitiveTypeTriangleStrip MTLPrimitiveType = 0
)

/* debug [enums.gen.go]: Processing enum MTLPurgeableState (4 cases) */
// MTLPurgeableState - The purgeable state of the resource.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPurgeableState
type MTLPurgeableState uint

const (
	// MTLPurgeableStateEmpty - A state that indicates to the system that it needs to consider   the contents of a resource as invalid, typically because you’re discarding it.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPurgeableState/empty
	MTLPurgeableStateEmpty MTLPurgeableState = 0
	// MTLPurgeableStateKeepCurrent - The current state is queried but doesn’t change.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPurgeableState/keepCurrent
	MTLPurgeableStateKeepCurrent MTLPurgeableState = 0
	// MTLPurgeableStateNonVolatile - The contents of the resource aren’t allowed to be discarded.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPurgeableState/nonVolatile
	MTLPurgeableStateNonVolatile MTLPurgeableState = 0
	// MTLPurgeableStateVolatile - The system is allowed to discard the resource to free up memory.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPurgeableState/volatile
	MTLPurgeableStateVolatile MTLPurgeableState = 0
)

/* debug [enums.gen.go]: Processing enum MTLReadWriteTextureTier (3 cases) */
// MTLReadWriteTextureTier - The support level for read-write texture formats.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLReadWriteTextureTier
type MTLReadWriteTextureTier uint

const (
	// MTLReadWriteTextureTier1 - Tier 1 read/write textures are supported.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLReadWriteTextureTier/tier1
	MTLReadWriteTextureTier1 MTLReadWriteTextureTier = 0
	// MTLReadWriteTextureTier2 - Tier 2 read/write textures are supported.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLReadWriteTextureTier/tier2
	MTLReadWriteTextureTier2 MTLReadWriteTextureTier = 0
	// MTLReadWriteTextureTierNone - Read-write textures are not supported.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLReadWriteTextureTier/tierNone
	MTLReadWriteTextureTierNone MTLReadWriteTextureTier = 0
)

/* debug [enums.gen.go]: Processing enum MTLRenderStages (5 cases) */
// MTLRenderStages - The stages in a render pass that triggers a synchronization command.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderStages
type MTLRenderStages uint

const (
	// MTLRenderStageFragment - The fragment rendering stage.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderStages/fragment
	MTLRenderStageFragment MTLRenderStages = 0
	// MTLRenderStageMesh - The mesh rendering stage.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderStages/mesh
	MTLRenderStageMesh MTLRenderStages = 0
	// MTLRenderStageObject - The object rendering stage.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderStages/object
	MTLRenderStageObject MTLRenderStages = 0
	// MTLRenderStageTile - The tile rendering stage.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderStages/tile
	MTLRenderStageTile MTLRenderStages = 0
	// MTLRenderStageVertex - The vertex rendering stage.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRenderStages/vertex
	MTLRenderStageVertex MTLRenderStages = 0
)

/* debug [enums.gen.go]: Processing enum MTLResourceOptions (11 cases) */
// MTLResourceOptions - Optional arguments used to set the behavior of a resource.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceOptions
type MTLResourceOptions uint

const (
	// MTLResourceCPUCacheModeWriteCombined - A write-combined CPU cache mode that is optimized for resources that the CPU writes into, but never reads.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceOptions/cpuCacheModeWriteCombined
	MTLResourceCPUCacheModeWriteCombined MTLResourceOptions = 0
	// MTLResourceHazardTrackingModeTracked - An option specifying that Metal prevents hazards when modifying this object’s contents.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceOptions/hazardTrackingModeTracked
	MTLResourceHazardTrackingModeTracked MTLResourceOptions = 0
	// MTLResourceHazardTrackingModeUntracked - An option specifying that the app must prevent hazards when modifying this object’s contents.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceOptions/hazardTrackingModeUntracked
	MTLResourceHazardTrackingModeUntracked MTLResourceOptions = 0
	// MTLResourceCPUCacheModeDefaultCache - The default CPU cache mode for the resource, which guarantees that read and write operations are executed in the expected order.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceOptions/MTLResourceCPUCacheModeDefaultCache
	MTLResourceCPUCacheModeDefaultCache MTLResourceOptions = 0
	// MTLResourceHazardTrackingModeDefault - An option specifying that the default tracking mode should be used.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceOptions/MTLResourceHazardTrackingModeDefault
	MTLResourceHazardTrackingModeDefault MTLResourceOptions = 0
	// MTLResourceOptionCPUCacheModeDefault - This constant was deprecated in iOS 9.0 and macOS 10.11.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceOptions/MTLResourceOptionCPUCacheModeDefault
	MTLResourceOptionCPUCacheModeDefault MTLResourceOptions = 0
	// MTLResourceOptionCPUCacheModeWriteCombined - This constant was deprecated in iOS 9.0 and macOS 10.11.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceOptions/optionCPUCacheModeWriteCombined
	MTLResourceOptionCPUCacheModeWriteCombined MTLResourceOptions = 0
	// MTLResourceStorageModeManaged - The CPU and GPU may maintain separate copies of the resource, and any changes must be explicitly synchronized.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceOptions/storageModeManaged
	MTLResourceStorageModeManaged MTLResourceOptions = 0
	// MTLResourceStorageModeMemoryless - The resource’s contents are only available to the GPU, and only exist temporarily during a render pass.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceOptions/storageModeMemoryless
	MTLResourceStorageModeMemoryless MTLResourceOptions = 0
	// MTLResourceStorageModePrivate - The resource is only available to the GPU.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceOptions/storageModePrivate
	MTLResourceStorageModePrivate MTLResourceOptions = 0
	// MTLResourceStorageModeShared - The CPU and GPU share access to the resource, allocated in system memory.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceOptions/storageModeShared
	MTLResourceStorageModeShared MTLResourceOptions = 0
)

/* debug [enums.gen.go]: Processing enum MTLResourceUsage (3 cases) */
// MTLResourceUsage - Options that describe how a graphics or compute function uses an argument buffer’s resource.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceUsage
type MTLResourceUsage uint

const (
	// MTLResourceUsageRead - An option that enables reading from the resource.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceUsage/read
	MTLResourceUsageRead MTLResourceUsage = 0
	// MTLResourceUsageSample - An option that enables sampling from the resource.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceUsage/sample
	MTLResourceUsageSample MTLResourceUsage = 0
	// MTLResourceUsageWrite - An option that enables writing to the resource.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceUsage/write
	MTLResourceUsageWrite MTLResourceUsage = 0
)

/* debug [enums.gen.go]: Processing enum MTLSamplerAddressMode (6 cases) */
// MTLSamplerAddressMode - Modes that determine the texture coordinate at each pixel when a fetch falls outside the bounds of a texture.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerAddressMode
type MTLSamplerAddressMode uint

const (
	// MTLSamplerAddressModeClampToBorderColor - Out-of-range texture coordinates return the value specified by the   property.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerAddressMode/clampToBorderColor
	MTLSamplerAddressModeClampToBorderColor MTLSamplerAddressMode = 0
	// MTLSamplerAddressModeClampToEdge - Texture coordinates are clamped between   and  , inclusive.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerAddressMode/clampToEdge
	MTLSamplerAddressModeClampToEdge MTLSamplerAddressMode = 0
	// MTLSamplerAddressModeClampToZero - Out-of-range texture coordinates return transparent zero   for images with an alpha channel and return opaque zero   for images without an alpha channel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerAddressMode/clampToZero
	MTLSamplerAddressModeClampToZero MTLSamplerAddressMode = 0
	// MTLSamplerAddressModeMirrorClampToEdge - Between   and  , the texture coordinates are mirrored across the axis; outside   and  , texture coordinates are clamped.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerAddressMode/mirrorClampToEdge
	MTLSamplerAddressModeMirrorClampToEdge MTLSamplerAddressMode = 0
	// MTLSamplerAddressModeMirrorRepeat - Between   and  , the texture coordinates are mirrored across the axis; outside   and  , the image is repeated.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerAddressMode/mirrorRepeat
	MTLSamplerAddressModeMirrorRepeat MTLSamplerAddressMode = 0
	// MTLSamplerAddressModeRepeat - Texture coordinates wrap to the other side of the texture, effectively keeping only the fractional part of the texture coordinate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerAddressMode/repeat
	MTLSamplerAddressModeRepeat MTLSamplerAddressMode = 0
)

/* debug [enums.gen.go]: Processing enum MTLSamplerBorderColor (3 cases) */
// MTLSamplerBorderColor - Values that determine the border color for clamped texture values when the sampler address mode is 
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerBorderColor
type MTLSamplerBorderColor uint

const (
	// MTLSamplerBorderColorOpaqueBlack - An opaque black color   for texture values outside the border
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerBorderColor/opaqueBlack
	MTLSamplerBorderColorOpaqueBlack MTLSamplerBorderColor = 0
	// MTLSamplerBorderColorOpaqueWhite - An opaque white color   for texture values outside the border.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerBorderColor/opaqueWhite
	MTLSamplerBorderColorOpaqueWhite MTLSamplerBorderColor = 0
	// MTLSamplerBorderColorTransparentBlack - A transparent black color   for texture values outside the border.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerBorderColor/transparentBlack
	MTLSamplerBorderColorTransparentBlack MTLSamplerBorderColor = 0
)

/* debug [enums.gen.go]: Processing enum MTLSamplerMinMagFilter (2 cases) */
// MTLSamplerMinMagFilter - Filtering options for determining which pixel value is returned within a mipmap level.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerMinMagFilter
type MTLSamplerMinMagFilter uint

const (
	// MTLSamplerMinMagFilterLinear - Select two pixels in each dimension and interpolate linearly between them.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerMinMagFilter/linear
	MTLSamplerMinMagFilterLinear MTLSamplerMinMagFilter = 0
	// MTLSamplerMinMagFilterNearest - Select the single pixel nearest to the sample point.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerMinMagFilter/nearest
	MTLSamplerMinMagFilterNearest MTLSamplerMinMagFilter = 0
)

/* debug [enums.gen.go]: Processing enum MTLSamplerMipFilter (3 cases) */
// MTLSamplerMipFilter - Filtering options for determining what pixel value is returned with multiple mipmap levels.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerMipFilter
type MTLSamplerMipFilter uint

const (
	// MTLSamplerMipFilterLinear - If the filter falls between mipmap levels, both levels are sampled and the results are determined by linear interpolation between levels.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerMipFilter/linear
	MTLSamplerMipFilterLinear MTLSamplerMipFilter = 0
	// MTLSamplerMipFilterNearest - The nearest mipmap level is selected.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerMipFilter/nearest
	MTLSamplerMipFilterNearest MTLSamplerMipFilter = 0
	// MTLSamplerMipFilterNotMipmapped - The texture is sampled from mipmap level  , and other mipmap levels are ignored.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerMipFilter/notMipmapped
	MTLSamplerMipFilterNotMipmapped MTLSamplerMipFilter = 0
)

/* debug [enums.gen.go]: Processing enum MTLSamplerReductionMode (3 cases) */
// MTLSamplerReductionMode - Configures how the sampler aggregates contributing samples to a final value.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerReductionMode
type MTLSamplerReductionMode uint

const (
	// MTLSamplerReductionModeMaximum - A reduction mode that finds the maximum contributing sample value by separately evaluating each channel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerReductionMode/maximum
	MTLSamplerReductionModeMaximum MTLSamplerReductionMode = 0
	// MTLSamplerReductionModeMinimum - A reduction mode that finds the minimum contributing sample value by separately evaluating each channel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerReductionMode/minimum
	MTLSamplerReductionModeMinimum MTLSamplerReductionMode = 0
	// MTLSamplerReductionModeWeightedAverage - A reduction mode that adds together the product of each contributing sample value by its weight.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerReductionMode/weightedAverage
	MTLSamplerReductionModeWeightedAverage MTLSamplerReductionMode = 0
)

/* debug [enums.gen.go]: Processing enum MTLShaderValidation (3 cases) */
// MTLShaderValidation - Indicates whether shader validation in an enabled or disabled state, or neither state.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLShaderValidation
type MTLShaderValidation uint

const (
	// MTLShaderValidationDefault - The default value when the property isn’t set.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLShaderValidation/default
	MTLShaderValidationDefault MTLShaderValidation = 0
	// MTLShaderValidationDisabled - Disables shader validation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLShaderValidation/disabled
	MTLShaderValidationDisabled MTLShaderValidation = 0
	// MTLShaderValidationEnabled - Enables shader validation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLShaderValidation/enabled
	MTLShaderValidationEnabled MTLShaderValidation = 0
)

/* debug [enums.gen.go]: Processing enum MTLSparsePageSize (3 cases) */
// MTLSparsePageSize - The page size options, in kilobytes, for sparse textures.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSparsePageSize
type MTLSparsePageSize uint

const (
	// MTLSparsePageSize16 - Represents a sparse texture’s page size of 16 kilobytes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSparsePageSize/size16
	MTLSparsePageSize16 MTLSparsePageSize = 0
	// MTLSparsePageSize256 - Represents a sparse texture’s page size of 256 kilobytes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSparsePageSize/size256
	MTLSparsePageSize256 MTLSparsePageSize = 0
	// MTLSparsePageSize64 - Represents a sparse texture’s page size of 64 kilobytes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSparsePageSize/size64
	MTLSparsePageSize64 MTLSparsePageSize = 0
)

/* debug [enums.gen.go]: Processing enum MTLSparseTextureMappingMode (2 cases) */
// MTLSparseTextureMappingMode - Options for sparse texture mapping.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSparseTextureMappingMode
type MTLSparseTextureMappingMode uint

const (
	// MTLSparseTextureMappingModeMap - A request to map sparse tiles from the heap to a region in the texture.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSparseTextureMappingMode/map
	MTLSparseTextureMappingModeMap MTLSparseTextureMappingMode = 0
	// MTLSparseTextureMappingModeUnmap - A request to remove any mappings for a region in the texture.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSparseTextureMappingMode/unmap
	MTLSparseTextureMappingModeUnmap MTLSparseTextureMappingMode = 0
)

/* debug [enums.gen.go]: Processing enum MTLSparseTextureRegionAlignmentMode (2 cases) */
// MTLSparseTextureRegionAlignmentMode - Options used when converting between a pixel-based region within a texture to a tile-based region.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSparseTextureRegionAlignmentMode
type MTLSparseTextureRegionAlignmentMode uint

const (
	// MTLSparseTextureRegionAlignmentModeInward - The tile region ignores partially covered tiles.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSparseTextureRegionAlignmentMode/inward
	MTLSparseTextureRegionAlignmentModeInward MTLSparseTextureRegionAlignmentMode = 0
	// MTLSparseTextureRegionAlignmentModeOutward - The tile region includes any partially covered tiles.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSparseTextureRegionAlignmentMode/outward
	MTLSparseTextureRegionAlignmentModeOutward MTLSparseTextureRegionAlignmentMode = 0
)

/* debug [enums.gen.go]: Processing enum MTLStages (11 cases) */
// MTLStages - Describes stages of GPU work.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStages
type MTLStages uint

const (
	// MTLStageAccelerationStructure - Represents all acceleration structure operations.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStages/accelerationStructure
	MTLStageAccelerationStructure MTLStages = 0
	// MTLStageAll - Convenience mask representing all stages of GPU work.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStages/all
	MTLStageAll MTLStages = 0
	// MTLStageBlit - Represents all blit operations in a pass.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStages/blit
	MTLStageBlit MTLStages = 0
	// MTLStageDispatch - Represents all compute dispatches in a compute pass.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStages/dispatch
	MTLStageDispatch MTLStages = 0
	// MTLStageFragment - Represents all fragment shader stage work in a render pass.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStages/fragment
	MTLStageFragment MTLStages = 0
	// MTLStageMachineLearning - Represents all machine learning network dispatch operations.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStages/machineLearning
	MTLStageMachineLearning MTLStages = 0
	// MTLStageMesh - Represents all mesh shader stage work work in a render pass.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStages/mesh
	MTLStageMesh MTLStages = 0
	// MTLStageObject - Represents all object shader stage work in a render pass.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStages/object
	MTLStageObject MTLStages = 0
	// MTLStageResourceState - Represents all sparse and placement sparse resource mapping updates.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStages/resourceState
	MTLStageResourceState MTLStages = 0
	// MTLStageTile - Represents all tile shading stage work in a render pass.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStages/tile
	MTLStageTile MTLStages = 0
	// MTLStageVertex - Represents all vertex shader stage work in a render pass.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStages/vertex
	MTLStageVertex MTLStages = 0
)

/* debug [enums.gen.go]: Processing enum MTLStencilOperation (8 cases) */
// MTLStencilOperation - The operation performed on a currently stored stencil value when a comparison test passes or fails.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStencilOperation
type MTLStencilOperation uint

const (
	// MTLStencilOperationDecrementClamp - If the current stencil value is not zero, decrease the stencil value by one. Otherwise, if the current stencil value is zero, do not change the stencil value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStencilOperation/decrementClamp
	MTLStencilOperationDecrementClamp MTLStencilOperation = 0
	// MTLStencilOperationDecrementWrap - If the current stencil value is not zero, decrease the stencil value by one. Otherwise, if the current stencil value is zero, set the stencil value to the maximum representable value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStencilOperation/decrementWrap
	MTLStencilOperationDecrementWrap MTLStencilOperation = 0
	// MTLStencilOperationIncrementClamp - If the current stencil value is not the maximum representable value, increase the stencil value by one. Otherwise, if the current stencil value is the maximum representable value, do not change the stencil value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStencilOperation/incrementClamp
	MTLStencilOperationIncrementClamp MTLStencilOperation = 0
	// MTLStencilOperationIncrementWrap - If the current stencil value is not the maximum representable value, increase the stencil value by one. Otherwise, if the current stencil value is the maximum representable value, set the stencil value to zero.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStencilOperation/incrementWrap
	MTLStencilOperationIncrementWrap MTLStencilOperation = 0
	// MTLStencilOperationInvert - Perform a logical bitwise invert operation on the current stencil value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStencilOperation/invert
	MTLStencilOperationInvert MTLStencilOperation = 0
	// MTLStencilOperationKeep - Keep the current stencil value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStencilOperation/keep
	MTLStencilOperationKeep MTLStencilOperation = 0
	// MTLStencilOperationReplace - Replace the stencil value with the stencil reference value, which is set by the   method of  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStencilOperation/replace
	MTLStencilOperationReplace MTLStencilOperation = 0
	// MTLStencilOperationZero - Set the stencil value to zero.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStencilOperation/zero
	MTLStencilOperationZero MTLStencilOperation = 0
)

/* debug [enums.gen.go]: Processing enum MTLStepFunction (9 cases) */
// MTLStepFunction - The frequency and locations at which a function fetches attribute data.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStepFunction
type MTLStepFunction uint

const (
	// MTLStepFunctionConstant - The function fetches attribute data once.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStepFunction/constant
	MTLStepFunctionConstant MTLStepFunction = 0
	// MTLStepFunctionPerInstance - The function fetches data based on the instance index.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStepFunction/perInstance
	MTLStepFunctionPerInstance MTLStepFunction = 0
	// MTLStepFunctionPerPatch - The post-tessellation function fetches data based on the patch index of the patch.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStepFunction/perPatch
	MTLStepFunctionPerPatch MTLStepFunction = 0
	// MTLStepFunctionPerPatchControlPoint - The post-tessellation function fetches data based on the control-point indices associated with the patch.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStepFunction/perPatchControlPoint
	MTLStepFunctionPerPatchControlPoint MTLStepFunction = 0
	// MTLStepFunctionPerVertex - The vertex function fetches data for every vertex.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStepFunction/perVertex
	MTLStepFunctionPerVertex MTLStepFunction = 0
	// MTLStepFunctionThreadPositionInGridX - The compute function fetches data based on the thread’s   coordinate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStepFunction/threadPositionInGridX
	MTLStepFunctionThreadPositionInGridX MTLStepFunction = 0
	// MTLStepFunctionThreadPositionInGridXIndexed - The compute function fetches data by using the thread’s   coordinate to look up a value in the index buffer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStepFunction/threadPositionInGridXIndexed
	MTLStepFunctionThreadPositionInGridXIndexed MTLStepFunction = 0
	// MTLStepFunctionThreadPositionInGridY - The compute function fetches data based on the thread’s   coordinate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStepFunction/threadPositionInGridY
	MTLStepFunctionThreadPositionInGridY MTLStepFunction = 0
	// MTLStepFunctionThreadPositionInGridYIndexed - The compute function fetches data by using the thread’s   coordinate to look up a value in the index buffer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStepFunction/threadPositionInGridYIndexed
	MTLStepFunctionThreadPositionInGridYIndexed MTLStepFunction = 0
)

/* debug [enums.gen.go]: Processing enum MTLStitchedLibraryOptions (3 cases) */
// MTLStitchedLibraryOptions enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStitchedLibraryOptions
type MTLStitchedLibraryOptions uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStitchedLibraryOptions/failOnBinaryArchiveMiss
	MTLStitchedLibraryOptionFailOnBinaryArchiveMiss MTLStitchedLibraryOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStitchedLibraryOptions/MTLStitchedLibraryOptionNone
	MTLStitchedLibraryOptionNone MTLStitchedLibraryOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStitchedLibraryOptions/storeLibraryInMetalPipelinesScript
	MTLStitchedLibraryOptionStoreLibraryInMetalPipelinesScript MTLStitchedLibraryOptions = 0
)

/* debug [enums.gen.go]: Processing enum MTLStorageMode (4 cases) */
// MTLStorageMode - Options for the memory location and access permissions for a resource.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStorageMode
type MTLStorageMode uint

const (
	// MTLStorageModeManaged - The CPU and GPU may maintain separate copies of the resource, and any changes must be explicitly synchronized.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStorageMode/managed
	MTLStorageModeManaged MTLStorageMode = 0
	// MTLStorageModeMemoryless - The resource’s contents are only available to the GPU, and only exist temporarily during a render pass.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStorageMode/memoryless
	MTLStorageModeMemoryless MTLStorageMode = 0
	// MTLStorageModePrivate - The resource is only available to the GPU.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStorageMode/private
	MTLStorageModePrivate MTLStorageMode = 0
	// MTLStorageModeShared - The CPU and GPU share access to the resource, allocated in system memory.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStorageMode/shared
	MTLStorageModeShared MTLStorageMode = 0
)

/* debug [enums.gen.go]: Processing enum MTLStoreAction (6 cases) */
// MTLStoreAction - Types of actions performed for an attachment at the end of a rendering pass.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStoreAction
type MTLStoreAction uint

const (
	// MTLStoreActionCustomSampleDepthStore - The GPU stores depth data in a sample-position–agnostic representation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStoreAction/customSampleDepthStore
	MTLStoreActionCustomSampleDepthStore MTLStoreAction = 0
	// MTLStoreActionDontCare - The GPU has permission to discard the rendered contents of the attachment at the end of the render pass, replacing them with arbitrary data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStoreAction/dontCare
	MTLStoreActionDontCare MTLStoreAction = 0
	// MTLStoreActionMultisampleResolve - The GPU resolves the multisampled data to one sample per pixel and stores the data to the resolve texture, discarding the multisample data afterwards.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStoreAction/multisampleResolve
	MTLStoreActionMultisampleResolve MTLStoreAction = 0
	// MTLStoreActionStore - The GPU stores the rendered contents to the texture.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStoreAction/store
	MTLStoreActionStore MTLStoreAction = 0
	// MTLStoreActionStoreAndMultisampleResolve - The GPU stores the multisample data to the multisample texture, resolves the data to a sample per pixel, and stores the data to the resolve texture.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStoreAction/storeAndMultisampleResolve
	MTLStoreActionStoreAndMultisampleResolve MTLStoreAction = 0
	// MTLStoreActionUnknown - The system selects a store action when it encodes the render pass.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStoreAction/unknown
	MTLStoreActionUnknown MTLStoreAction = 0
)

/* debug [enums.gen.go]: Processing enum MTLStoreActionOptions (2 cases) */
// MTLStoreActionOptions - Options that modify a store action.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStoreActionOptions
type MTLStoreActionOptions uint

const (
	// MTLStoreActionOptionCustomSamplePositions - An option that stores data in a sample-position–agnostic representation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStoreActionOptions/customSamplePositions
	MTLStoreActionOptionCustomSamplePositions MTLStoreActionOptions = 0
	// MTLStoreActionOptionNone - An option that doesn’t modify the intended behavior of a store action.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStoreActionOptions/MTLStoreActionOptionNone
	MTLStoreActionOptionNone MTLStoreActionOptions = 0
)

/* debug [enums.gen.go]: Processing enum MTLTensorDataType (10 cases) */
// MTLTensorDataType - The possible data types for the elements of a tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorDataType
type MTLTensorDataType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorDataType/bfloat16
	MTLTensorDataTypeBFloat16 MTLTensorDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorDataType/float16
	MTLTensorDataTypeFloat16 MTLTensorDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorDataType/float32
	MTLTensorDataTypeFloat32 MTLTensorDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorDataType/int16
	MTLTensorDataTypeInt16 MTLTensorDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorDataType/int32
	MTLTensorDataTypeInt32 MTLTensorDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorDataType/int8
	MTLTensorDataTypeInt8 MTLTensorDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorDataType/none
	MTLTensorDataTypeNone MTLTensorDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorDataType/uint16
	MTLTensorDataTypeUInt16 MTLTensorDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorDataType/uint32
	MTLTensorDataTypeUInt32 MTLTensorDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorDataType/uint8
	MTLTensorDataTypeUInt8 MTLTensorDataType = 0
)

/* debug [enums.gen.go]: Processing enum MTLTensorError (3 cases) */
// MTLTensorError - The error codes that Metal can raise when you create a tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorError-swift.struct/Code
type MTLTensorError uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorError-swift.struct/Code/internalError
	MTLTensorErrorInternalError MTLTensorError = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorError-swift.struct/Code/invalidDescriptor
	MTLTensorErrorInvalidDescriptor MTLTensorError = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorError-swift.struct/Code/none
	MTLTensorErrorNone MTLTensorError = 0
)

/* debug [enums.gen.go]: Processing enum MTLTensorUsage (3 cases) */
// MTLTensorUsage - The type that represents the different contexts for a tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorUsage
type MTLTensorUsage uint

const (
	// MTLTensorUsageCompute - A tensor context that applies to compute encoders.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorUsage/compute
	MTLTensorUsageCompute MTLTensorUsage = 0
	// MTLTensorUsageMachineLearning - A tensor context that applies to machine learning encoders.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorUsage/machineLearning
	MTLTensorUsageMachineLearning MTLTensorUsage = 0
	// MTLTensorUsageRender - A tensor context that applies to render encoders.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorUsage/render
	MTLTensorUsageRender MTLTensorUsage = 0
)

/* debug [enums.gen.go]: Processing enum MTLTessellationControlPointIndexType (3 cases) */
// MTLTessellationControlPointIndexType - Options for specifying the size of the control point indices in a control point index buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTessellationControlPointIndexType
type MTLTessellationControlPointIndexType uint

const (
	// MTLTessellationControlPointIndexTypeNone - No size. This value should only be used when drawing patches without a control point index buffer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTessellationControlPointIndexType/none
	MTLTessellationControlPointIndexTypeNone MTLTessellationControlPointIndexType = 0
	// MTLTessellationControlPointIndexTypeUInt16 - The size of a 16-bit unsigned integer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTessellationControlPointIndexType/uint16
	MTLTessellationControlPointIndexTypeUInt16 MTLTessellationControlPointIndexType = 0
	// MTLTessellationControlPointIndexTypeUInt32 - The size of a 32-bit unsigned integer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTessellationControlPointIndexType/uint32
	MTLTessellationControlPointIndexTypeUInt32 MTLTessellationControlPointIndexType = 0
)

/* debug [enums.gen.go]: Processing enum MTLTessellationFactorFormat (1 cases) */
// MTLTessellationFactorFormat - Options for specifying the format of the tessellation factors in a tessellation factor buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTessellationFactorFormat
type MTLTessellationFactorFormat uint

const (
	// MTLTessellationFactorFormatHalf - A 16-bit floating-point format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTessellationFactorFormat/half
	MTLTessellationFactorFormatHalf MTLTessellationFactorFormat = 0
)

/* debug [enums.gen.go]: Processing enum MTLTessellationFactorStepFunction (4 cases) */
// MTLTessellationFactorStepFunction - Options for specifying the step function that determines the tessellation factors for a patch from the tessellation factor buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTessellationFactorStepFunction
type MTLTessellationFactorStepFunction uint

const (
	// MTLTessellationFactorStepFunctionConstant - A constant step function. For all instances, the tessellation factor for all patches in a patch draw call is at the   location in the tessellation factor buffer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTessellationFactorStepFunction/constant
	MTLTessellationFactorStepFunctionConstant MTLTessellationFactorStepFunction = 0
	// MTLTessellationFactorStepFunctionPerInstance - A per-instance step function. For a given instance ID, the tessellation factor for a patch in a patch draw call is at the   location in the tessellation factor buffer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTessellationFactorStepFunction/perInstance
	MTLTessellationFactorStepFunctionPerInstance MTLTessellationFactorStepFunction = 0
	// MTLTessellationFactorStepFunctionPerPatch - A per-patch step function. For all instances, the tessellation factor for all patches in a patch draw call is at the   location in the tessellation factor buffer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTessellationFactorStepFunction/perPatch
	MTLTessellationFactorStepFunctionPerPatch MTLTessellationFactorStepFunction = 0
	// MTLTessellationFactorStepFunctionPerPatchAndPerInstance - A per-patch and per-instance step function. For a given instance ID, the tessellation factor for a patch in a patch draw call is at the   location in the tessellation factor buffer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTessellationFactorStepFunction/perPatchAndPerInstance
	MTLTessellationFactorStepFunctionPerPatchAndPerInstance MTLTessellationFactorStepFunction = 0
)

/* debug [enums.gen.go]: Processing enum MTLTessellationPartitionMode (4 cases) */
// MTLTessellationPartitionMode - Options for choosing the partition mode that the tessellator applies when deriving
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTessellationPartitionMode
type MTLTessellationPartitionMode uint

const (
	// MTLTessellationPartitionModeFractionalEven - A fractional even partitioning mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTessellationPartitionMode/fractionalEven
	MTLTessellationPartitionModeFractionalEven MTLTessellationPartitionMode = 0
	// MTLTessellationPartitionModeFractionalOdd - A fractional odd partitioning mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTessellationPartitionMode/fractionalOdd
	MTLTessellationPartitionModeFractionalOdd MTLTessellationPartitionMode = 0
	// MTLTessellationPartitionModeInteger - An integer partitioning mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTessellationPartitionMode/integer
	MTLTessellationPartitionModeInteger MTLTessellationPartitionMode = 0
	// MTLTessellationPartitionModePow2 - A power of two partitioning mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTessellationPartitionMode/pow2
	MTLTessellationPartitionModePow2 MTLTessellationPartitionMode = 0
)

/* debug [enums.gen.go]: Processing enum MTLTextureCompressionType (2 cases) */
// MTLTextureCompressionType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureCompressionType
type MTLTextureCompressionType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureCompressionType/lossless
	MTLTextureCompressionTypeLossless MTLTextureCompressionType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureCompressionType/lossy
	MTLTextureCompressionTypeLossy MTLTextureCompressionType = 0
)

/* debug [enums.gen.go]: Processing enum MTLTextureSparseTier (3 cases) */
// MTLTextureSparseTier - Enumerates the different support levels for sparse textures.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureSparseTier
type MTLTextureSparseTier uint

const (
	// MTLTextureSparseTier1 - Indicates support for sparse textures tier 1.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureSparseTier/tier1
	MTLTextureSparseTier1 MTLTextureSparseTier = 0
	// MTLTextureSparseTier2 - Indicates support for sparse textures tier 2.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureSparseTier/tier2
	MTLTextureSparseTier2 MTLTextureSparseTier = 0
	// MTLTextureSparseTierNone - Indicates that the texture is not sparse.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureSparseTier/tierNone
	MTLTextureSparseTierNone MTLTextureSparseTier = 0
)

/* debug [enums.gen.go]: Processing enum MTLTextureSwizzle (6 cases) */
// MTLTextureSwizzle - A set of options to choose from when creating a texture swizzle pattern.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureSwizzle
type MTLTextureSwizzle uint

const (
	// MTLTextureSwizzleAlpha - The alpha channel of the source pixel is copied to the destination channel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureSwizzle/alpha
	MTLTextureSwizzleAlpha MTLTextureSwizzle = 0
	// MTLTextureSwizzleBlue - The blue channel of the source pixel is copied to the destination channel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureSwizzle/blue
	MTLTextureSwizzleBlue MTLTextureSwizzle = 0
	// MTLTextureSwizzleGreen - The green channel of the source pixel is copied to the destination channel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureSwizzle/green
	MTLTextureSwizzleGreen MTLTextureSwizzle = 0
	// MTLTextureSwizzleOne - A value of   is copied to the destination channel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureSwizzle/one
	MTLTextureSwizzleOne MTLTextureSwizzle = 0
	// MTLTextureSwizzleRed - The red channel of the source pixel is copied to the destination channel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureSwizzle/red
	MTLTextureSwizzleRed MTLTextureSwizzle = 0
	// MTLTextureSwizzleZero - A value of   is copied to the destination channel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureSwizzle/zero
	MTLTextureSwizzleZero MTLTextureSwizzle = 0
)

/* debug [enums.gen.go]: Processing enum MTLTextureType (10 cases) */
// MTLTextureType - The dimension of each image, including whether multiple images are arranged into an array or a cube.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureType
type MTLTextureType uint

const (
	// MTLTextureType1D - A one-dimensional texture image.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureType/type1D
	MTLTextureType1D MTLTextureType = 0
	// MTLTextureType1DArray - An array of one-dimensional texture images.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureType/type1DArray
	MTLTextureType1DArray MTLTextureType = 0
	// MTLTextureType2D - A two-dimensional texture image.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureType/type2D
	MTLTextureType2D MTLTextureType = 0
	// MTLTextureType2DArray - An array of two-dimensional texture images.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureType/type2DArray
	MTLTextureType2DArray MTLTextureType = 0
	// MTLTextureType2DMultisample - A two-dimensional texture image that uses more than one sample for each pixel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureType/type2DMultisample
	MTLTextureType2DMultisample MTLTextureType = 0
	// MTLTextureType2DMultisampleArray - An array of two-dimensional texture images that use more than one sample for each pixel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureType/type2DMultisampleArray
	MTLTextureType2DMultisampleArray MTLTextureType = 0
	// MTLTextureType3D - A three-dimensional texture image.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureType/type3D
	MTLTextureType3D MTLTextureType = 0
	// MTLTextureTypeCube - A cube texture with six two-dimensional images.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureType/typeCube
	MTLTextureTypeCube MTLTextureType = 0
	// MTLTextureTypeCubeArray - An array of cube textures, each with six two-dimensional images.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureType/typeCubeArray
	MTLTextureTypeCubeArray MTLTextureType = 0
	// MTLTextureTypeTextureBuffer - A texture buffer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureType/typeTextureBuffer
	MTLTextureTypeTextureBuffer MTLTextureType = 0
)

/* debug [enums.gen.go]: Processing enum MTLTextureUsage (6 cases) */
// MTLTextureUsage - An enumeration for the various options that determine how you can use a texture.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureUsage
type MTLTextureUsage uint

const (
	// MTLTextureUsagePixelFormatView - An option to create texture views with a different component layout.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureUsage/pixelFormatView
	MTLTextureUsagePixelFormatView MTLTextureUsage = 0
	// MTLTextureUsageRenderTarget - An option for rendering to the texture in a render pass.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureUsage/renderTarget
	MTLTextureUsageRenderTarget MTLTextureUsage = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureUsage/shaderAtomic
	MTLTextureUsageShaderAtomic MTLTextureUsage = 0
	// MTLTextureUsageShaderRead - An option for reading or sampling from the texture in a shader.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureUsage/shaderRead
	MTLTextureUsageShaderRead MTLTextureUsage = 0
	// MTLTextureUsageShaderWrite - An option for writing to the texture in a shader.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureUsage/shaderWrite
	MTLTextureUsageShaderWrite MTLTextureUsage = 0
	// MTLTextureUsageUnknown - An option for a texture whose usage is unknown.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureUsage/unknown
	MTLTextureUsageUnknown MTLTextureUsage = 0
)

/* debug [enums.gen.go]: Processing enum MTLTransformType (2 cases) */
// MTLTransformType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTransformType
type MTLTransformType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTransformType/component
	MTLTransformTypeComponent MTLTransformType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTransformType/packedFloat4x3
	MTLTransformTypePackedFloat4x3 MTLTransformType = 0
)

/* debug [enums.gen.go]: Processing enum MTLTriangleFillMode (2 cases) */
// MTLTriangleFillMode - Specifies how to rasterize triangle and triangle strip primitives.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTriangleFillMode
type MTLTriangleFillMode uint

const (
	// MTLTriangleFillModeFill - Rasterize triangle and triangle strip primitives as filled triangles.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTriangleFillMode/fill
	MTLTriangleFillModeFill MTLTriangleFillMode = 0
	// MTLTriangleFillModeLines - Rasterize triangle and triangle strip primitives as lines.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTriangleFillMode/lines
	MTLTriangleFillModeLines MTLTriangleFillMode = 0
)

/* debug [enums.gen.go]: Processing enum MTLVertexFormat (54 cases) */
// MTLVertexFormat - Values that specify the organization of function vertex data.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat
type MTLVertexFormat uint

const (
	// MTLVertexFormatChar - One signed 8-bit two’s complement value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/char
	MTLVertexFormatChar MTLVertexFormat = 0
	// MTLVertexFormatChar2 - Two signed 8-bit two’s complement values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/char2
	MTLVertexFormatChar2 MTLVertexFormat = 0
	// MTLVertexFormatChar2Normalized - Two signed normalized 8-bit two’s complement values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/char2Normalized
	MTLVertexFormatChar2Normalized MTLVertexFormat = 0
	// MTLVertexFormatChar3 - Three signed 8-bit two’s complement values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/char3
	MTLVertexFormatChar3 MTLVertexFormat = 0
	// MTLVertexFormatChar3Normalized - Three signed normalized 8-bit two’s complement values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/char3Normalized
	MTLVertexFormatChar3Normalized MTLVertexFormat = 0
	// MTLVertexFormatChar4 - Four signed 8-bit two’s complement values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/char4
	MTLVertexFormatChar4 MTLVertexFormat = 0
	// MTLVertexFormatChar4Normalized - Four signed normalized 8-bit two’s complement values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/char4Normalized
	MTLVertexFormatChar4Normalized MTLVertexFormat = 0
	// MTLVertexFormatCharNormalized - One signed normalized 8-bit two’s complement value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/charNormalized
	MTLVertexFormatCharNormalized MTLVertexFormat = 0
	// MTLVertexFormatFloat - One single-precision floating-point value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/float
	MTLVertexFormatFloat MTLVertexFormat = 0
	// MTLVertexFormatFloat2 - Two single-precision floating-point values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/float2
	MTLVertexFormatFloat2 MTLVertexFormat = 0
	// MTLVertexFormatFloat3 - Three single-precision floating-point values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/float3
	MTLVertexFormatFloat3 MTLVertexFormat = 0
	// MTLVertexFormatFloat4 - Four single-precision floating-point values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/float4
	MTLVertexFormatFloat4 MTLVertexFormat = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/floatRG11B10
	MTLVertexFormatFloatRG11B10 MTLVertexFormat = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/floatRGB9E5
	MTLVertexFormatFloatRGB9E5 MTLVertexFormat = 0
	// MTLVertexFormatHalf - One half-precision floating-point value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/half
	MTLVertexFormatHalf MTLVertexFormat = 0
	// MTLVertexFormatHalf2 - Two half-precision floating-point values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/half2
	MTLVertexFormatHalf2 MTLVertexFormat = 0
	// MTLVertexFormatHalf3 - Three half-precision floating-point values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/half3
	MTLVertexFormatHalf3 MTLVertexFormat = 0
	// MTLVertexFormatHalf4 - Four half-precision floating-point values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/half4
	MTLVertexFormatHalf4 MTLVertexFormat = 0
	// MTLVertexFormatInt - One signed 32-bit two’s complement value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/int
	MTLVertexFormatInt MTLVertexFormat = 0
	// MTLVertexFormatInt1010102Normalized - One packed 32-bit value with four normalized signed two’s complement integer values, arranged as 10 bits, 10 bits, 10 bits, and 2 bits.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/int1010102Normalized
	MTLVertexFormatInt1010102Normalized MTLVertexFormat = 0
	// MTLVertexFormatInt2 - Two signed 32-bit two’s complement values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/int2
	MTLVertexFormatInt2 MTLVertexFormat = 0
	// MTLVertexFormatInt3 - Three signed 32-bit two’s complement values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/int3
	MTLVertexFormatInt3 MTLVertexFormat = 0
	// MTLVertexFormatInt4 - Four signed 32-bit two’s complement values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/int4
	MTLVertexFormatInt4 MTLVertexFormat = 0
	// MTLVertexFormatInvalid - An invalid vertex format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/invalid
	MTLVertexFormatInvalid MTLVertexFormat = 0
	// MTLVertexFormatShort - One signed 16-bit two’s complement value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/short
	MTLVertexFormatShort MTLVertexFormat = 0
	// MTLVertexFormatShort2 - Two signed 16-bit two’s complement values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/short2
	MTLVertexFormatShort2 MTLVertexFormat = 0
	// MTLVertexFormatShort2Normalized - Two signed normalized 16-bit two’s complement values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/short2Normalized
	MTLVertexFormatShort2Normalized MTLVertexFormat = 0
	// MTLVertexFormatShort3 - Three signed 16-bit two’s complement values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/short3
	MTLVertexFormatShort3 MTLVertexFormat = 0
	// MTLVertexFormatShort3Normalized - Three signed normalized 16-bit two’s complement values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/short3Normalized
	MTLVertexFormatShort3Normalized MTLVertexFormat = 0
	// MTLVertexFormatShort4 - Four signed 16-bit two’s complement values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/short4
	MTLVertexFormatShort4 MTLVertexFormat = 0
	// MTLVertexFormatShort4Normalized - Four signed normalized 16-bit two’s complement values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/short4Normalized
	MTLVertexFormatShort4Normalized MTLVertexFormat = 0
	// MTLVertexFormatShortNormalized - One signed normalized 16-bit two’s complement value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/shortNormalized
	MTLVertexFormatShortNormalized MTLVertexFormat = 0
	// MTLVertexFormatUChar - One unsigned 8-bit value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/uchar
	MTLVertexFormatUChar MTLVertexFormat = 0
	// MTLVertexFormatUChar2 - Two unsigned 8-bit values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/uchar2
	MTLVertexFormatUChar2 MTLVertexFormat = 0
	// MTLVertexFormatUChar2Normalized - Two unsigned normalized 8-bit values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/uchar2Normalized
	MTLVertexFormatUChar2Normalized MTLVertexFormat = 0
	// MTLVertexFormatUChar3 - Three unsigned 8-bit values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/uchar3
	MTLVertexFormatUChar3 MTLVertexFormat = 0
	// MTLVertexFormatUChar3Normalized - Three unsigned normalized 8-bit values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/uchar3Normalized
	MTLVertexFormatUChar3Normalized MTLVertexFormat = 0
	// MTLVertexFormatUChar4 - Four unsigned 8-bit values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/uchar4
	MTLVertexFormatUChar4 MTLVertexFormat = 0
	// MTLVertexFormatUChar4Normalized - Four unsigned normalized 8-bit values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/uchar4Normalized
	MTLVertexFormatUChar4Normalized MTLVertexFormat = 0
	// MTLVertexFormatUChar4Normalized_BGRA - Four unsigned normalized 8-bit values, arranged as blue, green, red, and alpha components.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/uchar4Normalized_bgra
	MTLVertexFormatUChar4Normalized_BGRA MTLVertexFormat = 0
	// MTLVertexFormatUCharNormalized - One unsigned normalized 8-bit value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/ucharNormalized
	MTLVertexFormatUCharNormalized MTLVertexFormat = 0
	// MTLVertexFormatUInt - One unsigned 32-bit value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/uint
	MTLVertexFormatUInt MTLVertexFormat = 0
	// MTLVertexFormatUInt1010102Normalized - One packed 32-bit value with four normalized unsigned integer values, arranged as 10 bits, 10 bits, 10 bits, and 2 bits.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/uint1010102Normalized
	MTLVertexFormatUInt1010102Normalized MTLVertexFormat = 0
	// MTLVertexFormatUInt2 - Two unsigned 32-bit values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/uint2
	MTLVertexFormatUInt2 MTLVertexFormat = 0
	// MTLVertexFormatUInt3 - Three unsigned 32-bit values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/uint3
	MTLVertexFormatUInt3 MTLVertexFormat = 0
	// MTLVertexFormatUInt4 - Four unsigned 32-bit values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/uint4
	MTLVertexFormatUInt4 MTLVertexFormat = 0
	// MTLVertexFormatUShort - One unsigned 16-bit value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/ushort
	MTLVertexFormatUShort MTLVertexFormat = 0
	// MTLVertexFormatUShort2 - Two unsigned 16-bit values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/ushort2
	MTLVertexFormatUShort2 MTLVertexFormat = 0
	// MTLVertexFormatUShort2Normalized - Two unsigned normalized 16-bit values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/ushort2Normalized
	MTLVertexFormatUShort2Normalized MTLVertexFormat = 0
	// MTLVertexFormatUShort3 - Three unsigned 16-bit values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/ushort3
	MTLVertexFormatUShort3 MTLVertexFormat = 0
	// MTLVertexFormatUShort3Normalized - Three unsigned normalized 16-bit values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/ushort3Normalized
	MTLVertexFormatUShort3Normalized MTLVertexFormat = 0
	// MTLVertexFormatUShort4 - Four unsigned 16-bit values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/ushort4
	MTLVertexFormatUShort4 MTLVertexFormat = 0
	// MTLVertexFormatUShort4Normalized - Four unsigned normalized 16-bit values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/ushort4Normalized
	MTLVertexFormatUShort4Normalized MTLVertexFormat = 0
	// MTLVertexFormatUShortNormalized - One unsigned normalized 16-bit value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/ushortNormalized
	MTLVertexFormatUShortNormalized MTLVertexFormat = 0
)

/* debug [enums.gen.go]: Processing enum MTLVertexStepFunction (5 cases) */
// MTLVertexStepFunction - The frequency with which the vertex function or post-tessellation vertex function fetches attribute data.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexStepFunction
type MTLVertexStepFunction uint

const (
	// MTLVertexStepFunctionConstant - The vertex function fetches attribute data once and uses that data for every vertex.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexStepFunction/constant
	MTLVertexStepFunctionConstant MTLVertexStepFunction = 0
	// MTLVertexStepFunctionPerInstance - The vertex function regularly fetches new attribute data for a number of instances that is determined by  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexStepFunction/perInstance
	MTLVertexStepFunctionPerInstance MTLVertexStepFunction = 0
	// MTLVertexStepFunctionPerPatch - The post-tessellation vertex function fetches data based on the patch index of the patch.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexStepFunction/perPatch
	MTLVertexStepFunctionPerPatch MTLVertexStepFunction = 0
	// MTLVertexStepFunctionPerPatchControlPoint - The post-tessellation vertex function fetches data based on the control-point indices associated with the patch.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexStepFunction/perPatchControlPoint
	MTLVertexStepFunctionPerPatchControlPoint MTLVertexStepFunction = 0
	// MTLVertexStepFunctionPerVertex - The vertex function fetches and uses new attribute data for every vertex.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexStepFunction/perVertex
	MTLVertexStepFunctionPerVertex MTLVertexStepFunction = 0
)

/* debug [enums.gen.go]: Processing enum MTLVisibilityResultMode (3 cases) */
// MTLVisibilityResultMode - The mode that determines what, if anything, the GPU writes to the results buffer, after the GPU executes the render pass.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVisibilityResultMode
type MTLVisibilityResultMode uint

const (
	// MTLVisibilityResultModeBoolean - The result records whether any samples passed depth and stencil tests.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVisibilityResultMode/boolean
	MTLVisibilityResultModeBoolean MTLVisibilityResultMode = 0
	// MTLVisibilityResultModeCounting - The result records how many samples passed depth and stencil tests.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVisibilityResultMode/counting
	MTLVisibilityResultModeCounting MTLVisibilityResultMode = 0
	// MTLVisibilityResultModeDisabled - The result doesn’t contain any data because visibility testing was disabled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVisibilityResultMode/disabled
	MTLVisibilityResultModeDisabled MTLVisibilityResultMode = 0
)

/* debug [enums.gen.go]: Processing enum MTLVisibilityResultType (2 cases) */
// MTLVisibilityResultType - This enumeration controls if Metal accumulates visibility results between render encoders or resets them.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVisibilityResultType
type MTLVisibilityResultType uint

const (
	// MTLVisibilityResultTypeAccumulate - Accumulate visibility results data across multiple render passes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVisibilityResultType/accumulate
	MTLVisibilityResultTypeAccumulate MTLVisibilityResultType = 0
	// MTLVisibilityResultTypeReset - Reset visibility result data when you create a render command encoder.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVisibilityResultType/reset
	MTLVisibilityResultTypeReset MTLVisibilityResultType = 0
)

/* debug [enums.gen.go]: Processing enum MTLWinding (2 cases) */
// MTLWinding - The vertex winding rule that determines a front-facing primitive.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLWinding
type MTLWinding uint

const (
	// MTLWindingClockwise - Primitives whose vertices are specified in clockwise order are front-facing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLWinding/clockwise
	MTLWindingClockwise MTLWinding = 0
	// MTLWindingCounterClockwise - Primitives whose vertices are specified in counter-clockwise order are front-facing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLWinding/counterClockwise
	MTLWindingCounterClockwise MTLWinding = 0
)


