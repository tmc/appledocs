// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FragmentedMovieMinder] class.
var (
	FragmentedMovieMinderClass     _FragmentedMovieMinderClass
	FragmentedMovieMinderClassOnce sync.Once
)

func getFragmentedMovieMinderClass() _FragmentedMovieMinderClass {
	FragmentedMovieMinderClassOnce.Do(func() {
		FragmentedMovieMinderClass = _FragmentedMovieMinderClass{objc.GetClass("AVFragmentedMovieMinder")}
	})
	return FragmentedMovieMinderClass
}

type _FragmentedMovieMinderClass struct {
	class objc.Class
}

// An interface definition for the [FragmentedMovieMinder] class.
type IFragmentedMovieMinder interface {
	objectivec.IObject
	MindingInterval() unsafe.Pointer
	SetMindingInterval(value unsafe.Pointer)
	Movies() AVFragmentedMovie
	SetMovies(value IAVFragmentedMovie)
}

// An object that checks whether a fragmented movie appends additional movie fragments.
//
// This class is identical to except that it’s capable of minding only assets of type .


// An object that checks whether a fragmented movie appends additional movie fragments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVFragmentedMovieMinder
type FragmentedMovieMinder struct {
	objectivec.Object
}

// FragmentedMovieMinderFrom constructs a [FragmentedMovieMinder] from an unsafe.Pointer.
//
// An object that checks whether a fragmented movie appends additional movie fragments.
func FragmentedMovieMinderFrom(ptr unsafe.Pointer) FragmentedMovieMinder {
	return FragmentedMovieMinder{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FragmentedMovieMinderClass) Alloc() FragmentedMovieMinder {
	rv := objc.Send[FragmentedMovieMinder](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FragmentedMovieMinderClass) New() FragmentedMovieMinder {
	rv := objc.Send[FragmentedMovieMinder](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FragmentedMovieMinder) Init() FragmentedMovieMinder {
	rv := objc.Send[FragmentedMovieMinder](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FragmentedMovieMinder) Autorelease() FragmentedMovieMinder {
	rv := objc.Send[FragmentedMovieMinder](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFragmentedMovieMinder creates a new FragmentedMovieMinder instance.
func NewFragmentedMovieMinder() FragmentedMovieMinder {
	return getFragmentedMovieMinderClass().New()
}



// The amount of time between checks for additional movie fragments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avfragmentedmovieminder/mindinginterval
func (f_ FragmentedMovieMinder) MindingInterval() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("mindingInterval"))
	return rv
}


// The amount of time between checks for additional movie fragments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avfragmentedmovieminder/mindinginterval
func (f_ FragmentedMovieMinder) SetMindingInterval(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setMindingInterval:"), value)
}


// An array containing the fragmented movie objects being minded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avfragmentedmovieminder/movies
func (f_ FragmentedMovieMinder) Movies() AVFragmentedMovie {
	rv := objc.Send[AVFragmentedMovie](f_.ID, objc.Sel("movies"))
	return rv
}


// An array containing the fragmented movie objects being minded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avfragmentedmovieminder/movies
func (f_ FragmentedMovieMinder) SetMovies(value IAVFragmentedMovie) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setMovies:"), value)
}



