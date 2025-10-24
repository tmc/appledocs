//go:build darwin && ios

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// iOS-only methods for CaptureSynchronizedSampleBufferData


// iOS-only properties

// A value indicating why the capture output failed to deliver sample buffers, if applicable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSynchronizedSampleBufferData/droppedReason
func (c_ CaptureSynchronizedSampleBufferData) DroppedReason() CaptureOutputDataDroppedReason {
	rv := objc.Send[CaptureOutputDataDroppedReason](c_.ID, objc.Sel("droppedReason"))
	return rv
}

// The depth data captured at this synchronization point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSynchronizedSampleBufferData/sampleBuffer
func (c_ CaptureSynchronizedSampleBufferData) SampleBuffer() SampleBufferRef /* not a class type */ {
	rv := objc.Send[SampleBufferRef](c_.ID, objc.Sel("sampleBuffer"))
	return rv
}

// A Boolean value indicating whether sample buffers were discarded between capture and processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSynchronizedSampleBufferData/sampleBufferWasDropped
func (c_ CaptureSynchronizedSampleBufferData) SampleBufferWasDropped() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("sampleBufferWasDropped"))
	return rv
}





