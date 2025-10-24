// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MatrixNeuron] class.
var (
	MatrixNeuronClass     _MatrixNeuronClass
	MatrixNeuronClassOnce sync.Once
)

func getMatrixNeuronClass() _MatrixNeuronClass {
	MatrixNeuronClassOnce.Do(func() {
		MatrixNeuronClass = _MatrixNeuronClass{objc.GetClass("MPSMatrixNeuron")}
	})
	return MatrixNeuronClass
}

type _MatrixNeuronClass struct {
	class objc.Class
}

// An interface definition for the [MatrixNeuron] class.
type IMatrixNeuron interface {
	IMatrixUnaryKernel
	// properties:
	Alpha() float64
	SetAlpha(value float64)
	SourceInputFeatureChannels() int
	SetSourceInputFeatureChannels(value int)
	SourceNumberOfFeatureVectors() int
	SetSourceNumberOfFeatureVectors(value int)
	// methods:
}

// A neuron activation kernel that operates on matrices.


// A neuron activation kernel that operates on matrices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixNeuron
type MatrixNeuron struct {
	MatrixUnaryKernel
}

// MatrixNeuronFrom constructs a [MatrixNeuron] from an unsafe.Pointer.
//
// A neuron activation kernel that operates on matrices.
func MatrixNeuronFrom(ptr unsafe.Pointer) MatrixNeuron {
	return MatrixNeuron{
		MatrixUnaryKernel: MatrixUnaryKernelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MatrixNeuronClass) Alloc() MatrixNeuron {
	rv := objc.Send[MatrixNeuron](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MatrixNeuronClass) New() MatrixNeuron {
	rv := objc.Send[MatrixNeuron](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MatrixNeuron) Init() MatrixNeuron {
	rv := objc.Send[MatrixNeuron](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MatrixNeuron) Autorelease() MatrixNeuron {
	rv := objc.Send[MatrixNeuron](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMatrixNeuron creates a new MatrixNeuron instance.
func NewMatrixNeuron() MatrixNeuron {
	return getMatrixNeuronClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/alpha
func (m_ MatrixNeuron) Alpha() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("alpha"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/alpha
func (m_ MatrixNeuron) SetAlpha(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlpha:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/sourceinputfeaturechannels
func (m_ MatrixNeuron) SourceInputFeatureChannels() int {
	rv := objc.Send[int](m_.ID, objc.Sel("sourceInputFeatureChannels"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/sourceinputfeaturechannels
func (m_ MatrixNeuron) SetSourceInputFeatureChannels(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceInputFeatureChannels:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/sourcenumberoffeaturevectors
func (m_ MatrixNeuron) SourceNumberOfFeatureVectors() int {
	rv := objc.Send[int](m_.ID, objc.Sel("sourceNumberOfFeatureVectors"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/sourcenumberoffeaturevectors
func (m_ MatrixNeuron) SetSourceNumberOfFeatureVectors(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceNumberOfFeatureVectors:"), value)
}



