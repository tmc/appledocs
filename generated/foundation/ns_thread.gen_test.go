// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewThread

// ExampleNewThread demonstrates how to create a Thread instance.
// Returns an initialized   object.
func ExampleNewThread() {
	_ = foundation.NewThread()
	// Output:
}
// ExampleThread_Cancel demonstrates using Cancel on a Thread instance.
// Changes the cancelled state of the receiver to indicate that it should exit.
func ExampleThread_Cancel() {
	obj := foundation.NewThread()
	obj.Cancel()
	// Output:
	}

// ExampleThread_Main demonstrates using Main on a Thread instance.
// The main entry point routine for the thread.
func ExampleThread_Main() {
	obj := foundation.NewThread()
	obj.Main()
	// Output:
	}

// ExampleThread_Start demonstrates using Start on a Thread instance.
// Starts the receiver.
func ExampleThread_Start() {
	obj := foundation.NewThread()
	obj.Start()
	// Output:
	}

