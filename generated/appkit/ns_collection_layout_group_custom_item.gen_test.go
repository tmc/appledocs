// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewCollectionLayoutGroupCustomItem

// ExampleNewCollectionLayoutGroupCustomItemWithFrame demonstrates how to create a CollectionLayoutGroupCustomItem instance using NewCollectionLayoutGroupCustomItemWithFrame.
// Creates a custom item with the specified frame.
func ExampleNewCollectionLayoutGroupCustomItemWithFrame() {
	_ = appkit.NewCollectionLayoutGroupCustomItemWithFrame(
		appkit.Rect /* not a class type */{}, // frame Rect /* not a class type */
	)
	// Output:
}
// ExampleNewCollectionLayoutGroupCustomItemWithFrameZIndex demonstrates how to create a CollectionLayoutGroupCustomItem instance using NewCollectionLayoutGroupCustomItemWithFrameZIndex.
// Creates a custom item with the specified frame and vertical stacking order in relation to other items in the group.
func ExampleNewCollectionLayoutGroupCustomItemWithFrameZIndex() {
	_ = appkit.NewCollectionLayoutGroupCustomItemWithFrameZIndex(
		appkit.Rect /* not a class type */{}, // frame Rect /* not a class type */
		0, // zIndex int
	)
	// Output:
}
