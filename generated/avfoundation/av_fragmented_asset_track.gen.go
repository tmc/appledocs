// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [FragmentedAssetTrack] class.
var (
	FragmentedAssetTrackClass     _FragmentedAssetTrackClass
	FragmentedAssetTrackClassOnce sync.Once
)

func getFragmentedAssetTrackClass() _FragmentedAssetTrackClass {
	FragmentedAssetTrackClassOnce.Do(func() {
		FragmentedAssetTrackClass = _FragmentedAssetTrackClass{objc.GetClass("AVFragmentedAssetTrack")}
	})
	return FragmentedAssetTrackClass
}

type _FragmentedAssetTrackClass struct {
	class objc.Class
}

// An interface definition for the [FragmentedAssetTrack] class.
type IFragmentedAssetTrack interface {
	IAssetTrack
}

// An object that provides the track-level interface to inspect a fragmented asset’s media tracks.
//
// This class subclasses . It has no methods or properties of its own.


// An object that provides the track-level interface to inspect a fragmented asset’s media tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVFragmentedAssetTrack
type FragmentedAssetTrack struct {
	AssetTrack
}

// FragmentedAssetTrackFrom constructs a [FragmentedAssetTrack] from an unsafe.Pointer.
//
// An object that provides the track-level interface to inspect a fragmented asset’s media tracks.
func FragmentedAssetTrackFrom(ptr unsafe.Pointer) FragmentedAssetTrack {
	return FragmentedAssetTrack{
		AssetTrack: AssetTrackFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (fc _FragmentedAssetTrackClass) Alloc() FragmentedAssetTrack {
	rv := objc.Send[FragmentedAssetTrack](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FragmentedAssetTrackClass) New() FragmentedAssetTrack {
	rv := objc.Send[FragmentedAssetTrack](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FragmentedAssetTrack) Init() FragmentedAssetTrack {
	rv := objc.Send[FragmentedAssetTrack](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FragmentedAssetTrack) Autorelease() FragmentedAssetTrack {
	rv := objc.Send[FragmentedAssetTrack](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFragmentedAssetTrack creates a new FragmentedAssetTrack instance.
func NewFragmentedAssetTrack() FragmentedAssetTrack {
	return getFragmentedAssetTrackClass().New()
}




