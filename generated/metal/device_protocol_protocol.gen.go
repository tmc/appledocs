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
	AccelerationStructureSizesWithDescriptor(descriptor IMTLAccelerationStructureDescriptor) MTLAccelerationStructureSizes/* debug [protocol_interface/required_method]: AccelerationStructureSizesWithDescriptor */
	FunctionHandleWithFunction(function unsafe.Pointer) unsafe.Pointer/* debug [protocol_interface/required_method]: FunctionHandleWithFunction */
	FunctionHandleWithBinaryFunction(function unsafe.Pointer) unsafe.Pointer/* debug [protocol_interface/required_method]: FunctionHandleWithBinaryFunction */
	GetDefaultSamplePositionsCount(positions objc.IObject /* cross-framework: MTLSamplePosition */, count uint)/* debug [protocol_interface/required_method]: GetDefaultSamplePositionsCount */
	HeapAccelerationStructureSizeAndAlignWithDescriptor(descriptor IMTLAccelerationStructureDescriptor) MTLSizeAndAlign/* debug [protocol_interface/required_method]: HeapAccelerationStructureSizeAndAlignWithDescriptor */
	HeapAccelerationStructureSizeAndAlignWithSize(size uint) MTLSizeAndAlign/* debug [protocol_interface/required_method]: HeapAccelerationStructureSizeAndAlignWithSize */
	HeapBufferSizeAndAlignWithLengthOptions(length uint, options ResourceOptions) MTLSizeAndAlign/* debug [protocol_interface/required_method]: HeapBufferSizeAndAlignWithLengthOptions */
	HeapTextureSizeAndAlignWithDescriptor(desc IMTLTextureDescriptor) MTLSizeAndAlign/* debug [protocol_interface/required_method]: HeapTextureSizeAndAlignWithDescriptor */
	NewAccelerationStructureWithDescriptor(descriptor IMTLAccelerationStructureDescriptor) unsafe.Pointer/* debug [protocol_interface/required_method]: NewAccelerationStructureWithDescriptor */
	NewAccelerationStructureWithSize(size uint) unsafe.Pointer/* debug [protocol_interface/required_method]: NewAccelerationStructureWithSize */
	NewArchiveWithURLError(url objc.IObject /* cross-framework: NSURL */, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewArchiveWithURLError */
	NewArgumentEncoderWithArguments(arguments []ArgumentDescriptor) unsafe.Pointer/* debug [protocol_interface/required_method]: NewArgumentEncoderWithArguments */
	NewArgumentEncoderWithBufferBinding(bufferBinding unsafe.Pointer) unsafe.Pointer/* debug [protocol_interface/required_method]: NewArgumentEncoderWithBufferBinding */
	NewArgumentTableWithDescriptorError(descriptor IMTL4ArgumentTableDescriptor, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewArgumentTableWithDescriptorError */
	NewBinaryArchiveWithDescriptorError(descriptor IMTLBinaryArchiveDescriptor, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewBinaryArchiveWithDescriptorError */
	NewBufferWithBytesLengthOptions(pointer objectivec.IObject, length uint, options ResourceOptions) unsafe.Pointer/* debug [protocol_interface/required_method]: NewBufferWithBytesLengthOptions */
	NewBufferWithBytesNoCopyLengthOptionsDeallocator(pointer objectivec.IObject, length uint, options ResourceOptions, deallocator unsafe.Pointer) unsafe.Pointer/* debug [protocol_interface/required_method]: NewBufferWithBytesNoCopyLengthOptionsDeallocator */
	NewBufferWithLengthOptions(length uint, options ResourceOptions) unsafe.Pointer/* debug [protocol_interface/required_method]: NewBufferWithLengthOptions */
	NewBufferWithLengthOptionsPlacementSparsePageSize(length uint, options ResourceOptions, placementSparsePageSize SparsePageSize) unsafe.Pointer/* debug [protocol_interface/required_method]: NewBufferWithLengthOptionsPlacementSparsePageSize */
	NewCommandAllocator() unsafe.Pointer/* debug [protocol_interface/required_method]: NewCommandAllocator */
	NewCommandAllocatorWithDescriptorError(descriptor IMTL4CommandAllocatorDescriptor, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewCommandAllocatorWithDescriptorError */
	NewCommandBuffer() unsafe.Pointer/* debug [protocol_interface/required_method]: NewCommandBuffer */
	NewCommandQueue() unsafe.Pointer/* debug [protocol_interface/required_method]: NewCommandQueue */
	NewCommandQueueWithDescriptor(descriptor IMTLCommandQueueDescriptor) unsafe.Pointer/* debug [protocol_interface/required_method]: NewCommandQueueWithDescriptor */
	NewCommandQueueWithMaxCommandBufferCount(maxCommandBufferCount uint) unsafe.Pointer/* debug [protocol_interface/required_method]: NewCommandQueueWithMaxCommandBufferCount */
	NewCompilerWithDescriptorError(descriptor IMTL4CompilerDescriptor, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewCompilerWithDescriptorError */
	NewComputePipelineStateWithDescriptorOptionsCompletionHandler(descriptor IMTLComputePipelineDescriptor, options PipelineOption, completionHandler NewComputePipelineStateWithReflectionCompletionHandler /* not a class type */)/* debug [protocol_interface/required_method]: NewComputePipelineStateWithDescriptorOptionsCompletionHandler */
	NewComputePipelineStateWithDescriptorOptionsReflectionError(descriptor IMTLComputePipelineDescriptor, options PipelineOption, reflection AutoreleasedComputePipelineReflection /* typedef */, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewComputePipelineStateWithDescriptorOptionsReflectionError */
	NewComputePipelineStateWithFunctionError(computeFunction unsafe.Pointer, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewComputePipelineStateWithFunctionError */
	NewComputePipelineStateWithFunctionCompletionHandler(computeFunction unsafe.Pointer, completionHandler NewComputePipelineStateCompletionHandler /* not a class type */)/* debug [protocol_interface/required_method]: NewComputePipelineStateWithFunctionCompletionHandler */
	NewComputePipelineStateWithFunctionOptionsCompletionHandler(computeFunction unsafe.Pointer, options PipelineOption, completionHandler NewComputePipelineStateWithReflectionCompletionHandler /* not a class type */)/* debug [protocol_interface/required_method]: NewComputePipelineStateWithFunctionOptionsCompletionHandler */
	NewComputePipelineStateWithFunctionOptionsReflectionError(computeFunction unsafe.Pointer, options PipelineOption, reflection AutoreleasedComputePipelineReflection /* typedef */, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewComputePipelineStateWithFunctionOptionsReflectionError */
	NewCounterHeapWithDescriptorError(descriptor IMTL4CounterHeapDescriptor, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewCounterHeapWithDescriptorError */
	NewCounterSampleBufferWithDescriptorError(descriptor IMTLCounterSampleBufferDescriptor, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewCounterSampleBufferWithDescriptorError */
	NewDefaultLibrary() unsafe.Pointer/* debug [protocol_interface/required_method]: NewDefaultLibrary */
	NewDefaultLibraryWithBundleError(bundle foundation.Bundle, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewDefaultLibraryWithBundleError */
	NewDepthStencilStateWithDescriptor(descriptor IMTLDepthStencilDescriptor) unsafe.Pointer/* debug [protocol_interface/required_method]: NewDepthStencilStateWithDescriptor */
	NewDynamicLibraryError(library unsafe.Pointer, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewDynamicLibraryError */
	NewDynamicLibraryWithURLError(url objc.IObject /* cross-framework: NSURL */, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewDynamicLibraryWithURLError */
	NewEvent() unsafe.Pointer/* debug [protocol_interface/required_method]: NewEvent */
	NewFence() unsafe.Pointer/* debug [protocol_interface/required_method]: NewFence */
	NewHeapWithDescriptor(descriptor IMTLHeapDescriptor) unsafe.Pointer/* debug [protocol_interface/required_method]: NewHeapWithDescriptor */
	NewIndirectCommandBufferWithDescriptorMaxCommandCountOptions(descriptor IMTLIndirectCommandBufferDescriptor, maxCount uint, options ResourceOptions) unsafe.Pointer/* debug [protocol_interface/required_method]: NewIndirectCommandBufferWithDescriptorMaxCommandCountOptions */
	NewIOCommandQueueWithDescriptorError(descriptor IMTLIOCommandQueueDescriptor, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewIOCommandQueueWithDescriptorError */
	NewIOFileHandleWithURLError(url objc.IObject /* cross-framework: NSURL */, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewIOFileHandleWithURLError */
	NewIOFileHandleWithURLCompressionMethodError(url objc.IObject /* cross-framework: NSURL */, compressionMethod CompressionMethod, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewIOFileHandleWithURLCompressionMethodError */
	NewIOHandleWithURLError(url objc.IObject /* cross-framework: NSURL */, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewIOHandleWithURLError */
	NewIOHandleWithURLCompressionMethodError(url objc.IObject /* cross-framework: NSURL */, compressionMethod CompressionMethod, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewIOHandleWithURLCompressionMethodError */
	NewLibraryWithDataError(data objectivec.IObject, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewLibraryWithDataError */
	NewLibraryWithFileError(filepath objc.IObject /* cross-framework: NSString */, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewLibraryWithFileError */
	NewLibraryWithSourceOptionsError(source objc.IObject /* cross-framework: NSString */, options IMTLCompileOptions, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewLibraryWithSourceOptionsError */
	NewLibraryWithSourceOptionsCompletionHandler(source objc.IObject /* cross-framework: NSString */, options IMTLCompileOptions, completionHandler NewLibraryCompletionHandler /* not a class type */)/* debug [protocol_interface/required_method]: NewLibraryWithSourceOptionsCompletionHandler */
	NewLibraryWithStitchedDescriptorError(descriptor IMTLStitchedLibraryDescriptor, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewLibraryWithStitchedDescriptorError */
	NewLibraryWithStitchedDescriptorCompletionHandler(descriptor IMTLStitchedLibraryDescriptor, completionHandler NewLibraryCompletionHandler /* not a class type */)/* debug [protocol_interface/required_method]: NewLibraryWithStitchedDescriptorCompletionHandler */
	NewLibraryWithURLError(url objc.IObject /* cross-framework: NSURL */, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewLibraryWithURLError */
	NewLogStateWithDescriptorError(descriptor IMTLLogStateDescriptor, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewLogStateWithDescriptorError */
	NewMTL4CommandQueue() unsafe.Pointer/* debug [protocol_interface/required_method]: NewMTL4CommandQueue */
	NewMTL4CommandQueueWithDescriptorError(descriptor IMTL4CommandQueueDescriptor, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewMTL4CommandQueueWithDescriptorError */
	NewPipelineDataSetSerializerWithDescriptor(descriptor IMTL4PipelineDataSetSerializerDescriptor) unsafe.Pointer/* debug [protocol_interface/required_method]: NewPipelineDataSetSerializerWithDescriptor */
	NewRasterizationRateMapWithDescriptor(descriptor IMTLRasterizationRateMapDescriptor) unsafe.Pointer/* debug [protocol_interface/required_method]: NewRasterizationRateMapWithDescriptor */
	NewRenderPipelineStateWithDescriptorError(descriptor IMTLRenderPipelineDescriptor, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewRenderPipelineStateWithDescriptorError */
	NewRenderPipelineStateWithDescriptorCompletionHandler(descriptor IMTLRenderPipelineDescriptor, completionHandler NewRenderPipelineStateCompletionHandler /* not a class type */)/* debug [protocol_interface/required_method]: NewRenderPipelineStateWithDescriptorCompletionHandler */
	NewRenderPipelineStateWithMeshDescriptorOptionsCompletionHandler(descriptor IMTLMeshRenderPipelineDescriptor, options PipelineOption, completionHandler NewRenderPipelineStateWithReflectionCompletionHandler /* not a class type */)/* debug [protocol_interface/required_method]: NewRenderPipelineStateWithMeshDescriptorOptionsCompletionHandler */
	NewRenderPipelineStateWithDescriptorOptionsCompletionHandler(descriptor IMTLRenderPipelineDescriptor, options PipelineOption, completionHandler NewRenderPipelineStateWithReflectionCompletionHandler /* not a class type */)/* debug [protocol_interface/required_method]: NewRenderPipelineStateWithDescriptorOptionsCompletionHandler */
	NewRenderPipelineStateWithDescriptorOptionsReflectionError(descriptor IMTLRenderPipelineDescriptor, options PipelineOption, reflection AutoreleasedRenderPipelineReflection /* typedef */, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewRenderPipelineStateWithDescriptorOptionsReflectionError */
	NewRenderPipelineStateWithTileDescriptorOptionsCompletionHandler(descriptor IMTLTileRenderPipelineDescriptor, options PipelineOption, completionHandler NewRenderPipelineStateWithReflectionCompletionHandler /* not a class type */)/* debug [protocol_interface/required_method]: NewRenderPipelineStateWithTileDescriptorOptionsCompletionHandler */
	NewRenderPipelineStateWithTileDescriptorOptionsReflectionError(descriptor IMTLTileRenderPipelineDescriptor, options PipelineOption, reflection AutoreleasedRenderPipelineReflection /* typedef */, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewRenderPipelineStateWithTileDescriptorOptionsReflectionError */
	NewResidencySetWithDescriptorError(desc IMTLResidencySetDescriptor, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewResidencySetWithDescriptorError */
	NewSamplerStateWithDescriptor(descriptor IMTLSamplerDescriptor) unsafe.Pointer/* debug [protocol_interface/required_method]: NewSamplerStateWithDescriptor */
	NewSharedEvent() unsafe.Pointer/* debug [protocol_interface/required_method]: NewSharedEvent */
	NewSharedEventWithHandle(sharedEventHandle IMTLSharedEventHandle) unsafe.Pointer/* debug [protocol_interface/required_method]: NewSharedEventWithHandle */
	NewSharedTextureWithDescriptor(descriptor IMTLTextureDescriptor) unsafe.Pointer/* debug [protocol_interface/required_method]: NewSharedTextureWithDescriptor */
	NewSharedTextureWithHandle(sharedHandle IMTLSharedTextureHandle) unsafe.Pointer/* debug [protocol_interface/required_method]: NewSharedTextureWithHandle */
	NewTensorWithDescriptorError(descriptor IMTLTensorDescriptor, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewTensorWithDescriptorError */
	NewTextureWithDescriptor(descriptor IMTLTextureDescriptor) unsafe.Pointer/* debug [protocol_interface/required_method]: NewTextureWithDescriptor */
	NewTextureWithDescriptorIosurfacePlane(descriptor IMTLTextureDescriptor, iosurface SurfaceRef /* not a class type */, plane uint) unsafe.Pointer/* debug [protocol_interface/required_method]: NewTextureWithDescriptorIosurfacePlane */
	NewTextureViewPoolWithDescriptorError(descriptor IMTLResourceViewPoolDescriptor, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewTextureViewPoolWithDescriptorError */
	MinimumLinearTextureAlignmentForPixelFormat(format PixelFormat) uint/* debug [protocol_interface/required_method]: MinimumLinearTextureAlignmentForPixelFormat */
	MinimumTextureBufferAlignmentForPixelFormat(format PixelFormat) uint/* debug [protocol_interface/required_method]: MinimumTextureBufferAlignmentForPixelFormat */
	NewRenderPipelineStateWithMeshDescriptorOptionsReflectionError(descriptor IMTLMeshRenderPipelineDescriptor, options PipelineOption, reflection AutoreleasedRenderPipelineReflection /* typedef */, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewRenderPipelineStateWithMeshDescriptorOptionsReflectionError */
	QueryTimestampFrequency() uint64/* debug [protocol_interface/required_method]: QueryTimestampFrequency */
	SampleTimestampsGpuTimestamp(cpuTimestamp Timestamp /* typedef */, gpuTimestamp Timestamp /* typedef */)/* debug [protocol_interface/required_method]: SampleTimestampsGpuTimestamp */
	SizeOfCounterHeapEntry(type_ MTL4CounterHeapType) uint/* debug [protocol_interface/required_method]: SizeOfCounterHeapEntry */
	SparseTileSizeWithTextureTypePixelFormatSampleCountSparsePageSize(textureType TextureType, pixelFormat PixelFormat, sampleCount uint, sparsePageSize SparsePageSize) MTLSize/* debug [protocol_interface/required_method]: SparseTileSizeWithTextureTypePixelFormatSampleCountSparsePageSize */
	SparseTileSizeWithTextureTypePixelFormatSampleCount(textureType TextureType, pixelFormat PixelFormat, sampleCount uint) MTLSize/* debug [protocol_interface/required_method]: SparseTileSizeWithTextureTypePixelFormatSampleCount */
	SparseTileSizeInBytesForSparsePageSize(sparsePageSize SparsePageSize) uint/* debug [protocol_interface/required_method]: SparseTileSizeInBytesForSparsePageSize */
	SupportsCounterSampling(samplingPoint CounterSamplingPoint) bool/* debug [protocol_interface/required_method]: SupportsCounterSampling */
	SupportsFamily(gpuFamily GPUFamily) bool/* debug [protocol_interface/required_method]: SupportsFamily */
	SupportsFeatureSet(featureSet FeatureSet) bool/* debug [protocol_interface/required_method]: SupportsFeatureSet */
	SupportsRasterizationRateMapWithLayerCount(layerCount uint) bool/* debug [protocol_interface/required_method]: SupportsRasterizationRateMapWithLayerCount */
	SupportsTextureSampleCount(sampleCount uint) bool/* debug [protocol_interface/required_method]: SupportsTextureSampleCount */
	SupportsVertexAmplificationCount(count uint) bool/* debug [protocol_interface/required_method]: SupportsVertexAmplificationCount */
	TensorSizeAndAlignWithDescriptor(descriptor IMTLTensorDescriptor) MTLSizeAndAlign/* debug [protocol_interface/required_method]: TensorSizeAndAlignWithDescriptor */
	// Optional methods
	ConvertSparsePixelRegionsToTileRegionsWithTileSizeAlignmentModeNumRegions(pixelRegions []objc.IObject /* cross-framework: MTLRegion */, tileRegions []objc.IObject /* cross-framework: MTLRegion */, tileSize objc.IObject /* cross-framework: MTLSize */, mode SparseTextureRegionAlignmentMode, numRegions uint)
	HasConvertSparsePixelRegionsToTileRegionsWithTileSizeAlignmentModeNumRegions() bool
	ConvertSparseTileRegionsToPixelRegionsWithTileSizeNumRegions(tileRegions []objc.IObject /* cross-framework: MTLRegion */, pixelRegions []objc.IObject /* cross-framework: MTLRegion */, tileSize objc.IObject /* cross-framework: MTLSize */, numRegions uint)
	HasConvertSparseTileRegionsToPixelRegionsWithTileSizeNumRegions() bool
}
