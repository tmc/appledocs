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

// The class instance for the [MovieErrorLog] class.
var (
	MovieErrorLogClass     _MovieErrorLogClass
	MovieErrorLogClassOnce sync.Once
)

func getMovieErrorLogClass() _MovieErrorLogClass {
	MovieErrorLogClassOnce.Do(func() {
		MovieErrorLogClass = _MovieErrorLogClass{objc.GetClass("MPMovieErrorLog")}
	})
	return MovieErrorLogClass
}

type _MovieErrorLogClass struct {
	class objc.Class
}

// An interface definition for the [MovieErrorLog] class.
type IMovieErrorLog interface {
	objectivec.IObject
	ImageCropRect() coregraphics.CGRect
	SetImageCropRect(value coregraphics.CGRect)
	Events() unsafe.Pointer
	SetEvents(value unsafe.Pointer)
	ExtendedLogData() foundation.Data
	SetExtendedLogData(value foundation.IData)
	ExtendedLogDataStringEncoding() uint
	SetExtendedLogDataStringEncoding(value uint)
	ShowsRouteButton() bool
	SetShowsRouteButton(value bool)
}

// Data describing network resource playback failures for the associated movie player, including timestamps indicating when each failure occurred.
//
// All movie error log properties are read-only.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieErrorLog
type MovieErrorLog struct {
	objectivec.Object
}

// MovieErrorLogFrom constructs a [MovieErrorLog] from an unsafe.Pointer.
//
// Data describing network resource playback failures for the associated movie player, including timestamps indicating when each failure occurred.
func MovieErrorLogFrom(ptr unsafe.Pointer) MovieErrorLog {
	return MovieErrorLog{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MovieErrorLogClass) Alloc() MovieErrorLog {
	rv := objc.Send[MovieErrorLog](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MovieErrorLogClass) New() MovieErrorLog {
	rv := objc.Send[MovieErrorLog](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MovieErrorLog) Init() MovieErrorLog {
	rv := objc.Send[MovieErrorLog](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MovieErrorLog) Autorelease() MovieErrorLog {
	rv := objc.Send[MovieErrorLog](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMovieErrorLog creates a new MovieErrorLog instance.
func NewMovieErrorLog() MovieErrorLog {
	return getMovieErrorLogClass().New()
}


// The bounds, in points, of the content area for the full size image associated with the media item artwork.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitemartwork/imagecroprect
func (m_ MovieErrorLog) ImageCropRect() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](m_.ID, objc.Sel("imageCropRect"))
	return rv
}


// SetImageCropRect sets the value of the imageCropRect property.
// The bounds, in points, of the content area for the full size image associated with the media item artwork.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitemartwork/imagecroprect
func (m_ MovieErrorLog) SetImageCropRect(value coregraphics.CGRect) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setImageCropRect:"), value)
}

// The events in the movie error log.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieerrorlog/events
func (m_ MovieErrorLog) Events() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("events"))
	return rv
}


// SetEvents sets the value of the events property.
// The events in the movie error log.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieerrorlog/events
func (m_ MovieErrorLog) SetEvents(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEvents:"), value)
}

// A textual version of the web server error log.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieerrorlog/extendedlogdata
func (m_ MovieErrorLog) ExtendedLogData() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("extendedLogData"))
	return rv
}


// SetExtendedLogData sets the value of the extendedLogData property.
// A textual version of the web server error log.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieerrorlog/extendedlogdata
func (m_ MovieErrorLog) SetExtendedLogData(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExtendedLogData:"), value)
}

// The string encoding for the extended log data property.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieerrorlog/extendedlogdatastringencoding
func (m_ MovieErrorLog) ExtendedLogDataStringEncoding() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("extendedLogDataStringEncoding"))
	return rv
}


// SetExtendedLogDataStringEncoding sets the value of the extendedLogDataStringEncoding property.
// The string encoding for the extended log data property.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieerrorlog/extendedlogdatastringencoding
func (m_ MovieErrorLog) SetExtendedLogDataStringEncoding(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExtendedLogDataStringEncoding:"), value)
}

// A Boolean value that indicates whether the route button is visible in the volume view.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/showsroutebutton
func (m_ MovieErrorLog) ShowsRouteButton() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsRouteButton"))
	return rv
}


// SetShowsRouteButton sets the value of the showsRouteButton property.
// A Boolean value that indicates whether the route button is visible in the volume view.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/showsroutebutton
func (m_ MovieErrorLog) SetShowsRouteButton(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsRouteButton:"), value)
}



