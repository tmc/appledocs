// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewTextView

// ExampleNewTextViewUsingTextLayoutManager demonstrates how to create a TextView instance using NewTextViewUsingTextLayoutManager.
func ExampleNewTextViewUsingTextLayoutManager() {
	_ = appkit.NewTextViewUsingTextLayoutManager(
		false, // usingTextLayoutManager bool
	)
	// Output:
}
// ExampleNewTextViewWithFrame demonstrates how to create a TextView instance using NewTextViewWithFrame.
// Initializes a text view.
func ExampleNewTextViewWithFrame() {
	_ = appkit.NewTextViewWithFrame(
		appkit.Rect /* not a class type */{}, // frameRect Rect /* not a class type */
	)
	// Output:
}
// ExampleTextView_BreakUndoCoalescing demonstrates using BreakUndoCoalescing on a TextView instance.
// Informs the receiver that it should begin coalescing successive typing operations in a new undo grouping.
func ExampleTextView_BreakUndoCoalescing() {
	obj := appkit.NewTextView()
	obj.BreakUndoCoalescing()
	// Output:
	}

// ExampleTextView_CleanUpAfterDragOperation demonstrates using CleanUpAfterDragOperation on a TextView instance.
// Releases the drag information still existing after the dragging session has completed.
func ExampleTextView_CleanUpAfterDragOperation() {
	obj := appkit.NewTextView()
	obj.CleanUpAfterDragOperation()
	// Output:
	}

// ExampleTextView_DidChangeText demonstrates using DidChangeText on a TextView instance.
// Sends out necessary notifications when a text change completes.
func ExampleTextView_DidChangeText() {
	obj := appkit.NewTextView()
	obj.DidChangeText()
	// Output:
	}

// ExampleTextView_InvalidateTextContainerOrigin demonstrates using InvalidateTextContainerOrigin on a TextView instance.
// Invalidates the calculated origin of the text container.
func ExampleTextView_InvalidateTextContainerOrigin() {
	obj := appkit.NewTextView()
	obj.InvalidateTextContainerOrigin()
	// Output:
	}

// ExampleTextView_UpdateCandidates demonstrates using UpdateCandidates on a TextView instance.
func ExampleTextView_UpdateCandidates() {
	obj := appkit.NewTextView()
	obj.UpdateCandidates()
	// Output:
	}

// ExampleTextView_UpdateDragTypeRegistration demonstrates using UpdateDragTypeRegistration on a TextView instance.
// Updates the acceptable drag types of all text views associated with the receiver’s layout manager.
func ExampleTextView_UpdateDragTypeRegistration() {
	obj := appkit.NewTextView()
	obj.UpdateDragTypeRegistration()
	// Output:
	}

// ExampleTextView_UpdateFontPanel demonstrates using UpdateFontPanel on a TextView instance.
// Updates the Font panel to contain the font attributes of the selection.
func ExampleTextView_UpdateFontPanel() {
	obj := appkit.NewTextView()
	obj.UpdateFontPanel()
	// Output:
	}

// ExampleTextView_UpdateQuickLookPreviewPanel demonstrates using UpdateQuickLookPreviewPanel on a TextView instance.
// Notifies the QuickLook panel that an update may be required.
func ExampleTextView_UpdateQuickLookPreviewPanel() {
	obj := appkit.NewTextView()
	obj.UpdateQuickLookPreviewPanel()
	// Output:
	}

// ExampleTextView_UpdateRuler demonstrates using UpdateRuler on a TextView instance.
// Updates the ruler view in the receiver’s enclosing scroll view to reflect the selection’s paragraph and marker attributes.
func ExampleTextView_UpdateRuler() {
	obj := appkit.NewTextView()
	obj.UpdateRuler()
	// Output:
	}

// ExampleTextView_UpdateTextTouchBarItems demonstrates using UpdateTextTouchBarItems on a TextView instance.
func ExampleTextView_UpdateTextTouchBarItems() {
	obj := appkit.NewTextView()
	obj.UpdateTextTouchBarItems()
	// Output:
	}

// ExampleTextView_UpdateTouchBarItemIdentifiers demonstrates using UpdateTouchBarItemIdentifiers on a TextView instance.
func ExampleTextView_UpdateTouchBarItemIdentifiers() {
	obj := appkit.NewTextView()
	obj.UpdateTouchBarItemIdentifiers()
	// Output:
	}

