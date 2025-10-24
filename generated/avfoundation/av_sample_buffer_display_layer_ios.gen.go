//go:build darwin && ios

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for SampleBufferDisplayLayer


// iOS-only properties

// A Boolean value that indicates whether video playback prevents the system from automatically backgrounding an app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/preventsAutomaticBackgroundingDuringVideoPlayback
func (s_ SampleBufferDisplayLayer) PreventsAutomaticBackgroundingDuringVideoPlayback() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("preventsAutomaticBackgroundingDuringVideoPlayback"))
	return rv
}
func (s_ SampleBufferDisplayLayer) SetPreventsAutomaticBackgroundingDuringVideoPlayback(value bool) {
	s_.ID.Send(objc.RegisterName("setPreventsAutomaticBackgroundingDuringVideoPlayback:"), value)
}





