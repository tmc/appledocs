// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [topLeft] class.
var topLeftClass = _topLeftClass{objc.GetClass("topLeft")}

type _topLeftClass struct {
	class objc.Class
}

// An interface definition for the [topLeft] class.
type ItopLeft interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRectangleFeature/topLeft-c.ivar

type topLeft struct {
	objectivec.Object
}

// topLeftFrom constructs a [topLeft] from an unsafe.Pointer.
func topLeftFrom(ptr unsafe.Pointer) topLeft {
	return topLeft{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (tc _topLeftClass) Alloc() topLeft {
	rv := objc.Send[topLeft](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (tc _topLeftClass) New() topLeft {
	rv := objc.Send[topLeft](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ topLeft) Init() topLeft {
	rv := objc.Send[topLeft](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ topLeft) Autorelease() topLeft {
	rv := objc.Send[topLeft](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewtopLeft creates a new topLeft instance.
func NewtopLeft() topLeft {
	return topLeftClass.New()
}




