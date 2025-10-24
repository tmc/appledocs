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

// A layer that normalizes all features of one channel.


// A layer that normalizes all features of one channel.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcinstancenormalizationlayer/beta
func (c_ CInstanceNormalizationLayer) Beta() IMLCTensor {
	rv := objc.Send[CTensor](c_.ID, objc.Sel("beta"))
	return rv
}


// The beta tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcinstancenormalizationlayer/beta
func (c_ CInstanceNormalizationLayer) SetBeta(value IMLCTensor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBeta:"), value)
}


// The beta tensor parameter you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcinstancenormalizationlayer/betaparameter
func (c_ CInstanceNormalizationLayer) BetaParameter() objc.IObject /* cross-framework: CTensorParameter */ {
	rv := objc.Send[CTensorParameter](c_.ID, objc.Sel("betaParameter"))
	return rv
}


// The beta tensor parameter you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcinstancenormalizationlayer/betaparameter
func (c_ CInstanceNormalizationLayer) SetBetaParameter(value objc.IObject /* cross-framework: CTensorParameter */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBetaParameter:"), value)
}


// The number of feature channels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcinstancenormalizationlayer/featurechannelcount
func (c_ CInstanceNormalizationLayer) FeatureChannelCount() int {
	rv := objc.Send[int](c_.ID, objc.Sel("featureChannelCount"))
	return rv
}


// The number of feature channels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcinstancenormalizationlayer/featurechannelcount
func (c_ CInstanceNormalizationLayer) SetFeatureChannelCount(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFeatureChannelCount:"), value)
}


// The gamma tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcinstancenormalizationlayer/gamma
func (c_ CInstanceNormalizationLayer) Gamma() IMLCTensor {
	rv := objc.Send[CTensor](c_.ID, objc.Sel("gamma"))
	return rv
}


// The gamma tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcinstancenormalizationlayer/gamma
func (c_ CInstanceNormalizationLayer) SetGamma(value IMLCTensor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGamma:"), value)
}


// The gamma tensor parameter you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcinstancenormalizationlayer/gammaparameter
func (c_ CInstanceNormalizationLayer) GammaParameter() objc.IObject /* cross-framework: CTensorParameter */ {
	rv := objc.Send[CTensorParameter](c_.ID, objc.Sel("gammaParameter"))
	return rv
}


// The gamma tensor parameter you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcinstancenormalizationlayer/gammaparameter
func (c_ CInstanceNormalizationLayer) SetGammaParameter(value objc.IObject /* cross-framework: CTensorParameter */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGammaParameter:"), value)
}


// The running mean tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcinstancenormalizationlayer/mean
func (c_ CInstanceNormalizationLayer) Mean() IMLCTensor {
	rv := objc.Send[CTensor](c_.ID, objc.Sel("mean"))
	return rv
}


// The running mean tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcinstancenormalizationlayer/mean
func (c_ CInstanceNormalizationLayer) SetMean(value IMLCTensor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMean:"), value)
}


// The momentum value for the running mean and variance computation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcinstancenormalizationlayer/momentum
func (c_ CInstanceNormalizationLayer) Momentum() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("momentum"))
	return rv
}


// The momentum value for the running mean and variance computation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcinstancenormalizationlayer/momentum
func (c_ CInstanceNormalizationLayer) SetMomentum(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMomentum:"), value)
}


// The running variance tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcinstancenormalizationlayer/variance
func (c_ CInstanceNormalizationLayer) Variance() IMLCTensor {
	rv := objc.Send[CTensor](c_.ID, objc.Sel("variance"))
	return rv
}


// The running variance tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcinstancenormalizationlayer/variance
func (c_ CInstanceNormalizationLayer) SetVariance(value IMLCTensor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVariance:"), value)
}


// The variance epsilon you use for numerical stability.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcinstancenormalizationlayer/varianceepsilon
func (c_ CInstanceNormalizationLayer) VarianceEpsilon() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("varianceEpsilon"))
	return rv
}


// The variance epsilon you use for numerical stability.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcinstancenormalizationlayer/varianceepsilon
func (c_ CInstanceNormalizationLayer) SetVarianceEpsilon(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVarianceEpsilon:"), value)
}



