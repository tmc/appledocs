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

/* debug [class.gen.go]: Generating class MPMovieErrorLog */


/* debug [class_header]: Header for MPMovieErrorLog */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MovieErrorLog */
// An interface definition for the [MovieErrorLog] class.
type IMovieErrorLog interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MovieErrorLog */
	// properties:
	ImageCropRect() corefoundation.CGRect
	SetImageCropRect(value corefoundation.CGRect)
	ShowsRouteButton() bool
	SetShowsRouteButton(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MovieErrorLog */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MovieErrorLog */
// Alloc allocates a new instance without initialization.
func (mc _MovieErrorLogClass) Alloc() MovieErrorLog {
	rv := objc.Send[MovieErrorLog](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MovieErrorLog */
// Data describing network resource playback failures for the associated movie player, including timestamps indicating when each failure occurred.
//
// All movie error log properties are read-only.


// Data describing network resource playback failures for the associated movie player, including timestamps indicating when each failure occurred.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MovieErrorLog *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MovieErrorLog */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MovieErrorLog */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MovieErrorLog */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MovieErrorLog */

// The bounds, in points, of the content area for the full size image associated with the media item artwork.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitemartwork/imagecroprect
func (m_ MovieErrorLog) ImageCropRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](m_.ID, objc.Sel("imageCropRect"))
	return rv
}/* debug [instance_properties/getter]: imageCropRect */


// The bounds, in points, of the content area for the full size image associated with the media item artwork.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitemartwork/imagecroprect
func (m_ MovieErrorLog) SetImageCropRect(value corefoundation.CGRect) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setImageCropRect:"), value)
}/* debug [instance_properties/setter]: imageCropRect */


// A Boolean value that indicates whether the route button is visible in the volume view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/showsroutebutton
func (m_ MovieErrorLog) ShowsRouteButton() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsRouteButton"))
	return rv
}/* debug [instance_properties/getter]: showsRouteButton */


// A Boolean value that indicates whether the route button is visible in the volume view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/showsroutebutton
func (m_ MovieErrorLog) SetShowsRouteButton(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsRouteButton:"), value)
}/* debug [instance_properties/setter]: showsRouteButton */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPMovieErrorLog */


