// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNNReduceFeatureChannelsAndWeightsSum */


/* debug [class_header]: Header for MPSNNReduceFeatureChannelsAndWeightsSum */
// The class instance for the [ReduceFeatureChannelsAndWeightsSum] class.
var (
	ReduceFeatureChannelsAndWeightsSumClass     _ReduceFeatureChannelsAndWeightsSumClass
	ReduceFeatureChannelsAndWeightsSumClassOnce sync.Once
)

func getReduceFeatureChannelsAndWeightsSumClass() _ReduceFeatureChannelsAndWeightsSumClass {
	ReduceFeatureChannelsAndWeightsSumClassOnce.Do(func() {
		ReduceFeatureChannelsAndWeightsSumClass = _ReduceFeatureChannelsAndWeightsSumClass{objc.GetClass("MPSNNReduceFeatureChannelsAndWeightsSum")}
	})
	return ReduceFeatureChannelsAndWeightsSumClass
}

type _ReduceFeatureChannelsAndWeightsSumClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ReduceFeatureChannelsAndWeightsSum */
// An interface definition for the [ReduceFeatureChannelsAndWeightsSum] class.
type IReduceFeatureChannelsAndWeightsSum interface {
	IReduceBinary
	
/* debug [class_interface_properties]: Properties for ReduceFeatureChannelsAndWeightsSum */
	// properties:
	DoWeightedSumByNonZeroWeights() objectivec.IObject
	SetDoWeightedSumByNonZeroWeights(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ReduceFeatureChannelsAndWeightsSum */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ReduceFeatureChannelsAndWeightsSum */
// Alloc allocates a new instance without initialization.
func (rc _ReduceFeatureChannelsAndWeightsSumClass) Alloc() ReduceFeatureChannelsAndWeightsSum {
	rv := objc.Send[ReduceFeatureChannelsAndWeightsSum](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReduceFeatureChannelsAndWeightsSumClass) New() ReduceFeatureChannelsAndWeightsSum {
	rv := objc.Send[ReduceFeatureChannelsAndWeightsSum](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReduceFeatureChannelsAndWeightsSum) Init() ReduceFeatureChannelsAndWeightsSum {
	rv := objc.Send[ReduceFeatureChannelsAndWeightsSum](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReduceFeatureChannelsAndWeightsSum) Autorelease() ReduceFeatureChannelsAndWeightsSum {
	rv := objc.Send[ReduceFeatureChannelsAndWeightsSum](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReduceFeatureChannelsAndWeightsSum creates a new ReduceFeatureChannelsAndWeightsSum instance.
func NewReduceFeatureChannelsAndWeightsSum() ReduceFeatureChannelsAndWeightsSum {
	return getReduceFeatureChannelsAndWeightsSumClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ReduceFeatureChannelsAndWeightsSum */
// A reduction filter that returns the weighted sum of all values for each feature channel in an image.


// A reduction filter that returns the weighted sum of all values for each feature channel in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceFeatureChannelsAndWeightsSum
type ReduceFeatureChannelsAndWeightsSum struct {
	ReduceBinary
}

// ReduceFeatureChannelsAndWeightsSumFrom constructs a [ReduceFeatureChannelsAndWeightsSum] from an unsafe.Pointer.
//
// A reduction filter that returns the weighted sum of all values for each feature channel in an image.
func ReduceFeatureChannelsAndWeightsSumFrom(ptr unsafe.Pointer) ReduceFeatureChannelsAndWeightsSum {
	return ReduceFeatureChannelsAndWeightsSum{
		ReduceBinary: ReduceBinaryFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ReduceFeatureChannelsAndWeightsSum */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelsandweightssum/3197835-initwithcoder
func NewReduceFeatureChannelsAndWeightsSumWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) ReduceFeatureChannelsAndWeightsSum {
	instance := getReduceFeatureChannelsAndWeightsSumClass().Alloc()
	rv := objc.Send[ReduceFeatureChannelsAndWeightsSum](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewReduceFeatureChannelsAndWeightsSumWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelsandweightssum/2942562-initwithdevice
func NewReduceFeatureChannelsAndWeightsSumWithDevice(device unsafe.Pointer) ReduceFeatureChannelsAndWeightsSum {
	instance := getReduceFeatureChannelsAndWeightsSumClass().Alloc()
	rv := objc.Send[ReduceFeatureChannelsAndWeightsSum](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewReduceFeatureChannelsAndWeightsSumWithDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelsandweightssum/2942551-initwithdevice
func NewReduceFeatureChannelsAndWeightsSumWithDeviceDoWeightedSumByNonZeroWeights(device unsafe.Pointer, doWeightedSumByNonZeroWeights bool) ReduceFeatureChannelsAndWeightsSum {
	instance := getReduceFeatureChannelsAndWeightsSumClass().Alloc()
	rv := objc.Send[ReduceFeatureChannelsAndWeightsSum](instance.ID, objc.Sel("initWithDevice:doWeightedSumByNonZeroWeights:"), device, doWeightedSumByNonZeroWeights)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewReduceFeatureChannelsAndWeightsSumWithDeviceDoWeightedSumByNonZeroWeights */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ReduceFeatureChannelsAndWeightsSum */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ReduceFeatureChannelsAndWeightsSum */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ReduceFeatureChannelsAndWeightsSum */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ReduceFeatureChannelsAndWeightsSum */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelsandweightssum/2942543-doweightedsumbynonzeroweights
func (r_ ReduceFeatureChannelsAndWeightsSum) DoWeightedSumByNonZeroWeights() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](r_.ID, objc.Sel("doWeightedSumByNonZeroWeights"))
	return rv
}/* debug [instance_properties/getter]: doWeightedSumByNonZeroWeights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelsandweightssum/2942543-doweightedsumbynonzeroweights
func (r_ ReduceFeatureChannelsAndWeightsSum) SetDoWeightedSumByNonZeroWeights(value objectivec.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setDoWeightedSumByNonZeroWeights:"), value)
}/* debug [instance_properties/setter]: doWeightedSumByNonZeroWeights */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNReduceFeatureChannelsAndWeightsSum */


