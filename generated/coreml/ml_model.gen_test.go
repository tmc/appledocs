// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml_test

import (
	"github.com/tmc/appledocs/generated/coreml"
)

// Suppress unused import errors
var _ = coreml.NewModel

// ExampleModel_NewState demonstrates using NewState on a Model instance.
// Creates a new state object.
func ExampleModel_NewState() {
	obj := coreml.NewModel()
	_ = obj.NewState()
	// Output:
	}

// ExampleModel_Prediction demonstrates using Prediction on a Model instance.
func ExampleModel_Prediction() {
	obj := coreml.NewModel()
	obj.Prediction()
	// Output:
	}

