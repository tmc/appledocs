// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewXPCListener

// ExampleXPCListener_Activate demonstrates using Activate on a XPCListener instance.
// Activates the listener.
func ExampleXPCListener_Activate() {
	obj := foundation.NewXPCListener()
	obj.Activate()
	// Output:
	}

// ExampleXPCListener_Invalidate demonstrates using Invalidate on a XPCListener instance.
// Invalidates the listener.
func ExampleXPCListener_Invalidate() {
	obj := foundation.NewXPCListener()
	obj.Invalidate()
	// Output:
	}

// ExampleXPCListener_Resume demonstrates using Resume on a XPCListener instance.
// Starts processing of incoming requests.
func ExampleXPCListener_Resume() {
	obj := foundation.NewXPCListener()
	obj.Resume()
	// Output:
	}

// ExampleXPCListener_Suspend demonstrates using Suspend on a XPCListener instance.
// Suspends the listener.
func ExampleXPCListener_Suspend() {
	obj := foundation.NewXPCListener()
	obj.Suspend()
	// Output:
	}

