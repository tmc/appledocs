// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLCOptimizer */


/* debug [class_header]: Header for MLCOptimizer */
// The class instance for the [COptimizer] class.
var (
	COptimizerClass     _COptimizerClass
	COptimizerClassOnce sync.Once
)

func getCOptimizerClass() _COptimizerClass {
	COptimizerClassOnce.Do(func() {
		COptimizerClass = _COptimizerClass{objc.GetClass("MLCOptimizer")}
	})
	return COptimizerClass
}

type _COptimizerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for COptimizer */
// An interface definition for the [COptimizer] class.
type ICOptimizer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for COptimizer */
	// properties:
	AppliesGradientClipping() bool
	SetAppliesGradientClipping(value bool)
	CustomGlobalNorm() float32
	GradientClipMax() float32
	GradientClipMin() float32
	GradientClippingType() CGradientClippingType
	GradientRescale() float32
	LearningRate() float32
	SetLearningRate(value float32)
	MaximumClippingNorm() float32
	RegularizationScale() float32
	RegularizationType() CRegularizationType
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for COptimizer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for COptimizer */
// Alloc allocates a new instance without initialization.
func (cc _COptimizerClass) Alloc() COptimizer {
	rv := objc.Send[COptimizer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _COptimizerClass) New() COptimizer {
	rv := objc.Send[COptimizer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ COptimizer) Init() COptimizer {
	rv := objc.Send[COptimizer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ COptimizer) Autorelease() COptimizer {
	rv := objc.Send[COptimizer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCOptimizer creates a new COptimizer instance.
func NewCOptimizer() COptimizer {
	return getCOptimizerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for COptimizer */
// The base class for all framework optimizers.


// The base class for all framework optimizers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCOptimizer
type COptimizer struct {
	objectivec.Object
}

// COptimizerFrom constructs a [COptimizer] from an unsafe.Pointer.
//
// The base class for all framework optimizers.
func COptimizerFrom(ptr unsafe.Pointer) COptimizer {
	return COptimizer{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for COptimizer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for COptimizer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for COptimizer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for COptimizer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for COptimizer */

// A Boolean value that indicates whether you apply gradient clipping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCOptimizer/appliesGradientClipping
func (c_ COptimizer) AppliesGradientClipping() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("appliesGradientClipping"))
	return rv
}/* debug [instance_properties/getter]: appliesGradientClipping */


// A Boolean value that indicates whether you apply gradient clipping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCOptimizer/appliesGradientClipping
func (c_ COptimizer) SetAppliesGradientClipping(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAppliesGradientClipping:"), value)
}/* debug [instance_properties/setter]: appliesGradientClipping */


// A custom norm the system uses in place of the global norm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCOptimizer/customGlobalNorm
func (c_ COptimizer) CustomGlobalNorm() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("customGlobalNorm"))
	return rv
}/* debug [instance_properties/getter]: customGlobalNorm */


// The maximum gradient value before the optimizer rescales the gradient, if you enable gradient clipping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCOptimizer/gradientClipMax
func (c_ COptimizer) GradientClipMax() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("gradientClipMax"))
	return rv
}/* debug [instance_properties/getter]: gradientClipMax */


// The minimum gradient value before the optimizer rescales the gradient, if you enable gradient clipping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCOptimizer/gradientClipMin
func (c_ COptimizer) GradientClipMin() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("gradientClipMin"))
	return rv
}/* debug [instance_properties/getter]: gradientClipMin */


// The type of clipping the system applies to the gradient.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCOptimizer/gradientClippingType
func (c_ COptimizer) GradientClippingType() CGradientClippingType {
	rv := objc.Send[CGradientClippingType](c_.ID, objc.Sel("gradientClippingType"))
	return rv
}/* debug [instance_properties/getter]: gradientClippingType */


// The rescale value the optimizer applies to gradients during updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCOptimizer/gradientRescale
func (c_ COptimizer) GradientRescale() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("gradientRescale"))
	return rv
}/* debug [instance_properties/getter]: gradientRescale */


// The learning rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCOptimizer/learningRate
func (c_ COptimizer) LearningRate() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("learningRate"))
	return rv
}/* debug [instance_properties/getter]: learningRate */


// The learning rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCOptimizer/learningRate
func (c_ COptimizer) SetLearningRate(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLearningRate:"), value)
}/* debug [instance_properties/setter]: learningRate */


// The maximum clipping value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCOptimizer/maximumClippingNorm
func (c_ COptimizer) MaximumClippingNorm() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("maximumClippingNorm"))
	return rv
}/* debug [instance_properties/getter]: maximumClippingNorm */


// The regularization scale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCOptimizer/regularizationScale
func (c_ COptimizer) RegularizationScale() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("regularizationScale"))
	return rv
}/* debug [instance_properties/getter]: regularizationScale */


// The regularization type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCOptimizer/regularizationType
func (c_ COptimizer) RegularizationType() CRegularizationType {
	rv := objc.Send[CRegularizationType](c_.ID, objc.Sel("regularizationType"))
	return rv
}/* debug [instance_properties/getter]: regularizationType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCOptimizer */



