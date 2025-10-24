//go:build darwin && ios

// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for AudioApplication


// iOS-only properties

// A value that indicates an app’s permission to add audio to calls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/microphoneInjectionPermission-swift.property
func (a_ AudioApplication) MicrophoneInjectionPermission() AudioApplicationMicrophoneInjectionPermission {
	rv := objc.Send[AudioApplicationMicrophoneInjectionPermission](a_.ID, objc.Sel("microphoneInjectionPermission"))
	return rv
}





