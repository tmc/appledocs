// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSImageReduceColumnMin */


/* debug [class_header]: Header for MPSImageReduceColumnMin */
// The class instance for the [ImageReduceColumnMin] class.
var (
	ImageReduceColumnMinClass     _ImageReduceColumnMinClass
	ImageReduceColumnMinClassOnce sync.Once
)

func getImageReduceColumnMinClass() _ImageReduceColumnMinClass {
	ImageReduceColumnMinClassOnce.Do(func() {
		ImageReduceColumnMinClass = _ImageReduceColumnMinClass{objc.GetClass("MPSImageReduceColumnMin")}
	})
	return ImageReduceColumnMinClass
}

type _ImageReduceColumnMinClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageReduceColumnMin */
// An interface definition for the [ImageReduceColumnMin] class.
type IImageReduceColumnMin interface {
	IImageReduceUnary
	
/* debug [class_interface_properties]: Properties for ImageReduceColumnMin */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageReduceColumnMin */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageReduceColumnMin */
// Alloc allocates a new instance without initialization.
func (ic _ImageReduceColumnMinClass) Alloc() ImageReduceColumnMin {
	rv := objc.Send[ImageReduceColumnMin](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageReduceColumnMinClass) New() ImageReduceColumnMin {
	rv := objc.Send[ImageReduceColumnMin](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageReduceColumnMin) Init() ImageReduceColumnMin {
	rv := objc.Send[ImageReduceColumnMin](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageReduceColumnMin) Autorelease() ImageReduceColumnMin {
	rv := objc.Send[ImageReduceColumnMin](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageReduceColumnMin creates a new ImageReduceColumnMin instance.
func NewImageReduceColumnMin() ImageReduceColumnMin {
	return getImageReduceColumnMinClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageReduceColumnMin */
// A filter that returns the minimum value for each column in an image.


// A filter that returns the minimum value for each column in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageReduceColumnMin
type ImageReduceColumnMin struct {
	ImageReduceUnary
}

// ImageReduceColumnMinFrom constructs a [ImageReduceColumnMin] from an unsafe.Pointer.
//
// A filter that returns the minimum value for each column in an image.
func ImageReduceColumnMinFrom(ptr unsafe.Pointer) ImageReduceColumnMin {
	return ImageReduceColumnMin{
		ImageReduceUnary: ImageReduceUnaryFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageReduceColumnMin */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagereducecolumnmin/2942333-initwithdevice
func NewImageReduceColumnMinWithDevice(device unsafe.Pointer) ImageReduceColumnMin {
	instance := getImageReduceColumnMinClass().Alloc()
	rv := objc.Send[ImageReduceColumnMin](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageReduceColumnMinWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageReduceColumnMin */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageReduceColumnMin */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageReduceColumnMin */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageReduceColumnMin */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageReduceColumnMin */


