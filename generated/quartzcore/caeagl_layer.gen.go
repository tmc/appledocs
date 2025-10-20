// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [EAGLLayer] class.
var (
	EAGLLayerClass     _EAGLLayerClass
	EAGLLayerClassOnce sync.Once
)

func getEAGLLayerClass() _EAGLLayerClass {
	EAGLLayerClassOnce.Do(func() {
		EAGLLayerClass = _EAGLLayerClass{objc.GetClass("CAEAGLLayer")}
	})
	return EAGLLayerClass
}

type _EAGLLayerClass struct {
	class objc.Class
}

// An interface definition for the [EAGLLayer] class.
type IEAGLLayer interface {
	ILayer
}

// A layer that supports drawing OpenGL content in iOS and tvOS applications.
//
// If you plan to use OpenGL for your rendering, use this class as the backing layer for your views by returning it from your view’s class method. The returned object is a wrapper for a Core Animation surface that is fully compatible with OpenGL ES function calls. Prior to designating the layer’s associated view as the render target for a graphics context, you can change the rendering attributes you want using the property. This property lets you configure the color format for the rendering surface and whether the surface retains its contents. For a list of keys (and corresponding values) you can include in this dictionary (along with their default values), see the . Because an OpenGL ES rendering surface is presented to the user using Core Animation, any effects and animations you apply to the layer affect the 3D content you render. However, for best performance, do the following: Set the layer’s opaque attribute to . Set the layer bounds to match the dimensions of the display. Make sure the layer is not transformed. Avoid drawing other layers on top of the object. If you must draw other, non OpenGL content, you might find the performance cost acceptable if you place transparent 2D content on top of the GL content and also make sure that the OpenGL content is opaque and not transformed. When drawing landscape content on a portrait display, you should rotate the content yourself rather than using the transform to rotate it.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEAGLLayer
type EAGLLayer struct {
	Layer
}

// EAGLLayerFrom constructs a [EAGLLayer] from an unsafe.Pointer.
//
// A layer that supports drawing OpenGL content in iOS and tvOS applications.
func EAGLLayerFrom(ptr unsafe.Pointer) EAGLLayer {
	return EAGLLayer{
		Layer: LayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ec _EAGLLayerClass) Alloc() EAGLLayer {
	rv := objc.Send[EAGLLayer](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _EAGLLayerClass) New() EAGLLayer {
	rv := objc.Send[EAGLLayer](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EAGLLayer) Init() EAGLLayer {
	rv := objc.Send[EAGLLayer](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EAGLLayer) Autorelease() EAGLLayer {
	rv := objc.Send[EAGLLayer](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEAGLLayer creates a new EAGLLayer instance.
func NewEAGLLayer() EAGLLayer {
	return getEAGLLayerClass().New()
}


// A Boolean value that determines whether the layer presents its content using a Core Animation transaction.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEAGLLayer/presentsWithTransaction
func (e_ EAGLLayer) PresentsWithTransaction() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("presentsWithTransaction"))
	return rv
}


// SetPresentsWithTransaction sets the value of the presentsWithTransaction property.
// A Boolean value that determines whether the layer presents its content using a Core Animation transaction.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEAGLLayer/presentsWithTransaction
func (e_ EAGLLayer) SetPresentsWithTransaction(value bool) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setPresentsWithTransaction:"), value)
}


