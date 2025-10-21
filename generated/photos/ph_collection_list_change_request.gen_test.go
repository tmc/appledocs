// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos_test

import (
	"github.com/tmc/appledocs/generated/photos"
)

// Suppress unused import errors
var _ = photos.NewPHCollectionListChangeRequest

// ExampleNewPHCollectionListChangeRequestForCollectionList demonstrates how to create a PHCollectionListChangeRequest instance using NewPHCollectionListChangeRequestForCollectionList.
// Creates a request for modifying the specified collection list.
func ExampleNewPHCollectionListChangeRequestForCollectionList() {
	_ = photos.NewPHCollectionListChangeRequestForCollectionList(
		photos.PHCollectionList{}, // collectionList PHCollectionList
	)
	// Output:
}
