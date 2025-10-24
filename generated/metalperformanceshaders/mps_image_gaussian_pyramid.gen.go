// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSImageGaussianPyramid */


/* debug [class_header]: Header for MPSImageGaussianPyramid */
// The class instance for the [ImageGaussianPyramid] class.
var (
	ImageGaussianPyramidClass     _ImageGaussianPyramidClass
	ImageGaussianPyramidClassOnce sync.Once
)

func getImageGaussianPyramidClass() _ImageGaussianPyramidClass {
	ImageGaussianPyramidClassOnce.Do(func() {
		ImageGaussianPyramidClass = _ImageGaussianPyramidClass{objc.GetClass("MPSImageGaussianPyramid")}
	})
	return ImageGaussianPyramidClass
}

type _ImageGaussianPyramidClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageGaussianPyramid */
// An interface definition for the [ImageGaussianPyramid] class.
type IImageGaussianPyramid interface {
	IImagePyramid
	
/* debug [class_interface_properties]: Properties for ImageGaussianPyramid */
	// properties:
	ClipRect() objc.IObject /* cross-framework: MTLRegion */
	SetClipRect(value objc.IObject /* cross-framework: MTLRegion */)
	Offset() objc.IObject /* cross-framework: MPSOffset */
	SetOffset(value objc.IObject /* cross-framework: MPSOffset */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageGaussianPyramid */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageGaussianPyramid */
// Alloc allocates a new instance without initialization.
func (ic _ImageGaussianPyramidClass) Alloc() ImageGaussianPyramid {
	rv := objc.Send[ImageGaussianPyramid](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageGaussianPyramidClass) New() ImageGaussianPyramid {
	rv := objc.Send[ImageGaussianPyramid](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageGaussianPyramid) Init() ImageGaussianPyramid {
	rv := objc.Send[ImageGaussianPyramid](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageGaussianPyramid) Autorelease() ImageGaussianPyramid {
	rv := objc.Send[ImageGaussianPyramid](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageGaussianPyramid creates a new ImageGaussianPyramid instance.
func NewImageGaussianPyramid() ImageGaussianPyramid {
	return getImageGaussianPyramidClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageGaussianPyramid */
// A filter that convolves an image with a Gaussian pyramid.
//
// The Gaussian image pyramid kernel is enqueued as an in-place operation using the method. All mip-map levels (after level 1) present in the provided image are filled using the provided filter kernel. The parameter is not used. The Gaussian image pyramid kernel ignores the and properties, and fills the entirety of the mip-map levels. Recall the size of the nth mip-map level as: Where and are the width and height of the 0th level, respectively (i.e. the image dimensions themselves). The Gaussian image pyramid is constructed as follows: First, the 0th level mip-map of the input image is filtered with the specified convolution kernel. The default convolution filter kernel is , where . You may also modify this kernel with a parameter of resulting in , where , or you may provide a completely custom kernel. Afterwards, the image is down-sampled by removing all odd rows and columns, which defines the next level in the Gaussian image pyramid. This procedure is continued until every mip-map level present in the image is filled with all the pyramid levels.


// A filter that convolves an image with a Gaussian pyramid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageGaussianPyramid
type ImageGaussianPyramid struct {
	ImagePyramid
}

// ImageGaussianPyramidFrom constructs a [ImageGaussianPyramid] from an unsafe.Pointer.
//
// A filter that convolves an image with a Gaussian pyramid.
func ImageGaussianPyramidFrom(ptr unsafe.Pointer) ImageGaussianPyramid {
	return ImageGaussianPyramid{
		ImagePyramid: ImagePyramidFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageGaussianPyramid *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageGaussianPyramid */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageGaussianPyramid */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageGaussianPyramid */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageGaussianPyramid */

// An optional clip rectangle to use when writing data. Only the pixels in the rectangle will be overwritten.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/cliprect
func (i_ ImageGaussianPyramid) ClipRect() objc.IObject /* cross-framework: MTLRegion */ {
	rv := objc.Send[Region](i_.ID, objc.Sel("clipRect"))
	return rv
}/* debug [instance_properties/getter]: clipRect */


// An optional clip rectangle to use when writing data. Only the pixels in the rectangle will be overwritten.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/cliprect
func (i_ ImageGaussianPyramid) SetClipRect(value objc.IObject /* cross-framework: MTLRegion */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setClipRect:"), value)
}/* debug [instance_properties/setter]: clipRect */


// The position of the destination clip rectangle origin relative to the source buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/offset
func (i_ ImageGaussianPyramid) Offset() objc.IObject /* cross-framework: MPSOffset */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("offset"))
	return rv
}/* debug [instance_properties/getter]: offset */


// The position of the destination clip rectangle origin relative to the source buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/offset
func (i_ ImageGaussianPyramid) SetOffset(value objc.IObject /* cross-framework: MPSOffset */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setOffset:"), value)
}/* debug [instance_properties/setter]: offset */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageGaussianPyramid */



