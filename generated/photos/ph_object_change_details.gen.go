// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [PHObjectChangeDetails] class.
var (
	PHObjectChangeDetailsClass     _PHObjectChangeDetailsClass
	PHObjectChangeDetailsClassOnce sync.Once
)

func getPHObjectChangeDetailsClass() _PHObjectChangeDetailsClass {
	PHObjectChangeDetailsClassOnce.Do(func() {
		PHObjectChangeDetailsClass = _PHObjectChangeDetailsClass{objc.GetClass("PHObjectChangeDetails")}
	})
	return PHObjectChangeDetailsClass
}

type _PHObjectChangeDetailsClass struct {
	class objc.Class
}

// An interface definition for the [PHObjectChangeDetails] class.
type IPHObjectChangeDetails interface {
	objectivec.IObject
}

// A description of changes that occurred in an asset or collection object.
//
// A object provides detailed information about differences between two states of an asset or collection object—one that you previously obtained and an updated state that would result if you fetched that entity again. You observe changes by adopting the protocol and registering your observer with the shared object. When Photos notifies your observer of a change, you get change details by passing the object you’re interested in to the method. For an asset collection or collection list, a object describe changes only to the collection’s properties. If you’re instead interested in changes to the collection’s membership, fetch the collection’s contents and use the method to track changes to the fetch result.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHObjectChangeDetails
type PHObjectChangeDetails struct {
	objectivec.Object
}

// PHObjectChangeDetailsFrom constructs a [PHObjectChangeDetails] from an unsafe.Pointer.
//
// A description of changes that occurred in an asset or collection object.
func PHObjectChangeDetailsFrom(ptr unsafe.Pointer) PHObjectChangeDetails {
	return PHObjectChangeDetails{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHObjectChangeDetailsClass) Alloc() PHObjectChangeDetails {
	rv := objc.Send[PHObjectChangeDetails](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHObjectChangeDetailsClass) New() PHObjectChangeDetails {
	rv := objc.Send[PHObjectChangeDetails](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHObjectChangeDetails) Init() PHObjectChangeDetails {
	rv := objc.Send[PHObjectChangeDetails](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHObjectChangeDetails) Autorelease() PHObjectChangeDetails {
	rv := objc.Send[PHObjectChangeDetails](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHObjectChangeDetails creates a new PHObjectChangeDetails instance.
func NewPHObjectChangeDetails() PHObjectChangeDetails {
	return getPHObjectChangeDetailsClass().New()
}




