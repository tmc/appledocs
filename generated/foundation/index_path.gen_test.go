// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewIndexPath

// ExampleNewIndexPathForItemInSection demonstrates how to create a IndexPath instance using NewIndexPathForItemInSection.
// Initializes an index path with the indexes of a specific item and section in a collection view.
func ExampleNewIndexPathForItemInSection() {
	_ = foundation.NewIndexPathForItemInSection(
		0, // item int
		0, // section int
	)
	// Output:
}
// ExampleNewIndexPathForRowInSection demonstrates how to create a IndexPath instance using NewIndexPathForRowInSection.
// Initializes an index path with the indexes of a specific row and section in a table view.
func ExampleNewIndexPathForRowInSection() {
	_ = foundation.NewIndexPathForRowInSection(
		0, // row int
		0, // section int
	)
	// Output:
}
// ExampleNewIndexPathWithIndex demonstrates how to create a IndexPath instance using NewIndexPathWithIndex.
// Initializes an index path with a single node.
func ExampleNewIndexPathWithIndex() {
	_ = foundation.NewIndexPathWithIndex(
		0, // index uint
	)
	// Output:
}
// ExampleNewIndexPathWithIndexesLength demonstrates how to create a IndexPath instance using NewIndexPathWithIndexesLength.
// Initializes an index path with the given nodes and length.
func ExampleNewIndexPathWithIndexesLength() {
	_ = foundation.NewIndexPathWithIndexesLength(
		[]foundation.uint{}, // indexes []uint
		0, // length uint
	)
	// Output:
}
