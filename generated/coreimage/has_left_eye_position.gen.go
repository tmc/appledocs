// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [hasLeftEyePosition] class.
var hasLeftEyePositionClass = _hasLeftEyePositionClass{objc.GetClass("hasLeftEyePosition")}

type _hasLeftEyePositionClass struct {
	class objc.Class
}

// An interface definition for the [hasLeftEyePosition] class.
type IhasLeftEyePosition interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasLeftEyePosition-c.ivar

type hasLeftEyePosition struct {
	objectivec.Object
}

// hasLeftEyePositionFrom constructs a [hasLeftEyePosition] from an unsafe.Pointer.
func hasLeftEyePositionFrom(ptr unsafe.Pointer) hasLeftEyePosition {
	return hasLeftEyePosition{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (hc _hasLeftEyePositionClass) Alloc() hasLeftEyePosition {
	rv := objc.Send[hasLeftEyePosition](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (hc _hasLeftEyePositionClass) New() hasLeftEyePosition {
	rv := objc.Send[hasLeftEyePosition](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ hasLeftEyePosition) Init() hasLeftEyePosition {
	rv := objc.Send[hasLeftEyePosition](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ hasLeftEyePosition) Autorelease() hasLeftEyePosition {
	rv := objc.Send[hasLeftEyePosition](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewhasLeftEyePosition creates a new hasLeftEyePosition instance.
func NewhasLeftEyePosition() hasLeftEyePosition {
	return hasLeftEyePositionClass.New()
}




