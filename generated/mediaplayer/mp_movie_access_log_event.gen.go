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

/* debug [class.gen.go]: Generating class MPMovieAccessLogEvent */


/* debug [class_header]: Header for MPMovieAccessLogEvent */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MovieAccessLogEvent */
// An interface definition for the [MovieAccessLogEvent] class.
type IMovieAccessLogEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MovieAccessLogEvent */
	// properties:
	ImageCropRect() corefoundation.CGRect
	SetImageCropRect(value corefoundation.CGRect)
	ShowsRouteButton() bool
	SetShowsRouteButton(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MovieAccessLogEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MovieAccessLogEvent */
// Alloc allocates a new instance without initialization.
func (mc _MovieAccessLogEventClass) Alloc() MovieAccessLogEvent {
	rv := objc.Send[MovieAccessLogEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MovieAccessLogEvent */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MovieAccessLogEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MovieAccessLogEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MovieAccessLogEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MovieAccessLogEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MovieAccessLogEvent */

// The bounds, in points, of the content area for the full size image associated with the media item artwork.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitemartwork/imagecroprect
func (m_ MovieAccessLogEvent) ImageCropRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](m_.ID, objc.Sel("imageCropRect"))
	return rv
}/* debug [instance_properties/getter]: imageCropRect */


// The bounds, in points, of the content area for the full size image associated with the media item artwork.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitemartwork/imagecroprect
func (m_ MovieAccessLogEvent) SetImageCropRect(value corefoundation.CGRect) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setImageCropRect:"), value)
}/* debug [instance_properties/setter]: imageCropRect */


// A Boolean value that indicates whether the route button is visible in the volume view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/showsroutebutton
func (m_ MovieAccessLogEvent) ShowsRouteButton() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsRouteButton"))
	return rv
}/* debug [instance_properties/getter]: showsRouteButton */


// A Boolean value that indicates whether the route button is visible in the volume view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/showsroutebutton
func (m_ MovieAccessLogEvent) SetShowsRouteButton(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsRouteButton:"), value)
}/* debug [instance_properties/setter]: showsRouteButton */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPMovieAccessLogEvent */


