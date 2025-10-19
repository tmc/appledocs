// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TrackingSeparatorToolbarItem] class.
var trackingSeparatorToolbarItemClass = _TrackingSeparatorToolbarItemClass{objc.GetClass("NSTrackingSeparatorToolbarItem")}

type _TrackingSeparatorToolbarItemClass struct {
	class objc.Class
}

// An interface definition for the [TrackingSeparatorToolbarItem] class.
type ITrackingSeparatorToolbarItem interface {
	IToolbarItem
}

// A toolbar separator that aligns with the vertical split view in the same window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTrackingSeparatorToolbarItem

type TrackingSeparatorToolbarItem struct {
	ToolbarItem
}

// TrackingSeparatorToolbarItemFrom constructs a [TrackingSeparatorToolbarItem] from an unsafe.Pointer.
//
// A toolbar separator that aligns with the vertical split view in the same window.
func TrackingSeparatorToolbarItemFrom(ptr unsafe.Pointer) TrackingSeparatorToolbarItem {
	return TrackingSeparatorToolbarItem{
		ToolbarItem: ToolbarItemFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (tc _TrackingSeparatorToolbarItemClass) Alloc() TrackingSeparatorToolbarItem {
	rv := objc.Send[TrackingSeparatorToolbarItem](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (tc _TrackingSeparatorToolbarItemClass) New() TrackingSeparatorToolbarItem {
	rv := objc.Send[TrackingSeparatorToolbarItem](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TrackingSeparatorToolbarItem) Init() TrackingSeparatorToolbarItem {
	rv := objc.Send[TrackingSeparatorToolbarItem](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TrackingSeparatorToolbarItem) Autorelease() TrackingSeparatorToolbarItem {
	rv := objc.Send[TrackingSeparatorToolbarItem](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTrackingSeparatorToolbarItem creates a new TrackingSeparatorToolbarItem instance.
func NewTrackingSeparatorToolbarItem() TrackingSeparatorToolbarItem {
	return trackingSeparatorToolbarItemClass.New()
}




