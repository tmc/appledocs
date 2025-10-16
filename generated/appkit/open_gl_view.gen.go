
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [OpenGLView] class.
var OpenGLViewClass _OpenGLViewClass

func init() {
	OpenGLViewClass = _OpenGLViewClass{objc.GetClass("NSOpenGLView")}
}

type _OpenGLViewClass struct {
	objc.Class
}

// An interface definition for the [OpenGLView] class.
type IOpenGLView interface {
	ID() objc.ID
	ClearGLContext()
	InitWithFramePixelFormat(frameRect unsafe.Pointer, format unsafe.Pointer) unsafe.Pointer
	PrepareOpenGL()
	Reshape()
	Update()
}

type OpenGLView struct {
	id objc.ID
}

func OpenGLViewFrom(ptr unsafe.Pointer) OpenGLView {
	return OpenGLView{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (o_ OpenGLView) ID() objc.ID {
	return o_.id
}

// Alloc allocates a new instance without initialization.
func (oc _OpenGLViewClass) Alloc() OpenGLView {
	rv := objc.Send[OpenGLView](objc.ID(oc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (oc _OpenGLViewClass) New() OpenGLView {
	rv := objc.Send[OpenGLView](objc.ID(oc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewOpenGLView creates and returns a new initialized instance.
func NewOpenGLView() OpenGLView {
	return OpenGLViewClass.New()
}

// Init initializes the instance.
func (o_ OpenGLView) Init() OpenGLView {
	rv := objc.Send[OpenGLView](o_.ID(), selInit)
	return rv
}
// Returns a default   object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSOpenGLView/defaultPixelFormat()
func (oc _OpenGLViewClass) DefaultPixelFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.Class), objc.RegisterName("defaultPixelFormat"))
	return rv
}
// Releases the   object associated with the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSOpenGLView/clearGLContext()
func (o_ OpenGLView) ClearGLContext() {
	objc.Send[objc.ID](o_.ID(), objc.RegisterName("clearGLContext"))
}
// Returns an   object initialized with the specified frame rectangle and pixel format. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSOpenGLView/init(frame:pixelFormat:)
func (o_ OpenGLView) InitWithFramePixelFormat(frameRect unsafe.Pointer, format unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID(), objc.RegisterName("initWithFrame:pixelFormat:"), frameRect, format)
	return rv
}
// Used by subclasses to initialize OpenGL state. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSOpenGLView/prepareOpenGL()
func (o_ OpenGLView) PrepareOpenGL() {
	objc.Send[objc.ID](o_.ID(), objc.RegisterName("prepareOpenGL"))
}
// Called by Cocoa when the view’s visible rectangle or bounds change. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSOpenGLView/reshape()
func (o_ OpenGLView) Reshape() {
	objc.Send[objc.ID](o_.ID(), objc.RegisterName("reshape"))
}
// Called by Cocoa when the view’s window moves or when the view itself moves or is resized. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSOpenGLView/update()
func (o_ OpenGLView) Update() {
	objc.Send[objc.ID](o_.ID(), objc.RegisterName("update"))
}
// The   object associated with the receiver. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSOpenGLView/openGLContext
func (o_ OpenGLView) OpenGLContext() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID(), objc.RegisterName("openGLContext"))
	return rv
}
// SetOpenGLContext sets the value of the openGLContext property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSOpenGLView/openGLContext
func (o_ OpenGLView) SetOpenGLContext(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID(), objc.RegisterName("setOpenGLContext:"), value)
}
// The   object associated with the receiver. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSOpenGLView/pixelFormat
func (o_ OpenGLView) PixelFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID(), objc.RegisterName("pixelFormat"))
	return rv
}
// SetPixelFormat sets the value of the pixelFormat property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSOpenGLView/pixelFormat
func (o_ OpenGLView) SetPixelFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID(), objc.RegisterName("setPixelFormat:"), value)
}
// A Boolean value indicating whether the view wants an OpenGL backing surface with a resolution greater than 1 pixel per point. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSOpenGLView/wantsBestResolutionOpenGLSurface
func (o_ OpenGLView) WantsBestResolutionOpenGLSurface() bool {
	rv := objc.Send[bool](o_.ID(), objc.RegisterName("wantsBestResolutionOpenGLSurface"))
	return rv
}
// SetWantsBestResolutionOpenGLSurface sets the value of the wantsBestResolutionOpenGLSurface property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSOpenGLView/wantsBestResolutionOpenGLSurface
func (o_ OpenGLView) SetWantsBestResolutionOpenGLSurface(value bool) {
	objc.Send[objc.ID](o_.ID(), objc.RegisterName("setWantsBestResolutionOpenGLSurface:"), value)
}
// Enables extended dynamic range values on the screen. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSOpenGLView/wantsExtendedDynamicRangeOpenGLSurface
func (o_ OpenGLView) WantsExtendedDynamicRangeOpenGLSurface() bool {
	rv := objc.Send[bool](o_.ID(), objc.RegisterName("wantsExtendedDynamicRangeOpenGLSurface"))
	return rv
}
// SetWantsExtendedDynamicRangeOpenGLSurface sets the value of the wantsExtendedDynamicRangeOpenGLSurface property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSOpenGLView/wantsExtendedDynamicRangeOpenGLSurface
func (o_ OpenGLView) SetWantsExtendedDynamicRangeOpenGLSurface(value bool) {
	objc.Send[objc.ID](o_.ID(), objc.RegisterName("setWantsExtendedDynamicRangeOpenGLSurface:"), value)
}
