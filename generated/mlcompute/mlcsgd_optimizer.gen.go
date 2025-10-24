// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MLCSGDOptimizer */


/* debug [class_header]: Header for MLCSGDOptimizer */
// The class instance for the [CSGDOptimizer] class.
var (
	CSGDOptimizerClass     _CSGDOptimizerClass
	CSGDOptimizerClassOnce sync.Once
)

func getCSGDOptimizerClass() _CSGDOptimizerClass {
	CSGDOptimizerClassOnce.Do(func() {
		CSGDOptimizerClass = _CSGDOptimizerClass{objc.GetClass("MLCSGDOptimizer")}
	})
	return CSGDOptimizerClass
}

type _CSGDOptimizerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CSGDOptimizer */
// An interface definition for the [CSGDOptimizer] class.
type ICSGDOptimizer interface {
	ICOptimizer
	
/* debug [class_interface_properties]: Properties for CSGDOptimizer */
	// properties:
	MomentumScale() float32
	UsesNesterovMomentum() bool
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CSGDOptimizer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CSGDOptimizer */
// Alloc allocates a new instance without initialization.
func (cc _CSGDOptimizerClass) Alloc() CSGDOptimizer {
	rv := objc.Send[CSGDOptimizer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CSGDOptimizerClass) New() CSGDOptimizer {
	rv := objc.Send[CSGDOptimizer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CSGDOptimizer) Init() CSGDOptimizer {
	rv := objc.Send[CSGDOptimizer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CSGDOptimizer) Autorelease() CSGDOptimizer {
	rv := objc.Send[CSGDOptimizer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCSGDOptimizer creates a new CSGDOptimizer instance.
func NewCSGDOptimizer() CSGDOptimizer {
	return getCSGDOptimizerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CSGDOptimizer */
// An optimizer that represents the stochastic gradient decent algorithm.


// An optimizer that represents the stochastic gradient decent algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCSGDOptimizer
type CSGDOptimizer struct {
	COptimizer
}

// CSGDOptimizerFrom constructs a [CSGDOptimizer] from an unsafe.Pointer.
//
// An optimizer that represents the stochastic gradient decent algorithm.
func CSGDOptimizerFrom(ptr unsafe.Pointer) CSGDOptimizer {
	return CSGDOptimizer{
		COptimizer: COptimizerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CSGDOptimizer */

// Creates an SGD optimizer with the descriptor you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCSGDOptimizer/init(descriptor:)
func NewCSGDOptimizerWithDescriptor(optimizerDescriptor IMLCOptimizerDescriptor) CSGDOptimizer {
	rv := objc.Send[CSGDOptimizer](objc.ID(getCSGDOptimizerClass().class), objc.Sel("optimizerWithDescriptor:"), optimizerDescriptor)
	return rv
}/* debug [class_init_methods/constructor]: NewCSGDOptimizerWithDescriptor */


// Create an SGD optimizer with the descriptor, momentum scale, and option to enable Nesterov momentum that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCSGDOptimizer/init(descriptor:momentumScale:usesNesterovMomentum:)
func NewCSGDOptimizerWithDescriptorMomentumScaleUsesNesterovMomentum(optimizerDescriptor IMLCOptimizerDescriptor, momentumScale float32, usesNesterovMomentum bool) CSGDOptimizer {
	rv := objc.Send[CSGDOptimizer](objc.ID(getCSGDOptimizerClass().class), objc.Sel("optimizerWithDescriptor:momentumScale:usesNesterovMomentum:"), optimizerDescriptor, momentumScale, usesNesterovMomentum)
	return rv
}/* debug [class_init_methods/constructor]: NewCSGDOptimizerWithDescriptorMomentumScaleUsesNesterovMomentum */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CSGDOptimizer */

// Creates an SGD optimizer with the descriptor you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCSGDOptimizer/init(descriptor:)
func (cc _CSGDOptimizerClass) OptimizerWithDescriptor(optimizerDescriptor IMLCOptimizerDescriptor) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("optimizerWithDescriptor:"), optimizerDescriptor)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=OptimizerWithDescriptor) */


// Create an SGD optimizer with the descriptor, momentum scale, and option to enable Nesterov momentum that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCSGDOptimizer/init(descriptor:momentumScale:usesNesterovMomentum:)
func (cc _CSGDOptimizerClass) OptimizerWithDescriptorMomentumScaleUsesNesterovMomentum(optimizerDescriptor IMLCOptimizerDescriptor, momentumScale float32, usesNesterovMomentum bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("optimizerWithDescriptor:momentumScale:usesNesterovMomentum:"), optimizerDescriptor, momentumScale, usesNesterovMomentum)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=OptimizerWithDescriptorMomentumScaleUsesNesterovMomentum) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CSGDOptimizer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CSGDOptimizer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CSGDOptimizer */

// A hyper-parameter specifying the momentum factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCSGDOptimizer/momentumScale
func (c_ CSGDOptimizer) MomentumScale() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("momentumScale"))
	return rv
}/* debug [instance_properties/getter]: momentumScale */


// A Boolean that indicates whether you enable Nesterov momentum.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCSGDOptimizer/usesNesterovMomentum
func (c_ CSGDOptimizer) UsesNesterovMomentum() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("usesNesterovMomentum"))
	return rv
}/* debug [instance_properties/getter]: usesNesterovMomentum */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCSGDOptimizer */


