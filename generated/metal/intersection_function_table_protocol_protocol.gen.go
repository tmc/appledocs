// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/corefoundation"
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
	SetBufferOffsetAtIndex(buffer unsafe.Pointer, offset uint, index uint)/* debug [protocol_interface/required_method]: SetBufferOffsetAtIndex */
	SetBuffersOffsetsWithRange(buffers []objc.ID, offsets uint, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetBuffersOffsetsWithRange */
	SetFunctionAtIndex(function unsafe.Pointer, index uint)/* debug [protocol_interface/required_method]: SetFunctionAtIndex */
	SetFunctionsWithRange(functions []objc.ID, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetFunctionsWithRange */
	SetOpaqueCurveIntersectionFunctionWithSignatureAtIndex(signature IntersectionFunctionSignature, index uint)/* debug [protocol_interface/required_method]: SetOpaqueCurveIntersectionFunctionWithSignatureAtIndex */
	SetOpaqueCurveIntersectionFunctionWithSignatureWithRange(signature IntersectionFunctionSignature, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetOpaqueCurveIntersectionFunctionWithSignatureWithRange */
	SetOpaqueTriangleIntersectionFunctionWithSignatureAtIndex(signature IntersectionFunctionSignature, index uint)/* debug [protocol_interface/required_method]: SetOpaqueTriangleIntersectionFunctionWithSignatureAtIndex */
	SetOpaqueTriangleIntersectionFunctionWithSignatureWithRange(signature IntersectionFunctionSignature, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetOpaqueTriangleIntersectionFunctionWithSignatureWithRange */
	SetVisibleFunctionTableAtBufferIndex(functionTable unsafe.Pointer, bufferIndex uint)/* debug [protocol_interface/required_method]: SetVisibleFunctionTableAtBufferIndex */
	SetVisibleFunctionTablesWithBufferRange(functionTables []objc.ID, bufferRange corefoundation.Range)/* debug [protocol_interface/required_method]: SetVisibleFunctionTablesWithBufferRange */
}
