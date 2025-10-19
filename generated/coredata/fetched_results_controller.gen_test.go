// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata_test

import (
	"github.com/tmc/appledocs/generated/coredata"
)


// ExampleNewFetchedResultsControllerWithFetchRequestManagedObjectContextSectionNameKeyPathCacheName demonstrates how to create a FetchedResultsController instance using NewFetchedResultsControllerWithFetchRequestManagedObjectContextSectionNameKeyPathCacheName.
// Returns a fetch request controller initialized using the given arguments.
func ExampleNewFetchedResultsControllerWithFetchRequestManagedObjectContextSectionNameKeyPathCacheName() {
	_ = coredata.NewFetchedResultsControllerWithFetchRequestManagedObjectContextSectionNameKeyPathCacheName(
		nil, // fetchRequest unsafe.Pointer
		nil, // context unsafe.Pointer
		"sectionNameKeyPath", // sectionNameKeyPath string
		"name", // name string
	)
	// Output:
}


