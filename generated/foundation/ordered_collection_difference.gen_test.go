// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewOrderedCollectionDifference

// ExampleNewOrderedCollectionDifferenceWithChanges demonstrates how to create a OrderedCollectionDifference instance using NewOrderedCollectionDifferenceWithChanges.
// Creates an ordered collection difference using an array of ordered collection changes.
func ExampleNewOrderedCollectionDifferenceWithChanges() {
	_ = foundation.NewOrderedCollectionDifferenceWithChanges(
		[]foundation.OrderedCollectionChange{}, // changes []OrderedCollectionChange
	)
	// Output:
}
