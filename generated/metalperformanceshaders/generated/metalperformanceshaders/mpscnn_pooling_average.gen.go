// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNPoolingAverage */


/* debug [class_header]: Header for MPSCNNPoolingAverage */
// The class instance for the [CNNPoolingAverage] class.
var (
	CNNPoolingAverageClass     _CNNPoolingAverageClass
	CNNPoolingAverageClassOnce sync.Once
)

func getCNNPoolingAverageClass() _CNNPoolingAverageClass {
	CNNPoolingAverageClassOnce.Do(func() {
		CNNPoolingAverageClass = _CNNPoolingAverageClass{objc.GetClass("MPSCNNPoolingAverage")}
	})
	return CNNPoolingAverageClass
}

type _CNNPoolingAverageClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNPoolingAverage */
// An interface definition for the [CNNPoolingAverage] class.
type ICNNPoolingAverage interface {
	ICNNPooling
	
/* debug [class_interface_properties]: Properties for CNNPoolingAverage */
	// properties:
	ZeroPadSizeX() objectivec.IObject
	SetZeroPadSizeX(value objectivec.IObject)
	ZeroPadSizeY() objectivec.IObject
	SetZeroPadSizeY(value objectivec.IObject)
	EdgeMode() ImageEdgeMode
	SetEdgeMode(value ImageEdgeMode)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNPoolingAverage */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNPoolingAverage */
// Alloc allocates a new instance without initialization.
func (cc _CNNPoolingAverageClass) Alloc() CNNPoolingAverage {
	rv := objc.Send[CNNPoolingAverage](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNPoolingAverageClass) New() CNNPoolingAverage {
	rv := objc.Send[CNNPoolingAverage](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNPoolingAverage) Init() CNNPoolingAverage {
	rv := objc.Send[CNNPoolingAverage](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNPoolingAverage) Autorelease() CNNPoolingAverage {
	rv := objc.Send[CNNPoolingAverage](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNPoolingAverage creates a new CNNPoolingAverage instance.
func NewCNNPoolingAverage() CNNPoolingAverage {
	return getCNNPoolingAverageClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNPoolingAverage */
// An average pooling filter.
//
// For each pixel in an image, the filter returns the average value of the pixels in the filter region defined by . When the value of the property is set to , the filtering window is shrunk to remain within the source image borders. For pixels close to the image borders, the filtering window will be smaller in order to fit inside the source image and less values will be used to compute the average value. In case the filtering window is entirely outside the source image border, the output value will be .


// An average pooling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingAverage
type CNNPoolingAverage struct {
	CNNPooling
}

// CNNPoolingAverageFrom constructs a [CNNPoolingAverage] from an unsafe.Pointer.
//
// An average pooling filter.
func CNNPoolingAverageFrom(ptr unsafe.Pointer) CNNPoolingAverage {
	return CNNPoolingAverage{
		CNNPooling: CNNPoolingFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNPoolingAverage */

// Initializes an average pooling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingaverage/2866999-initwithcoder
func NewCNNPoolingAverageWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) CNNPoolingAverage {
	instance := getCNNPoolingAverageClass().Alloc()
	rv := objc.Send[CNNPoolingAverage](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNPoolingAverageWithCoderDevice */


// Initializes an average pooling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingaverage/2875216-initwithdevice
func NewCNNPoolingAverageWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device unsafe.Pointer, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) CNNPoolingAverage {
	instance := getCNNPoolingAverageClass().Alloc()
	rv := objc.Send[CNNPoolingAverage](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:"), device, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNPoolingAverageWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNPoolingAverage */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNPoolingAverage */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNPoolingAverage */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNPoolingAverage */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingaverage/2875207-zeropadsizex
func (c_ CNNPoolingAverage) ZeroPadSizeX() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("zeroPadSizeX"))
	return rv
}/* debug [instance_properties/getter]: zeroPadSizeX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingaverage/2875207-zeropadsizex
func (c_ CNNPoolingAverage) SetZeroPadSizeX(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setZeroPadSizeX:"), value)
}/* debug [instance_properties/setter]: zeroPadSizeX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingaverage/2875221-zeropadsizey
func (c_ CNNPoolingAverage) ZeroPadSizeY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("zeroPadSizeY"))
	return rv
}/* debug [instance_properties/getter]: zeroPadSizeY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingaverage/2875221-zeropadsizey
func (c_ CNNPoolingAverage) SetZeroPadSizeY(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setZeroPadSizeY:"), value)
}/* debug [instance_properties/setter]: zeroPadSizeY */


// The edge mode to use when texture reads stray off the edge of an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/edgemode
func (c_ CNNPoolingAverage) EdgeMode() ImageEdgeMode {
	rv := objc.Send[ImageEdgeMode](c_.ID, objc.Sel("edgeMode"))
	return rv
}/* debug [instance_properties/getter]: edgeMode */


// The edge mode to use when texture reads stray off the edge of an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/edgemode
func (c_ CNNPoolingAverage) SetEdgeMode(value ImageEdgeMode) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEdgeMode:"), value)
}/* debug [instance_properties/setter]: edgeMode */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNPoolingAverage */


