// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSImageReduceRowMean */


/* debug [class_header]: Header for MPSImageReduceRowMean */
// The class instance for the [ImageReduceRowMean] class.
var (
	ImageReduceRowMeanClass     _ImageReduceRowMeanClass
	ImageReduceRowMeanClassOnce sync.Once
)

func getImageReduceRowMeanClass() _ImageReduceRowMeanClass {
	ImageReduceRowMeanClassOnce.Do(func() {
		ImageReduceRowMeanClass = _ImageReduceRowMeanClass{objc.GetClass("MPSImageReduceRowMean")}
	})
	return ImageReduceRowMeanClass
}

type _ImageReduceRowMeanClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageReduceRowMean */
// An interface definition for the [ImageReduceRowMean] class.
type IImageReduceRowMean interface {
	IImageReduceUnary
	
/* debug [class_interface_properties]: Properties for ImageReduceRowMean */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageReduceRowMean */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageReduceRowMean */
// Alloc allocates a new instance without initialization.
func (ic _ImageReduceRowMeanClass) Alloc() ImageReduceRowMean {
	rv := objc.Send[ImageReduceRowMean](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageReduceRowMeanClass) New() ImageReduceRowMean {
	rv := objc.Send[ImageReduceRowMean](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageReduceRowMean) Init() ImageReduceRowMean {
	rv := objc.Send[ImageReduceRowMean](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageReduceRowMean) Autorelease() ImageReduceRowMean {
	rv := objc.Send[ImageReduceRowMean](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageReduceRowMean creates a new ImageReduceRowMean instance.
func NewImageReduceRowMean() ImageReduceRowMean {
	return getImageReduceRowMeanClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageReduceRowMean */
// A filter that returns the mean value for each row in an image.


// A filter that returns the mean value for each row in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageReduceRowMean
type ImageReduceRowMean struct {
	ImageReduceUnary
}

// ImageReduceRowMeanFrom constructs a [ImageReduceRowMean] from an unsafe.Pointer.
//
// A filter that returns the mean value for each row in an image.
func ImageReduceRowMeanFrom(ptr unsafe.Pointer) ImageReduceRowMean {
	return ImageReduceRowMean{
		ImageReduceUnary: ImageReduceUnaryFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageReduceRowMean */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagereducerowmean/2942322-initwithdevice
func NewImageReduceRowMeanWithDevice(device unsafe.Pointer) ImageReduceRowMean {
	instance := getImageReduceRowMeanClass().Alloc()
	rv := objc.Send[ImageReduceRowMean](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageReduceRowMeanWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageReduceRowMean */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageReduceRowMean */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageReduceRowMean */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageReduceRowMean */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageReduceRowMean */


