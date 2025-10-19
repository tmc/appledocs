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

// An interface definition for the [TabViewItem] class.
type ITabViewItem interface {
	objectivec.IObject
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
// Alloc allocates a new instance without initialization.
func (tc _TabViewItemClass) Alloc() TabViewItem {
	rv := objc.Send[TabViewItem](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (tc _TabViewItemClass) New() TabViewItem {
	rv := objc.Send[TabViewItem](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TabViewItem) Init() TabViewItem {
	rv := objc.Send[TabViewItem](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TabViewItem) Autorelease() TabViewItem {
	rv := objc.Send[TabViewItem](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTabViewItem creates a new TabViewItem instance.
func NewTabViewItem() TabViewItem {
	return tabViewItemClass.New()
}




