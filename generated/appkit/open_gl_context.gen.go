// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [OpenGLContext] class.
var openGLContextClass = _OpenGLContextClass{objc.GetClass("NSOpenGLContext")}

type _OpenGLContextClass struct {
	class objc.Class
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



