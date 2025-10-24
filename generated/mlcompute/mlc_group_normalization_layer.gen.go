// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MLCGroupNormalizationLayer */


/* debug [class_header]: Header for MLCGroupNormalizationLayer */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CGroupNormalizationLayer */
// An interface definition for the [CGroupNormalizationLayer] class.
type ICGroupNormalizationLayer interface {
	ICLayer
	
/* debug [class_interface_properties]: Properties for CGroupNormalizationLayer */
	// properties:
	Beta() IMLCTensor
	BetaParameter() IMLCTensorParameter
	FeatureChannelCount() uint
	Gamma() IMLCTensor
	GammaParameter() IMLCTensorParameter
	GroupCount() uint
	VarianceEpsilon() float32
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CGroupNormalizationLayer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CGroupNormalizationLayer */
// Alloc allocates a new instance without initialization.
func (cc _CGroupNormalizationLayerClass) Alloc() CGroupNormalizationLayer {
	rv := objc.Send[CGroupNormalizationLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CGroupNormalizationLayer */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CGroupNormalizationLayer */

// Creates a group normalization layer with the number of feature channels and groups, beta and gamma tensors, and variance epsilon you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGroupNormalizationLayer/init(featureChannelCount:groupCount:beta:gamma:varianceEpsilon:)
func NewCGroupNormalizationLayerWithFeatureChannelCountGroupCountBetaGammaVarianceEpsilon(featureChannelCount uint, groupCount uint, beta IMLCTensor, gamma IMLCTensor, varianceEpsilon float32) CGroupNormalizationLayer {
	rv := objc.Send[CGroupNormalizationLayer](objc.ID(getCGroupNormalizationLayerClass().class), objc.Sel("layerWithFeatureChannelCount:groupCount:beta:gamma:varianceEpsilon:"), featureChannelCount, groupCount, beta, gamma, varianceEpsilon)
	return rv
}/* debug [class_init_methods/constructor]: NewCGroupNormalizationLayerWithFeatureChannelCountGroupCountBetaGammaVarianceEpsilon */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CGroupNormalizationLayer */

// Creates a group normalization layer with the number of feature channels and groups, beta and gamma tensors, and variance epsilon you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGroupNormalizationLayer/init(featureChannelCount:groupCount:beta:gamma:varianceEpsilon:)
func (cc _CGroupNormalizationLayerClass) LayerWithFeatureChannelCountGroupCountBetaGammaVarianceEpsilon(featureChannelCount uint, groupCount uint, beta IMLCTensor, gamma IMLCTensor, varianceEpsilon float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("layerWithFeatureChannelCount:groupCount:beta:gamma:varianceEpsilon:"), featureChannelCount, groupCount, beta, gamma, varianceEpsilon)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayerWithFeatureChannelCountGroupCountBetaGammaVarianceEpsilon) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CGroupNormalizationLayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CGroupNormalizationLayer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CGroupNormalizationLayer */

// The beta tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGroupNormalizationLayer/beta
func (c_ CGroupNormalizationLayer) Beta() IMLCTensor {
	rv := objc.Send[CTensor](c_.ID, objc.Sel("beta"))
	return rv
}/* debug [instance_properties/getter]: beta */


// The beta tensor parameter you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGroupNormalizationLayer/betaParameter
func (c_ CGroupNormalizationLayer) BetaParameter() IMLCTensorParameter {
	rv := objc.Send[CTensorParameter](c_.ID, objc.Sel("betaParameter"))
	return rv
}/* debug [instance_properties/getter]: betaParameter */


// The number of feature channels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGroupNormalizationLayer/featureChannelCount
func (c_ CGroupNormalizationLayer) FeatureChannelCount() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("featureChannelCount"))
	return rv
}/* debug [instance_properties/getter]: featureChannelCount */


// The gamma tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGroupNormalizationLayer/gamma
func (c_ CGroupNormalizationLayer) Gamma() IMLCTensor {
	rv := objc.Send[CTensor](c_.ID, objc.Sel("gamma"))
	return rv
}/* debug [instance_properties/getter]: gamma */


// The gamma tensor parameter you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGroupNormalizationLayer/gammaParameter
func (c_ CGroupNormalizationLayer) GammaParameter() IMLCTensorParameter {
	rv := objc.Send[CTensorParameter](c_.ID, objc.Sel("gammaParameter"))
	return rv
}/* debug [instance_properties/getter]: gammaParameter */


// The number of groups into which you separate the channels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGroupNormalizationLayer/groupCount
func (c_ CGroupNormalizationLayer) GroupCount() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("groupCount"))
	return rv
}/* debug [instance_properties/getter]: groupCount */


// The variance epsilon you use for numerical stability.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGroupNormalizationLayer/varianceEpsilon
func (c_ CGroupNormalizationLayer) VarianceEpsilon() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("varianceEpsilon"))
	return rv
}/* debug [instance_properties/getter]: varianceEpsilon */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCGroupNormalizationLayer */


