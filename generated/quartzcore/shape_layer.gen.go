// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ShapeLayer] class.
var (
	shapeLayerClass     _ShapeLayerClass
	shapeLayerClassOnce sync.Once
)

func getShapeLayerClass() _ShapeLayerClass {
	shapeLayerClassOnce.Do(func() {
		shapeLayerClass = _ShapeLayerClass{objc.GetClass("CAShapeLayer")}
	})
	return shapeLayerClass
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


