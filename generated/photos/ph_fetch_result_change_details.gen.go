// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHFetchResultChangeDetails] class.
var (
	PHFetchResultChangeDetailsClass     _PHFetchResultChangeDetailsClass
	PHFetchResultChangeDetailsClassOnce sync.Once
)

func getPHFetchResultChangeDetailsClass() _PHFetchResultChangeDetailsClass {
	PHFetchResultChangeDetailsClassOnce.Do(func() {
		PHFetchResultChangeDetailsClass = _PHFetchResultChangeDetailsClass{objc.GetClass("PHFetchResultChangeDetails")}
	})
	return PHFetchResultChangeDetailsClass
}

type _PHFetchResultChangeDetailsClass struct {
	class objc.Class
}

// An interface definition for the [PHFetchResultChangeDetails] class.
type IPHFetchResultChangeDetails interface {
	objectivec.IObject
}

// A description of changes that occurred in the set of asset or collection objects listed in a fetch result.
//
// A object provides detailed information about the differences between two fetch results—one that you previously obtained and an updated one that would result if you performed the same fetch again. The change details object provides information useful for updating a UI that lists the contents of a fetch result, such as the indexes of added, removed, and rearranged objects.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHFetchResultChangeDetails
type PHFetchResultChangeDetails struct {
	objectivec.Object
}

// PHFetchResultChangeDetailsFrom constructs a [PHFetchResultChangeDetails] from an unsafe.Pointer.
//
// A description of changes that occurred in the set of asset or collection objects listed in a fetch result.
func PHFetchResultChangeDetailsFrom(ptr unsafe.Pointer) PHFetchResultChangeDetails {
	return PHFetchResultChangeDetails{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHFetchResultChangeDetailsClass) Alloc() PHFetchResultChangeDetails {
	rv := objc.Send[PHFetchResultChangeDetails](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHFetchResultChangeDetailsClass) New() PHFetchResultChangeDetails {
	rv := objc.Send[PHFetchResultChangeDetails](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHFetchResultChangeDetails) Init() PHFetchResultChangeDetails {
	rv := objc.Send[PHFetchResultChangeDetails](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHFetchResultChangeDetails) Autorelease() PHFetchResultChangeDetails {
	rv := objc.Send[PHFetchResultChangeDetails](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHFetchResultChangeDetails creates a new PHFetchResultChangeDetails instance.
func NewPHFetchResultChangeDetails() PHFetchResultChangeDetails {
	return getPHFetchResultChangeDetailsClass().New()
}




