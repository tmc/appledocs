// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [bounds] class.
var (
	boundsClass     _boundsClass
	boundsClassOnce sync.Once
)

func getboundsClass() _boundsClass {
	boundsClassOnce.Do(func() {
		boundsClass = _boundsClass{objc.GetClass("bounds")}
	})
	return boundsClass
}

type _boundsClass struct {
	class objc.Class
}

// An interface definition for the [bounds] class.
type Ibounds interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRectangleFeature/bounds-c.ivar
type bounds struct {
	objectivec.Object
}

// boundsFrom constructs a [bounds] from an unsafe.Pointer.
func boundsFrom(ptr unsafe.Pointer) bounds {
	return bounds{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (bc _boundsClass) Alloc() bounds {
	rv := objc.Send[bounds](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _boundsClass) New() bounds {
	rv := objc.Send[bounds](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ bounds) Init() bounds {
	rv := objc.Send[bounds](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ bounds) Autorelease() bounds {
	rv := objc.Send[bounds](b_.ID, objc.Sel("autorelease"))
	return rv
}

// Newbounds creates a new bounds instance.
func Newbounds() bounds {
	return getboundsClass().New()
}




