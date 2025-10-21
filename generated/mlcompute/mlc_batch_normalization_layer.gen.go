// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [CBatchNormalizationLayer] class.
type ICBatchNormalizationLayer interface {
	ICLayer
}

// A layer that normalizes a batch of inputs.
//
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

// Alloc allocates a new instance without initialization.
func (cc _CBatchNormalizationLayerClass) Alloc() CBatchNormalizationLayer {
	rv := objc.Send[CBatchNormalizationLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




