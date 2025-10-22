// Code generated from Apple documentation for CreateML. DO NOT EDIT.

// Package createml provides Go bindings for the CreateML framework.
//
// Create machine learning models for use in your app.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to CreateML without requiring cgo.

// Create machine learning models for use in your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CreateML

package createml

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/CreateML.framework/CreateML"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

