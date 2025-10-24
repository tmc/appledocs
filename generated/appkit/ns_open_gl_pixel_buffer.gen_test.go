// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewOpenGLPixelBuffer

// ExampleNewOpenGLPixelBufferWithCGLPBufferObj demonstrates how to create a OpenGLPixelBuffer instance using NewOpenGLPixelBufferWithCGLPBufferObj.
// Initializes and returns an OpenGL pixel buffer object that encapsulates an existing CGL pixel buffer object.
func ExampleNewOpenGLPixelBufferWithCGLPBufferObj() {
	_ = appkit.NewOpenGLPixelBufferWithCGLPBufferObj(
		appkit.LPBufferObj /* not a class type */{}, // pbuffer LPBufferObj /* not a class type */
	)
	// Output:
}
