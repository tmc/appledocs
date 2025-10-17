// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [OpenGLPixelFormat] class.
var OpenGLPixelFormatClass objc.Class

func init() {
	OpenGLPixelFormatClass = objc.GetClass("NSOpenGLPixelFormat")
}

type OpenGLPixelFormat struct {
	objc.ID
}

func OpenGLPixelFormatFrom(ptr unsafe.Pointer) OpenGLPixelFormat {
	return OpenGLPixelFormat{
		ID: objc.ID(ptr),
	}
}



