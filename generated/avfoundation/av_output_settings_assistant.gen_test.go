// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation_test

import (
	"github.com/tmc/appledocs/generated/avfoundation"
)

// Suppress unused import errors
var _ = avfoundation.NewOutputSettingsAssistant

// ExampleNewOutputSettingsAssistantWithPreset demonstrates how to create a OutputSettingsAssistant instance using NewOutputSettingsAssistantWithPreset.
// Creates an output setting assistant with a preset configuration.
func ExampleNewOutputSettingsAssistantWithPreset() {
	_ = avfoundation.NewOutputSettingsAssistantWithPreset(
		avfoundation.OutputSettingsPreset{}, // presetIdentifier OutputSettingsPreset
	)
	// Output:
}
