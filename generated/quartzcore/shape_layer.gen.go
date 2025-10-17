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



