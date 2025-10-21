// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [OpenGLLayer] class.
var (
	OpenGLLayerClass     _OpenGLLayerClass
	OpenGLLayerClassOnce sync.Once
)

func getOpenGLLayerClass() _OpenGLLayerClass {
	OpenGLLayerClassOnce.Do(func() {
		OpenGLLayerClass = _OpenGLLayerClass{objc.GetClass("CAOpenGLLayer")}
	})
	return OpenGLLayerClass
}

type _OpenGLLayerClass struct {
	class objc.Class
}

// An interface definition for the [OpenGLLayer] class.
type IOpenGLLayer interface {
	ILayer
	CanDrawInCGLContextPixelFormatForLayerTimeDisplayTime(ctx unsafe.Pointer, pf unsafe.Pointer, t unsafe.Pointer, ts unsafe.Pointer) bool
	CopyCGLContextForPixelFormat(pf unsafe.Pointer) unsafe.Pointer
	CopyCGLPixelFormatForDisplayMask(mask unsafe.Pointer) unsafe.Pointer
	DrawInCGLContextPixelFormatForLayerTimeDisplayTime(ctx unsafe.Pointer, pf unsafe.Pointer, t unsafe.Pointer, ts unsafe.Pointer)
	ReleaseCGLContext(ctx unsafe.Pointer)
	ReleaseCGLPixelFormat(pf unsafe.Pointer)
}

// A layer that provides a layer suitable for rendering OpenGL content.
//
// To provide OpenGL content you subclass and override . You can specify that the OpenGL content is static by setting the property to .
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getOpenGLLayerClass().New()
}


// Returns whether the receiver should draw OpenGL content for the specified time.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAOpenGLLayer/canDraw(inCGLContext:pixelFormat:forLayerTime:displayTime:)
func (o_ OpenGLLayer) CanDrawInCGLContextPixelFormatForLayerTimeDisplayTime(ctx unsafe.Pointer, pf unsafe.Pointer, t unsafe.Pointer, ts unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("canDrawInCGLContext:pixelFormat:forLayerTime:displayTime:"), ctx, pf, t, ts)
	return rv
}

// Returns the rendering context the receiver requires for the specified pixel format.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAOpenGLLayer/copyCGLContext(forPixelFormat:)
func (o_ OpenGLLayer) CopyCGLContextForPixelFormat(pf unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("copyCGLContextForPixelFormat:"), pf)
	return rv
}

// Returns the OpenGL pixel format suitable for rendering to the set of displays specified by the display mask.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAOpenGLLayer/copyCGLPixelFormat(forDisplayMask:)
func (o_ OpenGLLayer) CopyCGLPixelFormatForDisplayMask(mask unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("copyCGLPixelFormatForDisplayMask:"), mask)
	return rv
}

// Draws the OpenGL content for the specified time.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAOpenGLLayer/draw(inCGLContext:pixelFormat:forLayerTime:displayTime:)
func (o_ OpenGLLayer) DrawInCGLContextPixelFormatForLayerTimeDisplayTime(ctx unsafe.Pointer, pf unsafe.Pointer, t unsafe.Pointer, ts unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("drawInCGLContext:pixelFormat:forLayerTime:displayTime:"), ctx, pf, t, ts)
}

// Releases the specified rendering context.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAOpenGLLayer/releaseCGLContext(_:)
func (o_ OpenGLLayer) ReleaseCGLContext(ctx unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("releaseCGLContext:"), ctx)
}

// Releases the specified OpenGL pixel format object.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAOpenGLLayer/releaseCGLPixelFormat(_:)
func (o_ OpenGLLayer) ReleaseCGLPixelFormat(pf unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("releaseCGLPixelFormat:"), pf)
}

// The layer’s colorspace in Core Graphics.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAOpenGLLayer/colorspace
func (o_ OpenGLLayer) Colorspace() coregraphics.CGColorSpaceRef {
	rv := objc.Send[coregraphics.CGColorSpaceRef](o_.ID, objc.Sel("colorspace"))
	return rv
}


// SetColorspace sets the value of the colorspace property.
// The layer’s colorspace in Core Graphics.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAOpenGLLayer/colorspace
func (o_ OpenGLLayer) SetColorspace(value coregraphics.CGColorSpaceRef) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setColorspace:"), value)
}

// Determines when the contents of the layer are updated.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAOpenGLLayer/isAsynchronous
func (o_ OpenGLLayer) Asynchronous() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("asynchronous"))
	return rv
}


// SetAsynchronous sets the value of the asynchronous property.
// Determines when the contents of the layer are updated.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAOpenGLLayer/isAsynchronous
func (o_ OpenGLLayer) SetAsynchronous(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAsynchronous:"), value)
}

// Tells whether or not the layer supports content with extended dynamic range.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAOpenGLLayer/wantsExtendedDynamicRangeContent
func (o_ OpenGLLayer) WantsExtendedDynamicRangeContent() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("wantsExtendedDynamicRangeContent"))
	return rv
}


// SetWantsExtendedDynamicRangeContent sets the value of the wantsExtendedDynamicRangeContent property.
// Tells whether or not the layer supports content with extended dynamic range.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAOpenGLLayer/wantsExtendedDynamicRangeContent
func (o_ OpenGLLayer) SetWantsExtendedDynamicRangeContent(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setWantsExtendedDynamicRangeContent:"), value)
}

// Determines when the contents of the layer are updated.
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caopengllayer/isasynchronous
func (o_ OpenGLLayer) IsAsynchronous() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("isAsynchronous"))
	return rv
}


// SetIsAsynchronous sets the value of the isAsynchronous property.
// Determines when the contents of the layer are updated.

//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caopengllayer/isasynchronous
func (o_ OpenGLLayer) SetIsAsynchronous(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setIsAsynchronous:"), value)
}



