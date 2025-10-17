// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TabView] class.
var tabViewClass = _TabViewClass{objc.GetClass("NSTabView")}

type _TabViewClass struct {
	class objc.Class
}

// A multipage interface that displays one page at a time. [Full Topic]
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

// Adds the specified tab item. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/addTabViewItem(_:)
func (t_ TabView) AddTabViewItem(tabViewItem unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("addTabViewItem:"), tabViewItem)
}
// Returns the index of the specified item in the tab view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/indexOfTabViewItem(_:)
func (t_ TabView) IndexOfTabViewItem(tabViewItem unsafe.Pointer) int {
	rv := objc.Send[int](t_.ID, objc.Sel("indexOfTabViewItem:"), tabViewItem)
	return rv
}
// Returns the index of the item that matches the specified identifier or if the item is not found. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/indexOfTabViewItem(withIdentifier:)
func (t_ TabView) IndexOfTabViewItemWithIdentifier(identifier objc.ID) int {
	rv := objc.Send[int](t_.ID, objc.Sel("indexOfTabViewItemWithIdentifier:"), identifier)
	return rv
}
// Inserts the specified item into the tab view’s array of tab view items at the specified index. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/insertTabViewItem(_:at:)
func (t_ TabView) InsertTabViewItemAtIndex(tabViewItem unsafe.Pointer, index int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("insertTabViewItem:atIndex:"), tabViewItem, index)
}
// Removes the specified item from the tab view’s array of tab view items. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/removeTabViewItem(_:)
func (t_ TabView) RemoveTabViewItem(tabViewItem unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("removeTabViewItem:"), tabViewItem)
}
// This action method selects the first tab view item. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/selectFirstTabViewItem(_:)
func (t_ TabView) SelectFirstTabViewItem(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("selectFirstTabViewItem:"), sender)
}
// This action method selects the last tab view item. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/selectLastTabViewItem(_:)
func (t_ TabView) SelectLastTabViewItem(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("selectLastTabViewItem:"), sender)
}
// This action method selects the next tab view item in the sequence. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/selectNextTabViewItem(_:)
func (t_ TabView) SelectNextTabViewItem(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("selectNextTabViewItem:"), sender)
}
// This action method selects the previous tab view item in the sequence. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/selectPreviousTabViewItem(_:)
func (t_ TabView) SelectPreviousTabViewItem(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("selectPreviousTabViewItem:"), sender)
}
// Selects the specified tab view item. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/selectTabViewItem(_:)
func (t_ TabView) SelectTabViewItem(tabViewItem unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("selectTabViewItem:"), tabViewItem)
}
// Selects the tab view item specified by . [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/selectTabViewItem(at:)
func (t_ TabView) SelectTabViewItemAtIndex(index int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("selectTabViewItemAtIndex:"), index)
}
// Selects the tab view item specified by . [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/selectTabViewItem(withIdentifier:)
func (t_ TabView) SelectTabViewItemWithIdentifier(identifier objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("selectTabViewItemWithIdentifier:"), identifier)
}
// Returns the tab view item at in the tab view’s array of items. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/tabViewItem(at:)-7r3at
func (t_ TabView) TabViewItemAtIndex(index int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("tabViewItemAtIndex:"), index)
	return rv
}
// Returns the tab view item at the specified point. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/tabViewItem(at:)-8gnqw
func (t_ TabView) TabViewItemAtPoint(point unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("tabViewItemAtPoint:"), point)
	return rv
}
// Sets the selected tab view item to the selected item obtained from the sender. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/takeSelectedTabViewItemFromSender(_:)
func (t_ TabView) TakeSelectedTabViewItemFromSender(sender objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("takeSelectedTabViewItemFromSender:"), sender)
}


