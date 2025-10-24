// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CGramMatrixLayer] class.
var (
	CGramMatrixLayerClass     _CGramMatrixLayerClass
	CGramMatrixLayerClassOnce sync.Once
)

func getCGramMatrixLayerClass() _CGramMatrixLayerClass {
	CGramMatrixLayerClassOnce.Do(func() {
		CGramMatrixLayerClass = _CGramMatrixLayerClass{objc.GetClass("MLCGramMatrixLayer")}
	})
	return CGramMatrixLayerClass
}

type _CGramMatrixLayerClass struct {
	class objc.Class
}

// An interface definition for the [CGramMatrixLayer] class.
type ICGramMatrixLayer interface {
	ICLayer
	// properties:
	Scale() float32
	SetScale(value float32)
	// methods:
}

// A layer that computes the uncentered cross-correlation values between the spacial planes of each feature channel of a tensor.
//
// For example, if the input tensor batch function is: The computation performed by this layer is: Interpret this operation as computing all combinations of fully connected layers between the different spatial planes of the input tensor. The layer performs this operation independently for each tensor in a batch. Then the layer stores these results in the feature channel and x-coordinate indices of the output batch. Legend:


// A layer that computes the uncentered cross-correlation values between the spacial planes of each feature channel of a tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGramMatrixLayer
type CGramMatrixLayer struct {
	CLayer
}

// CGramMatrixLayerFrom constructs a [CGramMatrixLayer] from an unsafe.Pointer.
//
// A layer that computes the uncentered cross-correlation values between the spacial planes of each feature channel of a tensor.
func CGramMatrixLayerFrom(ptr unsafe.Pointer) CGramMatrixLayer {
	return CGramMatrixLayer{
		CLayer: CLayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CGramMatrixLayerClass) Alloc() CGramMatrixLayer {
	rv := objc.Send[CGramMatrixLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CGramMatrixLayerClass) New() CGramMatrixLayer {
	rv := objc.Send[CGramMatrixLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CGramMatrixLayer) Init() CGramMatrixLayer {
	rv := objc.Send[CGramMatrixLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CGramMatrixLayer) Autorelease() CGramMatrixLayer {
	rv := objc.Send[CGramMatrixLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCGramMatrixLayer creates a new CGramMatrixLayer instance.
func NewCGramMatrixLayer() CGramMatrixLayer {
	return getCGramMatrixLayerClass().New()
}



// The scaling factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcgrammatrixlayer/scale
func (c_ CGramMatrixLayer) Scale() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("scale"))
	return rv
}


// The scaling factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcgrammatrixlayer/scale
func (c_ CGramMatrixLayer) SetScale(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScale:"), value)
}



