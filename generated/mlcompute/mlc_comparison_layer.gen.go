// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CComparisonLayer] class.
var (
	CComparisonLayerClass     _CComparisonLayerClass
	CComparisonLayerClassOnce sync.Once
)

func getCComparisonLayerClass() _CComparisonLayerClass {
	CComparisonLayerClassOnce.Do(func() {
		CComparisonLayerClass = _CComparisonLayerClass{objc.GetClass("MLCComparisonLayer")}
	})
	return CComparisonLayerClass
}

type _CComparisonLayerClass struct {
	class objc.Class
}

// An interface definition for the [CComparisonLayer] class.
type ICComparisonLayer interface {
	ICLayer
	// properties:
	Operation() CComparisonOperation /* not a class type */
	SetOperation(value CComparisonOperation /* not a class type */)
	// methods:
}

// A layer that performs elementwise comparison of two tensors.
//
// The layer returns a tensor with the shape equal to the largest shape of operations. It fills with the Boolean value , where corresponds to the you specify.


// A layer that performs elementwise comparison of two tensors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCComparisonLayer
type CComparisonLayer struct {
	CLayer
}

// CComparisonLayerFrom constructs a [CComparisonLayer] from an unsafe.Pointer.
//
// A layer that performs elementwise comparison of two tensors.
func CComparisonLayerFrom(ptr unsafe.Pointer) CComparisonLayer {
	return CComparisonLayer{
		CLayer: CLayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CComparisonLayerClass) Alloc() CComparisonLayer {
	rv := objc.Send[CComparisonLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CComparisonLayerClass) New() CComparisonLayer {
	rv := objc.Send[CComparisonLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CComparisonLayer) Init() CComparisonLayer {
	rv := objc.Send[CComparisonLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CComparisonLayer) Autorelease() CComparisonLayer {
	rv := objc.Send[CComparisonLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCComparisonLayer creates a new CComparisonLayer instance.
func NewCComparisonLayer() CComparisonLayer {
	return getCComparisonLayerClass().New()
}



// The comparison layer’s operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlccomparisonlayer/operation
func (c_ CComparisonLayer) Operation() CComparisonOperation /* not a class type */ {
	rv := objc.Send[CComparisonOperation](c_.ID, objc.Sel("operation"))
	return rv
}


// The comparison layer’s operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlccomparisonlayer/operation
func (c_ CComparisonLayer) SetOperation(value CComparisonOperation /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOperation:"), value)
}



