// Code generated from Apple documentation for ClassKit. DO NOT EDIT.

package classkit_test

import (
	"github.com/tmc/appledocs/generated/classkit"
)

// Suppress unused import errors
var _ = classkit.NewSContext

// ExampleSContext_BecomeActive demonstrates using BecomeActive on a SContext instance.
// Tells a context to become the active context.
func ExampleSContext_BecomeActive() {
	obj := classkit.NewSContext()
	obj.BecomeActive()
	// Output:
	}

// ExampleSContext_CreateNewActivity demonstrates using CreateNewActivity on a SContext instance.
// Creates and returns a new activity instance for the context.
func ExampleSContext_CreateNewActivity() {
	obj := classkit.NewSContext()
	_ = obj.CreateNewActivity()
	// Output:
	}

// ExampleSContext_RemoveFromParent demonstrates using RemoveFromParent on a SContext instance.
// Removes the context from its parent.
func ExampleSContext_RemoveFromParent() {
	obj := classkit.NewSContext()
	obj.RemoveFromParent()
	// Output:
	}

// ExampleSContext_ResetProgressReportingCapabilities demonstrates using ResetProgressReportingCapabilities on a SContext instance.
// Resets the set of capabilities for the context.
func ExampleSContext_ResetProgressReportingCapabilities() {
	obj := classkit.NewSContext()
	obj.ResetProgressReportingCapabilities()
	// Output:
	}

// ExampleSContext_ResignActive demonstrates using ResignActive on a SContext instance.
// Tells a context to stop being the active context.
func ExampleSContext_ResignActive() {
	obj := classkit.NewSContext()
	obj.ResignActive()
	// Output:
	}

