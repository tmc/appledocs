// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewWindow

// ExampleWindow_BecomeKeyWindow demonstrates using BecomeKeyWindow on a Window instance.
// Informs the window that it has become the key window.
func ExampleWindow_BecomeKeyWindow() {
	obj := appkit.NewWindow()
	obj.BecomeKeyWindow()
	// Output:
	}

// ExampleWindow_BecomeMainWindow demonstrates using BecomeMainWindow on a Window instance.
// Informs the window that it has become the main window.
func ExampleWindow_BecomeMainWindow() {
	obj := appkit.NewWindow()
	obj.BecomeMainWindow()
	// Output:
	}

// ExampleWindow_Center demonstrates using Center on a Window instance.
// Sets the window’s location to the center of the screen.
func ExampleWindow_Center() {
	obj := appkit.NewWindow()
	obj.Center()
	// Output:
	}

// ExampleWindow_Close demonstrates using Close on a Window instance.
// Removes the window from the screen.
func ExampleWindow_Close() {
	obj := appkit.NewWindow()
	obj.Close()
	// Output:
	}

// ExampleWindow_DisableCursorRects demonstrates using DisableCursorRects on a Window instance.
// Disables all cursor rectangle management within the window.
func ExampleWindow_DisableCursorRects() {
	obj := appkit.NewWindow()
	obj.DisableCursorRects()
	// Output:
	}

// ExampleWindow_DiscardCursorRects demonstrates using DiscardCursorRects on a Window instance.
// Invalidates all cursor rectangles in the window.
//
// Note: This example is not executed because DiscardCursorRects crashes when called on bare NSObject
// (it's a protocol/category method that should be overridden by subclasses).
func ExampleWindow_DiscardCursorRects() {
	obj := appkit.NewWindow()
	obj.DiscardCursorRects()
	}

// ExampleWindow_Display demonstrates using Display on a Window instance.
// Passes a display message down the window’s view hierarchy, thus redrawing all views within the window.
func ExampleWindow_Display() {
	obj := appkit.NewWindow()
	obj.Display()
	// Output:
	}

// ExampleWindow_DisplayIfNeeded demonstrates using DisplayIfNeeded on a Window instance.
// Passes a display message down the window’s view hierarchy, thus redrawing all views that need displaying.
func ExampleWindow_DisplayIfNeeded() {
	obj := appkit.NewWindow()
	obj.DisplayIfNeeded()
	// Output:
	}

// ExampleWindow_EnableCursorRects demonstrates using EnableCursorRects on a Window instance.
// Reenables cursor rectangle management within the window after a   message.
func ExampleWindow_EnableCursorRects() {
	obj := appkit.NewWindow()
	obj.EnableCursorRects()
	// Output:
	}

// ExampleWindow_InvalidateShadow demonstrates using InvalidateShadow on a Window instance.
// Invalidates the window shadow so that it is recomputed based on the current window shape.
func ExampleWindow_InvalidateShadow() {
	obj := appkit.NewWindow()
	obj.InvalidateShadow()
	// Output:
	}

// ExampleWindow_MakeKeyWindow demonstrates using MakeKeyWindow on a Window instance.
// Makes the window the key window.
func ExampleWindow_MakeKeyWindow() {
	obj := appkit.NewWindow()
	obj.MakeKeyWindow()
	// Output:
	}

// ExampleWindow_MakeMainWindow demonstrates using MakeMainWindow on a Window instance.
// Makes the window the main window.
func ExampleWindow_MakeMainWindow() {
	obj := appkit.NewWindow()
	obj.MakeMainWindow()
	// Output:
	}

// ExampleWindow_OrderFrontRegardless demonstrates using OrderFrontRegardless on a Window instance.
// Moves the window to the front of its level, even if its application isn’t active, without changing either the key window or the main window.
func ExampleWindow_OrderFrontRegardless() {
	obj := appkit.NewWindow()
	obj.OrderFrontRegardless()
	// Output:
	}

// ExampleWindow_RecalculateKeyViewLoop demonstrates using RecalculateKeyViewLoop on a Window instance.
// Marks the key view loop as “dirty” and in need of recalculation.
func ExampleWindow_RecalculateKeyViewLoop() {
	obj := appkit.NewWindow()
	obj.RecalculateKeyViewLoop()
	// Output:
	}

// ExampleWindow_ResetCursorRects demonstrates using ResetCursorRects on a Window instance.
// Clears the window’s cursor rectangles and the cursor rectangles of the   objects in its view hierarchy.
func ExampleWindow_ResetCursorRects() {
	obj := appkit.NewWindow()
	obj.ResetCursorRects()
	// Output:
	}

// ExampleWindow_ResignKeyWindow demonstrates using ResignKeyWindow on a Window instance.
// Resigns the window’s key window status.
func ExampleWindow_ResignKeyWindow() {
	obj := appkit.NewWindow()
	obj.ResignKeyWindow()
	// Output:
	}

// ExampleWindow_ResignMainWindow demonstrates using ResignMainWindow on a Window instance.
// Resigns the window’s main window status.
func ExampleWindow_ResignMainWindow() {
	obj := appkit.NewWindow()
	obj.ResignMainWindow()
	// Output:
	}

// ExampleWindow_UnregisterDraggedTypes demonstrates using UnregisterDraggedTypes on a Window instance.
// Unregisters the window as a possible destination for dragging operations.
func ExampleWindow_UnregisterDraggedTypes() {
	obj := appkit.NewWindow()
	obj.UnregisterDraggedTypes()
	// Output:
	}

