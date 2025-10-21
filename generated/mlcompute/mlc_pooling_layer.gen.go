// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CPoolingLayer] class.
var (
	CPoolingLayerClass     _CPoolingLayerClass
	CPoolingLayerClassOnce sync.Once
)

func getCPoolingLayerClass() _CPoolingLayerClass {
	CPoolingLayerClassOnce.Do(func() {
		CPoolingLayerClass = _CPoolingLayerClass{objc.GetClass("MLCPoolingLayer")}
	})
	return CPoolingLayerClass
}

type _CPoolingLayerClass struct {
	class objc.Class
}

// An interface definition for the [CPoolingLayer] class.
type ICPoolingLayer interface {
	ICLayer
}

// A layer that summarizes the average presence of a feature.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPoolingLayer
type CPoolingLayer struct {
	CLayer
}

// CPoolingLayerFrom constructs a [CPoolingLayer] from an unsafe.Pointer.
//
// A layer that summarizes the average presence of a feature.
func CPoolingLayerFrom(ptr unsafe.Pointer) CPoolingLayer {
	return CPoolingLayer{
		CLayer: CLayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CPoolingLayerClass) Alloc() CPoolingLayer {
	rv := objc.Send[CPoolingLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CPoolingLayerClass) New() CPoolingLayer {
	rv := objc.Send[CPoolingLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CPoolingLayer) Init() CPoolingLayer {
	rv := objc.Send[CPoolingLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CPoolingLayer) Autorelease() CPoolingLayer {
	rv := objc.Send[CPoolingLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCPoolingLayer creates a new CPoolingLayer instance.
func NewCPoolingLayer() CPoolingLayer {
	return getCPoolingLayerClass().New()
}




