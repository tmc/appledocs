// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSImageBox */


/* debug [class_header]: Header for MPSImageBox */
// The class instance for the [ImageBox] class.
var (
	ImageBoxClass     _ImageBoxClass
	ImageBoxClassOnce sync.Once
)

func getImageBoxClass() _ImageBoxClass {
	ImageBoxClassOnce.Do(func() {
		ImageBoxClass = _ImageBoxClass{objc.GetClass("MPSImageBox")}
	})
	return ImageBoxClass
}

type _ImageBoxClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageBox */
// An interface definition for the [ImageBox] class.
type IImageBox interface {
	IUnaryImageKernel
	
/* debug [class_interface_properties]: Properties for ImageBox */
	// properties:
	KernelHeight() objectivec.IObject
	SetKernelHeight(value objectivec.IObject)
	KernelWidth() objectivec.IObject
	SetKernelWidth(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageBox */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageBox */
// Alloc allocates a new instance without initialization.
func (ic _ImageBoxClass) Alloc() ImageBox {
	rv := objc.Send[ImageBox](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageBoxClass) New() ImageBox {
	rv := objc.Send[ImageBox](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageBox) Init() ImageBox {
	rv := objc.Send[ImageBox](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageBox) Autorelease() ImageBox {
	rv := objc.Send[ImageBox](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageBox creates a new ImageBox instance.
func NewImageBox() ImageBox {
	return getImageBoxClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageBox */
// A filter that convolves an image with a given kernel of odd width and height.
//
// The kernel elements all have equal weight, achieving a blur effect (each result is the unweighted average of the surrounding pixels). This allows for much faster algorithms, especially for larger blur radii. The box height and width must be odd numbers. The box blur is a separable filter and the Metal Performance Shaders framework will act accordingly to give best performance for multi-dimensional blurs.


// A filter that convolves an image with a given kernel of odd width and height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageBox
type ImageBox struct {
	UnaryImageKernel
}

// ImageBoxFrom constructs a [ImageBox] from an unsafe.Pointer.
//
// A filter that convolves an image with a given kernel of odd width and height.
func ImageBoxFrom(ptr unsafe.Pointer) ImageBox {
	return ImageBox{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageBox */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagebox/2866153-initwithcoder
func NewImageBoxWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) ImageBox {
	instance := getImageBoxClass().Alloc()
	rv := objc.Send[ImageBox](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageBoxWithCoderDevice */


// Initializes a box filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagebox/1618789-initwithdevice
func NewImageBoxWithDeviceKernelWidthKernelHeight(device unsafe.Pointer, kernelWidth uint, kernelHeight uint) ImageBox {
	instance := getImageBoxClass().Alloc()
	rv := objc.Send[ImageBox](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), device, kernelWidth, kernelHeight)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageBoxWithDeviceKernelWidthKernelHeight */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageBox */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageBox */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageBox */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageBox */

// The height of the filter window. Must be an odd number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagebox/1618739-kernelheight
func (i_ ImageBox) KernelHeight() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("kernelHeight"))
	return rv
}/* debug [instance_properties/getter]: kernelHeight */


// The height of the filter window. Must be an odd number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagebox/1618739-kernelheight
func (i_ ImageBox) SetKernelHeight(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setKernelHeight:"), value)
}/* debug [instance_properties/setter]: kernelHeight */


// The width of the filter window. Must be an odd number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagebox/1618834-kernelwidth
func (i_ ImageBox) KernelWidth() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("kernelWidth"))
	return rv
}/* debug [instance_properties/getter]: kernelWidth */


// The width of the filter window. Must be an odd number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagebox/1618834-kernelwidth
func (i_ ImageBox) SetKernelWidth(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setKernelWidth:"), value)
}/* debug [instance_properties/setter]: kernelWidth */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageBox */


