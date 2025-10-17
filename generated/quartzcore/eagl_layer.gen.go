// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [EAGLLayer] class.
var eAGLLayerClass = _EAGLLayerClass{objc.GetClass("CAEAGLLayer")}

type _EAGLLayerClass struct {
	class objc.Class
}

// A layer that supports drawing OpenGL content in iOS and tvOS applications. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEAGLLayer

type EAGLLayer struct {
	Layer
}

// EAGLLayerFrom constructs a [EAGLLayer] from an unsafe.Pointer.
//
// A layer that supports drawing OpenGL content in iOS and tvOS applications.
func EAGLLayerFrom(ptr unsafe.Pointer) EAGLLayer {
	return EAGLLayer{
		Layer: LayerFrom(ptr),
	}
}



