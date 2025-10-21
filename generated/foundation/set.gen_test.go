// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewSet

// ExampleNewSet demonstrates how to create a Set instance.
// Initializes a newly allocated set.
func ExampleNewSet() {
	_ = foundation.NewSet()
	// Output:
}
// ExampleNewSetWithCoder demonstrates how to create a Set instance using NewSetWithCoder.
func ExampleNewSetWithCoder() {
	_ = foundation.NewSetWithCoder(
		foundation.NSCoder{}, // coder NSCoder
	)
	// Output:
}
// ExampleNewSetWithCollectionViewIndexPath demonstrates how to create a Set instance using NewSetWithCollectionViewIndexPath.
func ExampleNewSetWithCollectionViewIndexPath() {
	_ = foundation.NewSetWithCollectionViewIndexPath(
		foundation.NSIndexPath{}, // indexPath NSIndexPath
	)
	// Output:
}
// ExampleNewSetWithCollectionViewIndexPaths demonstrates how to create a Set instance using NewSetWithCollectionViewIndexPaths.
func ExampleNewSetWithCollectionViewIndexPaths() {
	_ = foundation.NewSetWithCollectionViewIndexPaths(
		[]foundation.IndexPath{}, // indexPaths []IndexPath
	)
	// Output:
}
