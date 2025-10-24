// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNNInitialGradient */


/* debug [class_header]: Header for MPSNNInitialGradient */
// The class instance for the [InitialGradient] class.
var (
	InitialGradientClass     _InitialGradientClass
	InitialGradientClassOnce sync.Once
)

func getInitialGradientClass() _InitialGradientClass {
	InitialGradientClassOnce.Do(func() {
		InitialGradientClass = _InitialGradientClass{objc.GetClass("MPSNNInitialGradient")}
	})
	return InitialGradientClass
}

type _InitialGradientClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for InitialGradient */
// An interface definition for the [InitialGradient] class.
type IInitialGradient interface {
	ICNNKernel
	
/* debug [class_interface_properties]: Properties for InitialGradient */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for InitialGradient */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for InitialGradient */
// Alloc allocates a new instance without initialization.
func (ic _InitialGradientClass) Alloc() InitialGradient {
	rv := objc.Send[InitialGradient](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _InitialGradientClass) New() InitialGradient {
	rv := objc.Send[InitialGradient](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ InitialGradient) Init() InitialGradient {
	rv := objc.Send[InitialGradient](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ InitialGradient) Autorelease() InitialGradient {
	rv := objc.Send[InitialGradient](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewInitialGradient creates a new InitialGradient instance.
func NewInitialGradient() InitialGradient {
	return getInitialGradientClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for InitialGradient */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNInitialGradient
type InitialGradient struct {
	CNNKernel
}

// InitialGradientFrom constructs a [InitialGradient] from an unsafe.Pointer.
func InitialGradientFrom(ptr unsafe.Pointer) InitialGradient {
	return InitialGradient{
		CNNKernel: CNNKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for InitialGradient */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnninitialgradient/3131809-initwithdevice
func NewInitialGradientWithDevice(device unsafe.Pointer) InitialGradient {
	instance := getInitialGradientClass().Alloc()
	rv := objc.Send[InitialGradient](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewInitialGradientWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for InitialGradient */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for InitialGradient */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for InitialGradient */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for InitialGradient */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNInitialGradient */


