// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VNTrackingRequest */


/* debug [class_header]: Header for VNTrackingRequest */
// The class instance for the [TrackingRequest] class.
var (
	TrackingRequestClass     _TrackingRequestClass
	TrackingRequestClassOnce sync.Once
)

func getTrackingRequestClass() _TrackingRequestClass {
	TrackingRequestClassOnce.Do(func() {
		TrackingRequestClass = _TrackingRequestClass{objc.GetClass("VNTrackingRequest")}
	})
	return TrackingRequestClass
}

type _TrackingRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TrackingRequest */
// An interface definition for the [TrackingRequest] class.
type ITrackingRequest interface {
	IImageBasedRequest
	
/* debug [class_interface_properties]: Properties for TrackingRequest */
	// properties:
	InputObservation() IVNDetectedObjectObservation
	SetInputObservation(value IVNDetectedObjectObservation)
	LastFrame() bool
	SetLastFrame(value bool)
	TrackingLevel() RequestTrackingLevel
	SetTrackingLevel(value RequestTrackingLevel)
	IsLastFrame() bool
	SetIsLastFrame(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TrackingRequest */
	// methods:
	SupportedNumberOfTrackersAndReturnError(error_ objectivec.IObject) uint
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TrackingRequest */
// Alloc allocates a new instance without initialization.
func (tc _TrackingRequestClass) Alloc() TrackingRequest {
	rv := objc.Send[TrackingRequest](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TrackingRequestClass) New() TrackingRequest {
	rv := objc.Send[TrackingRequest](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TrackingRequest) Init() TrackingRequest {
	rv := objc.Send[TrackingRequest](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TrackingRequest) Autorelease() TrackingRequest {
	rv := objc.Send[TrackingRequest](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTrackingRequest creates a new TrackingRequest instance.
func NewTrackingRequest() TrackingRequest {
	return getTrackingRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TrackingRequest */
// The abstract superclass for image-analysis requests that track unique features across multiple images or video frames.
//
// Instantiate a tracking request subclass to perform object tracking across multiple frames of an image. After initialization, configure the degree of accuracy by setting , and provide observations you’d like to track by setting the initial bounding box.


// The abstract superclass for image-analysis requests that track unique features across multiple images or video frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTrackingRequest
type TrackingRequest struct {
	ImageBasedRequest
}

// TrackingRequestFrom constructs a [TrackingRequest] from an unsafe.Pointer.
//
// The abstract superclass for image-analysis requests that track unique features across multiple images or video frames.
func TrackingRequestFrom(ptr unsafe.Pointer) TrackingRequest {
	return TrackingRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TrackingRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TrackingRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TrackingRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TrackingRequest */

// Returns the maximum number of simultaneous trackers for the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTrackingRequest/supportedNumber(ofTrackersAndReturnError:)
func (t_ TrackingRequest) SupportedNumberOfTrackersAndReturnError(error_ objectivec.IObject) uint {
	rv := objc.Send[uint](t_.ID, objc.Sel("supportedNumberOfTrackersAndReturnError:"), error_)
	return rv
}/* debug [instance_methods/method]: SupportedNumberOfTrackersAndReturnError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TrackingRequest */

// The observation object defining a region to track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTrackingRequest/inputObservation
func (t_ TrackingRequest) InputObservation() IVNDetectedObjectObservation {
	rv := objc.Send[DetectedObjectObservation](t_.ID, objc.Sel("inputObservation"))
	return rv
}/* debug [instance_properties/getter]: inputObservation */


// The observation object defining a region to track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTrackingRequest/inputObservation
func (t_ TrackingRequest) SetInputObservation(value IVNDetectedObjectObservation) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setInputObservation:"), value)
}/* debug [instance_properties/setter]: inputObservation */


// A Boolean that indicates the last frame in a tracking sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTrackingRequest/isLastFrame
func (t_ TrackingRequest) LastFrame() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("lastFrame"))
	return rv
}/* debug [instance_properties/getter]: lastFrame */


// A Boolean that indicates the last frame in a tracking sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTrackingRequest/isLastFrame
func (t_ TrackingRequest) SetLastFrame(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLastFrame:"), value)
}/* debug [instance_properties/setter]: lastFrame */


// A value for specifying whether to prioritize speed or location accuracy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTrackingRequest/trackingLevel
func (t_ TrackingRequest) TrackingLevel() RequestTrackingLevel {
	rv := objc.Send[RequestTrackingLevel](t_.ID, objc.Sel("trackingLevel"))
	return rv
}/* debug [instance_properties/getter]: trackingLevel */


// A value for specifying whether to prioritize speed or location accuracy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTrackingRequest/trackingLevel
func (t_ TrackingRequest) SetTrackingLevel(value RequestTrackingLevel) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTrackingLevel:"), value)
}/* debug [instance_properties/setter]: trackingLevel */


// A Boolean that indicates the last frame in a tracking sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vntrackingrequest/islastframe
func (t_ TrackingRequest) IsLastFrame() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isLastFrame"))
	return rv
}/* debug [instance_properties/getter]: isLastFrame */


// A Boolean that indicates the last frame in a tracking sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vntrackingrequest/islastframe
func (t_ TrackingRequest) SetIsLastFrame(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsLastFrame:"), value)
}/* debug [instance_properties/setter]: isLastFrame */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNTrackingRequest */



