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

// The class instance for the [MovieErrorLogEvent] class.
var (
	MovieErrorLogEventClass     _MovieErrorLogEventClass
	MovieErrorLogEventClassOnce sync.Once
)

func getMovieErrorLogEventClass() _MovieErrorLogEventClass {
	MovieErrorLogEventClassOnce.Do(func() {
		MovieErrorLogEventClass = _MovieErrorLogEventClass{objc.GetClass("MPMovieErrorLogEvent")}
	})
	return MovieErrorLogEventClass
}

type _MovieErrorLogEventClass struct {
	class objc.Class
}

// An interface definition for the [MovieErrorLogEvent] class.
type IMovieErrorLogEvent interface {
	objectivec.IObject
}

// A single piece of information for a movie error log.
//
// All movie error log event properties are read-only. For a description of movie error logs, see .
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieErrorLogEvent
type MovieErrorLogEvent struct {
	objectivec.Object
}

// MovieErrorLogEventFrom constructs a [MovieErrorLogEvent] from an unsafe.Pointer.
//
// A single piece of information for a movie error log.
func MovieErrorLogEventFrom(ptr unsafe.Pointer) MovieErrorLogEvent {
	return MovieErrorLogEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MovieErrorLogEventClass) Alloc() MovieErrorLogEvent {
	rv := objc.Send[MovieErrorLogEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MovieErrorLogEventClass) New() MovieErrorLogEvent {
	rv := objc.Send[MovieErrorLogEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MovieErrorLogEvent) Init() MovieErrorLogEvent {
	rv := objc.Send[MovieErrorLogEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MovieErrorLogEvent) Autorelease() MovieErrorLogEvent {
	rv := objc.Send[MovieErrorLogEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMovieErrorLogEvent creates a new MovieErrorLogEvent instance.
func NewMovieErrorLogEvent() MovieErrorLogEvent {
	return getMovieErrorLogEventClass().New()
}


// The bounds, in points, of the content area for the full size image associated with the media item artwork.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitemartwork/imagecroprect
func (m_ MovieErrorLogEvent) ImageCropRect() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](m_.ID, objc.Sel("imageCropRect"))
	return rv
}


// SetImageCropRect sets the value of the imageCropRect property.
// The bounds, in points, of the content area for the full size image associated with the media item artwork.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitemartwork/imagecroprect
func (m_ MovieErrorLogEvent) SetImageCropRect(value coregraphics.CGRect) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setImageCropRect:"), value)
}

// The date and time when the error occurred.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieerrorlogevent/date
func (m_ MovieErrorLogEvent) Date() foundation.Date {
	rv := objc.Send[foundation.Date](m_.ID, objc.Sel("date"))
	return rv
}


// SetDate sets the value of the date property.
// The date and time when the error occurred.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieerrorlogevent/date
func (m_ MovieErrorLogEvent) SetDate(value foundation.IDate) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDate:"), value)
}

// A description of the error.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieerrorlogevent/errorcomment
func (m_ MovieErrorLogEvent) ErrorComment() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("errorComment"))
	return rv
}


// SetErrorComment sets the value of the errorComment property.
// A description of the error.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieerrorlogevent/errorcomment
func (m_ MovieErrorLogEvent) SetErrorComment(value appkit.string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setErrorComment:"), value)
}

// The network domain of the error.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieerrorlogevent/errordomain
func (m_ MovieErrorLogEvent) ErrorDomain() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("errorDomain"))
	return rv
}


// SetErrorDomain sets the value of the errorDomain property.
// The network domain of the error.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieerrorlogevent/errordomain
func (m_ MovieErrorLogEvent) SetErrorDomain(value appkit.string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setErrorDomain:"), value)
}

// A unique error code identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieerrorlogevent/errorstatuscode
func (m_ MovieErrorLogEvent) ErrorStatusCode() int {
	rv := objc.Send[int](m_.ID, objc.Sel("errorStatusCode"))
	return rv
}


// SetErrorStatusCode sets the value of the errorStatusCode property.
// A unique error code identifier.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieerrorlogevent/errorstatuscode
func (m_ MovieErrorLogEvent) SetErrorStatusCode(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setErrorStatusCode:"), value)
}

// A globally unique identifier (GUID) for the playback session.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieerrorlogevent/playbacksessionid
func (m_ MovieErrorLogEvent) PlaybackSessionID() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("playbackSessionID"))
	return rv
}


// SetPlaybackSessionID sets the value of the playbackSessionID property.
// A globally unique identifier (GUID) for the playback session.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieerrorlogevent/playbacksessionid
func (m_ MovieErrorLogEvent) SetPlaybackSessionID(value appkit.string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPlaybackSessionID:"), value)
}

// The IP address of the web server that was the source of the error.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieerrorlogevent/serveraddress
func (m_ MovieErrorLogEvent) ServerAddress() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("serverAddress"))
	return rv
}


// SetServerAddress sets the value of the serverAddress property.
// The IP address of the web server that was the source of the error.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieerrorlogevent/serveraddress
func (m_ MovieErrorLogEvent) SetServerAddress(value appkit.string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerAddress:"), value)
}

// The URI of the item playing when the error occurred.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieerrorlogevent/uri
func (m_ MovieErrorLogEvent) Uri() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("uri"))
	return rv
}


// SetUri sets the value of the uri property.
// The URI of the item playing when the error occurred.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieerrorlogevent/uri
func (m_ MovieErrorLogEvent) SetUri(value appkit.string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUri:"), value)
}

// A Boolean value that indicates whether the route button is visible in the volume view.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/showsroutebutton
func (m_ MovieErrorLogEvent) ShowsRouteButton() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsRouteButton"))
	return rv
}


// SetShowsRouteButton sets the value of the showsRouteButton property.
// A Boolean value that indicates whether the route button is visible in the volume view.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/showsroutebutton
func (m_ MovieErrorLogEvent) SetShowsRouteButton(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsRouteButton:"), value)
}



