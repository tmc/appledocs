// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewURLSessionTask

// ExampleNewURLSessionTask demonstrates how to create a URLSessionTask instance.
// Initializes an empty URL sesson task.
func ExampleNewURLSessionTask() {
	_ = foundation.NewURLSessionTask()
	// Output:
}

// ExampleURLSessionTask_Cancel demonstrates using Cancel on a URLSessionTask instance.
// Cancels the task.
func ExampleURLSessionTask_Cancel() {
	obj := foundation.NewURLSessionTask()
	obj.Cancel()
	// Output:
}

// ExampleURLSessionTask_Resume demonstrates using Resume on a URLSessionTask instance.
// Resumes the task, if it is suspended.
func ExampleURLSessionTask_Resume() {
	obj := foundation.NewURLSessionTask()
	obj.Resume()
	// Output:
}

// ExampleURLSessionTask_Suspend demonstrates using Suspend on a URLSessionTask instance.
// Temporarily suspends a task.
func ExampleURLSessionTask_Suspend() {
	obj := foundation.NewURLSessionTask()
	obj.Suspend()
	// Output:
}
