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

// The bounding box of the path.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/bounds
func (b_ BezierPath) Bounds() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](b_.ID, objc.Sel("bounds"))
	return rv
}


// SetBounds sets the value of the bounds property.
// The bounding box of the path.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/bounds
func (b_ BezierPath) SetBounds(value coregraphics.CGRect) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBounds:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/cgpath
func (b_ BezierPath) CgPath() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("cgPath"))
	return rv
}


// SetCgPath sets the value of the cgPath property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/cgpath
func (b_ BezierPath) SetCgPath(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setCgPath:"), value)
}

// The bounding box of the path, including any control points.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/controlpointbounds
func (b_ BezierPath) ControlPointBounds() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](b_.ID, objc.Sel("controlPointBounds"))
	return rv
}


// SetControlPointBounds sets the value of the controlPointBounds property.
// The bounding box of the path, including any control points.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/controlpointbounds
func (b_ BezierPath) SetControlPointBounds(value coregraphics.CGRect) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setControlPointBounds:"), value)
}

// The current point (the trailing point or ending point in the most recently added segment).
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/currentpoint
func (b_ BezierPath) CurrentPoint() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](b_.ID, objc.Sel("currentPoint"))
	return rv
}


// SetCurrentPoint sets the value of the currentPoint property.
// The current point (the trailing point or ending point in the most recently added segment).

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/currentpoint
func (b_ BezierPath) SetCurrentPoint(value coregraphics.CGPoint) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setCurrentPoint:"), value)
}

// The total number of path elements in the path.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/elementcount
func (b_ BezierPath) ElementCount() int {
	rv := objc.Send[int](b_.ID, objc.Sel("elementCount"))
	return rv
}


// SetElementCount sets the value of the elementCount property.
// The total number of path elements in the path.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/elementcount
func (b_ BezierPath) SetElementCount(value int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setElementCount:"), value)
}

// A flattened version of the path object.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/flattened
func (b_ BezierPath) Flattened() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("flattened"))
	return rv
}


// SetFlattened sets the value of the flattened property.
// A flattened version of the path object.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/flattened
func (b_ BezierPath) SetFlattened(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setFlattened:"), value)
}

// A Boolean value that indicates whether the path is empty.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/isempty
func (b_ BezierPath) IsEmpty() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isEmpty"))
	return rv
}


// SetIsEmpty sets the value of the isEmpty property.
// A Boolean value that indicates whether the path is empty.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/isempty
func (b_ BezierPath) SetIsEmpty(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsEmpty:"), value)
}

// The line cap style for the path.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/linecapstyle-swift.property
func (b_ BezierPath) LineCapStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("lineCapStyle"))
	return rv
}


// SetLineCapStyle sets the value of the lineCapStyle property.
// The line cap style for the path.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/linecapstyle-swift.property
func (b_ BezierPath) SetLineCapStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setLineCapStyle:"), value)
}

// The line join style for the path.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/linejoinstyle-swift.property
func (b_ BezierPath) LineJoinStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("lineJoinStyle"))
	return rv
}


// SetLineJoinStyle sets the value of the lineJoinStyle property.
// The line join style for the path.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/linejoinstyle-swift.property
func (b_ BezierPath) SetLineJoinStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setLineJoinStyle:"), value)
}

// The width of stroked path lines.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/linewidth
func (b_ BezierPath) LineWidth() float64 {
	rv := objc.Send[float64](b_.ID, objc.Sel("lineWidth"))
	return rv
}


// SetLineWidth sets the value of the lineWidth property.
// The width of stroked path lines.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/linewidth
func (b_ BezierPath) SetLineWidth(value float64) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setLineWidth:"), value)
}

// The limit at which miter joins are converted to bevel joins.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/miterlimit
func (b_ BezierPath) MiterLimit() float64 {
	rv := objc.Send[float64](b_.ID, objc.Sel("miterLimit"))
	return rv
}


// SetMiterLimit sets the value of the miterLimit property.
// The limit at which miter joins are converted to bevel joins.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/miterlimit
func (b_ BezierPath) SetMiterLimit(value float64) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setMiterLimit:"), value)
}

// A path containing the reversed contents of the current path object.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/reversed
func (b_ BezierPath) Reversed() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("reversed"))
	return rv
}


// SetReversed sets the value of the reversed property.
// A path containing the reversed contents of the current path object.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/reversed
func (b_ BezierPath) SetReversed(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setReversed:"), value)
}

// The winding rule used to fill the path.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/windingrule-swift.property
func (b_ BezierPath) WindingRule() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("windingRule"))
	return rv
}


// SetWindingRule sets the value of the windingRule property.
// The winding rule used to fill the path.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/windingrule-swift.property
func (b_ BezierPath) SetWindingRule(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setWindingRule:"), value)
}



