// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"
)

// PIndirectRenderCommand is the MTLIndirectRenderCommand protocol interface.
//
// A render command in an indirect command buffer.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.metal/documentation/Metal/MTLIndirectRenderCommand
type PIndirectRenderCommand interface {
	// Required methods
	ClearBarrier()
	DrawIndexedPatchesPatchStartPatchCountPatchIndexBufferPatchIndexBufferOffsetControlPointIndexBufferControlPointIndexBufferOffsetInstanceCountBaseInstanceTessellationFactorBufferTessellationFactorBufferOffsetTessellationFactorBufferInstanceStride(numberOfPatchControlPoints uint, patchStart uint, patchCount uint, patchIndexBuffer unsafe.Pointer, patchIndexBufferOffset uint, controlPointIndexBuffer unsafe.Pointer, controlPointIndexBufferOffset uint, instanceCount uint, baseInstance uint, buffer unsafe.Pointer, offset uint, instanceStride uint)
	DrawIndexedPrimitivesIndexCountIndexTypeIndexBufferIndexBufferOffsetInstanceCountBaseVertexBaseInstance(primitiveType PrimitiveType, indexCount uint, indexType IndexType, indexBuffer unsafe.Pointer, indexBufferOffset uint, instanceCount uint, baseVertex int, baseInstance uint)
	DrawMeshThreadgroupsThreadsPerObjectThreadgroupThreadsPerMeshThreadgroup(threadgroupsPerGrid Size, threadsPerObjectThreadgroup Size, threadsPerMeshThreadgroup Size)
	DrawMeshThreadsThreadsPerObjectThreadgroupThreadsPerMeshThreadgroup(threadsPerGrid Size, threadsPerObjectThreadgroup Size, threadsPerMeshThreadgroup Size)
	DrawPatchesPatchStartPatchCountPatchIndexBufferPatchIndexBufferOffsetInstanceCountBaseInstanceTessellationFactorBufferTessellationFactorBufferOffsetTessellationFactorBufferInstanceStride(numberOfPatchControlPoints uint, patchStart uint, patchCount uint, patchIndexBuffer unsafe.Pointer, patchIndexBufferOffset uint, instanceCount uint, baseInstance uint, buffer unsafe.Pointer, offset uint, instanceStride uint)
	DrawPrimitivesVertexStartVertexCountInstanceCountBaseInstance(primitiveType PrimitiveType, vertexStart uint, vertexCount uint, instanceCount uint, baseInstance uint)
	Reset()
	SetBarrier()
	SetCullMode(cullMode CullMode)
	SetDepthBiasSlopeScaleClamp(depthBias float32, slopeScale float32, clamp float32)
	SetDepthClipMode(depthClipMode DepthClipMode)
	SetDepthStencilState(depthStencilState unsafe.Pointer)
	SetFragmentBufferOffsetAtIndex(buffer unsafe.Pointer, offset uint, index uint)
	SetFrontFacingWinding(frontFacingWindning Winding)
	SetMeshBufferOffsetAtIndex(buffer unsafe.Pointer, offset uint, index uint)
	SetObjectBufferOffsetAtIndex(buffer unsafe.Pointer, offset uint, index uint)
	SetObjectThreadgroupMemoryLengthAtIndex(length uint, index uint)
	SetRenderPipelineState(pipelineState unsafe.Pointer)
	SetTriangleFillMode(fillMode TriangleFillMode)
	SetVertexBufferOffsetAtIndex(buffer unsafe.Pointer, offset uint, index uint)
	SetVertexBufferOffsetAttributeStrideAtIndex(buffer unsafe.Pointer, offset uint, stride uint, index uint)
}
