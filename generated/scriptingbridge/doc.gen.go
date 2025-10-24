
// Code generated from Apple documentation for ScriptingBridge. DO NOT EDIT.

// Package scriptingbridge provides Go bindings for the ScriptingBridge framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to ScriptingBridge without requiring cgo.
package scriptingbridge

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/ScriptingBridge.framework/ScriptingBridge"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

