//go:build darwin && ios

// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for SpeechSynthesizer


// iOS-only properties

// A Boolean value that specifies whether to send synthesized speech to an active call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/mixToTelephonyUplink
func (s_ SpeechSynthesizer) MixToTelephonyUplink() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("mixToTelephonyUplink"))
	return rv
}
func (s_ SpeechSynthesizer) SetMixToTelephonyUplink(value bool) {
	s_.ID.Send(objc.RegisterName("setMixToTelephonyUplink:"), value)
}

// An array of audio session channels to route generated speech.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/outputChannels
func (s_ SpeechSynthesizer) OutputChannels() []AudioSessionChannelDescription {
	rv := objc.Send[[]AudioSessionChannelDescription](s_.ID, objc.Sel("outputChannels"))
	return rv
}
func (s_ SpeechSynthesizer) SetOutputChannels(value []AudioSessionChannelDescription) {
	s_.ID.Send(objc.RegisterName("setOutputChannels:"), value)
}

// A Boolean value that specifies whether the app manages the audio session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/usesApplicationAudioSession
func (s_ SpeechSynthesizer) UsesApplicationAudioSession() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("usesApplicationAudioSession"))
	return rv
}
func (s_ SpeechSynthesizer) SetUsesApplicationAudioSession(value bool) {
	s_.ID.Send(objc.RegisterName("setUsesApplicationAudioSession:"), value)
}





