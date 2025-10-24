// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewTask

// ExampleNewTask demonstrates how to create a Task instance.
// Returns an initialized process object with the environment of the current process.
func ExampleNewTask() {
	_ = foundation.NewTask()
	// Output:
}
// ExampleTask_Interrupt demonstrates using Interrupt on a Task instance.
// Sends an interrupt signal to the receiver and all of its subtasks.
func ExampleTask_Interrupt() {
	obj := foundation.NewTask()
	obj.Interrupt()
	// Output:
	}

// ExampleTask_Resume demonstrates using Resume on a Task instance.
// Resumes execution of a suspended task.
func ExampleTask_Resume() {
	obj := foundation.NewTask()
	_ = obj.Resume()
	// Output:
	}

// ExampleTask_Suspend demonstrates using Suspend on a Task instance.
// Suspends execution of the receiver task.
func ExampleTask_Suspend() {
	obj := foundation.NewTask()
	_ = obj.Suspend()
	// Output:
	}

// ExampleTask_Terminate demonstrates using Terminate on a Task instance.
// Sends a terminate signal to the receiver and all of its subtasks.
func ExampleTask_Terminate() {
	obj := foundation.NewTask()
	obj.Terminate()
	// Output:
	}

// ExampleTask_WaitUntilExit demonstrates using WaitUntilExit on a Task instance.
// Blocks the process until the receiver is finished.
func ExampleTask_WaitUntilExit() {
	obj := foundation.NewTask()
	obj.WaitUntilExit()
	// Output:
	}

