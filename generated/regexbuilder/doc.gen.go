// Code generated from Apple documentation for RegexBuilder. DO NOT EDIT.

// Package regexbuilder provides Go bindings for the RegexBuilder framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to RegexBuilder without requiring cgo.
package regexbuilder

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/RegexBuilder.framework/RegexBuilder"

func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}
