// Code generated from Apple documentation for PaperKit. DO NOT EDIT.

// Package paperkit provides Go bindings for the PaperKit framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to PaperKit without requiring cgo.
package paperkit

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/PaperKit.framework/PaperKit"

func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}
