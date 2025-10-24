// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TrackOpticalFlowRequest] class.
var (
	TrackOpticalFlowRequestClass     _TrackOpticalFlowRequestClass
	TrackOpticalFlowRequestClassOnce sync.Once
)

func getTrackOpticalFlowRequestClass() _TrackOpticalFlowRequestClass {
	TrackOpticalFlowRequestClassOnce.Do(func() {
		TrackOpticalFlowRequestClass = _TrackOpticalFlowRequestClass{objc.GetClass("VNTrackOpticalFlowRequest")}
	})
	return TrackOpticalFlowRequestClass
}

type _TrackOpticalFlowRequestClass struct {
	class objc.Class
}

// An interface definition for the [TrackOpticalFlowRequest] class.
type ITrackOpticalFlowRequest interface {
	IStatefulRequest
	// properties:
	ComputationAccuracy() unsafe.Pointer
	SetComputationAccuracy(value unsafe.Pointer)
	KeepNetworkOutput() bool
	SetKeepNetworkOutput(value bool)
	OutputPixelFormat() uint32 /* not a class type */
	SetOutputPixelFormat(value uint32 /* not a class type */)
	Results() IVNPixelBufferObservation
	SetResults(value IVNPixelBufferObservation)
	// methods:
}

// An object that determines the direction change of vectors for each pixel from a previous to current image.
//
// This request works at the pixel level, so both images must have the same dimensions to successfully perform the request. Setting a region of interest isolates where to perform the change determination.


// An object that determines the direction change of vectors for each pixel from a previous to current image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTrackOpticalFlowRequest
type TrackOpticalFlowRequest struct {
	StatefulRequest
}

// TrackOpticalFlowRequestFrom constructs a [TrackOpticalFlowRequest] from an unsafe.Pointer.
//
// An object that determines the direction change of vectors for each pixel from a previous to current image.
func TrackOpticalFlowRequestFrom(ptr unsafe.Pointer) TrackOpticalFlowRequest {
	return TrackOpticalFlowRequest{
		StatefulRequest: StatefulRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (tc _TrackOpticalFlowRequestClass) Alloc() TrackOpticalFlowRequest {
	rv := objc.Send[TrackOpticalFlowRequest](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TrackOpticalFlowRequestClass) New() TrackOpticalFlowRequest {
	rv := objc.Send[TrackOpticalFlowRequest](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TrackOpticalFlowRequest) Init() TrackOpticalFlowRequest {
	rv := objc.Send[TrackOpticalFlowRequest](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TrackOpticalFlowRequest) Autorelease() TrackOpticalFlowRequest {
	rv := objc.Send[TrackOpticalFlowRequest](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTrackOpticalFlowRequest creates a new TrackOpticalFlowRequest instance.
func NewTrackOpticalFlowRequest() TrackOpticalFlowRequest {
	return getTrackOpticalFlowRequestClass().New()
}



// The level of accuracy to compute the optical flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vntrackopticalflowrequest/computationaccuracy-swift.property
func (t_ TrackOpticalFlowRequest) ComputationAccuracy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("computationAccuracy"))
	return rv
}


// The level of accuracy to compute the optical flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vntrackopticalflowrequest/computationaccuracy-swift.property
func (t_ TrackOpticalFlowRequest) SetComputationAccuracy(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setComputationAccuracy:"), value)
}


// A Boolean value that indicates the raw pixel buffer continues to emit from the network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vntrackopticalflowrequest/keepnetworkoutput
func (t_ TrackOpticalFlowRequest) KeepNetworkOutput() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("keepNetworkOutput"))
	return rv
}


// A Boolean value that indicates the raw pixel buffer continues to emit from the network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vntrackopticalflowrequest/keepnetworkoutput
func (t_ TrackOpticalFlowRequest) SetKeepNetworkOutput(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setKeepNetworkOutput:"), value)
}


// The pixel format type of the output value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vntrackopticalflowrequest/outputpixelformat
func (t_ TrackOpticalFlowRequest) OutputPixelFormat() uint32 /* not a class type */ {
	rv := objc.Send[uint32](t_.ID, objc.Sel("outputPixelFormat"))
	return rv
}


// The pixel format type of the output value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vntrackopticalflowrequest/outputpixelformat
func (t_ TrackOpticalFlowRequest) SetOutputPixelFormat(value uint32 /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setOutputPixelFormat:"), value)
}


// The optical flow results the request observes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vntrackopticalflowrequest/results
func (t_ TrackOpticalFlowRequest) Results() IVNPixelBufferObservation {
	rv := objc.Send[PixelBufferObservation](t_.ID, objc.Sel("results"))
	return rv
}


// The optical flow results the request observes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vntrackopticalflowrequest/results
func (t_ TrackOpticalFlowRequest) SetResults(value IVNPixelBufferObservation) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setResults:"), value)
}



