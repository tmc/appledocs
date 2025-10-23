// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TabViewController] class.
var (
	TabViewControllerClass     _TabViewControllerClass
	TabViewControllerClassOnce sync.Once
)

func getTabViewControllerClass() _TabViewControllerClass {
	TabViewControllerClassOnce.Do(func() {
		TabViewControllerClass = _TabViewControllerClass{objc.GetClass("NSTabViewController")}
	})
	return TabViewControllerClass
}

type _TabViewControllerClass struct {
	class objc.Class
}

// An interface definition for the [TabViewController] class.
type ITabViewController interface {
	IViewController
	CanPropagateSelectedChildViewControllerTitle() bool
	SetCanPropagateSelectedChildViewControllerTitle(value bool)
	SelectedTabViewItemIndex() int
	SetSelectedTabViewItemIndex(value int)
	TabStyle() unsafe.Pointer
	SetTabStyle(value unsafe.Pointer)
	TabView() ITabView
	SetTabView(value ITabView)
	TabViewItems() TabViewItem
	SetTabViewItems(value TabViewItem)
	TransitionOptions() unsafe.Pointer
	SetTransitionOptions(value unsafe.Pointer)
	Children() IViewController
	SetChildren(value IViewController)
	AddTabViewItem(tabViewItem TabViewItem)
	InsertTabViewItemAtIndex(tabViewItem TabViewItem, index int)
	ToolbarAllowedItemIdentifiers(toolbar IToolbar) []string
}

// A container view controller that manages a tab view interface, which organizes multiple pages of content but displays only one page at a time.
//
// Each page of content is managed by a separate child view controller. Navigation between child view controllers is accomplished with the help of an object, which the tab view controller manages. When the user selects a new tab, the tab view controller displays the content associated with the associated child view controller, replacing the previous content. Each tab is represented by an object, which contains the name of the tab and stores a pointer to the child view controller that manages the tab’s content. Normally, you configure the tab view items at design time using Interface Builder, but you can also add them programmatically using the methods of this class. Always assign a child view controller to new tab view items before adding those items to the tab view interface. Another way to add tabs programmatically is to add child view controllers directly to the tab view controller. When you call the or method of this class, the tab view controller automatically creates a default object for the specified view controller. You can fetch the newly created item using the method and configure it. Removing a child view controller with the method similarly removes the corresponding tab view item. The tab view controller lazily loads the views associated with each child view controller, creating them only after the corresponding tab is selected. When the tab view controller’s view is first displayed, only the view for the initially selected tab is loaded. The property determines the appearance of the tab controls. A tab view controller can display a segmented control or display tabs in the window’s toolbar. You can also provide your own control for displaying tabs. The tab view controller automatically coordinates interactions between designated control and the corresponding object.


// A container view controller that manages a tab view interface, which organizes multiple pages of content but displays only one page at a time.
//
// [Full Topic]
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



// Adds the specified tab to the end of the tab view controller’s list of tabs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewController/addTabViewItem(_:)
func (t_ TabViewController) AddTabViewItem(tabViewItem TabViewItem) {
	objc.Send[objc.ID](t_.ID, objc.Sel("addTabViewItem:"), tabViewItem)
}


// Inserts a tab view into the tab view controller’s list of tabs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewController/insertTabViewItem(_:at:)
func (t_ TabViewController) InsertTabViewItemAtIndex(tabViewItem TabViewItem, index int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("insertTabViewItem:atIndex:"), tabViewItem, index)
}


// Returns the array of identifier strings for the allowed toolbar items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewController/toolbarAllowedItemIdentifiers(_:)
func (t_ TabViewController) ToolbarAllowedItemIdentifiers(toolbar IToolbar) []string {
	rv := objc.Send[[]string](t_.ID, objc.Sel("toolbarAllowedItemIdentifiers:"), toolbar)
	return rv
}


// A Boolean value indicating whether the tab view controller gets its title from the selected child view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewcontroller/canpropagateselectedchildviewcontrollertitle
func (t_ TabViewController) CanPropagateSelectedChildViewControllerTitle() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("canPropagateSelectedChildViewControllerTitle"))
	return rv
}


// A Boolean value indicating whether the tab view controller gets its title from the selected child view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewcontroller/canpropagateselectedchildviewcontrollertitle
func (t_ TabViewController) SetCanPropagateSelectedChildViewControllerTitle(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCanPropagateSelectedChildViewControllerTitle:"), value)
}


// The index of the selected tab.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewcontroller/selectedtabviewitemindex
func (t_ TabViewController) SelectedTabViewItemIndex() int {
	rv := objc.Send[int](t_.ID, objc.Sel("selectedTabViewItemIndex"))
	return rv
}


// The index of the selected tab.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewcontroller/selectedtabviewitemindex
func (t_ TabViewController) SetSelectedTabViewItemIndex(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectedTabViewItemIndex:"), value)
}


// The style used to display the tabs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewcontroller/tabstyle-swift.property
func (t_ TabViewController) TabStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("tabStyle"))
	return rv
}


// The style used to display the tabs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewcontroller/tabstyle-swift.property
func (t_ TabViewController) SetTabStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTabStyle:"), value)
}


// The tab view that manages the views of the interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewcontroller/tabview
func (t_ TabViewController) TabView() ITabView {
	rv := objc.Send[TabView](t_.ID, objc.Sel("tabView"))
	return rv
}


// The tab view that manages the views of the interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewcontroller/tabview
func (t_ TabViewController) SetTabView(value ITabView) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTabView:"), value)
}


// The array of tab view items used to manage each of the child view controllers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewcontroller/tabviewitems
func (t_ TabViewController) TabViewItems() TabViewItem {
	rv := objc.Send[TabViewItem](t_.ID, objc.Sel("tabViewItems"))
	return rv
}


// The array of tab view items used to manage each of the child view controllers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewcontroller/tabviewitems
func (t_ TabViewController) SetTabViewItems(value TabViewItem) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTabViewItems:"), value)
}


// The animation options to use when switching between tabs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewcontroller/transitionoptions
func (t_ TabViewController) TransitionOptions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("transitionOptions"))
	return rv
}


// The animation options to use when switching between tabs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewcontroller/transitionoptions
func (t_ TabViewController) SetTransitionOptions(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTransitionOptions:"), value)
}


// An array of view controllers that are hierarchical children of the view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/children
func (t_ TabViewController) Children() IViewController {
	rv := objc.Send[ViewController](t_.ID, objc.Sel("children"))
	return rv
}


// An array of view controllers that are hierarchical children of the view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/children
func (t_ TabViewController) SetChildren(value IViewController) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setChildren:"), value)
}



