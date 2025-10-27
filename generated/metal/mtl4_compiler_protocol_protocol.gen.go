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
	NewDynamicLibraryError(library unsafe.Pointer, error_ foundation.foundation.INSError) unsafe.Pointer
	NewDynamicLibraryWithURLError(url foundation.foundation.INSURL, error_ foundation.foundation.INSError) unsafe.Pointer
	NewLibraryWithDescriptorError(descriptor IMTL4LibraryDescriptor, error_ foundation.foundation.INSError) unsafe.Pointer
	NewBinaryFunctionWithDescriptorCompilerTaskOptionsCompletionHandler(descriptor IMTL4BinaryFunctionDescriptor, compilerTaskOptions IMTL4CompilerTaskOptions, completionHandler objectivec.IObject) unsafe.Pointer
	NewBinaryFunctionWithDescriptorCompilerTaskOptionsError(descriptor IMTL4BinaryFunctionDescriptor, compilerTaskOptions IMTL4CompilerTaskOptions, error_ foundation.foundation.INSError) unsafe.Pointer
	NewComputePipelineStateWithDescriptorCompilerTaskOptionsCompletionHandler(descriptor IMTL4ComputePipelineDescriptor, compilerTaskOptions IMTL4CompilerTaskOptions, completionHandler NewComputePipelineStateCompletionHandler /* not a class type */) unsafe.Pointer
	NewComputePipelineStateWithDescriptorCompilerTaskOptionsError(descriptor IMTL4ComputePipelineDescriptor, compilerTaskOptions IMTL4CompilerTaskOptions, error_ foundation.foundation.INSError) unsafe.Pointer
	NewComputePipelineStateWithDescriptorDynamicLinkingDescriptorCompilerTaskOptionsCompletionHandler(descriptor IMTL4ComputePipelineDescriptor, dynamicLinkingDescriptor IMTL4PipelineStageDynamicLinkingDescriptor, compilerTaskOptions IMTL4CompilerTaskOptions, completionHandler NewComputePipelineStateCompletionHandler /* not a class type */) unsafe.Pointer
	NewComputePipelineStateWithDescriptorDynamicLinkingDescriptorCompilerTaskOptionsError(descriptor IMTL4ComputePipelineDescriptor, dynamicLinkingDescriptor IMTL4PipelineStageDynamicLinkingDescriptor, compilerTaskOptions IMTL4CompilerTaskOptions, error_ foundation.foundation.INSError) unsafe.Pointer
	NewDynamicLibraryCompletionHandler(library unsafe.Pointer, completionHandler NewDynamicLibraryCompletionHandler /* not a class type */) unsafe.Pointer
	NewDynamicLibraryWithURLCompletionHandler(url foundation.foundation.INSURL, completionHandler NewDynamicLibraryCompletionHandler /* not a class type */) unsafe.Pointer
	NewLibraryWithDescriptorCompletionHandler(descriptor IMTL4LibraryDescriptor, completionHandler NewLibraryCompletionHandler /* not a class type */) unsafe.Pointer
	NewMachineLearningPipelineStateWithDescriptorCompletionHandler(descriptor IMTL4MachineLearningPipelineDescriptor, completionHandler objectivec.IObject) unsafe.Pointer
	NewMachineLearningPipelineStateWithDescriptorError(descriptor IMTL4MachineLearningPipelineDescriptor, error_ foundation.foundation.INSError) unsafe.Pointer
	NewRenderPipelineStateBySpecializationWithDescriptorPipelineCompletionHandler(descriptor IMTL4PipelineDescriptor, pipeline unsafe.Pointer, completionHandler NewRenderPipelineStateCompletionHandler /* not a class type */) unsafe.Pointer
	NewRenderPipelineStateBySpecializationWithDescriptorPipelineError(descriptor IMTL4PipelineDescriptor, pipeline unsafe.Pointer, error_ foundation.foundation.INSError) unsafe.Pointer
	NewRenderPipelineStateWithDescriptorCompilerTaskOptionsCompletionHandler(descriptor IMTL4PipelineDescriptor, compilerTaskOptions IMTL4CompilerTaskOptions, completionHandler NewRenderPipelineStateCompletionHandler /* not a class type */) unsafe.Pointer
	NewRenderPipelineStateWithDescriptorCompilerTaskOptionsError(descriptor IMTL4PipelineDescriptor, compilerTaskOptions IMTL4CompilerTaskOptions, error_ foundation.foundation.INSError) unsafe.Pointer
	NewRenderPipelineStateWithDescriptorDynamicLinkingDescriptorCompilerTaskOptionsCompletionHandler(descriptor IMTL4PipelineDescriptor, dynamicLinkingDescriptor IMTL4RenderPipelineDynamicLinkingDescriptor, compilerTaskOptions IMTL4CompilerTaskOptions, completionHandler NewRenderPipelineStateCompletionHandler /* not a class type */) unsafe.Pointer
	NewRenderPipelineStateWithDescriptorDynamicLinkingDescriptorCompilerTaskOptionsError(descriptor IMTL4PipelineDescriptor, dynamicLinkingDescriptor IMTL4RenderPipelineDynamicLinkingDescriptor, compilerTaskOptions IMTL4CompilerTaskOptions, error_ foundation.foundation.INSError) unsafe.Pointer
}
