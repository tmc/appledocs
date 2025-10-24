// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore_test

import (
	"github.com/tmc/appledocs/generated/imagecapturecore"
)

// Suppress unused import errors
var _ = imagecapturecore.NewICDeviceBrowser

// ExampleNewICDeviceBrowser demonstrates how to create a ICDeviceBrowser instance.
// Creates an ImageCaptureCore device browser.
func ExampleNewICDeviceBrowser() {
	_ = imagecapturecore.NewICDeviceBrowser()
	// Output:
}
// ExampleICDeviceBrowser_Stop demonstrates using Stop on a ICDeviceBrowser instance.
// Tells the delegate to stop looking for devices.
func ExampleICDeviceBrowser_Stop() {
	obj := imagecapturecore.NewICDeviceBrowser()
	obj.Stop()
	// Output:
	}

// ExampleICDeviceBrowser_Start demonstrates using Start on a ICDeviceBrowser instance.
// Tells the delegate to start looking for devices.
func ExampleICDeviceBrowser_Start() {
	obj := imagecapturecore.NewICDeviceBrowser()
	obj.Start()
	// Output:
	}

