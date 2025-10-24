//go:build darwin && ios

// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for MoviePlayerViewController


// iOS-only properties

// The movie player controller object used to present the movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMoviePlayerViewController/moviePlayer
func (m_ MoviePlayerViewController) MoviePlayer() IMPMoviePlayerController {
	rv := objc.Send[MoviePlayerController](m_.ID, objc.Sel("moviePlayer"))
	return rv
}




