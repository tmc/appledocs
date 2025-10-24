// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MLCFullyConnectedLayer */


/* debug [class_header]: Header for MLCFullyConnectedLayer */
// The class instance for the [CFullyConnectedLayer] class.
var (
	CFullyConnectedLayerClass     _CFullyConnectedLayerClass
	CFullyConnectedLayerClassOnce sync.Once
)

func getCFullyConnectedLayerClass() _CFullyConnectedLayerClass {
	CFullyConnectedLayerClassOnce.Do(func() {
		CFullyConnectedLayerClass = _CFullyConnectedLayerClass{objc.GetClass("MLCFullyConnectedLayer")}
	})
	return CFullyConnectedLayerClass
}

type _CFullyConnectedLayerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CFullyConnectedLayer */
// An interface definition for the [CFullyConnectedLayer] class.
type ICFullyConnectedLayer interface {
	ICLayer
	
/* debug [class_interface_properties]: Properties for CFullyConnectedLayer */
	// properties:
	Biases() IMLCTensor
	BiasesParameter() IMLCTensorParameter
	Descriptor() IMLCConvolutionDescriptor
	Weights() IMLCTensor
	WeightsParameter() IMLCTensorParameter
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CFullyConnectedLayer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CFullyConnectedLayer */
// Alloc allocates a new instance without initialization.
func (cc _CFullyConnectedLayerClass) Alloc() CFullyConnectedLayer {
	rv := objc.Send[CFullyConnectedLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CFullyConnectedLayerClass) New() CFullyConnectedLayer {
	rv := objc.Send[CFullyConnectedLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CFullyConnectedLayer) Init() CFullyConnectedLayer {
	rv := objc.Send[CFullyConnectedLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CFullyConnectedLayer) Autorelease() CFullyConnectedLayer {
	rv := objc.Send[CFullyConnectedLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCFullyConnectedLayer creates a new CFullyConnectedLayer instance.
func NewCFullyConnectedLayer() CFullyConnectedLayer {
	return getCFullyConnectedLayerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CFullyConnectedLayer */
// A layer that connects each input to each output within its layer.
//
// This is also known as a dense layer.


// A layer that connects each input to each output within its layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCFullyConnectedLayer
type CFullyConnectedLayer struct {
	CLayer
}

// CFullyConnectedLayerFrom constructs a [CFullyConnectedLayer] from an unsafe.Pointer.
//
// A layer that connects each input to each output within its layer.
func CFullyConnectedLayerFrom(ptr unsafe.Pointer) CFullyConnectedLayer {
	return CFullyConnectedLayer{
		CLayer: CLayerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CFullyConnectedLayer */

// Creates a fully connected layer with the weights, biases, and convolution descriptor you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCFullyConnectedLayer/init(weights:biases:descriptor:)
func NewCFullyConnectedLayerWithWeightsBiasesDescriptor(weights IMLCTensor, biases IMLCTensor, descriptor IMLCConvolutionDescriptor) CFullyConnectedLayer {
	rv := objc.Send[CFullyConnectedLayer](objc.ID(getCFullyConnectedLayerClass().class), objc.Sel("layerWithWeights:biases:descriptor:"), weights, biases, descriptor)
	return rv
}/* debug [class_init_methods/constructor]: NewCFullyConnectedLayerWithWeightsBiasesDescriptor */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CFullyConnectedLayer */

// Creates a fully connected layer with the weights, biases, and convolution descriptor you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCFullyConnectedLayer/init(weights:biases:descriptor:)
func (cc _CFullyConnectedLayerClass) LayerWithWeightsBiasesDescriptor(weights IMLCTensor, biases IMLCTensor, descriptor IMLCConvolutionDescriptor) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("layerWithWeights:biases:descriptor:"), weights, biases, descriptor)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayerWithWeightsBiasesDescriptor) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CFullyConnectedLayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CFullyConnectedLayer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CFullyConnectedLayer */

// The biases tensor you use for the fully connected layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCFullyConnectedLayer/biases
func (c_ CFullyConnectedLayer) Biases() IMLCTensor {
	rv := objc.Send[CTensor](c_.ID, objc.Sel("biases"))
	return rv
}/* debug [instance_properties/getter]: biases */


// The biases tensor parameter you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCFullyConnectedLayer/biasesParameter
func (c_ CFullyConnectedLayer) BiasesParameter() IMLCTensorParameter {
	rv := objc.Send[CTensorParameter](c_.ID, objc.Sel("biasesParameter"))
	return rv
}/* debug [instance_properties/getter]: biasesParameter */


// The configuration object you use to create the fully connected layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCFullyConnectedLayer/descriptor
func (c_ CFullyConnectedLayer) Descriptor() IMLCConvolutionDescriptor {
	rv := objc.Send[CConvolutionDescriptor](c_.ID, objc.Sel("descriptor"))
	return rv
}/* debug [instance_properties/getter]: descriptor */


// The weights tensor you use for the fully connected layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCFullyConnectedLayer/weights
func (c_ CFullyConnectedLayer) Weights() IMLCTensor {
	rv := objc.Send[CTensor](c_.ID, objc.Sel("weights"))
	return rv
}/* debug [instance_properties/getter]: weights */


// The weights tensor parameter you use for optimizer updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCFullyConnectedLayer/weightsParameter
func (c_ CFullyConnectedLayer) WeightsParameter() IMLCTensorParameter {
	rv := objc.Send[CTensorParameter](c_.ID, objc.Sel("weightsParameter"))
	return rv
}/* debug [instance_properties/getter]: weightsParameter */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCFullyConnectedLayer */


