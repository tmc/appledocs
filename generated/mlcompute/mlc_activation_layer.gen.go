// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MLCActivationLayer */


/* debug [class_header]: Header for MLCActivationLayer */
// The class instance for the [CActivationLayer] class.
var (
	CActivationLayerClass     _CActivationLayerClass
	CActivationLayerClassOnce sync.Once
)

func getCActivationLayerClass() _CActivationLayerClass {
	CActivationLayerClassOnce.Do(func() {
		CActivationLayerClass = _CActivationLayerClass{objc.GetClass("MLCActivationLayer")}
	})
	return CActivationLayerClass
}

type _CActivationLayerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CActivationLayer */
// An interface definition for the [CActivationLayer] class.
type ICActivationLayer interface {
	ICLayer
	
/* debug [class_interface_properties]: Properties for CActivationLayer */
	// properties:
	Descriptor() IMLCActivationDescriptor
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CActivationLayer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CActivationLayer */
// Alloc allocates a new instance without initialization.
func (cc _CActivationLayerClass) Alloc() CActivationLayer {
	rv := objc.Send[CActivationLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CActivationLayerClass) New() CActivationLayer {
	rv := objc.Send[CActivationLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CActivationLayer) Init() CActivationLayer {
	rv := objc.Send[CActivationLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CActivationLayer) Autorelease() CActivationLayer {
	rv := objc.Send[CActivationLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCActivationLayer creates a new CActivationLayer instance.
func NewCActivationLayer() CActivationLayer {
	return getCActivationLayerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CActivationLayer */
// A layer that applies an activation function to the source tensor and produces an output.
//
// To construct an activation layer, create an activation descriptor and then pass it to the initializer.


// A layer that applies an activation function to the source tensor and produces an output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer
type CActivationLayer struct {
	CLayer
}

// CActivationLayerFrom constructs a [CActivationLayer] from an unsafe.Pointer.
//
// A layer that applies an activation function to the source tensor and produces an output.
func CActivationLayerFrom(ptr unsafe.Pointer) CActivationLayer {
	return CActivationLayer{
		CLayer: CLayerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CActivationLayer */

// Creates an activation layer with the descriptor you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/init(descriptor:)
func NewCActivationLayerWithDescriptor(descriptor IMLCActivationDescriptor) CActivationLayer {
	rv := objc.Send[CActivationLayer](objc.ID(getCActivationLayerClass().class), objc.Sel("layerWithDescriptor:"), descriptor)
	return rv
}/* debug [class_init_methods/constructor]: NewCActivationLayerWithDescriptor */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CActivationLayer */

// Creates an instance of a CELU activation layer using the alpha value you specify for the CELU formation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/celu(a:)
func (cc _CActivationLayerClass) CeluLayerWithA(a float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("celuLayerWithA:"), a)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CeluLayerWithA) */


// Creates an instance of a clamp activation layer using the minimum and maximum values you specify for the clamp formation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/clamp(min:max:)
func (cc _CActivationLayerClass) ClampLayerWithMinValueMaxValue(minValue float32, maxValue float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("clampLayerWithMinValue:maxValue:"), minValue, maxValue)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ClampLayerWithMinValueMaxValue) */


// Creates an instance of an ELU activation layer using the alpha value you specify for the ELU formation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/elu(a:)
func (cc _CActivationLayerClass) EluLayerWithA(a float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("eluLayerWithA:"), a)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=EluLayerWithA) */


// Creates an instance of a hard shrink activation layer using the lambda value you specify for the hard shrink formation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/hardShrink(a:)
func (cc _CActivationLayerClass) HardShrinkLayerWithA(a float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("hardShrinkLayerWithA:"), a)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=HardShrinkLayerWithA) */


// Creates an activation layer with the descriptor you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/init(descriptor:)
func (cc _CActivationLayerClass) LayerWithDescriptor(descriptor IMLCActivationDescriptor) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("layerWithDescriptor:"), descriptor)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayerWithDescriptor) */


// Creates an instance of a leaky ReLU activation layer using the angle of the negative slope you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/leakyReLU(negativeSlope:)
func (cc _CActivationLayerClass) LeakyReLULayerWithNegativeSlope(negativeSlope float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("leakyReLULayerWithNegativeSlope:"), negativeSlope)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LeakyReLULayerWithNegativeSlope) */


// Creates an instance of a linear activation layer using the scale factor and bias value you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/linear(scale:bias:)
func (cc _CActivationLayerClass) LinearLayerWithScaleBias(scale float32, bias float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("linearLayerWithScale:bias:"), scale, bias)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LinearLayerWithScaleBias) */


// Creates an instance of a ReLUN activation layer using the alpha and beta values you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/relun(a:b:)
func (cc _CActivationLayerClass) RelunLayerWithAB(a float32, b float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("relunLayerWithA:b:"), a, b)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RelunLayerWithAB) */


// Creates an instance of a soft plus activation layer using the beta value you specify for the soft plus formation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/softPlus(beta:)
func (cc _CActivationLayerClass) SoftPlusLayerWithBeta(beta float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("softPlusLayerWithBeta:"), beta)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SoftPlusLayerWithBeta) */


// Creates an instance of a soft shrink activation layer using the lambda value you specify for the soft shrink formation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/softShrink(a:)
func (cc _CActivationLayerClass) SoftShrinkLayerWithA(a float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("softShrinkLayerWithA:"), a)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SoftShrinkLayerWithA) */


// Creates an instance of a threshold activation layer using the threshold and replacement values you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/threshold(_:replacement:)
func (cc _CActivationLayerClass) ThresholdLayerWithThresholdReplacement(threshold float32, replacement float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("thresholdLayerWithThreshold:replacement:"), threshold, replacement)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ThresholdLayerWithThresholdReplacement) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CActivationLayer */

// Creates an instance of an absolute activation layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/absolute
func (cc _CActivationLayerClass) AbsoluteLayer() CActivationLayer {
	rv := objc.Send[CActivationLayer](objc.ID(cc.class), objc.Sel("absoluteLayer"))
	return rv
}/* debug [class_properties_class/property]: absoluteLayer */

// Creates an instance of a CELU activation layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/celu
func (cc _CActivationLayerClass) CeluLayer() CActivationLayer {
	rv := objc.Send[CActivationLayer](objc.ID(cc.class), objc.Sel("celuLayer"))
	return rv
}/* debug [class_properties_class/property]: celuLayer */

// Creates an instance of a parametric ELU activation layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/elu
func (cc _CActivationLayerClass) EluLayer() CActivationLayer {
	rv := objc.Send[CActivationLayer](objc.ID(cc.class), objc.Sel("eluLayer"))
	return rv
}/* debug [class_properties_class/property]: eluLayer */

// Creates an instance of a GELU activation layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/gelu
func (cc _CActivationLayerClass) GeluLayer() CActivationLayer {
	rv := objc.Send[CActivationLayer](objc.ID(cc.class), objc.Sel("geluLayer"))
	return rv
}/* debug [class_properties_class/property]: geluLayer */

// Creates an instance of a hard shrink activation layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/hardShrink
func (cc _CActivationLayerClass) HardShrinkLayer() CActivationLayer {
	rv := objc.Send[CActivationLayer](objc.ID(cc.class), objc.Sel("hardShrinkLayer"))
	return rv
}/* debug [class_properties_class/property]: hardShrinkLayer */

// Creates an instance of a hard sigmoid activation layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/hardSigmoid
func (cc _CActivationLayerClass) HardSigmoidLayer() CActivationLayer {
	rv := objc.Send[CActivationLayer](objc.ID(cc.class), objc.Sel("hardSigmoidLayer"))
	return rv
}/* debug [class_properties_class/property]: hardSigmoidLayer */

// Creates an instance of a hard swish activation layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/hardSwish
func (cc _CActivationLayerClass) HardSwishLayer() CActivationLayer {
	rv := objc.Send[CActivationLayer](objc.ID(cc.class), objc.Sel("hardSwishLayer"))
	return rv
}/* debug [class_properties_class/property]: hardSwishLayer */

// Creates an instance of a leaky ReLU activation layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/leakyReLU
func (cc _CActivationLayerClass) LeakyReLULayer() CActivationLayer {
	rv := objc.Send[CActivationLayer](objc.ID(cc.class), objc.Sel("leakyReLULayer"))
	return rv
}/* debug [class_properties_class/property]: leakyReLULayer */

// Creates an instance of a log sigmoid activation layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/logSigmoid
func (cc _CActivationLayerClass) LogSigmoidLayer() CActivationLayer {
	rv := objc.Send[CActivationLayer](objc.ID(cc.class), objc.Sel("logSigmoidLayer"))
	return rv
}/* debug [class_properties_class/property]: logSigmoidLayer */

// Creates an instance of a ReLU activation layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/relu
func (cc _CActivationLayerClass) ReluLayer() CActivationLayer {
	rv := objc.Send[CActivationLayer](objc.ID(cc.class), objc.Sel("reluLayer"))
	return rv
}/* debug [class_properties_class/property]: reluLayer */

// Creates an instance of a ReLU6 activation layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/relu6
func (cc _CActivationLayerClass) Relu6Layer() CActivationLayer {
	rv := objc.Send[CActivationLayer](objc.ID(cc.class), objc.Sel("relu6Layer"))
	return rv
}/* debug [class_properties_class/property]: relu6Layer */

// Creates an instance of a SELU activation layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/selu
func (cc _CActivationLayerClass) SeluLayer() CActivationLayer {
	rv := objc.Send[CActivationLayer](objc.ID(cc.class), objc.Sel("seluLayer"))
	return rv
}/* debug [class_properties_class/property]: seluLayer */

// Creates an instance of a sigmoid activation layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/sigmoid
func (cc _CActivationLayerClass) SigmoidLayer() CActivationLayer {
	rv := objc.Send[CActivationLayer](objc.ID(cc.class), objc.Sel("sigmoidLayer"))
	return rv
}/* debug [class_properties_class/property]: sigmoidLayer */

// Creates an instance of a parametric soft plus activation layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/softPlus
func (cc _CActivationLayerClass) SoftPlusLayer() CActivationLayer {
	rv := objc.Send[CActivationLayer](objc.ID(cc.class), objc.Sel("softPlusLayer"))
	return rv
}/* debug [class_properties_class/property]: softPlusLayer */

// Creates an instance of a soft shrink activation layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/softShrink
func (cc _CActivationLayerClass) SoftShrinkLayer() CActivationLayer {
	rv := objc.Send[CActivationLayer](objc.ID(cc.class), objc.Sel("softShrinkLayer"))
	return rv
}/* debug [class_properties_class/property]: softShrinkLayer */

// Creates an instance of a parametric soft sign activation layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/softSign
func (cc _CActivationLayerClass) SoftSignLayer() CActivationLayer {
	rv := objc.Send[CActivationLayer](objc.ID(cc.class), objc.Sel("softSignLayer"))
	return rv
}/* debug [class_properties_class/property]: softSignLayer */

// Creates an instance of a hyperbolic tangent activation layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/tanh
func (cc _CActivationLayerClass) TanhLayer() CActivationLayer {
	rv := objc.Send[CActivationLayer](objc.ID(cc.class), objc.Sel("tanhLayer"))
	return rv
}/* debug [class_properties_class/property]: tanhLayer */

// Creates an instance of a tanh shrink activation layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/tanhShrink
func (cc _CActivationLayerClass) TanhShrinkLayer() CActivationLayer {
	rv := objc.Send[CActivationLayer](objc.ID(cc.class), objc.Sel("tanhShrinkLayer"))
	return rv
}/* debug [class_properties_class/property]: tanhShrinkLayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CActivationLayer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CActivationLayer */

// Creates an instance of an absolute activation layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/absolute
func (c_ CActivationLayer) AbsoluteLayer() IMLCActivationLayer {
	rv := objc.Send[CActivationLayer](c_.ID, objc.Sel("absoluteLayer"))
	return rv
}/* debug [instance_properties/getter]: absoluteLayer */


// Creates an instance of a CELU activation layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/celu
func (c_ CActivationLayer) CeluLayer() IMLCActivationLayer {
	rv := objc.Send[CActivationLayer](c_.ID, objc.Sel("celuLayer"))
	return rv
}/* debug [instance_properties/getter]: celuLayer */


// The configuration object you use to create an activation layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/descriptor
func (c_ CActivationLayer) Descriptor() IMLCActivationDescriptor {
	rv := objc.Send[CActivationDescriptor](c_.ID, objc.Sel("descriptor"))
	return rv
}/* debug [instance_properties/getter]: descriptor */


// Creates an instance of a parametric ELU activation layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/elu
func (c_ CActivationLayer) EluLayer() IMLCActivationLayer {
	rv := objc.Send[CActivationLayer](c_.ID, objc.Sel("eluLayer"))
	return rv
}/* debug [instance_properties/getter]: eluLayer */


// Creates an instance of a GELU activation layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/gelu
func (c_ CActivationLayer) GeluLayer() IMLCActivationLayer {
	rv := objc.Send[CActivationLayer](c_.ID, objc.Sel("geluLayer"))
	return rv
}/* debug [instance_properties/getter]: geluLayer */


// Creates an instance of a hard shrink activation layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/hardShrink
func (c_ CActivationLayer) HardShrinkLayer() IMLCActivationLayer {
	rv := objc.Send[CActivationLayer](c_.ID, objc.Sel("hardShrinkLayer"))
	return rv
}/* debug [instance_properties/getter]: hardShrinkLayer */


// Creates an instance of a hard sigmoid activation layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/hardSigmoid
func (c_ CActivationLayer) HardSigmoidLayer() IMLCActivationLayer {
	rv := objc.Send[CActivationLayer](c_.ID, objc.Sel("hardSigmoidLayer"))
	return rv
}/* debug [instance_properties/getter]: hardSigmoidLayer */


// Creates an instance of a hard swish activation layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/hardSwish
func (c_ CActivationLayer) HardSwishLayer() IMLCActivationLayer {
	rv := objc.Send[CActivationLayer](c_.ID, objc.Sel("hardSwishLayer"))
	return rv
}/* debug [instance_properties/getter]: hardSwishLayer */


// Creates an instance of a leaky ReLU activation layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/leakyReLU
func (c_ CActivationLayer) LeakyReLULayer() IMLCActivationLayer {
	rv := objc.Send[CActivationLayer](c_.ID, objc.Sel("leakyReLULayer"))
	return rv
}/* debug [instance_properties/getter]: leakyReLULayer */


// Creates an instance of a log sigmoid activation layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/logSigmoid
func (c_ CActivationLayer) LogSigmoidLayer() IMLCActivationLayer {
	rv := objc.Send[CActivationLayer](c_.ID, objc.Sel("logSigmoidLayer"))
	return rv
}/* debug [instance_properties/getter]: logSigmoidLayer */


// Creates an instance of a ReLU activation layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/relu
func (c_ CActivationLayer) ReluLayer() IMLCActivationLayer {
	rv := objc.Send[CActivationLayer](c_.ID, objc.Sel("reluLayer"))
	return rv
}/* debug [instance_properties/getter]: reluLayer */


// Creates an instance of a ReLU6 activation layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/relu6
func (c_ CActivationLayer) Relu6Layer() IMLCActivationLayer {
	rv := objc.Send[CActivationLayer](c_.ID, objc.Sel("relu6Layer"))
	return rv
}/* debug [instance_properties/getter]: relu6Layer */


// Creates an instance of a SELU activation layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/selu
func (c_ CActivationLayer) SeluLayer() IMLCActivationLayer {
	rv := objc.Send[CActivationLayer](c_.ID, objc.Sel("seluLayer"))
	return rv
}/* debug [instance_properties/getter]: seluLayer */


// Creates an instance of a sigmoid activation layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/sigmoid
func (c_ CActivationLayer) SigmoidLayer() IMLCActivationLayer {
	rv := objc.Send[CActivationLayer](c_.ID, objc.Sel("sigmoidLayer"))
	return rv
}/* debug [instance_properties/getter]: sigmoidLayer */


// Creates an instance of a parametric soft plus activation layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/softPlus
func (c_ CActivationLayer) SoftPlusLayer() IMLCActivationLayer {
	rv := objc.Send[CActivationLayer](c_.ID, objc.Sel("softPlusLayer"))
	return rv
}/* debug [instance_properties/getter]: softPlusLayer */


// Creates an instance of a soft shrink activation layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/softShrink
func (c_ CActivationLayer) SoftShrinkLayer() IMLCActivationLayer {
	rv := objc.Send[CActivationLayer](c_.ID, objc.Sel("softShrinkLayer"))
	return rv
}/* debug [instance_properties/getter]: softShrinkLayer */


// Creates an instance of a parametric soft sign activation layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/softSign
func (c_ CActivationLayer) SoftSignLayer() IMLCActivationLayer {
	rv := objc.Send[CActivationLayer](c_.ID, objc.Sel("softSignLayer"))
	return rv
}/* debug [instance_properties/getter]: softSignLayer */


// Creates an instance of a hyperbolic tangent activation layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/tanh
func (c_ CActivationLayer) TanhLayer() IMLCActivationLayer {
	rv := objc.Send[CActivationLayer](c_.ID, objc.Sel("tanhLayer"))
	return rv
}/* debug [instance_properties/getter]: tanhLayer */


// Creates an instance of a tanh shrink activation layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer/tanhShrink
func (c_ CActivationLayer) TanhShrinkLayer() IMLCActivationLayer {
	rv := objc.Send[CActivationLayer](c_.ID, objc.Sel("tanhShrinkLayer"))
	return rv
}/* debug [instance_properties/getter]: tanhShrinkLayer */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCActivationLayer */


