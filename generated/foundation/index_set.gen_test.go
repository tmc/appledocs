// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewIndexSet

// ExampleNewIndexSetWithIndex demonstrates how to create a IndexSet instance using NewIndexSetWithIndex.
// Initializes an allocated   object with an index.
func ExampleNewIndexSetWithIndex() {
	_ = foundation.NewIndexSetWithIndex(
		0, // value uint
	)
	// Output:
}
// ExampleNewIndexSetWithIndexSet demonstrates how to create a IndexSet instance using NewIndexSetWithIndexSet.
// Initializes an allocated   object with an index set.
func ExampleNewIndexSetWithIndexSet() {
	_ = foundation.NewIndexSetWithIndexSet(
		foundation.NSIndexSet{}, // indexSet NSIndexSet
	)
	// Output:
}
// ExampleNewIndexSetWithIndexesInRange demonstrates how to create a IndexSet instance using NewIndexSetWithIndexesInRange.
// Initializes an allocated   object with an index range.
func ExampleNewIndexSetWithIndexesInRange() {
	_ = foundation.NewIndexSetWithIndexesInRange(
		foundation.Range{}, // range Range
	)
	// Output:
}
