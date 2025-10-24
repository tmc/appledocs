// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/corefoundation"
)

// PVisibleFunctionTable is the MTLVisibleFunctionTable protocol interface.
//
// A table of shader functions visible to your app that you can pass into compute commands to customize the behavior of a shader.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.metal/documentation/Metal/MTLVisibleFunctionTable
type PVisibleFunctionTable interface {
	// Required methods
	SetFunctionAtIndex(function unsafe.Pointer, index uint)/* debug [protocol_interface/required_method]: SetFunctionAtIndex */
	SetFunctionsWithRange(functions []objc.ID, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetFunctionsWithRange */
}

