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

/* debug [class.gen.go]: Generating class MPMovieErrorLogEvent */


/* debug [class_header]: Header for MPMovieErrorLogEvent */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MovieErrorLogEvent */
// An interface definition for the [MovieErrorLogEvent] class.
type IMovieErrorLogEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MovieErrorLogEvent */
	// properties:
	ImageCropRect() corefoundation.CGRect
	SetImageCropRect(value corefoundation.CGRect)
	ShowsRouteButton() bool
	SetShowsRouteButton(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MovieErrorLogEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MovieErrorLogEvent */
// Alloc allocates a new instance without initialization.
func (mc _MovieErrorLogEventClass) Alloc() MovieErrorLogEvent {
	rv := objc.Send[MovieErrorLogEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MovieErrorLogEvent */
// A single piece of information for a movie error log.
//
// All movie error log event properties are read-only. For a description of movie error logs, see .


// A single piece of information for a movie error log.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MovieErrorLogEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MovieErrorLogEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MovieErrorLogEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MovieErrorLogEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MovieErrorLogEvent */

// The bounds, in points, of the content area for the full size image associated with the media item artwork.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitemartwork/imagecroprect
func (m_ MovieErrorLogEvent) ImageCropRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](m_.ID, objc.Sel("imageCropRect"))
	return rv
}/* debug [instance_properties/getter]: imageCropRect */


// The bounds, in points, of the content area for the full size image associated with the media item artwork.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitemartwork/imagecroprect
func (m_ MovieErrorLogEvent) SetImageCropRect(value corefoundation.CGRect) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setImageCropRect:"), value)
}/* debug [instance_properties/setter]: imageCropRect */


// A Boolean value that indicates whether the route button is visible in the volume view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/showsroutebutton
func (m_ MovieErrorLogEvent) ShowsRouteButton() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsRouteButton"))
	return rv
}/* debug [instance_properties/getter]: showsRouteButton */


// A Boolean value that indicates whether the route button is visible in the volume view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/showsroutebutton
func (m_ MovieErrorLogEvent) SetShowsRouteButton(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsRouteButton:"), value)
}/* debug [instance_properties/setter]: showsRouteButton */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPMovieErrorLogEvent */


