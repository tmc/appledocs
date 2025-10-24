// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNPoolingGradient */


/* debug [class_header]: Header for MPSCNNPoolingGradient */
// The class instance for the [CNNPoolingGradient] class.
var (
	CNNPoolingGradientClass     _CNNPoolingGradientClass
	CNNPoolingGradientClassOnce sync.Once
)

func getCNNPoolingGradientClass() _CNNPoolingGradientClass {
	CNNPoolingGradientClassOnce.Do(func() {
		CNNPoolingGradientClass = _CNNPoolingGradientClass{objc.GetClass("MPSCNNPoolingGradient")}
	})
	return CNNPoolingGradientClass
}

type _CNNPoolingGradientClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNPoolingGradient */
// An interface definition for the [CNNPoolingGradient] class.
type ICNNPoolingGradient interface {
	ICNNGradientKernel
	
/* debug [class_interface_properties]: Properties for CNNPoolingGradient */
	// properties:
	SourceSize() Size get set /* not a class type */
	SetSourceSize(value Size get set /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNPoolingGradient */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNPoolingGradient */
// Alloc allocates a new instance without initialization.
func (cc _CNNPoolingGradientClass) Alloc() CNNPoolingGradient {
	rv := objc.Send[CNNPoolingGradient](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNPoolingGradientClass) New() CNNPoolingGradient {
	rv := objc.Send[CNNPoolingGradient](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNPoolingGradient) Init() CNNPoolingGradient {
	rv := objc.Send[CNNPoolingGradient](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNPoolingGradient) Autorelease() CNNPoolingGradient {
	rv := objc.Send[CNNPoolingGradient](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNPoolingGradient creates a new CNNPoolingGradient instance.
func NewCNNPoolingGradient() CNNPoolingGradient {
	return getCNNPoolingGradientClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNPoolingGradient */
// A gradient pooling kernel.


// A gradient pooling kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingGradient
type CNNPoolingGradient struct {
	CNNGradientKernel
}

// CNNPoolingGradientFrom constructs a [CNNPoolingGradient] from an unsafe.Pointer.
//
// A gradient pooling kernel.
func CNNPoolingGradientFrom(ptr unsafe.Pointer) CNNPoolingGradient {
	return CNNPoolingGradient{
		CNNGradientKernel: CNNGradientKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNPoolingGradient */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradient/2942350-initwithcoder
func NewCNNPoolingGradientWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) CNNPoolingGradient {
	instance := getCNNPoolingGradientClass().Alloc()
	rv := objc.Send[CNNPoolingGradient](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNPoolingGradientWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradient/2942337-initwithdevice
func NewCNNPoolingGradientWithDeviceKernelWidthKernelHeight(device unsafe.Pointer, kernelWidth uint, kernelHeight uint) CNNPoolingGradient {
	instance := getCNNPoolingGradientClass().Alloc()
	rv := objc.Send[CNNPoolingGradient](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), device, kernelWidth, kernelHeight)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNPoolingGradientWithDeviceKernelWidthKernelHeight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradient/2942347-initwithdevice
func NewCNNPoolingGradientWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device unsafe.Pointer, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) CNNPoolingGradient {
	instance := getCNNPoolingGradientClass().Alloc()
	rv := objc.Send[CNNPoolingGradient](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:"), device, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNPoolingGradientWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNPoolingGradient */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNPoolingGradient */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNPoolingGradient */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNPoolingGradient */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradient/2942343-sourcesize
func (c_ CNNPoolingGradient) SourceSize() Size get set /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("sourceSize"))
	return rv
}/* debug [instance_properties/getter]: sourceSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradient/2942343-sourcesize
func (c_ CNNPoolingGradient) SetSourceSize(value Size get set /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSourceSize:"), value)
}/* debug [instance_properties/setter]: sourceSize */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNPoolingGradient */


