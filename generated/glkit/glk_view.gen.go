// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [GLKView] class.
var (
	GLKViewClass     _GLKViewClass
	GLKViewClassOnce sync.Once
)

func getGLKViewClass() _GLKViewClass {
	GLKViewClassOnce.Do(func() {
		GLKViewClass = _GLKViewClass{objc.GetClass("GLKView")}
	})
	return GLKViewClass
}

type _GLKViewClass struct {
	class objc.Class
}

// An interface definition for the [GLKView] class.
type IGLKView interface {
	appkit.IView
	BindDrawable()
	DeleteDrawable()
	Display()
}

// A default implementation for views that draw their content using OpenGL ES.
//
// The class simplifies the effort required to create an OpenGL ES application by directly managing a framebuffer object on your behalf; your application simply needs to draw into the framebuffer when the contents need to be updated. To use this class in your application, create a new object and provide it an OpenGL ES context. Then, modify the view’s , , , and properties to configure the format of the drawable’s framebuffer object. After this, the view automatically creates or updates the framebuffer object whenever the view must be redrawn. A object uses the regular view drawing cycle for a object, calling its method whenever the contents of the view need to be updated. Before calling its method, the view makes its object the current OpenGL ES context and binds its framebuffer object to the OpenGL ES context as the target for rendering commands. Your application’s implementation of the method should call one or more OpenGL ES functions to render an image into the framebuffer object. Then, the view resolves any multisampling that you may have enabled and delivers the finished results. The class can be used in conjunction with a object to create an animation rendering loop that redraws the contents of the view at a specified frame rate.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKView
type GLKView struct {
	appkit.View
}

// GLKViewFrom constructs a [GLKView] from an unsafe.Pointer.
//
// A default implementation for views that draw their content using OpenGL ES.
func GLKViewFrom(ptr unsafe.Pointer) GLKView {
	return GLKView{
		View: appkit.ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (gc _GLKViewClass) Alloc() GLKView {
	rv := objc.Send[GLKView](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GLKViewClass) New() GLKView {
	rv := objc.Send[GLKView](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GLKView) Init() GLKView {
	rv := objc.Send[GLKView](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GLKView) Autorelease() GLKView {
	rv := objc.Send[GLKView](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGLKView creates a new GLKView instance.
func NewGLKView() GLKView {
	return getGLKViewClass().New()
}




// Initializes a new view.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKView/init(frame:context:)
func NewGLKViewWithFrameContext(frame coregraphics.CGRect, context unsafe.Pointer) GLKView {
	instance := getGLKViewClass().Alloc()
	rv := objc.Send[GLKView](instance.ID, objc.Sel("initWithFrame:context:"), frame, context)
	rv.Autorelease()
	return rv
}


// Binds the underlying framebuffer object to OpenGL ES.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKView/bindDrawable()
func (g_ GLKView) BindDrawable() {
	objc.Send[objc.ID](g_.ID, objc.Sel("bindDrawable"))
}

// Deletes the drawable object associated with the view.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKView/deleteDrawable()
func (g_ GLKView) DeleteDrawable() {
	objc.Send[objc.ID](g_.ID, objc.Sel("deleteDrawable"))
}

// Redraws the view’s contents immediately.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKView/display()
func (g_ GLKView) Display() {
	objc.Send[objc.ID](g_.ID, objc.Sel("display"))
}

// The OpenGL ES context used when drawing the view’s contents.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKView/context
func (g_ GLKView) Context() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("context"))
	return rv
}


// SetContext sets the value of the context property.
// The OpenGL ES context used when drawing the view’s contents.

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKView/context
func (g_ GLKView) SetContext(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setContext:"), value)
}

// The view’s delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKView/delegate
func (g_ GLKView) Delegate() objc.ID {
	rv := objc.Send[objc.ID](g_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The view’s delegate.

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKView/delegate
func (g_ GLKView) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDelegate:"), value)
}

// The format of the color renderbuffer.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKView/drawableColorFormat
func (g_ GLKView) DrawableColorFormat() GLKViewDrawableColorFormat {
	rv := objc.Send[GLKViewDrawableColorFormat](g_.ID, objc.Sel("drawableColorFormat"))
	return rv
}


// SetDrawableColorFormat sets the value of the drawableColorFormat property.
// The format of the color renderbuffer.

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKView/drawableColorFormat
func (g_ GLKView) SetDrawableColorFormat(value GLKViewDrawableColorFormat) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDrawableColorFormat:"), value)
}

// The format of the depth renderbuffer
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKView/drawableDepthFormat
func (g_ GLKView) DrawableDepthFormat() GLKViewDrawableDepthFormat {
	rv := objc.Send[GLKViewDrawableDepthFormat](g_.ID, objc.Sel("drawableDepthFormat"))
	return rv
}


// SetDrawableDepthFormat sets the value of the drawableDepthFormat property.
// The format of the depth renderbuffer

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKView/drawableDepthFormat
func (g_ GLKView) SetDrawableDepthFormat(value GLKViewDrawableDepthFormat) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDrawableDepthFormat:"), value)
}

// The height, in pixels, of the underlying framebuffer object.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKView/drawableHeight
func (g_ GLKView) DrawableHeight() int {
	rv := objc.Send[int](g_.ID, objc.Sel("drawableHeight"))
	return rv
}

// The format of the multisampling buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKView/drawableMultisample
func (g_ GLKView) DrawableMultisample() GLKViewDrawableMultisample {
	rv := objc.Send[GLKViewDrawableMultisample](g_.ID, objc.Sel("drawableMultisample"))
	return rv
}


// SetDrawableMultisample sets the value of the drawableMultisample property.
// The format of the multisampling buffer.

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKView/drawableMultisample
func (g_ GLKView) SetDrawableMultisample(value IGLKViewDrawableMultisample) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDrawableMultisample:"), value)
}

// The format of the stencil renderbuffer.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKView/drawableStencilFormat
func (g_ GLKView) DrawableStencilFormat() GLKViewDrawableStencilFormat {
	rv := objc.Send[GLKViewDrawableStencilFormat](g_.ID, objc.Sel("drawableStencilFormat"))
	return rv
}


// SetDrawableStencilFormat sets the value of the drawableStencilFormat property.
// The format of the stencil renderbuffer.

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKView/drawableStencilFormat
func (g_ GLKView) SetDrawableStencilFormat(value GLKViewDrawableStencilFormat) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDrawableStencilFormat:"), value)
}

// The width, in pixels, of the underlying framebuffer object.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKView/drawableWidth
func (g_ GLKView) DrawableWidth() int {
	rv := objc.Send[int](g_.ID, objc.Sel("drawableWidth"))
	return rv
}

// A Boolean value that indicates whether the view responds to messages that invalidate the view’s contents.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKView/enableSetNeedsDisplay
func (g_ GLKView) EnableSetNeedsDisplay() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("enableSetNeedsDisplay"))
	return rv
}


// SetEnableSetNeedsDisplay sets the value of the enableSetNeedsDisplay property.
// A Boolean value that indicates whether the view responds to messages that invalidate the view’s contents.

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKView/enableSetNeedsDisplay
func (g_ GLKView) SetEnableSetNeedsDisplay(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setEnableSetNeedsDisplay:"), value)
}

// Draws the contents of the view and returns them as a new image object.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKView/snapshot
func (g_ GLKView) Snapshot() appkit.Image {
	rv := objc.Send[appkit.Image](g_.ID, objc.Sel("snapshot"))
	return rv
}


