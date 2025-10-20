// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

// An image-analysis request that tracks the movement of a previously identified object across multiple images or video frames.
//
// Use this type of request to track the bounding boxes around objects previously identified in an image. Vision attempts to locate the same object from the input observation throughout the sequence.
//
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

// Alloc allocates a new instance without initialization.
func (tc _TrackObjectRequestClass) Alloc() TrackObjectRequest {
	rv := objc.Send[TrackObjectRequest](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




