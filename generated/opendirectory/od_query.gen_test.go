// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory_test

import (
	"github.com/tmc/appledocs/generated/opendirectory"
)

// Suppress unused import errors
var _ = opendirectory.NewODQuery

// ExampleODQuery_Synchronize demonstrates using Synchronize on a ODQuery instance.
// Restarts a query, disposing of any results it has obtained.
func ExampleODQuery_Synchronize() {
	obj := opendirectory.NewODQuery()
	obj.Synchronize()
	// Output:
	}

