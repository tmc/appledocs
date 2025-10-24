//go:build darwin && ios

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CaptureSynchronizedDataCollection


// Returns data captured by the specified capture output, using subscript syntax.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSynchronizedDataCollection/subscript(_:)
func (c_ CaptureSynchronizedDataCollection) ObjectForKeyedSubscript(key IAVCaptureOutput) ICaptureSynchronizedData {
	rv := objc.Send[CaptureSynchronizedData](c_.ID, objc.Sel("objectForKeyedSubscript:"), key)
	return rv
}

// Returns synchronized data captured by the specified capture output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSynchronizedDataCollection/synchronizedData(for:)
func (c_ CaptureSynchronizedDataCollection) SynchronizedDataForCaptureOutput(captureOutput IAVCaptureOutput) ICaptureSynchronizedData {
	rv := objc.Send[CaptureSynchronizedData](c_.ID, objc.Sel("synchronizedDataForCaptureOutput:"), captureOutput)
	return rv
}

// iOS-only properties

// The number of synchronized data objects in the collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSynchronizedDataCollection/count
func (c_ CaptureSynchronizedDataCollection) Count() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("count"))
	return rv
}





