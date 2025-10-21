// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

// Enum types and constants
// MTL4AlphaToOneState - Enumeration for controlling alpha-to-one state of a pipeline state object.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4AlphaToOneState
type MTL4AlphaToOneState uint

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

// MTLAccelerationStructureInstanceOptions - Options for adjusting the behavior of an instanced acceleration structure.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureInstanceOptions
type AccelerationStructureInstanceOptions uint

// MTLAccelerationStructureRefitOptions enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureRefitOptions
type AccelerationStructureRefitOptions uint

const (
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureRefitOptions/perPrimitiveData
AccelerationStructureRefitOptionPerPrimitiveData AccelerationStructureRefitOptions = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureRefitOptions/vertexData
AccelerationStructureRefitOptionVertexData AccelerationStructureRefitOptions = 0
)

// MTLAccelerationStructureUsage - Options that affect how Metal builds an acceleration structure and the behavior
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureUsage
type AccelerationStructureUsage uint

// MTLBlendFactor - The source and destination blend factors are often needed to complete specification of a blend operation. In most cases, the blend factor for both RGB values (
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendFactor
type BlendFactor uint

const (
// BlendFactorBlendAlpha - Blend factor of alpha value.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendFactor/blendAlpha
BlendFactorBlendAlpha BlendFactor = 0
// BlendFactorBlendColor - Blend factor of RGB values.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendFactor/blendColor
BlendFactorBlendColor BlendFactor = 0
// BlendFactorDestinationAlpha - Blend factor of destination alpha.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendFactor/destinationAlpha
BlendFactorDestinationAlpha BlendFactor = 0
// BlendFactorDestinationColor - Blend factor of destination values.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendFactor/destinationColor
BlendFactorDestinationColor BlendFactor = 0
// BlendFactorOne - Blend factor of one.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendFactor/one
BlendFactorOne BlendFactor = 0
// BlendFactorOneMinusBlendAlpha - Blend factor of one minus alpha value.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendFactor/oneMinusBlendAlpha
BlendFactorOneMinusBlendAlpha BlendFactor = 0
// BlendFactorOneMinusBlendColor - Blend factor of one minus RGB values.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendFactor/oneMinusBlendColor
BlendFactorOneMinusBlendColor BlendFactor = 0
// BlendFactorOneMinusDestinationAlpha - Blend factor of one minus destination alpha.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendFactor/oneMinusDestinationAlpha
BlendFactorOneMinusDestinationAlpha BlendFactor = 0
// BlendFactorOneMinusDestinationColor - Blend factor of one minus destination values.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendFactor/oneMinusDestinationColor
BlendFactorOneMinusDestinationColor BlendFactor = 0
// BlendFactorOneMinusSource1Alpha - Blend factor of one minus source alpha. This option supports dual-source blending and reads from the second color output of the fragment function.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendFactor/oneMinusSource1Alpha
BlendFactorOneMinusSource1Alpha BlendFactor = 0
// BlendFactorOneMinusSource1Color - Blend factor of one minus source values. This option supports dual-source blending and reads from the second color output of the fragment function.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendFactor/oneMinusSource1Color
BlendFactorOneMinusSource1Color BlendFactor = 0
// BlendFactorOneMinusSourceAlpha - Blend factor of one minus source alpha.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendFactor/oneMinusSourceAlpha
BlendFactorOneMinusSourceAlpha BlendFactor = 0
// BlendFactorOneMinusSourceColor - Blend factor of one minus source values.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendFactor/oneMinusSourceColor
BlendFactorOneMinusSourceColor BlendFactor = 0
// BlendFactorSource1Alpha - Blend factor of source alpha. This option supports dual-source blending and reads from the second color output of the fragment function.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendFactor/source1Alpha
BlendFactorSource1Alpha BlendFactor = 0
// BlendFactorSource1Color - Blend factor of source values. This option supports dual-source blending and reads from the second color output of the fragment function.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendFactor/source1Color
BlendFactorSource1Color BlendFactor = 0
// BlendFactorSourceAlpha - Blend factor of source alpha.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendFactor/sourceAlpha
BlendFactorSourceAlpha BlendFactor = 0
// BlendFactorSourceAlphaSaturated - Blend factor of the minimum of either source alpha or one minus destination alpha.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendFactor/sourceAlphaSaturated
BlendFactorSourceAlphaSaturated BlendFactor = 0
// BlendFactorSourceColor - Blend factor of source values.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendFactor/sourceColor
BlendFactorSourceColor BlendFactor = 0
// BlendFactorUnspecialized - Defers assigning the blend factor.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendFactor/unspecialized
BlendFactorUnspecialized BlendFactor = 0
// BlendFactorZero - Blend factor of zero.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendFactor/zero
BlendFactorZero BlendFactor = 0
)

// MTLBlendOperation - For every pixel, 
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendOperation
type BlendOperation uint

const (
// BlendOperationAdd - Add portions of both source and destination pixel values.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendOperation/add
BlendOperationAdd BlendOperation = 0
// BlendOperationMax - Maximum of the source and destination pixel values.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendOperation/max
BlendOperationMax BlendOperation = 0
// BlendOperationMin - Minimum of the source and destination pixel values.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendOperation/min
BlendOperationMin BlendOperation = 0
// BlendOperationReverseSubtract - Subtract a portion of the source values from a portion of the destination pixel values.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendOperation/reverseSubtract
BlendOperationReverseSubtract BlendOperation = 0
// BlendOperationSubtract - Subtract a portion of the destination pixel values from a portion of the source.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendOperation/subtract
BlendOperationSubtract BlendOperation = 0
// BlendOperationUnspecialized - Defers assigning the blend operation.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlendOperation/unspecialized
BlendOperationUnspecialized BlendOperation = 0
)

// MTLBlitOption - The options that enable behavior for some blit operations.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlitOption
type BlitOption uint

const (
// BlitOptionDepthFromDepthStencil - A blit option that copies the depth portion of a combined depth and stencil texture to or from a buffer.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlitOption/depthFromDepthStencil
BlitOptionDepthFromDepthStencil BlitOption = 0
)

// MTLBufferSparseTier - Enumerates the different support levels for sparse buffers.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBufferSparseTier
type BufferSparseTier uint

const (
// BufferSparseTier1 - Indicates support for sparse buffers tier 1.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBufferSparseTier/tier1
BufferSparseTier1 BufferSparseTier = 0
)

// MTLCaptureDestination - The kinds of destinations for captured command data.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCaptureDestination
type CaptureDestination uint

const (
// CaptureDestinationDeveloperTools - An option specifying that data should be captured to Xcode and that execution should stop in Xcode after the data is captured.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCaptureDestination/developerTools
CaptureDestinationDeveloperTools CaptureDestination = 0
)

// MTLCaptureError - Errors returned by capture sessions.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCaptureError
type CaptureError uint

// MTLColorWriteMask - Values used to specify a mask to permit or restrict writing to color channels of a color value.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLColorWriteMask
type ColorWriteMask uint

const (
// ColorWriteMaskNone - All color channels are disabled.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLColorWriteMask/MTLColorWriteMaskNone
ColorWriteMaskNone ColorWriteMask = 0
// ColorWriteMaskAll - All color channels are enabled.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLColorWriteMask/all
ColorWriteMaskAll ColorWriteMask = 0
// ColorWriteMaskAlpha - The alpha color channel is enabled.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLColorWriteMask/alpha
ColorWriteMaskAlpha ColorWriteMask = 0
// ColorWriteMaskBlue - The blue color channel is enabled.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLColorWriteMask/blue
ColorWriteMaskBlue ColorWriteMask = 0
// ColorWriteMaskGreen - The green color channel is enabled.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLColorWriteMask/green
ColorWriteMaskGreen ColorWriteMask = 0
// ColorWriteMaskRed - The red color channel is enabled.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLColorWriteMask/red
ColorWriteMaskRed ColorWriteMask = 0
// ColorWriteMaskUnspecialized - Defers assigning the color write mask.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLColorWriteMask/unspecialized
ColorWriteMaskUnspecialized ColorWriteMask = 0
)

// MTLCommandBufferError - Error codes that indicate why a GPU is unable to finish running a command buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandBufferError-swift.struct/Code
type CommandBufferError uint

const (
// CommandBufferErrorNotPermitted - An error code that indicates a process doesn’t have access to a GPU device.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandBufferError-swift.struct/Code/notPermitted
CommandBufferErrorNotPermitted CommandBufferError = 0
)

// MTLCommandBufferErrorOption - Options for reporting errors from a command buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandBufferErrorOption
type CommandBufferErrorOption uint

const (
// CommandBufferErrorOptionNone - An option that clears a command buffer’s error options.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandBufferErrorOption/MTLCommandBufferErrorOptionNone
CommandBufferErrorOptionNone CommandBufferErrorOption = 0
// CommandBufferErrorOptionEncoderExecutionStatus - An option that instructs a command buffer to save additional details about a GPU runtime error.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandBufferErrorOption/encoderExecutionStatus
CommandBufferErrorOptionEncoderExecutionStatus CommandBufferErrorOption = 0
)

// MTLCommandBufferStatus - The discrete states for a command buffer that represent its life cycle stages.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandBufferStatus
type CommandBufferStatus uint

const (
// CommandBufferStatusCommitted - A command buffer’s third state, which indicates the command queue is preparing to schedule the command buffer by resolving its dependencies.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandBufferStatus/committed
CommandBufferStatusCommitted CommandBufferStatus = 0
// CommandBufferStatusCompleted - A command buffer’s successful, final state, which indicates the GPU finished running the command buffer’s commands without any problems.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandBufferStatus/completed
CommandBufferStatusCompleted CommandBufferStatus = 0
// CommandBufferStatusEnqueued - A command buffer’s second state, which indicates its command queue is reserving a place for it.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandBufferStatus/enqueued
CommandBufferStatusEnqueued CommandBufferStatus = 0
// CommandBufferStatusError - A command buffer’s unsuccessful, final state, which indicates the GPU stopped running the buffer’s commands because of a runtime issue.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandBufferStatus/error
CommandBufferStatusError CommandBufferStatus = 0
// CommandBufferStatusNotEnqueued - A command buffer’s initial state, which indicates its command queue isn’t reserving a place for it.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandBufferStatus/notEnqueued
CommandBufferStatusNotEnqueued CommandBufferStatus = 0
// CommandBufferStatusScheduled - A command buffer’s fourth state, which indicates the command buffer has its resources ready and is waiting for the GPU to run its commands.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandBufferStatus/scheduled
CommandBufferStatusScheduled CommandBufferStatus = 0
)

// MTLCommandEncoderErrorState - Possible error conditions for the command encoder’s commands.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCommandEncoderErrorState
type CommandEncoderErrorState uint

// MTLCullMode - The mode that determines whether to perform culling and which type of primitive to cull.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCullMode
type CullMode uint

const (
// CullModeBack - Culls back-facing primitives.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCullMode/back
CullModeBack CullMode = 0
// CullModeFront - Culls front-facing primitives.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCullMode/front
CullModeFront CullMode = 0
// CullModeNone - Does not cull any primitives.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCullMode/none
CullModeNone CullMode = 0
)

// MTLCurveBasis enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCurveBasis
type CurveBasis uint

// MTLCurveEndCaps enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCurveEndCaps
type CurveEndCaps uint

// MTLCurveType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCurveType
type CurveType uint

// MTLDepthClipMode - The mode that determines how to deal with fragments outside of the near or far planes.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDepthClipMode
type DepthClipMode uint

const (
// DepthClipModeClamp - Clamp fragments outside the near or far planes.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDepthClipMode/clamp
DepthClipModeClamp DepthClipMode = 0
// DepthClipModeClip - Clip fragments outside the near or far planes.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDepthClipMode/clip
DepthClipModeClip DepthClipMode = 0
)

// MTLDeviceLocation - Indicates the location of the GPU relative to the system it’s connect to.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDeviceLocation
type DeviceLocation uint

const (
// DeviceLocationBuiltIn - A location that indicates the GPU is permanently connected to the system internally.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDeviceLocation/builtIn
DeviceLocationBuiltIn DeviceLocation = 0
// DeviceLocationExternal - A GPU location that indicates a person connected the GPU to the system with an external interface, such as Thunderbolt.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDeviceLocation/external
DeviceLocationExternal DeviceLocation = 0
// DeviceLocationSlot - A GPU location that indicates a person connected the GPU to a system’s internal slot.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDeviceLocation/slot
DeviceLocationSlot DeviceLocation = 0
// DeviceLocationUnspecified - A value that indicates the system can’t determine how the GPU connects to it.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDeviceLocation/unspecified
DeviceLocationUnspecified DeviceLocation = 0
)

// MTLDispatchType - The type of dispatch method to use when calling encoded functions.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDispatchType
type DispatchType uint

const (
// DispatchTypeSerial - Sets a command encoder to dispatch encoded commands serially during your pass.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDispatchType/serial
DispatchTypeSerial DispatchType = 0
)

// MTLDynamicLibraryError - Error codes that Metal can generate when creating dynamic libraries.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDynamicLibraryError-swift.struct/Code
type DynamicLibraryError uint

const (
// DynamicLibraryErrorCompilationFailure - An error code that indicates Metal couldn’t compile a dynamic library.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDynamicLibraryError-swift.struct/Code/compilationFailure
DynamicLibraryErrorCompilationFailure DynamicLibraryError = 0
// DynamicLibraryErrorDependencyLoadFailure - An error code that indicates a dynamic library couldn’t link to other dynamic libraries.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDynamicLibraryError-swift.struct/Code/dependencyLoadFailure
DynamicLibraryErrorDependencyLoadFailure DynamicLibraryError = 0
// DynamicLibraryErrorInvalidFile - An error code that indicates an app is using an invalid reference to a library file, typically related to a URL.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDynamicLibraryError-swift.struct/Code/invalidFile
DynamicLibraryErrorInvalidFile DynamicLibraryError = 0
// DynamicLibraryErrorNone - An error code that represents the absence of any problems.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDynamicLibraryError-swift.struct/Code/none
DynamicLibraryErrorNone DynamicLibraryError = 0
// DynamicLibraryErrorUnresolvedInstallName - An error code that indicates Metal couldn’t resolve the installation name for a new dynamic library.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDynamicLibraryError-swift.struct/Code/unresolvedInstallName
DynamicLibraryErrorUnresolvedInstallName DynamicLibraryError = 0
// DynamicLibraryErrorUnsupported - An error code that indicates the GPU device doesn’t support dynamic libraries.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDynamicLibraryError-swift.struct/Code/unsupported
DynamicLibraryErrorUnsupported DynamicLibraryError = 0
)

// MTLFeatureSet - The device feature sets that define specific platform, hardware, and software configurations.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet
type FeatureSet uint

const (
// FeatureSet_iOS_GPUFamily1_v1 - The GPU family 1, version 1 feature set for iOS.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/iOS_GPUFamily1_v1
FeatureSet_iOS_GPUFamily1_v1 FeatureSet = 0
// FeatureSet_iOS_GPUFamily1_v3 - The GPU family 1, version 3 feature set for iOS.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/iOS_GPUFamily1_v3
FeatureSet_iOS_GPUFamily1_v3 FeatureSet = 0
// FeatureSet_iOS_GPUFamily1_v5 - The GPU family 1, version 5 feature set for iOS.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/iOS_GPUFamily1_v5
FeatureSet_iOS_GPUFamily1_v5 FeatureSet = 0
// FeatureSet_iOS_GPUFamily2_v1 - The GPU family 2, version 1 feature set for iOS.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/iOS_GPUFamily2_v1
FeatureSet_iOS_GPUFamily2_v1 FeatureSet = 0
// FeatureSet_iOS_GPUFamily2_v2 - The GPU family 2, version 2 feature set for iOS.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/iOS_GPUFamily2_v2
FeatureSet_iOS_GPUFamily2_v2 FeatureSet = 0
// FeatureSet_iOS_GPUFamily2_v5 - The GPU family 2, version 5 feature set for iOS.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/iOS_GPUFamily2_v5
FeatureSet_iOS_GPUFamily2_v5 FeatureSet = 0
// FeatureSet_iOS_GPUFamily3_v1 - The GPU family 3, version 1 feature set for iOS.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/iOS_GPUFamily3_v1
FeatureSet_iOS_GPUFamily3_v1 FeatureSet = 0
// FeatureSet_iOS_GPUFamily3_v2 - The GPU family 3, version 2 feature set for iOS.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/iOS_GPUFamily3_v2
FeatureSet_iOS_GPUFamily3_v2 FeatureSet = 0
// FeatureSet_iOS_GPUFamily3_v3 - The GPU family 3, version 3 feature set for iOS.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/iOS_GPUFamily3_v3
FeatureSet_iOS_GPUFamily3_v3 FeatureSet = 0
// FeatureSet_iOS_GPUFamily3_v4 - The GPU family 3, version 4 feature set for iOS.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/iOS_GPUFamily3_v4
FeatureSet_iOS_GPUFamily3_v4 FeatureSet = 0
// FeatureSet_iOS_GPUFamily4_v1 - The GPU family 4, version 1 feature set for iOS.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/iOS_GPUFamily4_v1
FeatureSet_iOS_GPUFamily4_v1 FeatureSet = 0
// FeatureSet_iOS_GPUFamily4_v2 - The GPU family 4, version 2 feature set for iOS.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/iOS_GPUFamily4_v2
FeatureSet_iOS_GPUFamily4_v2 FeatureSet = 0
// FeatureSet_iOS_GPUFamily5_v1 - The GPU family 5, version 1 feature set for iOS.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/iOS_GPUFamily5_v1
FeatureSet_iOS_GPUFamily5_v1 FeatureSet = 0
// FeatureSet_macOS_GPUFamily1_v1 - The GPU family 1, version 1 feature set for macOS.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/macOS_GPUFamily1_v1
FeatureSet_macOS_GPUFamily1_v1 FeatureSet = 0
// FeatureSet_macOS_GPUFamily1_v2 - The GPU family 1, version 2 feature set for macOS.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/macOS_GPUFamily1_v2
FeatureSet_macOS_GPUFamily1_v2 FeatureSet = 0
// FeatureSet_macOS_GPUFamily1_v3 - The GPU family 1, version 3 feature set for macOS.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/macOS_GPUFamily1_v3
FeatureSet_macOS_GPUFamily1_v3 FeatureSet = 0
// FeatureSet_macOS_GPUFamily1_v4 - The GPU family 1, version 4 feature set for macOS.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/macOS_GPUFamily1_v4
FeatureSet_macOS_GPUFamily1_v4 FeatureSet = 0
// FeatureSet_macOS_GPUFamily2_v1 - The GPU family 2, version 1 feature set for macOS.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/macOS_GPUFamily2_v1
FeatureSet_macOS_GPUFamily2_v1 FeatureSet = 0
// FeatureSet_macOS_ReadWriteTextureTier2 - The read-write texture, tier 2 feature set for macOS.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/macOS_ReadWriteTextureTier2
FeatureSet_macOS_ReadWriteTextureTier2 FeatureSet = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/osx_GPUFamily1_v1
FeatureSet_OSX_GPUFamily1_v1 FeatureSet = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/osx_GPUFamily1_v2
FeatureSet_OSX_GPUFamily1_v2 FeatureSet = 0
// FeatureSet_tvOS_GPUFamily1_v2 - The GPU family 1, version 2 feature set for tvOS.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/tvOS_GPUFamily1_v2
FeatureSet_tvOS_GPUFamily1_v2 FeatureSet = 0
// FeatureSet_tvOS_GPUFamily1_v3 - The GPU family 1, version 3 feature set for tvOS.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/tvOS_GPUFamily1_v3
FeatureSet_tvOS_GPUFamily1_v3 FeatureSet = 0
// FeatureSet_tvOS_GPUFamily1_v4 - The GPU family 1, version 4 feature set for tvOS.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/tvOS_GPUFamily1_v4
FeatureSet_tvOS_GPUFamily1_v4 FeatureSet = 0
// FeatureSet_tvOS_GPUFamily2_v1 - The GPU family 2, version 1 feature set for tvOS.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/tvOS_GPUFamily2_v1
FeatureSet_tvOS_GPUFamily2_v1 FeatureSet = 0
// FeatureSet_tvOS_GPUFamily2_v2 - The GPU family 2, version 2 feature set for tvOS.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/tvOS_GPUFamily2_v2
FeatureSet_tvOS_GPUFamily2_v2 FeatureSet = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFeatureSet/tvos_GPUFamily1_v1-swift.type.property
FeatureSet_TVOS_GPUFamily1_v1 FeatureSet = 0
)

// MTLFunctionLogType - Options for different kinds of function logs.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionLogType
type FunctionLogType uint

const (
// FunctionLogTypeValidation - A message related to usage validation.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionLogType/validation
FunctionLogTypeValidation FunctionLogType = 0
)

// MTLFunctionOptions - Options that define how Metal creates the function object.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionOptions
type FunctionOptions uint

// MTLHazardTrackingMode - The options you use to specify the hazard tracking mode.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLHazardTrackingMode
type HazardTrackingMode uint

const (
// HazardTrackingModeDefault - An option specifying that the default tracking mode should be used.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLHazardTrackingMode/default
HazardTrackingModeDefault HazardTrackingMode = 0
// HazardTrackingModeUntracked - An option specifying that the app must prevent hazards when modifying this object’s contents.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLHazardTrackingMode/untracked
HazardTrackingModeUntracked HazardTrackingMode = 0
)

// MTLIOCommandQueueType - Designates the queue type for a new input/output command queue.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOCommandQueueType
type IOCommandQueueType uint

const (
// IOCommandQueueTypeConcurrent - Sets a new input/output command queue’s type to a queue that runs commands concurrently.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOCommandQueueType/concurrent
IOCommandQueueTypeConcurrent IOCommandQueueType = 0
// IOCommandQueueTypeSerial - Sets a new input/output command queue’s type to a queue that runs commands serially.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOCommandQueueType/serial
IOCommandQueueTypeSerial IOCommandQueueType = 0
)

// MTLIOCompressionMethod - The compression codecs that Metal supports for input/output handles.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOCompressionMethod
type IOCompressionMethod uint

const (
// IOCompressionMethodLZ4 - Indicates that a file uses the LZ4 compression algorithm codec.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOCompressionMethod/lz4
IOCompressionMethodLZ4 IOCompressionMethod = 0
// IOCompressionMethodLZBitmap - Indicates that a file uses the LZBitmap compression algorithm codec.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOCompressionMethod/lzBitmap
IOCompressionMethodLZBitmap IOCompressionMethod = 0
// IOCompressionMethodLZFSE - Indicates that a file uses the LZFSE compression algorithm codec.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOCompressionMethod/lzfse
IOCompressionMethodLZFSE IOCompressionMethod = 0
// IOCompressionMethodLZMA - Indicates that a file uses the LZMA compression algorithm codec.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOCompressionMethod/lzma
IOCompressionMethodLZMA IOCompressionMethod = 0
// IOCompressionMethodZlib - Indicates that a file uses the zlib compression algorithm codec.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOCompressionMethod/zlib
IOCompressionMethodZlib IOCompressionMethod = 0
)

// MTLIOCompressionStatus - Represents the final state of a compression context.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOCompressionStatus
type IOCompressionStatus uint

const (
// IOCompressionStatusComplete - Indicates the compression API successfully flushed and destroyed a compression context.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOCompressionStatus/complete
IOCompressionStatusComplete IOCompressionStatus = 0
// IOCompressionStatusError - Indicates the compression API had an error while flushing and destroying a compression context.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOCompressionStatus/error
IOCompressionStatusError IOCompressionStatus = 0
)

// MTLIOError - The error codes for creating an input/output file handle.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOError-swift.struct/Code
type IOError uint

const (
// IOErrorInternal - An error code that represents a problem internal to the Metal framework.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOError-swift.struct/Code/internal
IOErrorInternal IOError = 0
// IOErrorURLInvalid - An error code that represents a problem with a file URL.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOError-swift.struct/Code/urlInvalid
IOErrorURLInvalid IOError = 0
)

// MTLIOPriority - Designates the priority for a new input/output command queue.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOPriority
type IOPriority uint

const (
// IOPriorityHigh - Sets a new input/output command queue’s priority to a high priority.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOPriority/high
IOPriorityHigh IOPriority = 0
// IOPriorityLow - Designates the low priority for a new input/output command queue.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOPriority/low
IOPriorityLow IOPriority = 0
// IOPriorityNormal - Designates the normal priority for a new input/output command queue.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOPriority/normal
IOPriorityNormal IOPriority = 0
)

// MTLIOStatus - Represents the state of an input/output command buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOStatus
type IOStatus uint

const (
// IOStatusCancelled - Indicates the GPU has successfully abandoned the input/output command buffer.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOStatus/cancelled
IOStatusCancelled IOStatus = 0
// IOStatusComplete - Indicates the GPU has successfully finished executing the input/output command buffer.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOStatus/complete
IOStatusComplete IOStatus = 0
// IOStatusError - Indicates the GPU experienced a problem with the input/output command buffer.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOStatus/error
IOStatusError IOStatus = 0
// IOStatusPending - Indicates the GPU hasn’t finished executing the input/output command buffer.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIOStatus/pending
IOStatusPending IOStatus = 0
)

// MTLIndexType - The index type for an index buffer that references vertices of geometric primitives.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndexType
type IndexType uint

const (
// IndexTypeUInt16 - A 16-bit unsigned integer used as a primitive index.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndexType/uint16
IndexTypeUInt16 IndexType = 0
// IndexTypeUInt32 - A 32-bit unsigned integer used as a primitive index.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndexType/uint32
IndexTypeUInt32 IndexType = 0
)

// MTLIndirectCommandType - The types of commands that you can encode into the indirect command buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandType
type IndirectCommandType uint

const (
// IndirectCommandTypeConcurrentDispatch - A compute command using a grid aligned to threadgroup boundaries.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandType/concurrentDispatch
IndirectCommandTypeConcurrentDispatch IndirectCommandType = 0
// IndirectCommandTypeConcurrentDispatchThreads - A compute command using an arbitrarily sized grid.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandType/concurrentDispatchThreads
IndirectCommandTypeConcurrentDispatchThreads IndirectCommandType = 0
// IndirectCommandTypeDraw - A draw call command.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandType/draw
IndirectCommandTypeDraw IndirectCommandType = 0
// IndirectCommandTypeDrawIndexed - An indexed draw call command.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandType/drawIndexed
IndirectCommandTypeDrawIndexed IndirectCommandType = 0
// IndirectCommandTypeDrawIndexedPatches - An indexed draw call command for tessellated patches.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandType/drawIndexedPatches
IndirectCommandTypeDrawIndexedPatches IndirectCommandType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandType/drawMeshThreadgroups
IndirectCommandTypeDrawMeshThreadgroups IndirectCommandType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandType/drawMeshThreads
IndirectCommandTypeDrawMeshThreads IndirectCommandType = 0
// IndirectCommandTypeDrawPatches - A draw call command for tessellated patches.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandType/drawPatches
IndirectCommandTypeDrawPatches IndirectCommandType = 0
)

// MTLIntersectionFunctionSignature - Constants for specifying different types of custom intersection functions.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionSignature
type IntersectionFunctionSignature uint

// MTLLoadAction - Types of actions performed for an attachment at the start of a rendering pass.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLoadAction
type LoadAction uint

const (
// LoadActionClear - The GPU writes a value to every pixel in the attachment at the start of the render pass.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLoadAction/clear
LoadActionClear LoadAction = 0
// LoadActionDontCare - The GPU has permission to discard the existing contents of the attachment at the start of the render pass, replacing them with arbitrary data.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLoadAction/dontCare
LoadActionDontCare LoadAction = 0
// LoadActionLoad - The GPU preserves the existing contents of the attachment at the start of the render pass.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLoadAction/load
LoadActionLoad LoadAction = 0
)

// MTLLogLevel - The supported log levels for shader logging.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLogLevel
type LogLevel uint

const (
// LogLevelDebug - The log level that captures diagnostic information.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLogLevel/debug
LogLevelDebug LogLevel = 0
// LogLevelError - The log level that captures error information.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLogLevel/error
LogLevelError LogLevel = 0
// LogLevelFault - The log level that captures fault information.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLogLevel/fault
LogLevelFault LogLevel = 0
// LogLevelInfo - The log level that captures additional information.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLogLevel/info
LogLevelInfo LogLevel = 0
// LogLevelNotice - The log level that captures notifications.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLogLevel/notice
LogLevelNotice LogLevel = 0
// LogLevelUndefined - The log level when the log level hasn’t been configured.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLogLevel/undefined
LogLevelUndefined LogLevel = 0
)

// MTLMathFloatingPointFunctions - Indicates which FP32 math functions Metal uses.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMathFloatingPointFunctions
type MathFloatingPointFunctions uint

const (
// MathFloatingPointFunctionsFast - An indication that Metal uses the fast version of the 32b floating-point math functions.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMathFloatingPointFunctions/fast
MathFloatingPointFunctionsFast MathFloatingPointFunctions = 0
// MathFloatingPointFunctionsPrecise - An indication that Metal uses the precise version of the 32b floating-point math functions.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMathFloatingPointFunctions/precise
MathFloatingPointFunctionsPrecise MathFloatingPointFunctions = 0
)

// MTLMultisampleDepthResolveFilter - Filtering options for controlling an MSAA depth resolve operation.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMultisampleDepthResolveFilter
type MultisampleDepthResolveFilter uint

const (
// MultisampleDepthResolveFilterMax - The GPU compares all depth samples in the pixel and selects the sample with the largest value.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMultisampleDepthResolveFilter/max
MultisampleDepthResolveFilterMax MultisampleDepthResolveFilter = 0
// MultisampleDepthResolveFilterMin - The GPU compares all depth samples in the pixel and selects the sample with the smallest value.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMultisampleDepthResolveFilter/min
MultisampleDepthResolveFilterMin MultisampleDepthResolveFilter = 0
// MultisampleDepthResolveFilterSample0 - No filter is applied.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMultisampleDepthResolveFilter/sample0
MultisampleDepthResolveFilterSample0 MultisampleDepthResolveFilter = 0
)

// MTLMultisampleStencilResolveFilter - Constants used to control the multisample stencil resolve operation.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMultisampleStencilResolveFilter
type MultisampleStencilResolveFilter uint

const (
// MultisampleStencilResolveFilterDepthResolvedSample - Chooses the stencil sample corresponding to the depth sample selected by the depth resolve filter.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMultisampleStencilResolveFilter/depthResolvedSample
MultisampleStencilResolveFilterDepthResolvedSample MultisampleStencilResolveFilter = 0
// MultisampleStencilResolveFilterSample0 - Chooses the first stencil sample in the pixel.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMultisampleStencilResolveFilter/sample0
MultisampleStencilResolveFilterSample0 MultisampleStencilResolveFilter = 0
)

// MTLMutability - The options that determine the mutability of a buffer’s contents.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMutability
type Mutability uint

const (
// MutabilityDefault - The default behavior, based on the buffer’s type.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMutability/default
MutabilityDefault Mutability = 0
// MutabilityImmutable - An option that states that you can’t modify the buffer’s contents.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMutability/immutable
MutabilityImmutable Mutability = 0
// MutabilityMutable - An option that states that you can modify the buffer’s contents.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMutability/mutable
MutabilityMutable Mutability = 0
)

// MTLPipelineOption - Options that determine how Metal prepares the pipeline.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPipelineOption
type PipelineOption uint

const (
// PipelineOptionNone - Don’t provide any reflection information.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPipelineOption/MTLPipelineOptionNone
PipelineOptionNone PipelineOption = 0
// PipelineOptionArgumentInfo - An option instance that provides argument information for textures and threadgroup memory.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPipelineOption/argumentInfo
PipelineOptionArgumentInfo PipelineOption = 0
// PipelineOptionBindingInfo - An option that provides binding information for pipeline state resources.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPipelineOption/bindingInfo
PipelineOptionBindingInfo PipelineOption = 0
// PipelineOptionBufferTypeInfo - An option instance that provides detailed buffer type information for buffer arguments.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPipelineOption/bufferTypeInfo
PipelineOptionBufferTypeInfo PipelineOption = 0
// PipelineOptionFailOnBinaryArchiveMiss - An option that specifies that Metal only creates the pipeline state object if the compiled shader is present inside a linked binary archive.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPipelineOption/failOnBinaryArchiveMiss
PipelineOptionFailOnBinaryArchiveMiss PipelineOption = 0
)

// MTLPixelFormat - The data formats that describe the organization and characteristics of individual pixels in a texture.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat
type PixelFormat uint

const (
// PixelFormatA8Unorm - Ordinary format with one 8-bit normalized unsigned integer component.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/a8Unorm
PixelFormatA8Unorm PixelFormat = 0
// PixelFormatASTC_4x4_sRGB - ASTC-compressed format with low-dynamic-range content, conversion between sRGB and linear space, a block width of 4, and a block height of 4.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/astc_4x4_srgb
PixelFormatASTC_4x4_sRGB PixelFormat = 0
// PixelFormatBC1_RGBA - Compressed format with two 16-bit color components and one 32-bit descriptor component.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/bc1_rgba
PixelFormatBC1_RGBA PixelFormat = 0
// PixelFormatBGRA8Unorm - Ordinary format with four 8-bit normalized unsigned integer components in BGRA order.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/bgra8Unorm
PixelFormatBGRA8Unorm PixelFormat = 0
// PixelFormatDepth24Unorm_Stencil8 - A 32-bit combined depth and stencil pixel format with a 24-bit normalized unsigned integer for depth and an 8-bit unsigned integer for stencil.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/depth24Unorm_stencil8
PixelFormatDepth24Unorm_Stencil8 PixelFormat = 0
// PixelFormatInvalid - The default value of the pixel format for the  . You cannot create a texture with this value.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/invalid
PixelFormatInvalid PixelFormat = 0
// PixelFormatR8Sint - Ordinary format with one 8-bit signed integer component.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/r8Sint
PixelFormatR8Sint PixelFormat = 0
// PixelFormatR8Snorm - Ordinary format with one 8-bit normalized signed integer component.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/r8Snorm
PixelFormatR8Snorm PixelFormat = 0
// PixelFormatR8Uint - Ordinary format with one 8-bit unsigned integer component.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/r8Uint
PixelFormatR8Uint PixelFormat = 0
// PixelFormatR8Unorm - Ordinary format with one 8-bit normalized unsigned integer component.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/r8Unorm
PixelFormatR8Unorm PixelFormat = 0
// PixelFormatR8Unorm_sRGB - Ordinary format with one 8-bit normalized unsigned integer component with conversion between sRGB and linear space.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/r8Unorm_srgb
PixelFormatR8Unorm_sRGB PixelFormat = 0
// PixelFormatRGBA8Unorm - Ordinary format with four 8-bit normalized unsigned integer components in RGBA order.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/rgba8Unorm
PixelFormatRGBA8Unorm PixelFormat = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/unspecialized
PixelFormatUnspecialized PixelFormat = 0
)

// MTLPrimitiveTopologyClass - The primitive topologies available for rendering.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPrimitiveTopologyClass
type PrimitiveTopologyClass uint

const (
// PrimitiveTopologyClassLine - A line primitive.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPrimitiveTopologyClass/line
PrimitiveTopologyClassLine PrimitiveTopologyClass = 0
// PrimitiveTopologyClassPoint - A point primitive.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPrimitiveTopologyClass/point
PrimitiveTopologyClassPoint PrimitiveTopologyClass = 0
// PrimitiveTopologyClassTriangle - A triangle primitive.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPrimitiveTopologyClass/triangle
PrimitiveTopologyClassTriangle PrimitiveTopologyClass = 0
// PrimitiveTopologyClassUnspecified - An unspecified primitive.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPrimitiveTopologyClass/unspecified
PrimitiveTopologyClassUnspecified PrimitiveTopologyClass = 0
)

// MTLPrimitiveType - The geometric primitive type for drawing commands.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPrimitiveType
type PrimitiveType uint

const (
// PrimitiveTypeLine - Rasterize a line between each separate pair of vertices, resulting in a series of unconnected lines. If there are an odd number of vertices, the last vertex is ignored.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPrimitiveType/line
PrimitiveTypeLine PrimitiveType = 0
// PrimitiveTypeLineStrip - Rasterize a line between each pair of adjacent vertices, resulting in a series of connected lines (also called a polyline).
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPrimitiveType/lineStrip
PrimitiveTypeLineStrip PrimitiveType = 0
// PrimitiveTypePoint - Rasterize a point at each vertex. The vertex shader must provide  , or the point size is undefined.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPrimitiveType/point
PrimitiveTypePoint PrimitiveType = 0
// PrimitiveTypeTriangle - For every separate set of three vertices, rasterize a triangle. If the number of vertices is not a multiple of three, either one or two vertices is ignored.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPrimitiveType/triangle
PrimitiveTypeTriangle PrimitiveType = 0
// PrimitiveTypeTriangleStrip - For every three adjacent vertices, rasterize a triangle.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPrimitiveType/triangleStrip
PrimitiveTypeTriangleStrip PrimitiveType = 0
)

// MTLPurgeableState - The purgeable state of the resource.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPurgeableState
type PurgeableState uint

const (
// PurgeableStateEmpty - A state that indicates to the system that it needs to consider   the contents of a resource as invalid, typically because you’re discarding it.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPurgeableState/empty
PurgeableStateEmpty PurgeableState = 0
// PurgeableStateKeepCurrent - The current state is queried but doesn’t change.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPurgeableState/keepCurrent
PurgeableStateKeepCurrent PurgeableState = 0
// PurgeableStateNonVolatile - The contents of the resource aren’t allowed to be discarded.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPurgeableState/nonVolatile
PurgeableStateNonVolatile PurgeableState = 0
// PurgeableStateVolatile - The system is allowed to discard the resource to free up memory.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPurgeableState/volatile
PurgeableStateVolatile PurgeableState = 0
)

// MTLReadWriteTextureTier - The support level for read-write texture formats.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLReadWriteTextureTier
type ReadWriteTextureTier uint

// MTLResourceOptions - Optional arguments used to set the behavior of a resource.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceOptions
type ResourceOptions uint

const (
// ResourceOptionCPUCacheModeWriteCombined - This constant was deprecated in iOS 9.0 and macOS 10.11.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceOptions/optionCPUCacheModeWriteCombined
ResourceOptionCPUCacheModeWriteCombined ResourceOptions = 0
// ResourceStorageModePrivate - The resource is only available to the GPU.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceOptions/storageModePrivate
ResourceStorageModePrivate ResourceOptions = 0
)

// MTLResourceUsage - Options that describe how a graphics or compute function uses an argument buffer’s resource.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceUsage
type ResourceUsage uint

const (
// ResourceUsageRead - An option that enables reading from the resource.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceUsage/read
ResourceUsageRead ResourceUsage = 0
)

// MTLSamplerReductionMode - Configures how the sampler aggregates contributing samples to a final value.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerReductionMode
type SamplerReductionMode uint

const (
// SamplerReductionModeMinimum - A reduction mode that finds the minimum contributing sample value by separately evaluating each channel.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplerReductionMode/minimum
SamplerReductionModeMinimum SamplerReductionMode = 0
)

// MTLShaderValidation - Indicates whether shader validation in an enabled or disabled state, or neither state.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLShaderValidation
type ShaderValidation uint

// MTLSparsePageSize - The page size options, in kilobytes, for sparse textures.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSparsePageSize
type SparsePageSize uint

const (
// SparsePageSize16 - Represents a sparse texture’s page size of 16 kilobytes.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSparsePageSize/size16
SparsePageSize16 SparsePageSize = 0
// SparsePageSize256 - Represents a sparse texture’s page size of 256 kilobytes.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSparsePageSize/size256
SparsePageSize256 SparsePageSize = 0
// SparsePageSize64 - Represents a sparse texture’s page size of 64 kilobytes.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSparsePageSize/size64
SparsePageSize64 SparsePageSize = 0
)

// MTLSparseTextureRegionAlignmentMode - Options used when converting between a pixel-based region within a texture to a tile-based region.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSparseTextureRegionAlignmentMode
type SparseTextureRegionAlignmentMode uint

const (
// SparseTextureRegionAlignmentModeOutward - The tile region includes any partially covered tiles.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSparseTextureRegionAlignmentMode/outward
SparseTextureRegionAlignmentModeOutward SparseTextureRegionAlignmentMode = 0
)

// MTLStencilOperation - The operation performed on a currently stored stencil value when a comparison test passes or fails.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStencilOperation
type StencilOperation uint

const (
// StencilOperationDecrementClamp - If the current stencil value is not zero, decrease the stencil value by one. Otherwise, if the current stencil value is zero, do not change the stencil value.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStencilOperation/decrementClamp
StencilOperationDecrementClamp StencilOperation = 0
// StencilOperationDecrementWrap - If the current stencil value is not zero, decrease the stencil value by one. Otherwise, if the current stencil value is zero, set the stencil value to the maximum representable value.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStencilOperation/decrementWrap
StencilOperationDecrementWrap StencilOperation = 0
// StencilOperationIncrementClamp - If the current stencil value is not the maximum representable value, increase the stencil value by one. Otherwise, if the current stencil value is the maximum representable value, do not change the stencil value.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStencilOperation/incrementClamp
StencilOperationIncrementClamp StencilOperation = 0
// StencilOperationIncrementWrap - If the current stencil value is not the maximum representable value, increase the stencil value by one. Otherwise, if the current stencil value is the maximum representable value, set the stencil value to zero.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStencilOperation/incrementWrap
StencilOperationIncrementWrap StencilOperation = 0
// StencilOperationInvert - Perform a logical bitwise invert operation on the current stencil value.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStencilOperation/invert
StencilOperationInvert StencilOperation = 0
// StencilOperationKeep - Keep the current stencil value.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStencilOperation/keep
StencilOperationKeep StencilOperation = 0
// StencilOperationReplace - Replace the stencil value with the stencil reference value, which is set by the   method of  .
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStencilOperation/replace
StencilOperationReplace StencilOperation = 0
// StencilOperationZero - Set the stencil value to zero.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStencilOperation/zero
StencilOperationZero StencilOperation = 0
)

// MTLStitchedLibraryOptions enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStitchedLibraryOptions
type StitchedLibraryOptions uint

const (
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStitchedLibraryOptions/MTLStitchedLibraryOptionNone
StitchedLibraryOptionNone StitchedLibraryOptions = 0
)

// MTLStorageMode - Options for the memory location and access permissions for a resource.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStorageMode
type StorageMode uint

const (
// StorageModeManaged - The CPU and GPU may maintain separate copies of the resource, and any changes must be explicitly synchronized.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStorageMode/managed
StorageModeManaged StorageMode = 0
// StorageModeMemoryless - The resource’s contents are only available to the GPU, and only exist temporarily during a render pass.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStorageMode/memoryless
StorageModeMemoryless StorageMode = 0
// StorageModePrivate - The resource is only available to the GPU.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStorageMode/private
StorageModePrivate StorageMode = 0
// StorageModeShared - The CPU and GPU share access to the resource, allocated in system memory.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStorageMode/shared
StorageModeShared StorageMode = 0
)

// MTLStoreAction - Types of actions performed for an attachment at the end of a rendering pass.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStoreAction
type StoreAction uint

const (
// StoreActionCustomSampleDepthStore - The GPU stores depth data in a sample-position–agnostic representation.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStoreAction/customSampleDepthStore
StoreActionCustomSampleDepthStore StoreAction = 0
// StoreActionDontCare - The GPU has permission to discard the rendered contents of the attachment at the end of the render pass, replacing them with arbitrary data.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStoreAction/dontCare
StoreActionDontCare StoreAction = 0
// StoreActionMultisampleResolve - The GPU resolves the multisampled data to one sample per pixel and stores the data to the resolve texture, discarding the multisample data afterwards.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStoreAction/multisampleResolve
StoreActionMultisampleResolve StoreAction = 0
// StoreActionStore - The GPU stores the rendered contents to the texture.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStoreAction/store
StoreActionStore StoreAction = 0
// StoreActionStoreAndMultisampleResolve - The GPU stores the multisample data to the multisample texture, resolves the data to a sample per pixel, and stores the data to the resolve texture.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStoreAction/storeAndMultisampleResolve
StoreActionStoreAndMultisampleResolve StoreAction = 0
// StoreActionUnknown - The system selects a store action when it encodes the render pass.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStoreAction/unknown
StoreActionUnknown StoreAction = 0
)

// MTLStoreActionOptions - Options that modify a store action.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStoreActionOptions
type StoreActionOptions uint

const (
// StoreActionOptionNone - An option that doesn’t modify the intended behavior of a store action.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStoreActionOptions/MTLStoreActionOptionNone
StoreActionOptionNone StoreActionOptions = 0
// StoreActionOptionCustomSamplePositions - An option that stores data in a sample-position–agnostic representation.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStoreActionOptions/customSamplePositions
StoreActionOptionCustomSamplePositions StoreActionOptions = 0
)

// MTLTensorDataType - The possible data types for the elements of a tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorDataType
type TensorDataType uint

const (
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorDataType/bfloat16
TensorDataTypeBFloat16 TensorDataType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorDataType/int16
TensorDataTypeInt16 TensorDataType = 0
)

// MTLTensorError - The error codes that Metal can raise when you create a tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorError-swift.struct/Code
type TensorError uint

const (
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorError-swift.struct/Code/none
TensorErrorNone TensorError = 0
)

// MTLTensorUsage - The type that represents the different contexts for a tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorUsage
type TensorUsage uint

const (
// TensorUsageCompute - A tensor context that applies to compute encoders.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorUsage/compute
TensorUsageCompute TensorUsage = 0
// TensorUsageMachineLearning - A tensor context that applies to machine learning encoders.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorUsage/machineLearning
TensorUsageMachineLearning TensorUsage = 0
)

// MTLTessellationControlPointIndexType - Options for specifying the size of the control point indices in a control point index buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTessellationControlPointIndexType
type TessellationControlPointIndexType uint

const (
// TessellationControlPointIndexTypeNone - No size. This value should only be used when drawing patches without a control point index buffer.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTessellationControlPointIndexType/none
TessellationControlPointIndexTypeNone TessellationControlPointIndexType = 0
// TessellationControlPointIndexTypeUInt16 - The size of a 16-bit unsigned integer.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTessellationControlPointIndexType/uint16
TessellationControlPointIndexTypeUInt16 TessellationControlPointIndexType = 0
// TessellationControlPointIndexTypeUInt32 - The size of a 32-bit unsigned integer.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTessellationControlPointIndexType/uint32
TessellationControlPointIndexTypeUInt32 TessellationControlPointIndexType = 0
)

// MTLTessellationFactorFormat - Options for specifying the format of the tessellation factors in a tessellation factor buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTessellationFactorFormat
type TessellationFactorFormat uint

const (
// TessellationFactorFormatHalf - A 16-bit floating-point format.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTessellationFactorFormat/half
TessellationFactorFormatHalf TessellationFactorFormat = 0
)

// MTLTessellationFactorStepFunction - Options for specifying the step function that determines the tessellation factors for a patch from the tessellation factor buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTessellationFactorStepFunction
type TessellationFactorStepFunction uint

const (
// TessellationFactorStepFunctionConstant - A constant step function. For all instances, the tessellation factor for all patches in a patch draw call is at the   location in the tessellation factor buffer.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTessellationFactorStepFunction/constant
TessellationFactorStepFunctionConstant TessellationFactorStepFunction = 0
// TessellationFactorStepFunctionPerInstance - A per-instance step function. For a given instance ID, the tessellation factor for a patch in a patch draw call is at the   location in the tessellation factor buffer.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTessellationFactorStepFunction/perInstance
TessellationFactorStepFunctionPerInstance TessellationFactorStepFunction = 0
// TessellationFactorStepFunctionPerPatch - A per-patch step function. For all instances, the tessellation factor for all patches in a patch draw call is at the   location in the tessellation factor buffer.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTessellationFactorStepFunction/perPatch
TessellationFactorStepFunctionPerPatch TessellationFactorStepFunction = 0
// TessellationFactorStepFunctionPerPatchAndPerInstance - A per-patch and per-instance step function. For a given instance ID, the tessellation factor for a patch in a patch draw call is at the   location in the tessellation factor buffer.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTessellationFactorStepFunction/perPatchAndPerInstance
TessellationFactorStepFunctionPerPatchAndPerInstance TessellationFactorStepFunction = 0
)

// MTLTessellationPartitionMode - Options for choosing the partition mode that the tessellator applies when deriving
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTessellationPartitionMode
type TessellationPartitionMode uint

const (
// TessellationPartitionModeFractionalEven - A fractional even partitioning mode.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTessellationPartitionMode/fractionalEven
TessellationPartitionModeFractionalEven TessellationPartitionMode = 0
// TessellationPartitionModeFractionalOdd - A fractional odd partitioning mode.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTessellationPartitionMode/fractionalOdd
TessellationPartitionModeFractionalOdd TessellationPartitionMode = 0
// TessellationPartitionModeInteger - An integer partitioning mode.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTessellationPartitionMode/integer
TessellationPartitionModeInteger TessellationPartitionMode = 0
// TessellationPartitionModePow2 - A power of two partitioning mode.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTessellationPartitionMode/pow2
TessellationPartitionModePow2 TessellationPartitionMode = 0
)

// MTLTextureCompressionType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureCompressionType
type TextureCompressionType uint

const (
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureCompressionType/lossy
TextureCompressionTypeLossy TextureCompressionType = 0
)

// MTLTextureSparseTier - Enumerates the different support levels for sparse textures.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureSparseTier
type TextureSparseTier uint

const (
// TextureSparseTierNone - Indicates that the texture is not sparse.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureSparseTier/tierNone
TextureSparseTierNone TextureSparseTier = 0
)

// MTLTextureSwizzle - A set of options to choose from when creating a texture swizzle pattern.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureSwizzle
type TextureSwizzle uint

const (
// TextureSwizzleAlpha - The alpha channel of the source pixel is copied to the destination channel.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureSwizzle/alpha
TextureSwizzleAlpha TextureSwizzle = 0
// TextureSwizzleBlue - The blue channel of the source pixel is copied to the destination channel.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureSwizzle/blue
TextureSwizzleBlue TextureSwizzle = 0
// TextureSwizzleGreen - The green channel of the source pixel is copied to the destination channel.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureSwizzle/green
TextureSwizzleGreen TextureSwizzle = 0
// TextureSwizzleRed - The red channel of the source pixel is copied to the destination channel.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureSwizzle/red
TextureSwizzleRed TextureSwizzle = 0
)

// MTLTextureType - The dimension of each image, including whether multiple images are arranged into an array or a cube.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureType
type TextureType uint

const (
// TextureType1D - A one-dimensional texture image.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureType/type1D
TextureType1D TextureType = 0
// TextureType1DArray - An array of one-dimensional texture images.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureType/type1DArray
TextureType1DArray TextureType = 0
// TextureType2D - A two-dimensional texture image.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureType/type2D
TextureType2D TextureType = 0
// TextureType2DArray - An array of two-dimensional texture images.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureType/type2DArray
TextureType2DArray TextureType = 0
// TextureType2DMultisample - A two-dimensional texture image that uses more than one sample for each pixel.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureType/type2DMultisample
TextureType2DMultisample TextureType = 0
// TextureType2DMultisampleArray - An array of two-dimensional texture images that use more than one sample for each pixel.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureType/type2DMultisampleArray
TextureType2DMultisampleArray TextureType = 0
// TextureType3D - A three-dimensional texture image.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureType/type3D
TextureType3D TextureType = 0
// TextureTypeCube - A cube texture with six two-dimensional images.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureType/typeCube
TextureTypeCube TextureType = 0
// TextureTypeCubeArray - An array of cube textures, each with six two-dimensional images.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureType/typeCubeArray
TextureTypeCubeArray TextureType = 0
// TextureTypeTextureBuffer - A texture buffer.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureType/typeTextureBuffer
TextureTypeTextureBuffer TextureType = 0
)

// MTLTextureUsage - An enumeration for the various options that determine how you can use a texture.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureUsage
type TextureUsage uint

const (
// TextureUsagePixelFormatView - An option to create texture views with a different component layout.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureUsage/pixelFormatView
TextureUsagePixelFormatView TextureUsage = 0
// TextureUsageRenderTarget - An option for rendering to the texture in a render pass.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureUsage/renderTarget
TextureUsageRenderTarget TextureUsage = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureUsage/shaderAtomic
TextureUsageShaderAtomic TextureUsage = 0
// TextureUsageShaderRead - An option for reading or sampling from the texture in a shader.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureUsage/shaderRead
TextureUsageShaderRead TextureUsage = 0
// TextureUsageShaderWrite - An option for writing to the texture in a shader.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureUsage/shaderWrite
TextureUsageShaderWrite TextureUsage = 0
// TextureUsageUnknown - An option for a texture whose usage is unknown.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureUsage/unknown
TextureUsageUnknown TextureUsage = 0
)

// MTLTransformType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTransformType
type TransformType uint

// MTLTriangleFillMode - Specifies how to rasterize triangle and triangle strip primitives.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTriangleFillMode
type TriangleFillMode uint

const (
// TriangleFillModeFill - Rasterize triangle and triangle strip primitives as filled triangles.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTriangleFillMode/fill
TriangleFillModeFill TriangleFillMode = 0
// TriangleFillModeLines - Rasterize triangle and triangle strip primitives as lines.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTriangleFillMode/lines
TriangleFillModeLines TriangleFillMode = 0
)

// MTLVertexFormat - Values that specify the organization of function vertex data.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat
type VertexFormat uint

const (
// VertexFormatChar - One signed 8-bit two’s complement value.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/char
VertexFormatChar VertexFormat = 0
// VertexFormatChar2 - Two signed 8-bit two’s complement values.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/char2
VertexFormatChar2 VertexFormat = 0
// VertexFormatChar2Normalized - Two signed normalized 8-bit two’s complement values.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/char2Normalized
VertexFormatChar2Normalized VertexFormat = 0
// VertexFormatChar3 - Three signed 8-bit two’s complement values.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/char3
VertexFormatChar3 VertexFormat = 0
// VertexFormatChar3Normalized - Three signed normalized 8-bit two’s complement values.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/char3Normalized
VertexFormatChar3Normalized VertexFormat = 0
// VertexFormatChar4 - Four signed 8-bit two’s complement values.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/char4
VertexFormatChar4 VertexFormat = 0
// VertexFormatChar4Normalized - Four signed normalized 8-bit two’s complement values.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/char4Normalized
VertexFormatChar4Normalized VertexFormat = 0
// VertexFormatCharNormalized - One signed normalized 8-bit two’s complement value.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/charNormalized
VertexFormatCharNormalized VertexFormat = 0
// VertexFormatFloat - One single-precision floating-point value.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/float
VertexFormatFloat VertexFormat = 0
// VertexFormatFloat2 - Two single-precision floating-point values.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/float2
VertexFormatFloat2 VertexFormat = 0
// VertexFormatFloat3 - Three single-precision floating-point values.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/float3
VertexFormatFloat3 VertexFormat = 0
// VertexFormatFloat4 - Four single-precision floating-point values.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/float4
VertexFormatFloat4 VertexFormat = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/floatRG11B10
VertexFormatFloatRG11B10 VertexFormat = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/floatRGB9E5
VertexFormatFloatRGB9E5 VertexFormat = 0
// VertexFormatHalf - One half-precision floating-point value.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/half
VertexFormatHalf VertexFormat = 0
// VertexFormatHalf2 - Two half-precision floating-point values.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/half2
VertexFormatHalf2 VertexFormat = 0
// VertexFormatHalf3 - Three half-precision floating-point values.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/half3
VertexFormatHalf3 VertexFormat = 0
// VertexFormatHalf4 - Four half-precision floating-point values.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/half4
VertexFormatHalf4 VertexFormat = 0
// VertexFormatInt - One signed 32-bit two’s complement value.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/int
VertexFormatInt VertexFormat = 0
// VertexFormatInt1010102Normalized - One packed 32-bit value with four normalized signed two’s complement integer values, arranged as 10 bits, 10 bits, 10 bits, and 2 bits.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/int1010102Normalized
VertexFormatInt1010102Normalized VertexFormat = 0
// VertexFormatInt2 - Two signed 32-bit two’s complement values.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/int2
VertexFormatInt2 VertexFormat = 0
// VertexFormatInt3 - Three signed 32-bit two’s complement values.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/int3
VertexFormatInt3 VertexFormat = 0
// VertexFormatInt4 - Four signed 32-bit two’s complement values.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/int4
VertexFormatInt4 VertexFormat = 0
// VertexFormatInvalid - An invalid vertex format.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/invalid
VertexFormatInvalid VertexFormat = 0
// VertexFormatShort - One signed 16-bit two’s complement value.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/short
VertexFormatShort VertexFormat = 0
// VertexFormatShort2 - Two signed 16-bit two’s complement values.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/short2
VertexFormatShort2 VertexFormat = 0
// VertexFormatShort2Normalized - Two signed normalized 16-bit two’s complement values.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/short2Normalized
VertexFormatShort2Normalized VertexFormat = 0
// VertexFormatShort3 - Three signed 16-bit two’s complement values.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/short3
VertexFormatShort3 VertexFormat = 0
// VertexFormatShort3Normalized - Three signed normalized 16-bit two’s complement values.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/short3Normalized
VertexFormatShort3Normalized VertexFormat = 0
// VertexFormatShort4 - Four signed 16-bit two’s complement values.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/short4
VertexFormatShort4 VertexFormat = 0
// VertexFormatShort4Normalized - Four signed normalized 16-bit two’s complement values.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/short4Normalized
VertexFormatShort4Normalized VertexFormat = 0
// VertexFormatShortNormalized - One signed normalized 16-bit two’s complement value.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/shortNormalized
VertexFormatShortNormalized VertexFormat = 0
// VertexFormatUChar - One unsigned 8-bit value.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/uchar
VertexFormatUChar VertexFormat = 0
// VertexFormatUChar2 - Two unsigned 8-bit values.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/uchar2
VertexFormatUChar2 VertexFormat = 0
// VertexFormatUChar2Normalized - Two unsigned normalized 8-bit values.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/uchar2Normalized
VertexFormatUChar2Normalized VertexFormat = 0
// VertexFormatUChar3 - Three unsigned 8-bit values.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/uchar3
VertexFormatUChar3 VertexFormat = 0
// VertexFormatUChar3Normalized - Three unsigned normalized 8-bit values.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/uchar3Normalized
VertexFormatUChar3Normalized VertexFormat = 0
// VertexFormatUChar4 - Four unsigned 8-bit values.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/uchar4
VertexFormatUChar4 VertexFormat = 0
// VertexFormatUChar4Normalized - Four unsigned normalized 8-bit values.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/uchar4Normalized
VertexFormatUChar4Normalized VertexFormat = 0
// VertexFormatUChar4Normalized_BGRA - Four unsigned normalized 8-bit values, arranged as blue, green, red, and alpha components.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/uchar4Normalized_bgra
VertexFormatUChar4Normalized_BGRA VertexFormat = 0
// VertexFormatUCharNormalized - One unsigned normalized 8-bit value.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/ucharNormalized
VertexFormatUCharNormalized VertexFormat = 0
// VertexFormatUInt - One unsigned 32-bit value.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/uint
VertexFormatUInt VertexFormat = 0
// VertexFormatUInt1010102Normalized - One packed 32-bit value with four normalized unsigned integer values, arranged as 10 bits, 10 bits, 10 bits, and 2 bits.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/uint1010102Normalized
VertexFormatUInt1010102Normalized VertexFormat = 0
// VertexFormatUInt2 - Two unsigned 32-bit values.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/uint2
VertexFormatUInt2 VertexFormat = 0
// VertexFormatUInt3 - Three unsigned 32-bit values.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/uint3
VertexFormatUInt3 VertexFormat = 0
// VertexFormatUInt4 - Four unsigned 32-bit values.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/uint4
VertexFormatUInt4 VertexFormat = 0
// VertexFormatUShort - One unsigned 16-bit value.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/ushort
VertexFormatUShort VertexFormat = 0
// VertexFormatUShort2 - Two unsigned 16-bit values.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/ushort2
VertexFormatUShort2 VertexFormat = 0
// VertexFormatUShort2Normalized - Two unsigned normalized 16-bit values.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/ushort2Normalized
VertexFormatUShort2Normalized VertexFormat = 0
// VertexFormatUShort3 - Three unsigned 16-bit values.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/ushort3
VertexFormatUShort3 VertexFormat = 0
// VertexFormatUShort3Normalized - Three unsigned normalized 16-bit values.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/ushort3Normalized
VertexFormatUShort3Normalized VertexFormat = 0
// VertexFormatUShort4 - Four unsigned 16-bit values.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/ushort4
VertexFormatUShort4 VertexFormat = 0
// VertexFormatUShort4Normalized - Four unsigned normalized 16-bit values.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/ushort4Normalized
VertexFormatUShort4Normalized VertexFormat = 0
// VertexFormatUShortNormalized - One unsigned normalized 16-bit value.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexFormat/ushortNormalized
VertexFormatUShortNormalized VertexFormat = 0
)

// MTLVertexStepFunction - The frequency with which the vertex function or post-tessellation vertex function fetches attribute data.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexStepFunction
type VertexStepFunction uint

const (
// VertexStepFunctionConstant - The vertex function fetches attribute data once and uses that data for every vertex.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexStepFunction/constant
VertexStepFunctionConstant VertexStepFunction = 0
// VertexStepFunctionPerInstance - The vertex function regularly fetches new attribute data for a number of instances that is determined by  .
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexStepFunction/perInstance
VertexStepFunctionPerInstance VertexStepFunction = 0
// VertexStepFunctionPerPatch - The post-tessellation vertex function fetches data based on the patch index of the patch.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexStepFunction/perPatch
VertexStepFunctionPerPatch VertexStepFunction = 0
// VertexStepFunctionPerPatchControlPoint - The post-tessellation vertex function fetches data based on the control-point indices associated with the patch.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexStepFunction/perPatchControlPoint
VertexStepFunctionPerPatchControlPoint VertexStepFunction = 0
// VertexStepFunctionPerVertex - The vertex function fetches and uses new attribute data for every vertex.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexStepFunction/perVertex
VertexStepFunctionPerVertex VertexStepFunction = 0
)

// MTLVisibilityResultMode - The mode that determines what, if anything, the GPU writes to the results buffer, after the GPU executes the render pass.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVisibilityResultMode
type VisibilityResultMode uint

const (
// VisibilityResultModeBoolean - The result records whether any samples passed depth and stencil tests.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVisibilityResultMode/boolean
VisibilityResultModeBoolean VisibilityResultMode = 0
// VisibilityResultModeCounting - The result records how many samples passed depth and stencil tests.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVisibilityResultMode/counting
VisibilityResultModeCounting VisibilityResultMode = 0
// VisibilityResultModeDisabled - The result doesn’t contain any data because visibility testing was disabled.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVisibilityResultMode/disabled
VisibilityResultModeDisabled VisibilityResultMode = 0
)

// MTLVisibilityResultType - This enumeration controls if Metal accumulates visibility results between render encoders or resets them.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVisibilityResultType
type VisibilityResultType uint

const (
// VisibilityResultTypeAccumulate - Accumulate visibility results data across multiple render passes.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVisibilityResultType/accumulate
VisibilityResultTypeAccumulate VisibilityResultType = 0
// VisibilityResultTypeReset - Reset visibility result data when you create a render command encoder.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVisibilityResultType/reset
VisibilityResultTypeReset VisibilityResultType = 0
)

// MTLWinding - The vertex winding rule that determines a front-facing primitive.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLWinding
type Winding uint

const (
// WindingClockwise - Primitives whose vertices are specified in clockwise order are front-facing.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLWinding/clockwise
WindingClockwise Winding = 0
// WindingCounterClockwise - Primitives whose vertices are specified in counter-clockwise order are front-facing.
//
	// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLWinding/counterClockwise
WindingCounterClockwise Winding = 0
)


