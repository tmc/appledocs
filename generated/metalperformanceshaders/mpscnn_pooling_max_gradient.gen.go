// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNPoolingMaxGradient */


/* debug [class_header]: Header for MPSCNNPoolingMaxGradient */
// The class instance for the [CNNPoolingMaxGradient] class.
var (
	CNNPoolingMaxGradientClass     _CNNPoolingMaxGradientClass
	CNNPoolingMaxGradientClassOnce sync.Once
)

func getCNNPoolingMaxGradientClass() _CNNPoolingMaxGradientClass {
	CNNPoolingMaxGradientClassOnce.Do(func() {
		CNNPoolingMaxGradientClass = _CNNPoolingMaxGradientClass{objc.GetClass("MPSCNNPoolingMaxGradient")}
	})
	return CNNPoolingMaxGradientClass
}

type _CNNPoolingMaxGradientClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNPoolingMaxGradient */
// An interface definition for the [CNNPoolingMaxGradient] class.
type ICNNPoolingMaxGradient interface {
	ICNNPoolingGradient
	
/* debug [class_interface_properties]: Properties for CNNPoolingMaxGradient */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNPoolingMaxGradient */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNPoolingMaxGradient */
// Alloc allocates a new instance without initialization.
func (cc _CNNPoolingMaxGradientClass) Alloc() CNNPoolingMaxGradient {
	rv := objc.Send[CNNPoolingMaxGradient](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNPoolingMaxGradientClass) New() CNNPoolingMaxGradient {
	rv := objc.Send[CNNPoolingMaxGradient](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNPoolingMaxGradient) Init() CNNPoolingMaxGradient {
	rv := objc.Send[CNNPoolingMaxGradient](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNPoolingMaxGradient) Autorelease() CNNPoolingMaxGradient {
	rv := objc.Send[CNNPoolingMaxGradient](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNPoolingMaxGradient creates a new CNNPoolingMaxGradient instance.
func NewCNNPoolingMaxGradient() CNNPoolingMaxGradient {
	return getCNNPoolingMaxGradientClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNPoolingMaxGradient */
// A gradient max pooling filter.


// A gradient max pooling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingMaxGradient
type CNNPoolingMaxGradient struct {
	CNNPoolingGradient
}

// CNNPoolingMaxGradientFrom constructs a [CNNPoolingMaxGradient] from an unsafe.Pointer.
//
// A gradient max pooling filter.
func CNNPoolingMaxGradientFrom(ptr unsafe.Pointer) CNNPoolingMaxGradient {
	return CNNPoolingMaxGradient{
		CNNPoolingGradient: CNNPoolingGradientFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNPoolingMaxGradient */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingmaxgradient/2942342-initwithcoder
func NewCNNPoolingMaxGradientWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) CNNPoolingMaxGradient {
	instance := getCNNPoolingMaxGradientClass().Alloc()
	rv := objc.Send[CNNPoolingMaxGradient](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNPoolingMaxGradientWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingmaxgradient/2942348-initwithdevice
func NewCNNPoolingMaxGradientWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device unsafe.Pointer, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) CNNPoolingMaxGradient {
	instance := getCNNPoolingMaxGradientClass().Alloc()
	rv := objc.Send[CNNPoolingMaxGradient](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:"), device, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNPoolingMaxGradientWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNPoolingMaxGradient */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNPoolingMaxGradient */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNPoolingMaxGradient */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNPoolingMaxGradient */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNPoolingMaxGradient */


