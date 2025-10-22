// Code generated from Apple documentation for CreateMLComponents. DO NOT EDIT.

// Package createmlcomponents provides Go bindings for the CreateMLComponents framework.
//
// Create more customizable machine learning models in your app.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to CreateMLComponents without requiring cgo.

// Create more customizable machine learning models in your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CreateMLComponents

package createmlcomponents

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/CreateMLComponents.framework/CreateMLComponents"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

