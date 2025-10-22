// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [OpenGLLayer] class.
var (
	OpenGLLayerClass     _OpenGLLayerClass
	OpenGLLayerClassOnce sync.Once
)

func getOpenGLLayerClass() _OpenGLLayerClass {
	OpenGLLayerClassOnce.Do(func() {
		OpenGLLayerClass = _OpenGLLayerClass{objc.GetClass("NSOpenGLLayer")}
	})
	return OpenGLLayerClass
}

type _OpenGLLayerClass struct {
	class objc.Class
}

// An interface definition for the [OpenGLLayer] class.
type IOpenGLLayer interface {
	objectivec.IObject
	OpenGLContext() NSOpenGLContext
	SetOpenGLContext(value IOpenGLContext)
	OpenGLPixelFormat() NSOpenGLPixelFormat
	SetOpenGLPixelFormat(value NSOpenGLPixelFormat)
	View() NSView
	SetView(value IView)
}

// A subclass of that is suitable for rendering OpenGL into layers.
//
// Unlike , uses AppKit types.


// A subclass of that is suitable for rendering OpenGL into layers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLLayer

type OpenGLLayer struct {
	objectivec.Object
}

// OpenGLLayerFrom constructs a [OpenGLLayer] from an unsafe.Pointer.
//
// A subclass of that is suitable for rendering OpenGL into layers.
func OpenGLLayerFrom(ptr unsafe.Pointer) OpenGLLayer {
	return OpenGLLayer{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (oc _OpenGLLayerClass) Alloc() OpenGLLayer {
	rv := objc.Send[OpenGLLayer](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _OpenGLLayerClass) New() OpenGLLayer {
	rv := objc.Send[OpenGLLayer](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OpenGLLayer) Init() OpenGLLayer {
	rv := objc.Send[OpenGLLayer](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OpenGLLayer) Autorelease() OpenGLLayer {
	rv := objc.Send[OpenGLLayer](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOpenGLLayer creates a new OpenGLLayer instance.
func NewOpenGLLayer() OpenGLLayer {
	return getOpenGLLayerClass().New()
}



// The layer’s OpenGL context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopengllayer/openglcontext

func (o_ OpenGLLayer) OpenGLContext() NSOpenGLContext {
	rv := objc.Send[NSOpenGLContext](o_.ID, objc.Sel("openGLContext"))
	return rv
}


// The layer’s OpenGL context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopengllayer/openglcontext

func (o_ OpenGLLayer) SetOpenGLContext(value IOpenGLContext) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setOpenGLContext:"), value)
}


// Provides access to the layer’s associated OpenGL pixel format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopengllayer/openglpixelformat

func (o_ OpenGLLayer) OpenGLPixelFormat() NSOpenGLPixelFormat {
	rv := objc.Send[NSOpenGLPixelFormat](o_.ID, objc.Sel("openGLPixelFormat"))
	return rv
}


// Provides access to the layer’s associated OpenGL pixel format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopengllayer/openglpixelformat

func (o_ OpenGLLayer) SetOpenGLPixelFormat(value NSOpenGLPixelFormat) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setOpenGLPixelFormat:"), value)
}


// Returns the view associated with the layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopengllayer/view

func (o_ OpenGLLayer) View() NSView {
	rv := objc.Send[NSView](o_.ID, objc.Sel("view"))
	return rv
}


// Returns the view associated with the layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopengllayer/view

func (o_ OpenGLLayer) SetView(value IView) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setView:"), value)
}



