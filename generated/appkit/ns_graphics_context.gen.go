// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/coreimage"
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
	Attributes() unsafe.Pointer
	SetAttributes(value unsafe.Pointer)
	CgContext() coreimage.Context
	SetCgContext(value coreimage.Context)
	CiContext() coreimage.Context
	SetCiContext(value coreimage.Context)
	ColorRenderingIntent() unsafe.Pointer
	SetColorRenderingIntent(value unsafe.Pointer)
	CompositingOperation() NSCompositingOperation
	SetCompositingOperation(value NSCompositingOperation)
	GraphicsPort() unsafe.Pointer
	SetGraphicsPort(value unsafe.Pointer)
	ImageInterpolation() unsafe.Pointer
	SetImageInterpolation(value unsafe.Pointer)
	IsDrawingToScreen() bool
	SetIsDrawingToScreen(value bool)
	IsFlipped() bool
	SetIsFlipped(value bool)
	PatternPhase() coregraphics.CGPoint
	SetPatternPhase(value coregraphics.CGPoint)
	ShouldAntialias() bool
	SetShouldAntialias(value bool)
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



// Returns the current graphics context of the current thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/current
func (gc _GraphicsContextClass) CurrentContext() GraphicsContext {
	rv := objc.Send[GraphicsContext](objc.ID(gc.class), objc.Sel("currentContext"))
	return rv
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


// The attributes used to create this instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/attributes
func (g_ GraphicsContext) Attributes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("attributes"))
	return rv
}


// The attributes used to create this instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/attributes
func (g_ GraphicsContext) SetAttributes(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setAttributes:"), value)
}


// The Core Graphics context, which is a low-level, platform-specific graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/cgcontext
func (g_ GraphicsContext) CgContext() coreimage.Context {
	rv := objc.Send[coreimage.Context](g_.ID, objc.Sel("cgContext"))
	return rv
}


// The Core Graphics context, which is a low-level, platform-specific graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/cgcontext
func (g_ GraphicsContext) SetCgContext(value coreimage.Context) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setCgContext:"), value)
}


// A context for Core Image objects that you can use to render into the graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/cicontext
func (g_ GraphicsContext) CiContext() coreimage.Context {
	rv := objc.Send[coreimage.Context](g_.ID, objc.Sel("ciContext"))
	return rv
}


// A context for Core Image objects that you can use to render into the graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/cicontext
func (g_ GraphicsContext) SetCiContext(value coreimage.Context) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setCiContext:"), value)
}


// The color rendering intent in the graphics context’s graphics state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/colorrenderingintent
func (g_ GraphicsContext) ColorRenderingIntent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("colorRenderingIntent"))
	return rv
}


// The color rendering intent in the graphics context’s graphics state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/colorrenderingintent
func (g_ GraphicsContext) SetColorRenderingIntent(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setColorRenderingIntent:"), value)
}


// The graphics context’s global compositing operation setting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/compositingoperation
func (g_ GraphicsContext) CompositingOperation() NSCompositingOperation {
	rv := objc.Send[NSCompositingOperation](g_.ID, objc.Sel("compositingOperation"))
	return rv
}


// The graphics context’s global compositing operation setting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/compositingoperation
func (g_ GraphicsContext) SetCompositingOperation(value NSCompositingOperation) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setCompositingOperation:"), value)
}


// The low-level, platform-specific graphics context represented by the graphic port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/graphicsport
func (g_ GraphicsContext) GraphicsPort() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("graphicsPort"))
	return rv
}


// The low-level, platform-specific graphics context represented by the graphic port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/graphicsport
func (g_ GraphicsContext) SetGraphicsPort(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGraphicsPort:"), value)
}


// A constant that specifies the graphics context’s interpolation, or image smoothing, behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/imageinterpolation
func (g_ GraphicsContext) ImageInterpolation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("imageInterpolation"))
	return rv
}


// A constant that specifies the graphics context’s interpolation, or image smoothing, behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/imageinterpolation
func (g_ GraphicsContext) SetImageInterpolation(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setImageInterpolation:"), value)
}


// A Boolean value that indicates whether the drawing destination is the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/isdrawingtoscreen
func (g_ GraphicsContext) IsDrawingToScreen() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("isDrawingToScreen"))
	return rv
}


// A Boolean value that indicates whether the drawing destination is the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/isdrawingtoscreen
func (g_ GraphicsContext) SetIsDrawingToScreen(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIsDrawingToScreen:"), value)
}


// A Boolean value that indicates the graphics context’s flipped state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/isflipped
func (g_ GraphicsContext) IsFlipped() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("isFlipped"))
	return rv
}


// A Boolean value that indicates the graphics context’s flipped state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/isflipped
func (g_ GraphicsContext) SetIsFlipped(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIsFlipped:"), value)
}


// The amount to offset the pattern color when filling the graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/patternphase
func (g_ GraphicsContext) PatternPhase() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](g_.ID, objc.Sel("patternPhase"))
	return rv
}


// The amount to offset the pattern color when filling the graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/patternphase
func (g_ GraphicsContext) SetPatternPhase(value coregraphics.CGPoint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPatternPhase:"), value)
}


// A Boolean value that indicates whether the graphics context uses antialiasing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/shouldantialias
func (g_ GraphicsContext) ShouldAntialias() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("shouldAntialias"))
	return rv
}


// A Boolean value that indicates whether the graphics context uses antialiasing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgraphicscontext/shouldantialias
func (g_ GraphicsContext) SetShouldAntialias(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setShouldAntialias:"), value)
}



