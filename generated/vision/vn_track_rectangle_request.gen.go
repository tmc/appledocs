// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [TrackRectangleRequest] class.
type ITrackRectangleRequest interface {
	ITrackingRequest
}

// An image-analysis request that tracks movement of a previously identified rectangular object across multiple images or video frames.
//
// Use this type of request to track the bounding boxes of rectangles throughout a sequence of images. Vision returns locations for rectangles found in all orientations and sizes.
//
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

// Alloc allocates a new instance without initialization.
func (tc _TrackRectangleRequestClass) Alloc() TrackRectangleRequest {
	rv := objc.Send[TrackRectangleRequest](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates a new rectangle tracking request with a rectangle observation.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTrackRectangleRequest/init(rectangleObservation:completionHandler:)
func NewTrackRectangleRequestWithRectangleObservationCompletionHandler(observation unsafe.Pointer, completionHandler unsafe.Pointer) TrackRectangleRequest {
	instance := getTrackRectangleRequestClass().Alloc()
	rv := objc.Send[TrackRectangleRequest](instance.ID, objc.Sel("initWithRectangleObservation:completionHandler:"), observation, completionHandler)
	rv.Autorelease()
	return rv
}



