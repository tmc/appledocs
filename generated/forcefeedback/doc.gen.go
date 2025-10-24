
// Code generated from Apple documentation for ForceFeedback. DO NOT EDIT.

// Package forcefeedback provides Go bindings for the ForceFeedback framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to ForceFeedback without requiring cgo.
package forcefeedback

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/ForceFeedback.framework/ForceFeedback"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

