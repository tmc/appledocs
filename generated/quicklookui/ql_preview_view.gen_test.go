// Code generated from Apple documentation for QuickLookUI. DO NOT EDIT.

package quicklookui_test

import (
	"github.com/tmc/appledocs/generated/quicklookui"
)

// Suppress unused import errors
var _ = quicklookui.NewPreviewView

// ExampleNewPreviewViewWithFrame demonstrates how to create a PreviewView instance using NewPreviewViewWithFrame.
// Creates a preview view with the provided frame.
func ExampleNewPreviewViewWithFrame() {
	_ = quicklookui.NewPreviewViewWithFrame(
		quicklookui.Rect /* not a class type */{}, // frame Rect /* not a class type */
	)
	// Output:
}
// ExampleNewPreviewViewWithFrameStyle demonstrates how to create a PreviewView instance using NewPreviewViewWithFrameStyle.
// Creates a preview view with the provided frame and style.
func ExampleNewPreviewViewWithFrameStyle() {
	_ = quicklookui.NewPreviewViewWithFrameStyle(
		quicklookui.Rect /* not a class type */{}, // frame Rect /* not a class type */
		quicklookui.PreviewViewStyle{}, // style PreviewViewStyle
	)
	// Output:
}
// ExamplePreviewView_Close demonstrates using Close on a PreviewView instance.
// Closes the view, releasing the current preview item.
func ExamplePreviewView_Close() {
	obj := quicklookui.NewPreviewView()
	obj.Close()
	// Output:
	}

// ExamplePreviewView_RefreshPreviewItem demonstrates using RefreshPreviewItem on a PreviewView instance.
// Updates the preview to display the currently previewed item.
func ExamplePreviewView_RefreshPreviewItem() {
	obj := quicklookui.NewPreviewView()
	obj.RefreshPreviewItem()
	// Output:
	}



