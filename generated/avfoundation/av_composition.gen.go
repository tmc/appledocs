// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Composition] class.
var (
	CompositionClass     _CompositionClass
	CompositionClassOnce sync.Once
)

func getCompositionClass() _CompositionClass {
	CompositionClassOnce.Do(func() {
		CompositionClass = _CompositionClass{objc.GetClass("AVComposition")}
	})
	return CompositionClass
}

type _CompositionClass struct {
	class objc.Class
}

// An interface definition for the [Composition] class.
type IComposition interface {
	objectivec.IObject
}

// A parent class referenced by other AVFoundation classes.


// A parent class referenced by other AVFoundation classes. [Full Topic]
type Composition struct {
	objectivec.Object
}

// CompositionFrom constructs a [Composition] from an unsafe.Pointer.
//
// A parent class referenced by other AVFoundation classes.
func CompositionFrom(ptr unsafe.Pointer) Composition {
	return Composition{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CompositionClass) Alloc() Composition {
	rv := objc.Send[Composition](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CompositionClass) New() Composition {
	rv := objc.Send[Composition](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ Composition) Init() Composition {
	rv := objc.Send[Composition](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ Composition) Autorelease() Composition {
	rv := objc.Send[Composition](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewComposition creates a new Composition instance.
func NewComposition() Composition {
	return getCompositionClass().New()
}




