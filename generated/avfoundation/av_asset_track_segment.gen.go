// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAssetTrackSegment */


/* debug [class_header]: Header for AVAssetTrackSegment */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AssetTrackSegment */
// An interface definition for the [AssetTrackSegment] class.
type IAssetTrackSegment interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AssetTrackSegment */
	// properties:
	Empty() bool
	TimeMapping() TimeMapping /* not a class type */
	IsEmpty() bool
	SetIsEmpty(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AssetTrackSegment */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AssetTrackSegment */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AssetTrackSegment */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AssetTrackSegment *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AssetTrackSegment */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AssetTrackSegment */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AssetTrackSegment */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AssetTrackSegment */

// A Boolean value that indicates whether the segment is empty.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrackSegment/isEmpty
func (a_ AssetTrackSegment) Empty() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("empty"))
	return rv
}/* debug [instance_properties/getter]: empty */


// The time range of the track that this segment presents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrackSegment/timeMapping
func (a_ AssetTrackSegment) TimeMapping() TimeMapping /* not a class type */ {
	rv := objc.Send[TimeMapping](a_.ID, objc.Sel("timeMapping"))
	return rv
}/* debug [instance_properties/getter]: timeMapping */


// A Boolean value that indicates whether the segment is empty.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettracksegment/isempty
func (a_ AssetTrackSegment) IsEmpty() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isEmpty"))
	return rv
}/* debug [instance_properties/getter]: isEmpty */


// A Boolean value that indicates whether the segment is empty.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettracksegment/isempty
func (a_ AssetTrackSegment) SetIsEmpty(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsEmpty:"), value)
}/* debug [instance_properties/setter]: isEmpty */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAssetTrackSegment */



