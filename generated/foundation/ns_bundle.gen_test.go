// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewBundle

// ExampleBundle_LoadAppleScriptObjectiveCScripts demonstrates using LoadAppleScriptObjectiveCScripts on a Bundle instance.
func ExampleBundle_LoadAppleScriptObjectiveCScripts() {
	obj := foundation.NewBundle()
	obj.LoadAppleScriptObjectiveCScripts()
	// Output:
	}

// ExampleBundle_Unload demonstrates using Unload on a Bundle instance.
// Unloads the code associated with the receiver.
func ExampleBundle_Unload() {
	obj := foundation.NewBundle()
	_ = obj.Unload()
	// Output:
	}

