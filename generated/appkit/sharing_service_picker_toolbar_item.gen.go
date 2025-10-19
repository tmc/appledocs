// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SharingServicePickerToolbarItem] class.
var (
	sharingServicePickerToolbarItemClass     _SharingServicePickerToolbarItemClass
	sharingServicePickerToolbarItemClassOnce sync.Once
)

func getSharingServicePickerToolbarItemClass() _SharingServicePickerToolbarItemClass {
	sharingServicePickerToolbarItemClassOnce.Do(func() {
		sharingServicePickerToolbarItemClass = _SharingServicePickerToolbarItemClass{objc.GetClass("NSSharingServicePickerToolbarItem")}
	})
	return sharingServicePickerToolbarItemClass
}

type _SharingServicePickerToolbarItemClass struct {
	class objc.Class
}

// An interface definition for the [SharingServicePickerToolbarItem] class.
type ISharingServicePickerToolbarItem interface {
	IToolbarItem
}

// A toolbar item that displays the macOS share sheet.
//
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




