// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MLCEmbeddingLayer */


/* debug [class_header]: Header for MLCEmbeddingLayer */
// The class instance for the [CEmbeddingLayer] class.
var (
	CEmbeddingLayerClass     _CEmbeddingLayerClass
	CEmbeddingLayerClassOnce sync.Once
)

func getCEmbeddingLayerClass() _CEmbeddingLayerClass {
	CEmbeddingLayerClassOnce.Do(func() {
		CEmbeddingLayerClass = _CEmbeddingLayerClass{objc.GetClass("MLCEmbeddingLayer")}
	})
	return CEmbeddingLayerClass
}

type _CEmbeddingLayerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CEmbeddingLayer */
// An interface definition for the [CEmbeddingLayer] class.
type ICEmbeddingLayer interface {
	ICLayer
	
/* debug [class_interface_properties]: Properties for CEmbeddingLayer */
	// properties:
	Descriptor() IMLCEmbeddingDescriptor
	Weights() IMLCTensor
	WeightsParameter() IMLCTensorParameter
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CEmbeddingLayer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CEmbeddingLayer */
// Alloc allocates a new instance without initialization.
func (cc _CEmbeddingLayerClass) Alloc() CEmbeddingLayer {
	rv := objc.Send[CEmbeddingLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CEmbeddingLayerClass) New() CEmbeddingLayer {
	rv := objc.Send[CEmbeddingLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CEmbeddingLayer) Init() CEmbeddingLayer {
	rv := objc.Send[CEmbeddingLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CEmbeddingLayer) Autorelease() CEmbeddingLayer {
	rv := objc.Send[CEmbeddingLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCEmbeddingLayer creates a new CEmbeddingLayer instance.
func NewCEmbeddingLayer() CEmbeddingLayer {
	return getCEmbeddingLayerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CEmbeddingLayer */
// A layer that stores a word embedding.


// A layer that stores a word embedding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCEmbeddingLayer
type CEmbeddingLayer struct {
	CLayer
}

// CEmbeddingLayerFrom constructs a [CEmbeddingLayer] from an unsafe.Pointer.
//
// A layer that stores a word embedding.
func CEmbeddingLayerFrom(ptr unsafe.Pointer) CEmbeddingLayer {
	return CEmbeddingLayer{
		CLayer: CLayerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CEmbeddingLayer */

// Creates an embedding layer with the descriptor and word embedding weights tensor you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCEmbeddingLayer/init(descriptor:weights:)
func NewCEmbeddingLayerWithDescriptorWeights(descriptor IMLCEmbeddingDescriptor, weights IMLCTensor) CEmbeddingLayer {
	rv := objc.Send[CEmbeddingLayer](objc.ID(getCEmbeddingLayerClass().class), objc.Sel("layerWithDescriptor:weights:"), descriptor, weights)
	return rv
}/* debug [class_init_methods/constructor]: NewCEmbeddingLayerWithDescriptorWeights */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CEmbeddingLayer */

// Creates an embedding layer with the descriptor and word embedding weights tensor you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCEmbeddingLayer/init(descriptor:weights:)
func (cc _CEmbeddingLayerClass) LayerWithDescriptorWeights(descriptor IMLCEmbeddingDescriptor, weights IMLCTensor) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("layerWithDescriptor:weights:"), descriptor, weights)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayerWithDescriptorWeights) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CEmbeddingLayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CEmbeddingLayer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CEmbeddingLayer */

// The configuration object you use to create the embedding layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCEmbeddingLayer/descriptor
func (c_ CEmbeddingLayer) Descriptor() IMLCEmbeddingDescriptor {
	rv := objc.Send[CEmbeddingDescriptor](c_.ID, objc.Sel("descriptor"))
	return rv
}/* debug [instance_properties/getter]: descriptor */


// The weights tensor that contains the word embedding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCEmbeddingLayer/weights
func (c_ CEmbeddingLayer) Weights() IMLCTensor {
	rv := objc.Send[CTensor](c_.ID, objc.Sel("weights"))
	return rv
}/* debug [instance_properties/getter]: weights */


// The tensor parameter that describes the weights for the optimizer update.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCEmbeddingLayer/weightsParameter
func (c_ CEmbeddingLayer) WeightsParameter() IMLCTensorParameter {
	rv := objc.Send[CTensorParameter](c_.ID, objc.Sel("weightsParameter"))
	return rv
}/* debug [instance_properties/getter]: weightsParameter */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCEmbeddingLayer */


