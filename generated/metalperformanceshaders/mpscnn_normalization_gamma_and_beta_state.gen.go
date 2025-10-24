// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNNormalizationGammaAndBetaState] class.
var (
	CNNNormalizationGammaAndBetaStateClass     _CNNNormalizationGammaAndBetaStateClass
	CNNNormalizationGammaAndBetaStateClassOnce sync.Once
)

func getCNNNormalizationGammaAndBetaStateClass() _CNNNormalizationGammaAndBetaStateClass {
	CNNNormalizationGammaAndBetaStateClassOnce.Do(func() {
		CNNNormalizationGammaAndBetaStateClass = _CNNNormalizationGammaAndBetaStateClass{objc.GetClass("MPSCNNNormalizationGammaAndBetaState")}
	})
	return CNNNormalizationGammaAndBetaStateClass
}

type _CNNNormalizationGammaAndBetaStateClass struct {
	class objc.Class
}





// An interface definition for the [CNNNormalizationGammaAndBetaState] class.
type ICNNNormalizationGammaAndBetaState interface {
	IState
	

	// properties:
	Gamma() Buffer get /* not a class type */
	SetGamma(value Buffer get /* not a class type */)
	Beta() Buffer get /* not a class type */
	SetBeta(value Buffer get /* not a class type */)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNNormalizationGammaAndBetaStateClass) Alloc() CNNNormalizationGammaAndBetaState {
	rv := objc.Send[CNNNormalizationGammaAndBetaState](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNormalizationGammaAndBetaStateClass) New() CNNNormalizationGammaAndBetaState {
	rv := objc.Send[CNNNormalizationGammaAndBetaState](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNormalizationGammaAndBetaState) Init() CNNNormalizationGammaAndBetaState {
	rv := objc.Send[CNNNormalizationGammaAndBetaState](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNormalizationGammaAndBetaState) Autorelease() CNNNormalizationGammaAndBetaState {
	rv := objc.Send[CNNNormalizationGammaAndBetaState](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNormalizationGammaAndBetaState creates a new CNNNormalizationGammaAndBetaState instance.
func NewCNNNormalizationGammaAndBetaState() CNNNormalizationGammaAndBetaState {
	return getCNNNormalizationGammaAndBetaStateClass().New()
}





// An object that stores gamma and beta terms used to apply a scale and bias in instance- or batch-normalization operations.


// An object that stores gamma and beta terms used to apply a scale and bias in instance- or batch-normalization operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNormalizationGammaAndBetaState
type CNNNormalizationGammaAndBetaState struct {
	State
}

// CNNNormalizationGammaAndBetaStateFrom constructs a [CNNNormalizationGammaAndBetaState] from an unsafe.Pointer.
//
// An object that stores gamma and beta terms used to apply a scale and bias in instance- or batch-normalization operations.
func CNNNormalizationGammaAndBetaStateFrom(ptr unsafe.Pointer) CNNNormalizationGammaAndBetaState {
	return CNNNormalizationGammaAndBetaState{
		State: StateFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationgammaandbetastate/2953936-initwithgamma
func NewCNNNormalizationGammaAndBetaStateWithGammaBeta(gamma unsafe.Pointer, beta unsafe.Pointer) CNNNormalizationGammaAndBetaState {
	instance := getCNNNormalizationGammaAndBetaStateClass().Alloc()
	rv := objc.Send[CNNNormalizationGammaAndBetaState](instance.ID, objc.Sel("initWithGamma:beta:"), gamma, beta)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationgammaandbetastate/2953937-temporarystate
func (cc _CNNNormalizationGammaAndBetaStateClass) TemporaryState() {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("temporaryState"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationgammaandbetastate/2953937-temporarystatewithcommandbuffer
func (cc _CNNNormalizationGammaAndBetaStateClass) TemporaryStateWithCommandBufferNumberOfFeatureChannels(commandBuffer unsafe.Pointer, numberOfFeatureChannels uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("temporaryStateWithCommandBuffer:numberOfFeatureChannels:"), commandBuffer, numberOfFeatureChannels)
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationgammaandbetastate/2953934-gamma
func (c_ CNNNormalizationGammaAndBetaState) Gamma() Buffer get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("gamma"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationgammaandbetastate/2953934-gamma
func (c_ CNNNormalizationGammaAndBetaState) SetGamma(value Buffer get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGamma:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationgammaandbetastate/2953938-beta
func (c_ CNNNormalizationGammaAndBetaState) Beta() Buffer get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("beta"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationgammaandbetastate/2953938-beta
func (c_ CNNNormalizationGammaAndBetaState) SetBeta(value Buffer get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBeta:"), value)
}







