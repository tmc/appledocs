// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata_test

import (
	"github.com/tmc/appledocs/generated/coredata"
)


// ExampleNewPersistentStoreWithPersistentStoreCoordinatorConfigurationNameURLOptions demonstrates how to create a PersistentStore instance using NewPersistentStoreWithPersistentStoreCoordinatorConfigurationNameURLOptions.
// Returns a store initialized with the given arguments.
func ExampleNewPersistentStoreWithPersistentStoreCoordinatorConfigurationNameURLOptions() {
	_ = coredata.NewPersistentStoreWithPersistentStoreCoordinatorConfigurationNameURLOptions(
		nil, // root unsafe.Pointer
		"name", // name string
		nil, // url unsafe.Pointer
		nil, // options unsafe.Pointer
	)
	// Output:
}


