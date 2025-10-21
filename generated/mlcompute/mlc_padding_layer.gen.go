// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CPaddingLayer] class.
var (
	CPaddingLayerClass     _CPaddingLayerClass
	CPaddingLayerClassOnce sync.Once
)

func getCPaddingLayerClass() _CPaddingLayerClass {
	CPaddingLayerClassOnce.Do(func() {
		CPaddingLayerClass = _CPaddingLayerClass{objc.GetClass("MLCPaddingLayer")}
	})
	return CPaddingLayerClass
}

type _CPaddingLayerClass struct {
	class objc.Class
}

// An interface definition for the [CPaddingLayer] class.
type ICPaddingLayer interface {
	ICLayer
}

// A layer that pads a tensor with the padding sizes you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPaddingLayer
type CPaddingLayer struct {
	CLayer
}

// CPaddingLayerFrom constructs a [CPaddingLayer] from an unsafe.Pointer.
//
// A layer that pads a tensor with the padding sizes you specify.
func CPaddingLayerFrom(ptr unsafe.Pointer) CPaddingLayer {
	return CPaddingLayer{
		CLayer: CLayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CPaddingLayerClass) Alloc() CPaddingLayer {
	rv := objc.Send[CPaddingLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CPaddingLayerClass) New() CPaddingLayer {
	rv := objc.Send[CPaddingLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CPaddingLayer) Init() CPaddingLayer {
	rv := objc.Send[CPaddingLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CPaddingLayer) Autorelease() CPaddingLayer {
	rv := objc.Send[CPaddingLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCPaddingLayer creates a new CPaddingLayer instance.
func NewCPaddingLayer() CPaddingLayer {
	return getCPaddingLayerClass().New()
}




