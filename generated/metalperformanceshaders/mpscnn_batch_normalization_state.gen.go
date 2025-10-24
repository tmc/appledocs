// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [CNNBatchNormalizationState] class.
var (
	CNNBatchNormalizationStateClass     _CNNBatchNormalizationStateClass
	CNNBatchNormalizationStateClassOnce sync.Once
)

func getCNNBatchNormalizationStateClass() _CNNBatchNormalizationStateClass {
	CNNBatchNormalizationStateClassOnce.Do(func() {
		CNNBatchNormalizationStateClass = _CNNBatchNormalizationStateClass{objc.GetClass("MPSCNNBatchNormalizationState")}
	})
	return CNNBatchNormalizationStateClass
}

type _CNNBatchNormalizationStateClass struct {
	class objc.Class
}





// An interface definition for the [CNNBatchNormalizationState] class.
type ICNNBatchNormalizationState interface {
	IGradientState
	

	// properties:
	BatchNormalization() IMPSCNNBatchNormalization
	SetBatchNormalization(value IMPSCNNBatchNormalization)


	

	// methods:
	Reset()
	Variance()
	Mean()
	Beta()
	GradientForBeta()
	Gamma()
	GradientForGamma()


}





// Alloc allocates a new instance without initialization.
func (cc _CNNBatchNormalizationStateClass) Alloc() CNNBatchNormalizationState {
	rv := objc.Send[CNNBatchNormalizationState](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNBatchNormalizationStateClass) New() CNNBatchNormalizationState {
	rv := objc.Send[CNNBatchNormalizationState](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNBatchNormalizationState) Init() CNNBatchNormalizationState {
	rv := objc.Send[CNNBatchNormalizationState](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNBatchNormalizationState) Autorelease() CNNBatchNormalizationState {
	rv := objc.Send[CNNBatchNormalizationState](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNBatchNormalizationState creates a new CNNBatchNormalizationState instance.
func NewCNNBatchNormalizationState() CNNBatchNormalizationState {
	return getCNNBatchNormalizationStateClass().New()
}





// An object that stores data required to execute batch normalization.


// An object that stores data required to execute batch normalization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationState
type CNNBatchNormalizationState struct {
	GradientState
}

// CNNBatchNormalizationStateFrom constructs a [CNNBatchNormalizationState] from an unsafe.Pointer.
//
// An object that stores data required to execute batch normalization.
func CNNBatchNormalizationStateFrom(ptr unsafe.Pointer) CNNBatchNormalizationState {
	return CNNBatchNormalizationState{
		GradientState: GradientStateFrom(ptr),
	}
}




















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationstate/2942587-reset
func (c_ CNNBatchNormalizationState) Reset() {
	objc.Send[objc.ID](c_.ID, objc.Sel("reset"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationstate/2942603-variance
func (c_ CNNBatchNormalizationState) Variance() {
	objc.Send[objc.ID](c_.ID, objc.Sel("variance"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationstate/2942612-mean
func (c_ CNNBatchNormalizationState) Mean() {
	objc.Send[objc.ID](c_.ID, objc.Sel("mean"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationstate/2951888-beta
func (c_ CNNBatchNormalizationState) Beta() {
	objc.Send[objc.ID](c_.ID, objc.Sel("beta"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationstate/2951890-gradientforbeta
func (c_ CNNBatchNormalizationState) GradientForBeta() {
	objc.Send[objc.ID](c_.ID, objc.Sel("gradientForBeta"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationstate/2951892-gamma
func (c_ CNNBatchNormalizationState) Gamma() {
	objc.Send[objc.ID](c_.ID, objc.Sel("gamma"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationstate/2951893-gradientforgamma
func (c_ CNNBatchNormalizationState) GradientForGamma() {
	objc.Send[objc.ID](c_.ID, objc.Sel("gradientForGamma"))
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationstate/2953969-batchnormalization
func (c_ CNNBatchNormalizationState) BatchNormalization() IMPSCNNBatchNormalization {
	rv := objc.Send[CNNBatchNormalization](c_.ID, objc.Sel("batchNormalization"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationstate/2953969-batchnormalization
func (c_ CNNBatchNormalizationState) SetBatchNormalization(value IMPSCNNBatchNormalization) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBatchNormalization:"), value)
}








