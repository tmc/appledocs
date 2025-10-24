// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSImageStatisticsMeanAndVariance */


/* debug [class_header]: Header for MPSImageStatisticsMeanAndVariance */
// The class instance for the [ImageStatisticsMeanAndVariance] class.
var (
	ImageStatisticsMeanAndVarianceClass     _ImageStatisticsMeanAndVarianceClass
	ImageStatisticsMeanAndVarianceClassOnce sync.Once
)

func getImageStatisticsMeanAndVarianceClass() _ImageStatisticsMeanAndVarianceClass {
	ImageStatisticsMeanAndVarianceClassOnce.Do(func() {
		ImageStatisticsMeanAndVarianceClass = _ImageStatisticsMeanAndVarianceClass{objc.GetClass("MPSImageStatisticsMeanAndVariance")}
	})
	return ImageStatisticsMeanAndVarianceClass
}

type _ImageStatisticsMeanAndVarianceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageStatisticsMeanAndVariance */
// An interface definition for the [ImageStatisticsMeanAndVariance] class.
type IImageStatisticsMeanAndVariance interface {
	IUnaryImageKernel
	
/* debug [class_interface_properties]: Properties for ImageStatisticsMeanAndVariance */
	// properties:
	ClipRectSource() Region get set /* not a class type */
	SetClipRectSource(value Region get set /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageStatisticsMeanAndVariance */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageStatisticsMeanAndVariance */
// Alloc allocates a new instance without initialization.
func (ic _ImageStatisticsMeanAndVarianceClass) Alloc() ImageStatisticsMeanAndVariance {
	rv := objc.Send[ImageStatisticsMeanAndVariance](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageStatisticsMeanAndVarianceClass) New() ImageStatisticsMeanAndVariance {
	rv := objc.Send[ImageStatisticsMeanAndVariance](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageStatisticsMeanAndVariance) Init() ImageStatisticsMeanAndVariance {
	rv := objc.Send[ImageStatisticsMeanAndVariance](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageStatisticsMeanAndVariance) Autorelease() ImageStatisticsMeanAndVariance {
	rv := objc.Send[ImageStatisticsMeanAndVariance](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageStatisticsMeanAndVariance creates a new ImageStatisticsMeanAndVariance instance.
func NewImageStatisticsMeanAndVariance() ImageStatisticsMeanAndVariance {
	return getImageStatisticsMeanAndVarianceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageStatisticsMeanAndVariance */
// A kernel that computes the mean and variance for a given region of an image.
//
// The mean and variance values are written to the destination image at the following pixel locations: Mean value is written at pixel location Variance value is written at pixel location


// A kernel that computes the mean and variance for a given region of an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageStatisticsMeanAndVariance
type ImageStatisticsMeanAndVariance struct {
	UnaryImageKernel
}

// ImageStatisticsMeanAndVarianceFrom constructs a [ImageStatisticsMeanAndVariance] from an unsafe.Pointer.
//
// A kernel that computes the mean and variance for a given region of an image.
func ImageStatisticsMeanAndVarianceFrom(ptr unsafe.Pointer) ImageStatisticsMeanAndVariance {
	return ImageStatisticsMeanAndVariance{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageStatisticsMeanAndVariance */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagestatisticsmeanandvariance/2867044-initwithcoder
func NewImageStatisticsMeanAndVarianceWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) ImageStatisticsMeanAndVariance {
	instance := getImageStatisticsMeanAndVarianceClass().Alloc()
	rv := objc.Send[ImageStatisticsMeanAndVariance](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageStatisticsMeanAndVarianceWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagestatisticsmeanandvariance/2867165-initwithdevice
func NewImageStatisticsMeanAndVarianceWithDevice(device unsafe.Pointer) ImageStatisticsMeanAndVariance {
	instance := getImageStatisticsMeanAndVarianceClass().Alloc()
	rv := objc.Send[ImageStatisticsMeanAndVariance](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageStatisticsMeanAndVarianceWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageStatisticsMeanAndVariance */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageStatisticsMeanAndVariance */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageStatisticsMeanAndVariance */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageStatisticsMeanAndVariance */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagestatisticsmeanandvariance/2867131-cliprectsource
func (i_ ImageStatisticsMeanAndVariance) ClipRectSource() Region get set /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("clipRectSource"))
	return rv
}/* debug [instance_properties/getter]: clipRectSource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagestatisticsmeanandvariance/2867131-cliprectsource
func (i_ ImageStatisticsMeanAndVariance) SetClipRectSource(value Region get set /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setClipRectSource:"), value)
}/* debug [instance_properties/setter]: clipRectSource */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageStatisticsMeanAndVariance */


