// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CFullyConnectedLayer] class.
var (
	CFullyConnectedLayerClass     _CFullyConnectedLayerClass
	CFullyConnectedLayerClassOnce sync.Once
)

func getCFullyConnectedLayerClass() _CFullyConnectedLayerClass {
	CFullyConnectedLayerClassOnce.Do(func() {
		CFullyConnectedLayerClass = _CFullyConnectedLayerClass{objc.GetClass("MLCFullyConnectedLayer")}
	})
	return CFullyConnectedLayerClass
}

type _CFullyConnectedLayerClass struct {
	class objc.Class
}

// An interface definition for the [CFullyConnectedLayer] class.
type ICFullyConnectedLayer interface {
	ICLayer
}

// A layer that connects each input to each output within its layer.
//
// This is also known as a dense layer.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCFullyConnectedLayer
type CFullyConnectedLayer struct {
	CLayer
}

// CFullyConnectedLayerFrom constructs a [CFullyConnectedLayer] from an unsafe.Pointer.
//
// A layer that connects each input to each output within its layer.
func CFullyConnectedLayerFrom(ptr unsafe.Pointer) CFullyConnectedLayer {
	return CFullyConnectedLayer{
		CLayer: CLayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CFullyConnectedLayerClass) Alloc() CFullyConnectedLayer {
	rv := objc.Send[CFullyConnectedLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CFullyConnectedLayerClass) New() CFullyConnectedLayer {
	rv := objc.Send[CFullyConnectedLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CFullyConnectedLayer) Init() CFullyConnectedLayer {
	rv := objc.Send[CFullyConnectedLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CFullyConnectedLayer) Autorelease() CFullyConnectedLayer {
	rv := objc.Send[CFullyConnectedLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCFullyConnectedLayer creates a new CFullyConnectedLayer instance.
func NewCFullyConnectedLayer() CFullyConnectedLayer {
	return getCFullyConnectedLayerClass().New()
}




