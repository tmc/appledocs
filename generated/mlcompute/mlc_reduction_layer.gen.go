// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CReductionLayer] class.
var (
	CReductionLayerClass     _CReductionLayerClass
	CReductionLayerClassOnce sync.Once
)

func getCReductionLayerClass() _CReductionLayerClass {
	CReductionLayerClassOnce.Do(func() {
		CReductionLayerClass = _CReductionLayerClass{objc.GetClass("MLCReductionLayer")}
	})
	return CReductionLayerClass
}

type _CReductionLayerClass struct {
	class objc.Class
}

// An interface definition for the [CReductionLayer] class.
type ICReductionLayer interface {
	ICLayer
}

// A layer that reduces tensor values across a specific dimension to a scalar value.
//
// Use this layer to perform reduction operations on a given dimension. The output of this layer is a tensor of the same shape as the source tensor, except the layer sets the dimension to .
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCReductionLayer
type CReductionLayer struct {
	CLayer
}

// CReductionLayerFrom constructs a [CReductionLayer] from an unsafe.Pointer.
//
// A layer that reduces tensor values across a specific dimension to a scalar value.
func CReductionLayerFrom(ptr unsafe.Pointer) CReductionLayer {
	return CReductionLayer{
		CLayer: CLayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CReductionLayerClass) Alloc() CReductionLayer {
	rv := objc.Send[CReductionLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CReductionLayerClass) New() CReductionLayer {
	rv := objc.Send[CReductionLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CReductionLayer) Init() CReductionLayer {
	rv := objc.Send[CReductionLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CReductionLayer) Autorelease() CReductionLayer {
	rv := objc.Send[CReductionLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCReductionLayer creates a new CReductionLayer instance.
func NewCReductionLayer() CReductionLayer {
	return getCReductionLayerClass().New()
}


// The dimension to perform the reduction operation on.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcreductionlayer/dimension
func (c_ CReductionLayer) Dimension() int {
	rv := objc.Send[int](c_.ID, objc.Sel("dimension"))
	return rv
}


// SetDimension sets the value of the dimension property.
// The dimension to perform the reduction operation on.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcreductionlayer/dimension
func (c_ CReductionLayer) SetDimension(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDimension:"), value)
}

// The dimensions to perform the reduction operation on.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcreductionlayer/dimensions-9oph6
func (c_ CReductionLayer) Dimensions() int {
	rv := objc.Send[int](c_.ID, objc.Sel("dimensions"))
	return rv
}


// SetDimensions sets the value of the dimensions property.
// The dimensions to perform the reduction operation on.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcreductionlayer/dimensions-9oph6
func (c_ CReductionLayer) SetDimensions(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDimensions:"), value)
}

// The function reduction type the system uses for reduction.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcreductionlayer/reductiontype
func (c_ CReductionLayer) ReductionType() CReductionType {
	rv := objc.Send[CReductionType](c_.ID, objc.Sel("reductionType"))
	return rv
}


// SetReductionType sets the value of the reductionType property.
// The function reduction type the system uses for reduction.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcreductionlayer/reductiontype
func (c_ CReductionLayer) SetReductionType(value CReductionType) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setReductionType:"), value)
}



