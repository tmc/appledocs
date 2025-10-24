// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/corefoundation"
)

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

// An interface definition for the [MoviePlayerViewController] class.
type IMoviePlayerViewController interface {
	appkit.IViewController
	// properties:
	ImageCropRect() objc.IObject /* cross-framework: Rect */
	SetImageCropRect(value objc.IObject /* cross-framework: Rect */)
	MoviePlayer() IMPMoviePlayerController
	SetMoviePlayer(value IMPMoviePlayerController)
	ShowsRouteButton() bool
	SetShowsRouteButton(value bool)
	// methods:
}

// A simple view controller for displaying full-screen movies.
//
// Unlike using an object on its own to present a movie immediately, you can incorporate a movie player view controller wherever you would normally use a view controller. For example, you can present it using a tab bar or navigation bar-based interface, taking advantage of the transitions offered by those interfaces. To present a movie player view controller modally, you typically use the method. This method is part of a category on the class and is implemented by the Media Player framework. The method presents a movie player view controller using the standard transition animations for presenting video content. To dismiss a modally presented movie player view controller, call the method.


// A simple view controller for displaying full-screen movies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlayerViewController
type MoviePlayerViewController struct {
	appkit.ViewController
}

// MoviePlayerViewControllerFrom constructs a [MoviePlayerViewController] from an unsafe.Pointer.
//
// A simple view controller for displaying full-screen movies.
func MoviePlayerViewControllerFrom(ptr unsafe.Pointer) MoviePlayerViewController {
	return MoviePlayerViewController{
		ViewController: appkit.ViewControllerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MoviePlayerViewControllerClass) Alloc() MoviePlayerViewController {
	rv := objc.Send[MoviePlayerViewController](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The bounds, in points, of the content area for the full size image associated with the media item artwork.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitemartwork/imagecroprect
func (m_ MoviePlayerViewController) ImageCropRect() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](m_.ID, objc.Sel("imageCropRect"))
	return rv
}


// The bounds, in points, of the content area for the full size image associated with the media item artwork.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitemartwork/imagecroprect
func (m_ MoviePlayerViewController) SetImageCropRect(value objc.IObject /* cross-framework: Rect */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setImageCropRect:"), value)
}


// The movie player controller object used to present the movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayerviewcontroller/movieplayer
func (m_ MoviePlayerViewController) MoviePlayer() IMPMoviePlayerController {
	rv := objc.Send[MoviePlayerController](m_.ID, objc.Sel("moviePlayer"))
	return rv
}


// The movie player controller object used to present the movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayerviewcontroller/movieplayer
func (m_ MoviePlayerViewController) SetMoviePlayer(value IMPMoviePlayerController) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMoviePlayer:"), value)
}


// A Boolean value that indicates whether the route button is visible in the volume view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/showsroutebutton
func (m_ MoviePlayerViewController) ShowsRouteButton() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsRouteButton"))
	return rv
}


// A Boolean value that indicates whether the route button is visible in the volume view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/showsroutebutton
func (m_ MoviePlayerViewController) SetShowsRouteButton(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsRouteButton:"), value)
}



