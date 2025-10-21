// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CSliceLayer] class.
var (
	CSliceLayerClass     _CSliceLayerClass
	CSliceLayerClassOnce sync.Once
)

func getCSliceLayerClass() _CSliceLayerClass {
	CSliceLayerClassOnce.Do(func() {
		CSliceLayerClass = _CSliceLayerClass{objc.GetClass("MLCSliceLayer")}
	})
	return CSliceLayerClass
}

type _CSliceLayerClass struct {
	class objc.Class
}

// An interface definition for the [CSliceLayer] class.
type ICSliceLayer interface {
	ICLayer
}

// A layer that extracts a slice from a tensor.
//
// The framework supports positive stride. Use a slice layer to slice a given source. Slicing won’t decrease the tensor dimension. The start, end, and stride vectors must be of the same size, equal to the source tensor dimension.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCSliceLayer
type CSliceLayer struct {
	CLayer
}

// CSliceLayerFrom constructs a [CSliceLayer] from an unsafe.Pointer.
//
// A layer that extracts a slice from a tensor.
func CSliceLayerFrom(ptr unsafe.Pointer) CSliceLayer {
	return CSliceLayer{
		CLayer: CLayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CSliceLayerClass) Alloc() CSliceLayer {
	rv := objc.Send[CSliceLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CSliceLayerClass) New() CSliceLayer {
	rv := objc.Send[CSliceLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CSliceLayer) Init() CSliceLayer {
	rv := objc.Send[CSliceLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CSliceLayer) Autorelease() CSliceLayer {
	rv := objc.Send[CSliceLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCSliceLayer creates a new CSliceLayer instance.
func NewCSliceLayer() CSliceLayer {
	return getCSliceLayerClass().New()
}




