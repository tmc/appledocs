// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSImageMultiply */


/* debug [class_header]: Header for MPSImageMultiply */
// The class instance for the [ImageMultiply] class.
var (
	ImageMultiplyClass     _ImageMultiplyClass
	ImageMultiplyClassOnce sync.Once
)

func getImageMultiplyClass() _ImageMultiplyClass {
	ImageMultiplyClassOnce.Do(func() {
		ImageMultiplyClass = _ImageMultiplyClass{objc.GetClass("MPSImageMultiply")}
	})
	return ImageMultiplyClass
}

type _ImageMultiplyClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageMultiply */
// An interface definition for the [ImageMultiply] class.
type IImageMultiply interface {
	IImageArithmetic
	
/* debug [class_interface_properties]: Properties for ImageMultiply */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageMultiply */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageMultiply */
// Alloc allocates a new instance without initialization.
func (ic _ImageMultiplyClass) Alloc() ImageMultiply {
	rv := objc.Send[ImageMultiply](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageMultiplyClass) New() ImageMultiply {
	rv := objc.Send[ImageMultiply](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageMultiply) Init() ImageMultiply {
	rv := objc.Send[ImageMultiply](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageMultiply) Autorelease() ImageMultiply {
	rv := objc.Send[ImageMultiply](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageMultiply creates a new ImageMultiply instance.
func NewImageMultiply() ImageMultiply {
	return getImageMultiplyClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageMultiply */
// A filter that returns the element-wise product of its two input images.


// A filter that returns the element-wise product of its two input images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageMultiply
type ImageMultiply struct {
	ImageArithmetic
}

// ImageMultiplyFrom constructs a [ImageMultiply] from an unsafe.Pointer.
//
// A filter that returns the element-wise product of its two input images.
func ImageMultiplyFrom(ptr unsafe.Pointer) ImageMultiply {
	return ImageMultiply{
		ImageArithmetic: ImageArithmeticFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageMultiply */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagemultiply/2866600-initwithdevice
func NewImageMultiplyWithDevice(device unsafe.Pointer) ImageMultiply {
	instance := getImageMultiplyClass().Alloc()
	rv := objc.Send[ImageMultiply](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageMultiplyWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageMultiply */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageMultiply */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageMultiply */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageMultiply */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageMultiply */


