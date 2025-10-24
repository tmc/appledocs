// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MPMoviePlayerViewController */


/* debug [class_header]: Header for MPMoviePlayerViewController */
// The class instance for the [MoviePlayerViewController] class.
var (
	MoviePlayerViewControllerClass     _MoviePlayerViewControllerClass
	MoviePlayerViewControllerClassOnce sync.Once
)

func getMoviePlayerViewControllerClass() _MoviePlayerViewControllerClass {
	MoviePlayerViewControllerClassOnce.Do(func() {
		MoviePlayerViewControllerClass = _MoviePlayerViewControllerClass{objc.GetClass("MPMoviePlayerViewController")}
	})
	return MoviePlayerViewControllerClass
}

type _MoviePlayerViewControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MoviePlayerViewController */
// An interface definition for the [MoviePlayerViewController] class.
type IMoviePlayerViewController interface {
	IViewController
	
/* debug [class_interface_properties]: Properties for MoviePlayerViewController */
	// properties:
	ImageCropRect() corefoundation.CGRect
	SetImageCropRect(value corefoundation.CGRect)
	ShowsRouteButton() bool
	SetShowsRouteButton(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MoviePlayerViewController */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MoviePlayerViewController */
// Alloc allocates a new instance without initialization.
func (mc _MoviePlayerViewControllerClass) Alloc() MoviePlayerViewController {
	rv := objc.Send[MoviePlayerViewController](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MoviePlayerViewControllerClass) New() MoviePlayerViewController {
	rv := objc.Send[MoviePlayerViewController](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MoviePlayerViewController) Init() MoviePlayerViewController {
	rv := objc.Send[MoviePlayerViewController](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MoviePlayerViewController) Autorelease() MoviePlayerViewController {
	rv := objc.Send[MoviePlayerViewController](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMoviePlayerViewController creates a new MoviePlayerViewController instance.
func NewMoviePlayerViewController() MoviePlayerViewController {
	return getMoviePlayerViewControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MoviePlayerViewController */
// A simple view controller for displaying full-screen movies.
//
// Unlike using an object on its own to present a movie immediately, you can incorporate a movie player view controller wherever you would normally use a view controller. For example, you can present it using a tab bar or navigation bar-based interface, taking advantage of the transitions offered by those interfaces. To present a movie player view controller modally, you typically use the method. This method is part of a category on the class and is implemented by the Media Player framework. The method presents a movie player view controller using the standard transition animations for presenting video content. To dismiss a modally presented movie player view controller, call the method.


// A simple view controller for displaying full-screen movies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlayerViewController
type MoviePlayerViewController struct {
	ViewController
}

// MoviePlayerViewControllerFrom constructs a [MoviePlayerViewController] from an unsafe.Pointer.
//
// A simple view controller for displaying full-screen movies.
func MoviePlayerViewControllerFrom(ptr unsafe.Pointer) MoviePlayerViewController {
	return MoviePlayerViewController{
		ViewController: ViewControllerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MoviePlayerViewController */

// Returns a movie player view controller initialized with the specified movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlayerViewController/init(contentURL:)
func NewMoviePlayerViewControllerWithContentURL(contentURL objc.IObject /* cross-framework: NSURL */) MoviePlayerViewController {
	instance := getMoviePlayerViewControllerClass().Alloc()
	rv := objc.Send[MoviePlayerViewController](instance.ID, objc.Sel("initWithContentURL:"), contentURL)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMoviePlayerViewControllerWithContentURL */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MoviePlayerViewController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MoviePlayerViewController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MoviePlayerViewController */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MoviePlayerViewController */

// The bounds, in points, of the content area for the full size image associated with the media item artwork.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitemartwork/imagecroprect
func (m_ MoviePlayerViewController) ImageCropRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](m_.ID, objc.Sel("imageCropRect"))
	return rv
}/* debug [instance_properties/getter]: imageCropRect */


// The bounds, in points, of the content area for the full size image associated with the media item artwork.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitemartwork/imagecroprect
func (m_ MoviePlayerViewController) SetImageCropRect(value corefoundation.CGRect) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setImageCropRect:"), value)
}/* debug [instance_properties/setter]: imageCropRect */


// A Boolean value that indicates whether the route button is visible in the volume view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/showsroutebutton
func (m_ MoviePlayerViewController) ShowsRouteButton() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsRouteButton"))
	return rv
}/* debug [instance_properties/getter]: showsRouteButton */


// A Boolean value that indicates whether the route button is visible in the volume view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/showsroutebutton
func (m_ MoviePlayerViewController) SetShowsRouteButton(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsRouteButton:"), value)
}/* debug [instance_properties/setter]: showsRouteButton */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPMoviePlayerViewController */


