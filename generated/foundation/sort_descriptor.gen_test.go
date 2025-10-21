// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewSortDescriptor


// ExampleNewSortDescriptorWithKeyAscending demonstrates how to create a SortDescriptor instance using NewSortDescriptorWithKeyAscending.
// Creates a sort descriptor with a specified string key path and sort order.
func ExampleNewSortDescriptorWithKeyAscending() {
	_ = foundation.NewSortDescriptorWithKeyAscending(
		"key", // key string
		false, // ascending bool
	)
	// Output:
}


// ExampleNewSortDescriptorWithKeyAscendingSelector demonstrates how to create a SortDescriptor instance using NewSortDescriptorWithKeyAscendingSelector.
// Creates a sort descriptor with a specified string key path, ordering, and comparison selector.
func ExampleNewSortDescriptorWithKeyAscendingSelector() {
	_ = foundation.NewSortDescriptorWithKeyAscendingSelector(
		"key", // key string
		false, // ascending bool
		0, // selector objc.SEL
	)
	// Output:
}



