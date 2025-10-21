// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CReshapeLayer] class.
var (
	CReshapeLayerClass     _CReshapeLayerClass
	CReshapeLayerClassOnce sync.Once
)

func getCReshapeLayerClass() _CReshapeLayerClass {
	CReshapeLayerClassOnce.Do(func() {
		CReshapeLayerClass = _CReshapeLayerClass{objc.GetClass("MLCReshapeLayer")}
	})
	return CReshapeLayerClass
}

type _CReshapeLayerClass struct {
	class objc.Class
}

// An interface definition for the [CReshapeLayer] class.
type ICReshapeLayer interface {
	ICLayer
}

// A layer that reshapes a tensor with the shape you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCReshapeLayer
type CReshapeLayer struct {
	CLayer
}

// CReshapeLayerFrom constructs a [CReshapeLayer] from an unsafe.Pointer.
//
// A layer that reshapes a tensor with the shape you specify.
func CReshapeLayerFrom(ptr unsafe.Pointer) CReshapeLayer {
	return CReshapeLayer{
		CLayer: CLayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CReshapeLayerClass) Alloc() CReshapeLayer {
	rv := objc.Send[CReshapeLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CReshapeLayerClass) New() CReshapeLayer {
	rv := objc.Send[CReshapeLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CReshapeLayer) Init() CReshapeLayer {
	rv := objc.Send[CReshapeLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CReshapeLayer) Autorelease() CReshapeLayer {
	rv := objc.Send[CReshapeLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCReshapeLayer creates a new CReshapeLayer instance.
func NewCReshapeLayer() CReshapeLayer {
	return getCReshapeLayerClass().New()
}


// An array that contains the size of each dimension.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcreshapelayer/shape-8k50y
func (c_ CReshapeLayer) Shape() int {
	rv := objc.Send[int](c_.ID, objc.Sel("shape"))
	return rv
}


// SetShape sets the value of the shape property.
// An array that contains the size of each dimension.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcreshapelayer/shape-8k50y
func (c_ CReshapeLayer) SetShape(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShape:"), value)
}



