// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	

	// properties:
	Colorspace() ColorSpaceRef /* not a class type */
	SetColorspace(value ColorSpaceRef /* not a class type */)
	Asynchronous() bool
	SetAsynchronous(value bool)
	WantsExtendedDynamicRangeContent() bool
	SetWantsExtendedDynamicRangeContent(value bool)
	IsAsynchronous() bool
	SetIsAsynchronous(value bool)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (oc _OpenGLLayerClass) Alloc() OpenGLLayer {
	rv := objc.Send[OpenGLLayer](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// A layer that provides a layer suitable for rendering OpenGL content.
//
// To provide OpenGL content you subclass and override . You can specify that the OpenGL content is static by setting the property to .


// A layer that provides a layer suitable for rendering OpenGL content.
//
// [Full Topic]
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

























// The layer’s colorspace in Core Graphics.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAOpenGLLayer/colorspace
func (o_ OpenGLLayer) Colorspace() ColorSpaceRef /* not a class type */ {
	rv := objc.Send[ColorSpaceRef](o_.ID, objc.Sel("colorspace"))
	return rv
}


// The layer’s colorspace in Core Graphics.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAOpenGLLayer/colorspace
func (o_ OpenGLLayer) SetColorspace(value ColorSpaceRef /* not a class type */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setColorspace:"), value)
}


// Determines when the contents of the layer are updated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAOpenGLLayer/isAsynchronous
func (o_ OpenGLLayer) Asynchronous() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("asynchronous"))
	return rv
}


// Determines when the contents of the layer are updated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAOpenGLLayer/isAsynchronous
func (o_ OpenGLLayer) SetAsynchronous(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAsynchronous:"), value)
}


// Tells whether or not the layer supports content with extended dynamic range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAOpenGLLayer/wantsExtendedDynamicRangeContent
func (o_ OpenGLLayer) WantsExtendedDynamicRangeContent() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("wantsExtendedDynamicRangeContent"))
	return rv
}


// Tells whether or not the layer supports content with extended dynamic range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAOpenGLLayer/wantsExtendedDynamicRangeContent
func (o_ OpenGLLayer) SetWantsExtendedDynamicRangeContent(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setWantsExtendedDynamicRangeContent:"), value)
}


// Determines when the contents of the layer are updated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caopengllayer/isasynchronous
func (o_ OpenGLLayer) IsAsynchronous() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("isAsynchronous"))
	return rv
}


// Determines when the contents of the layer are updated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caopengllayer/isasynchronous
func (o_ OpenGLLayer) SetIsAsynchronous(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setIsAsynchronous:"), value)
}








