// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PDevice is the MTLDevice protocol interface.
//
// The main Metal interface to a GPU that apps use to draw graphics and run computations in parallel.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.11+
//   - tvOS +
//   - visionOS 1.0+
//
// See: doc://com.apple.metal/documentation/Metal/MTLDevice
type PDevice interface {
	// Required methods
	AccelerationStructureSizesWithDescriptor(descriptor IMTLAccelerationStructureDescriptor) MTLAccelerationStructureSizes
	FunctionHandleWithFunction(function unsafe.Pointer) unsafe.Pointer
	FunctionHandleWithBinaryFunction(function unsafe.Pointer) unsafe.Pointer
	GetDefaultSamplePositionsCount(positions SamplePosition, count uint)
	HeapAccelerationStructureSizeAndAlignWithDescriptor(descriptor IMTLAccelerationStructureDescriptor) MTLSizeAndAlign
	HeapAccelerationStructureSizeAndAlignWithSize(size uint) MTLSizeAndAlign
	HeapBufferSizeAndAlignWithLengthOptions(length uint, options ResourceOptions) MTLSizeAndAlign
	HeapTextureSizeAndAlignWithDescriptor(desc IMTLTextureDescriptor) MTLSizeAndAlign
	NewAccelerationStructureWithDescriptor(descriptor IMTLAccelerationStructureDescriptor) unsafe.Pointer
	NewAccelerationStructureWithSize(size uint) unsafe.Pointer
	NewArchiveWithURLError(url foundation.foundation.INSURL, error_ foundation.foundation.INSError) unsafe.Pointer
	NewArgumentEncoderWithArguments(arguments []ArgumentDescriptor) unsafe.Pointer
	NewArgumentEncoderWithBufferBinding(bufferBinding unsafe.Pointer) unsafe.Pointer
	NewArgumentTableWithDescriptorError(descriptor IMTL4ArgumentTableDescriptor, error_ foundation.foundation.INSError) unsafe.Pointer
	NewBinaryArchiveWithDescriptorError(descriptor IMTLBinaryArchiveDescriptor, error_ foundation.foundation.INSError) unsafe.Pointer
	NewBufferWithBytesLengthOptions(pointer objectivec.IObject, length uint, options ResourceOptions) unsafe.Pointer
	NewBufferWithBytesNoCopyLengthOptionsDeallocator(pointer objectivec.IObject, length uint, options ResourceOptions, deallocator unsafe.Pointer) unsafe.Pointer
	NewBufferWithLengthOptions(length uint, options ResourceOptions) unsafe.Pointer
	NewBufferWithLengthOptionsPlacementSparsePageSize(length uint, options ResourceOptions, placementSparsePageSize SparsePageSize) unsafe.Pointer
	NewCommandAllocator() unsafe.Pointer
	NewCommandAllocatorWithDescriptorError(descriptor IMTL4CommandAllocatorDescriptor, error_ foundation.foundation.INSError) unsafe.Pointer
	NewCommandBuffer() unsafe.Pointer
	NewCommandQueue() unsafe.Pointer
	NewCommandQueueWithDescriptor(descriptor IMTLCommandQueueDescriptor) unsafe.Pointer
	NewCommandQueueWithMaxCommandBufferCount(maxCommandBufferCount uint) unsafe.Pointer
	NewCompilerWithDescriptorError(descriptor IMTL4CompilerDescriptor, error_ foundation.foundation.INSError) unsafe.Pointer
	NewComputePipelineStateWithDescriptorOptionsCompletionHandler(descriptor IMTLComputePipelineDescriptor, options PipelineOption, completionHandler NewComputePipelineStateWithReflectionCompletionHandler /* not a class type */)
	NewComputePipelineStateWithDescriptorOptionsReflectionError(descriptor IMTLComputePipelineDescriptor, options PipelineOption, reflection AutoreleasedComputePipelineReflection, error_ foundation.foundation.INSError) unsafe.Pointer
	NewComputePipelineStateWithFunctionError(computeFunction unsafe.Pointer, error_ foundation.foundation.INSError) unsafe.Pointer
	NewComputePipelineStateWithFunctionCompletionHandler(computeFunction unsafe.Pointer, completionHandler NewComputePipelineStateCompletionHandler /* not a class type */)
	NewComputePipelineStateWithFunctionOptionsCompletionHandler(computeFunction unsafe.Pointer, options PipelineOption, completionHandler NewComputePipelineStateWithReflectionCompletionHandler /* not a class type */)
	NewComputePipelineStateWithFunctionOptionsReflectionError(computeFunction unsafe.Pointer, options PipelineOption, reflection AutoreleasedComputePipelineReflection, error_ foundation.foundation.INSError) unsafe.Pointer
	NewCounterHeapWithDescriptorError(descriptor IMTL4CounterHeapDescriptor, error_ foundation.foundation.INSError) unsafe.Pointer
	NewCounterSampleBufferWithDescriptorError(descriptor IMTLCounterSampleBufferDescriptor, error_ foundation.foundation.INSError) unsafe.Pointer
	NewDefaultLibrary() unsafe.Pointer
	NewDefaultLibraryWithBundleError(bundle foundation.Bundle, error_ foundation.foundation.INSError) unsafe.Pointer
	NewDepthStencilStateWithDescriptor(descriptor IMTLDepthStencilDescriptor) unsafe.Pointer
	NewDynamicLibraryError(library unsafe.Pointer, error_ foundation.foundation.INSError) unsafe.Pointer
	NewDynamicLibraryWithURLError(url foundation.foundation.INSURL, error_ foundation.foundation.INSError) unsafe.Pointer
	NewEvent() unsafe.Pointer
	NewFence() unsafe.Pointer
	NewHeapWithDescriptor(descriptor IMTLHeapDescriptor) unsafe.Pointer
	NewIndirectCommandBufferWithDescriptorMaxCommandCountOptions(descriptor IMTLIndirectCommandBufferDescriptor, maxCount uint, options ResourceOptions) unsafe.Pointer
	NewIOCommandQueueWithDescriptorError(descriptor IMTLIOCommandQueueDescriptor, error_ foundation.foundation.INSError) unsafe.Pointer
	NewIOFileHandleWithURLError(url foundation.foundation.INSURL, error_ foundation.foundation.INSError) unsafe.Pointer
	NewIOFileHandleWithURLCompressionMethodError(url foundation.foundation.INSURL, compressionMethod CompressionMethod, error_ foundation.foundation.INSError) unsafe.Pointer
	NewIOHandleWithURLError(url foundation.foundation.INSURL, error_ foundation.foundation.INSError) unsafe.Pointer
	NewIOHandleWithURLCompressionMethodError(url foundation.foundation.INSURL, compressionMethod CompressionMethod, error_ foundation.foundation.INSError) unsafe.Pointer
	NewLibraryWithDataError(data objectivec.IObject, error_ foundation.foundation.INSError) unsafe.Pointer
	NewLibraryWithFileError(filepath foundation.foundation.INSString, error_ foundation.foundation.INSError) unsafe.Pointer
	NewLibraryWithSourceOptionsError(source foundation.foundation.INSString, options IMTLCompileOptions, error_ foundation.foundation.INSError) unsafe.Pointer
	NewLibraryWithSourceOptionsCompletionHandler(source foundation.foundation.INSString, options IMTLCompileOptions, completionHandler NewLibraryCompletionHandler /* not a class type */)
	NewLibraryWithStitchedDescriptorError(descriptor IMTLStitchedLibraryDescriptor, error_ foundation.foundation.INSError) unsafe.Pointer
	NewLibraryWithStitchedDescriptorCompletionHandler(descriptor IMTLStitchedLibraryDescriptor, completionHandler NewLibraryCompletionHandler /* not a class type */)
	NewLibraryWithURLError(url foundation.foundation.INSURL, error_ foundation.foundation.INSError) unsafe.Pointer
	NewLogStateWithDescriptorError(descriptor IMTLLogStateDescriptor, error_ foundation.foundation.INSError) unsafe.Pointer
	NewMTL4CommandQueue() unsafe.Pointer
	NewMTL4CommandQueueWithDescriptorError(descriptor IMTL4CommandQueueDescriptor, error_ foundation.foundation.INSError) unsafe.Pointer
	NewPipelineDataSetSerializerWithDescriptor(descriptor IMTL4PipelineDataSetSerializerDescriptor) unsafe.Pointer
	NewRasterizationRateMapWithDescriptor(descriptor IMTLRasterizationRateMapDescriptor) unsafe.Pointer
	NewRenderPipelineStateWithDescriptorError(descriptor IMTLRenderPipelineDescriptor, error_ foundation.foundation.INSError) unsafe.Pointer
	NewRenderPipelineStateWithDescriptorCompletionHandler(descriptor IMTLRenderPipelineDescriptor, completionHandler NewRenderPipelineStateCompletionHandler /* not a class type */)
	NewRenderPipelineStateWithMeshDescriptorOptionsCompletionHandler(descriptor IMTLMeshRenderPipelineDescriptor, options PipelineOption, completionHandler NewRenderPipelineStateWithReflectionCompletionHandler /* not a class type */)
	NewRenderPipelineStateWithDescriptorOptionsCompletionHandler(descriptor IMTLRenderPipelineDescriptor, options PipelineOption, completionHandler NewRenderPipelineStateWithReflectionCompletionHandler /* not a class type */)
	NewRenderPipelineStateWithDescriptorOptionsReflectionError(descriptor IMTLRenderPipelineDescriptor, options PipelineOption, reflection AutoreleasedRenderPipelineReflection, error_ foundation.foundation.INSError) unsafe.Pointer
	NewRenderPipelineStateWithTileDescriptorOptionsCompletionHandler(descriptor IMTLTileRenderPipelineDescriptor, options PipelineOption, completionHandler NewRenderPipelineStateWithReflectionCompletionHandler /* not a class type */)
	NewRenderPipelineStateWithTileDescriptorOptionsReflectionError(descriptor IMTLTileRenderPipelineDescriptor, options PipelineOption, reflection AutoreleasedRenderPipelineReflection, error_ foundation.foundation.INSError) unsafe.Pointer
	NewResidencySetWithDescriptorError(desc IMTLResidencySetDescriptor, error_ foundation.foundation.INSError) unsafe.Pointer
	NewSamplerStateWithDescriptor(descriptor IMTLSamplerDescriptor) unsafe.Pointer
	NewSharedEvent() unsafe.Pointer
	NewSharedEventWithHandle(sharedEventHandle IMTLSharedEventHandle) unsafe.Pointer
	NewSharedTextureWithDescriptor(descriptor IMTLTextureDescriptor) unsafe.Pointer
	NewSharedTextureWithHandle(sharedHandle IMTLSharedTextureHandle) unsafe.Pointer
	NewTensorWithDescriptorError(descriptor IMTLTensorDescriptor, error_ foundation.foundation.INSError) unsafe.Pointer
	NewTextureWithDescriptor(descriptor IMTLTextureDescriptor) unsafe.Pointer
	NewTextureWithDescriptorIosurfacePlane(descriptor IMTLTextureDescriptor, iosurface SurfaceRef /* not a class type */, plane uint) unsafe.Pointer
	NewTextureViewPoolWithDescriptorError(descriptor IMTLResourceViewPoolDescriptor, error_ foundation.foundation.INSError) unsafe.Pointer
	MinimumLinearTextureAlignmentForPixelFormat(format PixelFormat) uint
	MinimumTextureBufferAlignmentForPixelFormat(format PixelFormat) uint
	NewRenderPipelineStateWithMeshDescriptorOptionsReflectionError(descriptor IMTLMeshRenderPipelineDescriptor, options PipelineOption, reflection AutoreleasedRenderPipelineReflection, error_ foundation.foundation.INSError) unsafe.Pointer
	QueryTimestampFrequency() uint64
	SampleTimestampsGpuTimestamp(cpuTimestamp Timestamp, gpuTimestamp Timestamp)
	SizeOfCounterHeapEntry(type_ MTL4CounterHeapType) uint
	SparseTileSizeWithTextureTypePixelFormatSampleCountSparsePageSize(textureType TextureType, pixelFormat PixelFormat, sampleCount uint, sparsePageSize SparsePageSize) MTLSize
	SparseTileSizeWithTextureTypePixelFormatSampleCount(textureType TextureType, pixelFormat PixelFormat, sampleCount uint) MTLSize
	SparseTileSizeInBytesForSparsePageSize(sparsePageSize SparsePageSize) uint
	SupportsCounterSampling(samplingPoint CounterSamplingPoint) bool
	SupportsFamily(gpuFamily GPUFamily) bool
	SupportsFeatureSet(featureSet FeatureSet) bool
	SupportsRasterizationRateMapWithLayerCount(layerCount uint) bool
	SupportsTextureSampleCount(sampleCount uint) bool
	SupportsVertexAmplificationCount(count uint) bool
	TensorSizeAndAlignWithDescriptor(descriptor IMTLTensorDescriptor) MTLSizeAndAlign
	// Optional methods
	ConvertSparsePixelRegionsToTileRegionsWithTileSizeAlignmentModeNumRegions(pixelRegions []MTLRegion, tileRegions []MTLRegion, tileSize Size, mode SparseTextureRegionAlignmentMode, numRegions uint)
	HasConvertSparsePixelRegionsToTileRegionsWithTileSizeAlignmentModeNumRegions() bool
	ConvertSparseTileRegionsToPixelRegionsWithTileSizeNumRegions(tileRegions []MTLRegion, pixelRegions []MTLRegion, tileSize Size, numRegions uint)
	HasConvertSparseTileRegionsToPixelRegionsWithTileSizeNumRegions() bool
}
