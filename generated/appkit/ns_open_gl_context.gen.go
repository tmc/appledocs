// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [OpenGLContext] class.
var (
	OpenGLContextClass     _OpenGLContextClass
	OpenGLContextClassOnce sync.Once
)

func getOpenGLContextClass() _OpenGLContextClass {
	OpenGLContextClassOnce.Do(func() {
		OpenGLContextClass = _OpenGLContextClass{objc.GetClass("NSOpenGLContext")}
	})
	return OpenGLContextClass
}

type _OpenGLContextClass struct {
	class objc.Class
}

// An interface definition for the [OpenGLContext] class.
type IOpenGLContext interface {
	objectivec.IObject
	CGLContextObj() unsafe.Pointer
	CurrentVirtualScreen() unsafe.Pointer
	SetCurrentVirtualScreen(value unsafe.Pointer)
	PixelFormat() IOpenGLPixelFormat
	View() IView
	SetView(value IView)
}

// An object that represents an OpenGL graphics context, into which all OpenGL calls are rendered.
//
// An OpenGL context is created using an object that specifies the context’s buffer types and other attributes. A context can be full-screen, offscreen, or associated with an object. A context draws into its , which is the frame buffer that is the target of OpenGL drawing operations. Every object wraps a low-level, platform-specific Core OpenGL (CGL) context. Your application can retrieve the CGL context by calling the method. For more information on the underling CGL context, see .


// An object that represents an OpenGL graphics context, into which all OpenGL calls are rendered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLContext
type OpenGLContext struct {
	objectivec.Object
}

// OpenGLContextFrom constructs a [OpenGLContext] from an unsafe.Pointer.
//
// An object that represents an OpenGL graphics context, into which all OpenGL calls are rendered.
func OpenGLContextFrom(ptr unsafe.Pointer) OpenGLContext {
	return OpenGLContext{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (oc _OpenGLContextClass) Alloc() OpenGLContext {
	rv := objc.Send[OpenGLContext](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _OpenGLContextClass) New() OpenGLContext {
	rv := objc.Send[OpenGLContext](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OpenGLContext) Init() OpenGLContext {
	rv := objc.Send[OpenGLContext](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OpenGLContext) Autorelease() OpenGLContext {
	rv := objc.Send[OpenGLContext](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOpenGLContext creates a new OpenGLContext instance.
func NewOpenGLContext() OpenGLContext {
	return getOpenGLContextClass().New()
}



// Initializes and returns an OpenGL context object using an existing CGL context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLContext/init(cglContextObj:)
func NewOpenGLContextWithCGLContextObj(context unsafe.Pointer) OpenGLContext {
	instance := getOpenGLContextClass().Alloc()
	rv := objc.Send[OpenGLContext](instance.ID, objc.Sel("initWithCGLContextObj:"), context)
	rv.Autorelease()
	return rv
}


// Returns an OpenGL context object initialized with the specified pixel format information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLContext/init(format:share:)
func NewOpenGLContextWithFormatShareContext(format IOpenGLPixelFormat, share IOpenGLContext) OpenGLContext {
	instance := getOpenGLContextClass().Alloc()
	rv := objc.Send[OpenGLContext](instance.ID, objc.Sel("initWithFormat:shareContext:"), format, share)
	rv.Autorelease()
	return rv
}



// Clears the current context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLContext/clearCurrentContext()
func (oc _OpenGLContextClass) ClearCurrentContext() {
	objc.Send[objc.ID](objc.ID(oc.class), objc.Sel("clearCurrentContext"))
}


// Returns the current OpenGL graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLContext/current
func (oc _OpenGLContextClass) CurrentContext() OpenGLContext {
	rv := objc.Send[OpenGLContext](objc.ID(oc.class), objc.Sel("currentContext"))
	return rv
}

// Returns the low-level, platform-specific Core OpenGL (CGL) context object represented by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLContext/cglContextObj
func (o_ OpenGLContext) CGLContextObj() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("CGLContextObj"))
	return rv
}


// Returns the current OpenGL graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLContext/current
func (o_ OpenGLContext) CurrentContext() IOpenGLContext {
	rv := objc.Send[OpenGLContext](o_.ID, objc.Sel("currentContext"))
	return rv
}


// Returns the current virtual screen for the OpenGL context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLContext/currentVirtualScreen
func (o_ OpenGLContext) CurrentVirtualScreen() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("currentVirtualScreen"))
	return rv
}


// Returns the current virtual screen for the OpenGL context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLContext/currentVirtualScreen
func (o_ OpenGLContext) SetCurrentVirtualScreen(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setCurrentVirtualScreen:"), value)
}


// The pixel format of the OpenGL context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLContext/pixelFormat
func (o_ OpenGLContext) PixelFormat() IOpenGLPixelFormat {
	rv := objc.Send[OpenGLPixelFormat](o_.ID, objc.Sel("pixelFormat"))
	return rv
}


// Returns the OpenGL context’s view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLContext/view
func (o_ OpenGLContext) View() IView {
	rv := objc.Send[View](o_.ID, objc.Sel("view"))
	return rv
}


// Returns the OpenGL context’s view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLContext/view
func (o_ OpenGLContext) SetView(value IView) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setView:"), value)
}


