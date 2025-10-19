// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TrackingArea] class.
var (
	trackingAreaClass     _TrackingAreaClass
	trackingAreaClassOnce sync.Once
)

func getTrackingAreaClass() _TrackingAreaClass {
	trackingAreaClassOnce.Do(func() {
		trackingAreaClass = _TrackingAreaClass{objc.GetClass("NSTrackingArea")}
	})
	return trackingAreaClass
}

type _TrackingAreaClass struct {
	class objc.Class
}

// An interface definition for the [TrackingArea] class.
type ITrackingArea interface {
	objectivec.IObject
}

// A region of a view that generates mouse-tracking and cursor-update events when the pointer is over that region.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTrackingArea
type TrackingArea struct {
	objectivec.Object
}

// TrackingAreaFrom constructs a [TrackingArea] from an unsafe.Pointer.
//
// A region of a view that generates mouse-tracking and cursor-update events when the pointer is over that region.
func TrackingAreaFrom(ptr unsafe.Pointer) TrackingArea {
	return TrackingArea{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TrackingAreaClass) Alloc() TrackingArea {
	rv := objc.Send[TrackingArea](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TrackingAreaClass) New() TrackingArea {
	rv := objc.Send[TrackingArea](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TrackingArea) Init() TrackingArea {
	rv := objc.Send[TrackingArea](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TrackingArea) Autorelease() TrackingArea {
	rv := objc.Send[TrackingArea](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTrackingArea creates a new TrackingArea instance.
func NewTrackingArea() TrackingArea {
	return getTrackingAreaClass().New()
}




