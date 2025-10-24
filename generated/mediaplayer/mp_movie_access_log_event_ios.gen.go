//go:build darwin && ios

// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for MovieAccessLogEvent


// iOS-only properties

// The accumulated duration of the media played, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieAccessLogEvent/durationWatched
func (m_ MovieAccessLogEvent) DurationWatched() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("durationWatched"))
	return rv
}

// The throughput required to play the stream, as advertised by the web server, in bits per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieAccessLogEvent/indicatedBitrate
func (m_ MovieAccessLogEvent) IndicatedBitrate() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("indicatedBitrate"))
	return rv
}

// The accumulated number of bytes transferred.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieAccessLogEvent/numberOfBytesTransferred
func (m_ MovieAccessLogEvent) NumberOfBytesTransferred() int64 {
	rv := objc.Send[int64](m_.ID, objc.Sel("numberOfBytesTransferred"))
	return rv
}

// The total number of dropped video frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieAccessLogEvent/numberOfDroppedVideoFrames
func (m_ MovieAccessLogEvent) NumberOfDroppedVideoFrames() int {
	rv := objc.Send[int](m_.ID, objc.Sel("numberOfDroppedVideoFrames"))
	return rv
}

// A count of media segments downloaded from the web server to your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieAccessLogEvent/numberOfSegmentsDownloaded
func (m_ MovieAccessLogEvent) NumberOfSegmentsDownloaded() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("numberOfSegmentsDownloaded"))
	return rv
}

// A count of changes to the property over the last uninterrupted period of playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieAccessLogEvent/numberOfServerAddressChanges
func (m_ MovieAccessLogEvent) NumberOfServerAddressChanges() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("numberOfServerAddressChanges"))
	return rv
}

// The total number of playback stalls encountered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieAccessLogEvent/numberOfStalls
func (m_ MovieAccessLogEvent) NumberOfStalls() int {
	rv := objc.Send[int](m_.ID, objc.Sel("numberOfStalls"))
	return rv
}

// The empirical throughput across all media downloaded for the movie player, in bits per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieAccessLogEvent/observedBitrate
func (m_ MovieAccessLogEvent) ObservedBitrate() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("observedBitrate"))
	return rv
}

// A GUID that identifies the playback session to use in HTTP requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieAccessLogEvent/playbackSessionID
func (m_ MovieAccessLogEvent) PlaybackSessionID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("playbackSessionID"))
	return rv
}

// The timestamp for when playback began for the movie log access event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieAccessLogEvent/playbackStartDate
func (m_ MovieAccessLogEvent) PlaybackStartDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](m_.ID, objc.Sel("playbackStartDate"))
	return rv
}

// An offset into the playlist where the last uninterrupted period of playback began, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieAccessLogEvent/playbackStartOffset
func (m_ MovieAccessLogEvent) PlaybackStartOffset() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("playbackStartOffset"))
	return rv
}

// The accumulated duration of the media downloaded, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieAccessLogEvent/segmentsDownloadedDuration
func (m_ MovieAccessLogEvent) SegmentsDownloadedDuration() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("segmentsDownloadedDuration"))
	return rv
}

// The IPv4 or IPv6 address of the web server that was the source of the last delivered media segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieAccessLogEvent/serverAddress
func (m_ MovieAccessLogEvent) ServerAddress() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("serverAddress"))
	return rv
}

// The URI of the playback item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieAccessLogEvent/uri
func (m_ MovieAccessLogEvent) URI() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("URI"))
	return rv
}





