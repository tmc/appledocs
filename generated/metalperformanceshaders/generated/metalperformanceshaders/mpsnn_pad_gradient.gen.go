// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNNPadGradient */


/* debug [class_header]: Header for MPSNNPadGradient */
// The class instance for the [PadGradient] class.
var (
	PadGradientClass     _PadGradientClass
	PadGradientClassOnce sync.Once
)

func getPadGradientClass() _PadGradientClass {
	PadGradientClassOnce.Do(func() {
		PadGradientClass = _PadGradientClass{objc.GetClass("MPSNNPadGradient")}
	})
	return PadGradientClass
}

type _PadGradientClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PadGradient */
// An interface definition for the [PadGradient] class.
type IPadGradient interface {
	ICNNGradientKernel
	
/* debug [class_interface_properties]: Properties for PadGradient */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PadGradient */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PadGradient */
// Alloc allocates a new instance without initialization.
func (pc _PadGradientClass) Alloc() PadGradient {
	rv := objc.Send[PadGradient](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PadGradientClass) New() PadGradient {
	rv := objc.Send[PadGradient](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PadGradient) Init() PadGradient {
	rv := objc.Send[PadGradient](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PadGradient) Autorelease() PadGradient {
	rv := objc.Send[PadGradient](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPadGradient creates a new PadGradient instance.
func NewPadGradient() PadGradient {
	return getPadGradientClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PadGradient */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPadGradient
type PadGradient struct {
	CNNGradientKernel
}

// PadGradientFrom constructs a [PadGradient] from an unsafe.Pointer.
func PadGradientFrom(ptr unsafe.Pointer) PadGradient {
	return PadGradient{
		CNNGradientKernel: CNNGradientKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PadGradient */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpadgradient/3037435-initwithcoder
func NewPadGradientWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) PadGradient {
	instance := getPadGradientClass().Alloc()
	rv := objc.Send[PadGradient](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPadGradientWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpadgradient/3037436-initwithdevice
func NewPadGradientWithDevice(device unsafe.Pointer) PadGradient {
	instance := getPadGradientClass().Alloc()
	rv := objc.Send[PadGradient](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPadGradientWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PadGradient */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PadGradient */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PadGradient */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PadGradient */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNPadGradient */


