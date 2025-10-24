// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNNormalizationMeanAndVarianceState] class.
var (
	CNNNormalizationMeanAndVarianceStateClass     _CNNNormalizationMeanAndVarianceStateClass
	CNNNormalizationMeanAndVarianceStateClassOnce sync.Once
)

func getCNNNormalizationMeanAndVarianceStateClass() _CNNNormalizationMeanAndVarianceStateClass {
	CNNNormalizationMeanAndVarianceStateClassOnce.Do(func() {
		CNNNormalizationMeanAndVarianceStateClass = _CNNNormalizationMeanAndVarianceStateClass{objc.GetClass("MPSCNNNormalizationMeanAndVarianceState")}
	})
	return CNNNormalizationMeanAndVarianceStateClass
}

type _CNNNormalizationMeanAndVarianceStateClass struct {
	class objc.Class
}





// An interface definition for the [CNNNormalizationMeanAndVarianceState] class.
type ICNNNormalizationMeanAndVarianceState interface {
	IState
	

	// properties:
	Mean() Buffer get /* not a class type */
	SetMean(value Buffer get /* not a class type */)
	Variance() Buffer get /* not a class type */
	SetVariance(value Buffer get /* not a class type */)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNNormalizationMeanAndVarianceStateClass) Alloc() CNNNormalizationMeanAndVarianceState {
	rv := objc.Send[CNNNormalizationMeanAndVarianceState](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNormalizationMeanAndVarianceStateClass) New() CNNNormalizationMeanAndVarianceState {
	rv := objc.Send[CNNNormalizationMeanAndVarianceState](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNormalizationMeanAndVarianceState) Init() CNNNormalizationMeanAndVarianceState {
	rv := objc.Send[CNNNormalizationMeanAndVarianceState](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNormalizationMeanAndVarianceState) Autorelease() CNNNormalizationMeanAndVarianceState {
	rv := objc.Send[CNNNormalizationMeanAndVarianceState](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNormalizationMeanAndVarianceState creates a new CNNNormalizationMeanAndVarianceState instance.
func NewCNNNormalizationMeanAndVarianceState() CNNNormalizationMeanAndVarianceState {
	return getCNNNormalizationMeanAndVarianceStateClass().New()
}





// An object that stores mean and variance terms used to execute batch normalization.


// An object that stores mean and variance terms used to execute batch normalization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNormalizationMeanAndVarianceState
type CNNNormalizationMeanAndVarianceState struct {
	State
}

// CNNNormalizationMeanAndVarianceStateFrom constructs a [CNNNormalizationMeanAndVarianceState] from an unsafe.Pointer.
//
// An object that stores mean and variance terms used to execute batch normalization.
func CNNNormalizationMeanAndVarianceStateFrom(ptr unsafe.Pointer) CNNNormalizationMeanAndVarianceState {
	return CNNNormalizationMeanAndVarianceState{
		State: StateFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationmeanandvariancestate/3002363-initwithmean
func NewCNNNormalizationMeanAndVarianceStateWithMeanVariance(mean unsafe.Pointer, variance unsafe.Pointer) CNNNormalizationMeanAndVarianceState {
	instance := getCNNNormalizationMeanAndVarianceStateClass().Alloc()
	rv := objc.Send[CNNNormalizationMeanAndVarianceState](instance.ID, objc.Sel("initWithMean:variance:"), mean, variance)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationmeanandvariancestate/3002365-temporarystate
func (cc _CNNNormalizationMeanAndVarianceStateClass) TemporaryState() {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("temporaryState"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationmeanandvariancestate/3002365-temporarystatewithcommandbuffer
func (cc _CNNNormalizationMeanAndVarianceStateClass) TemporaryStateWithCommandBufferNumberOfFeatureChannels(commandBuffer unsafe.Pointer, numberOfFeatureChannels uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("temporaryStateWithCommandBuffer:numberOfFeatureChannels:"), commandBuffer, numberOfFeatureChannels)
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationmeanandvariancestate/3002364-mean
func (c_ CNNNormalizationMeanAndVarianceState) Mean() Buffer get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("mean"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationmeanandvariancestate/3002364-mean
func (c_ CNNNormalizationMeanAndVarianceState) SetMean(value Buffer get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMean:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationmeanandvariancestate/3002366-variance
func (c_ CNNNormalizationMeanAndVarianceState) Variance() Buffer get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("variance"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationmeanandvariancestate/3002366-variance
func (c_ CNNNormalizationMeanAndVarianceState) SetVariance(value Buffer get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVariance:"), value)
}







