// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSImageAdd */


/* debug [class_header]: Header for MPSImageAdd */
// The class instance for the [ImageAdd] class.
var (
	ImageAddClass     _ImageAddClass
	ImageAddClassOnce sync.Once
)

func getImageAddClass() _ImageAddClass {
	ImageAddClassOnce.Do(func() {
		ImageAddClass = _ImageAddClass{objc.GetClass("MPSImageAdd")}
	})
	return ImageAddClass
}

type _ImageAddClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageAdd */
// An interface definition for the [ImageAdd] class.
type IImageAdd interface {
	IImageArithmetic
	
/* debug [class_interface_properties]: Properties for ImageAdd */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageAdd */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageAdd */
// Alloc allocates a new instance without initialization.
func (ic _ImageAddClass) Alloc() ImageAdd {
	rv := objc.Send[ImageAdd](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageAddClass) New() ImageAdd {
	rv := objc.Send[ImageAdd](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageAdd) Init() ImageAdd {
	rv := objc.Send[ImageAdd](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageAdd) Autorelease() ImageAdd {
	rv := objc.Send[ImageAdd](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageAdd creates a new ImageAdd instance.
func NewImageAdd() ImageAdd {
	return getImageAddClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageAdd */
// A filter that returns the element-wise sum of its two input images.


// A filter that returns the element-wise sum of its two input images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageAdd
type ImageAdd struct {
	ImageArithmetic
}

// ImageAddFrom constructs a [ImageAdd] from an unsafe.Pointer.
//
// A filter that returns the element-wise sum of its two input images.
func ImageAddFrom(ptr unsafe.Pointer) ImageAdd {
	return ImageAdd{
		ImageArithmetic: ImageArithmeticFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageAdd */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageadd/2866610-initwithdevice
func NewImageAddWithDevice(device unsafe.Pointer) ImageAdd {
	instance := getImageAddClass().Alloc()
	rv := objc.Send[ImageAdd](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageAddWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageAdd */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageAdd */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageAdd */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageAdd */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageAdd */


