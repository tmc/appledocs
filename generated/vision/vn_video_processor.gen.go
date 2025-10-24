// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VNVideoProcessor */


/* debug [class_header]: Header for VNVideoProcessor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VideoProcessor */
// An interface definition for the [VideoProcessor] class.
type IVideoProcessor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for VideoProcessor */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VideoProcessor */
	// methods:
	AddRequestProcessingOptionsError(request IVNRequest, processingOptions IVNVideoProcessorRequestProcessingOptions, error_ objectivec.IObject) bool
	AnalyzeTimeRangeError(timeRange TimeRange /* not a class type */, error_ objectivec.IObject) bool
	Cancel()
	RemoveRequestError(request IVNRequest, error_ objectivec.IObject) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VideoProcessor */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VideoProcessor */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VideoProcessor */

// Creates a video processor to perform Vision requests against the specified video asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNVideoProcessor/init(url:)
func NewVideoProcessorWithURL(videoURL objc.IObject /* cross-framework: NSURL */) VideoProcessor {
	instance := getVideoProcessorClass().Alloc()
	rv := objc.Send[VideoProcessor](instance.ID, objc.Sel("initWithURL:"), videoURL)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewVideoProcessorWithURL */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VideoProcessor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VideoProcessor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VideoProcessor */

// Adds a request with processing options to the video processor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNVideoProcessor/addRequest(_:processingOptions:)
func (v_ VideoProcessor) AddRequestProcessingOptionsError(request IVNRequest, processingOptions IVNVideoProcessorRequestProcessingOptions, error_ objectivec.IObject) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("addRequest:processingOptions:error:"), request, processingOptions, error_)
	return rv
}/* debug [instance_methods/method]: AddRequestProcessingOptionsError */


// Analyzes a time range of video content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNVideoProcessor/analyze(_:)
func (v_ VideoProcessor) AnalyzeTimeRangeError(timeRange TimeRange /* not a class type */, error_ objectivec.IObject) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("analyzeTimeRange:error:"), timeRange, error_)
	return rv
}/* debug [instance_methods/method]: AnalyzeTimeRangeError */


// Cancels the video processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNVideoProcessor/cancel()
func (v_ VideoProcessor) Cancel() {
	objc.Send[objc.ID](v_.ID, objc.Sel("cancel"))
}/* debug [instance_methods/method]: Cancel */


// Removes a Vision request from the video processor’s request queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNVideoProcessor/removeRequest(_:)
func (v_ VideoProcessor) RemoveRequestError(request IVNRequest, error_ objectivec.IObject) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("removeRequest:error:"), request, error_)
	return rv
}/* debug [instance_methods/method]: RemoveRequestError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VideoProcessor */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNVideoProcessor */


