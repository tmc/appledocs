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
	// properties:
	AllowsFullHeightLayout() bool
	SetAllowsFullHeightLayout(value bool)
	AutomaticMaximumThickness() float64
	SetAutomaticMaximumThickness(value float64)
	AutomaticallyAdjustsSafeAreaInsets() bool
	SetAutomaticallyAdjustsSafeAreaInsets(value bool)
	Behavior() SplitViewItemBehavior
	BottomAlignedAccessoryViewControllers() []SplitViewItemAccessoryViewController
	SetBottomAlignedAccessoryViewControllers(value []SplitViewItemAccessoryViewController)
	CanCollapse() bool
	SetCanCollapse(value bool)
	CanCollapseFromWindowResize() bool
	SetCanCollapseFromWindowResize(value bool)
	CollapseBehavior() SplitViewItemCollapseBehavior
	SetCollapseBehavior(value SplitViewItemCollapseBehavior)
	HoldingPriority() objc.IObject /* cross-framework: LayoutPriority */
	SetHoldingPriority(value objc.IObject /* cross-framework: LayoutPriority */)
	Collapsed() bool
	SetCollapsed(value bool)
	SpringLoaded() bool
	SetSpringLoaded(value bool)
	MaximumThickness() float64
	SetMaximumThickness(value float64)
	MinimumThickness() float64
	SetMinimumThickness(value float64)
	PreferredThicknessFraction() float64
	SetPreferredThicknessFraction(value float64)
	TitlebarSeparatorStyle() TitlebarSeparatorStyle
	SetTitlebarSeparatorStyle(value TitlebarSeparatorStyle)
	TopAlignedAccessoryViewControllers() []SplitViewItemAccessoryViewController
	SetTopAlignedAccessoryViewControllers(value []SplitViewItemAccessoryViewController)
	ViewController() IViewController
	SetViewController(value IViewController)
	IsCollapsed() bool
	SetIsCollapsed(value bool)
	IsSpringLoaded() bool
	SetIsSpringLoaded(value bool)
	// methods:
	AddBottomAlignedAccessoryViewController(childViewController ISplitViewItemAccessoryViewController)
	AddTopAlignedAccessoryViewController(childViewController ISplitViewItemAccessoryViewController)
	InsertBottomAlignedAccessoryViewControllerAtIndex(childViewController ISplitViewItemAccessoryViewController, index int)
	InsertTopAlignedAccessoryViewControllerAtIndex(childViewController ISplitViewItemAccessoryViewController, index int)
	RemoveBottomAlignedAccessoryViewControllerAtIndex(index int)
	RemoveTopAlignedAccessoryViewControllerAtIndex(index int)
}

// An item in a split view controller.
//
// A split view item represents a single pane in a split view controller ( ). Each split view item contains information about a child view controller in the split view controller, like its preferred thickness, holding priority, and collapsed state. To add one or more accessory views to the top or bottom of a split view item, such as a search field above a list, use the and properties to specify types.


// An item in a split view controller.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/init(contentListWithViewController:)
func NewSplitViewItemContentListWithViewController(viewController IViewController) SplitViewItem {
	rv := objc.Send[SplitViewItem](objc.ID(getSplitViewItemClass().class), objc.Sel("contentListWithViewController:"), viewController)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/init(inspectorWithViewController:)
func NewSplitViewItemInspectorWithViewController(viewController IViewController) SplitViewItem {
	rv := objc.Send[SplitViewItem](objc.ID(getSplitViewItemClass().class), objc.Sel("inspectorWithViewController:"), viewController)
	return rv
}


// Creates a split view item that represents a sidebar for the specified view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/init(sidebarWithViewController:)
func NewSplitViewItemSidebarWithViewController(viewController IViewController) SplitViewItem {
	rv := objc.Send[SplitViewItem](objc.ID(getSplitViewItemClass().class), objc.Sel("sidebarWithViewController:"), viewController)
	return rv
}


// Creates a split view item that represents the specified view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/init(viewController:)
func NewSplitViewItemWithViewController(viewController IViewController) SplitViewItem {
	rv := objc.Send[SplitViewItem](objc.ID(getSplitViewItemClass().class), objc.Sel("splitViewItemWithViewController:"), viewController)
	return rv
}



// Creates a split view item that represents a content list for the specified view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/init(contentListWithViewController:)
func (sc _SplitViewItemClass) ContentListWithViewController(viewController IViewController) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("contentListWithViewController:"), viewController)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/init(inspectorWithViewController:)
func (sc _SplitViewItemClass) InspectorWithViewController(viewController IViewController) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("inspectorWithViewController:"), viewController)
	return rv
}


// Creates a split view item that represents a sidebar for the specified view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/init(sidebarWithViewController:)
func (sc _SplitViewItemClass) SidebarWithViewController(viewController IViewController) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("sidebarWithViewController:"), viewController)
	return rv
}


// Creates a split view item that represents the specified view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/init(viewController:)
func (sc _SplitViewItemClass) SplitViewItemWithViewController(viewController IViewController) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("splitViewItemWithViewController:"), viewController)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/addBottomAlignedAccessoryViewController(_:)
func (s_ SplitViewItem) AddBottomAlignedAccessoryViewController(childViewController ISplitViewItemAccessoryViewController) {
	objc.Send[objc.ID](s_.ID, objc.Sel("addBottomAlignedAccessoryViewController:"), childViewController)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/addTopAlignedAccessoryViewController(_:)
func (s_ SplitViewItem) AddTopAlignedAccessoryViewController(childViewController ISplitViewItemAccessoryViewController) {
	objc.Send[objc.ID](s_.ID, objc.Sel("addTopAlignedAccessoryViewController:"), childViewController)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/insertBottomAlignedAccessoryViewController(_:at:)
func (s_ SplitViewItem) InsertBottomAlignedAccessoryViewControllerAtIndex(childViewController ISplitViewItemAccessoryViewController, index int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("insertBottomAlignedAccessoryViewController:atIndex:"), childViewController, index)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/insertTopAlignedAccessoryViewController(_:at:)
func (s_ SplitViewItem) InsertTopAlignedAccessoryViewControllerAtIndex(childViewController ISplitViewItemAccessoryViewController, index int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("insertTopAlignedAccessoryViewController:atIndex:"), childViewController, index)
}


// NOTE: you can use this method, or , whichever is easier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/removeBottomAlignedAccessoryViewController(at:)
func (s_ SplitViewItem) RemoveBottomAlignedAccessoryViewControllerAtIndex(index int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeBottomAlignedAccessoryViewControllerAtIndex:"), index)
}


// NOTE: you can use this method, or , whichever is easier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/removeTopAlignedAccessoryViewController(at:)
func (s_ SplitViewItem) RemoveTopAlignedAccessoryViewControllerAtIndex(index int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeTopAlignedAccessoryViewControllerAtIndex:"), index)
}


// A Boolean value that indicates whether full-height sidebars appear in the window after you set a style mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/allowsFullHeightLayout
func (s_ SplitViewItem) AllowsFullHeightLayout() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("allowsFullHeightLayout"))
	return rv
}


// A Boolean value that indicates whether full-height sidebars appear in the window after you set a style mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/allowsFullHeightLayout
func (s_ SplitViewItem) SetAllowsFullHeightLayout(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAllowsFullHeightLayout:"), value)
}


// The maximum thickness of the split view item when it resizes due to automatic sizing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/automaticMaximumThickness
func (s_ SplitViewItem) AutomaticMaximumThickness() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("automaticMaximumThickness"))
	return rv
}


// The maximum thickness of the split view item when it resizes due to automatic sizing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/automaticMaximumThickness
func (s_ SplitViewItem) SetAutomaticMaximumThickness(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAutomaticMaximumThickness:"), value)
}


// When YES, other items such as sidebars or inspectors may appear overlaid on top of this item’s and this item’s will be adjusted with respect to overlaid content. Defaults to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/automaticallyAdjustsSafeAreaInsets
func (s_ SplitViewItem) AutomaticallyAdjustsSafeAreaInsets() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("automaticallyAdjustsSafeAreaInsets"))
	return rv
}


// When YES, other items such as sidebars or inspectors may appear overlaid on top of this item’s and this item’s will be adjusted with respect to overlaid content. Defaults to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/automaticallyAdjustsSafeAreaInsets
func (s_ SplitViewItem) SetAutomaticallyAdjustsSafeAreaInsets(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAutomaticallyAdjustsSafeAreaInsets:"), value)
}


// The standard behavior type of the split view item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/behavior-swift.property
func (s_ SplitViewItem) Behavior() SplitViewItemBehavior {
	rv := objc.Send[SplitViewItemBehavior](s_.ID, objc.Sel("behavior"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/bottomAlignedAccessoryViewControllers
func (s_ SplitViewItem) BottomAlignedAccessoryViewControllers() []SplitViewItemAccessoryViewController {
	rv := objc.Send[[]SplitViewItemAccessoryViewController](s_.ID, objc.Sel("bottomAlignedAccessoryViewControllers"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/bottomAlignedAccessoryViewControllers
func (s_ SplitViewItem) SetBottomAlignedAccessoryViewControllers(value []SplitViewItemAccessoryViewController) {
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
	objc.Send[objc.ID](s_.ID, objc.Sel("setBottomAlignedAccessoryViewControllers:"), nsArray)
}


// A Boolean value that determines whether a user interaction can collapse the child view controller that corresponds to the split view item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/canCollapse
func (s_ SplitViewItem) CanCollapse() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("canCollapse"))
	return rv
}


// A Boolean value that determines whether a user interaction can collapse the child view controller that corresponds to the split view item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/canCollapse
func (s_ SplitViewItem) SetCanCollapse(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCanCollapse:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/canCollapseFromWindowResize
func (s_ SplitViewItem) CanCollapseFromWindowResize() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("canCollapseFromWindowResize"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/canCollapseFromWindowResize
func (s_ SplitViewItem) SetCanCollapseFromWindowResize(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCanCollapseFromWindowResize:"), value)
}


// The resizing behavior when the split view item toggles its collapsed state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/collapseBehavior-swift.property
func (s_ SplitViewItem) CollapseBehavior() SplitViewItemCollapseBehavior {
	rv := objc.Send[SplitViewItemCollapseBehavior](s_.ID, objc.Sel("collapseBehavior"))
	return rv
}


// The resizing behavior when the split view item toggles its collapsed state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/collapseBehavior-swift.property
func (s_ SplitViewItem) SetCollapseBehavior(value SplitViewItemCollapseBehavior) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCollapseBehavior:"), value)
}


// The priority for a split view item to hold its size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/holdingPriority
func (s_ SplitViewItem) HoldingPriority() objc.IObject /* cross-framework: LayoutPriority */ {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("holdingPriority"))
	return rv
}


// The priority for a split view item to hold its size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/holdingPriority
func (s_ SplitViewItem) SetHoldingPriority(value objc.IObject /* cross-framework: LayoutPriority */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHoldingPriority:"), value)
}


// A Boolean value that determines whether the child view controller that corresponds to the split view item is in a collapsed state in the split view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/isCollapsed
func (s_ SplitViewItem) Collapsed() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("collapsed"))
	return rv
}


// A Boolean value that determines whether the child view controller that corresponds to the split view item is in a collapsed state in the split view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/isCollapsed
func (s_ SplitViewItem) SetCollapsed(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCollapsed:"), value)
}


// A Boolean value that determines whether the split view item can temporarily expand during a drag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/isSpringLoaded
func (s_ SplitViewItem) SpringLoaded() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("springLoaded"))
	return rv
}


// A Boolean value that determines whether the split view item can temporarily expand during a drag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/isSpringLoaded
func (s_ SplitViewItem) SetSpringLoaded(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSpringLoaded:"), value)
}


// The maximum thickness of the split view item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/maximumThickness
func (s_ SplitViewItem) MaximumThickness() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("maximumThickness"))
	return rv
}


// The maximum thickness of the split view item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/maximumThickness
func (s_ SplitViewItem) SetMaximumThickness(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMaximumThickness:"), value)
}


// The minimum thickness of the split view item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/minimumThickness
func (s_ SplitViewItem) MinimumThickness() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("minimumThickness"))
	return rv
}


// The minimum thickness of the split view item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/minimumThickness
func (s_ SplitViewItem) SetMinimumThickness(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMinimumThickness:"), value)
}


// The preferred thickness of the split view item relative to the split view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/preferredThicknessFraction
func (s_ SplitViewItem) PreferredThicknessFraction() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("preferredThicknessFraction"))
	return rv
}


// The preferred thickness of the split view item relative to the split view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/preferredThicknessFraction
func (s_ SplitViewItem) SetPreferredThicknessFraction(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPreferredThicknessFraction:"), value)
}


// The type of separator that the app displays between the title bar and content of a window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/titlebarSeparatorStyle
func (s_ SplitViewItem) TitlebarSeparatorStyle() TitlebarSeparatorStyle {
	rv := objc.Send[TitlebarSeparatorStyle](s_.ID, objc.Sel("titlebarSeparatorStyle"))
	return rv
}


// The type of separator that the app displays between the title bar and content of a window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/titlebarSeparatorStyle
func (s_ SplitViewItem) SetTitlebarSeparatorStyle(value TitlebarSeparatorStyle) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTitlebarSeparatorStyle:"), value)
}


// The following methods allow you to add accessory views to the top/bottom of this splitViewItem. See for more details.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/topAlignedAccessoryViewControllers
func (s_ SplitViewItem) TopAlignedAccessoryViewControllers() []SplitViewItemAccessoryViewController {
	rv := objc.Send[[]SplitViewItemAccessoryViewController](s_.ID, objc.Sel("topAlignedAccessoryViewControllers"))
	return rv
}


// The following methods allow you to add accessory views to the top/bottom of this splitViewItem. See for more details.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/topAlignedAccessoryViewControllers
func (s_ SplitViewItem) SetTopAlignedAccessoryViewControllers(value []SplitViewItemAccessoryViewController) {
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
	objc.Send[objc.ID](s_.ID, objc.Sel("setTopAlignedAccessoryViewControllers:"), nsArray)
}


// The view controller that the split view item represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/viewController
func (s_ SplitViewItem) ViewController() IViewController {
	rv := objc.Send[ViewController](s_.ID, objc.Sel("viewController"))
	return rv
}


// The view controller that the split view item represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/viewController
func (s_ SplitViewItem) SetViewController(value IViewController) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setViewController:"), value)
}


// A Boolean value that determines whether the child view controller that corresponds to the split view item is in a collapsed state in the split view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitem/iscollapsed
func (s_ SplitViewItem) IsCollapsed() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isCollapsed"))
	return rv
}


// A Boolean value that determines whether the child view controller that corresponds to the split view item is in a collapsed state in the split view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitem/iscollapsed
func (s_ SplitViewItem) SetIsCollapsed(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsCollapsed:"), value)
}


// A Boolean value that determines whether the split view item can temporarily expand during a drag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitem/isspringloaded
func (s_ SplitViewItem) IsSpringLoaded() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isSpringLoaded"))
	return rv
}


// A Boolean value that determines whether the split view item can temporarily expand during a drag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitem/isspringloaded
func (s_ SplitViewItem) SetIsSpringLoaded(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsSpringLoaded:"), value)
}


