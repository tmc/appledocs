// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSImageAreaMax */


/* debug [class_header]: Header for MPSImageAreaMax */
// The class instance for the [ImageAreaMax] class.
var (
	ImageAreaMaxClass     _ImageAreaMaxClass
	ImageAreaMaxClassOnce sync.Once
)

func getImageAreaMaxClass() _ImageAreaMaxClass {
	ImageAreaMaxClassOnce.Do(func() {
		ImageAreaMaxClass = _ImageAreaMaxClass{objc.GetClass("MPSImageAreaMax")}
	})
	return ImageAreaMaxClass
}

type _ImageAreaMaxClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageAreaMax */
// An interface definition for the [ImageAreaMax] class.
type IImageAreaMax interface {
	IUnaryImageKernel
	
/* debug [class_interface_properties]: Properties for ImageAreaMax */
	// properties:
	KernelHeight() objectivec.IObject
	SetKernelHeight(value objectivec.IObject)
	KernelWidth() objectivec.IObject
	SetKernelWidth(value objectivec.IObject)
	EdgeMode() ImageEdgeMode
	SetEdgeMode(value ImageEdgeMode)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageAreaMax */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageAreaMax */
// Alloc allocates a new instance without initialization.
func (ic _ImageAreaMaxClass) Alloc() ImageAreaMax {
	rv := objc.Send[ImageAreaMax](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageAreaMaxClass) New() ImageAreaMax {
	rv := objc.Send[ImageAreaMax](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageAreaMax) Init() ImageAreaMax {
	rv := objc.Send[ImageAreaMax](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageAreaMax) Autorelease() ImageAreaMax {
	rv := objc.Send[ImageAreaMax](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageAreaMax creates a new ImageAreaMax instance.
func NewImageAreaMax() ImageAreaMax {
	return getImageAreaMaxClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageAreaMax */
// A filter that finds the maximum pixel value in a rectangular region centered around each pixel in the source image.
//
// If there are multiple channels in the source image, each channel is processed independently. The property value is assumed to always be for this filter.


// A filter that finds the maximum pixel value in a rectangular region centered around each pixel in the source image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageAreaMax
type ImageAreaMax struct {
	UnaryImageKernel
}

// ImageAreaMaxFrom constructs a [ImageAreaMax] from an unsafe.Pointer.
//
// A filter that finds the maximum pixel value in a rectangular region centered around each pixel in the source image.
func ImageAreaMaxFrom(ptr unsafe.Pointer) ImageAreaMax {
	return ImageAreaMax{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageAreaMax */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageareamax/2866327-initwithcoder
func NewImageAreaMaxWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) ImageAreaMax {
	instance := getImageAreaMaxClass().Alloc()
	rv := objc.Send[ImageAreaMax](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageAreaMaxWithCoderDevice */


// Initializes the kernel with a specified width and height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageareamax/1618281-initwithdevice
func NewImageAreaMaxWithDeviceKernelWidthKernelHeight(device unsafe.Pointer, kernelWidth uint, kernelHeight uint) ImageAreaMax {
	instance := getImageAreaMaxClass().Alloc()
	rv := objc.Send[ImageAreaMax](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), device, kernelWidth, kernelHeight)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageAreaMaxWithDeviceKernelWidthKernelHeight */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageAreaMax */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageAreaMax */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageAreaMax */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageAreaMax */

// The height of the filter window. Must be an odd number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageareamax/1618277-kernelheight
func (i_ ImageAreaMax) KernelHeight() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("kernelHeight"))
	return rv
}/* debug [instance_properties/getter]: kernelHeight */


// The height of the filter window. Must be an odd number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageareamax/1618277-kernelheight
func (i_ ImageAreaMax) SetKernelHeight(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setKernelHeight:"), value)
}/* debug [instance_properties/setter]: kernelHeight */


// The width of the filter window. Must be an odd number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageareamax/1618282-kernelwidth
func (i_ ImageAreaMax) KernelWidth() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("kernelWidth"))
	return rv
}/* debug [instance_properties/getter]: kernelWidth */


// The width of the filter window. Must be an odd number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageareamax/1618282-kernelwidth
func (i_ ImageAreaMax) SetKernelWidth(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setKernelWidth:"), value)
}/* debug [instance_properties/setter]: kernelWidth */


// The edge mode to use when texture reads stray off the edge of an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/edgemode
func (i_ ImageAreaMax) EdgeMode() ImageEdgeMode {
	rv := objc.Send[ImageEdgeMode](i_.ID, objc.Sel("edgeMode"))
	return rv
}/* debug [instance_properties/getter]: edgeMode */


// The edge mode to use when texture reads stray off the edge of an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/edgemode
func (i_ ImageAreaMax) SetEdgeMode(value ImageEdgeMode) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setEdgeMode:"), value)
}/* debug [instance_properties/setter]: edgeMode */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageAreaMax */


