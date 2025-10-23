// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/audiotoolbox"
	"github.com/tmc/appledocs/generated/foundation"
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
	DelaysRateChangeUntilHasSufficientMediaData() bool
	SetDelaysRateChangeUntilHasSufficientMediaData(value bool)
	IntendedSpatialAudioExperience() audiotoolbox.SpatialAudioExperience
	SetIntendedSpatialAudioExperience(value audiotoolbox.SpatialAudioExperience)
	Rate() float32
	SetRate(value float32)
	Renderers() unsafe.Pointer
	SetRenderers(value unsafe.Pointer)
	Timebase() unsafe.Pointer
	SetTimebase(value unsafe.Pointer)
	AddBoundaryTimeObserverForTimesQueueUsingBlock(times []foundation.Value, queue unsafe.Pointer, block unsafe.Pointer) objc.ID
	SetRateTime(rate float32, time unsafe.Pointer)
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

// Alloc allocates a new instance without initialization.
func (sc _SampleBufferRenderSynchronizerClass) Alloc() SampleBufferRenderSynchronizer {
	rv := objc.Send[SampleBufferRenderSynchronizer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Requests invocation of a block when specified times are traversed during normal rendering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/addBoundaryTimeObserver(forTimes:queue:using:)
func (s_ SampleBufferRenderSynchronizer) AddBoundaryTimeObserverForTimesQueueUsingBlock(times []foundation.Value, queue unsafe.Pointer, block unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("addBoundaryTimeObserverForTimes:queue:usingBlock:"), times, queue, block)
	return rv
}


// Sets the renderer’s time and rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/setRate(_:time:)
func (s_ SampleBufferRenderSynchronizer) SetRateTime(rate float32, time unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setRate:time:"), rate, time)
}


// A Boolean value that Indicates whether the playback should start immediately on rate change requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrendersynchronizer/delaysratechangeuntilhassufficientmediadata
func (s_ SampleBufferRenderSynchronizer) DelaysRateChangeUntilHasSufficientMediaData() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("delaysRateChangeUntilHasSufficientMediaData"))
	return rv
}


// A Boolean value that Indicates whether the playback should start immediately on rate change requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrendersynchronizer/delaysratechangeuntilhassufficientmediadata
func (s_ SampleBufferRenderSynchronizer) SetDelaysRateChangeUntilHasSufficientMediaData(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelaysRateChangeUntilHasSufficientMediaData:"), value)
}


// The synchronizer’s intended Spatial Audio experience.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrendersynchronizer/intendedspatialaudioexperience-3z7d3
func (s_ SampleBufferRenderSynchronizer) IntendedSpatialAudioExperience() audiotoolbox.SpatialAudioExperience {
	rv := objc.Send[audiotoolbox.SpatialAudioExperience](s_.ID, objc.Sel("intendedSpatialAudioExperience"))
	return rv
}


// The synchronizer’s intended Spatial Audio experience.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrendersynchronizer/intendedspatialaudioexperience-3z7d3
func (s_ SampleBufferRenderSynchronizer) SetIntendedSpatialAudioExperience(value audiotoolbox.SpatialAudioExperience) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIntendedSpatialAudioExperience:"), value)
}


// The current playback rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrendersynchronizer/rate
func (s_ SampleBufferRenderSynchronizer) Rate() float32 {
	rv := objc.Send[float32](s_.ID, objc.Sel("rate"))
	return rv
}


// The current playback rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrendersynchronizer/rate
func (s_ SampleBufferRenderSynchronizer) SetRate(value float32) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setRate:"), value)
}


// An array of queued sample buffer renderers currently attached to the synchronizer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrendersynchronizer/renderers
func (s_ SampleBufferRenderSynchronizer) Renderers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("renderers"))
	return rv
}


// An array of queued sample buffer renderers currently attached to the synchronizer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrendersynchronizer/renderers
func (s_ SampleBufferRenderSynchronizer) SetRenderers(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setRenderers:"), value)
}


// The synchronizer’s rendering timebase which determines how it interprets timestamps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrendersynchronizer/timebase
func (s_ SampleBufferRenderSynchronizer) Timebase() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("timebase"))
	return rv
}


// The synchronizer’s rendering timebase which determines how it interprets timestamps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrendersynchronizer/timebase
func (s_ SampleBufferRenderSynchronizer) SetTimebase(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTimebase:"), value)
}



