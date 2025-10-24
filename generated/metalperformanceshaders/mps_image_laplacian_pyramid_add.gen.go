// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSImageLaplacianPyramidAdd */


/* debug [class_header]: Header for MPSImageLaplacianPyramidAdd */
// The class instance for the [ImageLaplacianPyramidAdd] class.
var (
	ImageLaplacianPyramidAddClass     _ImageLaplacianPyramidAddClass
	ImageLaplacianPyramidAddClassOnce sync.Once
)

func getImageLaplacianPyramidAddClass() _ImageLaplacianPyramidAddClass {
	ImageLaplacianPyramidAddClassOnce.Do(func() {
		ImageLaplacianPyramidAddClass = _ImageLaplacianPyramidAddClass{objc.GetClass("MPSImageLaplacianPyramidAdd")}
	})
	return ImageLaplacianPyramidAddClass
}

type _ImageLaplacianPyramidAddClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageLaplacianPyramidAdd */
// An interface definition for the [ImageLaplacianPyramidAdd] class.
type IImageLaplacianPyramidAdd interface {
	IImageLaplacianPyramid
	
/* debug [class_interface_properties]: Properties for ImageLaplacianPyramidAdd */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageLaplacianPyramidAdd */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageLaplacianPyramidAdd */
// Alloc allocates a new instance without initialization.
func (ic _ImageLaplacianPyramidAddClass) Alloc() ImageLaplacianPyramidAdd {
	rv := objc.Send[ImageLaplacianPyramidAdd](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageLaplacianPyramidAddClass) New() ImageLaplacianPyramidAdd {
	rv := objc.Send[ImageLaplacianPyramidAdd](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageLaplacianPyramidAdd) Init() ImageLaplacianPyramidAdd {
	rv := objc.Send[ImageLaplacianPyramidAdd](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageLaplacianPyramidAdd) Autorelease() ImageLaplacianPyramidAdd {
	rv := objc.Send[ImageLaplacianPyramidAdd](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageLaplacianPyramidAdd creates a new ImageLaplacianPyramidAdd instance.
func NewImageLaplacianPyramidAdd() ImageLaplacianPyramidAdd {
	return getImageLaplacianPyramidAddClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageLaplacianPyramidAdd */
// A filter that convolves an image with an additive Laplacian pyramid.


// A filter that convolves an image with an additive Laplacian pyramid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageLaplacianPyramidAdd
type ImageLaplacianPyramidAdd struct {
	ImageLaplacianPyramid
}

// ImageLaplacianPyramidAddFrom constructs a [ImageLaplacianPyramidAdd] from an unsafe.Pointer.
//
// A filter that convolves an image with an additive Laplacian pyramid.
func ImageLaplacianPyramidAddFrom(ptr unsafe.Pointer) ImageLaplacianPyramidAdd {
	return ImageLaplacianPyramidAdd{
		ImageLaplacianPyramid: ImageLaplacianPyramidFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageLaplacianPyramidAdd *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageLaplacianPyramidAdd */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageLaplacianPyramidAdd */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageLaplacianPyramidAdd */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageLaplacianPyramidAdd */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageLaplacianPyramidAdd */



