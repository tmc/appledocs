// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit_test

import (
	"github.com/tmc/appledocs/generated/screencapturekit"
)

// Suppress unused import errors
var _ = screencapturekit.NewStreamConfiguration

// ExampleNewStreamConfigurationWithPreset demonstrates how to create a StreamConfiguration instance using NewStreamConfigurationWithPreset.
func ExampleNewStreamConfigurationWithPreset() {
	_ = screencapturekit.NewStreamConfigurationWithPreset(
		screencapturekit.StreamConfigurationPreset{}, // preset StreamConfigurationPreset
	)
	// Output:
}
