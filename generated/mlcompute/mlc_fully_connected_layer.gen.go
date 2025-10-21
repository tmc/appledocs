// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [CFullyConnectedLayer] class.
type ICFullyConnectedLayer interface {
	ICLayer
}

// A layer that connects each input to each output within its layer.
//
// This is also known as a dense layer.
//
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

// Alloc allocates a new instance without initialization.
func (cc _CFullyConnectedLayerClass) Alloc() CFullyConnectedLayer {
	rv := objc.Send[CFullyConnectedLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The biases tensor parameter you use for optimizer updates.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcfullyconnectedlayer/biasesparameter
func (c_ CFullyConnectedLayer) BiasesParameter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("biasesParameter"))
	return rv
}


// SetBiasesParameter sets the value of the biasesParameter property.
// The biases tensor parameter you use for optimizer updates.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcfullyconnectedlayer/biasesparameter
func (c_ CFullyConnectedLayer) SetBiasesParameter(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBiasesParameter:"), value)
}

// The weights tensor you use for the fully connected layer.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcfullyconnectedlayer/weights
func (c_ CFullyConnectedLayer) Weights() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("weights"))
	return rv
}


// SetWeights sets the value of the weights property.
// The weights tensor you use for the fully connected layer.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcfullyconnectedlayer/weights
func (c_ CFullyConnectedLayer) SetWeights(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWeights:"), value)
}

// The biases tensor you use for the fully connected layer.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcfullyconnectedlayer/biases
func (c_ CFullyConnectedLayer) Biases() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("biases"))
	return rv
}


// SetBiases sets the value of the biases property.
// The biases tensor you use for the fully connected layer.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcfullyconnectedlayer/biases
func (c_ CFullyConnectedLayer) SetBiases(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBiases:"), value)
}

// The configuration object you use to create the fully connected layer.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcfullyconnectedlayer/descriptor
func (c_ CFullyConnectedLayer) Descriptor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("descriptor"))
	return rv
}


// SetDescriptor sets the value of the descriptor property.
// The configuration object you use to create the fully connected layer.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcfullyconnectedlayer/descriptor
func (c_ CFullyConnectedLayer) SetDescriptor(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDescriptor:"), value)
}

// The weights tensor parameter you use for optimizer updates.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcfullyconnectedlayer/weightsparameter
func (c_ CFullyConnectedLayer) WeightsParameter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("weightsParameter"))
	return rv
}


// SetWeightsParameter sets the value of the weightsParameter property.
// The weights tensor parameter you use for optimizer updates.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcfullyconnectedlayer/weightsparameter
func (c_ CFullyConnectedLayer) SetWeightsParameter(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWeightsParameter:"), value)
}



