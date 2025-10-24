// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MLCRMSPropOptimizer */


/* debug [class_header]: Header for MLCRMSPropOptimizer */
// The class instance for the [CRMSPropOptimizer] class.
var (
	CRMSPropOptimizerClass     _CRMSPropOptimizerClass
	CRMSPropOptimizerClassOnce sync.Once
)

func getCRMSPropOptimizerClass() _CRMSPropOptimizerClass {
	CRMSPropOptimizerClassOnce.Do(func() {
		CRMSPropOptimizerClass = _CRMSPropOptimizerClass{objc.GetClass("MLCRMSPropOptimizer")}
	})
	return CRMSPropOptimizerClass
}

type _CRMSPropOptimizerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CRMSPropOptimizer */
// An interface definition for the [CRMSPropOptimizer] class.
type ICRMSPropOptimizer interface {
	ICOptimizer
	
/* debug [class_interface_properties]: Properties for CRMSPropOptimizer */
	// properties:
	Alpha() float32
	Epsilon() float32
	IsCentered() bool
	MomentumScale() float32
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CRMSPropOptimizer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CRMSPropOptimizer */
// Alloc allocates a new instance without initialization.
func (cc _CRMSPropOptimizerClass) Alloc() CRMSPropOptimizer {
	rv := objc.Send[CRMSPropOptimizer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CRMSPropOptimizerClass) New() CRMSPropOptimizer {
	rv := objc.Send[CRMSPropOptimizer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CRMSPropOptimizer) Init() CRMSPropOptimizer {
	rv := objc.Send[CRMSPropOptimizer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CRMSPropOptimizer) Autorelease() CRMSPropOptimizer {
	rv := objc.Send[CRMSPropOptimizer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCRMSPropOptimizer creates a new CRMSPropOptimizer instance.
func NewCRMSPropOptimizer() CRMSPropOptimizer {
	return getCRMSPropOptimizerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CRMSPropOptimizer */
// An optimizer that represents the root mean square propagation algorithm.


// An optimizer that represents the root mean square propagation algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCRMSPropOptimizer
type CRMSPropOptimizer struct {
	COptimizer
}

// CRMSPropOptimizerFrom constructs a [CRMSPropOptimizer] from an unsafe.Pointer.
//
// An optimizer that represents the root mean square propagation algorithm.
func CRMSPropOptimizerFrom(ptr unsafe.Pointer) CRMSPropOptimizer {
	return CRMSPropOptimizer{
		COptimizer: COptimizerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CRMSPropOptimizer */

// Creates an RMSProp optimizer with the descriptor you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCRMSPropOptimizer/init(descriptor:)
func NewCRMSPropOptimizerWithDescriptor(optimizerDescriptor IMLCOptimizerDescriptor) CRMSPropOptimizer {
	rv := objc.Send[CRMSPropOptimizer](objc.ID(getCRMSPropOptimizerClass().class), objc.Sel("optimizerWithDescriptor:"), optimizerDescriptor)
	return rv
}/* debug [class_init_methods/constructor]: NewCRMSPropOptimizerWithDescriptor */


// Creates an RMSProp optimizer with the descriptor, momentum scale, smoothing, epsilon, and option to compute the centered RMSProp that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCRMSPropOptimizer/init(descriptor:momentumScale:alpha:epsilon:isCentered:)
func NewCRMSPropOptimizerWithDescriptorMomentumScaleAlphaEpsilonIsCentered(optimizerDescriptor IMLCOptimizerDescriptor, momentumScale float32, alpha float32, epsilon float32, isCentered bool) CRMSPropOptimizer {
	rv := objc.Send[CRMSPropOptimizer](objc.ID(getCRMSPropOptimizerClass().class), objc.Sel("optimizerWithDescriptor:momentumScale:alpha:epsilon:isCentered:"), optimizerDescriptor, momentumScale, alpha, epsilon, isCentered)
	return rv
}/* debug [class_init_methods/constructor]: NewCRMSPropOptimizerWithDescriptorMomentumScaleAlphaEpsilonIsCentered */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CRMSPropOptimizer */

// Creates an RMSProp optimizer with the descriptor you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCRMSPropOptimizer/init(descriptor:)
func (cc _CRMSPropOptimizerClass) OptimizerWithDescriptor(optimizerDescriptor IMLCOptimizerDescriptor) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("optimizerWithDescriptor:"), optimizerDescriptor)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=OptimizerWithDescriptor) */


// Creates an RMSProp optimizer with the descriptor, momentum scale, smoothing, epsilon, and option to compute the centered RMSProp that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCRMSPropOptimizer/init(descriptor:momentumScale:alpha:epsilon:isCentered:)
func (cc _CRMSPropOptimizerClass) OptimizerWithDescriptorMomentumScaleAlphaEpsilonIsCentered(optimizerDescriptor IMLCOptimizerDescriptor, momentumScale float32, alpha float32, epsilon float32, isCentered bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("optimizerWithDescriptor:momentumScale:alpha:epsilon:isCentered:"), optimizerDescriptor, momentumScale, alpha, epsilon, isCentered)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=OptimizerWithDescriptorMomentumScaleAlphaEpsilonIsCentered) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CRMSPropOptimizer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CRMSPropOptimizer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CRMSPropOptimizer */

// The constant for smoothing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCRMSPropOptimizer/alpha
func (c_ CRMSPropOptimizer) Alpha() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("alpha"))
	return rv
}/* debug [instance_properties/getter]: alpha */


// The epsilon value you use to improve numerical stability.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCRMSPropOptimizer/epsilon
func (c_ CRMSPropOptimizer) Epsilon() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("epsilon"))
	return rv
}/* debug [instance_properties/getter]: epsilon */


// A Boolean that indicates whether you compute the centered RMSProp.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCRMSPropOptimizer/isCentered
func (c_ CRMSPropOptimizer) IsCentered() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCentered"))
	return rv
}/* debug [instance_properties/getter]: isCentered */


// A hyper-parameter that specifies the momentum factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCRMSPropOptimizer/momentumScale
func (c_ CRMSPropOptimizer) MomentumScale() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("momentumScale"))
	return rv
}/* debug [instance_properties/getter]: momentumScale */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCRMSPropOptimizer */


