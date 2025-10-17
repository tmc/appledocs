// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [OpenGLPixelBuffer] class.
var openGLPixelBufferClass = _OpenGLPixelBufferClass{objc.GetClass("NSOpenGLPixelBuffer")}

type _OpenGLPixelBufferClass struct {
	class objc.Class
}

// An interface definition for the [OpenGLPixelBuffer] class.
type IOpenGLPixelBuffer interface {
	objectivec.IObject
}

// An object that provides access to accelerated offscreen rendering. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLPixelBuffer

type OpenGLPixelBuffer struct {
	objectivec.Object
}

// OpenGLPixelBufferFrom constructs a [OpenGLPixelBuffer] from an unsafe.Pointer.
//
// An object that provides access to accelerated offscreen rendering.
func OpenGLPixelBufferFrom(ptr unsafe.Pointer) OpenGLPixelBuffer {
	return OpenGLPixelBuffer{objectivec.Object{objc.ID(ptr)}}
}



