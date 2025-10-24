// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewProgress

// ExampleNewProgressWithTotalUnitCount demonstrates how to create a Progress instance using NewProgressWithTotalUnitCount.
// Creates and returns a progress instance.
func ExampleNewProgressWithTotalUnitCount() {
	_ = foundation.NewProgressWithTotalUnitCount(
		10, // unitCount int64
	)
	// Output:
}
// ExampleProgress_Cancel demonstrates using Cancel on a Progress instance.
// Cancels progress tracking.
func ExampleProgress_Cancel() {
	obj := foundation.NewProgress()
	obj.Cancel()
	// Output:
	}

// ExampleProgress_Pause demonstrates using Pause on a Progress instance.
// Pauses progress tracking.
func ExampleProgress_Pause() {
	obj := foundation.NewProgress()
	obj.Pause()
	// Output:
	}

// ExampleProgress_Publish demonstrates using Publish on a Progress instance.
// Publishes the progress object for other processes to observe it.
func ExampleProgress_Publish() {
	obj := foundation.NewProgress()
	obj.Publish()
	// Output:
	}

// ExampleProgress_Resume demonstrates using Resume on a Progress instance.
// Resumes progress tracking.
func ExampleProgress_Resume() {
	obj := foundation.NewProgress()
	obj.Resume()
	// Output:
	}

// ExampleProgress_Unpublish demonstrates using Unpublish on a Progress instance.
// Removes a progress object from publication, making it unobservable by other processes.
func ExampleProgress_Unpublish() {
	obj := foundation.NewProgress()
	obj.Unpublish()
	// Output:
	}

