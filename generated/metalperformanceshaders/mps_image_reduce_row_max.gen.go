// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSImageReduceRowMax */


/* debug [class_header]: Header for MPSImageReduceRowMax */
// The class instance for the [ImageReduceRowMax] class.
var (
	ImageReduceRowMaxClass     _ImageReduceRowMaxClass
	ImageReduceRowMaxClassOnce sync.Once
)

func getImageReduceRowMaxClass() _ImageReduceRowMaxClass {
	ImageReduceRowMaxClassOnce.Do(func() {
		ImageReduceRowMaxClass = _ImageReduceRowMaxClass{objc.GetClass("MPSImageReduceRowMax")}
	})
	return ImageReduceRowMaxClass
}

type _ImageReduceRowMaxClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageReduceRowMax */
// An interface definition for the [ImageReduceRowMax] class.
type IImageReduceRowMax interface {
	IImageReduceUnary
	
/* debug [class_interface_properties]: Properties for ImageReduceRowMax */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageReduceRowMax */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageReduceRowMax */
// Alloc allocates a new instance without initialization.
func (ic _ImageReduceRowMaxClass) Alloc() ImageReduceRowMax {
	rv := objc.Send[ImageReduceRowMax](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageReduceRowMaxClass) New() ImageReduceRowMax {
	rv := objc.Send[ImageReduceRowMax](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageReduceRowMax) Init() ImageReduceRowMax {
	rv := objc.Send[ImageReduceRowMax](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageReduceRowMax) Autorelease() ImageReduceRowMax {
	rv := objc.Send[ImageReduceRowMax](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageReduceRowMax creates a new ImageReduceRowMax instance.
func NewImageReduceRowMax() ImageReduceRowMax {
	return getImageReduceRowMaxClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageReduceRowMax */
// A filter that returns the maximum value for each row in an image.


// A filter that returns the maximum value for each row in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageReduceRowMax
type ImageReduceRowMax struct {
	ImageReduceUnary
}

// ImageReduceRowMaxFrom constructs a [ImageReduceRowMax] from an unsafe.Pointer.
//
// A filter that returns the maximum value for each row in an image.
func ImageReduceRowMaxFrom(ptr unsafe.Pointer) ImageReduceRowMax {
	return ImageReduceRowMax{
		ImageReduceUnary: ImageReduceUnaryFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageReduceRowMax */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagereducerowmax/2942328-initwithdevice
func NewImageReduceRowMaxWithDevice(device unsafe.Pointer) ImageReduceRowMax {
	instance := getImageReduceRowMaxClass().Alloc()
	rv := objc.Send[ImageReduceRowMax](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageReduceRowMaxWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageReduceRowMax */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageReduceRowMax */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageReduceRowMax */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageReduceRowMax */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageReduceRowMax */


