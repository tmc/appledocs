// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [SampleBufferRenderSynchronizer] class.
var (
	SampleBufferRenderSynchronizerClass     _SampleBufferRenderSynchronizerClass
	SampleBufferRenderSynchronizerClassOnce sync.Once
)

func getSampleBufferRenderSynchronizerClass() _SampleBufferRenderSynchronizerClass {
	SampleBufferRenderSynchronizerClassOnce.Do(func() {
		SampleBufferRenderSynchronizerClass = _SampleBufferRenderSynchronizerClass{objc.GetClass("AVSampleBufferRenderSynchronizer")}
	})
	return SampleBufferRenderSynchronizerClass
}

type _SampleBufferRenderSynchronizerClass struct {
	class objc.Class
}





// An interface definition for the [SampleBufferRenderSynchronizer] class.
type ISampleBufferRenderSynchronizer interface {
	objectivec.IObject
	

	// properties:
	DelaysRateChangeUntilHasSufficientMediaData() bool
	SetDelaysRateChangeUntilHasSufficientMediaData(value bool)
	Rate() float32
	SetRate(value float32)
	Renderers() []objc.ID
	Timebase() TimebaseRef /* not a class type */


	

	// methods:
	AddBoundaryTimeObserverForTimesQueueUsingBlock(times []foundation.Value, queue objectivec.IObject, block unsafe.Pointer) objc.ID
	AddPeriodicTimeObserverForIntervalQueueUsingBlock(interval objectivec.IObject, queue objectivec.IObject, block unsafe.Pointer) objc.ID
	AddRenderer(renderer unsafe.Pointer)
	CurrentTime() objectivec.IObject
	RemoveRendererAtTimeCompletionHandler(renderer unsafe.Pointer, time objectivec.IObject, completionHandler unsafe.Pointer)
	RemoveTimeObserver(observer objectivec.IObject)
	SetRateTime(rate float32, time objectivec.IObject)
	SetRateTimeAtHostTime(rate float32, time objectivec.IObject, hostTime objectivec.IObject)


}





// Alloc allocates a new instance without initialization.
func (sc _SampleBufferRenderSynchronizerClass) Alloc() SampleBufferRenderSynchronizer {
	rv := objc.Send[SampleBufferRenderSynchronizer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SampleBufferRenderSynchronizerClass) New() SampleBufferRenderSynchronizer {
	rv := objc.Send[SampleBufferRenderSynchronizer](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SampleBufferRenderSynchronizer) Init() SampleBufferRenderSynchronizer {
	rv := objc.Send[SampleBufferRenderSynchronizer](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SampleBufferRenderSynchronizer) Autorelease() SampleBufferRenderSynchronizer {
	rv := objc.Send[SampleBufferRenderSynchronizer](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSampleBufferRenderSynchronizer creates a new SampleBufferRenderSynchronizer instance.
func NewSampleBufferRenderSynchronizer() SampleBufferRenderSynchronizer {
	return getSampleBufferRenderSynchronizerClass().New()
}





// An object used to synchronize multiple queued sample buffers to a single timeline.
//
// This class synchronizes multiple objects that conform to to a single timeline.


// An object used to synchronize multiple queued sample buffers to a single timeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer
type SampleBufferRenderSynchronizer struct {
	objectivec.Object
}

// SampleBufferRenderSynchronizerFrom constructs a [SampleBufferRenderSynchronizer] from an unsafe.Pointer.
//
// An object used to synchronize multiple queued sample buffers to a single timeline.
func SampleBufferRenderSynchronizerFrom(ptr unsafe.Pointer) SampleBufferRenderSynchronizer {
	return SampleBufferRenderSynchronizer{objectivec.Object{objc.ID(ptr)}}
}




















// Requests invocation of a block when specified times are traversed during normal rendering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/addBoundaryTimeObserver(forTimes:queue:using:)
func (s_ SampleBufferRenderSynchronizer) AddBoundaryTimeObserverForTimesQueueUsingBlock(times []foundation.Value, queue objectivec.IObject, block unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("addBoundaryTimeObserverForTimes:queue:usingBlock:"), times, queue, block)
	return rv
}


// Requests invocation of a block during rendering at specified time intervals.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/addPeriodicTimeObserver(forInterval:queue:using:)
func (s_ SampleBufferRenderSynchronizer) AddPeriodicTimeObserverForIntervalQueueUsingBlock(interval objectivec.IObject, queue objectivec.IObject, block unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("addPeriodicTimeObserverForInterval:queue:usingBlock:"), interval, queue, block)
	return rv
}


// Adds a renderer to the list of renderers under the synchronizer’s control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/addRenderer(_:)
func (s_ SampleBufferRenderSynchronizer) AddRenderer(renderer unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("addRenderer:"), renderer)
}


// Returns the current time of the synchronizer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/currentTime()
func (s_ SampleBufferRenderSynchronizer) CurrentTime() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("currentTime"))
	return rv
}


// Removes a renderer from the synchronizer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/removeRenderer(_:at:completionHandler:)
func (s_ SampleBufferRenderSynchronizer) RemoveRendererAtTimeCompletionHandler(renderer unsafe.Pointer, time objectivec.IObject, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeRenderer:atTime:completionHandler:"), renderer, time, completionHandler)
}


// Cancels the specified time observer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/removeTimeObserver(_:)
func (s_ SampleBufferRenderSynchronizer) RemoveTimeObserver(observer objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeTimeObserver:"), observer)
}


// Sets the renderer’s time and rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/setRate(_:time:)
func (s_ SampleBufferRenderSynchronizer) SetRateTime(rate float32, time objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setRate:time:"), rate, time)
}


// Sets the playback rate and the relationship between the current time and host time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/setRate(_:time:atHostTime:)
func (s_ SampleBufferRenderSynchronizer) SetRateTimeAtHostTime(rate float32, time objectivec.IObject, hostTime objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setRate:time:atHostTime:"), rate, time, hostTime)
}







// A Boolean value that Indicates whether the playback should start immediately on rate change requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/delaysRateChangeUntilHasSufficientMediaData
func (s_ SampleBufferRenderSynchronizer) DelaysRateChangeUntilHasSufficientMediaData() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("delaysRateChangeUntilHasSufficientMediaData"))
	return rv
}


// A Boolean value that Indicates whether the playback should start immediately on rate change requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/delaysRateChangeUntilHasSufficientMediaData
func (s_ SampleBufferRenderSynchronizer) SetDelaysRateChangeUntilHasSufficientMediaData(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelaysRateChangeUntilHasSufficientMediaData:"), value)
}


// The current playback rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/rate
func (s_ SampleBufferRenderSynchronizer) Rate() float32 {
	rv := objc.Send[float32](s_.ID, objc.Sel("rate"))
	return rv
}


// The current playback rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/rate
func (s_ SampleBufferRenderSynchronizer) SetRate(value float32) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setRate:"), value)
}


// An array of queued sample buffer renderers currently attached to the synchronizer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/renderers
func (s_ SampleBufferRenderSynchronizer) Renderers() []objc.ID {
	rv := objc.Send[[]objc.ID](s_.ID, objc.Sel("renderers"))
	return rv
}


// The synchronizer’s rendering timebase which determines how it interprets timestamps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/timebase
func (s_ SampleBufferRenderSynchronizer) Timebase() TimebaseRef /* not a class type */ {
	rv := objc.Send[TimebaseRef](s_.ID, objc.Sel("timebase"))
	return rv
}







