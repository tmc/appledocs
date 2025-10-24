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
	// properties:
	AccumulateWeightGradients() bool
	SetAccumulateWeightGradients(value bool)
	InputFeatureChannels() int
	SetInputFeatureChannels(value int)
	OutputFeatureChannels() int
	SetOutputFeatureChannels(value int)
	RecurrentOutputIsTemporary() bool
	SetRecurrentOutputIsTemporary(value bool)
	StoreAllIntermediateStates() bool
	SetStoreAllIntermediateStates(value bool)
	TrainingStateIsTemporary() bool
	SetTrainingStateIsTemporary(value bool)
	// methods:
}

// A layer for training recurrent neural networks on Metal Performance Shaders matrices.


// A layer for training recurrent neural networks on Metal Performance Shaders matrices.
//
// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/accumulateweightgradients
func (r_ RNNMatrixTrainingLayer) AccumulateWeightGradients() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("accumulateWeightGradients"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/accumulateweightgradients
func (r_ RNNMatrixTrainingLayer) SetAccumulateWeightGradients(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setAccumulateWeightGradients:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/inputfeaturechannels
func (r_ RNNMatrixTrainingLayer) InputFeatureChannels() int {
	rv := objc.Send[int](r_.ID, objc.Sel("inputFeatureChannels"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/inputfeaturechannels
func (r_ RNNMatrixTrainingLayer) SetInputFeatureChannels(value int) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setInputFeatureChannels:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/outputfeaturechannels
func (r_ RNNMatrixTrainingLayer) OutputFeatureChannels() int {
	rv := objc.Send[int](r_.ID, objc.Sel("outputFeatureChannels"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/outputfeaturechannels
func (r_ RNNMatrixTrainingLayer) SetOutputFeatureChannels(value int) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setOutputFeatureChannels:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/recurrentoutputistemporary
func (r_ RNNMatrixTrainingLayer) RecurrentOutputIsTemporary() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("recurrentOutputIsTemporary"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/recurrentoutputistemporary
func (r_ RNNMatrixTrainingLayer) SetRecurrentOutputIsTemporary(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRecurrentOutputIsTemporary:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/storeallintermediatestates
func (r_ RNNMatrixTrainingLayer) StoreAllIntermediateStates() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("storeAllIntermediateStates"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/storeallintermediatestates
func (r_ RNNMatrixTrainingLayer) SetStoreAllIntermediateStates(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setStoreAllIntermediateStates:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/trainingstateistemporary
func (r_ RNNMatrixTrainingLayer) TrainingStateIsTemporary() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("trainingStateIsTemporary"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/trainingstateistemporary
func (r_ RNNMatrixTrainingLayer) SetTrainingStateIsTemporary(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setTrainingStateIsTemporary:"), value)
}



