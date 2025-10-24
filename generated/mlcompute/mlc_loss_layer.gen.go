// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MLCLossLayer */


/* debug [class_header]: Header for MLCLossLayer */
// The class instance for the [CLossLayer] class.
var (
	CLossLayerClass     _CLossLayerClass
	CLossLayerClassOnce sync.Once
)

func getCLossLayerClass() _CLossLayerClass {
	CLossLayerClassOnce.Do(func() {
		CLossLayerClass = _CLossLayerClass{objc.GetClass("MLCLossLayer")}
	})
	return CLossLayerClass
}

type _CLossLayerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CLossLayer */
// An interface definition for the [CLossLayer] class.
type ICLossLayer interface {
	ICLayer
	
/* debug [class_interface_properties]: Properties for CLossLayer */
	// properties:
	Descriptor() IMLCLossDescriptor
	Weights() IMLCTensor
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CLossLayer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CLossLayer */
// Alloc allocates a new instance without initialization.
func (cc _CLossLayerClass) Alloc() CLossLayer {
	rv := objc.Send[CLossLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CLossLayerClass) New() CLossLayer {
	rv := objc.Send[CLossLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CLossLayer) Init() CLossLayer {
	rv := objc.Send[CLossLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CLossLayer) Autorelease() CLossLayer {
	rv := objc.Send[CLossLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCLossLayer creates a new CLossLayer instance.
func NewCLossLayer() CLossLayer {
	return getCLossLayerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CLossLayer */
// A layer that estimates the inaccuracies of the model to reduce the loss on the next evaluation.


// A layer that estimates the inaccuracies of the model to reduce the loss on the next evaluation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossLayer
type CLossLayer struct {
	CLayer
}

// CLossLayerFrom constructs a [CLossLayer] from an unsafe.Pointer.
//
// A layer that estimates the inaccuracies of the model to reduce the loss on the next evaluation.
func CLossLayerFrom(ptr unsafe.Pointer) CLossLayer {
	return CLossLayer{
		CLayer: CLayerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CLossLayer */

// Creates a loss layer with the descriptor you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossLayer/init(descriptor:)
func NewCLossLayerWithDescriptor(lossDescriptor IMLCLossDescriptor) CLossLayer {
	rv := objc.Send[CLossLayer](objc.ID(getCLossLayerClass().class), objc.Sel("layerWithDescriptor:"), lossDescriptor)
	return rv
}/* debug [class_init_methods/constructor]: NewCLossLayerWithDescriptor */


// Creates a loss layer with the descriptor and weights you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossLayer/init(descriptor:weights:)
func NewCLossLayerWithDescriptorWeights(lossDescriptor IMLCLossDescriptor, weights IMLCTensor) CLossLayer {
	rv := objc.Send[CLossLayer](objc.ID(getCLossLayerClass().class), objc.Sel("layerWithDescriptor:weights:"), lossDescriptor, weights)
	return rv
}/* debug [class_init_methods/constructor]: NewCLossLayerWithDescriptorWeights */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CLossLayer */

// Creates a categorical cross entropy loss layer with the reduction type, label smoothing, number of classes, and weight you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossLayer/categoricalCrossEntropy(reductionType:labelSmoothing:classCount:weight:)
func (cc _CLossLayerClass) CategoricalCrossEntropyLossWithReductionTypeLabelSmoothingClassCountWeight(reductionType CReductionType, labelSmoothing float32, classCount uint, weight float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("categoricalCrossEntropyLossWithReductionType:labelSmoothing:classCount:weight:"), reductionType, labelSmoothing, classCount, weight)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CategoricalCrossEntropyLossWithReductionTypeLabelSmoothingClassCountWeight) */


// Creates a categorical cross entropy loss layer with the reduction type, label smoothing, number of classes, and weights you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossLayer/categoricalCrossEntropy(reductionType:labelSmoothing:classCount:weights:)
func (cc _CLossLayerClass) CategoricalCrossEntropyLossWithReductionTypeLabelSmoothingClassCountWeights(reductionType CReductionType, labelSmoothing float32, classCount uint, weights IMLCTensor) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("categoricalCrossEntropyLossWithReductionType:labelSmoothing:classCount:weights:"), reductionType, labelSmoothing, classCount, weights)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CategoricalCrossEntropyLossWithReductionTypeLabelSmoothingClassCountWeights) */


// Creates a cosine distance loss layer with the reduction type and weight you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossLayer/cosineDistance(reductionType:weight:)
func (cc _CLossLayerClass) CosineDistanceLossWithReductionTypeWeight(reductionType CReductionType, weight float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("cosineDistanceLossWithReductionType:weight:"), reductionType, weight)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CosineDistanceLossWithReductionTypeWeight) */


// Creates a cosine distance loss layer with the reduction type and weights you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossLayer/cosineDistance(reductionType:weights:)
func (cc _CLossLayerClass) CosineDistanceLossWithReductionTypeWeights(reductionType CReductionType, weights IMLCTensor) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("cosineDistanceLossWithReductionType:weights:"), reductionType, weights)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CosineDistanceLossWithReductionTypeWeights) */


// Creates a hinge loss layer with the reduction type and weight you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossLayer/hingeLoss(reductionType:weight:)
func (cc _CLossLayerClass) HingeLossWithReductionTypeWeight(reductionType CReductionType, weight float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("hingeLossWithReductionType:weight:"), reductionType, weight)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=HingeLossWithReductionTypeWeight) */


// Creates a hinge loss layer with the reduction type and weights you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossLayer/hingeLoss(reductionType:weights:)
func (cc _CLossLayerClass) HingeLossWithReductionTypeWeights(reductionType CReductionType, weights IMLCTensor) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("hingeLossWithReductionType:weights:"), reductionType, weights)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=HingeLossWithReductionTypeWeights) */


// Creates a huber loss layer with the reduction type, delta, and weight you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossLayer/huberLoss(reductionType:delta:weight:)
func (cc _CLossLayerClass) HuberLossWithReductionTypeDeltaWeight(reductionType CReductionType, delta float32, weight float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("huberLossWithReductionType:delta:weight:"), reductionType, delta, weight)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=HuberLossWithReductionTypeDeltaWeight) */


// Creates a huber loss layer with the reduction type, delta, and weights you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossLayer/huberLoss(reductionType:delta:weights:)
func (cc _CLossLayerClass) HuberLossWithReductionTypeDeltaWeights(reductionType CReductionType, delta float32, weights IMLCTensor) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("huberLossWithReductionType:delta:weights:"), reductionType, delta, weights)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=HuberLossWithReductionTypeDeltaWeights) */


// Creates a loss layer with the descriptor you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossLayer/init(descriptor:)
func (cc _CLossLayerClass) LayerWithDescriptor(lossDescriptor IMLCLossDescriptor) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("layerWithDescriptor:"), lossDescriptor)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayerWithDescriptor) */


// Creates a loss layer with the descriptor and weights you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossLayer/init(descriptor:weights:)
func (cc _CLossLayerClass) LayerWithDescriptorWeights(lossDescriptor IMLCLossDescriptor, weights IMLCTensor) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("layerWithDescriptor:weights:"), lossDescriptor, weights)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayerWithDescriptorWeights) */


// Creates a mean absolute loss layer with the reduction type and weight.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossLayer/meanAbsoluteError(reductionType:weight:)
func (cc _CLossLayerClass) MeanAbsoluteErrorLossWithReductionTypeWeight(reductionType CReductionType, weight float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("meanAbsoluteErrorLossWithReductionType:weight:"), reductionType, weight)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=MeanAbsoluteErrorLossWithReductionTypeWeight) */


// Creates a mean absolute loss layer with the reduction type and weights you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossLayer/meanAbsoluteError(reductionType:weights:)
func (cc _CLossLayerClass) MeanAbsoluteErrorLossWithReductionTypeWeights(reductionType CReductionType, weights IMLCTensor) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("meanAbsoluteErrorLossWithReductionType:weights:"), reductionType, weights)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=MeanAbsoluteErrorLossWithReductionTypeWeights) */


// Creates a mean squared loss layer with the reduction type and weight you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossLayer/meanSquaredError(reductionType:weight:)
func (cc _CLossLayerClass) MeanSquaredErrorLossWithReductionTypeWeight(reductionType CReductionType, weight float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("meanSquaredErrorLossWithReductionType:weight:"), reductionType, weight)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=MeanSquaredErrorLossWithReductionTypeWeight) */


// Creates a mean squared loss layer with the reduction type and weights you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossLayer/meanSquaredError(reductionType:weights:)
func (cc _CLossLayerClass) MeanSquaredErrorLossWithReductionTypeWeights(reductionType CReductionType, weights IMLCTensor) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("meanSquaredErrorLossWithReductionType:weights:"), reductionType, weights)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=MeanSquaredErrorLossWithReductionTypeWeights) */


// Creates a sigmoid cross entropy loss layer with the reduction type, label smoothing, and weight you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossLayer/sigmoidCrossEntropy(reductionType:labelSmoothing:weight:)
func (cc _CLossLayerClass) SigmoidCrossEntropyLossWithReductionTypeLabelSmoothingWeight(reductionType CReductionType, labelSmoothing float32, weight float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("sigmoidCrossEntropyLossWithReductionType:labelSmoothing:weight:"), reductionType, labelSmoothing, weight)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SigmoidCrossEntropyLossWithReductionTypeLabelSmoothingWeight) */


// Creates a sigmoid cross entropy loss layer with the reduction type, label smoothing, and weights you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossLayer/sigmoidCrossEntropy(reductionType:labelSmoothing:weights:)
func (cc _CLossLayerClass) SigmoidCrossEntropyLossWithReductionTypeLabelSmoothingWeights(reductionType CReductionType, labelSmoothing float32, weights IMLCTensor) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("sigmoidCrossEntropyLossWithReductionType:labelSmoothing:weights:"), reductionType, labelSmoothing, weights)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SigmoidCrossEntropyLossWithReductionTypeLabelSmoothingWeights) */


// Creates a softmax cross entropy loss layer with the reduction type, label smoothing, number of classes, and weight you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossLayer/softmaxCrossEntropy(reductionType:labelSmoothing:classCount:weight:)
func (cc _CLossLayerClass) SoftmaxCrossEntropyLossWithReductionTypeLabelSmoothingClassCountWeight(reductionType CReductionType, labelSmoothing float32, classCount uint, weight float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("softmaxCrossEntropyLossWithReductionType:labelSmoothing:classCount:weight:"), reductionType, labelSmoothing, classCount, weight)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SoftmaxCrossEntropyLossWithReductionTypeLabelSmoothingClassCountWeight) */


// Creates a softmax cross entropy loss layer with the reduction type, label smoothing, number of classes, and weights you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossLayer/softmaxCrossEntropy(reductionType:labelSmoothing:classCount:weights:)
func (cc _CLossLayerClass) SoftmaxCrossEntropyLossWithReductionTypeLabelSmoothingClassCountWeights(reductionType CReductionType, labelSmoothing float32, classCount uint, weights IMLCTensor) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("softmaxCrossEntropyLossWithReductionType:labelSmoothing:classCount:weights:"), reductionType, labelSmoothing, classCount, weights)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SoftmaxCrossEntropyLossWithReductionTypeLabelSmoothingClassCountWeights) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CLossLayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CLossLayer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CLossLayer */

// The configuration object you use to create the loss layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossLayer/descriptor
func (c_ CLossLayer) Descriptor() IMLCLossDescriptor {
	rv := objc.Send[CLossDescriptor](c_.ID, objc.Sel("descriptor"))
	return rv
}/* debug [instance_properties/getter]: descriptor */


// The loss label weights tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossLayer/weights
func (c_ CLossLayer) Weights() IMLCTensor {
	rv := objc.Send[CTensor](c_.ID, objc.Sel("weights"))
	return rv
}/* debug [instance_properties/getter]: weights */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCLossLayer */


