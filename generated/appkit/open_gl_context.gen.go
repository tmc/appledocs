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
	openGLContextClass     _OpenGLContextClass
	openGLContextClassOnce sync.Once
)

func getOpenGLContextClass() _OpenGLContextClass {
	openGLContextClassOnce.Do(func() {
		openGLContextClass = _OpenGLContextClass{objc.GetClass("NSOpenGLContext")}
	})
	return openGLContextClass
}

type _OpenGLContextClass struct {
	class objc.Class
}

// An interface definition for the [OpenGLContext] class.
type IOpenGLContext interface {
	objectivec.IObject
}

// An object that represents an OpenGL graphics context, into which all OpenGL calls are rendered. [Full Topic]
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




