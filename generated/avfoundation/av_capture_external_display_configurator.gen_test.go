// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation_test

import (
	"github.com/tmc/appledocs/generated/avfoundation"
)

// Suppress unused import errors
var _ = avfoundation.NewCaptureExternalDisplayConfigurator

// ExampleCaptureExternalDisplayConfigurator_Stop demonstrates using Stop on a CaptureExternalDisplayConfigurator instance.
// Forces the external display configurator to asynchronously stop configuring the external display.
func ExampleCaptureExternalDisplayConfigurator_Stop() {
	obj := avfoundation.NewCaptureExternalDisplayConfigurator()
	obj.Stop()
	// Output:
	}

