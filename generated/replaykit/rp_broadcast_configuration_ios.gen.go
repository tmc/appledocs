//go:build darwin && ios

// Code generated from Apple documentation for ReplayKit. DO NOT EDIT.

package replaykit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for RPBroadcastConfiguration


// iOS-only properties

// The duration of movie clips sent the to the movie clip handler extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastConfiguration/clipDuration
func (r_ RPBroadcastConfiguration) ClipDuration() float64 {
	rv := objc.Send[float64](r_.ID, objc.Sel("clipDuration"))
	return rv
}
func (r_ RPBroadcastConfiguration) SetClipDuration(value float64) {
	r_.ID.Send(objc.RegisterName("setClipDuration:"), value)
}

// The compression properties for encoding movie clips that are to be overwritten.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastConfiguration/videoCompressionProperties
func (r_ RPBroadcastConfiguration) VideoCompressionProperties() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](r_.ID, objc.Sel("videoCompressionProperties"))
	return rv
}
func (r_ RPBroadcastConfiguration) SetVideoCompressionProperties(value foundation.IDictionary) {
	r_.ID.Send(objc.RegisterName("setVideoCompressionProperties:"), value)
}





