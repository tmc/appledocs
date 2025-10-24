// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ShapeLayer] class.
var (
	ShapeLayerClass     _ShapeLayerClass
	ShapeLayerClassOnce sync.Once
)

func getShapeLayerClass() _ShapeLayerClass {
	ShapeLayerClassOnce.Do(func() {
		ShapeLayerClass = _ShapeLayerClass{objc.GetClass("CAShapeLayer")}
	})
	return ShapeLayerClass
}

type _ShapeLayerClass struct {
	class objc.Class
}

// An interface definition for the [ShapeLayer] class.
type IShapeLayer interface {
	ILayer
	// properties:
	FillRule() ShapeLayerFillRule /* not a class type */
	SetFillRule(value ShapeLayerFillRule /* not a class type */)
	Path() PathRef /* not a class type */
	SetPath(value PathRef /* not a class type */)
	FillColor() objectivec.IObject
	SetFillColor(value objectivec.IObject)
	LineCap() ShapeLayerLineCap /* not a class type */
	SetLineCap(value ShapeLayerLineCap /* not a class type */)
	LineDashPattern() objc.IObject /* cross-framework: NSNumber */
	SetLineDashPattern(value objc.IObject /* cross-framework: NSNumber */)
	LineDashPhase() float64
	SetLineDashPhase(value float64)
	LineJoin() ShapeLayerLineJoin /* not a class type */
	SetLineJoin(value ShapeLayerLineJoin /* not a class type */)
	LineWidth() float64
	SetLineWidth(value float64)
	MiterLimit() float64
	SetMiterLimit(value float64)
	StrokeColor() objectivec.IObject
	SetStrokeColor(value objectivec.IObject)
	StrokeEnd() float64
	SetStrokeEnd(value float64)
	StrokeStart() float64
	SetStrokeStart(value float64)
	// methods:
}

// A layer that draws a cubic Bezier spline in its coordinate space.
//
// The shape is composited between the layer’s contents and its first sublayer. The shape will be drawn antialiased, and whenever possible it will be mapped into screen space before being rasterized to preserve resolution independence. However, certain kinds of image processing operations, such as CoreImage filters, applied to the layer or its ancestors may force rasterization in a local coordinate space. The following code shows how you can build complex, composite paths and display them using a shape layer. In this example, a series of progressively transformed ellipses form a simple flower shape. The shape layer that displays the path has its set to which stops the overlapping “petals” from filling with the yellow . The following figure shows the resulting shape layer.


// A layer that draws a cubic Bezier spline in its coordinate space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer
type ShapeLayer struct {
	Layer
}

// ShapeLayerFrom constructs a [ShapeLayer] from an unsafe.Pointer.
//
// A layer that draws a cubic Bezier spline in its coordinate space.
func ShapeLayerFrom(ptr unsafe.Pointer) ShapeLayer {
	return ShapeLayer{
		Layer: LayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _ShapeLayerClass) Alloc() ShapeLayer {
	rv := objc.Send[ShapeLayer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _ShapeLayerClass) New() ShapeLayer {
	rv := objc.Send[ShapeLayer](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ShapeLayer) Init() ShapeLayer {
	rv := objc.Send[ShapeLayer](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ShapeLayer) Autorelease() ShapeLayer {
	rv := objc.Send[ShapeLayer](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewShapeLayer creates a new ShapeLayer instance.
func NewShapeLayer() ShapeLayer {
	return getShapeLayerClass().New()
}



// The fill rule used when filling the shape’s path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/fillRule
func (s_ ShapeLayer) FillRule() ShapeLayerFillRule /* not a class type */ {
	rv := objc.Send[ShapeLayerFillRule](s_.ID, objc.Sel("fillRule"))
	return rv
}


// The fill rule used when filling the shape’s path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/fillRule
func (s_ ShapeLayer) SetFillRule(value ShapeLayerFillRule /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setFillRule:"), value)
}


// The path defining the shape to be rendered. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/path
func (s_ ShapeLayer) Path() PathRef /* not a class type */ {
	rv := objc.Send[PathRef](s_.ID, objc.Sel("path"))
	return rv
}


// The path defining the shape to be rendered. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/path
func (s_ ShapeLayer) SetPath(value PathRef /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPath:"), value)
}


// The color used to fill the shape’s path. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayer/fillcolor
func (s_ ShapeLayer) FillColor() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("fillColor"))
	return rv
}


// The color used to fill the shape’s path. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayer/fillcolor
func (s_ ShapeLayer) SetFillColor(value objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setFillColor:"), value)
}


// Specifies the line cap style for the shape’s path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayer/linecap
func (s_ ShapeLayer) LineCap() ShapeLayerLineCap /* not a class type */ {
	rv := objc.Send[ShapeLayerLineCap](s_.ID, objc.Sel("lineCap"))
	return rv
}


// Specifies the line cap style for the shape’s path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayer/linecap
func (s_ ShapeLayer) SetLineCap(value ShapeLayerLineCap /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLineCap:"), value)
}


// The dash pattern applied to the shape’s path when stroked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayer/linedashpattern
func (s_ ShapeLayer) LineDashPattern() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](s_.ID, objc.Sel("lineDashPattern"))
	return rv
}


// The dash pattern applied to the shape’s path when stroked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayer/linedashpattern
func (s_ ShapeLayer) SetLineDashPattern(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLineDashPattern:"), value)
}


// The dash phase applied to the shape’s path when stroked. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayer/linedashphase
func (s_ ShapeLayer) LineDashPhase() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("lineDashPhase"))
	return rv
}


// The dash phase applied to the shape’s path when stroked. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayer/linedashphase
func (s_ ShapeLayer) SetLineDashPhase(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLineDashPhase:"), value)
}


// Specifies the line join style for the shape’s path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayer/linejoin
func (s_ ShapeLayer) LineJoin() ShapeLayerLineJoin /* not a class type */ {
	rv := objc.Send[ShapeLayerLineJoin](s_.ID, objc.Sel("lineJoin"))
	return rv
}


// Specifies the line join style for the shape’s path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayer/linejoin
func (s_ ShapeLayer) SetLineJoin(value ShapeLayerLineJoin /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLineJoin:"), value)
}


// Specifies the line width of the shape’s path. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayer/linewidth
func (s_ ShapeLayer) LineWidth() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("lineWidth"))
	return rv
}


// Specifies the line width of the shape’s path. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayer/linewidth
func (s_ ShapeLayer) SetLineWidth(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLineWidth:"), value)
}


// The miter limit used when stroking the shape’s path. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayer/miterlimit
func (s_ ShapeLayer) MiterLimit() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("miterLimit"))
	return rv
}


// The miter limit used when stroking the shape’s path. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayer/miterlimit
func (s_ ShapeLayer) SetMiterLimit(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMiterLimit:"), value)
}


// The color used to stroke the shape’s path. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayer/strokecolor
func (s_ ShapeLayer) StrokeColor() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("strokeColor"))
	return rv
}


// The color used to stroke the shape’s path. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayer/strokecolor
func (s_ ShapeLayer) SetStrokeColor(value objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setStrokeColor:"), value)
}


// The relative location at which to stop stroking the path. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayer/strokeend
func (s_ ShapeLayer) StrokeEnd() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("strokeEnd"))
	return rv
}


// The relative location at which to stop stroking the path. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayer/strokeend
func (s_ ShapeLayer) SetStrokeEnd(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setStrokeEnd:"), value)
}


// The relative location at which to begin stroking the path. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayer/strokestart
func (s_ ShapeLayer) StrokeStart() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("strokeStart"))
	return rv
}


// The relative location at which to begin stroking the path. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cashapelayer/strokestart
func (s_ ShapeLayer) SetStrokeStart(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setStrokeStart:"), value)
}



