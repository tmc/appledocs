// Code generated from Apple documentation for CoreML. DO NOT EDIT.

// Package coreml provides Go bindings for the CoreML framework.
//
// Integrate machine learning models into your app. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to CoreML without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML
package coreml

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/CoreML.framework/CoreML"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


