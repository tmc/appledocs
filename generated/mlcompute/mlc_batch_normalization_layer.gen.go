// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MLCBatchNormalizationLayer */


/* debug [class_header]: Header for MLCBatchNormalizationLayer */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CBatchNormalizationLayer */
// An interface definition for the [CBatchNormalizationLayer] class.
type ICBatchNormalizationLayer interface {
	ICLayer
	
/* debug [class_interface_properties]: Properties for CBatchNormalizationLayer */
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

	
/* debug [class_interface_methods]: Methods for CBatchNormalizationLayer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CBatchNormalizationLayer */
// Alloc allocates a new instance without initialization.
func (cc _CBatchNormalizationLayerClass) Alloc() CBatchNormalizationLayer {
	rv := objc.Send[CBatchNormalizationLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CBatchNormalizationLayer */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CBatchNormalizationLayer */

// Creates a batch normalization layer with the number of feature channels, tensors, and variance epsilon you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCBatchNormalizationLayer/init(featureChannelCount:mean:variance:beta:gamma:varianceEpsilon:)
func NewCBatchNormalizationLayerWithFeatureChannelCountMeanVarianceBetaGammaVarianceEpsilon(featureChannelCount uint, mean IMLCTensor, variance IMLCTensor, beta IMLCTensor, gamma IMLCTensor, varianceEpsilon float32) CBatchNormalizationLayer {
	rv := objc.Send[CBatchNormalizationLayer](objc.ID(getCBatchNormalizationLayerClass().class), objc.Sel("layerWithFeatureChannelCount:mean:variance:beta:gamma:varianceEpsilon:"), featureChannelCount, mean, variance, beta, gamma, varianceEpsilon)
	return rv
}/* debug [class_init_methods/constructor]: NewCBatchNormalizationLayerWithFeatureChannelCountMeanVarianceBetaGammaVarianceEpsilon */


// Creates a batch normalization layer with the number of feature channels, tensors, variance epsilon, and momentum you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCBatchNormalizationLayer/init(featureChannelCount:mean:variance:beta:gamma:varianceEpsilon:momentum:)
func NewCBatchNormalizationLayerWithFeatureChannelCountMeanVarianceBetaGammaVarianceEpsilonMomentum(featureChannelCount uint, mean IMLCTensor, variance IMLCTensor, beta IMLCTensor, gamma IMLCTensor, varianceEpsilon float32, momentum float32) CBatchNormalizationLayer {
	rv := objc.Send[CBatchNormalizationLayer](objc.ID(getCBatchNormalizationLayerClass().class), objc.Sel("layerWithFeatureChannelCount:mean:variance:beta:gamma:varianceEpsilon:momentum:"), featureChannelCount, mean, variance, beta, gamma, varianceEpsilon, momentum)
	return rv
}/* debug [class_init_methods/constructor]: NewCBatchNormalizationLayerWithFeatureChannelCountMeanVarianceBetaGammaVarianceEpsilonMomentum */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CBatchNormalizationLayer */

// Creates a batch normalization layer with the number of feature channels, tensors, and variance epsilon you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCBatchNormalizationLayer/init(featureChannelCount:mean:variance:beta:gamma:varianceEpsilon:)
func (cc _CBatchNormalizationLayerClass) LayerWithFeatureChannelCountMeanVarianceBetaGammaVarianceEpsilon(featureChannelCount uint, mean IMLCTensor, variance IMLCTensor, beta IMLCTensor, gamma IMLCTensor, varianceEpsilon float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("layerWithFeatureChannelCount:mean:variance:beta:gamma:varianceEpsilon:"), featureChannelCount, mean, variance, beta, gamma, varianceEpsilon)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayerWithFeatureChannelCountMeanVarianceBetaGammaVarianceEpsilon) */


// Creates a batch normalization layer with the number of feature channels, tensors, variance epsilon, and momentum you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCBatchNormalizationLayer/init(featureChannelCount:mean:variance:beta:gamma:varianceEpsilon:momentum:)
func (cc _CBatchNormalizationLayerClass) LayerWithFeatureChannelCountMeanVarianceBetaGammaVarianceEpsilonMomentum(featureChannelCount uint, mean IMLCTensor, variance IMLCTensor, beta IMLCTensor, gamma IMLCTensor, varianceEpsilon float32, momentum float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("layerWithFeatureChannelCount:mean:variance:beta:gamma:varianceEpsilon:momentum:"), featureChannelCount, mean, variance, beta, gamma, varianceEpsilon, momentum)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayerWithFeatureChannelCountMeanVarianceBetaGammaVarianceEpsilonMomentum) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CBatchNormalizationLayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CBatchNormalizationLayer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CBatchNormalizationLayer */

// The beta tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCBatchNormalizationLayer/beta
func (c_ CBatchNormalizationLayer) Beta() IMLCTensor {
	rv := objc.Send[CTensor](c_.ID, objc.Sel("beta"))
	return rv
}/* debug [instance_properties/getter]: beta */


// The beta tensor parameter you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCBatchNormalizationLayer/betaParameter
func (c_ CBatchNormalizationLayer) BetaParameter() IMLCTensorParameter {
	rv := objc.Send[CTensorParameter](c_.ID, objc.Sel("betaParameter"))
	return rv
}/* debug [instance_properties/getter]: betaParameter */


// The number of feature channels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCBatchNormalizationLayer/featureChannelCount
func (c_ CBatchNormalizationLayer) FeatureChannelCount() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("featureChannelCount"))
	return rv
}/* debug [instance_properties/getter]: featureChannelCount */


// The gamma tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCBatchNormalizationLayer/gamma
func (c_ CBatchNormalizationLayer) Gamma() IMLCTensor {
	rv := objc.Send[CTensor](c_.ID, objc.Sel("gamma"))
	return rv
}/* debug [instance_properties/getter]: gamma */


// The gamma tensor parameter you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCBatchNormalizationLayer/gammaParameter
func (c_ CBatchNormalizationLayer) GammaParameter() IMLCTensorParameter {
	rv := objc.Send[CTensorParameter](c_.ID, objc.Sel("gammaParameter"))
	return rv
}/* debug [instance_properties/getter]: gammaParameter */


// The mean tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCBatchNormalizationLayer/mean
func (c_ CBatchNormalizationLayer) Mean() IMLCTensor {
	rv := objc.Send[CTensor](c_.ID, objc.Sel("mean"))
	return rv
}/* debug [instance_properties/getter]: mean */


// The value you use for the running mean and variance computation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCBatchNormalizationLayer/momentum
func (c_ CBatchNormalizationLayer) Momentum() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("momentum"))
	return rv
}/* debug [instance_properties/getter]: momentum */


// The variance tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCBatchNormalizationLayer/variance
func (c_ CBatchNormalizationLayer) Variance() IMLCTensor {
	rv := objc.Send[CTensor](c_.ID, objc.Sel("variance"))
	return rv
}/* debug [instance_properties/getter]: variance */


// The variance epsilon you use for numerical stability.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCBatchNormalizationLayer/varianceEpsilon
func (c_ CBatchNormalizationLayer) VarianceEpsilon() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("varianceEpsilon"))
	return rv
}/* debug [instance_properties/getter]: varianceEpsilon */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCBatchNormalizationLayer */


