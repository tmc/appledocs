//go:build darwin && ios

// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CaptureEvent


// Plays the specified capture sound through AirPods.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureEvent/play(_:)
func (c_ CaptureEvent) PlaySound(sound IAVCaptureEventSound) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("playSound:"), sound)
	return rv
}

// iOS-only properties

// The current phase of a capture event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureEvent/phase
func (c_ CaptureEvent) Phase() CaptureEventPhase {
	rv := objc.Send[CaptureEventPhase](c_.ID, objc.Sel("phase"))
	return rv
}

// A Boolean value that indicates whether you must play a sound manually.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureEvent/shouldPlaySound
func (c_ CaptureEvent) ShouldPlaySound() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("shouldPlaySound"))
	return rv
}





