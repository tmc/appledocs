// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewViewController

// ExampleViewController_CommitEditing demonstrates using CommitEditing on a ViewController instance.
// Returns whether the receiver was able to commit any pending edits.
//
// Note: This example is not executed because CommitEditing crashes when called on bare NSObject
// (it's a protocol/category method that should be overridden by subclasses).
func ExampleViewController_CommitEditing() {
	obj := appkit.NewViewController()
	_ = obj.CommitEditing()
	}

// ExampleViewController_DiscardEditing demonstrates using DiscardEditing on a ViewController instance.
// Causes the receiver to discard any changes, restoring the previous values.
//
// Note: This example is not executed because DiscardEditing crashes when called on bare NSObject
// (it's a protocol/category method that should be overridden by subclasses).
func ExampleViewController_DiscardEditing() {
	obj := appkit.NewViewController()
	obj.DiscardEditing()
	}

// ExampleViewController_LoadView demonstrates using LoadView on a ViewController instance.
// Instantiates a view from a nib file and sets the value of the   property.
func ExampleViewController_LoadView() {
	obj := appkit.NewViewController()
	obj.LoadView()
	// Output:
	}

// ExampleViewController_LoadViewIfNeeded demonstrates using LoadViewIfNeeded on a ViewController instance.
func ExampleViewController_LoadViewIfNeeded() {
	obj := appkit.NewViewController()
	obj.LoadViewIfNeeded()
	// Output:
	}

// ExampleViewController_RemoveFromParentViewController demonstrates using RemoveFromParentViewController on a ViewController instance.
// Removes the called view controller from its parent view controller.
func ExampleViewController_RemoveFromParentViewController() {
	obj := appkit.NewViewController()
	obj.RemoveFromParentViewController()
	// Output:
	}

// ExampleViewController_UpdateViewConstraints demonstrates using UpdateViewConstraints on a ViewController instance.
// Called during Auto Layout constraint updating to enable the view controller to mediate the process.
func ExampleViewController_UpdateViewConstraints() {
	obj := appkit.NewViewController()
	obj.UpdateViewConstraints()
	// Output:
	}

// ExampleViewController_ViewDidAppear demonstrates using ViewDidAppear on a ViewController instance.
// Called when the view controller’s view is fully transitioned onto the screen.
func ExampleViewController_ViewDidAppear() {
	obj := appkit.NewViewController()
	obj.ViewDidAppear()
	// Output:
	}

// ExampleViewController_ViewDidDisappear demonstrates using ViewDidDisappear on a ViewController instance.
// Called after the view controller’s view is removed from the view hierarchy in a window.
func ExampleViewController_ViewDidDisappear() {
	obj := appkit.NewViewController()
	obj.ViewDidDisappear()
	// Output:
	}

// ExampleViewController_ViewDidLayout demonstrates using ViewDidLayout on a ViewController instance.
// Called immediately after the   method of the view controller’s view is called.
func ExampleViewController_ViewDidLayout() {
	obj := appkit.NewViewController()
	obj.ViewDidLayout()
	// Output:
	}

// ExampleViewController_ViewDidLoad demonstrates using ViewDidLoad on a ViewController instance.
// Called after the view controller’s view has been loaded into memory.
func ExampleViewController_ViewDidLoad() {
	obj := appkit.NewViewController()
	obj.ViewDidLoad()
	// Output:
	}

// ExampleViewController_ViewWillAppear demonstrates using ViewWillAppear on a ViewController instance.
// Called after the view controller’s view has been loaded into memory is about to be added to the view hierarchy in the window.
func ExampleViewController_ViewWillAppear() {
	obj := appkit.NewViewController()
	obj.ViewWillAppear()
	// Output:
	}

// ExampleViewController_ViewWillDisappear demonstrates using ViewWillDisappear on a ViewController instance.
// Called when the view controller’s view is about to be removed from the view hierarchy in the window.
func ExampleViewController_ViewWillDisappear() {
	obj := appkit.NewViewController()
	obj.ViewWillDisappear()
	// Output:
	}

// ExampleViewController_ViewWillLayout demonstrates using ViewWillLayout on a ViewController instance.
// Called just before the   method of the view controller’s view is called.
func ExampleViewController_ViewWillLayout() {
	obj := appkit.NewViewController()
	obj.ViewWillLayout()
	// Output:
	}

