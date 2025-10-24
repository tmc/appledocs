// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [VideoPerformanceMetrics] class.
var (
	VideoPerformanceMetricsClass     _VideoPerformanceMetricsClass
	VideoPerformanceMetricsClassOnce sync.Once
)

func getVideoPerformanceMetricsClass() _VideoPerformanceMetricsClass {
	VideoPerformanceMetricsClassOnce.Do(func() {
		VideoPerformanceMetricsClass = _VideoPerformanceMetricsClass{objc.GetClass("AVVideoPerformanceMetrics")}
	})
	return VideoPerformanceMetricsClass
}

type _VideoPerformanceMetricsClass struct {
	class objc.Class
}





// An interface definition for the [VideoPerformanceMetrics] class.
type IVideoPerformanceMetrics interface {
	objectivec.IObject
	

	// properties:
	NumberOfCorruptedVideoFrames() objectivec.IObject
	NumberOfDisplayCompositedVideoFrames() objectivec.IObject
	NumberOfDroppedVideoFrames() objectivec.IObject
	NumberOfNonDisplayCompositedVideoFrames() objectivec.IObject
	NumberOfCorruptedFrames() int
	NumberOfDroppedFrames() int
	NumberOfFramesDisplayedUsingOptimizedCompositing() int
	TotalFrameDelay() float64
	TotalNumberOfVideoFrames() objectivec.IObject
	TotalAccumulatedFrameDelay() float64
	TotalNumberOfFrames() int


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (vc _VideoPerformanceMetricsClass) Alloc() VideoPerformanceMetrics {
	rv := objc.Send[VideoPerformanceMetrics](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VideoPerformanceMetricsClass) New() VideoPerformanceMetrics {
	rv := objc.Send[VideoPerformanceMetrics](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VideoPerformanceMetrics) Init() VideoPerformanceMetrics {
	rv := objc.Send[VideoPerformanceMetrics](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VideoPerformanceMetrics) Autorelease() VideoPerformanceMetrics {
	rv := objc.Send[VideoPerformanceMetrics](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVideoPerformanceMetrics creates a new VideoPerformanceMetrics instance.
func NewVideoPerformanceMetrics() VideoPerformanceMetrics {
	return getVideoPerformanceMetricsClass().New()
}





// An object that provides metrics related to video playback quality.


// An object that provides metrics related to video playback quality.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoPerformanceMetrics
type VideoPerformanceMetrics struct {
	objectivec.Object
}

// VideoPerformanceMetricsFrom constructs a [VideoPerformanceMetrics] from an unsafe.Pointer.
//
// An object that provides metrics related to video playback quality.
func VideoPerformanceMetricsFrom(ptr unsafe.Pointer) VideoPerformanceMetrics {
	return VideoPerformanceMetrics{objectivec.Object{objc.ID(ptr)}}
}

























// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoPerformanceMetrics/numberOfCorruptedVideoFrames
func (v_ VideoPerformanceMetrics) NumberOfCorruptedVideoFrames() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](v_.ID, objc.Sel("numberOfCorruptedVideoFrames"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoPerformanceMetrics/numberOfDisplayCompositedVideoFrames
func (v_ VideoPerformanceMetrics) NumberOfDisplayCompositedVideoFrames() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](v_.ID, objc.Sel("numberOfDisplayCompositedVideoFrames"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoPerformanceMetrics/numberOfDroppedVideoFrames
func (v_ VideoPerformanceMetrics) NumberOfDroppedVideoFrames() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](v_.ID, objc.Sel("numberOfDroppedVideoFrames"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoPerformanceMetrics/numberOfNonDisplayCompositedVideoFrames
func (v_ VideoPerformanceMetrics) NumberOfNonDisplayCompositedVideoFrames() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](v_.ID, objc.Sel("numberOfNonDisplayCompositedVideoFrames"))
	return rv
}


// The total number of corrupted frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoPerformanceMetrics/numberOfCorruptedFrames
func (v_ VideoPerformanceMetrics) NumberOfCorruptedFrames() int {
	rv := objc.Send[int](v_.ID, objc.Sel("numberOfCorruptedFrames"))
	return rv
}


// The total number of frames the system drops prior to decoding or from missing the display deadline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoPerformanceMetrics/numberOfDroppedFrames
func (v_ VideoPerformanceMetrics) NumberOfDroppedFrames() int {
	rv := objc.Send[int](v_.ID, objc.Sel("numberOfDroppedFrames"))
	return rv
}


// The total number of full screen frames rendered in a special power-efficient mode that didn’t require compositing with other UI elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoPerformanceMetrics/numberOfFramesDisplayedUsingOptimizedCompositing
func (v_ VideoPerformanceMetrics) NumberOfFramesDisplayedUsingOptimizedCompositing() int {
	rv := objc.Send[int](v_.ID, objc.Sel("numberOfFramesDisplayedUsingOptimizedCompositing"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoPerformanceMetrics/totalFrameDelay
func (v_ VideoPerformanceMetrics) TotalFrameDelay() float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("totalFrameDelay"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoPerformanceMetrics/totalNumberOfVideoFrames
func (v_ VideoPerformanceMetrics) TotalNumberOfVideoFrames() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](v_.ID, objc.Sel("totalNumberOfVideoFrames"))
	return rv
}


// The accumulated amount of time between the prescribed presentation times of displayed video frames and their actual time of display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoPerformanceMetrics/totalAccumulatedFrameDelay
func (v_ VideoPerformanceMetrics) TotalAccumulatedFrameDelay() float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("totalAccumulatedFrameDelay"))
	return rv
}


// The total number of frames that display if no frames drop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoPerformanceMetrics/totalNumberOfFrames
func (v_ VideoPerformanceMetrics) TotalNumberOfFrames() int {
	rv := objc.Send[int](v_.ID, objc.Sel("totalNumberOfFrames"))
	return rv
}










