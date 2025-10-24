// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSImageStatisticsMean */


/* debug [class_header]: Header for MPSImageStatisticsMean */
// The class instance for the [ImageStatisticsMean] class.
var (
	ImageStatisticsMeanClass     _ImageStatisticsMeanClass
	ImageStatisticsMeanClassOnce sync.Once
)

func getImageStatisticsMeanClass() _ImageStatisticsMeanClass {
	ImageStatisticsMeanClassOnce.Do(func() {
		ImageStatisticsMeanClass = _ImageStatisticsMeanClass{objc.GetClass("MPSImageStatisticsMean")}
	})
	return ImageStatisticsMeanClass
}

type _ImageStatisticsMeanClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageStatisticsMean */
// An interface definition for the [ImageStatisticsMean] class.
type IImageStatisticsMean interface {
	IUnaryImageKernel
	
/* debug [class_interface_properties]: Properties for ImageStatisticsMean */
	// properties:
	ClipRectSource() Region get set /* not a class type */
	SetClipRectSource(value Region get set /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageStatisticsMean */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageStatisticsMean */
// Alloc allocates a new instance without initialization.
func (ic _ImageStatisticsMeanClass) Alloc() ImageStatisticsMean {
	rv := objc.Send[ImageStatisticsMean](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageStatisticsMeanClass) New() ImageStatisticsMean {
	rv := objc.Send[ImageStatisticsMean](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageStatisticsMean) Init() ImageStatisticsMean {
	rv := objc.Send[ImageStatisticsMean](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageStatisticsMean) Autorelease() ImageStatisticsMean {
	rv := objc.Send[ImageStatisticsMean](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageStatisticsMean creates a new ImageStatisticsMean instance.
func NewImageStatisticsMean() ImageStatisticsMean {
	return getImageStatisticsMeanClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageStatisticsMean */
// A kernel that computes the mean for a given region of an image.


// A kernel that computes the mean for a given region of an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageStatisticsMean
type ImageStatisticsMean struct {
	UnaryImageKernel
}

// ImageStatisticsMeanFrom constructs a [ImageStatisticsMean] from an unsafe.Pointer.
//
// A kernel that computes the mean for a given region of an image.
func ImageStatisticsMeanFrom(ptr unsafe.Pointer) ImageStatisticsMean {
	return ImageStatisticsMean{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageStatisticsMean */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagestatisticsmean/2867124-initwithcoder
func NewImageStatisticsMeanWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) ImageStatisticsMean {
	instance := getImageStatisticsMeanClass().Alloc()
	rv := objc.Send[ImageStatisticsMean](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageStatisticsMeanWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagestatisticsmean/2867156-initwithdevice
func NewImageStatisticsMeanWithDevice(device unsafe.Pointer) ImageStatisticsMean {
	instance := getImageStatisticsMeanClass().Alloc()
	rv := objc.Send[ImageStatisticsMean](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageStatisticsMeanWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageStatisticsMean */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageStatisticsMean */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageStatisticsMean */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageStatisticsMean */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagestatisticsmean/2867093-cliprectsource
func (i_ ImageStatisticsMean) ClipRectSource() Region get set /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("clipRectSource"))
	return rv
}/* debug [instance_properties/getter]: clipRectSource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagestatisticsmean/2867093-cliprectsource
func (i_ ImageStatisticsMean) SetClipRectSource(value Region get set /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setClipRectSource:"), value)
}/* debug [instance_properties/setter]: clipRectSource */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageStatisticsMean */


