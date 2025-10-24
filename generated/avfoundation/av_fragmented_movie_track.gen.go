// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [FragmentedMovieTrack] class.
var (
	FragmentedMovieTrackClass     _FragmentedMovieTrackClass
	FragmentedMovieTrackClassOnce sync.Once
)

func getFragmentedMovieTrackClass() _FragmentedMovieTrackClass {
	FragmentedMovieTrackClassOnce.Do(func() {
		FragmentedMovieTrackClass = _FragmentedMovieTrackClass{objc.GetClass("AVFragmentedMovieTrack")}
	})
	return FragmentedMovieTrackClass
}

type _FragmentedMovieTrackClass struct {
	class objc.Class
}





// An interface definition for the [FragmentedMovieTrack] class.
type IFragmentedMovieTrack interface {
	IMovieTrack
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (fc _FragmentedMovieTrackClass) Alloc() FragmentedMovieTrack {
	rv := objc.Send[FragmentedMovieTrack](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FragmentedMovieTrackClass) New() FragmentedMovieTrack {
	rv := objc.Send[FragmentedMovieTrack](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FragmentedMovieTrack) Init() FragmentedMovieTrack {
	rv := objc.Send[FragmentedMovieTrack](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FragmentedMovieTrack) Autorelease() FragmentedMovieTrack {
	rv := objc.Send[FragmentedMovieTrack](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFragmentedMovieTrack creates a new FragmentedMovieTrack instance.
func NewFragmentedMovieTrack() FragmentedMovieTrack {
	return getFragmentedMovieTrackClass().New()
}





// An object that represents a track in a fragmented movie.


// An object that represents a track in a fragmented movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVFragmentedMovieTrack
type FragmentedMovieTrack struct {
	MovieTrack
}

// FragmentedMovieTrackFrom constructs a [FragmentedMovieTrack] from an unsafe.Pointer.
//
// An object that represents a track in a fragmented movie.
func FragmentedMovieTrackFrom(ptr unsafe.Pointer) FragmentedMovieTrack {
	return FragmentedMovieTrack{
		MovieTrack: MovieTrackFrom(ptr),
	}
}































