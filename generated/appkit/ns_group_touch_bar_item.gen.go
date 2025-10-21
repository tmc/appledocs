// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [GroupTouchBarItem] class.
var (
	GroupTouchBarItemClass     _GroupTouchBarItemClass
	GroupTouchBarItemClassOnce sync.Once
)

func getGroupTouchBarItemClass() _GroupTouchBarItemClass {
	GroupTouchBarItemClassOnce.Do(func() {
		GroupTouchBarItemClass = _GroupTouchBarItemClass{objc.GetClass("NSGroupTouchBarItem")}
	})
	return GroupTouchBarItemClass
}

type _GroupTouchBarItemClass struct {
	class objc.Class
}

// An interface definition for the [GroupTouchBarItem] class.
type IGroupTouchBarItem interface {
	ITouchBarItem
}

// A bar item that provides a bar to contain other items.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem
type GroupTouchBarItem struct {
	TouchBarItem
}

// GroupTouchBarItemFrom constructs a [GroupTouchBarItem] from an unsafe.Pointer.
//
// A bar item that provides a bar to contain other items.
func GroupTouchBarItemFrom(ptr unsafe.Pointer) GroupTouchBarItem {
	return GroupTouchBarItem{
		TouchBarItem: TouchBarItemFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (gc _GroupTouchBarItemClass) Alloc() GroupTouchBarItem {
	rv := objc.Send[GroupTouchBarItem](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GroupTouchBarItemClass) New() GroupTouchBarItem {
	rv := objc.Send[GroupTouchBarItem](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GroupTouchBarItem) Init() GroupTouchBarItem {
	rv := objc.Send[GroupTouchBarItem](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GroupTouchBarItem) Autorelease() GroupTouchBarItem {
	rv := objc.Send[GroupTouchBarItem](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGroupTouchBarItem creates a new GroupTouchBarItem instance.
func NewGroupTouchBarItem() GroupTouchBarItem {
	return getGroupTouchBarItemClass().New()
}




// Initializes and returns a group item whose bar is constructed from the supplied items.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/init(identifier:items:)
func NewGroupTouchBarItemGroupItemWithIdentifierItems(identifier unsafe.Pointer, items unsafe.Pointer) GroupTouchBarItem {
	rv := objc.Send[GroupTouchBarItem](objc.ID(getGroupTouchBarItemClass().class), objc.Sel("groupItemWithIdentifier:items:"), identifier, items)
	return rv
}



// Initializes and returns a group item whose bar is constructed from the supplied items, and with the specified compression options.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/init(identifier:items:allowedCompressionOptions:)
func NewGroupTouchBarItemGroupItemWithIdentifierItemsAllowedCompressionOptions(identifier unsafe.Pointer, items unsafe.Pointer, allowedCompressionOptions unsafe.Pointer) GroupTouchBarItem {
	rv := objc.Send[GroupTouchBarItem](objc.ID(getGroupTouchBarItemClass().class), objc.Sel("groupItemWithIdentifier:items:allowedCompressionOptions:"), identifier, items, allowedCompressionOptions)
	return rv
}


// Initializes and returns a group item whose bar is constructed from the supplied items.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/init(identifier:items:)
func (gc _GroupTouchBarItemClass) GroupItemWithIdentifierItems(identifier unsafe.Pointer, items unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("groupItemWithIdentifier:items:"), identifier, items)
	return rv
}

// Initializes and returns a group item whose bar is constructed from the supplied items, and with the specified compression options.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/init(identifier:items:allowedCompressionOptions:)
func (gc _GroupTouchBarItemClass) GroupItemWithIdentifierItemsAllowedCompressionOptions(identifier unsafe.Pointer, items unsafe.Pointer, allowedCompressionOptions unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("groupItemWithIdentifier:items:allowedCompressionOptions:"), identifier, items, allowedCompressionOptions)
	return rv
}

// A bar that holds this group’s items.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/groupTouchBar
func (g_ GroupTouchBarItem) GroupTouchBar() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("groupTouchBar"))
	return rv
}


// SetGroupTouchBar sets the value of the groupTouchBar property.
// A bar that holds this group’s items.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/groupTouchBar
func (g_ GroupTouchBarItem) SetGroupTouchBar(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGroupTouchBar:"), value)
}


