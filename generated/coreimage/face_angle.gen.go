// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [faceAngle] class.
var (
	faceAngleClass     _faceAngleClass
	faceAngleClassOnce sync.Once
)

func getfaceAngleClass() _faceAngleClass {
	faceAngleClassOnce.Do(func() {
		faceAngleClass = _faceAngleClass{objc.GetClass("faceAngle")}
	})
	return faceAngleClass
}

type _faceAngleClass struct {
	class objc.Class
}

// An interface definition for the [faceAngle] class.
type IfaceAngle interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/faceAngle-c.ivar
type faceAngle struct {
	objectivec.Object
}

// faceAngleFrom constructs a [faceAngle] from an unsafe.Pointer.
func faceAngleFrom(ptr unsafe.Pointer) faceAngle {
	return faceAngle{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _faceAngleClass) Alloc() faceAngle {
	rv := objc.Send[faceAngle](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _faceAngleClass) New() faceAngle {
	rv := objc.Send[faceAngle](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ faceAngle) Init() faceAngle {
	rv := objc.Send[faceAngle](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ faceAngle) Autorelease() faceAngle {
	rv := objc.Send[faceAngle](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewfaceAngle creates a new faceAngle instance.
func NewfaceAngle() faceAngle {
	return getfaceAngleClass().New()
}




