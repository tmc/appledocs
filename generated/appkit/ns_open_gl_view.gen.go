// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [OpenGLView] class.
var (
	OpenGLViewClass     _OpenGLViewClass
	OpenGLViewClassOnce sync.Once
)

func getOpenGLViewClass() _OpenGLViewClass {
	OpenGLViewClassOnce.Do(func() {
		OpenGLViewClass = _OpenGLViewClass{objc.GetClass("NSOpenGLView")}
	})
	return OpenGLViewClass
}

type _OpenGLViewClass struct {
	class objc.Class
}





// An interface definition for the [OpenGLView] class.
type IOpenGLView interface {
	IView
	

	// properties:
	OpenGLContext() IOpenGLContext
	SetOpenGLContext(value IOpenGLContext)
	PixelFormat() IOpenGLPixelFormat
	SetPixelFormat(value IOpenGLPixelFormat)
	WantsBestResolutionOpenGLSurface() bool
	SetWantsBestResolutionOpenGLSurface(value bool)
	WantsExtendedDynamicRangeOpenGLSurface() bool
	SetWantsExtendedDynamicRangeOpenGLSurface(value bool)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (oc _OpenGLViewClass) Alloc() OpenGLView {
	rv := objc.Send[OpenGLView](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// A view that displays OpenGL content in a view.
//
// An object maintains an and object into which OpenGL calls can be rendered. The view provides methods for accessing and managing the and objects, as well as notifications of visible region changes. An object cannot have subviews. You can, however, divide a single into multiple rendering areas using the function. When creating an object in Interface Builder, you use the inspector window to specify the pixel format attributes you want for the view. Only those attributes listed in the Interface Builder inspector are set when the view is instantiated.


// A view that displays OpenGL content in a view.
//
// [Full Topic]
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






// Returns an object initialized with the specified frame rectangle and pixel format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLView/init(frame:pixelFormat:)
func NewOpenGLViewWithFramePixelFormat(frameRect corefoundation.CGRect, format IOpenGLPixelFormat) OpenGLView {
	instance := getOpenGLViewClass().Alloc()
	rv := objc.Send[OpenGLView](instance.ID, objc.Sel("initWithFrame:pixelFormat:"), frameRect, format)
	rv.Autorelease()
	return rv
}







// Returns a default object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLView/defaultPixelFormat()
func (oc _OpenGLViewClass) DefaultPixelFormat() IOpenGLPixelFormat {
	rv := objc.Send[OpenGLPixelFormat](objc.ID(oc.class), objc.Sel("defaultPixelFormat"))
	return rv
}

















// The object associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLView/openGLContext
func (o_ OpenGLView) OpenGLContext() IOpenGLContext {
	rv := objc.Send[OpenGLContext](o_.ID, objc.Sel("openGLContext"))
	return rv
}


// The object associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLView/openGLContext
func (o_ OpenGLView) SetOpenGLContext(value IOpenGLContext) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setOpenGLContext:"), value)
}


// The object associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLView/pixelFormat
func (o_ OpenGLView) PixelFormat() IOpenGLPixelFormat {
	rv := objc.Send[OpenGLPixelFormat](o_.ID, objc.Sel("pixelFormat"))
	return rv
}


// The object associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLView/pixelFormat
func (o_ OpenGLView) SetPixelFormat(value IOpenGLPixelFormat) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setPixelFormat:"), value)
}


// A Boolean value indicating whether the view wants an OpenGL backing surface with a resolution greater than 1 pixel per point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLView/wantsBestResolutionOpenGLSurface
func (o_ OpenGLView) WantsBestResolutionOpenGLSurface() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("wantsBestResolutionOpenGLSurface"))
	return rv
}


// A Boolean value indicating whether the view wants an OpenGL backing surface with a resolution greater than 1 pixel per point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLView/wantsBestResolutionOpenGLSurface
func (o_ OpenGLView) SetWantsBestResolutionOpenGLSurface(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setWantsBestResolutionOpenGLSurface:"), value)
}


// Enables extended dynamic range values on the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLView/wantsExtendedDynamicRangeOpenGLSurface
func (o_ OpenGLView) WantsExtendedDynamicRangeOpenGLSurface() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("wantsExtendedDynamicRangeOpenGLSurface"))
	return rv
}


// Enables extended dynamic range values on the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLView/wantsExtendedDynamicRangeOpenGLSurface
func (o_ OpenGLView) SetWantsExtendedDynamicRangeOpenGLSurface(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setWantsExtendedDynamicRangeOpenGLSurface:"), value)
}







