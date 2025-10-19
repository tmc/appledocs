// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVPlayerItemTrack] class.
var aVPlayerItemTrackClass = _AVPlayerItemTrackClass{objc.GetClass("AVPlayerItemTrack")}

type _AVPlayerItemTrackClass struct {
	class objc.Class
}

// An object that represents the presentation state of an asset track during playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemTrack

type AVPlayerItemTrack struct {
	objectivec.Object
}

// AVPlayerItemTrackFrom constructs a [AVPlayerItemTrack] from an unsafe.Pointer.
//
// An object that represents the presentation state of an asset track during playback.
func AVPlayerItemTrackFrom(ptr unsafe.Pointer) AVPlayerItemTrack {
	return AVPlayerItemTrack{objectivec.Object{objc.ID(ptr)}}
}



