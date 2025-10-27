// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
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
	DispatchThreadsPerTile(threadsPerTile Size)
	DrawIndexedPrimitivesIndexCountIndexTypeIndexBufferIndexBufferLength(primitiveType PrimitiveType, indexCount uint, indexType IndexType, indexBuffer GPUAddress, indexBufferLength uint)
	DrawIndexedPrimitivesIndexCountIndexTypeIndexBufferIndexBufferLengthInstanceCount(primitiveType PrimitiveType, indexCount uint, indexType IndexType, indexBuffer GPUAddress, indexBufferLength uint, instanceCount uint)
	DrawIndexedPrimitivesIndexCountIndexTypeIndexBufferIndexBufferLengthInstanceCountBaseVertexBaseInstance(primitiveType PrimitiveType, indexCount uint, indexType IndexType, indexBuffer GPUAddress, indexBufferLength uint, instanceCount uint, baseVertex int, baseInstance uint)
	DrawIndexedPrimitivesIndexTypeIndexBufferIndexBufferLengthIndirectBuffer(primitiveType PrimitiveType, indexType IndexType, indexBuffer GPUAddress, indexBufferLength uint, indirectBuffer GPUAddress)
	DrawMeshThreadgroupsWithIndirectBufferThreadsPerObjectThreadgroupThreadsPerMeshThreadgroup(indirectBuffer GPUAddress, threadsPerObjectThreadgroup Size, threadsPerMeshThreadgroup Size)
	DrawMeshThreadgroupsThreadsPerObjectThreadgroupThreadsPerMeshThreadgroup(threadgroupsPerGrid Size, threadsPerObjectThreadgroup Size, threadsPerMeshThreadgroup Size)
	DrawMeshThreadsThreadsPerObjectThreadgroupThreadsPerMeshThreadgroup(threadsPerGrid Size, threadsPerObjectThreadgroup Size, threadsPerMeshThreadgroup Size)
	DrawPrimitivesIndirectBuffer(primitiveType PrimitiveType, indirectBuffer GPUAddress)
	DrawPrimitivesVertexStartVertexCount(primitiveType PrimitiveType, vertexStart uint, vertexCount uint)
	DrawPrimitivesVertexStartVertexCountInstanceCount(primitiveType PrimitiveType, vertexStart uint, vertexCount uint, instanceCount uint)
	DrawPrimitivesVertexStartVertexCountInstanceCountBaseInstance(primitiveType PrimitiveType, vertexStart uint, vertexCount uint, instanceCount uint, baseInstance uint)
	ExecuteCommandsInBufferIndirectBuffer(indirectCommandBuffer unsafe.Pointer, indirectRangeBuffer GPUAddress)
	ExecuteCommandsInBufferWithRange(indirectCommandBuffer unsafe.Pointer, executionRange foundation.Range)
	SetArgumentTableAtStages(argumentTable unsafe.Pointer, stages RenderStages)
	SetBlendColorRedGreenBlueAlpha(red float32, green float32, blue float32, alpha float32)
	SetColorAttachmentMap(mapping IMTLLogicalToPhysicalColorAttachmentMap)
	SetColorStoreActionAtIndex(storeAction StoreAction, colorAttachmentIndex uint)
	SetCullMode(cullMode CullMode)
	SetDepthBiasSlopeScaleClamp(depthBias float32, slopeScale float32, clamp float32)
	SetDepthClipMode(depthClipMode DepthClipMode)
	SetDepthStencilState(depthStencilState unsafe.Pointer)
	SetDepthStoreAction(storeAction StoreAction)
	SetDepthTestMinBoundMaxBound(minBound float32, maxBound float32)
	SetFrontFacingWinding(frontFacingWinding Winding)
	SetObjectThreadgroupMemoryLengthAtIndex(length uint, index uint)
	SetRenderPipelineState(pipelineState unsafe.Pointer)
	SetScissorRect(rect ScissorRect)
	SetScissorRectsCount(scissorRects []MTLScissorRect, count uint)
	SetStencilReferenceValue(referenceValue uint32 /* not a class type */)
	SetStencilFrontReferenceValueBackReferenceValue(frontReferenceValue uint32 /* not a class type */, backReferenceValue uint32 /* not a class type */)
	SetStencilStoreAction(storeAction StoreAction)
	SetThreadgroupMemoryLengthOffsetAtIndex(length uint, offset uint, index uint)
	SetTriangleFillMode(fillMode TriangleFillMode)
	SetVertexAmplificationCountViewMappings(count uint, viewMappings VertexAmplificationViewMapping)
	SetViewport(viewport Viewport)
	SetViewportsCount(viewports []MTLViewport, count uint)
	SetVisibilityResultModeOffset(mode VisibilityResultMode, offset uint)
	WriteTimestampWithGranularityAfterStageIntoHeapAtIndex(granularity MTL4TimestampGranularity, stage RenderStages, counterHeap unsafe.Pointer, index uint)
}
