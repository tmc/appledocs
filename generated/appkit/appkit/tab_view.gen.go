// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TabView] class.
var TabViewClass objc.Class

func init() {
	TabViewClass = objc.GetClass("NSTabView")
}

type TabView struct {
	objc.ID
}

func TabViewFrom(ptr unsafe.Pointer) TabView {
	return TabView{
		ID: objc.ID(ptr),
	}
}


// Adds the specified tab item. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/addTabViewItem(_:)
func (t_ TabView) AddTabViewItem(tabViewItem unsafe.Pointer) {
	sel := objc.RegisterName("addTabViewItem:")
	t_.ID.Send(sel, tabViewItem)
}
// Returns the index of the specified item in the tab view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/indexOfTabViewItem(_:)
func (t_ TabView) IndexOfTabViewItem(tabViewItem unsafe.Pointer) int {
	sel := objc.RegisterName("indexOfTabViewItem:")
	ret := t_.ID.Send(sel, tabViewItem)
	return int(ret)
}
// Returns the index of the item that matches the specified identifier or   if the item is not found. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/indexOfTabViewItem(withIdentifier:)
func (t_ TabView) IndexOfTabViewItemWithIdentifier(identifier objc.ID) int {
	sel := objc.RegisterName("indexOfTabViewItemWithIdentifier:")
	ret := t_.ID.Send(sel, identifier)
	return int(ret)
}
// Inserts the specified item into the tab view’s array of tab view items at the specified index. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/insertTabViewItem(_:at:)
func (t_ TabView) InsertTabViewItemAtIndex(tabViewItem unsafe.Pointer, index int) {
	sel := objc.RegisterName("insertTabViewItem:atIndex:")
	t_.ID.Send(sel, tabViewItem, index)
}
// Removes the specified item from the tab view’s array of tab view items. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/removeTabViewItem(_:)
func (t_ TabView) RemoveTabViewItem(tabViewItem unsafe.Pointer) {
	sel := objc.RegisterName("removeTabViewItem:")
	t_.ID.Send(sel, tabViewItem)
}
// This action method selects the first tab view item. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/selectFirstTabViewItem(_:)
func (t_ TabView) SelectFirstTabViewItem(sender objc.ID) {
	sel := objc.RegisterName("selectFirstTabViewItem:")
	t_.ID.Send(sel, sender)
}
// This action method selects the last tab view item. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/selectLastTabViewItem(_:)
func (t_ TabView) SelectLastTabViewItem(sender objc.ID) {
	sel := objc.RegisterName("selectLastTabViewItem:")
	t_.ID.Send(sel, sender)
}
// This action method selects the next tab view item in the sequence. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/selectNextTabViewItem(_:)
func (t_ TabView) SelectNextTabViewItem(sender objc.ID) {
	sel := objc.RegisterName("selectNextTabViewItem:")
	t_.ID.Send(sel, sender)
}
// This action method selects the previous tab view item in the sequence. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/selectPreviousTabViewItem(_:)
func (t_ TabView) SelectPreviousTabViewItem(sender objc.ID) {
	sel := objc.RegisterName("selectPreviousTabViewItem:")
	t_.ID.Send(sel, sender)
}
// Selects the specified tab view item. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/selectTabViewItem(_:)
func (t_ TabView) SelectTabViewItem(tabViewItem unsafe.Pointer) {
	sel := objc.RegisterName("selectTabViewItem:")
	t_.ID.Send(sel, tabViewItem)
}
// Selects the tab view item specified by  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/selectTabViewItem(at:)
func (t_ TabView) SelectTabViewItemAtIndex(index int) {
	sel := objc.RegisterName("selectTabViewItemAtIndex:")
	t_.ID.Send(sel, index)
}
// Selects the tab view item specified by  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/selectTabViewItem(withIdentifier:)
func (t_ TabView) SelectTabViewItemWithIdentifier(identifier objc.ID) {
	sel := objc.RegisterName("selectTabViewItemWithIdentifier:")
	t_.ID.Send(sel, identifier)
}
// Returns the tab view item at   in the tab view’s array of items. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/tabViewItem(at:)-7r3at
func (t_ TabView) TabViewItemAtIndex(index int) unsafe.Pointer {
	sel := objc.RegisterName("tabViewItemAtIndex:")
	ret := t_.ID.Send(sel, index)
	return unsafe.Pointer(ret)
}
// Returns the tab view item at the specified point. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/tabViewItem(at:)-8gnqw
func (t_ TabView) TabViewItemAtPoint(point foundation.Point) unsafe.Pointer {
	sel := objc.RegisterName("tabViewItemAtPoint:")
	ret := t_.ID.Send(sel, point)
	return unsafe.Pointer(ret)
}
// Sets the selected tab view item to the selected item obtained from the sender. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTabView/takeSelectedTabViewItemFromSender(_:)
func (t_ TabView) TakeSelectedTabViewItemFromSender(sender objc.ID) {
	sel := objc.RegisterName("takeSelectedTabViewItemFromSender:")
	t_.ID.Send(sel, sender)
}

