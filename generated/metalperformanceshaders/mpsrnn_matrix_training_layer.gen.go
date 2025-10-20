// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [RNNMatrixTrainingLayer] class.
var (
	RNNMatrixTrainingLayerClass     _RNNMatrixTrainingLayerClass
	RNNMatrixTrainingLayerClassOnce sync.Once
)

func getRNNMatrixTrainingLayerClass() _RNNMatrixTrainingLayerClass {
	RNNMatrixTrainingLayerClassOnce.Do(func() {
		RNNMatrixTrainingLayerClass = _RNNMatrixTrainingLayerClass{objc.GetClass("MPSRNNMatrixTrainingLayer")}
	})
	return RNNMatrixTrainingLayerClass
}

type _RNNMatrixTrainingLayerClass struct {
	class objc.Class
}

// An interface definition for the [RNNMatrixTrainingLayer] class.
type IRNNMatrixTrainingLayer interface {
	IKernel
}

// A layer for training recurrent neural networks on Metal Performance Shaders matrices.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixTrainingLayer
type RNNMatrixTrainingLayer struct {
	Kernel
}

// RNNMatrixTrainingLayerFrom constructs a [RNNMatrixTrainingLayer] from an unsafe.Pointer.
//
// A layer for training recurrent neural networks on Metal Performance Shaders matrices.
func RNNMatrixTrainingLayerFrom(ptr unsafe.Pointer) RNNMatrixTrainingLayer {
	return RNNMatrixTrainingLayer{
		Kernel: KernelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (rc _RNNMatrixTrainingLayerClass) Alloc() RNNMatrixTrainingLayer {
	rv := objc.Send[RNNMatrixTrainingLayer](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RNNMatrixTrainingLayerClass) New() RNNMatrixTrainingLayer {
	rv := objc.Send[RNNMatrixTrainingLayer](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RNNMatrixTrainingLayer) Init() RNNMatrixTrainingLayer {
	rv := objc.Send[RNNMatrixTrainingLayer](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RNNMatrixTrainingLayer) Autorelease() RNNMatrixTrainingLayer {
	rv := objc.Send[RNNMatrixTrainingLayer](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRNNMatrixTrainingLayer creates a new RNNMatrixTrainingLayer instance.
func NewRNNMatrixTrainingLayer() RNNMatrixTrainingLayer {
	return getRNNMatrixTrainingLayerClass().New()
}
