// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
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
}

// A layer that draws a cubic Bezier spline in its coordinate space.
//
// The shape is composited between the layer’s contents and its first sublayer. The shape will be drawn antialiased, and whenever possible it will be mapped into screen space before being rasterized to preserve resolution independence. However, certain kinds of image processing operations, such as CoreImage filters, applied to the layer or its ancestors may force rasterization in a local coordinate space. The following code shows how you can build complex, composite paths and display them using a shape layer. In this example, a series of progressively transformed ellipses form a simple flower shape. The shape layer that displays the path has its set to which stops the overlapping “petals” from filling with the yellow . The following figure shows the resulting shape layer.
//
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


// The color used to fill the shape’s path. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/fillColor
func (s_ ShapeLayer) FillColor() coregraphics.CGColorRef {
	rv := objc.Send[coregraphics.CGColorRef](s_.ID, objc.Sel("fillColor"))
	return rv
}


// SetFillColor sets the value of the fillColor property.
// The color used to fill the shape’s path. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/fillColor
func (s_ ShapeLayer) SetFillColor(value coregraphics.CGColorRef) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setFillColor:"), value)
}

// The fill rule used when filling the shape’s path.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/fillRule
func (s_ ShapeLayer) FillRule() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("fillRule"))
	return rv
}


// SetFillRule sets the value of the fillRule property.
// The fill rule used when filling the shape’s path.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/fillRule
func (s_ ShapeLayer) SetFillRule(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setFillRule:"), value)
}

// Specifies the line cap style for the shape’s path.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/lineCap
func (s_ ShapeLayer) LineCap() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("lineCap"))
	return rv
}


// SetLineCap sets the value of the lineCap property.
// Specifies the line cap style for the shape’s path.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/lineCap
func (s_ ShapeLayer) SetLineCap(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLineCap:"), value)
}

// The dash pattern applied to the shape’s path when stroked.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/lineDashPattern
func (s_ ShapeLayer) LineDashPattern() []NSNumber {
	rv := objc.Send[[]NSNumber](s_.ID, objc.Sel("lineDashPattern"))
	return rv
}


// SetLineDashPattern sets the value of the lineDashPattern property.
// The dash pattern applied to the shape’s path when stroked.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/lineDashPattern
func (s_ ShapeLayer) SetLineDashPattern(value []NSNumber) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](s_.ID, objc.Sel("setLineDashPattern:"), nsArray)
}

// The dash phase applied to the shape’s path when stroked. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/lineDashPhase
func (s_ ShapeLayer) LineDashPhase() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("lineDashPhase"))
	return rv
}


// SetLineDashPhase sets the value of the lineDashPhase property.
// The dash phase applied to the shape’s path when stroked. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/lineDashPhase
func (s_ ShapeLayer) SetLineDashPhase(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLineDashPhase:"), value)
}

// Specifies the line join style for the shape’s path.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/lineJoin
func (s_ ShapeLayer) LineJoin() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("lineJoin"))
	return rv
}


// SetLineJoin sets the value of the lineJoin property.
// Specifies the line join style for the shape’s path.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/lineJoin
func (s_ ShapeLayer) SetLineJoin(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLineJoin:"), value)
}

// Specifies the line width of the shape’s path. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/lineWidth
func (s_ ShapeLayer) LineWidth() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("lineWidth"))
	return rv
}


// SetLineWidth sets the value of the lineWidth property.
// Specifies the line width of the shape’s path. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/lineWidth
func (s_ ShapeLayer) SetLineWidth(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLineWidth:"), value)
}

// The miter limit used when stroking the shape’s path. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/miterLimit
func (s_ ShapeLayer) MiterLimit() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("miterLimit"))
	return rv
}


// SetMiterLimit sets the value of the miterLimit property.
// The miter limit used when stroking the shape’s path. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/miterLimit
func (s_ ShapeLayer) SetMiterLimit(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMiterLimit:"), value)
}

// The path defining the shape to be rendered. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/path
func (s_ ShapeLayer) Path() coregraphics.CGPathRef {
	rv := objc.Send[coregraphics.CGPathRef](s_.ID, objc.Sel("path"))
	return rv
}


// SetPath sets the value of the path property.
// The path defining the shape to be rendered. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/path
func (s_ ShapeLayer) SetPath(value coregraphics.CGPathRef) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPath:"), value)
}

// The color used to stroke the shape’s path. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/strokeColor
func (s_ ShapeLayer) StrokeColor() coregraphics.CGColorRef {
	rv := objc.Send[coregraphics.CGColorRef](s_.ID, objc.Sel("strokeColor"))
	return rv
}


// SetStrokeColor sets the value of the strokeColor property.
// The color used to stroke the shape’s path. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/strokeColor
func (s_ ShapeLayer) SetStrokeColor(value coregraphics.CGColorRef) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setStrokeColor:"), value)
}

// The relative location at which to stop stroking the path. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/strokeEnd
func (s_ ShapeLayer) StrokeEnd() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("strokeEnd"))
	return rv
}


// SetStrokeEnd sets the value of the strokeEnd property.
// The relative location at which to stop stroking the path. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/strokeEnd
func (s_ ShapeLayer) SetStrokeEnd(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setStrokeEnd:"), value)
}

// The relative location at which to begin stroking the path. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/strokeStart
func (s_ ShapeLayer) StrokeStart() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("strokeStart"))
	return rv
}


// SetStrokeStart sets the value of the strokeStart property.
// The relative location at which to begin stroking the path. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/strokeStart
func (s_ ShapeLayer) SetStrokeStart(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setStrokeStart:"), value)
}



