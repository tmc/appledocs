// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PMTL4Compiler is the MTL4Compiler protocol interface.
//
// A abstraction for a pipeline state and shader function compiler.
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//
// See: doc://com.apple.metal/documentation/Metal/MTL4Compiler
type PMTL4Compiler interface {
	// Required methods
	NewDynamicLibraryError(library unsafe.Pointer, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewDynamicLibraryError */
	NewDynamicLibraryWithURLError(url objc.IObject /* cross-framework: NSURL */, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewDynamicLibraryWithURLError */
	NewLibraryWithDescriptorError(descriptor IMTL4LibraryDescriptor, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewLibraryWithDescriptorError */
	NewBinaryFunctionWithDescriptorCompilerTaskOptionsCompletionHandler(descriptor IMTL4BinaryFunctionDescriptor, compilerTaskOptions IMTL4CompilerTaskOptions, completionHandler objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewBinaryFunctionWithDescriptorCompilerTaskOptionsCompletionHandler */
	NewBinaryFunctionWithDescriptorCompilerTaskOptionsError(descriptor IMTL4BinaryFunctionDescriptor, compilerTaskOptions IMTL4CompilerTaskOptions, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewBinaryFunctionWithDescriptorCompilerTaskOptionsError */
	NewComputePipelineStateWithDescriptorCompilerTaskOptionsCompletionHandler(descriptor IMTL4ComputePipelineDescriptor, compilerTaskOptions IMTL4CompilerTaskOptions, completionHandler NewComputePipelineStateCompletionHandler /* not a class type */) unsafe.Pointer/* debug [protocol_interface/required_method]: NewComputePipelineStateWithDescriptorCompilerTaskOptionsCompletionHandler */
	NewComputePipelineStateWithDescriptorCompilerTaskOptionsError(descriptor IMTL4ComputePipelineDescriptor, compilerTaskOptions IMTL4CompilerTaskOptions, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewComputePipelineStateWithDescriptorCompilerTaskOptionsError */
	NewComputePipelineStateWithDescriptorDynamicLinkingDescriptorCompilerTaskOptionsCompletionHandler(descriptor IMTL4ComputePipelineDescriptor, dynamicLinkingDescriptor IMTL4PipelineStageDynamicLinkingDescriptor, compilerTaskOptions IMTL4CompilerTaskOptions, completionHandler NewComputePipelineStateCompletionHandler /* not a class type */) unsafe.Pointer/* debug [protocol_interface/required_method]: NewComputePipelineStateWithDescriptorDynamicLinkingDescriptorCompilerTaskOptionsCompletionHandler */
	NewComputePipelineStateWithDescriptorDynamicLinkingDescriptorCompilerTaskOptionsError(descriptor IMTL4ComputePipelineDescriptor, dynamicLinkingDescriptor IMTL4PipelineStageDynamicLinkingDescriptor, compilerTaskOptions IMTL4CompilerTaskOptions, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewComputePipelineStateWithDescriptorDynamicLinkingDescriptorCompilerTaskOptionsError */
	NewDynamicLibraryCompletionHandler(library unsafe.Pointer, completionHandler NewDynamicLibraryCompletionHandler /* not a class type */) unsafe.Pointer/* debug [protocol_interface/required_method]: NewDynamicLibraryCompletionHandler */
	NewDynamicLibraryWithURLCompletionHandler(url objc.IObject /* cross-framework: NSURL */, completionHandler NewDynamicLibraryCompletionHandler /* not a class type */) unsafe.Pointer/* debug [protocol_interface/required_method]: NewDynamicLibraryWithURLCompletionHandler */
	NewLibraryWithDescriptorCompletionHandler(descriptor IMTL4LibraryDescriptor, completionHandler NewLibraryCompletionHandler /* not a class type */) unsafe.Pointer/* debug [protocol_interface/required_method]: NewLibraryWithDescriptorCompletionHandler */
	NewMachineLearningPipelineStateWithDescriptorCompletionHandler(descriptor IMTL4MachineLearningPipelineDescriptor, completionHandler objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewMachineLearningPipelineStateWithDescriptorCompletionHandler */
	NewMachineLearningPipelineStateWithDescriptorError(descriptor IMTL4MachineLearningPipelineDescriptor, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewMachineLearningPipelineStateWithDescriptorError */
	NewRenderPipelineStateBySpecializationWithDescriptorPipelineCompletionHandler(descriptor IMTL4PipelineDescriptor, pipeline unsafe.Pointer, completionHandler NewRenderPipelineStateCompletionHandler /* not a class type */) unsafe.Pointer/* debug [protocol_interface/required_method]: NewRenderPipelineStateBySpecializationWithDescriptorPipelineCompletionHandler */
	NewRenderPipelineStateBySpecializationWithDescriptorPipelineError(descriptor IMTL4PipelineDescriptor, pipeline unsafe.Pointer, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewRenderPipelineStateBySpecializationWithDescriptorPipelineError */
	NewRenderPipelineStateWithDescriptorCompilerTaskOptionsCompletionHandler(descriptor IMTL4PipelineDescriptor, compilerTaskOptions IMTL4CompilerTaskOptions, completionHandler NewRenderPipelineStateCompletionHandler /* not a class type */) unsafe.Pointer/* debug [protocol_interface/required_method]: NewRenderPipelineStateWithDescriptorCompilerTaskOptionsCompletionHandler */
	NewRenderPipelineStateWithDescriptorCompilerTaskOptionsError(descriptor IMTL4PipelineDescriptor, compilerTaskOptions IMTL4CompilerTaskOptions, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewRenderPipelineStateWithDescriptorCompilerTaskOptionsError */
	NewRenderPipelineStateWithDescriptorDynamicLinkingDescriptorCompilerTaskOptionsCompletionHandler(descriptor IMTL4PipelineDescriptor, dynamicLinkingDescriptor IMTL4RenderPipelineDynamicLinkingDescriptor, compilerTaskOptions IMTL4CompilerTaskOptions, completionHandler NewRenderPipelineStateCompletionHandler /* not a class type */) unsafe.Pointer/* debug [protocol_interface/required_method]: NewRenderPipelineStateWithDescriptorDynamicLinkingDescriptorCompilerTaskOptionsCompletionHandler */
	NewRenderPipelineStateWithDescriptorDynamicLinkingDescriptorCompilerTaskOptionsError(descriptor IMTL4PipelineDescriptor, dynamicLinkingDescriptor IMTL4RenderPipelineDynamicLinkingDescriptor, compilerTaskOptions IMTL4CompilerTaskOptions, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewRenderPipelineStateWithDescriptorDynamicLinkingDescriptorCompilerTaskOptionsError */
}
