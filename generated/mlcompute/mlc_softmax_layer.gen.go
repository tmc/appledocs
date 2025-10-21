// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CSoftmaxLayer] class.
var (
	CSoftmaxLayerClass     _CSoftmaxLayerClass
	CSoftmaxLayerClassOnce sync.Once
)

func getCSoftmaxLayerClass() _CSoftmaxLayerClass {
	CSoftmaxLayerClassOnce.Do(func() {
		CSoftmaxLayerClass = _CSoftmaxLayerClass{objc.GetClass("MLCSoftmaxLayer")}
	})
	return CSoftmaxLayerClass
}

type _CSoftmaxLayerClass struct {
	class objc.Class
}

// An interface definition for the [CSoftmaxLayer] class.
type ICSoftmaxLayer interface {
	ICLayer
}

// A layer that outputs a probability distribution as attention weights.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCSoftmaxLayer
type CSoftmaxLayer struct {
	CLayer
}

// CSoftmaxLayerFrom constructs a [CSoftmaxLayer] from an unsafe.Pointer.
//
// A layer that outputs a probability distribution as attention weights.
func CSoftmaxLayerFrom(ptr unsafe.Pointer) CSoftmaxLayer {
	return CSoftmaxLayer{
		CLayer: CLayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CSoftmaxLayerClass) Alloc() CSoftmaxLayer {
	rv := objc.Send[CSoftmaxLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CSoftmaxLayerClass) New() CSoftmaxLayer {
	rv := objc.Send[CSoftmaxLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CSoftmaxLayer) Init() CSoftmaxLayer {
	rv := objc.Send[CSoftmaxLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CSoftmaxLayer) Autorelease() CSoftmaxLayer {
	rv := objc.Send[CSoftmaxLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCSoftmaxLayer creates a new CSoftmaxLayer instance.
func NewCSoftmaxLayer() CSoftmaxLayer {
	return getCSoftmaxLayerClass().New()
}




