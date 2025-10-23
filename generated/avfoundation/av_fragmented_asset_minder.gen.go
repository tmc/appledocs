// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FragmentedAssetMinder] class.
var (
	FragmentedAssetMinderClass     _FragmentedAssetMinderClass
	FragmentedAssetMinderClassOnce sync.Once
)

func getFragmentedAssetMinderClass() _FragmentedAssetMinderClass {
	FragmentedAssetMinderClassOnce.Do(func() {
		FragmentedAssetMinderClass = _FragmentedAssetMinderClass{objc.GetClass("AVFragmentedAssetMinder")}
	})
	return FragmentedAssetMinderClass
}

type _FragmentedAssetMinderClass struct {
	class objc.Class
}

// An interface definition for the [FragmentedAssetMinder] class.
type IFragmentedAssetMinder interface {
	objectivec.IObject
	// properties:
	Assets() FragmentMinding /* not a class type */
	SetAssets(value FragmentMinding /* not a class type */)
	MindingInterval() unsafe.Pointer
	SetMindingInterval(value unsafe.Pointer)
	// methods:
}

// An object that periodically checks whether the system adds new fragments to a fragmented asset.


// An object that periodically checks whether the system adds new fragments to a fragmented asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVFragmentedAssetMinder
type FragmentedAssetMinder struct {
	objectivec.Object
}

// FragmentedAssetMinderFrom constructs a [FragmentedAssetMinder] from an unsafe.Pointer.
//
// An object that periodically checks whether the system adds new fragments to a fragmented asset.
func FragmentedAssetMinderFrom(ptr unsafe.Pointer) FragmentedAssetMinder {
	return FragmentedAssetMinder{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FragmentedAssetMinderClass) Alloc() FragmentedAssetMinder {
	rv := objc.Send[FragmentedAssetMinder](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FragmentedAssetMinderClass) New() FragmentedAssetMinder {
	rv := objc.Send[FragmentedAssetMinder](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FragmentedAssetMinder) Init() FragmentedAssetMinder {
	rv := objc.Send[FragmentedAssetMinder](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FragmentedAssetMinder) Autorelease() FragmentedAssetMinder {
	rv := objc.Send[FragmentedAssetMinder](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFragmentedAssetMinder creates a new FragmentedAssetMinder instance.
func NewFragmentedAssetMinder() FragmentedAssetMinder {
	return getFragmentedAssetMinderClass().New()
}



// The minded array of fragmented assets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avfragmentedassetminder/assets
func (f_ FragmentedAssetMinder) Assets() FragmentMinding /* not a class type */ {
	rv := objc.Send[FragmentMinding](f_.ID, objc.Sel("assets"))
	return rv
}


// The minded array of fragmented assets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avfragmentedassetminder/assets
func (f_ FragmentedAssetMinder) SetAssets(value FragmentMinding /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setAssets:"), value)
}


// An interval that specifies when to perform a check for additional fragments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avfragmentedassetminder/mindinginterval
func (f_ FragmentedAssetMinder) MindingInterval() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("mindingInterval"))
	return rv
}


// An interval that specifies when to perform a check for additional fragments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avfragmentedassetminder/mindinginterval
func (f_ FragmentedAssetMinder) SetMindingInterval(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setMindingInterval:"), value)
}



