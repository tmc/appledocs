// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TabViewController] class.
var (
	tabViewControllerClass     _TabViewControllerClass
	tabViewControllerClassOnce sync.Once
)

func getTabViewControllerClass() _TabViewControllerClass {
	tabViewControllerClassOnce.Do(func() {
		tabViewControllerClass = _TabViewControllerClass{objc.GetClass("NSTabViewController")}
	})
	return tabViewControllerClass
}

type _TabViewControllerClass struct {
	class objc.Class
}

// An interface definition for the [TabViewController] class.
type ITabViewController interface {
	IViewController
}

// A container view controller that manages a tab view interface, which organizes multiple pages of content but displays only one page at a time.
//
// Each page of content is managed by a separate child view controller. Navigation between child view controllers is accomplished with the help of an object, which the tab view controller manages. When the user selects a new tab, the tab view controller displays the content associated with the associated child view controller, replacing the previous content. Each tab is represented by an object, which contains the name of the tab and stores a pointer to the child view controller that manages the tab’s content. Normally, you configure the tab view items at design time using Interface Builder, but you can also add them programmatically using the methods of this class. Always assign a child view controller to new tab view items before adding those items to the tab view interface. Another way to add tabs programmatically is to add child view controllers directly to the tab view controller. When you call the or method of this class, the tab view controller automatically creates a default object for the specified view controller. You can fetch the newly created item using the method and configure it. Removing a child view controller with the method similarly removes the corresponding tab view item. The tab view controller lazily loads the views associated with each child view controller, creating them only after the corresponding tab is selected. When the tab view controller’s view is first displayed, only the view for the initially selected tab is loaded. The property determines the appearance of the tab controls. A tab view controller can display a segmented control or display tabs in the window’s toolbar. You can also provide your own control for displaying tabs. The tab view controller automatically coordinates interactions between designated control and the corresponding object.
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

// Alloc allocates a new instance without initialization.
func (tc _TabViewControllerClass) Alloc() TabViewController {
	rv := objc.Send[TabViewController](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TabViewControllerClass) New() TabViewController {
	rv := objc.Send[TabViewController](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TabViewController) Init() TabViewController {
	rv := objc.Send[TabViewController](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TabViewController) Autorelease() TabViewController {
	rv := objc.Send[TabViewController](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTabViewController creates a new TabViewController instance.
func NewTabViewController() TabViewController {
	return getTabViewControllerClass().New()
}




