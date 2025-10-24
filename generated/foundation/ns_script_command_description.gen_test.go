// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewScriptCommandDescription

// ExampleScriptCommandDescription_CreateCommandInstance demonstrates using CreateCommandInstance on a ScriptCommandDescription instance.
// Creates and returns an instance of the command object described by the receiver.
func ExampleScriptCommandDescription_CreateCommandInstance() {
	obj := foundation.NewScriptCommandDescription()
	_ = obj.CreateCommandInstance()
	// Output:
}
