// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)


// ExampleNewWindowWithWindowRef demonstrates how to create a Window instance using NewWindowWithWindowRef.
// Returns a Cocoa window created from a Carbon window.
func ExampleNewWindowWithWindowRef() {
	_ = appkit.NewWindowWithWindowRef(
		nil, // windowRef unsafe.Pointer
	)
	// Output:
}

// ExampleNewWindowWithContentRectStyleMaskBackingDefer demonstrates how to create a Window instance using NewWindowWithContentRectStyleMaskBackingDefer.
// Initializes the window with the specified values.
func ExampleNewWindowWithContentRectStyleMaskBackingDefer() {
	_ = appkit.NewWindowWithContentRectStyleMaskBackingDefer(
		nil, // contentRect unsafe.Pointer
		appkit.WindowStyleMask(0), // style WindowStyleMask
		appkit.BackingStoreType(0), // backingStoreType BackingStoreType
		false, // flag bool
	)
	// Output:
}

// ExampleNewWindowWithContentRectStyleMaskBackingDeferScreen demonstrates how to create a Window instance using NewWindowWithContentRectStyleMaskBackingDeferScreen.
// Initializes an allocated window with the specified values.
func ExampleNewWindowWithContentRectStyleMaskBackingDeferScreen() {
	_ = appkit.NewWindowWithContentRectStyleMaskBackingDeferScreen(
		nil, // contentRect unsafe.Pointer
		appkit.WindowStyleMask(0), // style WindowStyleMask
		appkit.BackingStoreType(0), // backingStoreType BackingStoreType
		false, // flag bool
		nil, // screen unsafe.Pointer
	)
	// Output:
}

// ExampleNewWindowWithContentViewController demonstrates how to create a Window instance using NewWindowWithContentViewController.
// Creates a titled window that contains the specified content view controller.
func ExampleNewWindowWithContentViewController() {
	_ = appkit.NewWindowWithContentViewController(
		nil, // contentViewController unsafe.Pointer
	)
	// Output:
}


