// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewCollectionViewDiffableDataSource

// ExampleNewCollectionViewDiffableDataSourceWithCollectionViewItemProvider demonstrates how to create a CollectionViewDiffableDataSource instance using NewCollectionViewDiffableDataSourceWithCollectionViewItemProvider.
// Creates a diffable data source with the specified item provider, and connects it to the specified collection view.
func ExampleNewCollectionViewDiffableDataSourceWithCollectionViewItemProvider() {
	_ = appkit.NewCollectionViewDiffableDataSourceWithCollectionViewItemProvider(
		appkit.CollectionView /* not a class type */{}, // collectionView CollectionView /* not a class type */
		appkit.CollectionViewDiffableDataSourceItemProvider /* not a class type */{}, // itemProvider CollectionViewDiffableDataSourceItemProvider /* not a class type */
	)
	// Output:
}
// ExampleCollectionViewDiffableDataSource_Snapshot demonstrates using Snapshot on a CollectionViewDiffableDataSource instance.
// Returns a representation of the current state of the data in the collection view.
func ExampleCollectionViewDiffableDataSource_Snapshot() {
	obj := appkit.NewCollectionViewDiffableDataSource()
	_ = obj.Snapshot()
	// Output:
	}

