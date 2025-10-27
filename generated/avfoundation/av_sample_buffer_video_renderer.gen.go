// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [SampleBufferVideoRenderer] class.
var (
	SampleBufferVideoRendererClass     _SampleBufferVideoRendererClass
	SampleBufferVideoRendererClassOnce sync.Once
)

func getSampleBufferVideoRendererClass() _SampleBufferVideoRendererClass {
	SampleBufferVideoRendererClassOnce.Do(func() {
		SampleBufferVideoRendererClass = _SampleBufferVideoRendererClass{objc.GetClass("AVSampleBufferVideoRenderer")}
	})
	return SampleBufferVideoRendererClass
}

type _SampleBufferVideoRendererClass struct {
	class objc.Class
}





// An interface definition for the [SampleBufferVideoRenderer] class.
type ISampleBufferVideoRenderer interface {
	objectivec.IObject
	

	// properties:
	Error() foundation.foundation.INSError
	RecommendedPixelBufferAttributes() foundation.IDictionary
	RequiresFlushToResumeDecoding() bool
	Status() QueuedSampleBufferRenderingStatus
	PresentationTimeExpectation() objectivec.IObject
	SetPresentationTimeExpectation(value objectivec.IObject)


	

	// methods:
	CopyDisplayedPixelBuffer() PixelBufferRef /* not a class type */
	ExpectMinimumUpcomingSampleBufferPresentationTime(minimumUpcomingPresentationTime objectivec.IObject)
	ExpectMonotonicallyIncreasingUpcomingSampleBufferPresentationTimes()
	FlushWithRemovalOfDisplayedImageCompletionHandler(removeDisplayedImage bool, handler unsafe.Pointer)
	LoadVideoPerformanceMetricsWithCompletionHandler(completionHandler unsafe.Pointer)
	ResetUpcomingSampleBufferPresentationTimeExpectations()


}





// Alloc allocates a new instance without initialization.
func (sc _SampleBufferVideoRendererClass) Alloc() SampleBufferVideoRenderer {
	rv := objc.Send[SampleBufferVideoRenderer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SampleBufferVideoRendererClass) New() SampleBufferVideoRenderer {
	rv := objc.Send[SampleBufferVideoRenderer](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SampleBufferVideoRenderer) Init() SampleBufferVideoRenderer {
	rv := objc.Send[SampleBufferVideoRenderer](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SampleBufferVideoRenderer) Autorelease() SampleBufferVideoRenderer {
	rv := objc.Send[SampleBufferVideoRenderer](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSampleBufferVideoRenderer creates a new SampleBufferVideoRenderer instance.
func NewSampleBufferVideoRenderer() SampleBufferVideoRenderer {
	return getSampleBufferVideoRendererClass().New()
}





// An object that enqueues video sample buffers for rendering.


// An object that enqueues video sample buffers for rendering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferVideoRenderer
type SampleBufferVideoRenderer struct {
	objectivec.Object
}

// SampleBufferVideoRendererFrom constructs a [SampleBufferVideoRenderer] from an unsafe.Pointer.
//
// An object that enqueues video sample buffers for rendering.
func SampleBufferVideoRendererFrom(ptr unsafe.Pointer) SampleBufferVideoRenderer {
	return SampleBufferVideoRenderer{objectivec.Object{objc.ID(ptr)}}
}




















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferVideoRenderer/displayedPixelBuffer()
func (s_ SampleBufferVideoRenderer) CopyDisplayedPixelBuffer() PixelBufferRef /* not a class type */ {
	rv := objc.Send[PixelBufferRef](s_.ID, objc.Sel("copyDisplayedPixelBuffer"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferVideoRenderer/expectMinimumUpcomingSampleBufferPresentationTime:
func (s_ SampleBufferVideoRenderer) ExpectMinimumUpcomingSampleBufferPresentationTime(minimumUpcomingPresentationTime objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("expectMinimumUpcomingSampleBufferPresentationTime:"), minimumUpcomingPresentationTime)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferVideoRenderer/expectMonotonicallyIncreasingUpcomingSampleBufferPresentationTimes
func (s_ SampleBufferVideoRenderer) ExpectMonotonicallyIncreasingUpcomingSampleBufferPresentationTimes() {
	objc.Send[objc.ID](s_.ID, objc.Sel("expectMonotonicallyIncreasingUpcomingSampleBufferPresentationTimes"))
}


// Tells the video renderer to discard pending enqueued sample buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferVideoRenderer/flush(removingDisplayedImage:completionHandler:)
func (s_ SampleBufferVideoRenderer) FlushWithRemovalOfDisplayedImageCompletionHandler(removeDisplayedImage bool, handler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("flushWithRemovalOfDisplayedImage:completionHandler:"), removeDisplayedImage, handler)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferVideoRenderer/loadVideoPerformanceMetrics(completionHandler:)
func (s_ SampleBufferVideoRenderer) LoadVideoPerformanceMetricsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("loadVideoPerformanceMetricsWithCompletionHandler:"), completionHandler)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferVideoRenderer/resetUpcomingSampleBufferPresentationTimeExpectations
func (s_ SampleBufferVideoRenderer) ResetUpcomingSampleBufferPresentationTimeExpectations() {
	objc.Send[objc.ID](s_.ID, objc.Sel("resetUpcomingSampleBufferPresentationTimeExpectations"))
}







// An object the describes the error that caused the rendering failure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferVideoRenderer/error
func (s_ SampleBufferVideoRenderer) Error() foundation.foundation.INSError {
	rv := objc.Send[foundation.NSError](s_.ID, objc.Sel("error"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferVideoRenderer/recommendedPixelBufferAttributes-6326f
func (s_ SampleBufferVideoRenderer) RecommendedPixelBufferAttributes() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](s_.ID, objc.Sel("recommendedPixelBufferAttributes"))
	return rv
}


// A Boolean value that Indicates whether the renderer requires flushing to continue decoding frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferVideoRenderer/requiresFlushToResumeDecoding
func (s_ SampleBufferVideoRenderer) RequiresFlushToResumeDecoding() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("requiresFlushToResumeDecoding"))
	return rv
}


// A status value that indicates whether this object can enqueue and render sample buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferVideoRenderer/status
func (s_ SampleBufferVideoRenderer) Status() QueuedSampleBufferRenderingStatus {
	rv := objc.Send[QueuedSampleBufferRenderingStatus](s_.ID, objc.Sel("status"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebuffervideorenderer/presentationtimeexpectation-swift.property
func (s_ SampleBufferVideoRenderer) PresentationTimeExpectation() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("presentationTimeExpectation"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebuffervideorenderer/presentationtimeexpectation-swift.property
func (s_ SampleBufferVideoRenderer) SetPresentationTimeExpectation(value objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPresentationTimeExpectation:"), value)
}








