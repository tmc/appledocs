// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVPlayerItemIntegratedTimeline] class.
var aVPlayerItemIntegratedTimelineClass = _AVPlayerItemIntegratedTimelineClass{objc.GetClass("AVPlayerItemIntegratedTimeline")}

type _AVPlayerItemIntegratedTimelineClass struct {
	class objc.Class
}

// An object that models the timeline and playback sequence of a primary player item and scheduled interstitial events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimeline

type AVPlayerItemIntegratedTimeline struct {
	objectivec.Object
}

// AVPlayerItemIntegratedTimelineFrom constructs a [AVPlayerItemIntegratedTimeline] from an unsafe.Pointer.
//
// An object that models the timeline and playback sequence of a primary player item and scheduled interstitial events.
func AVPlayerItemIntegratedTimelineFrom(ptr unsafe.Pointer) AVPlayerItemIntegratedTimeline {
	return AVPlayerItemIntegratedTimeline{objectivec.Object{objc.ID(ptr)}}
}



