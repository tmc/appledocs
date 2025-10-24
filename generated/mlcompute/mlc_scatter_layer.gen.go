// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CScatterLayer] class.
var (
	CScatterLayerClass     _CScatterLayerClass
	CScatterLayerClassOnce sync.Once
)

func getCScatterLayerClass() _CScatterLayerClass {
	CScatterLayerClassOnce.Do(func() {
		CScatterLayerClass = _CScatterLayerClass{objc.GetClass("MLCScatterLayer")}
	})
	return CScatterLayerClass
}

type _CScatterLayerClass struct {
	class objc.Class
}

// An interface definition for the [CScatterLayer] class.
type ICScatterLayer interface {
	ICLayer
	// properties:
	Dimension() int
	SetDimension(value int)
	ReductionType() CReductionType /* not a class type */
	SetReductionType(value CReductionType /* not a class type */)
	// methods:
}

// A layer that updates the output at an index you specify.


// A layer that updates the output at an index you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCScatterLayer
type CScatterLayer struct {
	CLayer
}

// CScatterLayerFrom constructs a [CScatterLayer] from an unsafe.Pointer.
//
// A layer that updates the output at an index you specify.
func CScatterLayerFrom(ptr unsafe.Pointer) CScatterLayer {
	return CScatterLayer{
		CLayer: CLayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CScatterLayerClass) Alloc() CScatterLayer {
	rv := objc.Send[CScatterLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CScatterLayerClass) New() CScatterLayer {
	rv := objc.Send[CScatterLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CScatterLayer) Init() CScatterLayer {
	rv := objc.Send[CScatterLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CScatterLayer) Autorelease() CScatterLayer {
	rv := objc.Send[CScatterLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCScatterLayer creates a new CScatterLayer instance.
func NewCScatterLayer() CScatterLayer {
	return getCScatterLayerClass().New()
}



// The dimension to index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcscatterlayer/dimension
func (c_ CScatterLayer) Dimension() int {
	rv := objc.Send[int](c_.ID, objc.Sel("dimension"))
	return rv
}


// The dimension to index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcscatterlayer/dimension
func (c_ CScatterLayer) SetDimension(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDimension:"), value)
}


// The reduction type that applies to all values in the source tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcscatterlayer/reductiontype
func (c_ CScatterLayer) ReductionType() CReductionType /* not a class type */ {
	rv := objc.Send[CReductionType](c_.ID, objc.Sel("reductionType"))
	return rv
}


// The reduction type that applies to all values in the source tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcscatterlayer/reductiontype
func (c_ CScatterLayer) SetReductionType(value CReductionType /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setReductionType:"), value)
}



