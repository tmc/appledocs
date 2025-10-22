// Code generated from Apple documentation for Automator. DO NOT EDIT.

// Package automator provides Go bindings for the Automator framework.
//
// Develop actions that the Automator app can load and run. View, edit, and run Automator workflows in your app.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to Automator without requiring cgo.

// Develop actions that the Automator app can load and run. View, edit, and run Automator workflows in your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator

package automator

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/Automator.framework/Automator"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

