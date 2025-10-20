// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [TabView] class.
var (
	tabViewClass     _TabViewClass
	tabViewClassOnce sync.Once
)

func getTabViewClass() _TabViewClass {
	tabViewClassOnce.Do(func() {
		tabViewClass = _TabViewClass{objc.GetClass("NSTabView")}
	})
	return tabViewClass
}

type _TabViewClass struct {
	class objc.Class
}

// An interface definition for the [TabView] class.
type ITabView interface {
	IView
	AddTabViewItem(tabViewItem unsafe.Pointer)
	IndexOfTabViewItem(tabViewItem unsafe.Pointer) int
	IndexOfTabViewItemWithIdentifier(identifier objc.ID) int
	InsertTabViewItemAtIndex(tabViewItem unsafe.Pointer, index int)
	RemoveTabViewItem(tabViewItem unsafe.Pointer)
	SelectFirstTabViewItem(sender objc.ID)
	SelectLastTabViewItem(sender objc.ID)
	SelectNextTabViewItem(sender objc.ID)
	SelectPreviousTabViewItem(sender objc.ID)
	SelectTabViewItem(tabViewItem unsafe.Pointer)
	SelectTabViewItemAtIndex(index int)
	SelectTabViewItemWithIdentifier(identifier objc.ID)
	TabViewItemAtIndex(index int) unsafe.Pointer
	TabViewItemAtPoint(point coregraphics.CGPoint) unsafe.Pointer
	TakeSelectedTabViewItemFromSender(sender objc.ID)
}

// A multipage interface that displays one page at a time.
//
// A tab view contains a row of tabs that give the appearance of folder tabs, as shown in the . The user selects the desired page by clicking the appropriate tab or using the arrow keys to move between pages. Each page displays a view hierarchy provided by your app.
//
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

// Alloc allocates a new instance without initialization.
func (tc _TabViewClass) Alloc() TabView {
	rv := objc.Send[TabView](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Adds the specified tab item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/addTabViewItem(_:)
func (t_ TabView) AddTabViewItem(tabViewItem unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("addTabViewItem:"), tabViewItem)
}

// Returns the index of the specified item in the tab view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/indexOfTabViewItem(_:)
func (t_ TabView) IndexOfTabViewItem(tabViewItem unsafe.Pointer) int {
	rv := objc.Send[int](t_.ID, objc.Sel("indexOfTabViewItem:"), tabViewItem)
	return rv
}

// Returns the index of the item that matches the specified identifier or if the item is not found.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/indexOfTabViewItem(withIdentifier:)
func (t_ TabView) IndexOfTabViewItemWithIdentifier(identifier objc.ID) int {
	rv := objc.Send[int](t_.ID, objc.Sel("indexOfTabViewItemWithIdentifier:"), identifier)
	return rv
}

// Inserts the specified item into the tab view’s array of tab view items at the specified index.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/insertTabViewItem(_:at:)
func (t_ TabView) InsertTabViewItemAtIndex(tabViewItem unsafe.Pointer, index int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("insertTabViewItem:atIndex:"), tabViewItem, index)
}

// Removes the specified item from the tab view’s array of tab view items.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/removeTabViewItem(_:)
func (t_ TabView) RemoveTabViewItem(tabViewItem unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("removeTabViewItem:"), tabViewItem)
}

// This action method selects the first tab view item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/selectFirstTabViewItem(_:)
func (t_ TabView) SelectFirstTabViewItem(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("selectFirstTabViewItem:"), sender)
}

// This action method selects the last tab view item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/selectLastTabViewItem(_:)
func (t_ TabView) SelectLastTabViewItem(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("selectLastTabViewItem:"), sender)
}

// This action method selects the next tab view item in the sequence.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/selectNextTabViewItem(_:)
func (t_ TabView) SelectNextTabViewItem(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("selectNextTabViewItem:"), sender)
}

// This action method selects the previous tab view item in the sequence.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/selectPreviousTabViewItem(_:)
func (t_ TabView) SelectPreviousTabViewItem(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("selectPreviousTabViewItem:"), sender)
}

// Selects the specified tab view item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/selectTabViewItem(_:)
func (t_ TabView) SelectTabViewItem(tabViewItem unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("selectTabViewItem:"), tabViewItem)
}

// Selects the tab view item specified by .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/selectTabViewItem(at:)
func (t_ TabView) SelectTabViewItemAtIndex(index int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("selectTabViewItemAtIndex:"), index)
}

// Selects the tab view item specified by .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/selectTabViewItem(withIdentifier:)
func (t_ TabView) SelectTabViewItemWithIdentifier(identifier objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("selectTabViewItemWithIdentifier:"), identifier)
}

// Returns the tab view item at in the tab view’s array of items.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/tabViewItem(at:)-7r3at
func (t_ TabView) TabViewItemAtIndex(index int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("tabViewItemAtIndex:"), index)
	return rv
}

// Returns the tab view item at the specified point.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/tabViewItem(at:)-8gnqw
func (t_ TabView) TabViewItemAtPoint(point coregraphics.CGPoint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("tabViewItemAtPoint:"), point)
	return rv
}

// Sets the selected tab view item to the selected item obtained from the sender.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/takeSelectedTabViewItemFromSender(_:)
func (t_ TabView) TakeSelectedTabViewItemFromSender(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("takeSelectedTabViewItemFromSender:"), sender)
}

// A Boolean value that indicates if the tab view allows truncating for labels that don’t fit on a tab.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/allowsTruncatedLabels
func (t_ TabView) AllowsTruncatedLabels() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsTruncatedLabels"))
	return rv
}


// SetAllowsTruncatedLabels sets the value of the allowsTruncatedLabels property.
// A Boolean value that indicates if the tab view allows truncating for labels that don’t fit on a tab.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/allowsTruncatedLabels
func (t_ TabView) SetAllowsTruncatedLabels(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsTruncatedLabels:"), value)
}
// The rectangle describing the content area of the tab view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/contentRect
func (t_ TabView) ContentRect() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](t_.ID, objc.Sel("contentRect"))
	return rv
}

// The size of the tab view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/controlSize
func (t_ TabView) ControlSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("controlSize"))
	return rv
}


// SetControlSize sets the value of the controlSize property.
// The size of the tab view.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/controlSize
func (t_ TabView) SetControlSize(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setControlSize:"), value)
}
// The tab view’s control tint.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/controlTint
func (t_ TabView) ControlTint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("controlTint"))
	return rv
}


// SetControlTint sets the value of the controlTint property.
// The tab view’s control tint.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/controlTint
func (t_ TabView) SetControlTint(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setControlTint:"), value)
}
// The tab view’s delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/delegate
func (t_ TabView) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The tab view’s delegate.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/delegate
func (t_ TabView) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDelegate:"), value)
}
// A Boolean value that indicates if the tab view draws a background color when its type is .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/drawsBackground
func (t_ TabView) DrawsBackground() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("drawsBackground"))
	return rv
}


// SetDrawsBackground sets the value of the drawsBackground property.
// A Boolean value that indicates if the tab view draws a background color when its type is .

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/drawsBackground
func (t_ TabView) SetDrawsBackground(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDrawsBackground:"), value)
}
// The font used for the tab view’s label text.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/font
func (t_ TabView) Font() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("font"))
	return rv
}


// SetFont sets the value of the font property.
// The font used for the tab view’s label text.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/font
func (t_ TabView) SetFont(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFont:"), value)
}
// The minimum size necessary for the tab view to display tabs in a useful way.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/minimumSize
func (t_ TabView) MinimumSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](t_.ID, objc.Sel("minimumSize"))
	return rv
}

// The number of items in the tab view’s array of tab view items.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/numberOfTabViewItems
func (t_ TabView) NumberOfTabViewItems() int {
	rv := objc.Send[int](t_.ID, objc.Sel("numberOfTabViewItems"))
	return rv
}

// The tab view item for the currently selected tab.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/selectedTabViewItem
func (t_ TabView) SelectedTabViewItem() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("selectedTabViewItem"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/tabPosition-swift.property
func (t_ TabView) TabPosition() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("tabPosition"))
	return rv
}


// SetTabPosition sets the value of the tabPosition property.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/tabPosition-swift.property
func (t_ TabView) SetTabPosition(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTabPosition:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/tabViewBorderType-swift.property
func (t_ TabView) TabViewBorderType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("tabViewBorderType"))
	return rv
}


// SetTabViewBorderType sets the value of the tabViewBorderType property.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/tabViewBorderType-swift.property
func (t_ TabView) SetTabViewBorderType(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTabViewBorderType:"), value)
}
// The tab view’s array of tab view items.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/tabViewItems
func (t_ TabView) TabViewItems() []__kindof NSTabViewItem {
	rv := objc.Send[[]__kindof NSTabViewItem](t_.ID, objc.Sel("tabViewItems"))
	return rv
}


// SetTabViewItems sets the value of the tabViewItems property.
// The tab view’s array of tab view items.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/tabViewItems
func (t_ TabView) SetTabViewItems(value []__kindof NSTabViewItem) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTabViewItems:"), value)
}
// The tab type to display the tabs.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/tabViewType
func (t_ TabView) TabViewType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("tabViewType"))
	return rv
}


// SetTabViewType sets the value of the tabViewType property.
// The tab type to display the tabs.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/tabViewType
func (t_ TabView) SetTabViewType(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTabViewType:"), value)
}


