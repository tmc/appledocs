// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PageController] class.
var (
	PageControllerClass     _PageControllerClass
	PageControllerClassOnce sync.Once
)

func getPageControllerClass() _PageControllerClass {
	PageControllerClassOnce.Do(func() {
		PageControllerClass = _PageControllerClass{objc.GetClass("NSPageController")}
	})
	return PageControllerClass
}

type _PageControllerClass struct {
	class objc.Class
}

// An interface definition for the [PageController] class.
type IPageController interface {
	IViewController
	NavigateForwardToObject(object objc.ID)
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


// Navigates to the specific object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageController/navigateForward(to:)
func (p_ PageController) NavigateForwardToObject(object objc.ID) {
	objc.Send[objc.ID](p_.ID, objc.Sel("navigateForwardToObject:"), object)
}

// An array containing the objects displayed in the page controller’s view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageController/arrangedObjects
func (p_ PageController) ArrangedObjects() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("arrangedObjects"))
	return rv
}


// SetArrangedObjects sets the value of the arrangedObjects property.
// An array containing the objects displayed in the page controller’s view.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageController/arrangedObjects
func (p_ PageController) SetArrangedObjects(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setArrangedObjects:"), value)
}
// The currently selected object in the arranged objects array.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageController/selectedIndex
func (p_ PageController) SelectedIndex() int {
	rv := objc.Send[int](p_.ID, objc.Sel("selectedIndex"))
	return rv
}


// SetSelectedIndex sets the value of the selectedIndex property.
// The currently selected object in the arranged objects array.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageController/selectedIndex
func (p_ PageController) SetSelectedIndex(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSelectedIndex:"), value)
}
// The view controller associated with the selected object..
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageController/selectedViewController
func (p_ PageController) SelectedViewController() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("selectedViewController"))
	return rv
}



