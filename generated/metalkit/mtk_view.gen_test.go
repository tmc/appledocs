// Code generated from Apple documentation for MetalKit. DO NOT EDIT.

package metalkit_test

import (
	"github.com/tmc/appledocs/generated/metalkit"
)


// ExampleNewMTKViewWithCoder demonstrates how to create a MTKView instance using NewMTKViewWithCoder.
// Initializes a view from data in a given unarchiver.
func ExampleNewMTKViewWithCoder() {
	_ = metalkit.NewMTKViewWithCoder(
		nil, // coder unsafe.Pointer
	)
	// Output:
}

// ExampleNewMTKViewWithFrameDevice demonstrates how to create a MTKView instance using NewMTKViewWithFrameDevice.
// Initializes a view with the specified frame rectangle and Metal device.
func ExampleNewMTKViewWithFrameDevice() {
	_ = metalkit.NewMTKViewWithFrameDevice(
		nil, // frameRect unsafe.Pointer
		nil, // device unsafe.Pointer
	)
	// Output:
}


