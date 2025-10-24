// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSImageStatisticsMinAndMax */


/* debug [class_header]: Header for MPSImageStatisticsMinAndMax */
// The class instance for the [ImageStatisticsMinAndMax] class.
var (
	ImageStatisticsMinAndMaxClass     _ImageStatisticsMinAndMaxClass
	ImageStatisticsMinAndMaxClassOnce sync.Once
)

func getImageStatisticsMinAndMaxClass() _ImageStatisticsMinAndMaxClass {
	ImageStatisticsMinAndMaxClassOnce.Do(func() {
		ImageStatisticsMinAndMaxClass = _ImageStatisticsMinAndMaxClass{objc.GetClass("MPSImageStatisticsMinAndMax")}
	})
	return ImageStatisticsMinAndMaxClass
}

type _ImageStatisticsMinAndMaxClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageStatisticsMinAndMax */
// An interface definition for the [ImageStatisticsMinAndMax] class.
type IImageStatisticsMinAndMax interface {
	IUnaryImageKernel
	
/* debug [class_interface_properties]: Properties for ImageStatisticsMinAndMax */
	// properties:
	ClipRectSource() Region get set /* not a class type */
	SetClipRectSource(value Region get set /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageStatisticsMinAndMax */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageStatisticsMinAndMax */
// Alloc allocates a new instance without initialization.
func (ic _ImageStatisticsMinAndMaxClass) Alloc() ImageStatisticsMinAndMax {
	rv := objc.Send[ImageStatisticsMinAndMax](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageStatisticsMinAndMaxClass) New() ImageStatisticsMinAndMax {
	rv := objc.Send[ImageStatisticsMinAndMax](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageStatisticsMinAndMax) Init() ImageStatisticsMinAndMax {
	rv := objc.Send[ImageStatisticsMinAndMax](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageStatisticsMinAndMax) Autorelease() ImageStatisticsMinAndMax {
	rv := objc.Send[ImageStatisticsMinAndMax](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageStatisticsMinAndMax creates a new ImageStatisticsMinAndMax instance.
func NewImageStatisticsMinAndMax() ImageStatisticsMinAndMax {
	return getImageStatisticsMinAndMaxClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageStatisticsMinAndMax */
// A kernel that computes the minimum and maximum pixel values for a given region of an image.
//
// The minimum and maximum values are written to the destination image at the following pixel locations: Minimum value is written at pixel location Maximum value is written at pixel location


// A kernel that computes the minimum and maximum pixel values for a given region of an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageStatisticsMinAndMax
type ImageStatisticsMinAndMax struct {
	UnaryImageKernel
}

// ImageStatisticsMinAndMaxFrom constructs a [ImageStatisticsMinAndMax] from an unsafe.Pointer.
//
// A kernel that computes the minimum and maximum pixel values for a given region of an image.
func ImageStatisticsMinAndMaxFrom(ptr unsafe.Pointer) ImageStatisticsMinAndMax {
	return ImageStatisticsMinAndMax{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageStatisticsMinAndMax */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagestatisticsminandmax/2867026-initwithcoder
func NewImageStatisticsMinAndMaxWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) ImageStatisticsMinAndMax {
	instance := getImageStatisticsMinAndMaxClass().Alloc()
	rv := objc.Send[ImageStatisticsMinAndMax](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageStatisticsMinAndMaxWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagestatisticsminandmax/2867125-initwithdevice
func NewImageStatisticsMinAndMaxWithDevice(device unsafe.Pointer) ImageStatisticsMinAndMax {
	instance := getImageStatisticsMinAndMaxClass().Alloc()
	rv := objc.Send[ImageStatisticsMinAndMax](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageStatisticsMinAndMaxWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageStatisticsMinAndMax */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageStatisticsMinAndMax */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageStatisticsMinAndMax */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageStatisticsMinAndMax */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagestatisticsminandmax/2867045-cliprectsource
func (i_ ImageStatisticsMinAndMax) ClipRectSource() Region get set /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("clipRectSource"))
	return rv
}/* debug [instance_properties/getter]: clipRectSource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagestatisticsminandmax/2867045-cliprectsource
func (i_ ImageStatisticsMinAndMax) SetClipRectSource(value Region get set /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setClipRectSource:"), value)
}/* debug [instance_properties/setter]: clipRectSource */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageStatisticsMinAndMax */


