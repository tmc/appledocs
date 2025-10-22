// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CTransposeLayer] class.
var (
	CTransposeLayerClass     _CTransposeLayerClass
	CTransposeLayerClassOnce sync.Once
)

func getCTransposeLayerClass() _CTransposeLayerClass {
	CTransposeLayerClassOnce.Do(func() {
		CTransposeLayerClass = _CTransposeLayerClass{objc.GetClass("MLCTransposeLayer")}
	})
	return CTransposeLayerClass
}

type _CTransposeLayerClass struct {
	class objc.Class
}

// An interface definition for the [CTransposeLayer] class.
type ICTransposeLayer interface {
	ICLayer
	Dimensions() int
	SetDimensions(value int)
}

// A layer that permutes the dimensions you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTransposeLayer
type CTransposeLayer struct {
	CLayer
}

// CTransposeLayerFrom constructs a [CTransposeLayer] from an unsafe.Pointer.
//
// A layer that permutes the dimensions you specify.
func CTransposeLayerFrom(ptr unsafe.Pointer) CTransposeLayer {
	return CTransposeLayer{
		CLayer: CLayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CTransposeLayerClass) Alloc() CTransposeLayer {
	rv := objc.Send[CTransposeLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CTransposeLayerClass) New() CTransposeLayer {
	rv := objc.Send[CTransposeLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CTransposeLayer) Init() CTransposeLayer {
	rv := objc.Send[CTransposeLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CTransposeLayer) Autorelease() CTransposeLayer {
	rv := objc.Send[CTransposeLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCTransposeLayer creates a new CTransposeLayer instance.
func NewCTransposeLayer() CTransposeLayer {
	return getCTransposeLayerClass().New()
}


// An array that contains an input axis source for each output axis, which represents the ordering of dimensions.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlctransposelayer/dimensions-71ed6
func (c_ CTransposeLayer) Dimensions() int {
	rv := objc.Send[int](c_.ID, objc.Sel("dimensions"))
	return rv
}


// SetDimensions sets the value of the dimensions property.
// An array that contains an input axis source for each output axis, which represents the ordering of dimensions.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlctransposelayer/dimensions-71ed6
func (c_ CTransposeLayer) SetDimensions(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDimensions:"), value)
}



