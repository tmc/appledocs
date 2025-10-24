// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/vision"
)

/* debug [class.gen.go]: Generating class NSTabView */


/* debug [class_header]: Header for NSTabView */
// The class instance for the [TabView] class.
var (
	TabViewClass     _TabViewClass
	TabViewClassOnce sync.Once
)

func getTabViewClass() _TabViewClass {
	TabViewClassOnce.Do(func() {
		TabViewClass = _TabViewClass{objc.GetClass("NSTabView")}
	})
	return TabViewClass
}

type _TabViewClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TabView */
// An interface definition for the [TabView] class.
type ITabView interface {
	IView
	
/* debug [class_interface_properties]: Properties for TabView */
	// properties:
	AllowsTruncatedLabels() bool
	SetAllowsTruncatedLabels(value bool)
	ContentRect() Rect /* not a class type */
	ControlSize() ControlSize
	SetControlSize(value ControlSize)
	ControlTint() ControlTint
	SetControlTint(value ControlTint)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	Font() IFont
	SetFont(value IFont)
	MinimumSize() Size /* not a class type */
	NumberOfTabViewItems() int
	SelectedTabViewItem() ITabViewItem
	TabPosition() TabPosition
	SetTabPosition(value TabPosition)
	TabViewBorderType() TabViewBorderType
	SetTabViewBorderType(value TabViewBorderType)
	TabViewItems() []TabViewItem
	SetTabViewItems(value []TabViewItem)
	TabViewType() TabViewType
	SetTabViewType(value TabViewType)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TabView */
	// methods:
	AddTabViewItem(tabViewItem ITabViewItem)
	IndexOfTabViewItem(tabViewItem ITabViewItem) int
	IndexOfTabViewItemWithIdentifier(identifier objc.IObject) int
	InsertTabViewItemAtIndex(tabViewItem ITabViewItem, index int)
	RemoveTabViewItem(tabViewItem ITabViewItem)
	SelectFirstTabViewItem(sender objc.IObject)
	SelectLastTabViewItem(sender objc.IObject)
	SelectNextTabViewItem(sender objc.IObject)
	SelectPreviousTabViewItem(sender objc.IObject)
	SelectTabViewItem(tabViewItem ITabViewItem)
	SelectTabViewItemAtIndex(index int)
	SelectTabViewItemWithIdentifier(identifier objc.IObject)
	TabViewItemAtIndex(index int) ITabViewItem
	TabViewItemAtPoint(point vision.Point) ITabViewItem
	TakeSelectedTabViewItemFromSender(sender objc.IObject)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TabView */
// Alloc allocates a new instance without initialization.
func (tc _TabViewClass) Alloc() TabView {
	rv := objc.Send[TabView](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TabViewClass) New() TabView {
	rv := objc.Send[TabView](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TabView) Init() TabView {
	rv := objc.Send[TabView](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TabView) Autorelease() TabView {
	rv := objc.Send[TabView](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTabView creates a new TabView instance.
func NewTabView() TabView {
	return getTabViewClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TabView */
// A multipage interface that displays one page at a time.
//
// A tab view contains a row of tabs that give the appearance of folder tabs, as shown in the . The user selects the desired page by clicking the appropriate tab or using the arrow keys to move between pages. Each page displays a view hierarchy provided by your app.


// A multipage interface that displays one page at a time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView
type TabView struct {
	View
}

// TabViewFrom constructs a [TabView] from an unsafe.Pointer.
//
// A multipage interface that displays one page at a time.
func TabViewFrom(ptr unsafe.Pointer) TabView {
	return TabView{
		View: ViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TabView *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TabView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TabView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TabView */

// Adds the specified tab item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/addTabViewItem(_:)
func (t_ TabView) AddTabViewItem(tabViewItem ITabViewItem) {
	objc.Send[objc.ID](t_.ID, objc.Sel("addTabViewItem:"), tabViewItem)
}/* debug [instance_methods/method]: AddTabViewItem */


// Returns the index of the specified item in the tab view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/indexOfTabViewItem(_:)
func (t_ TabView) IndexOfTabViewItem(tabViewItem ITabViewItem) int {
	rv := objc.Send[int](t_.ID, objc.Sel("indexOfTabViewItem:"), tabViewItem)
	return rv
}/* debug [instance_methods/method]: IndexOfTabViewItem */


// Returns the index of the item that matches the specified identifier or if the item is not found.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/indexOfTabViewItem(withIdentifier:)
func (t_ TabView) IndexOfTabViewItemWithIdentifier(identifier objc.IObject) int {
	rv := objc.Send[int](t_.ID, objc.Sel("indexOfTabViewItemWithIdentifier:"), identifier)
	return rv
}/* debug [instance_methods/method]: IndexOfTabViewItemWithIdentifier */


// Inserts the specified item into the tab view’s array of tab view items at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/insertTabViewItem(_:at:)
func (t_ TabView) InsertTabViewItemAtIndex(tabViewItem ITabViewItem, index int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("insertTabViewItem:atIndex:"), tabViewItem, index)
}/* debug [instance_methods/method]: InsertTabViewItemAtIndex */


// Removes the specified item from the tab view’s array of tab view items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/removeTabViewItem(_:)
func (t_ TabView) RemoveTabViewItem(tabViewItem ITabViewItem) {
	objc.Send[objc.ID](t_.ID, objc.Sel("removeTabViewItem:"), tabViewItem)
}/* debug [instance_methods/method]: RemoveTabViewItem */


// This action method selects the first tab view item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/selectFirstTabViewItem(_:)
func (t_ TabView) SelectFirstTabViewItem(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("selectFirstTabViewItem:"), sender)
}/* debug [instance_methods/method]: SelectFirstTabViewItem */


// This action method selects the last tab view item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/selectLastTabViewItem(_:)
func (t_ TabView) SelectLastTabViewItem(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("selectLastTabViewItem:"), sender)
}/* debug [instance_methods/method]: SelectLastTabViewItem */


// This action method selects the next tab view item in the sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/selectNextTabViewItem(_:)
func (t_ TabView) SelectNextTabViewItem(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("selectNextTabViewItem:"), sender)
}/* debug [instance_methods/method]: SelectNextTabViewItem */


// This action method selects the previous tab view item in the sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/selectPreviousTabViewItem(_:)
func (t_ TabView) SelectPreviousTabViewItem(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("selectPreviousTabViewItem:"), sender)
}/* debug [instance_methods/method]: SelectPreviousTabViewItem */


// Selects the specified tab view item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/selectTabViewItem(_:)
func (t_ TabView) SelectTabViewItem(tabViewItem ITabViewItem) {
	objc.Send[objc.ID](t_.ID, objc.Sel("selectTabViewItem:"), tabViewItem)
}/* debug [instance_methods/method]: SelectTabViewItem */


// Selects the tab view item specified by .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/selectTabViewItem(at:)
func (t_ TabView) SelectTabViewItemAtIndex(index int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("selectTabViewItemAtIndex:"), index)
}/* debug [instance_methods/method]: SelectTabViewItemAtIndex */


// Selects the tab view item specified by .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/selectTabViewItem(withIdentifier:)
func (t_ TabView) SelectTabViewItemWithIdentifier(identifier objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("selectTabViewItemWithIdentifier:"), identifier)
}/* debug [instance_methods/method]: SelectTabViewItemWithIdentifier */


// Returns the tab view item at in the tab view’s array of items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/tabViewItem(at:)-7r3at
func (t_ TabView) TabViewItemAtIndex(index int) ITabViewItem {
	rv := objc.Send[TabViewItem](t_.ID, objc.Sel("tabViewItemAtIndex:"), index)
	return rv
}/* debug [instance_methods/method]: TabViewItemAtIndex */


// Returns the tab view item at the specified point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/tabViewItem(at:)-8gnqw
func (t_ TabView) TabViewItemAtPoint(point vision.Point) ITabViewItem {
	rv := objc.Send[TabViewItem](t_.ID, objc.Sel("tabViewItemAtPoint:"), point)
	return rv
}/* debug [instance_methods/method]: TabViewItemAtPoint */


// Sets the selected tab view item to the selected item obtained from the sender.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/takeSelectedTabViewItemFromSender(_:)
func (t_ TabView) TakeSelectedTabViewItemFromSender(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("takeSelectedTabViewItemFromSender:"), sender)
}/* debug [instance_methods/method]: TakeSelectedTabViewItemFromSender */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TabView */

// A Boolean value that indicates if the tab view allows truncating for labels that don’t fit on a tab.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/allowsTruncatedLabels
func (t_ TabView) AllowsTruncatedLabels() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsTruncatedLabels"))
	return rv
}/* debug [instance_properties/getter]: allowsTruncatedLabels */


// A Boolean value that indicates if the tab view allows truncating for labels that don’t fit on a tab.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/allowsTruncatedLabels
func (t_ TabView) SetAllowsTruncatedLabels(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsTruncatedLabels:"), value)
}/* debug [instance_properties/setter]: allowsTruncatedLabels */


// The rectangle describing the content area of the tab view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/contentRect
func (t_ TabView) ContentRect() Rect /* not a class type */ {
	rv := objc.Send[Rect](t_.ID, objc.Sel("contentRect"))
	return rv
}/* debug [instance_properties/getter]: contentRect */


// The size of the tab view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/controlSize
func (t_ TabView) ControlSize() ControlSize {
	rv := objc.Send[ControlSize](t_.ID, objc.Sel("controlSize"))
	return rv
}/* debug [instance_properties/getter]: controlSize */


// The size of the tab view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/controlSize
func (t_ TabView) SetControlSize(value ControlSize) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setControlSize:"), value)
}/* debug [instance_properties/setter]: controlSize */


// The tab view’s control tint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/controlTint
func (t_ TabView) ControlTint() ControlTint {
	rv := objc.Send[ControlTint](t_.ID, objc.Sel("controlTint"))
	return rv
}/* debug [instance_properties/getter]: controlTint */


// The tab view’s control tint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/controlTint
func (t_ TabView) SetControlTint(value ControlTint) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setControlTint:"), value)
}/* debug [instance_properties/setter]: controlTint */


// The tab view’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/delegate
func (t_ TabView) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The tab view’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/delegate
func (t_ TabView) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// A Boolean value that indicates if the tab view draws a background color when its type is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/drawsBackground
func (t_ TabView) DrawsBackground() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("drawsBackground"))
	return rv
}/* debug [instance_properties/getter]: drawsBackground */


// A Boolean value that indicates if the tab view draws a background color when its type is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/drawsBackground
func (t_ TabView) SetDrawsBackground(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDrawsBackground:"), value)
}/* debug [instance_properties/setter]: drawsBackground */


// The font used for the tab view’s label text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/font
func (t_ TabView) Font() IFont {
	rv := objc.Send[Font](t_.ID, objc.Sel("font"))
	return rv
}/* debug [instance_properties/getter]: font */


// The font used for the tab view’s label text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/font
func (t_ TabView) SetFont(value IFont) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFont:"), value)
}/* debug [instance_properties/setter]: font */


// The minimum size necessary for the tab view to display tabs in a useful way.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/minimumSize
func (t_ TabView) MinimumSize() Size /* not a class type */ {
	rv := objc.Send[Size](t_.ID, objc.Sel("minimumSize"))
	return rv
}/* debug [instance_properties/getter]: minimumSize */


// The number of items in the tab view’s array of tab view items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/numberOfTabViewItems
func (t_ TabView) NumberOfTabViewItems() int {
	rv := objc.Send[int](t_.ID, objc.Sel("numberOfTabViewItems"))
	return rv
}/* debug [instance_properties/getter]: numberOfTabViewItems */


// The tab view item for the currently selected tab.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/selectedTabViewItem
func (t_ TabView) SelectedTabViewItem() ITabViewItem {
	rv := objc.Send[TabViewItem](t_.ID, objc.Sel("selectedTabViewItem"))
	return rv
}/* debug [instance_properties/getter]: selectedTabViewItem */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/tabPosition-swift.property
func (t_ TabView) TabPosition() TabPosition {
	rv := objc.Send[TabPosition](t_.ID, objc.Sel("tabPosition"))
	return rv
}/* debug [instance_properties/getter]: tabPosition */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/tabPosition-swift.property
func (t_ TabView) SetTabPosition(value TabPosition) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTabPosition:"), value)
}/* debug [instance_properties/setter]: tabPosition */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/tabViewBorderType-swift.property
func (t_ TabView) TabViewBorderType() TabViewBorderType {
	rv := objc.Send[TabViewBorderType](t_.ID, objc.Sel("tabViewBorderType"))
	return rv
}/* debug [instance_properties/getter]: tabViewBorderType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/tabViewBorderType-swift.property
func (t_ TabView) SetTabViewBorderType(value TabViewBorderType) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTabViewBorderType:"), value)
}/* debug [instance_properties/setter]: tabViewBorderType */


// The tab view’s array of tab view items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/tabViewItems
func (t_ TabView) TabViewItems() []TabViewItem {
	rv := objc.Send[[]TabViewItem](t_.ID, objc.Sel("tabViewItems"))
	return rv
}/* debug [instance_properties/getter]: tabViewItems */


// The tab view’s array of tab view items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/tabViewItems
func (t_ TabView) SetTabViewItems(value []TabViewItem) {
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
}/* debug [instance_properties/setter]: tabViewItems */


// The tab type to display the tabs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/tabViewType
func (t_ TabView) TabViewType() TabViewType {
	rv := objc.Send[TabViewType](t_.ID, objc.Sel("tabViewType"))
	return rv
}/* debug [instance_properties/getter]: tabViewType */


// The tab type to display the tabs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/tabViewType
func (t_ TabView) SetTabViewType(value TabViewType) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTabViewType:"), value)
}/* debug [instance_properties/setter]: tabViewType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSTabView */



