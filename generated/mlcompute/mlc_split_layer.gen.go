// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CSplitLayer] class.
var (
	CSplitLayerClass     _CSplitLayerClass
	CSplitLayerClassOnce sync.Once
)

func getCSplitLayerClass() _CSplitLayerClass {
	CSplitLayerClassOnce.Do(func() {
		CSplitLayerClass = _CSplitLayerClass{objc.GetClass("MLCSplitLayer")}
	})
	return CSplitLayerClass
}

type _CSplitLayerClass struct {
	class objc.Class
}

// An interface definition for the [CSplitLayer] class.
type ICSplitLayer interface {
	ICLayer
}

// A layer that splits a tensor value into a list of subtensors.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCSplitLayer
type CSplitLayer struct {
	CLayer
}

// CSplitLayerFrom constructs a [CSplitLayer] from an unsafe.Pointer.
//
// A layer that splits a tensor value into a list of subtensors.
func CSplitLayerFrom(ptr unsafe.Pointer) CSplitLayer {
	return CSplitLayer{
		CLayer: CLayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CSplitLayerClass) Alloc() CSplitLayer {
	rv := objc.Send[CSplitLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CSplitLayerClass) New() CSplitLayer {
	rv := objc.Send[CSplitLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CSplitLayer) Init() CSplitLayer {
	rv := objc.Send[CSplitLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CSplitLayer) Autorelease() CSplitLayer {
	rv := objc.Send[CSplitLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCSplitLayer creates a new CSplitLayer instance.
func NewCSplitLayer() CSplitLayer {
	return getCSplitLayerClass().New()
}




