// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [TrackingRequest] class.
type ITrackingRequest interface {
	IImageBasedRequest
	

	// properties:
	InputObservation() IVNDetectedObjectObservation
	SetInputObservation(value IVNDetectedObjectObservation)
	LastFrame() bool
	SetLastFrame(value bool)
	TrackingLevel() RequestTrackingLevel
	SetTrackingLevel(value RequestTrackingLevel)
	IsLastFrame() bool
	SetIsLastFrame(value bool)


	

	// methods:
	SupportedNumberOfTrackersAndReturnError(error_ objectivec.IObject) uint


}





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




















// Returns the maximum number of simultaneous trackers for the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTrackingRequest/supportedNumber(ofTrackersAndReturnError:)
func (t_ TrackingRequest) SupportedNumberOfTrackersAndReturnError(error_ objectivec.IObject) uint {
	rv := objc.Send[uint](t_.ID, objc.Sel("supportedNumberOfTrackersAndReturnError:"), error_)
	return rv
}







// The observation object defining a region to track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTrackingRequest/inputObservation
func (t_ TrackingRequest) InputObservation() IVNDetectedObjectObservation {
	rv := objc.Send[DetectedObjectObservation](t_.ID, objc.Sel("inputObservation"))
	return rv
}


// The observation object defining a region to track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTrackingRequest/inputObservation
func (t_ TrackingRequest) SetInputObservation(value IVNDetectedObjectObservation) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setInputObservation:"), value)
}


// A Boolean that indicates the last frame in a tracking sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTrackingRequest/isLastFrame
func (t_ TrackingRequest) LastFrame() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("lastFrame"))
	return rv
}


// A Boolean that indicates the last frame in a tracking sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTrackingRequest/isLastFrame
func (t_ TrackingRequest) SetLastFrame(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLastFrame:"), value)
}


// A value for specifying whether to prioritize speed or location accuracy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTrackingRequest/trackingLevel
func (t_ TrackingRequest) TrackingLevel() RequestTrackingLevel {
	rv := objc.Send[RequestTrackingLevel](t_.ID, objc.Sel("trackingLevel"))
	return rv
}


// A value for specifying whether to prioritize speed or location accuracy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTrackingRequest/trackingLevel
func (t_ TrackingRequest) SetTrackingLevel(value RequestTrackingLevel) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTrackingLevel:"), value)
}


// A Boolean that indicates the last frame in a tracking sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vntrackingrequest/islastframe
func (t_ TrackingRequest) IsLastFrame() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isLastFrame"))
	return rv
}


// A Boolean that indicates the last frame in a tracking sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vntrackingrequest/islastframe
func (t_ TrackingRequest) SetIsLastFrame(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsLastFrame:"), value)
}








