// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSImageReduceRowMin */


/* debug [class_header]: Header for MPSImageReduceRowMin */
// The class instance for the [ImageReduceRowMin] class.
var (
	ImageReduceRowMinClass     _ImageReduceRowMinClass
	ImageReduceRowMinClassOnce sync.Once
)

func getImageReduceRowMinClass() _ImageReduceRowMinClass {
	ImageReduceRowMinClassOnce.Do(func() {
		ImageReduceRowMinClass = _ImageReduceRowMinClass{objc.GetClass("MPSImageReduceRowMin")}
	})
	return ImageReduceRowMinClass
}

type _ImageReduceRowMinClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageReduceRowMin */
// An interface definition for the [ImageReduceRowMin] class.
type IImageReduceRowMin interface {
	IImageReduceUnary
	
/* debug [class_interface_properties]: Properties for ImageReduceRowMin */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageReduceRowMin */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageReduceRowMin */
// Alloc allocates a new instance without initialization.
func (ic _ImageReduceRowMinClass) Alloc() ImageReduceRowMin {
	rv := objc.Send[ImageReduceRowMin](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageReduceRowMinClass) New() ImageReduceRowMin {
	rv := objc.Send[ImageReduceRowMin](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageReduceRowMin) Init() ImageReduceRowMin {
	rv := objc.Send[ImageReduceRowMin](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageReduceRowMin) Autorelease() ImageReduceRowMin {
	rv := objc.Send[ImageReduceRowMin](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageReduceRowMin creates a new ImageReduceRowMin instance.
func NewImageReduceRowMin() ImageReduceRowMin {
	return getImageReduceRowMinClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageReduceRowMin */
// A filter that returns the minimum value for each row in an image.


// A filter that returns the minimum value for each row in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageReduceRowMin
type ImageReduceRowMin struct {
	ImageReduceUnary
}

// ImageReduceRowMinFrom constructs a [ImageReduceRowMin] from an unsafe.Pointer.
//
// A filter that returns the minimum value for each row in an image.
func ImageReduceRowMinFrom(ptr unsafe.Pointer) ImageReduceRowMin {
	return ImageReduceRowMin{
		ImageReduceUnary: ImageReduceUnaryFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageReduceRowMin */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagereducerowmin/2942325-initwithdevice
func NewImageReduceRowMinWithDevice(device unsafe.Pointer) ImageReduceRowMin {
	instance := getImageReduceRowMinClass().Alloc()
	rv := objc.Send[ImageReduceRowMin](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageReduceRowMinWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageReduceRowMin */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageReduceRowMin */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageReduceRowMin */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageReduceRowMin */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageReduceRowMin */


