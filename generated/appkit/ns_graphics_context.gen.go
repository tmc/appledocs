// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/coregraphics"
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
	SaveGraphicsState()
}

// An object that represents a graphics context.
//
// You can think of a graphics context as a destination to which drawing and graphics state operations are sent for execution. Each graphics context contains its own graphics environment and state. The class is an abstract superclass for destination-specific graphics contexts. You obtain instances of concrete subclasses with the class methods , , , , and . At any time there is the notion of the current context. The current context for the current thread may be set using . Graphics contexts are maintained on a stack. You push a graphics context onto the stack by sending it a message, and pop it off the stack by sending it a message. By sending to a graphics context object you remove it from the stack, and the next graphics context on the stack becomes the current graphics context.
//
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
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/init(attributes:)
func NewGraphicsContextWithAttributes(attributes unsafe.Pointer) GraphicsContext {
	rv := objc.Send[GraphicsContext](objc.ID(getGraphicsContextClass().class), objc.Sel("graphicsContextWithAttributes:"), attributes)
	return rv
}



// Creates a new graphics context using the specified bitmap image representation object as the context destination.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/init(bitmapImageRep:)
func NewGraphicsContextWithBitmapImageRep(bitmapRep unsafe.Pointer) GraphicsContext {
	rv := objc.Send[GraphicsContext](objc.ID(getGraphicsContextClass().class), objc.Sel("graphicsContextWithBitmapImageRep:"), bitmapRep)
	return rv
}


// Creates a graphics context using the specified attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/init(attributes:)
func (gc _GraphicsContextClass) GraphicsContextWithAttributes(attributes unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("graphicsContextWithAttributes:"), attributes)
	return rv
}

// Creates a new graphics context using the specified bitmap image representation object as the context destination.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/init(bitmapImageRep:)
func (gc _GraphicsContextClass) GraphicsContextWithBitmapImageRep(bitmapRep unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("graphicsContextWithBitmapImageRep:"), bitmapRep)
	return rv
}

// Saves the graphics state of the current graphics context.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/saveGraphicsState()-swift.type.method
func (gc _GraphicsContextClass) SaveGraphicsState() {
	objc.Send[objc.ID](objc.ID(gc.class), objc.Sel("saveGraphicsState"))
}

// Makes the graphics context of the specified graphics state current, and resets graphics state.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/setGraphicsState(_:)
func (gc _GraphicsContextClass) SetGraphicsState(gState int) {
	objc.Send[objc.ID](objc.ID(gc.class), objc.Sel("setGraphicsState:"), gState)
}

// Returns the current graphics context of the current thread.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/current
func (gc _GraphicsContextClass) CurrentContext() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("currentContext"))
	return rv
}
// Saves the current graphics state and creates a new graphics state on the top of the stack.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/saveGraphicsState()-swift.method
func (g_ GraphicsContext) SaveGraphicsState() {
	objc.Send[objc.ID](g_.ID, objc.Sel("saveGraphicsState"))
}

// The Core Graphics context, which is a low-level, platform-specific graphics context.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/cgContext
func (g_ GraphicsContext) CGContext() coregraphics.CGContextRef {
	rv := objc.Send[coregraphics.CGContextRef](g_.ID, objc.Sel("CGContext"))
	return rv
}

// A context for Core Image objects that you can use to render into the graphics context.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/ciContext
func (g_ GraphicsContext) CIContext() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("CIContext"))
	return rv
}

// The color rendering intent in the graphics context’s graphics state.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/colorRenderingIntent
func (g_ GraphicsContext) ColorRenderingIntent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("colorRenderingIntent"))
	return rv
}


// SetColorRenderingIntent sets the value of the colorRenderingIntent property.
// The color rendering intent in the graphics context’s graphics state.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/colorRenderingIntent
func (g_ GraphicsContext) SetColorRenderingIntent(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setColorRenderingIntent:"), value)
}

// The graphics context’s global compositing operation setting.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/compositingOperation
func (g_ GraphicsContext) CompositingOperation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("compositingOperation"))
	return rv
}


// SetCompositingOperation sets the value of the compositingOperation property.
// The graphics context’s global compositing operation setting.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/compositingOperation
func (g_ GraphicsContext) SetCompositingOperation(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setCompositingOperation:"), value)
}

// Returns the current graphics context of the current thread.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/current
func (g_ GraphicsContext) CurrentContext() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("currentContext"))
	return rv
}


// SetCurrentContext sets the value of the currentContext property.
// Returns the current graphics context of the current thread.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/current
func (g_ GraphicsContext) SetCurrentContext(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setCurrentContext:"), value)
}

// A Boolean value that indicates whether the drawing destination is the screen.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/isDrawingToScreen
func (g_ GraphicsContext) DrawingToScreen() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("drawingToScreen"))
	return rv
}

// A Boolean value that indicates the graphics context’s flipped state.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/isFlipped
func (g_ GraphicsContext) Flipped() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("flipped"))
	return rv
}

// The amount to offset the pattern color when filling the graphics context.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/patternPhase
func (g_ GraphicsContext) PatternPhase() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](g_.ID, objc.Sel("patternPhase"))
	return rv
}


// SetPatternPhase sets the value of the patternPhase property.
// The amount to offset the pattern color when filling the graphics context.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/patternPhase
func (g_ GraphicsContext) SetPatternPhase(value coregraphics.CGPoint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPatternPhase:"), value)
}


