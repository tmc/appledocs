// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CInstanceNormalizationLayer] class.
var (
	CInstanceNormalizationLayerClass     _CInstanceNormalizationLayerClass
	CInstanceNormalizationLayerClassOnce sync.Once
)

func getCInstanceNormalizationLayerClass() _CInstanceNormalizationLayerClass {
	CInstanceNormalizationLayerClassOnce.Do(func() {
		CInstanceNormalizationLayerClass = _CInstanceNormalizationLayerClass{objc.GetClass("MLCInstanceNormalizationLayer")}
	})
	return CInstanceNormalizationLayerClass
}

type _CInstanceNormalizationLayerClass struct {
	class objc.Class
}

// An interface definition for the [CInstanceNormalizationLayer] class.
type ICInstanceNormalizationLayer interface {
	ICLayer
	Beta() MLCTensor
	SetBeta(value IMLCTensor)
	BetaParameter() MLCTensorParameter
	SetBetaParameter(value IMLCTensorParameter)
	FeatureChannelCount() int
	SetFeatureChannelCount(value int)
	Gamma() MLCTensor
	SetGamma(value IMLCTensor)
	GammaParameter() MLCTensorParameter
	SetGammaParameter(value IMLCTensorParameter)
	Mean() MLCTensor
	SetMean(value IMLCTensor)
	Momentum() float32
	SetMomentum(value float32)
	Variance() MLCTensor
	SetVariance(value IMLCTensor)
	VarianceEpsilon() float32
	SetVarianceEpsilon(value float32)
}

// A layer that normalizes all features of one channel.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCInstanceNormalizationLayer
type CInstanceNormalizationLayer struct {
	CLayer
}

// CInstanceNormalizationLayerFrom constructs a [CInstanceNormalizationLayer] from an unsafe.Pointer.
//
// A layer that normalizes all features of one channel.
func CInstanceNormalizationLayerFrom(ptr unsafe.Pointer) CInstanceNormalizationLayer {
	return CInstanceNormalizationLayer{
		CLayer: CLayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CInstanceNormalizationLayerClass) Alloc() CInstanceNormalizationLayer {
	rv := objc.Send[CInstanceNormalizationLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CInstanceNormalizationLayerClass) New() CInstanceNormalizationLayer {
	rv := objc.Send[CInstanceNormalizationLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CInstanceNormalizationLayer) Init() CInstanceNormalizationLayer {
	rv := objc.Send[CInstanceNormalizationLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CInstanceNormalizationLayer) Autorelease() CInstanceNormalizationLayer {
	rv := objc.Send[CInstanceNormalizationLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCInstanceNormalizationLayer creates a new CInstanceNormalizationLayer instance.
func NewCInstanceNormalizationLayer() CInstanceNormalizationLayer {
	return getCInstanceNormalizationLayerClass().New()
}


// The beta tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcinstancenormalizationlayer/beta
func (c_ CInstanceNormalizationLayer) Beta() MLCTensor {
	rv := objc.Send[MLCTensor](c_.ID, objc.Sel("beta"))
	return rv
}


// SetBeta sets the value of the beta property.
// The beta tensor.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcinstancenormalizationlayer/beta
func (c_ CInstanceNormalizationLayer) SetBeta(value IMLCTensor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBeta:"), value)
}

// The beta tensor parameter you use for optimizer updates.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcinstancenormalizationlayer/betaparameter
func (c_ CInstanceNormalizationLayer) BetaParameter() MLCTensorParameter {
	rv := objc.Send[MLCTensorParameter](c_.ID, objc.Sel("betaParameter"))
	return rv
}


// SetBetaParameter sets the value of the betaParameter property.
// The beta tensor parameter you use for optimizer updates.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcinstancenormalizationlayer/betaparameter
func (c_ CInstanceNormalizationLayer) SetBetaParameter(value IMLCTensorParameter) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBetaParameter:"), value)
}

// The number of feature channels.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcinstancenormalizationlayer/featurechannelcount
func (c_ CInstanceNormalizationLayer) FeatureChannelCount() int {
	rv := objc.Send[int](c_.ID, objc.Sel("featureChannelCount"))
	return rv
}


// SetFeatureChannelCount sets the value of the featureChannelCount property.
// The number of feature channels.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcinstancenormalizationlayer/featurechannelcount
func (c_ CInstanceNormalizationLayer) SetFeatureChannelCount(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFeatureChannelCount:"), value)
}

// The gamma tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcinstancenormalizationlayer/gamma
func (c_ CInstanceNormalizationLayer) Gamma() MLCTensor {
	rv := objc.Send[MLCTensor](c_.ID, objc.Sel("gamma"))
	return rv
}


// SetGamma sets the value of the gamma property.
// The gamma tensor.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcinstancenormalizationlayer/gamma
func (c_ CInstanceNormalizationLayer) SetGamma(value IMLCTensor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGamma:"), value)
}

// The gamma tensor parameter you use for optimizer updates.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcinstancenormalizationlayer/gammaparameter
func (c_ CInstanceNormalizationLayer) GammaParameter() MLCTensorParameter {
	rv := objc.Send[MLCTensorParameter](c_.ID, objc.Sel("gammaParameter"))
	return rv
}


// SetGammaParameter sets the value of the gammaParameter property.
// The gamma tensor parameter you use for optimizer updates.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcinstancenormalizationlayer/gammaparameter
func (c_ CInstanceNormalizationLayer) SetGammaParameter(value IMLCTensorParameter) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGammaParameter:"), value)
}

// The running mean tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcinstancenormalizationlayer/mean
func (c_ CInstanceNormalizationLayer) Mean() MLCTensor {
	rv := objc.Send[MLCTensor](c_.ID, objc.Sel("mean"))
	return rv
}


// SetMean sets the value of the mean property.
// The running mean tensor.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcinstancenormalizationlayer/mean
func (c_ CInstanceNormalizationLayer) SetMean(value IMLCTensor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMean:"), value)
}

// The momentum value for the running mean and variance computation.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcinstancenormalizationlayer/momentum
func (c_ CInstanceNormalizationLayer) Momentum() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("momentum"))
	return rv
}


// SetMomentum sets the value of the momentum property.
// The momentum value for the running mean and variance computation.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcinstancenormalizationlayer/momentum
func (c_ CInstanceNormalizationLayer) SetMomentum(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMomentum:"), value)
}

// The running variance tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcinstancenormalizationlayer/variance
func (c_ CInstanceNormalizationLayer) Variance() MLCTensor {
	rv := objc.Send[MLCTensor](c_.ID, objc.Sel("variance"))
	return rv
}


// SetVariance sets the value of the variance property.
// The running variance tensor.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcinstancenormalizationlayer/variance
func (c_ CInstanceNormalizationLayer) SetVariance(value IMLCTensor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVariance:"), value)
}

// The variance epsilon you use for numerical stability.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcinstancenormalizationlayer/varianceepsilon
func (c_ CInstanceNormalizationLayer) VarianceEpsilon() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("varianceEpsilon"))
	return rv
}


// SetVarianceEpsilon sets the value of the varianceEpsilon property.
// The variance epsilon you use for numerical stability.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcinstancenormalizationlayer/varianceepsilon
func (c_ CInstanceNormalizationLayer) SetVarianceEpsilon(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVarianceEpsilon:"), value)
}



