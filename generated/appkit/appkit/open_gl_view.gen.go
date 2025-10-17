// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [OpenGLView] class.
var OpenGLViewClass objc.Class

func init() {
	OpenGLViewClass = objc.GetClass("NSOpenGLView")
}

type OpenGLView struct {
	objc.ID
}

func OpenGLViewFrom(ptr unsafe.Pointer) OpenGLView {
	return OpenGLView{
		ID: objc.ID(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (oc OpenGLView) Alloc() OpenGLView {
	ret := objc.ID(OpenGLViewClass).Send(objc.RegisterName("alloc"))
	return OpenGLView{ret}
}

// Init initializes the instance.
func (o_ OpenGLView) Init() OpenGLView {
	ret := o_.ID.Send(objc.RegisterName("init"))
	return OpenGLView{ret}
}
// Returns an   object initialized with the specified frame rectangle and pixel format. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSOpenGLView/init(frame:pixelFormat:)
func NewOpenGLViewWithFramePixelFormat(frameRect foundation.Rect, format unsafe.Pointer) OpenGLView {
	instance := OpenGLView{}.Alloc()
	sel := objc.RegisterName("initWithFrame:pixelFormat:")
	ret := instance.ID.Send(sel, frameRect, format)
	instance = OpenGLView{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}


// Returns a default   object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSOpenGLView/defaultPixelFormat()
func (oc OpenGLView) DefaultPixelFormat() unsafe.Pointer {
	sel := objc.RegisterName("defaultPixelFormat")
	ret := objc.ID(OpenGLViewClass).Send(sel)
	return unsafe.Pointer(ret)
}
// Releases the   object associated with the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSOpenGLView/clearGLContext()
func (o_ OpenGLView) ClearGLContext() {
	sel := objc.RegisterName("clearGLContext")
	o_.ID.Send(sel)
}
// Used by subclasses to initialize OpenGL state. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSOpenGLView/prepareOpenGL()
func (o_ OpenGLView) PrepareOpenGL() {
	sel := objc.RegisterName("prepareOpenGL")
	o_.ID.Send(sel)
}
// Called by Cocoa when the view’s visible rectangle or bounds change. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSOpenGLView/reshape()
func (o_ OpenGLView) Reshape() {
	sel := objc.RegisterName("reshape")
	o_.ID.Send(sel)
}
// Called by Cocoa when the view’s window moves or when the view itself moves or is resized. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSOpenGLView/update()
func (o_ OpenGLView) Update() {
	sel := objc.RegisterName("update")
	o_.ID.Send(sel)
}

