// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata_test

import (
	"github.com/tmc/appledocs/generated/coredata"
)


// ExampleNewBatchInsertRequestWithEntityNameDictionaryHandler demonstrates how to create a BatchInsertRequest instance using NewBatchInsertRequestWithEntityNameDictionaryHandler.
// Creates a batch-insertion request for a named managed entity, and specifies a closure that provides data dictionaries for insertion.
func ExampleNewBatchInsertRequestWithEntityNameDictionaryHandler() {
	_ = coredata.NewBatchInsertRequestWithEntityNameDictionaryHandler(
		"entityName", // entityName string
		nil, // handler unsafe.Pointer
	)
	// Output:
}

// ExampleNewBatchInsertRequestWithEntityNameManagedObjectHandler demonstrates how to create a BatchInsertRequest instance using NewBatchInsertRequestWithEntityNameManagedObjectHandler.
// Creates a batch-insertion request for a named managed entity, and specifies a closure that inserts data into the entity.
func ExampleNewBatchInsertRequestWithEntityNameManagedObjectHandler() {
	_ = coredata.NewBatchInsertRequestWithEntityNameManagedObjectHandler(
		"entityName", // entityName string
		nil, // handler unsafe.Pointer
	)
	// Output:
}

// ExampleNewBatchInsertRequestWithEntityManagedObjectHandler demonstrates how to create a BatchInsertRequest instance using NewBatchInsertRequestWithEntityManagedObjectHandler.
// Creates a batch-insertion request for a managed entity, and specifies a closure that inserts data into the entity.
func ExampleNewBatchInsertRequestWithEntityManagedObjectHandler() {
	_ = coredata.NewBatchInsertRequestWithEntityManagedObjectHandler(
		nil, // entity unsafe.Pointer
		nil, // handler unsafe.Pointer
	)
	// Output:
}


