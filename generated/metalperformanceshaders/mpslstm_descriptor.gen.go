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
	IRNNDescriptor
	

	// properties:
	CellGateRecurrentWeights() CNNConvolutionDataSource get set /* not a class type */
	SetCellGateRecurrentWeights(value CNNConvolutionDataSource get set /* not a class type */)
	CellGateMemoryWeights() CNNConvolutionDataSource get set /* not a class type */
	SetCellGateMemoryWeights(value CNNConvolutionDataSource get set /* not a class type */)
	InputGateInputWeights() CNNConvolutionDataSource get set /* not a class type */
	SetInputGateInputWeights(value CNNConvolutionDataSource get set /* not a class type */)
	OutputGateMemoryWeights() CNNConvolutionDataSource get set /* not a class type */
	SetOutputGateMemoryWeights(value CNNConvolutionDataSource get set /* not a class type */)
	ForgetGateMemoryWeights() CNNConvolutionDataSource get set /* not a class type */
	SetForgetGateMemoryWeights(value CNNConvolutionDataSource get set /* not a class type */)
	CellToOutputNeuronParamB() objectivec.IObject
	SetCellToOutputNeuronParamB(value objectivec.IObject)
	OutputGateInputWeights() CNNConvolutionDataSource get set /* not a class type */
	SetOutputGateInputWeights(value CNNConvolutionDataSource get set /* not a class type */)
	MemoryWeightsAreDiagonal() objectivec.IObject
	SetMemoryWeightsAreDiagonal(value objectivec.IObject)
	InputGateMemoryWeights() CNNConvolutionDataSource get set /* not a class type */
	SetInputGateMemoryWeights(value CNNConvolutionDataSource get set /* not a class type */)
	ForgetGateInputWeights() CNNConvolutionDataSource get set /* not a class type */
	SetForgetGateInputWeights(value CNNConvolutionDataSource get set /* not a class type */)
	ForgetGateRecurrentWeights() CNNConvolutionDataSource get set /* not a class type */
	SetForgetGateRecurrentWeights(value CNNConvolutionDataSource get set /* not a class type */)
	CellToOutputNeuronType() CNNNeuronType get set /* not a class type */
	SetCellToOutputNeuronType(value CNNNeuronType get set /* not a class type */)
	CellGateInputWeights() CNNConvolutionDataSource get set /* not a class type */
	SetCellGateInputWeights(value CNNConvolutionDataSource get set /* not a class type */)
	CellToOutputNeuronParamA() objectivec.IObject
	SetCellToOutputNeuronParamA(value objectivec.IObject)
	InputGateRecurrentWeights() CNNConvolutionDataSource get set /* not a class type */
	SetInputGateRecurrentWeights(value CNNConvolutionDataSource get set /* not a class type */)
	OutputGateRecurrentWeights() CNNConvolutionDataSource get set /* not a class type */
	SetOutputGateRecurrentWeights(value CNNConvolutionDataSource get set /* not a class type */)
	CellToOutputNeuronParamC() objectivec.IObject
	SetCellToOutputNeuronParamC(value objectivec.IObject)
	InputFeatureChannels() int
	SetInputFeatureChannels(value int)
	OutputFeatureChannels() int
	SetOutputFeatureChannels(value int)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (lc _LSTMDescriptorClass) Alloc() LSTMDescriptor {
	rv := objc.Send[LSTMDescriptor](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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










// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865681-createlstmdescriptor
func (lc _LSTMDescriptorClass) CreateLSTMDescriptor() {
	objc.Send[objc.ID](objc.ID(lc.class), objc.Sel("createLSTMDescriptor"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865681-createlstmdescriptorwithinputfea
func (lc _LSTMDescriptorClass) CreateLSTMDescriptorWithInputFeatureChannelsOutputFeatureChannels(inputFeatureChannels uint, outputFeatureChannels uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(lc.class), objc.Sel("createLSTMDescriptorWithInputFeatureChannels:outputFeatureChannels:"), inputFeatureChannels, outputFeatureChannels)
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865679-cellgaterecurrentweights
func (l_ LSTMDescriptor) CellGateRecurrentWeights() CNNConvolutionDataSource get set /* not a class type */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("cellGateRecurrentWeights"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865679-cellgaterecurrentweights
func (l_ LSTMDescriptor) SetCellGateRecurrentWeights(value CNNConvolutionDataSource get set /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCellGateRecurrentWeights:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865683-cellgatememoryweights
func (l_ LSTMDescriptor) CellGateMemoryWeights() CNNConvolutionDataSource get set /* not a class type */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("cellGateMemoryWeights"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865683-cellgatememoryweights
func (l_ LSTMDescriptor) SetCellGateMemoryWeights(value CNNConvolutionDataSource get set /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCellGateMemoryWeights:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865684-inputgateinputweights
func (l_ LSTMDescriptor) InputGateInputWeights() CNNConvolutionDataSource get set /* not a class type */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("inputGateInputWeights"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865684-inputgateinputweights
func (l_ LSTMDescriptor) SetInputGateInputWeights(value CNNConvolutionDataSource get set /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setInputGateInputWeights:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865688-outputgatememoryweights
func (l_ LSTMDescriptor) OutputGateMemoryWeights() CNNConvolutionDataSource get set /* not a class type */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("outputGateMemoryWeights"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865688-outputgatememoryweights
func (l_ LSTMDescriptor) SetOutputGateMemoryWeights(value CNNConvolutionDataSource get set /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setOutputGateMemoryWeights:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865689-forgetgatememoryweights
func (l_ LSTMDescriptor) ForgetGateMemoryWeights() CNNConvolutionDataSource get set /* not a class type */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("forgetGateMemoryWeights"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865689-forgetgatememoryweights
func (l_ LSTMDescriptor) SetForgetGateMemoryWeights(value CNNConvolutionDataSource get set /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setForgetGateMemoryWeights:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865694-celltooutputneuronparamb
func (l_ LSTMDescriptor) CellToOutputNeuronParamB() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](l_.ID, objc.Sel("cellToOutputNeuronParamB"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865694-celltooutputneuronparamb
func (l_ LSTMDescriptor) SetCellToOutputNeuronParamB(value objectivec.IObject) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCellToOutputNeuronParamB:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865701-outputgateinputweights
func (l_ LSTMDescriptor) OutputGateInputWeights() CNNConvolutionDataSource get set /* not a class type */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("outputGateInputWeights"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865701-outputgateinputweights
func (l_ LSTMDescriptor) SetOutputGateInputWeights(value CNNConvolutionDataSource get set /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setOutputGateInputWeights:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865712-memoryweightsarediagonal
func (l_ LSTMDescriptor) MemoryWeightsAreDiagonal() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](l_.ID, objc.Sel("memoryWeightsAreDiagonal"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865712-memoryweightsarediagonal
func (l_ LSTMDescriptor) SetMemoryWeightsAreDiagonal(value objectivec.IObject) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setMemoryWeightsAreDiagonal:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865731-inputgatememoryweights
func (l_ LSTMDescriptor) InputGateMemoryWeights() CNNConvolutionDataSource get set /* not a class type */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("inputGateMemoryWeights"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865731-inputgatememoryweights
func (l_ LSTMDescriptor) SetInputGateMemoryWeights(value CNNConvolutionDataSource get set /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setInputGateMemoryWeights:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865734-forgetgateinputweights
func (l_ LSTMDescriptor) ForgetGateInputWeights() CNNConvolutionDataSource get set /* not a class type */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("forgetGateInputWeights"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865734-forgetgateinputweights
func (l_ LSTMDescriptor) SetForgetGateInputWeights(value CNNConvolutionDataSource get set /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setForgetGateInputWeights:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865735-forgetgaterecurrentweights
func (l_ LSTMDescriptor) ForgetGateRecurrentWeights() CNNConvolutionDataSource get set /* not a class type */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("forgetGateRecurrentWeights"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865735-forgetgaterecurrentweights
func (l_ LSTMDescriptor) SetForgetGateRecurrentWeights(value CNNConvolutionDataSource get set /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setForgetGateRecurrentWeights:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865736-celltooutputneurontype
func (l_ LSTMDescriptor) CellToOutputNeuronType() CNNNeuronType get set /* not a class type */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("cellToOutputNeuronType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865736-celltooutputneurontype
func (l_ LSTMDescriptor) SetCellToOutputNeuronType(value CNNNeuronType get set /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCellToOutputNeuronType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865741-cellgateinputweights
func (l_ LSTMDescriptor) CellGateInputWeights() CNNConvolutionDataSource get set /* not a class type */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("cellGateInputWeights"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865741-cellgateinputweights
func (l_ LSTMDescriptor) SetCellGateInputWeights(value CNNConvolutionDataSource get set /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCellGateInputWeights:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865744-celltooutputneuronparama
func (l_ LSTMDescriptor) CellToOutputNeuronParamA() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](l_.ID, objc.Sel("cellToOutputNeuronParamA"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865744-celltooutputneuronparama
func (l_ LSTMDescriptor) SetCellToOutputNeuronParamA(value objectivec.IObject) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCellToOutputNeuronParamA:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865747-inputgaterecurrentweights
func (l_ LSTMDescriptor) InputGateRecurrentWeights() CNNConvolutionDataSource get set /* not a class type */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("inputGateRecurrentWeights"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865747-inputgaterecurrentweights
func (l_ LSTMDescriptor) SetInputGateRecurrentWeights(value CNNConvolutionDataSource get set /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setInputGateRecurrentWeights:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865750-outputgaterecurrentweights
func (l_ LSTMDescriptor) OutputGateRecurrentWeights() CNNConvolutionDataSource get set /* not a class type */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("outputGateRecurrentWeights"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865750-outputgaterecurrentweights
func (l_ LSTMDescriptor) SetOutputGateRecurrentWeights(value CNNConvolutionDataSource get set /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setOutputGateRecurrentWeights:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2935551-celltooutputneuronparamc
func (l_ LSTMDescriptor) CellToOutputNeuronParamC() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](l_.ID, objc.Sel("cellToOutputNeuronParamC"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2935551-celltooutputneuronparamc
func (l_ LSTMDescriptor) SetCellToOutputNeuronParamC(value objectivec.IObject) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCellToOutputNeuronParamC:"), value)
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








