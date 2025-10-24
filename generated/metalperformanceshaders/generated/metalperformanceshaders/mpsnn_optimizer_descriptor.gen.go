// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNNOptimizerDescriptor */


/* debug [class_header]: Header for MPSNNOptimizerDescriptor */
// The class instance for the [OptimizerDescriptor] class.
var (
	OptimizerDescriptorClass     _OptimizerDescriptorClass
	OptimizerDescriptorClassOnce sync.Once
)

func getOptimizerDescriptorClass() _OptimizerDescriptorClass {
	OptimizerDescriptorClassOnce.Do(func() {
		OptimizerDescriptorClass = _OptimizerDescriptorClass{objc.GetClass("MPSNNOptimizerDescriptor")}
	})
	return OptimizerDescriptorClass
}

type _OptimizerDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for OptimizerDescriptor */
// An interface definition for the [OptimizerDescriptor] class.
type IOptimizerDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for OptimizerDescriptor */
	// properties:
	ApplyGradientClipping() objectivec.IObject
	SetApplyGradientClipping(value objectivec.IObject)
	GradientClipMax() objectivec.IObject
	SetGradientClipMax(value objectivec.IObject)
	GradientClipMin() objectivec.IObject
	SetGradientClipMin(value objectivec.IObject)
	GradientRescale() objectivec.IObject
	SetGradientRescale(value objectivec.IObject)
	LearningRate() objectivec.IObject
	SetLearningRate(value objectivec.IObject)
	RegularizationScale() objectivec.IObject
	SetRegularizationScale(value objectivec.IObject)
	RegularizationType() RegularizationType get set /* not a class type */
	SetRegularizationType(value RegularizationType get set /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for OptimizerDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for OptimizerDescriptor */
// Alloc allocates a new instance without initialization.
func (oc _OptimizerDescriptorClass) Alloc() OptimizerDescriptor {
	rv := objc.Send[OptimizerDescriptor](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (oc _OptimizerDescriptorClass) New() OptimizerDescriptor {
	rv := objc.Send[OptimizerDescriptor](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OptimizerDescriptor) Init() OptimizerDescriptor {
	rv := objc.Send[OptimizerDescriptor](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OptimizerDescriptor) Autorelease() OptimizerDescriptor {
	rv := objc.Send[OptimizerDescriptor](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOptimizerDescriptor creates a new OptimizerDescriptor instance.
func NewOptimizerDescriptor() OptimizerDescriptor {
	return getOptimizerDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for OptimizerDescriptor */
// An object that specifies properties used by an optimizer kernel.


// An object that specifies properties used by an optimizer kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizerDescriptor
type OptimizerDescriptor struct {
	objectivec.Object
}

// OptimizerDescriptorFrom constructs a [OptimizerDescriptor] from an unsafe.Pointer.
//
// An object that specifies properties used by an optimizer kernel.
func OptimizerDescriptorFrom(ptr unsafe.Pointer) OptimizerDescriptor {
	return OptimizerDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for OptimizerDescriptor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerdescriptor/2966726-initwithlearningrate
func NewOptimizerDescriptorWithLearningRateGradientRescaleApplyGradientClippingGradientClipMaxGradientClipMinRegularizationTypeRegularizationScale(learningRate float32, gradientRescale float32, applyGradientClipping bool, gradientClipMax float32, gradientClipMin float32, regularizationType RegularizationType, regularizationScale float32) OptimizerDescriptor {
	instance := getOptimizerDescriptorClass().Alloc()
	rv := objc.Send[OptimizerDescriptor](instance.ID, objc.Sel("initWithLearningRate:gradientRescale:applyGradientClipping:gradientClipMax:gradientClipMin:regularizationType:regularizationScale:"), learningRate, gradientRescale, applyGradientClipping, gradientClipMax, gradientClipMin, regularizationType, regularizationScale)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewOptimizerDescriptorWithLearningRateGradientRescaleApplyGradientClippingGradientClipMaxGradientClipMinRegularizationTypeRegularizationScale */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerdescriptor/2966727-initwithlearningrate
func NewOptimizerDescriptorWithLearningRateGradientRescaleRegularizationTypeRegularizationScale(learningRate float32, gradientRescale float32, regularizationType RegularizationType, regularizationScale float32) OptimizerDescriptor {
	instance := getOptimizerDescriptorClass().Alloc()
	rv := objc.Send[OptimizerDescriptor](instance.ID, objc.Sel("initWithLearningRate:gradientRescale:regularizationType:regularizationScale:"), learningRate, gradientRescale, regularizationType, regularizationScale)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewOptimizerDescriptorWithLearningRateGradientRescaleRegularizationTypeRegularizationScale */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for OptimizerDescriptor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerdescriptor/2966729-optimizerdescriptorwithlearningr
func (oc _OptimizerDescriptorClass) OptimizerDescriptorWithLearningRateGradientRescaleApplyGradientClippingGradientClipMaxGradientClipMinRegularizationTypeRegularizationScale(learningRate float32, gradientRescale float32, applyGradientClipping bool, gradientClipMax float32, gradientClipMin float32, regularizationType RegularizationType, regularizationScale float32) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(oc.class), objc.Sel("optimizerDescriptorWithLearningRate:gradientRescale:applyGradientClipping:gradientClipMax:gradientClipMin:regularizationType:regularizationScale:"), learningRate, gradientRescale, applyGradientClipping, gradientClipMax, gradientClipMin, regularizationType, regularizationScale)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=OptimizerDescriptorWithLearningRateGradientRescaleApplyGradientClippingGradientClipMaxGradientClipMinRegularizationTypeRegularizationScale) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerdescriptor/2966730-optimizerdescriptorwithlearningr
func (oc _OptimizerDescriptorClass) OptimizerDescriptorWithLearningRateGradientRescaleRegularizationTypeRegularizationScale(learningRate float32, gradientRescale float32, regularizationType RegularizationType, regularizationScale float32) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(oc.class), objc.Sel("optimizerDescriptorWithLearningRate:gradientRescale:regularizationType:regularizationScale:"), learningRate, gradientRescale, regularizationType, regularizationScale)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=OptimizerDescriptorWithLearningRateGradientRescaleRegularizationTypeRegularizationScale) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for OptimizerDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for OptimizerDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for OptimizerDescriptor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerdescriptor/2966722-applygradientclipping
func (o_ OptimizerDescriptor) ApplyGradientClipping() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](o_.ID, objc.Sel("applyGradientClipping"))
	return rv
}/* debug [instance_properties/getter]: applyGradientClipping */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerdescriptor/2966722-applygradientclipping
func (o_ OptimizerDescriptor) SetApplyGradientClipping(value objectivec.IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setApplyGradientClipping:"), value)
}/* debug [instance_properties/setter]: applyGradientClipping */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerdescriptor/2966723-gradientclipmax
func (o_ OptimizerDescriptor) GradientClipMax() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](o_.ID, objc.Sel("gradientClipMax"))
	return rv
}/* debug [instance_properties/getter]: gradientClipMax */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerdescriptor/2966723-gradientclipmax
func (o_ OptimizerDescriptor) SetGradientClipMax(value objectivec.IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setGradientClipMax:"), value)
}/* debug [instance_properties/setter]: gradientClipMax */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerdescriptor/2966724-gradientclipmin
func (o_ OptimizerDescriptor) GradientClipMin() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](o_.ID, objc.Sel("gradientClipMin"))
	return rv
}/* debug [instance_properties/getter]: gradientClipMin */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerdescriptor/2966724-gradientclipmin
func (o_ OptimizerDescriptor) SetGradientClipMin(value objectivec.IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setGradientClipMin:"), value)
}/* debug [instance_properties/setter]: gradientClipMin */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerdescriptor/2966725-gradientrescale
func (o_ OptimizerDescriptor) GradientRescale() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](o_.ID, objc.Sel("gradientRescale"))
	return rv
}/* debug [instance_properties/getter]: gradientRescale */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerdescriptor/2966725-gradientrescale
func (o_ OptimizerDescriptor) SetGradientRescale(value objectivec.IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setGradientRescale:"), value)
}/* debug [instance_properties/setter]: gradientRescale */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerdescriptor/2966728-learningrate
func (o_ OptimizerDescriptor) LearningRate() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](o_.ID, objc.Sel("learningRate"))
	return rv
}/* debug [instance_properties/getter]: learningRate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerdescriptor/2966728-learningrate
func (o_ OptimizerDescriptor) SetLearningRate(value objectivec.IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setLearningRate:"), value)
}/* debug [instance_properties/setter]: learningRate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerdescriptor/2966731-regularizationscale
func (o_ OptimizerDescriptor) RegularizationScale() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](o_.ID, objc.Sel("regularizationScale"))
	return rv
}/* debug [instance_properties/getter]: regularizationScale */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerdescriptor/2966731-regularizationscale
func (o_ OptimizerDescriptor) SetRegularizationScale(value objectivec.IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setRegularizationScale:"), value)
}/* debug [instance_properties/setter]: regularizationScale */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerdescriptor/2966732-regularizationtype
func (o_ OptimizerDescriptor) RegularizationType() RegularizationType get set /* not a class type */ {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("regularizationType"))
	return rv
}/* debug [instance_properties/getter]: regularizationType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerdescriptor/2966732-regularizationtype
func (o_ OptimizerDescriptor) SetRegularizationType(value RegularizationType get set /* not a class type */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setRegularizationType:"), value)
}/* debug [instance_properties/setter]: regularizationType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNOptimizerDescriptor */


