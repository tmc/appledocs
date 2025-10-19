// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVSampleBufferRenderSynchronizer] class.
var aVSampleBufferRenderSynchronizerClass = _AVSampleBufferRenderSynchronizerClass{objc.GetClass("AVSampleBufferRenderSynchronizer")}

type _AVSampleBufferRenderSynchronizerClass struct {
	class objc.Class
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


