// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [OpenGLLayer] class.
var OpenGLLayerClass objc.Class

func init() {
	OpenGLLayerClass = objc.GetClass("NSOpenGLLayer")
}

type OpenGLLayer struct {
	objc.ID
}

func OpenGLLayerFrom(ptr unsafe.Pointer) OpenGLLayer {
	return OpenGLLayer{
		ID: objc.ID(ptr),
	}
}




