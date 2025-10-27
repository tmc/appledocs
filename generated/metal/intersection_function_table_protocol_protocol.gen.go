// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
)

// PIntersectionFunctionTable is the MTLIntersectionFunctionTable protocol interface.
//
// A table of intersection functions that Metal calls to perform ray-tracing intersection tests.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.metal/documentation/Metal/MTLIntersectionFunctionTable
type PIntersectionFunctionTable interface {
	// Required methods
	SetBufferOffsetAtIndex(buffer unsafe.Pointer, offset uint, index uint)
	SetBuffersOffsetsWithRange(buffers []objc.ID, offsets uint, range_ foundation.Range)
	SetFunctionAtIndex(function unsafe.Pointer, index uint)
	SetFunctionsWithRange(functions []objc.ID, range_ foundation.Range)
	SetOpaqueCurveIntersectionFunctionWithSignatureAtIndex(signature IntersectionFunctionSignature, index uint)
	SetOpaqueCurveIntersectionFunctionWithSignatureWithRange(signature IntersectionFunctionSignature, range_ foundation.Range)
	SetOpaqueTriangleIntersectionFunctionWithSignatureAtIndex(signature IntersectionFunctionSignature, index uint)
	SetOpaqueTriangleIntersectionFunctionWithSignatureWithRange(signature IntersectionFunctionSignature, range_ foundation.Range)
	SetVisibleFunctionTableAtBufferIndex(functionTable unsafe.Pointer, bufferIndex uint)
	SetVisibleFunctionTablesWithBufferRange(functionTables []objc.ID, bufferRange foundation.Range)
}
