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
	

	// properties:
	CanPropagateSelectedChildViewControllerTitle() bool
	SetCanPropagateSelectedChildViewControllerTitle(value bool)
	SelectedTabViewItemIndex() int
	SetSelectedTabViewItemIndex(value int)
	TabStyle() TabViewControllerTabStyle
	SetTabStyle(value TabViewControllerTabStyle)
	TabView() ITabView
	SetTabView(value ITabView)
	TabViewItems() []TabViewItem
	SetTabViewItems(value []TabViewItem)
	TransitionOptions() ViewControllerTransitionOptions
	SetTransitionOptions(value ViewControllerTransitionOptions)
	Children() IViewController
	SetChildren(value IViewController)


	

	// methods:
	AddTabViewItem(tabViewItem ITabViewItem)
	InsertTabViewItemAtIndex(tabViewItem ITabViewItem, index int)
	RemoveTabViewItem(tabViewItem ITabViewItem)
	TabViewDidSelectTabViewItem(tabView ITabView, tabViewItem ITabViewItem)
	TabViewShouldSelectTabViewItem(tabView ITabView, tabViewItem ITabViewItem) bool
	TabViewWillSelectTabViewItem(tabView ITabView, tabViewItem ITabViewItem)
	TabViewItemForViewController(viewController IViewController) ITabViewItem
	ToolbarItemForItemIdentifierWillBeInsertedIntoToolbar(toolbar IToolbar, itemIdentifier ToolbarItemIdentifier, flag bool) IToolbarItem
	ToolbarAllowedItemIdentifiers(toolbar IToolbar) []string
	ToolbarDefaultItemIdentifiers(toolbar IToolbar) []string
	ToolbarSelectableItemIdentifiers(toolbar IToolbar) []string
	ViewDidLoad()


}





// Alloc allocates a new instance without initialization.
func (tc _TabViewControllerClass) Alloc() TabViewController {
	rv := objc.Send[TabViewController](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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




















// Adds the specified tab to the end of the tab view controller’s list of tabs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewController/addTabViewItem(_:)
func (t_ TabViewController) AddTabViewItem(tabViewItem ITabViewItem) {
	objc.Send[objc.ID](t_.ID, objc.Sel("addTabViewItem:"), tabViewItem)
}


// Inserts a tab view into the tab view controller’s list of tabs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewController/insertTabViewItem(_:at:)
func (t_ TabViewController) InsertTabViewItemAtIndex(tabViewItem ITabViewItem, index int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("insertTabViewItem:atIndex:"), tabViewItem, index)
}


// Removes the specified tab view item from the tab view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewController/removeTabViewItem(_:)
func (t_ TabViewController) RemoveTabViewItem(tabViewItem ITabViewItem) {
	objc.Send[objc.ID](t_.ID, objc.Sel("removeTabViewItem:"), tabViewItem)
}


// Informs the tab view controller that the specified tab was selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewController/tabView(_:didSelect:)
func (t_ TabViewController) TabViewDidSelectTabViewItem(tabView ITabView, tabViewItem ITabViewItem) {
	objc.Send[objc.ID](t_.ID, objc.Sel("tabView:didSelectTabViewItem:"), tabView, tabViewItem)
}


// Asks the tab view controller if the specified tab should be selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewController/tabView(_:shouldSelect:)
func (t_ TabViewController) TabViewShouldSelectTabViewItem(tabView ITabView, tabViewItem ITabViewItem) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("tabView:shouldSelectTabViewItem:"), tabView, tabViewItem)
	return rv
}


// Informs the tab view controller that the specified tab is about to be selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewController/tabView(_:willSelect:)
func (t_ TabViewController) TabViewWillSelectTabViewItem(tabView ITabView, tabViewItem ITabViewItem) {
	objc.Send[objc.ID](t_.ID, objc.Sel("tabView:willSelectTabViewItem:"), tabView, tabViewItem)
}


// Returns the tab view item for the specified child view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewController/tabViewItem(for:)
func (t_ TabViewController) TabViewItemForViewController(viewController IViewController) ITabViewItem {
	rv := objc.Send[TabViewItem](t_.ID, objc.Sel("tabViewItemForViewController:"), viewController)
	return rv
}


// Returns the toolbar item for the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewController/toolbar(_:itemForItemIdentifier:willBeInsertedIntoToolbar:)
func (t_ TabViewController) ToolbarItemForItemIdentifierWillBeInsertedIntoToolbar(toolbar IToolbar, itemIdentifier ToolbarItemIdentifier, flag bool) IToolbarItem {
	rv := objc.Send[ToolbarItem](t_.ID, objc.Sel("toolbar:itemForItemIdentifier:willBeInsertedIntoToolbar:"), toolbar, itemIdentifier, flag)
	return rv
}


// Returns the array of identifier strings for the allowed toolbar items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewController/toolbarAllowedItemIdentifiers(_:)
func (t_ TabViewController) ToolbarAllowedItemIdentifiers(toolbar IToolbar) []string {
	rv := objc.Send[[]string](t_.ID, objc.Sel("toolbarAllowedItemIdentifiers:"), toolbar)
	return rv
}


// Returns the array of identifier strings for the default toolbar items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewController/toolbarDefaultItemIdentifiers(_:)
func (t_ TabViewController) ToolbarDefaultItemIdentifiers(toolbar IToolbar) []string {
	rv := objc.Send[[]string](t_.ID, objc.Sel("toolbarDefaultItemIdentifiers:"), toolbar)
	return rv
}


// Returns the array of identifier strings for the selectable toolbar items
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewController/toolbarSelectableItemIdentifiers(_:)
func (t_ TabViewController) ToolbarSelectableItemIdentifiers(toolbar IToolbar) []string {
	rv := objc.Send[[]string](t_.ID, objc.Sel("toolbarSelectableItemIdentifiers:"), toolbar)
	return rv
}


// Called after the view controller’s view has been loaded into memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewController/viewDidLoad()
func (t_ TabViewController) ViewDidLoad() {
	objc.Send[objc.ID](t_.ID, objc.Sel("viewDidLoad"))
}







// A Boolean value indicating whether the tab view controller gets its title from the selected child view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewController/canPropagateSelectedChildViewControllerTitle
func (t_ TabViewController) CanPropagateSelectedChildViewControllerTitle() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("canPropagateSelectedChildViewControllerTitle"))
	return rv
}


// A Boolean value indicating whether the tab view controller gets its title from the selected child view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewController/canPropagateSelectedChildViewControllerTitle
func (t_ TabViewController) SetCanPropagateSelectedChildViewControllerTitle(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCanPropagateSelectedChildViewControllerTitle:"), value)
}


// The index of the selected tab.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewController/selectedTabViewItemIndex
func (t_ TabViewController) SelectedTabViewItemIndex() int {
	rv := objc.Send[int](t_.ID, objc.Sel("selectedTabViewItemIndex"))
	return rv
}


// The index of the selected tab.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewController/selectedTabViewItemIndex
func (t_ TabViewController) SetSelectedTabViewItemIndex(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectedTabViewItemIndex:"), value)
}


// The style used to display the tabs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewController/tabStyle-swift.property
func (t_ TabViewController) TabStyle() TabViewControllerTabStyle {
	rv := objc.Send[TabViewControllerTabStyle](t_.ID, objc.Sel("tabStyle"))
	return rv
}


// The style used to display the tabs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewController/tabStyle-swift.property
func (t_ TabViewController) SetTabStyle(value TabViewControllerTabStyle) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTabStyle:"), value)
}


// The tab view that manages the views of the interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewController/tabView
func (t_ TabViewController) TabView() ITabView {
	rv := objc.Send[TabView](t_.ID, objc.Sel("tabView"))
	return rv
}


// The tab view that manages the views of the interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewController/tabView
func (t_ TabViewController) SetTabView(value ITabView) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTabView:"), value)
}


// The array of tab view items used to manage each of the child view controllers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewController/tabViewItems
func (t_ TabViewController) TabViewItems() []TabViewItem {
	rv := objc.Send[[]TabViewItem](t_.ID, objc.Sel("tabViewItems"))
	return rv
}


// The array of tab view items used to manage each of the child view controllers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewController/tabViewItems
func (t_ TabViewController) SetTabViewItems(value []TabViewItem) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](t_.ID, objc.Sel("setTabViewItems:"), nsArray)
}


// The animation options to use when switching between tabs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewController/transitionOptions
func (t_ TabViewController) TransitionOptions() ViewControllerTransitionOptions {
	rv := objc.Send[ViewControllerTransitionOptions](t_.ID, objc.Sel("transitionOptions"))
	return rv
}


// The animation options to use when switching between tabs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewController/transitionOptions
func (t_ TabViewController) SetTransitionOptions(value ViewControllerTransitionOptions) {
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








