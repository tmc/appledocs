// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MLCLayerNormalizationLayer */


/* debug [class_header]: Header for MLCLayerNormalizationLayer */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CLayerNormalizationLayer */
// An interface definition for the [CLayerNormalizationLayer] class.
type ICLayerNormalizationLayer interface {
	ICLayer
	
/* debug [class_interface_properties]: Properties for CLayerNormalizationLayer */
	// properties:
	Beta() IMLCTensor
	BetaParameter() IMLCTensorParameter
	Gamma() IMLCTensor
	GammaParameter() IMLCTensorParameter
	NormalizedShape() []foundation.Number
	VarianceEpsilon() float32
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CLayerNormalizationLayer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CLayerNormalizationLayer */
// Alloc allocates a new instance without initialization.
func (cc _CLayerNormalizationLayerClass) Alloc() CLayerNormalizationLayer {
	rv := objc.Send[CLayerNormalizationLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CLayerNormalizationLayer */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CLayerNormalizationLayer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CLayerNormalizationLayer */

// Creates a normalization layer with a shape, beta and gamma tensors, and variance epsilon you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLayerNormalizationLayer/layerWithNormalizedShape:beta:gamma:varianceEpsilon:
func (cc _CLayerNormalizationLayerClass) LayerWithNormalizedShapeBetaGammaVarianceEpsilon(normalizedShape []foundation.Number, beta IMLCTensor, gamma IMLCTensor, varianceEpsilon float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("layerWithNormalizedShape:beta:gamma:varianceEpsilon:"), normalizedShape, beta, gamma, varianceEpsilon)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayerWithNormalizedShapeBetaGammaVarianceEpsilon) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CLayerNormalizationLayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CLayerNormalizationLayer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CLayerNormalizationLayer */

// The beta tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLayerNormalizationLayer/beta
func (c_ CLayerNormalizationLayer) Beta() IMLCTensor {
	rv := objc.Send[CTensor](c_.ID, objc.Sel("beta"))
	return rv
}/* debug [instance_properties/getter]: beta */


// The beta tensor parameter you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLayerNormalizationLayer/betaParameter
func (c_ CLayerNormalizationLayer) BetaParameter() IMLCTensorParameter {
	rv := objc.Send[CTensorParameter](c_.ID, objc.Sel("betaParameter"))
	return rv
}/* debug [instance_properties/getter]: betaParameter */


// The gamma tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLayerNormalizationLayer/gamma
func (c_ CLayerNormalizationLayer) Gamma() IMLCTensor {
	rv := objc.Send[CTensor](c_.ID, objc.Sel("gamma"))
	return rv
}/* debug [instance_properties/getter]: gamma */


// The gamma tensor parameter you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLayerNormalizationLayer/gammaParameter
func (c_ CLayerNormalizationLayer) GammaParameter() IMLCTensorParameter {
	rv := objc.Send[CTensorParameter](c_.ID, objc.Sel("gammaParameter"))
	return rv
}/* debug [instance_properties/getter]: gammaParameter */


// The shape of the axes where normalization occurs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLayerNormalizationLayer/normalizedShape-2cz6k
func (c_ CLayerNormalizationLayer) NormalizedShape() []foundation.Number {
	rv := objc.Send[[]foundation.Number](c_.ID, objc.Sel("normalizedShape"))
	return rv
}/* debug [instance_properties/getter]: normalizedShape */


// The variance epsilon you use for numerical stability.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLayerNormalizationLayer/varianceEpsilon
func (c_ CLayerNormalizationLayer) VarianceEpsilon() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("varianceEpsilon"))
	return rv
}/* debug [instance_properties/getter]: varianceEpsilon */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCLayerNormalizationLayer */



