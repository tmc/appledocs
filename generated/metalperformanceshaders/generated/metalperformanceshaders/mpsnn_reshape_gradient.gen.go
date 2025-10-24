// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNNReshapeGradient */


/* debug [class_header]: Header for MPSNNReshapeGradient */
// The class instance for the [ReshapeGradient] class.
var (
	ReshapeGradientClass     _ReshapeGradientClass
	ReshapeGradientClassOnce sync.Once
)

func getReshapeGradientClass() _ReshapeGradientClass {
	ReshapeGradientClassOnce.Do(func() {
		ReshapeGradientClass = _ReshapeGradientClass{objc.GetClass("MPSNNReshapeGradient")}
	})
	return ReshapeGradientClass
}

type _ReshapeGradientClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ReshapeGradient */
// An interface definition for the [ReshapeGradient] class.
type IReshapeGradient interface {
	ICNNGradientKernel
	
/* debug [class_interface_properties]: Properties for ReshapeGradient */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ReshapeGradient */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ReshapeGradient */
// Alloc allocates a new instance without initialization.
func (rc _ReshapeGradientClass) Alloc() ReshapeGradient {
	rv := objc.Send[ReshapeGradient](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReshapeGradientClass) New() ReshapeGradient {
	rv := objc.Send[ReshapeGradient](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReshapeGradient) Init() ReshapeGradient {
	rv := objc.Send[ReshapeGradient](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReshapeGradient) Autorelease() ReshapeGradient {
	rv := objc.Send[ReshapeGradient](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReshapeGradient creates a new ReshapeGradient instance.
func NewReshapeGradient() ReshapeGradient {
	return getReshapeGradientClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ReshapeGradient */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReshapeGradient
type ReshapeGradient struct {
	CNNGradientKernel
}

// ReshapeGradientFrom constructs a [ReshapeGradient] from an unsafe.Pointer.
func ReshapeGradientFrom(ptr unsafe.Pointer) ReshapeGradient {
	return ReshapeGradient{
		CNNGradientKernel: CNNGradientKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ReshapeGradient */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreshapegradient/3037438-initwithcoder
func NewReshapeGradientWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) ReshapeGradient {
	instance := getReshapeGradientClass().Alloc()
	rv := objc.Send[ReshapeGradient](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewReshapeGradientWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreshapegradient/3037439-initwithdevice
func NewReshapeGradientWithDevice(device unsafe.Pointer) ReshapeGradient {
	instance := getReshapeGradientClass().Alloc()
	rv := objc.Send[ReshapeGradient](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewReshapeGradientWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ReshapeGradient */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ReshapeGradient */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ReshapeGradient */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ReshapeGradient */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNReshapeGradient */


