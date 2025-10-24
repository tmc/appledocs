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
// ExampleNewSetWithCollectionViewIndexPaths demonstrates how to create a Set instance using NewSetWithCollectionViewIndexPaths.
func ExampleNewSetWithCollectionViewIndexPaths() {
	_ = foundation.NewSetWithCollectionViewIndexPaths(
		[]foundation.IIndexPath{}, // indexPaths []IIndexPath
	)
	// Output:
}
// ExampleSet_AnyObject demonstrates using AnyObject on a Set instance.
// Returns one of the objects in the set, or   if the set contains no objects.
func ExampleSet_AnyObject() {
	obj := foundation.NewSet()
	_ = obj.AnyObject()
	// Output:
	}

// ExampleSet_ObjectEnumerator demonstrates using ObjectEnumerator on a Set instance.
// Returns an enumerator object that lets you access each object in the set.
func ExampleSet_ObjectEnumerator() {
	obj := foundation.NewSet()
	_ = obj.ObjectEnumerator()
	// Output:
	}

