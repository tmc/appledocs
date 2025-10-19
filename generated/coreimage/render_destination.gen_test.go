// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage_test

import (
	"github.com/tmc/appledocs/generated/coreimage"
)


// ExampleNewRenderDestinationWithBitmapDataWidthHeightBytesPerRowFormat demonstrates how to create a RenderDestination instance using NewRenderDestinationWithBitmapDataWidthHeightBytesPerRowFormat.
// Creates a render destination based on a client-managed buffer.
func ExampleNewRenderDestinationWithBitmapDataWidthHeightBytesPerRowFormat() {
	_ = coreimage.NewRenderDestinationWithBitmapDataWidthHeightBytesPerRowFormat(
		nil, // data unsafe.Pointer
		0, // width uint
		0, // height uint
		0, // bytesPerRow uint
		nil, // format unsafe.Pointer
	)
	// Output:
}

// ExampleNewRenderDestinationWithGLTextureTargetWidthHeight demonstrates how to create a RenderDestination instance using NewRenderDestinationWithGLTextureTargetWidthHeight.
// Creates a render destination based on an OpenGL texture.
func ExampleNewRenderDestinationWithGLTextureTargetWidthHeight() {
	_ = coreimage.NewRenderDestinationWithGLTextureTargetWidthHeight(
		nil, // texture unsafe.Pointer
		nil, // target unsafe.Pointer
		0, // width uint
		0, // height uint
	)
	// Output:
}

// ExampleNewRenderDestinationWithIOSurface demonstrates how to create a RenderDestination instance using NewRenderDestinationWithIOSurface.
// Creates a render destination based on an   object.
func ExampleNewRenderDestinationWithIOSurface() {
	_ = coreimage.NewRenderDestinationWithIOSurface(
		nil, // surface unsafe.Pointer
	)
	// Output:
}

// ExampleNewRenderDestinationWithMTLTextureCommandBuffer demonstrates how to create a RenderDestination instance using NewRenderDestinationWithMTLTextureCommandBuffer.
// Creates a render destination based on a Metal texture.
func ExampleNewRenderDestinationWithMTLTextureCommandBuffer() {
	_ = coreimage.NewRenderDestinationWithMTLTextureCommandBuffer(
		nil, // texture unsafe.Pointer
		nil, // commandBuffer unsafe.Pointer
	)
	// Output:
}

// ExampleNewRenderDestinationWithPixelBuffer demonstrates how to create a RenderDestination instance using NewRenderDestinationWithPixelBuffer.
// Creates a render destination based on a Core Video pixel buffer.
func ExampleNewRenderDestinationWithPixelBuffer() {
	_ = coreimage.NewRenderDestinationWithPixelBuffer(
		nil, // pixelBuffer unsafe.Pointer
	)
	// Output:
}

// ExampleNewRenderDestinationWithWidthHeightPixelFormatCommandBufferMtlTextureProvider demonstrates how to create a RenderDestination instance using NewRenderDestinationWithWidthHeightPixelFormatCommandBufferMtlTextureProvider.
// Creates a render destination based on a Metal texture with specified pixel format.
func ExampleNewRenderDestinationWithWidthHeightPixelFormatCommandBufferMtlTextureProvider() {
	_ = coreimage.NewRenderDestinationWithWidthHeightPixelFormatCommandBufferMtlTextureProvider(
		0, // width uint
		0, // height uint
		nil, // pixelFormat unsafe.Pointer
		nil, // commandBuffer unsafe.Pointer
		nil, // block unsafe.Pointer
	)
	// Output:
}


