// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PageController] class.
var pageControllerClass = _PageControllerClass{objc.GetClass("NSPageController")}

type _PageControllerClass struct {
	class objc.Class
}

// An interface definition for the [PageController] class.
type IPageController interface {
	IViewController
}

// An object that controls swipe navigation and animations between views or view content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageController

type PageController struct {
	ViewController
}

// PageControllerFrom constructs a [PageController] from an unsafe.Pointer.
//
// An object that controls swipe navigation and animations between views or view content.
func PageControllerFrom(ptr unsafe.Pointer) PageController {
	return PageController{
		ViewController: ViewControllerFrom(ptr),
	}
}



