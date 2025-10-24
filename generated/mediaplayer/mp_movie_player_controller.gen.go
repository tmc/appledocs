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
	// properties:
	ImageCropRect() objc.IObject /* cross-framework: Rect */
	SetImageCropRect(value objc.IObject /* cross-framework: Rect */)
	AccessLog() IMPMovieAccessLog
	SetAccessLog(value IMPMovieAccessLog)
	AllowsAirPlay() bool
	SetAllowsAirPlay(value bool)
	BackgroundView() objc.IObject /* cross-framework: View */
	SetBackgroundView(value objc.IObject /* cross-framework: View */)
	ControlStyle() MovieControlStyle /* not a class type */
	SetControlStyle(value MovieControlStyle /* not a class type */)
	EndPlaybackTime() float64
	SetEndPlaybackTime(value float64)
	ErrorLog() IMPMovieErrorLog
	SetErrorLog(value IMPMovieErrorLog)
	InitialPlaybackTime() float64
	SetInitialPlaybackTime(value float64)
	IsAirPlayVideoActive() bool
	SetIsAirPlayVideoActive(value bool)
	IsFullscreen() bool
	SetIsFullscreen(value bool)
	PlayableDuration() float64
	SetPlayableDuration(value float64)
	RepeatMode() MovieRepeatMode /* not a class type */
	SetRepeatMode(value MovieRepeatMode /* not a class type */)
	ShouldAutoplay() bool
	SetShouldAutoplay(value bool)
	TimedMetadata() unsafe.Pointer
	SetTimedMetadata(value unsafe.Pointer)
	UseApplicationAudioSession() bool
	SetUseApplicationAudioSession(value bool)
	View() objc.IObject /* cross-framework: View */
	SetView(value objc.IObject /* cross-framework: View */)
	MPMoviePlayerPlaybackDidFinishReasonUserInfoKey() objc.IObject /* cross-framework: NSString */
	ShowsRouteButton() bool
	SetShowsRouteButton(value bool)
	// methods:
}

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



// The bounds, in points, of the content area for the full size image associated with the media item artwork.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitemartwork/imagecroprect
func (m_ MoviePlayerController) ImageCropRect() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](m_.ID, objc.Sel("imageCropRect"))
	return rv
}


// The bounds, in points, of the content area for the full size image associated with the media item artwork.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitemartwork/imagecroprect
func (m_ MoviePlayerController) SetImageCropRect(value objc.IObject /* cross-framework: Rect */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setImageCropRect:"), value)
}


// A snapshot of the network playback log for the movie player if it is playing a network stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/accesslog
func (m_ MoviePlayerController) AccessLog() IMPMovieAccessLog {
	rv := objc.Send[MovieAccessLog](m_.ID, objc.Sel("accessLog"))
	return rv
}


// A snapshot of the network playback log for the movie player if it is playing a network stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/accesslog
func (m_ MoviePlayerController) SetAccessLog(value IMPMovieAccessLog) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAccessLog:"), value)
}


// Specifies whether the movie player allows AirPlay movie playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/allowsairplay
func (m_ MoviePlayerController) AllowsAirPlay() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowsAirPlay"))
	return rv
}


// Specifies whether the movie player allows AirPlay movie playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/allowsairplay
func (m_ MoviePlayerController) SetAllowsAirPlay(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllowsAirPlay:"), value)
}


// A customizable view that is displayed behind the movie content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/backgroundview
func (m_ MoviePlayerController) BackgroundView() objc.IObject /* cross-framework: View */ {
	rv := objc.Send[appkit.View](m_.ID, objc.Sel("backgroundView"))
	return rv
}


// A customizable view that is displayed behind the movie content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/backgroundview
func (m_ MoviePlayerController) SetBackgroundView(value objc.IObject /* cross-framework: View */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBackgroundView:"), value)
}


// The style of the playback controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/controlstyle
func (m_ MoviePlayerController) ControlStyle() MovieControlStyle /* not a class type */ {
	rv := objc.Send[MovieControlStyle](m_.ID, objc.Sel("controlStyle"))
	return rv
}


// The style of the playback controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/controlstyle
func (m_ MoviePlayerController) SetControlStyle(value MovieControlStyle /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setControlStyle:"), value)
}


// The end time (measured in seconds) for playback of the movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/endplaybacktime
func (m_ MoviePlayerController) EndPlaybackTime() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("endPlaybackTime"))
	return rv
}


// The end time (measured in seconds) for playback of the movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/endplaybacktime
func (m_ MoviePlayerController) SetEndPlaybackTime(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndPlaybackTime:"), value)
}


// A snapshot of the playback failure error log for the movie player if it is playing a network stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/errorlog
func (m_ MoviePlayerController) ErrorLog() IMPMovieErrorLog {
	rv := objc.Send[MovieErrorLog](m_.ID, objc.Sel("errorLog"))
	return rv
}


// A snapshot of the playback failure error log for the movie player if it is playing a network stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/errorlog
func (m_ MoviePlayerController) SetErrorLog(value IMPMovieErrorLog) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setErrorLog:"), value)
}


// The time, specified in seconds within the video timeline, when playback should start.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/initialplaybacktime
func (m_ MoviePlayerController) InitialPlaybackTime() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("initialPlaybackTime"))
	return rv
}


// The time, specified in seconds within the video timeline, when playback should start.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/initialplaybacktime
func (m_ MoviePlayerController) SetInitialPlaybackTime(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInitialPlaybackTime:"), value)
}


// Indicates whether the movie player is currently playing video via AirPlay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/isairplayvideoactive
func (m_ MoviePlayerController) IsAirPlayVideoActive() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isAirPlayVideoActive"))
	return rv
}


// Indicates whether the movie player is currently playing video via AirPlay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/isairplayvideoactive
func (m_ MoviePlayerController) SetIsAirPlayVideoActive(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsAirPlayVideoActive:"), value)
}


// A Boolean that indicates whether the movie player is in full-screen mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/isfullscreen
func (m_ MoviePlayerController) IsFullscreen() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isFullscreen"))
	return rv
}


// A Boolean that indicates whether the movie player is in full-screen mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/isfullscreen
func (m_ MoviePlayerController) SetIsFullscreen(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsFullscreen:"), value)
}


// The amount of currently playable content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/playableduration
func (m_ MoviePlayerController) PlayableDuration() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("playableDuration"))
	return rv
}


// The amount of currently playable content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/playableduration
func (m_ MoviePlayerController) SetPlayableDuration(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPlayableDuration:"), value)
}


// Determines how the movie player repeats the playback of the movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/repeatmode
func (m_ MoviePlayerController) RepeatMode() MovieRepeatMode /* not a class type */ {
	rv := objc.Send[MovieRepeatMode](m_.ID, objc.Sel("repeatMode"))
	return rv
}


// Determines how the movie player repeats the playback of the movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/repeatmode
func (m_ MoviePlayerController) SetRepeatMode(value MovieRepeatMode /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRepeatMode:"), value)
}


// A Boolean that indicates whether a movie should begin playback automatically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/shouldautoplay
func (m_ MoviePlayerController) ShouldAutoplay() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("shouldAutoplay"))
	return rv
}


// A Boolean that indicates whether a movie should begin playback automatically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/shouldautoplay
func (m_ MoviePlayerController) SetShouldAutoplay(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShouldAutoplay:"), value)
}


// Obtains the most recent time-based metadata provided by the streamed movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/timedmetadata
func (m_ MoviePlayerController) TimedMetadata() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("timedMetadata"))
	return rv
}


// Obtains the most recent time-based metadata provided by the streamed movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/timedmetadata
func (m_ MoviePlayerController) SetTimedMetadata(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedMetadata:"), value)
}


// A Boolean value that indicates whether the movie player should use the app’s audio session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/useapplicationaudiosession
func (m_ MoviePlayerController) UseApplicationAudioSession() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("useApplicationAudioSession"))
	return rv
}


// A Boolean value that indicates whether the movie player should use the app’s audio session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/useapplicationaudiosession
func (m_ MoviePlayerController) SetUseApplicationAudioSession(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUseApplicationAudioSession:"), value)
}


// The view containing the movie content and controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/view
func (m_ MoviePlayerController) View() objc.IObject /* cross-framework: View */ {
	rv := objc.Send[appkit.View](m_.ID, objc.Sel("view"))
	return rv
}


// The view containing the movie content and controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayercontroller/view
func (m_ MoviePlayerController) SetView(value objc.IObject /* cross-framework: View */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setView:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayerplaybackdidfinishreasonuserinfokey
func (m_ MoviePlayerController) MPMoviePlayerPlaybackDidFinishReasonUserInfoKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("MPMoviePlayerPlaybackDidFinishReasonUserInfoKey"))
	return rv
}


// A Boolean value that indicates whether the route button is visible in the volume view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/showsroutebutton
func (m_ MoviePlayerController) ShowsRouteButton() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsRouteButton"))
	return rv
}


// A Boolean value that indicates whether the route button is visible in the volume view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/showsroutebutton
func (m_ MoviePlayerController) SetShowsRouteButton(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsRouteButton:"), value)
}


