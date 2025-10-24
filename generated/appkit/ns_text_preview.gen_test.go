// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewTextPreview

// ExampleNewTextPreviewWithSnapshotImagePresentationFrame demonstrates how to create a TextPreview instance using NewTextPreviewWithSnapshotImagePresentationFrame.
// Creates a text preview using the specified image.
func ExampleNewTextPreviewWithSnapshotImagePresentationFrame() {
	_ = appkit.NewTextPreviewWithSnapshotImagePresentationFrame(
		appkit.ImageRef /* not a class type */{}, // snapshotImage ImageRef /* not a class type */
		appkit.Rect /* not a class type */{}, // presentationFrame Rect /* not a class type */
	)
	// Output:
}
