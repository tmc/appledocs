// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CLayerNormalizationLayer] class.
var (
	CLayerNormalizationLayerClass     _CLayerNormalizationLayerClass
	CLayerNormalizationLayerClassOnce sync.Once
)

func getCLayerNormalizationLayerClass() _CLayerNormalizationLayerClass {
	CLayerNormalizationLayerClassOnce.Do(func() {
		CLayerNormalizationLayerClass = _CLayerNormalizationLayerClass{objc.GetClass("MLCLayerNormalizationLayer")}
	})
	return CLayerNormalizationLayerClass
}

type _CLayerNormalizationLayerClass struct {
	class objc.Class
}

// An interface definition for the [CLayerNormalizationLayer] class.
type ICLayerNormalizationLayer interface {
	ICLayer
}

// A layer that applies layer normalization over inputs.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLayerNormalizationLayer
type CLayerNormalizationLayer struct {
	CLayer
}

// CLayerNormalizationLayerFrom constructs a [CLayerNormalizationLayer] from an unsafe.Pointer.
//
// A layer that applies layer normalization over inputs.
func CLayerNormalizationLayerFrom(ptr unsafe.Pointer) CLayerNormalizationLayer {
	return CLayerNormalizationLayer{
		CLayer: CLayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CLayerNormalizationLayerClass) Alloc() CLayerNormalizationLayer {
	rv := objc.Send[CLayerNormalizationLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CLayerNormalizationLayerClass) New() CLayerNormalizationLayer {
	rv := objc.Send[CLayerNormalizationLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CLayerNormalizationLayer) Init() CLayerNormalizationLayer {
	rv := objc.Send[CLayerNormalizationLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CLayerNormalizationLayer) Autorelease() CLayerNormalizationLayer {
	rv := objc.Send[CLayerNormalizationLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCLayerNormalizationLayer creates a new CLayerNormalizationLayer instance.
func NewCLayerNormalizationLayer() CLayerNormalizationLayer {
	return getCLayerNormalizationLayerClass().New()
}




