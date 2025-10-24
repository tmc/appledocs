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

// The duration of the movie, measured in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlayerController/duration
func (m_ MoviePlayerController) Duration() float64 {
	rv := objc.Send[TimeInterval](m_.ID, objc.Sel("duration"))
	return rv
}

// Indicates whether the movie player is currently playing video via AirPlay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlayerController/isAirPlayVideoActive
func (m_ MoviePlayerController) AirPlayVideoActive() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("airPlayVideoActive"))
	return rv
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
func (m_ MoviePlayerController) MovieSourceType() MovieSourceType /* not a class type */ {
	rv := objc.Send[MovieSourceType](m_.ID, objc.Sel("movieSourceType"))
	return rv
}
func (m_ MoviePlayerController) SetMovieSourceType(value MovieSourceType /* not a class type */) {
	m_.ID.Send(objc.RegisterName("setMovieSourceType:"), value)
}

// The width and height of the movie frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlayerController/naturalSize
func (m_ MoviePlayerController) NaturalSize() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[corefoundation.Size](m_.ID, objc.Sel("naturalSize"))
	return rv
}

// The current playback state of the movie player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlayerController/playbackState
func (m_ MoviePlayerController) PlaybackState() MoviePlaybackState /* not a class type */ {
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

// The scaling mode to use when displaying the movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlayerController/scalingMode
func (m_ MoviePlayerController) ScalingMode() MovieScalingMode /* not a class type */ {
	rv := objc.Send[MovieScalingMode](m_.ID, objc.Sel("scalingMode"))
	return rv
}
func (m_ MoviePlayerController) SetScalingMode(value MovieScalingMode /* not a class type */) {
	m_.ID.Send(objc.RegisterName("setScalingMode:"), value)
}





