// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class CAGradientLayer */


/* debug [class_header]: Header for CAGradientLayer */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GradientLayer */
// An interface definition for the [GradientLayer] class.
type IGradientLayer interface {
	ILayer
	
/* debug [class_interface_properties]: Properties for GradientLayer */
	// properties:
	Colors() objc.IObject /* cross-framework: NSArray */
	SetColors(value objc.IObject /* cross-framework: NSArray */)
	EndPoint() corefoundation.CGPoint
	SetEndPoint(value corefoundation.CGPoint)
	Locations() []foundation.Number
	SetLocations(value []foundation.Number)
	StartPoint() corefoundation.CGPoint
	SetStartPoint(value corefoundation.CGPoint)
	Type() GradientLayerType /* typedef */
	SetType(value GradientLayerType /* typedef */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GradientLayer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GradientLayer */
// Alloc allocates a new instance without initialization.
func (gc _GradientLayerClass) Alloc() GradientLayer {
	rv := objc.Send[GradientLayer](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GradientLayer */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GradientLayer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GradientLayer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GradientLayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GradientLayer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GradientLayer */

// An array of objects defining the color of each gradient stop. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAGradientLayer/colors
func (g_ GradientLayer) Colors() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](g_.ID, objc.Sel("colors"))
	return rv
}/* debug [instance_properties/getter]: colors */


// An array of objects defining the color of each gradient stop. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAGradientLayer/colors
func (g_ GradientLayer) SetColors(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setColors:"), value)
}/* debug [instance_properties/setter]: colors */


// The end point of the gradient when drawn in the layer’s coordinate space. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAGradientLayer/endPoint
func (g_ GradientLayer) EndPoint() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](g_.ID, objc.Sel("endPoint"))
	return rv
}/* debug [instance_properties/getter]: endPoint */


// The end point of the gradient when drawn in the layer’s coordinate space. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAGradientLayer/endPoint
func (g_ GradientLayer) SetEndPoint(value corefoundation.CGPoint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setEndPoint:"), value)
}/* debug [instance_properties/setter]: endPoint */


// An optional array of NSNumber objects defining the location of each gradient stop. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAGradientLayer/locations
func (g_ GradientLayer) Locations() []foundation.Number {
	rv := objc.Send[[]foundation.Number](g_.ID, objc.Sel("locations"))
	return rv
}/* debug [instance_properties/getter]: locations */


// An optional array of NSNumber objects defining the location of each gradient stop. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAGradientLayer/locations
func (g_ GradientLayer) SetLocations(value []foundation.Number) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](g_.ID, objc.Sel("setLocations:"), nsArray)
}/* debug [instance_properties/setter]: locations */


// The start point of the gradient when drawn in the layer’s coordinate space. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAGradientLayer/startPoint
func (g_ GradientLayer) StartPoint() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](g_.ID, objc.Sel("startPoint"))
	return rv
}/* debug [instance_properties/getter]: startPoint */


// The start point of the gradient when drawn in the layer’s coordinate space. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAGradientLayer/startPoint
func (g_ GradientLayer) SetStartPoint(value corefoundation.CGPoint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStartPoint:"), value)
}/* debug [instance_properties/setter]: startPoint */


// Style of gradient drawn by the layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAGradientLayer/type
func (g_ GradientLayer) Type() GradientLayerType /* typedef */ {
	rv := objc.Send[foundation.NSString](g_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// Style of gradient drawn by the layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAGradientLayer/type
func (g_ GradientLayer) SetType(value GradientLayerType /* typedef */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setType:"), value)
}/* debug [instance_properties/setter]: type */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CAGradientLayer */



