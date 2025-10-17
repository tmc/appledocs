// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SplitViewController] class.
var splitViewControllerClass = _SplitViewControllerClass{objc.GetClass("NSSplitViewController")}

type _SplitViewControllerClass struct {
	class objc.Class
}

// An object that manages an array of adjacent child views, and has a split view object for managing dividers between those views. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewController

type SplitViewController struct {
	ViewController
}

// SplitViewControllerFrom constructs a [SplitViewController] from an unsafe.Pointer.
//
// An object that manages an array of adjacent child views, and has a split view object for managing dividers between those views.
func SplitViewControllerFrom(ptr unsafe.Pointer) SplitViewController {
	return SplitViewController{
		ViewController: ViewControllerFrom(ptr),
	}
}



