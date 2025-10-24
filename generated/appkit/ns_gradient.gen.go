// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSGradient */


/* debug [class_header]: Header for NSGradient */
// The class instance for the [Gradient] class.
var (
	GradientClass     _GradientClass
	GradientClassOnce sync.Once
)

func getGradientClass() _GradientClass {
	GradientClassOnce.Do(func() {
		GradientClass = _GradientClass{objc.GetClass("NSGradient")}
	})
	return GradientClass
}

type _GradientClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Gradient */
// An interface definition for the [Gradient] class.
type IGradient interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Gradient */
	// properties:
	ColorSpace() IColorSpace
	SetColorSpace(value IColorSpace)
	NumberOfColorStops() int
	SetNumberOfColorStops(value int)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Gradient */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Gradient */
// Alloc allocates a new instance without initialization.
func (gc _GradientClass) Alloc() Gradient {
	rv := objc.Send[Gradient](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GradientClass) New() Gradient {
	rv := objc.Send[Gradient](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ Gradient) Init() Gradient {
	rv := objc.Send[Gradient](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ Gradient) Autorelease() Gradient {
	rv := objc.Send[Gradient](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGradient creates a new Gradient instance.
func NewGradient() Gradient {
	return getGradientClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Gradient */
// An object that can draw gradient fill colors
//
// This class provides convenience methods for drawing radial or linear (axial) gradients for rectangles and objects. It also supports primitive methods that let you customize the shape of the gradient fill. A gradient consists of two or more color changes over the range of the gradient shape. When creating a gradient object, you specify the colors and their locations relative to the start and end of the gradient. This combination of color and location is known as a . During drawing, the object uses the color stop information to compute color changes for you and passes that information to the Quartz shading functions. Because the class uses Quartz shadings, drawing is handled by computing the colors at a given point mathematically. This technique results in smooth gradients regardless of the resolution of the target device. For more information about gradients and their appearance, see in .


// An object that can draw gradient fill colors
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGradient
type Gradient struct {
	objectivec.Object
}

// GradientFrom constructs a [Gradient] from an unsafe.Pointer.
//
// An object that can draw gradient fill colors
func GradientFrom(ptr unsafe.Pointer) Gradient {
	return Gradient{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Gradient */

// Initializes a newly allocated gradient object with two colors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGradient/init(starting:ending:)
func NewGradientWithStartingColorEndingColor(startingColor IColor, endingColor IColor) Gradient {
	instance := getGradientClass().Alloc()
	rv := objc.Send[Gradient](instance.ID, objc.Sel("initWithStartingColor:endingColor:"), startingColor, endingColor)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewGradientWithStartingColorEndingColor */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Gradient */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Gradient */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Gradient */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Gradient */

// The color space of the colors associated with the gradient.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgradient/colorspace
func (g_ Gradient) ColorSpace() IColorSpace {
	rv := objc.Send[ColorSpace](g_.ID, objc.Sel("colorSpace"))
	return rv
}/* debug [instance_properties/getter]: colorSpace */


// The color space of the colors associated with the gradient.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgradient/colorspace
func (g_ Gradient) SetColorSpace(value IColorSpace) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setColorSpace:"), value)
}/* debug [instance_properties/setter]: colorSpace */


// The number of color stops associated with the gradient.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgradient/numberofcolorstops
func (g_ Gradient) NumberOfColorStops() int {
	rv := objc.Send[int](g_.ID, objc.Sel("numberOfColorStops"))
	return rv
}/* debug [instance_properties/getter]: numberOfColorStops */


// The number of color stops associated with the gradient.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgradient/numberofcolorstops
func (g_ Gradient) SetNumberOfColorStops(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setNumberOfColorStops:"), value)
}/* debug [instance_properties/setter]: numberOfColorStops */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSGradient */


