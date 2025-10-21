// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MovieAccessLogEvent] class.
var (
	MovieAccessLogEventClass     _MovieAccessLogEventClass
	MovieAccessLogEventClassOnce sync.Once
)

func getMovieAccessLogEventClass() _MovieAccessLogEventClass {
	MovieAccessLogEventClassOnce.Do(func() {
		MovieAccessLogEventClass = _MovieAccessLogEventClass{objc.GetClass("MPMovieAccessLogEvent")}
	})
	return MovieAccessLogEventClass
}

type _MovieAccessLogEventClass struct {
	class objc.Class
}

// An interface definition for the [MovieAccessLogEvent] class.
type IMovieAccessLogEvent interface {
	objectivec.IObject
}

// A single piece of information for a movie access log.
//
// For a description of movie access logs, see .
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieAccessLogEvent
type MovieAccessLogEvent struct {
	objectivec.Object
}

// MovieAccessLogEventFrom constructs a [MovieAccessLogEvent] from an unsafe.Pointer.
//
// A single piece of information for a movie access log.
func MovieAccessLogEventFrom(ptr unsafe.Pointer) MovieAccessLogEvent {
	return MovieAccessLogEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MovieAccessLogEventClass) Alloc() MovieAccessLogEvent {
	rv := objc.Send[MovieAccessLogEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MovieAccessLogEventClass) New() MovieAccessLogEvent {
	rv := objc.Send[MovieAccessLogEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MovieAccessLogEvent) Init() MovieAccessLogEvent {
	rv := objc.Send[MovieAccessLogEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MovieAccessLogEvent) Autorelease() MovieAccessLogEvent {
	rv := objc.Send[MovieAccessLogEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMovieAccessLogEvent creates a new MovieAccessLogEvent instance.
func NewMovieAccessLogEvent() MovieAccessLogEvent {
	return getMovieAccessLogEventClass().New()
}


// A Boolean value that indicates whether the route button is visible in the volume view.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/showsroutebutton
func (m_ MovieAccessLogEvent) ShowsRouteButton() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsRouteButton"))
	return rv
}


// SetShowsRouteButton sets the value of the showsRouteButton property.
// A Boolean value that indicates whether the route button is visible in the volume view.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/showsroutebutton
func (m_ MovieAccessLogEvent) SetShowsRouteButton(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsRouteButton:"), value)
}

// The throughput required to play the stream, as advertised by the web server, in bits per second.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieaccesslogevent/indicatedbitrate
func (m_ MovieAccessLogEvent) IndicatedBitrate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("indicatedBitrate"))
	return rv
}


// SetIndicatedBitrate sets the value of the indicatedBitrate property.
// The throughput required to play the stream, as advertised by the web server, in bits per second.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieaccesslogevent/indicatedbitrate
func (m_ MovieAccessLogEvent) SetIndicatedBitrate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIndicatedBitrate:"), value)
}

// The URI of the playback item.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieaccesslogevent/uri
func (m_ MovieAccessLogEvent) Uri() string {
	rv := objc.Send[string](m_.ID, objc.Sel("uri"))
	return rv
}


// SetUri sets the value of the uri property.
// The URI of the playback item.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieaccesslogevent/uri
func (m_ MovieAccessLogEvent) SetUri(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUri:"), objc.String(value))
}

// A GUID that identifies the playback session to use in HTTP requests.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieaccesslogevent/playbacksessionid
func (m_ MovieAccessLogEvent) PlaybackSessionID() string {
	rv := objc.Send[string](m_.ID, objc.Sel("playbackSessionID"))
	return rv
}


// SetPlaybackSessionID sets the value of the playbackSessionID property.
// A GUID that identifies the playback session to use in HTTP requests.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieaccesslogevent/playbacksessionid
func (m_ MovieAccessLogEvent) SetPlaybackSessionID(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPlaybackSessionID:"), objc.String(value))
}

// The empirical throughput across all media downloaded for the movie player, in bits per second.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieaccesslogevent/observedbitrate
func (m_ MovieAccessLogEvent) ObservedBitrate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("observedBitrate"))
	return rv
}


// SetObservedBitrate sets the value of the observedBitrate property.
// The empirical throughput across all media downloaded for the movie player, in bits per second.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieaccesslogevent/observedbitrate
func (m_ MovieAccessLogEvent) SetObservedBitrate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setObservedBitrate:"), value)
}

// The bounds, in points, of the content area for the full size image associated with the media item artwork.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitemartwork/imagecroprect
func (m_ MovieAccessLogEvent) ImageCropRect() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](m_.ID, objc.Sel("imageCropRect"))
	return rv
}


// SetImageCropRect sets the value of the imageCropRect property.
// The bounds, in points, of the content area for the full size image associated with the media item artwork.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitemartwork/imagecroprect
func (m_ MovieAccessLogEvent) SetImageCropRect(value coregraphics.CGRect) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setImageCropRect:"), value)
}

// The accumulated duration of the media played, in seconds.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieAccessLogEvent/durationWatched
func (m_ MovieAccessLogEvent) DurationWatched() foundation.TimeInterval {
	rv := objc.Send[foundation.TimeInterval](m_.ID, objc.Sel("durationWatched"))
	return rv
}

// The accumulated number of bytes transferred.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieAccessLogEvent/numberOfBytesTransferred
func (m_ MovieAccessLogEvent) NumberOfBytesTransferred() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("numberOfBytesTransferred"))
	return rv
}

// The total number of dropped video frames.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieAccessLogEvent/numberOfDroppedVideoFrames
func (m_ MovieAccessLogEvent) NumberOfDroppedVideoFrames() int {
	rv := objc.Send[int](m_.ID, objc.Sel("numberOfDroppedVideoFrames"))
	return rv
}

// A count of media segments downloaded from the web server to your app.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieAccessLogEvent/numberOfSegmentsDownloaded
func (m_ MovieAccessLogEvent) NumberOfSegmentsDownloaded() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("numberOfSegmentsDownloaded"))
	return rv
}

// A count of changes to the property over the last uninterrupted period of playback.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieAccessLogEvent/numberOfServerAddressChanges
func (m_ MovieAccessLogEvent) NumberOfServerAddressChanges() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("numberOfServerAddressChanges"))
	return rv
}

// The total number of playback stalls encountered.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieAccessLogEvent/numberOfStalls
func (m_ MovieAccessLogEvent) NumberOfStalls() int {
	rv := objc.Send[int](m_.ID, objc.Sel("numberOfStalls"))
	return rv
}

// The timestamp for when playback began for the movie log access event.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieAccessLogEvent/playbackStartDate
func (m_ MovieAccessLogEvent) PlaybackStartDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("playbackStartDate"))
	return rv
}

// An offset into the playlist where the last uninterrupted period of playback began, in seconds.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieAccessLogEvent/playbackStartOffset
func (m_ MovieAccessLogEvent) PlaybackStartOffset() foundation.TimeInterval {
	rv := objc.Send[foundation.TimeInterval](m_.ID, objc.Sel("playbackStartOffset"))
	return rv
}

// The accumulated duration of the media downloaded, in seconds.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieAccessLogEvent/segmentsDownloadedDuration
func (m_ MovieAccessLogEvent) SegmentsDownloadedDuration() foundation.TimeInterval {
	rv := objc.Send[foundation.TimeInterval](m_.ID, objc.Sel("segmentsDownloadedDuration"))
	return rv
}

// The IPv4 or IPv6 address of the web server that was the source of the last delivered media segment.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieAccessLogEvent/serverAddress
func (m_ MovieAccessLogEvent) ServerAddress() string {
	rv := objc.Send[string](m_.ID, objc.Sel("serverAddress"))
	return rv
}



