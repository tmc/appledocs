// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CLayerNormalizationLayer] class.
var (
	CLayerNormalizationLayerClass     _CLayerNormalizationLayerClass
	CLayerNormalizationLayerClassOnce sync.Once
)

func getCLayerNormalizationLayerClass() _CLayerNormalizationLayerClass {
	CLayerNormalizationLayerClassOnce.Do(func() {
		CLayerNormalizationLayerClass = _CLayerNormalizationLayerClass{objc.GetClass("MLCLayerNormalizationLayer")}
	})
	return CLayerNormalizationLayerClass
}

type _CLayerNormalizationLayerClass struct {
	class objc.Class
}

// An interface definition for the [CLayerNormalizationLayer] class.
type ICLayerNormalizationLayer interface {
	ICLayer
	// properties:
	Beta() IMLCTensor
	SetBeta(value IMLCTensor)
	BetaParameter() objc.IObject /* cross-framework: CTensorParameter */
	SetBetaParameter(value objc.IObject /* cross-framework: CTensorParameter */)
	Gamma() IMLCTensor
	SetGamma(value IMLCTensor)
	GammaParameter() objc.IObject /* cross-framework: CTensorParameter */
	SetGammaParameter(value objc.IObject /* cross-framework: CTensorParameter */)
	NormalizedShape() int
	SetNormalizedShape(value int)
	VarianceEpsilon() float32
	SetVarianceEpsilon(value float32)
	// methods:
}

// A layer that applies layer normalization over inputs.


// A layer that applies layer normalization over inputs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLayerNormalizationLayer
type CLayerNormalizationLayer struct {
	CLayer
}

// CLayerNormalizationLayerFrom constructs a [CLayerNormalizationLayer] from an unsafe.Pointer.
//
// A layer that applies layer normalization over inputs.
func CLayerNormalizationLayerFrom(ptr unsafe.Pointer) CLayerNormalizationLayer {
	return CLayerNormalizationLayer{
		CLayer: CLayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CLayerNormalizationLayerClass) Alloc() CLayerNormalizationLayer {
	rv := objc.Send[CLayerNormalizationLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CLayerNormalizationLayerClass) New() CLayerNormalizationLayer {
	rv := objc.Send[CLayerNormalizationLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CLayerNormalizationLayer) Init() CLayerNormalizationLayer {
	rv := objc.Send[CLayerNormalizationLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CLayerNormalizationLayer) Autorelease() CLayerNormalizationLayer {
	rv := objc.Send[CLayerNormalizationLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCLayerNormalizationLayer creates a new CLayerNormalizationLayer instance.
func NewCLayerNormalizationLayer() CLayerNormalizationLayer {
	return getCLayerNormalizationLayerClass().New()
}



// The beta tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclayernormalizationlayer/beta
func (c_ CLayerNormalizationLayer) Beta() IMLCTensor {
	rv := objc.Send[CTensor](c_.ID, objc.Sel("beta"))
	return rv
}


// The beta tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclayernormalizationlayer/beta
func (c_ CLayerNormalizationLayer) SetBeta(value IMLCTensor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBeta:"), value)
}


// The beta tensor parameter you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclayernormalizationlayer/betaparameter
func (c_ CLayerNormalizationLayer) BetaParameter() objc.IObject /* cross-framework: CTensorParameter */ {
	rv := objc.Send[CTensorParameter](c_.ID, objc.Sel("betaParameter"))
	return rv
}


// The beta tensor parameter you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclayernormalizationlayer/betaparameter
func (c_ CLayerNormalizationLayer) SetBetaParameter(value objc.IObject /* cross-framework: CTensorParameter */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBetaParameter:"), value)
}


// The gamma tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclayernormalizationlayer/gamma
func (c_ CLayerNormalizationLayer) Gamma() IMLCTensor {
	rv := objc.Send[CTensor](c_.ID, objc.Sel("gamma"))
	return rv
}


// The gamma tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclayernormalizationlayer/gamma
func (c_ CLayerNormalizationLayer) SetGamma(value IMLCTensor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGamma:"), value)
}


// The gamma tensor parameter you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclayernormalizationlayer/gammaparameter
func (c_ CLayerNormalizationLayer) GammaParameter() objc.IObject /* cross-framework: CTensorParameter */ {
	rv := objc.Send[CTensorParameter](c_.ID, objc.Sel("gammaParameter"))
	return rv
}


// The gamma tensor parameter you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclayernormalizationlayer/gammaparameter
func (c_ CLayerNormalizationLayer) SetGammaParameter(value objc.IObject /* cross-framework: CTensorParameter */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGammaParameter:"), value)
}


// The shape of the axes where normalization occurs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclayernormalizationlayer/normalizedshape-8ujvv
func (c_ CLayerNormalizationLayer) NormalizedShape() int {
	rv := objc.Send[int](c_.ID, objc.Sel("normalizedShape"))
	return rv
}


// The shape of the axes where normalization occurs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclayernormalizationlayer/normalizedshape-8ujvv
func (c_ CLayerNormalizationLayer) SetNormalizedShape(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNormalizedShape:"), value)
}


// The variance epsilon you use for numerical stability.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclayernormalizationlayer/varianceepsilon
func (c_ CLayerNormalizationLayer) VarianceEpsilon() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("varianceEpsilon"))
	return rv
}


// The variance epsilon you use for numerical stability.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclayernormalizationlayer/varianceepsilon
func (c_ CLayerNormalizationLayer) SetVarianceEpsilon(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVarianceEpsilon:"), value)
}



