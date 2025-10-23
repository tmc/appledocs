// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [GCControllerButtonInput] class.
var (
	GCControllerButtonInputClass     _GCControllerButtonInputClass
	GCControllerButtonInputClassOnce sync.Once
)

func getGCControllerButtonInputClass() _GCControllerButtonInputClass {
	GCControllerButtonInputClassOnce.Do(func() {
		GCControllerButtonInputClass = _GCControllerButtonInputClass{objc.GetClass("GCControllerButtonInput")}
	})
	return GCControllerButtonInputClass
}

type _GCControllerButtonInputClass struct {
	class objc.Class
}

// An interface definition for the [GCControllerButtonInput] class.
type IGCControllerButtonInput interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A parent class referenced by other GameController classes.


// A parent class referenced by other GameController classes. [Full Topic]
type GCControllerButtonInput struct {
	objectivec.Object
}

// GCControllerButtonInputFrom constructs a [GCControllerButtonInput] from an unsafe.Pointer.
//
// A parent class referenced by other GameController classes.
func GCControllerButtonInputFrom(ptr unsafe.Pointer) GCControllerButtonInput {
	return GCControllerButtonInput{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (gc _GCControllerButtonInputClass) Alloc() GCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GCControllerButtonInputClass) New() GCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCControllerButtonInput) Init() GCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCControllerButtonInput) Autorelease() GCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCControllerButtonInput creates a new GCControllerButtonInput instance.
func NewGCControllerButtonInput() GCControllerButtonInput {
	return getGCControllerButtonInputClass().New()
}




