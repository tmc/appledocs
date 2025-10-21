// Code generated from Apple documentation for NaturalLanguage. DO NOT EDIT.

// Package naturallanguage provides Go bindings for the NaturalLanguage framework.
//
// Analyze natural language text and deduce its language-specific metadata. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to NaturalLanguage without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage
package naturallanguage

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/NaturalLanguage.framework/NaturalLanguage"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


