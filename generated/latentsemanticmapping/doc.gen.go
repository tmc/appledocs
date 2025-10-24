
// Code generated from Apple documentation for LatentSemanticMapping. DO NOT EDIT.

// Package latentsemanticmapping provides Go bindings for the LatentSemanticMapping framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to LatentSemanticMapping without requiring cgo.
package latentsemanticmapping

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/LatentSemanticMapping.framework/LatentSemanticMapping"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

