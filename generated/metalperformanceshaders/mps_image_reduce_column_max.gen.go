// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSImageReduceColumnMax */


/* debug [class_header]: Header for MPSImageReduceColumnMax */
// The class instance for the [ImageReduceColumnMax] class.
var (
	ImageReduceColumnMaxClass     _ImageReduceColumnMaxClass
	ImageReduceColumnMaxClassOnce sync.Once
)

func getImageReduceColumnMaxClass() _ImageReduceColumnMaxClass {
	ImageReduceColumnMaxClassOnce.Do(func() {
		ImageReduceColumnMaxClass = _ImageReduceColumnMaxClass{objc.GetClass("MPSImageReduceColumnMax")}
	})
	return ImageReduceColumnMaxClass
}

type _ImageReduceColumnMaxClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageReduceColumnMax */
// An interface definition for the [ImageReduceColumnMax] class.
type IImageReduceColumnMax interface {
	IImageReduceUnary
	
/* debug [class_interface_properties]: Properties for ImageReduceColumnMax */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageReduceColumnMax */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageReduceColumnMax */
// Alloc allocates a new instance without initialization.
func (ic _ImageReduceColumnMaxClass) Alloc() ImageReduceColumnMax {
	rv := objc.Send[ImageReduceColumnMax](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageReduceColumnMaxClass) New() ImageReduceColumnMax {
	rv := objc.Send[ImageReduceColumnMax](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageReduceColumnMax) Init() ImageReduceColumnMax {
	rv := objc.Send[ImageReduceColumnMax](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageReduceColumnMax) Autorelease() ImageReduceColumnMax {
	rv := objc.Send[ImageReduceColumnMax](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageReduceColumnMax creates a new ImageReduceColumnMax instance.
func NewImageReduceColumnMax() ImageReduceColumnMax {
	return getImageReduceColumnMaxClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageReduceColumnMax */
// A filter that returns the maximum value for each column in an image.


// A filter that returns the maximum value for each column in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageReduceColumnMax
type ImageReduceColumnMax struct {
	ImageReduceUnary
}

// ImageReduceColumnMaxFrom constructs a [ImageReduceColumnMax] from an unsafe.Pointer.
//
// A filter that returns the maximum value for each column in an image.
func ImageReduceColumnMaxFrom(ptr unsafe.Pointer) ImageReduceColumnMax {
	return ImageReduceColumnMax{
		ImageReduceUnary: ImageReduceUnaryFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageReduceColumnMax */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagereducecolumnmax/2942318-initwithdevice
func NewImageReduceColumnMaxWithDevice(device unsafe.Pointer) ImageReduceColumnMax {
	instance := getImageReduceColumnMaxClass().Alloc()
	rv := objc.Send[ImageReduceColumnMax](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageReduceColumnMaxWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageReduceColumnMax */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageReduceColumnMax */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageReduceColumnMax */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageReduceColumnMax */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageReduceColumnMax */


