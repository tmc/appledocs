// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SplitViewItem] class.
var (
	SplitViewItemClass     _SplitViewItemClass
	SplitViewItemClassOnce sync.Once
)

func getSplitViewItemClass() _SplitViewItemClass {
	SplitViewItemClassOnce.Do(func() {
		SplitViewItemClass = _SplitViewItemClass{objc.GetClass("NSSplitViewItem")}
	})
	return SplitViewItemClass
}

type _SplitViewItemClass struct {
	class objc.Class
}

// An interface definition for the [SplitViewItem] class.
type ISplitViewItem interface {
	objectivec.IObject
	InsertBottomAlignedAccessoryViewControllerAtIndex(childViewController unsafe.Pointer, index int)
}

// An item in a split view controller.
//
// A split view item represents a single pane in a split view controller ( ). Each split view item contains information about a child view controller in the split view controller, like its preferred thickness, holding priority, and collapsed state. To add one or more accessory views to the top or bottom of a split view item, such as a search field above a list, use the and properties to specify types.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem
type SplitViewItem struct {
	objectivec.Object
}

// SplitViewItemFrom constructs a [SplitViewItem] from an unsafe.Pointer.
//
// An item in a split view controller.
func SplitViewItemFrom(ptr unsafe.Pointer) SplitViewItem {
	return SplitViewItem{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SplitViewItemClass) Alloc() SplitViewItem {
	rv := objc.Send[SplitViewItem](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SplitViewItemClass) New() SplitViewItem {
	rv := objc.Send[SplitViewItem](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SplitViewItem) Init() SplitViewItem {
	rv := objc.Send[SplitViewItem](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SplitViewItem) Autorelease() SplitViewItem {
	rv := objc.Send[SplitViewItem](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSplitViewItem creates a new SplitViewItem instance.
func NewSplitViewItem() SplitViewItem {
	return getSplitViewItemClass().New()
}


// Creates a split view item that represents a content list for the specified view controller.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/init(contentListWithViewController:)
func NewSplitViewItemContentListWithViewController(viewController unsafe.Pointer) SplitViewItem {
	rv := objc.Send[SplitViewItem](objc.ID(getSplitViewItemClass().class), objc.Sel("contentListWithViewController:"), viewController)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/init(inspectorWithViewController:)
func NewSplitViewItemInspectorWithViewController(viewController unsafe.Pointer) SplitViewItem {
	rv := objc.Send[SplitViewItem](objc.ID(getSplitViewItemClass().class), objc.Sel("inspectorWithViewController:"), viewController)
	return rv
}


// Creates a split view item that represents a content list for the specified view controller.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/init(contentListWithViewController:)
func (sc _SplitViewItemClass) ContentListWithViewController(viewController unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("contentListWithViewController:"), viewController)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/init(inspectorWithViewController:)
func (sc _SplitViewItemClass) InspectorWithViewController(viewController unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("inspectorWithViewController:"), viewController)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/insertBottomAlignedAccessoryViewController(_:at:)
func (s_ SplitViewItem) InsertBottomAlignedAccessoryViewControllerAtIndex(childViewController unsafe.Pointer, index int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("insertBottomAlignedAccessoryViewController:atIndex:"), childViewController, index)
}

// A Boolean value that indicates whether full-height sidebars appear in the window after you set a style mask.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/allowsFullHeightLayout
func (s_ SplitViewItem) AllowsFullHeightLayout() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("allowsFullHeightLayout"))
	return rv
}


// SetAllowsFullHeightLayout sets the value of the allowsFullHeightLayout property.
// A Boolean value that indicates whether full-height sidebars appear in the window after you set a style mask.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/allowsFullHeightLayout
func (s_ SplitViewItem) SetAllowsFullHeightLayout(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAllowsFullHeightLayout:"), value)
}
// The standard behavior type of the split view item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/behavior-swift.property
func (s_ SplitViewItem) Behavior() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("behavior"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/bottomAlignedAccessoryViewControllers
func (s_ SplitViewItem) BottomAlignedAccessoryViewControllers() []SplitViewItemAccessoryViewController {
	rv := objc.Send[[]SplitViewItemAccessoryViewController](s_.ID, objc.Sel("bottomAlignedAccessoryViewControllers"))
	return rv
}


// SetBottomAlignedAccessoryViewControllers sets the value of the bottomAlignedAccessoryViewControllers property.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/bottomAlignedAccessoryViewControllers
func (s_ SplitViewItem) SetBottomAlignedAccessoryViewControllers(value []SplitViewItemAccessoryViewController) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setBottomAlignedAccessoryViewControllers:"), value)
}
// The minimum thickness of the split view item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/minimumThickness
func (s_ SplitViewItem) MinimumThickness() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("minimumThickness"))
	return rv
}


// SetMinimumThickness sets the value of the minimumThickness property.
// The minimum thickness of the split view item.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/minimumThickness
func (s_ SplitViewItem) SetMinimumThickness(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMinimumThickness:"), value)
}
// The type of separator that the app displays between the title bar and content of a window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/titlebarSeparatorStyle
func (s_ SplitViewItem) TitlebarSeparatorStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("titlebarSeparatorStyle"))
	return rv
}


// SetTitlebarSeparatorStyle sets the value of the titlebarSeparatorStyle property.
// The type of separator that the app displays between the title bar and content of a window.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/titlebarSeparatorStyle
func (s_ SplitViewItem) SetTitlebarSeparatorStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTitlebarSeparatorStyle:"), value)
}
// The following methods allow you to add accessory views to the top/bottom of this splitViewItem. See for more details.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/topAlignedAccessoryViewControllers
func (s_ SplitViewItem) TopAlignedAccessoryViewControllers() []SplitViewItemAccessoryViewController {
	rv := objc.Send[[]SplitViewItemAccessoryViewController](s_.ID, objc.Sel("topAlignedAccessoryViewControllers"))
	return rv
}


// SetTopAlignedAccessoryViewControllers sets the value of the topAlignedAccessoryViewControllers property.
// The following methods allow you to add accessory views to the top/bottom of this splitViewItem. See for more details.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/topAlignedAccessoryViewControllers
func (s_ SplitViewItem) SetTopAlignedAccessoryViewControllers(value []SplitViewItemAccessoryViewController) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTopAlignedAccessoryViewControllers:"), value)
}
// The view controller that the split view item represents.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/viewController
func (s_ SplitViewItem) ViewController() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("viewController"))
	return rv
}


// SetViewController sets the value of the viewController property.
// The view controller that the split view item represents.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/viewController
func (s_ SplitViewItem) SetViewController(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setViewController:"), value)
}

