// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSImageLaplacianPyramid */


/* debug [class_header]: Header for MPSImageLaplacianPyramid */
// The class instance for the [ImageLaplacianPyramid] class.
var (
	ImageLaplacianPyramidClass     _ImageLaplacianPyramidClass
	ImageLaplacianPyramidClassOnce sync.Once
)

func getImageLaplacianPyramidClass() _ImageLaplacianPyramidClass {
	ImageLaplacianPyramidClassOnce.Do(func() {
		ImageLaplacianPyramidClass = _ImageLaplacianPyramidClass{objc.GetClass("MPSImageLaplacianPyramid")}
	})
	return ImageLaplacianPyramidClass
}

type _ImageLaplacianPyramidClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageLaplacianPyramid */
// An interface definition for the [ImageLaplacianPyramid] class.
type IImageLaplacianPyramid interface {
	IImagePyramid
	
/* debug [class_interface_properties]: Properties for ImageLaplacianPyramid */
	// properties:
	LaplacianBias() objectivec.IObject
	SetLaplacianBias(value objectivec.IObject)
	LaplacianScale() objectivec.IObject
	SetLaplacianScale(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageLaplacianPyramid */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageLaplacianPyramid */
// Alloc allocates a new instance without initialization.
func (ic _ImageLaplacianPyramidClass) Alloc() ImageLaplacianPyramid {
	rv := objc.Send[ImageLaplacianPyramid](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageLaplacianPyramidClass) New() ImageLaplacianPyramid {
	rv := objc.Send[ImageLaplacianPyramid](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageLaplacianPyramid) Init() ImageLaplacianPyramid {
	rv := objc.Send[ImageLaplacianPyramid](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageLaplacianPyramid) Autorelease() ImageLaplacianPyramid {
	rv := objc.Send[ImageLaplacianPyramid](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageLaplacianPyramid creates a new ImageLaplacianPyramid instance.
func NewImageLaplacianPyramid() ImageLaplacianPyramid {
	return getImageLaplacianPyramidClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageLaplacianPyramid */
// A filter that convolves an image with a Laplacian filter.


// A filter that convolves an image with a Laplacian filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageLaplacianPyramid
type ImageLaplacianPyramid struct {
	ImagePyramid
}

// ImageLaplacianPyramidFrom constructs a [ImageLaplacianPyramid] from an unsafe.Pointer.
//
// A filter that convolves an image with a Laplacian filter.
func ImageLaplacianPyramidFrom(ptr unsafe.Pointer) ImageLaplacianPyramid {
	return ImageLaplacianPyramid{
		ImagePyramid: ImagePyramidFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageLaplacianPyramid *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageLaplacianPyramid */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageLaplacianPyramid */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageLaplacianPyramid */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageLaplacianPyramid */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagelaplacianpyramid/2966645-laplacianbias
func (i_ ImageLaplacianPyramid) LaplacianBias() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("laplacianBias"))
	return rv
}/* debug [instance_properties/getter]: laplacianBias */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagelaplacianpyramid/2966645-laplacianbias
func (i_ ImageLaplacianPyramid) SetLaplacianBias(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setLaplacianBias:"), value)
}/* debug [instance_properties/setter]: laplacianBias */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagelaplacianpyramid/2966646-laplacianscale
func (i_ ImageLaplacianPyramid) LaplacianScale() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("laplacianScale"))
	return rv
}/* debug [instance_properties/getter]: laplacianScale */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagelaplacianpyramid/2966646-laplacianscale
func (i_ ImageLaplacianPyramid) SetLaplacianScale(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setLaplacianScale:"), value)
}/* debug [instance_properties/setter]: laplacianScale */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageLaplacianPyramid */



