// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHChange] class.
var (
	PHChangeClass     _PHChangeClass
	PHChangeClassOnce sync.Once
)

func getPHChangeClass() _PHChangeClass {
	PHChangeClassOnce.Do(func() {
		PHChangeClass = _PHChangeClass{objc.GetClass("PHChange")}
	})
	return PHChangeClass
}

type _PHChangeClass struct {
	class objc.Class
}

// An interface definition for the [PHChange] class.
type IPHChange interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A description of a change that occurred in the photo library.
//
// Photos provides objects to notify your app of changes to the assets and collections managed by the Photos app. To receive change information, adopt the protocol and register your observer with the shared object. After Photos provides a change object, you use its methods to get a change details object. Call the or method, passing an asset or collection object you’ve previously fetched or a fetch result containing several such objects. The resulting or object describes any changes that have happened to the object or fetch result since you last fetched it.


// A description of a change that occurred in the photo library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHChange
type PHChange struct {
	objectivec.Object
}

// PHChangeFrom constructs a [PHChange] from an unsafe.Pointer.
//
// A description of a change that occurred in the photo library.
func PHChangeFrom(ptr unsafe.Pointer) PHChange {
	return PHChange{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHChangeClass) Alloc() PHChange {
	rv := objc.Send[PHChange](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHChangeClass) New() PHChange {
	rv := objc.Send[PHChange](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHChange) Init() PHChange {
	rv := objc.Send[PHChange](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHChange) Autorelease() PHChange {
	rv := objc.Send[PHChange](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHChange creates a new PHChange instance.
func NewPHChange() PHChange {
	return getPHChangeClass().New()
}




