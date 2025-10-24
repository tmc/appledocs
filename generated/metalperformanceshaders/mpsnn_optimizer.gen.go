// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [Optimizer] class.
var (
	OptimizerClass     _OptimizerClass
	OptimizerClassOnce sync.Once
)

func getOptimizerClass() _OptimizerClass {
	OptimizerClassOnce.Do(func() {
		OptimizerClass = _OptimizerClass{objc.GetClass("MPSNNOptimizer")}
	})
	return OptimizerClass
}

type _OptimizerClass struct {
	class objc.Class
}





// An interface definition for the [Optimizer] class.
type IOptimizer interface {
	IKernel
	

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
	RegularizationType() RegularizationType get /* not a class type */
	SetRegularizationType(value RegularizationType get /* not a class type */)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (oc _OptimizerClass) Alloc() Optimizer {
	rv := objc.Send[Optimizer](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (oc _OptimizerClass) New() Optimizer {
	rv := objc.Send[Optimizer](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ Optimizer) Init() Optimizer {
	rv := objc.Send[Optimizer](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ Optimizer) Autorelease() Optimizer {
	rv := objc.Send[Optimizer](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOptimizer creates a new Optimizer instance.
func NewOptimizer() Optimizer {
	return getOptimizerClass().New()
}





// The base class for optimization layers.


// The base class for optimization layers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNOptimizer
type Optimizer struct {
	Kernel
}

// OptimizerFrom constructs a [Optimizer] from an unsafe.Pointer.
//
// The base class for optimization layers.
func OptimizerFrom(ptr unsafe.Pointer) Optimizer {
	return Optimizer{
		Kernel: KernelFrom(ptr),
	}
}

























// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizer/2966705-applygradientclipping
func (o_ Optimizer) ApplyGradientClipping() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](o_.ID, objc.Sel("applyGradientClipping"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizer/2966705-applygradientclipping
func (o_ Optimizer) SetApplyGradientClipping(value objectivec.IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setApplyGradientClipping:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizer/2966706-gradientclipmax
func (o_ Optimizer) GradientClipMax() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](o_.ID, objc.Sel("gradientClipMax"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizer/2966706-gradientclipmax
func (o_ Optimizer) SetGradientClipMax(value objectivec.IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setGradientClipMax:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizer/2966707-gradientclipmin
func (o_ Optimizer) GradientClipMin() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](o_.ID, objc.Sel("gradientClipMin"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizer/2966707-gradientclipmin
func (o_ Optimizer) SetGradientClipMin(value objectivec.IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setGradientClipMin:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizer/2966708-gradientrescale
func (o_ Optimizer) GradientRescale() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](o_.ID, objc.Sel("gradientRescale"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizer/2966708-gradientrescale
func (o_ Optimizer) SetGradientRescale(value objectivec.IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setGradientRescale:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizer/2966709-learningrate
func (o_ Optimizer) LearningRate() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](o_.ID, objc.Sel("learningRate"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizer/2966709-learningrate
func (o_ Optimizer) SetLearningRate(value objectivec.IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setLearningRate:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizer/2966710-regularizationscale
func (o_ Optimizer) RegularizationScale() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](o_.ID, objc.Sel("regularizationScale"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizer/2966710-regularizationscale
func (o_ Optimizer) SetRegularizationScale(value objectivec.IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setRegularizationScale:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizer/2966711-regularizationtype
func (o_ Optimizer) RegularizationType() RegularizationType get /* not a class type */ {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("regularizationType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizer/2966711-regularizationtype
func (o_ Optimizer) SetRegularizationType(value RegularizationType get /* not a class type */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setRegularizationType:"), value)
}








