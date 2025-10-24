// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewScriptCommand

// ExampleScriptCommand_ExecuteCommand demonstrates using ExecuteCommand on a ScriptCommand instance.
// Executes the command if it is valid and returns the result, if any.
func ExampleScriptCommand_ExecuteCommand() {
	obj := foundation.NewScriptCommand()
	_ = obj.ExecuteCommand()
	// Output:
	}

// ExampleScriptCommand_PerformDefaultImplementation demonstrates using PerformDefaultImplementation on a ScriptCommand instance.
// Overridden by subclasses to provide a default implementation for the command represented by the receiver.
func ExampleScriptCommand_PerformDefaultImplementation() {
	obj := foundation.NewScriptCommand()
	_ = obj.PerformDefaultImplementation()
	// Output:
	}

// ExampleScriptCommand_SuspendExecution demonstrates using SuspendExecution on a ScriptCommand instance.
// Suspends the execution of the receiver.
func ExampleScriptCommand_SuspendExecution() {
	obj := foundation.NewScriptCommand()
	obj.SuspendExecution()
	// Output:
	}

