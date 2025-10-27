// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PCustomModel is the MLCustomModel protocol interface.
//
// An interface that defines the behavior of a custom model.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+
//
// See: doc://com.apple.coreml/documentation/CoreML/MLCustomModel
type PCustomModel interface {
	// Required methods
	InitWithModelDescriptionParameterDictionaryError(modelDescription IMLModelDescription, parameters foundation.IDictionary, error_ foundation.foundation.INSError) objectivec.IObject
	PredictionFromFeaturesOptionsError(input unsafe.Pointer, options IMLPredictionOptions, error_ foundation.foundation.INSError) unsafe.Pointer
	// Optional methods
	PredictionsFromBatchOptionsError(inputBatch unsafe.Pointer, options IMLPredictionOptions, error_ foundation.foundation.INSError) unsafe.Pointer
	HasPredictionsFromBatchOptionsError() bool
}
