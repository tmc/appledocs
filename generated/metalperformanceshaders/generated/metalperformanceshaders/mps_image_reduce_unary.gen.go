// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSImageReduceUnary */


/* debug [class_header]: Header for MPSImageReduceUnary */
// The class instance for the [ImageReduceUnary] class.
var (
	ImageReduceUnaryClass     _ImageReduceUnaryClass
	ImageReduceUnaryClassOnce sync.Once
)

func getImageReduceUnaryClass() _ImageReduceUnaryClass {
	ImageReduceUnaryClassOnce.Do(func() {
		ImageReduceUnaryClass = _ImageReduceUnaryClass{objc.GetClass("MPSImageReduceUnary")}
	})
	return ImageReduceUnaryClass
}

type _ImageReduceUnaryClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageReduceUnary */
// An interface definition for the [ImageReduceUnary] class.
type IImageReduceUnary interface {
	IUnaryImageKernel
	
/* debug [class_interface_properties]: Properties for ImageReduceUnary */
	// properties:
	ClipRectSource() Region get set /* not a class type */
	SetClipRectSource(value Region get set /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageReduceUnary */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageReduceUnary */
// Alloc allocates a new instance without initialization.
func (ic _ImageReduceUnaryClass) Alloc() ImageReduceUnary {
	rv := objc.Send[ImageReduceUnary](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageReduceUnaryClass) New() ImageReduceUnary {
	rv := objc.Send[ImageReduceUnary](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageReduceUnary) Init() ImageReduceUnary {
	rv := objc.Send[ImageReduceUnary](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageReduceUnary) Autorelease() ImageReduceUnary {
	rv := objc.Send[ImageReduceUnary](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageReduceUnary creates a new ImageReduceUnary instance.
func NewImageReduceUnary() ImageReduceUnary {
	return getImageReduceUnaryClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageReduceUnary */
// The base class for reduction filters that take a single source as input.


// The base class for reduction filters that take a single source as input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageReduceUnary
type ImageReduceUnary struct {
	UnaryImageKernel
}

// ImageReduceUnaryFrom constructs a [ImageReduceUnary] from an unsafe.Pointer.
//
// The base class for reduction filters that take a single source as input.
func ImageReduceUnaryFrom(ptr unsafe.Pointer) ImageReduceUnary {
	return ImageReduceUnary{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageReduceUnary *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageReduceUnary */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageReduceUnary */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageReduceUnary */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageReduceUnary */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagereduceunary/2942332-cliprectsource
func (i_ ImageReduceUnary) ClipRectSource() Region get set /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("clipRectSource"))
	return rv
}/* debug [instance_properties/getter]: clipRectSource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagereduceunary/2942332-cliprectsource
func (i_ ImageReduceUnary) SetClipRectSource(value Region get set /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setClipRectSource:"), value)
}/* debug [instance_properties/setter]: clipRectSource */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageReduceUnary */



