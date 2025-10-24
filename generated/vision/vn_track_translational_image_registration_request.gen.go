// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VNTrackTranslationalImageRegistrationRequest */


/* debug [class_header]: Header for VNTrackTranslationalImageRegistrationRequest */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TrackTranslationalImageRegistrationRequest */
// An interface definition for the [TrackTranslationalImageRegistrationRequest] class.
type ITrackTranslationalImageRegistrationRequest interface {
	IStatefulRequest
	
/* debug [class_interface_properties]: Properties for TrackTranslationalImageRegistrationRequest */
	// properties:
	Results() []ImageTranslationAlignmentObservation
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TrackTranslationalImageRegistrationRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TrackTranslationalImageRegistrationRequest */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TrackTranslationalImageRegistrationRequest */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TrackTranslationalImageRegistrationRequest */

// Creates a new request that tracks the translational registration of two images, with a system callback on completion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTrackTranslationalImageRegistrationRequest/init(completionHandler:)
func NewTrackTranslationalImageRegistrationRequestWithCompletionHandler(completionHandler RequestCompletionHandler /* not a class type */) TrackTranslationalImageRegistrationRequest {
	instance := getTrackTranslationalImageRegistrationRequestClass().Alloc()
	rv := objc.Send[TrackTranslationalImageRegistrationRequest](instance.ID, objc.Sel("initWithCompletionHandler:"), completionHandler)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTrackTranslationalImageRegistrationRequestWithCompletionHandler */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TrackTranslationalImageRegistrationRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TrackTranslationalImageRegistrationRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TrackTranslationalImageRegistrationRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TrackTranslationalImageRegistrationRequest */

// The observed translational image alignment request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTrackTranslationalImageRegistrationRequest/results
func (t_ TrackTranslationalImageRegistrationRequest) Results() []ImageTranslationAlignmentObservation {
	rv := objc.Send[[]ImageTranslationAlignmentObservation](t_.ID, objc.Sel("results"))
	return rv
}/* debug [instance_properties/getter]: results */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNTrackTranslationalImageRegistrationRequest */


