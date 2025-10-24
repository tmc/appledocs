// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewCollectionViewDiffableDataSource

// ExampleCollectionViewDiffableDataSource_Snapshot demonstrates using Snapshot on a CollectionViewDiffableDataSource instance.
// Returns a representation of the current state of the data in the collection view.
func ExampleCollectionViewDiffableDataSource_Snapshot() {
	obj := appkit.NewCollectionViewDiffableDataSource()
	_ = obj.Snapshot()
	// Output:
	}

