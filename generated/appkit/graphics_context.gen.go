// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [GraphicsContext] class.
var (
	graphicsContextClass     _GraphicsContextClass
	graphicsContextClassOnce sync.Once
)

func getGraphicsContextClass() _GraphicsContextClass {
	graphicsContextClassOnce.Do(func() {
		graphicsContextClass = _GraphicsContextClass{objc.GetClass("NSGraphicsContext")}
	})
	return graphicsContextClass
}

type _GraphicsContextClass struct {
	class objc.Class
}

// An interface definition for the [GraphicsContext] class.
type IGraphicsContext interface {
	objectivec.IObject
}

// An object that represents a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext

type GraphicsContext struct {
	objectivec.Object
}

// GraphicsContextFrom constructs a [GraphicsContext] from an unsafe.Pointer.
//
// An object that represents a graphics context.
func GraphicsContextFrom(ptr unsafe.Pointer) GraphicsContext {
	return GraphicsContext{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (gc _GraphicsContextClass) Alloc() GraphicsContext {
	rv := objc.Send[GraphicsContext](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GraphicsContextClass) New() GraphicsContext {
	rv := objc.Send[GraphicsContext](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GraphicsContext) Init() GraphicsContext {
	rv := objc.Send[GraphicsContext](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GraphicsContext) Autorelease() GraphicsContext {
	rv := objc.Send[GraphicsContext](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGraphicsContext creates a new GraphicsContext instance.
func NewGraphicsContext() GraphicsContext {
	return getGraphicsContextClass().New()
}




