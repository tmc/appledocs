// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSImageErode */


/* debug [class_header]: Header for MPSImageErode */
// The class instance for the [ImageErode] class.
var (
	ImageErodeClass     _ImageErodeClass
	ImageErodeClassOnce sync.Once
)

func getImageErodeClass() _ImageErodeClass {
	ImageErodeClassOnce.Do(func() {
		ImageErodeClass = _ImageErodeClass{objc.GetClass("MPSImageErode")}
	})
	return ImageErodeClass
}

type _ImageErodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageErode */
// An interface definition for the [ImageErode] class.
type IImageErode interface {
	IImageDilate
	
/* debug [class_interface_properties]: Properties for ImageErode */
	// properties:
	EdgeMode() ImageEdgeMode
	SetEdgeMode(value ImageEdgeMode)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageErode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageErode */
// Alloc allocates a new instance without initialization.
func (ic _ImageErodeClass) Alloc() ImageErode {
	rv := objc.Send[ImageErode](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageErodeClass) New() ImageErode {
	rv := objc.Send[ImageErode](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageErode) Init() ImageErode {
	rv := objc.Send[ImageErode](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageErode) Autorelease() ImageErode {
	rv := objc.Send[ImageErode](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageErode creates a new ImageErode instance.
func NewImageErode() ImageErode {
	return getImageErodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageErode */
// A filter that finds the minimum pixel value in a rectangular region by applying an erosion function.
//
// An behaves like the filter, except that Metal calculates the intensity at each position relative to a different value before determining which is the maximum pixel value, allowing for shaped, nonrectangular morphological probes. The code example below shows pseudocode for the calculation that returns each pixel value: The definition of the filter is different from its counterpart ( ). This allows and to use the same filter, making open and close operators easier to write. A filter that contains all zeros is identical to an filter. Metal handles the center filter element as to avoid causing a general lightening of the image, and it handles the property as for this filter.


// A filter that finds the minimum pixel value in a rectangular region by applying an erosion function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageErode
type ImageErode struct {
	ImageDilate
}

// ImageErodeFrom constructs a [ImageErode] from an unsafe.Pointer.
//
// A filter that finds the minimum pixel value in a rectangular region by applying an erosion function.
func ImageErodeFrom(ptr unsafe.Pointer) ImageErode {
	return ImageErode{
		ImageDilate: ImageDilateFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageErode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageErode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageErode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageErode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageErode */

// The edge mode to use when texture reads stray off the edge of an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/edgemode
func (i_ ImageErode) EdgeMode() ImageEdgeMode {
	rv := objc.Send[ImageEdgeMode](i_.ID, objc.Sel("edgeMode"))
	return rv
}/* debug [instance_properties/getter]: edgeMode */


// The edge mode to use when texture reads stray off the edge of an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/edgemode
func (i_ ImageErode) SetEdgeMode(value ImageEdgeMode) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setEdgeMode:"), value)
}/* debug [instance_properties/setter]: edgeMode */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageErode */



