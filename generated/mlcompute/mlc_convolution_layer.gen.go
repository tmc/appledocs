// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CConvolutionLayer] class.
var (
	CConvolutionLayerClass     _CConvolutionLayerClass
	CConvolutionLayerClassOnce sync.Once
)

func getCConvolutionLayerClass() _CConvolutionLayerClass {
	CConvolutionLayerClassOnce.Do(func() {
		CConvolutionLayerClass = _CConvolutionLayerClass{objc.GetClass("MLCConvolutionLayer")}
	})
	return CConvolutionLayerClass
}

type _CConvolutionLayerClass struct {
	class objc.Class
}

// An interface definition for the [CConvolutionLayer] class.
type ICConvolutionLayer interface {
	ICLayer
}

// A layer that applies a convolution over a signal.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCConvolutionLayer
type CConvolutionLayer struct {
	CLayer
}

// CConvolutionLayerFrom constructs a [CConvolutionLayer] from an unsafe.Pointer.
//
// A layer that applies a convolution over a signal.
func CConvolutionLayerFrom(ptr unsafe.Pointer) CConvolutionLayer {
	return CConvolutionLayer{
		CLayer: CLayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CConvolutionLayerClass) Alloc() CConvolutionLayer {
	rv := objc.Send[CConvolutionLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CConvolutionLayerClass) New() CConvolutionLayer {
	rv := objc.Send[CConvolutionLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CConvolutionLayer) Init() CConvolutionLayer {
	rv := objc.Send[CConvolutionLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CConvolutionLayer) Autorelease() CConvolutionLayer {
	rv := objc.Send[CConvolutionLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCConvolutionLayer creates a new CConvolutionLayer instance.
func NewCConvolutionLayer() CConvolutionLayer {
	return getCConvolutionLayerClass().New()
}




