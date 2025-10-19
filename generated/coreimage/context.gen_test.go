// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage_test

import (
	"github.com/tmc/appledocs/generated/coreimage"
)


// ExampleNewContextWithEAGLContextOptions demonstrates how to create a Context instance using NewContextWithEAGLContextOptions.
// Creates a Core Image context from an EAGL context using the specified options.
func ExampleNewContextWithEAGLContextOptions() {
	_ = coreimage.NewContextWithEAGLContextOptions(
		nil, // eaglContext unsafe.Pointer
		nil, // options unsafe.Pointer
	)
	// Output:
}

// ExampleNewContextWithMTLCommandQueueOptions demonstrates how to create a Context instance using NewContextWithMTLCommandQueueOptions.
func ExampleNewContextWithMTLCommandQueueOptions() {
	_ = coreimage.NewContextWithMTLCommandQueueOptions(
		nil, // commandQueue unsafe.Pointer
		nil, // options unsafe.Pointer
	)
	// Output:
}

// ExampleNewContextWithMTLDeviceOptions demonstrates how to create a Context instance using NewContextWithMTLDeviceOptions.
// Creates a Core Image context using the specified Metal device and options.
func ExampleNewContextWithMTLDeviceOptions() {
	_ = coreimage.NewContextWithMTLDeviceOptions(
		nil, // device unsafe.Pointer
		nil, // options unsafe.Pointer
	)
	// Output:
}

// ExampleNewContextWithOptions demonstrates how to create a Context instance using NewContextWithOptions.
// Initializes a context without a specific rendering destination, using the specified options.
func ExampleNewContextWithOptions() {
	_ = coreimage.NewContextWithOptions(
		nil, // options unsafe.Pointer
	)
	// Output:
}

// ExampleNewContext demonstrates how to create a Context instance.
// Initializes a context without a specific rendering destination, using default options.
func ExampleNewContext() {
	_ = coreimage.NewContext()
	// Output:
}

// ExampleNewContextWithCGLContextPixelFormatColorSpaceOptions demonstrates how to create a Context instance using NewContextWithCGLContextPixelFormatColorSpaceOptions.
// Creates a Core Image context from a CGL context, using the specified options, color space, and pixel format object.
func ExampleNewContextWithCGLContextPixelFormatColorSpaceOptions() {
	_ = coreimage.NewContextWithCGLContextPixelFormatColorSpaceOptions(
		nil, // cglctx unsafe.Pointer
		nil, // pixelFormat unsafe.Pointer
		nil, // colorSpace unsafe.Pointer
		nil, // options unsafe.Pointer
	)
	// Output:
}

// ExampleNewContextForOfflineGPUAtIndex demonstrates how to create a Context instance using NewContextForOfflineGPUAtIndex.
// Creates an OpenGL-based Core Image context using a GPU that is not currently driving a display.
func ExampleNewContextForOfflineGPUAtIndex() {
	_ = coreimage.NewContextForOfflineGPUAtIndex(
		nil, // index unsafe.Pointer
	)
	// Output:
}

// ExampleNewContextForOfflineGPUAtIndexColorSpaceOptionsSharedContext demonstrates how to create a Context instance using NewContextForOfflineGPUAtIndexColorSpaceOptionsSharedContext.
// Creates an OpenGL-based Core Image context using a GPU that is not currently driving a display, with the specified options.
func ExampleNewContextForOfflineGPUAtIndexColorSpaceOptionsSharedContext() {
	_ = coreimage.NewContextForOfflineGPUAtIndexColorSpaceOptionsSharedContext(
		nil, // index unsafe.Pointer
		nil, // colorSpace unsafe.Pointer
		nil, // options unsafe.Pointer
		nil, // sharedContext unsafe.Pointer
	)
	// Output:
}

// ExampleNewContextWithMTLCommandQueue demonstrates how to create a Context instance using NewContextWithMTLCommandQueue.
func ExampleNewContextWithMTLCommandQueue() {
	_ = coreimage.NewContextWithMTLCommandQueue(
		nil, // commandQueue unsafe.Pointer
	)
	// Output:
}

// ExampleNewContextWithMTLDevice demonstrates how to create a Context instance using NewContextWithMTLDevice.
// Creates a Core Image context using the specified Metal device.
func ExampleNewContextWithMTLDevice() {
	_ = coreimage.NewContextWithMTLDevice(
		nil, // device unsafe.Pointer
	)
	// Output:
}

// ExampleNewContextWithCGContextOptions demonstrates how to create a Context instance using NewContextWithCGContextOptions.
// Creates a Core Image context from a Quartz context, using the specified options.
func ExampleNewContextWithCGContextOptions() {
	_ = coreimage.NewContextWithCGContextOptions(
		nil, // cgctx unsafe.Pointer
		nil, // options unsafe.Pointer
	)
	// Output:
}

// ExampleNewContextWithEAGLContext demonstrates how to create a Context instance using NewContextWithEAGLContext.
// Creates a Core Image context from an EAGL context.
func ExampleNewContextWithEAGLContext() {
	_ = coreimage.NewContextWithEAGLContext(
		nil, // eaglContext unsafe.Pointer
	)
	// Output:
}


