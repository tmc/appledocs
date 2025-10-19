// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ShapeLayer] class.
var shapeLayerClass = _ShapeLayerClass{objc.GetClass("CAShapeLayer")}

type _ShapeLayerClass struct {
	class objc.Class
}

// An interface definition for the [ShapeLayer] class.
type IShapeLayer interface {
	ILayer
}

// A layer that draws a cubic Bezier spline in its coordinate space. [Full Topic]
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

// New creates and returns a new instance with a +1 retain count.
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
	return shapeLayerClass.New()
}




