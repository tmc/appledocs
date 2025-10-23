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
	// properties:
	Bounds() coregraphics.CGRect
	CGPath() coregraphics.PathRef /* not a class type */
	SetCGPath(value coregraphics.PathRef /* not a class type */)
	ControlPointBounds() coregraphics.CGRect
	CurrentPoint() coregraphics.CGPoint
	ElementCount() int /* primitive/slice/pointer. */
	Flatness() float64 /* primitive/slice/pointer. */
	SetFlatness(value float64 /* primitive/slice/pointer. */)
	BezierPathByFlatteningPath() IBezierPath
	Empty() bool /* primitive/slice/pointer. */
	LineCapStyle() LineCapStyle
	SetLineCapStyle(value LineCapStyle)
	LineJoinStyle() LineJoinStyle
	SetLineJoinStyle(value LineJoinStyle)
	LineWidth() float64 /* primitive/slice/pointer. */
	SetLineWidth(value float64 /* primitive/slice/pointer. */)
	MiterLimit() float64 /* primitive/slice/pointer. */
	SetMiterLimit(value float64 /* primitive/slice/pointer. */)
	BezierPathByReversingPath() IBezierPath
	WindingRule() WindingRule
	SetWindingRule(value WindingRule)
	Flattened() IBezierPath
	SetFlattened(value IBezierPath)
	IsEmpty() bool /* primitive/slice/pointer. */
	SetIsEmpty(value bool /* primitive/slice/pointer. */)
	Reversed() IBezierPath
	SetReversed(value IBezierPath)
	// methods:
	AddClip()
	AppendBezierPath(path IBezierPath)
	AppendBezierPathWithCGGlyphInFont(glyph objc.IObject /* cross-framework Glyph */, font IFont)
	AppendBezierPathWithCGGlyphsCountInFont(glyphs objc.IObject /* cross-framework Glyph */, count int /* primitive/slice/pointer. */, font IFont)
	AppendBezierPathWithArcFromPointToPointRadius(point1 coregraphics.CGPoint, point2 coregraphics.CGPoint, radius float64 /* primitive/slice/pointer. */)
	AppendBezierPathWithArcWithCenterRadiusStartAngleEndAngle(center coregraphics.CGPoint, radius float64 /* primitive/slice/pointer. */, startAngle float64 /* primitive/slice/pointer. */, endAngle float64 /* primitive/slice/pointer. */)
	AppendBezierPathWithArcWithCenterRadiusStartAngleEndAngleClockwise(center coregraphics.CGPoint, radius float64 /* primitive/slice/pointer. */, startAngle float64 /* primitive/slice/pointer. */, endAngle float64 /* primitive/slice/pointer. */, clockwise bool /* primitive/slice/pointer. */)
	AppendBezierPathWithOvalInRect(rect coregraphics.CGRect)
	AppendBezierPathWithPointsCount(points PointArray /* not a class type */, count int /* primitive/slice/pointer. */)
	AppendBezierPathWithRect(rect coregraphics.CGRect)
	AppendBezierPathWithRoundedRectXRadiusYRadius(rect coregraphics.CGRect, xRadius float64 /* primitive/slice/pointer. */, yRadius float64 /* primitive/slice/pointer. */)
	ClosePath()
	ContainsPoint(point coregraphics.CGPoint) bool /* primitive/slice/pointer. */
	CurveToPointControlPoint1ControlPoint2(endPoint coregraphics.CGPoint, controlPoint1 coregraphics.CGPoint, controlPoint2 coregraphics.CGPoint)
	CurveToPointControlPoint(endPoint coregraphics.CGPoint, controlPoint coregraphics.CGPoint)
	ElementAtIndex(index int /* primitive/slice/pointer. */) BezierPathElement
	ElementAtIndexAssociatedPoints(index int /* primitive/slice/pointer. */, points PointArray /* not a class type */) BezierPathElement
	Fill()
	GetLineDashCountPhase(pattern coregraphics.float64 /* primitive/slice/pointer. */, count Integer /* not a class type */, phase coregraphics.float64 /* primitive/slice/pointer. */)
	LineToPoint(point coregraphics.CGPoint)
	MoveToPoint(point coregraphics.CGPoint)
	RelativeCurveToPointControlPoint1ControlPoint2(endPoint coregraphics.CGPoint, controlPoint1 coregraphics.CGPoint, controlPoint2 coregraphics.CGPoint)
	RelativeCurveToPointControlPoint(endPoint coregraphics.CGPoint, controlPoint coregraphics.CGPoint)
	RelativeLineToPoint(point coregraphics.CGPoint)
	RelativeMoveToPoint(point coregraphics.CGPoint)
	RemoveAllPoints()
	SetAssociatedPointsAtIndex(points PointArray /* not a class type */, index int /* primitive/slice/pointer. */)
	SetClip()
	SetLineDashCountPhase(pattern coregraphics.float64 /* primitive/slice/pointer. */, count int /* primitive/slice/pointer. */, phase float64 /* primitive/slice/pointer. */)
	Stroke()
	TransformUsingAffineTransform(transform objc.IObject /* cross-framework AffineTransform */)
}

// An object that can create paths using PostScript-style commands.
//
// Paths consist of straight and curved line segments joined together. Paths can form recognizable shapes such as rectangles, ovals, arcs, and glyphs; they can also form complex polygons using either straight or curved line segments. A single path can be closed by connecting its two endpoints, or it can be left open. An object can contain multiple disconnected paths, whether they are closed or open. Each of these paths is referred to as a subpath. The subpaths of a Bézier path object must be manipulated as a group. The only way to manipulate subpaths individually is to create separate objects for each. For a given object, you can stroke the path’s outline or fill the region occupied by the path. You can also use the path as a clipping region for views or other regions. Using methods of , you can also perform hit detection on the filled or stroked path. Hit detection is needed to implement interactive graphics, as in rubber banding and dragging operations. The current graphics context is automatically saved and restored for all drawing operations involving Bézier path objects, so your application does not need to worry about the graphics settings changing across invocations.


// An object that can create paths using PostScript-style commands.
//
// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/init(cgPath:)
func NewBezierPathWithCGPath(cgPath coregraphics.PathRef /* not a class type */) BezierPath {
	rv := objc.Send[BezierPath](objc.ID(getBezierPathClass().class), objc.Sel("bezierPathWithCGPath:"), cgPath)
	return rv
}


// Creates and returns a new Bézier path object initialized with an oval path inscribed in the specified rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/init(ovalIn:)
func NewBezierPathWithOvalInRect(rect coregraphics.CGRect) BezierPath {
	rv := objc.Send[BezierPath](objc.ID(getBezierPathClass().class), objc.Sel("bezierPathWithOvalInRect:"), rect)
	return rv
}


// Creates and returns a new Bézier path object initialized with a rectangular path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/init(rect:)
func NewBezierPathWithRect(rect coregraphics.CGRect) BezierPath {
	rv := objc.Send[BezierPath](objc.ID(getBezierPathClass().class), objc.Sel("bezierPathWithRect:"), rect)
	return rv
}


// Creates and returns a new Bézier path object initialized with a rounded rectangular path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/init(roundedRect:xRadius:yRadius:)
func NewBezierPathWithRoundedRectXRadiusYRadius(rect coregraphics.CGRect, xRadius float64 /* primitive/slice/pointer. */, yRadius float64 /* primitive/slice/pointer. */) BezierPath {
	rv := objc.Send[BezierPath](objc.ID(getBezierPathClass().class), objc.Sel("bezierPathWithRoundedRect:xRadius:yRadius:"), rect, xRadius, yRadius)
	return rv
}



// Creates and returns a new Bézier path object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/bezierPath
func (bc _BezierPathClass) BezierPath() IBezierPath {
	rv := objc.Send[BezierPath](objc.ID(bc.class), objc.Sel("bezierPath"))
	return rv
}


// Intersects the specified rectangle with the clipping path of the current graphics context and makes the resulting shape the current clipping path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/clip(_:)
func (bc _BezierPathClass) ClipRect(rect coregraphics.CGRect) {
	objc.Send[objc.ID](objc.ID(bc.class), objc.Sel("clipRect:"), rect)
}


// Draws a set of packed glyphs at the specified point in the current coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/drawPackedGlyphs(_:at:)
func (bc _BezierPathClass) DrawPackedGlyphsAtPoint(packedGlyphs unsafe.Pointer, point coregraphics.CGPoint) {
	objc.Send[objc.ID](objc.ID(bc.class), objc.Sel("drawPackedGlyphs:atPoint:"), packedGlyphs, point)
}


// Fills the specified rectangular path with the current fill color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/fill(_:)
func (bc _BezierPathClass) FillRect(rect coregraphics.CGRect) {
	objc.Send[objc.ID](objc.ID(bc.class), objc.Sel("fillRect:"), rect)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/init(cgPath:)
func (bc _BezierPathClass) BezierPathWithCGPath(cgPath coregraphics.PathRef /* not a class type */) IBezierPath {
	rv := objc.Send[BezierPath](objc.ID(bc.class), objc.Sel("bezierPathWithCGPath:"), cgPath)
	return rv
}


// Creates and returns a new Bézier path object initialized with an oval path inscribed in the specified rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/init(ovalIn:)
func (bc _BezierPathClass) BezierPathWithOvalInRect(rect coregraphics.CGRect) IBezierPath {
	rv := objc.Send[BezierPath](objc.ID(bc.class), objc.Sel("bezierPathWithOvalInRect:"), rect)
	return rv
}


// Creates and returns a new Bézier path object initialized with a rectangular path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/init(rect:)
func (bc _BezierPathClass) BezierPathWithRect(rect coregraphics.CGRect) IBezierPath {
	rv := objc.Send[BezierPath](objc.ID(bc.class), objc.Sel("bezierPathWithRect:"), rect)
	return rv
}


// Creates and returns a new Bézier path object initialized with a rounded rectangular path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/init(roundedRect:xRadius:yRadius:)
func (bc _BezierPathClass) BezierPathWithRoundedRectXRadiusYRadius(rect coregraphics.CGRect, xRadius float64 /* primitive/slice/pointer. */, yRadius float64 /* primitive/slice/pointer. */) IBezierPath {
	rv := objc.Send[BezierPath](objc.ID(bc.class), objc.Sel("bezierPathWithRoundedRect:xRadius:yRadius:"), rect, xRadius, yRadius)
	return rv
}


// Strokes the path of the specified rectangle using the current stroke color and the default drawing attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/stroke(_:)
func (bc _BezierPathClass) StrokeRect(rect coregraphics.CGRect) {
	objc.Send[objc.ID](objc.ID(bc.class), objc.Sel("strokeRect:"), rect)
}


// Strokes a line between two points using the current stroke color and the default drawing attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/strokeLine(from:to:)
func (bc _BezierPathClass) StrokeLineFromPointToPoint(point1 coregraphics.CGPoint, point2 coregraphics.CGPoint) {
	objc.Send[objc.ID](objc.ID(bc.class), objc.Sel("strokeLineFromPoint:toPoint:"), point1, point2)
}


// The default flatness value for all paths.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/defaultFlatness
func (bc _BezierPathClass) DefaultFlatness() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](objc.ID(bc.class), objc.Sel("defaultFlatness"))
	return rv
}

// Returns the default line cap style for all paths.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/defaultLineCapStyle
func (bc _BezierPathClass) DefaultLineCapStyle() LineCapStyle {
	rv := objc.Send[LineCapStyle](objc.ID(bc.class), objc.Sel("defaultLineCapStyle"))
	return rv
}

// Returns the default line join style for all paths.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/defaultLineJoinStyle
func (bc _BezierPathClass) DefaultLineJoinStyle() LineJoinStyle {
	rv := objc.Send[LineJoinStyle](objc.ID(bc.class), objc.Sel("defaultLineJoinStyle"))
	return rv
}

// Returns the default line width for the all paths.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/defaultLineWidth
func (bc _BezierPathClass) DefaultLineWidth() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](objc.ID(bc.class), objc.Sel("defaultLineWidth"))
	return rv
}

// Returns the default miter limit for all paths.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/defaultMiterLimit
func (bc _BezierPathClass) DefaultMiterLimit() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](objc.ID(bc.class), objc.Sel("defaultMiterLimit"))
	return rv
}

// Returns the default winding rule used to fill all paths.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/defaultWindingRule
func (bc _BezierPathClass) DefaultWindingRule() WindingRule {
	rv := objc.Send[WindingRule](objc.ID(bc.class), objc.Sel("defaultWindingRule"))
	return rv
}

// Intersects the area enclosed by the path with the clipping path of the current graphics context and makes the resulting shape the current clipping path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/addClip()
func (b_ BezierPath) AddClip() {
	objc.Send[objc.ID](b_.ID, objc.Sel("addClip"))
}


// Appends the contents of the specified path object to the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/append(_:)
func (b_ BezierPath) AppendBezierPath(path IBezierPath) {
	objc.Send[objc.ID](b_.ID, objc.Sel("appendBezierPath:"), path)
}


// Appends an outline of the specified glyph to the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/append(withCGGlyph:in:)
func (b_ BezierPath) AppendBezierPathWithCGGlyphInFont(glyph objc.IObject /* cross-framework Glyph */, font IFont) {
	objc.Send[objc.ID](b_.ID, objc.Sel("appendBezierPathWithCGGlyph:inFont:"), glyph, font)
}


// Appends the outlines of the specified glyphs to the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/append(withCGGlyphs:count:in:)
func (b_ BezierPath) AppendBezierPathWithCGGlyphsCountInFont(glyphs objc.IObject /* cross-framework Glyph */, count int /* primitive/slice/pointer. */, font IFont) {
	objc.Send[objc.ID](b_.ID, objc.Sel("appendBezierPathWithCGGlyphs:count:inFont:"), glyphs, count, font)
}


// Appends an arc to the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/appendArc(from:to:radius:)
func (b_ BezierPath) AppendBezierPathWithArcFromPointToPointRadius(point1 coregraphics.CGPoint, point2 coregraphics.CGPoint, radius float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("appendBezierPathWithArcFromPoint:toPoint:radius:"), point1, point2, radius)
}


// Appends an arc of a circle to the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/appendArc(withCenter:radius:startAngle:endAngle:)
func (b_ BezierPath) AppendBezierPathWithArcWithCenterRadiusStartAngleEndAngle(center coregraphics.CGPoint, radius float64 /* primitive/slice/pointer. */, startAngle float64 /* primitive/slice/pointer. */, endAngle float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("appendBezierPathWithArcWithCenter:radius:startAngle:endAngle:"), center, radius, startAngle, endAngle)
}


// Appends an arc of a circle to the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/appendArc(withCenter:radius:startAngle:endAngle:clockwise:)
func (b_ BezierPath) AppendBezierPathWithArcWithCenterRadiusStartAngleEndAngleClockwise(center coregraphics.CGPoint, radius float64 /* primitive/slice/pointer. */, startAngle float64 /* primitive/slice/pointer. */, endAngle float64 /* primitive/slice/pointer. */, clockwise bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("appendBezierPathWithArcWithCenter:radius:startAngle:endAngle:clockwise:"), center, radius, startAngle, endAngle, clockwise)
}


// Appends an oval path to the path, inscribing the oval in the specified rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/appendOval(in:)
func (b_ BezierPath) AppendBezierPathWithOvalInRect(rect coregraphics.CGRect) {
	objc.Send[objc.ID](b_.ID, objc.Sel("appendBezierPathWithOvalInRect:"), rect)
}


// Appends a series of line segments to the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/appendPoints(_:count:)
func (b_ BezierPath) AppendBezierPathWithPointsCount(points PointArray /* not a class type */, count int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("appendBezierPathWithPoints:count:"), points, count)
}


// Appends a rectangular path to the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/appendRect(_:)
func (b_ BezierPath) AppendBezierPathWithRect(rect coregraphics.CGRect) {
	objc.Send[objc.ID](b_.ID, objc.Sel("appendBezierPathWithRect:"), rect)
}


// Appends a rounded rectangular path to the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/appendRoundedRect(_:xRadius:yRadius:)
func (b_ BezierPath) AppendBezierPathWithRoundedRectXRadiusYRadius(rect coregraphics.CGRect, xRadius float64 /* primitive/slice/pointer. */, yRadius float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("appendBezierPathWithRoundedRect:xRadius:yRadius:"), rect, xRadius, yRadius)
}


// Closes the most recently added subpath.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/close()
func (b_ BezierPath) ClosePath() {
	objc.Send[objc.ID](b_.ID, objc.Sel("closePath"))
}


// Returns a Boolean value that indicates whether the path contains the specified point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/contains(_:)
func (b_ BezierPath) ContainsPoint(point coregraphics.CGPoint) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](b_.ID, objc.Sel("containsPoint:"), point)
	return rv
}


// Adds a Bezier cubic curve to the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/curve(to:controlPoint1:controlPoint2:)
func (b_ BezierPath) CurveToPointControlPoint1ControlPoint2(endPoint coregraphics.CGPoint, controlPoint1 coregraphics.CGPoint, controlPoint2 coregraphics.CGPoint) {
	objc.Send[objc.ID](b_.ID, objc.Sel("curveToPoint:controlPoint1:controlPoint2:"), endPoint, controlPoint1, controlPoint2)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/curve(to:controlPoint:)
func (b_ BezierPath) CurveToPointControlPoint(endPoint coregraphics.CGPoint, controlPoint coregraphics.CGPoint) {
	objc.Send[objc.ID](b_.ID, objc.Sel("curveToPoint:controlPoint:"), endPoint, controlPoint)
}


// Returns the type of path element at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/element(at:)
func (b_ BezierPath) ElementAtIndex(index int /* primitive/slice/pointer. */) BezierPathElement {
	rv := objc.Send[BezierPathElement](b_.ID, objc.Sel("elementAtIndex:"), index)
	return rv
}


// Gets the element type and (and optionally) the associated points for the path element at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/element(at:associatedPoints:)
func (b_ BezierPath) ElementAtIndexAssociatedPoints(index int /* primitive/slice/pointer. */, points PointArray /* not a class type */) BezierPathElement {
	rv := objc.Send[BezierPathElement](b_.ID, objc.Sel("elementAtIndex:associatedPoints:"), index, points)
	return rv
}


// Paints the region enclosed by the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/fill()
func (b_ BezierPath) Fill() {
	objc.Send[objc.ID](b_.ID, objc.Sel("fill"))
}


// Returns the line-stroking pattern for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/getLineDash(_:count:phase:)
func (b_ BezierPath) GetLineDashCountPhase(pattern coregraphics.float64 /* primitive/slice/pointer. */, count Integer /* not a class type */, phase coregraphics.float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("getLineDash:count:phase:"), pattern, count, phase)
}


// Appends a straight line to the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/line(to:)
func (b_ BezierPath) LineToPoint(point coregraphics.CGPoint) {
	objc.Send[objc.ID](b_.ID, objc.Sel("lineToPoint:"), point)
}


// Moves the path’s current point to the specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/move(to:)
func (b_ BezierPath) MoveToPoint(point coregraphics.CGPoint) {
	objc.Send[objc.ID](b_.ID, objc.Sel("moveToPoint:"), point)
}


// Adds a Bezier cubic curve to the path from the current point to a new location, which is specified as a relative distance from the current point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/relativeCurve(to:controlPoint1:controlPoint2:)
func (b_ BezierPath) RelativeCurveToPointControlPoint1ControlPoint2(endPoint coregraphics.CGPoint, controlPoint1 coregraphics.CGPoint, controlPoint2 coregraphics.CGPoint) {
	objc.Send[objc.ID](b_.ID, objc.Sel("relativeCurveToPoint:controlPoint1:controlPoint2:"), endPoint, controlPoint1, controlPoint2)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/relativeCurve(to:controlPoint:)
func (b_ BezierPath) RelativeCurveToPointControlPoint(endPoint coregraphics.CGPoint, controlPoint coregraphics.CGPoint) {
	objc.Send[objc.ID](b_.ID, objc.Sel("relativeCurveToPoint:controlPoint:"), endPoint, controlPoint)
}


// Appends a straight line segment to the path starting at the current point and moving towards the specified point, relative to the current location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/relativeLine(to:)
func (b_ BezierPath) RelativeLineToPoint(point coregraphics.CGPoint) {
	objc.Send[objc.ID](b_.ID, objc.Sel("relativeLineToPoint:"), point)
}


// Moves the path’s current point to a new point whose location is the specified distance from the current point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/relativeMove(to:)
func (b_ BezierPath) RelativeMoveToPoint(point coregraphics.CGPoint) {
	objc.Send[objc.ID](b_.ID, objc.Sel("relativeMoveToPoint:"), point)
}


// Removes all path elements from the path, effectively clearing the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/removeAllPoints()
func (b_ BezierPath) RemoveAllPoints() {
	objc.Send[objc.ID](b_.ID, objc.Sel("removeAllPoints"))
}


// Changes the points associated with the specified path element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/setAssociatedPoints(_:at:)
func (b_ BezierPath) SetAssociatedPointsAtIndex(points PointArray /* not a class type */, index int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAssociatedPoints:atIndex:"), points, index)
}


// Replaces the clipping path of the current graphics context with the area inside the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/setClip()
func (b_ BezierPath) SetClip() {
	objc.Send[objc.ID](b_.ID, objc.Sel("setClip"))
}


// Sets the line-stroking pattern for the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/setLineDash(_:count:phase:)
func (b_ BezierPath) SetLineDashCountPhase(pattern coregraphics.float64 /* primitive/slice/pointer. */, count int /* primitive/slice/pointer. */, phase float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setLineDash:count:phase:"), pattern, count, phase)
}


// Draws a line along the path using the current stroke color and drawing attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/stroke()
func (b_ BezierPath) Stroke() {
	objc.Send[objc.ID](b_.ID, objc.Sel("stroke"))
}


// Transforms all points in the path using the specified transform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/transform(using:)
func (b_ BezierPath) TransformUsingAffineTransform(transform objc.IObject /* cross-framework AffineTransform */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("transformUsingAffineTransform:"), transform)
}


// The bounding box of the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/bounds
func (b_ BezierPath) Bounds() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](b_.ID, objc.Sel("bounds"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/cgPath
func (b_ BezierPath) CGPath() coregraphics.PathRef /* not a class type */ {
	rv := objc.Send[coregraphics.PathRef](b_.ID, objc.Sel("CGPath"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/cgPath
func (b_ BezierPath) SetCGPath(value coregraphics.PathRef /* not a class type */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setCGPath:"), value)
}


// The bounding box of the path, including any control points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/controlPointBounds
func (b_ BezierPath) ControlPointBounds() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](b_.ID, objc.Sel("controlPointBounds"))
	return rv
}


// The current point (the trailing point or ending point in the most recently added segment).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/currentPoint
func (b_ BezierPath) CurrentPoint() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](b_.ID, objc.Sel("currentPoint"))
	return rv
}


// The default flatness value for all paths.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/defaultFlatness
func (b_ BezierPath) DefaultFlatness() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](b_.ID, objc.Sel("defaultFlatness"))
	return rv
}


// The default flatness value for all paths.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/defaultFlatness
func (b_ BezierPath) SetDefaultFlatness(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setDefaultFlatness:"), value)
}


// Returns the default line cap style for all paths.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/defaultLineCapStyle
func (b_ BezierPath) DefaultLineCapStyle() LineCapStyle {
	rv := objc.Send[LineCapStyle](b_.ID, objc.Sel("defaultLineCapStyle"))
	return rv
}


// Returns the default line cap style for all paths.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/defaultLineCapStyle
func (b_ BezierPath) SetDefaultLineCapStyle(value LineCapStyle) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setDefaultLineCapStyle:"), value)
}


// Returns the default line join style for all paths.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/defaultLineJoinStyle
func (b_ BezierPath) DefaultLineJoinStyle() LineJoinStyle {
	rv := objc.Send[LineJoinStyle](b_.ID, objc.Sel("defaultLineJoinStyle"))
	return rv
}


// Returns the default line join style for all paths.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/defaultLineJoinStyle
func (b_ BezierPath) SetDefaultLineJoinStyle(value LineJoinStyle) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setDefaultLineJoinStyle:"), value)
}


// Returns the default line width for the all paths.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/defaultLineWidth
func (b_ BezierPath) DefaultLineWidth() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](b_.ID, objc.Sel("defaultLineWidth"))
	return rv
}


// Returns the default line width for the all paths.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/defaultLineWidth
func (b_ BezierPath) SetDefaultLineWidth(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setDefaultLineWidth:"), value)
}


// Returns the default miter limit for all paths.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/defaultMiterLimit
func (b_ BezierPath) DefaultMiterLimit() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](b_.ID, objc.Sel("defaultMiterLimit"))
	return rv
}


// Returns the default miter limit for all paths.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/defaultMiterLimit
func (b_ BezierPath) SetDefaultMiterLimit(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setDefaultMiterLimit:"), value)
}


// Returns the default winding rule used to fill all paths.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/defaultWindingRule
func (b_ BezierPath) DefaultWindingRule() WindingRule {
	rv := objc.Send[WindingRule](b_.ID, objc.Sel("defaultWindingRule"))
	return rv
}


// Returns the default winding rule used to fill all paths.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/defaultWindingRule
func (b_ BezierPath) SetDefaultWindingRule(value WindingRule) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setDefaultWindingRule:"), value)
}


// The total number of path elements in the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/elementCount
func (b_ BezierPath) ElementCount() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](b_.ID, objc.Sel("elementCount"))
	return rv
}


// The accuracy with which curves are rendered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/flatness
func (b_ BezierPath) Flatness() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](b_.ID, objc.Sel("flatness"))
	return rv
}


// The accuracy with which curves are rendered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/flatness
func (b_ BezierPath) SetFlatness(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setFlatness:"), value)
}


// A flattened version of the path object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/flattened
func (b_ BezierPath) BezierPathByFlatteningPath() IBezierPath {
	rv := objc.Send[BezierPath](b_.ID, objc.Sel("bezierPathByFlatteningPath"))
	return rv
}


// A Boolean value that indicates whether the path is empty.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/isEmpty
func (b_ BezierPath) Empty() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](b_.ID, objc.Sel("empty"))
	return rv
}


// The line cap style for the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/lineCapStyle-swift.property
func (b_ BezierPath) LineCapStyle() LineCapStyle {
	rv := objc.Send[LineCapStyle](b_.ID, objc.Sel("lineCapStyle"))
	return rv
}


// The line cap style for the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/lineCapStyle-swift.property
func (b_ BezierPath) SetLineCapStyle(value LineCapStyle) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setLineCapStyle:"), value)
}


// The line join style for the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/lineJoinStyle-swift.property
func (b_ BezierPath) LineJoinStyle() LineJoinStyle {
	rv := objc.Send[LineJoinStyle](b_.ID, objc.Sel("lineJoinStyle"))
	return rv
}


// The line join style for the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/lineJoinStyle-swift.property
func (b_ BezierPath) SetLineJoinStyle(value LineJoinStyle) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setLineJoinStyle:"), value)
}


// The width of stroked path lines.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/lineWidth
func (b_ BezierPath) LineWidth() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](b_.ID, objc.Sel("lineWidth"))
	return rv
}


// The width of stroked path lines.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/lineWidth
func (b_ BezierPath) SetLineWidth(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setLineWidth:"), value)
}


// The limit at which miter joins are converted to bevel joins.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/miterLimit
func (b_ BezierPath) MiterLimit() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](b_.ID, objc.Sel("miterLimit"))
	return rv
}


// The limit at which miter joins are converted to bevel joins.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/miterLimit
func (b_ BezierPath) SetMiterLimit(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setMiterLimit:"), value)
}


// A path containing the reversed contents of the current path object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/reversed
func (b_ BezierPath) BezierPathByReversingPath() IBezierPath {
	rv := objc.Send[BezierPath](b_.ID, objc.Sel("bezierPathByReversingPath"))
	return rv
}


// The winding rule used to fill the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/windingRule-swift.property
func (b_ BezierPath) WindingRule() WindingRule {
	rv := objc.Send[WindingRule](b_.ID, objc.Sel("windingRule"))
	return rv
}


// The winding rule used to fill the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/windingRule-swift.property
func (b_ BezierPath) SetWindingRule(value WindingRule) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setWindingRule:"), value)
}


// A flattened version of the path object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/flattened
func (b_ BezierPath) Flattened() IBezierPath {
	rv := objc.Send[BezierPath](b_.ID, objc.Sel("flattened"))
	return rv
}


// A flattened version of the path object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/flattened
func (b_ BezierPath) SetFlattened(value IBezierPath) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setFlattened:"), value)
}


// A Boolean value that indicates whether the path is empty.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/isempty
func (b_ BezierPath) IsEmpty() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](b_.ID, objc.Sel("isEmpty"))
	return rv
}


// A Boolean value that indicates whether the path is empty.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/isempty
func (b_ BezierPath) SetIsEmpty(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsEmpty:"), value)
}


// A path containing the reversed contents of the current path object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/reversed
func (b_ BezierPath) Reversed() IBezierPath {
	rv := objc.Send[BezierPath](b_.ID, objc.Sel("reversed"))
	return rv
}


// A path containing the reversed contents of the current path object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/reversed
func (b_ BezierPath) SetReversed(value IBezierPath) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setReversed:"), value)
}


