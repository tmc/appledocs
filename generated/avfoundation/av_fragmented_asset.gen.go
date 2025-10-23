// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [FragmentedAsset] class.
var (
	FragmentedAssetClass     _FragmentedAssetClass
	FragmentedAssetClassOnce sync.Once
)

func getFragmentedAssetClass() _FragmentedAssetClass {
	FragmentedAssetClassOnce.Do(func() {
		FragmentedAssetClass = _FragmentedAssetClass{objc.GetClass("AVFragmentedAsset")}
	})
	return FragmentedAssetClass
}

type _FragmentedAssetClass struct {
	class objc.Class
}

// An interface definition for the [FragmentedAsset] class.
type IFragmentedAsset interface {
	IURLAsset
	// properties:
	CanContainFragments() bool /* primitive/slice/pointer. */
	SetCanContainFragments(value bool /* primitive/slice/pointer. */)
	Tracks() objc.IObject /* cross-framework: FragmentedAssetTrack */
	SetTracks(value objc.IObject /* cross-framework: FragmentedAssetTrack */)
	// methods:
}

// An asset with a duration that the system can extend without modifying its existing media data.
//
// By using an box in their box, QuickTime movie files and MPEG-4 files can indicate that they accommodate additional fragments. To determine whether a fragmented asset can monitor the addition of fragments, check the value of its property. Associate a fragmented asset with an instance of to know when the system appends new fragments. When it has an associated asset minder, posts notifications whenever it detects new fragments. It may also post and , as the documentation of those notifications explains.


// An asset with a duration that the system can extend without modifying its existing media data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVFragmentedAsset
type FragmentedAsset struct {
	URLAsset
}

// FragmentedAssetFrom constructs a [FragmentedAsset] from an unsafe.Pointer.
//
// An asset with a duration that the system can extend without modifying its existing media data.
func FragmentedAssetFrom(ptr unsafe.Pointer) FragmentedAsset {
	return FragmentedAsset{
		URLAsset: URLAssetFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (fc _FragmentedAssetClass) Alloc() FragmentedAsset {
	rv := objc.Send[FragmentedAsset](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FragmentedAssetClass) New() FragmentedAsset {
	rv := objc.Send[FragmentedAsset](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FragmentedAsset) Init() FragmentedAsset {
	rv := objc.Send[FragmentedAsset](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FragmentedAsset) Autorelease() FragmentedAsset {
	rv := objc.Send[FragmentedAsset](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFragmentedAsset creates a new FragmentedAsset instance.
func NewFragmentedAsset() FragmentedAsset {
	return getFragmentedAssetClass().New()
}



// A Boolean value that indicates whether you can extend the asset by fragments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/cancontainfragments
func (f_ FragmentedAsset) CanContainFragments() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](f_.ID, objc.Sel("canContainFragments"))
	return rv
}


// A Boolean value that indicates whether you can extend the asset by fragments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/cancontainfragments
func (f_ FragmentedAsset) SetCanContainFragments(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setCanContainFragments:"), value)
}


// The tracks an asset contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avfragmentedasset/tracks
func (f_ FragmentedAsset) Tracks() objc.IObject /* cross-framework: FragmentedAssetTrack */ {
	rv := objc.Send[FragmentedAssetTrack](f_.ID, objc.Sel("tracks"))
	return rv
}


// The tracks an asset contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avfragmentedasset/tracks
func (f_ FragmentedAsset) SetTracks(value objc.IObject /* cross-framework: FragmentedAssetTrack */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setTracks:"), value)
}



