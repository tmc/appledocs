// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CompositionTrackSegment] class.
var (
	CompositionTrackSegmentClass     _CompositionTrackSegmentClass
	CompositionTrackSegmentClassOnce sync.Once
)

func getCompositionTrackSegmentClass() _CompositionTrackSegmentClass {
	CompositionTrackSegmentClassOnce.Do(func() {
		CompositionTrackSegmentClass = _CompositionTrackSegmentClass{objc.GetClass("AVCompositionTrackSegment")}
	})
	return CompositionTrackSegmentClass
}

type _CompositionTrackSegmentClass struct {
	class objc.Class
}

// An interface definition for the [CompositionTrackSegment] class.
type ICompositionTrackSegment interface {
	IAssetTrackSegment
}

// A track segment that maps a time from the source media track to the composition track.
//
// You typically use this class to save a low-level representation of a composition.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrackSegment
type CompositionTrackSegment struct {
	AssetTrackSegment
}

// CompositionTrackSegmentFrom constructs a [CompositionTrackSegment] from an unsafe.Pointer.
//
// A track segment that maps a time from the source media track to the composition track.
func CompositionTrackSegmentFrom(ptr unsafe.Pointer) CompositionTrackSegment {
	return CompositionTrackSegment{
		AssetTrackSegment: AssetTrackSegmentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CompositionTrackSegmentClass) Alloc() CompositionTrackSegment {
	rv := objc.Send[CompositionTrackSegment](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CompositionTrackSegmentClass) New() CompositionTrackSegment {
	rv := objc.Send[CompositionTrackSegment](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CompositionTrackSegment) Init() CompositionTrackSegment {
	rv := objc.Send[CompositionTrackSegment](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CompositionTrackSegment) Autorelease() CompositionTrackSegment {
	rv := objc.Send[CompositionTrackSegment](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCompositionTrackSegment creates a new CompositionTrackSegment instance.
func NewCompositionTrackSegment() CompositionTrackSegment {
	return getCompositionTrackSegmentClass().New()
}




