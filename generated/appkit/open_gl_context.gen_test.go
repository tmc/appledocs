// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewOpenGLContext

// ExampleNewOpenGLContextWithCGLContextObj demonstrates how to create a OpenGLContext instance using NewOpenGLContextWithCGLContextObj.
// Initializes and returns an OpenGL context object using an existing CGL context.
func ExampleNewOpenGLContextWithCGLContextObj() {
	_ = appkit.NewOpenGLContextWithCGLContextObj(
		appkit.LContextObj /* not a class type */ {}, // context LContextObj /* not a class type */
	)
	// Output:
}
