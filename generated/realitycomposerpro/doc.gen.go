// Code generated from Apple documentation for RealityComposerPro. DO NOT EDIT.

// Package realitycomposerpro provides Go bindings for the RealityComposerPro framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to RealityComposerPro without requiring cgo.
package realitycomposerpro

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/RealityComposerPro.framework/RealityComposerPro"

func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}
