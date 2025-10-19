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

// An interface definition for the [OpenGLLayer] class.
type IOpenGLLayer interface {
	ILayer
	DrawInCGLContextPixelFormatForLayerTimeDisplayTime(ctx unsafe.Pointer, pf unsafe.Pointer, t unsafe.Pointer, ts unsafe.Pointer)
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
// Alloc allocates a new instance without initialization.
func (oc _OpenGLLayerClass) Alloc() OpenGLLayer {
	rv := objc.Send[OpenGLLayer](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (oc _OpenGLLayerClass) New() OpenGLLayer {
	rv := objc.Send[OpenGLLayer](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OpenGLLayer) Init() OpenGLLayer {
	rv := objc.Send[OpenGLLayer](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OpenGLLayer) Autorelease() OpenGLLayer {
	rv := objc.Send[OpenGLLayer](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOpenGLLayer creates a new OpenGLLayer instance.
func NewOpenGLLayer() OpenGLLayer {
	return openGLLayerClass.New()
}


// Draws the OpenGL content for the specified time. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAOpenGLLayer/draw(inCGLContext:pixelFormat:forLayerTime:displayTime:)
func (o_ OpenGLLayer) DrawInCGLContextPixelFormatForLayerTimeDisplayTime(ctx unsafe.Pointer, pf unsafe.Pointer, t unsafe.Pointer, ts unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("drawInCGLContext:pixelFormat:forLayerTime:displayTime:"), ctx, pf, t, ts)
}


