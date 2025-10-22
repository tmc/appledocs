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
	AlternateGroupID() int
	SetAlternateGroupID(value int)
	MediaDataStorage() AVMediaDataStorage
	SetMediaDataStorage(value IAVMediaDataStorage)
	MediaDecodeTimeRange() unsafe.Pointer
	SetMediaDecodeTimeRange(value unsafe.Pointer)
	MediaPresentationTimeRange() unsafe.Pointer
	SetMediaPresentationTimeRange(value unsafe.Pointer)
}

// A track in a movie that conforms to the QuickTime or ISO base media file format.


// A track in a movie that conforms to the QuickTime or ISO base media file format.
//
// [Full Topic]
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



// A value that identifies the track as a member of a particular alternate group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmovietrack/alternategroupid

func (m_ MovieTrack) AlternateGroupID() int {
	rv := objc.Send[int](m_.ID, objc.Sel("alternateGroupID"))
	return rv
}


// A value that identifies the track as a member of a particular alternate group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmovietrack/alternategroupid

func (m_ MovieTrack) SetAlternateGroupID(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlternateGroupID:"), value)
}


// The storage container for media data added to a track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmovietrack/mediadatastorage

func (m_ MovieTrack) MediaDataStorage() AVMediaDataStorage {
	rv := objc.Send[AVMediaDataStorage](m_.ID, objc.Sel("mediaDataStorage"))
	return rv
}


// The storage container for media data added to a track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmovietrack/mediadatastorage

func (m_ MovieTrack) SetMediaDataStorage(value IAVMediaDataStorage) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMediaDataStorage:"), value)
}


// A range of decode times for the track’s media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmovietrack/mediadecodetimerange

func (m_ MovieTrack) MediaDecodeTimeRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("mediaDecodeTimeRange"))
	return rv
}


// A range of decode times for the track’s media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmovietrack/mediadecodetimerange

func (m_ MovieTrack) SetMediaDecodeTimeRange(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMediaDecodeTimeRange:"), value)
}


// A range of presentation times for the track’s media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmovietrack/mediapresentationtimerange

func (m_ MovieTrack) MediaPresentationTimeRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("mediaPresentationTimeRange"))
	return rv
}


// A range of presentation times for the track’s media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmovietrack/mediapresentationtimerange

func (m_ MovieTrack) SetMediaPresentationTimeRange(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMediaPresentationTimeRange:"), value)
}



