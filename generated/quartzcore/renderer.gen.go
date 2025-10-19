// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Renderer] class.
var (
	rendererClass     _RendererClass
	rendererClassOnce sync.Once
)

func getRendererClass() _RendererClass {
	rendererClassOnce.Do(func() {
		rendererClass = _RendererClass{objc.GetClass("CARenderer")}
	})
	return rendererClass
}

type _RendererClass struct {
	class objc.Class
}

// An interface definition for the [Renderer] class.
type IRenderer interface {
	objectivec.IObject
	Render()
}

// A layer that allows an application to render a layer tree into a Core OpenGL context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CARenderer
type Renderer struct {
	objectivec.Object
}

// RendererFrom constructs a [Renderer] from an unsafe.Pointer.
//
// A layer that allows an application to render a layer tree into a Core OpenGL context.
func RendererFrom(ptr unsafe.Pointer) Renderer {
	return Renderer{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _RendererClass) Alloc() Renderer {
	rv := objc.Send[Renderer](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RendererClass) New() Renderer {
	rv := objc.Send[Renderer](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ Renderer) Init() Renderer {
	rv := objc.Send[Renderer](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ Renderer) Autorelease() Renderer {
	rv := objc.Send[Renderer](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRenderer creates a new Renderer instance.
func NewRenderer() Renderer {
	return getRendererClass().New()
}


// Render the update region of the current frame to the target context. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CARenderer/render()
func (r_ Renderer) Render() {
	objc.Send[objc.ID](r_.ID, objc.Sel("render"))
}


