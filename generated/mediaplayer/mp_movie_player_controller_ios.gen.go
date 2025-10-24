//go:build darwin && ios

// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for MoviePlayerController


// iOS-only properties

// A snapshot of the network playback log for the movie player if it is playing a network stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlayerController/accessLog
func (m_ MoviePlayerController) AccessLog() IMPMovieAccessLog {
	rv := objc.Send[MovieAccessLog](m_.ID, objc.Sel("accessLog"))
	return rv
}

// Specifies whether the movie player allows AirPlay movie playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlayerController/allowsAirPlay
func (m_ MoviePlayerController) AllowsAirPlay() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowsAirPlay"))
	return rv
}
func (m_ MoviePlayerController) SetAllowsAirPlay(value bool) {
	m_.ID.Send(objc.RegisterName("setAllowsAirPlay:"), value)
}

// A customizable view that is displayed behind the movie content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlayerController/backgroundView
func (m_ MoviePlayerController) BackgroundView() appkit.View {
	rv := objc.Send[appkit.View](m_.ID, objc.Sel("backgroundView"))
	return rv
}

// The URL that points to the movie file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlayerController/contentURL
func (m_ MoviePlayerController) ContentURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](m_.ID, objc.Sel("contentURL"))
	return rv
}
func (m_ MoviePlayerController) SetContentURL(value objc.IObject /* cross-framework: NSURL */) {
	m_.ID.Send(objc.RegisterName("setContentURL:"), value)
}

// The style of the playback controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlayerController/controlStyle
func (m_ MoviePlayerController) ControlStyle() MovieControlStyle {
	rv := objc.Send[MovieControlStyle](m_.ID, objc.Sel("controlStyle"))
	return rv
}
func (m_ MoviePlayerController) SetControlStyle(value MovieControlStyle) {
	m_.ID.Send(objc.RegisterName("setControlStyle:"), value)
}

// The duration of the movie, measured in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlayerController/duration
func (m_ MoviePlayerController) Duration() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("duration"))
	return rv
}

// The end time (measured in seconds) for playback of the movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlayerController/endPlaybackTime
func (m_ MoviePlayerController) EndPlaybackTime() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("endPlaybackTime"))
	return rv
}
func (m_ MoviePlayerController) SetEndPlaybackTime(value float64) {
	m_.ID.Send(objc.RegisterName("setEndPlaybackTime:"), value)
}

// A snapshot of the playback failure error log for the movie player if it is playing a network stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlayerController/errorLog
func (m_ MoviePlayerController) ErrorLog() IMPMovieErrorLog {
	rv := objc.Send[MovieErrorLog](m_.ID, objc.Sel("errorLog"))
	return rv
}

// The time, specified in seconds within the video timeline, when playback should start.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlayerController/initialPlaybackTime
func (m_ MoviePlayerController) InitialPlaybackTime() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("initialPlaybackTime"))
	return rv
}
func (m_ MoviePlayerController) SetInitialPlaybackTime(value float64) {
	m_.ID.Send(objc.RegisterName("setInitialPlaybackTime:"), value)
}

// Indicates whether the movie player is currently playing video via AirPlay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlayerController/isAirPlayVideoActive
func (m_ MoviePlayerController) AirPlayVideoActive() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("airPlayVideoActive"))
	return rv
}

// A Boolean that indicates whether the movie player is in full-screen mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlayerController/isFullscreen
func (m_ MoviePlayerController) Fullscreen() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("fullscreen"))
	return rv
}
func (m_ MoviePlayerController) SetFullscreen(value bool) {
	m_.ID.Send(objc.RegisterName("setFullscreen:"), value)
}

// The network load state of the movie player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlayerController/loadState
func (m_ MoviePlayerController) LoadState() MovieLoadState {
	rv := objc.Send[MovieLoadState](m_.ID, objc.Sel("loadState"))
	return rv
}

// The types of media available in the movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlayerController/movieMediaTypes
func (m_ MoviePlayerController) MovieMediaTypes() MovieMediaTypeMask {
	rv := objc.Send[MovieMediaTypeMask](m_.ID, objc.Sel("movieMediaTypes"))
	return rv
}

// The playback type of the movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlayerController/movieSourceType
func (m_ MoviePlayerController) MovieSourceType() MovieSourceType {
	rv := objc.Send[MovieSourceType](m_.ID, objc.Sel("movieSourceType"))
	return rv
}
func (m_ MoviePlayerController) SetMovieSourceType(value MovieSourceType) {
	m_.ID.Send(objc.RegisterName("setMovieSourceType:"), value)
}

// The width and height of the movie frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlayerController/naturalSize
func (m_ MoviePlayerController) NaturalSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](m_.ID, objc.Sel("naturalSize"))
	return rv
}

// The amount of currently playable content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlayerController/playableDuration
func (m_ MoviePlayerController) PlayableDuration() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("playableDuration"))
	return rv
}

// The current playback state of the movie player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlayerController/playbackState
func (m_ MoviePlayerController) PlaybackState() MoviePlaybackState {
	rv := objc.Send[MoviePlaybackState](m_.ID, objc.Sel("playbackState"))
	return rv
}

// A Boolean that indicates whether the first video frame of the movie is ready to be displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlayerController/readyForDisplay
func (m_ MoviePlayerController) ReadyForDisplay() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("readyForDisplay"))
	return rv
}

// Determines how the movie player repeats the playback of the movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlayerController/repeatMode
func (m_ MoviePlayerController) RepeatMode() MovieRepeatMode {
	rv := objc.Send[MovieRepeatMode](m_.ID, objc.Sel("repeatMode"))
	return rv
}
func (m_ MoviePlayerController) SetRepeatMode(value MovieRepeatMode) {
	m_.ID.Send(objc.RegisterName("setRepeatMode:"), value)
}

// The scaling mode to use when displaying the movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlayerController/scalingMode
func (m_ MoviePlayerController) ScalingMode() MovieScalingMode {
	rv := objc.Send[MovieScalingMode](m_.ID, objc.Sel("scalingMode"))
	return rv
}
func (m_ MoviePlayerController) SetScalingMode(value MovieScalingMode) {
	m_.ID.Send(objc.RegisterName("setScalingMode:"), value)
}

// A Boolean that indicates whether a movie should begin playback automatically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlayerController/shouldAutoplay
func (m_ MoviePlayerController) ShouldAutoplay() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("shouldAutoplay"))
	return rv
}
func (m_ MoviePlayerController) SetShouldAutoplay(value bool) {
	m_.ID.Send(objc.RegisterName("setShouldAutoplay:"), value)
}

// Obtains the most recent time-based metadata provided by the streamed movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlayerController/timedMetadata
func (m_ MoviePlayerController) TimedMetadata() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("timedMetadata"))
	return rv
}

// A Boolean value that indicates whether the movie player should use the app’s audio session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlayerController/useApplicationAudioSession
func (m_ MoviePlayerController) UseApplicationAudioSession() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("useApplicationAudioSession"))
	return rv
}
func (m_ MoviePlayerController) SetUseApplicationAudioSession(value bool) {
	m_.ID.Send(objc.RegisterName("setUseApplicationAudioSession:"), value)
}

// The view containing the movie content and controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlayerController/view
func (m_ MoviePlayerController) View() appkit.View {
	rv := objc.Send[appkit.View](m_.ID, objc.Sel("view"))
	return rv
}




