// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TrackingSeparatorToolbarItem] class.
var (
	TrackingSeparatorToolbarItemClass     _TrackingSeparatorToolbarItemClass
	TrackingSeparatorToolbarItemClassOnce sync.Once
)

func getTrackingSeparatorToolbarItemClass() _TrackingSeparatorToolbarItemClass {
	TrackingSeparatorToolbarItemClassOnce.Do(func() {
		TrackingSeparatorToolbarItemClass = _TrackingSeparatorToolbarItemClass{objc.GetClass("NSTrackingSeparatorToolbarItem")}
	})
	return TrackingSeparatorToolbarItemClass
}

type _TrackingSeparatorToolbarItemClass struct {
	class objc.Class
}

// An interface definition for the [TrackingSeparatorToolbarItem] class.
type ITrackingSeparatorToolbarItem interface {
	IToolbarItem
}

// A toolbar separator that aligns with the vertical split view in the same window.
//
// Use a to divide an into sections that visually align with the views on either side of the divider of the . This keeps s above the content that’s the for the item’s . The must be in the same window as the toolbar containing this item before showing the toolbar.
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getTrackingSeparatorToolbarItemClass().New()
}


// The vertical split view to align with the toolbar separator.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTrackingSeparatorToolbarItem/splitView
func (t_ TrackingSeparatorToolbarItem) SplitView() NSSplitView {
	rv := objc.Send[NSSplitView](t_.ID, objc.Sel("splitView"))
	return rv
}


// SetSplitView sets the value of the splitView property.
// The vertical split view to align with the toolbar separator.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTrackingSeparatorToolbarItem/splitView
func (t_ TrackingSeparatorToolbarItem) SetSplitView(value ISplitView) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSplitView:"), value)
}

// The object that defines the action method the toolbar item calls when clicked.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/target
func (t_ TrackingSeparatorToolbarItem) Target() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("target"))
	return rv
}


// SetTarget sets the value of the target property.
// The object that defines the action method the toolbar item calls when clicked.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/target
func (t_ TrackingSeparatorToolbarItem) SetTarget(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTarget:"), value)
}

// The index of the split view divider to align with the tracking separator.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstrackingseparatortoolbaritem/dividerindex
func (t_ TrackingSeparatorToolbarItem) DividerIndex() int {
	rv := objc.Send[int](t_.ID, objc.Sel("dividerIndex"))
	return rv
}


// SetDividerIndex sets the value of the dividerIndex property.
// The index of the split view divider to align with the tracking separator.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstrackingseparatortoolbaritem/dividerindex
func (t_ TrackingSeparatorToolbarItem) SetDividerIndex(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDividerIndex:"), value)
}



