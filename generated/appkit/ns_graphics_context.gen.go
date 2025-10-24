// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coreimage"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/vision"
)

/* debug [class.gen.go]: Generating class NSGraphicsContext */


/* debug [class_header]: Header for NSGraphicsContext */
// The class instance for the [GraphicsContext] class.
var (
	GraphicsContextClass     _GraphicsContextClass
	GraphicsContextClassOnce sync.Once
)

func getGraphicsContextClass() _GraphicsContextClass {
	GraphicsContextClassOnce.Do(func() {
		GraphicsContextClass = _GraphicsContextClass{objc.GetClass("NSGraphicsContext")}
	})
	return GraphicsContextClass
}

type _GraphicsContextClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GraphicsContext */
// An interface definition for the [GraphicsContext] class.
type IGraphicsContext interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for GraphicsContext */
	// properties:
	CompositingOperation() CompositingOperation
	SetCompositingOperation(value CompositingOperation)
	ImageInterpolation() ImageInterpolation
	SetImageInterpolation(value ImageInterpolation)
	PatternPhase() vision.Point
	SetPatternPhase(value vision.Point)
	ShouldAntialias() bool
	SetShouldAntialias(value bool)
	Attributes() objectivec.IObject
	SetAttributes(value objectivec.IObject)
	CgContext() coreimage.Context
	SetCgContext(value coreimage.Context)
	CiContext() coreimage.Context
	SetCiContext(value coreimage.Context)
	ColorRenderingIntent() ColorRenderingIntent /* not a class type */
	SetColorRenderingIntent(value ColorRenderingIntent /* not a class type */)
	GraphicsPort() objectivec.IObject
	SetGraphicsPort(value objectivec.IObject)
	IsDrawingToScreen() bool
	SetIsDrawingToScreen(value bool)
	IsFlipped() bool
	SetIsFlipped(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GraphicsContext */
	// methods:
	RestoreGraphicsState()
	SaveGraphicsState()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GraphicsContext */
// Alloc allocates a new instance without initialization.
func (gc _GraphicsContextClass) Alloc() GraphicsContext {
	rv := objc.Send[GraphicsContext](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GraphicsContextClass) New() GraphicsContext {
	rv := objc.Send[GraphicsContext](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GraphicsContext) Init() GraphicsContext {
	rv := objc.Send[GraphicsContext](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GraphicsContext) Autorelease() GraphicsContext {
	rv := objc.Send[GraphicsContext](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGraphicsContext creates a new GraphicsContext instance.
func NewGraphicsContext() GraphicsContext {
	return getGraphicsContextClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GraphicsContext */
// An object that represents a graphics context.
//
// You can think of a graphics context as a destination to which drawing and graphics state operations are sent for execution. Each graphics context contains its own graphics environment and state. The class is an abstract superclass for destination-specific graphics contexts. You obtain instances of concrete subclasses with the class methods , , , , and . At any time there is the notion of the current context. The current context for the current thread may be set using . Graphics contexts are maintained on a stack. You push a graphics context onto the stack by sending it a message, and pop it off the stack by sending it a message. By sending to a graphics context object you remove it from the stack, and the next graphics context on the stack becomes the current graphics context.


// An object that represents a graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext
type GraphicsContext struct {
	objectivec.Object
}

// GraphicsContextFrom constructs a [GraphicsContext] from an unsafe.Pointer.
//
// An object that represents a graphics context.
func GraphicsContextFrom(ptr unsafe.Pointer) GraphicsContext {
	return GraphicsContext{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GraphicsContext *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GraphicsContext */

// Pops a graphics context from the per-thread stack, makes it current, and sends the context a restore graphics state message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/restoreGraphicsState()-swift.type.method
func (gc _GraphicsContextClass) RestoreGraphicsState() {
	objc.Send[objc.ID](objc.ID(gc.class), objc.Sel("restoreGraphicsState"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RestoreGraphicsState) */


// Saves the graphics state of the current graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/saveGraphicsState()-swift.type.method
func (gc _GraphicsContextClass) SaveGraphicsState() {
	objc.Send[objc.ID](objc.ID(gc.class), objc.Sel("saveGraphicsState"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SaveGraphicsState) */


// Makes the graphics context of the specified graphics state current, and resets graphics state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/setGraphicsState(_:)
func (gc _GraphicsContextClass) SetGraphicsState(gState int) {
	objc.Send[objc.ID](objc.ID(gc.class), objc.Sel("setGraphicsState:"), gState)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SetGraphicsState) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GraphicsContext */

// Returns the current graphics context of the current thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/current
func (gc _GraphicsContextClass) CurrentContext() GraphicsContext {
	rv := objc.Send[GraphicsContext](objc.ID(gc.class), objc.Sel("currentContext"))
	return rv
}/* debug [class_properties_class/property]: currentContext */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GraphicsContext */

// Removes the context’s graphics state from the top of the graphics state stack and makes the next graphics state the current graphics state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/restoreGraphicsState()-swift.method
func (g_ GraphicsContext) RestoreGraphicsState() {
	objc.Send[objc.ID](g_.ID, objc.Sel("restoreGraphicsState"))
}/* debug [instance_methods/method]: RestoreGraphicsState */


// Saves the current graphics state and creates a new graphics state on the top of the stack.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/saveGraphicsState()-swift.method
func (g_ GraphicsContext) SaveGraphicsState() {
	objc.Send[objc.ID](g_.ID, objc.Sel("saveGraphicsState"))
}/* debug [instance_methods/method]: SaveGraphicsState */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GraphicsContext */

// The graphics context’s global compositing operation setting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/compositingOperation
func (g_ GraphicsContext) CompositingOperation() CompositingOperation {
	rv := objc.Send[CompositingOperation](g_.ID, objc.Sel("compositingOperation"))
	return rv
}/* debug [instance_properties/getter]: compositingOperation */


// The graphics context’s global compositing operation setting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/compositingOperation
func (g_ GraphicsContext) SetCompositingOperation(value CompositingOperation) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setCompositingOperation:"), value)
}/* debug [instance_properties/setter]: compositingOperation */


// Returns the current graphics context of the current thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/current
func (g_ GraphicsContext) CurrentContext() IGraphicsContext {
	rv := objc.Send[GraphicsContext](g_.ID, objc.Sel("currentContext"))
	return rv
}/* debug [instance_properties/getter]: currentContext */


// Returns the current graphics context of the current thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/current
func (g_ GraphicsContext) SetCurrentContext(value IGraphicsContext) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setCurrentContext:"), value)
}/* debug [instance_properties/setter]: currentContext */


// A constant that specifies the graphics context’s interpolation, or image smoothing, behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/imageInterpolation
func (g_ GraphicsContext) ImageInterpolation() ImageInterpolation {
	rv := objc.Send[ImageInterpolation](g_.ID, objc.Sel("imageInterpolation"))
	return rv
}/* debug [instance_properties/getter]: imageInterpolation */


// A constant that specifies the graphics context’s interpolation, or image smoothing, behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/imageInterpolation
func (g_ GraphicsContext) SetImageInterpolation(value ImageInterpolation) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setImageInterpolation:"), value)
}/* debug [instance_properties/setter]: imageInterpolation */


// The amount to offset the pattern color when filling the graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/patternPhase
func (g_ GraphicsContext) PatternPhase() vision.Point {
	rv := objc.Send[vision.Point](g_.ID, objc.Sel("patternPhase"))
	return rv
}/* debug [instance_properties/getter]: patternPhase */


// The amount to offset the pattern color when filling the graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/patternPhase
func (g_ GraphicsContext) SetPatternPhase(value vision.Point) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPatternPhase:"), value)
}/* debug [instance_properties/setter]: patternPhase */


// A Boolean value that indicates whether the graphics context uses antialiasing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/shouldAntialias
func (g_ GraphicsContext) ShouldAntialias() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("shouldAntialias"))
	return rv
}/* debug [instance_properties/getter]: shouldAntialias */


// A Boolean value that indicates whether the graphics context uses antialiasing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/shouldAntialias
func (g_ GraphicsContext) SetShouldAntialias(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setShouldAntialias:"), value)
}/* debug [instance_properties/setter]: shouldAntialias */


// The attributes used to create this instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/attributes
func (g_ GraphicsContext) Attributes() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](g_.ID, objc.Sel("attributes"))
	return rv
}/* debug [instance_properties/getter]: attributes */


// The attributes used to create this instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/attributes
func (g_ GraphicsContext) SetAttributes(value objectivec.IObject) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setAttributes:"), value)
}/* debug [instance_properties/setter]: attributes */


// The Core Graphics context, which is a low-level, platform-specific graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/cgcontext
func (g_ GraphicsContext) CgContext() coreimage.Context {
	rv := objc.Send[coreimage.Context](g_.ID, objc.Sel("cgContext"))
	return rv
}/* debug [instance_properties/getter]: cgContext */


// The Core Graphics context, which is a low-level, platform-specific graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/cgcontext
func (g_ GraphicsContext) SetCgContext(value coreimage.Context) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setCgContext:"), value)
}/* debug [instance_properties/setter]: cgContext */


// A context for Core Image objects that you can use to render into the graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/cicontext
func (g_ GraphicsContext) CiContext() coreimage.Context {
	rv := objc.Send[coreimage.Context](g_.ID, objc.Sel("ciContext"))
	return rv
}/* debug [instance_properties/getter]: ciContext */


// A context for Core Image objects that you can use to render into the graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/cicontext
func (g_ GraphicsContext) SetCiContext(value coreimage.Context) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setCiContext:"), value)
}/* debug [instance_properties/setter]: ciContext */


// The color rendering intent in the graphics context’s graphics state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/colorrenderingintent
func (g_ GraphicsContext) ColorRenderingIntent() ColorRenderingIntent /* not a class type */ {
	rv := objc.Send[ColorRenderingIntent](g_.ID, objc.Sel("colorRenderingIntent"))
	return rv
}/* debug [instance_properties/getter]: colorRenderingIntent */


// The color rendering intent in the graphics context’s graphics state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/colorrenderingintent
func (g_ GraphicsContext) SetColorRenderingIntent(value ColorRenderingIntent /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setColorRenderingIntent:"), value)
}/* debug [instance_properties/setter]: colorRenderingIntent */


// The low-level, platform-specific graphics context represented by the graphic port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/graphicsport
func (g_ GraphicsContext) GraphicsPort() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](g_.ID, objc.Sel("graphicsPort"))
	return rv
}/* debug [instance_properties/getter]: graphicsPort */


// The low-level, platform-specific graphics context represented by the graphic port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/graphicsport
func (g_ GraphicsContext) SetGraphicsPort(value objectivec.IObject) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGraphicsPort:"), value)
}/* debug [instance_properties/setter]: graphicsPort */


// A Boolean value that indicates whether the drawing destination is the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/isdrawingtoscreen
func (g_ GraphicsContext) IsDrawingToScreen() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("isDrawingToScreen"))
	return rv
}/* debug [instance_properties/getter]: isDrawingToScreen */


// A Boolean value that indicates whether the drawing destination is the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/isdrawingtoscreen
func (g_ GraphicsContext) SetIsDrawingToScreen(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIsDrawingToScreen:"), value)
}/* debug [instance_properties/setter]: isDrawingToScreen */


// A Boolean value that indicates the graphics context’s flipped state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/isflipped
func (g_ GraphicsContext) IsFlipped() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("isFlipped"))
	return rv
}/* debug [instance_properties/getter]: isFlipped */


// A Boolean value that indicates the graphics context’s flipped state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/isflipped
func (g_ GraphicsContext) SetIsFlipped(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIsFlipped:"), value)
}/* debug [instance_properties/setter]: isFlipped */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSGraphicsContext */



