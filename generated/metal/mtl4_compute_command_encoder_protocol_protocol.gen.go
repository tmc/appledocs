// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/corefoundation"
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
	BuildAccelerationStructureDescriptorScratchBuffer(accelerationStructure unsafe.Pointer, descriptor IMTL4AccelerationStructureDescriptor, scratchBuffer objc.IObject /* cross-framework: MTL4BufferRange */)/* debug [protocol_interface/required_method]: BuildAccelerationStructureDescriptorScratchBuffer */
	CopyAccelerationStructureToAccelerationStructure(sourceAccelerationStructure unsafe.Pointer, destinationAccelerationStructure unsafe.Pointer)/* debug [protocol_interface/required_method]: CopyAccelerationStructureToAccelerationStructure */
	CopyFromBufferSourceOffsetToBufferDestinationOffsetSize(sourceBuffer unsafe.Pointer, sourceOffset uint, destinationBuffer unsafe.Pointer, destinationOffset uint, size uint)/* debug [protocol_interface/required_method]: CopyFromBufferSourceOffsetToBufferDestinationOffsetSize */
	CopyFromTensorSourceOriginSourceDimensionsToTensorDestinationOriginDestinationDimensions(sourceTensor unsafe.Pointer, sourceOrigin IMTLTensorExtents, sourceDimensions IMTLTensorExtents, destinationTensor unsafe.Pointer, destinationOrigin IMTLTensorExtents, destinationDimensions IMTLTensorExtents)/* debug [protocol_interface/required_method]: CopyFromTensorSourceOriginSourceDimensionsToTensorDestinationOriginDestinationDimensions */
	CopyFromTextureToTexture(sourceTexture unsafe.Pointer, destinationTexture unsafe.Pointer)/* debug [protocol_interface/required_method]: CopyFromTextureToTexture */
	CopyFromTextureSourceSliceSourceLevelToTextureDestinationSliceDestinationLevelSliceCountLevelCount(sourceTexture unsafe.Pointer, sourceSlice uint, sourceLevel uint, destinationTexture unsafe.Pointer, destinationSlice uint, destinationLevel uint, sliceCount uint, levelCount uint)/* debug [protocol_interface/required_method]: CopyFromTextureSourceSliceSourceLevelToTextureDestinationSliceDestinationLevelSliceCountLevelCount */
	CopyFromTextureSourceSliceSourceLevelSourceOriginSourceSizeToTextureDestinationSliceDestinationLevelDestinationOrigin(sourceTexture unsafe.Pointer, sourceSlice uint, sourceLevel uint, sourceOrigin objc.IObject /* cross-framework: MTLOrigin */, sourceSize objc.IObject /* cross-framework: MTLSize */, destinationTexture unsafe.Pointer, destinationSlice uint, destinationLevel uint, destinationOrigin objc.IObject /* cross-framework: MTLOrigin */)/* debug [protocol_interface/required_method]: CopyFromTextureSourceSliceSourceLevelSourceOriginSourceSizeToTextureDestinationSliceDestinationLevelDestinationOrigin */
	CopyAndCompactAccelerationStructureToAccelerationStructure(sourceAccelerationStructure unsafe.Pointer, destinationAccelerationStructure unsafe.Pointer)/* debug [protocol_interface/required_method]: CopyAndCompactAccelerationStructureToAccelerationStructure */
	CopyFromBufferSourceOffsetSourceBytesPerRowSourceBytesPerImageSourceSizeToTextureDestinationSliceDestinationLevelDestinationOrigin(sourceBuffer unsafe.Pointer, sourceOffset uint, sourceBytesPerRow uint, sourceBytesPerImage uint, sourceSize objc.IObject /* cross-framework: MTLSize */, destinationTexture unsafe.Pointer, destinationSlice uint, destinationLevel uint, destinationOrigin objc.IObject /* cross-framework: MTLOrigin */)/* debug [protocol_interface/required_method]: CopyFromBufferSourceOffsetSourceBytesPerRowSourceBytesPerImageSourceSizeToTextureDestinationSliceDestinationLevelDestinationOrigin */
	CopyFromBufferSourceOffsetSourceBytesPerRowSourceBytesPerImageSourceSizeToTextureDestinationSliceDestinationLevelDestinationOriginOptions(sourceBuffer unsafe.Pointer, sourceOffset uint, sourceBytesPerRow uint, sourceBytesPerImage uint, sourceSize objc.IObject /* cross-framework: MTLSize */, destinationTexture unsafe.Pointer, destinationSlice uint, destinationLevel uint, destinationOrigin objc.IObject /* cross-framework: MTLOrigin */, options BlitOption)/* debug [protocol_interface/required_method]: CopyFromBufferSourceOffsetSourceBytesPerRowSourceBytesPerImageSourceSizeToTextureDestinationSliceDestinationLevelDestinationOriginOptions */
	CopyFromTextureSourceSliceSourceLevelSourceOriginSourceSizeToBufferDestinationOffsetDestinationBytesPerRowDestinationBytesPerImage(sourceTexture unsafe.Pointer, sourceSlice uint, sourceLevel uint, sourceOrigin objc.IObject /* cross-framework: MTLOrigin */, sourceSize objc.IObject /* cross-framework: MTLSize */, destinationBuffer unsafe.Pointer, destinationOffset uint, destinationBytesPerRow uint, destinationBytesPerImage uint)/* debug [protocol_interface/required_method]: CopyFromTextureSourceSliceSourceLevelSourceOriginSourceSizeToBufferDestinationOffsetDestinationBytesPerRowDestinationBytesPerImage */
	CopyFromTextureSourceSliceSourceLevelSourceOriginSourceSizeToBufferDestinationOffsetDestinationBytesPerRowDestinationBytesPerImageOptions(sourceTexture unsafe.Pointer, sourceSlice uint, sourceLevel uint, sourceOrigin objc.IObject /* cross-framework: MTLOrigin */, sourceSize objc.IObject /* cross-framework: MTLSize */, destinationBuffer unsafe.Pointer, destinationOffset uint, destinationBytesPerRow uint, destinationBytesPerImage uint, options BlitOption)/* debug [protocol_interface/required_method]: CopyFromTextureSourceSliceSourceLevelSourceOriginSourceSizeToBufferDestinationOffsetDestinationBytesPerRowDestinationBytesPerImageOptions */
	CopyIndirectCommandBufferSourceRangeDestinationDestinationIndex(source unsafe.Pointer, sourceRange corefoundation.Range, destination unsafe.Pointer, destinationIndex uint)/* debug [protocol_interface/required_method]: CopyIndirectCommandBufferSourceRangeDestinationDestinationIndex */
	DispatchThreadgroupsWithIndirectBufferThreadsPerThreadgroup(indirectBuffer GPUAddress /* typedef */, threadsPerThreadgroup objc.IObject /* cross-framework: MTLSize */)/* debug [protocol_interface/required_method]: DispatchThreadgroupsWithIndirectBufferThreadsPerThreadgroup */
	DispatchThreadgroupsThreadsPerThreadgroup(threadgroupsPerGrid objc.IObject /* cross-framework: MTLSize */, threadsPerThreadgroup objc.IObject /* cross-framework: MTLSize */)/* debug [protocol_interface/required_method]: DispatchThreadgroupsThreadsPerThreadgroup */
	DispatchThreadsWithIndirectBuffer(indirectBuffer GPUAddress /* typedef */)/* debug [protocol_interface/required_method]: DispatchThreadsWithIndirectBuffer */
	DispatchThreadsThreadsPerThreadgroup(threadsPerGrid objc.IObject /* cross-framework: MTLSize */, threadsPerThreadgroup objc.IObject /* cross-framework: MTLSize */)/* debug [protocol_interface/required_method]: DispatchThreadsThreadsPerThreadgroup */
	ExecuteCommandsInBufferIndirectBuffer(indirectCommandbuffer unsafe.Pointer, indirectRangeBuffer GPUAddress /* typedef */)/* debug [protocol_interface/required_method]: ExecuteCommandsInBufferIndirectBuffer */
	ExecuteCommandsInBufferWithRange(indirectCommandBuffer unsafe.Pointer, executionRange corefoundation.Range)/* debug [protocol_interface/required_method]: ExecuteCommandsInBufferWithRange */
	FillBufferRangeValue(buffer unsafe.Pointer, range_ corefoundation.Range, value uint8 /* not a class type */)/* debug [protocol_interface/required_method]: FillBufferRangeValue */
	GenerateMipmapsForTexture(texture unsafe.Pointer)/* debug [protocol_interface/required_method]: GenerateMipmapsForTexture */
	OptimizeContentsForCPUAccess(texture unsafe.Pointer)/* debug [protocol_interface/required_method]: OptimizeContentsForCPUAccess */
	OptimizeContentsForCPUAccessSliceLevel(texture unsafe.Pointer, slice uint, level uint)/* debug [protocol_interface/required_method]: OptimizeContentsForCPUAccessSliceLevel */
	OptimizeContentsForGPUAccess(texture unsafe.Pointer)/* debug [protocol_interface/required_method]: OptimizeContentsForGPUAccess */
	OptimizeContentsForGPUAccessSliceLevel(texture unsafe.Pointer, slice uint, level uint)/* debug [protocol_interface/required_method]: OptimizeContentsForGPUAccessSliceLevel */
	OptimizeIndirectCommandBufferWithRange(indirectCommandBuffer unsafe.Pointer, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: OptimizeIndirectCommandBufferWithRange */
	RefitAccelerationStructureDescriptorDestinationScratchBufferOptions(sourceAccelerationStructure unsafe.Pointer, descriptor IMTL4AccelerationStructureDescriptor, destinationAccelerationStructure unsafe.Pointer, scratchBuffer objc.IObject /* cross-framework: MTL4BufferRange */, options AccelerationStructureRefitOptions)/* debug [protocol_interface/required_method]: RefitAccelerationStructureDescriptorDestinationScratchBufferOptions */
	RefitAccelerationStructureDescriptorDestinationScratchBuffer(sourceAccelerationStructure unsafe.Pointer, descriptor IMTL4AccelerationStructureDescriptor, destinationAccelerationStructure unsafe.Pointer, scratchBuffer objc.IObject /* cross-framework: MTL4BufferRange */)/* debug [protocol_interface/required_method]: RefitAccelerationStructureDescriptorDestinationScratchBuffer */
	ResetCommandsInBufferWithRange(buffer unsafe.Pointer, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: ResetCommandsInBufferWithRange */
	SetArgumentTable(argumentTable unsafe.Pointer)/* debug [protocol_interface/required_method]: SetArgumentTable */
	SetComputePipelineState(state unsafe.Pointer)/* debug [protocol_interface/required_method]: SetComputePipelineState */
	SetImageblockWidthHeight(width uint, height uint)/* debug [protocol_interface/required_method]: SetImageblockWidthHeight */
	SetThreadgroupMemoryLengthAtIndex(length uint, index uint)/* debug [protocol_interface/required_method]: SetThreadgroupMemoryLengthAtIndex */
	Stages() Stages/* debug [protocol_interface/required_method]: Stages */
	WriteCompactedAccelerationStructureSizeToBuffer(accelerationStructure unsafe.Pointer, buffer objc.IObject /* cross-framework: MTL4BufferRange */)/* debug [protocol_interface/required_method]: WriteCompactedAccelerationStructureSizeToBuffer */
	WriteTimestampWithGranularityIntoHeapAtIndex(granularity MTL4TimestampGranularity, counterHeap unsafe.Pointer, index uint)/* debug [protocol_interface/required_method]: WriteTimestampWithGranularityIntoHeapAtIndex */
}
