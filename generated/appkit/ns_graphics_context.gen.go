// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [GraphicsContext] class.
type IGraphicsContext interface {
	objectivec.IObject
	// properties:
	Attributes() foundation.IDictionary /* already interface */
	CGContext() ContextRef /* not a class type */
	CIContext() objc.IObject /* cross-framework: Context */
	ColorRenderingIntent() ColorRenderingIntent
	SetColorRenderingIntent(value ColorRenderingIntent)
	CompositingOperation() CompositingOperation
	SetCompositingOperation(value CompositingOperation)
	GraphicsPort() unsafe.Pointer
	ImageInterpolation() ImageInterpolation
	SetImageInterpolation(value ImageInterpolation)
	DrawingToScreen() bool /* primitive/slice/pointer. */
	Flipped() bool /* primitive/slice/pointer. */
	PatternPhase() objc.IObject /* cross-framework: Point */
	SetPatternPhase(value objc.IObject /* cross-framework: Point */)
	ShouldAntialias() bool /* primitive/slice/pointer. */
	SetShouldAntialias(value bool /* primitive/slice/pointer. */)
	IsDrawingToScreen() bool /* primitive/slice/pointer. */
	SetIsDrawingToScreen(value bool /* primitive/slice/pointer. */)
	IsFlipped() bool /* primitive/slice/pointer. */
	SetIsFlipped(value bool /* primitive/slice/pointer. */)
	// methods:
	FlushGraphics()
	RestoreGraphicsState()
	SaveGraphicsState()
}

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

// Alloc allocates a new instance without initialization.
func (gc _GraphicsContextClass) Alloc() GraphicsContext {
	rv := objc.Send[GraphicsContext](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Creates a graphics context using the specified attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/init(attributes:)
func NewGraphicsContextWithAttributes(attributes foundation.IDictionary /* already interface */) GraphicsContext {
	rv := objc.Send[GraphicsContext](objc.ID(getGraphicsContextClass().class), objc.Sel("graphicsContextWithAttributes:"), attributes)
	return rv
}


// Creates a new graphics context using the specified bitmap image representation object as the context destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/init(bitmapImageRep:)
func NewGraphicsContextWithBitmapImageRep(bitmapRep IBitmapImageRep) GraphicsContext {
	rv := objc.Send[GraphicsContext](objc.ID(getGraphicsContextClass().class), objc.Sel("graphicsContextWithBitmapImageRep:"), bitmapRep)
	return rv
}


// Creates a new graphics context from the specified Core Graphics context and the initial flipped state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/init(cgContext:flipped:)
func NewGraphicsContextWithCGContextFlipped(graphicsPort ContextRef /* not a class type */, initialFlippedState bool /* primitive/slice/pointer. */) GraphicsContext {
	rv := objc.Send[GraphicsContext](objc.ID(getGraphicsContextClass().class), objc.Sel("graphicsContextWithCGContext:flipped:"), graphicsPort, initialFlippedState)
	return rv
}


// Creates a new graphics context from the specified graphics port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/init(graphicsPort:flipped:)
func NewGraphicsContextWithGraphicsPortFlipped(graphicsPort unsafe.Pointer, initialFlippedState bool /* primitive/slice/pointer. */) GraphicsContext {
	rv := objc.Send[GraphicsContext](objc.ID(getGraphicsContextClass().class), objc.Sel("graphicsContextWithGraphicsPort:flipped:"), graphicsPort, initialFlippedState)
	return rv
}


// Creates a new graphics context for drawing into a window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/init(window:)
func NewGraphicsContextWithWindow(window IWindow) GraphicsContext {
	rv := objc.Send[GraphicsContext](objc.ID(getGraphicsContextClass().class), objc.Sel("graphicsContextWithWindow:"), window)
	return rv
}



// Returns a Boolean value that indicates whether the current context is drawing to the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/currentContextDrawingToScreen()
func (gc _GraphicsContextClass) CurrentContextDrawingToScreen() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](objc.ID(gc.class), objc.Sel("currentContextDrawingToScreen"))
	return rv
}


// Creates a graphics context using the specified attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/init(attributes:)
func (gc _GraphicsContextClass) GraphicsContextWithAttributes(attributes foundation.IDictionary /* already interface */) IGraphicsContext {
	rv := objc.Send[GraphicsContext](objc.ID(gc.class), objc.Sel("graphicsContextWithAttributes:"), attributes)
	return rv
}


// Creates a new graphics context using the specified bitmap image representation object as the context destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/init(bitmapImageRep:)
func (gc _GraphicsContextClass) GraphicsContextWithBitmapImageRep(bitmapRep IBitmapImageRep) IGraphicsContext {
	rv := objc.Send[GraphicsContext](objc.ID(gc.class), objc.Sel("graphicsContextWithBitmapImageRep:"), bitmapRep)
	return rv
}


// Creates a new graphics context from the specified Core Graphics context and the initial flipped state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/init(cgContext:flipped:)
func (gc _GraphicsContextClass) GraphicsContextWithCGContextFlipped(graphicsPort ContextRef /* not a class type */, initialFlippedState bool /* primitive/slice/pointer. */) IGraphicsContext {
	rv := objc.Send[GraphicsContext](objc.ID(gc.class), objc.Sel("graphicsContextWithCGContext:flipped:"), graphicsPort, initialFlippedState)
	return rv
}


// Creates a new graphics context from the specified graphics port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/init(graphicsPort:flipped:)
func (gc _GraphicsContextClass) GraphicsContextWithGraphicsPortFlipped(graphicsPort unsafe.Pointer, initialFlippedState bool /* primitive/slice/pointer. */) IGraphicsContext {
	rv := objc.Send[GraphicsContext](objc.ID(gc.class), objc.Sel("graphicsContextWithGraphicsPort:flipped:"), graphicsPort, initialFlippedState)
	return rv
}


// Creates a new graphics context for drawing into a window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/init(window:)
func (gc _GraphicsContextClass) GraphicsContextWithWindow(window IWindow) IGraphicsContext {
	rv := objc.Send[GraphicsContext](objc.ID(gc.class), objc.Sel("graphicsContextWithWindow:"), window)
	return rv
}


// Pops a graphics context from the per-thread stack, makes it current, and sends the context a restore graphics state message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/restoreGraphicsState()-swift.type.method
func (gc _GraphicsContextClass) RestoreGraphicsState() {
	objc.Send[objc.ID](objc.ID(gc.class), objc.Sel("restoreGraphicsState"))
}


// Saves the graphics state of the current graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/saveGraphicsState()-swift.type.method
func (gc _GraphicsContextClass) SaveGraphicsState() {
	objc.Send[objc.ID](objc.ID(gc.class), objc.Sel("saveGraphicsState"))
}


// Makes the graphics context of the specified graphics state current, and resets graphics state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/setGraphicsState(_:)
func (gc _GraphicsContextClass) SetGraphicsState(gState int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](objc.ID(gc.class), objc.Sel("setGraphicsState:"), gState)
}


// Returns the current graphics context of the current thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/current
func (gc _GraphicsContextClass) CurrentContext() GraphicsContext {
	rv := objc.Send[GraphicsContext](objc.ID(gc.class), objc.Sel("currentContext"))
	return rv
}

// Forces any buffered operations or data to be sent to the graphics context’s destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/flushGraphics()
func (g_ GraphicsContext) FlushGraphics() {
	objc.Send[objc.ID](g_.ID, objc.Sel("flushGraphics"))
}


// Removes the context’s graphics state from the top of the graphics state stack and makes the next graphics state the current graphics state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/restoreGraphicsState()-swift.method
func (g_ GraphicsContext) RestoreGraphicsState() {
	objc.Send[objc.ID](g_.ID, objc.Sel("restoreGraphicsState"))
}


// Saves the current graphics state and creates a new graphics state on the top of the stack.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/saveGraphicsState()-swift.method
func (g_ GraphicsContext) SaveGraphicsState() {
	objc.Send[objc.ID](g_.ID, objc.Sel("saveGraphicsState"))
}


// The attributes used to create this instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/attributes
func (g_ GraphicsContext) Attributes() foundation.IDictionary /* already interface */ {
	rv := objc.Send[foundation.IDictionary](g_.ID, objc.Sel("attributes"))
	return rv
}


// The Core Graphics context, which is a low-level, platform-specific graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/cgContext
func (g_ GraphicsContext) CGContext() ContextRef /* not a class type */ {
	rv := objc.Send[ContextRef](g_.ID, objc.Sel("CGContext"))
	return rv
}


// A context for Core Image objects that you can use to render into the graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/ciContext
func (g_ GraphicsContext) CIContext() objc.IObject /* cross-framework: Context */ {
	rv := objc.Send[Context](g_.ID, objc.Sel("CIContext"))
	return rv
}


// The color rendering intent in the graphics context’s graphics state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/colorRenderingIntent
func (g_ GraphicsContext) ColorRenderingIntent() ColorRenderingIntent {
	rv := objc.Send[ColorRenderingIntent](g_.ID, objc.Sel("colorRenderingIntent"))
	return rv
}


// The color rendering intent in the graphics context’s graphics state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/colorRenderingIntent
func (g_ GraphicsContext) SetColorRenderingIntent(value ColorRenderingIntent) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setColorRenderingIntent:"), value)
}


// The graphics context’s global compositing operation setting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/compositingOperation
func (g_ GraphicsContext) CompositingOperation() CompositingOperation {
	rv := objc.Send[CompositingOperation](g_.ID, objc.Sel("compositingOperation"))
	return rv
}


// The graphics context’s global compositing operation setting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/compositingOperation
func (g_ GraphicsContext) SetCompositingOperation(value CompositingOperation) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setCompositingOperation:"), value)
}


// Returns the current graphics context of the current thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/current
func (g_ GraphicsContext) CurrentContext() IGraphicsContext {
	rv := objc.Send[GraphicsContext](g_.ID, objc.Sel("currentContext"))
	return rv
}


// Returns the current graphics context of the current thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/current
func (g_ GraphicsContext) SetCurrentContext(value IGraphicsContext) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setCurrentContext:"), value)
}


// The low-level, platform-specific graphics context represented by the graphic port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/graphicsPort
func (g_ GraphicsContext) GraphicsPort() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("graphicsPort"))
	return rv
}


// A constant that specifies the graphics context’s interpolation, or image smoothing, behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/imageInterpolation
func (g_ GraphicsContext) ImageInterpolation() ImageInterpolation {
	rv := objc.Send[ImageInterpolation](g_.ID, objc.Sel("imageInterpolation"))
	return rv
}


// A constant that specifies the graphics context’s interpolation, or image smoothing, behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/imageInterpolation
func (g_ GraphicsContext) SetImageInterpolation(value ImageInterpolation) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setImageInterpolation:"), value)
}


// A Boolean value that indicates whether the drawing destination is the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/isDrawingToScreen
func (g_ GraphicsContext) DrawingToScreen() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](g_.ID, objc.Sel("drawingToScreen"))
	return rv
}


// A Boolean value that indicates the graphics context’s flipped state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/isFlipped
func (g_ GraphicsContext) Flipped() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](g_.ID, objc.Sel("flipped"))
	return rv
}


// The amount to offset the pattern color when filling the graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/patternPhase
func (g_ GraphicsContext) PatternPhase() objc.IObject /* cross-framework: Point */ {
	rv := objc.Send[Point](g_.ID, objc.Sel("patternPhase"))
	return rv
}


// The amount to offset the pattern color when filling the graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/patternPhase
func (g_ GraphicsContext) SetPatternPhase(value objc.IObject /* cross-framework: Point */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPatternPhase:"), value)
}


// A Boolean value that indicates whether the graphics context uses antialiasing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/shouldAntialias
func (g_ GraphicsContext) ShouldAntialias() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](g_.ID, objc.Sel("shouldAntialias"))
	return rv
}


// A Boolean value that indicates whether the graphics context uses antialiasing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/shouldAntialias
func (g_ GraphicsContext) SetShouldAntialias(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setShouldAntialias:"), value)
}


// A Boolean value that indicates whether the drawing destination is the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/isdrawingtoscreen
func (g_ GraphicsContext) IsDrawingToScreen() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](g_.ID, objc.Sel("isDrawingToScreen"))
	return rv
}


// A Boolean value that indicates whether the drawing destination is the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/isdrawingtoscreen
func (g_ GraphicsContext) SetIsDrawingToScreen(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIsDrawingToScreen:"), value)
}


// A Boolean value that indicates the graphics context’s flipped state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/isflipped
func (g_ GraphicsContext) IsFlipped() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](g_.ID, objc.Sel("isFlipped"))
	return rv
}


// A Boolean value that indicates the graphics context’s flipped state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/isflipped
func (g_ GraphicsContext) SetIsFlipped(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIsFlipped:"), value)
}


