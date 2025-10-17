// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [OpenGLPixelBuffer] class.
var OpenGLPixelBufferClass objc.Class

func init() {
	OpenGLPixelBufferClass = objc.GetClass("NSOpenGLPixelBuffer")
}

type OpenGLPixelBuffer struct {
	objc.ID
}

func OpenGLPixelBufferFrom(ptr unsafe.Pointer) OpenGLPixelBuffer {
	return OpenGLPixelBuffer{
		ID: objc.ID(ptr),
	}
}




