// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage_test

import (
	"github.com/tmc/appledocs/generated/coreimage"
)


// ExampleNewKernelWithFunctionNameFromMetalLibraryDataError demonstrates how to create a Kernel instance using NewKernelWithFunctionNameFromMetalLibraryDataError.
// Creates a single kernel object using a Metal Shading Language (MSL) kernel function.
func ExampleNewKernelWithFunctionNameFromMetalLibraryDataError() {
	_ = coreimage.NewKernelWithFunctionNameFromMetalLibraryDataError(
		"name", // name string
		nil, // data unsafe.Pointer
		nil, // error unsafe.Pointer
	)
	// Output:
}

// ExampleNewKernelWithFunctionNameFromMetalLibraryDataOutputPixelFormatError demonstrates how to create a Kernel instance using NewKernelWithFunctionNameFromMetalLibraryDataOutputPixelFormatError.
// Creates a single kernel object using a Metal Shading Language kernel function with optional pixel format.
func ExampleNewKernelWithFunctionNameFromMetalLibraryDataOutputPixelFormatError() {
	_ = coreimage.NewKernelWithFunctionNameFromMetalLibraryDataOutputPixelFormatError(
		"name", // name string
		nil, // data unsafe.Pointer
		nil, // format unsafe.Pointer
		nil, // error unsafe.Pointer
	)
	// Output:
}

// ExampleNewKernelWithString demonstrates how to create a Kernel instance using NewKernelWithString.
// Creates a single kernel object.
func ExampleNewKernelWithString() {
	_ = coreimage.NewKernelWithString(
		"string", // string string
	)
	// Output:
}


