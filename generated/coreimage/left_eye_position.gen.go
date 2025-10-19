// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [leftEyePosition] class.
var (
	leftEyePositionClass     _leftEyePositionClass
	leftEyePositionClassOnce sync.Once
)

func getleftEyePositionClass() _leftEyePositionClass {
	leftEyePositionClassOnce.Do(func() {
		leftEyePositionClass = _leftEyePositionClass{objc.GetClass("leftEyePosition")}
	})
	return leftEyePositionClass
}

type _leftEyePositionClass struct {
	class objc.Class
}

// An interface definition for the [leftEyePosition] class.
type IleftEyePosition interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/leftEyePosition-c.ivar
type leftEyePosition struct {
	objectivec.Object
}

// leftEyePositionFrom constructs a [leftEyePosition] from an unsafe.Pointer.
func leftEyePositionFrom(ptr unsafe.Pointer) leftEyePosition {
	return leftEyePosition{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (lc _leftEyePositionClass) Alloc() leftEyePosition {
	rv := objc.Send[leftEyePosition](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (lc _leftEyePositionClass) New() leftEyePosition {
	rv := objc.Send[leftEyePosition](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ leftEyePosition) Init() leftEyePosition {
	rv := objc.Send[leftEyePosition](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ leftEyePosition) Autorelease() leftEyePosition {
	rv := objc.Send[leftEyePosition](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewleftEyePosition creates a new leftEyePosition instance.
func NewleftEyePosition() leftEyePosition {
	return getleftEyePositionClass().New()
}




