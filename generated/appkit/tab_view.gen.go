
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TabView] class.
var TabViewClass _TabViewClass

func init() {
	TabViewClass = _TabViewClass{objc.GetClass("NSTabView")}
}

type _TabViewClass struct {
	objc.Class
}

// An interface definition for the [TabView] class.
type ITabView interface {
	ID() objc.ID
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
	TabViewItemAtPoint(point unsafe.Pointer) unsafe.Pointer
	TakeSelectedTabViewItemFromSender(sender objc.ID)
}

type TabView struct {
	id objc.ID
}

func TabViewFrom(ptr unsafe.Pointer) TabView {
	return TabView{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TabView) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TabViewClass) Alloc() TabView {
	rv := objc.Send[TabView](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TabViewClass) New() TabView {
	rv := objc.Send[TabView](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTabView creates and returns a new initialized instance.
func NewTabView() TabView {
	return TabViewClass.New()
}

// Init initializes the instance.
func (t_ TabView) Init() TabView {
	rv := objc.Send[TabView](t_.ID(), selInit)
	return rv
}
// Adds the specified tab item. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/addTabViewItem(_:)
func (t_ TabView) AddTabViewItem(tabViewItem unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("addTabViewItem:"), tabViewItem)
}
// Returns the index of the specified item in the tab view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/indexOfTabViewItem(_:)
func (t_ TabView) IndexOfTabViewItem(tabViewItem unsafe.Pointer) int {
	rv := objc.Send[int](t_.ID(), objc.RegisterName("indexOfTabViewItem:"), tabViewItem)
	return rv
}
// Returns the index of the item that matches the specified identifier or   if the item is not found. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/indexOfTabViewItem(withIdentifier:)
func (t_ TabView) IndexOfTabViewItemWithIdentifier(identifier objc.ID) int {
	rv := objc.Send[int](t_.ID(), objc.RegisterName("indexOfTabViewItemWithIdentifier:"), identifier)
	return rv
}
// Inserts the specified item into the tab view’s array of tab view items at the specified index. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/insertTabViewItem(_:at:)
func (t_ TabView) InsertTabViewItemAtIndex(tabViewItem unsafe.Pointer, index int) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("insertTabViewItem:atIndex:"), tabViewItem, index)
}
// Removes the specified item from the tab view’s array of tab view items. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/removeTabViewItem(_:)
func (t_ TabView) RemoveTabViewItem(tabViewItem unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("removeTabViewItem:"), tabViewItem)
}
// This action method selects the first tab view item. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/selectFirstTabViewItem(_:)
func (t_ TabView) SelectFirstTabViewItem(sender objc.ID) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("selectFirstTabViewItem:"), sender)
}
// This action method selects the last tab view item. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/selectLastTabViewItem(_:)
func (t_ TabView) SelectLastTabViewItem(sender objc.ID) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("selectLastTabViewItem:"), sender)
}
// This action method selects the next tab view item in the sequence. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/selectNextTabViewItem(_:)
func (t_ TabView) SelectNextTabViewItem(sender objc.ID) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("selectNextTabViewItem:"), sender)
}
// This action method selects the previous tab view item in the sequence. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/selectPreviousTabViewItem(_:)
func (t_ TabView) SelectPreviousTabViewItem(sender objc.ID) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("selectPreviousTabViewItem:"), sender)
}
// Selects the specified tab view item. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/selectTabViewItem(_:)
func (t_ TabView) SelectTabViewItem(tabViewItem unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("selectTabViewItem:"), tabViewItem)
}
// Selects the tab view item specified by  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/selectTabViewItem(at:)
func (t_ TabView) SelectTabViewItemAtIndex(index int) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("selectTabViewItemAtIndex:"), index)
}
// Selects the tab view item specified by  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/selectTabViewItem(withIdentifier:)
func (t_ TabView) SelectTabViewItemWithIdentifier(identifier objc.ID) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("selectTabViewItemWithIdentifier:"), identifier)
}
// Returns the tab view item at   in the tab view’s array of items. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/tabViewItem(at:)-7r3at
func (t_ TabView) TabViewItemAtIndex(index int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("tabViewItemAtIndex:"), index)
	return rv
}
// Returns the tab view item at the specified point. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/tabViewItem(at:)-8gnqw
func (t_ TabView) TabViewItemAtPoint(point unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("tabViewItemAtPoint:"), point)
	return rv
}
// Sets the selected tab view item to the selected item obtained from the sender. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/takeSelectedTabViewItemFromSender(_:)
func (t_ TabView) TakeSelectedTabViewItemFromSender(sender objc.ID) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("takeSelectedTabViewItemFromSender:"), sender)
}
// A Boolean value that indicates if the tab view allows truncating for labels that don’t fit on a tab. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/allowsTruncatedLabels
func (t_ TabView) AllowsTruncatedLabels() bool {
	rv := objc.Send[bool](t_.ID(), objc.RegisterName("allowsTruncatedLabels"))
	return rv
}
// SetAllowsTruncatedLabels sets the value of the allowsTruncatedLabels property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/allowsTruncatedLabels
func (t_ TabView) SetAllowsTruncatedLabels(value bool) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setAllowsTruncatedLabels:"), value)
}
// The rectangle describing the content area of the tab view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/contentRect
func (t_ TabView) ContentRect() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("contentRect"))
	return rv
}
// The size of the tab view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/controlSize
func (t_ TabView) ControlSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("controlSize"))
	return rv
}
// SetControlSize sets the value of the controlSize property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/controlSize
func (t_ TabView) SetControlSize(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setControlSize:"), value)
}
// The tab view’s control tint. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/controlTint
func (t_ TabView) ControlTint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("controlTint"))
	return rv
}
// SetControlTint sets the value of the controlTint property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/controlTint
func (t_ TabView) SetControlTint(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setControlTint:"), value)
}
// The tab view’s delegate. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/delegate
func (t_ TabView) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("delegate"))
	return rv
}
// SetDelegate sets the value of the delegate property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/delegate
func (t_ TabView) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setDelegate:"), value)
}
// A Boolean value that indicates if the tab view draws a background color when its type is  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/drawsBackground
func (t_ TabView) DrawsBackground() bool {
	rv := objc.Send[bool](t_.ID(), objc.RegisterName("drawsBackground"))
	return rv
}
// SetDrawsBackground sets the value of the drawsBackground property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/drawsBackground
func (t_ TabView) SetDrawsBackground(value bool) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setDrawsBackground:"), value)
}
// The font used for the tab view’s label text. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/font
func (t_ TabView) Font() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("font"))
	return rv
}
// SetFont sets the value of the font property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/font
func (t_ TabView) SetFont(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setFont:"), value)
}
// The minimum size necessary for the tab view to display tabs in a useful way. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/minimumSize
func (t_ TabView) MinimumSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("minimumSize"))
	return rv
}
// The number of items in the tab view’s array of tab view items. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/numberOfTabViewItems
func (t_ TabView) NumberOfTabViewItems() int {
	rv := objc.Send[int](t_.ID(), objc.RegisterName("numberOfTabViewItems"))
	return rv
}
// The tab view item for the currently selected tab. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/selectedTabViewItem
func (t_ TabView) SelectedTabViewItem() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("selectedTabViewItem"))
	return rv
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/tabPosition-swift.property
func (t_ TabView) TabPosition() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("tabPosition"))
	return rv
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/tabPosition-swift.property
func (t_ TabView) SetTabPosition(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setTabPosition:"), value)
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/tabViewBorderType-swift.property
func (t_ TabView) TabViewBorderType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("tabViewBorderType"))
	return rv
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/tabViewBorderType-swift.property
func (t_ TabView) SetTabViewBorderType(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setTabViewBorderType:"), value)
}
// The tab view’s array of tab view items. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/tabViewItems
func (t_ TabView) TabViewItems() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("tabViewItems"))
	return rv
}
// SetTabViewItems sets the value of the tabViewItems property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/tabViewItems
func (t_ TabView) SetTabViewItems(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setTabViewItems:"), value)
}
// The tab type to display the tabs. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/tabViewType
func (t_ TabView) TabViewType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("tabViewType"))
	return rv
}
// SetTabViewType sets the value of the tabViewType property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/tabViewType
func (t_ TabView) SetTabViewType(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setTabViewType:"), value)
}
