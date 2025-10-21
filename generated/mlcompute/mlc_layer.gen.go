// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CLayer] class.
var (
	CLayerClass     _CLayerClass
	CLayerClassOnce sync.Once
)

func getCLayerClass() _CLayerClass {
	CLayerClassOnce.Do(func() {
		CLayerClass = _CLayerClass{objc.GetClass("MLCLayer")}
	})
	return CLayerClass
}

type _CLayerClass struct {
	class objc.Class
}

// An interface definition for the [CLayer] class.
type ICLayer interface {
	objectivec.IObject
}

// The base class for all framework layers.
//
// This class defines a polymorphic interface for subclasses. There are subclasses for each supported neural network layer type. Use the appropriate subclass initializer to create a layer object.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLayer
type CLayer struct {
	objectivec.Object
}

// CLayerFrom constructs a [CLayer] from an unsafe.Pointer.
//
// The base class for all framework layers.
func CLayerFrom(ptr unsafe.Pointer) CLayer {
	return CLayer{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CLayerClass) Alloc() CLayer {
	rv := objc.Send[CLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CLayerClass) New() CLayer {
	rv := objc.Send[CLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CLayer) Init() CLayer {
	rv := objc.Send[CLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CLayer) Autorelease() CLayer {
	rv := objc.Send[CLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCLayer creates a new CLayer instance.
func NewCLayer() CLayer {
	return getCLayerClass().New()
}




