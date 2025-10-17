// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [OpenGLLayer] class.
var openGLLayerClass = _OpenGLLayerClass{objc.GetClass("CAOpenGLLayer")}

type _OpenGLLayerClass struct {
	class objc.Class
}

// A layer that provides a layer suitable for rendering OpenGL content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAOpenGLLayer

type OpenGLLayer struct {
	Layer
}

// OpenGLLayerFrom constructs a [OpenGLLayer] from an unsafe.Pointer.
//
// A layer that provides a layer suitable for rendering OpenGL content.
func OpenGLLayerFrom(ptr unsafe.Pointer) OpenGLLayer {
	return OpenGLLayer{
		Layer: LayerFrom(ptr),
	}
}

// Draws the OpenGL content for the specified time. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAOpenGLLayer/draw(inCGLContext:pixelFormat:forLayerTime:displayTime:)
func (o_ OpenGLLayer) DrawInCGLContextPixelFormatForLayerTimeDisplayTime(ctx unsafe.Pointer, pf unsafe.Pointer, t unsafe.Pointer, ts unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("drawInCGLContext:pixelFormat:forLayerTime:displayTime:"), ctx, pf, t, ts)
}


