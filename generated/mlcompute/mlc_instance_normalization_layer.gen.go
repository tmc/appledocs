// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MLCInstanceNormalizationLayer */


/* debug [class_header]: Header for MLCInstanceNormalizationLayer */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CInstanceNormalizationLayer */
// An interface definition for the [CInstanceNormalizationLayer] class.
type ICInstanceNormalizationLayer interface {
	ICLayer
	
/* debug [class_interface_properties]: Properties for CInstanceNormalizationLayer */
	// properties:
	Beta() IMLCTensor
	BetaParameter() IMLCTensorParameter
	FeatureChannelCount() uint
	Gamma() IMLCTensor
	GammaParameter() IMLCTensorParameter
	Mean() IMLCTensor
	Momentum() float32
	Variance() IMLCTensor
	VarianceEpsilon() float32
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CInstanceNormalizationLayer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CInstanceNormalizationLayer */
// Alloc allocates a new instance without initialization.
func (cc _CInstanceNormalizationLayerClass) Alloc() CInstanceNormalizationLayer {
	rv := objc.Send[CInstanceNormalizationLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CInstanceNormalizationLayer */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CInstanceNormalizationLayer */

// Creates an instance normalization layer with the number of feature channels, beta and gamma tensors, and variance epsilon you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCInstanceNormalizationLayer/init(featureChannelCount:beta:gamma:varianceEpsilon:)
func NewCInstanceNormalizationLayerWithFeatureChannelCountBetaGammaVarianceEpsilon(featureChannelCount uint, beta IMLCTensor, gamma IMLCTensor, varianceEpsilon float32) CInstanceNormalizationLayer {
	rv := objc.Send[CInstanceNormalizationLayer](objc.ID(getCInstanceNormalizationLayerClass().class), objc.Sel("layerWithFeatureChannelCount:beta:gamma:varianceEpsilon:"), featureChannelCount, beta, gamma, varianceEpsilon)
	return rv
}/* debug [class_init_methods/constructor]: NewCInstanceNormalizationLayerWithFeatureChannelCountBetaGammaVarianceEpsilon */


// Creates an instance normalization layer with the number of feature channels, beta and gamma tensors, variance epsilon, and momentum you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCInstanceNormalizationLayer/init(featureChannelCount:beta:gamma:varianceEpsilon:momentum:)
func NewCInstanceNormalizationLayerWithFeatureChannelCountBetaGammaVarianceEpsilonMomentum(featureChannelCount uint, beta IMLCTensor, gamma IMLCTensor, varianceEpsilon float32, momentum float32) CInstanceNormalizationLayer {
	rv := objc.Send[CInstanceNormalizationLayer](objc.ID(getCInstanceNormalizationLayerClass().class), objc.Sel("layerWithFeatureChannelCount:beta:gamma:varianceEpsilon:momentum:"), featureChannelCount, beta, gamma, varianceEpsilon, momentum)
	return rv
}/* debug [class_init_methods/constructor]: NewCInstanceNormalizationLayerWithFeatureChannelCountBetaGammaVarianceEpsilonMomentum */


// Creates an instance normalization layer with the number of feature channels, mean, variance, beta and gamma tensors, variance epsilon, and momentum you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCInstanceNormalizationLayer/init(featureChannelCount:mean:variance:beta:gamma:varianceEpsilon:momentum:)
func NewCInstanceNormalizationLayerWithFeatureChannelCountMeanVarianceBetaGammaVarianceEpsilonMomentum(featureChannelCount uint, mean IMLCTensor, variance IMLCTensor, beta IMLCTensor, gamma IMLCTensor, varianceEpsilon float32, momentum float32) CInstanceNormalizationLayer {
	rv := objc.Send[CInstanceNormalizationLayer](objc.ID(getCInstanceNormalizationLayerClass().class), objc.Sel("layerWithFeatureChannelCount:mean:variance:beta:gamma:varianceEpsilon:momentum:"), featureChannelCount, mean, variance, beta, gamma, varianceEpsilon, momentum)
	return rv
}/* debug [class_init_methods/constructor]: NewCInstanceNormalizationLayerWithFeatureChannelCountMeanVarianceBetaGammaVarianceEpsilonMomentum */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CInstanceNormalizationLayer */

// Creates an instance normalization layer with the number of feature channels, beta and gamma tensors, and variance epsilon you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCInstanceNormalizationLayer/init(featureChannelCount:beta:gamma:varianceEpsilon:)
func (cc _CInstanceNormalizationLayerClass) LayerWithFeatureChannelCountBetaGammaVarianceEpsilon(featureChannelCount uint, beta IMLCTensor, gamma IMLCTensor, varianceEpsilon float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("layerWithFeatureChannelCount:beta:gamma:varianceEpsilon:"), featureChannelCount, beta, gamma, varianceEpsilon)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayerWithFeatureChannelCountBetaGammaVarianceEpsilon) */


// Creates an instance normalization layer with the number of feature channels, beta and gamma tensors, variance epsilon, and momentum you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCInstanceNormalizationLayer/init(featureChannelCount:beta:gamma:varianceEpsilon:momentum:)
func (cc _CInstanceNormalizationLayerClass) LayerWithFeatureChannelCountBetaGammaVarianceEpsilonMomentum(featureChannelCount uint, beta IMLCTensor, gamma IMLCTensor, varianceEpsilon float32, momentum float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("layerWithFeatureChannelCount:beta:gamma:varianceEpsilon:momentum:"), featureChannelCount, beta, gamma, varianceEpsilon, momentum)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayerWithFeatureChannelCountBetaGammaVarianceEpsilonMomentum) */


// Creates an instance normalization layer with the number of feature channels, mean, variance, beta and gamma tensors, variance epsilon, and momentum you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCInstanceNormalizationLayer/init(featureChannelCount:mean:variance:beta:gamma:varianceEpsilon:momentum:)
func (cc _CInstanceNormalizationLayerClass) LayerWithFeatureChannelCountMeanVarianceBetaGammaVarianceEpsilonMomentum(featureChannelCount uint, mean IMLCTensor, variance IMLCTensor, beta IMLCTensor, gamma IMLCTensor, varianceEpsilon float32, momentum float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("layerWithFeatureChannelCount:mean:variance:beta:gamma:varianceEpsilon:momentum:"), featureChannelCount, mean, variance, beta, gamma, varianceEpsilon, momentum)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayerWithFeatureChannelCountMeanVarianceBetaGammaVarianceEpsilonMomentum) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CInstanceNormalizationLayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CInstanceNormalizationLayer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CInstanceNormalizationLayer */

// The beta tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCInstanceNormalizationLayer/beta
func (c_ CInstanceNormalizationLayer) Beta() IMLCTensor {
	rv := objc.Send[CTensor](c_.ID, objc.Sel("beta"))
	return rv
}/* debug [instance_properties/getter]: beta */


// The beta tensor parameter you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCInstanceNormalizationLayer/betaParameter
func (c_ CInstanceNormalizationLayer) BetaParameter() IMLCTensorParameter {
	rv := objc.Send[CTensorParameter](c_.ID, objc.Sel("betaParameter"))
	return rv
}/* debug [instance_properties/getter]: betaParameter */


// The number of feature channels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCInstanceNormalizationLayer/featureChannelCount
func (c_ CInstanceNormalizationLayer) FeatureChannelCount() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("featureChannelCount"))
	return rv
}/* debug [instance_properties/getter]: featureChannelCount */


// The gamma tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCInstanceNormalizationLayer/gamma
func (c_ CInstanceNormalizationLayer) Gamma() IMLCTensor {
	rv := objc.Send[CTensor](c_.ID, objc.Sel("gamma"))
	return rv
}/* debug [instance_properties/getter]: gamma */


// The gamma tensor parameter you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCInstanceNormalizationLayer/gammaParameter
func (c_ CInstanceNormalizationLayer) GammaParameter() IMLCTensorParameter {
	rv := objc.Send[CTensorParameter](c_.ID, objc.Sel("gammaParameter"))
	return rv
}/* debug [instance_properties/getter]: gammaParameter */


// The running mean tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCInstanceNormalizationLayer/mean
func (c_ CInstanceNormalizationLayer) Mean() IMLCTensor {
	rv := objc.Send[CTensor](c_.ID, objc.Sel("mean"))
	return rv
}/* debug [instance_properties/getter]: mean */


// The momentum value for the running mean and variance computation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCInstanceNormalizationLayer/momentum
func (c_ CInstanceNormalizationLayer) Momentum() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("momentum"))
	return rv
}/* debug [instance_properties/getter]: momentum */


// The running variance tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCInstanceNormalizationLayer/variance
func (c_ CInstanceNormalizationLayer) Variance() IMLCTensor {
	rv := objc.Send[CTensor](c_.ID, objc.Sel("variance"))
	return rv
}/* debug [instance_properties/getter]: variance */


// The variance epsilon you use for numerical stability.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCInstanceNormalizationLayer/varianceEpsilon
func (c_ CInstanceNormalizationLayer) VarianceEpsilon() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("varianceEpsilon"))
	return rv
}/* debug [instance_properties/getter]: varianceEpsilon */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCInstanceNormalizationLayer */


