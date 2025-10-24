// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewRunningApplication

// ExampleRunningApplication_ForceTerminate demonstrates using ForceTerminate on a RunningApplication instance.
// Attempts to force the receiver to quit.
func ExampleRunningApplication_ForceTerminate() {
	obj := appkit.NewRunningApplication()
	_ = obj.ForceTerminate()
	// Output:
	}

// ExampleRunningApplication_Hide demonstrates using Hide on a RunningApplication instance.
// Attempts to hide or the application.
func ExampleRunningApplication_Hide() {
	obj := appkit.NewRunningApplication()
	_ = obj.Hide()
	// Output:
	}

// ExampleRunningApplication_Terminate demonstrates using Terminate on a RunningApplication instance.
// Attempts to quit the receiver normally.
func ExampleRunningApplication_Terminate() {
	obj := appkit.NewRunningApplication()
	_ = obj.Terminate()
	// Output:
	}

// ExampleRunningApplication_Unhide demonstrates using Unhide on a RunningApplication instance.
// Attempts to unhide or the application.
func ExampleRunningApplication_Unhide() {
	obj := appkit.NewRunningApplication()
	_ = obj.Unhide()
	// Output:
	}

