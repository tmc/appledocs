//go:build darwin && ios

// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for INPlayMediaIntent


// iOS-only properties

// The playback speed for a media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INPlayMediaIntent/playbackSpeed-6ngbq
func (i_ INPlayMediaIntent) PlaybackSpeed() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](i_.ID, objc.Sel("playbackSpeed"))
	return rv
}

// The resume playback setting at the time the user plays the media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INPlayMediaIntent/resumePlayback-9zfyp
func (i_ INPlayMediaIntent) ResumePlayback() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](i_.ID, objc.Sel("resumePlayback"))
	return rv
}





