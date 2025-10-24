// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNPooling */


/* debug [class_header]: Header for MPSCNNPooling */
// The class instance for the [CNNPooling] class.
var (
	CNNPoolingClass     _CNNPoolingClass
	CNNPoolingClassOnce sync.Once
)

func getCNNPoolingClass() _CNNPoolingClass {
	CNNPoolingClassOnce.Do(func() {
		CNNPoolingClass = _CNNPoolingClass{objc.GetClass("MPSCNNPooling")}
	})
	return CNNPoolingClass
}

type _CNNPoolingClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNPooling */
// An interface definition for the [CNNPooling] class.
type ICNNPooling interface {
	ICNNKernel
	
/* debug [class_interface_properties]: Properties for CNNPooling */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNPooling */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNPooling */
// Alloc allocates a new instance without initialization.
func (cc _CNNPoolingClass) Alloc() CNNPooling {
	rv := objc.Send[CNNPooling](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNPoolingClass) New() CNNPooling {
	rv := objc.Send[CNNPooling](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNPooling) Init() CNNPooling {
	rv := objc.Send[CNNPooling](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNPooling) Autorelease() CNNPooling {
	rv := objc.Send[CNNPooling](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNPooling creates a new CNNPooling instance.
func NewCNNPooling() CNNPooling {
	return getCNNPoolingClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNPooling */
// A pooling kernel.
//
// Pooling is a form of non-linear sub-sampling. Pooling partitions the input image into a set of rectangles (overlapping or non-overlapping) and, for each such sub-region, outputs a value. The pooling operation is used in computer vision to reduce the dimensionality of intermediate representations. The encode methods in the class can be used to encode an object to a object. The exact location of the pooling window for each output value is determined as follows: The pooling window center for the first (top left) output pixel of the clip rectangle is at spatial coordinates in the input image. From this, the top left corner of the pooling window is at , and extends pixels to the right and down direction, which means that the last pixel to be included into the pooling window is at , , so that for even kernel sizes the pooling window extends one pixel more into the left and up direction. The following pooling windows can be then easily deduced from the first one by simple shifting the source coordinates according to the values of the and properties. For example, the pooling window center for the output value at coordinate of the destination clip rectangle ( computed with regard to clipping rectangle origin) is at . Quite often it is desirable to distribute the pooling windows as evenly as possible in the input image. As explained above, if the is zero, then the center of the first pooling window is at the top left corner of the input image, which means that the left and top stripes of the pooling window are read from outside the input image boundaries (when filter size is larger than unity). Also it may mean that some values from the bottom and right stripes are not included at all in the pooling, resulting in loss of valuable information. A scheme used in some common libraries is to shift the source according to the following formula: , for odd for even Where is the size of the input image (or more accurately the size corresponding to the scaled value in source coordinates, which commonly coincides with the source image itself), is , and is . This offset distributes the pooling window centers evenly in the effective source , when the output size is rounded up with regards to stride ( ) and is commonly used in CNN libraries (for example uses this offset scheme in its maximum pooling implementation with - padding, for padding one can simply set to get the first pooling window inside the source image completely). For an object, the way the input image borders are handled can become important: if there are negative values in the source image near the borders of the image and the pooling window crosses the borders, then using a edge modemay cause the maximum pooling operation to override the negative input data values with zeros coming from outside the source image borders, resulting in large boundary effects. A simple way to avoid this is to use a edge mode, which for an object effectively causes all pooling windows to remain within the source image.


// A pooling kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPooling
type CNNPooling struct {
	CNNKernel
}

// CNNPoolingFrom constructs a [CNNPooling] from an unsafe.Pointer.
//
// A pooling kernel.
func CNNPoolingFrom(ptr unsafe.Pointer) CNNPooling {
	return CNNPooling{
		CNNKernel: CNNKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNPooling */

// Initializes a pooling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPooling/init(coder:device:)
func NewCNNPoolingWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) CNNPooling {
	instance := getCNNPoolingClass().Alloc()
	rv := objc.Send[CNNPooling](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNPoolingWithCoderDevice */


// Initializes a pooling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpooling/1648887-initwithdevice
func NewCNNPoolingWithDeviceKernelWidthKernelHeight(device unsafe.Pointer, kernelWidth uint, kernelHeight uint) CNNPooling {
	instance := getCNNPoolingClass().Alloc()
	rv := objc.Send[CNNPooling](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), device, kernelWidth, kernelHeight)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNPoolingWithDeviceKernelWidthKernelHeight */


// Initializes a pooling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpooling/1648902-initwithdevice
func NewCNNPoolingWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device unsafe.Pointer, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) CNNPooling {
	instance := getCNNPoolingClass().Alloc()
	rv := objc.Send[CNNPooling](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:"), device, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNPoolingWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNPooling */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNPooling */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNPooling */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNPooling */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNPooling */


