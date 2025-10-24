// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewSplitViewItem

// ExampleNewSplitViewItemContentListWithViewController demonstrates how to create a SplitViewItem instance using NewSplitViewItemContentListWithViewController.
// Creates a split view item that represents a content list for the specified view controller.
func ExampleNewSplitViewItemContentListWithViewController() {
	_ = appkit.NewSplitViewItemContentListWithViewController(
		appkit.NSViewController{}, // viewController NSViewController
	)
	// Output:
}

// ExampleNewSplitViewItemInspectorWithViewController demonstrates how to create a SplitViewItem instance using NewSplitViewItemInspectorWithViewController.
func ExampleNewSplitViewItemInspectorWithViewController() {
	_ = appkit.NewSplitViewItemInspectorWithViewController(
		appkit.NSViewController{}, // viewController NSViewController
	)
	// Output:
}

// ExampleNewSplitViewItemSidebarWithViewController demonstrates how to create a SplitViewItem instance using NewSplitViewItemSidebarWithViewController.
// Creates a split view item that represents a sidebar for the specified view controller.
func ExampleNewSplitViewItemSidebarWithViewController() {
	_ = appkit.NewSplitViewItemSidebarWithViewController(
		appkit.NSViewController{}, // viewController NSViewController
	)
	// Output:
}

// ExampleNewSplitViewItemWithViewController demonstrates how to create a SplitViewItem instance using NewSplitViewItemWithViewController.
// Creates a split view item that represents the specified view controller.
func ExampleNewSplitViewItemWithViewController() {
	_ = appkit.NewSplitViewItemWithViewController(
		appkit.NSViewController{}, // viewController NSViewController
	)
	// Output:
}
