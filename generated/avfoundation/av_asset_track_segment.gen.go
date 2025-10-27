// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [AssetTrackSegment] class.
var (
	AssetTrackSegmentClass     _AssetTrackSegmentClass
	AssetTrackSegmentClassOnce sync.Once
)

func getAssetTrackSegmentClass() _AssetTrackSegmentClass {
	AssetTrackSegmentClassOnce.Do(func() {
		AssetTrackSegmentClass = _AssetTrackSegmentClass{objc.GetClass("AVAssetTrackSegment")}
	})
	return AssetTrackSegmentClass
}

type _AssetTrackSegmentClass struct {
	class objc.Class
}





// An interface definition for the [AssetTrackSegment] class.
type IAssetTrackSegment interface {
	objectivec.IObject
	

	// properties:
	Empty() bool
	TimeMapping() objectivec.IObject
	IsEmpty() bool
	SetIsEmpty(value bool)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ac _AssetTrackSegmentClass) Alloc() AssetTrackSegment {
	rv := objc.Send[AssetTrackSegment](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AssetTrackSegmentClass) New() AssetTrackSegment {
	rv := objc.Send[AssetTrackSegment](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetTrackSegment) Init() AssetTrackSegment {
	rv := objc.Send[AssetTrackSegment](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetTrackSegment) Autorelease() AssetTrackSegment {
	rv := objc.Send[AssetTrackSegment](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetTrackSegment creates a new AssetTrackSegment instance.
func NewAssetTrackSegment() AssetTrackSegment {
	return getAssetTrackSegmentClass().New()
}





// An object that represents a time range segment of an asset track.


// An object that represents a time range segment of an asset track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrackSegment
type AssetTrackSegment struct {
	objectivec.Object
}

// AssetTrackSegmentFrom constructs a [AssetTrackSegment] from an unsafe.Pointer.
//
// An object that represents a time range segment of an asset track.
func AssetTrackSegmentFrom(ptr unsafe.Pointer) AssetTrackSegment {
	return AssetTrackSegment{objectivec.Object{objc.ID(ptr)}}
}

























// A Boolean value that indicates whether the segment is empty.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrackSegment/isEmpty
func (a_ AssetTrackSegment) Empty() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("empty"))
	return rv
}


// The time range of the track that this segment presents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrackSegment/timeMapping
func (a_ AssetTrackSegment) TimeMapping() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("timeMapping"))
	return rv
}


// A Boolean value that indicates whether the segment is empty.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettracksegment/isempty
func (a_ AssetTrackSegment) IsEmpty() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isEmpty"))
	return rv
}


// A Boolean value that indicates whether the segment is empty.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettracksegment/isempty
func (a_ AssetTrackSegment) SetIsEmpty(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsEmpty:"), value)
}








