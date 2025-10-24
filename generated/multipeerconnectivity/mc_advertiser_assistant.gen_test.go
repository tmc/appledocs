// Code generated from Apple documentation for MultipeerConnectivity. DO NOT EDIT.

package multipeerconnectivity_test

import (
	"github.com/tmc/appledocs/generated/multipeerconnectivity"
)

// Suppress unused import errors
var _ = multipeerconnectivity.NewMCAdvertiserAssistant

// ExampleMCAdvertiserAssistant_Start demonstrates using Start on a MCAdvertiserAssistant instance.
// Begins advertising the service provided by a local peer and starts the assistant.
func ExampleMCAdvertiserAssistant_Start() {
	obj := multipeerconnectivity.NewMCAdvertiserAssistant()
	obj.Start()
	// Output:
	}

// ExampleMCAdvertiserAssistant_Stop demonstrates using Stop on a MCAdvertiserAssistant instance.
// Stops advertising the service provided by a local peer and stops the assistant.
func ExampleMCAdvertiserAssistant_Stop() {
	obj := multipeerconnectivity.NewMCAdvertiserAssistant()
	obj.Stop()
	// Output:
	}

