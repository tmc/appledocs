// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

// Enum types and constants
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

// MLMultiArrayDataType - Constants that define the underlying element types a multiarray can store.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArrayDataType
type MLMultiArrayDataType uint

const (
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
)

// MLReshapeFrequencyHint enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLReshapeFrequencyHint
type MLReshapeFrequencyHint int

// MLSpecializationStrategy - The optimization strategy for the model specialization.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLSpecializationStrategy
type MLSpecializationStrategy int

// MLTaskState - The state of a machine learning task.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLTaskState
type MLTaskState uint


