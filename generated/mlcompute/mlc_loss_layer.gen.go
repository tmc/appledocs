// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [CLossLayer] class.
type ICLossLayer interface {
	ICLayer
}

// A layer that estimates the inaccuracies of the model to reduce the loss on the next evaluation.
//
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

// Alloc allocates a new instance without initialization.
func (cc _CLossLayerClass) Alloc() CLossLayer {
	rv := objc.Send[CLossLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Creates a mean squared loss layer with the reduction type and weights you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLossLayer/meanSquaredError(reductionType:weights:)
func (cc _CLossLayerClass) MeanSquaredErrorLossWithReductionTypeWeights(reductionType unsafe.Pointer, weights unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("meanSquaredErrorLossWithReductionType:weights:"), reductionType, weights)
	return rv
}

// The configuration object you use to create the loss layer.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclosslayer/descriptor
func (c_ CLossLayer) Descriptor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("descriptor"))
	return rv
}


// SetDescriptor sets the value of the descriptor property.
// The configuration object you use to create the loss layer.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclosslayer/descriptor
func (c_ CLossLayer) SetDescriptor(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDescriptor:"), value)
}

// The loss label weights tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclosslayer/weights
func (c_ CLossLayer) Weights() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("weights"))
	return rv
}


// SetWeights sets the value of the weights property.
// The loss label weights tensor.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlclosslayer/weights
func (c_ CLossLayer) SetWeights(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWeights:"), value)
}



