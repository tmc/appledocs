// Code generated from Apple documentation for ScriptingBridge. DO NOT EDIT.

package scriptingbridge_test

import (
	"github.com/tmc/appledocs/generated/scriptingbridge"
)

// Suppress unused import errors
var _ = scriptingbridge.NewSBApplication

// ExampleSBApplication_Activate demonstrates using Activate on a SBApplication instance.
// Moves the target application to the foreground immediately.
func ExampleSBApplication_Activate() {
	obj := scriptingbridge.NewSBApplication()
	obj.Activate()
	// Output:
	}

