// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TabViewController] class.
var tabViewControllerClass = _TabViewControllerClass{objc.GetClass("NSTabViewController")}

type _TabViewControllerClass struct {
	class objc.Class
}

// An interface definition for the [TabViewController] class.
type ITabViewController interface {
	IViewController
}

// A container view controller that manages a tab view interface, which organizes multiple pages of content but displays only one page at a time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewController

type TabViewController struct {
	ViewController
}

// TabViewControllerFrom constructs a [TabViewController] from an unsafe.Pointer.
//
// A container view controller that manages a tab view interface, which organizes multiple pages of content but displays only one page at a time.
func TabViewControllerFrom(ptr unsafe.Pointer) TabViewController {
	return TabViewController{
		ViewController: ViewControllerFrom(ptr),
	}
}



