// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [hasRightEyePosition] class.
var (
	HasRightEyePositionClass     _hasRightEyePositionClass
	HasRightEyePositionClassOnce sync.Once
)

func gethasRightEyePositionClass() _hasRightEyePositionClass {
	HasRightEyePositionClassOnce.Do(func() {
		HasRightEyePositionClass = _hasRightEyePositionClass{objc.GetClass("hasRightEyePosition")}
	})
	return HasRightEyePositionClass
}

type _hasRightEyePositionClass struct {
	class objc.Class
}

// An interface definition for the [hasRightEyePosition] class.
type IhasRightEyePosition interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasRightEyePosition-c.ivar
type hasRightEyePosition struct {
	objectivec.Object
}

// hasRightEyePositionFrom constructs a [hasRightEyePosition] from an unsafe.Pointer.
func hasRightEyePositionFrom(ptr unsafe.Pointer) hasRightEyePosition {
	return hasRightEyePosition{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _hasRightEyePositionClass) Alloc() hasRightEyePosition {
	rv := objc.Send[hasRightEyePosition](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _hasRightEyePositionClass) New() hasRightEyePosition {
	rv := objc.Send[hasRightEyePosition](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ hasRightEyePosition) Init() hasRightEyePosition {
	rv := objc.Send[hasRightEyePosition](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ hasRightEyePosition) Autorelease() hasRightEyePosition {
	rv := objc.Send[hasRightEyePosition](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewhasRightEyePosition creates a new hasRightEyePosition instance.
func NewhasRightEyePosition() hasRightEyePosition {
	return gethasRightEyePositionClass().New()
}




