// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CConcatenationLayer] class.
var (
	CConcatenationLayerClass     _CConcatenationLayerClass
	CConcatenationLayerClassOnce sync.Once
)

func getCConcatenationLayerClass() _CConcatenationLayerClass {
	CConcatenationLayerClassOnce.Do(func() {
		CConcatenationLayerClass = _CConcatenationLayerClass{objc.GetClass("MLCConcatenationLayer")}
	})
	return CConcatenationLayerClass
}

type _CConcatenationLayerClass struct {
	class objc.Class
}

// An interface definition for the [CConcatenationLayer] class.
type ICConcatenationLayer interface {
	ICLayer
	// properties:
	Dimension() int
	SetDimension(value int)
	// methods:
}

// A layer that combines tensors into a single tensor.


// A layer that combines tensors into a single tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConcatenationLayer
type CConcatenationLayer struct {
	CLayer
}

// CConcatenationLayerFrom constructs a [CConcatenationLayer] from an unsafe.Pointer.
//
// A layer that combines tensors into a single tensor.
func CConcatenationLayerFrom(ptr unsafe.Pointer) CConcatenationLayer {
	return CConcatenationLayer{
		CLayer: CLayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CConcatenationLayerClass) Alloc() CConcatenationLayer {
	rv := objc.Send[CConcatenationLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CConcatenationLayerClass) New() CConcatenationLayer {
	rv := objc.Send[CConcatenationLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CConcatenationLayer) Init() CConcatenationLayer {
	rv := objc.Send[CConcatenationLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CConcatenationLayer) Autorelease() CConcatenationLayer {
	rv := objc.Send[CConcatenationLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCConcatenationLayer creates a new CConcatenationLayer instance.
func NewCConcatenationLayer() CConcatenationLayer {
	return getCConcatenationLayerClass().New()
}



// The dimension, or axis, along which you concatenate tensors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcconcatenationlayer/dimension
func (c_ CConcatenationLayer) Dimension() int {
	rv := objc.Send[int](c_.ID, objc.Sel("dimension"))
	return rv
}


// The dimension, or axis, along which you concatenate tensors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcconcatenationlayer/dimension
func (c_ CConcatenationLayer) SetDimension(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDimension:"), value)
}



