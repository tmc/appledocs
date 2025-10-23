// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [GCGameControllerActivationContext] class.
var (
	GCGameControllerActivationContextClass     _GCGameControllerActivationContextClass
	GCGameControllerActivationContextClassOnce sync.Once
)

func getGCGameControllerActivationContextClass() _GCGameControllerActivationContextClass {
	GCGameControllerActivationContextClassOnce.Do(func() {
		GCGameControllerActivationContextClass = _GCGameControllerActivationContextClass{objc.GetClass("GCGameControllerActivationContext")}
	})
	return GCGameControllerActivationContextClass
}

type _GCGameControllerActivationContextClass struct {
	class objc.Class
}

// An interface definition for the [GCGameControllerActivationContext] class.
type IGCGameControllerActivationContext interface {
	objectivec.IObject
	// properties:
	PreviousApplicationBundleID() string /* primitive/slice/pointer. */
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCGameControllerActivationContext
type GCGameControllerActivationContext struct {
	objectivec.Object
}

// GCGameControllerActivationContextFrom constructs a [GCGameControllerActivationContext] from an unsafe.Pointer.
func GCGameControllerActivationContextFrom(ptr unsafe.Pointer) GCGameControllerActivationContext {
	return GCGameControllerActivationContext{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (gc _GCGameControllerActivationContextClass) Alloc() GCGameControllerActivationContext {
	rv := objc.Send[GCGameControllerActivationContext](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GCGameControllerActivationContextClass) New() GCGameControllerActivationContext {
	rv := objc.Send[GCGameControllerActivationContext](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCGameControllerActivationContext) Init() GCGameControllerActivationContext {
	rv := objc.Send[GCGameControllerActivationContext](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCGameControllerActivationContext) Autorelease() GCGameControllerActivationContext {
	rv := objc.Send[GCGameControllerActivationContext](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCGameControllerActivationContext creates a new GCGameControllerActivationContext instance.
func NewGCGameControllerActivationContext() GCGameControllerActivationContext {
	return getGCGameControllerActivationContextClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCGameControllerActivationContext/previousApplicationBundleID
func (g_ GCGameControllerActivationContext) PreviousApplicationBundleID() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](g_.ID, objc.Sel("previousApplicationBundleID"))
	return rv
}



