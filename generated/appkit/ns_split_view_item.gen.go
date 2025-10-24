// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSSplitViewItem */


/* debug [class_header]: Header for NSSplitViewItem */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SplitViewItem */
// An interface definition for the [SplitViewItem] class.
type ISplitViewItem interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SplitViewItem */
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
	HoldingPriority() LayoutPriority /* typedef */
	SetHoldingPriority(value LayoutPriority /* typedef */)
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SplitViewItem */
	// methods:
	AddBottomAlignedAccessoryViewController(childViewController ISplitViewItemAccessoryViewController)
	AddTopAlignedAccessoryViewController(childViewController ISplitViewItemAccessoryViewController)
	InsertBottomAlignedAccessoryViewControllerAtIndex(childViewController ISplitViewItemAccessoryViewController, index int)
	InsertTopAlignedAccessoryViewControllerAtIndex(childViewController ISplitViewItemAccessoryViewController, index int)
	RemoveBottomAlignedAccessoryViewControllerAtIndex(index int)
	RemoveTopAlignedAccessoryViewControllerAtIndex(index int)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SplitViewItem */
// Alloc allocates a new instance without initialization.
func (sc _SplitViewItemClass) Alloc() SplitViewItem {
	rv := objc.Send[SplitViewItem](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SplitViewItem */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SplitViewItem */

// Creates a split view item that represents a content list for the specified view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/init(contentListWithViewController:)
func NewSplitViewItemContentListWithViewController(viewController IViewController) SplitViewItem {
	rv := objc.Send[SplitViewItem](objc.ID(getSplitViewItemClass().class), objc.Sel("contentListWithViewController:"), viewController)
	return rv
}/* debug [class_init_methods/constructor]: NewSplitViewItemContentListWithViewController */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/init(inspectorWithViewController:)
func NewSplitViewItemInspectorWithViewController(viewController IViewController) SplitViewItem {
	rv := objc.Send[SplitViewItem](objc.ID(getSplitViewItemClass().class), objc.Sel("inspectorWithViewController:"), viewController)
	return rv
}/* debug [class_init_methods/constructor]: NewSplitViewItemInspectorWithViewController */


// Creates a split view item that represents a sidebar for the specified view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/init(sidebarWithViewController:)
func NewSplitViewItemSidebarWithViewController(viewController IViewController) SplitViewItem {
	rv := objc.Send[SplitViewItem](objc.ID(getSplitViewItemClass().class), objc.Sel("sidebarWithViewController:"), viewController)
	return rv
}/* debug [class_init_methods/constructor]: NewSplitViewItemSidebarWithViewController */


// Creates a split view item that represents the specified view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/init(viewController:)
func NewSplitViewItemWithViewController(viewController IViewController) SplitViewItem {
	rv := objc.Send[SplitViewItem](objc.ID(getSplitViewItemClass().class), objc.Sel("splitViewItemWithViewController:"), viewController)
	return rv
}/* debug [class_init_methods/constructor]: NewSplitViewItemWithViewController */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SplitViewItem */

// Creates a split view item that represents a content list for the specified view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/init(contentListWithViewController:)
func (sc _SplitViewItemClass) ContentListWithViewController(viewController IViewController) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(sc.class), objc.Sel("contentListWithViewController:"), viewController)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ContentListWithViewController) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/init(inspectorWithViewController:)
func (sc _SplitViewItemClass) InspectorWithViewController(viewController IViewController) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(sc.class), objc.Sel("inspectorWithViewController:"), viewController)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=InspectorWithViewController) */


// Creates a split view item that represents a sidebar for the specified view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/init(sidebarWithViewController:)
func (sc _SplitViewItemClass) SidebarWithViewController(viewController IViewController) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(sc.class), objc.Sel("sidebarWithViewController:"), viewController)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SidebarWithViewController) */


// Creates a split view item that represents the specified view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/init(viewController:)
func (sc _SplitViewItemClass) SplitViewItemWithViewController(viewController IViewController) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(sc.class), objc.Sel("splitViewItemWithViewController:"), viewController)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SplitViewItemWithViewController) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SplitViewItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SplitViewItem */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/addBottomAlignedAccessoryViewController(_:)
func (s_ SplitViewItem) AddBottomAlignedAccessoryViewController(childViewController ISplitViewItemAccessoryViewController) {
	objc.Send[objc.ID](s_.ID, objc.Sel("addBottomAlignedAccessoryViewController:"), childViewController)
}/* debug [instance_methods/method]: AddBottomAlignedAccessoryViewController */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/addTopAlignedAccessoryViewController(_:)
func (s_ SplitViewItem) AddTopAlignedAccessoryViewController(childViewController ISplitViewItemAccessoryViewController) {
	objc.Send[objc.ID](s_.ID, objc.Sel("addTopAlignedAccessoryViewController:"), childViewController)
}/* debug [instance_methods/method]: AddTopAlignedAccessoryViewController */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/insertBottomAlignedAccessoryViewController(_:at:)
func (s_ SplitViewItem) InsertBottomAlignedAccessoryViewControllerAtIndex(childViewController ISplitViewItemAccessoryViewController, index int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("insertBottomAlignedAccessoryViewController:atIndex:"), childViewController, index)
}/* debug [instance_methods/method]: InsertBottomAlignedAccessoryViewControllerAtIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/insertTopAlignedAccessoryViewController(_:at:)
func (s_ SplitViewItem) InsertTopAlignedAccessoryViewControllerAtIndex(childViewController ISplitViewItemAccessoryViewController, index int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("insertTopAlignedAccessoryViewController:atIndex:"), childViewController, index)
}/* debug [instance_methods/method]: InsertTopAlignedAccessoryViewControllerAtIndex */


// NOTE: you can use this method, or , whichever is easier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/removeBottomAlignedAccessoryViewController(at:)
func (s_ SplitViewItem) RemoveBottomAlignedAccessoryViewControllerAtIndex(index int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeBottomAlignedAccessoryViewControllerAtIndex:"), index)
}/* debug [instance_methods/method]: RemoveBottomAlignedAccessoryViewControllerAtIndex */


// NOTE: you can use this method, or , whichever is easier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/removeTopAlignedAccessoryViewController(at:)
func (s_ SplitViewItem) RemoveTopAlignedAccessoryViewControllerAtIndex(index int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeTopAlignedAccessoryViewControllerAtIndex:"), index)
}/* debug [instance_methods/method]: RemoveTopAlignedAccessoryViewControllerAtIndex */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SplitViewItem */

// A Boolean value that indicates whether full-height sidebars appear in the window after you set a style mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/allowsFullHeightLayout
func (s_ SplitViewItem) AllowsFullHeightLayout() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("allowsFullHeightLayout"))
	return rv
}/* debug [instance_properties/getter]: allowsFullHeightLayout */


// A Boolean value that indicates whether full-height sidebars appear in the window after you set a style mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/allowsFullHeightLayout
func (s_ SplitViewItem) SetAllowsFullHeightLayout(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAllowsFullHeightLayout:"), value)
}/* debug [instance_properties/setter]: allowsFullHeightLayout */


// The maximum thickness of the split view item when it resizes due to automatic sizing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/automaticMaximumThickness
func (s_ SplitViewItem) AutomaticMaximumThickness() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("automaticMaximumThickness"))
	return rv
}/* debug [instance_properties/getter]: automaticMaximumThickness */


// The maximum thickness of the split view item when it resizes due to automatic sizing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/automaticMaximumThickness
func (s_ SplitViewItem) SetAutomaticMaximumThickness(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAutomaticMaximumThickness:"), value)
}/* debug [instance_properties/setter]: automaticMaximumThickness */


// When YES, other items such as sidebars or inspectors may appear overlaid on top of this item’s and this item’s will be adjusted with respect to overlaid content. Defaults to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/automaticallyAdjustsSafeAreaInsets
func (s_ SplitViewItem) AutomaticallyAdjustsSafeAreaInsets() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("automaticallyAdjustsSafeAreaInsets"))
	return rv
}/* debug [instance_properties/getter]: automaticallyAdjustsSafeAreaInsets */


// When YES, other items such as sidebars or inspectors may appear overlaid on top of this item’s and this item’s will be adjusted with respect to overlaid content. Defaults to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/automaticallyAdjustsSafeAreaInsets
func (s_ SplitViewItem) SetAutomaticallyAdjustsSafeAreaInsets(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAutomaticallyAdjustsSafeAreaInsets:"), value)
}/* debug [instance_properties/setter]: automaticallyAdjustsSafeAreaInsets */


// The standard behavior type of the split view item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/behavior-swift.property
func (s_ SplitViewItem) Behavior() SplitViewItemBehavior {
	rv := objc.Send[SplitViewItemBehavior](s_.ID, objc.Sel("behavior"))
	return rv
}/* debug [instance_properties/getter]: behavior */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/bottomAlignedAccessoryViewControllers
func (s_ SplitViewItem) BottomAlignedAccessoryViewControllers() []SplitViewItemAccessoryViewController {
	rv := objc.Send[[]SplitViewItemAccessoryViewController](s_.ID, objc.Sel("bottomAlignedAccessoryViewControllers"))
	return rv
}/* debug [instance_properties/getter]: bottomAlignedAccessoryViewControllers */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/bottomAlignedAccessoryViewControllers
func (s_ SplitViewItem) SetBottomAlignedAccessoryViewControllers(value []SplitViewItemAccessoryViewController) {
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
}/* debug [instance_properties/setter]: bottomAlignedAccessoryViewControllers */


// A Boolean value that determines whether a user interaction can collapse the child view controller that corresponds to the split view item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/canCollapse
func (s_ SplitViewItem) CanCollapse() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("canCollapse"))
	return rv
}/* debug [instance_properties/getter]: canCollapse */


// A Boolean value that determines whether a user interaction can collapse the child view controller that corresponds to the split view item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/canCollapse
func (s_ SplitViewItem) SetCanCollapse(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCanCollapse:"), value)
}/* debug [instance_properties/setter]: canCollapse */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/canCollapseFromWindowResize
func (s_ SplitViewItem) CanCollapseFromWindowResize() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("canCollapseFromWindowResize"))
	return rv
}/* debug [instance_properties/getter]: canCollapseFromWindowResize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/canCollapseFromWindowResize
func (s_ SplitViewItem) SetCanCollapseFromWindowResize(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCanCollapseFromWindowResize:"), value)
}/* debug [instance_properties/setter]: canCollapseFromWindowResize */


// The resizing behavior when the split view item toggles its collapsed state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/collapseBehavior-swift.property
func (s_ SplitViewItem) CollapseBehavior() SplitViewItemCollapseBehavior {
	rv := objc.Send[SplitViewItemCollapseBehavior](s_.ID, objc.Sel("collapseBehavior"))
	return rv
}/* debug [instance_properties/getter]: collapseBehavior */


// The resizing behavior when the split view item toggles its collapsed state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/collapseBehavior-swift.property
func (s_ SplitViewItem) SetCollapseBehavior(value SplitViewItemCollapseBehavior) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCollapseBehavior:"), value)
}/* debug [instance_properties/setter]: collapseBehavior */


// The priority for a split view item to hold its size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/holdingPriority
func (s_ SplitViewItem) HoldingPriority() LayoutPriority /* typedef */ {
	rv := objc.Send[float32](s_.ID, objc.Sel("holdingPriority"))
	return rv
}/* debug [instance_properties/getter]: holdingPriority */


// The priority for a split view item to hold its size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/holdingPriority
func (s_ SplitViewItem) SetHoldingPriority(value LayoutPriority /* typedef */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHoldingPriority:"), value)
}/* debug [instance_properties/setter]: holdingPriority */


// A Boolean value that determines whether the child view controller that corresponds to the split view item is in a collapsed state in the split view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/isCollapsed
func (s_ SplitViewItem) Collapsed() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("collapsed"))
	return rv
}/* debug [instance_properties/getter]: collapsed */


// A Boolean value that determines whether the child view controller that corresponds to the split view item is in a collapsed state in the split view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/isCollapsed
func (s_ SplitViewItem) SetCollapsed(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCollapsed:"), value)
}/* debug [instance_properties/setter]: collapsed */


// A Boolean value that determines whether the split view item can temporarily expand during a drag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/isSpringLoaded
func (s_ SplitViewItem) SpringLoaded() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("springLoaded"))
	return rv
}/* debug [instance_properties/getter]: springLoaded */


// A Boolean value that determines whether the split view item can temporarily expand during a drag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/isSpringLoaded
func (s_ SplitViewItem) SetSpringLoaded(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSpringLoaded:"), value)
}/* debug [instance_properties/setter]: springLoaded */


// The maximum thickness of the split view item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/maximumThickness
func (s_ SplitViewItem) MaximumThickness() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("maximumThickness"))
	return rv
}/* debug [instance_properties/getter]: maximumThickness */


// The maximum thickness of the split view item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/maximumThickness
func (s_ SplitViewItem) SetMaximumThickness(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMaximumThickness:"), value)
}/* debug [instance_properties/setter]: maximumThickness */


// The minimum thickness of the split view item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/minimumThickness
func (s_ SplitViewItem) MinimumThickness() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("minimumThickness"))
	return rv
}/* debug [instance_properties/getter]: minimumThickness */


// The minimum thickness of the split view item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/minimumThickness
func (s_ SplitViewItem) SetMinimumThickness(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMinimumThickness:"), value)
}/* debug [instance_properties/setter]: minimumThickness */


// The preferred thickness of the split view item relative to the split view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/preferredThicknessFraction
func (s_ SplitViewItem) PreferredThicknessFraction() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("preferredThicknessFraction"))
	return rv
}/* debug [instance_properties/getter]: preferredThicknessFraction */


// The preferred thickness of the split view item relative to the split view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/preferredThicknessFraction
func (s_ SplitViewItem) SetPreferredThicknessFraction(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPreferredThicknessFraction:"), value)
}/* debug [instance_properties/setter]: preferredThicknessFraction */


// The type of separator that the app displays between the title bar and content of a window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/titlebarSeparatorStyle
func (s_ SplitViewItem) TitlebarSeparatorStyle() TitlebarSeparatorStyle {
	rv := objc.Send[TitlebarSeparatorStyle](s_.ID, objc.Sel("titlebarSeparatorStyle"))
	return rv
}/* debug [instance_properties/getter]: titlebarSeparatorStyle */


// The type of separator that the app displays between the title bar and content of a window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/titlebarSeparatorStyle
func (s_ SplitViewItem) SetTitlebarSeparatorStyle(value TitlebarSeparatorStyle) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTitlebarSeparatorStyle:"), value)
}/* debug [instance_properties/setter]: titlebarSeparatorStyle */


// The following methods allow you to add accessory views to the top/bottom of this splitViewItem. See for more details.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/topAlignedAccessoryViewControllers
func (s_ SplitViewItem) TopAlignedAccessoryViewControllers() []SplitViewItemAccessoryViewController {
	rv := objc.Send[[]SplitViewItemAccessoryViewController](s_.ID, objc.Sel("topAlignedAccessoryViewControllers"))
	return rv
}/* debug [instance_properties/getter]: topAlignedAccessoryViewControllers */


// The following methods allow you to add accessory views to the top/bottom of this splitViewItem. See for more details.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/topAlignedAccessoryViewControllers
func (s_ SplitViewItem) SetTopAlignedAccessoryViewControllers(value []SplitViewItemAccessoryViewController) {
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
}/* debug [instance_properties/setter]: topAlignedAccessoryViewControllers */


// The view controller that the split view item represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/viewController
func (s_ SplitViewItem) ViewController() IViewController {
	rv := objc.Send[ViewController](s_.ID, objc.Sel("viewController"))
	return rv
}/* debug [instance_properties/getter]: viewController */


// The view controller that the split view item represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/viewController
func (s_ SplitViewItem) SetViewController(value IViewController) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setViewController:"), value)
}/* debug [instance_properties/setter]: viewController */


// A Boolean value that determines whether the child view controller that corresponds to the split view item is in a collapsed state in the split view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitem/iscollapsed
func (s_ SplitViewItem) IsCollapsed() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isCollapsed"))
	return rv
}/* debug [instance_properties/getter]: isCollapsed */


// A Boolean value that determines whether the child view controller that corresponds to the split view item is in a collapsed state in the split view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitem/iscollapsed
func (s_ SplitViewItem) SetIsCollapsed(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsCollapsed:"), value)
}/* debug [instance_properties/setter]: isCollapsed */


// A Boolean value that determines whether the split view item can temporarily expand during a drag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitem/isspringloaded
func (s_ SplitViewItem) IsSpringLoaded() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isSpringLoaded"))
	return rv
}/* debug [instance_properties/getter]: isSpringLoaded */


// A Boolean value that determines whether the split view item can temporarily expand during a drag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitem/isspringloaded
func (s_ SplitViewItem) SetIsSpringLoaded(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsSpringLoaded:"), value)
}/* debug [instance_properties/setter]: isSpringLoaded */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSSplitViewItem */


