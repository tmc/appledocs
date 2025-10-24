// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight_test

import (
	"github.com/tmc/appledocs/generated/corespotlight"
)

// Suppress unused import errors
var _ = corespotlight.NewCSSearchQuery

// ExampleCSSearchQuery_Cancel demonstrates using Cancel on a CSSearchQuery instance.
// Cancels the current query operation.
func ExampleCSSearchQuery_Cancel() {
	obj := corespotlight.NewCSSearchQuery()
	obj.Cancel()
	// Output:
	}

// ExampleCSSearchQuery_Start demonstrates using Start on a CSSearchQuery instance.
// Starts searching the index for items that match the current query string and parameters.
func ExampleCSSearchQuery_Start() {
	obj := corespotlight.NewCSSearchQuery()
	obj.Start()
	// Output:
	}

