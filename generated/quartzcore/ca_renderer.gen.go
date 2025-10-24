// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Renderer] class.
var (
	RendererClass     _RendererClass
	RendererClassOnce sync.Once
)

func getRendererClass() _RendererClass {
	RendererClassOnce.Do(func() {
		RendererClass = _RendererClass{objc.GetClass("CARenderer")}
	})
	return RendererClass
}

type _RendererClass struct {
	class objc.Class
}

// An interface definition for the [Renderer] class.
type IRenderer interface {
	objectivec.IObject
	// properties:
	Bounds() objc.IObject /* cross-framework: Rect */
	SetBounds(value objc.IObject /* cross-framework: Rect */)
	Layer() ILayer
	SetLayer(value ILayer)
	// methods:
	Render()
}

// A layer that allows an application to render a layer tree into a Core OpenGL context.
//
// For real-time output you should use an instance of to host the layer-tree.


// A layer that allows an application to render a layer tree into a Core OpenGL context.
//
// [Full Topic]
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



// Render the update region of the current frame to the target context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CARenderer/render()
func (r_ Renderer) Render() {
	objc.Send[objc.ID](r_.ID, objc.Sel("render"))
}


// The bounds of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/carenderer/bounds
func (r_ Renderer) Bounds() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](r_.ID, objc.Sel("bounds"))
	return rv
}


// The bounds of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/carenderer/bounds
func (r_ Renderer) SetBounds(value objc.IObject /* cross-framework: Rect */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setBounds:"), value)
}


// The root layer of the layer-tree the receiver should render.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/carenderer/layer
func (r_ Renderer) Layer() ILayer {
	rv := objc.Send[Layer](r_.ID, objc.Sel("layer"))
	return rv
}


// The root layer of the layer-tree the receiver should render.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/carenderer/layer
func (r_ Renderer) SetLayer(value ILayer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setLayer:"), value)
}



