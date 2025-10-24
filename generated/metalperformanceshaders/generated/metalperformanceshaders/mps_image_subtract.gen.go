// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSImageSubtract */


/* debug [class_header]: Header for MPSImageSubtract */
// The class instance for the [ImageSubtract] class.
var (
	ImageSubtractClass     _ImageSubtractClass
	ImageSubtractClassOnce sync.Once
)

func getImageSubtractClass() _ImageSubtractClass {
	ImageSubtractClassOnce.Do(func() {
		ImageSubtractClass = _ImageSubtractClass{objc.GetClass("MPSImageSubtract")}
	})
	return ImageSubtractClass
}

type _ImageSubtractClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageSubtract */
// An interface definition for the [ImageSubtract] class.
type IImageSubtract interface {
	IImageArithmetic
	
/* debug [class_interface_properties]: Properties for ImageSubtract */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageSubtract */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageSubtract */
// Alloc allocates a new instance without initialization.
func (ic _ImageSubtractClass) Alloc() ImageSubtract {
	rv := objc.Send[ImageSubtract](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageSubtractClass) New() ImageSubtract {
	rv := objc.Send[ImageSubtract](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageSubtract) Init() ImageSubtract {
	rv := objc.Send[ImageSubtract](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageSubtract) Autorelease() ImageSubtract {
	rv := objc.Send[ImageSubtract](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageSubtract creates a new ImageSubtract instance.
func NewImageSubtract() ImageSubtract {
	return getImageSubtractClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageSubtract */
// A filter that returns the element-wise difference of its two input images.


// A filter that returns the element-wise difference of its two input images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageSubtract
type ImageSubtract struct {
	ImageArithmetic
}

// ImageSubtractFrom constructs a [ImageSubtract] from an unsafe.Pointer.
//
// A filter that returns the element-wise difference of its two input images.
func ImageSubtractFrom(ptr unsafe.Pointer) ImageSubtract {
	return ImageSubtract{
		ImageArithmetic: ImageArithmeticFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageSubtract */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagesubtract/2866613-initwithdevice
func NewImageSubtractWithDevice(device unsafe.Pointer) ImageSubtract {
	instance := getImageSubtractClass().Alloc()
	rv := objc.Send[ImageSubtract](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageSubtractWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageSubtract */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageSubtract */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageSubtract */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageSubtract */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageSubtract */


