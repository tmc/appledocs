// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewURLSession

// ExampleNewURLSession demonstrates how to create a URLSession instance.
func ExampleNewURLSession() {
	_ = foundation.NewURLSession()
	// Output:
}
// ExampleURLSession_FinishTasksAndInvalidate demonstrates using FinishTasksAndInvalidate on a URLSession instance.
// Invalidates the session, allowing any outstanding tasks to finish.
func ExampleURLSession_FinishTasksAndInvalidate() {
	obj := foundation.NewURLSession()
	obj.FinishTasksAndInvalidate()
	// Output:
	}

// ExampleURLSession_InvalidateAndCancel demonstrates using InvalidateAndCancel on a URLSession instance.
// Cancels all outstanding tasks and then invalidates the session.
func ExampleURLSession_InvalidateAndCancel() {
	obj := foundation.NewURLSession()
	obj.InvalidateAndCancel()
	// Output:
	}

