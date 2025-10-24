// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VNTrackRectangleRequest */


/* debug [class_header]: Header for VNTrackRectangleRequest */
// The class instance for the [TrackRectangleRequest] class.
var (
	TrackRectangleRequestClass     _TrackRectangleRequestClass
	TrackRectangleRequestClassOnce sync.Once
)

func getTrackRectangleRequestClass() _TrackRectangleRequestClass {
	TrackRectangleRequestClassOnce.Do(func() {
		TrackRectangleRequestClass = _TrackRectangleRequestClass{objc.GetClass("VNTrackRectangleRequest")}
	})
	return TrackRectangleRequestClass
}

type _TrackRectangleRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TrackRectangleRequest */
// An interface definition for the [TrackRectangleRequest] class.
type ITrackRectangleRequest interface {
	ITrackingRequest
	
/* debug [class_interface_properties]: Properties for TrackRectangleRequest */
	// properties:
	VNTrackRectangleRequestRevision1() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TrackRectangleRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TrackRectangleRequest */
// Alloc allocates a new instance without initialization.
func (tc _TrackRectangleRequestClass) Alloc() TrackRectangleRequest {
	rv := objc.Send[TrackRectangleRequest](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TrackRectangleRequestClass) New() TrackRectangleRequest {
	rv := objc.Send[TrackRectangleRequest](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TrackRectangleRequest) Init() TrackRectangleRequest {
	rv := objc.Send[TrackRectangleRequest](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TrackRectangleRequest) Autorelease() TrackRectangleRequest {
	rv := objc.Send[TrackRectangleRequest](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTrackRectangleRequest creates a new TrackRectangleRequest instance.
func NewTrackRectangleRequest() TrackRectangleRequest {
	return getTrackRectangleRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TrackRectangleRequest */
// An image-analysis request that tracks movement of a previously identified rectangular object across multiple images or video frames.
//
// Use this type of request to track the bounding boxes of rectangles throughout a sequence of images. Vision returns locations for rectangles found in all orientations and sizes.


// An image-analysis request that tracks movement of a previously identified rectangular object across multiple images or video frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTrackRectangleRequest
type TrackRectangleRequest struct {
	TrackingRequest
}

// TrackRectangleRequestFrom constructs a [TrackRectangleRequest] from an unsafe.Pointer.
//
// An image-analysis request that tracks movement of a previously identified rectangular object across multiple images or video frames.
func TrackRectangleRequestFrom(ptr unsafe.Pointer) TrackRectangleRequest {
	return TrackRectangleRequest{
		TrackingRequest: TrackingRequestFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TrackRectangleRequest */

// Creates a new rectangle tracking request with a rectangle observation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTrackRectangleRequest/init(rectangleObservation:)
func NewTrackRectangleRequestWithRectangleObservation(observation IVNRectangleObservation) TrackRectangleRequest {
	instance := getTrackRectangleRequestClass().Alloc()
	rv := objc.Send[TrackRectangleRequest](instance.ID, objc.Sel("initWithRectangleObservation:"), observation)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTrackRectangleRequestWithRectangleObservation */


// Creates a new rectangle tracking request with a rectangle observation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTrackRectangleRequest/init(rectangleObservation:completionHandler:)
func NewTrackRectangleRequestWithRectangleObservationCompletionHandler(observation IVNRectangleObservation, completionHandler RequestCompletionHandler /* not a class type */) TrackRectangleRequest {
	instance := getTrackRectangleRequestClass().Alloc()
	rv := objc.Send[TrackRectangleRequest](instance.ID, objc.Sel("initWithRectangleObservation:completionHandler:"), observation, completionHandler)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTrackRectangleRequestWithRectangleObservationCompletionHandler */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TrackRectangleRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TrackRectangleRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TrackRectangleRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TrackRectangleRequest */

// A constant for specifying revision 1 of the rectangling tracking request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vntrackrectanglerequestrevision1
func (t_ TrackRectangleRequest) VNTrackRectangleRequestRevision1() int {
	rv := objc.Send[int](t_.ID, objc.Sel("VNTrackRectangleRequestRevision1"))
	return rv
}/* debug [instance_properties/getter]: VNTrackRectangleRequestRevision1 */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNTrackRectangleRequest */


