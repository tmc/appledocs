// Code generated from Apple documentation for Speech. DO NOT EDIT.

// Package speech provides Go bindings for the Speech framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to Speech without requiring cgo.
package speech

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/Speech.framework/Speech"

func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}
