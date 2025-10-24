//go:build darwin && ios

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CaptureDataOutputSynchronizer


// Designates a delegate object to receive synchronized data and a dispatch queue for delivering that data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDataOutputSynchronizer/setDelegate(_:queue:)
func (c_ CaptureDataOutputSynchronizer) SetDelegateQueue(delegate unsafe.Pointer, delegateCallbackQueue objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegate:queue:"), delegate, delegateCallbackQueue)
}

// iOS-only properties

// The list of data outputs governed by this data output synchronizer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDataOutputSynchronizer/dataOutputs
func (c_ CaptureDataOutputSynchronizer) DataOutputs() []CaptureOutput {
	rv := objc.Send[[]CaptureOutput](c_.ID, objc.Sel("dataOutputs"))
	return rv
}

// A delegate object that receives synchronized capture data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDataOutputSynchronizer/delegate
func (c_ CaptureDataOutputSynchronizer) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("delegate"))
	return rv
}

// A dispatch queue for delivering synchronized capture data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDataOutputSynchronizer/delegateCallbackQueue
func (c_ CaptureDataOutputSynchronizer) DelegateCallbackQueue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("delegateCallbackQueue"))
	return rv
}




