// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [topLeft] class.
var (
	topLeftClass     _topLeftClass
	topLeftClassOnce sync.Once
)

func gettopLeftClass() _topLeftClass {
	topLeftClassOnce.Do(func() {
		topLeftClass = _topLeftClass{objc.GetClass("topLeft")}
	})
	return topLeftClass
}

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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return gettopLeftClass().New()
}




