// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AVFragmentedMovieMinder] class.
var aVFragmentedMovieMinderClass = _AVFragmentedMovieMinderClass{objc.GetClass("AVFragmentedMovieMinder")}

type _AVFragmentedMovieMinderClass struct {
	class objc.Class
}

// An object that checks whether a fragmented movie appends additional movie fragments. [Full Topic]
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



