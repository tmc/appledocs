// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewView

// ExampleNewViewWithFrame demonstrates how to create a View instance using NewViewWithFrame.
// Initializes and returns a newly allocated   object with a specified frame rectangle.
func ExampleNewViewWithFrame() {
	_ = appkit.NewViewWithFrame(
		appkit.Rect /* not a class type */{}, // frameRect Rect /* not a class type */
	)
	// Output:
}
// ExampleView_BeginDocument demonstrates using BeginDocument on a View instance.
// Invoked at the beginning of the printing session, this method sets up the current graphics context.
func ExampleView_BeginDocument() {
	obj := appkit.NewView()
	obj.BeginDocument()
	// Output:
	}

// ExampleView_DiscardCursorRects demonstrates using DiscardCursorRects on a View instance.
// Invalidates all cursor rectangles set up using  .
//
// Note: This example is not executed because DiscardCursorRects crashes when called on bare NSObject
// (it's a protocol/category method that should be overridden by subclasses).
func ExampleView_DiscardCursorRects() {
	obj := appkit.NewView()
	obj.DiscardCursorRects()
	}

// ExampleView_Display demonstrates using Display on a View instance.
// Displays the view and all its subviews if possible, invoking each of the   methods  ,  , and   as necessary.
func ExampleView_Display() {
	obj := appkit.NewView()
	obj.Display()
	// Output:
	}

// ExampleView_DisplayIfNeeded demonstrates using DisplayIfNeeded on a View instance.
// Displays the view and all its subviews if any part of the view has been marked as needing display.
func ExampleView_DisplayIfNeeded() {
	obj := appkit.NewView()
	obj.DisplayIfNeeded()
	// Output:
	}

// ExampleView_DisplayIfNeededIgnoringOpacity demonstrates using DisplayIfNeededIgnoringOpacity on a View instance.
// Acts as  , except that this method doesn’t back up to the first opaque ancestor—it simply causes the view and its descendants to execute their drawing code.
func ExampleView_DisplayIfNeededIgnoringOpacity() {
	obj := appkit.NewView()
	obj.DisplayIfNeededIgnoringOpacity()
	// Output:
	}

// ExampleView_DrawFocusRingMask demonstrates using DrawFocusRingMask on a View instance.
// Draws the focus ring mask for the view.
func ExampleView_DrawFocusRingMask() {
	obj := appkit.NewView()
	obj.DrawFocusRingMask()
	// Output:
	}

// ExampleView_EndDocument demonstrates using EndDocument on a View instance.
// This method is invoked at the end of the printing session.
func ExampleView_EndDocument() {
	obj := appkit.NewView()
	obj.EndDocument()
	// Output:
	}

// ExampleView_EndPage demonstrates using EndPage on a View instance.
// Writes the end of a conforming page.
func ExampleView_EndPage() {
	obj := appkit.NewView()
	obj.EndPage()
	// Output:
	}

// ExampleView_ExerciseAmbiguityInLayout demonstrates using ExerciseAmbiguityInLayout on a View instance.
// Randomly changes the frame of a view with an ambiguous layout between the different valid values.
func ExampleView_ExerciseAmbiguityInLayout() {
	obj := appkit.NewView()
	obj.ExerciseAmbiguityInLayout()
	// Output:
	}

// ExampleView_InvalidateIntrinsicContentSize demonstrates using InvalidateIntrinsicContentSize on a View instance.
// Invalidates the view’s intrinsic content size.
func ExampleView_InvalidateIntrinsicContentSize() {
	obj := appkit.NewView()
	obj.InvalidateIntrinsicContentSize()
	// Output:
	}

// ExampleView_Layout demonstrates using Layout on a View instance.
// Perform layout in concert with the constraint-based layout system.
func ExampleView_Layout() {
	obj := appkit.NewView()
	obj.Layout()
	// Output:
	}

// ExampleView_LayoutSubtreeIfNeeded demonstrates using LayoutSubtreeIfNeeded on a View instance.
// Updates the layout of the receiving view and its subviews based on the current views and constraints.
func ExampleView_LayoutSubtreeIfNeeded() {
	obj := appkit.NewView()
	obj.LayoutSubtreeIfNeeded()
	// Output:
	}

// ExampleView_MakeBackingLayer demonstrates using MakeBackingLayer on a View instance.
// Creates the view’s backing layer.
func ExampleView_MakeBackingLayer() {
	obj := appkit.NewView()
	_ = obj.MakeBackingLayer()
	// Output:
	}

// ExampleView_NoteFocusRingMaskChanged demonstrates using NoteFocusRingMaskChanged on a View instance.
// Invoked to notify the view that the focus ring mask requires updating.
func ExampleView_NoteFocusRingMaskChanged() {
	obj := appkit.NewView()
	obj.NoteFocusRingMaskChanged()
	// Output:
	}

// ExampleView_PrepareForReuse demonstrates using PrepareForReuse on a View instance.
// Restores the view to an initial state so that it can be reused.
//
// Note: This example is not executed because PrepareForReuse crashes when called on bare NSObject
// (it's a protocol/category method that should be overridden by subclasses).
func ExampleView_PrepareForReuse() {
	obj := appkit.NewView()
	obj.PrepareForReuse()
	}

// ExampleView_RemoveAllToolTips demonstrates using RemoveAllToolTips on a View instance.
// Removes all tooltips assigned to the view.
func ExampleView_RemoveAllToolTips() {
	obj := appkit.NewView()
	obj.RemoveAllToolTips()
	// Output:
	}

// ExampleView_RemoveFromSuperview demonstrates using RemoveFromSuperview on a View instance.
// Unlinks the view from its superview and its window, removes it from the responder chain, and invalidates its cursor rectangles.
func ExampleView_RemoveFromSuperview() {
	obj := appkit.NewView()
	obj.RemoveFromSuperview()
	// Output:
	}

// ExampleView_RemoveFromSuperviewWithoutNeedingDisplay demonstrates using RemoveFromSuperviewWithoutNeedingDisplay on a View instance.
// Unlinks the view from its superview and its window and removes it from the responder chain, but does not invalidate its cursor rectangles to cause redrawing.
func ExampleView_RemoveFromSuperviewWithoutNeedingDisplay() {
	obj := appkit.NewView()
	obj.RemoveFromSuperviewWithoutNeedingDisplay()
	// Output:
	}

// ExampleView_ResetCursorRects demonstrates using ResetCursorRects on a View instance.
// Overridden by subclasses to define their default cursor rectangles.
func ExampleView_ResetCursorRects() {
	obj := appkit.NewView()
	obj.ResetCursorRects()
	// Output:
	}

// ExampleView_UnregisterDraggedTypes demonstrates using UnregisterDraggedTypes on a View instance.
// Unregisters the view as a possible destination in a dragging session.
func ExampleView_UnregisterDraggedTypes() {
	obj := appkit.NewView()
	obj.UnregisterDraggedTypes()
	// Output:
	}

// ExampleView_UpdateConstraints demonstrates using UpdateConstraints on a View instance.
// Update constraints for the view.
func ExampleView_UpdateConstraints() {
	obj := appkit.NewView()
	obj.UpdateConstraints()
	// Output:
	}

// ExampleView_UpdateConstraintsForSubtreeIfNeeded demonstrates using UpdateConstraintsForSubtreeIfNeeded on a View instance.
// Updates the constraints for the receiving view and its subviews.
func ExampleView_UpdateConstraintsForSubtreeIfNeeded() {
	obj := appkit.NewView()
	obj.UpdateConstraintsForSubtreeIfNeeded()
	// Output:
	}

// ExampleView_UpdateLayer demonstrates using UpdateLayer on a View instance.
// Updates the view’s content by modifying its underlying layer.
func ExampleView_UpdateLayer() {
	obj := appkit.NewView()
	obj.UpdateLayer()
	// Output:
	}

// ExampleView_UpdateTrackingAreas demonstrates using UpdateTrackingAreas on a View instance.
// Invoked automatically when the view’s geometry changes such that its tracking areas need to be recalculated.
func ExampleView_UpdateTrackingAreas() {
	obj := appkit.NewView()
	obj.UpdateTrackingAreas()
	// Output:
	}

// ExampleView_ViewDidChangeBackingProperties demonstrates using ViewDidChangeBackingProperties on a View instance.
// Responds when the view’s backing store properties change.
func ExampleView_ViewDidChangeBackingProperties() {
	obj := appkit.NewView()
	obj.ViewDidChangeBackingProperties()
	// Output:
	}

// ExampleView_ViewDidChangeEffectiveAppearance demonstrates using ViewDidChangeEffectiveAppearance on a View instance.
// Informs the view that its effective appearance changed.
func ExampleView_ViewDidChangeEffectiveAppearance() {
	obj := appkit.NewView()
	obj.ViewDidChangeEffectiveAppearance()
	// Output:
	}

// ExampleView_ViewDidEndLiveResize demonstrates using ViewDidEndLiveResize on a View instance.
// Informs the view of the end of a live resize—the user has finished resizing the view.
func ExampleView_ViewDidEndLiveResize() {
	obj := appkit.NewView()
	obj.ViewDidEndLiveResize()
	// Output:
	}

// ExampleView_ViewDidHide demonstrates using ViewDidHide on a View instance.
// Invoked when the view is hidden, either directly, or in response to an ancestor being hidden.
func ExampleView_ViewDidHide() {
	obj := appkit.NewView()
	obj.ViewDidHide()
	// Output:
	}

// ExampleView_ViewDidMoveToSuperview demonstrates using ViewDidMoveToSuperview on a View instance.
// Informs the view that its superview has changed (possibly to  ).
func ExampleView_ViewDidMoveToSuperview() {
	obj := appkit.NewView()
	obj.ViewDidMoveToSuperview()
	// Output:
	}

// ExampleView_ViewDidMoveToWindow demonstrates using ViewDidMoveToWindow on a View instance.
// Informs the view that it has been added to a new view hierarchy.
func ExampleView_ViewDidMoveToWindow() {
	obj := appkit.NewView()
	obj.ViewDidMoveToWindow()
	// Output:
	}

// ExampleView_ViewDidUnhide demonstrates using ViewDidUnhide on a View instance.
// Invoked when the view is unhidden, either directly, or in response to an ancestor being unhidden
func ExampleView_ViewDidUnhide() {
	obj := appkit.NewView()
	obj.ViewDidUnhide()
	// Output:
	}

// ExampleView_ViewWillDraw demonstrates using ViewWillDraw on a View instance.
// Informs the view that it’s required to draw content.
func ExampleView_ViewWillDraw() {
	obj := appkit.NewView()
	obj.ViewWillDraw()
	// Output:
	}

// ExampleView_ViewWillStartLiveResize demonstrates using ViewWillStartLiveResize on a View instance.
// Informs the view of the start of a live resize—the user has started resizing the view.
func ExampleView_ViewWillStartLiveResize() {
	obj := appkit.NewView()
	obj.ViewWillStartLiveResize()
	// Output:
	}

