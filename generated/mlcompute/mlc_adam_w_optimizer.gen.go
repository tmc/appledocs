// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MLCAdamWOptimizer */


/* debug [class_header]: Header for MLCAdamWOptimizer */
// The class instance for the [CAdamWOptimizer] class.
var (
	CAdamWOptimizerClass     _CAdamWOptimizerClass
	CAdamWOptimizerClassOnce sync.Once
)

func getCAdamWOptimizerClass() _CAdamWOptimizerClass {
	CAdamWOptimizerClassOnce.Do(func() {
		CAdamWOptimizerClass = _CAdamWOptimizerClass{objc.GetClass("MLCAdamWOptimizer")}
	})
	return CAdamWOptimizerClass
}

type _CAdamWOptimizerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CAdamWOptimizer */
// An interface definition for the [CAdamWOptimizer] class.
type ICAdamWOptimizer interface {
	ICOptimizer
	
/* debug [class_interface_properties]: Properties for CAdamWOptimizer */
	// properties:
	Beta1() float32
	Beta2() float32
	Epsilon() float32
	TimeStep() uint
	UsesAMSGrad() bool
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CAdamWOptimizer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CAdamWOptimizer */
// Alloc allocates a new instance without initialization.
func (cc _CAdamWOptimizerClass) Alloc() CAdamWOptimizer {
	rv := objc.Send[CAdamWOptimizer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CAdamWOptimizerClass) New() CAdamWOptimizer {
	rv := objc.Send[CAdamWOptimizer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CAdamWOptimizer) Init() CAdamWOptimizer {
	rv := objc.Send[CAdamWOptimizer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CAdamWOptimizer) Autorelease() CAdamWOptimizer {
	rv := objc.Send[CAdamWOptimizer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCAdamWOptimizer creates a new CAdamWOptimizer instance.
func NewCAdamWOptimizer() CAdamWOptimizer {
	return getCAdamWOptimizerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CAdamWOptimizer */
// An optimizer that represents the Adam algorithm with weight decay.


// An optimizer that represents the Adam algorithm with weight decay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCAdamWOptimizer
type CAdamWOptimizer struct {
	COptimizer
}

// CAdamWOptimizerFrom constructs a [CAdamWOptimizer] from an unsafe.Pointer.
//
// An optimizer that represents the Adam algorithm with weight decay.
func CAdamWOptimizerFrom(ptr unsafe.Pointer) CAdamWOptimizer {
	return CAdamWOptimizer{
		COptimizer: COptimizerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CAdamWOptimizer */

// Creates a default optimizer with the descriptor you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCAdamWOptimizer/init(descriptor:)
func NewCAdamWOptimizerWithDescriptor(optimizerDescriptor IMLCOptimizerDescriptor) CAdamWOptimizer {
	rv := objc.Send[CAdamWOptimizer](objc.ID(getCAdamWOptimizerClass().class), objc.Sel("optimizerWithDescriptor:"), optimizerDescriptor)
	return rv
}/* debug [class_init_methods/constructor]: NewCAdamWOptimizerWithDescriptor */


// Creates an AdamW optimizer with the values you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCAdamWOptimizer/init(descriptor:beta1:beta2:epsilon:usesAMSGrad:timeStep:)
func NewCAdamWOptimizerWithDescriptorBeta1Beta2EpsilonUsesAMSGradTimeStep(optimizerDescriptor IMLCOptimizerDescriptor, beta1 float32, beta2 float32, epsilon float32, usesAMSGrad bool, timeStep uint) CAdamWOptimizer {
	rv := objc.Send[CAdamWOptimizer](objc.ID(getCAdamWOptimizerClass().class), objc.Sel("optimizerWithDescriptor:beta1:beta2:epsilon:usesAMSGrad:timeStep:"), optimizerDescriptor, beta1, beta2, epsilon, usesAMSGrad, timeStep)
	return rv
}/* debug [class_init_methods/constructor]: NewCAdamWOptimizerWithDescriptorBeta1Beta2EpsilonUsesAMSGradTimeStep */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CAdamWOptimizer */

// Creates a default optimizer with the descriptor you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCAdamWOptimizer/init(descriptor:)
func (cc _CAdamWOptimizerClass) OptimizerWithDescriptor(optimizerDescriptor IMLCOptimizerDescriptor) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("optimizerWithDescriptor:"), optimizerDescriptor)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=OptimizerWithDescriptor) */


// Creates an AdamW optimizer with the values you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCAdamWOptimizer/init(descriptor:beta1:beta2:epsilon:usesAMSGrad:timeStep:)
func (cc _CAdamWOptimizerClass) OptimizerWithDescriptorBeta1Beta2EpsilonUsesAMSGradTimeStep(optimizerDescriptor IMLCOptimizerDescriptor, beta1 float32, beta2 float32, epsilon float32, usesAMSGrad bool, timeStep uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("optimizerWithDescriptor:beta1:beta2:epsilon:usesAMSGrad:timeStep:"), optimizerDescriptor, beta1, beta2, epsilon, usesAMSGrad, timeStep)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=OptimizerWithDescriptorBeta1Beta2EpsilonUsesAMSGradTimeStep) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CAdamWOptimizer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CAdamWOptimizer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CAdamWOptimizer */

// The coefficent for computing running averages of gradient.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCAdamWOptimizer/beta1
func (c_ CAdamWOptimizer) Beta1() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("beta1"))
	return rv
}/* debug [instance_properties/getter]: beta1 */


// The coefficent for computing running averages of square of gradient.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCAdamWOptimizer/beta2
func (c_ CAdamWOptimizer) Beta2() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("beta2"))
	return rv
}/* debug [instance_properties/getter]: beta2 */


// The epsilon value for improving numerical stability.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCAdamWOptimizer/epsilon
func (c_ CAdamWOptimizer) Epsilon() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("epsilon"))
	return rv
}/* debug [instance_properties/getter]: epsilon */


// The initial timestep for the update.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCAdamWOptimizer/timeStep
func (c_ CAdamWOptimizer) TimeStep() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("timeStep"))
	return rv
}/* debug [instance_properties/getter]: timeStep */


// A Boolean value that indicates whether to use a variant of the algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCAdamWOptimizer/usesAMSGrad
func (c_ CAdamWOptimizer) UsesAMSGrad() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("usesAMSGrad"))
	return rv
}/* debug [instance_properties/getter]: usesAMSGrad */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCAdamWOptimizer */


