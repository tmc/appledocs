// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision_test

import (
	"github.com/tmc/appledocs/generated/vision"
)

// Suppress unused import errors
var _ = vision.NewCoreMLRequest

// ExampleNewCoreMLRequestWithModel demonstrates how to create a CoreMLRequest instance using NewCoreMLRequestWithModel.
// Creates a model container to use with an image analysis request based on the model you provide.
func ExampleNewCoreMLRequestWithModel() {
	_ = vision.NewCoreMLRequestWithModel(
		vision.VNCoreMLModel{}, // model VNCoreMLModel
	)
	// Output:
}
