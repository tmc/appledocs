// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [VideoProcessorRequestProcessingOptions] class.
var (
	VideoProcessorRequestProcessingOptionsClass     _VideoProcessorRequestProcessingOptionsClass
	VideoProcessorRequestProcessingOptionsClassOnce sync.Once
)

func getVideoProcessorRequestProcessingOptionsClass() _VideoProcessorRequestProcessingOptionsClass {
	VideoProcessorRequestProcessingOptionsClassOnce.Do(func() {
		VideoProcessorRequestProcessingOptionsClass = _VideoProcessorRequestProcessingOptionsClass{objc.GetClass("VNVideoProcessorRequestProcessingOptions")}
	})
	return VideoProcessorRequestProcessingOptionsClass
}

type _VideoProcessorRequestProcessingOptionsClass struct {
	class objc.Class
}





// An interface definition for the [VideoProcessorRequestProcessingOptions] class.
type IVideoProcessorRequestProcessingOptions interface {
	objectivec.IObject
	

	// properties:
	Cadence() IVNVideoProcessorCadence
	SetCadence(value IVNVideoProcessorCadence)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (vc _VideoProcessorRequestProcessingOptionsClass) Alloc() VideoProcessorRequestProcessingOptions {
	rv := objc.Send[VideoProcessorRequestProcessingOptions](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VideoProcessorRequestProcessingOptionsClass) New() VideoProcessorRequestProcessingOptions {
	rv := objc.Send[VideoProcessorRequestProcessingOptions](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VideoProcessorRequestProcessingOptions) Init() VideoProcessorRequestProcessingOptions {
	rv := objc.Send[VideoProcessorRequestProcessingOptions](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VideoProcessorRequestProcessingOptions) Autorelease() VideoProcessorRequestProcessingOptions {
	rv := objc.Send[VideoProcessorRequestProcessingOptions](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVideoProcessorRequestProcessingOptions creates a new VideoProcessorRequestProcessingOptions instance.
func NewVideoProcessorRequestProcessingOptions() VideoProcessorRequestProcessingOptions {
	return getVideoProcessorRequestProcessingOptionsClass().New()
}





// An object that defines a video processor’s configuration options.


// An object that defines a video processor’s configuration options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNVideoProcessor/RequestProcessingOptions
type VideoProcessorRequestProcessingOptions struct {
	objectivec.Object
}

// VideoProcessorRequestProcessingOptionsFrom constructs a [VideoProcessorRequestProcessingOptions] from an unsafe.Pointer.
//
// An object that defines a video processor’s configuration options.
func VideoProcessorRequestProcessingOptionsFrom(ptr unsafe.Pointer) VideoProcessorRequestProcessingOptions {
	return VideoProcessorRequestProcessingOptions{objectivec.Object{objc.ID(ptr)}}
}

























// The cadence the video processor maintains to process the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNVideoProcessor/RequestProcessingOptions/cadence
func (v_ VideoProcessorRequestProcessingOptions) Cadence() IVNVideoProcessorCadence {
	rv := objc.Send[VideoProcessorCadence](v_.ID, objc.Sel("cadence"))
	return rv
}


// The cadence the video processor maintains to process the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNVideoProcessor/RequestProcessingOptions/cadence
func (v_ VideoProcessorRequestProcessingOptions) SetCadence(value IVNVideoProcessorCadence) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setCadence:"), value)
}








