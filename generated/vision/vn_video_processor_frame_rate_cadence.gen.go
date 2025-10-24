// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [VideoProcessorFrameRateCadence] class.
var (
	VideoProcessorFrameRateCadenceClass     _VideoProcessorFrameRateCadenceClass
	VideoProcessorFrameRateCadenceClassOnce sync.Once
)

func getVideoProcessorFrameRateCadenceClass() _VideoProcessorFrameRateCadenceClass {
	VideoProcessorFrameRateCadenceClassOnce.Do(func() {
		VideoProcessorFrameRateCadenceClass = _VideoProcessorFrameRateCadenceClass{objc.GetClass("VNVideoProcessorFrameRateCadence")}
	})
	return VideoProcessorFrameRateCadenceClass
}

type _VideoProcessorFrameRateCadenceClass struct {
	class objc.Class
}





// An interface definition for the [VideoProcessorFrameRateCadence] class.
type IVideoProcessorFrameRateCadence interface {
	IVideoProcessorCadence
	

	// properties:
	FrameRate() int
	Cadence() IVNVideoProcessorCadence
	SetCadence(value IVNVideoProcessorCadence)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (vc _VideoProcessorFrameRateCadenceClass) Alloc() VideoProcessorFrameRateCadence {
	rv := objc.Send[VideoProcessorFrameRateCadence](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VideoProcessorFrameRateCadenceClass) New() VideoProcessorFrameRateCadence {
	rv := objc.Send[VideoProcessorFrameRateCadence](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VideoProcessorFrameRateCadence) Init() VideoProcessorFrameRateCadence {
	rv := objc.Send[VideoProcessorFrameRateCadence](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VideoProcessorFrameRateCadence) Autorelease() VideoProcessorFrameRateCadence {
	rv := objc.Send[VideoProcessorFrameRateCadence](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVideoProcessorFrameRateCadence creates a new VideoProcessorFrameRateCadence instance.
func NewVideoProcessorFrameRateCadence() VideoProcessorFrameRateCadence {
	return getVideoProcessorFrameRateCadenceClass().New()
}





// An object that defines a frame-based cadence for processing a video stream.


// An object that defines a frame-based cadence for processing a video stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNVideoProcessor/FrameRateCadence
type VideoProcessorFrameRateCadence struct {
	VideoProcessorCadence
}

// VideoProcessorFrameRateCadenceFrom constructs a [VideoProcessorFrameRateCadence] from an unsafe.Pointer.
//
// An object that defines a frame-based cadence for processing a video stream.
func VideoProcessorFrameRateCadenceFrom(ptr unsafe.Pointer) VideoProcessorFrameRateCadence {
	return VideoProcessorFrameRateCadence{
		VideoProcessorCadence: VideoProcessorCadenceFrom(ptr),
	}
}






// Creates a new frame-based cadence with a frame rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNVideoProcessor/FrameRateCadence/init(_:)
func NewVideoProcessorFrameRateCadenceWithFrameRate(frameRate int) VideoProcessorFrameRateCadence {
	instance := getVideoProcessorFrameRateCadenceClass().Alloc()
	rv := objc.Send[VideoProcessorFrameRateCadence](instance.ID, objc.Sel("initWithFrameRate:"), frameRate)
	rv.Autorelease()
	return rv
}






















// The frame rate at which to process video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNVideoProcessor/FrameRateCadence/frameRate
func (v_ VideoProcessorFrameRateCadence) FrameRate() int {
	rv := objc.Send[int](v_.ID, objc.Sel("frameRate"))
	return rv
}


// The cadence the video processor maintains to process the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvideoprocessor/requestprocessingoptions/cadence
func (v_ VideoProcessorFrameRateCadence) Cadence() IVNVideoProcessorCadence {
	rv := objc.Send[VideoProcessorCadence](v_.ID, objc.Sel("cadence"))
	return rv
}


// The cadence the video processor maintains to process the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvideoprocessor/requestprocessingoptions/cadence
func (v_ VideoProcessorFrameRateCadence) SetCadence(value IVNVideoProcessorCadence) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setCadence:"), value)
}







