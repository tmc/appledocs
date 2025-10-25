// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVSampleBufferRenderSynchronizer */


/* debug [class_header]: Header for AVSampleBufferRenderSynchronizer */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SampleBufferRenderSynchronizer */
// An interface definition for the [SampleBufferRenderSynchronizer] class.
type ISampleBufferRenderSynchronizer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SampleBufferRenderSynchronizer */
	// properties:
	DelaysRateChangeUntilHasSufficientMediaData() bool
	SetDelaysRateChangeUntilHasSufficientMediaData(value bool)
	Rate() float32
	SetRate(value float32)
	Renderers() []objc.ID
	Timebase() TimebaseRef /* not a class type */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SampleBufferRenderSynchronizer */
	// methods:
	AddBoundaryTimeObserverForTimesQueueUsingBlock(times []foundation.Value, queue objectivec.IObject, block unsafe.Pointer) objc.ID
	AddPeriodicTimeObserverForIntervalQueueUsingBlock(interval objc.IObject /* cross-framework: Time */, queue objectivec.IObject, block unsafe.Pointer) objc.ID
	AddRenderer(renderer unsafe.Pointer)
	CurrentTime() objc.IObject /* cross-framework: Time */
	RemoveRendererAtTimeCompletionHandler(renderer unsafe.Pointer, time objc.IObject /* cross-framework: Time */, completionHandler unsafe.Pointer)
	RemoveTimeObserver(observer objc.IObject)
	SetRateTime(rate float32, time objc.IObject /* cross-framework: Time */)
	SetRateTimeAtHostTime(rate float32, time objc.IObject /* cross-framework: Time */, hostTime objc.IObject /* cross-framework: Time */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SampleBufferRenderSynchronizer */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SampleBufferRenderSynchronizer */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SampleBufferRenderSynchronizer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SampleBufferRenderSynchronizer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SampleBufferRenderSynchronizer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SampleBufferRenderSynchronizer */

// Requests invocation of a block when specified times are traversed during normal rendering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/addBoundaryTimeObserver(forTimes:queue:using:)
func (s_ SampleBufferRenderSynchronizer) AddBoundaryTimeObserverForTimesQueueUsingBlock(times []foundation.Value, queue objectivec.IObject, block unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("addBoundaryTimeObserverForTimes:queue:usingBlock:"), times, queue, block)
	return rv
}/* debug [instance_methods/method]: AddBoundaryTimeObserverForTimesQueueUsingBlock */


// Requests invocation of a block during rendering at specified time intervals.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/addPeriodicTimeObserver(forInterval:queue:using:)
func (s_ SampleBufferRenderSynchronizer) AddPeriodicTimeObserverForIntervalQueueUsingBlock(interval objc.IObject /* cross-framework: Time */, queue objectivec.IObject, block unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("addPeriodicTimeObserverForInterval:queue:usingBlock:"), interval, queue, block)
	return rv
}/* debug [instance_methods/method]: AddPeriodicTimeObserverForIntervalQueueUsingBlock */


// Adds a renderer to the list of renderers under the synchronizer’s control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/addRenderer(_:)
func (s_ SampleBufferRenderSynchronizer) AddRenderer(renderer unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("addRenderer:"), renderer)
}/* debug [instance_methods/method]: AddRenderer */


// Returns the current time of the synchronizer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/currentTime()
func (s_ SampleBufferRenderSynchronizer) CurrentTime() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](s_.ID, objc.Sel("currentTime"))
	return rv
}/* debug [instance_methods/method]: CurrentTime */


// Removes a renderer from the synchronizer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/removeRenderer(_:at:completionHandler:)
func (s_ SampleBufferRenderSynchronizer) RemoveRendererAtTimeCompletionHandler(renderer unsafe.Pointer, time objc.IObject /* cross-framework: Time */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeRenderer:atTime:completionHandler:"), renderer, time, completionHandler)
}/* debug [instance_methods/method]: RemoveRendererAtTimeCompletionHandler */


// Cancels the specified time observer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/removeTimeObserver(_:)
func (s_ SampleBufferRenderSynchronizer) RemoveTimeObserver(observer objc.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeTimeObserver:"), observer)
}/* debug [instance_methods/method]: RemoveTimeObserver */


// Sets the renderer’s time and rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/setRate(_:time:)
func (s_ SampleBufferRenderSynchronizer) SetRateTime(rate float32, time objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setRate:time:"), rate, time)
}/* debug [instance_methods/method]: SetRateTime */


// Sets the playback rate and the relationship between the current time and host time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/setRate(_:time:atHostTime:)
func (s_ SampleBufferRenderSynchronizer) SetRateTimeAtHostTime(rate float32, time objc.IObject /* cross-framework: Time */, hostTime objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setRate:time:atHostTime:"), rate, time, hostTime)
}/* debug [instance_methods/method]: SetRateTimeAtHostTime */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SampleBufferRenderSynchronizer */

// A Boolean value that Indicates whether the playback should start immediately on rate change requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/delaysRateChangeUntilHasSufficientMediaData
func (s_ SampleBufferRenderSynchronizer) DelaysRateChangeUntilHasSufficientMediaData() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("delaysRateChangeUntilHasSufficientMediaData"))
	return rv
}/* debug [instance_properties/getter]: delaysRateChangeUntilHasSufficientMediaData */


// A Boolean value that Indicates whether the playback should start immediately on rate change requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/delaysRateChangeUntilHasSufficientMediaData
func (s_ SampleBufferRenderSynchronizer) SetDelaysRateChangeUntilHasSufficientMediaData(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelaysRateChangeUntilHasSufficientMediaData:"), value)
}/* debug [instance_properties/setter]: delaysRateChangeUntilHasSufficientMediaData */


// The current playback rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/rate
func (s_ SampleBufferRenderSynchronizer) Rate() float32 {
	rv := objc.Send[float32](s_.ID, objc.Sel("rate"))
	return rv
}/* debug [instance_properties/getter]: rate */


// The current playback rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/rate
func (s_ SampleBufferRenderSynchronizer) SetRate(value float32) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setRate:"), value)
}/* debug [instance_properties/setter]: rate */


// An array of queued sample buffer renderers currently attached to the synchronizer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/renderers
func (s_ SampleBufferRenderSynchronizer) Renderers() []objc.ID {
	rv := objc.Send[[]objc.ID](s_.ID, objc.Sel("renderers"))
	return rv
}/* debug [instance_properties/getter]: renderers */


// The synchronizer’s rendering timebase which determines how it interprets timestamps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/timebase
func (s_ SampleBufferRenderSynchronizer) Timebase() TimebaseRef /* not a class type */ {
	rv := objc.Send[TimebaseRef](s_.ID, objc.Sel("timebase"))
	return rv
}/* debug [instance_properties/getter]: timebase */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVSampleBufferRenderSynchronizer */


