// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [GCControllerInputState] class.
var (
	GCControllerInputStateClass     _GCControllerInputStateClass
	GCControllerInputStateClassOnce sync.Once
)

func getGCControllerInputStateClass() _GCControllerInputStateClass {
	GCControllerInputStateClassOnce.Do(func() {
		GCControllerInputStateClass = _GCControllerInputStateClass{objc.GetClass("GCControllerInputState")}
	})
	return GCControllerInputStateClass
}

type _GCControllerInputStateClass struct {
	class objc.Class
}

// An interface definition for the [GCControllerInputState] class.
type IGCControllerInputState interface {
	objectivec.IObject
}

// A parent class referenced by other GameController classes.
type GCControllerInputState struct {
	objectivec.Object
}

// GCControllerInputStateFrom constructs a [GCControllerInputState] from an unsafe.Pointer.
//
// A parent class referenced by other GameController classes.
func GCControllerInputStateFrom(ptr unsafe.Pointer) GCControllerInputState {
	return GCControllerInputState{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (gc _GCControllerInputStateClass) Alloc() GCControllerInputState {
	rv := objc.Send[GCControllerInputState](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GCControllerInputStateClass) New() GCControllerInputState {
	rv := objc.Send[GCControllerInputState](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCControllerInputState) Init() GCControllerInputState {
	rv := objc.Send[GCControllerInputState](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCControllerInputState) Autorelease() GCControllerInputState {
	rv := objc.Send[GCControllerInputState](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCControllerInputState creates a new GCControllerInputState instance.
func NewGCControllerInputState() GCControllerInputState {
	return getGCControllerInputStateClass().New()
}




