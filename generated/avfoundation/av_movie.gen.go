// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [Movie] class.
var (
	MovieClass     _MovieClass
	MovieClassOnce sync.Once
)

func getMovieClass() _MovieClass {
	MovieClassOnce.Do(func() {
		MovieClass = _MovieClass{objc.GetClass("AVMovie")}
	})
	return MovieClass
}

type _MovieClass struct {
	class objc.Class
}

// An interface definition for the [Movie] class.
type IMovie interface {
	IAsset
	// properties:
	CanContainMovieFragments() bool /* primitive/slice/pointer. */
	ContainsMovieFragments() bool /* primitive/slice/pointer. */
	Data() objc.IObject /* cross-framework: Data */
	SetData(value objc.IObject /* cross-framework: Data */)
	DefaultMediaDataStorage() objc.IObject /* cross-framework: MediaDataStorage */
	SetDefaultMediaDataStorage(value objc.IObject /* cross-framework: MediaDataStorage */)
	Tracks() objc.IObject /* cross-framework: MovieTrack */
	SetTracks(value objc.IObject /* cross-framework: MovieTrack */)
	Url() objc.IObject /* cross-framework: URL */
	SetUrl(value objc.IObject /* cross-framework: URL */)
	// methods:
}

// An object that represents an audiovisual container that conforms to the QuickTime movie file format or a related format like MPEG-4.
//
// supports operations involving the format-specific portions of the QuickTime movie model that doesn’t support. For instance, retrieving the movie header from an existing QuickTime movie file. You can also use to write a movie header into a new file, thereby creating a reference movie.


// An object that represents an audiovisual container that conforms to the QuickTime movie file format or a related format like MPEG-4.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMovie
type Movie struct {
	Asset
}

// MovieFrom constructs a [Movie] from an unsafe.Pointer.
//
// An object that represents an audiovisual container that conforms to the QuickTime movie file format or a related format like MPEG-4.
func MovieFrom(ptr unsafe.Pointer) Movie {
	return Movie{
		Asset: AssetFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MovieClass) Alloc() Movie {
	rv := objc.Send[Movie](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MovieClass) New() Movie {
	rv := objc.Send[Movie](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ Movie) Init() Movie {
	rv := objc.Send[Movie](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ Movie) Autorelease() Movie {
	rv := objc.Send[Movie](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMovie creates a new Movie instance.
func NewMovie() Movie {
	return getMovieClass().New()
}



// A Boolean value that indicates whether fragments can extend the movie file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMovie/canContainMovieFragments
func (m_ Movie) CanContainMovieFragments() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("canContainMovieFragments"))
	return rv
}


// A Boolean value that indicates whether at least one movie fragment extends the movie file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMovie/containsMovieFragments
func (m_ Movie) ContainsMovieFragments() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("containsMovieFragments"))
	return rv
}


// A data object that contains the movie file’s data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmovie/data
func (m_ Movie) Data() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("data"))
	return rv
}


// A data object that contains the movie file’s data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmovie/data
func (m_ Movie) SetData(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setData:"), value)
}


// The default storage container for media data added to a movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmovie/defaultmediadatastorage
func (m_ Movie) DefaultMediaDataStorage() objc.IObject /* cross-framework: MediaDataStorage */ {
	rv := objc.Send[MediaDataStorage](m_.ID, objc.Sel("defaultMediaDataStorage"))
	return rv
}


// The default storage container for media data added to a movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmovie/defaultmediadatastorage
func (m_ Movie) SetDefaultMediaDataStorage(value objc.IObject /* cross-framework: MediaDataStorage */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDefaultMediaDataStorage:"), value)
}


// The tracks that a movie contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmovie/tracks
func (m_ Movie) Tracks() objc.IObject /* cross-framework: MovieTrack */ {
	rv := objc.Send[MovieTrack](m_.ID, objc.Sel("tracks"))
	return rv
}


// The tracks that a movie contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmovie/tracks
func (m_ Movie) SetTracks(value objc.IObject /* cross-framework: MovieTrack */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTracks:"), value)
}


// A URL to a QuickTime or ISO base media file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmovie/url
func (m_ Movie) Url() objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](m_.ID, objc.Sel("url"))
	return rv
}


// A URL to a QuickTime or ISO base media file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmovie/url
func (m_ Movie) SetUrl(value objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUrl:"), value)
}



