// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [VideoProcessorCadence] class.
var (
	VideoProcessorCadenceClass     _VideoProcessorCadenceClass
	VideoProcessorCadenceClassOnce sync.Once
)

func getVideoProcessorCadenceClass() _VideoProcessorCadenceClass {
	VideoProcessorCadenceClassOnce.Do(func() {
		VideoProcessorCadenceClass = _VideoProcessorCadenceClass{objc.GetClass("VNVideoProcessorCadence")}
	})
	return VideoProcessorCadenceClass
}

type _VideoProcessorCadenceClass struct {
	class objc.Class
}





// An interface definition for the [VideoProcessorCadence] class.
type IVideoProcessorCadence interface {
	objectivec.IObject
	

	// properties:
	Cadence() IVNVideoProcessorCadence
	SetCadence(value IVNVideoProcessorCadence)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (vc _VideoProcessorCadenceClass) Alloc() VideoProcessorCadence {
	rv := objc.Send[VideoProcessorCadence](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VideoProcessorCadenceClass) New() VideoProcessorCadence {
	rv := objc.Send[VideoProcessorCadence](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VideoProcessorCadence) Init() VideoProcessorCadence {
	rv := objc.Send[VideoProcessorCadence](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VideoProcessorCadence) Autorelease() VideoProcessorCadence {
	rv := objc.Send[VideoProcessorCadence](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVideoProcessorCadence creates a new VideoProcessorCadence instance.
func NewVideoProcessorCadence() VideoProcessorCadence {
	return getVideoProcessorCadenceClass().New()
}





// An object that defines the cadence at which to process video.


// An object that defines the cadence at which to process video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNVideoProcessor/Cadence
type VideoProcessorCadence struct {
	objectivec.Object
}

// VideoProcessorCadenceFrom constructs a [VideoProcessorCadence] from an unsafe.Pointer.
//
// An object that defines the cadence at which to process video.
func VideoProcessorCadenceFrom(ptr unsafe.Pointer) VideoProcessorCadence {
	return VideoProcessorCadence{objectivec.Object{objc.ID(ptr)}}
}

























// The cadence the video processor maintains to process the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvideoprocessor/requestprocessingoptions/cadence
func (v_ VideoProcessorCadence) Cadence() IVNVideoProcessorCadence {
	rv := objc.Send[VideoProcessorCadence](v_.ID, objc.Sel("cadence"))
	return rv
}


// The cadence the video processor maintains to process the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvideoprocessor/requestprocessingoptions/cadence
func (v_ VideoProcessorCadence) SetCadence(value IVNVideoProcessorCadence) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setCadence:"), value)
}








