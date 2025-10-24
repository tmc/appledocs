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
	EvaluateOnCPUWithInputsOutputsError(inputs []MultiArray, outputs []MultiArray, error_ objectivec.IObject) bool/* debug [protocol_interface/required_method]: EvaluateOnCPUWithInputsOutputsError */
	InitWithParameterDictionaryError(parameters foundation.IDictionary, error_ objectivec.IObject) objectivec.IObject/* debug [protocol_interface/required_method]: InitWithParameterDictionaryError */
	OutputShapesForInputShapesError(inputShapes []foundation.Array, error_ objectivec.IObject) []foundation.Array/* debug [protocol_interface/required_method]: OutputShapesForInputShapesError */
	SetWeightDataError(weights []foundation.Data, error_ objectivec.IObject) bool/* debug [protocol_interface/required_method]: SetWeightDataError */
	// Optional methods
	EncodeToCommandBufferInputsOutputsError(commandBuffer unsafe.Pointer, inputs []objc.ID, outputs []objc.ID, error_ objectivec.IObject) bool
	HasEncodeToCommandBufferInputsOutputsError() bool
}
