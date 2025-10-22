// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [StatefulRequest] class.
var (
	StatefulRequestClass     _StatefulRequestClass
	StatefulRequestClassOnce sync.Once
)

func getStatefulRequestClass() _StatefulRequestClass {
	StatefulRequestClassOnce.Do(func() {
		StatefulRequestClass = _StatefulRequestClass{objc.GetClass("VNStatefulRequest")}
	})
	return StatefulRequestClass
}

type _StatefulRequestClass struct {
	class objc.Class
}

// An interface definition for the [StatefulRequest] class.
type IStatefulRequest interface {
	IImageBasedRequest
	FrameAnalysisSpacing() unsafe.Pointer
	SetFrameAnalysisSpacing(value unsafe.Pointer)
	MinimumLatencyFrameCount() int
	SetMinimumLatencyFrameCount(value int)
}

// An abstract request type that builds evidence of a condition over time.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNStatefulRequest
type StatefulRequest struct {
	ImageBasedRequest
}

// StatefulRequestFrom constructs a [StatefulRequest] from an unsafe.Pointer.
//
// An abstract request type that builds evidence of a condition over time.
func StatefulRequestFrom(ptr unsafe.Pointer) StatefulRequest {
	return StatefulRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _StatefulRequestClass) Alloc() StatefulRequest {
	rv := objc.Send[StatefulRequest](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _StatefulRequestClass) New() StatefulRequest {
	rv := objc.Send[StatefulRequest](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ StatefulRequest) Init() StatefulRequest {
	rv := objc.Send[StatefulRequest](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ StatefulRequest) Autorelease() StatefulRequest {
	rv := objc.Send[StatefulRequest](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewStatefulRequest creates a new StatefulRequest instance.
func NewStatefulRequest() StatefulRequest {
	return getStatefulRequestClass().New()
}




// Initializes a video-based request.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNStatefulRequest/init(frameAnalysisSpacing:completionHandler:)
func NewStatefulRequestWithFrameAnalysisSpacingCompletionHandler(frameAnalysisSpacing unsafe.Pointer, completionHandler unsafe.Pointer) StatefulRequest {
	instance := getStatefulRequestClass().Alloc()
	rv := objc.Send[StatefulRequest](instance.ID, objc.Sel("initWithFrameAnalysisSpacing:completionHandler:"), frameAnalysisSpacing, completionHandler)
	rv.Autorelease()
	return rv
}


// A time value that indicates the interval between analysis operations.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnstatefulrequest/frameanalysisspacing
func (s_ StatefulRequest) FrameAnalysisSpacing() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("frameAnalysisSpacing"))
	return rv
}


// SetFrameAnalysisSpacing sets the value of the frameAnalysisSpacing property.
// A time value that indicates the interval between analysis operations.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnstatefulrequest/frameanalysisspacing
func (s_ StatefulRequest) SetFrameAnalysisSpacing(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setFrameAnalysisSpacing:"), value)
}

// The minimum number of frames a request processes before reporting an observation.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnstatefulrequest/minimumlatencyframecount
func (s_ StatefulRequest) MinimumLatencyFrameCount() int {
	rv := objc.Send[int](s_.ID, objc.Sel("minimumLatencyFrameCount"))
	return rv
}


// SetMinimumLatencyFrameCount sets the value of the minimumLatencyFrameCount property.
// The minimum number of frames a request processes before reporting an observation.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnstatefulrequest/minimumlatencyframecount
func (s_ StatefulRequest) SetMinimumLatencyFrameCount(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMinimumLatencyFrameCount:"), value)
}


