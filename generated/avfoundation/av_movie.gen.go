// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

// An object that represents an audiovisual container that conforms to the QuickTime movie file format or a related format like MPEG-4.
//
// supports operations involving the format-specific portions of the QuickTime movie model that doesn’t support. For instance, retrieving the movie header from an existing QuickTime movie file. You can also use to write a movie header into a new file, thereby creating a reference movie.
//
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




// Creates a movie object from a movie header stored in a QuickTime movie file of ISO base media file.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMovie/init(url:options:)
func NewMovieWithURLOptions(URL unsafe.Pointer, options unsafe.Pointer) Movie {
	instance := getMovieClass().Alloc()
	rv := objc.Send[Movie](instance.ID, objc.Sel("initWithURL:options:"), URL, options)
	rv.Autorelease()
	return rv
}


// A Boolean value that indicates whether fragments can extend the movie file.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMovie/canContainMovieFragments
func (m_ Movie) CanContainMovieFragments() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("canContainMovieFragments"))
	return rv
}

// A data object that contains the movie file’s data.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMovie/data
func (m_ Movie) Data() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("data"))
	return rv
}


