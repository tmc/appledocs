// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
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
	// properties:
	ImageCropRect() objc.IObject /* cross-framework: Rect */
	SetImageCropRect(value objc.IObject /* cross-framework: Rect */)
	DurationWatched() float64
	SetDurationWatched(value float64)
	IndicatedBitrate() float64
	SetIndicatedBitrate(value float64)
	NumberOfBytesTransferred() unsafe.Pointer
	SetNumberOfBytesTransferred(value unsafe.Pointer)
	NumberOfDroppedVideoFrames() int
	SetNumberOfDroppedVideoFrames(value int)
	NumberOfSegmentsDownloaded() int
	SetNumberOfSegmentsDownloaded(value int)
	NumberOfServerAddressChanges() int
	SetNumberOfServerAddressChanges(value int)
	NumberOfStalls() int
	SetNumberOfStalls(value int)
	ObservedBitrate() float64
	SetObservedBitrate(value float64)
	PlaybackSessionID() objc.IObject /* cross-framework: NSString */
	SetPlaybackSessionID(value objc.IObject /* cross-framework: NSString */)
	PlaybackStartDate() objc.IObject /* cross-framework: Date */
	SetPlaybackStartDate(value objc.IObject /* cross-framework: Date */)
	PlaybackStartOffset() float64
	SetPlaybackStartOffset(value float64)
	SegmentsDownloadedDuration() float64
	SetSegmentsDownloadedDuration(value float64)
	ServerAddress() objc.IObject /* cross-framework: NSString */
	SetServerAddress(value objc.IObject /* cross-framework: NSString */)
	Uri() objc.IObject /* cross-framework: NSString */
	SetUri(value objc.IObject /* cross-framework: NSString */)
	ShowsRouteButton() bool
	SetShowsRouteButton(value bool)
	// methods:
}

// A single piece of information for a movie access log.
//
// For a description of movie access logs, see .


// A single piece of information for a movie access log.
//
// [Full Topic]
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



// The bounds, in points, of the content area for the full size image associated with the media item artwork.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitemartwork/imagecroprect
func (m_ MovieAccessLogEvent) ImageCropRect() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](m_.ID, objc.Sel("imageCropRect"))
	return rv
}


// The bounds, in points, of the content area for the full size image associated with the media item artwork.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitemartwork/imagecroprect
func (m_ MovieAccessLogEvent) SetImageCropRect(value objc.IObject /* cross-framework: Rect */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setImageCropRect:"), value)
}


// The accumulated duration of the media played, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieaccesslogevent/durationwatched
func (m_ MovieAccessLogEvent) DurationWatched() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("durationWatched"))
	return rv
}


// The accumulated duration of the media played, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieaccesslogevent/durationwatched
func (m_ MovieAccessLogEvent) SetDurationWatched(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDurationWatched:"), value)
}


// The throughput required to play the stream, as advertised by the web server, in bits per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieaccesslogevent/indicatedbitrate
func (m_ MovieAccessLogEvent) IndicatedBitrate() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("indicatedBitrate"))
	return rv
}


// The throughput required to play the stream, as advertised by the web server, in bits per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieaccesslogevent/indicatedbitrate
func (m_ MovieAccessLogEvent) SetIndicatedBitrate(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIndicatedBitrate:"), value)
}


// The accumulated number of bytes transferred.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieaccesslogevent/numberofbytestransferred
func (m_ MovieAccessLogEvent) NumberOfBytesTransferred() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("numberOfBytesTransferred"))
	return rv
}


// The accumulated number of bytes transferred.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieaccesslogevent/numberofbytestransferred
func (m_ MovieAccessLogEvent) SetNumberOfBytesTransferred(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNumberOfBytesTransferred:"), value)
}


// The total number of dropped video frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieaccesslogevent/numberofdroppedvideoframes
func (m_ MovieAccessLogEvent) NumberOfDroppedVideoFrames() int {
	rv := objc.Send[int](m_.ID, objc.Sel("numberOfDroppedVideoFrames"))
	return rv
}


// The total number of dropped video frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieaccesslogevent/numberofdroppedvideoframes
func (m_ MovieAccessLogEvent) SetNumberOfDroppedVideoFrames(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNumberOfDroppedVideoFrames:"), value)
}


// A count of media segments downloaded from the web server to your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieaccesslogevent/numberofsegmentsdownloaded
func (m_ MovieAccessLogEvent) NumberOfSegmentsDownloaded() int {
	rv := objc.Send[int](m_.ID, objc.Sel("numberOfSegmentsDownloaded"))
	return rv
}


// A count of media segments downloaded from the web server to your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieaccesslogevent/numberofsegmentsdownloaded
func (m_ MovieAccessLogEvent) SetNumberOfSegmentsDownloaded(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNumberOfSegmentsDownloaded:"), value)
}


// A count of changes to the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieaccesslogevent/numberofserveraddresschanges
func (m_ MovieAccessLogEvent) NumberOfServerAddressChanges() int {
	rv := objc.Send[int](m_.ID, objc.Sel("numberOfServerAddressChanges"))
	return rv
}


// A count of changes to the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieaccesslogevent/numberofserveraddresschanges
func (m_ MovieAccessLogEvent) SetNumberOfServerAddressChanges(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNumberOfServerAddressChanges:"), value)
}


// The total number of playback stalls encountered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieaccesslogevent/numberofstalls
func (m_ MovieAccessLogEvent) NumberOfStalls() int {
	rv := objc.Send[int](m_.ID, objc.Sel("numberOfStalls"))
	return rv
}


// The total number of playback stalls encountered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieaccesslogevent/numberofstalls
func (m_ MovieAccessLogEvent) SetNumberOfStalls(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNumberOfStalls:"), value)
}


// The empirical throughput across all media downloaded for the movie player, in bits per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieaccesslogevent/observedbitrate
func (m_ MovieAccessLogEvent) ObservedBitrate() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("observedBitrate"))
	return rv
}


// The empirical throughput across all media downloaded for the movie player, in bits per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieaccesslogevent/observedbitrate
func (m_ MovieAccessLogEvent) SetObservedBitrate(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setObservedBitrate:"), value)
}


// A GUID that identifies the playback session to use in HTTP requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieaccesslogevent/playbacksessionid
func (m_ MovieAccessLogEvent) PlaybackSessionID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("playbackSessionID"))
	return rv
}


// A GUID that identifies the playback session to use in HTTP requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieaccesslogevent/playbacksessionid
func (m_ MovieAccessLogEvent) SetPlaybackSessionID(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPlaybackSessionID:"), value)
}


// The timestamp for when playback began for the movie log access event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieaccesslogevent/playbackstartdate
func (m_ MovieAccessLogEvent) PlaybackStartDate() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](m_.ID, objc.Sel("playbackStartDate"))
	return rv
}


// The timestamp for when playback began for the movie log access event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieaccesslogevent/playbackstartdate
func (m_ MovieAccessLogEvent) SetPlaybackStartDate(value objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPlaybackStartDate:"), value)
}


// An offset into the playlist where the last uninterrupted period of playback began, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieaccesslogevent/playbackstartoffset
func (m_ MovieAccessLogEvent) PlaybackStartOffset() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("playbackStartOffset"))
	return rv
}


// An offset into the playlist where the last uninterrupted period of playback began, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieaccesslogevent/playbackstartoffset
func (m_ MovieAccessLogEvent) SetPlaybackStartOffset(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPlaybackStartOffset:"), value)
}


// The accumulated duration of the media downloaded, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieaccesslogevent/segmentsdownloadedduration
func (m_ MovieAccessLogEvent) SegmentsDownloadedDuration() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("segmentsDownloadedDuration"))
	return rv
}


// The accumulated duration of the media downloaded, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieaccesslogevent/segmentsdownloadedduration
func (m_ MovieAccessLogEvent) SetSegmentsDownloadedDuration(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSegmentsDownloadedDuration:"), value)
}


// The IPv4 or IPv6 address of the web server that was the source of the last delivered media segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieaccesslogevent/serveraddress
func (m_ MovieAccessLogEvent) ServerAddress() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("serverAddress"))
	return rv
}


// The IPv4 or IPv6 address of the web server that was the source of the last delivered media segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieaccesslogevent/serveraddress
func (m_ MovieAccessLogEvent) SetServerAddress(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerAddress:"), value)
}


// The URI of the playback item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieaccesslogevent/uri
func (m_ MovieAccessLogEvent) Uri() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("uri"))
	return rv
}


// The URI of the playback item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieaccesslogevent/uri
func (m_ MovieAccessLogEvent) SetUri(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUri:"), value)
}


// A Boolean value that indicates whether the route button is visible in the volume view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/showsroutebutton
func (m_ MovieAccessLogEvent) ShowsRouteButton() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsRouteButton"))
	return rv
}


// A Boolean value that indicates whether the route button is visible in the volume view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/showsroutebutton
func (m_ MovieAccessLogEvent) SetShowsRouteButton(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsRouteButton:"), value)
}



