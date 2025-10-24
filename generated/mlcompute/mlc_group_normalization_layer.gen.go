// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CGroupNormalizationLayer] class.
var (
	CGroupNormalizationLayerClass     _CGroupNormalizationLayerClass
	CGroupNormalizationLayerClassOnce sync.Once
)

func getCGroupNormalizationLayerClass() _CGroupNormalizationLayerClass {
	CGroupNormalizationLayerClassOnce.Do(func() {
		CGroupNormalizationLayerClass = _CGroupNormalizationLayerClass{objc.GetClass("MLCGroupNormalizationLayer")}
	})
	return CGroupNormalizationLayerClass
}

type _CGroupNormalizationLayerClass struct {
	class objc.Class
}

// An interface definition for the [CGroupNormalizationLayer] class.
type ICGroupNormalizationLayer interface {
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
	GroupCount() int
	SetGroupCount(value int)
	VarianceEpsilon() float32
	SetVarianceEpsilon(value float32)
	// methods:
}

// A layer that divides the channels into groups for normalization.


// A layer that divides the channels into groups for normalization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGroupNormalizationLayer
type CGroupNormalizationLayer struct {
	CLayer
}

// CGroupNormalizationLayerFrom constructs a [CGroupNormalizationLayer] from an unsafe.Pointer.
//
// A layer that divides the channels into groups for normalization.
func CGroupNormalizationLayerFrom(ptr unsafe.Pointer) CGroupNormalizationLayer {
	return CGroupNormalizationLayer{
		CLayer: CLayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CGroupNormalizationLayerClass) Alloc() CGroupNormalizationLayer {
	rv := objc.Send[CGroupNormalizationLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CGroupNormalizationLayerClass) New() CGroupNormalizationLayer {
	rv := objc.Send[CGroupNormalizationLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CGroupNormalizationLayer) Init() CGroupNormalizationLayer {
	rv := objc.Send[CGroupNormalizationLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CGroupNormalizationLayer) Autorelease() CGroupNormalizationLayer {
	rv := objc.Send[CGroupNormalizationLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCGroupNormalizationLayer creates a new CGroupNormalizationLayer instance.
func NewCGroupNormalizationLayer() CGroupNormalizationLayer {
	return getCGroupNormalizationLayerClass().New()
}



// The beta tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcgroupnormalizationlayer/beta
func (c_ CGroupNormalizationLayer) Beta() IMLCTensor {
	rv := objc.Send[CTensor](c_.ID, objc.Sel("beta"))
	return rv
}


// The beta tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcgroupnormalizationlayer/beta
func (c_ CGroupNormalizationLayer) SetBeta(value IMLCTensor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBeta:"), value)
}


// The beta tensor parameter you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcgroupnormalizationlayer/betaparameter
func (c_ CGroupNormalizationLayer) BetaParameter() objc.IObject /* cross-framework: CTensorParameter */ {
	rv := objc.Send[CTensorParameter](c_.ID, objc.Sel("betaParameter"))
	return rv
}


// The beta tensor parameter you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcgroupnormalizationlayer/betaparameter
func (c_ CGroupNormalizationLayer) SetBetaParameter(value objc.IObject /* cross-framework: CTensorParameter */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBetaParameter:"), value)
}


// The number of feature channels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcgroupnormalizationlayer/featurechannelcount
func (c_ CGroupNormalizationLayer) FeatureChannelCount() int {
	rv := objc.Send[int](c_.ID, objc.Sel("featureChannelCount"))
	return rv
}


// The number of feature channels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcgroupnormalizationlayer/featurechannelcount
func (c_ CGroupNormalizationLayer) SetFeatureChannelCount(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFeatureChannelCount:"), value)
}


// The gamma tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcgroupnormalizationlayer/gamma
func (c_ CGroupNormalizationLayer) Gamma() IMLCTensor {
	rv := objc.Send[CTensor](c_.ID, objc.Sel("gamma"))
	return rv
}


// The gamma tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcgroupnormalizationlayer/gamma
func (c_ CGroupNormalizationLayer) SetGamma(value IMLCTensor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGamma:"), value)
}


// The gamma tensor parameter you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcgroupnormalizationlayer/gammaparameter
func (c_ CGroupNormalizationLayer) GammaParameter() objc.IObject /* cross-framework: CTensorParameter */ {
	rv := objc.Send[CTensorParameter](c_.ID, objc.Sel("gammaParameter"))
	return rv
}


// The gamma tensor parameter you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcgroupnormalizationlayer/gammaparameter
func (c_ CGroupNormalizationLayer) SetGammaParameter(value objc.IObject /* cross-framework: CTensorParameter */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGammaParameter:"), value)
}


// The number of groups into which you separate the channels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcgroupnormalizationlayer/groupcount
func (c_ CGroupNormalizationLayer) GroupCount() int {
	rv := objc.Send[int](c_.ID, objc.Sel("groupCount"))
	return rv
}


// The number of groups into which you separate the channels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcgroupnormalizationlayer/groupcount
func (c_ CGroupNormalizationLayer) SetGroupCount(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGroupCount:"), value)
}


// The variance epsilon you use for numerical stability.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcgroupnormalizationlayer/varianceepsilon
func (c_ CGroupNormalizationLayer) VarianceEpsilon() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("varianceEpsilon"))
	return rv
}


// The variance epsilon you use for numerical stability.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcgroupnormalizationlayer/varianceepsilon
func (c_ CGroupNormalizationLayer) SetVarianceEpsilon(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVarianceEpsilon:"), value)
}



