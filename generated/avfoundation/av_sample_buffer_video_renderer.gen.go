// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVSampleBufferVideoRenderer */


/* debug [class_header]: Header for AVSampleBufferVideoRenderer */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SampleBufferVideoRenderer */
// An interface definition for the [SampleBufferVideoRenderer] class.
type ISampleBufferVideoRenderer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SampleBufferVideoRenderer */
	// properties:
	Error() Error
	RecommendedPixelBufferAttributes() foundation.IDictionary
	RequiresFlushToResumeDecoding() bool
	Status() QueuedSampleBufferRenderingStatus
	PresentationTimeExpectation() objectivec.IObject
	SetPresentationTimeExpectation(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SampleBufferVideoRenderer */
	// methods:
	CopyDisplayedPixelBuffer() PixelBufferRef /* not a class type */
	ExpectMinimumUpcomingSampleBufferPresentationTime(minimumUpcomingPresentationTime objc.IObject /* cross-framework: Time */)
	ExpectMonotonicallyIncreasingUpcomingSampleBufferPresentationTimes()
	FlushWithRemovalOfDisplayedImageCompletionHandler(removeDisplayedImage bool, handler unsafe.Pointer)
	LoadVideoPerformanceMetricsWithCompletionHandler(completionHandler unsafe.Pointer)
	ResetUpcomingSampleBufferPresentationTimeExpectations()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SampleBufferVideoRenderer */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SampleBufferVideoRenderer */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SampleBufferVideoRenderer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SampleBufferVideoRenderer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SampleBufferVideoRenderer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SampleBufferVideoRenderer */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferVideoRenderer/displayedPixelBuffer()
func (s_ SampleBufferVideoRenderer) CopyDisplayedPixelBuffer() PixelBufferRef /* not a class type */ {
	rv := objc.Send[PixelBufferRef](s_.ID, objc.Sel("copyDisplayedPixelBuffer"))
	return rv
}/* debug [instance_methods/method]: CopyDisplayedPixelBuffer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferVideoRenderer/expectMinimumUpcomingSampleBufferPresentationTime:
func (s_ SampleBufferVideoRenderer) ExpectMinimumUpcomingSampleBufferPresentationTime(minimumUpcomingPresentationTime objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("expectMinimumUpcomingSampleBufferPresentationTime:"), minimumUpcomingPresentationTime)
}/* debug [instance_methods/method]: ExpectMinimumUpcomingSampleBufferPresentationTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferVideoRenderer/expectMonotonicallyIncreasingUpcomingSampleBufferPresentationTimes
func (s_ SampleBufferVideoRenderer) ExpectMonotonicallyIncreasingUpcomingSampleBufferPresentationTimes() {
	objc.Send[objc.ID](s_.ID, objc.Sel("expectMonotonicallyIncreasingUpcomingSampleBufferPresentationTimes"))
}/* debug [instance_methods/method]: ExpectMonotonicallyIncreasingUpcomingSampleBufferPresentationTimes */


// Tells the video renderer to discard pending enqueued sample buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferVideoRenderer/flush(removingDisplayedImage:completionHandler:)
func (s_ SampleBufferVideoRenderer) FlushWithRemovalOfDisplayedImageCompletionHandler(removeDisplayedImage bool, handler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("flushWithRemovalOfDisplayedImage:completionHandler:"), removeDisplayedImage, handler)
}/* debug [instance_methods/method]: FlushWithRemovalOfDisplayedImageCompletionHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferVideoRenderer/loadVideoPerformanceMetrics(completionHandler:)
func (s_ SampleBufferVideoRenderer) LoadVideoPerformanceMetricsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("loadVideoPerformanceMetricsWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: LoadVideoPerformanceMetricsWithCompletionHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferVideoRenderer/resetUpcomingSampleBufferPresentationTimeExpectations
func (s_ SampleBufferVideoRenderer) ResetUpcomingSampleBufferPresentationTimeExpectations() {
	objc.Send[objc.ID](s_.ID, objc.Sel("resetUpcomingSampleBufferPresentationTimeExpectations"))
}/* debug [instance_methods/method]: ResetUpcomingSampleBufferPresentationTimeExpectations */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SampleBufferVideoRenderer */

// An object the describes the error that caused the rendering failure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferVideoRenderer/error
func (s_ SampleBufferVideoRenderer) Error() Error {
	rv := objc.Send[Error](s_.ID, objc.Sel("error"))
	return rv
}/* debug [instance_properties/getter]: error */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferVideoRenderer/recommendedPixelBufferAttributes-6326f
func (s_ SampleBufferVideoRenderer) RecommendedPixelBufferAttributes() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](s_.ID, objc.Sel("recommendedPixelBufferAttributes"))
	return rv
}/* debug [instance_properties/getter]: recommendedPixelBufferAttributes */


// A Boolean value that Indicates whether the renderer requires flushing to continue decoding frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferVideoRenderer/requiresFlushToResumeDecoding
func (s_ SampleBufferVideoRenderer) RequiresFlushToResumeDecoding() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("requiresFlushToResumeDecoding"))
	return rv
}/* debug [instance_properties/getter]: requiresFlushToResumeDecoding */


// A status value that indicates whether this object can enqueue and render sample buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferVideoRenderer/status
func (s_ SampleBufferVideoRenderer) Status() QueuedSampleBufferRenderingStatus {
	rv := objc.Send[QueuedSampleBufferRenderingStatus](s_.ID, objc.Sel("status"))
	return rv
}/* debug [instance_properties/getter]: status */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebuffervideorenderer/presentationtimeexpectation-swift.property
func (s_ SampleBufferVideoRenderer) PresentationTimeExpectation() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("presentationTimeExpectation"))
	return rv
}/* debug [instance_properties/getter]: presentationTimeExpectation */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebuffervideorenderer/presentationtimeexpectation-swift.property
func (s_ SampleBufferVideoRenderer) SetPresentationTimeExpectation(value objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPresentationTimeExpectation:"), value)
}/* debug [instance_properties/setter]: presentationTimeExpectation */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVSampleBufferVideoRenderer */



