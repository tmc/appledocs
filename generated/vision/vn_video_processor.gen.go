// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [VideoProcessor] class.
var (
	VideoProcessorClass     _VideoProcessorClass
	VideoProcessorClassOnce sync.Once
)

func getVideoProcessorClass() _VideoProcessorClass {
	VideoProcessorClassOnce.Do(func() {
		VideoProcessorClass = _VideoProcessorClass{objc.GetClass("VNVideoProcessor")}
	})
	return VideoProcessorClass
}

type _VideoProcessorClass struct {
	class objc.Class
}





// An interface definition for the [VideoProcessor] class.
type IVideoProcessor interface {
	objectivec.IObject
	

	// properties:


	

	// methods:
	AddRequestProcessingOptionsError(request IVNRequest, processingOptions IVNVideoProcessorRequestProcessingOptions, error_ foundation.foundation.INSError) bool
	AnalyzeTimeRangeError(timeRange objectivec.IObject, error_ foundation.foundation.INSError) bool
	Cancel()
	RemoveRequestError(request IVNRequest, error_ foundation.foundation.INSError) bool


}





// Alloc allocates a new instance without initialization.
func (vc _VideoProcessorClass) Alloc() VideoProcessor {
	rv := objc.Send[VideoProcessor](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VideoProcessorClass) New() VideoProcessor {
	rv := objc.Send[VideoProcessor](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VideoProcessor) Init() VideoProcessor {
	rv := objc.Send[VideoProcessor](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VideoProcessor) Autorelease() VideoProcessor {
	rv := objc.Send[VideoProcessor](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVideoProcessor creates a new VideoProcessor instance.
func NewVideoProcessor() VideoProcessor {
	return getVideoProcessorClass().New()
}





// An object that performs offline analysis of video content.


// An object that performs offline analysis of video content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNVideoProcessor
type VideoProcessor struct {
	objectivec.Object
}

// VideoProcessorFrom constructs a [VideoProcessor] from an unsafe.Pointer.
//
// An object that performs offline analysis of video content.
func VideoProcessorFrom(ptr unsafe.Pointer) VideoProcessor {
	return VideoProcessor{objectivec.Object{objc.ID(ptr)}}
}






// Creates a video processor to perform Vision requests against the specified video asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNVideoProcessor/init(url:)
func NewVideoProcessorWithURL(videoURL foundation.foundation.INSURL) VideoProcessor {
	instance := getVideoProcessorClass().Alloc()
	rv := objc.Send[VideoProcessor](instance.ID, objc.Sel("initWithURL:"), videoURL)
	rv.Autorelease()
	return rv
}

















// Adds a request with processing options to the video processor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNVideoProcessor/addRequest(_:processingOptions:)
func (v_ VideoProcessor) AddRequestProcessingOptionsError(request IVNRequest, processingOptions IVNVideoProcessorRequestProcessingOptions, error_ foundation.foundation.INSError) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("addRequest:processingOptions:error:"), request, processingOptions, error_)
	return rv
}


// Analyzes a time range of video content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNVideoProcessor/analyze(_:)
func (v_ VideoProcessor) AnalyzeTimeRangeError(timeRange objectivec.IObject, error_ foundation.foundation.INSError) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("analyzeTimeRange:error:"), timeRange, error_)
	return rv
}


// Cancels the video processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNVideoProcessor/cancel()
func (v_ VideoProcessor) Cancel() {
	objc.Send[objc.ID](v_.ID, objc.Sel("cancel"))
}


// Removes a Vision request from the video processor’s request queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNVideoProcessor/removeRequest(_:)
func (v_ VideoProcessor) RemoveRequestError(request IVNRequest, error_ foundation.foundation.INSError) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("removeRequest:error:"), request, error_)
	return rv
}












