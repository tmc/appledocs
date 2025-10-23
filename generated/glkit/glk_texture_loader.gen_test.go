// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit_test

import (
	"github.com/tmc/appledocs/generated/glkit"
)

// Suppress unused import errors
var _ = glkit.NewGLKTextureLoader

// ExampleNewGLKTextureLoaderWithShareContext demonstrates how to create a GLKTextureLoader instance using NewGLKTextureLoaderWithShareContext.
// Initializes a new texture loader object.
func ExampleNewGLKTextureLoaderWithShareContext() {
	_ = glkit.NewGLKTextureLoaderWithShareContext(
		glkit.OpenGLContext{}, // context OpenGLContext
	)
	// Output:
}
