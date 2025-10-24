// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [TrackHomographicImageRegistrationRequest] class.
var (
	TrackHomographicImageRegistrationRequestClass     _TrackHomographicImageRegistrationRequestClass
	TrackHomographicImageRegistrationRequestClassOnce sync.Once
)

func getTrackHomographicImageRegistrationRequestClass() _TrackHomographicImageRegistrationRequestClass {
	TrackHomographicImageRegistrationRequestClassOnce.Do(func() {
		TrackHomographicImageRegistrationRequestClass = _TrackHomographicImageRegistrationRequestClass{objc.GetClass("VNTrackHomographicImageRegistrationRequest")}
	})
	return TrackHomographicImageRegistrationRequestClass
}

type _TrackHomographicImageRegistrationRequestClass struct {
	class objc.Class
}





// An interface definition for the [TrackHomographicImageRegistrationRequest] class.
type ITrackHomographicImageRegistrationRequest interface {
	IStatefulRequest
	

	// properties:
	Results() []ImageHomographicAlignmentObservation


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (tc _TrackHomographicImageRegistrationRequestClass) Alloc() TrackHomographicImageRegistrationRequest {
	rv := objc.Send[TrackHomographicImageRegistrationRequest](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TrackHomographicImageRegistrationRequestClass) New() TrackHomographicImageRegistrationRequest {
	rv := objc.Send[TrackHomographicImageRegistrationRequest](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TrackHomographicImageRegistrationRequest) Init() TrackHomographicImageRegistrationRequest {
	rv := objc.Send[TrackHomographicImageRegistrationRequest](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TrackHomographicImageRegistrationRequest) Autorelease() TrackHomographicImageRegistrationRequest {
	rv := objc.Send[TrackHomographicImageRegistrationRequest](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTrackHomographicImageRegistrationRequest creates a new TrackHomographicImageRegistrationRequest instance.
func NewTrackHomographicImageRegistrationRequest() TrackHomographicImageRegistrationRequest {
	return getTrackHomographicImageRegistrationRequestClass().New()
}





// An image-analysis request, as a stateful request you track over time, that determines the perspective warp matrix necessary to align the content of two images.
//
// This request is similar to . However, as a , it automatically computes the registration against the previous frame.


// An image-analysis request, as a stateful request you track over time, that determines the perspective warp matrix necessary to align the content of two images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTrackHomographicImageRegistrationRequest
type TrackHomographicImageRegistrationRequest struct {
	StatefulRequest
}

// TrackHomographicImageRegistrationRequestFrom constructs a [TrackHomographicImageRegistrationRequest] from an unsafe.Pointer.
//
// An image-analysis request, as a stateful request you track over time, that determines the perspective warp matrix necessary to align the content of two images.
func TrackHomographicImageRegistrationRequestFrom(ptr unsafe.Pointer) TrackHomographicImageRegistrationRequest {
	return TrackHomographicImageRegistrationRequest{
		StatefulRequest: StatefulRequestFrom(ptr),
	}
}






// Creates a new request that tracks the homographic transformation of two images, with a system callback on completion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTrackHomographicImageRegistrationRequest/init(completionHandler:)
func NewTrackHomographicImageRegistrationRequestWithCompletionHandler(completionHandler RequestCompletionHandler /* not a class type */) TrackHomographicImageRegistrationRequest {
	instance := getTrackHomographicImageRegistrationRequestClass().Alloc()
	rv := objc.Send[TrackHomographicImageRegistrationRequest](instance.ID, objc.Sel("initWithCompletionHandler:"), completionHandler)
	rv.Autorelease()
	return rv
}






















// The observed homographic image alignment request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTrackHomographicImageRegistrationRequest/results
func (t_ TrackHomographicImageRegistrationRequest) Results() []ImageHomographicAlignmentObservation {
	rv := objc.Send[[]ImageHomographicAlignmentObservation](t_.ID, objc.Sel("results"))
	return rv
}







