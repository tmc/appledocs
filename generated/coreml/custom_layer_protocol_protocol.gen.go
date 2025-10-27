// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PCustomLayer is the MLCustomLayer protocol interface.
//
// An interface that defines the behavior of a custom layer in your neural network model.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 11.2+
//   - iPadOS 11.2+
//   - macOS 10.13.2+
//   - tvOS 11.2+
//   - visionOS 1.0+
//   - watchOS 4.2+
//
// See: doc://com.apple.coreml/documentation/CoreML/MLCustomLayer
type PCustomLayer interface {
	// Required methods
	EvaluateOnCPUWithInputsOutputsError(inputs []MultiArray, outputs []MultiArray, error_ foundation.foundation.INSError) bool
	InitWithParameterDictionaryError(parameters foundation.IDictionary, error_ foundation.foundation.INSError) objectivec.IObject
	OutputShapesForInputShapesError(inputShapes []foundation.Array, error_ foundation.foundation.INSError) []foundation.Array
	SetWeightDataError(weights []foundation.Data, error_ foundation.foundation.INSError) bool
	// Optional methods
	EncodeToCommandBufferInputsOutputsError(commandBuffer unsafe.Pointer, inputs []objc.ID, outputs []objc.ID, error_ foundation.foundation.INSError) bool
	HasEncodeToCommandBufferInputsOutputsError() bool
}
