// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SharingServicePickerTouchBarItem] class.
var (
	sharingServicePickerTouchBarItemClass     _SharingServicePickerTouchBarItemClass
	sharingServicePickerTouchBarItemClassOnce sync.Once
)

func getSharingServicePickerTouchBarItemClass() _SharingServicePickerTouchBarItemClass {
	sharingServicePickerTouchBarItemClassOnce.Do(func() {
		sharingServicePickerTouchBarItemClass = _SharingServicePickerTouchBarItemClass{objc.GetClass("NSSharingServicePickerTouchBarItem")}
	})
	return sharingServicePickerTouchBarItemClass
}

type _SharingServicePickerTouchBarItemClass struct {
	class objc.Class
}

// An interface definition for the [SharingServicePickerTouchBarItem] class.
type ISharingServicePickerTouchBarItem interface {
	ITouchBarItem
}

// A bar item that, along with its delegate, provides a list of objects eligible for sharing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerTouchBarItem

type SharingServicePickerTouchBarItem struct {
	TouchBarItem
}

// SharingServicePickerTouchBarItemFrom constructs a [SharingServicePickerTouchBarItem] from an unsafe.Pointer.
//
// A bar item that, along with its delegate, provides a list of objects eligible for sharing.
func SharingServicePickerTouchBarItemFrom(ptr unsafe.Pointer) SharingServicePickerTouchBarItem {
	return SharingServicePickerTouchBarItem{
		TouchBarItem: TouchBarItemFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (sc _SharingServicePickerTouchBarItemClass) Alloc() SharingServicePickerTouchBarItem {
	rv := objc.Send[SharingServicePickerTouchBarItem](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SharingServicePickerTouchBarItemClass) New() SharingServicePickerTouchBarItem {
	rv := objc.Send[SharingServicePickerTouchBarItem](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SharingServicePickerTouchBarItem) Init() SharingServicePickerTouchBarItem {
	rv := objc.Send[SharingServicePickerTouchBarItem](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SharingServicePickerTouchBarItem) Autorelease() SharingServicePickerTouchBarItem {
	rv := objc.Send[SharingServicePickerTouchBarItem](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSharingServicePickerTouchBarItem creates a new SharingServicePickerTouchBarItem instance.
func NewSharingServicePickerTouchBarItem() SharingServicePickerTouchBarItem {
	return getSharingServicePickerTouchBarItemClass().New()
}




