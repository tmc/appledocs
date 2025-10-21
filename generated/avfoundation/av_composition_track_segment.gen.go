// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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


// A Boolean value that indicates whether the segment is empty.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontracksegment/isempty
func (c_ CompositionTrackSegment) IsEmpty() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isEmpty"))
	return rv
}


// SetIsEmpty sets the value of the isEmpty property.
// A Boolean value that indicates whether the segment is empty.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontracksegment/isempty
func (c_ CompositionTrackSegment) SetIsEmpty(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsEmpty:"), value)
}

// An identifier of a track in the container file whose media this track segment presents.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontracksegment/sourcetrackid
func (c_ CompositionTrackSegment) SourceTrackID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("sourceTrackID"))
	return rv
}


// SetSourceTrackID sets the value of the sourceTrackID property.
// An identifier of a track in the container file whose media this track segment presents.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontracksegment/sourcetrackid
func (c_ CompositionTrackSegment) SetSourceTrackID(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSourceTrackID:"), value)
}

// A URL of the container file whose media this track segment presents.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontracksegment/sourceurl
func (c_ CompositionTrackSegment) SourceURL() foundation.URL {
	rv := objc.Send[foundation.URL](c_.ID, objc.Sel("sourceURL"))
	return rv
}


// SetSourceURL sets the value of the sourceURL property.
// A URL of the container file whose media this track segment presents.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontracksegment/sourceurl
func (c_ CompositionTrackSegment) SetSourceURL(value foundation.URL) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSourceURL:"), value)
}



