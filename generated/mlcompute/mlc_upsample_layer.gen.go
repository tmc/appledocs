// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CUpsampleLayer] class.
var (
	CUpsampleLayerClass     _CUpsampleLayerClass
	CUpsampleLayerClassOnce sync.Once
)

func getCUpsampleLayerClass() _CUpsampleLayerClass {
	CUpsampleLayerClassOnce.Do(func() {
		CUpsampleLayerClass = _CUpsampleLayerClass{objc.GetClass("MLCUpsampleLayer")}
	})
	return CUpsampleLayerClass
}

type _CUpsampleLayerClass struct {
	class objc.Class
}

// An interface definition for the [CUpsampleLayer] class.
type ICUpsampleLayer interface {
	ICLayer
}

// A layer that applies upsampling with the shape you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCUpsampleLayer
type CUpsampleLayer struct {
	CLayer
}

// CUpsampleLayerFrom constructs a [CUpsampleLayer] from an unsafe.Pointer.
//
// A layer that applies upsampling with the shape you specify.
func CUpsampleLayerFrom(ptr unsafe.Pointer) CUpsampleLayer {
	return CUpsampleLayer{
		CLayer: CLayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CUpsampleLayerClass) Alloc() CUpsampleLayer {
	rv := objc.Send[CUpsampleLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CUpsampleLayerClass) New() CUpsampleLayer {
	rv := objc.Send[CUpsampleLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CUpsampleLayer) Init() CUpsampleLayer {
	rv := objc.Send[CUpsampleLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CUpsampleLayer) Autorelease() CUpsampleLayer {
	rv := objc.Send[CUpsampleLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCUpsampleLayer creates a new CUpsampleLayer instance.
func NewCUpsampleLayer() CUpsampleLayer {
	return getCUpsampleLayerClass().New()
}




