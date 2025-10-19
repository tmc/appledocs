// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVSampleBufferRenderSynchronizer] class.
var (
	aVSampleBufferRenderSynchronizerClass     _AVSampleBufferRenderSynchronizerClass
	aVSampleBufferRenderSynchronizerClassOnce sync.Once
)

func getAVSampleBufferRenderSynchronizerClass() _AVSampleBufferRenderSynchronizerClass {
	aVSampleBufferRenderSynchronizerClassOnce.Do(func() {
		aVSampleBufferRenderSynchronizerClass = _AVSampleBufferRenderSynchronizerClass{objc.GetClass("AVSampleBufferRenderSynchronizer")}
	})
	return aVSampleBufferRenderSynchronizerClass
}

type _AVSampleBufferRenderSynchronizerClass struct {
	class objc.Class
}

// An interface definition for the [AVSampleBufferRenderSynchronizer] class.
type IAVSampleBufferRenderSynchronizer interface {
	objectivec.IObject
	AddBoundaryTimeObserverForTimesQueueUsingBlock(times unsafe.Pointer, queue unsafe.Pointer, block unsafe.Pointer) objc.ID
	SetRateTime(rate float32, time unsafe.Pointer)
}

// An object used to synchronize multiple queued sample buffers to a single timeline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer
type AVSampleBufferRenderSynchronizer struct {
	objectivec.Object
}

// AVSampleBufferRenderSynchronizerFrom constructs a [AVSampleBufferRenderSynchronizer] from an unsafe.Pointer.
//
// An object used to synchronize multiple queued sample buffers to a single timeline.
func AVSampleBufferRenderSynchronizerFrom(ptr unsafe.Pointer) AVSampleBufferRenderSynchronizer {
	return AVSampleBufferRenderSynchronizer{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AVSampleBufferRenderSynchronizerClass) Alloc() AVSampleBufferRenderSynchronizer {
	rv := objc.Send[AVSampleBufferRenderSynchronizer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AVSampleBufferRenderSynchronizerClass) New() AVSampleBufferRenderSynchronizer {
	rv := objc.Send[AVSampleBufferRenderSynchronizer](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVSampleBufferRenderSynchronizer) Init() AVSampleBufferRenderSynchronizer {
	rv := objc.Send[AVSampleBufferRenderSynchronizer](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVSampleBufferRenderSynchronizer) Autorelease() AVSampleBufferRenderSynchronizer {
	rv := objc.Send[AVSampleBufferRenderSynchronizer](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVSampleBufferRenderSynchronizer creates a new AVSampleBufferRenderSynchronizer instance.
func NewAVSampleBufferRenderSynchronizer() AVSampleBufferRenderSynchronizer {
	return getAVSampleBufferRenderSynchronizerClass().New()
}


// Requests invocation of a block when specified times are traversed during normal rendering. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/addBoundaryTimeObserver(forTimes:queue:using:)
func (a_ AVSampleBufferRenderSynchronizer) AddBoundaryTimeObserverForTimesQueueUsingBlock(times unsafe.Pointer, queue unsafe.Pointer, block unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("addBoundaryTimeObserverForTimes:queue:usingBlock:"), times, queue, block)
	return rv
}
// Sets the renderer’s time and rate. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/setRate(_:time:)
func (a_ AVSampleBufferRenderSynchronizer) SetRateTime(rate float32, time unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRate:time:"), rate, time)
}


