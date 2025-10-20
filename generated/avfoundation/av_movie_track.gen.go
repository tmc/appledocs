// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MovieTrack] class.
var (
	MovieTrackClass     _MovieTrackClass
	MovieTrackClassOnce sync.Once
)

func getMovieTrackClass() _MovieTrackClass {
	MovieTrackClassOnce.Do(func() {
		MovieTrackClass = _MovieTrackClass{objc.GetClass("AVMovieTrack")}
	})
	return MovieTrackClass
}

type _MovieTrackClass struct {
	class objc.Class
}

// An interface definition for the [MovieTrack] class.
type IMovieTrack interface {
	IAssetTrack
}

// A track in a movie that conforms to the QuickTime or ISO base media file format.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMovieTrack
type MovieTrack struct {
	AssetTrack
}

// MovieTrackFrom constructs a [MovieTrack] from an unsafe.Pointer.
//
// A track in a movie that conforms to the QuickTime or ISO base media file format.
func MovieTrackFrom(ptr unsafe.Pointer) MovieTrack {
	return MovieTrack{
		AssetTrack: AssetTrackFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MovieTrackClass) Alloc() MovieTrack {
	rv := objc.Send[MovieTrack](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MovieTrackClass) New() MovieTrack {
	rv := objc.Send[MovieTrack](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MovieTrack) Init() MovieTrack {
	rv := objc.Send[MovieTrack](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MovieTrack) Autorelease() MovieTrack {
	rv := objc.Send[MovieTrack](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMovieTrack creates a new MovieTrack instance.
func NewMovieTrack() MovieTrack {
	return getMovieTrackClass().New()
}




