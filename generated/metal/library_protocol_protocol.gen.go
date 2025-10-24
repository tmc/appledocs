// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
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
	NewFunctionWithDescriptorError(descriptor IMTLFunctionDescriptor, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewFunctionWithDescriptorError */
	NewFunctionWithDescriptorCompletionHandler(descriptor IMTLFunctionDescriptor, completionHandler unsafe.Pointer)/* debug [protocol_interface/required_method]: NewFunctionWithDescriptorCompletionHandler */
	NewFunctionWithName(functionName objc.IObject /* cross-framework: NSString */) unsafe.Pointer/* debug [protocol_interface/required_method]: NewFunctionWithName */
	NewFunctionWithNameConstantValuesError(name objc.IObject /* cross-framework: NSString */, constantValues IMTLFunctionConstantValues, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewFunctionWithNameConstantValuesError */
	NewFunctionWithNameConstantValuesCompletionHandler(name objc.IObject /* cross-framework: NSString */, constantValues IMTLFunctionConstantValues, completionHandler unsafe.Pointer)/* debug [protocol_interface/required_method]: NewFunctionWithNameConstantValuesCompletionHandler */
	NewIntersectionFunctionWithDescriptorError(descriptor IMTLIntersectionFunctionDescriptor, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewIntersectionFunctionWithDescriptorError */
	NewIntersectionFunctionWithDescriptorCompletionHandler(descriptor IMTLIntersectionFunctionDescriptor, completionHandler unsafe.Pointer)/* debug [protocol_interface/required_method]: NewIntersectionFunctionWithDescriptorCompletionHandler */
	ReflectionForFunctionWithName(functionName objc.IObject /* cross-framework: NSString */) FunctionReflection/* debug [protocol_interface/required_method]: ReflectionForFunctionWithName */
}
