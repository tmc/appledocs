// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mouthPosition] class.
var (
	MouthPositionClass     _mouthPositionClass
	MouthPositionClassOnce sync.Once
)

func getmouthPositionClass() _mouthPositionClass {
	MouthPositionClassOnce.Do(func() {
		MouthPositionClass = _mouthPositionClass{objc.GetClass("mouthPosition")}
	})
	return MouthPositionClass
}

type _mouthPositionClass struct {
	class objc.Class
}

// An interface definition for the [mouthPosition] class.
type ImouthPosition interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/mouthPosition-c.ivar
type mouthPosition struct {
	objectivec.Object
}

// mouthPositionFrom constructs a [mouthPosition] from an unsafe.Pointer.
func mouthPositionFrom(ptr unsafe.Pointer) mouthPosition {
	return mouthPosition{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mouthPositionClass) Alloc() mouthPosition {
	rv := objc.Send[mouthPosition](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mouthPositionClass) New() mouthPosition {
	rv := objc.Send[mouthPosition](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mouthPosition) Init() mouthPosition {
	rv := objc.Send[mouthPosition](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mouthPosition) Autorelease() mouthPosition {
	rv := objc.Send[mouthPosition](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmouthPosition creates a new mouthPosition instance.
func NewmouthPosition() mouthPosition {
	return getmouthPositionClass().New()
}




