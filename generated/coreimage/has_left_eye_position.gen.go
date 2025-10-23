// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [hasLeftEyePosition] class.
var (
	HasLeftEyePositionClass     _hasLeftEyePositionClass
	HasLeftEyePositionClassOnce sync.Once
)

func gethasLeftEyePositionClass() _hasLeftEyePositionClass {
	HasLeftEyePositionClassOnce.Do(func() {
		HasLeftEyePositionClass = _hasLeftEyePositionClass{objc.GetClass("hasLeftEyePosition")}
	})
	return HasLeftEyePositionClass
}

type _hasLeftEyePositionClass struct {
	class objc.Class
}

// An interface definition for the [hasLeftEyePosition] class.
type IhasLeftEyePosition interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return gethasLeftEyePositionClass().New()
}




