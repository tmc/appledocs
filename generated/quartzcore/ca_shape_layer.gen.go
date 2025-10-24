// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class CAShapeLayer */


/* debug [class_header]: Header for CAShapeLayer */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ShapeLayer */
// An interface definition for the [ShapeLayer] class.
type IShapeLayer interface {
	ILayer
	
/* debug [class_interface_properties]: Properties for ShapeLayer */
	// properties:
	FillColor() ColorRef /* not a class type */
	SetFillColor(value ColorRef /* not a class type */)
	FillRule() ShapeLayerFillRule /* typedef */
	SetFillRule(value ShapeLayerFillRule /* typedef */)
	LineCap() ShapeLayerLineCap /* typedef */
	SetLineCap(value ShapeLayerLineCap /* typedef */)
	LineDashPattern() []foundation.Number
	SetLineDashPattern(value []foundation.Number)
	LineDashPhase() float64
	SetLineDashPhase(value float64)
	LineJoin() ShapeLayerLineJoin /* typedef */
	SetLineJoin(value ShapeLayerLineJoin /* typedef */)
	LineWidth() float64
	SetLineWidth(value float64)
	MiterLimit() float64
	SetMiterLimit(value float64)
	Path() PathRef /* not a class type */
	SetPath(value PathRef /* not a class type */)
	StrokeColor() ColorRef /* not a class type */
	SetStrokeColor(value ColorRef /* not a class type */)
	StrokeEnd() float64
	SetStrokeEnd(value float64)
	StrokeStart() float64
	SetStrokeStart(value float64)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ShapeLayer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ShapeLayer */
// Alloc allocates a new instance without initialization.
func (sc _ShapeLayerClass) Alloc() ShapeLayer {
	rv := objc.Send[ShapeLayer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ShapeLayer */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ShapeLayer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ShapeLayer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ShapeLayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ShapeLayer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ShapeLayer */

// The color used to fill the shape’s path. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/fillColor
func (s_ ShapeLayer) FillColor() ColorRef /* not a class type */ {
	rv := objc.Send[ColorRef](s_.ID, objc.Sel("fillColor"))
	return rv
}/* debug [instance_properties/getter]: fillColor */


// The color used to fill the shape’s path. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/fillColor
func (s_ ShapeLayer) SetFillColor(value ColorRef /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setFillColor:"), value)
}/* debug [instance_properties/setter]: fillColor */


// The fill rule used when filling the shape’s path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/fillRule
func (s_ ShapeLayer) FillRule() ShapeLayerFillRule /* typedef */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("fillRule"))
	return rv
}/* debug [instance_properties/getter]: fillRule */


// The fill rule used when filling the shape’s path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/fillRule
func (s_ ShapeLayer) SetFillRule(value ShapeLayerFillRule /* typedef */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setFillRule:"), value)
}/* debug [instance_properties/setter]: fillRule */


// Specifies the line cap style for the shape’s path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/lineCap
func (s_ ShapeLayer) LineCap() ShapeLayerLineCap /* typedef */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("lineCap"))
	return rv
}/* debug [instance_properties/getter]: lineCap */


// Specifies the line cap style for the shape’s path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/lineCap
func (s_ ShapeLayer) SetLineCap(value ShapeLayerLineCap /* typedef */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLineCap:"), value)
}/* debug [instance_properties/setter]: lineCap */


// The dash pattern applied to the shape’s path when stroked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/lineDashPattern
func (s_ ShapeLayer) LineDashPattern() []foundation.Number {
	rv := objc.Send[[]foundation.Number](s_.ID, objc.Sel("lineDashPattern"))
	return rv
}/* debug [instance_properties/getter]: lineDashPattern */


// The dash pattern applied to the shape’s path when stroked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/lineDashPattern
func (s_ ShapeLayer) SetLineDashPattern(value []foundation.Number) {
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
}/* debug [instance_properties/setter]: lineDashPattern */


// The dash phase applied to the shape’s path when stroked. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/lineDashPhase
func (s_ ShapeLayer) LineDashPhase() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("lineDashPhase"))
	return rv
}/* debug [instance_properties/getter]: lineDashPhase */


// The dash phase applied to the shape’s path when stroked. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/lineDashPhase
func (s_ ShapeLayer) SetLineDashPhase(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLineDashPhase:"), value)
}/* debug [instance_properties/setter]: lineDashPhase */


// Specifies the line join style for the shape’s path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/lineJoin
func (s_ ShapeLayer) LineJoin() ShapeLayerLineJoin /* typedef */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("lineJoin"))
	return rv
}/* debug [instance_properties/getter]: lineJoin */


// Specifies the line join style for the shape’s path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/lineJoin
func (s_ ShapeLayer) SetLineJoin(value ShapeLayerLineJoin /* typedef */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLineJoin:"), value)
}/* debug [instance_properties/setter]: lineJoin */


// Specifies the line width of the shape’s path. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/lineWidth
func (s_ ShapeLayer) LineWidth() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("lineWidth"))
	return rv
}/* debug [instance_properties/getter]: lineWidth */


// Specifies the line width of the shape’s path. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/lineWidth
func (s_ ShapeLayer) SetLineWidth(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLineWidth:"), value)
}/* debug [instance_properties/setter]: lineWidth */


// The miter limit used when stroking the shape’s path. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/miterLimit
func (s_ ShapeLayer) MiterLimit() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("miterLimit"))
	return rv
}/* debug [instance_properties/getter]: miterLimit */


// The miter limit used when stroking the shape’s path. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/miterLimit
func (s_ ShapeLayer) SetMiterLimit(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMiterLimit:"), value)
}/* debug [instance_properties/setter]: miterLimit */


// The path defining the shape to be rendered. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/path
func (s_ ShapeLayer) Path() PathRef /* not a class type */ {
	rv := objc.Send[PathRef](s_.ID, objc.Sel("path"))
	return rv
}/* debug [instance_properties/getter]: path */


// The path defining the shape to be rendered. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/path
func (s_ ShapeLayer) SetPath(value PathRef /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPath:"), value)
}/* debug [instance_properties/setter]: path */


// The color used to stroke the shape’s path. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/strokeColor
func (s_ ShapeLayer) StrokeColor() ColorRef /* not a class type */ {
	rv := objc.Send[ColorRef](s_.ID, objc.Sel("strokeColor"))
	return rv
}/* debug [instance_properties/getter]: strokeColor */


// The color used to stroke the shape’s path. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/strokeColor
func (s_ ShapeLayer) SetStrokeColor(value ColorRef /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setStrokeColor:"), value)
}/* debug [instance_properties/setter]: strokeColor */


// The relative location at which to stop stroking the path. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/strokeEnd
func (s_ ShapeLayer) StrokeEnd() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("strokeEnd"))
	return rv
}/* debug [instance_properties/getter]: strokeEnd */


// The relative location at which to stop stroking the path. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/strokeEnd
func (s_ ShapeLayer) SetStrokeEnd(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setStrokeEnd:"), value)
}/* debug [instance_properties/setter]: strokeEnd */


// The relative location at which to begin stroking the path. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/strokeStart
func (s_ ShapeLayer) StrokeStart() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("strokeStart"))
	return rv
}/* debug [instance_properties/getter]: strokeStart */


// The relative location at which to begin stroking the path. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/strokeStart
func (s_ ShapeLayer) SetStrokeStart(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setStrokeStart:"), value)
}/* debug [instance_properties/setter]: strokeStart */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CAShapeLayer */



