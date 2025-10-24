//go:build darwin && ios

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CaptureDepthDataOutput


// Designates a delegate object to receive depth data and a dispatch queue for delivering that data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDepthDataOutput/setDelegate(_:callbackQueue:)
func (c_ CaptureDepthDataOutput) SetDelegateCallbackQueue(delegate unsafe.Pointer, callbackQueue objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegate:callbackQueue:"), delegate, callbackQueue)
}

// iOS-only properties

// A Boolean value that determines whether the capture output should discard any depth data that is not processed before the next depth data is captured.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDepthDataOutput/alwaysDiscardsLateDepthData
func (c_ CaptureDepthDataOutput) AlwaysDiscardsLateDepthData() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("alwaysDiscardsLateDepthData"))
	return rv
}
func (c_ CaptureDepthDataOutput) SetAlwaysDiscardsLateDepthData(value bool) {
	c_.ID.Send(objc.RegisterName("setAlwaysDiscardsLateDepthData:"), value)
}

// A delegate object that receives depth data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDepthDataOutput/delegate
func (c_ CaptureDepthDataOutput) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("delegate"))
	return rv
}

// A dispatch queue for delivering depth data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDepthDataOutput/delegateCallbackQueue
func (c_ CaptureDepthDataOutput) DelegateCallbackQueue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("delegateCallbackQueue"))
	return rv
}

// A Boolean value that determines whether the depth data output should filter depth data to smooth out noise and fill invalid values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDepthDataOutput/isFilteringEnabled
func (c_ CaptureDepthDataOutput) FilteringEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("filteringEnabled"))
	return rv
}
func (c_ CaptureDepthDataOutput) SetFilteringEnabled(value bool) {
	c_.ID.Send(objc.RegisterName("setFilteringEnabled:"), value)
}




