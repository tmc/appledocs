// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	IRNNDescriptor
	// properties:
	CellGateInputWeights() CNNConvolutionDataSource /* not a class type */
	SetCellGateInputWeights(value CNNConvolutionDataSource /* not a class type */)
	CellGateMemoryWeights() CNNConvolutionDataSource /* not a class type */
	SetCellGateMemoryWeights(value CNNConvolutionDataSource /* not a class type */)
	CellGateRecurrentWeights() CNNConvolutionDataSource /* not a class type */
	SetCellGateRecurrentWeights(value CNNConvolutionDataSource /* not a class type */)
	CellToOutputNeuronParamA() float32
	SetCellToOutputNeuronParamA(value float32)
	CellToOutputNeuronParamB() float32
	SetCellToOutputNeuronParamB(value float32)
	CellToOutputNeuronParamC() float32
	SetCellToOutputNeuronParamC(value float32)
	CellToOutputNeuronType() CNNNeuronType /* not a class type */
	SetCellToOutputNeuronType(value CNNNeuronType /* not a class type */)
	ForgetGateInputWeights() CNNConvolutionDataSource /* not a class type */
	SetForgetGateInputWeights(value CNNConvolutionDataSource /* not a class type */)
	ForgetGateMemoryWeights() CNNConvolutionDataSource /* not a class type */
	SetForgetGateMemoryWeights(value CNNConvolutionDataSource /* not a class type */)
	ForgetGateRecurrentWeights() CNNConvolutionDataSource /* not a class type */
	SetForgetGateRecurrentWeights(value CNNConvolutionDataSource /* not a class type */)
	InputGateInputWeights() CNNConvolutionDataSource /* not a class type */
	SetInputGateInputWeights(value CNNConvolutionDataSource /* not a class type */)
	InputGateMemoryWeights() CNNConvolutionDataSource /* not a class type */
	SetInputGateMemoryWeights(value CNNConvolutionDataSource /* not a class type */)
	InputGateRecurrentWeights() CNNConvolutionDataSource /* not a class type */
	SetInputGateRecurrentWeights(value CNNConvolutionDataSource /* not a class type */)
	MemoryWeightsAreDiagonal() bool
	SetMemoryWeightsAreDiagonal(value bool)
	OutputGateInputWeights() CNNConvolutionDataSource /* not a class type */
	SetOutputGateInputWeights(value CNNConvolutionDataSource /* not a class type */)
	OutputGateMemoryWeights() CNNConvolutionDataSource /* not a class type */
	SetOutputGateMemoryWeights(value CNNConvolutionDataSource /* not a class type */)
	OutputGateRecurrentWeights() CNNConvolutionDataSource /* not a class type */
	SetOutputGateRecurrentWeights(value CNNConvolutionDataSource /* not a class type */)
	InputFeatureChannels() int
	SetInputFeatureChannels(value int)
	OutputFeatureChannels() int
	SetOutputFeatureChannels(value int)
	// methods:
}

// A description of a long short-term memory block or layer.
//
// The recurrent neural network (RNN) layer initialized with transforms the input data (image or matrix), the memory cell data, and previous output with a set of filters. Each produces one feature map in the output data and memory cell according to the long short-term memory (LSTM) formula detailed below. You may provide the LSTM unit with a single input or a sequence of inputs.


// A description of a long short-term memory block or layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSLSTMDescriptor
type LSTMDescriptor struct {
	RNNDescriptor
}

// LSTMDescriptorFrom constructs a [LSTMDescriptor] from an unsafe.Pointer.
//
// A description of a long short-term memory block or layer.
func LSTMDescriptorFrom(ptr unsafe.Pointer) LSTMDescriptor {
	return LSTMDescriptor{
		RNNDescriptor: RNNDescriptorFrom(ptr),
	}
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/cellgateinputweights
func (l_ LSTMDescriptor) CellGateInputWeights() CNNConvolutionDataSource /* not a class type */ {
	rv := objc.Send[CNNConvolutionDataSource](l_.ID, objc.Sel("cellGateInputWeights"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/cellgateinputweights
func (l_ LSTMDescriptor) SetCellGateInputWeights(value CNNConvolutionDataSource /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCellGateInputWeights:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/cellgatememoryweights
func (l_ LSTMDescriptor) CellGateMemoryWeights() CNNConvolutionDataSource /* not a class type */ {
	rv := objc.Send[CNNConvolutionDataSource](l_.ID, objc.Sel("cellGateMemoryWeights"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/cellgatememoryweights
func (l_ LSTMDescriptor) SetCellGateMemoryWeights(value CNNConvolutionDataSource /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCellGateMemoryWeights:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/cellgaterecurrentweights
func (l_ LSTMDescriptor) CellGateRecurrentWeights() CNNConvolutionDataSource /* not a class type */ {
	rv := objc.Send[CNNConvolutionDataSource](l_.ID, objc.Sel("cellGateRecurrentWeights"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/cellgaterecurrentweights
func (l_ LSTMDescriptor) SetCellGateRecurrentWeights(value CNNConvolutionDataSource /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCellGateRecurrentWeights:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/celltooutputneuronparama
func (l_ LSTMDescriptor) CellToOutputNeuronParamA() float32 {
	rv := objc.Send[float32](l_.ID, objc.Sel("cellToOutputNeuronParamA"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/celltooutputneuronparama
func (l_ LSTMDescriptor) SetCellToOutputNeuronParamA(value float32) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCellToOutputNeuronParamA:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/celltooutputneuronparamb
func (l_ LSTMDescriptor) CellToOutputNeuronParamB() float32 {
	rv := objc.Send[float32](l_.ID, objc.Sel("cellToOutputNeuronParamB"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/celltooutputneuronparamb
func (l_ LSTMDescriptor) SetCellToOutputNeuronParamB(value float32) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCellToOutputNeuronParamB:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/celltooutputneuronparamc
func (l_ LSTMDescriptor) CellToOutputNeuronParamC() float32 {
	rv := objc.Send[float32](l_.ID, objc.Sel("cellToOutputNeuronParamC"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/celltooutputneuronparamc
func (l_ LSTMDescriptor) SetCellToOutputNeuronParamC(value float32) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCellToOutputNeuronParamC:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/celltooutputneurontype
func (l_ LSTMDescriptor) CellToOutputNeuronType() CNNNeuronType /* not a class type */ {
	rv := objc.Send[CNNNeuronType](l_.ID, objc.Sel("cellToOutputNeuronType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/celltooutputneurontype
func (l_ LSTMDescriptor) SetCellToOutputNeuronType(value CNNNeuronType /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCellToOutputNeuronType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/forgetgateinputweights
func (l_ LSTMDescriptor) ForgetGateInputWeights() CNNConvolutionDataSource /* not a class type */ {
	rv := objc.Send[CNNConvolutionDataSource](l_.ID, objc.Sel("forgetGateInputWeights"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/forgetgateinputweights
func (l_ LSTMDescriptor) SetForgetGateInputWeights(value CNNConvolutionDataSource /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setForgetGateInputWeights:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/forgetgatememoryweights
func (l_ LSTMDescriptor) ForgetGateMemoryWeights() CNNConvolutionDataSource /* not a class type */ {
	rv := objc.Send[CNNConvolutionDataSource](l_.ID, objc.Sel("forgetGateMemoryWeights"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/forgetgatememoryweights
func (l_ LSTMDescriptor) SetForgetGateMemoryWeights(value CNNConvolutionDataSource /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setForgetGateMemoryWeights:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/forgetgaterecurrentweights
func (l_ LSTMDescriptor) ForgetGateRecurrentWeights() CNNConvolutionDataSource /* not a class type */ {
	rv := objc.Send[CNNConvolutionDataSource](l_.ID, objc.Sel("forgetGateRecurrentWeights"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/forgetgaterecurrentweights
func (l_ LSTMDescriptor) SetForgetGateRecurrentWeights(value CNNConvolutionDataSource /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setForgetGateRecurrentWeights:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/inputgateinputweights
func (l_ LSTMDescriptor) InputGateInputWeights() CNNConvolutionDataSource /* not a class type */ {
	rv := objc.Send[CNNConvolutionDataSource](l_.ID, objc.Sel("inputGateInputWeights"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/inputgateinputweights
func (l_ LSTMDescriptor) SetInputGateInputWeights(value CNNConvolutionDataSource /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setInputGateInputWeights:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/inputgatememoryweights
func (l_ LSTMDescriptor) InputGateMemoryWeights() CNNConvolutionDataSource /* not a class type */ {
	rv := objc.Send[CNNConvolutionDataSource](l_.ID, objc.Sel("inputGateMemoryWeights"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/inputgatememoryweights
func (l_ LSTMDescriptor) SetInputGateMemoryWeights(value CNNConvolutionDataSource /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setInputGateMemoryWeights:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/inputgaterecurrentweights
func (l_ LSTMDescriptor) InputGateRecurrentWeights() CNNConvolutionDataSource /* not a class type */ {
	rv := objc.Send[CNNConvolutionDataSource](l_.ID, objc.Sel("inputGateRecurrentWeights"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/inputgaterecurrentweights
func (l_ LSTMDescriptor) SetInputGateRecurrentWeights(value CNNConvolutionDataSource /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setInputGateRecurrentWeights:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/memoryweightsarediagonal
func (l_ LSTMDescriptor) MemoryWeightsAreDiagonal() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("memoryWeightsAreDiagonal"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/memoryweightsarediagonal
func (l_ LSTMDescriptor) SetMemoryWeightsAreDiagonal(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setMemoryWeightsAreDiagonal:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/outputgateinputweights
func (l_ LSTMDescriptor) OutputGateInputWeights() CNNConvolutionDataSource /* not a class type */ {
	rv := objc.Send[CNNConvolutionDataSource](l_.ID, objc.Sel("outputGateInputWeights"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/outputgateinputweights
func (l_ LSTMDescriptor) SetOutputGateInputWeights(value CNNConvolutionDataSource /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setOutputGateInputWeights:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/outputgatememoryweights
func (l_ LSTMDescriptor) OutputGateMemoryWeights() CNNConvolutionDataSource /* not a class type */ {
	rv := objc.Send[CNNConvolutionDataSource](l_.ID, objc.Sel("outputGateMemoryWeights"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/outputgatememoryweights
func (l_ LSTMDescriptor) SetOutputGateMemoryWeights(value CNNConvolutionDataSource /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setOutputGateMemoryWeights:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/outputgaterecurrentweights
func (l_ LSTMDescriptor) OutputGateRecurrentWeights() CNNConvolutionDataSource /* not a class type */ {
	rv := objc.Send[CNNConvolutionDataSource](l_.ID, objc.Sel("outputGateRecurrentWeights"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/outputgaterecurrentweights
func (l_ LSTMDescriptor) SetOutputGateRecurrentWeights(value CNNConvolutionDataSource /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setOutputGateRecurrentWeights:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/inputfeaturechannels
func (l_ LSTMDescriptor) InputFeatureChannels() int {
	rv := objc.Send[int](l_.ID, objc.Sel("inputFeatureChannels"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/inputfeaturechannels
func (l_ LSTMDescriptor) SetInputFeatureChannels(value int) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setInputFeatureChannels:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/outputfeaturechannels
func (l_ LSTMDescriptor) OutputFeatureChannels() int {
	rv := objc.Send[int](l_.ID, objc.Sel("outputFeatureChannels"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/outputfeaturechannels
func (l_ LSTMDescriptor) SetOutputFeatureChannels(value int) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setOutputFeatureChannels:"), value)
}



