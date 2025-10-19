// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [OpenGLView] class.
var (
	openGLViewClass     _OpenGLViewClass
	openGLViewClassOnce sync.Once
)

func getOpenGLViewClass() _OpenGLViewClass {
	openGLViewClassOnce.Do(func() {
		openGLViewClass = _OpenGLViewClass{objc.GetClass("NSOpenGLView")}
	})
	return openGLViewClass
}

type _OpenGLViewClass struct {
	class objc.Class
}

// An interface definition for the [OpenGLView] class.
type IOpenGLView interface {
	IView
	ClearGLContext()
	PrepareOpenGL()
	Reshape()
	Update()
}

// A view that displays OpenGL content in a view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLView
type OpenGLView struct {
	View
}

// OpenGLViewFrom constructs a [OpenGLView] from an unsafe.Pointer.
//
// A view that displays OpenGL content in a view.
func OpenGLViewFrom(ptr unsafe.Pointer) OpenGLView {
	return OpenGLView{
		View: ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (oc _OpenGLViewClass) Alloc() OpenGLView {
	rv := objc.Send[OpenGLView](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _OpenGLViewClass) New() OpenGLView {
	rv := objc.Send[OpenGLView](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OpenGLView) Init() OpenGLView {
	rv := objc.Send[OpenGLView](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OpenGLView) Autorelease() OpenGLView {
	rv := objc.Send[OpenGLView](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOpenGLView creates a new OpenGLView instance.
func NewOpenGLView() OpenGLView {
	return getOpenGLViewClass().New()
}


// Returns an object initialized with the specified frame rectangle and pixel format. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLView/init(frame:pixelFormat:)
func NewOpenGLViewWithFramePixelFormat(frameRect unsafe.Pointer, format unsafe.Pointer) OpenGLView {
	instance := getOpenGLViewClass().Alloc()
	rv := objc.Send[OpenGLView](instance.ID, objc.Sel("initWithFrame:pixelFormat:"), frameRect, format)
	rv.Autorelease()
	return rv
}


// Returns a default object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLView/defaultPixelFormat()
func (oc _OpenGLViewClass) DefaultPixelFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("defaultPixelFormat"))
	return rv
}
// Releases the object associated with the view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLView/clearGLContext()
func (o_ OpenGLView) ClearGLContext() {
	objc.Send[objc.ID](o_.ID, objc.Sel("clearGLContext"))
}
// Used by subclasses to initialize OpenGL state. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLView/prepareOpenGL()
func (o_ OpenGLView) PrepareOpenGL() {
	objc.Send[objc.ID](o_.ID, objc.Sel("prepareOpenGL"))
}
// Called by Cocoa when the view’s visible rectangle or bounds change. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLView/reshape()
func (o_ OpenGLView) Reshape() {
	objc.Send[objc.ID](o_.ID, objc.Sel("reshape"))
}
// Called by Cocoa when the view’s window moves or when the view itself moves or is resized. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLView/update()
func (o_ OpenGLView) Update() {
	objc.Send[objc.ID](o_.ID, objc.Sel("update"))
}

