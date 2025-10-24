//go:build darwin && ios

// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for AudioSessionRouteDescription


// iOS-only properties

// An array of audio input port descriptions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionRouteDescription/inputs
func (a_ AudioSessionRouteDescription) Inputs() []AudioSessionPortDescription {
	rv := objc.Send[[]AudioSessionPortDescription](a_.ID, objc.Sel("inputs"))
	return rv
}

// An array of audio output port descriptions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionRouteDescription/outputs
func (a_ AudioSessionRouteDescription) Outputs() []AudioSessionPortDescription {
	rv := objc.Send[[]AudioSessionPortDescription](a_.ID, objc.Sel("outputs"))
	return rv
}





