// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal
import (

	"github.com/tmc/appledocs/generated/foundation"
)


// C struct types
// MTL4BufferRange
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4BufferRange
type MTL4BufferRange struct {
	BufferAddress GPUAddress
	Length uint64
}

// MTL4CopySparseBufferMappingOperation - Groups together arguments for an operation to copy a sparse buffer mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CopySparseBufferMappingOperation
type MTL4CopySparseBufferMappingOperation struct {
	DestinationOffset uint // The origin in the destination buffer, in tiles.
	SourceRange foundation.Range // The range in the source buffer, in tiles.
}

// MTL4CopySparseTextureMappingOperation - Groups together arguments for an operation to copy a sparse texture mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CopySparseTextureMappingOperation
type MTL4CopySparseTextureMappingOperation struct {
	DestinationLevel uint // The index of the mipmap level in the destination texture.
	DestinationOrigin Origin // The origin in the destination texture to copy into, in tiles.
	DestinationSlice uint // The index of the array slice in the destination texture to copy into.
	SourceLevel uint // The index of the mipmap level in the source texture.
	SourceRegion Region // The region in the source texture, in tiles.
	SourceSlice uint // The index of the array slice in the texture source of the copy operation.
}

// MTL4TimestampHeapEntry - Represents a timestamp data entry in a counter heap of type 
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4TimestampHeapEntry
type MTL4TimestampHeapEntry struct {
	Timestamp uint64
}

// MTL4UpdateSparseBufferMappingOperation - Groups together arguments for an operation to update a sparse buffer mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4UpdateSparseBufferMappingOperation
type MTL4UpdateSparseBufferMappingOperation struct {
	BufferRange foundation.Range // The range in the buffer, in tiles.
	HeapOffset uint // The starting offset in the heap, in tiles.
	Mode SparseTextureMappingMode // The mode of the mapping operation to perform.
}

// MTL4UpdateSparseTextureMappingOperation - Groups together arguments for an operation to update a sparse texture mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4UpdateSparseTextureMappingOperation
type MTL4UpdateSparseTextureMappingOperation struct {
	HeapOffset uint // The starting offset in the heap, in tiles.
	Mode SparseTextureMappingMode // The mode of the mapping operation to perform.
	TextureLevel uint // The index of the mipmap level in the texture to update.
	TextureRegion Region // The region in the texture to update, in tiles.
	TextureSlice uint // The index of the array slice in the texture to update.
}

// MTLAccelerationStructureInstanceDescriptor - A description of an instance in an instanced geometry acceleration structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureInstanceDescriptor
type MTLAccelerationStructureInstanceDescriptor struct {
	AccelerationStructureIndex uint32 // The index of the acceleration structure to use for the instance.
	IntersectionFunctionTableOffset uint32 // An offset for determining which function in the intersection function table Metal needs to call when testing a ray against the instance.
	Mask uint32 // A mask to use for the instance when testing a ray against the geometry.
	Options AccelerationStructureInstanceOptions // The options for the instance.
	TransformationMatrix PackedFloat4x3 // The transform for placing and orienting the instance in the scene.
}

// MTLAccelerationStructureMotionInstanceDescriptor - A description of an instance in an instanced geometry acceleration structure, with the instance including a user identifier and motion data for the instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionInstanceDescriptor
type MTLAccelerationStructureMotionInstanceDescriptor struct {
	AccelerationStructureIndex uint32 // The index of an acceleration structure which applies to the next   acceleration-structure motion instance you create with the descriptor.
	IntersectionFunctionTableOffset uint32 // An offset into the intersection-function table for ray tracing,   which applies to the next acceleration-structure motion instance you create   with the descriptor.
	Mask uint32 // A mask for testing ray-tracing rays with a scene’s geometry, which applies to   the next acceleration-structure motion instance you create with the descriptor.
	MotionEndBorderMode MotionBorderMode // A behavior that configures how a motion instance handles timestamps   after an ending time.
	MotionEndTime float32 // An ending time for the range of motion that the key-frame data represents.
	MotionStartBorderMode MotionBorderMode // A behavior that configures how a motion instance handles timestamps   before a starting time.
	MotionStartTime float32 // A starting time for the range of motion that the key-frame data represents.
	MotionTransformsCount uint32 // The number of motion data key-frames, which applies to   the next acceleration-structure motion instance you create with the descriptor.
	MotionTransformsStartIndex uint32 // The index of motion data that represents the first key-frame motion data, which applies to   the next acceleration-structure motion instance you create with the descriptor.
	Options AccelerationStructureInstanceOptions // An option set which applies to the next acceleration structure motion-instance   you create with the descriptor.
	UserID uint32 // An unique identifier, which applies to the next acceleration-structure motion instance   you create with the descriptor.
}

// MTLAccelerationStructureSizes - The expected sizes for a ray-tracing acceleration structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureSizes
type MTLAccelerationStructureSizes struct {
	AccelerationStructureSize uint // The size of the acceleration structure, in bytes.
	BuildScratchBufferSize uint // The amount of scratch memory, in bytes, the GPU devices needs to build the acceleration structure.
	RefitScratchBufferSize uint // The amount of scratch memory, in bytes, the GPU device needs to refit the acceleration structure.
}

// MTLAccelerationStructureUserIDInstanceDescriptor - A description of an instance in an instanced geometry acceleration structure, with the instance including a user identifier for the instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureUserIDInstanceDescriptor
type MTLAccelerationStructureUserIDInstanceDescriptor struct {
	AccelerationStructureIndex uint32 // The index of the acceleration structure to use for the instance.
	IntersectionFunctionTableOffset uint32 // An offset for determining which function in the intersection function table Metal calls when testing a ray against the instance.
	Mask uint32 // A mask to use for the instance when testing a ray against the geometry.
	Options AccelerationStructureInstanceOptions // The options for the instance.
	TransformationMatrix PackedFloat4x3 // The transform for placing and orienting the instance in the scene.
	UserID uint32 // The user identifier for the instance.
}

// MTLAxisAlignedBoundingBox - The bounds for an axis-aligned bounding box.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAxisAlignedBoundingBox-c.struct
type MTLAxisAlignedBoundingBox struct {
	Max PackedFloat3
	Min PackedFloat3
}

// MTLClearColor - An RGBA value used for a color pixel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLClearColor
type MTLClearColor struct {
	Alpha float64 // The alpha channel.
	Blue float64 // The blue color channel.
	Green float64 // The green color channel.
	Red float64 // The red color channel.
}

// MTLComponentTransform
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComponentTransform
type MTLComponentTransform struct {
	Pivot PackedFloat3
	Rotation PackedFloatQuaternion
	Scale PackedFloat3
	Shear PackedFloat3
	Translation PackedFloat3
}

// MTLCounterResultStageUtilization - The data structure for storing the data you resolve from a stage-utilization counter set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCounterResultStageUtilization
type MTLCounterResultStageUtilization struct {
	FragmentCycles uint64 // The number of cycles the GPU uses to run fragment shaders during a pass.
	PostTessellationVertexCycles uint64 // The number of cycles the GPU uses to run post-tessellation vertex shaders during a pass.
	RenderTargetCycles uint64 // The number of cycles the GPU uses to write data to render targets during a render pass.
	TessellationCycles uint64 // The number of cycles the GPU uses to run the tessellation stage during a pass.
	TotalCycles uint64 // The total number of cycles the GPU uses to run a pass.
	VertexCycles uint64 // The number of cycles the GPU uses to run vertex shaders during a pass.
}

// MTLCounterResultStatistic - The data structure for storing the data you resolve from a statistic counter set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCounterResultStatistic
type MTLCounterResultStatistic struct {
	ClipperInvocations uint64 // The number of primitives a render pass sends to the clip stage.
	ClipperPrimitivesOut uint64 // The number of primitives the clip stage produces during a render pass.
	ComputeKernelInvocations uint64 // The number of times a pass calls any compute kernel.
	FragmentInvocations uint64 // The number of times a render pass calls fragment shaders.
	FragmentsPassed uint64 // The number of fragments a render pass sends to the visibility and blend stages because they pass the scissor, depth, and stencil tests.
	PostTessellationVertexInvocations uint64 // The number of vertices a render pass sends to a post-tessellation vertex shader.
	TessellationInputPatches uint64 // The number of tessellation patches a render pass sends to the tessellation stage.
	VertexInvocations uint64 // The number of times a render pass calls any vertex shader.
}

// MTLCounterResultTimestamp - The data structure for storing the data you resolve from a timestamp counter set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCounterResultTimestamp
type MTLCounterResultTimestamp struct {
	Timestamp uint64 // A timestamp value from a GPU at a particular point in time during an operation, typically at the beginning or ending of a render stage.
}

// MTLDispatchThreadgroupsIndirectArguments - The data layout required for arguments needed to specify the size of threadgroups.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDispatchThreadgroupsIndirectArguments
type MTLDispatchThreadgroupsIndirectArguments struct {
	ThreadgroupsPerGrid uint32 // The number of threadgroups for the grid, in each dimension.
}

// MTLDispatchThreadsIndirectArguments
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDispatchThreadsIndirectArguments
type MTLDispatchThreadsIndirectArguments struct {
	ThreadsPerGrid uint32
	ThreadsPerThreadgroup uint32
}

// MTLDrawIndexedPrimitivesIndirectArguments - The data layout required for drawing indexed primitives via indirect buffer calls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDrawIndexedPrimitivesIndirectArguments
type MTLDrawIndexedPrimitivesIndirectArguments struct {
	BaseInstance uint32 // The first instance to draw.
	BaseVertex int32 // The first vertex to draw.
	IndexCount uint32 // For each instance, the number of indices to read from the index buffer.
	IndexStart uint32 // The first index to draw.
	InstanceCount uint32 // The number of instances to draw.
}

// MTLDrawPatchIndirectArguments - The data layout required for drawing patches via indirect buffer calls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDrawPatchIndirectArguments
type MTLDrawPatchIndirectArguments struct {
	BaseInstance uint32 // The first instance to draw.
	InstanceCount uint32 // The number of instances to draw.
	PatchCount uint32 // The number of patches in each instance.
	PatchStart uint32 // The patch start index.
}

// MTLDrawPrimitivesIndirectArguments - The data layout required for drawing primitives via indirect buffer calls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDrawPrimitivesIndirectArguments
type MTLDrawPrimitivesIndirectArguments struct {
	BaseInstance uint32 // The first instance to draw.
	InstanceCount uint32 // The number of instances to draw.
	VertexCount uint32 // The number of vertices to draw.
	VertexStart uint32 // The first vertex to draw.
}

// MTLIndirectAccelerationStructureInstanceDescriptor - A description of an instance in an instanced geometry acceleration structure that the GPU can populate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectAccelerationStructureInstanceDescriptor
type MTLIndirectAccelerationStructureInstanceDescriptor struct {
	AccelerationStructureID ResourceID
	IntersectionFunctionTableOffset uint32
	Mask uint32
	Options AccelerationStructureInstanceOptions
	TransformationMatrix PackedFloat4x3
	UserID uint32
}

// MTLIndirectAccelerationStructureMotionInstanceDescriptor - A description of an instance in an acceleration structure that the GPU can populate, with motion data for the instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectAccelerationStructureMotionInstanceDescriptor
type MTLIndirectAccelerationStructureMotionInstanceDescriptor struct {
	AccelerationStructureID ResourceID // The acceleration resource handle to use for this instance.
	IntersectionFunctionTableOffset uint32 // An offset for determining which function in the intersection function table Metal calls when testing a ray against the instance.
	Mask uint32 // An instance mask to ignore geometry during ray tracing.
	MotionEndBorderMode MotionBorderMode // The motion border mode describing what happens if Metal samples the acceleration structure after the motion end time.
	MotionEndTime float32 // The end time of the motion instance.
	MotionStartBorderMode MotionBorderMode // The motion border mode describing what happens if Metal samples the acceleration structure before the motion start time.
	MotionStartTime float32 // The start time of the motion instance.
	MotionTransformsCount uint32 // The number of motion transforms belonging to the motion instance.
	MotionTransformsStartIndex uint32 // The index of the first set of transforms describing one keyframe of the animation.
	Options AccelerationStructureInstanceOptions // The options for this instance.
	UserID uint32 // A user-assigned ID to help identify the instance.
}

// MTLIndirectCommandBufferExecutionRange - A range of commands in an indirect command buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferExecutionRange
type MTLIndirectCommandBufferExecutionRange struct {
	Length uint32 // The number of items in the command execution range.
	Location uint32 // The first index in the command execution range.
}

// MTLIntersectionFunctionBufferArguments
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionBufferArguments
type MTLIntersectionFunctionBufferArguments struct {
	IntersectionFunctionBuffer uint64
	IntersectionFunctionBufferSize uint64
	IntersectionFunctionStride uint64
}

// MTLMapIndirectArguments - The data layout for mapping sparse texture regions when using indirect commands.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMapIndirectArguments
type MTLMapIndirectArguments struct {
	MipMapLevel uint32 // The mipmap to change.
	RegionOriginX uint32 // The x coordinate of the region to change, measured in tile coordinates.
	RegionOriginY uint32 // The y coordinate of the region to change, measured in tile coordinates.
	RegionOriginZ uint32 // The z coordinate of the region to change, measured in tile coordinates.
	RegionSizeDepth uint32 // The depth of the region, measured in tile coordinates.
	RegionSizeHeight uint32 // The height of the region, measured in tile coordinates.
	RegionSizeWidth uint32 // The width of the region, measured in tile coordinates.
	SliceId uint32 // The texture slice to change.
}

// MTLOrigin - The coordinates for the front upper-left corner of a region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLOrigin
type MTLOrigin struct {
	X uint // The x coordinate of the origin.
	Y uint // The y coordinate of the origin.
	Z uint // The z coordinate of the origin.
}

// MTLPackedFloat3 - A structure that contains three 32-bit floating-point values with no additional padding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPackedFloat3-c.struct
type MTLPackedFloat3 struct {
	Elements float32
	X float32
	Y float32
	Z float32
}

// MTLPackedFloat4x3 - A structure that contains the top three rows of a 4x4 matrix of 32-bit floating-point values, in column-major order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPackedFloat4x3-c.struct
type MTLPackedFloat4x3 struct {
	Columns PackedFloat3
}

// MTLPackedFloatQuaternion
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPackedFloatQuaternion
type MTLPackedFloatQuaternion struct {
	W float32
	X float32
	Y float32
	Z float32
}

// MTLQuadTessellationFactorsHalf - The per-patch tessellation factors for a quad patch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLQuadTessellationFactorsHalf
type MTLQuadTessellationFactorsHalf struct {
	EdgeTessellationFactor uint16 // The edge tessellation factors, with each index value providing the tessellation factor for a particular edge.
	InsideTessellationFactor uint16 // The inside tessellation factors, with the value in index 0 providing the horizontal tessellation factor and the value in index 1 providing the vertical tessellation factor.
}

// MTLRegion - The bounds for a subset of an instance’s elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRegion
type MTLRegion struct {
	Origin Origin // The coordinates of the front upper-left corner of the region.
	Size Size // The dimensions of the region.
}

// MTLResourceID
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceID
type MTLResourceID struct {
}

// MTLSamplePosition - A subpixel sample position for use in multisample antialiasing (MSAA).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplePosition
type MTLSamplePosition struct {
	X float32 // The x position of the sample on the subpixel grid.
	Y float32 // The y position of the sample on the subpixel grid.
}

// MTLScissorRect - A rectangle for the scissor fragment test.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLScissorRect
type MTLScissorRect struct {
	Height uint // The height of the scissor rectangle, in pixels.
	Width uint // The width of the scissor rectangle, in pixels.
	X uint // The x window coordinate of the upper-left corner of the scissor rectangle.
	Y uint // The y window coordinate of the upper-left corner of the scissor rectangle.
}

// MTLSize - A type that represents one, two, or three dimensions of a type instance, such as an array or texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSize
type MTLSize struct {
	Depth uint // A value for the z-axis dimension.
	Height uint // A value for the y-axis dimension.
	Width uint // A value for the x-axis dimension.
}

// MTLSizeAndAlign - The size and alignment of a resource, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSizeAndAlign
type MTLSizeAndAlign struct {
	Align uint // The alignment of a resource, in bytes.
	Size uint // The size of a resource, in bytes.
}

// MTLStageInRegionIndirectArguments - The data layout required for the arguments needed to specify the stage-in region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStageInRegionIndirectArguments
type MTLStageInRegionIndirectArguments struct {
	StageInOrigin uint32 // The location of the upper-left corner of the block.
	StageInSize uint32 // The size of the block.
}

// MTLTextureSwizzleChannels - A pattern that modifies the data read or sampled from a texture by rearranging or duplicating the elements of a vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureSwizzleChannels
type MTLTextureSwizzleChannels struct {
	Alpha TextureSwizzle // The data copied to the fourth output channel.
	Blue TextureSwizzle // The data copied to the third output channel.
	Green TextureSwizzle // The data copied to the second output channel.
	Red TextureSwizzle // The data copied to the first output channel.
}

// MTLTriangleTessellationFactorsHalf - The per-patch tessellation factors for a triangle patch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTriangleTessellationFactorsHalf
type MTLTriangleTessellationFactorsHalf struct {
	EdgeTessellationFactor uint16 // The edge tessellation factors, with each index value providing the tessellation factor for a particular edge.
	InsideTessellationFactor uint16 // The inside tessellation factor.
}

// MTLVertexAmplificationViewMapping - An offset applied to a render target index and viewport index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexAmplificationViewMapping
type MTLVertexAmplificationViewMapping struct {
	RenderTargetArrayIndexOffset uint32 // An offset into the list of render targets.
	ViewportArrayIndexOffset uint32 // An offset into the list of viewports.
}

// MTLViewport - A 3D rectangular region for the viewport clipping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLViewport
type MTLViewport struct {
	Height float64 // The height of the viewport, in pixels.
	OriginX float64 // The x coordinate of the upper-left corner of the viewport.
	OriginY float64 // The y coordinate of the upper-left corner of the viewport.
	Width float64 // The width of the viewport, in pixels.
	Zfar float64 // The z coordinate of the far clipping plane of the viewport.
	Znear float64 // The z coordinate of the near clipping plane of the viewport.
}





