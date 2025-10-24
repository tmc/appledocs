// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNPoolingMax */


/* debug [class_header]: Header for MPSCNNPoolingMax */
// The class instance for the [CNNPoolingMax] class.
var (
	CNNPoolingMaxClass     _CNNPoolingMaxClass
	CNNPoolingMaxClassOnce sync.Once
)

func getCNNPoolingMaxClass() _CNNPoolingMaxClass {
	CNNPoolingMaxClassOnce.Do(func() {
		CNNPoolingMaxClass = _CNNPoolingMaxClass{objc.GetClass("MPSCNNPoolingMax")}
	})
	return CNNPoolingMaxClass
}

type _CNNPoolingMaxClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNPoolingMax */
// An interface definition for the [CNNPoolingMax] class.
type ICNNPoolingMax interface {
	ICNNPooling
	
/* debug [class_interface_properties]: Properties for CNNPoolingMax */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNPoolingMax */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNPoolingMax */
// Alloc allocates a new instance without initialization.
func (cc _CNNPoolingMaxClass) Alloc() CNNPoolingMax {
	rv := objc.Send[CNNPoolingMax](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNPoolingMaxClass) New() CNNPoolingMax {
	rv := objc.Send[CNNPoolingMax](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNPoolingMax) Init() CNNPoolingMax {
	rv := objc.Send[CNNPoolingMax](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNPoolingMax) Autorelease() CNNPoolingMax {
	rv := objc.Send[CNNPoolingMax](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNPoolingMax creates a new CNNPoolingMax instance.
func NewCNNPoolingMax() CNNPoolingMax {
	return getCNNPoolingMaxClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNPoolingMax */
// A max pooling filter.
//
// For each pixel in an image, the filter returns the maximum value of the pixels in the filter region defined by x .


// A max pooling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingMax
type CNNPoolingMax struct {
	CNNPooling
}

// CNNPoolingMaxFrom constructs a [CNNPoolingMax] from an unsafe.Pointer.
//
// A max pooling filter.
func CNNPoolingMaxFrom(ptr unsafe.Pointer) CNNPoolingMax {
	return CNNPoolingMax{
		CNNPooling: CNNPoolingFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNPoolingMax */

// Initializes a max pooling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingmax/2867097-initwithcoder
func NewCNNPoolingMaxWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) CNNPoolingMax {
	instance := getCNNPoolingMaxClass().Alloc()
	rv := objc.Send[CNNPoolingMax](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNPoolingMaxWithCoderDevice */


// Initializes a max pooling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingmax/2875151-initwithdevice
func NewCNNPoolingMaxWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device unsafe.Pointer, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) CNNPoolingMax {
	instance := getCNNPoolingMaxClass().Alloc()
	rv := objc.Send[CNNPoolingMax](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:"), device, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNPoolingMaxWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNPoolingMax */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNPoolingMax */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNPoolingMax */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNPoolingMax */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNPoolingMax */


