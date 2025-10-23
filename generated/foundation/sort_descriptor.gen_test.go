// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewSortDescriptor

// ExampleNewSortDescriptorWithCoder demonstrates how to create a SortDescriptor instance using NewSortDescriptorWithCoder.
// Creates a sort descriptor by decoding from the coder you specify.
func ExampleNewSortDescriptorWithCoder() {
	_ = foundation.NewSortDescriptorWithCoder(
		foundation.NSCoder{}, // coder NSCoder
	)
	// Output:
}
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
