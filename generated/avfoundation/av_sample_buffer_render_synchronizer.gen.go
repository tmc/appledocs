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
	AddBoundaryTimeObserverForTimesQueueUsingBlock(times unsafe.Pointer, queue unsafe.Pointer, block unsafe.Pointer) objc.ID
	SetRateTime(rate unsafe.Pointer, time unsafe.Pointer)
}

// An object used to synchronize multiple queued sample buffers to a single timeline.
//
// This class synchronizes multiple objects that conform to to a single timeline.
//
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
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/addBoundaryTimeObserver(forTimes:queue:using:)
func (s_ SampleBufferRenderSynchronizer) AddBoundaryTimeObserverForTimesQueueUsingBlock(times unsafe.Pointer, queue unsafe.Pointer, block unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("addBoundaryTimeObserverForTimes:queue:usingBlock:"), times, queue, block)
	return rv
}

// Sets the renderer’s time and rate.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/setRate(_:time:)
func (s_ SampleBufferRenderSynchronizer) SetRateTime(rate unsafe.Pointer, time unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setRate:time:"), rate, time)
}



