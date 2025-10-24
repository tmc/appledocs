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
	// properties:
	Beta() IMLCTensor
	SetBeta(value IMLCTensor)
	BetaParameter() objc.IObject /* cross-framework: CTensorParameter */
	SetBetaParameter(value objc.IObject /* cross-framework: CTensorParameter */)
	FeatureChannelCount() int
	SetFeatureChannelCount(value int)
	Gamma() IMLCTensor
	SetGamma(value IMLCTensor)
	GammaParameter() objc.IObject /* cross-framework: CTensorParameter */
	SetGammaParameter(value objc.IObject /* cross-framework: CTensorParameter */)
	Mean() IMLCTensor
	SetMean(value IMLCTensor)
	Momentum() float32
	SetMomentum(value float32)
	Variance() IMLCTensor
	SetVariance(value IMLCTensor)
	VarianceEpsilon() float32
	SetVarianceEpsilon(value float32)
	// methods:
}

// A layer that normalizes a batch of inputs.


// A layer that normalizes a batch of inputs.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcbatchnormalizationlayer/beta
func (c_ CBatchNormalizationLayer) Beta() IMLCTensor {
	rv := objc.Send[CTensor](c_.ID, objc.Sel("beta"))
	return rv
}


// The beta tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcbatchnormalizationlayer/beta
func (c_ CBatchNormalizationLayer) SetBeta(value IMLCTensor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBeta:"), value)
}


// The beta tensor parameter you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcbatchnormalizationlayer/betaparameter
func (c_ CBatchNormalizationLayer) BetaParameter() objc.IObject /* cross-framework: CTensorParameter */ {
	rv := objc.Send[CTensorParameter](c_.ID, objc.Sel("betaParameter"))
	return rv
}


// The beta tensor parameter you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcbatchnormalizationlayer/betaparameter
func (c_ CBatchNormalizationLayer) SetBetaParameter(value objc.IObject /* cross-framework: CTensorParameter */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBetaParameter:"), value)
}


// The number of feature channels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcbatchnormalizationlayer/featurechannelcount
func (c_ CBatchNormalizationLayer) FeatureChannelCount() int {
	rv := objc.Send[int](c_.ID, objc.Sel("featureChannelCount"))
	return rv
}


// The number of feature channels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcbatchnormalizationlayer/featurechannelcount
func (c_ CBatchNormalizationLayer) SetFeatureChannelCount(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFeatureChannelCount:"), value)
}


// The gamma tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcbatchnormalizationlayer/gamma
func (c_ CBatchNormalizationLayer) Gamma() IMLCTensor {
	rv := objc.Send[CTensor](c_.ID, objc.Sel("gamma"))
	return rv
}


// The gamma tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcbatchnormalizationlayer/gamma
func (c_ CBatchNormalizationLayer) SetGamma(value IMLCTensor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGamma:"), value)
}


// The gamma tensor parameter you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcbatchnormalizationlayer/gammaparameter
func (c_ CBatchNormalizationLayer) GammaParameter() objc.IObject /* cross-framework: CTensorParameter */ {
	rv := objc.Send[CTensorParameter](c_.ID, objc.Sel("gammaParameter"))
	return rv
}


// The gamma tensor parameter you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcbatchnormalizationlayer/gammaparameter
func (c_ CBatchNormalizationLayer) SetGammaParameter(value objc.IObject /* cross-framework: CTensorParameter */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGammaParameter:"), value)
}


// The mean tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcbatchnormalizationlayer/mean
func (c_ CBatchNormalizationLayer) Mean() IMLCTensor {
	rv := objc.Send[CTensor](c_.ID, objc.Sel("mean"))
	return rv
}


// The mean tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcbatchnormalizationlayer/mean
func (c_ CBatchNormalizationLayer) SetMean(value IMLCTensor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMean:"), value)
}


// The value you use for the running mean and variance computation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcbatchnormalizationlayer/momentum
func (c_ CBatchNormalizationLayer) Momentum() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("momentum"))
	return rv
}


// The value you use for the running mean and variance computation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcbatchnormalizationlayer/momentum
func (c_ CBatchNormalizationLayer) SetMomentum(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMomentum:"), value)
}


// The variance tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcbatchnormalizationlayer/variance
func (c_ CBatchNormalizationLayer) Variance() IMLCTensor {
	rv := objc.Send[CTensor](c_.ID, objc.Sel("variance"))
	return rv
}


// The variance tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcbatchnormalizationlayer/variance
func (c_ CBatchNormalizationLayer) SetVariance(value IMLCTensor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVariance:"), value)
}


// The variance epsilon you use for numerical stability.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcbatchnormalizationlayer/varianceepsilon
func (c_ CBatchNormalizationLayer) VarianceEpsilon() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("varianceEpsilon"))
	return rv
}


// The variance epsilon you use for numerical stability.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcbatchnormalizationlayer/varianceepsilon
func (c_ CBatchNormalizationLayer) SetVarianceEpsilon(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVarianceEpsilon:"), value)
}



