// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CMultiheadAttentionLayer] class.
var (
	CMultiheadAttentionLayerClass     _CMultiheadAttentionLayerClass
	CMultiheadAttentionLayerClassOnce sync.Once
)

func getCMultiheadAttentionLayerClass() _CMultiheadAttentionLayerClass {
	CMultiheadAttentionLayerClassOnce.Do(func() {
		CMultiheadAttentionLayerClass = _CMultiheadAttentionLayerClass{objc.GetClass("MLCMultiheadAttentionLayer")}
	})
	return CMultiheadAttentionLayerClass
}

type _CMultiheadAttentionLayerClass struct {
	class objc.Class
}

// An interface definition for the [CMultiheadAttentionLayer] class.
type ICMultiheadAttentionLayer interface {
	ICLayer
}

// A multihead, scaled dot-product attention layer that attends to one or more entries in the input key-value pairs.
//
// The dimensions of projections are as follows: ``
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCMultiheadAttentionLayer
type CMultiheadAttentionLayer struct {
	CLayer
}

// CMultiheadAttentionLayerFrom constructs a [CMultiheadAttentionLayer] from an unsafe.Pointer.
//
// A multihead, scaled dot-product attention layer that attends to one or more entries in the input key-value pairs.
func CMultiheadAttentionLayerFrom(ptr unsafe.Pointer) CMultiheadAttentionLayer {
	return CMultiheadAttentionLayer{
		CLayer: CLayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CMultiheadAttentionLayerClass) Alloc() CMultiheadAttentionLayer {
	rv := objc.Send[CMultiheadAttentionLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CMultiheadAttentionLayerClass) New() CMultiheadAttentionLayer {
	rv := objc.Send[CMultiheadAttentionLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CMultiheadAttentionLayer) Init() CMultiheadAttentionLayer {
	rv := objc.Send[CMultiheadAttentionLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CMultiheadAttentionLayer) Autorelease() CMultiheadAttentionLayer {
	rv := objc.Send[CMultiheadAttentionLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCMultiheadAttentionLayer creates a new CMultiheadAttentionLayer instance.
func NewCMultiheadAttentionLayer() CMultiheadAttentionLayer {
	return getCMultiheadAttentionLayerClass().New()
}




