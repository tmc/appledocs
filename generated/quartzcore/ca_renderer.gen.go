// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/corevideo"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CARenderer */


/* debug [class_header]: Header for CARenderer */
// The class instance for the [Renderer] class.
var (
	RendererClass     _RendererClass
	RendererClassOnce sync.Once
)

func getRendererClass() _RendererClass {
	RendererClassOnce.Do(func() {
		RendererClass = _RendererClass{objc.GetClass("CARenderer")}
	})
	return RendererClass
}

type _RendererClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Renderer */
// An interface definition for the [Renderer] class.
type IRenderer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Renderer */
	// properties:
	Bounds() corefoundation.CGRect
	SetBounds(value corefoundation.CGRect)
	Layer() ILayer
	SetLayer(value ILayer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Renderer */
	// methods:
	AddUpdateRect(r corefoundation.CGRect)
	BeginFrameAtTimeTimeStamp(t float64, ts corevideo.CVTimeStamp)
	EndFrame()
	NextFrameTime() float64
	Render()
	SetDestination(tex unsafe.Pointer)
	UpdateBounds() corefoundation.CGRect
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Renderer */
// Alloc allocates a new instance without initialization.
func (rc _RendererClass) Alloc() Renderer {
	rv := objc.Send[Renderer](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RendererClass) New() Renderer {
	rv := objc.Send[Renderer](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ Renderer) Init() Renderer {
	rv := objc.Send[Renderer](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ Renderer) Autorelease() Renderer {
	rv := objc.Send[Renderer](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRenderer creates a new Renderer instance.
func NewRenderer() Renderer {
	return getRendererClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Renderer */
// A layer that allows an application to render a layer tree into a Core OpenGL context.
//
// For real-time output you should use an instance of to host the layer-tree.


// A layer that allows an application to render a layer tree into a Core OpenGL context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CARenderer
type Renderer struct {
	objectivec.Object
}

// RendererFrom constructs a [Renderer] from an unsafe.Pointer.
//
// A layer that allows an application to render a layer tree into a Core OpenGL context.
func RendererFrom(ptr unsafe.Pointer) Renderer {
	return Renderer{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Renderer */

// Creates and returns a instance with the render target specified by the Core OpenGL context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CARenderer/init(cglContext:options:)
func NewRendererWithCGLContextOptions(ctx objectivec.IObject, dict objc.IObject /* cross-framework: NSDictionary */) Renderer {
	rv := objc.Send[Renderer](objc.ID(getRendererClass().class), objc.Sel("rendererWithCGLContext:options:"), ctx, dict)
	return rv
}/* debug [class_init_methods/constructor]: NewRendererWithCGLContextOptions */


// Creates a layer renderer from a Metal texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CARenderer/init(mtlTexture:options:)
func NewRendererWithMTLTextureOptions(tex unsafe.Pointer, dict objc.IObject /* cross-framework: NSDictionary */) Renderer {
	rv := objc.Send[Renderer](objc.ID(getRendererClass().class), objc.Sel("rendererWithMTLTexture:options:"), tex, dict)
	return rv
}/* debug [class_init_methods/constructor]: NewRendererWithMTLTextureOptions */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Renderer */

// Creates and returns a instance with the render target specified by the Core OpenGL context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CARenderer/init(cglContext:options:)
func (rc _RendererClass) RendererWithCGLContextOptions(ctx objectivec.IObject, dict objc.IObject /* cross-framework: NSDictionary */) IRenderer {
	rv := objc.Send[Renderer](objc.ID(rc.class), objc.Sel("rendererWithCGLContext:options:"), ctx, dict)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RendererWithCGLContextOptions) */


// Creates a layer renderer from a Metal texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CARenderer/init(mtlTexture:options:)
func (rc _RendererClass) RendererWithMTLTextureOptions(tex unsafe.Pointer, dict objc.IObject /* cross-framework: NSDictionary */) IRenderer {
	rv := objc.Send[Renderer](objc.ID(rc.class), objc.Sel("rendererWithMTLTexture:options:"), tex, dict)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RendererWithMTLTextureOptions) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Renderer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Renderer */

// Adds the rectangle to the update region of the current frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CARenderer/addUpdate(_:)
func (r_ Renderer) AddUpdateRect(r corefoundation.CGRect) {
	objc.Send[objc.ID](r_.ID, objc.Sel("addUpdateRect:"), r)
}/* debug [instance_methods/method]: AddUpdateRect */


// Begin rendering a frame at the specified time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CARenderer/beginFrame(atTime:timeStamp:)
func (r_ Renderer) BeginFrameAtTimeTimeStamp(t float64, ts corevideo.CVTimeStamp) {
	objc.Send[objc.ID](r_.ID, objc.Sel("beginFrameAtTime:timeStamp:"), t, ts)
}/* debug [instance_methods/method]: BeginFrameAtTimeTimeStamp */


// Release any data associated with the current frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CARenderer/endFrame()
func (r_ Renderer) EndFrame() {
	objc.Send[objc.ID](r_.ID, objc.Sel("endFrame"))
}/* debug [instance_methods/method]: EndFrame */


// Returns the time at which the next update should happen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CARenderer/nextFrameTime()
func (r_ Renderer) NextFrameTime() float64 {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("nextFrameTime"))
	return rv
}/* debug [instance_methods/method]: NextFrameTime */


// Render the update region of the current frame to the target context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CARenderer/render()
func (r_ Renderer) Render() {
	objc.Send[objc.ID](r_.ID, objc.Sel("render"))
}/* debug [instance_methods/method]: Render */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CARenderer/setDestination(_:)
func (r_ Renderer) SetDestination(tex unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setDestination:"), tex)
}/* debug [instance_methods/method]: SetDestination */


// Returns the bounds of the update region that contains all pixels that will be rendered by the current frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CARenderer/updateBounds()
func (r_ Renderer) UpdateBounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](r_.ID, objc.Sel("updateBounds"))
	return rv
}/* debug [instance_methods/method]: UpdateBounds */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Renderer */

// The bounds of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CARenderer/bounds
func (r_ Renderer) Bounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](r_.ID, objc.Sel("bounds"))
	return rv
}/* debug [instance_properties/getter]: bounds */


// The bounds of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CARenderer/bounds
func (r_ Renderer) SetBounds(value corefoundation.CGRect) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setBounds:"), value)
}/* debug [instance_properties/setter]: bounds */


// The root layer of the layer-tree the receiver should render.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CARenderer/layer
func (r_ Renderer) Layer() ILayer {
	rv := objc.Send[Layer](r_.ID, objc.Sel("layer"))
	return rv
}/* debug [instance_properties/getter]: layer */


// The root layer of the layer-tree the receiver should render.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CARenderer/layer
func (r_ Renderer) SetLayer(value ILayer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setLayer:"), value)
}/* debug [instance_properties/setter]: layer */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CARenderer */


