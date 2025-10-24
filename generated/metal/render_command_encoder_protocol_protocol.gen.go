// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/corefoundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PRenderCommandEncoder is the MTLRenderCommandEncoder protocol interface.
//
// An interface that encodes a render pass into a command buffer, including all its draw calls and configuration.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.11+
//   - tvOS +
//   - visionOS 1.0+
//
// See: doc://com.apple.metal/documentation/Metal/MTLRenderCommandEncoder
type PRenderCommandEncoder interface {
	// Required methods
	DispatchThreadsPerTile(threadsPerTile objc.IObject /* cross-framework: MTLSize */)/* debug [protocol_interface/required_method]: DispatchThreadsPerTile */
	DrawIndexedPatchesPatchIndexBufferPatchIndexBufferOffsetControlPointIndexBufferControlPointIndexBufferOffsetIndirectBufferIndirectBufferOffset(numberOfPatchControlPoints uint, patchIndexBuffer unsafe.Pointer, patchIndexBufferOffset uint, controlPointIndexBuffer unsafe.Pointer, controlPointIndexBufferOffset uint, indirectBuffer unsafe.Pointer, indirectBufferOffset uint)/* debug [protocol_interface/required_method]: DrawIndexedPatchesPatchIndexBufferPatchIndexBufferOffsetControlPointIndexBufferControlPointIndexBufferOffsetIndirectBufferIndirectBufferOffset */
	DrawIndexedPatchesPatchStartPatchCountPatchIndexBufferPatchIndexBufferOffsetControlPointIndexBufferControlPointIndexBufferOffsetInstanceCountBaseInstance(numberOfPatchControlPoints uint, patchStart uint, patchCount uint, patchIndexBuffer unsafe.Pointer, patchIndexBufferOffset uint, controlPointIndexBuffer unsafe.Pointer, controlPointIndexBufferOffset uint, instanceCount uint, baseInstance uint)/* debug [protocol_interface/required_method]: DrawIndexedPatchesPatchStartPatchCountPatchIndexBufferPatchIndexBufferOffsetControlPointIndexBufferControlPointIndexBufferOffsetInstanceCountBaseInstance */
	DrawIndexedPrimitivesIndexCountIndexTypeIndexBufferIndexBufferOffset(primitiveType PrimitiveType, indexCount uint, indexType IndexType, indexBuffer unsafe.Pointer, indexBufferOffset uint)/* debug [protocol_interface/required_method]: DrawIndexedPrimitivesIndexCountIndexTypeIndexBufferIndexBufferOffset */
	DrawIndexedPrimitivesIndexCountIndexTypeIndexBufferIndexBufferOffsetInstanceCount(primitiveType PrimitiveType, indexCount uint, indexType IndexType, indexBuffer unsafe.Pointer, indexBufferOffset uint, instanceCount uint)/* debug [protocol_interface/required_method]: DrawIndexedPrimitivesIndexCountIndexTypeIndexBufferIndexBufferOffsetInstanceCount */
	DrawIndexedPrimitivesIndexCountIndexTypeIndexBufferIndexBufferOffsetInstanceCountBaseVertexBaseInstance(primitiveType PrimitiveType, indexCount uint, indexType IndexType, indexBuffer unsafe.Pointer, indexBufferOffset uint, instanceCount uint, baseVertex int, baseInstance uint)/* debug [protocol_interface/required_method]: DrawIndexedPrimitivesIndexCountIndexTypeIndexBufferIndexBufferOffsetInstanceCountBaseVertexBaseInstance */
	DrawIndexedPrimitivesIndexTypeIndexBufferIndexBufferOffsetIndirectBufferIndirectBufferOffset(primitiveType PrimitiveType, indexType IndexType, indexBuffer unsafe.Pointer, indexBufferOffset uint, indirectBuffer unsafe.Pointer, indirectBufferOffset uint)/* debug [protocol_interface/required_method]: DrawIndexedPrimitivesIndexTypeIndexBufferIndexBufferOffsetIndirectBufferIndirectBufferOffset */
	DrawMeshThreadgroupsThreadsPerObjectThreadgroupThreadsPerMeshThreadgroup(threadgroupsPerGrid objc.IObject /* cross-framework: MTLSize */, threadsPerObjectThreadgroup objc.IObject /* cross-framework: MTLSize */, threadsPerMeshThreadgroup objc.IObject /* cross-framework: MTLSize */)/* debug [protocol_interface/required_method]: DrawMeshThreadgroupsThreadsPerObjectThreadgroupThreadsPerMeshThreadgroup */
	DrawMeshThreadgroupsWithIndirectBufferIndirectBufferOffsetThreadsPerObjectThreadgroupThreadsPerMeshThreadgroup(indirectBuffer unsafe.Pointer, indirectBufferOffset uint, threadsPerObjectThreadgroup objc.IObject /* cross-framework: MTLSize */, threadsPerMeshThreadgroup objc.IObject /* cross-framework: MTLSize */)/* debug [protocol_interface/required_method]: DrawMeshThreadgroupsWithIndirectBufferIndirectBufferOffsetThreadsPerObjectThreadgroupThreadsPerMeshThreadgroup */
	DrawMeshThreadsThreadsPerObjectThreadgroupThreadsPerMeshThreadgroup(threadsPerGrid objc.IObject /* cross-framework: MTLSize */, threadsPerObjectThreadgroup objc.IObject /* cross-framework: MTLSize */, threadsPerMeshThreadgroup objc.IObject /* cross-framework: MTLSize */)/* debug [protocol_interface/required_method]: DrawMeshThreadsThreadsPerObjectThreadgroupThreadsPerMeshThreadgroup */
	DrawPatchesPatchIndexBufferPatchIndexBufferOffsetIndirectBufferIndirectBufferOffset(numberOfPatchControlPoints uint, patchIndexBuffer unsafe.Pointer, patchIndexBufferOffset uint, indirectBuffer unsafe.Pointer, indirectBufferOffset uint)/* debug [protocol_interface/required_method]: DrawPatchesPatchIndexBufferPatchIndexBufferOffsetIndirectBufferIndirectBufferOffset */
	DrawPatchesPatchStartPatchCountPatchIndexBufferPatchIndexBufferOffsetInstanceCountBaseInstance(numberOfPatchControlPoints uint, patchStart uint, patchCount uint, patchIndexBuffer unsafe.Pointer, patchIndexBufferOffset uint, instanceCount uint, baseInstance uint)/* debug [protocol_interface/required_method]: DrawPatchesPatchStartPatchCountPatchIndexBufferPatchIndexBufferOffsetInstanceCountBaseInstance */
	DrawPrimitivesIndirectBufferIndirectBufferOffset(primitiveType PrimitiveType, indirectBuffer unsafe.Pointer, indirectBufferOffset uint)/* debug [protocol_interface/required_method]: DrawPrimitivesIndirectBufferIndirectBufferOffset */
	DrawPrimitivesVertexStartVertexCount(primitiveType PrimitiveType, vertexStart uint, vertexCount uint)/* debug [protocol_interface/required_method]: DrawPrimitivesVertexStartVertexCount */
	DrawPrimitivesVertexStartVertexCountInstanceCount(primitiveType PrimitiveType, vertexStart uint, vertexCount uint, instanceCount uint)/* debug [protocol_interface/required_method]: DrawPrimitivesVertexStartVertexCountInstanceCount */
	DrawPrimitivesVertexStartVertexCountInstanceCountBaseInstance(primitiveType PrimitiveType, vertexStart uint, vertexCount uint, instanceCount uint, baseInstance uint)/* debug [protocol_interface/required_method]: DrawPrimitivesVertexStartVertexCountInstanceCountBaseInstance */
	ExecuteCommandsInBufferIndirectBufferIndirectBufferOffset(indirectCommandbuffer unsafe.Pointer, indirectRangeBuffer unsafe.Pointer, indirectBufferOffset uint)/* debug [protocol_interface/required_method]: ExecuteCommandsInBufferIndirectBufferIndirectBufferOffset */
	ExecuteCommandsInBufferWithRange(indirectCommandBuffer unsafe.Pointer, executionRange corefoundation.Range)/* debug [protocol_interface/required_method]: ExecuteCommandsInBufferWithRange */
	MemoryBarrierWithScopeAfterStagesBeforeStages(scope BarrierScope, after RenderStages, before RenderStages)/* debug [protocol_interface/required_method]: MemoryBarrierWithScopeAfterStagesBeforeStages */
	MemoryBarrierWithResourcesCountAfterStagesBeforeStages(resources []objc.ID, count uint, after RenderStages, before RenderStages)/* debug [protocol_interface/required_method]: MemoryBarrierWithResourcesCountAfterStagesBeforeStages */
	SampleCountersInBufferAtSampleIndexWithBarrier(sampleBuffer unsafe.Pointer, sampleIndex uint, barrier bool)/* debug [protocol_interface/required_method]: SampleCountersInBufferAtSampleIndexWithBarrier */
	SetBlendColorRedGreenBlueAlpha(red float32, green float32, blue float32, alpha float32)/* debug [protocol_interface/required_method]: SetBlendColorRedGreenBlueAlpha */
	SetColorAttachmentMap(mapping IMTLLogicalToPhysicalColorAttachmentMap)/* debug [protocol_interface/required_method]: SetColorAttachmentMap */
	SetColorStoreActionAtIndex(storeAction StoreAction, colorAttachmentIndex uint)/* debug [protocol_interface/required_method]: SetColorStoreActionAtIndex */
	SetColorStoreActionOptionsAtIndex(storeActionOptions StoreActionOptions, colorAttachmentIndex uint)/* debug [protocol_interface/required_method]: SetColorStoreActionOptionsAtIndex */
	SetCullMode(cullMode CullMode)/* debug [protocol_interface/required_method]: SetCullMode */
	SetDepthBiasSlopeScaleClamp(depthBias float32, slopeScale float32, clamp float32)/* debug [protocol_interface/required_method]: SetDepthBiasSlopeScaleClamp */
	SetDepthClipMode(depthClipMode DepthClipMode)/* debug [protocol_interface/required_method]: SetDepthClipMode */
	SetDepthStencilState(depthStencilState unsafe.Pointer)/* debug [protocol_interface/required_method]: SetDepthStencilState */
	SetDepthStoreAction(storeAction StoreAction)/* debug [protocol_interface/required_method]: SetDepthStoreAction */
	SetDepthStoreActionOptions(storeActionOptions StoreActionOptions)/* debug [protocol_interface/required_method]: SetDepthStoreActionOptions */
	SetDepthTestMinBoundMaxBound(minBound float32, maxBound float32)/* debug [protocol_interface/required_method]: SetDepthTestMinBoundMaxBound */
	SetFragmentAccelerationStructureAtBufferIndex(accelerationStructure unsafe.Pointer, bufferIndex uint)/* debug [protocol_interface/required_method]: SetFragmentAccelerationStructureAtBufferIndex */
	SetFragmentBufferOffsetAtIndex(buffer unsafe.Pointer, offset uint, index uint)/* debug [protocol_interface/required_method]: SetFragmentBufferOffsetAtIndex */
	SetFragmentBuffersOffsetsWithRange(buffers []objc.ID, offsets uint, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetFragmentBuffersOffsetsWithRange */
	SetFragmentBytesLengthAtIndex(bytes objectivec.IObject, length uint, index uint)/* debug [protocol_interface/required_method]: SetFragmentBytesLengthAtIndex */
	SetFragmentIntersectionFunctionTableAtBufferIndex(intersectionFunctionTable unsafe.Pointer, bufferIndex uint)/* debug [protocol_interface/required_method]: SetFragmentIntersectionFunctionTableAtBufferIndex */
	SetFragmentIntersectionFunctionTablesWithBufferRange(intersectionFunctionTables []objc.ID, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetFragmentIntersectionFunctionTablesWithBufferRange */
	SetFragmentSamplerStateAtIndex(sampler unsafe.Pointer, index uint)/* debug [protocol_interface/required_method]: SetFragmentSamplerStateAtIndex */
	SetFragmentSamplerStateLodMinClampLodMaxClampAtIndex(sampler unsafe.Pointer, lodMinClamp float32, lodMaxClamp float32, index uint)/* debug [protocol_interface/required_method]: SetFragmentSamplerStateLodMinClampLodMaxClampAtIndex */
	SetFragmentSamplerStatesLodMinClampsLodMaxClampsWithRange(samplers []objc.ID, lodMinClamps []float32, lodMaxClamps []float32, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetFragmentSamplerStatesLodMinClampsLodMaxClampsWithRange */
	SetFragmentSamplerStatesWithRange(samplers []objc.ID, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetFragmentSamplerStatesWithRange */
	SetFragmentTextureAtIndex(texture unsafe.Pointer, index uint)/* debug [protocol_interface/required_method]: SetFragmentTextureAtIndex */
	SetFragmentTexturesWithRange(textures []objc.ID, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetFragmentTexturesWithRange */
	SetFragmentVisibleFunctionTableAtBufferIndex(functionTable unsafe.Pointer, bufferIndex uint)/* debug [protocol_interface/required_method]: SetFragmentVisibleFunctionTableAtBufferIndex */
	SetFragmentVisibleFunctionTablesWithBufferRange(functionTables []objc.ID, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetFragmentVisibleFunctionTablesWithBufferRange */
	SetFrontFacingWinding(frontFacingWinding Winding)/* debug [protocol_interface/required_method]: SetFrontFacingWinding */
	SetMeshBufferOffsetAtIndex(buffer unsafe.Pointer, offset uint, index uint)/* debug [protocol_interface/required_method]: SetMeshBufferOffsetAtIndex */
	SetMeshBuffersOffsetsWithRange(buffers []objc.ID, offsets uint, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetMeshBuffersOffsetsWithRange */
	SetMeshBytesLengthAtIndex(bytes objectivec.IObject, length uint, index uint)/* debug [protocol_interface/required_method]: SetMeshBytesLengthAtIndex */
	SetMeshSamplerStateAtIndex(sampler unsafe.Pointer, index uint)/* debug [protocol_interface/required_method]: SetMeshSamplerStateAtIndex */
	SetMeshSamplerStateLodMinClampLodMaxClampAtIndex(sampler unsafe.Pointer, lodMinClamp float32, lodMaxClamp float32, index uint)/* debug [protocol_interface/required_method]: SetMeshSamplerStateLodMinClampLodMaxClampAtIndex */
	SetMeshSamplerStatesLodMinClampsLodMaxClampsWithRange(samplers []objc.ID, lodMinClamps []float32, lodMaxClamps []float32, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetMeshSamplerStatesLodMinClampsLodMaxClampsWithRange */
	SetMeshSamplerStatesWithRange(samplers []objc.ID, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetMeshSamplerStatesWithRange */
	SetMeshTextureAtIndex(texture unsafe.Pointer, index uint)/* debug [protocol_interface/required_method]: SetMeshTextureAtIndex */
	SetMeshTexturesWithRange(textures []objc.ID, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetMeshTexturesWithRange */
	SetObjectBufferOffsetAtIndex(buffer unsafe.Pointer, offset uint, index uint)/* debug [protocol_interface/required_method]: SetObjectBufferOffsetAtIndex */
	SetObjectBuffersOffsetsWithRange(buffers []objc.ID, offsets uint, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetObjectBuffersOffsetsWithRange */
	SetObjectBytesLengthAtIndex(bytes objectivec.IObject, length uint, index uint)/* debug [protocol_interface/required_method]: SetObjectBytesLengthAtIndex */
	SetObjectSamplerStateAtIndex(sampler unsafe.Pointer, index uint)/* debug [protocol_interface/required_method]: SetObjectSamplerStateAtIndex */
	SetObjectSamplerStateLodMinClampLodMaxClampAtIndex(sampler unsafe.Pointer, lodMinClamp float32, lodMaxClamp float32, index uint)/* debug [protocol_interface/required_method]: SetObjectSamplerStateLodMinClampLodMaxClampAtIndex */
	SetObjectSamplerStatesLodMinClampsLodMaxClampsWithRange(samplers []objc.ID, lodMinClamps []float32, lodMaxClamps []float32, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetObjectSamplerStatesLodMinClampsLodMaxClampsWithRange */
	SetObjectSamplerStatesWithRange(samplers []objc.ID, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetObjectSamplerStatesWithRange */
	SetObjectTextureAtIndex(texture unsafe.Pointer, index uint)/* debug [protocol_interface/required_method]: SetObjectTextureAtIndex */
	SetObjectTexturesWithRange(textures []objc.ID, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetObjectTexturesWithRange */
	SetObjectThreadgroupMemoryLengthAtIndex(length uint, index uint)/* debug [protocol_interface/required_method]: SetObjectThreadgroupMemoryLengthAtIndex */
	SetRenderPipelineState(pipelineState unsafe.Pointer)/* debug [protocol_interface/required_method]: SetRenderPipelineState */
	SetScissorRect(rect objc.IObject /* cross-framework: MTLScissorRect */)/* debug [protocol_interface/required_method]: SetScissorRect */
	SetScissorRectsCount(scissorRects []objc.IObject /* cross-framework: MTLScissorRect */, count uint)/* debug [protocol_interface/required_method]: SetScissorRectsCount */
	SetStencilReferenceValue(referenceValue uint32 /* not a class type */)/* debug [protocol_interface/required_method]: SetStencilReferenceValue */
	SetStencilFrontReferenceValueBackReferenceValue(frontReferenceValue uint32 /* not a class type */, backReferenceValue uint32 /* not a class type */)/* debug [protocol_interface/required_method]: SetStencilFrontReferenceValueBackReferenceValue */
	SetStencilStoreAction(storeAction StoreAction)/* debug [protocol_interface/required_method]: SetStencilStoreAction */
	SetStencilStoreActionOptions(storeActionOptions StoreActionOptions)/* debug [protocol_interface/required_method]: SetStencilStoreActionOptions */
	SetTessellationFactorBufferOffsetInstanceStride(buffer unsafe.Pointer, offset uint, instanceStride uint)/* debug [protocol_interface/required_method]: SetTessellationFactorBufferOffsetInstanceStride */
	SetTessellationFactorScale(scale float32)/* debug [protocol_interface/required_method]: SetTessellationFactorScale */
	SetThreadgroupMemoryLengthOffsetAtIndex(length uint, offset uint, index uint)/* debug [protocol_interface/required_method]: SetThreadgroupMemoryLengthOffsetAtIndex */
	SetTileAccelerationStructureAtBufferIndex(accelerationStructure unsafe.Pointer, bufferIndex uint)/* debug [protocol_interface/required_method]: SetTileAccelerationStructureAtBufferIndex */
	SetTileBufferOffsetAtIndex(buffer unsafe.Pointer, offset uint, index uint)/* debug [protocol_interface/required_method]: SetTileBufferOffsetAtIndex */
	SetTileBuffersOffsetsWithRange(buffers []objc.ID, offsets uint, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetTileBuffersOffsetsWithRange */
	SetTileBytesLengthAtIndex(bytes objectivec.IObject, length uint, index uint)/* debug [protocol_interface/required_method]: SetTileBytesLengthAtIndex */
	SetTileIntersectionFunctionTableAtBufferIndex(intersectionFunctionTable unsafe.Pointer, bufferIndex uint)/* debug [protocol_interface/required_method]: SetTileIntersectionFunctionTableAtBufferIndex */
	SetTileIntersectionFunctionTablesWithBufferRange(intersectionFunctionTables []objc.ID, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetTileIntersectionFunctionTablesWithBufferRange */
	SetTileSamplerStateAtIndex(sampler unsafe.Pointer, index uint)/* debug [protocol_interface/required_method]: SetTileSamplerStateAtIndex */
	SetTileSamplerStateLodMinClampLodMaxClampAtIndex(sampler unsafe.Pointer, lodMinClamp float32, lodMaxClamp float32, index uint)/* debug [protocol_interface/required_method]: SetTileSamplerStateLodMinClampLodMaxClampAtIndex */
	SetTileSamplerStatesLodMinClampsLodMaxClampsWithRange(samplers []objc.ID, lodMinClamps []float32, lodMaxClamps []float32, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetTileSamplerStatesLodMinClampsLodMaxClampsWithRange */
	SetTileSamplerStatesWithRange(samplers []objc.ID, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetTileSamplerStatesWithRange */
	SetTileTextureAtIndex(texture unsafe.Pointer, index uint)/* debug [protocol_interface/required_method]: SetTileTextureAtIndex */
	SetTileTexturesWithRange(textures []objc.ID, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetTileTexturesWithRange */
	SetTileVisibleFunctionTableAtBufferIndex(functionTable unsafe.Pointer, bufferIndex uint)/* debug [protocol_interface/required_method]: SetTileVisibleFunctionTableAtBufferIndex */
	SetTileVisibleFunctionTablesWithBufferRange(functionTables []objc.ID, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetTileVisibleFunctionTablesWithBufferRange */
	SetTriangleFillMode(fillMode TriangleFillMode)/* debug [protocol_interface/required_method]: SetTriangleFillMode */
	SetVertexAccelerationStructureAtBufferIndex(accelerationStructure unsafe.Pointer, bufferIndex uint)/* debug [protocol_interface/required_method]: SetVertexAccelerationStructureAtBufferIndex */
	SetVertexAmplificationCountViewMappings(count uint, viewMappings objc.IObject /* cross-framework: MTLVertexAmplificationViewMapping */)/* debug [protocol_interface/required_method]: SetVertexAmplificationCountViewMappings */
	SetVertexBufferOffsetAttributeStrideAtIndex(buffer unsafe.Pointer, offset uint, stride uint, index uint)/* debug [protocol_interface/required_method]: SetVertexBufferOffsetAttributeStrideAtIndex */
	SetVertexBufferOffsetAtIndex(buffer unsafe.Pointer, offset uint, index uint)/* debug [protocol_interface/required_method]: SetVertexBufferOffsetAtIndex */
	SetVertexBuffersOffsetsAttributeStridesWithRange(buffers []objc.ID, offsets uint, strides uint, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetVertexBuffersOffsetsAttributeStridesWithRange */
	SetVertexBuffersOffsetsWithRange(buffers []objc.ID, offsets uint, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetVertexBuffersOffsetsWithRange */
	SetVertexBytesLengthAttributeStrideAtIndex(bytes objectivec.IObject, length uint, stride uint, index uint)/* debug [protocol_interface/required_method]: SetVertexBytesLengthAttributeStrideAtIndex */
	SetVertexBytesLengthAtIndex(bytes objectivec.IObject, length uint, index uint)/* debug [protocol_interface/required_method]: SetVertexBytesLengthAtIndex */
	SetVertexIntersectionFunctionTableAtBufferIndex(intersectionFunctionTable unsafe.Pointer, bufferIndex uint)/* debug [protocol_interface/required_method]: SetVertexIntersectionFunctionTableAtBufferIndex */
	SetVertexIntersectionFunctionTablesWithBufferRange(intersectionFunctionTables []objc.ID, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetVertexIntersectionFunctionTablesWithBufferRange */
	SetVertexSamplerStateAtIndex(sampler unsafe.Pointer, index uint)/* debug [protocol_interface/required_method]: SetVertexSamplerStateAtIndex */
	SetVertexSamplerStateLodMinClampLodMaxClampAtIndex(sampler unsafe.Pointer, lodMinClamp float32, lodMaxClamp float32, index uint)/* debug [protocol_interface/required_method]: SetVertexSamplerStateLodMinClampLodMaxClampAtIndex */
	SetVertexSamplerStatesLodMinClampsLodMaxClampsWithRange(samplers []objc.ID, lodMinClamps []float32, lodMaxClamps []float32, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetVertexSamplerStatesLodMinClampsLodMaxClampsWithRange */
	SetVertexSamplerStatesWithRange(samplers []objc.ID, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetVertexSamplerStatesWithRange */
	SetVertexTextureAtIndex(texture unsafe.Pointer, index uint)/* debug [protocol_interface/required_method]: SetVertexTextureAtIndex */
	SetVertexTexturesWithRange(textures []objc.ID, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetVertexTexturesWithRange */
	SetVertexVisibleFunctionTableAtBufferIndex(functionTable unsafe.Pointer, bufferIndex uint)/* debug [protocol_interface/required_method]: SetVertexVisibleFunctionTableAtBufferIndex */
	SetVertexVisibleFunctionTablesWithBufferRange(functionTables []objc.ID, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetVertexVisibleFunctionTablesWithBufferRange */
	SetViewport(viewport objc.IObject /* cross-framework: MTLViewport */)/* debug [protocol_interface/required_method]: SetViewport */
	SetViewportsCount(viewports []objc.IObject /* cross-framework: MTLViewport */, count uint)/* debug [protocol_interface/required_method]: SetViewportsCount */
	SetVisibilityResultModeOffset(mode VisibilityResultMode, offset uint)/* debug [protocol_interface/required_method]: SetVisibilityResultModeOffset */
	TextureBarrier()/* debug [protocol_interface/required_method]: TextureBarrier */
	UpdateFenceAfterStages(fence unsafe.Pointer, stages RenderStages)/* debug [protocol_interface/required_method]: UpdateFenceAfterStages */
	UseHeap(heap unsafe.Pointer)/* debug [protocol_interface/required_method]: UseHeap */
	UseHeapStages(heap unsafe.Pointer, stages RenderStages)/* debug [protocol_interface/required_method]: UseHeapStages */
	UseHeapsCount(heaps []objc.ID, count uint)/* debug [protocol_interface/required_method]: UseHeapsCount */
	UseHeapsCountStages(heaps []objc.ID, count uint, stages RenderStages)/* debug [protocol_interface/required_method]: UseHeapsCountStages */
	UseResourceUsage(resource unsafe.Pointer, usage ResourceUsage)/* debug [protocol_interface/required_method]: UseResourceUsage */
	UseResourceUsageStages(resource unsafe.Pointer, usage ResourceUsage, stages RenderStages)/* debug [protocol_interface/required_method]: UseResourceUsageStages */
	UseResourcesCountUsage(resources []objc.ID, count uint, usage ResourceUsage)/* debug [protocol_interface/required_method]: UseResourcesCountUsage */
	UseResourcesCountUsageStages(resources []objc.ID, count uint, usage ResourceUsage, stages RenderStages)/* debug [protocol_interface/required_method]: UseResourcesCountUsageStages */
	WaitForFenceBeforeStages(fence unsafe.Pointer, stages RenderStages)/* debug [protocol_interface/required_method]: WaitForFenceBeforeStages */
}
