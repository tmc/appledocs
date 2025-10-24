// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MLCAdamOptimizer */


/* debug [class_header]: Header for MLCAdamOptimizer */
// The class instance for the [CAdamOptimizer] class.
var (
	CAdamOptimizerClass     _CAdamOptimizerClass
	CAdamOptimizerClassOnce sync.Once
)

func getCAdamOptimizerClass() _CAdamOptimizerClass {
	CAdamOptimizerClassOnce.Do(func() {
		CAdamOptimizerClass = _CAdamOptimizerClass{objc.GetClass("MLCAdamOptimizer")}
	})
	return CAdamOptimizerClass
}

type _CAdamOptimizerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CAdamOptimizer */
// An interface definition for the [CAdamOptimizer] class.
type ICAdamOptimizer interface {
	ICOptimizer
	
/* debug [class_interface_properties]: Properties for CAdamOptimizer */
	// properties:
	Beta1() float32
	Beta2() float32
	Epsilon() float32
	TimeStep() uint
	UsesAMSGrad() bool
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CAdamOptimizer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CAdamOptimizer */
// Alloc allocates a new instance without initialization.
func (cc _CAdamOptimizerClass) Alloc() CAdamOptimizer {
	rv := objc.Send[CAdamOptimizer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CAdamOptimizerClass) New() CAdamOptimizer {
	rv := objc.Send[CAdamOptimizer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CAdamOptimizer) Init() CAdamOptimizer {
	rv := objc.Send[CAdamOptimizer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CAdamOptimizer) Autorelease() CAdamOptimizer {
	rv := objc.Send[CAdamOptimizer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCAdamOptimizer creates a new CAdamOptimizer instance.
func NewCAdamOptimizer() CAdamOptimizer {
	return getCAdamOptimizerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CAdamOptimizer */
// An optimizer that represents the adaptive moment estimation algorithm.


// An optimizer that represents the adaptive moment estimation algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCAdamOptimizer
type CAdamOptimizer struct {
	COptimizer
}

// CAdamOptimizerFrom constructs a [CAdamOptimizer] from an unsafe.Pointer.
//
// An optimizer that represents the adaptive moment estimation algorithm.
func CAdamOptimizerFrom(ptr unsafe.Pointer) CAdamOptimizer {
	return CAdamOptimizer{
		COptimizer: COptimizerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CAdamOptimizer */

// Creates an Adam optimizer with the descriptor you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCAdamOptimizer/init(descriptor:)
func NewCAdamOptimizerWithDescriptor(optimizerDescriptor IMLCOptimizerDescriptor) CAdamOptimizer {
	rv := objc.Send[CAdamOptimizer](objc.ID(getCAdamOptimizerClass().class), objc.Sel("optimizerWithDescriptor:"), optimizerDescriptor)
	return rv
}/* debug [class_init_methods/constructor]: NewCAdamOptimizerWithDescriptor */


// Creates an Adam optimizer with the values you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCAdamOptimizer/init(descriptor:beta1:beta2:epsilon:timeStep:)
func NewCAdamOptimizerWithDescriptorBeta1Beta2EpsilonTimeStep(optimizerDescriptor IMLCOptimizerDescriptor, beta1 float32, beta2 float32, epsilon float32, timeStep uint) CAdamOptimizer {
	rv := objc.Send[CAdamOptimizer](objc.ID(getCAdamOptimizerClass().class), objc.Sel("optimizerWithDescriptor:beta1:beta2:epsilon:timeStep:"), optimizerDescriptor, beta1, beta2, epsilon, timeStep)
	return rv
}/* debug [class_init_methods/constructor]: NewCAdamOptimizerWithDescriptorBeta1Beta2EpsilonTimeStep */


// Creates an Adam optimizer with the values you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCAdamOptimizer/init(descriptor:beta1:beta2:epsilon:usesAMSGrad:timeStep:)
func NewCAdamOptimizerWithDescriptorBeta1Beta2EpsilonUsesAMSGradTimeStep(optimizerDescriptor IMLCOptimizerDescriptor, beta1 float32, beta2 float32, epsilon float32, usesAMSGrad bool, timeStep uint) CAdamOptimizer {
	rv := objc.Send[CAdamOptimizer](objc.ID(getCAdamOptimizerClass().class), objc.Sel("optimizerWithDescriptor:beta1:beta2:epsilon:usesAMSGrad:timeStep:"), optimizerDescriptor, beta1, beta2, epsilon, usesAMSGrad, timeStep)
	return rv
}/* debug [class_init_methods/constructor]: NewCAdamOptimizerWithDescriptorBeta1Beta2EpsilonUsesAMSGradTimeStep */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CAdamOptimizer */

// Creates an Adam optimizer with the descriptor you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCAdamOptimizer/init(descriptor:)
func (cc _CAdamOptimizerClass) OptimizerWithDescriptor(optimizerDescriptor IMLCOptimizerDescriptor) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("optimizerWithDescriptor:"), optimizerDescriptor)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=OptimizerWithDescriptor) */


// Creates an Adam optimizer with the values you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCAdamOptimizer/init(descriptor:beta1:beta2:epsilon:timeStep:)
func (cc _CAdamOptimizerClass) OptimizerWithDescriptorBeta1Beta2EpsilonTimeStep(optimizerDescriptor IMLCOptimizerDescriptor, beta1 float32, beta2 float32, epsilon float32, timeStep uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("optimizerWithDescriptor:beta1:beta2:epsilon:timeStep:"), optimizerDescriptor, beta1, beta2, epsilon, timeStep)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=OptimizerWithDescriptorBeta1Beta2EpsilonTimeStep) */


// Creates an Adam optimizer with the values you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCAdamOptimizer/init(descriptor:beta1:beta2:epsilon:usesAMSGrad:timeStep:)
func (cc _CAdamOptimizerClass) OptimizerWithDescriptorBeta1Beta2EpsilonUsesAMSGradTimeStep(optimizerDescriptor IMLCOptimizerDescriptor, beta1 float32, beta2 float32, epsilon float32, usesAMSGrad bool, timeStep uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("optimizerWithDescriptor:beta1:beta2:epsilon:usesAMSGrad:timeStep:"), optimizerDescriptor, beta1, beta2, epsilon, usesAMSGrad, timeStep)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=OptimizerWithDescriptorBeta1Beta2EpsilonUsesAMSGradTimeStep) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CAdamOptimizer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CAdamOptimizer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CAdamOptimizer */

// The coefficent for computing running averages of gradient.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCAdamOptimizer/beta1
func (c_ CAdamOptimizer) Beta1() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("beta1"))
	return rv
}/* debug [instance_properties/getter]: beta1 */


// The coefficent for computing running averages of square of gradient.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCAdamOptimizer/beta2
func (c_ CAdamOptimizer) Beta2() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("beta2"))
	return rv
}/* debug [instance_properties/getter]: beta2 */


// The epsilon value for improving numerical stability.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCAdamOptimizer/epsilon
func (c_ CAdamOptimizer) Epsilon() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("epsilon"))
	return rv
}/* debug [instance_properties/getter]: epsilon */


// The initial timestep for the update.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCAdamOptimizer/timeStep
func (c_ CAdamOptimizer) TimeStep() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("timeStep"))
	return rv
}/* debug [instance_properties/getter]: timeStep */


// A Boolean value that indicates whether to use a variant of the algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCAdamOptimizer/usesAMSGrad
func (c_ CAdamOptimizer) UsesAMSGrad() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("usesAMSGrad"))
	return rv
}/* debug [instance_properties/getter]: usesAMSGrad */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCAdamOptimizer */


