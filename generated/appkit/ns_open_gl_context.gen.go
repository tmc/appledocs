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
	ClearDrawable()
	CopyAttributesFromContextWithMask(context unsafe.Pointer, mask unsafe.Pointer)
}

// An object that represents an OpenGL graphics context, into which all OpenGL calls are rendered.
//
// An OpenGL context is created using an object that specifies the context’s buffer types and other attributes. A context can be full-screen, offscreen, or associated with an object. A context draws into its , which is the frame buffer that is the target of OpenGL drawing operations. Every object wraps a low-level, platform-specific Core OpenGL (CGL) context. Your application can retrieve the CGL context by calling the method. For more information on the underling CGL context, see .
//
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

// Disassociates the OpenGL context from its viewport.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLContext/clearDrawable()
func (o_ OpenGLContext) ClearDrawable() {
	objc.Send[objc.ID](o_.ID, objc.Sel("clearDrawable"))
}

// Copies selected groups of state variables to the OpenGL context.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLContext/copyAttributesFromContext:withMask:
func (o_ OpenGLContext) CopyAttributesFromContextWithMask(context unsafe.Pointer, mask unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("copyAttributesFromContext:withMask:"), context, mask)
}

// Returns the low-level, platform-specific Core OpenGL (CGL) context object represented by the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLContext/cglContextObj
func (o_ OpenGLContext) CGLContextObj() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("CGLContextObj"))
	return rv
}

// Returns the current virtual screen for the OpenGL context.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLContext/currentVirtualScreen
func (o_ OpenGLContext) CurrentVirtualScreen() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("currentVirtualScreen"))
	return rv
}

// SetCurrentVirtualScreen sets the value of the currentVirtualScreen property.
// Returns the current virtual screen for the OpenGL context.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLContext/currentVirtualScreen
func (o_ OpenGLContext) SetCurrentVirtualScreen(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setCurrentVirtualScreen:"), value)
}
