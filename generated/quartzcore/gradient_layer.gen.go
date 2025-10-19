// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [GradientLayer] class.
var (
	gradientLayerClass     _GradientLayerClass
	gradientLayerClassOnce sync.Once
)

func getGradientLayerClass() _GradientLayerClass {
	gradientLayerClassOnce.Do(func() {
		gradientLayerClass = _GradientLayerClass{objc.GetClass("CAGradientLayer")}
	})
	return gradientLayerClass
}

type _GradientLayerClass struct {
	class objc.Class
}

// An interface definition for the [GradientLayer] class.
type IGradientLayer interface {
	ILayer
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

// Alloc allocates a new instance without initialization.
func (gc _GradientLayerClass) Alloc() GradientLayer {
	rv := objc.Send[GradientLayer](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GradientLayerClass) New() GradientLayer {
	rv := objc.Send[GradientLayer](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GradientLayer) Init() GradientLayer {
	rv := objc.Send[GradientLayer](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GradientLayer) Autorelease() GradientLayer {
	rv := objc.Send[GradientLayer](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGradientLayer creates a new GradientLayer instance.
func NewGradientLayer() GradientLayer {
	return getGradientLayerClass().New()
}




