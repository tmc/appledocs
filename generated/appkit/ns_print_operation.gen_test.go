// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewPrintOperation

// ExamplePrintOperation_CleanUpOperation demonstrates using CleanUpOperation on a PrintOperation instance.
// Called at the end of a print operation to remove the print operation as the current operation.
func ExamplePrintOperation_CleanUpOperation() {
	obj := appkit.NewPrintOperation()
	obj.CleanUpOperation()
	// Output:
	}

// ExamplePrintOperation_CreateContext demonstrates using CreateContext on a PrintOperation instance.
// Creates the graphics context object used for drawing during the operation.
func ExamplePrintOperation_CreateContext() {
	obj := appkit.NewPrintOperation()
	_ = obj.CreateContext()
	// Output:
	}

// ExamplePrintOperation_DeliverResult demonstrates using DeliverResult on a PrintOperation instance.
// Delivers the results of the print operation to the intended destination.
func ExamplePrintOperation_DeliverResult() {
	obj := appkit.NewPrintOperation()
	_ = obj.DeliverResult()
	// Output:
	}

// ExamplePrintOperation_DestroyContext demonstrates using DestroyContext on a PrintOperation instance.
// Destroys the print operation’s graphics context.
func ExamplePrintOperation_DestroyContext() {
	obj := appkit.NewPrintOperation()
	obj.DestroyContext()
	// Output:
	}

// ExamplePrintOperation_RunOperation demonstrates using RunOperation on a PrintOperation instance.
// Runs the print operation on the current thread.
func ExamplePrintOperation_RunOperation() {
	obj := appkit.NewPrintOperation()
	_ = obj.RunOperation()
	// Output:
	}

