// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CSelectionLayer] class.
var (
	CSelectionLayerClass     _CSelectionLayerClass
	CSelectionLayerClassOnce sync.Once
)

func getCSelectionLayerClass() _CSelectionLayerClass {
	CSelectionLayerClassOnce.Do(func() {
		CSelectionLayerClass = _CSelectionLayerClass{objc.GetClass("MLCSelectionLayer")}
	})
	return CSelectionLayerClass
}

type _CSelectionLayerClass struct {
	class objc.Class
}

// An interface definition for the [CSelectionLayer] class.
type ICSelectionLayer interface {
	ICLayer
}

// A layer for selecting elements from two tensors.
//
// A selection layer takes a condition tensor that acts as a mask. It determines whether the corresponding element or row in the output comes from tensor (if the element in the condition is ) or tensor (if ).
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCSelectionLayer
type CSelectionLayer struct {
	CLayer
}

// CSelectionLayerFrom constructs a [CSelectionLayer] from an unsafe.Pointer.
//
// A layer for selecting elements from two tensors.
func CSelectionLayerFrom(ptr unsafe.Pointer) CSelectionLayer {
	return CSelectionLayer{
		CLayer: CLayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CSelectionLayerClass) Alloc() CSelectionLayer {
	rv := objc.Send[CSelectionLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CSelectionLayerClass) New() CSelectionLayer {
	rv := objc.Send[CSelectionLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CSelectionLayer) Init() CSelectionLayer {
	rv := objc.Send[CSelectionLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CSelectionLayer) Autorelease() CSelectionLayer {
	rv := objc.Send[CSelectionLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCSelectionLayer creates a new CSelectionLayer instance.
func NewCSelectionLayer() CSelectionLayer {
	return getCSelectionLayerClass().New()
}




