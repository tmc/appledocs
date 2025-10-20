// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	SupportedNumberOfTrackersAndReturnError(error_ unsafe.Pointer) uint
}

// The abstract superclass for image-analysis requests that track unique features across multiple images or video frames.
//
// Instantiate a tracking request subclass to perform object tracking across multiple frames of an image. After initialization, configure the degree of accuracy by setting , and provide observations you’d like to track by setting the initial bounding box.
//
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

// Alloc allocates a new instance without initialization.
func (tc _TrackingRequestClass) Alloc() TrackingRequest {
	rv := objc.Send[TrackingRequest](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Returns the maximum number of simultaneous trackers for the request.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTrackingRequest/supportedNumber(ofTrackersAndReturnError:)
func (t_ TrackingRequest) SupportedNumberOfTrackersAndReturnError(error_ unsafe.Pointer) uint {
	rv := objc.Send[uint](t_.ID, objc.Sel("supportedNumberOfTrackersAndReturnError:"), error_)
	return rv
}



