// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewBackgroundActivityScheduler

// ExampleBackgroundActivityScheduler_Invalidate demonstrates using Invalidate on a BackgroundActivityScheduler instance.
// Prevents the background activity from being scheduled again.
func ExampleBackgroundActivityScheduler_Invalidate() {
	obj := foundation.NewBackgroundActivityScheduler()
	obj.Invalidate()
	// Output:
	}

