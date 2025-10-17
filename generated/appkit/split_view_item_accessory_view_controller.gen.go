// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SplitViewItemAccessoryViewController] class.
var splitViewItemAccessoryViewControllerClass = _SplitViewItemAccessoryViewControllerClass{objc.GetClass("NSSplitViewItemAccessoryViewController")}

type _SplitViewItemAccessoryViewControllerClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItemAccessoryViewController

type SplitViewItemAccessoryViewController struct {
	ViewController
}

// SplitViewItemAccessoryViewControllerFrom constructs a [SplitViewItemAccessoryViewController] from an unsafe.Pointer.
func SplitViewItemAccessoryViewControllerFrom(ptr unsafe.Pointer) SplitViewItemAccessoryViewController {
	return SplitViewItemAccessoryViewController{
		ViewController: ViewControllerFrom(ptr),
	}
}



