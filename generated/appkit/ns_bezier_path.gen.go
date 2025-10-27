// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
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
	CurrentPoint() corefoundation.CGPoint
	Bounds() corefoundation.CGRect
	SetBounds(value corefoundation.CGRect)
	CgPath() Path /* not a class type */
	SetCgPath(value Path /* not a class type */)
	ControlPointBounds() corefoundation.CGRect
	SetControlPointBounds(value corefoundation.CGRect)
	ElementCount() int
	SetElementCount(value int)
	Flatness() float64
	SetFlatness(value float64)
	Flattened() IBezierPath
	SetFlattened(value IBezierPath)
	IsEmpty() bool
	SetIsEmpty(value bool)
	LineCapStyle() objectivec.IObject
	SetLineCapStyle(value objectivec.IObject)
	LineJoinStyle() objectivec.IObject
	SetLineJoinStyle(value objectivec.IObject)
	LineWidth() float64
	SetLineWidth(value float64)
	MiterLimit() float64
	SetMiterLimit(value float64)
	Reversed() IBezierPath
	SetReversed(value IBezierPath)
	WindingRule() objectivec.IObject
	SetWindingRule(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (bc _BezierPathClass) Alloc() BezierPath {
	rv := objc.Send[BezierPath](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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

























// The current point (the trailing point or ending point in the most recently added segment).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/currentPoint
func (b_ BezierPath) CurrentPoint() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](b_.ID, objc.Sel("currentPoint"))
	return rv
}


// The bounding box of the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/bounds
func (b_ BezierPath) Bounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](b_.ID, objc.Sel("bounds"))
	return rv
}


// The bounding box of the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/bounds
func (b_ BezierPath) SetBounds(value corefoundation.CGRect) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBounds:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/cgpath
func (b_ BezierPath) CgPath() Path /* not a class type */ {
	rv := objc.Send[Path](b_.ID, objc.Sel("cgPath"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/cgpath
func (b_ BezierPath) SetCgPath(value Path /* not a class type */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setCgPath:"), value)
}


// The bounding box of the path, including any control points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/controlpointbounds
func (b_ BezierPath) ControlPointBounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](b_.ID, objc.Sel("controlPointBounds"))
	return rv
}


// The bounding box of the path, including any control points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/controlpointbounds
func (b_ BezierPath) SetControlPointBounds(value corefoundation.CGRect) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setControlPointBounds:"), value)
}


// The total number of path elements in the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/elementcount
func (b_ BezierPath) ElementCount() int {
	rv := objc.Send[int](b_.ID, objc.Sel("elementCount"))
	return rv
}


// The total number of path elements in the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/elementcount
func (b_ BezierPath) SetElementCount(value int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setElementCount:"), value)
}


// The accuracy with which curves are rendered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/flatness
func (b_ BezierPath) Flatness() float64 {
	rv := objc.Send[float64](b_.ID, objc.Sel("flatness"))
	return rv
}


// The accuracy with which curves are rendered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/flatness
func (b_ BezierPath) SetFlatness(value float64) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setFlatness:"), value)
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
func (b_ BezierPath) IsEmpty() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isEmpty"))
	return rv
}


// A Boolean value that indicates whether the path is empty.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/isempty
func (b_ BezierPath) SetIsEmpty(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsEmpty:"), value)
}


// The line cap style for the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/linecapstyle-swift.property
func (b_ BezierPath) LineCapStyle() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](b_.ID, objc.Sel("lineCapStyle"))
	return rv
}


// The line cap style for the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/linecapstyle-swift.property
func (b_ BezierPath) SetLineCapStyle(value objectivec.IObject) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setLineCapStyle:"), value)
}


// The line join style for the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/linejoinstyle-swift.property
func (b_ BezierPath) LineJoinStyle() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](b_.ID, objc.Sel("lineJoinStyle"))
	return rv
}


// The line join style for the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/linejoinstyle-swift.property
func (b_ BezierPath) SetLineJoinStyle(value objectivec.IObject) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setLineJoinStyle:"), value)
}


// The width of stroked path lines.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/linewidth
func (b_ BezierPath) LineWidth() float64 {
	rv := objc.Send[float64](b_.ID, objc.Sel("lineWidth"))
	return rv
}


// The width of stroked path lines.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/linewidth
func (b_ BezierPath) SetLineWidth(value float64) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setLineWidth:"), value)
}


// The limit at which miter joins are converted to bevel joins.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/miterlimit
func (b_ BezierPath) MiterLimit() float64 {
	rv := objc.Send[float64](b_.ID, objc.Sel("miterLimit"))
	return rv
}


// The limit at which miter joins are converted to bevel joins.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/miterlimit
func (b_ BezierPath) SetMiterLimit(value float64) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setMiterLimit:"), value)
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


// The winding rule used to fill the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/windingrule-swift.property
func (b_ BezierPath) WindingRule() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](b_.ID, objc.Sel("windingRule"))
	return rv
}


// The winding rule used to fill the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbezierpath/windingrule-swift.property
func (b_ BezierPath) SetWindingRule(value objectivec.IObject) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setWindingRule:"), value)
}








