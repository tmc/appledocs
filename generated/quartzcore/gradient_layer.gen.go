// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [GradientLayer] class.
var gradientLayerClass = _GradientLayerClass{objc.GetClass("CAGradientLayer")}

type _GradientLayerClass struct {
	class objc.Class
}

// A layer that draws a color gradient over its background color, filling the shape of the layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAGradientLayer

type GradientLayer struct {
	Layer
}

// GradientLayerFrom constructs a [GradientLayer] from an unsafe.Pointer.
//
// A layer that draws a color gradient over its background color, filling the shape of the layer.
func GradientLayerFrom(ptr unsafe.Pointer) GradientLayer {
	return GradientLayer{
		Layer: LayerFrom(ptr),
	}
}



