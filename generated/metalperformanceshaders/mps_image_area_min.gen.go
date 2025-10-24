// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSImageAreaMin */


/* debug [class_header]: Header for MPSImageAreaMin */
// The class instance for the [ImageAreaMin] class.
var (
	ImageAreaMinClass     _ImageAreaMinClass
	ImageAreaMinClassOnce sync.Once
)

func getImageAreaMinClass() _ImageAreaMinClass {
	ImageAreaMinClassOnce.Do(func() {
		ImageAreaMinClass = _ImageAreaMinClass{objc.GetClass("MPSImageAreaMin")}
	})
	return ImageAreaMinClass
}

type _ImageAreaMinClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageAreaMin */
// An interface definition for the [ImageAreaMin] class.
type IImageAreaMin interface {
	IImageAreaMax
	
/* debug [class_interface_properties]: Properties for ImageAreaMin */
	// properties:
	EdgeMode() ImageEdgeMode
	SetEdgeMode(value ImageEdgeMode)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageAreaMin */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageAreaMin */
// Alloc allocates a new instance without initialization.
func (ic _ImageAreaMinClass) Alloc() ImageAreaMin {
	rv := objc.Send[ImageAreaMin](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageAreaMinClass) New() ImageAreaMin {
	rv := objc.Send[ImageAreaMin](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageAreaMin) Init() ImageAreaMin {
	rv := objc.Send[ImageAreaMin](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageAreaMin) Autorelease() ImageAreaMin {
	rv := objc.Send[ImageAreaMin](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageAreaMin creates a new ImageAreaMin instance.
func NewImageAreaMin() ImageAreaMin {
	return getImageAreaMinClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageAreaMin */
// A filter that finds the minimum pixel value in a rectangular region centered around each pixel in the source image.
//
// An filter has the same methods and properties as the class. If there are multiple channels in the source image, each channel is processed independently. The property value is assumed to always be for this filter.


// A filter that finds the minimum pixel value in a rectangular region centered around each pixel in the source image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageAreaMin
type ImageAreaMin struct {
	ImageAreaMax
}

// ImageAreaMinFrom constructs a [ImageAreaMin] from an unsafe.Pointer.
//
// A filter that finds the minimum pixel value in a rectangular region centered around each pixel in the source image.
func ImageAreaMinFrom(ptr unsafe.Pointer) ImageAreaMin {
	return ImageAreaMin{
		ImageAreaMax: ImageAreaMaxFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageAreaMin *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageAreaMin */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageAreaMin */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageAreaMin */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageAreaMin */

// The edge mode to use when texture reads stray off the edge of an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/edgemode
func (i_ ImageAreaMin) EdgeMode() ImageEdgeMode {
	rv := objc.Send[ImageEdgeMode](i_.ID, objc.Sel("edgeMode"))
	return rv
}/* debug [instance_properties/getter]: edgeMode */


// The edge mode to use when texture reads stray off the edge of an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/edgemode
func (i_ ImageAreaMin) SetEdgeMode(value ImageEdgeMode) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setEdgeMode:"), value)
}/* debug [instance_properties/setter]: edgeMode */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageAreaMin */



