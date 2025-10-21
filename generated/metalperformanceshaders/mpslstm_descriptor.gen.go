// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [LSTMDescriptor] class.
var (
	LSTMDescriptorClass     _LSTMDescriptorClass
	LSTMDescriptorClassOnce sync.Once
)

func getLSTMDescriptorClass() _LSTMDescriptorClass {
	LSTMDescriptorClassOnce.Do(func() {
		LSTMDescriptorClass = _LSTMDescriptorClass{objc.GetClass("MPSLSTMDescriptor")}
	})
	return LSTMDescriptorClass
}

type _LSTMDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [LSTMDescriptor] class.
type ILSTMDescriptor interface {
	objectivec.IObject
}

// A description of a long short-term memory block or layer.
//
// The recurrent neural network (RNN) layer initialized with transforms the input data (image or matrix), the memory cell data, and previous output with a set of filters. Each produces one feature map in the output data and memory cell according to the long short-term memory (LSTM) formula detailed below. You may provide the LSTM unit with a single input or a sequence of inputs.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSLSTMDescriptor
type LSTMDescriptor struct {
	objectivec.Object
}

// LSTMDescriptorFrom constructs a [LSTMDescriptor] from an unsafe.Pointer.
//
// A description of a long short-term memory block or layer.
func LSTMDescriptorFrom(ptr unsafe.Pointer) LSTMDescriptor {
	return LSTMDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (lc _LSTMDescriptorClass) Alloc() LSTMDescriptor {
	rv := objc.Send[LSTMDescriptor](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (lc _LSTMDescriptorClass) New() LSTMDescriptor {
	rv := objc.Send[LSTMDescriptor](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LSTMDescriptor) Init() LSTMDescriptor {
	rv := objc.Send[LSTMDescriptor](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LSTMDescriptor) Autorelease() LSTMDescriptor {
	rv := objc.Send[LSTMDescriptor](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLSTMDescriptor creates a new LSTMDescriptor instance.
func NewLSTMDescriptor() LSTMDescriptor {
	return getLSTMDescriptorClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/celltooutputneurontype
func (l_ LSTMDescriptor) CellToOutputNeuronType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("cellToOutputNeuronType"))
	return rv
}


// SetCellToOutputNeuronType sets the value of the cellToOutputNeuronType property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/celltooutputneurontype
func (l_ LSTMDescriptor) SetCellToOutputNeuronType(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCellToOutputNeuronType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/outputgatememoryweights
func (l_ LSTMDescriptor) OutputGateMemoryWeights() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("outputGateMemoryWeights"))
	return rv
}


// SetOutputGateMemoryWeights sets the value of the outputGateMemoryWeights property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/outputgatememoryweights
func (l_ LSTMDescriptor) SetOutputGateMemoryWeights(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setOutputGateMemoryWeights:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/cellgaterecurrentweights
func (l_ LSTMDescriptor) CellGateRecurrentWeights() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("cellGateRecurrentWeights"))
	return rv
}


// SetCellGateRecurrentWeights sets the value of the cellGateRecurrentWeights property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/cellgaterecurrentweights
func (l_ LSTMDescriptor) SetCellGateRecurrentWeights(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCellGateRecurrentWeights:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/celltooutputneuronparamb
func (l_ LSTMDescriptor) CellToOutputNeuronParamB() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("cellToOutputNeuronParamB"))
	return rv
}


// SetCellToOutputNeuronParamB sets the value of the cellToOutputNeuronParamB property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/celltooutputneuronparamb
func (l_ LSTMDescriptor) SetCellToOutputNeuronParamB(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCellToOutputNeuronParamB:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/celltooutputneuronparama
func (l_ LSTMDescriptor) CellToOutputNeuronParamA() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("cellToOutputNeuronParamA"))
	return rv
}


// SetCellToOutputNeuronParamA sets the value of the cellToOutputNeuronParamA property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/celltooutputneuronparama
func (l_ LSTMDescriptor) SetCellToOutputNeuronParamA(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCellToOutputNeuronParamA:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/outputgaterecurrentweights
func (l_ LSTMDescriptor) OutputGateRecurrentWeights() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("outputGateRecurrentWeights"))
	return rv
}


// SetOutputGateRecurrentWeights sets the value of the outputGateRecurrentWeights property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/outputgaterecurrentweights
func (l_ LSTMDescriptor) SetOutputGateRecurrentWeights(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setOutputGateRecurrentWeights:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/cellgateinputweights
func (l_ LSTMDescriptor) CellGateInputWeights() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("cellGateInputWeights"))
	return rv
}


// SetCellGateInputWeights sets the value of the cellGateInputWeights property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/cellgateinputweights
func (l_ LSTMDescriptor) SetCellGateInputWeights(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCellGateInputWeights:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/outputgateinputweights
func (l_ LSTMDescriptor) OutputGateInputWeights() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("outputGateInputWeights"))
	return rv
}


// SetOutputGateInputWeights sets the value of the outputGateInputWeights property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/outputgateinputweights
func (l_ LSTMDescriptor) SetOutputGateInputWeights(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setOutputGateInputWeights:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/outputfeaturechannels
func (l_ LSTMDescriptor) OutputFeatureChannels() int {
	rv := objc.Send[int](l_.ID, objc.Sel("outputFeatureChannels"))
	return rv
}


// SetOutputFeatureChannels sets the value of the outputFeatureChannels property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/outputfeaturechannels
func (l_ LSTMDescriptor) SetOutputFeatureChannels(value int) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setOutputFeatureChannels:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/forgetgatememoryweights
func (l_ LSTMDescriptor) ForgetGateMemoryWeights() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("forgetGateMemoryWeights"))
	return rv
}


// SetForgetGateMemoryWeights sets the value of the forgetGateMemoryWeights property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/forgetgatememoryweights
func (l_ LSTMDescriptor) SetForgetGateMemoryWeights(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setForgetGateMemoryWeights:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/inputgatememoryweights
func (l_ LSTMDescriptor) InputGateMemoryWeights() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("inputGateMemoryWeights"))
	return rv
}


// SetInputGateMemoryWeights sets the value of the inputGateMemoryWeights property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/inputgatememoryweights
func (l_ LSTMDescriptor) SetInputGateMemoryWeights(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setInputGateMemoryWeights:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/celltooutputneuronparamc
func (l_ LSTMDescriptor) CellToOutputNeuronParamC() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("cellToOutputNeuronParamC"))
	return rv
}


// SetCellToOutputNeuronParamC sets the value of the cellToOutputNeuronParamC property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/celltooutputneuronparamc
func (l_ LSTMDescriptor) SetCellToOutputNeuronParamC(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCellToOutputNeuronParamC:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/inputgateinputweights
func (l_ LSTMDescriptor) InputGateInputWeights() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("inputGateInputWeights"))
	return rv
}


// SetInputGateInputWeights sets the value of the inputGateInputWeights property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/inputgateinputweights
func (l_ LSTMDescriptor) SetInputGateInputWeights(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setInputGateInputWeights:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/inputgaterecurrentweights
func (l_ LSTMDescriptor) InputGateRecurrentWeights() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("inputGateRecurrentWeights"))
	return rv
}


// SetInputGateRecurrentWeights sets the value of the inputGateRecurrentWeights property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/inputgaterecurrentweights
func (l_ LSTMDescriptor) SetInputGateRecurrentWeights(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setInputGateRecurrentWeights:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/inputfeaturechannels
func (l_ LSTMDescriptor) InputFeatureChannels() int {
	rv := objc.Send[int](l_.ID, objc.Sel("inputFeatureChannels"))
	return rv
}


// SetInputFeatureChannels sets the value of the inputFeatureChannels property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/inputfeaturechannels
func (l_ LSTMDescriptor) SetInputFeatureChannels(value int) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setInputFeatureChannels:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/memoryweightsarediagonal
func (l_ LSTMDescriptor) MemoryWeightsAreDiagonal() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("memoryWeightsAreDiagonal"))
	return rv
}


// SetMemoryWeightsAreDiagonal sets the value of the memoryWeightsAreDiagonal property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/memoryweightsarediagonal
func (l_ LSTMDescriptor) SetMemoryWeightsAreDiagonal(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setMemoryWeightsAreDiagonal:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/cellgatememoryweights
func (l_ LSTMDescriptor) CellGateMemoryWeights() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("cellGateMemoryWeights"))
	return rv
}


// SetCellGateMemoryWeights sets the value of the cellGateMemoryWeights property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/cellgatememoryweights
func (l_ LSTMDescriptor) SetCellGateMemoryWeights(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCellGateMemoryWeights:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/forgetgaterecurrentweights
func (l_ LSTMDescriptor) ForgetGateRecurrentWeights() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("forgetGateRecurrentWeights"))
	return rv
}


// SetForgetGateRecurrentWeights sets the value of the forgetGateRecurrentWeights property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/forgetgaterecurrentweights
func (l_ LSTMDescriptor) SetForgetGateRecurrentWeights(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setForgetGateRecurrentWeights:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/forgetgateinputweights
func (l_ LSTMDescriptor) ForgetGateInputWeights() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("forgetGateInputWeights"))
	return rv
}


// SetForgetGateInputWeights sets the value of the forgetGateInputWeights property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/forgetgateinputweights
func (l_ LSTMDescriptor) SetForgetGateInputWeights(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setForgetGateInputWeights:"), value)
}



