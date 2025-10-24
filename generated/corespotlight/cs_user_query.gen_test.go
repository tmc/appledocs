// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight_test

import (
	"github.com/tmc/appledocs/generated/corespotlight"
)

// Suppress unused import errors
var _ = corespotlight.NewCSUserQuery

// ExampleCSUserQuery_Cancel demonstrates using Cancel on a CSUserQuery instance.
// Cancels the current query operation.
func ExampleCSUserQuery_Cancel() {
	obj := corespotlight.NewCSUserQuery()
	obj.Cancel()
	// Output:
	}

// ExampleCSUserQuery_Start demonstrates using Start on a CSUserQuery instance.
// Starts searching the index for items that match the current query   string and parameters.
func ExampleCSUserQuery_Start() {
	obj := corespotlight.NewCSUserQuery()
	obj.Start()
	// Output:
	}

