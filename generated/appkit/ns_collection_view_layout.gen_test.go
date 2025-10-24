// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewCollectionViewLayout

// ExampleCollectionViewLayout_FinalizeAnimatedBoundsChange demonstrates using FinalizeAnimatedBoundsChange on a CollectionViewLayout instance.
// Cleans up after any animated changes to the collection view’s bounds or after the insertion or deletion of items.
func ExampleCollectionViewLayout_FinalizeAnimatedBoundsChange() {
	obj := appkit.NewCollectionViewLayout()
	obj.FinalizeAnimatedBoundsChange()
	// Output:
	}

// ExampleCollectionViewLayout_FinalizeCollectionViewUpdates demonstrates using FinalizeCollectionViewUpdates on a CollectionViewLayout instance.
// Performs needed steps after items are inserted, deleted, or moved within a collection view.
func ExampleCollectionViewLayout_FinalizeCollectionViewUpdates() {
	obj := appkit.NewCollectionViewLayout()
	obj.FinalizeCollectionViewUpdates()
	// Output:
	}

// ExampleCollectionViewLayout_FinalizeLayoutTransition demonstrates using FinalizeLayoutTransition on a CollectionViewLayout instance.
// Performs any final steps related to a layout transition before the transition animations actually occur.
func ExampleCollectionViewLayout_FinalizeLayoutTransition() {
	obj := appkit.NewCollectionViewLayout()
	obj.FinalizeLayoutTransition()
	// Output:
	}

// ExampleCollectionViewLayout_InvalidateLayout demonstrates using InvalidateLayout on a CollectionViewLayout instance.
// Invalidates all layout information and triggers a layout update.
func ExampleCollectionViewLayout_InvalidateLayout() {
	obj := appkit.NewCollectionViewLayout()
	obj.InvalidateLayout()
	// Output:
	}

// ExampleCollectionViewLayout_PrepareLayout demonstrates using PrepareLayout on a CollectionViewLayout instance.
// Prepares the layout object to begin laying out content.
//
// Note: This example is not executed because PrepareLayout crashes when called on bare NSObject
// (it's a protocol/category method that should be overridden by subclasses).
func ExampleCollectionViewLayout_PrepareLayout() {
	obj := appkit.NewCollectionViewLayout()
	obj.PrepareLayout()
	}

