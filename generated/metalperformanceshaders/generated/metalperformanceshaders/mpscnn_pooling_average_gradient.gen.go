// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNPoolingAverageGradient */


/* debug [class_header]: Header for MPSCNNPoolingAverageGradient */
// The class instance for the [CNNPoolingAverageGradient] class.
var (
	CNNPoolingAverageGradientClass     _CNNPoolingAverageGradientClass
	CNNPoolingAverageGradientClassOnce sync.Once
)

func getCNNPoolingAverageGradientClass() _CNNPoolingAverageGradientClass {
	CNNPoolingAverageGradientClassOnce.Do(func() {
		CNNPoolingAverageGradientClass = _CNNPoolingAverageGradientClass{objc.GetClass("MPSCNNPoolingAverageGradient")}
	})
	return CNNPoolingAverageGradientClass
}

type _CNNPoolingAverageGradientClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNPoolingAverageGradient */
// An interface definition for the [CNNPoolingAverageGradient] class.
type ICNNPoolingAverageGradient interface {
	ICNNPoolingGradient
	
/* debug [class_interface_properties]: Properties for CNNPoolingAverageGradient */
	// properties:
	ZeroPadSizeX() objectivec.IObject
	SetZeroPadSizeX(value objectivec.IObject)
	ZeroPadSizeY() objectivec.IObject
	SetZeroPadSizeY(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNPoolingAverageGradient */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNPoolingAverageGradient */
// Alloc allocates a new instance without initialization.
func (cc _CNNPoolingAverageGradientClass) Alloc() CNNPoolingAverageGradient {
	rv := objc.Send[CNNPoolingAverageGradient](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNPoolingAverageGradientClass) New() CNNPoolingAverageGradient {
	rv := objc.Send[CNNPoolingAverageGradient](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNPoolingAverageGradient) Init() CNNPoolingAverageGradient {
	rv := objc.Send[CNNPoolingAverageGradient](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNPoolingAverageGradient) Autorelease() CNNPoolingAverageGradient {
	rv := objc.Send[CNNPoolingAverageGradient](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNPoolingAverageGradient creates a new CNNPoolingAverageGradient instance.
func NewCNNPoolingAverageGradient() CNNPoolingAverageGradient {
	return getCNNPoolingAverageGradientClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNPoolingAverageGradient */
// A gradient average pooling filter.


// A gradient average pooling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingAverageGradient
type CNNPoolingAverageGradient struct {
	CNNPoolingGradient
}

// CNNPoolingAverageGradientFrom constructs a [CNNPoolingAverageGradient] from an unsafe.Pointer.
//
// A gradient average pooling filter.
func CNNPoolingAverageGradientFrom(ptr unsafe.Pointer) CNNPoolingAverageGradient {
	return CNNPoolingAverageGradient{
		CNNPoolingGradient: CNNPoolingGradientFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNPoolingAverageGradient */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingaveragegradient/2942345-initwithcoder
func NewCNNPoolingAverageGradientWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) CNNPoolingAverageGradient {
	instance := getCNNPoolingAverageGradientClass().Alloc()
	rv := objc.Send[CNNPoolingAverageGradient](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNPoolingAverageGradientWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingaveragegradient/2942339-initwithdevice
func NewCNNPoolingAverageGradientWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device unsafe.Pointer, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) CNNPoolingAverageGradient {
	instance := getCNNPoolingAverageGradientClass().Alloc()
	rv := objc.Send[CNNPoolingAverageGradient](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:"), device, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNPoolingAverageGradientWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNPoolingAverageGradient */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNPoolingAverageGradient */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNPoolingAverageGradient */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNPoolingAverageGradient */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingaveragegradient/2942341-zeropadsizex
func (c_ CNNPoolingAverageGradient) ZeroPadSizeX() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("zeroPadSizeX"))
	return rv
}/* debug [instance_properties/getter]: zeroPadSizeX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingaveragegradient/2942341-zeropadsizex
func (c_ CNNPoolingAverageGradient) SetZeroPadSizeX(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setZeroPadSizeX:"), value)
}/* debug [instance_properties/setter]: zeroPadSizeX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingaveragegradient/2942354-zeropadsizey
func (c_ CNNPoolingAverageGradient) ZeroPadSizeY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("zeroPadSizeY"))
	return rv
}/* debug [instance_properties/getter]: zeroPadSizeY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingaveragegradient/2942354-zeropadsizey
func (c_ CNNPoolingAverageGradient) SetZeroPadSizeY(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setZeroPadSizeY:"), value)
}/* debug [instance_properties/setter]: zeroPadSizeY */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNPoolingAverageGradient */


