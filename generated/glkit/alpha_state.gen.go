// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [alphaState] class.
var (
	AlphaStateClass     _alphaStateClass
	AlphaStateClassOnce sync.Once
)

func getalphaStateClass() _alphaStateClass {
	AlphaStateClassOnce.Do(func() {
		AlphaStateClass = _alphaStateClass{objc.GetClass("alphaState")}
	})
	return AlphaStateClass
}

type _alphaStateClass struct {
	class objc.Class
}

// An interface definition for the [alphaState] class.
type IalphaState interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfo/alphaState-c.ivar
type alphaState struct {
	objectivec.Object
}

// alphaStateFrom constructs a [alphaState] from an unsafe.Pointer.
func alphaStateFrom(ptr unsafe.Pointer) alphaState {
	return alphaState{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _alphaStateClass) Alloc() alphaState {
	rv := objc.Send[alphaState](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _alphaStateClass) New() alphaState {
	rv := objc.Send[alphaState](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ alphaState) Init() alphaState {
	rv := objc.Send[alphaState](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ alphaState) Autorelease() alphaState {
	rv := objc.Send[alphaState](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewalphaState creates a new alphaState instance.
func NewalphaState() alphaState {
	return getalphaStateClass().New()
}




