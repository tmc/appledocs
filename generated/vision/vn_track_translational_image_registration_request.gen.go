// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [TrackTranslationalImageRegistrationRequest] class.
var (
	TrackTranslationalImageRegistrationRequestClass     _TrackTranslationalImageRegistrationRequestClass
	TrackTranslationalImageRegistrationRequestClassOnce sync.Once
)

func getTrackTranslationalImageRegistrationRequestClass() _TrackTranslationalImageRegistrationRequestClass {
	TrackTranslationalImageRegistrationRequestClassOnce.Do(func() {
		TrackTranslationalImageRegistrationRequestClass = _TrackTranslationalImageRegistrationRequestClass{objc.GetClass("VNTrackTranslationalImageRegistrationRequest")}
	})
	return TrackTranslationalImageRegistrationRequestClass
}

type _TrackTranslationalImageRegistrationRequestClass struct {
	class objc.Class
}





// An interface definition for the [TrackTranslationalImageRegistrationRequest] class.
type ITrackTranslationalImageRegistrationRequest interface {
	IStatefulRequest
	

	// properties:
	Results() []ImageTranslationAlignmentObservation


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (tc _TrackTranslationalImageRegistrationRequestClass) Alloc() TrackTranslationalImageRegistrationRequest {
	rv := objc.Send[TrackTranslationalImageRegistrationRequest](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TrackTranslationalImageRegistrationRequestClass) New() TrackTranslationalImageRegistrationRequest {
	rv := objc.Send[TrackTranslationalImageRegistrationRequest](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TrackTranslationalImageRegistrationRequest) Init() TrackTranslationalImageRegistrationRequest {
	rv := objc.Send[TrackTranslationalImageRegistrationRequest](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TrackTranslationalImageRegistrationRequest) Autorelease() TrackTranslationalImageRegistrationRequest {
	rv := objc.Send[TrackTranslationalImageRegistrationRequest](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTrackTranslationalImageRegistrationRequest creates a new TrackTranslationalImageRegistrationRequest instance.
func NewTrackTranslationalImageRegistrationRequest() TrackTranslationalImageRegistrationRequest {
	return getTrackTranslationalImageRegistrationRequestClass().New()
}





// An image-analysis request, as a stateful request you track over time, that determines the affine transform necessary to align the content of two images.
//
// This request is similar to . However, as a , it automatically computes the registration against the previous frame.


// An image-analysis request, as a stateful request you track over time, that determines the affine transform necessary to align the content of two images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTrackTranslationalImageRegistrationRequest
type TrackTranslationalImageRegistrationRequest struct {
	StatefulRequest
}

// TrackTranslationalImageRegistrationRequestFrom constructs a [TrackTranslationalImageRegistrationRequest] from an unsafe.Pointer.
//
// An image-analysis request, as a stateful request you track over time, that determines the affine transform necessary to align the content of two images.
func TrackTranslationalImageRegistrationRequestFrom(ptr unsafe.Pointer) TrackTranslationalImageRegistrationRequest {
	return TrackTranslationalImageRegistrationRequest{
		StatefulRequest: StatefulRequestFrom(ptr),
	}
}






// Creates a new request that tracks the translational registration of two images, with a system callback on completion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTrackTranslationalImageRegistrationRequest/init(completionHandler:)
func NewTrackTranslationalImageRegistrationRequestWithCompletionHandler(completionHandler RequestCompletionHandler /* not a class type */) TrackTranslationalImageRegistrationRequest {
	instance := getTrackTranslationalImageRegistrationRequestClass().Alloc()
	rv := objc.Send[TrackTranslationalImageRegistrationRequest](instance.ID, objc.Sel("initWithCompletionHandler:"), completionHandler)
	rv.Autorelease()
	return rv
}






















// The observed translational image alignment request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTrackTranslationalImageRegistrationRequest/results
func (t_ TrackTranslationalImageRegistrationRequest) Results() []ImageTranslationAlignmentObservation {
	rv := objc.Send[[]ImageTranslationAlignmentObservation](t_.ID, objc.Sel("results"))
	return rv
}







