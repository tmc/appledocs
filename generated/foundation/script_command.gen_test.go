// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewScriptCommand

// ExampleNewScriptCommandWithCoder demonstrates how to create a ScriptCommand instance using NewScriptCommandWithCoder.
func ExampleNewScriptCommandWithCoder() {
	_ = foundation.NewScriptCommandWithCoder(
		foundation.NSCoder{}, // inCoder NSCoder
	)
	// Output:
}
// ExampleNewScriptCommandWithCommandDescription demonstrates how to create a ScriptCommand instance using NewScriptCommandWithCommandDescription.
// Returns an a script command object initialized from the passed command description.
func ExampleNewScriptCommandWithCommandDescription() {
	_ = foundation.NewScriptCommandWithCommandDescription(
		foundation.NSScriptCommandDescription{}, // commandDef NSScriptCommandDescription
	)
	// Output:
}
