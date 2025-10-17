// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [OpenGLPixelFormat] class.
var openGLPixelFormatClass = _OpenGLPixelFormatClass{objc.GetClass("NSOpenGLPixelFormat")}

type _OpenGLPixelFormatClass struct {
	class objc.Class
}

// An object that specifies the types of buffers and other attributes of the OpenGL context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLPixelFormat

type OpenGLPixelFormat struct {
	objectivec.Object
}

// OpenGLPixelFormatFrom constructs a [OpenGLPixelFormat] from an unsafe.Pointer.
//
// An object that specifies the types of buffers and other attributes of the OpenGL context.
func OpenGLPixelFormatFrom(ptr unsafe.Pointer) OpenGLPixelFormat {
	return OpenGLPixelFormat{objectivec.Object{objc.ID(ptr)}}
}



