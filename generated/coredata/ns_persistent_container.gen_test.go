// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata_test

import (
	"github.com/tmc/appledocs/generated/coredata"
)

// Suppress unused import errors
var _ = coredata.NewPersistentContainer

// ExamplePersistentContainer_NewBackgroundContext demonstrates using NewBackgroundContext on a PersistentContainer instance.
// Returns a new managed object context that executes on a private queue.
func ExamplePersistentContainer_NewBackgroundContext() {
	obj := coredata.NewPersistentContainer()
	_ = obj.NewBackgroundContext()
	// Output:
	}

