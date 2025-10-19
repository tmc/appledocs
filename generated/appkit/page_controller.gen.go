// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PageController] class.
var (
	pageControllerClass     _PageControllerClass
	pageControllerClassOnce sync.Once
)

func getPageControllerClass() _PageControllerClass {
	pageControllerClassOnce.Do(func() {
		pageControllerClass = _PageControllerClass{objc.GetClass("NSPageController")}
	})
	return pageControllerClass
}

type _PageControllerClass struct {
	class objc.Class
}

// An interface definition for the [PageController] class.
type IPageController interface {
	IViewController
}

// An object that controls swipe navigation and animations between views or view content.
//
// is useful for user interfaces which control navigating multiple pages as in a book or a web browser history. Page controller inherits from the class . You must assign the property to a view in your view hierarchy. The class does not vend a view and does insert itself into the responder chain. Conceptually, the page controller manages swiping between an array of pages, the . Using the property, you can determine how many pages forward or backward the user may navigate.
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

// Alloc allocates a new instance without initialization.
func (pc _PageControllerClass) Alloc() PageController {
	rv := objc.Send[PageController](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PageControllerClass) New() PageController {
	rv := objc.Send[PageController](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PageController) Init() PageController {
	rv := objc.Send[PageController](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PageController) Autorelease() PageController {
	rv := objc.Send[PageController](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPageController creates a new PageController instance.
func NewPageController() PageController {
	return getPageControllerClass().New()
}




