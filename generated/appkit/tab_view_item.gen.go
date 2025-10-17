// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TabViewItem] class.
var tabViewItemClass = _TabViewItemClass{objc.GetClass("NSTabViewItem")}

type _TabViewItemClass struct {
	class objc.Class
}

// An item in a tab view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem

type TabViewItem struct {
	objectivec.Object
}

// TabViewItemFrom constructs a [TabViewItem] from an unsafe.Pointer.
//
// An item in a tab view.
func TabViewItemFrom(ptr unsafe.Pointer) TabViewItem {
	return TabViewItem{objectivec.Object{objc.ID(ptr)}}
}



