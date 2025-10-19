// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [bottomRight] class.
var (
	bottomRightClass     _bottomRightClass
	bottomRightClassOnce sync.Once
)

func getbottomRightClass() _bottomRightClass {
	bottomRightClassOnce.Do(func() {
		bottomRightClass = _bottomRightClass{objc.GetClass("bottomRight")}
	})
	return bottomRightClass
}

type _bottomRightClass struct {
	class objc.Class
}

// An interface definition for the [bottomRight] class.
type IbottomRight interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRectangleFeature/bottomRight-c.ivar
type bottomRight struct {
	objectivec.Object
}

// bottomRightFrom constructs a [bottomRight] from an unsafe.Pointer.
func bottomRightFrom(ptr unsafe.Pointer) bottomRight {
	return bottomRight{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (bc _bottomRightClass) Alloc() bottomRight {
	rv := objc.Send[bottomRight](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _bottomRightClass) New() bottomRight {
	rv := objc.Send[bottomRight](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ bottomRight) Init() bottomRight {
	rv := objc.Send[bottomRight](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ bottomRight) Autorelease() bottomRight {
	rv := objc.Send[bottomRight](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewbottomRight creates a new bottomRight instance.
func NewbottomRight() bottomRight {
	return getbottomRightClass().New()
}




