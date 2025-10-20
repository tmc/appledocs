// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [hasFaceAngle] class.
var (
	HasFaceAngleClass     _hasFaceAngleClass
	HasFaceAngleClassOnce sync.Once
)

func gethasFaceAngleClass() _hasFaceAngleClass {
	HasFaceAngleClassOnce.Do(func() {
		HasFaceAngleClass = _hasFaceAngleClass{objc.GetClass("hasFaceAngle")}
	})
	return HasFaceAngleClass
}

type _hasFaceAngleClass struct {
	class objc.Class
}

// An interface definition for the [hasFaceAngle] class.
type IhasFaceAngle interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasFaceAngle-c.ivar
type hasFaceAngle struct {
	objectivec.Object
}

// hasFaceAngleFrom constructs a [hasFaceAngle] from an unsafe.Pointer.
func hasFaceAngleFrom(ptr unsafe.Pointer) hasFaceAngle {
	return hasFaceAngle{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _hasFaceAngleClass) Alloc() hasFaceAngle {
	rv := objc.Send[hasFaceAngle](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _hasFaceAngleClass) New() hasFaceAngle {
	rv := objc.Send[hasFaceAngle](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ hasFaceAngle) Init() hasFaceAngle {
	rv := objc.Send[hasFaceAngle](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ hasFaceAngle) Autorelease() hasFaceAngle {
	rv := objc.Send[hasFaceAngle](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewhasFaceAngle creates a new hasFaceAngle instance.
func NewhasFaceAngle() hasFaceAngle {
	return gethasFaceAngleClass().New()
}




