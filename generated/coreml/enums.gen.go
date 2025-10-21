// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

// Enum types and constants
// MLComputeUnits - The set of processing-unit configurations the model can use to make predictions.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLComputeUnits
type ComputeUnits uint

const (
	// ComputeUnitsAll - The option you choose to allow the model to use all compute units available, including the neural engine.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLComputeUnits/all
	ComputeUnitsAll ComputeUnits = 0
	// ComputeUnitsCPUAndGPU - The option you choose to allow the model to use both the CPU and GPU, but not the neural engine.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLComputeUnits/cpuAndGPU
	ComputeUnitsCPUAndGPU ComputeUnits = 0
	// ComputeUnitsCPUAndNeuralEngine - The option you choose to allow the model to use both the CPU and neural engine, but not the GPU.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLComputeUnits/cpuAndNeuralEngine
	ComputeUnitsCPUAndNeuralEngine ComputeUnits = 0
	// ComputeUnitsCPUOnly - The option you choose to limit the model to only use the CPU.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLComputeUnits/cpuOnly
	ComputeUnitsCPUOnly ComputeUnits = 0
)

// MLFeatureType - The possible types for feature values, input features, and output features.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureType
type FeatureType uint

const (
	// FeatureTypeDictionary - The type for dictionary features and feature values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureType/dictionary
	FeatureTypeDictionary FeatureType = 0
	// FeatureTypeDouble - The type for double features and feature values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureType/double
	FeatureTypeDouble FeatureType = 0
	// FeatureTypeImage - The type for image features and feature values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureType/image
	FeatureTypeImage FeatureType = 0
	// FeatureTypeInt64 - The type for integer features and feature values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureType/int64
	FeatureTypeInt64 FeatureType = 0
	// FeatureTypeInvalid - The type for invalid feature values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureType/invalid
	FeatureTypeInvalid FeatureType = 0
	// FeatureTypeMultiArray - The type for multidimensional array features and feature values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureType/multiArray
	FeatureTypeMultiArray FeatureType = 0
	// FeatureTypeSequence - The type for sequence features and feature values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureType/sequence
	FeatureTypeSequence FeatureType = 0
	// FeatureTypeState - MLState. Represents a model state that may be updated in each inference.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureType/state
	FeatureTypeState FeatureType = 0
	// FeatureTypeString - The type for string features and feature values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureType/string
	FeatureTypeString FeatureType = 0
)

// MLImageSizeConstraintType - The modes that determine how the model defines a feature’s image size constraint.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLImageSizeConstraintType
type ImageSizeConstraintType uint

const (
	// ImageSizeConstraintTypeEnumerated - The image feature accepts image sizes listed in an array.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLImageSizeConstraintType/enumerated
	ImageSizeConstraintTypeEnumerated ImageSizeConstraintType = 0
	// ImageSizeConstraintTypeRange - The image feature accepts image sizes defined by a range of widths and a range of heights.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLImageSizeConstraintType/range
	ImageSizeConstraintTypeRange ImageSizeConstraintType = 0
	// ImageSizeConstraintTypeUnspecified - The image size constraint is not configured and should be ignored.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLImageSizeConstraintType/unspecified
	ImageSizeConstraintTypeUnspecified ImageSizeConstraintType = 0
)

// MLModelError - Information about a Core ML model error.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelError-swift.struct/Code
type ModelError uint

const (
	// ModelErrorCustomLayer - An error code for problems related to custom layers.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelError-swift.struct/Code/customLayer
	ModelErrorCustomLayer ModelError = 0
	// ModelErrorCustomModel - An error code for problems related to custom models.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelError-swift.struct/Code/customModel
	ModelErrorCustomModel ModelError = 0
	// ModelErrorFeatureType - An error code for problems related to model features.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelError-swift.struct/Code/featureType
	ModelErrorFeatureType ModelError = 0
	// ModelErrorGeneric - An error code for runtime issues that don’t apply to the other error codes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelError-swift.struct/Code/generic
	ModelErrorGeneric ModelError = 0
	// ModelErrorIO - An error code for problems related to the system’s input or output.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelError-swift.struct/Code/io
	ModelErrorIO ModelError = 0
	// ModelErrorModelCollection - An error code for problems related to retrieving a model collection from the deployment system.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelError-swift.struct/Code/modelCollection
	ModelErrorModelCollection ModelError = 0
	// ModelErrorModelDecryption - An error code for problems related to decrypting models.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelError-swift.struct/Code/modelDecryption
	ModelErrorModelDecryption ModelError = 0
	// ModelErrorModelDecryptionKeyFetch - An error code for problems related to retrieving a model’s decryption key.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelError-swift.struct/Code/modelDecryptionKeyFetch
	ModelErrorModelDecryptionKeyFetch ModelError = 0
	// ModelErrorParameters - An error code for problems related to model parameters.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelError-swift.struct/Code/parameters
	ModelErrorParameters ModelError = 0
	// ModelErrorPredictionCancelled - An error code for problems related to canceling the prediction before it completes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelError-swift.struct/Code/predictionCancelled
	ModelErrorPredictionCancelled ModelError = 0
	// ModelErrorUpdate - An error code for problems related to on-device model updates.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelError-swift.struct/Code/update
	ModelErrorUpdate ModelError = 0
)

// MLMultiArrayDataType - Constants that define the underlying element types a multiarray can store.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArrayDataType
type MultiArrayDataType uint

const (
	// MultiArrayDataTypeDouble - Designates the multiarray’s elements as doubles.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArrayDataType/double
	MultiArrayDataTypeDouble MultiArrayDataType = 0
	// MultiArrayDataTypeFloat - Designates the multiarray’s elements as floats.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArrayDataType/float
	MultiArrayDataTypeFloat MultiArrayDataType = 0
	// MultiArrayDataTypeFloat16 - Designates the multiarray’s elements as 16-bit floats.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArrayDataType/float16
	MultiArrayDataTypeFloat16 MultiArrayDataType = 0
	// MultiArrayDataTypeFloat32 - Designates the multiarray’s elements as 32-bit floats.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArrayDataType/float32
	MultiArrayDataTypeFloat32 MultiArrayDataType = 0
	// MultiArrayDataTypeFloat64 - Designates the multiarray’s elements as 64-bit floats.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArrayDataType/float64
	MultiArrayDataTypeFloat64 MultiArrayDataType = 0
	// MultiArrayDataTypeInt32 - Designates the multiarray’s elements as 32-bit integers.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArrayDataType/int32
	MultiArrayDataTypeInt32 MultiArrayDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArrayDataType/int8
	MultiArrayDataTypeInt8 MultiArrayDataType = 0
)

// MLReshapeFrequencyHint enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLReshapeFrequencyHint
type ReshapeFrequencyHint uint

const (
	// ReshapeFrequencyHintFrequent - The input shape is expected to change frequently on each prediction sent to this loaded model instance. Core ML will try to minimize the latency associated with shape changes and avoid expensive shape-specific optimizations prior to prediction computation. While prediction computation may be slower for each specific shape, switching between shapes should be faster.   This is the default.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLReshapeFrequencyHint/MLReshapeFrequencyHintFrequent
	ReshapeFrequencyHintFrequent ReshapeFrequencyHint = 0
	// ReshapeFrequencyHintInfrequent - The input shape is expected to be stable and many/all predictions sent to this loaded model instance would use the same input shapes repeatedly. On the shape change, Core ML re-optimizes the internal engine for the new shape if possible. The re-optimization takes some time, but the subsequent predictions for the shape should run faster.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLReshapeFrequencyHint/MLReshapeFrequencyHintInfrequent
	ReshapeFrequencyHintInfrequent ReshapeFrequencyHint = 0
)

// MLSpecializationStrategy - The optimization strategy for the model specialization.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLSpecializationStrategy
type SpecializationStrategy uint

// MLTaskState - The state of a machine learning task.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLTaskState
type TaskState uint

const (
	// TaskStateCancelling - The state of a machine learning task that’s in mid-termination, before it could finish successfully.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLTaskState/cancelling
	TaskStateCancelling TaskState = 0
	// TaskStateCompleted - The state of a machine learning task that has finished successfully.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLTaskState/completed
	TaskStateCompleted TaskState = 0
	// TaskStateFailed - The state of a machine learning task that has terminated due to an error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLTaskState/failed
	TaskStateFailed TaskState = 0
	// TaskStateRunning - The state of a machine learning task that’s executing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLTaskState/running
	TaskStateRunning TaskState = 0
	// TaskStateSuspended - The state of a machine learning task that’s paused.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLTaskState/suspended
	TaskStateSuspended TaskState = 0
)

// MLUpdateProgressEvent - A type of event during a model update task.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateProgressEvent
type UpdateProgressEvent uint


