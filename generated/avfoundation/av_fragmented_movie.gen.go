// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [FragmentedMovie] class.
var (
	FragmentedMovieClass     _FragmentedMovieClass
	FragmentedMovieClassOnce sync.Once
)

func getFragmentedMovieClass() _FragmentedMovieClass {
	FragmentedMovieClassOnce.Do(func() {
		FragmentedMovieClass = _FragmentedMovieClass{objc.GetClass("AVFragmentedMovie")}
	})
	return FragmentedMovieClass
}

type _FragmentedMovieClass struct {
	class objc.Class
}

// An interface definition for the [FragmentedMovie] class.
type IFragmentedMovie interface {
	IMovie
}

// An object that represents a fragmented movie file.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVFragmentedMovie
type FragmentedMovie struct {
	Movie
}

// FragmentedMovieFrom constructs a [FragmentedMovie] from an unsafe.Pointer.
//
// An object that represents a fragmented movie file.
func FragmentedMovieFrom(ptr unsafe.Pointer) FragmentedMovie {
	return FragmentedMovie{
		Movie: MovieFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (fc _FragmentedMovieClass) Alloc() FragmentedMovie {
	rv := objc.Send[FragmentedMovie](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FragmentedMovieClass) New() FragmentedMovie {
	rv := objc.Send[FragmentedMovie](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FragmentedMovie) Init() FragmentedMovie {
	rv := objc.Send[FragmentedMovie](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FragmentedMovie) Autorelease() FragmentedMovie {
	rv := objc.Send[FragmentedMovie](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFragmentedMovie creates a new FragmentedMovie instance.
func NewFragmentedMovie() FragmentedMovie {
	return getFragmentedMovieClass().New()
}




