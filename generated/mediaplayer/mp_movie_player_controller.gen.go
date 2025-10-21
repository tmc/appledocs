// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MoviePlayerController] class.
type IMoviePlayerController interface {
	objectivec.IObject
}

// A type of movie player that manages the playback of a movie from a file or a network stream.
//
// Playback occurs in a view owned by the movie player and takes place either fullscreen or inline. You can incorporate a movie player’s view into a view hierarchy owned by your app, or use an MPMoviePlayerViewController object to manage the presentation for you. Movie players support wireless movie playback to AirPlay-enabled hardware such as Apple TV. AirPlay playback is enabled by default. To disable AirPlay in your app, set the property to . In iOS 8.0 and later, users access AirPlay compatible hardware through the Control Panel; no AirPlay control is displayed by the movie player. When you add a movie player’s view to your app’s view hierarchy, be sure to size the frame correctly, as shown here: Consider a movie player view to be an opaque structure. You can add your own custom subviews to layer content on top of the movie but you must never modify any of its existing subviews. In addition to layering content on top of a movie, you can provide custom background content by adding subviews to the view in the property. Custom subviews are supported in both inline and fullscreen playback modes but you must adjust the positions of your views when entering or exiting fullscreen mode. Use the and notifications to detect changes to and from fullscreen mode. This class supports programmatic control of movie playback, and user-based control via buttons supplied by the movie player. You can control most aspects of playback programmatically using the methods and properties of the protocol, to which this class conforms. The methods and properties of that protocol let you start and stop playback, seek forward and backward through the movie’s content, and even change the playback rate. In addition, the property of this class lets you display a set of standard system controls that allow the user to manipulate playback. You can also set the property for network-based content to start automatically. You typically specify the movie you want to play when you create a new object. However, you can also change the currently playing movie by changing the value in the property. Changing this property lets you reuse the same movie player controller object in multiple places. For performance reasons you may want to play movies as local files. Do this by first downloading them to a local directory. To facilitate the creation of video bookmarks or chapter links for a long movie, the class defines methods for generating thumbnail images at specific times within a movie. You can request a single thumbnail image using the method or request multiple thumbnail images using the method. To play a network stream whose URL requires access credentials, first create an appropriate object. Do this by calling, for example, the method, as shown here: In addition, create an appropriate object, as shown here. Make appropriate modifications for the realm you are accessing: Add the URL credential and the protection space to the object. Do this by calling, for example, the method, as shown here: With the credential and protection space information in place, you can then play the protected stream.
//
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

// Alloc allocates a new instance without initialization.
func (mc _MoviePlayerControllerClass) Alloc() MoviePlayerController {
	rv := objc.Send[MoviePlayerController](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The URL that points to the movie file.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlayerController/contentURL
func (m_ MoviePlayerController) ContentURL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("contentURL"))
	return rv
}


// SetContentURL sets the value of the contentURL property.
// The URL that points to the movie file.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlayerController/contentURL
func (m_ MoviePlayerController) SetContentURL(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setContentURL:"), value)
}

// Indicates whether the movie player is currently playing video via AirPlay.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlayerController/isAirPlayVideoActive
func (m_ MoviePlayerController) AirPlayVideoActive() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("airPlayVideoActive"))
	return rv
}

// The network load state of the movie player.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlayerController/loadState
func (m_ MoviePlayerController) LoadState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("loadState"))
	return rv
}

// The current playback state of the movie player.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlayerController/playbackState
func (m_ MoviePlayerController) PlaybackState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("playbackState"))
	return rv
}

// Determines how the movie player repeats the playback of the movie.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlayerController/repeatMode
func (m_ MoviePlayerController) RepeatMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("repeatMode"))
	return rv
}


// SetRepeatMode sets the value of the repeatMode property.
// Determines how the movie player repeats the playback of the movie.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlayerController/repeatMode
func (m_ MoviePlayerController) SetRepeatMode(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRepeatMode:"), value)
}

// A Boolean that indicates whether a movie should begin playback automatically.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlayerController/shouldAutoplay
func (m_ MoviePlayerController) ShouldAutoplay() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("shouldAutoplay"))
	return rv
}


// SetShouldAutoplay sets the value of the shouldAutoplay property.
// A Boolean that indicates whether a movie should begin playback automatically.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlayerController/shouldAutoplay
func (m_ MoviePlayerController) SetShouldAutoplay(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShouldAutoplay:"), value)
}

// Obtains the most recent time-based metadata provided by the streamed movie.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlayerController/timedMetadata
func (m_ MoviePlayerController) TimedMetadata() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("timedMetadata"))
	return rv
}



