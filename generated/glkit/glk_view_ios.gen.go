//go:build darwin && ios

// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/corefoundation"
)

// iOS-only methods for GLKView


// iOS-only properties

// The OpenGL ES context used when drawing the view’s contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKView/context
func (g_ GLKView) Context() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("context"))
	return rv
}
func (g_ GLKView) SetContext(value unsafe.Pointer) {
	g_.ID.Send(objc.RegisterName("setContext:"), value)
}

// The view’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKView/delegate
func (g_ GLKView) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("delegate"))
	return rv
}
func (g_ GLKView) SetDelegate(value unsafe.Pointer) {
	g_.ID.Send(objc.RegisterName("setDelegate:"), value)
}

// The format of the color renderbuffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKView/drawableColorFormat
func (g_ GLKView) DrawableColorFormat() GLKViewDrawableColorFormat {
	rv := objc.Send[GLKViewDrawableColorFormat](g_.ID, objc.Sel("drawableColorFormat"))
	return rv
}
func (g_ GLKView) SetDrawableColorFormat(value GLKViewDrawableColorFormat) {
	g_.ID.Send(objc.RegisterName("setDrawableColorFormat:"), value)
}

// The format of the depth renderbuffer
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKView/drawableDepthFormat
func (g_ GLKView) DrawableDepthFormat() GLKViewDrawableDepthFormat {
	rv := objc.Send[GLKViewDrawableDepthFormat](g_.ID, objc.Sel("drawableDepthFormat"))
	return rv
}
func (g_ GLKView) SetDrawableDepthFormat(value GLKViewDrawableDepthFormat) {
	g_.ID.Send(objc.RegisterName("setDrawableDepthFormat:"), value)
}

// The height, in pixels, of the underlying framebuffer object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKView/drawableHeight
func (g_ GLKView) DrawableHeight() int {
	rv := objc.Send[int](g_.ID, objc.Sel("drawableHeight"))
	return rv
}

// The format of the multisampling buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKView/drawableMultisample
func (g_ GLKView) DrawableMultisample() GLKViewDrawableMultisample {
	rv := objc.Send[GLKViewDrawableMultisample](g_.ID, objc.Sel("drawableMultisample"))
	return rv
}
func (g_ GLKView) SetDrawableMultisample(value GLKViewDrawableMultisample) {
	g_.ID.Send(objc.RegisterName("setDrawableMultisample:"), value)
}

// The format of the stencil renderbuffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKView/drawableStencilFormat
func (g_ GLKView) DrawableStencilFormat() GLKViewDrawableStencilFormat {
	rv := objc.Send[GLKViewDrawableStencilFormat](g_.ID, objc.Sel("drawableStencilFormat"))
	return rv
}
func (g_ GLKView) SetDrawableStencilFormat(value GLKViewDrawableStencilFormat) {
	g_.ID.Send(objc.RegisterName("setDrawableStencilFormat:"), value)
}

// The width, in pixels, of the underlying framebuffer object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKView/drawableWidth
func (g_ GLKView) DrawableWidth() int {
	rv := objc.Send[int](g_.ID, objc.Sel("drawableWidth"))
	return rv
}

// A Boolean value that indicates whether the view responds to messages that invalidate the view’s contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKView/enableSetNeedsDisplay
func (g_ GLKView) EnableSetNeedsDisplay() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("enableSetNeedsDisplay"))
	return rv
}
func (g_ GLKView) SetEnableSetNeedsDisplay(value bool) {
	g_.ID.Send(objc.RegisterName("setEnableSetNeedsDisplay:"), value)
}

// Draws the contents of the view and returns them as a new image object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKView/snapshot
func (g_ GLKView) Snapshot() appkit.Image {
	rv := objc.Send[appkit.Image](g_.ID, objc.Sel("snapshot"))
	return rv
}




