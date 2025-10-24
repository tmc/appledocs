//go:build darwin && ios

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// iOS-only methods for CaptureSynchronizedDepthData


// iOS-only properties

// The depth data captured at this synchronization point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSynchronizedDepthData/depthData
func (c_ CaptureSynchronizedDepthData) DepthData() IAVDepthData {
	rv := objc.Send[DepthData](c_.ID, objc.Sel("depthData"))
	return rv
}

// A Boolean value indicating whether depth data was discarded between capture and processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSynchronizedDepthData/depthDataWasDropped
func (c_ CaptureSynchronizedDepthData) DepthDataWasDropped() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("depthDataWasDropped"))
	return rv
}

// A value indicating why the capture output failed to deliver depth data, if applicable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSynchronizedDepthData/droppedReason
func (c_ CaptureSynchronizedDepthData) DroppedReason() CaptureOutputDataDroppedReason {
	rv := objc.Send[CaptureOutputDataDroppedReason](c_.ID, objc.Sel("droppedReason"))
	return rv
}





