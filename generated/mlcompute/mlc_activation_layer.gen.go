// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CActivationLayer] class.
var (
	CActivationLayerClass     _CActivationLayerClass
	CActivationLayerClassOnce sync.Once
)

func getCActivationLayerClass() _CActivationLayerClass {
	CActivationLayerClassOnce.Do(func() {
		CActivationLayerClass = _CActivationLayerClass{objc.GetClass("MLCActivationLayer")}
	})
	return CActivationLayerClass
}

type _CActivationLayerClass struct {
	class objc.Class
}

// An interface definition for the [CActivationLayer] class.
type ICActivationLayer interface {
	ICLayer
}

// A layer that applies an activation function to the source tensor and produces an output.
//
// To construct an activation layer, create an activation descriptor and then pass it to the initializer.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationLayer
type CActivationLayer struct {
	CLayer
}

// CActivationLayerFrom constructs a [CActivationLayer] from an unsafe.Pointer.
//
// A layer that applies an activation function to the source tensor and produces an output.
func CActivationLayerFrom(ptr unsafe.Pointer) CActivationLayer {
	return CActivationLayer{
		CLayer: CLayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CActivationLayerClass) Alloc() CActivationLayer {
	rv := objc.Send[CActivationLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CActivationLayerClass) New() CActivationLayer {
	rv := objc.Send[CActivationLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CActivationLayer) Init() CActivationLayer {
	rv := objc.Send[CActivationLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CActivationLayer) Autorelease() CActivationLayer {
	rv := objc.Send[CActivationLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCActivationLayer creates a new CActivationLayer instance.
func NewCActivationLayer() CActivationLayer {
	return getCActivationLayerClass().New()
}




