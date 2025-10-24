// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPMoviePlayerController */


/* debug [class_header]: Header for MPMoviePlayerController */
// The class instance for the [MoviePlayerController] class.
var (
	MoviePlayerControllerClass     _MoviePlayerControllerClass
	MoviePlayerControllerClassOnce sync.Once
)

func getMoviePlayerControllerClass() _MoviePlayerControllerClass {
	MoviePlayerControllerClassOnce.Do(func() {
		MoviePlayerControllerClass = _MoviePlayerControllerClass{objc.GetClass("MPMoviePlayerController")}
	})
	return MoviePlayerControllerClass
}

type _MoviePlayerControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MoviePlayerController */
// An interface definition for the [MoviePlayerController] class.
type IMoviePlayerController interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MoviePlayerController */
	// properties:
	ImageCropRect() corefoundation.CGRect
	SetImageCropRect(value corefoundation.CGRect)
	IsAirPlayVideoActive() bool
	SetIsAirPlayVideoActive(value bool)
	IsFullscreen() bool
	SetIsFullscreen(value bool)
	MPMoviePlayerPlaybackDidFinishReasonUserInfoKey() objc.IObject /* cross-framework: NSString */
	ShowsRouteButton() bool
	SetShowsRouteButton(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MoviePlayerController */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MoviePlayerController */
// Alloc allocates a new instance without initialization.
func (mc _MoviePlayerControllerClass) Alloc() MoviePlayerController {
	rv := objc.Send[MoviePlayerController](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MoviePlayerControllerClass) New() MoviePlayerController {
	rv := objc.Send[MoviePlayerController](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MoviePlayerController) Init() MoviePlayerController {
	rv := objc.Send[MoviePlayerController](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MoviePlayerController) Autorelease() MoviePlayerController {
	rv := objc.Send[MoviePlayerController](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMoviePlayerController creates a new MoviePlayerController instance.
func NewMoviePlayerController() MoviePlayerController {
	return getMoviePlayerControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MoviePlayerController */
// A type of movie player that manages the playback of a movie from a file or a network stream.
//
// Playback occurs in a view owned by the movie player and takes place either fullscreen or inline. You can incorporate a movie player’s view into a view hierarchy owned by your app, or use an MPMoviePlayerViewController object to manage the presentation for you. Movie players support wireless movie playback to AirPlay-enabled hardware such as Apple TV. AirPlay playback is enabled by default. To disable AirPlay in your app, set the property to . In iOS 8.0 and later, users access AirPlay compatible hardware through the Control Panel; no AirPlay control is displayed by the movie player. When you add a movie player’s view to your app’s view hierarchy, be sure to size the frame correctly, as shown here: Consider a movie player view to be an opaque structure. You can add your own custom subviews to layer content on top of the movie but you must never modify any of its existing subviews. In addition to layering content on top of a movie, you can provide custom background content by adding subviews to the view in the property. Custom subviews are supported in both inline and fullscreen playback modes but you must adjust the positions of your views when entering or exiting fullscreen mode. Use the and notifications to detect changes to and from fullscreen mode. This class supports programmatic control of movie playback, and user-based control via buttons supplied by the movie player. You can control most aspects of playback programmatically using the methods and properties of the protocol, to which this class conforms. The methods and properties of that protocol let you start and stop playback, seek forward and backward through the movie’s content, and even change the playback rate. In addition, the property of this class lets you display a set of standard system controls that allow the user to manipulate playback. You can also set the property for network-based content to start automatically. You typically specify the movie you want to play when you create a new object. However, you can also change the currently playing movie by changing the value in the property. Changing this property lets you reuse the same movie player controller object in multiple places. For performance reasons you may want to play movies as local files. Do this by first downloading them to a local directory. To facilitate the creation of video bookmarks or chapter links for a long movie, the class defines methods for generating thumbnail images at specific times within a movie. You can request a single thumbnail image using the method or request multiple thumbnail images using the method. To play a network stream whose URL requires access credentials, first create an appropriate object. Do this by calling, for example, the method, as shown here: In addition, create an appropriate object, as shown here. Make appropriate modifications for the realm you are accessing: Add the URL credential and the protection space to the object. Do this by calling, for example, the method, as shown here: With the credential and protection space information in place, you can then play the protected stream.


// A type of movie player that manages the playback of a movie from a file or a network stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlayerController
type MoviePlayerController struct {
	objectivec.Object
}

// MoviePlayerControllerFrom constructs a [MoviePlayerController] from an unsafe.Pointer.
//
// A type of movie player that manages the playback of a movie from a file or a network stream.
func MoviePlayerControllerFrom(ptr unsafe.Pointer) MoviePlayerController {
	return MoviePlayerController{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MoviePlayerController */

// Returns a object initialized with the movie at the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlayerController/init(contentURL:)
func NewMoviePlayerControllerWithContentURL(url objc.IObject /* cross-framework: NSURL */) MoviePlayerController {
	instance := getMoviePlayerControllerClass().Alloc()
	rv := objc.Send[MoviePlayerController](instance.ID, objc.Sel("initWithContentURL:"), url)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMoviePlayerControllerWithContentURL */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MoviePlayerController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MoviePlayerController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MoviePlayerController */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MoviePlayerController */

// The bounds, in points, of the content area for the full size image associated with the media item artwork.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitemartwork/imagecroprect
func (m_ MoviePlayerController) ImageCropRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](m_.ID, objc.Sel("imageCropRect"))
	return rv
}/* debug [instance_properties/getter]: imageCropRect */


// The bounds, in points, of the content area for the full size image associated with the media item artwork.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitemartwork/imagecroprect
func (m_ MoviePlayerController) SetImageCropRect(value corefoundation.CGRect) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setImageCropRect:"), value)
}/* debug [instance_properties/setter]: imageCropRect */


// Indicates whether the movie player is currently playing video via AirPlay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/isairplayvideoactive
func (m_ MoviePlayerController) IsAirPlayVideoActive() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isAirPlayVideoActive"))
	return rv
}/* debug [instance_properties/getter]: isAirPlayVideoActive */


// Indicates whether the movie player is currently playing video via AirPlay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/isairplayvideoactive
func (m_ MoviePlayerController) SetIsAirPlayVideoActive(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsAirPlayVideoActive:"), value)
}/* debug [instance_properties/setter]: isAirPlayVideoActive */


// A Boolean that indicates whether the movie player is in full-screen mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/isfullscreen
func (m_ MoviePlayerController) IsFullscreen() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isFullscreen"))
	return rv
}/* debug [instance_properties/getter]: isFullscreen */


// A Boolean that indicates whether the movie player is in full-screen mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/isfullscreen
func (m_ MoviePlayerController) SetIsFullscreen(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsFullscreen:"), value)
}/* debug [instance_properties/setter]: isFullscreen */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayerplaybackdidfinishreasonuserinfokey
func (m_ MoviePlayerController) MPMoviePlayerPlaybackDidFinishReasonUserInfoKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("MPMoviePlayerPlaybackDidFinishReasonUserInfoKey"))
	return rv
}/* debug [instance_properties/getter]: MPMoviePlayerPlaybackDidFinishReasonUserInfoKey */


// A Boolean value that indicates whether the route button is visible in the volume view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/showsroutebutton
func (m_ MoviePlayerController) ShowsRouteButton() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsRouteButton"))
	return rv
}/* debug [instance_properties/getter]: showsRouteButton */


// A Boolean value that indicates whether the route button is visible in the volume view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/showsroutebutton
func (m_ MoviePlayerController) SetShowsRouteButton(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsRouteButton:"), value)
}/* debug [instance_properties/setter]: showsRouteButton */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPMoviePlayerController */


