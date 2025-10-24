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

/* debug [class.gen.go]: Generating class MPMovieAccessLog */


/* debug [class_header]: Header for MPMovieAccessLog */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MovieAccessLog */
// An interface definition for the [MovieAccessLog] class.
type IMovieAccessLog interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MovieAccessLog */
	// properties:
	ImageCropRect() corefoundation.CGRect
	SetImageCropRect(value corefoundation.CGRect)
	AccessLog() IMPMovieAccessLog
	SetAccessLog(value IMPMovieAccessLog)
	ShowsRouteButton() bool
	SetShowsRouteButton(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MovieAccessLog */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MovieAccessLog */
// Alloc allocates a new instance without initialization.
func (mc _MovieAccessLogClass) Alloc() MovieAccessLog {
	rv := objc.Send[MovieAccessLog](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MovieAccessLog */
// Key metrics about network playback for an associated movie player that’s playing streamed content.
//
// The log presents these metrics as a collection of instances and also makes it available in a textual format. A movie access log describes one uninterrupted period of playback. A movie player (an instance of the class) can access this log from its property. All movie access log properties are read-only.


// Key metrics about network playback for an associated movie player that’s playing streamed content.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MovieAccessLog *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MovieAccessLog */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MovieAccessLog */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MovieAccessLog */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MovieAccessLog */

// The bounds, in points, of the content area for the full size image associated with the media item artwork.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitemartwork/imagecroprect
func (m_ MovieAccessLog) ImageCropRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](m_.ID, objc.Sel("imageCropRect"))
	return rv
}/* debug [instance_properties/getter]: imageCropRect */


// The bounds, in points, of the content area for the full size image associated with the media item artwork.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitemartwork/imagecroprect
func (m_ MovieAccessLog) SetImageCropRect(value corefoundation.CGRect) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setImageCropRect:"), value)
}/* debug [instance_properties/setter]: imageCropRect */


// A snapshot of the network playback log for the movie player if it is playing a network stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/accesslog
func (m_ MovieAccessLog) AccessLog() IMPMovieAccessLog {
	rv := objc.Send[MovieAccessLog](m_.ID, objc.Sel("accessLog"))
	return rv
}/* debug [instance_properties/getter]: accessLog */


// A snapshot of the network playback log for the movie player if it is playing a network stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/accesslog
func (m_ MovieAccessLog) SetAccessLog(value IMPMovieAccessLog) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAccessLog:"), value)
}/* debug [instance_properties/setter]: accessLog */


// A Boolean value that indicates whether the route button is visible in the volume view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/showsroutebutton
func (m_ MovieAccessLog) ShowsRouteButton() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsRouteButton"))
	return rv
}/* debug [instance_properties/getter]: showsRouteButton */


// A Boolean value that indicates whether the route button is visible in the volume view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/showsroutebutton
func (m_ MovieAccessLog) SetShowsRouteButton(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsRouteButton:"), value)
}/* debug [instance_properties/setter]: showsRouteButton */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPMovieAccessLog */


