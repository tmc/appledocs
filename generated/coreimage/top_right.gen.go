// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [topRight] class.
var topRightClass = _topRightClass{objc.GetClass("topRight")}

type _topRightClass struct {
	class objc.Class
}

// An interface definition for the [topRight] class.
type ItopRight interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRectangleFeature/topRight-c.ivar

type topRight struct {
	objectivec.Object
}

// topRightFrom constructs a [topRight] from an unsafe.Pointer.
func topRightFrom(ptr unsafe.Pointer) topRight {
	return topRight{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (tc _topRightClass) Alloc() topRight {
	rv := objc.Send[topRight](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (tc _topRightClass) New() topRight {
	rv := objc.Send[topRight](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ topRight) Init() topRight {
	rv := objc.Send[topRight](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ topRight) Autorelease() topRight {
	rv := objc.Send[topRight](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewtopRight creates a new topRight instance.
func NewtopRight() topRight {
	return topRightClass.New()
}




