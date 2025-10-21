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

// The class instance for the [MovieAccessLog] class.
var (
	MovieAccessLogClass     _MovieAccessLogClass
	MovieAccessLogClassOnce sync.Once
)

func getMovieAccessLogClass() _MovieAccessLogClass {
	MovieAccessLogClassOnce.Do(func() {
		MovieAccessLogClass = _MovieAccessLogClass{objc.GetClass("MPMovieAccessLog")}
	})
	return MovieAccessLogClass
}

type _MovieAccessLogClass struct {
	class objc.Class
}

// An interface definition for the [MovieAccessLog] class.
type IMovieAccessLog interface {
	objectivec.IObject
}

// Key metrics about network playback for an associated movie player that’s playing streamed content.
//
// The log presents these metrics as a collection of instances and also makes it available in a textual format. A movie access log describes one uninterrupted period of playback. A movie player (an instance of the class) can access this log from its property. All movie access log properties are read-only.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieAccessLog
type MovieAccessLog struct {
	objectivec.Object
}

// MovieAccessLogFrom constructs a [MovieAccessLog] from an unsafe.Pointer.
//
// Key metrics about network playback for an associated movie player that’s playing streamed content.
func MovieAccessLogFrom(ptr unsafe.Pointer) MovieAccessLog {
	return MovieAccessLog{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MovieAccessLogClass) Alloc() MovieAccessLog {
	rv := objc.Send[MovieAccessLog](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MovieAccessLogClass) New() MovieAccessLog {
	rv := objc.Send[MovieAccessLog](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MovieAccessLog) Init() MovieAccessLog {
	rv := objc.Send[MovieAccessLog](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MovieAccessLog) Autorelease() MovieAccessLog {
	rv := objc.Send[MovieAccessLog](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMovieAccessLog creates a new MovieAccessLog instance.
func NewMovieAccessLog() MovieAccessLog {
	return getMovieAccessLogClass().New()
}


// The bounds, in points, of the content area for the full size image associated with the media item artwork.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitemartwork/imagecroprect
func (m_ MovieAccessLog) ImageCropRect() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](m_.ID, objc.Sel("imageCropRect"))
	return rv
}


// SetImageCropRect sets the value of the imageCropRect property.
// The bounds, in points, of the content area for the full size image associated with the media item artwork.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitemartwork/imagecroprect
func (m_ MovieAccessLog) SetImageCropRect(value coregraphics.CGRect) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setImageCropRect:"), value)
}

// The events in the movie access log.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieaccesslog/events
func (m_ MovieAccessLog) Events() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("events"))
	return rv
}


// SetEvents sets the value of the events property.
// The events in the movie access log.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieaccesslog/events
func (m_ MovieAccessLog) SetEvents(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEvents:"), value)
}

// A textual version of the web server access log for the associated movie player.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieaccesslog/extendedlogdata
func (m_ MovieAccessLog) ExtendedLogData() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("extendedLogData"))
	return rv
}


// SetExtendedLogData sets the value of the extendedLogData property.
// A textual version of the web server access log for the associated movie player.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieaccesslog/extendedlogdata
func (m_ MovieAccessLog) SetExtendedLogData(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExtendedLogData:"), value)
}

// The string encoding for the
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieaccesslog/extendedlogdatastringencoding
func (m_ MovieAccessLog) ExtendedLogDataStringEncoding() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("extendedLogDataStringEncoding"))
	return rv
}


// SetExtendedLogDataStringEncoding sets the value of the extendedLogDataStringEncoding property.
// The string encoding for the

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieaccesslog/extendedlogdatastringencoding
func (m_ MovieAccessLog) SetExtendedLogDataStringEncoding(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExtendedLogDataStringEncoding:"), value)
}

// A snapshot of the network playback log for the movie player if it is playing a network stream.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/accesslog
func (m_ MovieAccessLog) AccessLog() MPMovieAccessLog {
	rv := objc.Send[MPMovieAccessLog](m_.ID, objc.Sel("accessLog"))
	return rv
}


// SetAccessLog sets the value of the accessLog property.
// A snapshot of the network playback log for the movie player if it is playing a network stream.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/accesslog
func (m_ MovieAccessLog) SetAccessLog(value IMPMovieAccessLog) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAccessLog:"), value)
}

// A Boolean value that indicates whether the route button is visible in the volume view.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/showsroutebutton
func (m_ MovieAccessLog) ShowsRouteButton() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsRouteButton"))
	return rv
}


// SetShowsRouteButton sets the value of the showsRouteButton property.
// A Boolean value that indicates whether the route button is visible in the volume view.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/showsroutebutton
func (m_ MovieAccessLog) SetShowsRouteButton(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsRouteButton:"), value)
}



