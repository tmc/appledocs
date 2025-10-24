// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CArithmeticLayer] class.
var (
	CArithmeticLayerClass     _CArithmeticLayerClass
	CArithmeticLayerClassOnce sync.Once
)

func getCArithmeticLayerClass() _CArithmeticLayerClass {
	CArithmeticLayerClassOnce.Do(func() {
		CArithmeticLayerClass = _CArithmeticLayerClass{objc.GetClass("MLCArithmeticLayer")}
	})
	return CArithmeticLayerClass
}

type _CArithmeticLayerClass struct {
	class objc.Class
}

// An interface definition for the [CArithmeticLayer] class.
type ICArithmeticLayer interface {
	ICLayer
	// properties:
	Operation() CArithmeticOperation /* not a class type */
	SetOperation(value CArithmeticOperation /* not a class type */)
	// methods:
}

// A layer that performs an arithmetic operation.


// A layer that performs an arithmetic operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCArithmeticLayer
type CArithmeticLayer struct {
	CLayer
}

// CArithmeticLayerFrom constructs a [CArithmeticLayer] from an unsafe.Pointer.
//
// A layer that performs an arithmetic operation.
func CArithmeticLayerFrom(ptr unsafe.Pointer) CArithmeticLayer {
	return CArithmeticLayer{
		CLayer: CLayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CArithmeticLayerClass) Alloc() CArithmeticLayer {
	rv := objc.Send[CArithmeticLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CArithmeticLayerClass) New() CArithmeticLayer {
	rv := objc.Send[CArithmeticLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CArithmeticLayer) Init() CArithmeticLayer {
	rv := objc.Send[CArithmeticLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CArithmeticLayer) Autorelease() CArithmeticLayer {
	rv := objc.Send[CArithmeticLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCArithmeticLayer creates a new CArithmeticLayer instance.
func NewCArithmeticLayer() CArithmeticLayer {
	return getCArithmeticLayerClass().New()
}



// The arithmetic layer’s operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcarithmeticlayer/operation
func (c_ CArithmeticLayer) Operation() CArithmeticOperation /* not a class type */ {
	rv := objc.Send[CArithmeticOperation](c_.ID, objc.Sel("operation"))
	return rv
}


// The arithmetic layer’s operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcarithmeticlayer/operation
func (c_ CArithmeticLayer) SetOperation(value CArithmeticOperation /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOperation:"), value)
}



