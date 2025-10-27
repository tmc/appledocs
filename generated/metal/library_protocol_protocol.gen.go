// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
)

// PLibrary is the MTLLibrary protocol interface.
//
// A collection of Metal shader functions.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.11+
//   - tvOS +
//   - visionOS 1.0+
//
// See: doc://com.apple.metal/documentation/Metal/MTLLibrary
type PLibrary interface {
	// Required methods
	NewFunctionWithDescriptorError(descriptor IMTLFunctionDescriptor, error_ foundation.foundation.INSError) unsafe.Pointer
	NewFunctionWithDescriptorCompletionHandler(descriptor IMTLFunctionDescriptor, completionHandler unsafe.Pointer)
	NewFunctionWithName(functionName foundation.foundation.INSString) unsafe.Pointer
	NewFunctionWithNameConstantValuesError(name foundation.foundation.INSString, constantValues IMTLFunctionConstantValues, error_ foundation.foundation.INSError) unsafe.Pointer
	NewFunctionWithNameConstantValuesCompletionHandler(name foundation.foundation.INSString, constantValues IMTLFunctionConstantValues, completionHandler unsafe.Pointer)
	NewIntersectionFunctionWithDescriptorError(descriptor IMTLIntersectionFunctionDescriptor, error_ foundation.foundation.INSError) unsafe.Pointer
	NewIntersectionFunctionWithDescriptorCompletionHandler(descriptor IMTLIntersectionFunctionDescriptor, completionHandler unsafe.Pointer)
	ReflectionForFunctionWithName(functionName foundation.foundation.INSString) IFunctionReflection
}
