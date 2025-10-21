// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CEmbeddingLayer] class.
var (
	CEmbeddingLayerClass     _CEmbeddingLayerClass
	CEmbeddingLayerClassOnce sync.Once
)

func getCEmbeddingLayerClass() _CEmbeddingLayerClass {
	CEmbeddingLayerClassOnce.Do(func() {
		CEmbeddingLayerClass = _CEmbeddingLayerClass{objc.GetClass("MLCEmbeddingLayer")}
	})
	return CEmbeddingLayerClass
}

type _CEmbeddingLayerClass struct {
	class objc.Class
}

// An interface definition for the [CEmbeddingLayer] class.
type ICEmbeddingLayer interface {
	ICLayer
}

// A layer that stores a word embedding.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCEmbeddingLayer
type CEmbeddingLayer struct {
	CLayer
}

// CEmbeddingLayerFrom constructs a [CEmbeddingLayer] from an unsafe.Pointer.
//
// A layer that stores a word embedding.
func CEmbeddingLayerFrom(ptr unsafe.Pointer) CEmbeddingLayer {
	return CEmbeddingLayer{
		CLayer: CLayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CEmbeddingLayerClass) Alloc() CEmbeddingLayer {
	rv := objc.Send[CEmbeddingLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CEmbeddingLayerClass) New() CEmbeddingLayer {
	rv := objc.Send[CEmbeddingLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CEmbeddingLayer) Init() CEmbeddingLayer {
	rv := objc.Send[CEmbeddingLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CEmbeddingLayer) Autorelease() CEmbeddingLayer {
	rv := objc.Send[CEmbeddingLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCEmbeddingLayer creates a new CEmbeddingLayer instance.
func NewCEmbeddingLayer() CEmbeddingLayer {
	return getCEmbeddingLayerClass().New()
}




