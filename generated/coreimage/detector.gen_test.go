// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage_test

import (
	"github.com/tmc/appledocs/generated/coreimage"
)


// ExampleNewDetectorOfTypeContextOptions demonstrates how to create a Detector instance using NewDetectorOfTypeContextOptions.
// Creates and returns a configured detector.
func ExampleNewDetectorOfTypeContextOptions() {
	_ = coreimage.NewDetectorOfTypeContextOptions(
		"type", // type string
		nil, // context unsafe.Pointer
		nil, // options unsafe.Pointer
	)
	// Output:
}


