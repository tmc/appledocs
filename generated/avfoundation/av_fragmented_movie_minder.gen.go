// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AVFragmentedMovieMinder] class.
var (
	aVFragmentedMovieMinderClass     _AVFragmentedMovieMinderClass
	aVFragmentedMovieMinderClassOnce sync.Once
)

func getAVFragmentedMovieMinderClass() _AVFragmentedMovieMinderClass {
	aVFragmentedMovieMinderClassOnce.Do(func() {
		aVFragmentedMovieMinderClass = _AVFragmentedMovieMinderClass{objc.GetClass("AVFragmentedMovieMinder")}
	})
	return aVFragmentedMovieMinderClass
}

type _AVFragmentedMovieMinderClass struct {
	class objc.Class
}

// An interface definition for the [AVFragmentedMovieMinder] class.
type IAVFragmentedMovieMinder interface {
	IAVFragmentedAssetMinder
}

// An object that checks whether a fragmented movie appends additional movie fragments.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVFragmentedMovieMinder
type AVFragmentedMovieMinder struct {
	AVFragmentedAssetMinder
}

// AVFragmentedMovieMinderFrom constructs a [AVFragmentedMovieMinder] from an unsafe.Pointer.
//
// An object that checks whether a fragmented movie appends additional movie fragments.
func AVFragmentedMovieMinderFrom(ptr unsafe.Pointer) AVFragmentedMovieMinder {
	return AVFragmentedMovieMinder{
		AVFragmentedAssetMinder: AVFragmentedAssetMinderFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AVFragmentedMovieMinderClass) Alloc() AVFragmentedMovieMinder {
	rv := objc.Send[AVFragmentedMovieMinder](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AVFragmentedMovieMinderClass) New() AVFragmentedMovieMinder {
	rv := objc.Send[AVFragmentedMovieMinder](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVFragmentedMovieMinder) Init() AVFragmentedMovieMinder {
	rv := objc.Send[AVFragmentedMovieMinder](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVFragmentedMovieMinder) Autorelease() AVFragmentedMovieMinder {
	rv := objc.Send[AVFragmentedMovieMinder](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVFragmentedMovieMinder creates a new AVFragmentedMovieMinder instance.
func NewAVFragmentedMovieMinder() AVFragmentedMovieMinder {
	return getAVFragmentedMovieMinderClass().New()
}




