// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
)

// PMTL4ComputeCommandEncoder is the MTL4ComputeCommandEncoder protocol interface.
//
// Encodes a compute pass and other memory operations into a command buffer.
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//
// See: doc://com.apple.metal/documentation/Metal/MTL4ComputeCommandEncoder
type PMTL4ComputeCommandEncoder interface {
	// Required methods
	BuildAccelerationStructureDescriptorScratchBuffer(accelerationStructure unsafe.Pointer, descriptor IMTL4AccelerationStructureDescriptor, scratchBuffer MTL4BufferRange)
	CopyAccelerationStructureToAccelerationStructure(sourceAccelerationStructure unsafe.Pointer, destinationAccelerationStructure unsafe.Pointer)
	CopyFromBufferSourceOffsetToBufferDestinationOffsetSize(sourceBuffer unsafe.Pointer, sourceOffset uint, destinationBuffer unsafe.Pointer, destinationOffset uint, size uint)
	CopyFromTensorSourceOriginSourceDimensionsToTensorDestinationOriginDestinationDimensions(sourceTensor unsafe.Pointer, sourceOrigin IMTLTensorExtents, sourceDimensions IMTLTensorExtents, destinationTensor unsafe.Pointer, destinationOrigin IMTLTensorExtents, destinationDimensions IMTLTensorExtents)
	CopyFromTextureToTexture(sourceTexture unsafe.Pointer, destinationTexture unsafe.Pointer)
	CopyFromTextureSourceSliceSourceLevelToTextureDestinationSliceDestinationLevelSliceCountLevelCount(sourceTexture unsafe.Pointer, sourceSlice uint, sourceLevel uint, destinationTexture unsafe.Pointer, destinationSlice uint, destinationLevel uint, sliceCount uint, levelCount uint)
	CopyFromTextureSourceSliceSourceLevelSourceOriginSourceSizeToTextureDestinationSliceDestinationLevelDestinationOrigin(sourceTexture unsafe.Pointer, sourceSlice uint, sourceLevel uint, sourceOrigin Origin, sourceSize Size, destinationTexture unsafe.Pointer, destinationSlice uint, destinationLevel uint, destinationOrigin Origin)
	CopyAndCompactAccelerationStructureToAccelerationStructure(sourceAccelerationStructure unsafe.Pointer, destinationAccelerationStructure unsafe.Pointer)
	CopyFromBufferSourceOffsetSourceBytesPerRowSourceBytesPerImageSourceSizeToTextureDestinationSliceDestinationLevelDestinationOrigin(sourceBuffer unsafe.Pointer, sourceOffset uint, sourceBytesPerRow uint, sourceBytesPerImage uint, sourceSize Size, destinationTexture unsafe.Pointer, destinationSlice uint, destinationLevel uint, destinationOrigin Origin)
	CopyFromBufferSourceOffsetSourceBytesPerRowSourceBytesPerImageSourceSizeToTextureDestinationSliceDestinationLevelDestinationOriginOptions(sourceBuffer unsafe.Pointer, sourceOffset uint, sourceBytesPerRow uint, sourceBytesPerImage uint, sourceSize Size, destinationTexture unsafe.Pointer, destinationSlice uint, destinationLevel uint, destinationOrigin Origin, options BlitOption)
	CopyFromTextureSourceSliceSourceLevelSourceOriginSourceSizeToBufferDestinationOffsetDestinationBytesPerRowDestinationBytesPerImage(sourceTexture unsafe.Pointer, sourceSlice uint, sourceLevel uint, sourceOrigin Origin, sourceSize Size, destinationBuffer unsafe.Pointer, destinationOffset uint, destinationBytesPerRow uint, destinationBytesPerImage uint)
	CopyFromTextureSourceSliceSourceLevelSourceOriginSourceSizeToBufferDestinationOffsetDestinationBytesPerRowDestinationBytesPerImageOptions(sourceTexture unsafe.Pointer, sourceSlice uint, sourceLevel uint, sourceOrigin Origin, sourceSize Size, destinationBuffer unsafe.Pointer, destinationOffset uint, destinationBytesPerRow uint, destinationBytesPerImage uint, options BlitOption)
	CopyIndirectCommandBufferSourceRangeDestinationDestinationIndex(source unsafe.Pointer, sourceRange foundation.Range, destination unsafe.Pointer, destinationIndex uint)
	DispatchThreadgroupsWithIndirectBufferThreadsPerThreadgroup(indirectBuffer GPUAddress, threadsPerThreadgroup Size)
	DispatchThreadgroupsThreadsPerThreadgroup(threadgroupsPerGrid Size, threadsPerThreadgroup Size)
	DispatchThreadsWithIndirectBuffer(indirectBuffer GPUAddress)
	DispatchThreadsThreadsPerThreadgroup(threadsPerGrid Size, threadsPerThreadgroup Size)
	ExecuteCommandsInBufferIndirectBuffer(indirectCommandbuffer unsafe.Pointer, indirectRangeBuffer GPUAddress)
	ExecuteCommandsInBufferWithRange(indirectCommandBuffer unsafe.Pointer, executionRange foundation.Range)
	FillBufferRangeValue(buffer unsafe.Pointer, range_ foundation.Range, value uint8 /* not a class type */)
	GenerateMipmapsForTexture(texture unsafe.Pointer)
	OptimizeContentsForCPUAccess(texture unsafe.Pointer)
	OptimizeContentsForCPUAccessSliceLevel(texture unsafe.Pointer, slice uint, level uint)
	OptimizeContentsForGPUAccess(texture unsafe.Pointer)
	OptimizeContentsForGPUAccessSliceLevel(texture unsafe.Pointer, slice uint, level uint)
	OptimizeIndirectCommandBufferWithRange(indirectCommandBuffer unsafe.Pointer, range_ foundation.Range)
	RefitAccelerationStructureDescriptorDestinationScratchBufferOptions(sourceAccelerationStructure unsafe.Pointer, descriptor IMTL4AccelerationStructureDescriptor, destinationAccelerationStructure unsafe.Pointer, scratchBuffer MTL4BufferRange, options AccelerationStructureRefitOptions)
	RefitAccelerationStructureDescriptorDestinationScratchBuffer(sourceAccelerationStructure unsafe.Pointer, descriptor IMTL4AccelerationStructureDescriptor, destinationAccelerationStructure unsafe.Pointer, scratchBuffer MTL4BufferRange)
	ResetCommandsInBufferWithRange(buffer unsafe.Pointer, range_ foundation.Range)
	SetArgumentTable(argumentTable unsafe.Pointer)
	SetComputePipelineState(state unsafe.Pointer)
	SetImageblockWidthHeight(width uint, height uint)
	SetThreadgroupMemoryLengthAtIndex(length uint, index uint)
	Stages() Stages
	WriteCompactedAccelerationStructureSizeToBuffer(accelerationStructure unsafe.Pointer, buffer MTL4BufferRange)
	WriteTimestampWithGranularityIntoHeapAtIndex(granularity MTL4TimestampGranularity, counterHeap unsafe.Pointer, index uint)
}
