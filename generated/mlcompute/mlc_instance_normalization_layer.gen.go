// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CInstanceNormalizationLayer] class.
var (
	CInstanceNormalizationLayerClass     _CInstanceNormalizationLayerClass
	CInstanceNormalizationLayerClassOnce sync.Once
)

func getCInstanceNormalizationLayerClass() _CInstanceNormalizationLayerClass {
	CInstanceNormalizationLayerClassOnce.Do(func() {
		CInstanceNormalizationLayerClass = _CInstanceNormalizationLayerClass{objc.GetClass("MLCInstanceNormalizationLayer")}
	})
	return CInstanceNormalizationLayerClass
}

type _CInstanceNormalizationLayerClass struct {
	class objc.Class
}

// An interface definition for the [CInstanceNormalizationLayer] class.
type ICInstanceNormalizationLayer interface {
	ICLayer
}

// A layer that normalizes all features of one channel.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCInstanceNormalizationLayer
type CInstanceNormalizationLayer struct {
	CLayer
}

// CInstanceNormalizationLayerFrom constructs a [CInstanceNormalizationLayer] from an unsafe.Pointer.
//
// A layer that normalizes all features of one channel.
func CInstanceNormalizationLayerFrom(ptr unsafe.Pointer) CInstanceNormalizationLayer {
	return CInstanceNormalizationLayer{
		CLayer: CLayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CInstanceNormalizationLayerClass) Alloc() CInstanceNormalizationLayer {
	rv := objc.Send[CInstanceNormalizationLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CInstanceNormalizationLayerClass) New() CInstanceNormalizationLayer {
	rv := objc.Send[CInstanceNormalizationLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CInstanceNormalizationLayer) Init() CInstanceNormalizationLayer {
	rv := objc.Send[CInstanceNormalizationLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CInstanceNormalizationLayer) Autorelease() CInstanceNormalizationLayer {
	rv := objc.Send[CInstanceNormalizationLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCInstanceNormalizationLayer creates a new CInstanceNormalizationLayer instance.
func NewCInstanceNormalizationLayer() CInstanceNormalizationLayer {
	return getCInstanceNormalizationLayerClass().New()
}




