// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSImageDivide */


/* debug [class_header]: Header for MPSImageDivide */
// The class instance for the [ImageDivide] class.
var (
	ImageDivideClass     _ImageDivideClass
	ImageDivideClassOnce sync.Once
)

func getImageDivideClass() _ImageDivideClass {
	ImageDivideClassOnce.Do(func() {
		ImageDivideClass = _ImageDivideClass{objc.GetClass("MPSImageDivide")}
	})
	return ImageDivideClass
}

type _ImageDivideClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageDivide */
// An interface definition for the [ImageDivide] class.
type IImageDivide interface {
	IImageArithmetic
	
/* debug [class_interface_properties]: Properties for ImageDivide */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageDivide */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageDivide */
// Alloc allocates a new instance without initialization.
func (ic _ImageDivideClass) Alloc() ImageDivide {
	rv := objc.Send[ImageDivide](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageDivideClass) New() ImageDivide {
	rv := objc.Send[ImageDivide](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageDivide) Init() ImageDivide {
	rv := objc.Send[ImageDivide](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageDivide) Autorelease() ImageDivide {
	rv := objc.Send[ImageDivide](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageDivide creates a new ImageDivide instance.
func NewImageDivide() ImageDivide {
	return getImageDivideClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageDivide */
// A filter that returns the element-wise quotient of its two input images.


// A filter that returns the element-wise quotient of its two input images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageDivide
type ImageDivide struct {
	ImageArithmetic
}

// ImageDivideFrom constructs a [ImageDivide] from an unsafe.Pointer.
//
// A filter that returns the element-wise quotient of its two input images.
func ImageDivideFrom(ptr unsafe.Pointer) ImageDivide {
	return ImageDivide{
		ImageArithmetic: ImageArithmeticFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageDivide */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedivide/2866606-initwithdevice
func NewImageDivideWithDevice(device unsafe.Pointer) ImageDivide {
	instance := getImageDivideClass().Alloc()
	rv := objc.Send[ImageDivide](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageDivideWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageDivide */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageDivide */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageDivide */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageDivide */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageDivide */


