// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLCOptimizerDescriptor */


/* debug [class_header]: Header for MLCOptimizerDescriptor */
// The class instance for the [COptimizerDescriptor] class.
var (
	COptimizerDescriptorClass     _COptimizerDescriptorClass
	COptimizerDescriptorClassOnce sync.Once
)

func getCOptimizerDescriptorClass() _COptimizerDescriptorClass {
	COptimizerDescriptorClassOnce.Do(func() {
		COptimizerDescriptorClass = _COptimizerDescriptorClass{objc.GetClass("MLCOptimizerDescriptor")}
	})
	return COptimizerDescriptorClass
}

type _COptimizerDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for COptimizerDescriptor */
// An interface definition for the [COptimizerDescriptor] class.
type ICOptimizerDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for COptimizerDescriptor */
	// properties:
	AppliesGradientClipping() bool
	CustomGlobalNorm() float32
	GradientClipMax() float32
	GradientClipMin() float32
	GradientClippingType() CGradientClippingType
	GradientRescale() float32
	LearningRate() float32
	MaximumClippingNorm() float32
	RegularizationScale() float32
	RegularizationType() CRegularizationType
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for COptimizerDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for COptimizerDescriptor */
// Alloc allocates a new instance without initialization.
func (cc _COptimizerDescriptorClass) Alloc() COptimizerDescriptor {
	rv := objc.Send[COptimizerDescriptor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _COptimizerDescriptorClass) New() COptimizerDescriptor {
	rv := objc.Send[COptimizerDescriptor](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ COptimizerDescriptor) Init() COptimizerDescriptor {
	rv := objc.Send[COptimizerDescriptor](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ COptimizerDescriptor) Autorelease() COptimizerDescriptor {
	rv := objc.Send[COptimizerDescriptor](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCOptimizerDescriptor creates a new COptimizerDescriptor instance.
func NewCOptimizerDescriptor() COptimizerDescriptor {
	return getCOptimizerDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for COptimizerDescriptor */
// A configuration object you use to create an optimizer.


// A configuration object you use to create an optimizer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCOptimizerDescriptor
type COptimizerDescriptor struct {
	objectivec.Object
}

// COptimizerDescriptorFrom constructs a [COptimizerDescriptor] from an unsafe.Pointer.
//
// A configuration object you use to create an optimizer.
func COptimizerDescriptorFrom(ptr unsafe.Pointer) COptimizerDescriptor {
	return COptimizerDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for COptimizerDescriptor */

// Creates a descriptor with the learning rate, gradient rescale, clipping option and values, and regularization type and scale that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCOptimizerDescriptor/init(learningRate:gradientRescale:appliesGradientClipping:gradientClipMax:gradientClipMin:regularizationType:regularizationScale:)
func NewCOptimizerDescriptorWithLearningRateGradientRescaleAppliesGradientClippingGradientClipMaxGradientClipMinRegularizationTypeRegularizationScale(learningRate float32, gradientRescale float32, appliesGradientClipping bool, gradientClipMax float32, gradientClipMin float32, regularizationType CRegularizationType, regularizationScale float32) COptimizerDescriptor {
	rv := objc.Send[COptimizerDescriptor](objc.ID(getCOptimizerDescriptorClass().class), objc.Sel("descriptorWithLearningRate:gradientRescale:appliesGradientClipping:gradientClipMax:gradientClipMin:regularizationType:regularizationScale:"), learningRate, gradientRescale, appliesGradientClipping, gradientClipMax, gradientClipMin, regularizationType, regularizationScale)
	return rv
}/* debug [class_init_methods/constructor]: NewCOptimizerDescriptorWithLearningRateGradientRescaleAppliesGradientClippingGradientClipMaxGradientClipMinRegularizationTypeRegularizationScale */


// Creates a descriptor with the learning rate, gradient rescale, clipping option and values, and regularization type and scale that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCOptimizerDescriptor/init(learningRate:gradientRescale:appliesGradientClipping:gradientClippingType:gradientClipMax:gradientClipMin:maximumClippingNorm:customGlobalNorm:regularizationType:regularizationScale:)
func NewCOptimizerDescriptorWithLearningRateGradientRescaleAppliesGradientClippingGradientClippingTypeGradientClipMaxGradientClipMinMaximumClippingNormCustomGlobalNormRegularizationTypeRegularizationScale(learningRate float32, gradientRescale float32, appliesGradientClipping bool, gradientClippingType CGradientClippingType, gradientClipMax float32, gradientClipMin float32, maximumClippingNorm float32, customGlobalNorm float32, regularizationType CRegularizationType, regularizationScale float32) COptimizerDescriptor {
	rv := objc.Send[COptimizerDescriptor](objc.ID(getCOptimizerDescriptorClass().class), objc.Sel("descriptorWithLearningRate:gradientRescale:appliesGradientClipping:gradientClippingType:gradientClipMax:gradientClipMin:maximumClippingNorm:customGlobalNorm:regularizationType:regularizationScale:"), learningRate, gradientRescale, appliesGradientClipping, gradientClippingType, gradientClipMax, gradientClipMin, maximumClippingNorm, customGlobalNorm, regularizationType, regularizationScale)
	return rv
}/* debug [class_init_methods/constructor]: NewCOptimizerDescriptorWithLearningRateGradientRescaleAppliesGradientClippingGradientClippingTypeGradientClipMaxGradientClipMinMaximumClippingNormCustomGlobalNormRegularizationTypeRegularizationScale */


// Creates an optimizer descriptor with the learning rate, gradient rescale, regularization type, and regulation scale that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCOptimizerDescriptor/init(learningRate:gradientRescale:regularizationType:regularizationScale:)
func NewCOptimizerDescriptorWithLearningRateGradientRescaleRegularizationTypeRegularizationScale(learningRate float32, gradientRescale float32, regularizationType CRegularizationType, regularizationScale float32) COptimizerDescriptor {
	rv := objc.Send[COptimizerDescriptor](objc.ID(getCOptimizerDescriptorClass().class), objc.Sel("descriptorWithLearningRate:gradientRescale:regularizationType:regularizationScale:"), learningRate, gradientRescale, regularizationType, regularizationScale)
	return rv
}/* debug [class_init_methods/constructor]: NewCOptimizerDescriptorWithLearningRateGradientRescaleRegularizationTypeRegularizationScale */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for COptimizerDescriptor */

// Creates a descriptor with the learning rate, gradient rescale, clipping option and values, and regularization type and scale that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCOptimizerDescriptor/init(learningRate:gradientRescale:appliesGradientClipping:gradientClipMax:gradientClipMin:regularizationType:regularizationScale:)
func (cc _COptimizerDescriptorClass) DescriptorWithLearningRateGradientRescaleAppliesGradientClippingGradientClipMaxGradientClipMinRegularizationTypeRegularizationScale(learningRate float32, gradientRescale float32, appliesGradientClipping bool, gradientClipMax float32, gradientClipMin float32, regularizationType CRegularizationType, regularizationScale float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("descriptorWithLearningRate:gradientRescale:appliesGradientClipping:gradientClipMax:gradientClipMin:regularizationType:regularizationScale:"), learningRate, gradientRescale, appliesGradientClipping, gradientClipMax, gradientClipMin, regularizationType, regularizationScale)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithLearningRateGradientRescaleAppliesGradientClippingGradientClipMaxGradientClipMinRegularizationTypeRegularizationScale) */


// Creates a descriptor with the learning rate, gradient rescale, clipping option and values, and regularization type and scale that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCOptimizerDescriptor/init(learningRate:gradientRescale:appliesGradientClipping:gradientClippingType:gradientClipMax:gradientClipMin:maximumClippingNorm:customGlobalNorm:regularizationType:regularizationScale:)
func (cc _COptimizerDescriptorClass) DescriptorWithLearningRateGradientRescaleAppliesGradientClippingGradientClippingTypeGradientClipMaxGradientClipMinMaximumClippingNormCustomGlobalNormRegularizationTypeRegularizationScale(learningRate float32, gradientRescale float32, appliesGradientClipping bool, gradientClippingType CGradientClippingType, gradientClipMax float32, gradientClipMin float32, maximumClippingNorm float32, customGlobalNorm float32, regularizationType CRegularizationType, regularizationScale float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("descriptorWithLearningRate:gradientRescale:appliesGradientClipping:gradientClippingType:gradientClipMax:gradientClipMin:maximumClippingNorm:customGlobalNorm:regularizationType:regularizationScale:"), learningRate, gradientRescale, appliesGradientClipping, gradientClippingType, gradientClipMax, gradientClipMin, maximumClippingNorm, customGlobalNorm, regularizationType, regularizationScale)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithLearningRateGradientRescaleAppliesGradientClippingGradientClippingTypeGradientClipMaxGradientClipMinMaximumClippingNormCustomGlobalNormRegularizationTypeRegularizationScale) */


// Creates an optimizer descriptor with the learning rate, gradient rescale, regularization type, and regulation scale that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCOptimizerDescriptor/init(learningRate:gradientRescale:regularizationType:regularizationScale:)
func (cc _COptimizerDescriptorClass) DescriptorWithLearningRateGradientRescaleRegularizationTypeRegularizationScale(learningRate float32, gradientRescale float32, regularizationType CRegularizationType, regularizationScale float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("descriptorWithLearningRate:gradientRescale:regularizationType:regularizationScale:"), learningRate, gradientRescale, regularizationType, regularizationScale)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithLearningRateGradientRescaleRegularizationTypeRegularizationScale) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for COptimizerDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for COptimizerDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for COptimizerDescriptor */

// A Boolean that indicates whether you apply gradient clipping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCOptimizerDescriptor/appliesGradientClipping
func (c_ COptimizerDescriptor) AppliesGradientClipping() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("appliesGradientClipping"))
	return rv
}/* debug [instance_properties/getter]: appliesGradientClipping */


// A custom norm the system uses in place of the global norm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCOptimizerDescriptor/customGlobalNorm
func (c_ COptimizerDescriptor) CustomGlobalNorm() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("customGlobalNorm"))
	return rv
}/* debug [instance_properties/getter]: customGlobalNorm */


// The maximum gradient value before the optimizer rescales the gradient, if you enabled gradient clipping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCOptimizerDescriptor/gradientClipMax
func (c_ COptimizerDescriptor) GradientClipMax() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("gradientClipMax"))
	return rv
}/* debug [instance_properties/getter]: gradientClipMax */


// The minimum gradient value before the optimizer rescales the gradient, if you enabled gradient clipping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCOptimizerDescriptor/gradientClipMin
func (c_ COptimizerDescriptor) GradientClipMin() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("gradientClipMin"))
	return rv
}/* debug [instance_properties/getter]: gradientClipMin */


// The type of clipping the system applies to the gradient.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCOptimizerDescriptor/gradientClippingType
func (c_ COptimizerDescriptor) GradientClippingType() CGradientClippingType {
	rv := objc.Send[CGradientClippingType](c_.ID, objc.Sel("gradientClippingType"))
	return rv
}/* debug [instance_properties/getter]: gradientClippingType */


// The rescale value the optimizer applies to gradients during updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCOptimizerDescriptor/gradientRescale
func (c_ COptimizerDescriptor) GradientRescale() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("gradientRescale"))
	return rv
}/* debug [instance_properties/getter]: gradientRescale */


// The learning rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCOptimizerDescriptor/learningRate
func (c_ COptimizerDescriptor) LearningRate() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("learningRate"))
	return rv
}/* debug [instance_properties/getter]: learningRate */


// The maximum clipping value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCOptimizerDescriptor/maximumClippingNorm
func (c_ COptimizerDescriptor) MaximumClippingNorm() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("maximumClippingNorm"))
	return rv
}/* debug [instance_properties/getter]: maximumClippingNorm */


// The regularization scale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCOptimizerDescriptor/regularizationScale
func (c_ COptimizerDescriptor) RegularizationScale() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("regularizationScale"))
	return rv
}/* debug [instance_properties/getter]: regularizationScale */


// The regularization type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCOptimizerDescriptor/regularizationType
func (c_ COptimizerDescriptor) RegularizationType() CRegularizationType {
	rv := objc.Send[CRegularizationType](c_.ID, objc.Sel("regularizationType"))
	return rv
}/* debug [instance_properties/getter]: regularizationType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCOptimizerDescriptor */


