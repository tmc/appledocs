// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	// properties:
	SplitView() ISplitView
	SetSplitView(value ISplitView)
	IsVertical() bool
	SetIsVertical(value bool)
	MinimumThicknessForInlineSidebars() float64
	SetMinimumThicknessForInlineSidebars(value float64)
	SplitViewItems() ISplitViewItem
	SetSplitViewItems(value ISplitViewItem)
	// methods:
}

// An object that manages an array of adjacent child views, and has a split view object for managing dividers between those views.
//
// A split view controller manages a set of child views that it displays next to each other in a side-by-side or top-to-bottom arrangement. A split view controller owns an array of split view items ( ), each of which has a view controller ( ) and corresponding view. The split view controller’s object manages those child views and the dividers between them. By default, a split view arranges its child views vertically from top to bottom. To specify a horizontal (side-by-side) arrangement, implement the property of the object to return . The split view controller serves as the delegate of its object. If you override a split view delegate method, your override must call . To use a split view controller, you must use Auto Layout for the child views and to support animations that collapse and reveal child views. For example, if you design a layout that contains two views, a content area and an optional sidebar, you employ Auto Layout constraints to specify whether the content area shrinks or remains the same size when the sidebar becomes visible. A split view controller employs lazy loading of its views. For example, adding a collapsed split view item as a new child doesn’t load the associated view until it shows. For more information about using in your app, see .


// An object that manages an array of adjacent child views, and has a split view object for managing dividers between those views.
//
// [Full Topic]
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



// The split view that the split view controller manages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewController/splitView
func (s_ SplitViewController) SplitView() ISplitView {
	rv := objc.Send[SplitView](s_.ID, objc.Sel("splitView"))
	return rv
}


// The split view that the split view controller manages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewController/splitView
func (s_ SplitViewController) SetSplitView(value ISplitView) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSplitView:"), value)
}


// A Boolean value that determines the geometric orientation of the split view’s dividers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitview/isvertical
func (s_ SplitViewController) IsVertical() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isVertical"))
	return rv
}


// A Boolean value that determines the geometric orientation of the split view’s dividers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitview/isvertical
func (s_ SplitViewController) SetIsVertical(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsVertical:"), value)
}


// The minimum thickness for a sidebar before it automatically collapses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewcontroller/minimumthicknessforinlinesidebars
func (s_ SplitViewController) MinimumThicknessForInlineSidebars() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("minimumThicknessForInlineSidebars"))
	return rv
}


// The minimum thickness for a sidebar before it automatically collapses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewcontroller/minimumthicknessforinlinesidebars
func (s_ SplitViewController) SetMinimumThicknessForInlineSidebars(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMinimumThicknessForInlineSidebars:"), value)
}


// The array of split view items that correspond to the split view controller’s child view controllers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewcontroller/splitviewitems
func (s_ SplitViewController) SplitViewItems() ISplitViewItem {
	rv := objc.Send[SplitViewItem](s_.ID, objc.Sel("splitViewItems"))
	return rv
}


// The array of split view items that correspond to the split view controller’s child view controllers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewcontroller/splitviewitems
func (s_ SplitViewController) SetSplitViewItems(value ISplitViewItem) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSplitViewItems:"), value)
}



