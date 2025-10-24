// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

/* debug [enums.gen.go]: Generating 10 enums for CoreML */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum MLModelError (11 cases) */
// MLModelError - Information about a Core ML model error.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelError-swift.struct/Code
type MLModelError uint

const (
	// MLModelErrorCustomLayer - An error code for problems related to custom layers.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelError-swift.struct/Code/customLayer
	MLModelErrorCustomLayer MLModelError = 0
	// MLModelErrorCustomModel - An error code for problems related to custom models.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelError-swift.struct/Code/customModel
	MLModelErrorCustomModel MLModelError = 0
	// MLModelErrorFeatureType - An error code for problems related to model features.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelError-swift.struct/Code/featureType
	MLModelErrorFeatureType MLModelError = 0
	// MLModelErrorGeneric - An error code for runtime issues that don’t apply to the other error codes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelError-swift.struct/Code/generic
	MLModelErrorGeneric MLModelError = 0
	// MLModelErrorIO - An error code for problems related to the system’s input or output.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelError-swift.struct/Code/io
	MLModelErrorIO MLModelError = 0
	// MLModelErrorModelCollection - An error code for problems related to retrieving a model collection from the deployment system.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelError-swift.struct/Code/modelCollection
	MLModelErrorModelCollection MLModelError = 0
	// MLModelErrorModelDecryption - An error code for problems related to decrypting models.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelError-swift.struct/Code/modelDecryption
	MLModelErrorModelDecryption MLModelError = 0
	// MLModelErrorModelDecryptionKeyFetch - An error code for problems related to retrieving a model’s decryption key.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelError-swift.struct/Code/modelDecryptionKeyFetch
	MLModelErrorModelDecryptionKeyFetch MLModelError = 0
	// MLModelErrorParameters - An error code for problems related to model parameters.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelError-swift.struct/Code/parameters
	MLModelErrorParameters MLModelError = 0
	// MLModelErrorPredictionCancelled - An error code for problems related to canceling the prediction before it completes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelError-swift.struct/Code/predictionCancelled
	MLModelErrorPredictionCancelled MLModelError = 0
	// MLModelErrorUpdate - An error code for problems related to on-device model updates.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelError-swift.struct/Code/update
	MLModelErrorUpdate MLModelError = 0
)

/* debug [enums.gen.go]: Processing enum MLComputeUnits (4 cases) */
// MLComputeUnits - The set of processing-unit configurations the model can use to make predictions.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLComputeUnits
type MLComputeUnits uint

const (
	// MLComputeUnitsAll - The option you choose to allow the model to use all compute units available, including the neural engine.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLComputeUnits/all
	MLComputeUnitsAll MLComputeUnits = 0
	// MLComputeUnitsCPUAndGPU - The option you choose to allow the model to use both the CPU and GPU, but not the neural engine.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLComputeUnits/cpuAndGPU
	MLComputeUnitsCPUAndGPU MLComputeUnits = 0
	// MLComputeUnitsCPUAndNeuralEngine - The option you choose to allow the model to use both the CPU and neural engine, but not the GPU.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLComputeUnits/cpuAndNeuralEngine
	MLComputeUnitsCPUAndNeuralEngine MLComputeUnits = 0
	// MLComputeUnitsCPUOnly - The option you choose to limit the model to only use the CPU.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLComputeUnits/cpuOnly
	MLComputeUnitsCPUOnly MLComputeUnits = 0
)

/* debug [enums.gen.go]: Processing enum MLFeatureType (9 cases) */
// MLFeatureType - The possible types for feature values, input features, and output features.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureType
type MLFeatureType uint

const (
	// MLFeatureTypeDictionary - The type for dictionary features and feature values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureType/dictionary
	MLFeatureTypeDictionary MLFeatureType = 0
	// MLFeatureTypeDouble - The type for double features and feature values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureType/double
	MLFeatureTypeDouble MLFeatureType = 0
	// MLFeatureTypeImage - The type for image features and feature values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureType/image
	MLFeatureTypeImage MLFeatureType = 0
	// MLFeatureTypeInt64 - The type for integer features and feature values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureType/int64
	MLFeatureTypeInt64 MLFeatureType = 0
	// MLFeatureTypeInvalid - The type for invalid feature values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureType/invalid
	MLFeatureTypeInvalid MLFeatureType = 0
	// MLFeatureTypeMultiArray - The type for multidimensional array features and feature values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureType/multiArray
	MLFeatureTypeMultiArray MLFeatureType = 0
	// MLFeatureTypeSequence - The type for sequence features and feature values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureType/sequence
	MLFeatureTypeSequence MLFeatureType = 0
	// MLFeatureTypeState - MLState. Represents a model state that may be updated in each inference.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureType/state
	MLFeatureTypeState MLFeatureType = 0
	// MLFeatureTypeString - The type for string features and feature values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureType/string
	MLFeatureTypeString MLFeatureType = 0
)

/* debug [enums.gen.go]: Processing enum MLImageSizeConstraintType (3 cases) */
// MLImageSizeConstraintType - The modes that determine how the model defines a feature’s image size constraint.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLImageSizeConstraintType
type MLImageSizeConstraintType uint

const (
	// MLImageSizeConstraintTypeEnumerated - The image feature accepts image sizes listed in an array.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLImageSizeConstraintType/enumerated
	MLImageSizeConstraintTypeEnumerated MLImageSizeConstraintType = 0
	// MLImageSizeConstraintTypeRange - The image feature accepts image sizes defined by a range of widths and a range of heights.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLImageSizeConstraintType/range
	MLImageSizeConstraintTypeRange MLImageSizeConstraintType = 0
	// MLImageSizeConstraintTypeUnspecified - The image size constraint is not configured and should be ignored.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLImageSizeConstraintType/unspecified
	MLImageSizeConstraintTypeUnspecified MLImageSizeConstraintType = 0
)

/* debug [enums.gen.go]: Processing enum MLMultiArrayDataType (7 cases) */
// MLMultiArrayDataType - Constants that define the underlying element types a multiarray can store.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArrayDataType
type MLMultiArrayDataType uint

const (
	// MLMultiArrayDataTypeDouble - Designates the multiarray’s elements as doubles.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArrayDataType/double
	MLMultiArrayDataTypeDouble MLMultiArrayDataType = 0
	// MLMultiArrayDataTypeFloat - Designates the multiarray’s elements as floats.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArrayDataType/float
	MLMultiArrayDataTypeFloat MLMultiArrayDataType = 0
	// MLMultiArrayDataTypeFloat16 - Designates the multiarray’s elements as 16-bit floats.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArrayDataType/float16
	MLMultiArrayDataTypeFloat16 MLMultiArrayDataType = 0
	// MLMultiArrayDataTypeFloat32 - Designates the multiarray’s elements as 32-bit floats.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArrayDataType/float32
	MLMultiArrayDataTypeFloat32 MLMultiArrayDataType = 0
	// MLMultiArrayDataTypeFloat64 - Designates the multiarray’s elements as 64-bit floats.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArrayDataType/float64
	MLMultiArrayDataTypeFloat64 MLMultiArrayDataType = 0
	// MLMultiArrayDataTypeInt32 - Designates the multiarray’s elements as 32-bit integers.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArrayDataType/int32
	MLMultiArrayDataTypeInt32 MLMultiArrayDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArrayDataType/int8
	MLMultiArrayDataTypeInt8 MLMultiArrayDataType = 0
)

/* debug [enums.gen.go]: Processing enum MLMultiArrayShapeConstraintType (3 cases) */
// MLMultiArrayShapeConstraintType - The possible types of shape constraints.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArrayShapeConstraintType
type MLMultiArrayShapeConstraintType uint

const (
	// MLMultiArrayShapeConstraintTypeEnumerated - The constraint is an array of allowed shapes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArrayShapeConstraintType/enumerated
	MLMultiArrayShapeConstraintTypeEnumerated MLMultiArrayShapeConstraintType = 0
	// MLMultiArrayShapeConstraintTypeRange - The constraint is a set of ranges allowed for the array shape.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArrayShapeConstraintType/range
	MLMultiArrayShapeConstraintTypeRange MLMultiArrayShapeConstraintType = 0
	// MLMultiArrayShapeConstraintTypeUnspecified - The constraint type is undefined.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArrayShapeConstraintType/unspecified
	MLMultiArrayShapeConstraintTypeUnspecified MLMultiArrayShapeConstraintType = 0
)

/* debug [enums.gen.go]: Processing enum MLReshapeFrequencyHint (2 cases) */
// MLReshapeFrequencyHint enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLReshapeFrequencyHint
type MLReshapeFrequencyHint int

const (
	// MLReshapeFrequencyHintFrequent - The input shape is expected to change frequently on each prediction sent to this loaded model instance. Core ML will try to minimize the latency associated with shape changes and avoid expensive shape-specific optimizations prior to prediction computation. While prediction computation may be slower for each specific shape, switching between shapes should be faster.   This is the default.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLReshapeFrequencyHint/MLReshapeFrequencyHintFrequent
	MLReshapeFrequencyHintFrequent MLReshapeFrequencyHint = 0
	// MLReshapeFrequencyHintInfrequent - The input shape is expected to be stable and many/all predictions sent to this loaded model instance would use the same input shapes repeatedly. On the shape change, Core ML re-optimizes the internal engine for the new shape if possible. The re-optimization takes some time, but the subsequent predictions for the shape should run faster.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLReshapeFrequencyHint/MLReshapeFrequencyHintInfrequent
	MLReshapeFrequencyHintInfrequent MLReshapeFrequencyHint = 0
)

/* debug [enums.gen.go]: Processing enum MLSpecializationStrategy (2 cases) */
// MLSpecializationStrategy - The optimization strategy for the model specialization.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLSpecializationStrategy
type MLSpecializationStrategy int

const (
	// MLSpecializationStrategyDefault - The strategy that works well for most applications.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLSpecializationStrategy/MLSpecializationStrategyDefault
	MLSpecializationStrategyDefault MLSpecializationStrategy = 0
	// MLSpecializationStrategyFastPrediction - Prefer the prediction latency at the potential cost of specialization time, memory footprint, and the disk space usage of specialized artifacts.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLSpecializationStrategy/MLSpecializationStrategyFastPrediction
	MLSpecializationStrategyFastPrediction MLSpecializationStrategy = 0
)

/* debug [enums.gen.go]: Processing enum MLTaskState (5 cases) */
// MLTaskState - The state of a machine learning task.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLTaskState
type MLTaskState uint

const (
	// MLTaskStateCancelling - The state of a machine learning task that’s in mid-termination, before it could finish successfully.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLTaskState/cancelling
	MLTaskStateCancelling MLTaskState = 0
	// MLTaskStateCompleted - The state of a machine learning task that has finished successfully.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLTaskState/completed
	MLTaskStateCompleted MLTaskState = 0
	// MLTaskStateFailed - The state of a machine learning task that has terminated due to an error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLTaskState/failed
	MLTaskStateFailed MLTaskState = 0
	// MLTaskStateRunning - The state of a machine learning task that’s executing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLTaskState/running
	MLTaskStateRunning MLTaskState = 0
	// MLTaskStateSuspended - The state of a machine learning task that’s paused.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLTaskState/suspended
	MLTaskStateSuspended MLTaskState = 0
)

/* debug [enums.gen.go]: Processing enum MLUpdateProgressEvent (3 cases) */
// MLUpdateProgressEvent - A type of event during a model update task.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateProgressEvent
type MLUpdateProgressEvent uint

const (
	// MLUpdateProgressEventEpochEnd - An event that represents the end of training epoch.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateProgressEvent/epochEnd
	MLUpdateProgressEventEpochEnd MLUpdateProgressEvent = 0
	// MLUpdateProgressEventMiniBatchEnd - An event that represents the end of a mini-batch within a training epoch.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateProgressEvent/miniBatchEnd
	MLUpdateProgressEventMiniBatchEnd MLUpdateProgressEvent = 0
	// MLUpdateProgressEventTrainingBegin - An event that represents the start of training.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateProgressEvent/trainingBegin
	MLUpdateProgressEventTrainingBegin MLUpdateProgressEvent = 0
)


