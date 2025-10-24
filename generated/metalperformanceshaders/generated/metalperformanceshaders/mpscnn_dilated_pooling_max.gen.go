// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNDilatedPoolingMax */


/* debug [class_header]: Header for MPSCNNDilatedPoolingMax */
// The class instance for the [CNNDilatedPoolingMax] class.
var (
	CNNDilatedPoolingMaxClass     _CNNDilatedPoolingMaxClass
	CNNDilatedPoolingMaxClassOnce sync.Once
)

func getCNNDilatedPoolingMaxClass() _CNNDilatedPoolingMaxClass {
	CNNDilatedPoolingMaxClassOnce.Do(func() {
		CNNDilatedPoolingMaxClass = _CNNDilatedPoolingMaxClass{objc.GetClass("MPSCNNDilatedPoolingMax")}
	})
	return CNNDilatedPoolingMaxClass
}

type _CNNDilatedPoolingMaxClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNDilatedPoolingMax */
// An interface definition for the [CNNDilatedPoolingMax] class.
type ICNNDilatedPoolingMax interface {
	ICNNPooling
	
/* debug [class_interface_properties]: Properties for CNNDilatedPoolingMax */
	// properties:
	DilationRateY() objectivec.IObject
	SetDilationRateY(value objectivec.IObject)
	DilationRateX() objectivec.IObject
	SetDilationRateX(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNDilatedPoolingMax */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNDilatedPoolingMax */
// Alloc allocates a new instance without initialization.
func (cc _CNNDilatedPoolingMaxClass) Alloc() CNNDilatedPoolingMax {
	rv := objc.Send[CNNDilatedPoolingMax](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNDilatedPoolingMaxClass) New() CNNDilatedPoolingMax {
	rv := objc.Send[CNNDilatedPoolingMax](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNDilatedPoolingMax) Init() CNNDilatedPoolingMax {
	rv := objc.Send[CNNDilatedPoolingMax](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNDilatedPoolingMax) Autorelease() CNNDilatedPoolingMax {
	rv := objc.Send[CNNDilatedPoolingMax](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNDilatedPoolingMax creates a new CNNDilatedPoolingMax instance.
func NewCNNDilatedPoolingMax() CNNDilatedPoolingMax {
	return getCNNDilatedPoolingMaxClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNDilatedPoolingMax */
// A dilated max pooling filter.
//
// For each pixel, returns the maximum value of pixels in the filter region by step size .


// A dilated max pooling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDilatedPoolingMax
type CNNDilatedPoolingMax struct {
	CNNPooling
}

// CNNDilatedPoolingMaxFrom constructs a [CNNDilatedPoolingMax] from an unsafe.Pointer.
//
// A dilated max pooling filter.
func CNNDilatedPoolingMaxFrom(ptr unsafe.Pointer) CNNDilatedPoolingMax {
	return CNNDilatedPoolingMax{
		CNNPooling: CNNPoolingFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNDilatedPoolingMax */

// Initializes a dilated max pooling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndilatedpoolingmax/2873025-initwithcoder
func NewCNNDilatedPoolingMaxWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) CNNDilatedPoolingMax {
	instance := getCNNDilatedPoolingMaxClass().Alloc()
	rv := objc.Send[CNNDilatedPoolingMax](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNDilatedPoolingMaxWithCoderDevice */


// Initializes a dilated max pooling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndilatedpoolingmax/2881192-initwithdevice
func NewCNNDilatedPoolingMaxWithDeviceKernelWidthKernelHeightDilationRateXDilationRateYStrideInPixelsXStrideInPixelsY(device unsafe.Pointer, kernelWidth uint, kernelHeight uint, dilationRateX uint, dilationRateY uint, strideInPixelsX uint, strideInPixelsY uint) CNNDilatedPoolingMax {
	instance := getCNNDilatedPoolingMaxClass().Alloc()
	rv := objc.Send[CNNDilatedPoolingMax](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:dilationRateX:dilationRateY:strideInPixelsX:strideInPixelsY:"), device, kernelWidth, kernelHeight, dilationRateX, dilationRateY, strideInPixelsX, strideInPixelsY)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNDilatedPoolingMaxWithDeviceKernelWidthKernelHeightDilationRateXDilationRateYStrideInPixelsXStrideInPixelsY */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNDilatedPoolingMax */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNDilatedPoolingMax */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNDilatedPoolingMax */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNDilatedPoolingMax */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndilatedpoolingmax/2881193-dilationratey
func (c_ CNNDilatedPoolingMax) DilationRateY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("dilationRateY"))
	return rv
}/* debug [instance_properties/getter]: dilationRateY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndilatedpoolingmax/2881193-dilationratey
func (c_ CNNDilatedPoolingMax) SetDilationRateY(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDilationRateY:"), value)
}/* debug [instance_properties/setter]: dilationRateY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndilatedpoolingmax/2881194-dilationratex
func (c_ CNNDilatedPoolingMax) DilationRateX() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("dilationRateX"))
	return rv
}/* debug [instance_properties/getter]: dilationRateX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndilatedpoolingmax/2881194-dilationratex
func (c_ CNNDilatedPoolingMax) SetDilationRateX(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDilationRateX:"), value)
}/* debug [instance_properties/setter]: dilationRateX */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNDilatedPoolingMax */


