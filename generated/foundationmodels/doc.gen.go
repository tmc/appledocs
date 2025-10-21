// Code generated from Apple documentation for FoundationModels. DO NOT EDIT.

// Package foundationmodels provides Go bindings for the FoundationModels framework.
//
// Perform tasks with the on-device model that specializes in language understanding, structured output, and tool calling. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to FoundationModels without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/FoundationModels
package foundationmodels

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/FoundationModels.framework/FoundationModels"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

