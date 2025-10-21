// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CBatchNormalizationLayer] class.
var (
	CBatchNormalizationLayerClass     _CBatchNormalizationLayerClass
	CBatchNormalizationLayerClassOnce sync.Once
)

func getCBatchNormalizationLayerClass() _CBatchNormalizationLayerClass {
	CBatchNormalizationLayerClassOnce.Do(func() {
		CBatchNormalizationLayerClass = _CBatchNormalizationLayerClass{objc.GetClass("MLCBatchNormalizationLayer")}
	})
	return CBatchNormalizationLayerClass
}

type _CBatchNormalizationLayerClass struct {
	class objc.Class
}

// An interface definition for the [CBatchNormalizationLayer] class.
type ICBatchNormalizationLayer interface {
	ICLayer
}

// A layer that normalizes a batch of inputs.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCBatchNormalizationLayer
type CBatchNormalizationLayer struct {
	CLayer
}

// CBatchNormalizationLayerFrom constructs a [CBatchNormalizationLayer] from an unsafe.Pointer.
//
// A layer that normalizes a batch of inputs.
func CBatchNormalizationLayerFrom(ptr unsafe.Pointer) CBatchNormalizationLayer {
	return CBatchNormalizationLayer{
		CLayer: CLayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CBatchNormalizationLayerClass) Alloc() CBatchNormalizationLayer {
	rv := objc.Send[CBatchNormalizationLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CBatchNormalizationLayerClass) New() CBatchNormalizationLayer {
	rv := objc.Send[CBatchNormalizationLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CBatchNormalizationLayer) Init() CBatchNormalizationLayer {
	rv := objc.Send[CBatchNormalizationLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CBatchNormalizationLayer) Autorelease() CBatchNormalizationLayer {
	rv := objc.Send[CBatchNormalizationLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCBatchNormalizationLayer creates a new CBatchNormalizationLayer instance.
func NewCBatchNormalizationLayer() CBatchNormalizationLayer {
	return getCBatchNormalizationLayerClass().New()
}


// The beta tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcbatchnormalizationlayer/beta
func (c_ CBatchNormalizationLayer) Beta() MLCTensor {
	rv := objc.Send[MLCTensor](c_.ID, objc.Sel("beta"))
	return rv
}


// SetBeta sets the value of the beta property.
// The beta tensor.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcbatchnormalizationlayer/beta
func (c_ CBatchNormalizationLayer) SetBeta(value IMLCTensor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBeta:"), value)
}

// The beta tensor parameter you use for optimizer updates.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcbatchnormalizationlayer/betaparameter
func (c_ CBatchNormalizationLayer) BetaParameter() MLCTensorParameter {
	rv := objc.Send[MLCTensorParameter](c_.ID, objc.Sel("betaParameter"))
	return rv
}


// SetBetaParameter sets the value of the betaParameter property.
// The beta tensor parameter you use for optimizer updates.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcbatchnormalizationlayer/betaparameter
func (c_ CBatchNormalizationLayer) SetBetaParameter(value IMLCTensorParameter) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBetaParameter:"), value)
}

// The number of feature channels.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcbatchnormalizationlayer/featurechannelcount
func (c_ CBatchNormalizationLayer) FeatureChannelCount() int {
	rv := objc.Send[int](c_.ID, objc.Sel("featureChannelCount"))
	return rv
}


// SetFeatureChannelCount sets the value of the featureChannelCount property.
// The number of feature channels.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcbatchnormalizationlayer/featurechannelcount
func (c_ CBatchNormalizationLayer) SetFeatureChannelCount(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFeatureChannelCount:"), value)
}

// The gamma tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcbatchnormalizationlayer/gamma
func (c_ CBatchNormalizationLayer) Gamma() MLCTensor {
	rv := objc.Send[MLCTensor](c_.ID, objc.Sel("gamma"))
	return rv
}


// SetGamma sets the value of the gamma property.
// The gamma tensor.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcbatchnormalizationlayer/gamma
func (c_ CBatchNormalizationLayer) SetGamma(value IMLCTensor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGamma:"), value)
}

// The gamma tensor parameter you use for optimizer updates.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcbatchnormalizationlayer/gammaparameter
func (c_ CBatchNormalizationLayer) GammaParameter() MLCTensorParameter {
	rv := objc.Send[MLCTensorParameter](c_.ID, objc.Sel("gammaParameter"))
	return rv
}


// SetGammaParameter sets the value of the gammaParameter property.
// The gamma tensor parameter you use for optimizer updates.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcbatchnormalizationlayer/gammaparameter
func (c_ CBatchNormalizationLayer) SetGammaParameter(value IMLCTensorParameter) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGammaParameter:"), value)
}

// The mean tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcbatchnormalizationlayer/mean
func (c_ CBatchNormalizationLayer) Mean() MLCTensor {
	rv := objc.Send[MLCTensor](c_.ID, objc.Sel("mean"))
	return rv
}


// SetMean sets the value of the mean property.
// The mean tensor.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcbatchnormalizationlayer/mean
func (c_ CBatchNormalizationLayer) SetMean(value IMLCTensor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMean:"), value)
}

// The value you use for the running mean and variance computation.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcbatchnormalizationlayer/momentum
func (c_ CBatchNormalizationLayer) Momentum() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("momentum"))
	return rv
}


// SetMomentum sets the value of the momentum property.
// The value you use for the running mean and variance computation.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcbatchnormalizationlayer/momentum
func (c_ CBatchNormalizationLayer) SetMomentum(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMomentum:"), value)
}

// The variance tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcbatchnormalizationlayer/variance
func (c_ CBatchNormalizationLayer) Variance() MLCTensor {
	rv := objc.Send[MLCTensor](c_.ID, objc.Sel("variance"))
	return rv
}


// SetVariance sets the value of the variance property.
// The variance tensor.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcbatchnormalizationlayer/variance
func (c_ CBatchNormalizationLayer) SetVariance(value IMLCTensor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVariance:"), value)
}

// The variance epsilon you use for numerical stability.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcbatchnormalizationlayer/varianceepsilon
func (c_ CBatchNormalizationLayer) VarianceEpsilon() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("varianceEpsilon"))
	return rv
}


// SetVarianceEpsilon sets the value of the varianceEpsilon property.
// The variance epsilon you use for numerical stability.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcbatchnormalizationlayer/varianceepsilon
func (c_ CBatchNormalizationLayer) SetVarianceEpsilon(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVarianceEpsilon:"), value)
}



