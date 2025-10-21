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


// A customizable view that is displayed behind the movie content.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/backgroundview
func (m_ MoviePlayerController) BackgroundView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("backgroundView"))
	return rv
}


// SetBackgroundView sets the value of the backgroundView property.
// A customizable view that is displayed behind the movie content.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/backgroundview
func (m_ MoviePlayerController) SetBackgroundView(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBackgroundView:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayerplaybackdidfinishreasonuserinfokey
func (m_ MoviePlayerController) MPMoviePlayerPlaybackDidFinishReasonUserInfoKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("MPMoviePlayerPlaybackDidFinishReasonUserInfoKey"))
	return rv
}

// A Boolean that indicates whether the first video frame of the movie is ready to be displayed.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/readyfordisplay
func (m_ MoviePlayerController) ReadyForDisplay() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("readyForDisplay"))
	return rv
}


// SetReadyForDisplay sets the value of the readyForDisplay property.
// A Boolean that indicates whether the first video frame of the movie is ready to be displayed.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/readyfordisplay
func (m_ MoviePlayerController) SetReadyForDisplay(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setReadyForDisplay:"), value)
}

// The view containing the movie content and controls.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/view
func (m_ MoviePlayerController) View() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("view"))
	return rv
}


// SetView sets the value of the view property.
// The view containing the movie content and controls.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/view
func (m_ MoviePlayerController) SetView(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setView:"), value)
}

// The types of media available in the movie.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/moviemediatypes
func (m_ MoviePlayerController) MovieMediaTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("movieMediaTypes"))
	return rv
}


// SetMovieMediaTypes sets the value of the movieMediaTypes property.
// The types of media available in the movie.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/moviemediatypes
func (m_ MoviePlayerController) SetMovieMediaTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMovieMediaTypes:"), value)
}

// The width and height of the movie frame.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/naturalsize
func (m_ MoviePlayerController) NaturalSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](m_.ID, objc.Sel("naturalSize"))
	return rv
}


// SetNaturalSize sets the value of the naturalSize property.
// The width and height of the movie frame.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/naturalsize
func (m_ MoviePlayerController) SetNaturalSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNaturalSize:"), value)
}

// The amount of currently playable content.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/playableduration
func (m_ MoviePlayerController) PlayableDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("playableDuration"))
	return rv
}


// SetPlayableDuration sets the value of the playableDuration property.
// The amount of currently playable content.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/playableduration
func (m_ MoviePlayerController) SetPlayableDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPlayableDuration:"), value)
}

// The style of the playback controls.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/controlstyle
func (m_ MoviePlayerController) ControlStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("controlStyle"))
	return rv
}


// SetControlStyle sets the value of the controlStyle property.
// The style of the playback controls.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/controlstyle
func (m_ MoviePlayerController) SetControlStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setControlStyle:"), value)
}

// The scaling mode to use when displaying the movie.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/scalingmode
func (m_ MoviePlayerController) ScalingMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("scalingMode"))
	return rv
}


// SetScalingMode sets the value of the scalingMode property.
// The scaling mode to use when displaying the movie.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/scalingmode
func (m_ MoviePlayerController) SetScalingMode(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setScalingMode:"), value)
}

// A Boolean value that indicates whether the movie player should use the app’s audio session.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/useapplicationaudiosession
func (m_ MoviePlayerController) UseApplicationAudioSession() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("useApplicationAudioSession"))
	return rv
}


// SetUseApplicationAudioSession sets the value of the useApplicationAudioSession property.
// A Boolean value that indicates whether the movie player should use the app’s audio session.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/useapplicationaudiosession
func (m_ MoviePlayerController) SetUseApplicationAudioSession(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUseApplicationAudioSession:"), value)
}

// The end time (measured in seconds) for playback of the movie.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/endplaybacktime
func (m_ MoviePlayerController) EndPlaybackTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("endPlaybackTime"))
	return rv
}


// SetEndPlaybackTime sets the value of the endPlaybackTime property.
// The end time (measured in seconds) for playback of the movie.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/endplaybacktime
func (m_ MoviePlayerController) SetEndPlaybackTime(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndPlaybackTime:"), value)
}

// A snapshot of the network playback log for the movie player if it is playing a network stream.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/accesslog
func (m_ MoviePlayerController) AccessLog() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("accessLog"))
	return rv
}


// SetAccessLog sets the value of the accessLog property.
// A snapshot of the network playback log for the movie player if it is playing a network stream.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/accesslog
func (m_ MoviePlayerController) SetAccessLog(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAccessLog:"), value)
}

// A Boolean value that indicates whether the route button is visible in the volume view.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/showsroutebutton
func (m_ MoviePlayerController) ShowsRouteButton() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsRouteButton"))
	return rv
}


// SetShowsRouteButton sets the value of the showsRouteButton property.
// A Boolean value that indicates whether the route button is visible in the volume view.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/showsroutebutton
func (m_ MoviePlayerController) SetShowsRouteButton(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsRouteButton:"), value)
}

// The bounds, in points, of the content area for the full size image associated with the media item artwork.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitemartwork/imagecroprect
func (m_ MoviePlayerController) ImageCropRect() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](m_.ID, objc.Sel("imageCropRect"))
	return rv
}


// SetImageCropRect sets the value of the imageCropRect property.
// The bounds, in points, of the content area for the full size image associated with the media item artwork.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitemartwork/imagecroprect
func (m_ MoviePlayerController) SetImageCropRect(value coregraphics.CGRect) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setImageCropRect:"), value)
}

// Specifies whether the movie player allows AirPlay movie playback.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/allowsairplay
func (m_ MoviePlayerController) AllowsAirPlay() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowsAirPlay"))
	return rv
}


// SetAllowsAirPlay sets the value of the allowsAirPlay property.
// Specifies whether the movie player allows AirPlay movie playback.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/allowsairplay
func (m_ MoviePlayerController) SetAllowsAirPlay(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllowsAirPlay:"), value)
}

// A Boolean that indicates whether the movie player is in full-screen mode.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/isfullscreen
func (m_ MoviePlayerController) IsFullscreen() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isFullscreen"))
	return rv
}


// SetIsFullscreen sets the value of the isFullscreen property.
// A Boolean that indicates whether the movie player is in full-screen mode.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/isfullscreen
func (m_ MoviePlayerController) SetIsFullscreen(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsFullscreen:"), value)
}

// Indicates whether the movie player is currently playing video via AirPlay.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/isairplayvideoactive
func (m_ MoviePlayerController) IsAirPlayVideoActive() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isAirPlayVideoActive"))
	return rv
}


// SetIsAirPlayVideoActive sets the value of the isAirPlayVideoActive property.
// Indicates whether the movie player is currently playing video via AirPlay.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/isairplayvideoactive
func (m_ MoviePlayerController) SetIsAirPlayVideoActive(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsAirPlayVideoActive:"), value)
}

// The time, specified in seconds within the video timeline, when playback should start.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/initialplaybacktime
func (m_ MoviePlayerController) InitialPlaybackTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("initialPlaybackTime"))
	return rv
}


// SetInitialPlaybackTime sets the value of the initialPlaybackTime property.
// The time, specified in seconds within the video timeline, when playback should start.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/initialplaybacktime
func (m_ MoviePlayerController) SetInitialPlaybackTime(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInitialPlaybackTime:"), value)
}

// A snapshot of the playback failure error log for the movie player if it is playing a network stream.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/errorlog
func (m_ MoviePlayerController) ErrorLog() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("errorLog"))
	return rv
}


// SetErrorLog sets the value of the errorLog property.
// A snapshot of the playback failure error log for the movie player if it is playing a network stream.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/errorlog
func (m_ MoviePlayerController) SetErrorLog(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setErrorLog:"), value)
}

// The duration of the movie, measured in seconds.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/duration
func (m_ MoviePlayerController) Duration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("duration"))
	return rv
}


// SetDuration sets the value of the duration property.
// The duration of the movie, measured in seconds.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/duration
func (m_ MoviePlayerController) SetDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDuration:"), value)
}

// The playback type of the movie.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/moviesourcetype
func (m_ MoviePlayerController) MovieSourceType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("movieSourceType"))
	return rv
}


// SetMovieSourceType sets the value of the movieSourceType property.
// The playback type of the movie.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/moviesourcetype
func (m_ MoviePlayerController) SetMovieSourceType(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMovieSourceType:"), value)
}

// The URL that points to the movie file.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlayerController/contentURL
func (m_ MoviePlayerController) ContentURL() foundation.URL {
	rv := objc.Send[foundation.URL](m_.ID, objc.Sel("contentURL"))
	return rv
}


// SetContentURL sets the value of the contentURL property.
// The URL that points to the movie file.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlayerController/contentURL
func (m_ MoviePlayerController) SetContentURL(value foundation.URL) {
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



