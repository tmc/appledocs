// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [OpenGLLayer] class.
var openGLLayerClass = _OpenGLLayerClass{objc.GetClass("NSOpenGLLayer")}

type _OpenGLLayerClass struct {
	class objc.Class
}

// An interface definition for the [OpenGLLayer] class.
type IOpenGLLayer interface {
	IOpenGLLayer
}

// A subclass of that is suitable for rendering OpenGL into layers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLLayer

type OpenGLLayer struct {
	objectivec.Object
}

// OpenGLLayerFrom constructs a [OpenGLLayer] from an unsafe.Pointer.
//
// A subclass of that is suitable for rendering OpenGL into layers.
func OpenGLLayerFrom(ptr unsafe.Pointer) OpenGLLayer {
	return OpenGLLayer{objectivec.Object{objc.ID(ptr)}}
}



