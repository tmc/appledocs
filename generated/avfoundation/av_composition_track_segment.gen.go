// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCompositionTrackSegment */


/* debug [class_header]: Header for AVCompositionTrackSegment */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CompositionTrackSegment */
// An interface definition for the [CompositionTrackSegment] class.
type ICompositionTrackSegment interface {
	IAssetTrackSegment
	
/* debug [class_interface_properties]: Properties for CompositionTrackSegment */
	// properties:
	Empty() bool
	SourceTrackID() PersistentTrackID /* not a class type */
	SourceURL() objc.IObject /* cross-framework: NSURL */
	IsEmpty() bool
	SetIsEmpty(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CompositionTrackSegment */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CompositionTrackSegment */
// Alloc allocates a new instance without initialization.
func (cc _CompositionTrackSegmentClass) Alloc() CompositionTrackSegment {
	rv := objc.Send[CompositionTrackSegment](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CompositionTrackSegment */
// A track segment that maps a time from the source media track to the composition track.
//
// You typically use this class to save a low-level representation of a composition.


// A track segment that maps a time from the source media track to the composition track.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CompositionTrackSegment */

// Creates an object that presents an empty composition track segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrackSegment/init(timeRange:)
func NewCompositionTrackSegmentWithTimeRange(timeRange TimeRange /* not a class type */) CompositionTrackSegment {
	instance := getCompositionTrackSegmentClass().Alloc()
	rv := objc.Send[CompositionTrackSegment](instance.ID, objc.Sel("initWithTimeRange:"), timeRange)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCompositionTrackSegmentWithTimeRange */


// Creates an object that presents a segment of a media file that the specified URL references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrackSegment/init(url:trackID:sourceTimeRange:targetTimeRange:)
func NewCompositionTrackSegmentWithURLTrackIDSourceTimeRangeTargetTimeRange(URL objc.IObject /* cross-framework: NSURL */, trackID PersistentTrackID /* not a class type */, sourceTimeRange TimeRange /* not a class type */, targetTimeRange TimeRange /* not a class type */) CompositionTrackSegment {
	instance := getCompositionTrackSegmentClass().Alloc()
	rv := objc.Send[CompositionTrackSegment](instance.ID, objc.Sel("initWithURL:trackID:sourceTimeRange:targetTimeRange:"), URL, trackID, sourceTimeRange, targetTimeRange)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCompositionTrackSegmentWithURLTrackIDSourceTimeRangeTargetTimeRange */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CompositionTrackSegment */

// Returns a new object that presents an empty composition track segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrackSegment/compositionTrackSegmentWithTimeRange:
func (cc _CompositionTrackSegmentClass) CompositionTrackSegmentWithTimeRange(timeRange TimeRange /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("compositionTrackSegmentWithTimeRange:"), timeRange)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CompositionTrackSegmentWithTimeRange) */


// Returns a new an object that presents a segment of a media file that the specified URL references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrackSegment/compositionTrackSegmentWithURL:trackID:sourceTimeRange:targetTimeRange:
func (cc _CompositionTrackSegmentClass) CompositionTrackSegmentWithURLTrackIDSourceTimeRangeTargetTimeRange(URL objc.IObject /* cross-framework: NSURL */, trackID PersistentTrackID /* not a class type */, sourceTimeRange TimeRange /* not a class type */, targetTimeRange TimeRange /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("compositionTrackSegmentWithURL:trackID:sourceTimeRange:targetTimeRange:"), URL, trackID, sourceTimeRange, targetTimeRange)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CompositionTrackSegmentWithURLTrackIDSourceTimeRangeTargetTimeRange) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CompositionTrackSegment */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CompositionTrackSegment */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CompositionTrackSegment */

// A Boolean value that indicates whether the segment is empty.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrackSegment/isEmpty
func (c_ CompositionTrackSegment) Empty() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("empty"))
	return rv
}/* debug [instance_properties/getter]: empty */


// An identifier of a track in the container file whose media this track segment presents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrackSegment/sourceTrackID
func (c_ CompositionTrackSegment) SourceTrackID() PersistentTrackID /* not a class type */ {
	rv := objc.Send[PersistentTrackID](c_.ID, objc.Sel("sourceTrackID"))
	return rv
}/* debug [instance_properties/getter]: sourceTrackID */


// A URL of the container file whose media this track segment presents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrackSegment/sourceURL
func (c_ CompositionTrackSegment) SourceURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](c_.ID, objc.Sel("sourceURL"))
	return rv
}/* debug [instance_properties/getter]: sourceURL */


// A Boolean value that indicates whether the segment is empty.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontracksegment/isempty
func (c_ CompositionTrackSegment) IsEmpty() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isEmpty"))
	return rv
}/* debug [instance_properties/getter]: isEmpty */


// A Boolean value that indicates whether the segment is empty.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontracksegment/isempty
func (c_ CompositionTrackSegment) SetIsEmpty(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsEmpty:"), value)
}/* debug [instance_properties/setter]: isEmpty */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCompositionTrackSegment */


