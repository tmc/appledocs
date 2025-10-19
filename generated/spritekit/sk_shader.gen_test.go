// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit_test

import (
	"github.com/tmc/appledocs/generated/spritekit"
)

// Suppress unused import errors
var _ = spritekit.NewSKShader


// ExampleNewSKShaderWithFileNamed demonstrates how to create a SKShader instance using NewSKShaderWithFileNamed.
// Creates a new shader object by loading the source for a fragment shader from a file stored in the app’s bundle.
func ExampleNewSKShaderWithFileNamed() {
	_ = spritekit.NewSKShaderWithFileNamed(
		"name", // name string
	)
	// Output:
}

// ExampleNewSKShaderWithSource demonstrates how to create a SKShader instance using NewSKShaderWithSource.
// Initializes a new shader object using the specified source code.
func ExampleNewSKShaderWithSource() {
	_ = spritekit.NewSKShaderWithSource(
		"source", // source string
	)
	// Output:
}



