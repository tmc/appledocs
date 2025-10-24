// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MLCConvolutionLayer */


/* debug [class_header]: Header for MLCConvolutionLayer */
// The class instance for the [CConvolutionLayer] class.
var (
	CConvolutionLayerClass     _CConvolutionLayerClass
	CConvolutionLayerClassOnce sync.Once
)

func getCConvolutionLayerClass() _CConvolutionLayerClass {
	CConvolutionLayerClassOnce.Do(func() {
		CConvolutionLayerClass = _CConvolutionLayerClass{objc.GetClass("MLCConvolutionLayer")}
	})
	return CConvolutionLayerClass
}

type _CConvolutionLayerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CConvolutionLayer */
// An interface definition for the [CConvolutionLayer] class.
type ICConvolutionLayer interface {
	ICLayer
	
/* debug [class_interface_properties]: Properties for CConvolutionLayer */
	// properties:
	Biases() IMLCTensor
	BiasesParameter() IMLCTensorParameter
	Descriptor() IMLCConvolutionDescriptor
	Weights() IMLCTensor
	WeightsParameter() IMLCTensorParameter
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CConvolutionLayer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CConvolutionLayer */
// Alloc allocates a new instance without initialization.
func (cc _CConvolutionLayerClass) Alloc() CConvolutionLayer {
	rv := objc.Send[CConvolutionLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CConvolutionLayerClass) New() CConvolutionLayer {
	rv := objc.Send[CConvolutionLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CConvolutionLayer) Init() CConvolutionLayer {
	rv := objc.Send[CConvolutionLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CConvolutionLayer) Autorelease() CConvolutionLayer {
	rv := objc.Send[CConvolutionLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCConvolutionLayer creates a new CConvolutionLayer instance.
func NewCConvolutionLayer() CConvolutionLayer {
	return getCConvolutionLayerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CConvolutionLayer */
// A layer that applies a convolution over a signal.


// A layer that applies a convolution over a signal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConvolutionLayer
type CConvolutionLayer struct {
	CLayer
}

// CConvolutionLayerFrom constructs a [CConvolutionLayer] from an unsafe.Pointer.
//
// A layer that applies a convolution over a signal.
func CConvolutionLayerFrom(ptr unsafe.Pointer) CConvolutionLayer {
	return CConvolutionLayer{
		CLayer: CLayerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CConvolutionLayer */

// Creates a convolution layer with the weights, biases, and descriptor you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConvolutionLayer/init(weights:biases:descriptor:)
func NewCConvolutionLayerWithWeightsBiasesDescriptor(weights IMLCTensor, biases IMLCTensor, descriptor IMLCConvolutionDescriptor) CConvolutionLayer {
	rv := objc.Send[CConvolutionLayer](objc.ID(getCConvolutionLayerClass().class), objc.Sel("layerWithWeights:biases:descriptor:"), weights, biases, descriptor)
	return rv
}/* debug [class_init_methods/constructor]: NewCConvolutionLayerWithWeightsBiasesDescriptor */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CConvolutionLayer */

// Creates a convolution layer with the weights, biases, and descriptor you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConvolutionLayer/init(weights:biases:descriptor:)
func (cc _CConvolutionLayerClass) LayerWithWeightsBiasesDescriptor(weights IMLCTensor, biases IMLCTensor, descriptor IMLCConvolutionDescriptor) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("layerWithWeights:biases:descriptor:"), weights, biases, descriptor)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayerWithWeightsBiasesDescriptor) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CConvolutionLayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CConvolutionLayer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CConvolutionLayer */

// The biases tensor you use for the convolution layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConvolutionLayer/biases
func (c_ CConvolutionLayer) Biases() IMLCTensor {
	rv := objc.Send[CTensor](c_.ID, objc.Sel("biases"))
	return rv
}/* debug [instance_properties/getter]: biases */


// The biases tensor parameter you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConvolutionLayer/biasesParameter
func (c_ CConvolutionLayer) BiasesParameter() IMLCTensorParameter {
	rv := objc.Send[CTensorParameter](c_.ID, objc.Sel("biasesParameter"))
	return rv
}/* debug [instance_properties/getter]: biasesParameter */


// The configuration object you use to create the convolution layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConvolutionLayer/descriptor
func (c_ CConvolutionLayer) Descriptor() IMLCConvolutionDescriptor {
	rv := objc.Send[CConvolutionDescriptor](c_.ID, objc.Sel("descriptor"))
	return rv
}/* debug [instance_properties/getter]: descriptor */


// The weights tensor you use for the convolution layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConvolutionLayer/weights
func (c_ CConvolutionLayer) Weights() IMLCTensor {
	rv := objc.Send[CTensor](c_.ID, objc.Sel("weights"))
	return rv
}/* debug [instance_properties/getter]: weights */


// The weights tensor parameter you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConvolutionLayer/weightsParameter
func (c_ CConvolutionLayer) WeightsParameter() IMLCTensorParameter {
	rv := objc.Send[CTensorParameter](c_.ID, objc.Sel("weightsParameter"))
	return rv
}/* debug [instance_properties/getter]: weightsParameter */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCConvolutionLayer */


