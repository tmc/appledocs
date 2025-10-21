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
}

// An object that determines the direction change of vectors for each pixel from a previous to current image.
//
// This request works at the pixel level, so both images must have the same dimensions to successfully perform the request. Setting a region of interest isolates where to perform the change determination.
//
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


// The optical flow results the request observes.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntrackopticalflowrequest/results
func (t_ TrackOpticalFlowRequest) Results() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("results"))
	return rv
}


// SetResults sets the value of the results property.
// The optical flow results the request observes.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntrackopticalflowrequest/results
func (t_ TrackOpticalFlowRequest) SetResults(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setResults:"), value)
}

// The pixel format type of the output value.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntrackopticalflowrequest/outputpixelformat
func (t_ TrackOpticalFlowRequest) OutputPixelFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("outputPixelFormat"))
	return rv
}


// SetOutputPixelFormat sets the value of the outputPixelFormat property.
// The pixel format type of the output value.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntrackopticalflowrequest/outputpixelformat
func (t_ TrackOpticalFlowRequest) SetOutputPixelFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setOutputPixelFormat:"), value)
}

// The level of accuracy to compute the optical flow.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntrackopticalflowrequest/computationaccuracy-swift.property
func (t_ TrackOpticalFlowRequest) ComputationAccuracy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("computationAccuracy"))
	return rv
}


// SetComputationAccuracy sets the value of the computationAccuracy property.
// The level of accuracy to compute the optical flow.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntrackopticalflowrequest/computationaccuracy-swift.property
func (t_ TrackOpticalFlowRequest) SetComputationAccuracy(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setComputationAccuracy:"), value)
}

// A Boolean value that indicates the raw pixel buffer continues to emit from the network.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntrackopticalflowrequest/keepnetworkoutput
func (t_ TrackOpticalFlowRequest) KeepNetworkOutput() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("keepNetworkOutput"))
	return rv
}


// SetKeepNetworkOutput sets the value of the keepNetworkOutput property.
// A Boolean value that indicates the raw pixel buffer continues to emit from the network.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntrackopticalflowrequest/keepnetworkoutput
func (t_ TrackOpticalFlowRequest) SetKeepNetworkOutput(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setKeepNetworkOutput:"), value)
}



