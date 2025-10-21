// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CGroupNormalizationLayer] class.
var (
	CGroupNormalizationLayerClass     _CGroupNormalizationLayerClass
	CGroupNormalizationLayerClassOnce sync.Once
)

func getCGroupNormalizationLayerClass() _CGroupNormalizationLayerClass {
	CGroupNormalizationLayerClassOnce.Do(func() {
		CGroupNormalizationLayerClass = _CGroupNormalizationLayerClass{objc.GetClass("MLCGroupNormalizationLayer")}
	})
	return CGroupNormalizationLayerClass
}

type _CGroupNormalizationLayerClass struct {
	class objc.Class
}

// An interface definition for the [CGroupNormalizationLayer] class.
type ICGroupNormalizationLayer interface {
	ICLayer
}

// A layer that divides the channels into groups for normalization.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGroupNormalizationLayer
type CGroupNormalizationLayer struct {
	CLayer
}

// CGroupNormalizationLayerFrom constructs a [CGroupNormalizationLayer] from an unsafe.Pointer.
//
// A layer that divides the channels into groups for normalization.
func CGroupNormalizationLayerFrom(ptr unsafe.Pointer) CGroupNormalizationLayer {
	return CGroupNormalizationLayer{
		CLayer: CLayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CGroupNormalizationLayerClass) Alloc() CGroupNormalizationLayer {
	rv := objc.Send[CGroupNormalizationLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CGroupNormalizationLayerClass) New() CGroupNormalizationLayer {
	rv := objc.Send[CGroupNormalizationLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CGroupNormalizationLayer) Init() CGroupNormalizationLayer {
	rv := objc.Send[CGroupNormalizationLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CGroupNormalizationLayer) Autorelease() CGroupNormalizationLayer {
	rv := objc.Send[CGroupNormalizationLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCGroupNormalizationLayer creates a new CGroupNormalizationLayer instance.
func NewCGroupNormalizationLayer() CGroupNormalizationLayer {
	return getCGroupNormalizationLayerClass().New()
}




