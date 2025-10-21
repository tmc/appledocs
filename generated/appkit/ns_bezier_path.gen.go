// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [BezierPath] class.
var (
	BezierPathClass     _BezierPathClass
	BezierPathClassOnce sync.Once
)

func getBezierPathClass() _BezierPathClass {
	BezierPathClassOnce.Do(func() {
		BezierPathClass = _BezierPathClass{objc.GetClass("NSBezierPath")}
	})
	return BezierPathClass
}

type _BezierPathClass struct {
	class objc.Class
}

// An interface definition for the [BezierPath] class.
type IBezierPath interface {
	objectivec.IObject
	RelativeCurveToPointControlPoint(endPoint coregraphics.CGPoint, controlPoint coregraphics.CGPoint)
}

// An object that can create paths using PostScript-style commands.
//
// Paths consist of straight and curved line segments joined together. Paths can form recognizable shapes such as rectangles, ovals, arcs, and glyphs; they can also form complex polygons using either straight or curved line segments. A single path can be closed by connecting its two endpoints, or it can be left open. An object can contain multiple disconnected paths, whether they are closed or open. Each of these paths is referred to as a subpath. The subpaths of a Bézier path object must be manipulated as a group. The only way to manipulate subpaths individually is to create separate objects for each. For a given object, you can stroke the path’s outline or fill the region occupied by the path. You can also use the path as a clipping region for views or other regions. Using methods of , you can also perform hit detection on the filled or stroked path. Hit detection is needed to implement interactive graphics, as in rubber banding and dragging operations. The current graphics context is automatically saved and restored for all drawing operations involving Bézier path objects, so your application does not need to worry about the graphics settings changing across invocations.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath
type BezierPath struct {
	objectivec.Object
}

// BezierPathFrom constructs a [BezierPath] from an unsafe.Pointer.
//
// An object that can create paths using PostScript-style commands.
func BezierPathFrom(ptr unsafe.Pointer) BezierPath {
	return BezierPath{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (bc _BezierPathClass) Alloc() BezierPath {
	rv := objc.Send[BezierPath](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BezierPathClass) New() BezierPath {
	rv := objc.Send[BezierPath](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BezierPath) Init() BezierPath {
	rv := objc.Send[BezierPath](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BezierPath) Autorelease() BezierPath {
	rv := objc.Send[BezierPath](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBezierPath creates a new BezierPath instance.
func NewBezierPath() BezierPath {
	return getBezierPathClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/relativeCurve(to:controlPoint:)
func (b_ BezierPath) RelativeCurveToPointControlPoint(endPoint coregraphics.CGPoint, controlPoint coregraphics.CGPoint) {
	objc.Send[objc.ID](b_.ID, objc.Sel("relativeCurveToPoint:controlPoint:"), endPoint, controlPoint)
}

// The accuracy with which curves are rendered.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/flatness
func (b_ BezierPath) Flatness() float64 {
	rv := objc.Send[float64](b_.ID, objc.Sel("flatness"))
	return rv
}


// SetFlatness sets the value of the flatness property.
// The accuracy with which curves are rendered.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/flatness
func (b_ BezierPath) SetFlatness(value float64) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setFlatness:"), value)
}



