// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [ReduceFeatureChannelsAndWeightsSum] class.
type IReduceFeatureChannelsAndWeightsSum interface {
	IReduceBinary
	

	// properties:
	DoWeightedSumByNonZeroWeights() objectivec.IObject
	SetDoWeightedSumByNonZeroWeights(value objectivec.IObject)


	

	// methods:


}





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






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelsandweightssum/3197835-initwithcoder
func NewReduceFeatureChannelsAndWeightsSumWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) ReduceFeatureChannelsAndWeightsSum {
	instance := getReduceFeatureChannelsAndWeightsSumClass().Alloc()
	rv := objc.Send[ReduceFeatureChannelsAndWeightsSum](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelsandweightssum/2942562-initwithdevice
func NewReduceFeatureChannelsAndWeightsSumWithDevice(device unsafe.Pointer) ReduceFeatureChannelsAndWeightsSum {
	instance := getReduceFeatureChannelsAndWeightsSumClass().Alloc()
	rv := objc.Send[ReduceFeatureChannelsAndWeightsSum](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelsandweightssum/2942551-initwithdevice
func NewReduceFeatureChannelsAndWeightsSumWithDeviceDoWeightedSumByNonZeroWeights(device unsafe.Pointer, doWeightedSumByNonZeroWeights bool) ReduceFeatureChannelsAndWeightsSum {
	instance := getReduceFeatureChannelsAndWeightsSumClass().Alloc()
	rv := objc.Send[ReduceFeatureChannelsAndWeightsSum](instance.ID, objc.Sel("initWithDevice:doWeightedSumByNonZeroWeights:"), device, doWeightedSumByNonZeroWeights)
	rv.Autorelease()
	return rv
}






















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelsandweightssum/2942543-doweightedsumbynonzeroweights
func (r_ ReduceFeatureChannelsAndWeightsSum) DoWeightedSumByNonZeroWeights() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](r_.ID, objc.Sel("doWeightedSumByNonZeroWeights"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelsandweightssum/2942543-doweightedsumbynonzeroweights
func (r_ ReduceFeatureChannelsAndWeightsSum) SetDoWeightedSumByNonZeroWeights(value objectivec.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setDoWeightedSumByNonZeroWeights:"), value)
}







