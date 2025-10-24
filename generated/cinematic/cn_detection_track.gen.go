// Code generated from Apple documentation for Cinematic. DO NOT EDIT.

package cinematic

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corevideo"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CNDetectionTrack */


/* debug [class_header]: Header for CNDetectionTrack */
// The class instance for the [CNDetectionTrack] class.
var (
	CNDetectionTrackClass     _CNDetectionTrackClass
	CNDetectionTrackClassOnce sync.Once
)

func getCNDetectionTrackClass() _CNDetectionTrackClass {
	CNDetectionTrackClassOnce.Do(func() {
		CNDetectionTrackClass = _CNDetectionTrackClass{objc.GetClass("CNDetectionTrack")}
	})
	return CNDetectionTrackClass
}

type _CNDetectionTrackClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNDetectionTrack */
// An interface definition for the [CNDetectionTrack] class.
type ICNDetectionTrack interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CNDetectionTrack */
	// properties:
	DetectionGroupID() CNDetectionGroupID /* typedef */
	DetectionID() CNDetectionID /* typedef */
	DetectionType() CNDetectionType
	Discrete() bool
	UserCreated() bool
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNDetectionTrack */
	// methods:
	DetectionAtOrBeforeTime(time objc.IObject /* cross-framework: Time */) ICNDetection
	DetectionNearestTime(time objc.IObject /* cross-framework: Time */) ICNDetection
	DetectionsInTimeRange(timeRange TimeRange /* not a class type */) []CNDetection
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNDetectionTrack */
// Alloc allocates a new instance without initialization.
func (cc _CNDetectionTrackClass) Alloc() CNDetectionTrack {
	rv := objc.Send[CNDetectionTrack](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNDetectionTrackClass) New() CNDetectionTrack {
	rv := objc.Send[CNDetectionTrack](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNDetectionTrack) Init() CNDetectionTrack {
	rv := objc.Send[CNDetectionTrack](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNDetectionTrack) Autorelease() CNDetectionTrack {
	rv := objc.Send[CNDetectionTrack](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNDetectionTrack creates a new CNDetectionTrack instance.
func NewCNDetectionTrack() CNDetectionTrack {
	return getCNDetectionTrackClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNDetectionTrack */
// An object representing a series of detections of the same subject over time.


// An object representing a series of detections of the same subject over time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetectionTrack-61x7g
type CNDetectionTrack struct {
	objectivec.Object
}

// CNDetectionTrackFrom constructs a [CNDetectionTrack] from an unsafe.Pointer.
//
// An object representing a series of detections of the same subject over time.
func CNDetectionTrackFrom(ptr unsafe.Pointer) CNDetectionTrack {
	return CNDetectionTrack{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNDetectionTrack *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNDetectionTrack */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNDetectionTrack */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNDetectionTrack */

// Returns the array of detections in the detection track before a given time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetectionTrack-61x7g/detectionAtOrBeforeTime:
func (c_ CNDetectionTrack) DetectionAtOrBeforeTime(time objc.IObject /* cross-framework: Time */) ICNDetection {
	rv := objc.Send[CNDetection](c_.ID, objc.Sel("detectionAtOrBeforeTime:"), time)
	return rv
}/* debug [instance_methods/method]: DetectionAtOrBeforeTime */


// Returns the array of detections in the detection track nearest a given time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetectionTrack-61x7g/detectionNearestTime:
func (c_ CNDetectionTrack) DetectionNearestTime(time objc.IObject /* cross-framework: Time */) ICNDetection {
	rv := objc.Send[CNDetection](c_.ID, objc.Sel("detectionNearestTime:"), time)
	return rv
}/* debug [instance_methods/method]: DetectionNearestTime */


// Returns the array of detections in the detection track within the given time range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetectionTrack-61x7g/detectionsInTimeRange:
func (c_ CNDetectionTrack) DetectionsInTimeRange(timeRange TimeRange /* not a class type */) []CNDetection {
	rv := objc.Send[[]CNDetection](c_.ID, objc.Sel("detectionsInTimeRange:"), timeRange)
	return rv
}/* debug [instance_methods/method]: DetectionsInTimeRange */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNDetectionTrack */

// The detection group ID of the subject detected by the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetectionTrack-61x7g/detectionGroupID
func (c_ CNDetectionTrack) DetectionGroupID() CNDetectionGroupID /* typedef */ {
	rv := objc.Send[int64](c_.ID, objc.Sel("detectionGroupID"))
	return rv
}/* debug [instance_properties/getter]: detectionGroupID */


// The unique ID of the subject detected during this track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetectionTrack-61x7g/detectionID
func (c_ CNDetectionTrack) DetectionID() CNDetectionID /* typedef */ {
	rv := objc.Send[int64](c_.ID, objc.Sel("detectionID"))
	return rv
}/* debug [instance_properties/getter]: detectionID */


// The type of object that’s detected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetectionTrack-61x7g/detectionType
func (c_ CNDetectionTrack) DetectionType() CNDetectionType {
	rv := objc.Send[CNDetectionType](c_.ID, objc.Sel("detectionType"))
	return rv
}/* debug [instance_properties/getter]: detectionType */


// A flag determining if the detection track has discrete detections, otherwise continuous.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetectionTrack-61x7g/discrete
func (c_ CNDetectionTrack) Discrete() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("discrete"))
	return rv
}/* debug [instance_properties/getter]: discrete */


// A flag indicating if the client created the detection track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetectionTrack-61x7g/userCreated
func (c_ CNDetectionTrack) UserCreated() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("userCreated"))
	return rv
}/* debug [instance_properties/getter]: userCreated */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNDetectionTrack */



