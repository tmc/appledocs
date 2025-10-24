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
	ClearBarrier()/* debug [protocol_interface/required_method]: ClearBarrier */
	DrawIndexedPatchesPatchStartPatchCountPatchIndexBufferPatchIndexBufferOffsetControlPointIndexBufferControlPointIndexBufferOffsetInstanceCountBaseInstanceTessellationFactorBufferTessellationFactorBufferOffsetTessellationFactorBufferInstanceStride(numberOfPatchControlPoints uint, patchStart uint, patchCount uint, patchIndexBuffer unsafe.Pointer, patchIndexBufferOffset uint, controlPointIndexBuffer unsafe.Pointer, controlPointIndexBufferOffset uint, instanceCount uint, baseInstance uint, buffer unsafe.Pointer, offset uint, instanceStride uint)/* debug [protocol_interface/required_method]: DrawIndexedPatchesPatchStartPatchCountPatchIndexBufferPatchIndexBufferOffsetControlPointIndexBufferControlPointIndexBufferOffsetInstanceCountBaseInstanceTessellationFactorBufferTessellationFactorBufferOffsetTessellationFactorBufferInstanceStride */
	DrawIndexedPrimitivesIndexCountIndexTypeIndexBufferIndexBufferOffsetInstanceCountBaseVertexBaseInstance(primitiveType PrimitiveType, indexCount uint, indexType IndexType, indexBuffer unsafe.Pointer, indexBufferOffset uint, instanceCount uint, baseVertex int, baseInstance uint)/* debug [protocol_interface/required_method]: DrawIndexedPrimitivesIndexCountIndexTypeIndexBufferIndexBufferOffsetInstanceCountBaseVertexBaseInstance */
	DrawMeshThreadgroupsThreadsPerObjectThreadgroupThreadsPerMeshThreadgroup(threadgroupsPerGrid objc.IObject /* cross-framework: MTLSize */, threadsPerObjectThreadgroup objc.IObject /* cross-framework: MTLSize */, threadsPerMeshThreadgroup objc.IObject /* cross-framework: MTLSize */)/* debug [protocol_interface/required_method]: DrawMeshThreadgroupsThreadsPerObjectThreadgroupThreadsPerMeshThreadgroup */
	DrawMeshThreadsThreadsPerObjectThreadgroupThreadsPerMeshThreadgroup(threadsPerGrid objc.IObject /* cross-framework: MTLSize */, threadsPerObjectThreadgroup objc.IObject /* cross-framework: MTLSize */, threadsPerMeshThreadgroup objc.IObject /* cross-framework: MTLSize */)/* debug [protocol_interface/required_method]: DrawMeshThreadsThreadsPerObjectThreadgroupThreadsPerMeshThreadgroup */
	DrawPatchesPatchStartPatchCountPatchIndexBufferPatchIndexBufferOffsetInstanceCountBaseInstanceTessellationFactorBufferTessellationFactorBufferOffsetTessellationFactorBufferInstanceStride(numberOfPatchControlPoints uint, patchStart uint, patchCount uint, patchIndexBuffer unsafe.Pointer, patchIndexBufferOffset uint, instanceCount uint, baseInstance uint, buffer unsafe.Pointer, offset uint, instanceStride uint)/* debug [protocol_interface/required_method]: DrawPatchesPatchStartPatchCountPatchIndexBufferPatchIndexBufferOffsetInstanceCountBaseInstanceTessellationFactorBufferTessellationFactorBufferOffsetTessellationFactorBufferInstanceStride */
	DrawPrimitivesVertexStartVertexCountInstanceCountBaseInstance(primitiveType PrimitiveType, vertexStart uint, vertexCount uint, instanceCount uint, baseInstance uint)/* debug [protocol_interface/required_method]: DrawPrimitivesVertexStartVertexCountInstanceCountBaseInstance */
	Reset()/* debug [protocol_interface/required_method]: Reset */
	SetBarrier()/* debug [protocol_interface/required_method]: SetBarrier */
	SetCullMode(cullMode CullMode)/* debug [protocol_interface/required_method]: SetCullMode */
	SetDepthBiasSlopeScaleClamp(depthBias float32, slopeScale float32, clamp float32)/* debug [protocol_interface/required_method]: SetDepthBiasSlopeScaleClamp */
	SetDepthClipMode(depthClipMode DepthClipMode)/* debug [protocol_interface/required_method]: SetDepthClipMode */
	SetDepthStencilState(depthStencilState unsafe.Pointer)/* debug [protocol_interface/required_method]: SetDepthStencilState */
	SetFragmentBufferOffsetAtIndex(buffer unsafe.Pointer, offset uint, index uint)/* debug [protocol_interface/required_method]: SetFragmentBufferOffsetAtIndex */
	SetFrontFacingWinding(frontFacingWindning Winding)/* debug [protocol_interface/required_method]: SetFrontFacingWinding */
	SetMeshBufferOffsetAtIndex(buffer unsafe.Pointer, offset uint, index uint)/* debug [protocol_interface/required_method]: SetMeshBufferOffsetAtIndex */
	SetObjectBufferOffsetAtIndex(buffer unsafe.Pointer, offset uint, index uint)/* debug [protocol_interface/required_method]: SetObjectBufferOffsetAtIndex */
	SetObjectThreadgroupMemoryLengthAtIndex(length uint, index uint)/* debug [protocol_interface/required_method]: SetObjectThreadgroupMemoryLengthAtIndex */
	SetRenderPipelineState(pipelineState unsafe.Pointer)/* debug [protocol_interface/required_method]: SetRenderPipelineState */
	SetTriangleFillMode(fillMode TriangleFillMode)/* debug [protocol_interface/required_method]: SetTriangleFillMode */
	SetVertexBufferOffsetAtIndex(buffer unsafe.Pointer, offset uint, index uint)/* debug [protocol_interface/required_method]: SetVertexBufferOffsetAtIndex */
	SetVertexBufferOffsetAttributeStrideAtIndex(buffer unsafe.Pointer, offset uint, stride uint, index uint)/* debug [protocol_interface/required_method]: SetVertexBufferOffsetAttributeStrideAtIndex */
}
