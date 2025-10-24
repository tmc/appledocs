// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewXPCConnection

// ExampleXPCConnection_Activate demonstrates using Activate on a XPCConnection instance.
// Activates the connection.
func ExampleXPCConnection_Activate() {
	obj := foundation.NewXPCConnection()
	obj.Activate()
	// Output:
	}

// ExampleXPCConnection_Invalidate demonstrates using Invalidate on a XPCConnection instance.
// Invalidates the connection.
func ExampleXPCConnection_Invalidate() {
	obj := foundation.NewXPCConnection()
	obj.Invalidate()
	// Output:
	}

// ExampleXPCConnection_Resume demonstrates using Resume on a XPCConnection instance.
// Starts or resumes handling of messages on a connection.
func ExampleXPCConnection_Resume() {
	obj := foundation.NewXPCConnection()
	obj.Resume()
	// Output:
	}

// ExampleXPCConnection_Suspend demonstrates using Suspend on a XPCConnection instance.
// Suspends the connection.
func ExampleXPCConnection_Suspend() {
	obj := foundation.NewXPCConnection()
	obj.Suspend()
	// Output:
	}

