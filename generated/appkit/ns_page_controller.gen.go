// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	ArrangedObjects() objc.IObject /* cross-framework: NSArray */
	SetArrangedObjects(value objc.IObject /* cross-framework: NSArray */)
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	SelectedIndex() int
	SetSelectedIndex(value int)
	SelectedViewController() IViewController
	TransitionStyle() PageControllerTransitionStyle
	SetTransitionStyle(value PageControllerTransitionStyle)
	// methods:
	CompleteTransition()
	NavigateBack(sender objectivec.IObject)
	NavigateForward(sender objectivec.IObject)
	NavigateForwardToObject(object objectivec.IObject)
	TakeSelectedIndexFrom(sender objectivec.IObject)
}

// An object that controls swipe navigation and animations between views or view content.
//
// is useful for user interfaces which control navigating multiple pages as in a book or a web browser history. Page controller inherits from the class . You must assign the property to a view in your view hierarchy. The class does not vend a view and does insert itself into the responder chain. Conceptually, the page controller manages swiping between an array of pages, the . Using the property, you can determine how many pages forward or backward the user may navigate.


// An object that controls swipe navigation and animations between views or view content.
//
// [Full Topic]
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



// Invoked when the page transition is completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageController/completeTransition()
func (p_ PageController) CompleteTransition() {
	objc.Send[objc.ID](p_.ID, objc.Sel("completeTransition"))
}


// Navigates backwards in the page controller’s arranged objects array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageController/navigateBack(_:)
func (p_ PageController) NavigateBack(sender objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("navigateBack:"), sender)
}


// Navigates to the next object in the page controller’s arranged objects array, if appropriate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageController/navigateForward(_:)
func (p_ PageController) NavigateForward(sender objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("navigateForward:"), sender)
}


// Navigates to the specific object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageController/navigateForward(to:)
func (p_ PageController) NavigateForwardToObject(object objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("navigateForwardToObject:"), object)
}


// Navigates to the selected index, which is taken from the sender.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageController/takeSelectedIndexFrom(_:)
func (p_ PageController) TakeSelectedIndexFrom(sender objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("takeSelectedIndexFrom:"), sender)
}


// An array containing the objects displayed in the page controller’s view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageController/arrangedObjects
func (p_ PageController) ArrangedObjects() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](p_.ID, objc.Sel("arrangedObjects"))
	return rv
}


// An array containing the objects displayed in the page controller’s view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageController/arrangedObjects
func (p_ PageController) SetArrangedObjects(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setArrangedObjects:"), value)
}


// The page controller’s delegate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageController/delegate
func (p_ PageController) Delegate() objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("delegate"))
	return rv
}


// The page controller’s delegate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageController/delegate
func (p_ PageController) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDelegate:"), value)
}


// The currently selected object in the arranged objects array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageController/selectedIndex
func (p_ PageController) SelectedIndex() int {
	rv := objc.Send[int](p_.ID, objc.Sel("selectedIndex"))
	return rv
}


// The currently selected object in the arranged objects array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageController/selectedIndex
func (p_ PageController) SetSelectedIndex(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSelectedIndex:"), value)
}


// The view controller associated with the selected object..
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageController/selectedViewController
func (p_ PageController) SelectedViewController() IViewController {
	rv := objc.Send[ViewController](p_.ID, objc.Sel("selectedViewController"))
	return rv
}


// The transition style the page controller uses when changing pages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageController/transitionStyle-swift.property
func (p_ PageController) TransitionStyle() PageControllerTransitionStyle {
	rv := objc.Send[PageControllerTransitionStyle](p_.ID, objc.Sel("transitionStyle"))
	return rv
}


// The transition style the page controller uses when changing pages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageController/transitionStyle-swift.property
func (p_ PageController) SetTransitionStyle(value PageControllerTransitionStyle) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTransitionStyle:"), value)
}



