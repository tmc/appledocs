//go:build darwin && ios

// Code generated from Apple documentation for ReplayKit. DO NOT EDIT.

package replaykit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for RPSystemBroadcastPickerView


// iOS-only properties

// A bundle identifier of a broadcast extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPSystemBroadcastPickerView/preferredExtension
func (r_ RPSystemBroadcastPickerView) PreferredExtension() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](r_.ID, objc.Sel("preferredExtension"))
	return rv
}
func (r_ RPSystemBroadcastPickerView) SetPreferredExtension(value objc.IObject /* cross-framework: NSString */) {
	r_.ID.Send(objc.RegisterName("setPreferredExtension:"), value)
}

// A Boolean value that indicates whether the microphone button is visible in the broadcast picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPSystemBroadcastPickerView/showsMicrophoneButton
func (r_ RPSystemBroadcastPickerView) ShowsMicrophoneButton() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("showsMicrophoneButton"))
	return rv
}
func (r_ RPSystemBroadcastPickerView) SetShowsMicrophoneButton(value bool) {
	r_.ID.Send(objc.RegisterName("setShowsMicrophoneButton:"), value)
}






