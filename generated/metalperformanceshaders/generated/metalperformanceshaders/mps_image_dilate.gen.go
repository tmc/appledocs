// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSImageDilate */


/* debug [class_header]: Header for MPSImageDilate */
// The class instance for the [ImageDilate] class.
var (
	ImageDilateClass     _ImageDilateClass
	ImageDilateClassOnce sync.Once
)

func getImageDilateClass() _ImageDilateClass {
	ImageDilateClassOnce.Do(func() {
		ImageDilateClass = _ImageDilateClass{objc.GetClass("MPSImageDilate")}
	})
	return ImageDilateClass
}

type _ImageDilateClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageDilate */
// An interface definition for the [ImageDilate] class.
type IImageDilate interface {
	IUnaryImageKernel
	
/* debug [class_interface_properties]: Properties for ImageDilate */
	// properties:
	KernelWidth() objectivec.IObject
	SetKernelWidth(value objectivec.IObject)
	KernelHeight() objectivec.IObject
	SetKernelHeight(value objectivec.IObject)
	EdgeMode() ImageEdgeMode
	SetEdgeMode(value ImageEdgeMode)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageDilate */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageDilate */
// Alloc allocates a new instance without initialization.
func (ic _ImageDilateClass) Alloc() ImageDilate {
	rv := objc.Send[ImageDilate](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageDilateClass) New() ImageDilate {
	rv := objc.Send[ImageDilate](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageDilate) Init() ImageDilate {
	rv := objc.Send[ImageDilate](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageDilate) Autorelease() ImageDilate {
	rv := objc.Send[ImageDilate](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageDilate creates a new ImageDilate instance.
func NewImageDilate() ImageDilate {
	return getImageDilateClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageDilate */
// A filter that finds the maximum pixel value in a rectangular region by applying a dilation function.
//
// An filter behaves like the filter, except Metal calculates the intensity at each position relative to a different value before determining which is the maximum pixel value, allowing for shaped, nonrectangular morphological probes. The code example below shows pseudocode for the calculation that returns each pixel value: A filter that contains all zeros is identical to an filter. Metal handles the center filter element as to avoid causing a general darkening of the image, and it handles the property as for this filter.


// A filter that finds the maximum pixel value in a rectangular region by applying a dilation function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageDilate
type ImageDilate struct {
	UnaryImageKernel
}

// ImageDilateFrom constructs a [ImageDilate] from an unsafe.Pointer.
//
// A filter that finds the maximum pixel value in a rectangular region by applying a dilation function.
func ImageDilateFrom(ptr unsafe.Pointer) ImageDilate {
	return ImageDilate{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageDilate */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedilate/2866325-initwithcoder
func NewImageDilateWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) ImageDilate {
	instance := getImageDilateClass().Alloc()
	rv := objc.Send[ImageDilate](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageDilateWithCoderDevice */


// Initializes the kernel with a specified width, height, and weight values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedilate/1618285-initwithdevice
func NewImageDilateWithDeviceKernelWidthKernelHeightValues(device unsafe.Pointer, kernelWidth uint, kernelHeight uint, values objectivec.IObject) ImageDilate {
	instance := getImageDilateClass().Alloc()
	rv := objc.Send[ImageDilate](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:values:"), device, kernelWidth, kernelHeight, values)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageDilateWithDeviceKernelWidthKernelHeightValues */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageDilate */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageDilate */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageDilate */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageDilate */

// The width of the filter window which must be an odd number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedilate/1618279-kernelwidth
func (i_ ImageDilate) KernelWidth() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("kernelWidth"))
	return rv
}/* debug [instance_properties/getter]: kernelWidth */


// The width of the filter window which must be an odd number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedilate/1618279-kernelwidth
func (i_ ImageDilate) SetKernelWidth(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setKernelWidth:"), value)
}/* debug [instance_properties/setter]: kernelWidth */


// The height of the filter window. which must be an odd number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedilate/1618280-kernelheight
func (i_ ImageDilate) KernelHeight() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("kernelHeight"))
	return rv
}/* debug [instance_properties/getter]: kernelHeight */


// The height of the filter window. which must be an odd number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedilate/1618280-kernelheight
func (i_ ImageDilate) SetKernelHeight(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setKernelHeight:"), value)
}/* debug [instance_properties/setter]: kernelHeight */


// The edge mode to use when texture reads stray off the edge of an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/edgemode
func (i_ ImageDilate) EdgeMode() ImageEdgeMode {
	rv := objc.Send[ImageEdgeMode](i_.ID, objc.Sel("edgeMode"))
	return rv
}/* debug [instance_properties/getter]: edgeMode */


// The edge mode to use when texture reads stray off the edge of an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/edgemode
func (i_ ImageDilate) SetEdgeMode(value ImageEdgeMode) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setEdgeMode:"), value)
}/* debug [instance_properties/setter]: edgeMode */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageDilate */


