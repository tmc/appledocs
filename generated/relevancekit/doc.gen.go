// Code generated from Apple documentation for RelevanceKit. DO NOT EDIT.

// Package relevancekit provides Go bindings for the RelevanceKit framework.
//
// Provide on-device intelligence with contextual clues that increase your widget’s visibility on Apple Watch. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to RelevanceKit without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/RelevanceKit
package relevancekit

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/RelevanceKit.framework/RelevanceKit"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


