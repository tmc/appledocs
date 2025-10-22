// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SplitViewController] class.
var (
	SplitViewControllerClass     _SplitViewControllerClass
	SplitViewControllerClassOnce sync.Once
)

func getSplitViewControllerClass() _SplitViewControllerClass {
	SplitViewControllerClassOnce.Do(func() {
		SplitViewControllerClass = _SplitViewControllerClass{objc.GetClass("NSSplitViewController")}
	})
	return SplitViewControllerClass
}

type _SplitViewControllerClass struct {
	class objc.Class
}

// An interface definition for the [SplitViewController] class.
type ISplitViewController interface {
	IViewController
	AddSplitViewItem(splitViewItem ISplitViewItem)
	InsertSplitViewItemAtIndex(splitViewItem ISplitViewItem, index int)
	RemoveSplitViewItem(splitViewItem ISplitViewItem)
	SplitViewItemForViewController(viewController IViewController) SplitViewItem
	ToggleInspector(sender objectivec.IObject)
	ToggleSidebar(sender objectivec.IObject)
	ViewDidLoad()
	MinimumThicknessForInlineSidebars() float64
	SetMinimumThicknessForInlineSidebars(value float64)
	SplitView() NSSplitView
	SetSplitView(value ISplitView)
	SplitViewItems() []SplitViewItem
	SetSplitViewItems(value []SplitViewItem)
	IsVertical() bool
	SetIsVertical(value bool)
}

// An object that manages an array of adjacent child views, and has a split view object for managing dividers between those views.
//
// A split view controller manages a set of child views that it displays next to each other in a side-by-side or top-to-bottom arrangement. A split view controller owns an array of split view items ( ), each of which has a view controller ( ) and corresponding view. The split view controller’s object manages those child views and the dividers between them. By default, a split view arranges its child views vertically from top to bottom. To specify a horizontal (side-by-side) arrangement, implement the property of the object to return . The split view controller serves as the delegate of its object. If you override a split view delegate method, your override must call . To use a split view controller, you must use Auto Layout for the child views and to support animations that collapse and reveal child views. For example, if you design a layout that contains two views, a content area and an optional sidebar, you employ Auto Layout constraints to specify whether the content area shrinks or remains the same size when the sidebar becomes visible. A split view controller employs lazy loading of its views. For example, adding a collapsed split view item as a new child doesn’t load the associated view until it shows. For more information about using in your app, see .
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

// Alloc allocates a new instance without initialization.
func (sc _SplitViewControllerClass) Alloc() SplitViewController {
	rv := objc.Send[SplitViewController](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SplitViewControllerClass) New() SplitViewController {
	rv := objc.Send[SplitViewController](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SplitViewController) Init() SplitViewController {
	rv := objc.Send[SplitViewController](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SplitViewController) Autorelease() SplitViewController {
	rv := objc.Send[SplitViewController](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSplitViewController creates a new SplitViewController instance.
func NewSplitViewController() SplitViewController {
	return getSplitViewControllerClass().New()
}


// Adds a split view item to the end of the array of split view items.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewController/addSplitViewItem(_:)
func (s_ SplitViewController) AddSplitViewItem(splitViewItem ISplitViewItem) {
	objc.Send[objc.ID](s_.ID, objc.Sel("addSplitViewItem:"), splitViewItem)
}

// Adds a split view item to the array of split view items at the specified index position.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewController/insertSplitViewItem(_:at:)
func (s_ SplitViewController) InsertSplitViewItemAtIndex(splitViewItem ISplitViewItem, index int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("insertSplitViewItem:atIndex:"), splitViewItem, index)
}

// Removes a specified split view item from the split view controller.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewController/removeSplitViewItem(_:)
func (s_ SplitViewController) RemoveSplitViewItem(splitViewItem ISplitViewItem) {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeSplitViewItem:"), splitViewItem)
}

// Returns the corresponding split view item for the specified child view controller of the split view controller.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewController/splitViewItem(for:)
func (s_ SplitViewController) SplitViewItemForViewController(viewController IViewController) SplitViewItem {
	rv := objc.Send[SplitViewItem](s_.ID, objc.Sel("splitViewItemForViewController:"), viewController)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewController/toggleInspector(_:)
func (s_ SplitViewController) ToggleInspector(sender objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("toggleInspector:"), sender)
}

// Collapses or expands the first sidebar in the split view controller using an animation.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewController/toggleSidebar(_:)
func (s_ SplitViewController) ToggleSidebar(sender objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("toggleSidebar:"), sender)
}

// Configures the split view controller after its view loads into memory.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewController/viewDidLoad()
func (s_ SplitViewController) ViewDidLoad() {
	objc.Send[objc.ID](s_.ID, objc.Sel("viewDidLoad"))
}

// The minimum thickness for a sidebar before it automatically collapses.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewController/minimumThicknessForInlineSidebars
func (s_ SplitViewController) MinimumThicknessForInlineSidebars() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("minimumThicknessForInlineSidebars"))
	return rv
}


// SetMinimumThicknessForInlineSidebars sets the value of the minimumThicknessForInlineSidebars property.
// The minimum thickness for a sidebar before it automatically collapses.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewController/minimumThicknessForInlineSidebars
func (s_ SplitViewController) SetMinimumThicknessForInlineSidebars(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMinimumThicknessForInlineSidebars:"), value)
}

// The split view that the split view controller manages.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewController/splitView
func (s_ SplitViewController) SplitView() NSSplitView {
	rv := objc.Send[NSSplitView](s_.ID, objc.Sel("splitView"))
	return rv
}


// SetSplitView sets the value of the splitView property.
// The split view that the split view controller manages.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewController/splitView
func (s_ SplitViewController) SetSplitView(value ISplitView) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSplitView:"), value)
}

// The array of split view items that correspond to the split view controller’s child view controllers.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewController/splitViewItems
func (s_ SplitViewController) SplitViewItems() []SplitViewItem {
	rv := objc.Send[[]SplitViewItem](s_.ID, objc.Sel("splitViewItems"))
	return rv
}


// SetSplitViewItems sets the value of the splitViewItems property.
// The array of split view items that correspond to the split view controller’s child view controllers.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewController/splitViewItems
func (s_ SplitViewController) SetSplitViewItems(value []SplitViewItem) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](s_.ID, objc.Sel("setSplitViewItems:"), nsArray)
}

// A Boolean value that determines the geometric orientation of the split view’s dividers.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitview/isvertical
func (s_ SplitViewController) IsVertical() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isVertical"))
	return rv
}


// SetIsVertical sets the value of the isVertical property.
// A Boolean value that determines the geometric orientation of the split view’s dividers.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitview/isvertical
func (s_ SplitViewController) SetIsVertical(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsVertical:"), value)
}



