//go:build darwin && ios

// Code generated from Apple documentation for CoreAudioKit. DO NOT EDIT.

package coreaudiokit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/audiotoolbox"
)

// iOS-only methods for InterAppAudioTransportView


// iOS-only properties

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/CAInterAppAudioTransportView/currentTimeLabelFont
func (i_ InterAppAudioTransportView) CurrentTimeLabelFont() appkit.Font {
	rv := objc.Send[appkit.Font](i_.ID, objc.Sel("currentTimeLabelFont"))
	return rv
}
func (i_ InterAppAudioTransportView) SetCurrentTimeLabelFont(value appkit.Font) {
	i_.ID.Send(objc.RegisterName("setCurrentTimeLabelFont:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/CAInterAppAudioTransportView/isConnected
func (i_ InterAppAudioTransportView) Connected() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("connected"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/CAInterAppAudioTransportView/isEnabled
func (i_ InterAppAudioTransportView) Enabled() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("enabled"))
	return rv
}
func (i_ InterAppAudioTransportView) SetEnabled(value bool) {
	i_.ID.Send(objc.RegisterName("setEnabled:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/CAInterAppAudioTransportView/isPlaying
func (i_ InterAppAudioTransportView) Playing() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("playing"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/CAInterAppAudioTransportView/isRecording
func (i_ InterAppAudioTransportView) Recording() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("recording"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/CAInterAppAudioTransportView/labelColor
func (i_ InterAppAudioTransportView) LabelColor() appkit.Color {
	rv := objc.Send[appkit.Color](i_.ID, objc.Sel("labelColor"))
	return rv
}
func (i_ InterAppAudioTransportView) SetLabelColor(value appkit.Color) {
	i_.ID.Send(objc.RegisterName("setLabelColor:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/CAInterAppAudioTransportView/pauseButtonColor
func (i_ InterAppAudioTransportView) PauseButtonColor() appkit.Color {
	rv := objc.Send[appkit.Color](i_.ID, objc.Sel("pauseButtonColor"))
	return rv
}
func (i_ InterAppAudioTransportView) SetPauseButtonColor(value appkit.Color) {
	i_.ID.Send(objc.RegisterName("setPauseButtonColor:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/CAInterAppAudioTransportView/playButtonColor
func (i_ InterAppAudioTransportView) PlayButtonColor() appkit.Color {
	rv := objc.Send[appkit.Color](i_.ID, objc.Sel("playButtonColor"))
	return rv
}
func (i_ InterAppAudioTransportView) SetPlayButtonColor(value appkit.Color) {
	i_.ID.Send(objc.RegisterName("setPlayButtonColor:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/CAInterAppAudioTransportView/recordButtonColor
func (i_ InterAppAudioTransportView) RecordButtonColor() appkit.Color {
	rv := objc.Send[appkit.Color](i_.ID, objc.Sel("recordButtonColor"))
	return rv
}
func (i_ InterAppAudioTransportView) SetRecordButtonColor(value appkit.Color) {
	i_.ID.Send(objc.RegisterName("setRecordButtonColor:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/CAInterAppAudioTransportView/rewindButtonColor
func (i_ InterAppAudioTransportView) RewindButtonColor() appkit.Color {
	rv := objc.Send[appkit.Color](i_.ID, objc.Sel("rewindButtonColor"))
	return rv
}
func (i_ InterAppAudioTransportView) SetRewindButtonColor(value appkit.Color) {
	i_.ID.Send(objc.RegisterName("setRewindButtonColor:"), value)
}





