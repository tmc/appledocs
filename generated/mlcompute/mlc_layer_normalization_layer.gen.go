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
	Beta() MLCTensor
	SetBeta(value IMLCTensor)
	BetaParameter() MLCTensorParameter
	SetBetaParameter(value IMLCTensorParameter)
	Gamma() MLCTensor
	SetGamma(value IMLCTensor)
	GammaParameter() MLCTensorParameter
	SetGammaParameter(value IMLCTensorParameter)
	NormalizedShape() int
	SetNormalizedShape(value int)
	VarianceEpsilon() float32
	SetVarianceEpsilon(value float32)
}

// A layer that applies layer normalization over inputs.
//
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
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclayernormalizationlayer/beta
func (c_ CLayerNormalizationLayer) Beta() MLCTensor {
	rv := objc.Send[MLCTensor](c_.ID, objc.Sel("beta"))
	return rv
}


// SetBeta sets the value of the beta property.
// The beta tensor.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclayernormalizationlayer/beta
func (c_ CLayerNormalizationLayer) SetBeta(value IMLCTensor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBeta:"), value)
}

// The beta tensor parameter you use for optimizer updates.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclayernormalizationlayer/betaparameter
func (c_ CLayerNormalizationLayer) BetaParameter() MLCTensorParameter {
	rv := objc.Send[MLCTensorParameter](c_.ID, objc.Sel("betaParameter"))
	return rv
}


// SetBetaParameter sets the value of the betaParameter property.
// The beta tensor parameter you use for optimizer updates.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclayernormalizationlayer/betaparameter
func (c_ CLayerNormalizationLayer) SetBetaParameter(value IMLCTensorParameter) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBetaParameter:"), value)
}

// The gamma tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclayernormalizationlayer/gamma
func (c_ CLayerNormalizationLayer) Gamma() MLCTensor {
	rv := objc.Send[MLCTensor](c_.ID, objc.Sel("gamma"))
	return rv
}


// SetGamma sets the value of the gamma property.
// The gamma tensor.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclayernormalizationlayer/gamma
func (c_ CLayerNormalizationLayer) SetGamma(value IMLCTensor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGamma:"), value)
}

// The gamma tensor parameter you use for optimizer updates.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclayernormalizationlayer/gammaparameter
func (c_ CLayerNormalizationLayer) GammaParameter() MLCTensorParameter {
	rv := objc.Send[MLCTensorParameter](c_.ID, objc.Sel("gammaParameter"))
	return rv
}


// SetGammaParameter sets the value of the gammaParameter property.
// The gamma tensor parameter you use for optimizer updates.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclayernormalizationlayer/gammaparameter
func (c_ CLayerNormalizationLayer) SetGammaParameter(value IMLCTensorParameter) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGammaParameter:"), value)
}

// The shape of the axes where normalization occurs.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclayernormalizationlayer/normalizedshape-8ujvv
func (c_ CLayerNormalizationLayer) NormalizedShape() int {
	rv := objc.Send[int](c_.ID, objc.Sel("normalizedShape"))
	return rv
}


// SetNormalizedShape sets the value of the normalizedShape property.
// The shape of the axes where normalization occurs.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclayernormalizationlayer/normalizedshape-8ujvv
func (c_ CLayerNormalizationLayer) SetNormalizedShape(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNormalizedShape:"), value)
}

// The variance epsilon you use for numerical stability.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclayernormalizationlayer/varianceepsilon
func (c_ CLayerNormalizationLayer) VarianceEpsilon() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("varianceEpsilon"))
	return rv
}


// SetVarianceEpsilon sets the value of the varianceEpsilon property.
// The variance epsilon you use for numerical stability.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclayernormalizationlayer/varianceepsilon
func (c_ CLayerNormalizationLayer) SetVarianceEpsilon(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVarianceEpsilon:"), value)
}



