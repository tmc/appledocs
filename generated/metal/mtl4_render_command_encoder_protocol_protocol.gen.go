// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/corefoundation"
)

// PMTL4RenderCommandEncoder is the MTL4RenderCommandEncoder protocol interface.
//
// Encodes a render pass into a command buffer, including all its draw calls and configuration.
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//
// See: doc://com.apple.metal/documentation/Metal/MTL4RenderCommandEncoder
type PMTL4RenderCommandEncoder interface {
	// Required methods
	DispatchThreadsPerTile(threadsPerTile objc.IObject /* cross-framework: MTLSize */)/* debug [protocol_interface/required_method]: DispatchThreadsPerTile */
	DrawIndexedPrimitivesIndexCountIndexTypeIndexBufferIndexBufferLength(primitiveType PrimitiveType, indexCount uint, indexType IndexType, indexBuffer GPUAddress /* typedef */, indexBufferLength uint)/* debug [protocol_interface/required_method]: DrawIndexedPrimitivesIndexCountIndexTypeIndexBufferIndexBufferLength */
	DrawIndexedPrimitivesIndexCountIndexTypeIndexBufferIndexBufferLengthInstanceCount(primitiveType PrimitiveType, indexCount uint, indexType IndexType, indexBuffer GPUAddress /* typedef */, indexBufferLength uint, instanceCount uint)/* debug [protocol_interface/required_method]: DrawIndexedPrimitivesIndexCountIndexTypeIndexBufferIndexBufferLengthInstanceCount */
	DrawIndexedPrimitivesIndexCountIndexTypeIndexBufferIndexBufferLengthInstanceCountBaseVertexBaseInstance(primitiveType PrimitiveType, indexCount uint, indexType IndexType, indexBuffer GPUAddress /* typedef */, indexBufferLength uint, instanceCount uint, baseVertex int, baseInstance uint)/* debug [protocol_interface/required_method]: DrawIndexedPrimitivesIndexCountIndexTypeIndexBufferIndexBufferLengthInstanceCountBaseVertexBaseInstance */
	DrawIndexedPrimitivesIndexTypeIndexBufferIndexBufferLengthIndirectBuffer(primitiveType PrimitiveType, indexType IndexType, indexBuffer GPUAddress /* typedef */, indexBufferLength uint, indirectBuffer GPUAddress /* typedef */)/* debug [protocol_interface/required_method]: DrawIndexedPrimitivesIndexTypeIndexBufferIndexBufferLengthIndirectBuffer */
	DrawMeshThreadgroupsWithIndirectBufferThreadsPerObjectThreadgroupThreadsPerMeshThreadgroup(indirectBuffer GPUAddress /* typedef */, threadsPerObjectThreadgroup objc.IObject /* cross-framework: MTLSize */, threadsPerMeshThreadgroup objc.IObject /* cross-framework: MTLSize */)/* debug [protocol_interface/required_method]: DrawMeshThreadgroupsWithIndirectBufferThreadsPerObjectThreadgroupThreadsPerMeshThreadgroup */
	DrawMeshThreadgroupsThreadsPerObjectThreadgroupThreadsPerMeshThreadgroup(threadgroupsPerGrid objc.IObject /* cross-framework: MTLSize */, threadsPerObjectThreadgroup objc.IObject /* cross-framework: MTLSize */, threadsPerMeshThreadgroup objc.IObject /* cross-framework: MTLSize */)/* debug [protocol_interface/required_method]: DrawMeshThreadgroupsThreadsPerObjectThreadgroupThreadsPerMeshThreadgroup */
	DrawMeshThreadsThreadsPerObjectThreadgroupThreadsPerMeshThreadgroup(threadsPerGrid objc.IObject /* cross-framework: MTLSize */, threadsPerObjectThreadgroup objc.IObject /* cross-framework: MTLSize */, threadsPerMeshThreadgroup objc.IObject /* cross-framework: MTLSize */)/* debug [protocol_interface/required_method]: DrawMeshThreadsThreadsPerObjectThreadgroupThreadsPerMeshThreadgroup */
	DrawPrimitivesIndirectBuffer(primitiveType PrimitiveType, indirectBuffer GPUAddress /* typedef */)/* debug [protocol_interface/required_method]: DrawPrimitivesIndirectBuffer */
	DrawPrimitivesVertexStartVertexCount(primitiveType PrimitiveType, vertexStart uint, vertexCount uint)/* debug [protocol_interface/required_method]: DrawPrimitivesVertexStartVertexCount */
	DrawPrimitivesVertexStartVertexCountInstanceCount(primitiveType PrimitiveType, vertexStart uint, vertexCount uint, instanceCount uint)/* debug [protocol_interface/required_method]: DrawPrimitivesVertexStartVertexCountInstanceCount */
	DrawPrimitivesVertexStartVertexCountInstanceCountBaseInstance(primitiveType PrimitiveType, vertexStart uint, vertexCount uint, instanceCount uint, baseInstance uint)/* debug [protocol_interface/required_method]: DrawPrimitivesVertexStartVertexCountInstanceCountBaseInstance */
	ExecuteCommandsInBufferIndirectBuffer(indirectCommandBuffer unsafe.Pointer, indirectRangeBuffer GPUAddress /* typedef */)/* debug [protocol_interface/required_method]: ExecuteCommandsInBufferIndirectBuffer */
	ExecuteCommandsInBufferWithRange(indirectCommandBuffer unsafe.Pointer, executionRange corefoundation.Range)/* debug [protocol_interface/required_method]: ExecuteCommandsInBufferWithRange */
	SetArgumentTableAtStages(argumentTable unsafe.Pointer, stages RenderStages)/* debug [protocol_interface/required_method]: SetArgumentTableAtStages */
	SetBlendColorRedGreenBlueAlpha(red float32, green float32, blue float32, alpha float32)/* debug [protocol_interface/required_method]: SetBlendColorRedGreenBlueAlpha */
	SetColorAttachmentMap(mapping IMTLLogicalToPhysicalColorAttachmentMap)/* debug [protocol_interface/required_method]: SetColorAttachmentMap */
	SetColorStoreActionAtIndex(storeAction StoreAction, colorAttachmentIndex uint)/* debug [protocol_interface/required_method]: SetColorStoreActionAtIndex */
	SetCullMode(cullMode CullMode)/* debug [protocol_interface/required_method]: SetCullMode */
	SetDepthBiasSlopeScaleClamp(depthBias float32, slopeScale float32, clamp float32)/* debug [protocol_interface/required_method]: SetDepthBiasSlopeScaleClamp */
	SetDepthClipMode(depthClipMode DepthClipMode)/* debug [protocol_interface/required_method]: SetDepthClipMode */
	SetDepthStencilState(depthStencilState unsafe.Pointer)/* debug [protocol_interface/required_method]: SetDepthStencilState */
	SetDepthStoreAction(storeAction StoreAction)/* debug [protocol_interface/required_method]: SetDepthStoreAction */
	SetDepthTestMinBoundMaxBound(minBound float32, maxBound float32)/* debug [protocol_interface/required_method]: SetDepthTestMinBoundMaxBound */
	SetFrontFacingWinding(frontFacingWinding Winding)/* debug [protocol_interface/required_method]: SetFrontFacingWinding */
	SetObjectThreadgroupMemoryLengthAtIndex(length uint, index uint)/* debug [protocol_interface/required_method]: SetObjectThreadgroupMemoryLengthAtIndex */
	SetRenderPipelineState(pipelineState unsafe.Pointer)/* debug [protocol_interface/required_method]: SetRenderPipelineState */
	SetScissorRect(rect objc.IObject /* cross-framework: MTLScissorRect */)/* debug [protocol_interface/required_method]: SetScissorRect */
	SetScissorRectsCount(scissorRects []objc.IObject /* cross-framework: MTLScissorRect */, count uint)/* debug [protocol_interface/required_method]: SetScissorRectsCount */
	SetStencilReferenceValue(referenceValue uint32 /* not a class type */)/* debug [protocol_interface/required_method]: SetStencilReferenceValue */
	SetStencilFrontReferenceValueBackReferenceValue(frontReferenceValue uint32 /* not a class type */, backReferenceValue uint32 /* not a class type */)/* debug [protocol_interface/required_method]: SetStencilFrontReferenceValueBackReferenceValue */
	SetStencilStoreAction(storeAction StoreAction)/* debug [protocol_interface/required_method]: SetStencilStoreAction */
	SetThreadgroupMemoryLengthOffsetAtIndex(length uint, offset uint, index uint)/* debug [protocol_interface/required_method]: SetThreadgroupMemoryLengthOffsetAtIndex */
	SetTriangleFillMode(fillMode TriangleFillMode)/* debug [protocol_interface/required_method]: SetTriangleFillMode */
	SetVertexAmplificationCountViewMappings(count uint, viewMappings objc.IObject /* cross-framework: MTLVertexAmplificationViewMapping */)/* debug [protocol_interface/required_method]: SetVertexAmplificationCountViewMappings */
	SetViewport(viewport objc.IObject /* cross-framework: MTLViewport */)/* debug [protocol_interface/required_method]: SetViewport */
	SetViewportsCount(viewports []objc.IObject /* cross-framework: MTLViewport */, count uint)/* debug [protocol_interface/required_method]: SetViewportsCount */
	SetVisibilityResultModeOffset(mode VisibilityResultMode, offset uint)/* debug [protocol_interface/required_method]: SetVisibilityResultModeOffset */
	WriteTimestampWithGranularityAfterStageIntoHeapAtIndex(granularity MTL4TimestampGranularity, stage RenderStages, counterHeap unsafe.Pointer, index uint)/* debug [protocol_interface/required_method]: WriteTimestampWithGranularityAfterStageIntoHeapAtIndex */
}
