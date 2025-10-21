// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CMatMulLayer] class.
var (
	CMatMulLayerClass     _CMatMulLayerClass
	CMatMulLayerClassOnce sync.Once
)

func getCMatMulLayerClass() _CMatMulLayerClass {
	CMatMulLayerClassOnce.Do(func() {
		CMatMulLayerClass = _CMatMulLayerClass{objc.GetClass("MLCMatMulLayer")}
	})
	return CMatMulLayerClass
}

type _CMatMulLayerClass struct {
	class objc.Class
}

// An interface definition for the [CMatMulLayer] class.
type ICMatMulLayer interface {
	ICLayer
}

// A layer that multiplies matrices.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCMatMulLayer
type CMatMulLayer struct {
	CLayer
}

// CMatMulLayerFrom constructs a [CMatMulLayer] from an unsafe.Pointer.
//
// A layer that multiplies matrices.
func CMatMulLayerFrom(ptr unsafe.Pointer) CMatMulLayer {
	return CMatMulLayer{
		CLayer: CLayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CMatMulLayerClass) Alloc() CMatMulLayer {
	rv := objc.Send[CMatMulLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CMatMulLayerClass) New() CMatMulLayer {
	rv := objc.Send[CMatMulLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CMatMulLayer) Init() CMatMulLayer {
	rv := objc.Send[CMatMulLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CMatMulLayer) Autorelease() CMatMulLayer {
	rv := objc.Send[CMatMulLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCMatMulLayer creates a new CMatMulLayer instance.
func NewCMatMulLayer() CMatMulLayer {
	return getCMatMulLayerClass().New()
}




