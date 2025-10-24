// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SharingServicePickerToolbarItem] class.
var (
	SharingServicePickerToolbarItemClass     _SharingServicePickerToolbarItemClass
	SharingServicePickerToolbarItemClassOnce sync.Once
)

func getSharingServicePickerToolbarItemClass() _SharingServicePickerToolbarItemClass {
	SharingServicePickerToolbarItemClassOnce.Do(func() {
		SharingServicePickerToolbarItemClass = _SharingServicePickerToolbarItemClass{objc.GetClass("NSSharingServicePickerToolbarItem")}
	})
	return SharingServicePickerToolbarItemClass
}

type _SharingServicePickerToolbarItemClass struct {
	class objc.Class
}

// An interface definition for the [SharingServicePickerToolbarItem] class.
type ISharingServicePickerToolbarItem interface {
	IToolbarItem
	// properties:
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	// methods:
}

// A toolbar item that displays the macOS share sheet.
//
// An object is a standard item you add to your window’s toolbar. When someone clicks it, the item displays the macOS share sheet. Use this item to share the selected or focal content from the current window. For example, you might share the photo someone is viewing, the currently selected text, or the window’s associated document. Provide the items to share using the associated object. For an app built using Mac Catalyst, provide the items from the object in the property.


// A toolbar item that displays the macOS share sheet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerToolbarItem
type SharingServicePickerToolbarItem struct {
	ToolbarItem
}

// SharingServicePickerToolbarItemFrom constructs a [SharingServicePickerToolbarItem] from an unsafe.Pointer.
//
// A toolbar item that displays the macOS share sheet.
func SharingServicePickerToolbarItemFrom(ptr unsafe.Pointer) SharingServicePickerToolbarItem {
	return SharingServicePickerToolbarItem{
		ToolbarItem: ToolbarItemFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SharingServicePickerToolbarItemClass) Alloc() SharingServicePickerToolbarItem {
	rv := objc.Send[SharingServicePickerToolbarItem](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SharingServicePickerToolbarItemClass) New() SharingServicePickerToolbarItem {
	rv := objc.Send[SharingServicePickerToolbarItem](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SharingServicePickerToolbarItem) Init() SharingServicePickerToolbarItem {
	rv := objc.Send[SharingServicePickerToolbarItem](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SharingServicePickerToolbarItem) Autorelease() SharingServicePickerToolbarItem {
	rv := objc.Send[SharingServicePickerToolbarItem](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSharingServicePickerToolbarItem creates a new SharingServicePickerToolbarItem instance.
func NewSharingServicePickerToolbarItem() SharingServicePickerToolbarItem {
	return getSharingServicePickerToolbarItemClass().New()
}



// The custom object from your app that provides the items to share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerToolbarItem/delegate
func (s_ SharingServicePickerToolbarItem) Delegate() objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("delegate"))
	return rv
}


// The custom object from your app that provides the items to share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerToolbarItem/delegate
func (s_ SharingServicePickerToolbarItem) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
}


