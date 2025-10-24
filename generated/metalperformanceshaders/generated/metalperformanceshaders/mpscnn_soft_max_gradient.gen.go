// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNSoftMaxGradient */


/* debug [class_header]: Header for MPSCNNSoftMaxGradient */
// The class instance for the [CNNSoftMaxGradient] class.
var (
	CNNSoftMaxGradientClass     _CNNSoftMaxGradientClass
	CNNSoftMaxGradientClassOnce sync.Once
)

func getCNNSoftMaxGradientClass() _CNNSoftMaxGradientClass {
	CNNSoftMaxGradientClassOnce.Do(func() {
		CNNSoftMaxGradientClass = _CNNSoftMaxGradientClass{objc.GetClass("MPSCNNSoftMaxGradient")}
	})
	return CNNSoftMaxGradientClass
}

type _CNNSoftMaxGradientClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNSoftMaxGradient */
// An interface definition for the [CNNSoftMaxGradient] class.
type ICNNSoftMaxGradient interface {
	ICNNGradientKernel
	
/* debug [class_interface_properties]: Properties for CNNSoftMaxGradient */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNSoftMaxGradient */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNSoftMaxGradient */
// Alloc allocates a new instance without initialization.
func (cc _CNNSoftMaxGradientClass) Alloc() CNNSoftMaxGradient {
	rv := objc.Send[CNNSoftMaxGradient](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNSoftMaxGradientClass) New() CNNSoftMaxGradient {
	rv := objc.Send[CNNSoftMaxGradient](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNSoftMaxGradient) Init() CNNSoftMaxGradient {
	rv := objc.Send[CNNSoftMaxGradient](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNSoftMaxGradient) Autorelease() CNNSoftMaxGradient {
	rv := objc.Send[CNNSoftMaxGradient](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNSoftMaxGradient creates a new CNNSoftMaxGradient instance.
func NewCNNSoftMaxGradient() CNNSoftMaxGradient {
	return getCNNSoftMaxGradientClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNSoftMaxGradient */
// A gradient softmax filter.


// A gradient softmax filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSoftMaxGradient
type CNNSoftMaxGradient struct {
	CNNGradientKernel
}

// CNNSoftMaxGradientFrom constructs a [CNNSoftMaxGradient] from an unsafe.Pointer.
//
// A gradient softmax filter.
func CNNSoftMaxGradientFrom(ptr unsafe.Pointer) CNNSoftMaxGradient {
	return CNNSoftMaxGradient{
		CNNGradientKernel: CNNGradientKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNSoftMaxGradient */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnsoftmaxgradient/2942618-initwithcoder
func NewCNNSoftMaxGradientWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) CNNSoftMaxGradient {
	instance := getCNNSoftMaxGradientClass().Alloc()
	rv := objc.Send[CNNSoftMaxGradient](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNSoftMaxGradientWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnsoftmaxgradient/2942620-initwithdevice
func NewCNNSoftMaxGradientWithDevice(device unsafe.Pointer) CNNSoftMaxGradient {
	instance := getCNNSoftMaxGradientClass().Alloc()
	rv := objc.Send[CNNSoftMaxGradient](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNSoftMaxGradientWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNSoftMaxGradient */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNSoftMaxGradient */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNSoftMaxGradient */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNSoftMaxGradient */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNSoftMaxGradient */


