// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [GradientLayer] class.
var (
	GradientLayerClass     _GradientLayerClass
	GradientLayerClassOnce sync.Once
)

func getGradientLayerClass() _GradientLayerClass {
	GradientLayerClassOnce.Do(func() {
		GradientLayerClass = _GradientLayerClass{objc.GetClass("CAGradientLayer")}
	})
	return GradientLayerClass
}

type _GradientLayerClass struct {
	class objc.Class
}

// An interface definition for the [GradientLayer] class.
type IGradientLayer interface {
	ILayer
	// properties:
	Colors() unsafe.Pointer
	SetColors(value unsafe.Pointer)
	EndPoint() objc.IObject /* cross-framework: Point */
	SetEndPoint(value objc.IObject /* cross-framework: Point */)
	Locations() objc.IObject /* cross-framework: NSNumber */
	SetLocations(value objc.IObject /* cross-framework: NSNumber */)
	StartPoint() objc.IObject /* cross-framework: Point */
	SetStartPoint(value objc.IObject /* cross-framework: Point */)
	Type() GradientLayerType /* not a class type */
	SetType(value GradientLayerType /* not a class type */)
	// methods:
}

// A layer that draws a color gradient over its background color, filling the shape of the layer.
//
// You use a gradient layer to create a color gradient containing an arbitrary number of colors. By default, the colors are spread uniformly across the layer, but you can optionally specify locations for control over the color positions through the gradient. The following code shows how to create a gradient layer containing four colors that are evenly distributed through the gradient. Rotating the layer by 90° ( doc://com.apple.documentation/documentation/corefoundation/cgfloat/1845230-pi radians) gives a horizontal gradient. The following figure shows the appearance of the gradient layer.


// A layer that draws a color gradient over its background color, filling the shape of the layer.
//
// [Full Topic]
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



// An array of
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cagradientlayer/colors
func (g_ GradientLayer) Colors() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("colors"))
	return rv
}


// An array of
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cagradientlayer/colors
func (g_ GradientLayer) SetColors(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setColors:"), value)
}


// The end point of the gradient when drawn in the layer’s coordinate space. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cagradientlayer/endpoint
func (g_ GradientLayer) EndPoint() objc.IObject /* cross-framework: Point */ {
	rv := objc.Send[corefoundation.Point](g_.ID, objc.Sel("endPoint"))
	return rv
}


// The end point of the gradient when drawn in the layer’s coordinate space. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cagradientlayer/endpoint
func (g_ GradientLayer) SetEndPoint(value objc.IObject /* cross-framework: Point */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setEndPoint:"), value)
}


// An optional array of NSNumber objects defining the location of each gradient stop. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cagradientlayer/locations
func (g_ GradientLayer) Locations() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](g_.ID, objc.Sel("locations"))
	return rv
}


// An optional array of NSNumber objects defining the location of each gradient stop. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cagradientlayer/locations
func (g_ GradientLayer) SetLocations(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setLocations:"), value)
}


// The start point of the gradient when drawn in the layer’s coordinate space. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cagradientlayer/startpoint
func (g_ GradientLayer) StartPoint() objc.IObject /* cross-framework: Point */ {
	rv := objc.Send[corefoundation.Point](g_.ID, objc.Sel("startPoint"))
	return rv
}


// The start point of the gradient when drawn in the layer’s coordinate space. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cagradientlayer/startpoint
func (g_ GradientLayer) SetStartPoint(value objc.IObject /* cross-framework: Point */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStartPoint:"), value)
}


// Style of gradient drawn by the layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cagradientlayer/type
func (g_ GradientLayer) Type() GradientLayerType /* not a class type */ {
	rv := objc.Send[GradientLayerType](g_.ID, objc.Sel("type"))
	return rv
}


// Style of gradient drawn by the layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cagradientlayer/type
func (g_ GradientLayer) SetType(value GradientLayerType /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setType:"), value)
}



