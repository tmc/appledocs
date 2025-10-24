// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [TrackObjectRequest] class.
var (
	TrackObjectRequestClass     _TrackObjectRequestClass
	TrackObjectRequestClassOnce sync.Once
)

func getTrackObjectRequestClass() _TrackObjectRequestClass {
	TrackObjectRequestClassOnce.Do(func() {
		TrackObjectRequestClass = _TrackObjectRequestClass{objc.GetClass("VNTrackObjectRequest")}
	})
	return TrackObjectRequestClass
}

type _TrackObjectRequestClass struct {
	class objc.Class
}





// An interface definition for the [TrackObjectRequest] class.
type ITrackObjectRequest interface {
	ITrackingRequest
	

	// properties:
	VNTrackObjectRequestRevision1() int
	VNTrackObjectRequestRevision2() int


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (tc _TrackObjectRequestClass) Alloc() TrackObjectRequest {
	rv := objc.Send[TrackObjectRequest](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TrackObjectRequestClass) New() TrackObjectRequest {
	rv := objc.Send[TrackObjectRequest](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TrackObjectRequest) Init() TrackObjectRequest {
	rv := objc.Send[TrackObjectRequest](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TrackObjectRequest) Autorelease() TrackObjectRequest {
	rv := objc.Send[TrackObjectRequest](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTrackObjectRequest creates a new TrackObjectRequest instance.
func NewTrackObjectRequest() TrackObjectRequest {
	return getTrackObjectRequestClass().New()
}





// An image-analysis request that tracks the movement of a previously identified object across multiple images or video frames.
//
// Use this type of request to track the bounding boxes around objects previously identified in an image. Vision attempts to locate the same object from the input observation throughout the sequence.


// An image-analysis request that tracks the movement of a previously identified object across multiple images or video frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTrackObjectRequest
type TrackObjectRequest struct {
	TrackingRequest
}

// TrackObjectRequestFrom constructs a [TrackObjectRequest] from an unsafe.Pointer.
//
// An image-analysis request that tracks the movement of a previously identified object across multiple images or video frames.
func TrackObjectRequestFrom(ptr unsafe.Pointer) TrackObjectRequest {
	return TrackObjectRequest{
		TrackingRequest: TrackingRequestFrom(ptr),
	}
}






// Creates a new object tracking request with a detected object observation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTrackObjectRequest/init(detectedObjectObservation:)
func NewTrackObjectRequestWithDetectedObjectObservation(observation IVNDetectedObjectObservation) TrackObjectRequest {
	instance := getTrackObjectRequestClass().Alloc()
	rv := objc.Send[TrackObjectRequest](instance.ID, objc.Sel("initWithDetectedObjectObservation:"), observation)
	rv.Autorelease()
	return rv
}


// Creates a new object tracking request with a detected object observation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTrackObjectRequest/init(detectedObjectObservation:completionHandler:)
func NewTrackObjectRequestWithDetectedObjectObservationCompletionHandler(observation IVNDetectedObjectObservation, completionHandler RequestCompletionHandler /* not a class type */) TrackObjectRequest {
	instance := getTrackObjectRequestClass().Alloc()
	rv := objc.Send[TrackObjectRequest](instance.ID, objc.Sel("initWithDetectedObjectObservation:completionHandler:"), observation, completionHandler)
	rv.Autorelease()
	return rv
}






















// A constant for specifying revision 1 of the object tracking request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vntrackobjectrequestrevision1
func (t_ TrackObjectRequest) VNTrackObjectRequestRevision1() int {
	rv := objc.Send[int](t_.ID, objc.Sel("VNTrackObjectRequestRevision1"))
	return rv
}


// A constant for specifying revision 2 of the object tracking request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vntrackobjectrequestrevision2
func (t_ TrackObjectRequest) VNTrackObjectRequestRevision2() int {
	rv := objc.Send[int](t_.ID, objc.Sel("VNTrackObjectRequestRevision2"))
	return rv
}







