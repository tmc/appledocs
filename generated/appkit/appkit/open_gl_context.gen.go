// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [OpenGLContext] class.
var OpenGLContextClass objc.Class

func init() {
	OpenGLContextClass = objc.GetClass("NSOpenGLContext")
}

type OpenGLContext struct {
	objc.ID
}

func OpenGLContextFrom(ptr unsafe.Pointer) OpenGLContext {
	return OpenGLContext{
		ID: objc.ID(ptr),
	}
}




