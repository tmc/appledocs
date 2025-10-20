// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage_test

import (
	"github.com/tmc/appledocs/generated/coreimage"
)

// Suppress unused import errors
var _ = coreimage.NewContext


// ExampleNewContextWithMTLDevice demonstrates how to create a Context instance using NewContextWithMTLDevice.
// Creates a Core Image context using the specified Metal device.
func ExampleNewContextWithMTLDevice() {
	_ = coreimage.NewContextWithMTLDevice(
		0, // device objc.ID
	)
	// Output:
}







// ExampleNewContext demonstrates how to create a Context instance.
// Initializes a context without a specific rendering destination, using default options.
func ExampleNewContext() {
	_ = coreimage.NewContext()
	// Output:
}




// ExampleNewContextWithMTLCommandQueue demonstrates how to create a Context instance using NewContextWithMTLCommandQueue.
func ExampleNewContextWithMTLCommandQueue() {
	_ = coreimage.NewContextWithMTLCommandQueue(
		0, // commandQueue objc.ID
	)
	// Output:
}


