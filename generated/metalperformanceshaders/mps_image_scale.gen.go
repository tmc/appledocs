// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSImageScale */


/* debug [class_header]: Header for MPSImageScale */
// The class instance for the [ImageScale] class.
var (
	ImageScaleClass     _ImageScaleClass
	ImageScaleClassOnce sync.Once
)

func getImageScaleClass() _ImageScaleClass {
	ImageScaleClassOnce.Do(func() {
		ImageScaleClass = _ImageScaleClass{objc.GetClass("MPSImageScale")}
	})
	return ImageScaleClass
}

type _ImageScaleClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageScale */
// An interface definition for the [ImageScale] class.
type IImageScale interface {
	IUnaryImageKernel
	
/* debug [class_interface_properties]: Properties for ImageScale */
	// properties:
	ScaleTransform() objectivec.IObject
	SetScaleTransform(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageScale */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageScale */
// Alloc allocates a new instance without initialization.
func (ic _ImageScaleClass) Alloc() ImageScale {
	rv := objc.Send[ImageScale](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageScaleClass) New() ImageScale {
	rv := objc.Send[ImageScale](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageScale) Init() ImageScale {
	rv := objc.Send[ImageScale](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageScale) Autorelease() ImageScale {
	rv := objc.Send[ImageScale](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageScale creates a new ImageScale instance.
func NewImageScale() ImageScale {
	return getImageScaleClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageScale */
// A filter that resizes and changes the aspect ratio of an image.


// A filter that resizes and changes the aspect ratio of an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageScale
type ImageScale struct {
	UnaryImageKernel
}

// ImageScaleFrom constructs a [ImageScale] from an unsafe.Pointer.
//
// A filter that resizes and changes the aspect ratio of an image.
func ImageScaleFrom(ptr unsafe.Pointer) ImageScale {
	return ImageScale{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageScale */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagescale/2881187-initwithcoder
func NewImageScaleWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) ImageScale {
	instance := getImageScaleClass().Alloc()
	rv := objc.Send[ImageScale](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageScaleWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagescale/2881186-initwithdevice
func NewImageScaleWithDevice(device unsafe.Pointer) ImageScale {
	instance := getImageScaleClass().Alloc()
	rv := objc.Send[ImageScale](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageScaleWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageScale */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageScale */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageScale */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageScale */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagescale/2881183-scaletransform
func (i_ ImageScale) ScaleTransform() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("scaleTransform"))
	return rv
}/* debug [instance_properties/getter]: scaleTransform */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagescale/2881183-scaletransform
func (i_ ImageScale) SetScaleTransform(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setScaleTransform:"), value)
}/* debug [instance_properties/setter]: scaleTransform */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageScale */


