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



