//go:build darwin && ios

// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for InputPickerInteraction


// Dismisses the input picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVInputPickerInteraction/dismiss()
func (i_ InputPickerInteraction) Dismiss() {
	objc.Send[objc.ID](i_.ID, objc.Sel("dismiss"))
}

// Presents the input picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVInputPickerInteraction/present()
func (i_ InputPickerInteraction) Present() {
	objc.Send[objc.ID](i_.ID, objc.Sel("present"))
}

// iOS-only properties

// The audio session for the picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVInputPickerInteraction/audioSession
func (i_ InputPickerInteraction) AudioSession() avfaudio.AudioSession {
	rv := objc.Send[avfaudio.AudioSession](i_.ID, objc.Sel("audioSession"))
	return rv
}
func (i_ InputPickerInteraction) SetAudioSession(value avfaudio.AudioSession) {
	i_.ID.Send(objc.RegisterName("setAudioSession:"), value)
}

// The input picker view’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVInputPickerInteraction/delegate-swift.property
func (i_ InputPickerInteraction) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("delegate"))
	return rv
}
func (i_ InputPickerInteraction) SetDelegate(value unsafe.Pointer) {
	i_.ID.Send(objc.RegisterName("setDelegate:"), value)
}

// A Boolean value that indicates whether the picker is currently visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVInputPickerInteraction/isPresented
func (i_ InputPickerInteraction) Presented() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("presented"))
	return rv
}




