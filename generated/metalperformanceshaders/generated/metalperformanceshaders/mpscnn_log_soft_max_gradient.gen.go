// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNLogSoftMaxGradient */


/* debug [class_header]: Header for MPSCNNLogSoftMaxGradient */
// The class instance for the [CNNLogSoftMaxGradient] class.
var (
	CNNLogSoftMaxGradientClass     _CNNLogSoftMaxGradientClass
	CNNLogSoftMaxGradientClassOnce sync.Once
)

func getCNNLogSoftMaxGradientClass() _CNNLogSoftMaxGradientClass {
	CNNLogSoftMaxGradientClassOnce.Do(func() {
		CNNLogSoftMaxGradientClass = _CNNLogSoftMaxGradientClass{objc.GetClass("MPSCNNLogSoftMaxGradient")}
	})
	return CNNLogSoftMaxGradientClass
}

type _CNNLogSoftMaxGradientClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNLogSoftMaxGradient */
// An interface definition for the [CNNLogSoftMaxGradient] class.
type ICNNLogSoftMaxGradient interface {
	ICNNGradientKernel
	
/* debug [class_interface_properties]: Properties for CNNLogSoftMaxGradient */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNLogSoftMaxGradient */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNLogSoftMaxGradient */
// Alloc allocates a new instance without initialization.
func (cc _CNNLogSoftMaxGradientClass) Alloc() CNNLogSoftMaxGradient {
	rv := objc.Send[CNNLogSoftMaxGradient](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNLogSoftMaxGradientClass) New() CNNLogSoftMaxGradient {
	rv := objc.Send[CNNLogSoftMaxGradient](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNLogSoftMaxGradient) Init() CNNLogSoftMaxGradient {
	rv := objc.Send[CNNLogSoftMaxGradient](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNLogSoftMaxGradient) Autorelease() CNNLogSoftMaxGradient {
	rv := objc.Send[CNNLogSoftMaxGradient](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNLogSoftMaxGradient creates a new CNNLogSoftMaxGradient instance.
func NewCNNLogSoftMaxGradient() CNNLogSoftMaxGradient {
	return getCNNLogSoftMaxGradientClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNLogSoftMaxGradient */
// A gradient logarithmic softmax filter.


// A gradient logarithmic softmax filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLogSoftMaxGradient
type CNNLogSoftMaxGradient struct {
	CNNGradientKernel
}

// CNNLogSoftMaxGradientFrom constructs a [CNNLogSoftMaxGradient] from an unsafe.Pointer.
//
// A gradient logarithmic softmax filter.
func CNNLogSoftMaxGradientFrom(ptr unsafe.Pointer) CNNLogSoftMaxGradient {
	return CNNLogSoftMaxGradient{
		CNNGradientKernel: CNNGradientKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNLogSoftMaxGradient */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlogsoftmaxgradient/2942619-initwithcoder
func NewCNNLogSoftMaxGradientWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) CNNLogSoftMaxGradient {
	instance := getCNNLogSoftMaxGradientClass().Alloc()
	rv := objc.Send[CNNLogSoftMaxGradient](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNLogSoftMaxGradientWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlogsoftmaxgradient/2942622-initwithdevice
func NewCNNLogSoftMaxGradientWithDevice(device unsafe.Pointer) CNNLogSoftMaxGradient {
	instance := getCNNLogSoftMaxGradientClass().Alloc()
	rv := objc.Send[CNNLogSoftMaxGradient](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNLogSoftMaxGradientWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNLogSoftMaxGradient */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNLogSoftMaxGradient */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNLogSoftMaxGradient */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNLogSoftMaxGradient */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNLogSoftMaxGradient */


