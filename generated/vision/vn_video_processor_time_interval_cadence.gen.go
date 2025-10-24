// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [VideoProcessorTimeIntervalCadence] class.
var (
	VideoProcessorTimeIntervalCadenceClass     _VideoProcessorTimeIntervalCadenceClass
	VideoProcessorTimeIntervalCadenceClassOnce sync.Once
)

func getVideoProcessorTimeIntervalCadenceClass() _VideoProcessorTimeIntervalCadenceClass {
	VideoProcessorTimeIntervalCadenceClassOnce.Do(func() {
		VideoProcessorTimeIntervalCadenceClass = _VideoProcessorTimeIntervalCadenceClass{objc.GetClass("VNVideoProcessorTimeIntervalCadence")}
	})
	return VideoProcessorTimeIntervalCadenceClass
}

type _VideoProcessorTimeIntervalCadenceClass struct {
	class objc.Class
}





// An interface definition for the [VideoProcessorTimeIntervalCadence] class.
type IVideoProcessorTimeIntervalCadence interface {
	IVideoProcessorCadence
	

	// properties:
	TimeInterval() float64
	Cadence() IVNVideoProcessorCadence
	SetCadence(value IVNVideoProcessorCadence)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (vc _VideoProcessorTimeIntervalCadenceClass) Alloc() VideoProcessorTimeIntervalCadence {
	rv := objc.Send[VideoProcessorTimeIntervalCadence](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VideoProcessorTimeIntervalCadenceClass) New() VideoProcessorTimeIntervalCadence {
	rv := objc.Send[VideoProcessorTimeIntervalCadence](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VideoProcessorTimeIntervalCadence) Init() VideoProcessorTimeIntervalCadence {
	rv := objc.Send[VideoProcessorTimeIntervalCadence](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VideoProcessorTimeIntervalCadence) Autorelease() VideoProcessorTimeIntervalCadence {
	rv := objc.Send[VideoProcessorTimeIntervalCadence](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVideoProcessorTimeIntervalCadence creates a new VideoProcessorTimeIntervalCadence instance.
func NewVideoProcessorTimeIntervalCadence() VideoProcessorTimeIntervalCadence {
	return getVideoProcessorTimeIntervalCadenceClass().New()
}





// An object that defines a time-based cadence for processing a video stream.


// An object that defines a time-based cadence for processing a video stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNVideoProcessor/TimeIntervalCadence
type VideoProcessorTimeIntervalCadence struct {
	VideoProcessorCadence
}

// VideoProcessorTimeIntervalCadenceFrom constructs a [VideoProcessorTimeIntervalCadence] from an unsafe.Pointer.
//
// An object that defines a time-based cadence for processing a video stream.
func VideoProcessorTimeIntervalCadenceFrom(ptr unsafe.Pointer) VideoProcessorTimeIntervalCadence {
	return VideoProcessorTimeIntervalCadence{
		VideoProcessorCadence: VideoProcessorCadenceFrom(ptr),
	}
}






// Creates a new time-based cadence with a time interval.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNVideoProcessor/TimeIntervalCadence/init(_:)
func NewVideoProcessorTimeIntervalCadenceWithTimeInterval(timeInterval float64) VideoProcessorTimeIntervalCadence {
	instance := getVideoProcessorTimeIntervalCadenceClass().Alloc()
	rv := objc.Send[VideoProcessorTimeIntervalCadence](instance.ID, objc.Sel("initWithTimeInterval:"), timeInterval)
	rv.Autorelease()
	return rv
}






















// The time interval of the cadence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNVideoProcessor/TimeIntervalCadence/timeInterval
func (v_ VideoProcessorTimeIntervalCadence) TimeInterval() float64 {
	rv := objc.Send[objc.ID](v_.ID, objc.Sel("timeInterval"))
	return rv
}


// The cadence the video processor maintains to process the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvideoprocessor/requestprocessingoptions/cadence
func (v_ VideoProcessorTimeIntervalCadence) Cadence() IVNVideoProcessorCadence {
	rv := objc.Send[VideoProcessorCadence](v_.ID, objc.Sel("cadence"))
	return rv
}


// The cadence the video processor maintains to process the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvideoprocessor/requestprocessingoptions/cadence
func (v_ VideoProcessorTimeIntervalCadence) SetCadence(value IVNVideoProcessorCadence) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setCadence:"), value)
}







