// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	objectivec.IObject
	Alpha() float64
	SetAlpha(value float64)
	SourceInputFeatureChannels() int
	SetSourceInputFeatureChannels(value int)
	SourceNumberOfFeatureVectors() int
	SetSourceNumberOfFeatureVectors(value int)
}

// A neuron activation kernel that operates on matrices.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixNeuron
type MatrixNeuron struct {
	objectivec.Object
}

// MatrixNeuronFrom constructs a [MatrixNeuron] from an unsafe.Pointer.
//
// A neuron activation kernel that operates on matrices.
func MatrixNeuronFrom(ptr unsafe.Pointer) MatrixNeuron {
	return MatrixNeuron{objectivec.Object{objc.ID(ptr)}}
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


//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/alpha
func (m_ MatrixNeuron) Alpha() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("alpha"))
	return rv
}


// SetAlpha sets the value of the alpha property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/alpha
func (m_ MatrixNeuron) SetAlpha(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlpha:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/sourceinputfeaturechannels
func (m_ MatrixNeuron) SourceInputFeatureChannels() int {
	rv := objc.Send[int](m_.ID, objc.Sel("sourceInputFeatureChannels"))
	return rv
}


// SetSourceInputFeatureChannels sets the value of the sourceInputFeatureChannels property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/sourceinputfeaturechannels
func (m_ MatrixNeuron) SetSourceInputFeatureChannels(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceInputFeatureChannels:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/sourcenumberoffeaturevectors
func (m_ MatrixNeuron) SourceNumberOfFeatureVectors() int {
	rv := objc.Send[int](m_.ID, objc.Sel("sourceNumberOfFeatureVectors"))
	return rv
}


// SetSourceNumberOfFeatureVectors sets the value of the sourceNumberOfFeatureVectors property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/sourcenumberoffeaturevectors
func (m_ MatrixNeuron) SetSourceNumberOfFeatureVectors(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceNumberOfFeatureVectors:"), value)
}



