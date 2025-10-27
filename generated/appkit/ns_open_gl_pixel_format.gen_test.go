// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewOpenGLPixelFormat

// ExampleNewOpenGLPixelFormatWithAttributes demonstrates how to create a OpenGLPixelFormat instance using NewOpenGLPixelFormatWithAttributes.
// Returns an OpenGL pixel format object initialized with specified pixel format attributes.
func ExampleNewOpenGLPixelFormatWithAttributes() {
	_ = appkit.NewOpenGLPixelFormatWithAttributes(
		appkit.OpenGLPixelFormatAttribute{}, // attribs OpenGLPixelFormatAttribute
	)
	// Output:
}
// ExampleNewOpenGLPixelFormatWithCGLPixelFormatObj demonstrates how to create a OpenGLPixelFormat instance using NewOpenGLPixelFormatWithCGLPixelFormatObj.
// Returns an OpenGL pixel format object initialized with using an existing CGL pixel format object.
func ExampleNewOpenGLPixelFormatWithCGLPixelFormatObj() {
	_ = appkit.NewOpenGLPixelFormatWithCGLPixelFormatObj(
		appkit.LPixelFormatObj /* not a class type */{}, // format LPixelFormatObj /* not a class type */
	)
	// Output:
}
