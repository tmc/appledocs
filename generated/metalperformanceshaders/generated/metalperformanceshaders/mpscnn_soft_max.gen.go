// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSCNNSoftMax */


/* debug [class_header]: Header for MPSCNNSoftMax */
// The class instance for the [CNNSoftMax] class.
var (
	CNNSoftMaxClass     _CNNSoftMaxClass
	CNNSoftMaxClassOnce sync.Once
)

func getCNNSoftMaxClass() _CNNSoftMaxClass {
	CNNSoftMaxClassOnce.Do(func() {
		CNNSoftMaxClass = _CNNSoftMaxClass{objc.GetClass("MPSCNNSoftMax")}
	})
	return CNNSoftMaxClass
}

type _CNNSoftMaxClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNSoftMax */
// An interface definition for the [CNNSoftMax] class.
type ICNNSoftMax interface {
	ICNNKernel
	
/* debug [class_interface_properties]: Properties for CNNSoftMax */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNSoftMax */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNSoftMax */
// Alloc allocates a new instance without initialization.
func (cc _CNNSoftMaxClass) Alloc() CNNSoftMax {
	rv := objc.Send[CNNSoftMax](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNSoftMaxClass) New() CNNSoftMax {
	rv := objc.Send[CNNSoftMax](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNSoftMax) Init() CNNSoftMax {
	rv := objc.Send[CNNSoftMax](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNSoftMax) Autorelease() CNNSoftMax {
	rv := objc.Send[CNNSoftMax](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNSoftMax creates a new CNNSoftMax instance.
func NewCNNSoftMax() CNNSoftMax {
	return getCNNSoftMaxClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNSoftMax */
// A neural transfer function that is useful for classification tasks.
//
// The softmax filter is applied across feature channels in a convolutional manner at all spatial locations. The softmax filter can be seen as the combination of an activation function (exponential) and a normalization operator. For each feature channel per pixel in an image in a feature map, the softmax filter computes the following: Where is the result channel in the pixel and is the number of feature channels.


// A neural transfer function that is useful for classification tasks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSoftMax
type CNNSoftMax struct {
	CNNKernel
}

// CNNSoftMaxFrom constructs a [CNNSoftMax] from an unsafe.Pointer.
//
// A neural transfer function that is useful for classification tasks.
func CNNSoftMaxFrom(ptr unsafe.Pointer) CNNSoftMax {
	return CNNSoftMax{
		CNNKernel: CNNKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNSoftMax *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNSoftMax */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNSoftMax */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNSoftMax */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNSoftMax */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNSoftMax */



