// Code generated from Apple documentation for PhotosUI. DO NOT EDIT.

package photosui_test

import (
	"github.com/tmc/appledocs/generated/photosui"
)

// Suppress unused import errors
var _ = photosui.NewPHPickerViewController

// ExampleNewPHPickerViewControllerWithConfiguration demonstrates how to create a PHPickerViewController instance using NewPHPickerViewControllerWithConfiguration.
// Creates a new picker view controller with the configuration you specify.
func ExampleNewPHPickerViewControllerWithConfiguration() {
	_ = photosui.NewPHPickerViewControllerWithConfiguration(
		photosui.PHPickerConfiguration{}, // configuration PHPickerConfiguration
	)
	// Output:
}
