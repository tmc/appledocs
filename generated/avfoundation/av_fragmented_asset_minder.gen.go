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
	Assets() []Asset
	MindingInterval() float64
	SetMindingInterval(value float64)


	

	// methods:
	AddFragmentedAsset(asset unsafe.Pointer)
	RemoveFragmentedAsset(asset unsafe.Pointer)


}





// Alloc allocates a new instance without initialization.
func (fc _FragmentedAssetMinderClass) Alloc() FragmentedAssetMinder {
	rv := objc.Send[FragmentedAssetMinder](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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






// Creates a fragmented asset minder that monitors the specified asset at the indicated minding interval.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVFragmentedAssetMinder/init(asset:mindingInterval:)
func NewFragmentedAssetMinderWithAssetMindingInterval(asset unsafe.Pointer, mindingInterval float64) FragmentedAssetMinder {
	instance := getFragmentedAssetMinderClass().Alloc()
	rv := objc.Send[FragmentedAssetMinder](instance.ID, objc.Sel("initWithAsset:mindingInterval:"), asset, mindingInterval)
	rv.Autorelease()
	return rv
}







// Creates a fragmented asset minder containing the specified asset and minding interval.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVFragmentedAssetMinder/fragmentedAssetMinderWithAsset:mindingInterval:
func (fc _FragmentedAssetMinderClass) FragmentedAssetMinderWithAssetMindingInterval(asset unsafe.Pointer, mindingInterval float64) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(fc.class), objc.Sel("fragmentedAssetMinderWithAsset:mindingInterval:"), asset, mindingInterval)
	return rv
}












// Adds a fragmented asset to the array of minded assets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVFragmentedAssetMinder/addFragmentedAsset(_:)
func (f_ FragmentedAssetMinder) AddFragmentedAsset(asset unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("addFragmentedAsset:"), asset)
}


// Removes a fragmented asset from the array of minded assets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVFragmentedAssetMinder/removeFragmentedAsset(_:)
func (f_ FragmentedAssetMinder) RemoveFragmentedAsset(asset unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("removeFragmentedAsset:"), asset)
}







// The minded array of fragmented assets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVFragmentedAssetMinder/assets
func (f_ FragmentedAssetMinder) Assets() []Asset {
	rv := objc.Send[[]Asset](f_.ID, objc.Sel("assets"))
	return rv
}


// An interval that specifies when to perform a check for additional fragments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVFragmentedAssetMinder/mindingInterval
func (f_ FragmentedAssetMinder) MindingInterval() float64 {
	rv := objc.Send[float64](f_.ID, objc.Sel("mindingInterval"))
	return rv
}


// An interval that specifies when to perform a check for additional fragments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVFragmentedAssetMinder/mindingInterval
func (f_ FragmentedAssetMinder) SetMindingInterval(value float64) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setMindingInterval:"), value)
}







