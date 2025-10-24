// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSImageReduceColumnMean */


/* debug [class_header]: Header for MPSImageReduceColumnMean */
// The class instance for the [ImageReduceColumnMean] class.
var (
	ImageReduceColumnMeanClass     _ImageReduceColumnMeanClass
	ImageReduceColumnMeanClassOnce sync.Once
)

func getImageReduceColumnMeanClass() _ImageReduceColumnMeanClass {
	ImageReduceColumnMeanClassOnce.Do(func() {
		ImageReduceColumnMeanClass = _ImageReduceColumnMeanClass{objc.GetClass("MPSImageReduceColumnMean")}
	})
	return ImageReduceColumnMeanClass
}

type _ImageReduceColumnMeanClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageReduceColumnMean */
// An interface definition for the [ImageReduceColumnMean] class.
type IImageReduceColumnMean interface {
	IImageReduceUnary
	
/* debug [class_interface_properties]: Properties for ImageReduceColumnMean */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageReduceColumnMean */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageReduceColumnMean */
// Alloc allocates a new instance without initialization.
func (ic _ImageReduceColumnMeanClass) Alloc() ImageReduceColumnMean {
	rv := objc.Send[ImageReduceColumnMean](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageReduceColumnMeanClass) New() ImageReduceColumnMean {
	rv := objc.Send[ImageReduceColumnMean](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageReduceColumnMean) Init() ImageReduceColumnMean {
	rv := objc.Send[ImageReduceColumnMean](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageReduceColumnMean) Autorelease() ImageReduceColumnMean {
	rv := objc.Send[ImageReduceColumnMean](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageReduceColumnMean creates a new ImageReduceColumnMean instance.
func NewImageReduceColumnMean() ImageReduceColumnMean {
	return getImageReduceColumnMeanClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageReduceColumnMean */
// A filter that returns the mean value for each column in an image.


// A filter that returns the mean value for each column in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageReduceColumnMean
type ImageReduceColumnMean struct {
	ImageReduceUnary
}

// ImageReduceColumnMeanFrom constructs a [ImageReduceColumnMean] from an unsafe.Pointer.
//
// A filter that returns the mean value for each column in an image.
func ImageReduceColumnMeanFrom(ptr unsafe.Pointer) ImageReduceColumnMean {
	return ImageReduceColumnMean{
		ImageReduceUnary: ImageReduceUnaryFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageReduceColumnMean */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagereducecolumnmean/2942331-initwithdevice
func NewImageReduceColumnMeanWithDevice(device unsafe.Pointer) ImageReduceColumnMean {
	instance := getImageReduceColumnMeanClass().Alloc()
	rv := objc.Send[ImageReduceColumnMean](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageReduceColumnMeanWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageReduceColumnMean */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageReduceColumnMean */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageReduceColumnMean */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageReduceColumnMean */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageReduceColumnMean */


