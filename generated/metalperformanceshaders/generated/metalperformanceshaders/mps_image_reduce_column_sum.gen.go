// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSImageReduceColumnSum */


/* debug [class_header]: Header for MPSImageReduceColumnSum */
// The class instance for the [ImageReduceColumnSum] class.
var (
	ImageReduceColumnSumClass     _ImageReduceColumnSumClass
	ImageReduceColumnSumClassOnce sync.Once
)

func getImageReduceColumnSumClass() _ImageReduceColumnSumClass {
	ImageReduceColumnSumClassOnce.Do(func() {
		ImageReduceColumnSumClass = _ImageReduceColumnSumClass{objc.GetClass("MPSImageReduceColumnSum")}
	})
	return ImageReduceColumnSumClass
}

type _ImageReduceColumnSumClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageReduceColumnSum */
// An interface definition for the [ImageReduceColumnSum] class.
type IImageReduceColumnSum interface {
	IImageReduceUnary
	
/* debug [class_interface_properties]: Properties for ImageReduceColumnSum */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageReduceColumnSum */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageReduceColumnSum */
// Alloc allocates a new instance without initialization.
func (ic _ImageReduceColumnSumClass) Alloc() ImageReduceColumnSum {
	rv := objc.Send[ImageReduceColumnSum](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageReduceColumnSumClass) New() ImageReduceColumnSum {
	rv := objc.Send[ImageReduceColumnSum](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageReduceColumnSum) Init() ImageReduceColumnSum {
	rv := objc.Send[ImageReduceColumnSum](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageReduceColumnSum) Autorelease() ImageReduceColumnSum {
	rv := objc.Send[ImageReduceColumnSum](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageReduceColumnSum creates a new ImageReduceColumnSum instance.
func NewImageReduceColumnSum() ImageReduceColumnSum {
	return getImageReduceColumnSumClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageReduceColumnSum */
// A filter that returns the sum of all values for a column in an image.


// A filter that returns the sum of all values for a column in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageReduceColumnSum
type ImageReduceColumnSum struct {
	ImageReduceUnary
}

// ImageReduceColumnSumFrom constructs a [ImageReduceColumnSum] from an unsafe.Pointer.
//
// A filter that returns the sum of all values for a column in an image.
func ImageReduceColumnSumFrom(ptr unsafe.Pointer) ImageReduceColumnSum {
	return ImageReduceColumnSum{
		ImageReduceUnary: ImageReduceUnaryFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageReduceColumnSum */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagereducecolumnsum/2942321-initwithdevice
func NewImageReduceColumnSumWithDevice(device unsafe.Pointer) ImageReduceColumnSum {
	instance := getImageReduceColumnSumClass().Alloc()
	rv := objc.Send[ImageReduceColumnSum](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageReduceColumnSumWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageReduceColumnSum */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageReduceColumnSum */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageReduceColumnSum */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageReduceColumnSum */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageReduceColumnSum */


