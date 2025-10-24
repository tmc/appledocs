// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSCNNInstanceNormalizationGradient */


/* debug [class_header]: Header for MPSCNNInstanceNormalizationGradient */
// The class instance for the [CNNInstanceNormalizationGradient] class.
var (
	CNNInstanceNormalizationGradientClass     _CNNInstanceNormalizationGradientClass
	CNNInstanceNormalizationGradientClassOnce sync.Once
)

func getCNNInstanceNormalizationGradientClass() _CNNInstanceNormalizationGradientClass {
	CNNInstanceNormalizationGradientClassOnce.Do(func() {
		CNNInstanceNormalizationGradientClass = _CNNInstanceNormalizationGradientClass{objc.GetClass("MPSCNNInstanceNormalizationGradient")}
	})
	return CNNInstanceNormalizationGradientClass
}

type _CNNInstanceNormalizationGradientClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNInstanceNormalizationGradient */
// An interface definition for the [CNNInstanceNormalizationGradient] class.
type ICNNInstanceNormalizationGradient interface {
	ICNNGradientKernel
	
/* debug [class_interface_properties]: Properties for CNNInstanceNormalizationGradient */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNInstanceNormalizationGradient */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNInstanceNormalizationGradient */
// Alloc allocates a new instance without initialization.
func (cc _CNNInstanceNormalizationGradientClass) Alloc() CNNInstanceNormalizationGradient {
	rv := objc.Send[CNNInstanceNormalizationGradient](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNInstanceNormalizationGradientClass) New() CNNInstanceNormalizationGradient {
	rv := objc.Send[CNNInstanceNormalizationGradient](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNInstanceNormalizationGradient) Init() CNNInstanceNormalizationGradient {
	rv := objc.Send[CNNInstanceNormalizationGradient](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNInstanceNormalizationGradient) Autorelease() CNNInstanceNormalizationGradient {
	rv := objc.Send[CNNInstanceNormalizationGradient](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNInstanceNormalizationGradient creates a new CNNInstanceNormalizationGradient instance.
func NewCNNInstanceNormalizationGradient() CNNInstanceNormalizationGradient {
	return getCNNInstanceNormalizationGradientClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNInstanceNormalizationGradient */
// A gradient instance normalization kernel.


// A gradient instance normalization kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalizationGradient
type CNNInstanceNormalizationGradient struct {
	CNNGradientKernel
}

// CNNInstanceNormalizationGradientFrom constructs a [CNNInstanceNormalizationGradient] from an unsafe.Pointer.
//
// A gradient instance normalization kernel.
func CNNInstanceNormalizationGradientFrom(ptr unsafe.Pointer) CNNInstanceNormalizationGradient {
	return CNNInstanceNormalizationGradient{
		CNNGradientKernel: CNNGradientKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNInstanceNormalizationGradient *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNInstanceNormalizationGradient */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNInstanceNormalizationGradient */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNInstanceNormalizationGradient */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNInstanceNormalizationGradient */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNInstanceNormalizationGradient */



