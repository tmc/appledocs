// Code generated from Apple documentation for ScriptingBridge. DO NOT EDIT.

// Package scriptingbridge provides Go bindings for the ScriptingBridge framework.
//
// Automate scriptable apps by sending and receiving Apple events. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to ScriptingBridge without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/ScriptingBridge
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

